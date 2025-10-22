package main

import (
	"strings"
	"text/template"
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

// crossFrameworkTypeRegistry maps type names to their framework package names
// This allows proper type resolution across frameworks instead of falling back to unsafe.Pointer
// Format: map[typeName]frameworkPackage (e.g., "Window" -> "appkit", "String" -> "foundation")
var crossFrameworkTypeRegistry = make(map[string]string)

// templateFuncs is the FuncMap available to all templates
// This is the central registration point for all template helper functions
var templateFuncs = template.FuncMap{
	// String utilities
	"join":        joinStrings,
	"lower":       lowerString,
	"trimspace":   trimSpaceString,
	"trimRight":   trimRightString,
	"trimPrefix":  trimPrefixString,
	"hasPrefix":   hasPrefixString,
	"commentLine": commentLine,
	"dict":        dict,

	// occ2go type mapping (wrapped to apply framework-specific mappings)
	"mapCTypeToGo": mapCTypeToGoWithFramework,

	// Parameter processing helpers
	"isGoKeyword":   isGoKeyword,
	"prepareParams": prepareParams,

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
	"classFileName":                       classFileName,
	"classTestFileName":                   classTestFileName,
	"protocolFileName":                    protocolFileName,
	"receiverName":                        receiverName,
	"selectorToGoName":                    selectorToGoName,
	"mapObjCTypeToGo":                     mapObjCTypeToGo,
	// formatMethodParams is now provided by GeneratorFuncs.Funcs()
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
	"sortMethodsByName":                   sortMethodsByName,
	"generateTestValue":                   generateTestValue,
	"generateTestValueWithPackage":        generateTestValueWithPackage,
	"canGenerateTestValue":                canGenerateTestValue,
	"wrapObjCReturn":                      wrapObjCReturn,
	"isEssentialSelector":                 isEssentialSelector,
	"convertDocURL":                       convertDocURL,

	// Property generation helpers
	"propertyToGoName":            propertyToGoName,
	"contains":                    sliceContainsString,
	"capitalize":                  capitalizeFirst,
	"propertyConflictsWithParent": propertyConflictsWithParent,
	"typeToInterfaceType":         typeToInterfaceType,

	// Import merging
	"mergeImports": mergeImports,

	// Type resolution
	"resolveType": resolveType,

	// Method filtering
	"isInheritedFromNSObject": isInheritedFromNSObject,

	// Cross-framework dependency detection
	"classDependsOnCoreGraphics": classDependsOnCoreGraphics,

	// Method name disambiguation
	"methodGoName": methodGoName,

	// Class-level helpers
	"getClassImports":        getClassImports,
	"getSortedClassImports":  getSortedClassImports,
	"getInterfaceParent":     getInterfaceParent,
	"getStructEmbeddedField": getStructEmbeddedField,
	"getFromConstructorBody": getFromConstructorBody,
	"getConstructorBody":     getConstructorBody,

	// Utility functions for template generation
	"sortedKeys":        sortedKeys,
	"stripObjCPrefix":   stripObjCPrefix,
	"cleanConstantName": cleanConstantName,
	"strContains":       stringsContains,
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
