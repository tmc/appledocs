package main

import (
	"strings"

	"github.com/tmc/appledocs/occ2go"
)

// funcs_naming.go contains name conversions and identifier formatting for Go code generation.
// This includes conversions from Objective-C naming conventions (NSButton, setTitle:, etc.)
// to idiomatic Go naming conventions (Button, SetTitle, etc.).

// classToInterfaceName converts an Objective-C class name to a Go interface name.
// Strips the NS/CG/CF prefix and adds an "I" prefix.
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
		"SM", // ServiceManagement framework
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

// classFileName converts a class name to a file name (snake_case).
// Uses the FULL class name including ObjC prefix to prevent duplicate files.
// Examples:
//
//	NSButton -> ns_button.gen.go
//	NSTableView -> ns_table_view.gen.go
//	ICCameraDevice -> ic_camera_device.gen.go
func classFileName(className string) string {
	// Use the full class name (e.g., ICCameraDevice -> ic_camera_device.gen.go)
	// NOT stripped prefix (CameraDevice -> camera_device.gen.go)
	// This prevents duplicate file generation - see appledocs-227
	return toSnakeCase(className) + ".gen.go"
}

// protocolFileName converts a protocol name to a file name (snake_case).
// Examples:
//
//	NSCopying -> copying_protocol.gen.go
//	NSTableViewDataSource -> table_view_data_source_protocol.gen.go
func protocolFileName(protocolName string) string {
	name := stripObjCPrefix(protocolName)
	return toSnakeCase(name) + "_protocol.gen.go"
}

// classTestFileName converts a class name to a test file name (snake_case).
// Examples:
//
//	NSButton -> button.gen_test.go
//	NSTableView -> table_view.gen_test.go
//	NSURLRequest -> url_request.gen_test.go
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
//
//	Button, false -> b_
//	Button, true -> bc
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
//
//	imageByInsertingIntermediate -> ImageByInsertingIntermediate (no params, no change)
//	imageByInsertingIntermediate: -> ImageByInsertingIntermediateWithCache (1 param named "cache")
//	setTitle:forState: -> SetTitleForState (already unique from selector parts)
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

// stripNSPrefix is an alias for stripObjCPrefix for backward compatibility.
func stripNSPrefix(className string) string {
	return stripObjCPrefix(className)
}

// initMethodToConstructorName converts an Objective-C init method selector to a Go constructor name.
// Handles both traditional init methods (initWithFrame:) and class factory methods (buttonWithTitle:).
// Examples:
//
//	NSButton, "init" -> NewButton
//	NSButton, "initWithFrame:" -> NewButtonWithFrame
//	NSButton, "buttonWithTitle:target:action:" -> NewButtonWithTitleTargetAction
//	CIBlendKernel, "kernelWithString:" -> NewBlendKernelWithString (avoids "KernelKernel")
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

// propertyToGoName converts an Objective-C property name to a Go method name.
// Capitalizes the first letter for export and handles special cases.
// Examples:
//
//	title -> Title
//	backgroundColor -> BackgroundColor
//	object -> GetObject (special case to avoid conflict with embedded Object type)
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

// capitalizeFirst capitalizes the first letter of a string for exported identifiers.
// Capitalizing makes any Go keyword a valid exported identifier.
// Examples:
//
//	title -> Title
//	backgroundColor -> BackgroundColor
//	false -> False (valid once capitalized)
//	type -> Type (valid once capitalized)
func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	// Special case: "object" property conflicts with embedded Object type
	if strings.ToLower(s) == "object" {
		return "GetObject"
	}

	// Capitalizing makes any keyword a valid exported identifier
	// true -> True, false -> False, type -> Type, etc. are all valid
	return strings.ToUpper(s[:1]) + s[1:]
}

// commentLine formats a string for use in a Go comment by collapsing whitespace.
// Replaces newlines with spaces and collapses multiple spaces into single spaces.
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

// methodGoName generates the final Go method name for an Objective-C method,
// with disambiguation if multiple methods would map to the same name.
// Examples:
//
//	setTitle: (unique) -> SetTitle
//	imageByInsertingIntermediate: (conflicts) -> ImageByInsertingIntermediateWithCache
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

// cleanConstantName removes common constant prefixes and applies naming conventions.
// Strips prefixes like kCI, kCG, kCA, kCF, kNS, k and ensures the result is exported.
// Examples:
//
//	kCGBlendModeNormal -> BlendModeNormal
//	kCAFillModeForwards -> FillModeForwards
//	kCIAttributeFilterName -> AttributeFilterName
func cleanConstantName(name string) string {
	// Remove common constant prefixes
	prefixes := []string{"kCI", "kCG", "kCA", "kCF", "kNS", "k"}

	for _, prefix := range prefixes {
		if strings.HasPrefix(name, prefix) {
			name = strings.TrimPrefix(name, prefix)
			break
		}
	}

	// Also strip framework-specific object prefixes from the result
	name = stripObjCPrefix(name)

	// Ensure first letter is uppercase for exported Go identifier
	if len(name) > 0 && name[0] >= 'a' && name[0] <= 'z' {
		name = strings.ToUpper(name[:1]) + name[1:]
	}

	return name
}

// propertyConflictsWithParent checks if a property accessor would conflict with an embedded parent field.
// Returns true if the capitalized property name matches the Go struct name of the superclass.
// This prevents method name conflicts in the generated code.
// Examples:
//
//	NSView has a "window" property, NSWindow is the parent -> Window() conflicts with embedded Window field
//	NSButton has a "view" property, NSView is the parent -> View() conflicts with embedded View field
func propertyConflictsWithParent(className, superClass, propertyName string) bool {
	if superClass == "" || propertyName == "" {
		return false
	}

	// Capitalize the property name to get the Go method name
	capitalizedProp := capitalizeFirst(propertyName)

	// Get the Go struct name for the superclass (strips NS/CG/CF prefix)
	superStructName := classToStructName(superClass)

	// Check if they match
	return capitalizedProp == superStructName
}
