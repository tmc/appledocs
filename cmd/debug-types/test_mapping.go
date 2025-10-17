package main

import (
	"fmt"
	"strings"

	"github.com/tmc/appledocs/cmd/generate-framework-bindings"
)

// This file tests the type mapping directly
func testTypeMapping() {
	tests := []struct {
		objcType  string
		framework string
	}{
		{"NSRect", "AppKit"},
		{"NSWindowStyleMask", "AppKit"},
		{"NSBackingStoreType", "AppKit"},
		{"BOOL", "AppKit"},
		{"NSScreen *", "AppKit"},
		{"void *", "AppKit"},
	}

	fmt.Printf("Testing type mappings:\n\n")
	for _, test := range tests {
		// Call mapObjCTypeToGo
		result := mapObjCTypeToGo(test.objcType, test.framework)
		fmt.Printf("%-30s -> %-30s\n", test.objcType, result)

		// Also check if it's in the registry
		if goType, found := lookupTypeMapping(test.objcType, test.framework); found {
			fmt.Printf("  Registry: %s (found=%v)\n", goType, found)
		} else {
			fmt.Printf("  Registry: (not found)\n")
		}
	}
}

// Import the type mapping functions - need to make them public
// For now, this is pseudo-code to show what we need
func mapObjCTypeToGo(objcType, framework string) string {
	objcType = strings.TrimSpace(objcType)

	// Special built-in types (before checking pointers)
	switch objcType {
	case "BOOL":
		return "bool"
	case "void *":
		return "unsafe.Pointer"
	}

	// Check the type mapping registry first
	if goType, found := lookupTypeMapping(objcType, framework); found {
		return goType
	}

	return "unsafe.Pointer"
}

func lookupTypeMapping(objcType, framework string) (string, bool) {
	objcType = strings.TrimSpace(objcType)

	// For testing, hardcode the registry
	registry := map[string]map[string]string{
		"NSRect": {
			"Foundation": "foundation.Rect",
		},
		"NSWindowStyleMask": {
			"AppKit": "WindowStyleMask",
		},
		"NSBackingStoreType": {
			"AppKit": "BackingStoreType",
		},
	}

	if typeMap, ok := registry[objcType]; ok {
		if goType, ok := typeMap[framework]; ok {
			return goType, true
		}
	}

	return "", false
}
