package main

import (
	"fmt"
	"strings"

	"github.com/tmc/appledocs/occ2go"
)

// mapCTypeToGoWithFramework wraps occ2go.MapCTypeToGo and applies framework-specific type mappings.
// This ensures C types like CGAffineTransform are properly qualified with their framework package.
func mapCTypeToGoWithFramework(cType, framework string) string {
	// Check if this C type is a typedef in the current framework BEFORE mapping
	// This handles CF*Ref types (CFTypeRef, CFAllocatorRef, etc.)
	strippedCType := stripObjCPrefix(cType)
	if currentFrameworkTypedefs[strippedCType] {
		// This is a typedef in the current framework - return the stripped name
		// E.g., CFTypeRef -> TypeRef, CFAllocatorRef -> AllocatorRef
		return strippedCType
	}

	// First apply occ2go's basic C type mapping
	goType := occ2go.MapCTypeToGo(cType, framework)

	// If the result is a Go primitive type (int, float32, string, etc.) or a slice,
	// return it directly without further processing. This prevents incorrectly
	// passing Go types back through MapCTypeToGo, which would map "float32" to
	// "unsafe.Pointer" (since MapCTypeToGo doesn't recognize "float32" as a C type).
	goPrimitives := map[string]bool{
		"":               true, // void
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
	}
	if goPrimitives[goType] || strings.HasPrefix(goType, "[]") {
		return goType
	}

	// For non-primitives (CGColorRef, NSWindow, etc.), apply framework-specific mapping
	// This handles prefix stripping (CGColorRef -> ColorRef) and cross-framework qualification
	mapped := mapObjCTypeToGo(goType, framework)

	return mapped
}

// mapObjCTypeToGo maps Objective-C types to Go types for darwinkit style.
// Examples:
//
//	NSString * -> string
//	id -> objc.Object
//	NSButton * -> Button (interface type in parameters)
//	NSRect -> foundation.Rect
//	NSWindowStyleMask -> WindowStyleMask
//	NSUInteger [] -> []uint (array syntax correction)
func mapObjCTypeToGo(objcType, framework string) string {
	objcType = strings.TrimSpace(objcType)

	// Handle C-style array syntax: "Type []" or "Type[]" -> "[]Type"
	// ObjC sometimes uses array syntax like "const NSUInteger[]" which should become "[]uint" in Go
	if strings.HasSuffix(objcType, "[]") || strings.HasSuffix(objcType, " []") {
		// Remove the array brackets
		arrayType := strings.TrimSuffix(strings.TrimSuffix(objcType, "[]"), " []")
		arrayType = strings.TrimSpace(arrayType)
		// Remove "const" keyword if present
		arrayType = strings.TrimPrefix(arrayType, "const ")
		arrayType = strings.TrimSpace(arrayType)
		// Recursively map the base type, then prepend []
		baseGoType := mapObjCTypeToGo(arrayType, framework)
		return "[]" + baseGoType
	}

	Debug.TypeMap("mapObjCTypeToGo entry", objcType, framework,
		"objcType", objcType,
		"framework", framework)

	// Strip self-package qualifications from Swift documentation
	// Swift docs often use module.Type format (e.g., uniformtypeidentifiers.UTType)
	// When generating the same framework, we should use unqualified names
	if framework != "" {
		// Build package prefix (e.g., "uniformtypeidentifiers." from "UniformTypeIdentifiers")
		packagePrefix := strings.ToLower(framework) + "."
		// Check for exact package.Type pattern
		if strings.HasPrefix(strings.ToLower(objcType), packagePrefix) {
			// Extract the type name after the dot
			// E.g., "uniformtypeidentifiers.UTType" -> "UTType"
			parts := strings.SplitN(objcType, ".", 2)
			if len(parts) == 2 {
				objcType = parts[1]
			}
		}
	}

	// Strip type qualifiers (__kindof, const, etc.) using occ2go utility
	objcType = occ2go.StripTypeQualifiers(objcType)

	// Handle id<Protocol> pattern (e.g., "id<NSFetchRequestResult>" -> "objc.ID")
	// This is Objective-C's protocol conformance syntax
	if occ2go.IsProtocolType(objcType) {
		return "objc.ID"
	}

	// Handle []id<Protocol> pattern (e.g., "[]id<NSFetchRequestResult>" -> "[]objc.ID")
	if occ2go.IsArrayType(objcType) {
		elementType := occ2go.GetArrayElementType(objcType)
		if occ2go.IsProtocolType(elementType) {
			return "[]objc.ID"
		}
	}

	// Handle array types that are already converted by occ2go (e.g., "[]void (^)(void)" -> "[]unsafe.Pointer")
	// This handles cases where occ2go has already converted NSArray<T> to []T
	// We need to recursively map the element type
	if occ2go.IsArrayType(objcType) {
		elementType := occ2go.GetArrayElementType(objcType)
		goElementType := mapObjCTypeToGo(elementType, framework)
		return "[]" + goElementType
	}

	// Handle Objective-C generic types (e.g., NSArray<NSString *>, NSArray<SCDisplay *>)
	if occ2go.IsGenericType(objcType) {
		// Check if this is NSDictionary - map to IDictionary
		if strings.Contains(objcType, "NSDictionary") {
			// NSDictionary<K, V> -> IDictionary
			// We can't represent the generic key/value types in Go, so use the interface
			if framework == "Foundation" {
				return "IDictionary"
			}
			return "foundation.IDictionary"
		}

		// Extract NSArray element type: NSArray<ElementType *> -> []ElementType
		// Handle both "NSArray<T>" and "NSArray<T> *" patterns
		elementType := occ2go.ExtractGenericElementType(objcType)
		if elementType != "" {
			// Successfully extracted NSArray element type
			Debug.Object("ExtractGenericElementType", objcType, elementType,
				"objcType", objcType,
				"elementType", elementType)

			// Special case: NSString -> string
			if elementType == "NSString" {
				return "[]string"
			}

			// Special case: NSObject with protocol conformance (e.g., NSObject<SomeProtocol>)
			// These should map to []objectivec.IObject
			if strings.HasPrefix(elementType, "NSObject<") {
				Debug.Object("NSObject< pattern match", objcType, elementType,
					"returning", "[]objectivec.IObject")
				return "[]objectivec.IObject"
			}

			// Strip common Apple prefixes from element types
			// This ensures NSArray<SCDisplay *> -> []Display, NSArray<NSButton *> -> []Button
			strippedType := stripObjCPrefix(elementType)
			if strippedType != elementType {
				// Prefix was stripped - check if this type exists in current framework
				if currentFrameworkClasses[strippedType] {
					// Type is defined in current framework, safe to use
					return "[]" + strippedType
				}
				// Cross-framework reference - check type mapping registry for explicit mapping
				// Foundation defines URL and Number, not NSURL and NSNumber
				// Try both the stripped type name and the original element type
				if goType, found := lookupTypeMapping(strippedType, framework); found {
					return "[]" + goType
				}
				if goType, found := lookupTypeMapping(elementType, framework); found {
					return "[]" + goType
				}
				// Last resort - fall back to unsafe.Pointer for array elements
				return "[]unsafe.Pointer"
			}

			// For other types, try to map them
			goElementType := mapObjCTypeToGo(elementType, framework)
			if goElementType == "unsafe.Pointer" {
				// If mapping failed, use the element type directly
				return "[]" + elementType
			}
			return "[]" + goElementType
		}
		// For other generic types, fall back to unsafe.Pointer
		return "unsafe.Pointer"
	}

	// Special built-in types (before checking pointers)
	switch objcType {
	case "id":
		return "objc.ID"
	case "Class":
		return "objc.Class"
	case "SEL":
		return "objc.SEL"
	case "BOOL":
		return "bool"
	case "NSInteger", "Int":
		return "int"
	case "NSUInteger", "UInt":
		return "uint"
	case "unsigned long long", "UInt64", "uint64_t":
		return "uint64"
	case "CGFloat", "Double":
		return "float64"
	case "Float":
		return "float32"
	case "void":
		return ""
	case "String", "String?":
		// Swift string types map to Go string
		return "string"
	}

	// Check the type mapping registry first (includes both with and without pointers)
	// This must come before the block check so that mapped block types (e.g., void (^)(void) -> func())
	// are handled correctly
	if goType, found := lookupTypeMapping(objcType, framework); found {
		Debug.TypeMap("found in registry", objcType, goType,
			"objcType", objcType,
			"goType", goType,
			"framework", framework)
		return goType
	}

	Debug.TypeMap("NOT found in registry", objcType, framework,
		"objcType", objcType,
		"framework", framework)

	// NOTE: Removed hardcoded block-to-unsafe.Pointer mapping here.
	// Block type mapping is now handled by occ2go.MapCTypeToGo which converts
	// blocks to Go function types (e.g., void (^)(void) -> func()).
	// This happens in the fallback call to occ2go.MapCTypeToGo below.

	// Handle pointers for types not in the registry
	isPointer := occ2go.IsPointerType(objcType)
	objcTypeNoPtr := occ2go.StripPointer(objcType)

	// Special case: NSString * -> string (most common string parameter type)
	if isPointer && objcTypeNoPtr == "NSString" {
		return "string"
	}

	// Also handle NSString without pointer (from property types in docs)
	if objcType == "NSString" {
		return "string"
	}

	// Check registry again for type without pointer
	if isPointer && objcTypeNoPtr != objcType {
		Debug.TypeMap("pointer check", objcType, objcTypeNoPtr,
			"objcType", objcType,
			"objcTypeNoPtr", objcTypeNoPtr,
			"checking", "lookupTypeMapping")
		if goType, found := lookupTypeMapping(objcTypeNoPtr, framework); found {
			Debug.TypeMap("pointer-to-enum found", objcType, goType,
				"objcType", objcType,
				"objcTypeNoPtr", objcTypeNoPtr,
				"goType", goType)
			return goType
		}
		Debug.TypeMap("pointer not found", objcType, objcTypeNoPtr,
			"objcType", objcType,
			"objcTypeNoPtr", objcTypeNoPtr)
	}

	// Also try stripping prefix for NON-pointer types (handles docs that omit the *)
	// This is especially common for return types and property types
	// IMPORTANT: Do NOT strip prefix for enums - they need to keep their NS prefix
	if !isPointer && objcType != "" {
		strippedType := stripObjCPrefix(objcType)
		Debug.TypeMap("non-pointer block", objcType, strippedType,
			"objcType", objcType,
			"strippedType", strippedType,
			"isPointer", isPointer,
			"framework", framework)
		if strippedType != objcType {
			// Check if this is an enum type - enums use stripped names to match the generated type definitions
			if currentFrameworkEnums[strippedType] {
				// This is an enum - return the stripped type to match generated enum type names
				Debug.TypeMap("is enum, returning stripped", objcType, strippedType,
					"objcType", objcType,
					"strippedType", strippedType)
				return strippedType
			}

			// Don't check currentFrameworkStructs here for early return!
			// Structs can be cross-framework (e.g., CGPoint defined in CoreFoundation but used in ObjectiveC).
			// Let resolveType() handle framework qualification below.

			Debug.TypeMap("not in currentFrameworkEnums", objcType, strippedType,
				"strippedType", strippedType,
				"enumsSize", len(currentFrameworkEnums))

			// Successfully stripped a prefix - check if this is a known type
			// in the current framework or type registry
			if mappedGoType, found := lookupTypeMapping(strippedType, framework); found {
				Debug.TypeMap("lookupTypeMapping found", strippedType, mappedGoType,
					"strippedType", strippedType,
					"mappedGoType", mappedGoType,
					"framework", framework)
				return mappedGoType
			}
			// Let it fall through to use strippedType and then resolve it
			// resolveType will check crossFrameworkTypeRegistry to properly qualify cross-framework types
			resolvedType := resolveType(framework, strippedType)
			Debug.TypeMap("resolveType called", strippedType, resolvedType,
				"framework", framework,
				"strippedType", strippedType,
				"resolvedType", resolvedType)
			// If resolveType returned unsafe.Pointer, return the ORIGINAL objcType
			// so it can be collected as an undefined type with its full name
			if resolvedType == "unsafe.Pointer" {
				return objcType
			}
			return resolvedType
		}
	}

	//  For pointer types to ObjC classes, try stripping prefix BEFORE falling back to MapCTypeToGo
	// This allows cross-framework type resolution to work properly
	goType := ""
	if isPointer && objcTypeNoPtr != "" {
		strippedType := stripObjCPrefix(objcTypeNoPtr)
		Debug.TypeMap("pointer stripping", objcType, strippedType,
			"objcType", objcType,
			"objcTypeNoPtr", objcTypeNoPtr,
			"strippedType", strippedType,
			"framework", framework)
		if strippedType != objcTypeNoPtr {
			// Successfully stripped a prefix - this is likely an ObjC class type
			// Use the stripped type and let resolveType find the right framework
			goType = strippedType
			Debug.TypeMap("set goType from stripped", objcType, goType,
				"goType", goType,
				"from", objcType)
		}
	}

	// Fall back to occ2go mapping if we haven't resolved it yet
	if goType == "" {
		goType = occ2go.MapCTypeToGo(objcType, framework)
		Debug.TypeMap("after occ2go.MapCTypeToGo", objcType, goType,
			"objcType", objcType,
			"goType", goType)

		// If occ2go.MapCTypeToGo returned a type with NS/CG/CA prefix, strip it
		// This handles enums that exist but weren't extracted (see bead appledocs-473)
		// Example: NSEnergyFormatterUnit -> EnergyFormatterUnit
		// EXCEPTION: Don't strip prefix for types defined in current framework (classes, enums, typedefs, structs)
		strippedGoType := stripObjCPrefix(goType)
		// Check if stripped type exists in current framework before stripping
		inCurrentFramework := currentFrameworkClasses[strippedGoType] ||
			currentFrameworkEnums[strippedGoType] ||
			currentFrameworkTypedefs[strippedGoType] ||
			currentFrameworkStructs[strippedGoType]

		if strippedGoType != goType && goType != "unsafe.Pointer" && !inCurrentFramework {
			Debug.TypeMap("stripping prefix from fallback", goType, strippedGoType,
				"from", goType,
				"to", strippedGoType)
			goType = strippedGoType
		} else {
			Debug.TypeMap("NOT stripping fallback", goType, strippedGoType,
				"goType", goType,
				"strippedGoType", strippedGoType,
				"equal", strippedGoType == goType,
				"isUnsafe", goType == "unsafe.Pointer",
				"inCurrentFramework", inCurrentFramework)
		}
	}

	// Never return empty string for a type - default to unsafe.Pointer
	if goType == "" {
		return "unsafe.Pointer"
	}

	// Resolve cross-framework types (e.g., CGAffineTransform -> coregraphics.CGAffineTransform)
	resolvedType := resolveType(framework, goType)
	Debug.TypeMap("before/after resolve", objcType, resolvedType,
		"objcType", objcType,
		"framework", framework,
		"beforeResolve", goType,
		"afterResolve", resolvedType)
	goType = resolvedType

	// Check for framework hierarchy violations - if resolved type references a higher-level framework,
	// map to generic objectivec.IObject instead to avoid import cycles
	// This handles both struct types (e.g., "replaykit.BroadcastConfiguration") and interface types
	// (e.g., "replaykit.IRPBroadcastConfiguration")
	if strings.Contains(goType, ".") && framework != "" {
		parts := strings.Split(goType, ".")
		if len(parts) >= 2 {
			targetFramework := parts[0]
			currentLevel := getFrameworkLevel(strings.ToLower(framework))
			targetLevel := getFrameworkLevel(targetFramework)

			Debug.Hierarchy("checking hierarchy", goType, framework,
				"goType", goType,
				"framework", framework,
				"currentLevel", currentLevel,
				"targetFramework", targetFramework,
				"targetLevel", targetLevel)

			if currentLevel >= 0 && targetLevel > currentLevel {
				// Hierarchy violation - map to objectivec.IObject
				Debug.Hierarchy("hierarchy violation", framework, targetFramework,
					"currentFramework", framework,
					"currentLevel", currentLevel,
					"targetFramework", targetFramework,
					"targetLevel", targetLevel,
					"mapping", "objectivec.IObject")
				return "objectivec.IObject"
			}
		}
	}

	Debug.TypeMap("exit", objcType, goType,
		"objcType", objcType,
		"framework", framework,
		"returning", goType)

	return goType
}

// getFrameworkLevel returns the hierarchy level for a framework (0-4), or -1 if unknown
func getFrameworkLevel(framework string) int {
	// Access the frameworkLevels map from framework_hierarchy.go
	// We need to import this or duplicate the levels here
	// For now, duplicate the essential levels
	levels := map[string]int{
		"objc":                   0,
		"objectivec":             0,
		"coregraphics":           1,
		"corefoundation":         1,
		"foundation":             1,
		"coretext":               1,
		"iosurface":              1,
		"uniformtypeidentifiers": 2, // UniformTypeIdentifiers depends on Foundation
		"coreimage":              2,
		"quartzcore":             2,
		"coreaudio":              2,
		"coremidi":               2,
		"imageio":                2,
		"coredata":               2,
		"corelocation":           2,
		"corespotlight":          2,
		"network":                2,
		"security":               2,
		"corebluetooth":          2,
		"corevideo":              2,
		"coreml":                 2,
		"vision":                 2,
		"naturallanguage":        2,
		"appkit":                 3,
		"uikit":                  3,
		"webkit":                 3,
		"pdfkit":                 3,
		"networkextension":       3,
		"avfoundation":           4,
		"avfaudio":               4,
		"avkit":                  4,
		"avrouting":              4,
		"audiotoolbox":           4,
		"cloudkit":               4,
		"contacts":               4,
		"contactsui":             4,
		"gameplaykit":            4,
		"intents":                4,
		"intentsui":              4,
		"metal":                  4,
		"metalkit":               4,
		"eventkit":               4,
		"healthkit":              4,
		"homekit":                4,
		"mapkit":                 4,
		"messages":               4,
		"storekit":               4,
		"usernotifications":      4,
		"replaykit":              4,
	}
	if level, ok := levels[strings.ToLower(framework)]; ok {
		return level
	}
	return -1
}

// resolveType resolves a type name to its fully qualified name, handling cross-framework dependencies.
// Takes the current framework context and a type name (e.g., "MutableAttributedString") and returns
// either the unqualified name (if it's in the same framework) or a qualified name (e.g., "foundation.MutableAttributedString").
// This helper is used in templates to properly reference types that may come from other frameworks.
//
// Examples:
//
//	resolveType("AppKit", "Button") -> "Button" (same framework)
//	resolveType("AppKit", "MutableAttributedString") -> "foundation.MutableAttributedString" (cross-framework)
//	resolveType("Foundation", "Array") -> "Array" (same framework)
func resolveType(framework, typeName string) string {
	Debug.TypeMap("resolveType entry", typeName, framework,
		"framework", framework,
		"typeName", typeName)

	if typeName == "" {
		return ""
	}

	// If the type is already qualified with a package name (e.g., "corefoundation.Point"),
	// return it as-is. This prevents resolveType from re-processing types that have
	// already been resolved by lookupTypeMapping.
	// See appledocs-514: CGPoint should stay as corefoundation.Point, not become coregraphics.CGPoint
	if strings.Contains(typeName, ".") {
		// Check if it's a qualified type (package.Type format)
		parts := strings.Split(typeName, ".")
		if len(parts) == 2 {
			pkgName := parts[0]
			// Verify it looks like a valid package name (all lowercase, no special chars except maybe digits)
			isValidPkg := true
			for _, ch := range pkgName {
				if !((ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')) {
					isValidPkg = false
					break
				}
			}
			if isValidPkg {
				Debug.TypeMap("already qualified, returning as-is", typeName, framework,
					"typeName", typeName,
					"package", pkgName,
					"framework", framework)
				return typeName
			}
		}
	}

	// Check for function types (e.g., "func()", "func(int) string")
	// Function types are Go primitives and should never be qualified
	if strings.HasPrefix(typeName, "func(") {
		Debug.TypeMap("returning function type unqualified", typeName, framework,
			"typeName", typeName,
			"framework", framework)
		return typeName
	}

	// Never qualify Go primitives - they should always be unqualified
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
	}
	if goPrimitives[typeName] {
		Debug.TypeMap("returning primitive unqualified", typeName, framework,
			"typeName", typeName,
			"framework", framework)
		return typeName
	}

	// ALSO check for capital-S String which should map to lowercase string
	// This happens when NSString typedef resolves to "String" instead of "string"
	if typeName == "String" {
		Debug.TypeMap("converting String to string", typeName, "string",
			"from", "String",
			"to", "string",
			"framework", framework)
		return "string"
	}

	// Strip self-package qualifications (e.g., foundation.NSString in Foundation -> NSString)
	// This prevents incorrect qualification like foundation.NSOrderedCollectionChange in the foundation package
	packagePrefix := strings.ToLower(framework) + "."
	if strings.HasPrefix(typeName, packagePrefix) {
		return strings.TrimPrefix(typeName, packagePrefix)
	}

	// Check if the type exists in current framework FIRST before adding qualifications
	// This prevents self-imports (e.g., coregraphics.CGAffineTransform in CoreGraphics)
	// Check classes, enums, typedefs, and structs - all stored with stripped ObjC prefixes
	strippedTypeName := stripObjCPrefix(typeName)
	if currentFrameworkClasses[strippedTypeName] || currentFrameworkEnums[strippedTypeName] || currentFrameworkTypedefs[strippedTypeName] || currentFrameworkStructs[strippedTypeName] {
		// DEBUG: Uncomment to debug same-framework type resolution
		// fmt.Fprintf(os.Stderr, "DEBUG resolveType: Found '%s' (stripped: '%s') in current framework '%s', returning as-is\n", typeName, strippedTypeName, framework)
		// It's in the current framework, return as-is
		return typeName
	}

	// Check if we know about this type from the cross-framework registry
	// This automatically handles ALL cross-framework types without hardcoding
	if frameworkPkg, found := crossFrameworkTypeRegistry[typeName]; found {
		Debug.TypeMap("registry lookup hit", typeName, frameworkPkg,
			"typeName", typeName,
			"frameworkPkg", frameworkPkg,
			"currentFramework", framework)
		// Don't qualify types with their own framework name (e.g., foundation.NSString in foundation package)
		if strings.ToLower(framework) == frameworkPkg {
			Debug.TypeMap("same framework, returning unqualified", typeName, frameworkPkg)
			return typeName
		}
		// NEVER qualify Go built-in primitives, even if they appear in cross-framework registry
		// This prevents errors like "appkit.string" when NSString resolves to "string"
		if isGoPrimitive(typeName) {
			return typeName
		}

		// Check for framework hierarchy violations BEFORE adding the package qualification
		// If the target framework is at a higher level than the current framework, return
		// objectivec.IObject instead to avoid import cycles (fixes appledocs-496)
		currentLevel := getFrameworkLevel(strings.ToLower(framework))
		targetLevel := getFrameworkLevel(frameworkPkg)
		if currentLevel >= 0 && targetLevel > currentLevel {
			return "objectivec.IObject"
		}

		Debug.TypeMap("cross-framework registry hit", typeName, frameworkPkg,
			"typeName", typeName,
			"framework", framework,
			"targetFramework", frameworkPkg)
		return frameworkPkg + "." + typeName
	}

	// Before falling back to unsafe.Pointer, check if this type belongs to the current framework
	// based on naming conventions. For example, in AppKit, types like NSView, NSButton, NSTextCheckingResult
	// should be returned as-is, not qualified with appkit.
	// This handles types that aren't classes (so not in currentFrameworkClasses) but are still
	// defined in the current framework's types.gen.go file.
	if framework != "" {
		// Check common framework prefixes
		frameworkPrefixes := map[string][]string{
			"AppKit":           {"NS", "AK"},
			"Foundation":       {"NS", "CF"},
			"CoreGraphics":     {"CG"},
			"QuartzCore":       {"CA"},
			"CoreImage":        {"CI"},
			"CoreData":         {"NS", "CD"},
			"AVFoundation":     {"AV"},
			"Metal":            {"MTL"},
			"MetalKit":         {"MTK"},
			"SpriteKit":        {"SK"},
			"SceneKit":         {"SCN"},
			"CoreML":           {"ML"},
			"Vision":           {"VN"},
			"CoreLocation":     {"CL"},
			"MapKit":           {"MK"},
			"PhotoKit":         {"PH"},
			"Photos":           {"PH"},
			"ScreenCaptureKit": {"SC"},
		}

		if prefixes, ok := frameworkPrefixes[framework]; ok {
			for _, prefix := range prefixes {
				if strings.HasPrefix(typeName, prefix) {
					// Type likely belongs to current framework, return as-is
					return typeName
				}
			}
		}
	}

	// Last resort fallbacks for common type patterns before unsafe.Pointer
	// These handle types that exist in docs but aren't extracted yet
	if strings.HasPrefix(typeName, "NS") || strings.HasPrefix(typeName, "CG") {
		// NS_OPTIONS types end with "Options" - map to uint
		if strings.HasSuffix(typeName, "Options") {
			return "uint"
		}
		// String constant types end with "Key", "Domain", or "Kind" - map to string
		if strings.HasSuffix(typeName, "Key") || strings.HasSuffix(typeName, "Domain") || strings.HasSuffix(typeName, "Kind") {
			return "string"
		}
		// ID suffix types - map to uint
		if strings.HasSuffix(typeName, "ID") {
			return "uint"
		}
		// Result suffix types - map to int (for enums like NSComparisonResult)
		if strings.HasSuffix(typeName, "Result") {
			return "int"
		}
	}

	// Truly unknown types fall back to unsafe.Pointer
	return "unsafe.Pointer"
}

// parameterToGoType converts an Objective-C parameter to a Go type.
// Uses the occ2go.MapCTypeToGo function with framework context.
func parameterToGoType(param occ2go.Parameter, framework string) string {
	paramType := strings.TrimSpace(param.Type)
	return occ2go.MapCTypeToGo(paramType, framework)
}

// wrapObjCReturn generates the return statement for converting objc.ID to Go types.
// It handles special cases like bool conversion and objc.Object mapping.
// Examples:
//
//	wrapObjCReturn("bool") -> "ret != 0"
//	wrapObjCReturn("objc.Object") -> "objc.ID(ret)"
//	wrapObjCReturn("int") -> "int(ret)"
func wrapObjCReturn(goType string) string {
	switch goType {
	case "bool":
		return "ret != 0"
	case "objc.ID":
		return "ret"
	case "unsafe.Pointer":
		return "unsafe.Pointer(ret)"
	default:
		// Default cast
		return fmt.Sprintf("%s(ret)", goType)
	}
}

// typeToStructName extracts the unqualified struct name from a Go type.
// This is used for objc.Send[T] calls which need the local struct name.
// Examples:
//
//	foundation.Data -> Data
//	NSData -> Data
//	Data -> Data
//	string -> string
//	objc.ID -> objc.ID
func typeToStructName(goType string) string {
	// Handle empty or basic types
	if goType == "" || goType == "string" || goType == "bool" || goType == "int" ||
		goType == "uint" || goType == "float32" || goType == "float64" ||
		strings.HasPrefix(goType, "objc.") || strings.HasPrefix(goType, "unsafe.") {
		return goType
	}

	// Strip framework prefix (e.g., "foundation.Data" -> "Data")
	if strings.Contains(goType, ".") {
		parts := strings.Split(goType, ".")
		goType = parts[len(parts)-1]
	}

	// Strip NS/CG/CA prefixes to get struct name
	return classToStructName(goType)
}
