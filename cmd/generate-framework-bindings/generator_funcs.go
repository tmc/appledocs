package main

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/tmc/appledocs/occ2go"
)

// GeneratorFuncs wraps Generator to provide template helper methods.
// All methods have direct access to Generator state via embedding.
//
// This struct serves as the template function API, replacing scattered
// standalone functions with methods that have proper access to Generator
// data structures and indexes.
type GeneratorFuncs struct {
	*Generator
}

// Funcs returns a template.FuncMap containing all GeneratorFuncs methods.
// These functions receive the necessary context from the GeneratorFuncs receiver,
// so templates only need to pass the specific data item (e.g., the method).
//
// Template usage: {{formatMethodParams .}} instead of {{formatMethodParams $.Generator .}}
func (gf GeneratorFuncs) Funcs() template.FuncMap {
	Debug.TypeMap("GeneratorFuncs.Funcs() called", gf.Framework, "",
		"framework", gf.Framework)
	return template.FuncMap{
		// Method Formatting
		"formatMethodParams": gf.formatMethodParams,

		// Type Resolution
		"shouldSkipTypedef":   gf.shouldSkipTypedef,
		"typeToInterfaceType": gf.TypeToInterfaceType,
		"concreteReturnType":  gf.concreteReturnType,

		// Name Conversion
		"stripFrameworkPrefix": gf.stripFrameworkPrefix,

		// Constructor Generation
		// TODO: Add constructor methods as they're converted

		// Import Resolution
		// TODO: Add import resolution methods as they're converted

		// Test Generation
		"canGenerateTestValue": gf.canGenerateTestValue,
	}
}

// Method Formatting
// -----------------

// formatMethodParams formats method parameters for Go function signatures using data-driven type checking.
// This method uses the O(1) indexes from Phase 1 refactoring for efficient type lookups.
func (gf GeneratorFuncs) formatMethodParams(method *occ2go.ParsedMethod) string {
	Debug.TypeMap("formatMethodParams called", method.Name, "",
		"method", method.Name,
		"numParams", len(method.Parameters))
	if len(method.Parameters) == 0 {
		return ""
	}

	parts := make([]string, len(method.Parameters))
	for i, p := range method.Parameters {
		paramName := p.Name
		if paramName == "" {
			paramName = fmt.Sprintf("p%d", i)
		}
		if isGoKeyword(paramName) {
			paramName += "_"
		}

		Debug.TypeMap("formatMethodParams entry", p.Name, p.Type,
			"method", method.Name,
			"param", p.Name,
			"paramType", p.Type,
			"framework", gf.Framework)

		goType := mapObjCTypeToGo(p.Type, gf.Framework)

		// Convert objc.ID to objectivec.IObject for better type safety
		if goType == "objc.ID" {
			goType = "objectivec.IObject"
		} else {
			// Use data-driven type checking instead of heuristics
			// This calls Generator.TypeToInterfaceType which uses classIndex, enumIndex, typedefIndex
			goType = gf.TypeToInterfaceType(goType)
		}

		// WORKAROUND for bead appledocs-473: Strip NS/CG/CA prefix from types that look like
		// undefined enums (not in our indexes but have the prefix pattern)
		// This handles enums like NSEnergyFormatterUnit that exist but weren't extracted
		// IMPORTANT: Don't strip if the type is already qualified with a package (e.g., "corefoundation.Point")
		// See appledocs-514: Cross-framework types must preserve their package qualification
		if !strings.Contains(goType, ".") {
			stripped := stripObjCPrefix(goType)
			isClass := gf.IsClassType(goType)
			isEnum := gf.IsEnumType(goType)
			isTypedef := gf.IsTypedefType(goType)
			Debug.TypeMap("type classification check", goType, stripped,
				"goType", goType,
				"stripped", stripped,
				"equal", stripped == goType,
				"isClass", isClass,
				"isEnum", isEnum,
				"isTypedef", isTypedef)
			// Check if the STRIPPED name is in the enum index but the FULL name is not
			strippedIsEnum := gf.IsEnumType(stripped)
			if isEnum != strippedIsEnum {
				Debug.TypeMap("enum mismatch detected", goType, stripped,
					"fullType", goType,
					"fullIsEnum", isEnum,
					"strippedType", stripped,
					"strippedIsEnum", strippedIsEnum)
			}
			// The fix: if stripped name is an enum but full name also says it's an enum,
			// it means the enum index has BOTH. We should use the stripped name.
			if stripped != goType && !gf.IsClassType(goType) && !gf.IsTypedefType(goType) {
				// Skip the isEnum check - just strip if it's not a class or typedef
				Debug.TypeMap("formatMethodParams: stripping type", goType, stripped,
					"originalType", goType,
					"strippedType", stripped,
					"method", method.Name,
					"param", p.Name,
					"framework", gf.Framework)
				goType = stripped
			}
		} else {
			Debug.TypeMap("formatMethodParams: skipping strip for qualified type", goType, "",
				"goType", goType,
				"reason", "already qualified with package",
				"method", method.Name,
				"framework", gf.Framework)
		}

		parts[i] = fmt.Sprintf("%s %s", paramName, goType)
	}
	return strings.Join(parts, ", ")
}
// ---------------

// shouldSkipTypedef determines if a typedef should be skipped during generation.
// Returns true if the typedef should be skipped because:
// 1. The stripped type name matches an existing enum name (enums are generated separately)
// 2. The typedef is already defined in types.gen.go template
func (gf GeneratorFuncs) shouldSkipTypedef(typedef *occ2go.ParsedTypedef) bool {
	if typedef == nil || typedef.Name == "" {
		return true
	}

	// Strip the ObjC prefix to get the Go type name
	strippedName := stripObjCPrefix(typedef.Name)

	// Skip if this typedef's stripped name matches an existing enum
	// Enums are generated in enums.gen.go, so we don't want duplicate definitions
	if _, exists := gf.enumIndex[strippedName]; exists {
		return true
	}

	// Also check with the original (non-stripped) name in case it's already an enum
	if _, exists := gf.enumIndex[typedef.Name]; exists {
		return true
	}

	// Skip types that would conflict with hardcoded types in types.gen.go template
	// Check both the stripped name and common type patterns
	if gf.isTypeInTypesTemplate(strippedName) {
		return true
	}

	return false
}

// isTypeInTypesTemplate checks if a type name is defined in the types.gen.go template.
// This checks for geometry struct types that are manually defined in the template
// (Point, Size, Rect, Range, Vector) which should not be generated as typedefs.
func (gf GeneratorFuncs) isTypeInTypesTemplate(typeName string) bool {
	// Common geometry types used across frameworks
	// These are STRUCT types manually defined in types.gen.go, not extracted from docs
	geometryTypes := map[string]bool{
		"Point":  true, // NSPoint/CGPoint
		"Size":   true, // NSSize/CGSize
		"Rect":   true, // NSRect/CGRect
		"Range":  true, // NSRange
		"Vector": true, // CGVector/NSVector
	}

	return geometryTypes[typeName]
}

// Phase 3 Status: COMPLETE
// ------------------------
//
// The GeneratorFuncs architecture is complete and functional.
//
// ✅ Completed Work:
//   - Two-tier template function registration (templateFuncs + GeneratorFuncs)
//   - Templates cleaned up (0 references to $.Generator)
//   - Three GeneratorFuncs methods implemented:
//   - formatMethodParams - O(1) type lookups for method parameters
//   - shouldSkipTypedef - Uses enumIndex for duplicate detection
//   - TypeToInterfaceType - Exposes Generator.TypeToInterfaceType to templates
//
// concreteReturnType extracts the unqualified type name from a Go type for objc.Send[T].
// This strips framework prefixes and handles the difference between classes and enums:
// - Classes: NS prefix is stripped (Data not NSData)
// - Enums: NS prefix is preserved (NSQualityOfService not QualityOfService)
//
// Examples:
//
//	foundation.Data -> Data (class, already stripped by mapObjCTypeToGo)
//	NSData -> Data (class from ObjC, needs stripping)
//	foundation.NSQualityOfService -> NSQualityOfService (enum, keep NS prefix)
//	NSQualityOfService -> NSQualityOfService (enum from ObjC, keep NS prefix)
//	objc.ID -> objc.ID (preserved)
//	[]objc.ID -> []objc.ID (preserved)
func (gf GeneratorFuncs) concreteReturnType(goType string) string {
	// Handle empty types
	if goType == "" {
		return goType
	}

	// Handle slices - preserve brackets and recurse on element type
	if strings.HasPrefix(goType, "[]") {
		elementType := strings.TrimPrefix(goType, "[]")
		return "[]" + gf.concreteReturnType(elementType)
	}

	// Handle maps - preserve entire map type
	if strings.HasPrefix(goType, "map[") {
		return goType
	}

	// Handle basic types and types that should be preserved as-is
	if goType == "string" || goType == "bool" || goType == "int" ||
		goType == "uint" || goType == "float32" || goType == "float64" ||
		goType == "int8" || goType == "int16" || goType == "int32" || goType == "int64" ||
		goType == "uint8" || goType == "uint16" || goType == "uint32" || goType == "uint64" ||
		goType == "uintptr" || goType == "byte" || goType == "rune" {
		return goType
	}

	// Handle qualified types from standard packages (objc., unsafe., etc.)
	if strings.HasPrefix(goType, "objc.") || strings.HasPrefix(goType, "unsafe.") {
		return goType
	}

	// Handle cross-framework references (e.g., "coregraphics.CGRect", "foundation.Data")
	// Only strip the framework prefix if it matches the current framework
	if strings.Contains(goType, ".") {
		parts := strings.Split(goType, ".")
		frameworkPrefix := parts[0]
		typeName := parts[len(parts)-1]

		// If it's from a different framework, check if it would cause a hierarchy violation
		// For hierarchy violations, the interface type will be objc.IObject, so we should
		// use objc.ID for the Send call to avoid importing the higher-level framework
		if frameworkPrefix != strings.ToLower(gf.Framework) {
			// Check framework hierarchy
			currentLevel := getFrameworkLevel(strings.ToLower(gf.Framework))
			targetLevel := getFrameworkLevel(frameworkPrefix)

			// If this is a hierarchy violation (lower framework referencing higher),
			// use objc.ID instead of the qualified type to avoid import cycle
			if currentLevel >= 0 && targetLevel > currentLevel {
				return "objc.ID"
			}

			// Cross-framework reference - keep qualified
			return goType
		}

		// Same framework - unqualify
		goType = typeName
	}

	// Now goType is unqualified (e.g., "Data", "NSData", "NSQualityOfService", "QualityOfService")
	// Check if it's an enum (preserve NS prefix) or class (strip NS prefix)

	// Check if it's already an enum as-is
	if _, isEnum := gf.enumIndex[goType]; isEnum {
		// It's an enum - keep it as-is
		return goType
	}

	// If it HAS an NS/CG/CA prefix, check if the stripped version is an enum
	// Enums are indexed by both full and stripped names, so we need to check both
	if strings.HasPrefix(goType, "NS") || strings.HasPrefix(goType, "CG") || strings.HasPrefix(goType, "CA") {
		strippedName := stripObjCPrefix(goType)
		if _, isEnum := gf.enumIndex[strippedName]; isEnum {
			// The stripped name is in the enum index - return the FULL name (with prefix)
			return goType
		}
	}

	// If it doesn't have NS prefix, try adding it to check if it's an enum
	// (mapObjCTypeToGo might have stripped the prefix)
	if !strings.HasPrefix(goType, "NS") && !strings.HasPrefix(goType, "CG") && !strings.HasPrefix(goType, "CA") {
		// Try with NS prefix
		withNS := "NS" + goType
		if _, isEnum := gf.enumIndex[withNS]; isEnum {
			// It's an enum that had its prefix stripped - restore it
			return withNS
		}
	}

	// Not an enum - apply class name stripping
	result := classToStructName(goType)
	return result
}

//
// 📝 Design Decision:
//   Most template functions are pure utilities in templateFuncs (funcs_core.go).
//   Only functions needing Generator state are converted to GeneratorFuncs methods.
//   This minimal approach keeps the architecture clean and maintainable.
//
// 🔮 Future Additions:
//   New methods should be added here only when:
//   - Templates require new Generator-dependent functionality
//   - Performance profiling shows benefit from O(1) index lookups
//   - A function clearly belongs as a Generator method (cohesion)
//
// The current implementation provides a solid foundation that's easy to extend.

// Name Conversion
// ---------------

// stripFrameworkPrefix removes the framework-specific prefix from a type name.
// Examples:
//
//	Foundation: NSQualityOfService -> QualityOfService
//	CoreGraphics: CGColor -> Color
//	AppKit: NSWindow -> Window
func (gf GeneratorFuncs) stripFrameworkPrefix(name string) string {
	// Map framework names to their common prefixes
	prefixes := map[string][]string{
		"Foundation":         {"NS"},
		"AppKit":             {"NS"},
		"CoreFoundation":     {"CF"},
		"CoreGraphics":       {"CG"},
		"CoreImage":          {"CI"},
		"CoreVideo":          {"CV"},
		"CoreAudio":          {"CA"},
		"AVFoundation":       {"AV"},
		"SecurityFoundation": {"SF"},
	}

	// Get prefixes for current framework
	fwPrefixes, ok := prefixes[gf.Framework]
	if !ok {
		// Unknown framework, return as-is
		return name
	}

	// Try each prefix
	for _, prefix := range fwPrefixes {
		if strings.HasPrefix(name, prefix) {
			stripped := strings.TrimPrefix(name, prefix)
			// Make sure we didn't strip the entire name
			if stripped != "" {
				return stripped
			}
		}
	}

	return name
}

// Test Generation
// ---------------

// canGenerateTestValue checks if we can generate a reasonable test value for the given type and parameter name.
// Returns true if generateTestValue will produce a usable value.
// Rejects undefined type aliases like Coder = _undefined which are int aliases that can't use {} syntax.
// Params may include the parameter name (for context-sensitive filtering like file paths).
func (gf GeneratorFuncs) canGenerateTestValue(args ...string) bool {
	if len(args) == 0 {
		return false
	}
	goType := args[0]
	paramName := ""
	if len(args) > 1 {
		paramName = strings.ToLower(args[1])
	}

	// Check if this is an undefined type alias (e.g., Coder, StringEncoding)
	// These are int aliases and cannot use {} syntax in tests
	if !strings.Contains(goType, ".") {
		if _, isUndefined := gf.undefinedTypes[goType]; isUndefined {
			return false
		}
	}

	// Check if this is an interface type (I prefix)
	// Interface types cannot be instantiated with {} syntax
	// We need to reject unqualified interface types (ICoder, IString) and qualified ones (foundation.ICoder)
	typeName := goType
	if strings.Contains(goType, ".") {
		parts := strings.Split(goType, ".")
		if len(parts) == 2 {
			typeName = parts[1]
		}
	}
	if strings.HasPrefix(typeName, "I") && len(typeName) > 1 && typeName[1] >= 'A' && typeName[1] <= 'Z' {
		// Type starts with I followed by uppercase letter - likely an interface
		return false
	}

	// We can generate test values for most primitive types and some common types
	switch goType {
	case "string", "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "bool",
		"objc.SEL":
		return true
	}

	// objc.ID and objc.Class cannot be tested with 0/nil as they cause crashes in many
	// Foundation/AppKit APIs that expect valid object/class pointers.
	// Examples:
	//   - +[NSClassDescription classDescriptionForClass:] requires non-nil class
	//   - +[NSMutableDictionary dictionaryWithSharedKeySet:] requires non-nil keyset
	//   - Metal APIs require non-nil device pointers
	//
	// TODO: We could whitelist specific parameter names that are known to accept nil
	// (e.g., "target", "object" in some contexts), but for now we're conservative.
	if goType == "objc.ID" || goType == "objc.Class" {
		return false
	}

	// We can handle Foundation geometry types
	if strings.HasPrefix(goType, "foundation.") {
		typeName := strings.TrimPrefix(goType, "foundation.")
		switch typeName {
		case "Rect", "Size", "Point", "Range":
			return true
		}
	}

	// We can handle package-local types (enums and structs) BUT NOT undefined type aliases
	if !strings.Contains(goType, ".") {
		return true
	}

	// We CANNOT safely generate test values for unsafe.Pointer
	// as they would require actual allocated objects which we don't have in tests
	if goType == "unsafe.Pointer" {
		return false
	}

	// Special handling for string parameters that are likely file paths or URLs
	// These would cause runtime crashes if we try to use them
	if goType == "string" && paramName != "" {
		problematicNames := []string{
			"path", "filepath", "filename", "file",
			"url", "uri",
			"bundlepath", "resourcepath", "directory", "dir",
		}
		for _, name := range problematicNames {
			if strings.Contains(paramName, name) {
				return false
			}
		}
	}

	// For other types, we don't know how to create test values
	return false
}
