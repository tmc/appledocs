package appledocs

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// FS provides filesystem access to Apple documentation.
type FS struct {
	fsys fs.FS
}

// Open creates a new FS rooted at the given directory.
// The directory should contain Apple documentation JSON files.
// Returns an error if the path does not exist or is not a directory.
func Open(root string) (*FS, error) {
	info, err := os.Stat(root)
	if err != nil {
		return nil, fmt.Errorf("appledocs.Open: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("appledocs.Open: %s is not a directory", root)
	}

	return &FS{
		fsys: os.DirFS(root),
	}, nil
}

// Open implements fs.FS.
func (f *FS) Open(name string) (fs.File, error) {
	return f.fsys.Open(name)
}

// ReadFile reads the named file from the documentation filesystem.
func (f *FS) ReadFile(name string) ([]byte, error) {
	// Try the exact path first
	file, err := f.fsys.Open(name)
	if err == nil {
		defer file.Close()
		return io.ReadAll(file)
	}

	// If the path ends with .json and doesn't exist, try language-specific variants
	// This handles files that were cached with their language parameter.
	// Try .json.language%3Dobjc and .json.language%3Dswift as fallbacks.
	if strings.HasSuffix(name, ".json") {
		// Try Objective-C variant first (more relevant for ObjC bindings)
		langPath := name + ".language%3Dobjc"
		file, langErr := f.fsys.Open(langPath)
		if langErr == nil {
			defer file.Close()
			return io.ReadAll(file)
		}

		// Try Swift variant as second fallback
		langPath = name + ".language%3Dswift"
		file, langErr = f.fsys.Open(langPath)
		if langErr == nil {
			defer file.Close()
			return io.ReadAll(file)
		}
	}

	// Return the original error if none worked
	return nil, err
}

// Stat returns file info for the named file.
func (f *FS) Stat(name string) (fs.FileInfo, error) {
	return fs.Stat(f.fsys, name)
}

// ReadDocument reads and parses a documentation JSON file.
// Returns an error if the file cannot be read or contains invalid JSON.
func (f *FS) ReadDocument(path string) (*Document, error) {
	data, err := f.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("appledocs: read document %s: %w", path, err)
	}

	var doc Document
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("appledocs: parse JSON in %s: %w", path, err)
	}

	return &doc, nil
}

// IsFramework returns true if the given name appears to be a framework.
// Framework files are JSON files at the root level.
func IsFramework(name string) bool {
	// Frameworks are .json files at the root (no directory separator)
	return strings.HasSuffix(name, ".json") && !strings.Contains(name, string(filepath.Separator))
}

// FrameworkName extracts the framework name from a path.
// For "Foundation.json" returns "Foundation".
// For "Foundation/NSString.json" returns "Foundation".
func FrameworkName(path string) string {
	// Remove .json suffix
	name := strings.TrimSuffix(path, ".json")
	// Take first path component
	if idx := strings.Index(name, string(filepath.Separator)); idx != -1 {
		return name[:idx]
	}
	return name
}

// GetFramework reads a framework's root documentation.
func GetFramework(fsys *FS, name string) (*Document, error) {
	path := name + ".json"
	return fsys.ReadDocument(path)
}

// ListFrameworks returns a sorted list of all available frameworks.
// A framework is identified by a .json file at the root level.
func ListFrameworks(fsys *FS) ([]string, error) {
	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return nil, fmt.Errorf("appledocs.ListFrameworks: %w", err)
	}

	var frameworks []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if IsFramework(entry.Name()) {
			name := FrameworkName(entry.Name())
			frameworks = append(frameworks, name)
		}
	}

	sort.Strings(frameworks)
	return frameworks, nil
}

// GetSymbol reads a symbol's documentation by path.
// The path should be relative to the framework, e.g., "Foundation/NSString".
func GetSymbol(fsys *FS, path string) (*Document, error) {
	// Ensure path ends with .json
	if !strings.HasSuffix(path, ".json") {
		path = path + ".json"
	}
	return fsys.ReadDocument(path)
}

// GetSymbolByURL reads a symbol's documentation by its documentation URL.
// For example: "doc://com.apple.foundation/documentation/Foundation/NSString"
func GetSymbolByURL(fsys *FS, url string) (*Document, error) {
	// Extract path from URL
	// URLs are like: doc://com.apple.foundation/documentation/Foundation/NSString
	const prefix = "/documentation/"
	idx := strings.Index(url, prefix)
	if idx == -1 {
		return nil, fmt.Errorf("invalid documentation URL: %s", url)
	}

	path := url[idx+len(prefix):]
	return GetSymbol(fsys, path)
}

// ListSymbols returns all symbols in a framework.
// The framework parameter should be the framework name (e.g., "Foundation").
// Returns an error if the framework directory does not exist.
func ListSymbols(fsys *FS, framework string) ([]string, error) {
	// Check if framework directory exists
	if _, err := fsys.Stat(framework); err != nil {
		return nil, fmt.Errorf("appledocs.ListSymbols(%s): %w", framework, err)
	}

	// Use a map to deduplicate symbols (in case both .json and language variants exist)
	symbolsMap := make(map[string]bool)
	err := fs.WalkDir(fsys, framework, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		// Match .json files and language-specific variants
		// (URL-encoded: .json.language%3Dobjc = .json.language=objc, .json.language%3Dswift = .json.language=swift)
		if strings.HasSuffix(path, ".json") ||
		   strings.HasSuffix(path, ".json.language%3Dobjc") ||
		   strings.HasSuffix(path, ".json.language%3Dswift") {
			// Remove framework prefix
			symbol := strings.TrimPrefix(path, framework+string(filepath.Separator))
			// Remove .json suffix and any language suffix
			symbol = strings.TrimSuffix(symbol, ".json.language%3Dswift")
			symbol = strings.TrimSuffix(symbol, ".json.language%3Dobjc")
			symbol = strings.TrimSuffix(symbol, ".json")
			symbolsMap[symbol] = true
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("appledocs.ListSymbols(%s): walk directory: %w", framework, err)
	}

	// Convert map to sorted slice
	symbols := make([]string, 0, len(symbolsMap))
	for symbol := range symbolsMap {
		symbols = append(symbols, symbol)
	}
	sort.Strings(symbols)
	return symbols, nil
}

// FrameworkInfo contains high-level framework information.
type FrameworkInfo struct {
	Name      string
	Title     string
	Abstract  string
	Platforms []Platform
	Modules   []Module
}

// GetFrameworkInfo returns high-level information about a framework.
func GetFrameworkInfo(fsys *FS, name string) (*FrameworkInfo, error) {
	doc, err := GetFramework(fsys, name)
	if err != nil {
		return nil, err
	}

	info := &FrameworkInfo{
		Name:      name,
		Title:     doc.Metadata.Title,
		Platforms: doc.Metadata.Platforms,
		Modules:   doc.Metadata.Modules,
	}

	// Extract abstract text
	var abstractParts []string
	for _, content := range doc.Abstract {
		if content.Text != "" {
			abstractParts = append(abstractParts, content.Text)
		}
	}
	info.Abstract = strings.Join(abstractParts, " ")

	return info, nil
}

// SymbolInfo contains high-level symbol information.
type SymbolInfo struct {
	Framework  string
	Name       string
	Title      string
	Kind       string
	SymbolKind string
	Role       string
	Abstract   string
	Platforms  []Platform
	URL        string
}

// GetSymbolInfo returns high-level information about a symbol.
func GetSymbolInfo(fsys *FS, path string) (*SymbolInfo, error) {
	doc, err := GetSymbol(fsys, path)
	if err != nil {
		return nil, err
	}

	framework := FrameworkName(path)
	name := strings.TrimPrefix(path, framework+string(filepath.Separator))
	name = strings.TrimSuffix(name, ".json")

	info := &SymbolInfo{
		Framework:  framework,
		Name:       name,
		Title:      doc.Metadata.Title,
		Kind:       doc.Kind,
		SymbolKind: doc.Metadata.SymbolKind,
		Role:       doc.Metadata.Role,
		Platforms:  doc.Metadata.Platforms,
		URL:        doc.Identifier.URL,
	}

	// Extract abstract text
	var abstractParts []string
	for _, content := range doc.Abstract {
		if content.Text != "" {
			abstractParts = append(abstractParts, content.Text)
		}
	}
	info.Abstract = strings.Join(abstractParts, " ")

	return info, nil
}

// SearchSymbols searches for symbols matching the given query.
// The search is case-insensitive and matches against symbol names and titles.
func SearchSymbols(fsys *FS, framework, query string) ([]string, error) {
	symbols, err := ListSymbols(fsys, framework)
	if err != nil {
		return nil, err
	}

	query = strings.ToLower(query)
	var matches []string

	for _, symbol := range symbols {
		if strings.Contains(strings.ToLower(symbol), query) {
			matches = append(matches, symbol)
		}
	}

	return matches, nil
}

// Symbols returns an iterator over all symbols in a framework.
// The iterator yields (symbolPath, Document) pairs.
// Symbols that cannot be read are silently skipped.
//
// Requires Go 1.23+ for range-over-func support.
//
// Example:
//
//	for path, doc := range appledocs.Symbols(fsys, "Foundation") {
//	    if doc.Metadata.SymbolKind == "class" {
//	        fmt.Println("Class:", doc.Metadata.Title)
//	    }
//	}
func Symbols(fsys *FS, framework string) func(yield func(string, *Document) bool) {
	return func(yield func(string, *Document) bool) {
		symbols, err := ListSymbols(fsys, framework)
		if err != nil {
			return
		}

		for _, symbol := range symbols {
			path := filepath.Join(framework, symbol+".json")
			doc, err := fsys.ReadDocument(path)
			if err != nil {
				continue // Skip symbols that can't be read
			}

			if !yield(path, doc) {
				return
			}
		}
	}
}

// SymbolEntry represents a symbol document with its framework.
type SymbolEntry struct {
	Framework string
	Path      string
	Doc       *Document
}

// AllSymbols returns an iterator over all symbols in all frameworks.
// The iterator yields SymbolEntry structs containing framework, path, and document.
// Frameworks and symbols that cannot be read are silently skipped.
//
// Requires Go 1.23+ for range-over-func support.
//
// Example:
//
//	for entry := range appledocs.AllSymbols(fsys) {
//	    fmt.Printf("%s/%s: %s\n", entry.Framework, entry.Path, entry.Doc.Metadata.Title)
//	}
func AllSymbols(fsys *FS) func(yield func(SymbolEntry) bool) {
	return func(yield func(SymbolEntry) bool) {
		frameworks, err := ListFrameworks(fsys)
		if err != nil {
			return
		}

		for _, framework := range frameworks {
			symbols, err := ListSymbols(fsys, framework)
			if err != nil {
				continue // Skip frameworks that can't be listed
			}

			for _, symbol := range symbols {
				path := filepath.Join(framework, symbol+".json")
				doc, err := fsys.ReadDocument(path)
				if err != nil {
					continue // Skip symbols that can't be read
				}

				entry := SymbolEntry{
					Framework: framework,
					Path:      path,
					Doc:       doc,
				}

				if !yield(entry) {
					return
				}
			}
		}
	}
}
