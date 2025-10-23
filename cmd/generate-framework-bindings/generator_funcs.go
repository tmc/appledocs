package main

import (
	"fmt"
	"os"
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
	return template.FuncMap{
		// Method Formatting
		"formatMethodParams": gf.formatMethodParams,

		// Type Resolution
		"shouldSkipTypedef":   gf.shouldSkipTypedef,
		"typeToInterfaceType": gf.TypeToInterfaceType,
		"concreteReturnType":  gf.concreteReturnType,

		// Name Conversion
		// TODO: Add name conversion methods as they're converted

		// Constructor Generation
		// TODO: Add constructor methods as they're converted

		// Import Resolution
		// TODO: Add import resolution methods as they're converted
	}
}

// Method Formatting
// -----------------

// formatMethodParams formats method parameters for Go function signatures using data-driven type checking.
// This method uses the O(1) indexes from Phase 1 refactoring for efficient type lookups.
func (gf GeneratorFuncs) formatMethodParams(method *occ2go.ParsedMethod) string {
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

		goType := mapObjCTypeToGo(p.Type, gf.Framework)

		// Convert objc.ID to objectivec.IObject for better type safety
		if goType == "objc.ID" {
			goType = "objectivec.IObject"
		} else {
			// Use data-driven type checking instead of heuristics
			// This calls Generator.TypeToInterfaceType which uses classIndex, enumIndex, typedefIndex
			goType = gf.TypeToInterfaceType(goType)
		}

		parts[i] = fmt.Sprintf("%s %s", paramName, goType)
	}
	return strings.Join(parts, ", ")
}

// Type Resolution
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
// This avoids hardcoding a list by using pattern matching and framework-specific rules.
func (gf GeneratorFuncs) isTypeInTypesTemplate(typeName string) bool {
	// Common geometry types used across frameworks
	geometryTypes := map[string]bool{
		"Point":  true, // NSPoint/CGPoint
		"Size":   true, // NSSize/CGSize
		"Rect":   true, // NSRect/CGRect
		"Range":  true, // NSRange
		"Vector": true, // CGVector/NSVector
	}

	// Geometry types are defined in types.gen.go for several frameworks
	if geometryTypes[typeName] {
		return true
	}

	// RectEdge is a special enum defined inline in types.gen.go for Foundation
	if gf.Framework == "Foundation" && typeName == "RectEdge" {
		return true
	}

	return false
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

		// Debug
		if strings.Contains(goType, "Data") || strings.Contains(goType, "Quality") {
			fmt.Fprintf(os.Stderr, "DEBUG framework check: frameworkPrefix=%q gf.Framework=%q strings.ToLower(gf.Framework)=%q\n",
				frameworkPrefix, gf.Framework, strings.ToLower(gf.Framework))
		}

		// If it's from a different framework, keep it fully qualified
		// For example, in Foundation: "coregraphics.CGRect" stays as "coregraphics.CGRect"
		if frameworkPrefix != strings.ToLower(gf.Framework) {
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
		if strings.Contains(goType, "Quality") {
			fmt.Fprintf(os.Stderr, "DEBUG: concreteReturnType found %q in enumIndex directly\n", goType)
		}
		return goType
	}

	// If it HAS an NS/CG/CA prefix, check if the stripped version is an enum
	// Enums are indexed by both full and stripped names, so we need to check both
	if strings.HasPrefix(goType, "NS") || strings.HasPrefix(goType, "CG") || strings.HasPrefix(goType, "CA") {
		strippedName := stripObjCPrefix(goType)
		if _, isEnum := gf.enumIndex[strippedName]; isEnum {
			// The stripped name is in the enum index - return the FULL name (with prefix)
			if strings.Contains(goType, "Quality") {
				fmt.Fprintf(os.Stderr, "DEBUG: concreteReturnType found stripped %q in enumIndex for %q\n", strippedName, goType)
			}
			return goType
		}
		if strings.Contains(goType, "Quality") {
			fmt.Fprintf(os.Stderr, "DEBUG: concreteReturnType did NOT find %q or %q in enumIndex (size=%d)\n", goType, strippedName, len(gf.enumIndex))
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
