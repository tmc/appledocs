package occ2go

import (
	"strings"
)

// naming.go contains Objective-C to Go naming convention conversions.
// These functions handle the transformation of Objective-C identifiers
// (class names, selectors, property names, etc.) into idiomatic Go names.

// goKeywords contains Go keywords and predeclared identifiers
var goKeywords = map[string]bool{
	// Reserved keywords
	"break": true, "case": true, "chan": true, "const": true, "continue": true,
	"default": true, "defer": true, "else": true, "fallthrough": true, "for": true,
	"func": true, "go": true, "goto": true, "if": true, "import": true,
	"interface": true, "map": true, "package": true, "range": true, "return": true,
	"select": true, "struct": true, "switch": true, "type": true, "var": true,
	// Predeclared identifiers
	"append": true, "bool": true, "byte": true, "cap": true, "close": true,
	"complex": true, "complex64": true, "complex128": true, "copy": true,
	"delete": true, "error": true, "false": true, "float32": true, "float64": true,
	"imag": true, "int": true, "int8": true, "int16": true, "int32": true, "int64": true,
	"iota": true, "len": true, "make": true, "new": true, "nil": true,
	"panic": true, "print": true, "println": true, "real": true, "recover": true,
	"rune": true, "string": true, "true": true, "uint": true, "uint8": true,
	"uint16": true, "uint32": true, "uint64": true, "uintptr": true,
}

// isGoKeyword checks if a string is a Go reserved keyword or predeclared identifier
func isGoKeyword(name string) bool {
	return goKeywords[name]
}

// ClassToInterfaceName converts an Objective-C class name to a Go interface name.
// Strips the NS/CG/CF prefix and adds an "I" prefix.
// Examples:
//
//	NSButton -> IButton
//	NSView -> IView
//	CGContext -> IContext
func ClassToInterfaceName(className string) string {
	if className == "" {
		return ""
	}

	// Strip common prefixes
	name := StripObjCPrefix(className)
	return "I" + name
}

// ClassToStructName converts an Objective-C class name to a Go struct name.
// Strips the NS/CG/CF prefix and capitalizes if it's a Go keyword.
// Examples:
//
//	NSButton -> Button
//	NSView -> View
//	CGContext -> Context
//	map -> Map (Go keyword, capitalized)
func ClassToStructName(className string) string {
	if className == "" {
		return ""
	}

	name := StripObjCPrefix(className)

	// Special case: "isa" should be "ISA" (all caps acronym)
	if name == "isa" {
		return "ISA"
	}

	// If the name is a Go keyword, capitalize it
	if isGoKeyword(name) {
		return strings.ToUpper(name[:1]) + name[1:]
	}

	return name
}

// ClassToVarName converts an Objective-C class name to a Go variable name for the class.
// Strips prefix and returns capitalized with 'Class' suffix to make it exported.
// Examples:
//
//	NSButton -> ButtonClass
//	NSView -> ViewClass
//	CGContext -> ContextClass
func ClassToVarName(className string) string {
	if className == "" {
		return ""
	}

	name := StripObjCPrefix(className)
	// Make it capitalized to be public (exported from the package)
	return strings.ToUpper(name[:1]) + name[1:] + "Class"
}

// StripObjCPrefix removes Objective-C prefixes algorithmically and invalid identifier characters from a class name
func StripObjCPrefix(className string) string {
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

// StripNSPrefix is an alias for StripObjCPrefix for backward compatibility.
func StripNSPrefix(className string) string {
	return StripObjCPrefix(className)
}

// MethodToGoName converts an Objective-C method selector to a Go method name.
// Examples:
//
//	setTitle: -> SetTitle
//	buttonWithTitle:image: -> ButtonWithTitleImage
//	initWithFrame: -> InitWithFrame
//	isEnabled -> IsEnabled
func MethodToGoName(selector string) string {
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

// InitMethodToConstructorName converts an Objective-C init method selector to a Go constructor name.
// Handles both traditional init methods (initWithFrame:) and class factory methods (buttonWithTitle:).
// Examples:
//
//	NSButton, "init" -> NewButton
//	NSButton, "initWithFrame:" -> NewButtonWithFrame
//	NSButton, "buttonWithTitle:target:action:" -> NewButtonWithTitleTargetAction
//	CIBlendKernel, "kernelWithString:" -> NewBlendKernelWithString (avoids "KernelKernel")
func InitMethodToConstructorName(className, selector string) string {
	structName := ClassToStructName(className)

	// Special case for plain "init"
	if selector == "init" {
		return "New" + structName
	}

	// Check if this is a traditional init method (starts with "init")
	if strings.HasPrefix(selector, "init") {
		// Strip "init" prefix
		selector = strings.TrimPrefix(selector, "init")

		// Convert selector to Go name (handles colons, capitalization)
		goName := SelectorToGoName("init" + selector)

		// Replace "Init" prefix with "New{ClassName}"
		if strings.HasPrefix(goName, "Init") {
			return "New" + structName + strings.TrimPrefix(goName, "Init")
		}

		return "New" + structName + goName
	}

	// This is a class factory method (e.g., buttonWithTitle:target:action:, kernelWithString:)
	// Convert the entire selector to Go name
	goName := SelectorToGoName(selector)

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

// PropertyToGoName converts an Objective-C property name to a Go method name.
// Capitalizes the first letter for export and handles special cases.
// Examples:
//
//	title -> Title
//	backgroundColor -> BackgroundColor
//	object -> GetObject (special case to avoid conflict with embedded Object type)
func PropertyToGoName(propName string) string {
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

// CapitalizeFirst capitalizes the first letter of a string for exported identifiers.
// Capitalizing makes any Go keyword a valid exported identifier.
// Examples:
//
//	title -> Title
//	backgroundColor -> BackgroundColor
//	false -> False (valid once capitalized)
//	type -> Type (valid once capitalized)
func CapitalizeFirst(s string) string {
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

// ToSnakeCase converts CamelCase to snake_case and strips invalid filename characters
func ToSnakeCase(s string) string {
	// First strip colons, backticks, and other invalid filename characters
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, "`", "")

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

// ClassFileName converts a class name to a file name (snake_case).
// Uses the FULL class name including ObjC prefix to prevent duplicate files.
// Examples:
//
//	NSButton -> ns_button.gen.go
//	NSTableView -> ns_table_view.gen.go
//	ICCameraDevice -> ic_camera_device.gen.go
func ClassFileName(className string) string {
	// Use the full class name (e.g., ICCameraDevice -> ic_camera_device.gen.go)
	// NOT stripped prefix (CameraDevice -> camera_device.gen.go)
	// This prevents duplicate file generation - see appledocs-227
	return ToSnakeCase(className) + ".gen.go"
}

// ProtocolFileName converts a protocol name to a file name (snake_case).
// Examples:
//
//	NSCopying -> copying_protocol.gen.go
//	NSTableViewDataSource -> table_view_data_source_protocol.gen.go
func ProtocolFileName(protocolName string) string {
	name := StripObjCPrefix(protocolName)
	return ToSnakeCase(name) + "_protocol.gen.go"
}

// ClassTestFileName converts a class name to a test file name (snake_case).
// Uses the FULL class name including ObjC prefix to match ClassFileName convention.
// Examples:
//
//	NSButton -> ns_button.gen_test.go
//	NSTableView -> ns_table_view.gen_test.go
//	NSURLRequest -> ns_url_request.gen_test.go
//	ICCameraDevice -> ic_camera_device.gen_test.go
func ClassTestFileName(className string) string {
	// Use full class name to match ClassFileName (see appledocs-549)
	return ToSnakeCase(className) + ".gen_test.go"
}

// ReceiverName generates a short receiver name for methods.
// Examples:
//
//	Button, false -> b_
//	Button, true -> bc
func ReceiverName(className string, isClass bool) string {
	name := StripObjCPrefix(className)
	if len(name) == 0 {
		return "x"
	}
	short := strings.ToLower(string(name[0]))
	if isClass {
		return short + "c"
	}
	return short + "_"
}

// CleanConstantName removes common constant prefixes and applies naming conventions.
// Strips prefixes like kCI, kCG, kCA, kCF, kNS, k and ensures the result is exported.
// Examples:
//
//	kCGBlendModeNormal -> BlendModeNormal
//	kCAFillModeForwards -> FillModeForwards
//	kCIAttributeFilterName -> AttributeFilterName
func CleanConstantName(name string) string {
	// Remove common constant prefixes
	prefixes := []string{"kCI", "kCG", "kCA", "kCF", "kNS", "k"}

	for _, prefix := range prefixes {
		if strings.HasPrefix(name, prefix) {
			name = strings.TrimPrefix(name, prefix)
			break
		}
	}

	// Also strip framework-specific object prefixes from the result
	name = StripObjCPrefix(name)

	// Ensure first letter is uppercase for exported Go identifier
	if len(name) > 0 && name[0] >= 'a' && name[0] <= 'z' {
		name = strings.ToUpper(name[:1]) + name[1:]
	}

	return name
}

// PropertyConflictsWithParent checks if a property accessor would conflict with an embedded parent field.
// Returns true if the capitalized property name matches the Go struct name of the superclass.
// This prevents method name conflicts in the generated code.
// Examples:
//
//	NSView has a "window" property, NSWindow is the parent -> Window() conflicts with embedded Window field
//	NSButton has a "view" property, NSView is the parent -> View() conflicts with embedded View field
func PropertyConflictsWithParent(className, superClass, propertyName string) bool {
	if className == "" || superClass == "" || propertyName == "" {
		return false
	}

	// Capitalize the property name to get the Go method name
	capitalizedProp := CapitalizeFirst(propertyName)

	// Get the Go struct name for the superclass (strips NS/CG/CF prefix)
	superStructName := ClassToStructName(superClass)

	// Check if they match
	return capitalizedProp == superStructName
}

// CommentLine formats a string for use in a Go comment by collapsing whitespace.
// Replaces newlines with spaces and collapses multiple spaces into single spaces.
func CommentLine(s string) string {
	if s == "" {
		return ""
	}
	// Replace newlines and tabs with spaces
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", " ")
	s = strings.ReplaceAll(s, "\t", " ")

	// Collapse multiple spaces
	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}

	return strings.TrimSpace(s)
}

// DisambiguateMethodName generates a unique Go method name for an Objective-C method
// by appending parameter labels from the selector. This matches Swift's approach.
// Examples:
//
//	imageByInsertingIntermediate -> ImageByInsertingIntermediate (no params, no change)
//	imageByInsertingIntermediate: -> ImageByInsertingIntermediateWithCache (1 param named "cache")
//	setTitle:forState: -> SetTitleForState (already unique from selector parts)
func DisambiguateMethodName(method *ParsedMethod) string {
	selector := method.Selector

	// If there are no parameters, just use the standard conversion
	if len(method.Parameters) == 0 {
		return SelectorToGoName(selector)
	}

	// Split selector by colons to get parameter labels
	parts := strings.Split(selector, ":")

	// Remove empty parts from the split (happens when selector ends with :)
	var nonEmptyParts []string
	for _, p := range parts {
		if p != "" {
			nonEmptyParts = append(nonEmptyParts, p)
		}
	}
	parts = nonEmptyParts

	// If we have parameter labels, build the disambiguated name
	// For single-parameter methods, append "With" + capitalized parameter name
	if len(parts) == 1 && len(method.Parameters) == 1 {
		baseName := SelectorToGoName(parts[0])
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
	return SelectorToGoName(selector)
}

// MethodGoName generates the final Go method name for an Objective-C method,
// with disambiguation if multiple methods would map to the same name.
// Examples:
//
//	setTitle: (unique) -> SetTitle
//	imageByInsertingIntermediate: (conflicts) -> ImageByInsertingIntermediateWithCache
func MethodGoName(method *ParsedMethod, allMethods []*ParsedMethod) string {
	// Generate the standard Go name
	goName := SelectorToGoName(method.Selector)

	// Count how many methods map to the same Go name
	count := 0
	for _, m := range allMethods {
		if SelectorToGoName(m.Selector) == goName {
			count++
		}
	}

	// If there's only one method with this name, no disambiguation needed
	if count == 1 {
		return goName
	}

	// Multiple methods map to the same Go name - disambiguate
	return DisambiguateMethodName(method)
}
