// Package appledocs provides simple access to Apple documentation JSON files.
//
// This package provides a minimalist API for reading Apple's documentation
// without requiring thousands of generated types. The JSON structure is
// consistent enough that map[string]interface{} with helper functions is
// sufficient for most use cases.
//
// # Basic Usage
//
//	fsys, _ := appledocs.Open("output/tutorials/data/documentation")
//	doc, _ := appledocs.LoadMap(fsys, "Foundation/NSString.json")
//	title := appledocs.Title(doc)  // "NSString"
//	kind := appledocs.SymbolKind(doc)  // "class"
package appledocs

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Open returns an fs.FS for the docs directory.
func Open(path string) (fs.FS, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("open docs: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("open docs: %s is not a directory", path)
	}
	return os.DirFS(path), nil
}

// Load unmarshals a JSON file into v.
func Load(fsys fs.FS, path string, v interface{}) error {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("parse %s: %w", path, err)
	}
	return nil
}

// LoadMap loads a JSON file as a map[string]interface{}.
func LoadMap(fsys fs.FS, path string) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := Load(fsys, path, &m)
	return m, err
}

// GetString safely extracts a string from a nested map using a path.
// Returns empty string if the path doesn't exist or value isn't a string.
//
//	title := GetString(doc, "metadata", "title")
func GetString(m map[string]interface{}, path ...string) string {
	current := m
	for i, key := range path {
		if i == len(path)-1 {
			// Last key - extract string
			if v, ok := current[key].(string); ok {
				return v
			}
			return ""
		}
		// Navigate deeper
		if next, ok := current[key].(map[string]interface{}); ok {
			current = next
		} else {
			return ""
		}
	}
	return ""
}

// GetInt safely extracts an int from a nested map.
func GetInt(m map[string]interface{}, path ...string) int {
	current := m
	for i, key := range path {
		if i == len(path)-1 {
			// JSON numbers are float64
			if v, ok := current[key].(float64); ok {
				return int(v)
			}
			return 0
		}
		if next, ok := current[key].(map[string]interface{}); ok {
			current = next
		} else {
			return 0
		}
	}
	return 0
}

// GetBool safely extracts a bool from a nested map.
func GetBool(m map[string]interface{}, path ...string) bool {
	current := m
	for i, key := range path {
		if i == len(path)-1 {
			if v, ok := current[key].(bool); ok {
				return v
			}
			return false
		}
		if next, ok := current[key].(map[string]interface{}); ok {
			current = next
		} else {
			return false
		}
	}
	return false
}

// GetMap safely extracts a nested map.
func GetMap(m map[string]interface{}, path ...string) map[string]interface{} {
	current := m
	for i, key := range path {
		if i == len(path)-1 {
			if v, ok := current[key].(map[string]interface{}); ok {
				return v
			}
			return nil
		}
		if next, ok := current[key].(map[string]interface{}); ok {
			current = next
		} else {
			return nil
		}
	}
	return current
}

// GetArray safely extracts an array.
func GetArray(m map[string]interface{}, path ...string) []interface{} {
	current := m
	for i, key := range path {
		if i == len(path)-1 {
			if v, ok := current[key].([]interface{}); ok {
				return v
			}
			return nil
		}
		if next, ok := current[key].(map[string]interface{}); ok {
			current = next
		} else {
			return nil
		}
	}
	return nil
}

// Common accessor helpers for Apple doc structure

// Kind returns the document kind ("symbol" or "article").
func Kind(m map[string]interface{}) string {
	return GetString(m, "kind")
}

// SymbolKind returns the symbol kind (class, method, property, etc).
func SymbolKind(m map[string]interface{}) string {
	return GetString(m, "metadata", "symbolKind")
}

// Title returns the document title.
func Title(m map[string]interface{}) string {
	return GetString(m, "metadata", "title")
}

// ExternalID returns the external identifier (e.g. "c:objc(cs)NSString").
func ExternalID(m map[string]interface{}) string {
	return GetString(m, "metadata", "externalID")
}

// Role returns the metadata role.
func Role(m map[string]interface{}) string {
	return GetString(m, "metadata", "role")
}

// URL returns the identifier URL.
func URL(m map[string]interface{}) string {
	return GetString(m, "identifier", "url")
}

// InterfaceLanguage returns the interface language ("swift" or "occ").
func InterfaceLanguage(m map[string]interface{}) string {
	return GetString(m, "identifier", "interfaceLanguage")
}

// References returns the references map.
func References(m map[string]interface{}) map[string]interface{} {
	return GetMap(m, "references")
}

// Metadata returns the metadata map.
func Metadata(m map[string]interface{}) map[string]interface{} {
	return GetMap(m, "metadata")
}

// Platforms returns platform information.
func Platforms(m map[string]interface{}) []interface{} {
	return GetArray(m, "metadata", "platforms")
}

// Modules returns module information.
func Modules(m map[string]interface{}) []interface{} {
	return GetArray(m, "metadata", "modules")
}

// IsFramework returns true if the path looks like a framework file.
func IsFramework(path string) bool {
	return strings.HasSuffix(path, ".json") && !strings.Contains(path, string(filepath.Separator))
}

// FrameworkName extracts the framework name from a path.
func FrameworkName(path string) string {
	name := strings.TrimSuffix(path, ".json")
	if idx := strings.Index(name, string(filepath.Separator)); idx != -1 {
		return name[:idx]
	}
	return name
}

// SymbolPath constructs the path to a symbol's JSON file.
// Example: SymbolPath("Foundation", "NSString") -> "Foundation/NSString.json"
func SymbolPath(framework, symbol string) string {
	return filepath.Join(framework, symbol) + ".json"
}
