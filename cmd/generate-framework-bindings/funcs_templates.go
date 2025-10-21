package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tmc/appledocs/occ2go"
)

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

// extractMethodName extracts the method name by removing the type prefix.
// Examples:
//
//	CGContextSetFillColor → SetFillColor
//	CGPathAddRect → AddRect
func extractMethodName(funcName, framework, typeName string) string {
	prefix := GetFrameworkPrefix(framework)
	if prefix == "" {
		return funcName
	}

	// Remove prefix + type name
	fullPrefix := prefix + typeName
	if strings.HasPrefix(funcName, fullPrefix) {
		return strings.TrimPrefix(funcName, fullPrefix)
	}

	return funcName
}

// extractTypeName extracts the type name from a function name.
// Examples:
//
//	CGContextSetFillColor → Context
//	CGPathAddRect → Path
//	CGColorCreate → Color
func extractTypeName(funcName, framework string) string {
	// Handle framework-specific prefixes
	prefix := GetFrameworkPrefix(framework)
	if prefix == "" {
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

// formatMethodParams formats method parameters for Go function signature.
// Returns: "title string, target objectivec.IObject, action objc.Selector"
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
		if os.Getenv("DEBUG_TYPEMAP") == "1" && strings.Contains(p.Type, "CellAttribute") {
			fmt.Fprintf(os.Stderr, "DEBUG formatMethodParams: p.Type=%s framework=%s\n", p.Type, framework)
		}
		goType := mapObjCTypeToGo(p.Type, framework)
		if os.Getenv("DEBUG_TYPEMAP") == "1" && strings.Contains(p.Type, "CellAttribute") {
			fmt.Fprintf(os.Stderr, "DEBUG formatMethodParams: after mapObjCTypeToGo goType=%s\n", goType)
		}

		// Convert objc.ID to objectivec.IObject for better type safety
		// This allows users to pass any Objective-C object wrapper instead of raw objc.ID
		if goType == "objc.ID" {
			goType = "objectivec.IObject"
		} else {
			// For other class types, use interface types
			goType = typeToInterfaceType(goType)
		}

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
	// Handle array types specially: []Foo{} should become []packageName.Foo{}, not packageName.[]Foo{}
	if strings.HasPrefix(testValue, "[]") {
		// Extract the element type from []Type{}
		elementType := strings.TrimPrefix(testValue, "[]")
		elementType = strings.TrimSuffix(elementType, "{}")
		return "[]" + packageName + "." + elementType + "{}"
	}
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

// isEssentialSelector checks if a selector is essential for object lifecycle.
// Delegates to occ2go.IsEssentialSelector.
func isEssentialSelector(selector string) bool {
	return occ2go.IsEssentialSelector(selector)
}

// convertDocURL converts Apple's doc:// scheme URLs to https:// URLs.
// If the URL doesn't start with "doc://", it returns it unchanged.
// Examples:
//
//	doc://com.apple.foundation/documentation/Foundation/NSString -> https://developer.apple.com/documentation/foundation/nsstring
//	https://developer.apple.com/... -> https://developer.apple.com/... (unchanged)
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
//   - Any instance init method (selector starting with "init")
//   - Any class method marked as an initializer in docs (IsInitializer = true)
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
