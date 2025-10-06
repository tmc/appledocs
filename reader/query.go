package reader

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
)

// GetFramework reads a framework's root documentation.
func GetFramework(fsys *FS, name string) (*Document, error) {
	path := name + ".json"
	return fsys.ReadDocument(path)
}

// ListFrameworks returns a sorted list of all available frameworks.
func ListFrameworks(fsys *FS) ([]string, error) {
	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return nil, fmt.Errorf("list frameworks: %w", err)
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
func ListSymbols(fsys *FS, framework string) ([]string, error) {
	// Check if framework directory exists
	if _, err := fsys.Stat(framework); err != nil {
		return nil, fmt.Errorf("list symbols for %s: %w", framework, err)
	}

	var symbols []string
	err := fs.WalkDir(fsys, framework, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".json") {
			// Remove framework prefix and .json suffix
			symbol := strings.TrimPrefix(path, framework+string(filepath.Separator))
			symbol = strings.TrimSuffix(symbol, ".json")
			symbols = append(symbols, symbol)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("list symbols for %s: %w", framework, err)
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
