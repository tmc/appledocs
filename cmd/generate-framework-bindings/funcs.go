package main

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/tmc/appledocs/occ2go"
)

// templateFuncs is the FuncMap available to all templates
var templateFuncs = template.FuncMap{
	// String utilities
	"join":        strings.Join,
	"lower":       strings.ToLower,
	"trimspace":   strings.TrimSpace,
	"trimRight":   strings.TrimRight,
	"hasPrefix":   strings.HasPrefix,
	"commentLine": commentLine,
	"dict":        dict,

	// occ2go type mapping
	"mapCTypeToGo": occ2go.MapCTypeToGo,

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
	"isConstructor":            isConstructor,
	"stripNSPrefix":            stripNSPrefix,
	"needsFoundationImport":    needsFoundationImport,
	"needsQuartzCoreImport":    needsQuartzCoreImport,
	"needsCustomImports":       needsCustomImports,
	"getRequiredImports":       getRequiredImports,
	"sortedImportPaths":        sortedImportPaths,
	"prepareClassMethods":         prepareClassMethods,
	"prepareInstanceMethods":      prepareInstanceMethods,
	"prepareInitMethods":          prepareInitMethods,
	"initMethodToConstructorName": initMethodToConstructorName,
	"classHasInit":                classHasInit,
	"sortMethodsByName":           sortMethodsByName,
	"wrapObjCReturn":              wrapObjCReturn,
	"isEssentialSelector":         isEssentialSelector,
	"convertDocURL":               convertDocURL,

	// Property generation helpers
	"propertyToGoName":         propertyToGoName,
	"contains":                 sliceContainsString,
	"capitalize":               capitalizeFirst,

	// Import merging
	"mergeImports": mergeImports,
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
		paramType = occ2go.MapCTypeToGo(paramType, framework)

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
		paramType = occ2go.MapCTypeToGo(paramType, framework)

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
		data.ReturnType = occ2go.MapCTypeToGo(fn.ReturnType, framework)
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
// Strips prefix and appends 'Class'.
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
	return name + "Class"
}

// stripObjCPrefix removes common Objective-C prefixes and invalid identifier characters from a class name
func stripObjCPrefix(className string) string {
	// First strip colons and other invalid identifier characters
	className = strings.ReplaceAll(className, ":", "")

	prefixes := []string{"NS", "CG", "CF", "CA", "CI", "CL", "CM", "CV", "CT"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(className, prefix) {
			// Make sure the next character is uppercase (to avoid stripping "NS" from "NSone" for example)
			if len(className) > len(prefix) {
				nextChar := className[len(prefix)]
				if nextChar >= 'A' && nextChar <= 'Z' {
					name := className[len(prefix):]
					// Check if result is a Go keyword and escape it
					if isGoKeyword(strings.ToLower(name)) {
						return name + "_"
					}
					return name
				}
			}
		}
	}

	// Check if the className (after stripping invalid chars) is a Go keyword
	if isGoKeyword(strings.ToLower(className)) {
		return className + "_"
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
// Examples:
//   NSButton -> button.gen.go
//   NSTableView -> table_view.gen.go
//   NSURLRequest -> url_request.gen.go
func classFileName(className string) string {
	name := stripObjCPrefix(className)
	return toSnakeCase(name) + ".gen.go"
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

// mapObjCTypeToGo maps Objective-C types to Go types for darwinkit style.
// Examples:
//   NSString * -> string
//   id -> objc.Object
//   NSButton * -> Button (interface type in parameters)
//   NSRect -> foundation.Rect
//   NSWindowStyleMask -> WindowStyleMask
func mapObjCTypeToGo(objcType, framework string) string {
	objcType = strings.TrimSpace(objcType)

	// Handle Objective-C generic types (e.g., NSArray<NSString *>)
	// These cannot be directly represented in Go, so map to unsafe.Pointer
	if strings.Contains(objcType, "<") {
		return "unsafe.Pointer"
	}

	// Handle Objective-C blocks (e.g., void (^)(NSModalResponse))
	// Blocks are closures that cannot be easily represented in Go, so map to unsafe.Pointer
	if strings.Contains(objcType, "^") {
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
	case "NSInteger":
		return "int"
	case "NSUInteger":
		return "uint"
	case "CGFloat":
		return "float64"
	case "void":
		return ""
	}

	// Check the type mapping registry first (includes both with and without pointers)
	if goType, found := lookupTypeMapping(objcType, framework); found {
		return goType
	}

	// Handle pointers for types not in the registry
	isPointer := strings.HasSuffix(objcType, "*")
	objcTypeNoPtr := strings.TrimSpace(strings.TrimSuffix(objcType, "*"))

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

// prepareClassMethods filters methods to return only class methods, deduplicated by selector.
// When multiple methods have the same selector, the first one is kept.
func prepareClassMethods(methods []*occ2go.ParsedMethod) []*occ2go.ParsedMethod {
	result := make([]*occ2go.ParsedMethod, 0)
	seen := make(map[string]bool)

	for _, m := range methods {
		if m.IsClassMethod && !seen[m.Selector] {
			result = append(result, m)
			seen[m.Selector] = true
		}
	}
	return result
}

// prepareInstanceMethods filters methods to return only instance methods,
// excluding ALL init methods (which are converted to constructors), deduplicated by selector.
// When multiple methods have the same selector, the first one is kept.
func prepareInstanceMethods(methods []*occ2go.ParsedMethod) []*occ2go.ParsedMethod {
	result := make([]*occ2go.ParsedMethod, 0)
	seen := make(map[string]bool)

	for _, m := range methods {
		if !m.IsClassMethod && !seen[m.Selector] {
			// Skip ALL init methods - they're converted to package-level constructors
			if strings.HasPrefix(m.Selector, "init") {
				continue
			}
			result = append(result, m)
			seen[m.Selector] = true
		}
	}
	return result
}

// prepareInitMethods filters methods to return only init methods (for constructor generation), deduplicated by selector.
// When multiple methods have the same selector, the first one is kept.
func prepareInitMethods(methods []*occ2go.ParsedMethod) []*occ2go.ParsedMethod {
	result := make([]*occ2go.ParsedMethod, 0)
	seen := make(map[string]bool)

	for _, m := range methods {
		if !m.IsClassMethod && strings.HasPrefix(m.Selector, "init") && !seen[m.Selector] {
			result = append(result, m)
			seen[m.Selector] = true
		}
	}
	return result
}

// initMethodToConstructorName converts an init method selector to a constructor function name.
// Examples:
//   "init" -> "New"
//   "initWithFrame:" -> "NewWithFrame"
//   "initWithContentRect:styleMask:backing:defer:" -> "NewWithContentRectStyleMaskBackingDefer"
func initMethodToConstructorName(className, selector string) string {
	structName := classToStructName(className)

	// Special case for plain "init"
	if selector == "init" {
		return "New" + structName
	}

	// Strip "init" prefix
	if strings.HasPrefix(selector, "init") {
		selector = strings.TrimPrefix(selector, "init")
	}

	// Convert selector to Go name (handles colons, capitalization)
	goName := occ2go.SelectorToGoName("init" + selector)

	// Replace "Init" prefix with "New{ClassName}"
	if strings.HasPrefix(goName, "Init") {
		return "New" + structName + strings.TrimPrefix(goName, "Init")
	}

	return "New" + structName + goName
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
	return strings.ToUpper(s[:1]) + s[1:]
}

// needsCustomImports checks if any methods use types that require custom imports
func needsCustomImports(methods []*occ2go.ParsedMethod, framework string) bool {
	return getRequiredImports(methods, framework) != nil
}

// getRequiredImports returns a map of import paths needed for methods.
// For example: {"github.com/progrium/darwinkit/macos/foundation": true}
func getRequiredImports(methods []*occ2go.ParsedMethod, framework string) map[string]bool {
	imports := make(map[string]bool)

	// Check all methods for types that need custom imports
	for _, m := range methods {
		// Check return type
		if m.ReturnType != "" {
			if importPath := getTypeImportPath(m.ReturnType, framework); importPath != "" {
				imports[importPath] = true
			}
		}

		// Check parameters
		for _, p := range m.Parameters {
			if importPath := getTypeImportPath(p.Type, framework); importPath != "" {
				imports[importPath] = true
			}
		}
	}

	if len(imports) == 0 {
		return nil
	}
	return imports
}

// sortedImportPaths returns a sorted slice of import paths for template iteration.
// This makes it easy for templates to range over imports in a consistent order.
func sortedImportPaths(imports map[string]bool) []string {
	if imports == nil || len(imports) == 0 {
		return []string{}
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

	return paths
}

// getClassRequiredImports returns a sorted slice of all required imports for a class (including both methods and properties).
// This is a convenience function for templates to get all imports at once.
func getClassRequiredImports(class interface{}, framework string) []string {
	// This is a bit of a hack, but we need to work with the parsed class data
	// For now, we return an empty slice - this would need proper type handling
	return []string{}
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

// classHasInit checks if a class has any init methods (with selector "init").
// This is used to determine if a test file should be generated.
func classHasInit(methods []*occ2go.ParsedMethod) bool {
	for _, m := range methods {
		if !m.IsClassMethod && m.Selector == "init" {
			return true
		}
	}
	return false
}
