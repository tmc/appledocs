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
	// Geometry types - returned as opaque pointers from Objective-C
	// They are actually passed by value when used as parameters, but returned as pointers
	{ObjCType: "NSRect", GoType: "unsafe.Pointer", Framework: "Foundation"},
	{ObjCType: "CGRect", GoType: "unsafe.Pointer", Framework: "Foundation"},
	{ObjCType: "NSSize", GoType: "unsafe.Pointer", Framework: "Foundation"},
	{ObjCType: "CGSize", GoType: "unsafe.Pointer", Framework: "Foundation"},
	{ObjCType: "NSPoint", GoType: "unsafe.Pointer", Framework: "Foundation"},
	{ObjCType: "CGPoint", GoType: "unsafe.Pointer", Framework: "Foundation"},
	{ObjCType: "NSRange", GoType: "unsafe.Pointer", Framework: "Foundation"},

	// Geometry types for AppKit - return actual struct values (not pointers)
	// These are returned by value from Objective-C methods and properties
	{ObjCType: "NSRect", GoType: "CGRect", Framework: "AppKit"},
	{ObjCType: "CGRect", GoType: "CGRect", Framework: "AppKit"},
	{ObjCType: "NSSize", GoType: "CGSize", Framework: "AppKit"},
	{ObjCType: "CGSize", GoType: "CGSize", Framework: "AppKit"},
	{ObjCType: "NSPoint", GoType: "CGPoint", Framework: "AppKit"},
	{ObjCType: "CGPoint", GoType: "CGPoint", Framework: "AppKit"},
	{ObjCType: "NSRange", GoType: "CGPoint", Framework: "AppKit"},  // NSRange maps to CGPoint for compatibility

	// Geometry types for CoreImage - also as unsafe.Pointer (no coregraphics imports)
	{ObjCType: "CGRect", GoType: "unsafe.Pointer", Framework: "CoreImage"},
	{ObjCType: "CGSize", GoType: "unsafe.Pointer", Framework: "CoreImage"},
	{ObjCType: "CGPoint", GoType: "unsafe.Pointer", Framework: "CoreImage"},
	{ObjCType: "CGAffineTransform", GoType: "unsafe.Pointer", Framework: "CoreImage"},

	// AppKit window and view types (enums)
	{ObjCType: "NSWindowStyleMask", GoType: "WindowStyleMask", Framework: "AppKit"},
	{ObjCType: "NSBackingStoreType", GoType: "BackingStoreType", Framework: "AppKit"},
	{ObjCType: "NSWindowOrderingMode", GoType: "WindowOrderingMode", Framework: "AppKit"},
	{ObjCType: "NSWindowLevel", GoType: "WindowLevel", Framework: "AppKit"},

	// AppKit string types
	{ObjCType: "NSString *", GoType: "string", Framework: "AppKit"},

	// Foundation date/time types - unqualified within Foundation
	{ObjCType: "NSTimeInterval", GoType: "TimeInterval", Framework: "Foundation"},
	// Foundation date/time types for AppKit - as float64 (no darwinkit imports)
	{ObjCType: "NSTimeInterval", GoType: "float64", Framework: "AppKit"},

	// CoreGraphics types
	{ObjCType: "CGFloat", GoType: "float64", Framework: "CoreGraphics"},
	{ObjCType: "CGEventRef", GoType: "unsafe.Pointer", Framework: "AppKit"},
	{ObjCType: "CGEventRef", GoType: "unsafe.Pointer", Framework: "CoreGraphics"},
	// Note: CGAffineTransform is now handled by occ2go.MapCTypeToGo and resolveType
	// Removed incorrect mapping that was stripping the "CG" prefix

	// Foundation edge enum - unqualified within Foundation
	{ObjCType: "NSRectEdge", GoType: "RectEdge", Framework: "Foundation"},
	// Foundation edge enum for AppKit - as int (no darwinkit imports)
	{ObjCType: "NSRectEdge", GoType: "int", Framework: "AppKit"},

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

// debugLogTypeMapping logs the actual type strings being looked up (for debugging)
// This helper is used during generation to understand what metadata types arrive
func debugLogTypeMapping(objcType, framework string, result string) {
	// This would normally log to stderr or a debug file
	// Enable with environment variable DEBUG_TYPE_MAPPING=1
}

// getAllMappedTypes returns all ObjC types in the registry for debugging
func getAllMappedTypes() []TypeMapping {
	return typeRegistry
}

// lookupTypeMappingDetails finds the full TypeMapping for a given Objective-C type.
func lookupTypeMappingDetails(objcType, framework string) *TypeMapping {
	objcType = strings.TrimSpace(objcType)

	// Direct lookup - try framework-specific first
	for i, mapping := range typeRegistry {
		if mapping.ObjCType == objcType && mapping.Framework != "" && mapping.Framework == framework {
			return &typeRegistry[i]
		}
	}

	// Then try framework-agnostic types
	for i, mapping := range typeRegistry {
		if mapping.ObjCType == objcType && mapping.Framework == "" {
			return &typeRegistry[i]
		}
	}

	// Then try any matching type
	for i, mapping := range typeRegistry {
		if mapping.ObjCType == objcType {
			return &typeRegistry[i]
		}
	}

	// Try without pointer suffix
	objcTypeNoPtr := strings.TrimSuffix(objcType, " *")
	if objcTypeNoPtr != objcType {
		for i, mapping := range typeRegistry {
			if mapping.ObjCType == objcTypeNoPtr && mapping.Framework != "" && mapping.Framework == framework {
				return &typeRegistry[i]
			}
		}

		for i, mapping := range typeRegistry {
			if mapping.ObjCType == objcTypeNoPtr && mapping.Framework == "" {
				return &typeRegistry[i]
			}
		}

		for i, mapping := range typeRegistry {
			if mapping.ObjCType == objcTypeNoPtr {
				return &typeRegistry[i]
			}
		}
	}

	return nil
}

// isFrameworkLocalType checks if an import path is for the same framework being generated.
// For example: if framework is "Foundation" and import is "github.com/progrium/darwinkit/macos/foundation",
// this returns true because both refer to Foundation.
func isFrameworkLocalType(importPath, framework string) bool {
	framework = strings.ToLower(framework)
	importPath = strings.ToLower(importPath)

	// Check if the framework name appears in the import path
	// This is a simplified check - in a real scenario we might want to be more precise
	return strings.Contains(importPath, framework)
}
