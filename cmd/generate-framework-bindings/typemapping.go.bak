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
var typeRegistry = []TypeMapping{}

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
			Debug.TypeMap("typeRegistry match", objcType, mapping.GoType,
				"objcType", objcType,
				"goType", mapping.GoType,
				"mappingFramework", mapping.Framework,
				"currentFramework", framework)
			// If the type is from a different framework, qualify it
			if mapping.Framework != "" && mapping.Framework != framework {
				// Check for framework hierarchy violations - if target framework is at a higher level,
				// return objectivec.IObject instead to avoid import cycles (fixes appledocs-496)
				currentLevel := getFrameworkLevel(strings.ToLower(framework))
				targetLevel := getFrameworkLevel(strings.ToLower(mapping.Framework))
				if currentLevel >= 0 && targetLevel > currentLevel {
					// Higher-level framework dependency detected - use generic interface type
					return "objectivec.IObject", true
				}

				// Don't qualify types that are already qualified (unsafe.Pointer, objc.ID, etc.)
				if strings.Contains(mapping.GoType, ".") {
					return mapping.GoType, true
				}
				return strings.ToLower(mapping.Framework) + "." + mapping.GoType, true
			}
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
				// If the type is from a different framework, qualify it
				if mapping.Framework != "" && mapping.Framework != framework {
					// Don't qualify types that are already qualified (unsafe.Pointer, objc.ID, etc.)
					if strings.Contains(mapping.GoType, ".") {
						return mapping.GoType, true
					}
					return strings.ToLower(mapping.Framework) + "." + mapping.GoType, true
				}
				return mapping.GoType, true
			}
		}
	}

	// Check cross-framework type registry (auto-discovered types)
	// Try with both ObjC names (NSColor, NSImageScaling) and Go names (Color, ImageScaling)
	// First try the stripped name (preferred) to avoid returning NSCellAttribute when we want CellAttribute
	strippedType := stripObjCPrefix(objcType)
	Debug.TypeMap("lookupTypeMapping", objcType, strippedType,
		"objcType", objcType,
		"strippedType", strippedType,
		"stripped", strippedType != objcType)

	// IMPORTANT: Check if the stripped type exists in the current framework BEFORE checking cross-framework registry
	// This prevents "Cursor" in CloudKit from resolving to appkit.Cursor instead of CKQueryCursor
	if strippedType != objcType {
		// Check if this stripped type is a class in the current framework
		if currentFrameworkClasses[strippedType] {
			Debug.TypeMap("lookupTypeMapping: found in current framework classes", objcType, strippedType,
				"framework", framework,
				"using", "unqualified name")
			return strippedType, true
		}

		// Check if it's an enum in the current framework
		// Enums are generated with stripped names (e.g., ComparisonResult not NSComparisonResult)
		// so we must return the stripped name to match the generated enum type
		Debug.TypeMap("checking currentFrameworkEnums", objcType, strippedType,
			"strippedType", strippedType,
			"enumsSize", len(currentFrameworkEnums))
		if currentFrameworkEnums[strippedType] {
			Debug.TypeMap("found in current framework enums", objcType, strippedType,
				"framework", framework,
				"returning", "STRIPPED name")
			return strippedType, true // Return STRIPPED name to match generated enum types
		}

		// Check if it's a struct in the current framework
		// Structs keep their full names with prefix (e.g., CGSize not Size)
		if currentFrameworkStructs[strippedType] {
			Debug.TypeMap("found in current framework structs", objcType, objcType,
				"framework", framework,
				"returning", "ORIGINAL name with prefix")
			return objcType, true // Return ORIGINAL name to preserve CG/NS prefix for structs
		}

		// HEURISTIC: Types with NS/CG/CA prefix that are NOT pointer types are likely enums
		// Classes are always used as pointers (*), but enums are value types
		// Only apply this if objcType does NOT contain " *" (not a pointer type)
		// IMPORTANT: Return STRIPPED name since enums are generated without prefixes
		// NOTE: This handles enums that exist but weren't extracted (see bead appledocs-473)
		Debug.TypeMap("HEURISTIC CHECK", objcType, strippedType,
			"hasPointer", strings.Contains(objcType, " *"))
		if !strings.Contains(objcType, " *") {
			// Type has a prefix and is not a pointer type - likely an enum
			Debug.TypeMap("HEURISTIC: non-pointer with prefix, likely enum", objcType, strippedType,
				"returning", "STRIPPED name")
			return strippedType, true // Return STRIPPED name to match generated enum types
		}

		// Check if it's a typedef in the current framework
		if currentFrameworkTypedefs[strippedType] {
			Debug.TypeMap("found in current framework typedefs", objcType, strippedType,
				"framework", framework,
				"using", "unqualified name")
			return strippedType, true
		}

		// Not in current framework, check cross-framework registry
		if frameworkPkg, found := crossFrameworkTypeRegistry[strippedType]; found {
			Debug.TypeMap("found in cross-framework registry (stripped)", objcType, strippedType,
				"frameworkPkg", frameworkPkg,
				"currentFramework", framework)
			// Don't qualify types with their own framework name
			if strings.ToLower(framework) == frameworkPkg {
				return strippedType, true
			}
			// NEVER qualify Go built-in primitives, even if they appear in cross-framework registry
			if isGoPrimitive(strippedType) {
				return strippedType, true
			}
			result := frameworkPkg + "." + strippedType
			Debug.TypeMap("returning qualified type", objcType, result,
				"result", result)
			return result, true
		}
		Debug.TypeMap("stripped type NOT FOUND in registry", objcType, strippedType)
	}

	// Then try the original name as fallback
	if frameworkPkg, found := crossFrameworkTypeRegistry[objcType]; found {
		Debug.TypeMap("found in cross-framework registry (original)", objcType, objcType,
			"frameworkPkg", frameworkPkg,
			"currentFramework", framework)
		// Don't qualify types with their own framework name
		if strings.ToLower(framework) == frameworkPkg {
			return objcType, true
		}
		// NEVER qualify Go built-in primitives, even if they appear in cross-framework registry
		if isGoPrimitive(objcType) {
			return objcType, true
		}
		result := frameworkPkg + "." + objcType
		Debug.TypeMap("returning qualified", objcType, result,
			"result", result)
		return result, true
	}
	Debug.TypeMap("NOT FOUND in registry", objcType, objcType)

	// Try without pointer suffix in cross-framework registry
	if objcTypeNoPtr != objcType {
		if frameworkPkg, found := crossFrameworkTypeRegistry[objcTypeNoPtr]; found {
			// Don't qualify types with their own framework name
			if strings.ToLower(framework) == frameworkPkg {
				return objcTypeNoPtr, true
			}
			// NEVER qualify Go built-in primitives, even if they appear in cross-framework registry
			if isGoPrimitive(objcTypeNoPtr) {
				return objcTypeNoPtr, true
			}
			return frameworkPkg + "." + objcTypeNoPtr, true
		}
	}

	return "", false
}

// isGoPrimitive checks if a type name is a Go built-in primitive type.
// These types should NEVER be qualified with a package name.
func isGoPrimitive(typeName string) bool {
	// Check for function types (e.g., "func()", "func(int) string")
	if strings.HasPrefix(typeName, "func(") {
		return true
	}

	goPrimitives := map[string]bool{
		"string":         true,
		"int":            true,
		"int8":           true,
		"int16":          true,
		"int32":          true,
		"int64":          true,
		"uint":           true,
		"uint8":          true,
		"uint16":         true,
		"uint32":         true,
		"uint64":         true,
		"float32":        true,
		"float64":        true,
		"bool":           true,
		"byte":           true,
		"rune":           true,
		"uintptr":        true,
		"unsafe.Pointer": true,
		"objc.ID":        true,
		"objc.Class":     true,
		"objc.SEL":       true,
	}
	return goPrimitives[typeName]
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
// UNUSED: Commented out as unreachable code
/*
func debugLogTypeMapping(objcType, framework string, result string) {
	// This would normally log to stderr or a debug file
	// Enable with environment variable DEBUG_TYPE_MAPPING=1
}
*/

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
