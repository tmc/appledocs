package main

import (
	"strings"
	"text/template"

	"github.com/tmc/appledocs/occ2go"
)

// currentFrameworkClasses holds the set of class names (after prefix stripping) defined in the current framework
// This is used to detect cross-framework type references
var currentFrameworkClasses = make(map[string]bool)

// currentFrameworkEnums holds the set of enum names (after prefix stripping) defined in the current framework
// This is used to detect cross-framework type references
var currentFrameworkEnums = make(map[string]bool)

// currentFrameworkTypedefs holds the set of typedef names (after prefix stripping) defined in the current framework
// This is used to detect cross-framework type references
var currentFrameworkTypedefs = make(map[string]bool)

// currentFrameworkStructs holds the set of struct names (after prefix stripping) defined in the current framework
// This is used to detect cross-framework type references
var currentFrameworkStructs = make(map[string]bool)

// crossFrameworkTypeRegistry maps type names to their framework package names
// This allows proper type resolution across frameworks instead of falling back to unsafe.Pointer
// Format: map[typeName]frameworkPackage (e.g., "Window" -> "appkit", "String" -> "foundation")
var crossFrameworkTypeRegistry = make(map[string]string)

// templateFuncs is the FuncMap available to all templates
// This is the central registration point for all template helper functions
var templateFuncs = template.FuncMap{
	// String utilities
	"join":                joinStrings,
	"lower":               lowerString,
	"title":               titleString,
	"trimspace":           trimSpaceString,
	"trimRight":           trimRightString,
	"trimPrefix":          trimPrefixString,
	"hasPrefix":           hasPrefixString,
	"commentLine":         commentLine,
	"dict":                dict,
	"isValidGoIdentifier": isValidGoIdentifier,

	// occ2go type mapping (wrapped to apply framework-specific mappings)
	"mapCTypeToGo": mapCTypeToGoWithFramework,

	// Parameter processing helpers
	"isGoKeyword":         isGoKeyword,
	"prepareParams":       prepareParams,
	"getRelaxedParamInfo": GetRelaxedParamInfo, // Access relaxed parameter metadata

	// Function data preparation
	"prepareFunctionData":    prepareFunctionData,
	"prepareFunctionDocData": prepareFunctionDocData,

	// Method generation
	"groupFunctionsByType": groupFunctionsByType,

	// Darwinkit-style class name conversions
	"classToInterfaceName":    classToInterfaceName,
	"classToStructName":       classToStructName,
	"classToVarName":          classToVarName,
	"methodToGoName":          methodToGoName,
	"objcSelectorFromMethod":  objcSelectorFromMethod,
	"parameterToGoType":       parameterToGoType,
	"generateConstructorName": generateConstructorName,
	"isPropertyGetter":        isPropertyGetter,
	"isPropertySetter":        isPropertySetter,

	// DarwinKit class generation helpers
	"classFileName":     classFileName,
	"classTestFileName": classTestFileName,
	"protocolFileName":  protocolFileName,
	"receiverName":      receiverName,
	"selectorToGoName":  selectorToGoName,
	"mapObjCTypeToGo":   mapObjCTypeToGo,
	// formatMethodParams is provided by GeneratorFuncs.Funcs() at runtime,
	// but we need a stub here for template parsing in init()
	"formatMethodParams":                  formatMethodParamsStub,
	"formatMethodParamNames":              formatMethodParamNames,
	"formatMethodParamNamesWithFramework": formatMethodParamNamesWithFramework,
	"isConstructor":                       isConstructor,
	"stripNSPrefix":                       stripNSPrefix,
	"needsFoundationImport":               needsFoundationImport,
	"needsQuartzCoreImport":               needsQuartzCoreImport,
	"needsCustomImports":                  needsCustomImports,
	"getRequiredImports":                  getRequiredImports,
	"getFunctionRequiredImports":          getFunctionRequiredImports,
	"sortedImportPaths":                   sortedImportPaths,
	"prepareClassMethods":                 prepareClassMethods,
	"prepareInstanceMethods":              prepareInstanceMethods,
	"filterPropertyMethods":               filterPropertyMethods,
	"prepareInitMethods":                  prepareInitMethods,
	"initMethodToConstructorName":         initMethodToConstructorName,
	"prepareInitMethodsWithClassName":     prepareInitMethodsWithClassName,
	"classHasInit":                        classHasInit,
	"shouldExcludeTestExample":            shouldExcludeTestExample,
	"shouldExcludeTestMethod":             shouldExcludeTestMethod,
	"isSafeToTestOnNSObject":              isSafeToTestOnNSObject,
	"sortMethodsByName":                   sortMethodsByName,
	"generateTestValue":                   generateTestValue,
	"generateTestValueWithPackage":        generateTestValueWithPackage,
	"canGenerateTestValue":                canGenerateTestValue,
	"wrapObjCReturn":                      wrapObjCReturn,
	"isEssentialSelector":                 isEssentialSelector,
	"convertDocURL":                       convertDocURL,
	"structsUseUnsafe":                    structsUseUnsafe,
	"structsUseObjc":                      structsUseObjc,

	// Property generation helpers
	"propertyToGoName":            propertyToGoName,
	"contains":                    sliceContainsString,
	"capitalize":                  capitalizeFirst,
	"propertyConflictsWithParent": propertyConflictsWithParent,
	// Note: typeToInterfaceType is now provided by GeneratorFuncs.TypeToInterfaceType

	// Import merging
	"mergeImports": mergeImports,

	// Type resolution
	"resolveType":           resolveType,
	"parseCFunctionPointer": parseCFunctionPointer,
	"typedefsNeedUnsafe":    typedefsNeedUnsafe,
	"typedefsNeedObjc":      typedefsNeedObjc,

	// Method filtering
	"isInheritedFromNSObject": isInheritedFromNSObject,

	// Cross-framework dependency detection
	"classDependsOnCoreGraphics": classDependsOnCoreGraphics,

	// Method name disambiguation
	"methodGoName":           methodGoName,
	"disambiguateMethodName": disambiguateMethodName,

	// Class-level helpers
	"getClassImports":        getClassImports,
	"getSortedClassImports":  getSortedClassImports,
	"getInterfaceParent":     getInterfaceParent,
	"getStructEmbeddedField": getStructEmbeddedField,
	"getFromConstructorBody": getFromConstructorBody,
	"getConstructorBody":     getConstructorBody,

	// Utility functions for template generation
	"sortedKeys":         sortedKeys,
	"stripObjCPrefix":    stripObjCPrefix,
	"cleanConstantName":  cleanConstantName,
	"strContains":        stringsContains,
	"enumUnderlyingType": enumUnderlyingType,
}

// FunctionData represents data for function template rendering.
type FunctionData struct {
	Name       string
	Comment    string
	Parameters []ParameterData
	ReturnType string
}

// ParameterData represents data for parameter template rendering.
type ParameterData struct {
	Name string
	Type string
}

// FunctionDocData represents data for function documentation template rendering.
type FunctionDocData struct {
	Name               string
	Abstract           string
	Framework          string
	Parameters         []ParameterData
	ReturnType         string
	DocURL             string
	HasAvailability    bool
	IsDeprecated       bool
	IntroducedPlatform string
	IntroducedVersion  string
	DeprecatedPlatform string
	DeprecatedVersion  string
	// Method-style API fields
	MethodName       string          // Method name without type prefix (e.g., "SetFillColor")
	MethodParameters []ParameterData // Parameters excluding the receiver
}

// goKeywords contains Go keywords and predeclared identifiers that need escaping in parameter names
var goKeywords = map[string]bool{
	// Reserved keywords
	"break": true, "case": true, "chan": true, "const": true, "continue": true,
	"default": true, "defer": true, "else": true, "fallthrough": true, "for": true,
	"func": true, "go": true, "goto": true, "if": true, "import": true,
	"interface": true, "map": true, "package": true, "range": true, "return": true,
	"select": true, "struct": true, "switch": true, "type": true, "var": true,
	// Predeclared identifiers
	"true": true, "false": true, "nil": true, "iota": true,
	// Special identifiers
	"init": true,
	// Built-in types
	"bool": true, "byte": true, "complex64": true, "complex128": true,
	"error": true, "float32": true, "float64": true,
	"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
	"rune": true, "string": true,
	"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true, "uintptr": true,
	// Built-in functions
	"append": true, "cap": true, "close": true, "complex": true, "copy": true,
	"delete": true, "imag": true, "len": true, "make": true, "new": true,
	"panic": true, "print": true, "println": true, "real": true, "recover": true,
	// Common Cocoa API conflicts
	"protocol": true,
}

// isGoKeyword checks if a string is a Go reserved keyword or predeclared identifier
func isGoKeyword(name string) bool {
	return goKeywords[name]
}

// String wrapper functions for template compatibility
func joinStrings(sep string, a []string) string {
	return strings.Join(a, sep)
}

func lowerString(s string) string {
	return strings.ToLower(s)
}

func titleString(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func trimSpaceString(s string) string {
	return strings.TrimSpace(s)
}

func trimRightString(s, cutset string) string {
	return strings.TrimRight(s, cutset)
}

func trimPrefixString(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

func hasPrefixString(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func stringsContains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// formatMethodParamsStub is a stub for template parsing in init().
// The real implementation is provided by GeneratorFuncs.formatMethodParams at runtime.
func formatMethodParamsStub(gen interface{}, method interface{}) string {
	panic("formatMethodParamsStub called - should be overridden by GeneratorFuncs.Funcs()")
}

// isValidGoIdentifier checks if a string is a valid Go identifier.
// Valid Go identifiers must:
// - Start with a letter (a-z, A-Z) or underscore
// - Contain only letters, digits, or underscores
// - Not be empty
// structsUseUnsafe checks if any struct fields map to unsafe.Pointer
func structsUseUnsafe(structs []*occ2go.ParsedStruct, framework string) bool {
	for _, s := range structs {
		for _, field := range s.Fields {
			// Map the C type to Go type to check if it uses unsafe.Pointer
			mappedType := mapCTypeToGoWithFramework(field.Type, framework)
			if strings.Contains(mappedType, "unsafe.Pointer") {
				return true
			}
		}
	}
	return false
}

// structsUseObjc checks if any struct fields use objc types (objc.ID, objc.Class, objc.SEL, etc.)
func structsUseObjc(structs []*occ2go.ParsedStruct, framework string) bool {
	for _, s := range structs {
		for _, field := range s.Fields {
			// Map the C type to Go type to check if it uses objc.* types
			mappedType := mapCTypeToGoWithFramework(field.Type, framework)
			if strings.Contains(mappedType, "objc.") {
				return true
			}
		}
	}
	return false
}

func isValidGoIdentifier(s string) bool {
	if s == "" {
		return false
	}

	// First character must be letter or underscore
	first := rune(s[0])
	if !((first >= 'a' && first <= 'z') || (first >= 'A' && first <= 'Z') || first == '_') {
		return false
	}

	// Remaining characters must be letter, digit, or underscore
	for _, r := range s[1:] {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_') {
			return false
		}
	}

	return true
}

// enumUnderlyingType determines the appropriate Go type for an enum.
// It's data-driven: first checks the enum's BaseType from documentation,
// then falls back to inspecting actual enum values for negative numbers.
func enumUnderlyingType(enum *occ2go.ParsedEnum) string {
	// Data-driven approach: use the BaseType from documentation if available
	if enum.BaseType != "" {
		switch enum.BaseType {
		case "NSInteger", "NSInt", "int", "Int", "signed long", "signed int", "int32_t", "int64_t":
			return "int"
		case "NSUInteger", "NSUInt", "uint", "UInt", "unsigned long", "unsigned int", "uint32_t", "uint64_t":
			return "uint"
		}
	}

	// Fallback: inspect actual values for negative numbers
	// This handles cases where BaseType is not set or is ambiguous
	for _, enumCase := range enum.Cases {
		if enumCase.IntValue < 0 {
			return "int"
		}
	}

	// Default to uint for non-negative enums
	return "uint"
}
