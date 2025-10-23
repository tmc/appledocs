package main

import (
	"fmt"
	"strings"

	"github.com/tmc/appledocs/occ2go"
)

// isClassType is now a method on Generator - see Generator.IsClassType()

// getClassImportsOld was moved to funcs_unused.go - see that file for the implementation

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
			return ""
		}
		return "objectivec.IObject"
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
