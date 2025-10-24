package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tmc/appledocs/occ2go"
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
// It gets populated at runtime by buildTypeRegistryFromParsedData
var typeRegistry = []TypeMapping{}

// staticTypeRegistry contains framework-specific type mappings that should
// always take precedence over dynamically discovered types
var staticTypeRegistry = []TypeMapping{
	// Foundation scalar types
	{ObjCType: "NSInteger", GoType: "int", Framework: ""},
	{ObjCType: "NSUInteger", GoType: "uint", Framework: ""},

	// Foundation struct types
	{ObjCType: "NSRange", GoType: "Range", Framework: "Foundation"},

	// Objective-C runtime types
	{ObjCType: "Class", GoType: "objc.Class", Framework: ""},

	// Foundation KVO types (enums and option sets)
	{ObjCType: "NSKeyValueChange", GoType: "uint", Framework: ""},
	{ObjCType: "KeyValueChange", GoType: "uint", Framework: ""},
	{ObjCType: "NSKeyValueObservingOptions", GoType: "uint", Framework: ""},
	{ObjCType: "KeyValueObservingOptions", GoType: "uint", Framework: ""},
	{ObjCType: "NSKeyValueSetMutationKind", GoType: "uint", Framework: ""},
	{ObjCType: "KeyValueSetMutationKind", GoType: "uint", Framework: ""},

	// Foundation/AppKit string constant types (typedef NSString *)
	{ObjCType: "NSAccessibilityAttributeName", GoType: "string", Framework: ""},
	{ObjCType: "AccessibilityAttributeName", GoType: "string", Framework: ""},
	{ObjCType: "NSBindingName", GoType: "string", Framework: ""},
	{ObjCType: "BindingName", GoType: "string", Framework: ""},

	// IOBluetooth opaque types
	{ObjCType: "BluetoothL2CAPChannelRef", GoType: "uintptr", Framework: ""},
	{ObjCType: "IOReturn", GoType: "int", Framework: ""},
	{ObjCType: "Return", GoType: "int", Framework: "ObjectiveC"}, // IOBluetooth return code

	// CoreFoundation geometry types (defined in CoreFoundation, used across many frameworks)
	// These types were previously mapped to coregraphics but are actually defined in corefoundation
	{ObjCType: "CGPoint", GoType: "corefoundation.CGPoint", Framework: ""},
	{ObjCType: "CGSize", GoType: "corefoundation.CGSize", Framework: ""},
	{ObjCType: "CGRect", GoType: "corefoundation.CGRect", Framework: ""},
	{ObjCType: "CGAffineTransform", GoType: "corefoundation.CGAffineTransform", Framework: ""},
	{ObjCType: "CGFloat", GoType: "float64", Framework: ""}, // Direct mapping - CGFloat is always 64-bit on modern macOS
	{ObjCType: "CFByteOrder", GoType: "corefoundation.ByteOrder", Framework: ""},
	{ObjCType: "ByteOrder", GoType: "corefoundation.ByteOrder", Framework: ""},

	// When IN CoreFoundation, use unqualified names
	{ObjCType: "CGPoint", GoType: "CGPoint", Framework: "CoreFoundation"},
	{ObjCType: "CGSize", GoType: "CGSize", Framework: "CoreFoundation"},
	{ObjCType: "CGRect", GoType: "CGRect", Framework: "CoreFoundation"},
	{ObjCType: "CGAffineTransform", GoType: "CGAffineTransform", Framework: "CoreFoundation"},
	{ObjCType: "CGFloat", GoType: "float64", Framework: "CoreFoundation"}, // Direct mapping - no typedef needed
	{ObjCType: "CFByteOrder", GoType: "ByteOrder", Framework: "CoreFoundation"},
	{ObjCType: "ByteOrder", GoType: "ByteOrder", Framework: "CoreFoundation"},

	// When IN CoreGraphics, also use the local names (they may be type aliases)
	{ObjCType: "CGPoint", GoType: "Point", Framework: "CoreGraphics"},
	{ObjCType: "CGSize", GoType: "Size", Framework: "CoreGraphics"},
	{ObjCType: "CGRect", GoType: "Rect", Framework: "CoreGraphics"},
	{ObjCType: "CGAffineTransform", GoType: "AffineTransform", Framework: "CoreGraphics"},
	{ObjCType: "CGFloat", GoType: "float64", Framework: "CoreGraphics"}, // Direct mapping
	{ObjCType: "CFByteOrder", GoType: "int32", Framework: "CoreGraphics"}, // Use primitive type
	{ObjCType: "ByteOrder", GoType: "int32", Framework: "CoreGraphics"}, // Use primitive type
	{ObjCType: "CGDisplayReservationInterval", GoType: "float32", Framework: "CoreGraphics"}, // Display fade interval
	{ObjCType: "CGDisplayFadeInterval", GoType: "float32", Framework: "CoreGraphics"}, // Display fade interval
	{ObjCType: "CGDirectDisplayID", GoType: "uint32", Framework: "CoreGraphics"}, // Display ID
	{ObjCType: "AffineTransformComponents", GoType: "uintptr", Framework: "CoreGraphics"}, // Opaque handle (undocumented struct)
	{ObjCType: "CGAffineTransformComponents", GoType: "uintptr", Framework: "CoreGraphics"}, // Opaque handle (undocumented struct)

	// Foundation time types
	// NSTimeInterval is a typedef for double (seconds since reference date)
	{ObjCType: "NSTimeInterval", GoType: "float64", Framework: ""},
	{ObjCType: "TimeInterval", GoType: "float64", Framework: "Foundation"},
	{ObjCType: "TimeInterval", GoType: "float64", Framework: "CoreGraphics"},
}

// manualFrameworkTypes maps framework names to types that are manually defined
// (e.g., in custom files like rect_types.go) and should be registered in currentFrameworkStructs
var manualFrameworkTypes = map[string][]string{
	"coregraphics": {"Point", "Size", "Rect", "AffineTransform", "Float"},
}

// lookupTypeMapping finds a type mapping for the given Objective-C type.
// Returns the Go type and whether a mapping was found.
func lookupTypeMapping(objcType, framework string) (string, bool) {
	objcType = strings.TrimSpace(objcType)

	// Debug entry
	// if strings.Contains(objcType, "HashTableCallBacks") {
	// 	fmt.Fprintf(os.Stderr, "[DEBUG] lookupTypeMapping ENTRY: objcType=%q framework=%q\n", objcType, framework)
	// }

	// Check static type registry first (framework-specific overrides)
	// First check for exact framework match
	for _, mapping := range staticTypeRegistry {
		if mapping.ObjCType == objcType && strings.EqualFold(mapping.Framework, framework) {
			// Debug for our types
			// if strings.Contains(objcType, "HashTable") || strings.Contains(objcType, "MapTable") || strings.Contains(objcType, "EdgeInsets") {
			// 	// fmt.Fprintf(os.Stderr, "[DEBUG] static registry framework-specific: objcType=%q mapping.GoType=%q equal=%v\n",
			// 	// 	objcType, mapping.GoType, mapping.GoType == objcType)
			// }
			// If GoType equals ObjCType (identity mapping), check if we should strip the prefix
			// CG-prefixed geometry types (CGPoint, CGSize, CGRect, etc.) keep their prefix
			// NS-prefixed types like NSHashTableCallBacks → HashTableCallBacks (strip prefix)
			if mapping.GoType == objcType {
				// Preserve CG prefix for geometry types
				if strings.HasPrefix(objcType, "CG") && !strings.HasSuffix(objcType, "Ref") && !strings.HasSuffix(objcType, "Callback") {
					// Keep CG prefix for geometry types (CGPoint, CGSize, CGRect, CGAffineTransform, etc.)
					return objcType, true
				}
				// For other types, strip the prefix
				stripped := occ2go.StripObjCPrefix(objcType)
				if stripped != objcType {
					// Debug
					if strings.Contains(objcType, "HashTable") || strings.Contains(objcType, "MapTable") || strings.Contains(objcType, "EdgeInsets") {
						fmt.Fprintf(os.Stderr, "[DEBUG] Stripping in static registry: objcType=%q stripped=%q\n", objcType, stripped)
					}
					return stripped, true
				}
			}
			return mapping.GoType, true
		}
	}
	// Then check for framework-agnostic entries (Framework == "")
	for _, mapping := range staticTypeRegistry {
		if mapping.ObjCType == objcType && mapping.Framework == "" {
			// If GoType equals ObjCType (identity mapping), strip the prefix
			if mapping.GoType == objcType {
				stripped := occ2go.StripObjCPrefix(objcType)
				if stripped != objcType {
					return stripped, true
				}
			}
			return mapping.GoType, true
		}
	}

	// Direct lookup - try framework-specific first
	for _, mapping := range typeRegistry {
		if mapping.ObjCType == objcType && mapping.Framework != "" && mapping.Framework == framework {
			// Debug for our types
			if strings.Contains(objcType, "HashTable") || strings.Contains(objcType, "MapTable") || strings.Contains(objcType, "EdgeInsets") {
				fmt.Fprintf(os.Stderr, "[DEBUG] framework-specific loop: objcType=%q mapping.GoType=%q equal=%v\n",
					objcType, mapping.GoType, mapping.GoType == objcType)
			}
			// If GoType equals ObjCType (identity mapping), strip the prefix
			if mapping.GoType == objcType {
				stripped := occ2go.StripObjCPrefix(objcType)
				if stripped != objcType {
					return stripped, true
				}
			}
			return mapping.GoType, true
		}
	}

	// Then try framework-agnostic types
	for _, mapping := range typeRegistry {
		if mapping.ObjCType == objcType && mapping.Framework == "" {
			// Debug for our types
			if strings.Contains(objcType, "HashTable") || strings.Contains(objcType, "MapTable") || strings.Contains(objcType, "EdgeInsets") {
				fmt.Fprintf(os.Stderr, "[DEBUG] framework-agnostic loop: objcType=%q mapping.GoType=%q equal=%v\n",
					objcType, mapping.GoType, mapping.GoType == objcType)
			}
			// If GoType equals ObjCType (identity mapping), strip the prefix
			if mapping.GoType == objcType {
				stripped := occ2go.StripObjCPrefix(objcType)
				if stripped != objcType {
					return stripped, true
				}
			}
			return mapping.GoType, true
		}
	}

	// Then try any matching type regardless of framework
	// (geometry types from Foundation are used in AppKit, etc.)
	for _, mapping := range typeRegistry {
		if mapping.ObjCType == objcType {
			// Debug for our types
			// if strings.Contains(objcType, "HashTable") || strings.Contains(objcType, "MapTable") || strings.Contains(objcType, "EdgeInsets") {
			// 	fmt.Fprintf(os.Stderr, "[DEBUG] typeRegistry loop: objcType=%q mapping.GoType=%q mapping.Framework=%q currentFramework=%q equal=%v\n",
			// 		objcType, mapping.GoType, mapping.Framework, framework, mapping.GoType == objcType)
			// }
			// If GoType equals ObjCType (identity mapping), strip the prefix
			if mapping.GoType == objcType {
				stripped := occ2go.StripObjCPrefix(objcType)
				if stripped != objcType {
					// Return stripped type without framework qualification
					return stripped, true
				}
			}
			Debug.TypeMap("typeRegistry match", objcType, mapping.GoType,
				"objcType", objcType,
				"goType", mapping.GoType,
				"mappingFramework", mapping.Framework,
				"currentFramework", framework)
			// If the type is from a different framework, qualify it
			if mapping.Framework != "" && mapping.Framework != framework {
				// Check for framework hierarchy violations - if target framework is at a higher level,
				// return objectivec.IObject (or IObject if we're IN objectivec) to avoid import cycles (fixes appledocs-496)
				currentLevel := getFrameworkLevel(strings.ToLower(framework))
				targetLevel := getFrameworkLevel(strings.ToLower(mapping.Framework))
				if currentLevel >= 0 && targetLevel > currentLevel {
					// Higher-level framework dependency detected - use generic interface type
					fallbackType := "objectivec.IObject"
					if strings.ToLower(framework) == "objectivec" {
						fallbackType = "IObject"
					}
					return fallbackType, true
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
		// CRITICAL: Check static type registry FIRST before checking enums/classes
		// This ensures explicit type mappings (e.g., BluetoothL2CAPChannelRef -> uintptr)
		// take precedence over auto-discovered framework types
		for _, mapping := range staticTypeRegistry {
			if mapping.ObjCType == strippedType && strings.EqualFold(mapping.Framework, framework) {
				Debug.TypeMap("found stripped type in static registry (framework-specific)", strippedType, mapping.GoType,
					"objcType", objcType,
					"strippedType", strippedType,
					"goType", mapping.GoType,
					"framework", framework)
				return mapping.GoType, true
			}
		}
		// Check framework-agnostic static mappings
		for _, mapping := range staticTypeRegistry {
			if mapping.ObjCType == strippedType && mapping.Framework == "" {
				Debug.TypeMap("found stripped type in static registry (framework-agnostic)", strippedType, mapping.GoType,
					"objcType", objcType,
					"strippedType", strippedType,
					"goType", mapping.GoType)
				return mapping.GoType, true
			}
		}

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
			// Debug for our types
			// if strings.Contains(objcType, "HashTable") || strings.Contains(objcType, "MapTable") || strings.Contains(objcType, "EdgeInsets") {
			// 	fmt.Fprintf(os.Stderr, "[DEBUG] found in currentFrameworkStructs: objcType=%q strippedType=%q\n", objcType, strippedType)
			// }
			Debug.TypeMap("found in current framework structs", objcType, objcType,
				"framework", framework,
				"returning", "ORIGINAL name with prefix")
			return objcType, true // Return ORIGINAL name to preserve CG/NS prefix for structs
		}
		// IMPORTANT: Before applying heuristic, check if this type is explicitly registered in crossFrameworkTypeRegistry
		// This handles geometry types (CGPoint, CGSize, CGRect) which are structs (not pointers) but should not be treated as enums
		// See appledocs-514: CGPoint is defined in CoreFoundation, not CoreGraphics
		if frameworkPkg, found := crossFrameworkTypeRegistry[strippedType]; found {
			Debug.TypeMap("found in cross-framework registry (before heuristic)", objcType, strippedType,
				"frameworkPkg", frameworkPkg,
				"currentFramework", framework)
			// Don't qualify if it's the same framework
			if strings.ToLower(framework) == frameworkPkg {
				return strippedType, true
			}
			// NEVER qualify Go built-in primitives, even if they appear in cross-framework registry
			if isGoPrimitive(strippedType) {
				return strippedType, true
			}
			// Check framework hierarchy - prevent lower-level frameworks from importing higher-level ones
			// (fixes appledocs-519: ObjectiveC shouldn't import UIKit types)
			currentLevel := getFrameworkLevel(strings.ToLower(framework))
			targetLevel := getFrameworkLevel(frameworkPkg)
			// For ObjectiveC framework (level 0), treat ANY unknown framework as higher-level
			// Most application frameworks aren't in the levels map, so we conservatively assume they're higher
			shouldUseIObject := false
			if currentLevel >= 0 && targetLevel > currentLevel {
				shouldUseIObject = true // Known higher-level framework
			} else if currentLevel == 0 && targetLevel == -1 {
				shouldUseIObject = true // ObjectiveC importing unknown framework - assume it's higher
			}
			if shouldUseIObject {
				// Higher-level framework dependency detected - use generic interface type
				fallbackType := "objectivec.IObject"
				if strings.ToLower(framework) == "objectivec" {
					fallbackType = "IObject"
				}
				Debug.TypeMap("framework layering violation detected", objcType, fallbackType,
					"currentFramework", framework,
					"currentLevel", currentLevel,
					"targetFramework", frameworkPkg,
					"targetLevel", targetLevel)
				return fallbackType, true
			}
			result := frameworkPkg + "." + strippedType
			Debug.TypeMap("returning qualified type (before heuristic)", objcType, result,
				"result", result)
			return result, true
		}

		// HEURISTIC: Types with NS/CG/CA/UI prefix that are NOT pointer types are likely enums
		// Classes are always used as pointers (*), but enums are value types
		// Only apply this if objcType does NOT contain " *" (not a pointer type)
		// IMPORTANT: Return STRIPPED name since enums are generated without prefixes
		// NOTE: This handles enums that exist but weren't extracted (see bead appledocs-473)
		Debug.TypeMap("HEURISTIC CHECK", objcType, strippedType,
			"hasPointer", strings.Contains(objcType, " *"))
		if !strings.Contains(objcType, " *") {
			// Check for framework layering violations BEFORE returning stripped name
			// UI-prefixed types are from UIKit (level 3), which is higher than ObjectiveC (level 0)
			// (fixes appledocs-519: ObjectiveC shouldn't import UIKit enums)
			if strings.HasPrefix(objcType, "UI") && len(objcType) > 2 && objcType[2] >= 'A' && objcType[2] <= 'Z' {
				currentLevel := getFrameworkLevel(strings.ToLower(framework))
				uikitLevel := getFrameworkLevel("uikit")
				if currentLevel >= 0 && uikitLevel > currentLevel {
					fallbackType := "objectivec.IObject"
					if strings.ToLower(framework) == "objectivec" {
						fallbackType = "IObject"
					}
					Debug.TypeMap("HEURISTIC: UI-prefixed type, framework layering violation", objcType, fallbackType,
						"currentFramework", framework,
						"currentLevel", currentLevel,
						"uikitLevel", uikitLevel)
					return fallbackType, true
				}
			}
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
			// Check framework hierarchy - prevent lower-level frameworks from importing higher-level ones
			// (fixes appledocs-519: ObjectiveC shouldn't import UIKit types)
			currentLevel := getFrameworkLevel(strings.ToLower(framework))
			targetLevel := getFrameworkLevel(frameworkPkg)
			// For ObjectiveC framework (level 0), treat ANY unknown framework as higher-level
			// Most application frameworks aren't in the levels map, so we conservatively assume they're higher
			shouldUseIObject := false
			if currentLevel >= 0 && targetLevel > currentLevel {
				shouldUseIObject = true // Known higher-level framework
			} else if currentLevel == 0 && targetLevel == -1 {
				shouldUseIObject = true // ObjectiveC importing unknown framework - assume it's higher
			}
			if shouldUseIObject {
				// Higher-level framework dependency detected - use generic interface type
				fallbackType := "objectivec.IObject"
				if strings.ToLower(framework) == "objectivec" {
					fallbackType = "IObject"
				}
				Debug.TypeMap("framework layering violation detected (stripped path)", objcType, fallbackType,
					"currentFramework", framework,
					"currentLevel", currentLevel,
					"targetFramework", frameworkPkg,
					"targetLevel", targetLevel)
				return fallbackType, true
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
// Note: This queries crossFrameworkTypeRegistry which is populated at runtime.
func getAllAppKitEnumTypes() []string {
	var result []string
	seen := make(map[string]bool)

	// Query the cross-framework registry for AppKit types
	for typeName, framework := range crossFrameworkTypeRegistry {
		if framework == "appkit" && !seen[typeName] {
			// Return the Go-style name (already stripped of NS prefix in registry)
			result = append(result, typeName)
			seen[typeName] = true
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
// Note: This queries crossFrameworkTypeRegistry which is populated at runtime.
func getAllMappedTypes() []TypeMapping {
	var result []TypeMapping
	seen := make(map[string]bool) // To avoid duplicates

	// Convert crossFrameworkTypeRegistry to TypeMapping structs
	// The registry contains both ObjC names (NSWindow) and Go names (Window)
	for typeName, framework := range crossFrameworkTypeRegistry {
		// Check if this looks like an ObjC name (has NS/CG/CA/etc prefix)
		stripped := stripObjCPrefix(typeName)
		if stripped != typeName {
			// This is an ObjC name
			objcName := typeName
			goName := stripped

			// Create a TypeMapping with both names
			key := objcName + ":" + framework
			if !seen[key] {
				result = append(result, TypeMapping{
					ObjCType:  objcName,
					GoType:    goName,
					Framework: framework,
				})
				seen[key] = true
			}
		} else {
			// This is already a Go name, check if we haven't already added it
			// via its ObjC counterpart
			key := typeName + ":" + framework
			if !seen[key] {
				// Add it as both ObjC and Go name (for types without prefix)
				result = append(result, TypeMapping{
					ObjCType:  typeName,
					GoType:    typeName,
					Framework: framework,
				})
				seen[key] = true
			}
		}
	}

	return result
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
