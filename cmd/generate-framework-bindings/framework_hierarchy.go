package main

import (
	"fmt"
	"os"
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
	"coreimage":       2,
	"quartzcore":      2,
	"coreaudio":       2,
	"coremidi":        2,
	"imageio":         2,
	"coredata":        2,
	"corelocation":    2,
	"corespotlight":   2,
	"network":         2,
	"security":        2,
	"corebluetooth":   2,
	"corevideo":       2,
	"coreml":          2,
	"vision":          2,
	"naturallanguage": 2,

	// Level 3: Application frameworks and UI
	"appkit":           3,
	"uikit":            3,
	"webkit":           3,
	"pdfkit":           3,
	"networkextension": 3,

	// Level 4: Higher-level application services
	"avfoundation":      4,
	"avfaudio":          4,
	"avkit":             4,
	"avrouting":         4,
	"audiotoolbox":      4,
	"cloudkit":          4,
	"contacts":          4,
	"contactsui":        4,
	"gameplaykit":       4,
	"intents":           4,
	"intentsui":         4,
	"metal":             4,
	"metalkit":          4,
	"eventkit":          4,
	"eventkit ui":       4,
	"healthkit":         4,
	"homekit":           4,
	"mapkit":            4,
	"messages":          4,
	"storekit":          4,
	"usernotifications": 4,
	"replaykit":         4,
}

// RelaxedParameterInfo tracks information about parameters that were relaxed due to hierarchy violations
type RelaxedParameterInfo struct {
	ClassName      string // e.g., "NSExtensionContext"
	MethodSelector string // e.g., "completeRequestWithBroadcastURL:broadcastConfiguration:setupInfo:"
	ParamName      string // e.g., "broadcastConfiguration"
	OriginalType   string // e.g., "RPBroadcastConfiguration *"
	ExpectedGoType string // e.g., "replaykit.BroadcastConfiguration"
}

// relaxedParametersMap tracks all relaxed parameters globally for comment generation
// Key format: "ClassName::MethodSelector::ParamName"
var relaxedParametersMap = make(map[string]*RelaxedParameterInfo)

// GetRelaxedParamInfo retrieves information about a relaxed parameter for use in templates
func GetRelaxedParamInfo(className, methodSelector, paramName string) *RelaxedParameterInfo {
	key := className + "::" + methodSelector + "::" + paramName
	return relaxedParametersMap[key]
}

// RelaxMethodParameters marks parameters that violate hierarchy with relaxed types.
// Instead of filtering out methods, we keep them but mark violating parameters
// to use objectivec.IObject with a comment indicating the expected type.
func RelaxMethodParameters(className string, methods []*occ2go.ParsedMethod, currentFramework string) {
	currentLevel, exists := frameworkLevels[strings.ToLower(currentFramework)]
	if !exists {
		// Unknown framework, don't relax
		return
	}

	for _, method := range methods {
		// Check return type for violations
		if violatesHierarchy(method.ReturnType, currentFramework, currentLevel) {
			// For now, skip methods with violating return types
			// We could relax these too, but it's more complex
			continue
		}

		// Check and relax parameters
		for i := range method.Parameters {
			param := &method.Parameters[i]
			if violatesHierarchy(param.Type, currentFramework, currentLevel) {
				// Get the expected Go type (with cross-framework reference)
				expectedType := mapObjCTypeToGo(param.Type, currentFramework)

				// Store relaxation info in global map for template use
				key := className + "::" + method.Selector + "::" + param.Name
				relaxedParametersMap[key] = &RelaxedParameterInfo{
					ClassName:      className,
					MethodSelector: method.Selector,
					ParamName:      param.Name,
					OriginalType:   param.Type,
					ExpectedGoType: expectedType,
				}

				// Relax the type to objectivec.IObject
				// We set the Type to a marker that the type mapper will recognize
				param.Type = "id" // Maps to objectivec.IObject
			}
		}
	}
}

// FilterMethodsByHierarchy filters methods that would create upward dependency violations.
// It removes methods that reference types from higher-level frameworks.
// NOTE: This is being phased out in favor of RelaxMethodParameters.
func FilterMethodsByHierarchy(methods []*occ2go.ParsedMethod, currentFramework string) []*occ2go.ParsedMethod {
	currentLevel, exists := frameworkLevels[strings.ToLower(currentFramework)]
	if !exists {
		// Unknown framework, don't filter
		return methods
	}

	filtered := make([]*occ2go.ParsedMethod, 0, len(methods))
	for _, method := range methods {
		// Only skip if return type violates hierarchy
		// Parameters will be relaxed by RelaxMethodParameters
		if violatesHierarchy(method.ReturnType, currentFramework, currentLevel) {
			if os.Getenv("DEBUG_HIERARCHY") == "1" {
				fmt.Fprintf(os.Stderr, "DEBUG: Skipping method %s due to return type %s\n", method.Name, method.ReturnType)
			}
			continue
		}
		filtered = append(filtered, method)
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

// violatesHierarchy checks if a type reference would create an upward dependency violation
func violatesHierarchy(objcType, currentFramework string, currentLevel int) bool {
	// Map the type to Go to see if it references another framework
	goType := mapObjCTypeToGo(objcType, currentFramework)

	debug := os.Getenv("DEBUG_HIERARCHY") == "1"
	if debug && (strings.Contains(objcType, "Hotspot") || strings.Contains(objcType, "Broadcast") || strings.Contains(objcType, "RPBroadcast") || strings.Contains(goType, "networkextension") || strings.Contains(goType, "replaykit")) {
		fmt.Fprintf(os.Stderr, "DEBUG violatesHierarchy: objcType=%s goType=%s currentFramework=%s currentLevel=%d\n",
			objcType, goType, currentFramework, currentLevel)
	}

	// Check if it's a cross-framework reference (contains '.')
	if !strings.Contains(goType, ".") {
		if debug && (strings.Contains(objcType, "Hotspot") || strings.Contains(goType, "networkextension")) {
			fmt.Fprintf(os.Stderr, "DEBUG violatesHierarchy: no cross-framework reference (no dot)\n")
		}
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
		if debug && (strings.Contains(objcType, "Hotspot") || strings.Contains(goType, "networkextension")) {
			fmt.Fprintf(os.Stderr, "DEBUG violatesHierarchy: unknown target framework %s\n", targetFramework)
		}
		return false
	}

	// Violation if target framework is higher level than current
	violation := targetLevel > currentLevel
	if debug && (strings.Contains(objcType, "Hotspot") || strings.Contains(goType, "networkextension")) {
		fmt.Fprintf(os.Stderr, "DEBUG violatesHierarchy: targetFramework=%s targetLevel=%d violation=%v\n",
			targetFramework, targetLevel, violation)
	}
	return violation
}
