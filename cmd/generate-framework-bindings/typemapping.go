package main

import (
	"strings"
)

// TypeMapping represents a mapping from an Objective-C type to a Go type.
type TypeMapping struct {
	// ObjCType is the Objective-C type to match (e.g., "NSRect", "NSWindowStyleMask")
	ObjCType string
	// GoType is the corresponding Go type (e.g., "foundation.Rect", "WindowStyleMask")
	GoType string
	// Framework is the framework this type belongs to (e.g., "AppKit", "Foundation", "CoreGraphics")
	Framework string
	// RequiresImport is the import path if this type needs to be imported (e.g., "github.com/progrium/darwinkit/macos/foundation")
	RequiresImport string
}

// typeRegistry contains all known Objective-C to Go type mappings
var typeRegistry = []TypeMapping{
	// Foundation geometry types
	{ObjCType: "NSRect", GoType: "foundation.Rect", Framework: "Foundation", RequiresImport: "github.com/progrium/darwinkit/macos/foundation"},
	{ObjCType: "CGRect", GoType: "foundation.Rect", Framework: "Foundation", RequiresImport: "github.com/progrium/darwinkit/macos/foundation"},
	{ObjCType: "NSSize", GoType: "foundation.Size", Framework: "Foundation", RequiresImport: "github.com/progrium/darwinkit/macos/foundation"},
	{ObjCType: "CGSize", GoType: "foundation.Size", Framework: "Foundation", RequiresImport: "github.com/progrium/darwinkit/macos/foundation"},
	{ObjCType: "NSPoint", GoType: "foundation.Point", Framework: "Foundation", RequiresImport: "github.com/progrium/darwinkit/macos/foundation"},
	{ObjCType: "CGPoint", GoType: "foundation.Point", Framework: "Foundation", RequiresImport: "github.com/progrium/darwinkit/macos/foundation"},
	{ObjCType: "NSRange", GoType: "foundation.Range", Framework: "Foundation", RequiresImport: "github.com/progrium/darwinkit/macos/foundation"},

	// AppKit window and view types (enums)
	{ObjCType: "NSWindowStyleMask", GoType: "WindowStyleMask", Framework: "AppKit"},
	{ObjCType: "NSBackingStoreType", GoType: "BackingStoreType", Framework: "AppKit"},
	{ObjCType: "NSWindowOrderingMode", GoType: "WindowOrderingMode", Framework: "AppKit"},
	{ObjCType: "NSWindowLevel", GoType: "WindowLevel", Framework: "AppKit"},

	// AppKit string types
	{ObjCType: "NSString *", GoType: "string", Framework: "AppKit"},

	// Foundation date/time types
	{ObjCType: "NSTimeInterval", GoType: "foundation.TimeInterval", Framework: "Foundation", RequiresImport: "github.com/progrium/darwinkit/macos/foundation"},

	// CoreGraphics types
	{ObjCType: "CGFloat", GoType: "float64", Framework: "CoreGraphics"},
	{ObjCType: "CGAffineTransform", GoType: "coregraphics.AffineTransform", Framework: "CoreGraphics", RequiresImport: "github.com/progrium/darwinkit/macos/coregraphics"},

	// Foundation edge enum
	{ObjCType: "NSRectEdge", GoType: "foundation.RectEdge", Framework: "Foundation", RequiresImport: "github.com/progrium/darwinkit/macos/foundation"},

	// AppKit event types
	{ObjCType: "NSEventType", GoType: "EventType", Framework: "AppKit"},
	{ObjCType: "NSEventModifierFlags", GoType: "EventModifierFlags", Framework: "AppKit"},
}

// lookupTypeMapping finds a type mapping for the given Objective-C type.
// Returns the Go type and whether a mapping was found.
func lookupTypeMapping(objcType, framework string) (string, bool) {
	objcType = strings.TrimSpace(objcType)

	// Direct lookup - try framework-specific first
	for _, mapping := range typeRegistry {
		if mapping.ObjCType == objcType && mapping.Framework != "" && mapping.Framework == framework {
			return mapping.GoType, true
		}
	}

	// Then try framework-agnostic types
	for _, mapping := range typeRegistry {
		if mapping.ObjCType == objcType && mapping.Framework == "" {
			return mapping.GoType, true
		}
	}

	// Then try any matching type regardless of framework
	// (geometry types from Foundation are used in AppKit, etc.)
	for _, mapping := range typeRegistry {
		if mapping.ObjCType == objcType {
			return mapping.GoType, true
		}
	}

	// Try without pointer suffix for object types
	objcTypeNoPtr := strings.TrimSuffix(objcType, " *")
	if objcTypeNoPtr != objcType {
		// Direct lookup for no-pointer version
		for _, mapping := range typeRegistry {
			if mapping.ObjCType == objcTypeNoPtr && mapping.Framework != "" && mapping.Framework == framework {
				return mapping.GoType, true
			}
		}

		// Then try framework-agnostic types
		for _, mapping := range typeRegistry {
			if mapping.ObjCType == objcTypeNoPtr && mapping.Framework == "" {
				return mapping.GoType, true
			}
		}

		// Then try any matching type
		for _, mapping := range typeRegistry {
			if mapping.ObjCType == objcTypeNoPtr {
				return mapping.GoType, true
			}
		}
	}

	return "", false
}

// getTypeImportPath returns the import path needed for a given Go type, if any.
func getTypeImportPath(objcType, framework string) string {
	objcType = strings.TrimSpace(objcType)

	for _, mapping := range typeRegistry {
		if mapping.ObjCType == objcType || (strings.TrimSuffix(objcType, " *") == mapping.ObjCType) {
			if mapping.Framework == "" || mapping.Framework == framework {
				return mapping.RequiresImport
			}
		}
	}

	return ""
}

// getAllAppKitEnumTypes returns all AppKit enum type names that we can generate.
// This helps in determining which types are available for generation.
func getAllAppKitEnumTypes() []string {
	var result []string
	seen := make(map[string]bool)

	for _, mapping := range typeRegistry {
		if mapping.Framework == "AppKit" && !strings.Contains(mapping.GoType, ".") {
			if !seen[mapping.GoType] {
				result = append(result, mapping.GoType)
				seen[mapping.GoType] = true
			}
		}
	}

	return result
}
