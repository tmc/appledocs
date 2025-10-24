package main

import (
	"strings"

	"github.com/tmc/appledocs/occ2go"
)

// funcs_naming.go contains name conversions and identifier formatting for Go code generation.
// This file provides template-facing wrappers that delegate to occ2go.naming.go for the actual logic.
// All naming convention logic has been moved to the occ2go package for reusability.

// classToInterfaceName converts an Objective-C class name to a Go interface name.
// Delegates to occ2go.ClassToInterfaceName.
func classToInterfaceName(className string) string {
	return occ2go.ClassToInterfaceName(className)
}

// classToStructName converts an Objective-C class name to a Go struct name.
// Delegates to occ2go.ClassToStructName.
func classToStructName(className string) string {
	return occ2go.ClassToStructName(className)
}

// classToVarName converts an Objective-C class name to a Go variable name for the class.
// Delegates to occ2go.ClassToVarName.
func classToVarName(className string) string {
	return occ2go.ClassToVarName(className)
}

// stripObjCPrefix removes Objective-C prefixes algorithmically.
// Delegates to occ2go.StripObjCPrefix.
func stripObjCPrefix(className string) string {
	return occ2go.StripObjCPrefix(className)
}

// MethodInfo represents parsed information about an Objective-C method
type MethodInfo struct {
	Selector   string
	IsInstance bool
	Parameters []occ2go.Parameter
	ReturnType string
}

// methodToGoName converts an Objective-C method selector to a Go method name.
// Delegates to occ2go.MethodToGoName.
func methodToGoName(method MethodInfo) string {
	return occ2go.MethodToGoName(method.Selector)
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
// Delegates to occ2go.ClassFileName.
func classFileName(className string) string {
	return occ2go.ClassFileName(className)
}

// classIOSFileName converts a class name to an iOS-specific file name (snake_case with _ios suffix).
func classIOSFileName(className string) string {
	base := occ2go.ClassFileName(className)
	// Replace .gen.go with _ios.gen.go
	return base[:len(base)-len(".gen.go")] + "_ios.gen.go"
}

// protocolFileName converts a protocol name to a file name (snake_case).
// Delegates to occ2go.ProtocolFileName.
func protocolFileName(protocolName string) string {
	return occ2go.ProtocolFileName(protocolName)
}

// classTestFileName converts a class name to a test file name (snake_case).
// Delegates to occ2go.ClassTestFileName.
func classTestFileName(className string) string {
	return occ2go.ClassTestFileName(className)
}

// receiverName generates a short receiver name for methods.
// Delegates to occ2go.ReceiverName.
func receiverName(className string, isClass bool) string {
	return occ2go.ReceiverName(className, isClass)
}

// selectorToGoName converts an Objective-C selector to Go name.
// This wraps the occ2go.SelectorToGoName function.
func selectorToGoName(selector string) string {
	return occ2go.SelectorToGoName(selector)
}

// disambiguateMethodName generates a unique Go method name for an Objective-C method.
// Delegates to occ2go.DisambiguateMethodName.
func disambiguateMethodName(method *occ2go.ParsedMethod) string {
	return occ2go.DisambiguateMethodName(method)
}

// stripNSPrefix is an alias for stripObjCPrefix for backward compatibility.
// Delegates to occ2go.StripNSPrefix.
func stripNSPrefix(className string) string {
	return occ2go.StripNSPrefix(className)
}

// initMethodToConstructorName converts an Objective-C init method selector to a Go constructor name.
// Delegates to occ2go.InitMethodToConstructorName.
func initMethodToConstructorName(className, selector string) string {
	return occ2go.InitMethodToConstructorName(className, selector)
}

// propertyToGoName converts an Objective-C property name to a Go method name.
// Delegates to occ2go.PropertyToGoName.
func propertyToGoName(propName string) string {
	return occ2go.PropertyToGoName(propName)
}

// capitalizeFirst capitalizes the first letter of a string for exported identifiers.
// Delegates to occ2go.CapitalizeFirst.
func capitalizeFirst(s string) string {
	return occ2go.CapitalizeFirst(s)
}

// commentLine formats a string for use in a Go comment by collapsing whitespace.
// Delegates to occ2go.CommentLine.
func commentLine(s string) string {
	return occ2go.CommentLine(s)
}

// methodGoName generates the final Go method name for an Objective-C method,
// with disambiguation if multiple methods would map to the same name.
// Delegates to occ2go.MethodGoName.
func methodGoName(method *occ2go.ParsedMethod, allMethods []*occ2go.ParsedMethod) string {
	return occ2go.MethodGoName(method, allMethods)
}

// cleanConstantName removes common constant prefixes and applies naming conventions.
// Delegates to occ2go.CleanConstantName.
func cleanConstantName(name string) string {
	return occ2go.CleanConstantName(name)
}

// propertyConflictsWithParent checks if a property accessor would conflict with an embedded parent field.
// Delegates to occ2go.PropertyConflictsWithParent.
func propertyConflictsWithParent(className, superClass, propertyName string) bool {
	return occ2go.PropertyConflictsWithParent(className, superClass, propertyName)
}
