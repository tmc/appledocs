package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/tmc/appledocs/occ2go"
)

// currentFrameworkClasses holds the set of class names (after prefix stripping) defined in the current framework
// This is used to detect cross-framework type references
var currentFrameworkClasses = make(map[string]bool)

// crossFrameworkTypeRegistry maps type names to their framework package names
// This allows proper type resolution across frameworks instead of falling back to unsafe.Pointer
// Format: map[typeName]frameworkPackage (e.g., "Window" -> "appkit", "String" -> "foundation")
var crossFrameworkTypeRegistry = make(map[string]string)

// templateFuncs is the FuncMap available to all templates
var templateFuncs = template.FuncMap{
	// String utilities
	"join":        strings.Join,
	"lower":       strings.ToLower,
	"trimspace":   strings.TrimSpace,
	"trimRight":   strings.TrimRight,
	"trimPrefix":  strings.TrimPrefix,
	"hasPrefix":   strings.HasPrefix,
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
	"classFileName":            classFileName,
	"classTestFileName":        classTestFileName,
	"protocolFileName":         protocolFileName,
	"receiverName":             receiverName,
	"selectorToGoName":         selectorToGoName,
	"mapObjCTypeToGo":          mapObjCTypeToGo,
	"formatMethodParams":       formatMethodParams,
	"formatMethodParamNames":   formatMethodParamNames,
	"formatMethodParamNamesWithFramework": formatMethodParamNamesWithFramework,
	"isConstructor":            isConstructor,
	"stripNSPrefix":            stripNSPrefix,
	"needsFoundationImport":    needsFoundationImport,
	"needsQuartzCoreImport":    needsQuartzCoreImport,
	"needsCustomImports":       needsCustomImports,
	"getRequiredImports":       getRequiredImports,
	"getFunctionRequiredImports": getFunctionRequiredImports,
	"sortedImportPaths":        sortedImportPaths,
	"prepareClassMethods":         prepareClassMethods,
	"prepareInstanceMethods":      prepareInstanceMethods,
	"filterPropertyMethods":       filterPropertyMethods,
	"prepareInitMethods":          prepareInitMethods,
	"initMethodToConstructorName":       initMethodToConstructorName,
	"prepareInitMethodsWithClassName":   prepareInitMethodsWithClassName,
	"classHasInit":                      classHasInit,
	"shouldExcludeTestExample":          shouldExcludeTestExample,
	"shouldExcludeTestMethod":           shouldExcludeTestMethod,
	"sortMethodsByName":              sortMethodsByName,
	"generateTestValue":              generateTestValue,
	"generateTestValueWithPackage":   generateTestValueWithPackage,
	"canGenerateTestValue":           canGenerateTestValue,
	"wrapObjCReturn":                 wrapObjCReturn,
	"isEssentialSelector":         isEssentialSelector,
	"convertDocURL":               convertDocURL,

	// Property generation helpers
	"propertyToGoName":         propertyToGoName,
	"contains":                 sliceContainsString,
	"capitalize":               capitalizeFirst,

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
	"getClassImports":       getClassImports,
	"getInterfaceParent":    getInterfaceParent,
	"getStructEmbeddedField": getStructEmbeddedField,
	"getFromConstructorBody": getFromConstructorBody,
	"getConstructorBody":     getConstructorBody,

	// Utility functions for template generation
	"sortedKeys":      sortedKeys,
	"stripObjCPrefix": stripObjCPrefix,
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

var goKeywords = map[string]bool{
	"break": true, "case": true, "chan": true, "const": true, "continue": true,
	"default": true, "defer": true, "else": true, "fallthrough": true, "for": true,
	"func": true, "go": true, "goto": true, "if": true, "import": true,
	"interface": true, "map": true, "package": true, "range": true, "return": true,
	"select": true, "struct": true, "switch": true, "type": true, "var": true,
	// Special identifiers that cannot be used as type/variable names
	"init": true,
	// Built-in types that shadow if used as parameter names
	"bool": true, "byte": true, "complex64": true, "complex128": true,
	"error": true, "float32": true, "float64": true,
	"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
	"rune": true, "string": true,
	"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true, "uintptr": true,
}

// isGoKeyword checks if a string is a Go reserved keyword.
func isGoKeyword(name string) bool {
	return goKeywords[name]
}

// prepareParams is a template helper that processes function parameters.
// It returns a slice of ParameterData with proper type mapping and unique naming.
func prepareParams(params []occ2go.Parameter, framework string) []ParameterData {
	result := make([]ParameterData, 0, len(params))
	usedNames := make(map[string]int)

	for i, p := range params {
		// Clean and map parameter type
		paramType := strings.TrimRight(p.Type, ",;)")
		paramType = strings.TrimSpace(paramType)
		paramType = mapCTypeToGoWithFramework(paramType, framework)

		// Generate parameter name if missing
		paramName := p.Name
		if paramName == "" {
			paramName = fmt.Sprintf("p%d", i)
		}

		// Escape Go keywords
		if isGoKeyword(paramName) {
			paramName = paramName + "_"
		}

		// Make parameter name unique if it's already used in this function
		if count, exists := usedNames[paramName]; exists {
			paramName = fmt.Sprintf("%s%d", paramName, count)
			usedNames[p.Name]++
		} else {
			usedNames[paramName] = 1
		}

		result = append(result, ParameterData{
			Name: paramName,
			Type: paramType,
		})
	}

	return result
}

// prepareFunctionDocData converts a ParsedFunction into FunctionDocData for documentation template rendering.
func prepareFunctionDocData(fn *occ2go.ParsedFunction, framework string) FunctionDocData {
	baseData := prepareFunctionData(fn, framework)

	data := FunctionDocData{
		Name:       fn.Name,
		Abstract:   fn.Abstract,
		Framework:  framework,
		Parameters: baseData.Parameters,
		ReturnType: baseData.ReturnType,
		DocURL:     fn.DocURL,
	}

	// Extract availability info for macOS only
	if !fn.Availability.IsEmpty() {
		for _, platform := range fn.Availability.Platforms() {
			if platform == "macOS" {
				data.HasAvailability = true
				data.IntroducedPlatform = platform
				data.IntroducedVersion = fn.Availability.IntroducedAt[platform]

				if deprecatedAt, ok := fn.Availability.DeprecatedAt[platform]; ok {
					data.IsDeprecated = true
					data.DeprecatedPlatform = platform
					data.DeprecatedVersion = deprecatedAt
				}
				break
			}
		}
	}

	// Compute method-style API fields
	typeName := extractTypeName(fn.Name, framework)
	if typeName != "" {
		data.MethodName = extractMethodName(fn.Name, framework, typeName)
		// Method parameters exclude the first parameter (receiver)
		if len(baseData.Parameters) > 0 {
			data.MethodParameters = baseData.Parameters[1:]
		}
	}

	return data
}

// extractMethodName extracts the method name by removing the type prefix.
// Examples:
//
//	CGContextSetFillColor → SetFillColor
//	CGPathAddRect → AddRect
func extractMethodName(funcName, framework, typeName string) string {
	var prefix string
	switch framework {
	case "CoreGraphics":
		prefix = "CG"
	case "CoreFoundation":
		prefix = "CF"
	case "CoreAudio":
		prefix = "CA"
	default:
		return funcName
	}

	// Remove prefix + type name
	fullPrefix := prefix + typeName
	if strings.HasPrefix(funcName, fullPrefix) {
		return strings.TrimPrefix(funcName, fullPrefix)
	}

	return funcName
}

// prepareFunctionData converts a ParsedFunction into FunctionData for template rendering.
func prepareFunctionData(fn *occ2go.ParsedFunction, framework string) FunctionData {
	data := FunctionData{
		Name:       fn.Name,
		Parameters: make([]ParameterData, 0, len(fn.Parameters)),
	}

	// Generate comment from availability info
	if !fn.Availability.IsEmpty() {
		for _, platform := range fn.Availability.Platforms() {
			version := fn.Availability.IntroducedAt[platform]
			status := ""
			if fn.Availability.Beta {
				status = " (Beta)"
			} else if deprecatedAt, ok := fn.Availability.DeprecatedAt[platform]; ok {
				status = fmt.Sprintf(" (Deprecated in %s)", deprecatedAt)
			}
			if platform == "macOS" {
				data.Comment = fmt.Sprintf("is available on %s %s+%s", platform, version, status)
				break
			}
		}
	}

	// Process parameters with unique naming
	usedNames := make(map[string]int)
	for i, p := range fn.Parameters {
		paramType := strings.TrimSpace(strings.TrimRight(p.Type, ",;)"))
		paramType = mapCTypeToGoWithFramework(paramType, framework)

		paramName := p.Name
		if paramName == "" {
			paramName = fmt.Sprintf("p%d", i)
		}
		if isGoKeyword(paramName) {
			paramName = paramName + "_"
		}

		if count, exists := usedNames[paramName]; exists {
			paramName = fmt.Sprintf("%s%d", paramName, count)
			usedNames[p.Name]++
		} else {
			usedNames[paramName] = 1
		}

		data.Parameters = append(data.Parameters, ParameterData{
			Name: paramName,
			Type: paramType,
		})
	}

	// Process return type
	if fn.ReturnType != "" && fn.ReturnType != "void" {
		data.ReturnType = mapCTypeToGoWithFramework(fn.ReturnType, framework)
	}

	return data
}

// groupFunctionsByType groups functions by their receiver type based on naming conventions.
// For example: CGContextXXX → Context type, CGPathXXX → Path type
func groupFunctionsByType(functions []*occ2go.ParsedFunction, framework string) map[string][]*occ2go.ParsedFunction {
	groups := make(map[string][]*occ2go.ParsedFunction)

	for _, fn := range functions {
		// Extract type prefix (e.g., CGContext from CGContextSetFillColor)
		typeName := extractTypeName(fn.Name, framework)
		if typeName != "" {
			groups[typeName] = append(groups[typeName], fn)
		}
	}

	return groups
}

// extractTypeName extracts the type name from a function name.
// Examples:
//
//	CGContextSetFillColor → Context
//	CGPathAddRect → Path
//	CGColorCreate → Color
func extractTypeName(funcName, framework string) string {
	// Handle framework-specific prefixes
	var prefix string
	switch framework {
	case "CoreGraphics":
		prefix = "CG"
	case "CoreFoundation":
		prefix = "CF"
	case "CoreAudio":
		prefix = "CA"
	default:
		return ""
	}

	// Function must start with framework prefix
	if !strings.HasPrefix(funcName, prefix) {
		return ""
	}

	// Extract the type name (first capitalized word after prefix)
	remainder := strings.TrimPrefix(funcName, prefix)

	// Find where the type name ends (next capital letter or end of string)
	typeEnd := 0
	for i, ch := range remainder {
		if i > 0 && ch >= 'A' && ch <= 'Z' {
			typeEnd = i
			break
		}
	}

	if typeEnd == 0 {
		// No capital letter found, use whole remainder
		return remainder
	}

	typeName := remainder[:typeEnd]

	// Filter out common non-type prefixes
	skipPrefixes := []string{"Get", "Set", "Create", "Make", "Copy", "Release", "Retain"}
	for _, skip := range skipPrefixes {
		if typeName == skip {
			return ""
		}
	}

	return typeName
}

// classToInterfaceName converts an Objective-C class name to a Go interface name.
// Strips the NS/CG/CF prefix and prepends 'I'.
// Examples:
//
//	NSButton -> IButton
//	NSView -> IView
//	CGContext -> IContext
func classToInterfaceName(className string) string {
	if className == "" {
		return ""
	}

	// Strip common prefixes
	name := stripObjCPrefix(className)
	return "I" + name
}

// classToStructName converts an Objective-C class name to a Go struct name.
// Strips the NS/CG/CF prefix.
// Examples:
//
//	NSButton -> Button
//	NSView -> View
//	CGContext -> Context
func classToStructName(className string) string {
	if className == "" {
		return ""
	}

	return stripObjCPrefix(className)
}

// classToVarName converts an Objective-C class name to a Go variable name for the class.
// Strips prefix and returns capitalized with 'Class' suffix to make it exported.
// Examples:
//
//	NSButton -> ButtonClass
//	NSView -> ViewClass
//	CGContext -> ContextClass
func classToVarName(className string) string {
	if className == "" {
		return ""
	}

	name := stripObjCPrefix(className)
	// Make it capitalized to be public (exported from the package)
	return strings.ToUpper(name[:1]) + name[1:] + "Class"
}

// stripObjCPrefix removes Objective-C prefixes algorithmically and invalid identifier characters from a class name
func stripObjCPrefix(className string) string {
	// First strip colons and other invalid identifier characters
	className = strings.ReplaceAll(className, ":", "")

	// Known Apple framework prefixes
	// Order matters: try longer prefixes first (e.g., "MPSNN" before "MPS")
	knownPrefixes := []string{
		// 4-letter prefixes
		"MPSNN",
		// 3-letter prefixes
		"MPS", "MTL", "MTK",
		// 2-letter prefixes (most common)
		"NS", "CG", "CF", "CA", "CI", "CL", "CM", "CV", "CT", "SC", "AV", "UI", "WK", "SK",
		"PK", "AR", "ML", "VN", "NL", "AS", "LA", "MP", "HC", "HM", "GK", "QL", "AU", "IO",
	}

	// Try each known prefix
	for _, prefix := range knownPrefixes {
		if len(className) > len(prefix) && strings.HasPrefix(className, prefix) {
			// Check if next character is uppercase (start of actual name)
			nextChar := className[len(prefix)]
			if nextChar >= 'A' && nextChar <= 'Z' {
				return className[len(prefix):]
			}
		}
	}

	return className
}

// MethodInfo represents parsed information about an Objective-C method
type MethodInfo struct {
	Selector   string
	IsInstance bool
	Parameters []occ2go.Parameter
	ReturnType string
}

// methodToGoName converts an Objective-C method selector to a Go method name.
// Examples:
//
//	setTitle: -> SetTitle
//	buttonWithTitle:image: -> ButtonWithTitleImage
//	initWithFrame: -> InitWithFrame
//	isEnabled -> IsEnabled
func methodToGoName(method MethodInfo) string {
	selector := method.Selector
	if selector == "" {
		return ""
	}

	// Remove trailing colons
	selector = strings.TrimSuffix(selector, ":")

	// Split by colon to get parts
	parts := strings.Split(selector, ":")

	// Capitalize each part
	var result strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		// Capitalize first letter
		if len(part) > 0 {
			result.WriteString(strings.ToUpper(part[:1]))
			if len(part) > 1 {
				result.WriteString(part[1:])
			}
		}
	}

	return result.String()
}

// objcSelectorFromMethod builds the Objective-C selector string from method info.
// Examples:
//
//	setTitle: -> "setTitle:"
//	buttonWithTitle:image: -> "buttonWithTitle:image:"
//	init -> "init"
func objcSelectorFromMethod(method MethodInfo) string {
	return method.Selector
}

// parameterToGoType converts an Objective-C parameter to a Go type.
// Uses the occ2go.MapCTypeToGo function with framework context.
func parameterToGoType(param occ2go.Parameter, framework string) string {
	paramType := strings.TrimSpace(param.Type)
	return occ2go.MapCTypeToGo(paramType, framework)
}

// generateConstructorName generates a Go constructor function name from an init method.
// Examples:
//
//	init -> New
//	initWithFrame: -> NewWithFrame
//	initWithTitle:image: -> NewWithTitleImage
func generateConstructorName(method MethodInfo) string {
	selector := method.Selector

	// Special case for plain "init"
	if selector == "init" {
		return "New"
	}

	// Strip "init" prefix if present
	if strings.HasPrefix(selector, "init") {
		selector = strings.TrimPrefix(selector, "init")
	}

	// Remove trailing colons
	selector = strings.TrimSuffix(selector, ":")

	// If selector starts with "With", keep it
	if strings.HasPrefix(selector, "With") {
		selector = strings.TrimPrefix(selector, "With")
		return "NewWith" + methodToGoName(MethodInfo{Selector: selector})
	}

	// Convert to Go name
	return "New" + methodToGoName(MethodInfo{Selector: selector})
}

// isPropertyGetter determines if a method is likely a property getter.
// A getter:
// - Takes no parameters
// - Returns a non-void value
// - Does not start with "init", "alloc", "new", "copy", "mutableCopy"
func isPropertyGetter(method MethodInfo) bool {
	// Must have no parameters
	if len(method.Parameters) > 0 {
		return false
	}

	// Must return something
	if method.ReturnType == "" || method.ReturnType == "void" {
		return false
	}

	// Check selector patterns that indicate it's not a getter
	selector := method.Selector
	nonGetterPrefixes := []string{"init", "alloc", "new", "copy", "mutableCopy"}
	for _, prefix := range nonGetterPrefixes {
		if strings.HasPrefix(selector, prefix) {
			return false
		}
	}

	return true
}

// isPropertySetter determines if a method is likely a property setter.
// A setter:
// - Takes exactly one parameter
// - Returns void
// - Starts with "set" followed by a capital letter
func isPropertySetter(method MethodInfo) bool {
	// Must take exactly one parameter
	if len(method.Parameters) != 1 {
		return false
	}

	// Must return void
	if method.ReturnType != "" && method.ReturnType != "void" {
		return false
	}

	// Must start with "set" followed by uppercase letter
	selector := method.Selector
	if len(selector) < 4 {
		return false
	}

	if !strings.HasPrefix(selector, "set") {
		return false
	}

	// Check that the 4th character is uppercase
	if len(selector) > 3 && selector[3] >= 'A' && selector[3] <= 'Z' {
		return true
	}

	return false
}

// classFileName converts a class name to a file name (snake_case).
// Uses the FULL class name including ObjC prefix to prevent duplicate files.
// Examples:
//   NSButton -> ns_button.gen.go
//   NSTableView -> ns_table_view.gen.go
//   ICCameraDevice -> ic_camera_device.gen.go
func classFileName(className string) string {
	// Use the full class name (e.g., ICCameraDevice -> ic_camera_device.gen.go)
	// NOT stripped prefix (CameraDevice -> camera_device.gen.go)
	// This prevents duplicate file generation - see appledocs-227
	return toSnakeCase(className) + ".gen.go"
}

// protocolFileName converts a protocol name to a file name (snake_case).
// Examples:
//   NSCopying -> copying_protocol.gen.go
//   NSTableViewDataSource -> table_view_data_source_protocol.gen.go
func protocolFileName(protocolName string) string {
	name := stripObjCPrefix(protocolName)
	return toSnakeCase(name) + "_protocol.gen.go"
}

// classTestFileName converts a class name to a test file name (snake_case).
// Examples:
//   NSButton -> button.gen_test.go
//   NSTableView -> table_view.gen_test.go
//   NSURLRequest -> url_request.gen_test.go
func classTestFileName(className string) string {
	name := stripObjCPrefix(className)
	return toSnakeCase(name) + ".gen_test.go"
}

// toSnakeCase converts CamelCase to snake_case and strips invalid filename characters
func toSnakeCase(s string) string {
	// First strip colons and other invalid filename characters
	s = strings.ReplaceAll(s, ":", "")

	var result strings.Builder
	for i, ch := range s {
		if i > 0 && ch >= 'A' && ch <= 'Z' {
			// Check if previous character was lowercase or next character is lowercase
			prevIsLower := i > 0 && s[i-1] >= 'a' && s[i-1] <= 'z'
			nextIsLower := i < len(s)-1 && s[i+1] >= 'a' && s[i+1] <= 'z'
			if prevIsLower || nextIsLower {
				result.WriteByte('_')
			}
		}
		result.WriteRune(ch)
	}
	return strings.ToLower(result.String())
}

// receiverName generates a short receiver name for methods.
// Examples:
//   Button, false -> b_
//   Button, true -> bc
func receiverName(className string, isClass bool) string {
	name := stripObjCPrefix(className)
	if len(name) == 0 {
		return "x"
	}
	short := strings.ToLower(string(name[0]))
	if isClass {
		return short + "c"
	}
	return short + "_"
}

// selectorToGoName converts an Objective-C selector to Go name.
// This wraps the occ2go.SelectorToGoName function.
func selectorToGoName(selector string) string {
	return occ2go.SelectorToGoName(selector)
}

// disambiguateMethodName generates a unique Go method name for an Objective-C method
// by appending parameter labels from the selector. This matches Swift's approach.
// Examples:
//   imageByInsertingIntermediate -> ImageByInsertingIntermediate (no params, no change)
//   imageByInsertingIntermediate: -> ImageByInsertingIntermediateWithCache (1 param named "cache")
//   setTitle:forState: -> SetTitleForState (already unique from selector parts)
func disambiguateMethodName(method *occ2go.ParsedMethod) string {
	selector := method.Selector

	// If there are no parameters, just use the standard conversion
	if len(method.Parameters) == 0 {
		return selectorToGoName(selector)
	}

	// Split selector by colons to get parameter labels
	parts := strings.Split(selector, ":")

	// If selector doesn't end with colon, the last part is not a parameter label
	if !strings.HasSuffix(selector, ":") {
		parts = parts[:len(parts)-1]
	}

	// If we have parameter labels, build the disambiguated name
	// For single-parameter methods, append "With" + capitalized parameter name
	if len(parts) == 1 && len(method.Parameters) == 1 {
		baseName := selectorToGoName(parts[0])
		// Get the parameter name from the method parameters
		paramName := method.Parameters[0].Name
		if paramName == "" {
			// Fallback: use the selector part
			paramName = parts[0]
		}
		// Capitalize parameter name
		paramName = strings.ToUpper(paramName[:1]) + paramName[1:]
		return baseName + "With" + paramName
	}

	// For multi-parameter methods, the selector already contains the labels
	// Just use the standard conversion which will include all parts
	return selectorToGoName(selector)
}

// mapCTypeToGoWithFramework wraps occ2go.MapCTypeToGo and applies framework-specific type mappings.
// This ensures C types like CGAffineTransform are properly qualified with their framework package.
func mapCTypeToGoWithFramework(cType, framework string) string {
	// First apply occ2go's basic C type mapping
	goType := occ2go.MapCTypeToGo(cType, framework)

	// Then apply our framework-specific mapping to add package qualifiers
	// For example, CGAffineTransform -> coregraphics.CGAffineTransform
	mapped := mapObjCTypeToGo(goType, framework)

	return mapped
}

// mapObjCTypeToGo maps Objective-C types to Go types for darwinkit style.
// Examples:
//   NSString * -> string
//   id -> objc.Object
//   NSButton * -> Button (interface type in parameters)
//   NSRect -> foundation.Rect
//   NSWindowStyleMask -> WindowStyleMask
func mapObjCTypeToGo(objcType, framework string) string {
	objcType = strings.TrimSpace(objcType)



	// Strip __kindof qualifier (e.g., "__kindof NSView *" -> "NSView *")
	// __kindof is an Objective-C type qualifier meaning "this type or any subclass"
	// In Go, we just use the base type
	objcType = strings.TrimPrefix(objcType, "__kindof ")

	// Handle array types with __kindof in element type (e.g., "[]__kindof AVCaptureControl" -> "[]AVCaptureControl")
	// This can occur when occ2go parser has already converted NSArray<__kindof T> to []__kindof T
	if strings.HasPrefix(objcType, "[]__kindof ") {
		objcType = "[]" + strings.TrimPrefix(objcType, "[]__kindof ")
	}

	// Handle id<Protocol> pattern (e.g., "id<NSFetchRequestResult>" -> "objc.ID")
	// This is Objective-C's protocol conformance syntax
	if strings.HasPrefix(objcType, "id<") && strings.Contains(objcType, ">") {
		return "objc.ID"
	}

	// Handle []id<Protocol> pattern (e.g., "[]id<NSFetchRequestResult>" -> "[]objc.ID")
	if strings.HasPrefix(objcType, "[]id<") && strings.Contains(objcType, ">") {
		return "[]objc.ID"
	}

	// Handle array types that are already converted by occ2go (e.g., "[]void (^)(void)" -> "[]unsafe.Pointer")
	// This handles cases where occ2go has already converted NSArray<T> to []T
	// We need to recursively map the element type
	if strings.HasPrefix(objcType, "[]") {
		elementType := strings.TrimPrefix(objcType, "[]")
		goElementType := mapObjCTypeToGo(elementType, framework)
		return "[]" + goElementType
	}

	// Handle Objective-C generic types (e.g., NSArray<NSString *>, NSArray<SCDisplay *>)
	if strings.Contains(objcType, "<") {
		// Extract NSArray element type: NSArray<ElementType *> -> []ElementType
		if strings.HasPrefix(objcType, "NSArray<") && strings.HasSuffix(objcType, ">") {
			// Extract element type between < and >
			start := strings.Index(objcType, "<") + 1
			end := strings.LastIndex(objcType, ">")
			if start > 0 && end > start {
				elementType := strings.TrimSpace(objcType[start:end])
				// Strip __kindof qualifier from element type
				elementType = strings.TrimPrefix(elementType, "__kindof ")
				// Remove trailing * from pointer types
				elementType = strings.TrimSpace(strings.TrimSuffix(elementType, "*"))

				// Strip protocol conformance syntax: NSView<NSCollectionViewElement> -> NSView
				// Objective-C uses Type<Protocol> syntax for protocol conformance, but in Go we just use the base type
				// The protocol conformance is checked at runtime by Objective-C, not at compile time
				if protocolStart := strings.Index(elementType, "<"); protocolStart > 0 {
					if strings.HasSuffix(elementType, ">") {
						elementType = strings.TrimSpace(elementType[:protocolStart])
					}
				}

				// Special case: NSString -> string
				if elementType == "NSString" {
					return "[]string"
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
					// Cross-framework reference - try resolveType before falling back to unsafe.Pointer
					resolvedType := resolveType(framework, elementType)
					if resolvedType != "unsafe.Pointer" {
						// resolveType found a valid cross-framework type (foundation.NSString, coregraphics.CGRect, etc.)
						return "[]" + resolvedType
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
		}
		// For other generic types (NSDictionary, etc.), fall back to unsafe.Pointer
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
		return goType
	}

	// Handle Objective-C blocks (e.g., void (^)(NSModalResponse))
	// Blocks are closures that cannot be easily represented in Go, so map to unsafe.Pointer
	// This comes after type registry check so explicitly mapped blocks can use proper Go types
	if strings.Contains(objcType, "^") {
		return "unsafe.Pointer"
	}

	// Handle pointers for types not in the registry
	isPointer := strings.HasSuffix(objcType, "*")
	objcTypeNoPtr := strings.TrimSpace(strings.TrimSuffix(objcType, "*"))

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
		if goType, found := lookupTypeMapping(objcTypeNoPtr, framework); found {
			return goType
		}
	}

	// Fall back to occ2go mapping
	goType := occ2go.MapCTypeToGo(objcType, framework)

	// Never return empty string for a type - default to unsafe.Pointer
	if goType == "" {
		return "unsafe.Pointer"
	}

	// Resolve cross-framework types (e.g., CGAffineTransform -> coregraphics.CGAffineTransform)
	goType = resolveType(framework, goType)

	return goType
}

// formatMethodParams formats method parameters for Go function signature.
// Returns: "title string, target objc.IObject, action objc.Selector"
func formatMethodParams(method *occ2go.ParsedMethod, framework string) string {
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
		goType := mapObjCTypeToGo(p.Type, framework)
		parts[i] = fmt.Sprintf("%s %s", paramName, goType)
	}
	return strings.Join(parts, ", ")
}

// formatMethodParamNames formats method parameter names for calling.
// Returns: "title, target, action"
func formatMethodParamNames(method *occ2go.ParsedMethod) string {
	return formatMethodParamNamesWithFramework(method, "")
}

// formatMethodParamNamesWithFramework formats method parameter names for objc.Send calls.
// If framework is provided, it will wrap string parameters with objc.String().
func formatMethodParamNamesWithFramework(method *occ2go.ParsedMethod, framework string) string {
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

		// Wrap string parameters with objc.String() to convert Go strings to NSString*
		if framework != "" {
			goType := mapObjCTypeToGo(p.Type, framework)
			if goType == "string" {
				paramName = "objc.String(" + paramName + ")"
			}
		}

		parts[i] = paramName
	}
	return strings.Join(parts, ", ")
}

// isConstructor checks if a method is a constructor (returns instance of class).
func isConstructor(method *occ2go.ParsedMethod, className string) bool {
	// Class methods that start with class name or common constructor prefixes
	if !method.IsClassMethod {
		return false
	}

	// Check if it's a factory method that returns the class type
	selector := strings.ToLower(method.Selector)
	classNameLower := strings.ToLower(stripObjCPrefix(className))

	if strings.HasPrefix(selector, classNameLower) {
		return true
	}

	// Common factory method patterns
	factoryPrefixes := []string{"new", "create", "make", "alloc"}
	for _, prefix := range factoryPrefixes {
		if strings.HasPrefix(selector, prefix) {
			return true
		}
	}

	return false
}

// stripNSPrefix is an alias for stripObjCPrefix for clarity in templates
func stripNSPrefix(className string) string {
	return stripObjCPrefix(className)
}

// needsFoundationImport checks if any types require foundation import
func needsFoundationImport(methods []*occ2go.ParsedMethod) bool {
	for _, m := range methods {
		// Check return type
		if strings.Contains(m.ReturnType, "NS") &&
		   !strings.Contains(m.ReturnType, "NSInteger") &&
		   !strings.Contains(m.ReturnType, "NSUInteger") {
			return true
		}
		// Check parameters
		for _, p := range m.Parameters {
			if strings.Contains(p.Type, "NS") &&
			   !strings.Contains(p.Type, "NSInteger") &&
			   !strings.Contains(p.Type, "NSUInteger") {
				return true
			}
		}
	}
	return false
}

// needsQuartzCoreImport checks if any types require quartzcore import
func needsQuartzCoreImport(methods []*occ2go.ParsedMethod) bool {
	for _, m := range methods {
		// Check return type
		if strings.Contains(m.ReturnType, "CA") || strings.Contains(m.ReturnType, "CI") {
			return true
		}
		// Check parameters
		for _, p := range m.Parameters {
			if strings.Contains(p.Type, "CA") || strings.Contains(p.Type, "CI") {
				return true
			}
		}
	}
	return false
}

// prepareClassMethods filters methods to return only class methods, deduplicated by Go method name.
// When multiple methods would generate the same Go method name (e.g., foo and foo:),
// the methods are disambiguated by appending parameter labels and the Name field is updated.
func prepareClassMethods(methods []*occ2go.ParsedMethod) []*occ2go.ParsedMethod {
	result := make([]*occ2go.ParsedMethod, 0)
	seenSelectors := make(map[string]bool)

	// First pass: collect all class methods
	var classMethods []*occ2go.ParsedMethod
	for _, m := range methods {
		if m.IsClassMethod && !seenSelectors[m.Selector] {
			classMethods = append(classMethods, m)
			seenSelectors[m.Selector] = true
		}
	}

	// Second pass: detect Go method name collisions and disambiguate
	goNameCounts := make(map[string]int)

	// Count how many methods map to each Go name
	for _, m := range classMethods {
		goName := selectorToGoName(m.Selector)
		goNameCounts[goName]++
	}

	// Third pass: build result with disambiguation, updating .Name as needed
	seenGoNames := make(map[string]bool)
	for _, m := range classMethods {
		goName := selectorToGoName(m.Selector)

		// If this Go name has duplicates, disambiguate using parameter labels
		if goNameCounts[goName] > 1 {
			// Create a copy of the method and update its Name field
			methodCopy := *m
			methodCopy.Name = disambiguateMethodName(m)
			goName = methodCopy.Name

			// Skip if we've already seen this exact Go name (shouldn't happen after disambiguation)
			if seenGoNames[goName] {
				continue
			}

			result = append(result, &methodCopy)
			seenGoNames[goName] = true
		} else {
			// No collision, use original method
			if seenGoNames[goName] {
				continue
			}
			result = append(result, m)
			seenGoNames[goName] = true
		}
	}

	return result
}

// prepareInstanceMethods filters methods to return only instance methods,
// excluding ALL init methods (which are converted to constructors), deduplicated by Go method name.
// When multiple methods would generate the same Go method name (e.g., foo and foo:),
// the methods are disambiguated by appending parameter labels and the Name field is updated.
func prepareInstanceMethods(methods []*occ2go.ParsedMethod) []*occ2go.ParsedMethod {
	result := make([]*occ2go.ParsedMethod, 0)
	seenSelectors := make(map[string]bool)

	// First pass: collect all instance methods (excluding init)
	var instanceMethods []*occ2go.ParsedMethod
	for _, m := range methods {
		if !m.IsClassMethod && !seenSelectors[m.Selector] {
			// Skip ALL init methods - they're converted to package-level constructors
			if strings.HasPrefix(m.Selector, "init") {
				continue
			}
			instanceMethods = append(instanceMethods, m)
			seenSelectors[m.Selector] = true
		}
	}

	// Second pass: detect Go method name collisions and disambiguate
	goNameCounts := make(map[string]int)

	// Count how many methods map to each Go name
	for _, m := range instanceMethods {
		goName := selectorToGoName(m.Selector)
		goNameCounts[goName]++
	}

	// Third pass: build result with disambiguation, updating .Name as needed
	seenGoNames := make(map[string]bool)
	for _, m := range instanceMethods {
		goName := selectorToGoName(m.Selector)

		// If this Go name has duplicates, disambiguate using parameter labels
		if goNameCounts[goName] > 1 {
			// Create a copy of the method and update its Name field
			methodCopy := *m
			methodCopy.Name = disambiguateMethodName(m)
			goName = methodCopy.Name

			// Skip if we've already seen this exact Go name (shouldn't happen after disambiguation)
			if seenGoNames[goName] {
				continue
			}

			result = append(result, &methodCopy)
			seenGoNames[goName] = true
		} else {
			// No collision, use original method
			if seenGoNames[goName] {
				continue
			}
			result = append(result, m)
			seenGoNames[goName] = true
		}
	}

	return result
}

// filterPropertyMethods removes methods that are generated from properties (getters/setters)
// to prevent duplicate generation. Property methods are generated separately in the properties section.
func filterPropertyMethods(class *occ2go.ParsedClass) []*occ2go.ParsedMethod {
	if class == nil {
		return nil
	}

	// Build set of property selectors (getter and setter)
	// Also build a set of property setter METHOD names (without colon) that would collide with generated setters
	propertySelectors := make(map[string]bool)
	propertySetterMethods := make(map[string]bool) // Maps "setFoo" -> true if property "foo" exists
	for _, prop := range class.Properties {
		// Getter selector is just the property name
		propertySelectors[prop.Name] = true

		// Setter selector is "set<CapitalizedName>:"
		capitalizedName := strings.ToUpper(prop.Name[:1]) + prop.Name[1:]
		setterSelector := "set" + capitalizedName + ":"
		propertySelectors[setterSelector] = true

		// Also track the setter method name without colon for collision detection
		// E.g., property "accessibilityFrameInParentSpace" -> method "setAccessibilityFrameInParentSpace"
		setterMethodName := "set" + capitalizedName
		propertySetterMethods[setterMethodName] = true
	}

	// Filter methods, excluding those that match property selectors
	var filtered []*occ2go.ParsedMethod

	// Build a map of method selectors to detect setter-like collisions
	methodSelectors := make(map[string]*occ2go.ParsedMethod)
	for _, m := range class.Methods {
		methodSelectors[m.Selector] = m
	}

	for _, m := range class.Methods {
		if !m.IsClassMethod && !propertySelectors[m.Selector] {
			// Skip parameterless set* methods that would collide with property setters
			// Example: method setAccessibilityFrameInParentSpace() collides with property accessibilityFrameInParentSpace's setter
			if strings.HasPrefix(m.Selector, "set") && len(m.Parameters) == 0 && !strings.HasSuffix(m.Selector, ":") {
				if propertySetterMethods[m.Selector] {
					// Skip this method because it collides with a generated property setter
					continue
				}

				// Also check if there's a setter version with a colon (method overload case)
				setterVersion := m.Selector + ":"
				if setterMethod, exists := methodSelectors[setterVersion]; exists && len(setterMethod.Parameters) > 0 {
					// Skip this parameterless version as it collides with the parameterized method
					continue
				}
			}
			filtered = append(filtered, m)
		} else if m.IsClassMethod {
			// Always include class methods
			filtered = append(filtered, m)
		}
	}

	// Now call prepareInstanceMethods on the filtered list
	return prepareInstanceMethods(filtered)
}

// prepareInitMethods filters methods to return only init methods (for constructor generation), deduplicated by constructor name.
// Includes both:
//  - Instance methods with selectors starting with "init" (traditional init methods)
//  - Class methods marked as initializers in documentation (factory methods like buttonWithTitle:target:action:)
// When multiple methods would generate the same constructor name (e.g., initWithContentsOfURL: and arrayWithContentsOfURL:),
// the instance method is preferred.
func prepareInitMethods(methods []*occ2go.ParsedMethod) []*occ2go.ParsedMethod {
	result := make([]*occ2go.ParsedMethod, 0)
	seen := make(map[string]*occ2go.ParsedMethod)

	// First pass: collect all potential init methods with their constructor names
	for _, m := range methods {
		isInit := false

		// Traditional instance init methods
		if !m.IsClassMethod && strings.HasPrefix(m.Selector, "init") {
			isInit = true
		}

		// Class factory methods marked as initializers in docs (e.g., buttonWithTitle:target:action:)
		if m.IsClassMethod && m.IsInitializer {
			isInit = true
		}

		if isInit {
			// We don't have the className here, so we'll use a simple dedup strategy:
			// Prefer instance methods over class methods with similar signatures
			key := m.Selector
			existing, exists := seen[key]

			if !exists {
				seen[key] = m
			} else {
				// If we have both an instance and class method, prefer instance
				// Instance methods take precedence because they're the "real" initializers
				if !m.IsClassMethod && existing.IsClassMethod {
					seen[key] = m
				}
			}
		}
	}

	// Second pass: deduplicate by actual constructor name
	constructorNames := make(map[string]bool)
	for _, m := range methods {
		if seenMethod, exists := seen[m.Selector]; exists && seenMethod == m {
			// Generate constructor name (we need className, but we don't have it here)
			// So we'll do a simpler check: deduplicate by parameter signature
			paramSig := fmt.Sprintf("%d", len(m.Parameters))
			for _, p := range m.Parameters {
				paramSig += ":" + p.Type
			}
			constructorKey := m.Selector + paramSig

			if !constructorNames[constructorKey] {
				result = append(result, m)
				constructorNames[constructorKey] = true
			}
		}
	}
	return result
}

// initMethodToConstructorName converts an init method selector to a constructor function name.
// Handles both traditional init methods and class factory methods.
// Examples:
//   "init" -> "NewButton"
//   "initWithFrame:" -> "NewButtonWithFrame"
//   "buttonWithTitle:target:action:" -> "NewButtonWithTitleTargetAction"
//   "checkboxWithTitle:target:action:" -> "NewCheckboxWithTitleTargetAction"
func initMethodToConstructorName(className, selector string) string {
	structName := classToStructName(className)

	// Special case for plain "init"
	if selector == "init" {
		return "New" + structName
	}

	// Check if this is a traditional init method (starts with "init")
	if strings.HasPrefix(selector, "init") {
		// Strip "init" prefix
		selector = strings.TrimPrefix(selector, "init")

		// Convert selector to Go name (handles colons, capitalization)
		goName := occ2go.SelectorToGoName("init" + selector)

		// Replace "Init" prefix with "New{ClassName}"
		if strings.HasPrefix(goName, "Init") {
			return "New" + structName + strings.TrimPrefix(goName, "Init")
		}

		return "New" + structName + goName
	}

	// This is a class factory method (e.g., buttonWithTitle:target:action:, kernelWithString:)
	// Convert the entire selector to Go name
	goName := occ2go.SelectorToGoName(selector)

	// Check if the selector and class name share a common suffix/prefix to avoid doubling
	// Example: BlendKernel.kernelWithString: -> "New" + "BlendKernel" + "KernelWithString"
	// We want: NewBlendKernelWithString (not NewBlendKernelKernelWithString)
	// Strategy: If structName ends with the same word that goName starts with, merge them

	// Try to find the common overlap
	for i := 1; i <= len(structName) && i <= len(goName); i++ {
		if strings.HasSuffix(structName, goName[:i]) {
			// Found overlap: structName ends with first i chars of goName
			// Return structName + rest of goName
			return "New" + structName + goName[i:]
		}
	}

	// No overlap, just concatenate
	return "New" + structName + goName
}

// prepareInitMethodsWithClassName properly deduplicates init methods by their generated constructor names.
// This fixes cases where both instance and class factory methods would generate the same constructor name
// (e.g., initWithContentsOfURL: and arrayWithContentsOfURL: both map to NewArrayWithContentsOfURL).
// Instance methods are preferred over class factory methods when deduplicating.
func prepareInitMethodsWithClassName(className string, methods []*occ2go.ParsedMethod) []*occ2go.ParsedMethod {
	// First get all potential init methods
	initMethods := prepareInitMethods(methods)

	// Deduplicate by constructor name
	seen := make(map[string]*occ2go.ParsedMethod)
	for _, m := range initMethods {
		constructorName := initMethodToConstructorName(className, m.Selector)

		existing, exists := seen[constructorName]
		if !exists {
			seen[constructorName] = m
		} else {
			// If we have both an instance and class method mapping to the same name,
			// prefer the instance method
			if !m.IsClassMethod && existing.IsClassMethod {
				seen[constructorName] = m
			}
		}
	}

	// Convert map back to slice
	result := make([]*occ2go.ParsedMethod, 0, len(seen))
	for _, m := range seen {
		result = append(result, m)
	}

	return result
}

// classDependsOnCoreGraphics returns true if any method in the class uses CoreGraphics types
// that actually require the coregraphics import (i.e., not mapped to unsafe.Pointer).
func classDependsOnCoreGraphics(methods []*occ2go.ParsedMethod, framework string) bool {
	if framework == "CoreGraphics" {
		return false
	}

	for _, m := range methods {
		// Check return type - use the same mapping logic as getRequiredImports
		if m.ReturnType != "" && m.ReturnType != "void" {
			goType := mapObjCTypeToGo(m.ReturnType, framework)
			// Check if this actually needs coregraphics import (not unsafe.Pointer)
			if strings.HasPrefix(goType, "coregraphics.") {
				return true
			}
		}

		// Check parameters - use the same mapping logic as getRequiredImports
		for _, p := range m.Parameters {
			goType := mapObjCTypeToGo(p.Type, framework)
			// Check if this actually needs coregraphics import (not unsafe.Pointer)
			if strings.HasPrefix(goType, "coregraphics.") {
				return true
			}
		}
	}

	return false
}

// sortMethodsByName sorts methods by name for consistent output
func sortMethodsByName(methods []*occ2go.ParsedMethod) []*occ2go.ParsedMethod {
	sorted := make([]*occ2go.ParsedMethod, len(methods))
	copy(sorted, methods)

	// Simple bubble sort by Name
	for i := 0; i < len(sorted)-1; i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[i].Name > sorted[j].Name {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	return sorted
}

// wrapObjCReturn generates the return statement for converting objc.ID to Go types.
// It handles special cases like bool conversion and objc.Object mapping.
// Examples:
//   wrapObjCReturn("bool") -> "ret != 0"
//   wrapObjCReturn("objc.Object") -> "objc.ID(ret)"
//   wrapObjCReturn("int") -> "int(ret)"
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

// dict creates a map from alternating key-value pairs.
// Usage: {{template "name" (dict "key1" .Value1 "key2" .Value2)}}
func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("dict requires an even number of arguments")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

// propertyToGoName converts an Objective-C property name to a Go method name.
// Examples:
//   title -> Title
//   isEnabled -> IsEnabled
//   backgroundColor -> BackgroundColor
func propertyToGoName(propName string) string {
	if propName == "" {
		return ""
	}
	// Special case: "object" property conflicts with embedded Object type
	// Generate "GetObject()" instead of "Object()"
	if strings.ToLower(propName) == "object" {
		return "GetObject"
	}
	// Capitalize first letter
	return strings.ToUpper(propName[:1]) + propName[1:]
}

// sliceContainsString checks if a string slice contains a specific string.
func sliceContainsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// capitalizeFirst capitalizes the first letter of a string.
// Examples:
//   title -> Title
//   backgroundColor -> BackgroundColor
func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	// Special case: "object" property conflicts with embedded Object type
	// Generate "GetObject" instead of "Object"
	if strings.ToLower(s) == "object" {
		return "GetObject"
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// needsCustomImports checks if any methods use types that require custom imports
func needsCustomImports(methods []*occ2go.ParsedMethod, framework string) bool {
	return getRequiredImports(methods, framework) != nil
}

// getRequiredImports returns a map of import paths needed for methods.
// It maps Objective-C types to Go types first, then checks if those Go types need imports.
// This prevents adding imports for types that get mapped to unsafe.Pointer or built-in types.
// For example: {"github.com/tmc/appledocs/generated/coregraphics": true}
func getRequiredImports(methods []*occ2go.ParsedMethod, framework string) map[string]bool {
	imports := make(map[string]bool)

	// Check all methods for types that need custom imports
	for _, m := range methods {
		// Check return type - map to Go first, then check if it needs an import
		if m.ReturnType != "" && m.ReturnType != "void" {
			goType := mapObjCTypeToGo(m.ReturnType, framework)
			// Check if this Go type needs an import (e.g., coregraphics.CGAffineTransform)
			if importPath := getGoTypeImportPath(goType); importPath != "" {
				imports[importPath] = true
			}
		}

		// Check parameters - map to Go first, then check if they need imports
		for _, p := range m.Parameters {
			goType := mapObjCTypeToGo(p.Type, framework)
			// Check if this Go type needs an import
			if importPath := getGoTypeImportPath(goType); importPath != "" {
				imports[importPath] = true
			}
		}
	}

	if len(imports) == 0 {
		return nil
	}
	return imports
}

// getGoTypeImportPath takes a Go type string (after mapping from Objective-C) and returns
// the import path if it requires one, or empty string if it doesn't.
// Examples:
//   "coregraphics.CGAffineTransform" -> "github.com/tmc/appledocs/generated/coregraphics"
//   "foundation.Rect" -> "github.com/tmc/appledocs/generated/foundation"
//   "unsafe.Pointer" -> ""
//   "int" -> ""
//   "bool" -> ""
func getGoTypeImportPath(goType string) string {
	// Built-in types and types from std library don't need custom imports
	if goType == "" || goType == "unsafe.Pointer" {
		return ""
	}

	// Check for framework-prefixed types (e.g., "coregraphics.CGAffineTransform")
	if strings.Contains(goType, ".") {
		parts := strings.SplitN(goType, ".", 2)
		if len(parts) == 2 {
			packageName := parts[0]
			// Map package names to import paths
			switch packageName {
			case "coregraphics":
				return "github.com/tmc/appledocs/generated/coregraphics"
			case "foundation":
				return "github.com/tmc/appledocs/generated/foundation"
			case "quartzcore":
				return "github.com/tmc/appledocs/generated/quartzcore"
			case "appkit":
				return "github.com/tmc/appledocs/generated/appkit"
			case "usernotifications":
				return "github.com/tmc/appledocs/generated/usernotifications"
			case "objc":
				// objc is already imported by default in the template
				return ""
			}
		}
	}

	return ""
}

// ImportInfo holds information about an import for template rendering
type ImportInfo struct {
	PackageName string
	ImportPath  string
}

// sortedImportPaths returns a sorted slice of import paths for template iteration.
// This makes it easy for templates to range over imports in a consistent order.
func sortedImportPaths(imports map[string]bool) []ImportInfo {
	if imports == nil || len(imports) == 0 {
		return []ImportInfo{}
	}

	// Convert map keys to slice
	paths := make([]string, 0, len(imports))
	for path := range imports {
		paths = append(paths, path)
	}

	// Simple sort by string value for consistency
	for i := 0; i < len(paths)-1; i++ {
		for j := i + 1; j < len(paths); j++ {
			if paths[i] > paths[j] {
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}

	// Convert paths to ImportInfo structs
	result := make([]ImportInfo, 0, len(paths))
	for _, path := range paths {
		packageName := extractPackageNameFromImportPath(path)
		result = append(result, ImportInfo{
			PackageName: packageName,
			ImportPath:  path,
		})
	}

	return result
}

// extractPackageNameFromImportPath extracts the package name from an import path.
// For example: "github.com/tmc/appledocs/generated/coregraphics" -> "coregraphics"
func extractPackageNameFromImportPath(importPath string) string {
	parts := strings.Split(importPath, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return importPath
}

// getClassRequiredImports returns a sorted slice of all required imports for a class (including both methods and properties).
// This is a convenience function for templates to get all imports at once.
func getClassRequiredImports(class interface{}, framework string) []string {
	imports := make(map[string]bool)
	currentFrameworkImportPath := "github.com/tmc/appledocs/generated/" + strings.ToLower(framework)

	// Try to extract class data using reflection
	classVal := reflect.ValueOf(class)
	if classVal.Kind() == reflect.Ptr {
		classVal = classVal.Elem()
	}

	// Try to access Methods field
	if classVal.Kind() == reflect.Struct {
		methodsField := classVal.FieldByName("Methods")
		if methodsField.IsValid() && methodsField.Kind() == reflect.Slice {
			for i := 0; i < methodsField.Len(); i++ {
				method := methodsField.Index(i).Interface()
				if parsedMethod, ok := method.(*occ2go.ParsedMethod); ok {
					// Check return type
					if parsedMethod.ReturnType != "" {
						goType := mapObjCTypeToGo(parsedMethod.ReturnType, framework)
						if importPath := getGoTypeImportPath(goType); importPath != "" {
							if importPath != currentFrameworkImportPath {
								imports[importPath] = true
							}
						}
					}

					// Check parameters
					for _, param := range parsedMethod.Parameters {
						goType := mapObjCTypeToGo(param.Type, framework)
						if importPath := getGoTypeImportPath(goType); importPath != "" {
							if importPath != currentFrameworkImportPath {
								imports[importPath] = true
							}
						}
					}
				}
			}
		}

		// Try to access Properties field
		propertiesField := classVal.FieldByName("Properties")
		if propertiesField.IsValid() && propertiesField.Kind() == reflect.Slice {
			for i := 0; i < propertiesField.Len(); i++ {
				property := propertiesField.Index(i).Interface()
				if parsedProp, ok := property.(*occ2go.ParsedProperty); ok {
					goType := mapObjCTypeToGo(parsedProp.Type, framework)
					if importPath := getGoTypeImportPath(goType); importPath != "" {
						if importPath != currentFrameworkImportPath {
							imports[importPath] = true
						}
					}
				}
			}
		}
	}

	// Convert map to sorted slice
	result := make([]string, 0, len(imports))
	for imp := range imports {
		result = append(result, imp)
	}
	sort.Strings(result)
	return result
}

// getFunctionRequiredImports analyzes standalone functions to collect required import paths.
// It examines each function's return type and parameters, applying type mappings and extracting
// import paths from qualified type names (e.g., "coregraphics.CGAffineTransform").
func getFunctionRequiredImports(functions []*occ2go.ParsedFunction, framework string) map[string]bool {
	imports := make(map[string]bool)

	// Build the current framework's import path to filter it out
	currentFrameworkImportPath := "github.com/tmc/appledocs/generated/" + strings.ToLower(framework)

	for _, fn := range functions {
		// Check return type
		if fn.ReturnType != "" && fn.ReturnType != "void" {
			// Use mapCTypeToGoWithFramework to get the same result as prepareFunctionData
			goType := mapCTypeToGoWithFramework(fn.ReturnType, framework)
			if importPath := getGoTypeImportPath(goType); importPath != "" {
				// Don't import the current framework itself
				if importPath != currentFrameworkImportPath {
					imports[importPath] = true
				}
			}
		}

		// Check all parameters
		for _, param := range fn.Parameters {
			// Use mapCTypeToGoWithFramework to get the same result as prepareFunctionData
			goType := mapCTypeToGoWithFramework(param.Type, framework)
			if importPath := getGoTypeImportPath(goType); importPath != "" {
				// Don't import the current framework itself
				if importPath != currentFrameworkImportPath {
					imports[importPath] = true
				}
			}
		}
	}

	return imports
}

// mergeImports merges two import maps into a single deduplicated map.
// Useful for combining imports from class methods and instance methods.
func mergeImports(map1, map2 map[string]bool) map[string]bool {
	result := make(map[string]bool)

	// Add all imports from map1
	for path := range map1 {
		result[path] = true
	}

	// Add all imports from map2
	for path := range map2 {
		result[path] = true
	}

	return result
}

// commentLine formats a string as a single-line Go comment, handling newlines and special characters.
// Multi-line text is collapsed to a single line with spaces.
func commentLine(s string) string {
	if s == "" {
		return ""
	}
	// Replace newlines with spaces
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", " ")

	// Collapse multiple spaces
	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}

	return strings.TrimSpace(s)
}

// isEssentialSelector checks if a selector is one of the essential methods
// we generate manually (alloc, new, init, autorelease, etc.) for classes with init methods.
// These should be skipped when generating methods from Apple's documentation to avoid duplicates.
func isEssentialSelector(selector string) bool {
	essentialSelectors := []string{
		"alloc",
		"allocWithZone:",
		"new",
		"init",
		"autorelease",
		"copy",
		"copyWithZone:",
		"mutableCopy",
		"mutableCopyWithZone:",
	}

	for _, essential := range essentialSelectors {
		if selector == essential {
			return true
		}
	}
	return false
}

// convertDocURL converts Apple's doc:// scheme URLs to https:// URLs.
// If the URL doesn't start with "doc://", it returns it unchanged.
// Examples:
//   doc://com.apple.foundation/documentation/Foundation/NSString -> https://developer.apple.com/documentation/foundation/nsstring
//   https://developer.apple.com/... -> https://developer.apple.com/... (unchanged)
func convertDocURL(url string) string {
	if url == "" {
		return ""
	}

	// If it's already an https URL or doesn't start with doc://, return as-is
	if !strings.HasPrefix(url, "doc://") {
		return url
	}

	// Strip doc:// prefix
	url = strings.TrimPrefix(url, "doc://")

	// Remove the domain part (e.g., com.apple.foundation, com.apple.objectivec)
	parts := strings.SplitN(url, "/", 2)
	if len(parts) < 2 {
		// Malformed URL, return empty to skip it
		return ""
	}

	// Construct https URL
	return "https://developer.apple.com/" + parts[1]
}

// classHasInit checks if a class has any init methods (including factory initializers).
// This is used to determine if a test file should be generated.
// Returns true if the class has:
//  - Any instance init method (selector starting with "init")
//  - Any class method marked as an initializer in docs (IsInitializer = true)
func classHasInit(methods []*occ2go.ParsedMethod) bool {
	for _, m := range methods {
		// Instance init methods
		if !m.IsClassMethod && strings.HasPrefix(m.Selector, "init") {
			return true
		}
		// Class factory methods marked as initializers
		if m.IsClassMethod && m.IsInitializer {
			return true
		}
	}
	return false
}

// generateTestValue generates an example test value for a given Go type.
// Returns the Go code as a string that can be used in test examples.
// Handles cross-framework dependencies properly.
// The packageName parameter is used to properly qualify package-local types in test files.
func generateTestValue(goType, framework, paramName string) string {
	// Handle primitive types
	switch goType {
	case "string":
		// Special handling for specific parameter names that need valid values
		paramLower := strings.ToLower(paramName)
		if strings.Contains(paramLower, "querystring") || strings.Contains(paramLower, "query") && strings.Contains(paramLower, "string") {
			// Metadata query string - needs valid syntax
			return `"kMDItemFSName == '*.txt'"`
		}
		if strings.Contains(paramLower, "path") {
			// Path parameters need to be absolute paths
			return `"/tmp/test"`
		}
		if strings.Contains(paramLower, "url") || strings.Contains(paramLower, "urlstring") {
			// URL strings
			return `"https://example.com"`
		}
		return fmt.Sprintf(`"%s"`, paramName)
	case "int", "int8", "int16", "int32", "int64":
		return "0"
	case "uint", "uint8", "uint16", "uint32", "uint64":
		return "0"
	case "float32", "float64":
		return "0.0"
	case "bool":
		return "false"
	case "objc.ID":
		return "0"
	case "objc.SEL":
		return "0"
	case "objc.Class":
		return "0"
	case "unsafe.Pointer":
		return "nil"
	}

	// Handle CoreGraphics geometry types - only when explicitly qualified
	// These are cross-framework types that need full qualification
	if goType == "coregraphics.CGRect" || goType == "CGRect" && framework != "Foundation" {
		return "coregraphics.CGRect{}"
	}
	if goType == "coregraphics.CGSize" || goType == "CGSize" && framework != "Foundation" {
		return "coregraphics.CGSize{}"
	}
	if goType == "coregraphics.CGPoint" || goType == "CGPoint" && framework != "Foundation" {
		return "coregraphics.CGPoint{}"
	}

	// Handle framework-specific types
	// Foundation types - when already qualified with "foundation.", return as-is
	// The template will NOT add a prefix because it checks hasPrefix("foundation.")
	if strings.HasPrefix(goType, "foundation.") {
		typeName := strings.TrimPrefix(goType, "foundation.")
		switch typeName {
		case "Rect":
			return "foundation.Rect{}"
		case "Size":
			return "foundation.Size{}"
		case "Point":
			return "foundation.Point{}"
		case "Range":
			return "foundation.Range{}"
		case "TimeInterval":
			// TimeInterval is a type alias for float64, not a struct
			return "foundation.TimeInterval(0.0)"
		default:
			// Other foundation types - try to use zero value or constructor
			return fmt.Sprintf("%s{}", goType)
		}
	}

	// Handle package-local types (no dot) - these need to be qualified with packageName in test files
	// The template WILL add the package prefix for these
	if !strings.Contains(goType, ".") {
		// Special cases for known type aliases (not structs)
		if goType == "TimeInterval" {
			return "TimeInterval(0.0)"
		}
		// Could be an enum or a struct from the same package
		// Use struct literal syntax instead of type cast
		// We return the unqualified type name; the template will add the package prefix
		return fmt.Sprintf("%s{}", goType)
	}

	// Default: try zero value for the type
	return fmt.Sprintf("%s{}", goType)
}

// generateTestValueWithPackage generates a test value and applies the proper package prefix
// based on the Go type and target package context. This simplifies the template logic by
// handling all the package qualification rules in one place.
//
// Parameters:
//   - goType: The mapped Go type (e.g., "coregraphics.CGRect", "Rect", "string")
//   - framework: The framework context for type mapping
//   - packageName: The package name where the test will be generated (e.g., "foundation")
//   - paramName: The parameter name for context-sensitive test values
//
// Returns a fully-qualified test value expression ready to use in generated code.
func generateTestValueWithPackage(goType, framework, packageName, paramName string) string {
	testValue := generateTestValue(goType, framework, paramName)

	// Check if the test value already starts with a package qualifier
	// (e.g., "coregraphics.CGRect{}", "foundation.Range{}")
	if strings.HasPrefix(testValue, "foundation.") ||
		strings.HasPrefix(testValue, "coregraphics.") ||
		strings.HasPrefix(testValue, "objc.") ||
		strings.HasPrefix(testValue, "unsafe.") {
		// Already fully qualified - return as-is
		return testValue
	}

	// Check if this is a primitive type or stdlib type that doesn't need qualification
	switch goType {
	case "string", "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "bool":
		return testValue
	}

	// Check if goType already has a package prefix (belt and suspenders check)
	if strings.HasPrefix(goType, "foundation.") ||
		strings.HasPrefix(goType, "coregraphics.") ||
		strings.HasPrefix(goType, "objc.") ||
		strings.HasPrefix(goType, "unsafe.") {
		// Type is already qualified, return test value as-is
		return testValue
	}

	// Package-local type - needs to be qualified with packageName
	return packageName + "." + testValue
}

// canGenerateTestValue checks if we can generate a reasonable test value for the given type and parameter name.
// Returns true if generateTestValue will produce a usable value.
// Params may include the parameter name (for context-sensitive filtering like file paths).
func canGenerateTestValue(args ...string) bool {
	if len(args) == 0 {
		return false
	}
	goType := args[0]
	paramName := ""
	if len(args) > 1 {
		paramName = strings.ToLower(args[1])
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

	// We can handle package-local types (enums and structs)
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

// resolveType resolves a type name to its fully qualified name, handling cross-framework dependencies.
// Takes the current framework context and a type name (e.g., "MutableAttributedString") and returns
// either the unqualified name (if it's in the same framework) or a qualified name (e.g., "foundation.MutableAttributedString").
// This helper is used in templates to properly reference types that may come from other frameworks.
//
// Examples:
//   resolveType("AppKit", "Button") -> "Button" (same framework)
//   resolveType("AppKit", "MutableAttributedString") -> "foundation.MutableAttributedString" (cross-framework)
//   resolveType("Foundation", "Array") -> "Array" (same framework)
func resolveType(framework, typeName string) string {
	if typeName == "" {
		return ""
	}

	// Common Foundation base classes that other frameworks inherit from
	foundationTypes := map[string]bool{
		"MutableAttributedString": true,
		"AttributedString":        true,
		"Array":                   true,
		"MutableArray":            true,
		"Dictionary":              true,
		"MutableDictionary":       true,
		"Set":                     true,
		"MutableSet":              true,
		"String":                  true,
		"MutableString":           true,
		"Data":                    true,
		"MutableData":             true,
		"Date":                    true,
		"URL":                     true,
		"NSURL":                   true, // Include NS-prefixed version
		"URLRequest":              true,
		"MutableURLRequest":       true,
		"Value":                   true,
		"Number":                  true,
		"NSNumber":                true, // Include NS-prefixed version
		"URLSession":              true,
		"URLSessionTask":          true,
		"URLSessionDataTask":      true,
		"URLSessionUploadTask":    true,
		"URLSessionDownloadTask":  true,
		"URLSessionStreamTask":    true,
		"Enumerator":              true, // For OSLog.OSLogEnumerator
		"Operation":               true,
		"OperationQueue":          true,
		"Expression":              true, // NSExpression - used by CoreData
		"ExtensionContext":        true, // NSExtensionContext - used by AuthenticationServices
		"Coder":                   true, // NSCoder - base class for archiving
		"KeyedArchiver":           true, // NSKeyedArchiver
		"KeyedUnarchiver":         true, // NSKeyedUnarchiver - used by MetalPerformanceShaders
	}

	// QuartzCore types used by other frameworks
	quartzCoreTypes := map[string]bool{
		"Layer":          true, // CALayer
		"Animation":      true, // CAAnimation
		"MediaTiming":    true, // CAMediaTiming protocol
		"Transaction":    true, // CATransaction
		"TransformLayer": true, // CATransformLayer
		"OpenGLLayer":    true, // CAOpenGLLayer - used by NSOpenGLLayer in AppKit
	}

	// AppKit types used by other frameworks (common base classes)
	appKitTypes := map[string]bool{
		"Responder":            true, // NSResponder
		"View":                 true, // NSView
		"Control":              true, // NSControl
		"Window":               true, // NSWindow
		"ViewController":       true, // NSViewController
		"NavigationController": true, // NSNavigationController (though less common on macOS)
		"Panel":                true, // NSPanel
		"Application":          true, // NSApplication
		"Document":             true, // NSDocument
		"WindowController":     true, // NSWindowController
		"Menu":                 true, // NSMenu
		"MenuItem":             true, // NSMenuItem
	}

	// CoreGraphics types used by other frameworks
	coreGraphicsTypes := map[string]bool{
		// Struct types
		"CGAffineTransform": true,
		"CGPoint":           true,
		"CGSize":            true,
		"CGRect":            true,
		"CGVector":          true,
		"CGFloat":           true,
		// Opaque ref types
		"CGColorRef":         true,
		"CGColorSpaceRef":    true,
		"CGContextRef":       true,
		"CGImageRef":         true,
		"CGImageSourceRef":   true,
		"CGImageDestinationRef": true,
		"CGPathRef":          true,
		"CGLayerRef":         true,
		"CGFontRef":          true,
		"CGDataProviderRef":  true,
		"CGDataConsumerRef":  true,
		"CGFunctionRef":      true,
		"CGShadingRef":       true,
		"CGGradientRef":      true,
		"CGPatternRef":       true,
		"CGPDFDocumentRef":   true,
		"CGPDFPageRef":       true,
	}

	// Check if the type exists in current framework FIRST before adding qualifications
	// This prevents self-imports (e.g., coregraphics.CGAffineTransform in CoreGraphics)
	// Strip the ObjC prefix before checking, since currentFrameworkClasses contains stripped names
	strippedTypeName := stripObjCPrefix(typeName)
	if currentFrameworkClasses[strippedTypeName] {
		// DEBUG: Uncomment to debug same-framework type resolution
		// fmt.Fprintf(os.Stderr, "DEBUG resolveType: Found '%s' (stripped: '%s') in current framework '%s', returning as-is\n", typeName, strippedTypeName, framework)
		// It's in the current framework, return as-is
		return typeName
	}

	// If we're in CoreGraphics framework, all types are local
	if framework == "CoreGraphics" && coreGraphicsTypes[typeName] {
		return typeName
	}

	// If this is a known CoreGraphics type and we're not in CoreGraphics, qualify it
	if coreGraphicsTypes[typeName] {
		// Make sure the type has the CG prefix for proper type reference
		if !strings.HasPrefix(typeName, "CG") {
			return "coregraphics.CG" + typeName
		}
		return "coregraphics." + typeName
	}

	// If we're in QuartzCore framework, all types are local
	if framework == "QuartzCore" && quartzCoreTypes[typeName] {
		return typeName
	}

	// If this is a known QuartzCore type and we're not in QuartzCore, qualify it
	if quartzCoreTypes[typeName] {
		return "quartzcore." + typeName
	}

	// If we're in Foundation framework, all types are local
	if framework == "Foundation" && foundationTypes[typeName] {
		return typeName
	}

	// If we're in AppKit (or other frameworks that embed NSObject), Foundation types are also local
	// since NSObject/Foundation is embedded in the object hierarchy
	// This includes most UI/system frameworks that depend on Foundation
	if (framework == "AppKit" || framework == "QuartzCore" || framework == "CoreData" ||
		framework == "Accessibility" || framework == "Accounts" || framework == "AddressBook" ||
		framework == "AdServices" || framework == "AdSupport" || framework == "Automator" ||
		framework == "CallKit" || framework == "ClassKit" || framework == "CloudKit" ||
		framework == "Collaboration" || framework == "Contacts" || framework == "ContactsUI" ||
		framework == "CoreLocationUI" || framework == "CryptoKit" || framework == "Darwin" ||
		framework == "DeviceCheck" || framework == "DocumentPickerUI" || framework == "EventKit" ||
		framework == "EventKitUI" || framework == "ExtensionKit" || framework == "FileProvider" ||
		framework == "FileProviderUI" || framework == "GameController" || framework == "GameKit" ||
		framework == "GLKit" || framework == "HealthKit" || framework == "HealthKitUI" ||
		framework == "HomeKit" || framework == "IOSurface" || framework == "LocalAuthentication" ||
		framework == "MapKit" || framework == "MediaAccessibility" || framework == "MediaKit" ||
		framework == "MessageUI" || framework == "Messages" || framework == "Metal" ||
		framework == "MetalKit" || framework == "MetalPerformanceShaders" || framework == "ModelIO" ||
		framework == "MultipeerConnectivity" || framework == "NaturalLanguage" || framework == "Network" ||
		framework == "NotificationCenter" || framework == "PDFKit" || framework == "PencilKit" ||
		framework == "Photos" || framework == "PhotosUI" || framework == "PlaygroundSupport" ||
		framework == "PushKit" || framework == "QuickLook" || framework == "RealityKit" ||
		framework == "SafariServices" || framework == "SceneKit" || framework == "ScreenTime" ||
		framework == "Security" || framework == "SensorKit" || framework == "ServiceManagement" ||
		framework == "SharedWithYou" || framework == "SharedWithYouCore" || framework == "ShazamKit" ||
		framework == "SiriKit" || framework == "Social" || framework == "SoundAnalysis" ||
		framework == "Speech" || framework == "SpriteKit" || framework == "StoreKit" ||
		framework == "SwiftUI" || framework == "SystemConfiguration" || framework == "ThreadNetwork" ||
		framework == "UserNotifications" || framework == "UserNotificationsUI" || framework == "VideoSubscriberAccount" ||
		framework == "VideoToolbox" || framework == "Vision" || framework == "VisionKit" ||
		framework == "WatchConnectivity" || framework == "WatchKit" || framework == "WebKit" ||
		framework == "WidgetKit") && foundationTypes[typeName] {
		return typeName
	}

	// If this is a known Foundation type and we're not in Foundation/AppKit, qualify it
	if foundationTypes[typeName] {
		return "foundation." + typeName
	}

	// If we're in AppKit framework, all types are local
	if framework == "AppKit" && appKitTypes[typeName] {
		return typeName
	}

	// If this is a known AppKit type and we're not in AppKit, qualify it
	if appKitTypes[typeName] {
		return "appkit." + typeName
	}

	// Check if we know about this type from the cross-framework registry
	if frameworkPkg, found := crossFrameworkTypeRegistry[typeName]; found {
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
			"AppKit":        {"NS", "AK"},
			"Foundation":    {"NS", "CF"},
			"CoreGraphics":  {"CG"},
			"QuartzCore":    {"CA"},
			"CoreImage":     {"CI"},
			"CoreData":      {"NS", "CD"},
			"AVFoundation":  {"AV"},
			"Metal":         {"MTL"},
			"MetalKit":      {"MTK"},
			"SpriteKit":     {"SK"},
			"SceneKit":      {"SCN"},
			"CoreML":        {"ML"},
			"Vision":        {"VN"},
			"CoreLocation":  {"CL"},
			"MapKit":        {"MK"},
			"PhotoKit":      {"PH"},
			"Photos":        {"PH"},
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

	// Last resort: fall back to unsafe.Pointer for truly unknown types
	// This should be rare with a well-populated registry
	return "unsafe.Pointer"
}

// methodGoName generates the Go method name for a given Objective-C method.
// This function handles disambiguation when multiple Objective-C methods would map
// to the same Go name (e.g., foo and foo:). It takes a method and the list of all
// methods in the same category (instance or class) to detect collisions.
//
// Examples:
//   imageByInsertingIntermediate -> ImageByInsertingIntermediate (no collision)
//   imageByInsertingIntermediate: (when there's also imageByInsertingIntermediate) -> ImageByInsertingIntermediateWithCache
func methodGoName(method *occ2go.ParsedMethod, allMethods []*occ2go.ParsedMethod) string {
	// Generate the standard Go name
	goName := selectorToGoName(method.Selector)

	// Count how many methods map to the same Go name
	count := 0
	for _, m := range allMethods {
		if selectorToGoName(m.Selector) == goName {
			count++
		}
	}

	// If there's only one method with this name, no disambiguation needed
	if count == 1 {
		return goName
	}

	// Multiple methods map to the same Go name - disambiguate
	return disambiguateMethodName(method)
}

// isInheritedFromNSObject checks if a selector is likely inherited from NSObject.
// These methods should not be redeclared in child interfaces.
// Returns true for common NSObject methods that appear in most/all classes.
func isInheritedFromNSObject(selector string) bool {
	// Common NSObject instance methods that should not be redeclared
	nsobjectMethods := map[string]bool{
		// Memory management
		"retain":              true,
		"release":             true,
		"autorelease":         true,
		"retainCount":         true,

		// Object identity and comparison
		"isEqual:":            true,
		"isEqualTo:":          true,
		"hash":                true,
		"isKindOfClass:":      true,
		"isMemberOfClass:":    true,
		"conformsToProtocol:": true,
		"respondsToSelector:": true,

		// Description and debugging
		"description":         true,
		"debugDescription":    true,
		"className":           true,
		"superclass":          true,

		// KVC (Key-Value Coding)
		"valueForKey:":                    true,
		"setValue:forKey:":                true,
		"valueForKeyPath:":                true,
		"setValue:forKeyPath:":            true,
		"valueForUndefinedKey:":           true,
		"setValue:forUndefinedKey:":       true,
		"setNilValueForKey:":              true,
		"dictionaryWithValuesForKeys:":    true,
		"setValuesForKeysWithDictionary:": true,
		"valuesForKeys:":                  true,
		"takeValuesFromDictionary:":       true,

		// KVO (Key-Value Observing)
		"addObserver:forKeyPath:options:context:":                          true,
		"removeObserver:forKeyPath:":                                        true,
		"removeObserver:forKeyPath:context:":                                true,
		"willChangeValueForKey:":                                            true,
		"didChangeValueForKey:":                                             true,
		"willChange:valuesAtIndexes:forKey:":                                true,
		"didChange:valuesAtIndexes:forKey:":                                 true,
		"willChangeValueForKey:withSetMutation:usingObjects:":               true,
		"didChangeValueForKey:withSetMutation:usingObjects:":                true,
		"observationInfo":                                                   true,
		"setObservationInfo:":                                               true,
		"observeValueForKeyPath:ofObject:change:context:":                   true,
		"keyPathsForValuesAffectingValueForKey:":                            true,
		"automaticallyNotifiesObserversForKey:":                             true,

		// Notifications
		"postNotification:":                       true,
		"postNotificationName:object:":            true,
		"postNotificationName:object:userInfo:":   true,

		// Copying
		"copy":                true,
		"mutableCopy":         true,
		"copyWithZone:":       true,
		"mutableCopyWithZone:": true,

		// Archiving
		"classForCoder":       true,
		"replacementObjectForCoder:": true,
		"awakeAfterUsingCoder:":      true,

		// Forwarding
		"forwardInvocation:":  true,
		"forwardingTargetForSelector:": true,
		"methodSignatureForSelector:":  true,
		"doesNotRecognizeSelector:":    true,

		// Scripting
		"scriptingIsEqualTo:":         true,
		"scriptingIsLessThanOrEqualTo:": true,
		"scriptingIsLessThan:":        true,
		"scriptingIsGreaterThanOrEqualTo:": true,
		"scriptingIsGreaterThan:":     true,
		"scriptingBeginsWith:":        true,
		"scriptingEndsWith:":          true,
		"scriptingContains:":          true,

		// Performance
		"performSelector:":                    true,
		"performSelector:withObject:":         true,
		"performSelector:withObject:withObject:": true,
		"performSelectorOnMainThread:withObject:waitUntilDone:": true,
		"performSelector:onThread:withObject:waitUntilDone:":    true,
		"performSelectorInBackground:withObject:":               true,
		"performSelector:withObject:afterDelay:":                true,
		"performSelector:withObject:afterDelay:inModes:":        true,
		"cancelPreviousPerformRequestsWithTarget:":              true,
		"cancelPreviousPerformRequestsWithTarget:selector:object:": true,
	}

	return nsobjectMethods[selector]
}

// ClassImports holds the import paths needed for a class
type ClassImports struct {
	NeedsObjectiveC             bool
	NeedsFoundation             bool
	NeedsQuartzCore             bool
	NeedsCoreGraphics           bool
	NeedsCloudKit               bool
	NeedsAppKit                 bool
	NeedsUserNotifications      bool
	NeedsUniformTypeIdentifiers bool
}

// typeReferencesFramework checks if a Go type string contains a reference to a specific framework.
// It handles container types like []pkg.Type, map[string]pkg.Type, etc. by checking if the
// framework package name appears with a dot (pkg.) anywhere in the type string.
func typeReferencesFramework(goType, framework string) bool {
	// Check for "framework." pattern which indicates the framework is used as a package qualifier
	return strings.Contains(goType, framework+".")
}

// getClassImports analyzes a class and its methods to determine which framework imports are needed.
// This consolidates the complex import detection logic from the template into a single helper function.
// Returns a ClassImports struct with boolean flags for each potential import.
func getClassImports(class *occ2go.ParsedClass, framework, outputModule string) ClassImports {
	imports := ClassImports{}

	if class == nil {
		return imports
	}

	// Determine struct name for self-referential check
	structName := classToStructName(class.Name)

	// Check struct embedding for import needs by using getStructEmbeddedField
	// This ensures we catch all cases where objectivec.Object is embedded
	// Don't import a framework into itself
	embeddedField := getStructEmbeddedField(class, framework)
	if strings.HasPrefix(embeddedField, "objectivec.") {
		imports.NeedsObjectiveC = true
	} else if strings.HasPrefix(embeddedField, "foundation.") && framework != "Foundation" {
		imports.NeedsFoundation = true
	} else if strings.HasPrefix(embeddedField, "quartzcore.") && framework != "QuartzCore" {
		imports.NeedsQuartzCore = true
	} else if strings.HasPrefix(embeddedField, "appkit.") && framework != "AppKit" {
		imports.NeedsAppKit = true
	}

	// Also check superclass for import needs (for interface embedding)
	if framework != "ObjectiveC" && class.SuperClass != "" {
		superStructName := classToStructName(class.SuperClass)
		superResolved := resolveType(framework, superStructName)
		isSelfReferential := (superStructName == structName)

		// If superclass has a framework prefix, we need that import
		// Don't import a framework into itself
		if strings.HasPrefix(superResolved, "foundation.") && framework != "Foundation" {
			imports.NeedsFoundation = true
		} else if strings.HasPrefix(superResolved, "quartzcore.") && framework != "QuartzCore" {
			imports.NeedsQuartzCore = true
		} else if strings.HasPrefix(superResolved, "appkit.") && framework != "AppKit" {
			imports.NeedsAppKit = true
		} else if strings.HasPrefix(superResolved, "objectivec.") {
			imports.NeedsObjectiveC = true
		} else if class.SuperClass == "NSObject" || superStructName == "Object" || isSelfReferential {
			imports.NeedsObjectiveC = true
		}
	} else if framework != "ObjectiveC" {
		// No superclass specified, default to objectivec
		imports.NeedsObjectiveC = true
	}

	// Check methods for CoreGraphics dependencies
	if classDependsOnCoreGraphics(class.Methods, framework) {
		imports.NeedsCoreGraphics = true
	}

	// Check properties for CoreGraphics dependencies
	if !imports.NeedsCoreGraphics {
		for _, prop := range class.Properties {
			goType := mapObjCTypeToGo(prop.Type, framework)
			if strings.HasPrefix(goType, "coregraphics.") {
				imports.NeedsCoreGraphics = true
				break
			}
		}
	}

	// Check method parameters and return types for AppKit, QuartzCore, and CloudKit dependencies
	for _, method := range class.Methods {
		// Check return type
		if method.ReturnType != "" {
			goType := mapObjCTypeToGo(method.ReturnType, framework)
			// Don't import a framework into itself
			if typeReferencesFramework(goType, "appkit") && framework != "AppKit" {
				imports.NeedsAppKit = true
			} else if typeReferencesFramework(goType, "quartzcore") && framework != "QuartzCore" {
				imports.NeedsQuartzCore = true
			} else if typeReferencesFramework(goType, "cloudkit") && framework != "CloudKit" {
				imports.NeedsCloudKit = true
			}
		}
		// Check parameters
		for _, param := range method.Parameters {
			goType := mapObjCTypeToGo(param.Type, framework)
			// Don't import a framework into itself
			if typeReferencesFramework(goType, "foundation") && framework != "Foundation" {
				imports.NeedsFoundation = true
			} else if typeReferencesFramework(goType, "appkit") && framework != "AppKit" {
				imports.NeedsAppKit = true
			} else if typeReferencesFramework(goType, "quartzcore") && framework != "QuartzCore" {
				imports.NeedsQuartzCore = true
			} else if typeReferencesFramework(goType, "cloudkit") && framework != "CloudKit" {
				imports.NeedsCloudKit = true
			}
		}
		if imports.NeedsFoundation || imports.NeedsAppKit || imports.NeedsQuartzCore || imports.NeedsCloudKit {
			break
		}
	}

	// Check method parameters and return types for Foundation dependencies
	if !imports.NeedsFoundation {
		for _, method := range class.Methods {
			// Check return type
			goReturnType := mapObjCTypeToGo(method.ReturnType, framework)
			if typeReferencesFramework(goReturnType, "foundation") {
				imports.NeedsFoundation = true
				break
			}
			// Check parameter types
			for _, param := range method.Parameters {
				goParamType := mapObjCTypeToGo(param.Type, framework)
				if typeReferencesFramework(goParamType, "foundation") {
					imports.NeedsFoundation = true
					break
				}
			}
			if imports.NeedsFoundation {
				break
			}
		}
	}

	// Check properties for Foundation dependencies
	if !imports.NeedsFoundation {
		for _, prop := range class.Properties {
			goType := mapObjCTypeToGo(prop.Type, framework)
			if strings.HasPrefix(goType, "foundation.") {
				imports.NeedsFoundation = true
				break
			}
		}
	}

	// Check method parameters and return types for UserNotifications dependencies
	if !imports.NeedsUserNotifications {
		for _, method := range class.Methods {
			// Check return type (use Contains to handle slices like []usernotifications.NotificationAction)
			goReturnType := mapObjCTypeToGo(method.ReturnType, framework)
			if strings.Contains(goReturnType, "usernotifications.") {
				imports.NeedsUserNotifications = true
				break
			}
			// Check parameter types
			for _, param := range method.Parameters {
				goParamType := mapObjCTypeToGo(param.Type, framework)
				if strings.Contains(goParamType, "usernotifications.") {
					imports.NeedsUserNotifications = true
					break
				}
			}
			if imports.NeedsUserNotifications {
				break
			}
		}
	}

	// Check properties for UserNotifications dependencies
	if !imports.NeedsUserNotifications {
		for _, prop := range class.Properties {
			goType := mapObjCTypeToGo(prop.Type, framework)
			if strings.Contains(goType, "usernotifications.") {
				imports.NeedsUserNotifications = true
				break
			}
		}
	}

	// Check method parameters and return types for UniformTypeIdentifiers dependencies
	if !imports.NeedsUniformTypeIdentifiers {
		for _, method := range class.Methods {
			// Check return type
			goReturnType := mapObjCTypeToGo(method.ReturnType, framework)
			if strings.Contains(goReturnType, "uniformtypeidentifiers.") {
				imports.NeedsUniformTypeIdentifiers = true
				break
			}
			// Check parameter types
			for _, param := range method.Parameters {
				goParamType := mapObjCTypeToGo(param.Type, framework)
				if strings.Contains(goParamType, "uniformtypeidentifiers.") {
					imports.NeedsUniformTypeIdentifiers = true
					break
				}
			}
			if imports.NeedsUniformTypeIdentifiers {
				break
			}
		}
	}

	// Check properties for UniformTypeIdentifiers dependencies
	if !imports.NeedsUniformTypeIdentifiers {
		for _, prop := range class.Properties {
			goType := mapObjCTypeToGo(prop.Type, framework)
			if strings.Contains(goType, "uniformtypeidentifiers.") {
				imports.NeedsUniformTypeIdentifiers = true
				break
			}
		}
	}

	return imports
}

// getInterfaceParent determines the parent interface for a class interface definition.
// This consolidates the complex interface hierarchy resolution logic from the template.
// Returns the fully-qualified parent interface name (e.g., "foundation.IMutableAttributedString" or "objectivec.IObject").
func getInterfaceParent(class *occ2go.ParsedClass, framework string) string {
	if class == nil {
		return "objectivec.IObject"
	}

	className := class.Name
	structName := classToStructName(className)

	// Special cases for ObjectiveC framework
	if framework == "ObjectiveC" {
		if className == "NSObject" {
			return "objc.IObject"
		}
		return "IObject"
	}

	// Check if class has a superclass (other than NSObject)
	if class.SuperClass != "" && class.SuperClass != "NSObject" {
		superStructName := classToStructName(class.SuperClass)

		// Check for self-referential case (class inherits from itself - edge case)
		if superStructName == structName {
			return "objectivec.IObject"
		}

		// Resolve superclass to its qualified type
		superResolved := resolveType(framework, superStructName)

		// If superclass resolves to unsafe.Pointer, it means the parent class doesn't exist
		// Fall back to objectivec.IObject instead
		if superResolved == "unsafe.Pointer" {
			return "objectivec.IObject"
		}

		// Build interface name based on resolved framework
		if strings.HasPrefix(superResolved, "foundation.") {
			typeName := strings.TrimPrefix(superResolved, "foundation.")
			return "foundation.I" + typeName
		}

		if strings.HasPrefix(superResolved, "quartzcore.") {
			typeName := strings.TrimPrefix(superResolved, "quartzcore.")
			return "quartzcore.I" + typeName
		}

		if strings.HasPrefix(superResolved, "appkit.") {
			typeName := strings.TrimPrefix(superResolved, "appkit.")
			return "appkit.I" + typeName
		}

		if strings.HasPrefix(superResolved, "objectivec.") {
			typeName := strings.TrimPrefix(superResolved, "objectivec.")
			return "objectivec.I" + typeName
		}

		// Local type in same framework
		return "I" + superStructName
	}

	// Default: inherit from objectivec.IObject
	return "objectivec.IObject"
}

// getStructEmbeddedField determines what field should be embedded in the struct definition.
// This consolidates the struct embedding logic from the template.
// Returns the embedded field type name (e.g., "objectivec.Object", "foundation.MutableAttributedString", or "Button").
func getStructEmbeddedField(class *occ2go.ParsedClass, framework string) string {
	if class == nil {
		return "objectivec.Object"
	}

	className := class.Name
	structName := classToStructName(className)

	// Special case: ObjectiveC NSObject uses objc.ID directly
	if framework == "ObjectiveC" && className == "NSObject" {
		return "objc.ID"
	}

	// Check if class has a superclass (other than NSObject)
	if class.SuperClass != "" {
		superStructName := classToStructName(class.SuperClass)

		// Check for self-referential case (class inherits from itself - edge case)
		if superStructName == structName {
			if framework == "ObjectiveC" {
				return "Object"
			}
			return "objectivec.Object"
		}

		// Check if superclass is NSObject or Object
		if class.SuperClass == "NSObject" || superStructName == "Object" {
			if framework == "ObjectiveC" {
				return "Object"
			}
			return "objectivec.Object"
		}

		// Resolve superclass to its qualified type
		superResolved := resolveType(framework, superStructName)

		// If superclass resolves to unsafe.Pointer, it means the parent class doesn't exist
		// Fall back to objectivec.Object instead
		if superResolved == "unsafe.Pointer" {
			if framework == "ObjectiveC" {
				return "Object"
			}
			return "objectivec.Object"
		}

		return superResolved
	}

	// No superclass or superclass is NSObject - use base Object
	if framework == "ObjectiveC" {
		return "Object"
	}
	return "objectivec.Object"
}

// getFromConstructorBody generates the body of the XFrom(ptr unsafe.Pointer) constructor.
// This consolidates the From constructor generation logic from the template.
// Returns the constructor body as a string (without the function signature or surrounding braces).
func getFromConstructorBody(class *occ2go.ParsedClass, framework string) string {
	if class == nil {
		return "return " + classToStructName("") + "{objectivec.Object{objc.ID(ptr)}}"
	}

	className := class.Name
	structName := classToStructName(className)

	// Special case: ObjectiveC NSObject uses objc.ID directly
	if framework == "ObjectiveC" && className == "NSObject" {
		return fmt.Sprintf("return %s{objc.ID(ptr)}", structName)
	}

	// Check if class has a superclass (other than NSObject)
	if class.SuperClass != "" {
		superStructName := classToStructName(class.SuperClass)

		// Check for self-referential case
		if superStructName == structName {
			if framework == "ObjectiveC" {
				return fmt.Sprintf("return %s{Object{objc.ID(ptr)}}", structName)
			}
			return fmt.Sprintf("return %s{objectivec.Object{objc.ID(ptr)}}", structName)
		}

		// Check if superclass is NSObject or Object
		if class.SuperClass == "NSObject" || superStructName == "Object" {
			if framework == "ObjectiveC" {
				return fmt.Sprintf("return %s{Object{objc.ID(ptr)}}", structName)
			}
			return fmt.Sprintf("return %s{objectivec.Object{objc.ID(ptr)}}", structName)
		}

		// Has a non-NSObject superclass - need to construct with named field
		superResolved := resolveType(framework, superStructName)

		// If superclass resolves to unsafe.Pointer, it means the parent class doesn't exist
		// Fall back to objectivec.Object instead
		if superResolved == "unsafe.Pointer" {
			if framework == "ObjectiveC" {
				return fmt.Sprintf("return %s{Object{objc.ID(ptr)}}", structName)
			}
			return fmt.Sprintf("return %s{objectivec.Object{objc.ID(ptr)}}", structName)
		}

		return fmt.Sprintf("return %s{\n\t\t%s: %sFrom(ptr),\n\t}", structName, superStructName, superResolved)
	}

	// No superclass - use base Object
	if framework == "ObjectiveC" {
		return fmt.Sprintf("return %s{Object{objc.ID(ptr)}}", structName)
	}
	return fmt.Sprintf("return %s{objectivec.Object{objc.ID(ptr)}}", structName)
}

// getConstructorBody generates the body of a package-level constructor function.
// It handles the differences between class methods (convenience constructors that return
// autoreleased objects) and instance methods (init* methods that require explicit Autorelease()).
//
// Parameters:
//   - method: The init method to generate constructor body for
//   - structName: The Go struct name (e.g., "Window")
//   - paramNames: Comma-separated parameter names to pass to the Objective-C method
//
// Returns a complete function body including proper memory management.
func getConstructorBody(method *occ2go.ParsedMethod, structName, paramNames string) string {
	selector := method.Selector

	// Build the parameter list for objc.Send
	params := ""
	if paramNames != "" {
		params = ", " + paramNames
	}

	if method.IsClassMethod {
		// Class methods (convenience constructors) return autoreleased objects - don't call Autorelease()
		return fmt.Sprintf("\trv := objc.Send[%s](objc.ID(get%sClass().class), objc.Sel(\"%s\")%s)\n\treturn rv",
			structName, structName, selector, params)
	}

	// Instance methods (init*) require Autorelease() to balance the +1 from alloc
	return fmt.Sprintf("\tinstance := get%sClass().Alloc()\n\trv := objc.Send[%s](instance.ID, objc.Sel(\"%s\")%s)\n\trv.Autorelease()\n\treturn rv",
		structName, structName, selector, params)
}

// buildCrossFrameworkTypeRegistry scans generated frameworks and populates the type registry.
// This allows proper type resolution instead of falling back to unsafe.Pointer.
//
// It scans the output directory for generated frameworks and extracts class names,
// building a map of type name -> framework package name.
//
// Example registry entries:
//   "Window" -> "appkit"
//   "String" -> "foundation"
//   "Layer" -> "quartzcore"
func buildCrossFrameworkTypeRegistry(outputDir string) error {
	// Check if output directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		// Output directory doesn't exist yet, registry will be empty
		return nil
	}

	// Scan all subdirectories (frameworks)
	entries, err := os.ReadDir(outputDir)
	if err != nil {
		return fmt.Errorf("failed to read output directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		frameworkPkg := strings.ToLower(entry.Name())
		frameworkDir := filepath.Join(outputDir, entry.Name())

		// Look for types.gen.go which contains type definitions
		typesFile := filepath.Join(frameworkDir, "types.gen.go")
		if _, err := os.Stat(typesFile); os.IsNotExist(err) {
			continue
		}

		// Parse types.gen.go to extract type names
		data, err := os.ReadFile(typesFile)
		if err != nil {
			continue // Skip on error
		}

		// Extract type definitions (e.g., "type Window struct")
		// Simple regex to match "type TypeName" declarations
		typeRegex := regexp.MustCompile(`(?m)^type\s+([A-Z][A-Za-z0-9_]*)\s+(?:struct|interface|unsafe\.Pointer)`)
		matches := typeRegex.FindAllSubmatch(data, -1)

		for _, match := range matches {
			if len(match) > 1 {
				typeName := string(match[1])
				// Add to registry if not already present (first framework wins)
				if _, exists := crossFrameworkTypeRegistry[typeName]; !exists {
					crossFrameworkTypeRegistry[typeName] = frameworkPkg
				}
			}
		}
	}

	// After scanning all frameworks, override common Foundation types to ensure they're always mapped to foundation
	// This prevents other frameworks (accessibility, authenticationservices, etc.) from claiming Foundation types
	// just because they come first alphabetically
	foundationCoreTypes := []string{
		"NSURL", "URL",
		"NSNumber", "Number",
		"NSString", "String",
		"NSArray", "Array",
		"NSDictionary", "Dictionary",
		"NSData", "Data",
		"NSDate", "Date",
		"NSSet", "Set",
	}
	for _, typeName := range foundationCoreTypes {
		crossFrameworkTypeRegistry[typeName] = "foundation"
	}

	return nil
}

// sortedKeys returns the keys of a map in sorted order, suitable for templates.
// This is useful for generating consistent output across template generations.
func sortedKeys(m map[string]bool) []string {
	if m == nil {
		return []string{}
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// Simple bubble sort for consistency
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if keys[i] > keys[j] {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}

	return keys
}
