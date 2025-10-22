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
		// TODO: Add type resolution methods as they're converted

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
