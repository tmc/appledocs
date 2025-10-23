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
	return template.FuncMap{
		// Method Formatting
		"formatMethodParams": gf.formatMethodParams,

		// Type Resolution
		"shouldSkipTypedef": gf.shouldSkipTypedef,

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

	// TimeInterval is defined for Foundation and ObjectiveC
	if (gf.Framework == "Foundation" || gf.Framework == "ObjectiveC") && typeName == "TimeInterval" {
		return true
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

// TODO: Add more GeneratorFuncs methods here as they're converted from standalone functions

// Name Conversion
// ---------------
// TODO: Add name conversion methods

// Constructor Generation
// ----------------------
// TODO: Add constructor methods

// Import Resolution
// -----------------
// TODO: Add import resolution methods
