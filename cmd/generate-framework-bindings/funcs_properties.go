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

// typeToInterfaceType is a template wrapper function that calls Generator.TypeToInterfaceType.
// The Generator is expected to be in the template's root context (.).
// This wrapper extracts the Generator from the template data and delegates to the method.
func typeToInterfaceType(gen interface{}, goType string) string {
	if g, ok := gen.(*Generator); ok {
		return g.TypeToInterfaceType(goType)
	}
	// Fallback for cases where Generator isn't available (shouldn't happen in practice)
	return typeToInterfaceTypeHeuristic(goType)
}

// typeToInterfaceTypeHeuristic is the old heuristic-based implementation, kept as a fallback.
// DEPRECATED: Use Generator.TypeToInterfaceType instead for data-driven type checking.
func typeToInterfaceTypeHeuristic(goType string) string {
	// Handle qualified types (e.g., "foundation.Coder" -> "foundation.ICoder")
	if strings.Contains(goType, ".") {
		parts := strings.SplitN(goType, ".", 2)
		if len(parts) == 2 {
			pkg := parts[0]
			typeName := parts[1]

			// Don't convert runtime types (objc.ID, unsafe.Pointer, etc.)
			if pkg == "objc" || pkg == "unsafe" || pkg == "objectivec" {
				return goType
			}

			// Don't convert CoreGraphics types (structs and refs)
			if strings.HasPrefix(typeName, "CG") {
				return goType
			}

			// Recursively convert the type part
			interfaceType := typeToInterfaceTypeHeuristic(typeName)
			return pkg + "." + interfaceType
		}
		return goType
	}

	// Don't convert primitives, slices, pointers, or special types
	if strings.HasPrefix(goType, "[]") ||
		strings.HasPrefix(goType, "*") ||
		strings.HasPrefix(goType, "map[") ||
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

	// DEPRECATED HEURISTIC: This fallback function is only used when Generator context
	// isn't available. The hardcoded lists have been removed - use Generator.TypeToInterfaceType()
	// instead for data-driven type checking that queries actual Enums, Typedefs, and Classes.

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
