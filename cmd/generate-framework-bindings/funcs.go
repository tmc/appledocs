package main

import (
	"fmt"
)

// dict creates a map from alternating key-value pairs.
// Usage: {{template "name" (dict "key1" .Value1 "key2" .Value2)}}
func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("dict requires an even number of arguments")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

// sliceContainsString checks if a string slice contains a specific string.
func sliceContainsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// ClassImports holds the import paths needed for a class
type ClassImports struct {
	// DEPRECATED: Legacy boolean fields for backward compatibility.
	// Use ImportPaths map instead for dynamic, extensible import detection.
	// These fields are populated automatically from ImportPaths for now,
	// but will be removed in a future version.
	NeedsObjectiveC             bool // Deprecated: Use ImportPaths["objectivec"]
	NeedsFoundation             bool // Deprecated: Use ImportPaths["foundation"]
	NeedsQuartzCore             bool // Deprecated: Use ImportPaths["quartzcore"]
	NeedsCoreGraphics           bool // Deprecated: Use ImportPaths["coregraphics"]
	NeedsCloudKit               bool // Deprecated: Use ImportPaths["cloudkit"]
	NeedsAppKit                 bool // Deprecated: Use ImportPaths["appkit"]
	NeedsUserNotifications      bool // Deprecated: Use ImportPaths["usernotifications"]
	NeedsUniformTypeIdentifiers bool // Deprecated: Use ImportPaths["uniformtypeidentifiers"]

	// ImportPaths is a dynamic map of all import paths needed (package name -> import path)
	// This replaces the need for hard-coded boolean fields above.
	// Keys are package names (e.g., "foundation", "appkit"), values are full import paths.
	ImportPaths map[string]string
}

// sortedKeys returns the keys of a map in sorted order, suitable for templates.
// This is useful for generating consistent output across template generations.
func sortedKeys(m map[string]bool) []string {
	if m == nil {
		return []string{}
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// Simple bubble sort for consistency
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if keys[i] > keys[j] {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}

	return keys
}

// ImportInfo holds information about an import for template rendering
type ImportInfo struct {
	PackageName string
	ImportPath  string
}
