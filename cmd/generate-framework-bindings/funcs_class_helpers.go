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
		// Use unqualified IObject when in the objectivec package
		if framework == "ObjectiveC" {
			return "IObject"
		}
		return "objectivec.IObject"
	}

	className := class.Name
	structName := classToStructName(className)

	// Special cases for ObjectiveC framework
	if framework == "ObjectiveC" {
		if className == "NSObject" {
			return ""
		}
		// When we're IN the objectivec package, use unqualified IObject
		return "IObject"
	}

	// Check if class has a superclass (other than NSObject)
	if class.SuperClass != "" && class.SuperClass != "NSObject" {
		superStructName := classToStructName(class.SuperClass)

		// Check for self-referential case (class inherits from itself - edge case)
		if superStructName == structName {
			if framework == "ObjectiveC" {
				return "IObject"
			}
			return "objectivec.IObject"
		}

		// Resolve superclass to its qualified type
		superResolved := resolveType(framework, superStructName)

		// If superclass resolves to unsafe.Pointer, it means the parent class doesn't exist
		// Fall back to IObject instead
		if superResolved == "unsafe.Pointer" {
			if framework == "ObjectiveC" {
				return "IObject"
			}
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

	// Default: inherit from IObject
	if framework == "ObjectiveC" {
		return "IObject"
	}
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

// iosOnlyMethod returns true if the method is only available on iOS (not macOS).
func iosOnlyMethod(method *occ2go.ParsedMethod) bool {
	if method == nil {
		return false
	}
	return method.Availability.IOSOnly()
}

// iosOnlyProperty returns true if the property is only available on iOS (not macOS).
func iosOnlyProperty(property *occ2go.ParsedProperty) bool {
	if property == nil {
		return false
	}
	return property.Availability.IOSOnly()
}

// classHasIOSOnlyMembers returns true if the class has any iOS-only methods or properties.
// Deprecated: Use hasIOSOnlyMethods instead for template consistency.
func classHasIOSOnlyMembers(class *occ2go.ParsedClass) bool {
	return hasIOSOnlyMethods(class)
}

// hasIOSOnlyMethods returns true if the class has any iOS-only methods or properties.
// This is the preferred template function name.
func hasIOSOnlyMethods(class *occ2go.ParsedClass) bool {
	if class == nil {
		return false
	}

	// Check methods
	for _, method := range class.Methods {
		if iosOnlyMethod(method) {
			return true
		}
	}

	// Check properties
	for _, property := range class.Properties {
		if iosOnlyProperty(property) {
			return true
		}
	}

	return false
}

// embeddedTypeNameshadowsParentMethod checks if an embedded type's name would shadow
// an inherited method from objectivec.IObject. This happens when a class embeds a parent
// type whose name matches a method name (e.g., ClassDescription type shadows ClassDescription() method).
func embeddedTypeNameShadowsParentMethod(class *occ2go.ParsedClass, framework string) bool {
	if class == nil {
		return false
	}

	// Get the embedded type name
	embeddedTypeName := getStructEmbeddedTypeName(class, framework)
	if embeddedTypeName == "" {
		return false
	}

	// Known IObject methods that could be shadowed by type names
	// (method name without parentheses)
	shadowableMethodNames := map[string]bool{
		"ClassDescription": true,
		"ObjectSpecifier":  true,
		// Add more as discovered
	}

	return shadowableMethodNames[embeddedTypeName]
}

// getStructEmbeddedTypeName returns the name of the embedded type without package qualification.
// For example, if embedding "foundation.ClassDescription", returns "ClassDescription".
func getStructEmbeddedTypeName(class *occ2go.ParsedClass, framework string) string {
	if class == nil {
		return ""
	}

	// Get the parent struct name
	if class.SuperClass == "" || class.SuperClass == "NSObject" {
		return "Object"
	}

	return classToStructName(class.SuperClass)
}

// getShadowedMethodSignature returns the method signature that is shadowed by the embedded type name.
// Returns empty string if no shadowing occurs.
func getShadowedMethodSignature(class *occ2go.ParsedClass, framework string) string {
	if !embeddedTypeNameShadowsParentMethod(class, framework) {
		return ""
	}

	embeddedTypeName := getStructEmbeddedTypeName(class, framework)

	// Return the method signature that needs to be forwarded
	// For ClassDescription, the signature is: ClassDescription() IObject
	switch embeddedTypeName {
	case "ClassDescription":
		return "ClassDescription() IObject"
	case "ObjectSpecifier":
		return "ObjectSpecifier() IObject"
	default:
		return ""
	}
}

// getShadowedMethodName returns just the method name (without return type) that is shadowed.
// Returns empty string if no shadowing occurs.
func getShadowedMethodName(class *occ2go.ParsedClass, framework string) string {
	if !embeddedTypeNameShadowsParentMethod(class, framework) {
		return ""
	}

	embeddedTypeName := getStructEmbeddedTypeName(class, framework)

	// Return just the method name for calling
	return embeddedTypeName
}

// methodConflictsWithParent checks if a method would conflict with an inherited method
// from objectivec.IObject. This happens when the parent has a method with the same name
// but potentially different signature.
func methodConflictsWithParent(method *occ2go.ParsedMethod, class *occ2go.ParsedClass, framework string) bool {
	if method == nil || class == nil {
		return false
	}

	// Known objectivec.IObject and parent interface methods that are commonly redeclared
	// Map method name to true if it exists in a parent interface
	objectMethods := map[string]bool{
		"ForwardInvocation":                     true, // IObject
		"AttributeKeys":                         true, // IObject
		"ToManyRelationshipKeys":                true, // IObject
		"ToOneRelationshipKeys":                 true, // IObject
		"Count":                                 true, // IObject (collections)
		"ObjectSpecifier":                       true, // IObject
		"AddObserverSelectorNameObject":         true, // INotificationCenter
		"PostNotificationNameObject":            true, // INotificationCenter
		"PostNotificationNameObjectUserInfo":    true, // INotificationCenter
		"RemoveObserverNameObject":              true, // INotificationCenter
		// Add more as discovered
	}

	methodName := method.Name
	if methodName == "" && method.Selector != "" {
		// Try to derive name from selector
		methodName = occ2go.SelectorToGoName(method.Selector)
	}

	return objectMethods[methodName]
}

// propertyConflictsWithParentMethod checks if a property accessor would conflict
// with a method inherited from objectivec.IObject.
func propertyConflictsWithParentMethod(property *occ2go.ParsedProperty, class *occ2go.ParsedClass, framework string) bool {
	if property == nil || class == nil {
		return false
	}

	// Known objectivec.IObject and parent interface methods/properties
	objectMethods := map[string]bool{
		"AttributeKeys":          true, // IObject
		"ToManyRelationshipKeys": true, // IObject
		"ToOneRelationshipKeys":  true, // IObject
		"Count":                  true, // IObject (collections)
		"ObjectSpecifier":        true, // IObject
		"Attribution":            true, // IURLRequest
		"NetworkServiceType":     true, // IURLRequest
		"Delegate":               true, // IURLSession*
	}

	propertyName := property.Name
	if propertyName == "" {
		return false
	}

	// Check if getter name conflicts
	getterName := occ2go.PropertyToGoName(propertyName)
	return objectMethods[getterName]
}
