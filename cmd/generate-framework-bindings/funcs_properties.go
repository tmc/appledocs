package main

import (
	"strings"
)

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

// typeToInterfaceType converts a concrete type to its interface type for setter parameters.
// For example: "Image" becomes "IImage", "Window" becomes "IWindow".
// Types that don't have interfaces (primitives, slices, enums, typedefs, etc.) are returned unchanged.
func typeToInterfaceType(goType string) string {
	// Don't convert primitives, slices, pointers, or special types
	if strings.HasPrefix(goType, "[]") ||
		strings.HasPrefix(goType, "*") ||
		strings.HasPrefix(goType, "map[") ||
		strings.Contains(goType, ".") || // qualified types like "objc.ID"
		goType == "string" ||
		goType == "int" ||
		goType == "int64" ||
		goType == "uint" ||
		goType == "uint64" ||
		goType == "float32" ||
		goType == "float64" ||
		goType == "bool" ||
		goType == "unsafe.Pointer" ||
		strings.HasPrefix(goType, "CG") || // CoreGraphics types (structs and refs)
		strings.HasPrefix(goType, "NS") && (strings.HasSuffix(goType, "Integer") || strings.HasSuffix(goType, "UInteger")) {
		return goType
	}

	// Don't convert enum-like types (these are typically uint-based type aliases)
	// Common patterns for enums: *Position, *Scaling, *Flags, *Options, *Mask, *State, *Style, *Type, *Mode
	enumSuffixes := []string{
		"Position", "Scaling", "Flags", "Options", "Mask", "State", "Style",
		"Type", "Mode", "Direction", "Alignment", "Format", "Status", "Kind",
		"Level", "Priority", "Policy", "Strategy", "Behavior", "Attribute",
		"Orientation", "Gamut",
	}
	for _, suffix := range enumSuffixes {
		if strings.HasSuffix(goType, suffix) {
			return goType
		}
	}

	// If it already starts with I and next char is uppercase, it's already an interface
	if strings.HasPrefix(goType, "I") && len(goType) > 1 && goType[1] >= 'A' && goType[1] <= 'Z' {
		return goType
	}

	// Convert to interface type: "Image" -> "IImage"
	// This works for class types like Image, Window, View, etc.
	// If the type still has an ObjC prefix (NS, CG, CA), strip it first
	// so we get "IAccessibilityElement" not "INSAccessibilityElement"
	interfaceType := goType
	if strings.HasPrefix(goType, "NS") && len(goType) > 2 && goType[2] >= 'A' && goType[2] <= 'Z' {
		interfaceType = goType[2:]
	} else if strings.HasPrefix(goType, "CG") && len(goType) > 2 && goType[2] >= 'A' && goType[2] <= 'Z' {
		interfaceType = goType[2:]
	} else if strings.HasPrefix(goType, "CA") && len(goType) > 2 && goType[2] >= 'A' && goType[2] <= 'Z' {
		interfaceType = goType[2:]
	}

	return "I" + interfaceType
}
