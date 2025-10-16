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
	"join":      strings.Join,
	"lower":     strings.ToLower,
	"trimspace": strings.TrimSpace,
	"trimRight": strings.TrimRight,

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

// stripObjCPrefix removes common Objective-C prefixes from a class name
func stripObjCPrefix(className string) string {
	prefixes := []string{"NS", "CG", "CF", "CA", "CI", "CL", "CM", "CV", "CT"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(className, prefix) {
			// Make sure the next character is uppercase (to avoid stripping "NS" from "NSone" for example)
			if len(className) > len(prefix) {
				nextChar := className[len(prefix)]
				if nextChar >= 'A' && nextChar <= 'Z' {
					return className[len(prefix):]
				}
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
