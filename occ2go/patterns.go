package occ2go

import (
	"strings"
)

// patterns.go contains utilities for detecting Objective-C method patterns.
// These functions identify constructors, memory management methods, and other
// special method categories based on selector names and conventions.

// IsConstructorSelector checks if a selector is a constructor/initializer.
// Constructor selectors include init methods and common factory methods.
// Examples:
//
//	init -> true
//	initWithFrame: -> true
//	alloc -> true
//	new -> true
//	buttonWithTitle: -> false (would need class context to determine)
func IsConstructorSelector(selector string) bool {
	// Instance init methods
	if strings.HasPrefix(selector, "init") {
		return true
	}

	// Essential allocation/creation methods
	essentialConstructors := []string{
		"alloc",
		"allocWithZone:",
		"new",
	}

	for _, essential := range essentialConstructors {
		if selector == essential {
			return true
		}
	}

	return false
}

// IsFactoryMethodForClass checks if a class method is a factory constructor.
// Factory methods are class methods that return instances of the class.
// Examples:
//
//	buttonWithTitle: (for NSButton) -> true
//	imageNamed: (for NSImage) -> true
//	alloc -> true
//	description -> false
func IsFactoryMethodForClass(selector string, className string, isClassMethod bool) bool {
	// Must be a class method
	if !isClassMethod {
		return false
	}

	// Check if selector starts with the class name (without prefix)
	selectorLower := strings.ToLower(selector)
	classNameLower := strings.ToLower(StripObjCPrefix(className))

	if strings.HasPrefix(selectorLower, classNameLower) {
		return true
	}

	// Common factory method prefixes
	factoryPrefixes := []string{"new", "create", "make", "alloc"}
	for _, prefix := range factoryPrefixes {
		if strings.HasPrefix(selectorLower, prefix) {
			return true
		}
	}

	return false
}

// IsEssentialSelector checks if a selector is essential for object lifecycle.
// Essential selectors include allocation, initialization, and memory management.
// Examples:
//
//	alloc -> true
//	init -> true
//	copy -> true
//	dealloc -> false (not exposed in Go bindings)
func IsEssentialSelector(selector string) bool {
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

// IsInheritedFromNSObject checks if a selector is likely inherited from NSObject.
// These methods should not be redeclared in child interfaces to avoid duplication.
// Examples:
//
//	retain -> true
//	isEqual: -> true
//	description -> true
//	customMethod -> false
func IsInheritedFromNSObject(selector string) bool {
	// Common NSObject instance methods that should not be redeclared
	nsobjectMethods := map[string]bool{
		// Memory management
		"retain":      true,
		"release":     true,
		"autorelease": true,
		"retainCount": true,

		// Object identity and comparison
		"isEqual:":            true,
		"isEqualTo:":          true,
		"hash":                true,
		"isKindOfClass:":      true,
		"isMemberOfClass:":    true,
		"conformsToProtocol:": true,
		"respondsToSelector:": true,

		// Description and debugging
		"description":      true,
		"debugDescription": true,
		"className":        true,
		"superclass":       true,

		// KVC (Key-Value Coding)
		"valueForKey:":                    true,
		"setValue:forKey:":                true,
		"valueForKeyPath:":                true,
		"setValue:forKeyPath:":            true,
		"valueForUndefinedKey:":           true,
		"setValue:forUndefinedKey:":       true,
		"setNilValueForKey:":              true,
		"dictionaryWithValuesForKeys:":    true,
		"setValuesForKeysWithDictionary:": true,
		"valuesForKeys:":                  true,
		"takeValuesFromDictionary:":       true,

		// KVO (Key-Value Observing)
		"addObserver:forKeyPath:options:context:":    true,
		"removeObserver:forKeyPath:":                 true,
		"removeObserver:forKeyPath:context:":         true,
		"willChangeValueForKey:":                     true,
		"didChangeValueForKey:":                      true,
		"willChange:valuesAtIndexes:forKey:":         true,
		"didChange:valuesAtIndexes:forKey:":          true,
		"observeValueForKeyPath:ofObject:change:context:": true,

		// NSCoding
		"encodeWithCoder:":    true,
		"initWithCoder:":      true,
		"awakeAfterUsingCoder:": true,

		// NSCopying
		"copy":           true,
		"copyWithZone:":  true,
		"mutableCopy":    true,
		"mutableCopyWithZone:": true,

		// Object lifecycle
		"performSelector:":                   true,
		"performSelector:withObject:":        true,
		"performSelector:withObject:withObject:": true,
		"performSelectorOnMainThread:withObject:waitUntilDone:": true,
		"performSelectorInBackground:withObject:": true,

		// Class methods
		"class":       true,
		"alloc":       true,
		"new":         true,
		"initialize":  true,
		"load":        true,
	}

	return nsobjectMethods[selector]
}

// IsMemoryManagementSelector checks if a selector is related to memory management.
// These methods handle retain/release/autorelease which are typically not exposed in Go.
// Examples:
//
//	retain -> true
//	release -> true
//	autorelease -> true
//	dealloc -> true
//	init -> false (initialization, not memory management)
func IsMemoryManagementSelector(selector string) bool {
	memoryMethods := map[string]bool{
		"retain":      true,
		"release":     true,
		"autorelease": true,
		"dealloc":     true,
		"retainCount": true,
	}

	return memoryMethods[selector]
}

// IsPropertyAccessor checks if a selector is likely a property accessor.
// Property accessors follow the pattern "propertyName" (getter) or "setPropertyName:" (setter).
// This is a heuristic check based on naming conventions.
// Examples:
//
//	title -> true (likely getter)
//	setTitle: -> true (likely setter)
//	buttonWithTitle: -> false (not a property accessor)
func IsPropertyAccessor(selector string) bool {
	// Setter pattern: set[PropertyName]:
	if strings.HasPrefix(selector, "set") && strings.HasSuffix(selector, ":") {
		// Check if it has exactly one parameter (setters have one parameter)
		colonCount := strings.Count(selector, ":")
		return colonCount == 1
	}

	// Getter pattern: simple selector with no colons
	// Note: This is a heuristic - not all parameterless methods are getters
	return !strings.Contains(selector, ":")
}

// IsBooleanGetter checks if a selector is a boolean getter (is/has/can pattern).
// Examples:
//
//	isEnabled -> true
//	hasChildren -> true
//	canBecomeKey -> true
//	title -> false
func IsBooleanGetter(selector string) bool {
	prefixes := []string{"is", "has", "can", "should", "will", "did"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(selector, prefix) && len(selector) > len(prefix) {
			// Check if next character is uppercase (e.g., isEnabled, not island)
			nextChar := selector[len(prefix)]
			if nextChar >= 'A' && nextChar <= 'Z' {
				return true
			}
		}
	}
	return false
}

// HasInitMethods checks if a method list contains any init methods.
// Init methods include instance init methods and class factory initializers.
// Examples:
//
//	[{Selector: "init"}] -> true
//	[{Selector: "initWithFrame:"}] -> true
//	[{Selector: "description"}] -> false
func HasInitMethods(methods []*ParsedMethod) bool {
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

// FilterInitMethods returns only the init/constructor methods from a list.
// This includes instance init methods and class factory methods.
func FilterInitMethods(methods []*ParsedMethod) []*ParsedMethod {
	var initMethods []*ParsedMethod
	for _, m := range methods {
		// Instance init methods
		if !m.IsClassMethod && strings.HasPrefix(m.Selector, "init") {
			initMethods = append(initMethods, m)
			continue
		}
		// Class factory methods that are constructors
		if m.IsClassMethod && (m.IsInitializer || IsEssentialSelector(m.Selector)) {
			initMethods = append(initMethods, m)
			continue
		}
	}
	return initMethods
}

// ShouldSkipMethod determines if a method should be skipped during code generation.
// Methods are skipped if they are inherited from NSObject, are memory management methods,
// or are otherwise not suitable for Go bindings.
// Examples:
//
//	retain -> true (skip, memory management)
//	description -> true (skip, inherited from NSObject)
//	customMethod -> false (don't skip)
func ShouldSkipMethod(selector string) bool {
	// Skip NSObject inherited methods
	if IsInheritedFromNSObject(selector) {
		return true
	}

	// Skip memory management methods (handled automatically in Go)
	if IsMemoryManagementSelector(selector) {
		return true
	}

	return false
}
