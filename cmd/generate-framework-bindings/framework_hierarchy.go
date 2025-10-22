package main

import (
	"strings"

	"github.com/tmc/appledocs/occ2go"
)

// frameworkLevels defines the dependency hierarchy of Apple frameworks.
// Lower numbers are more foundational. A framework should only depend on
// frameworks at the same or lower level, never higher levels.
var frameworkLevels = map[string]int{
	// Level 0: Core Objective-C runtime
	"objc":       0,
	"objectivec": 0,

	// Level 1: Core graphics, foundation, and system services
	"coregraphics":   1,
	"corefoundation": 1,
	"foundation":     1,
	"coretext":       1,
	"iosurface":      1,

	// Level 2: Media and core services
	"coreimage":    2,
	"quartzcore":   2,
	"coreaudio":    2,
	"coremidi":     2,
	"imageio":      2,
	"coredata":     2,
	"corelocation": 2,
	"corespotlight": 2,

	// Level 3: Application frameworks and UI
	"appkit":  3,
	"uikit":   3,
	"webkit":  3,

	// Level 4: Higher-level application services
	"avfoundation": 4,
	"avfaudio":     4,
	"avkit":        4,
	"audiotoolbox": 4,
	"cloudkit":     4,
	"contacts":     4,
	"contactsui":   4,
	"metal":        4,
	"metalkit":     4,
}

// FilterMethodsByHierarchy filters methods that would create upward dependency violations.
// It removes methods that reference types from higher-level frameworks.
func FilterMethodsByHierarchy(methods []*occ2go.ParsedMethod, currentFramework string) []*occ2go.ParsedMethod {
	currentLevel, exists := frameworkLevels[strings.ToLower(currentFramework)]
	if !exists {
		// Unknown framework, don't filter
		return methods
	}

	filtered := make([]*occ2go.ParsedMethod, 0, len(methods))
	for _, method := range methods {
		if !shouldSkipMethod(method, currentFramework, currentLevel) {
			filtered = append(filtered, method)
		}
	}

	return filtered
}

// FilterPropertiesByHierarchy filters properties that would create upward dependency violations.
// It removes properties that reference types from higher-level frameworks.
func FilterPropertiesByHierarchy(properties []*occ2go.ParsedProperty, currentFramework string) []*occ2go.ParsedProperty {
	currentLevel, exists := frameworkLevels[strings.ToLower(currentFramework)]
	if !exists {
		// Unknown framework, don't filter
		return properties
	}

	filtered := make([]*occ2go.ParsedProperty, 0, len(properties))
	for _, prop := range properties {
		if !shouldSkipProperty(prop, currentFramework, currentLevel) {
			filtered = append(filtered, prop)
		}
	}

	return filtered
}

// shouldSkipProperty returns true if a property references a type from a higher-level framework
func shouldSkipProperty(prop *occ2go.ParsedProperty, currentFramework string, currentLevel int) bool {
	// Use ObjCType if available (more accurate), fall back to Type
	propType := prop.Type
	if prop.ObjCType != "" {
		propType = prop.ObjCType
	}

	return violatesHierarchy(propType, currentFramework, currentLevel)
}

// shouldSkipMethod returns true if a method references types from a higher-level framework
func shouldSkipMethod(method *occ2go.ParsedMethod, currentFramework string, currentLevel int) bool {
	// Check return type
	if violatesHierarchy(method.ReturnType, currentFramework, currentLevel) {
		return true
	}

	// Check all parameters
	for _, param := range method.Parameters {
		if violatesHierarchy(param.Type, currentFramework, currentLevel) {
			return true
		}
	}

	return false
}

// violatesHierarchy checks if a type reference would create an upward dependency violation
func violatesHierarchy(objcType, currentFramework string, currentLevel int) bool {
	// Map the type to Go to see if it references another framework
	goType := mapObjCTypeToGo(objcType, currentFramework)

	// Check if it's a cross-framework reference (contains '.')
	if !strings.Contains(goType, ".") {
		return false
	}

	// Extract the framework package name
	parts := strings.Split(goType, ".")
	if len(parts) < 2 {
		return false
	}

	targetFramework := parts[0]
	targetLevel, exists := frameworkLevels[targetFramework]
	if !exists {
		// Unknown target framework, allow it (might be a new framework we haven't categorized)
		return false
	}

	// Violation if target framework is higher level than current
	return targetLevel > currentLevel
}
