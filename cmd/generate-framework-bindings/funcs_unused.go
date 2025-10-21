package main

// This file contains unused/unreachable functions that have been moved from funcs.go
// These are kept for reference and can be removed in future cleanup

/*
import (
	"reflect"
	"sort"
	"strings"

	"github.com/tmc/appledocs/occ2go"
)

// getClassRequiredImports returns a sorted slice of all required imports for a class (including both methods and properties).
// This is a convenience function for templates to get all imports at once.
// UNUSED: Commented out as unreachable code
func getClassRequiredImports(class interface{}, framework string) []string {
	imports := make(map[string]bool)
	currentFrameworkImportPath := "github.com/tmc/appledocs/generated/" + strings.ToLower(framework)

	// Try to extract class data using reflection
	classVal := reflect.ValueOf(class)
	if classVal.Kind() == reflect.Ptr {
		classVal = classVal.Elem()
	}

	// Try to access Methods field
	if classVal.Kind() == reflect.Struct {
		methodsField := classVal.FieldByName("Methods")
		if methodsField.IsValid() && methodsField.Kind() == reflect.Slice {
			for i := 0; i < methodsField.Len(); i++ {
				method := methodsField.Index(i).Interface()
				if parsedMethod, ok := method.(*occ2go.ParsedMethod); ok {
					// Check return type
					if parsedMethod.ReturnType != "" {
						goType := mapObjCTypeToGo(parsedMethod.ReturnType, framework)
						if importPath := getGoTypeImportPath(goType); importPath != "" {
							if importPath != currentFrameworkImportPath {
								imports[importPath] = true
							}
						}
					}

					// Check parameters
					for _, param := range parsedMethod.Parameters {
						goType := mapObjCTypeToGo(param.Type, framework)
						if importPath := getGoTypeImportPath(goType); importPath != "" {
							if importPath != currentFrameworkImportPath {
								imports[importPath] = true
							}
						}
					}
				}
			}
		}

		// Try to access Properties field
		propertiesField := classVal.FieldByName("Properties")
		if propertiesField.IsValid() && propertiesField.Kind() == reflect.Slice {
			for i := 0; i < propertiesField.Len(); i++ {
				property := propertiesField.Index(i).Interface()
				if parsedProp, ok := property.(*occ2go.ParsedProperty); ok {
					goType := mapObjCTypeToGo(parsedProp.Type, framework)
					if importPath := getGoTypeImportPath(goType); importPath != "" {
						if importPath != currentFrameworkImportPath {
							imports[importPath] = true
						}
					}
				}
			}
		}
	}

	// Convert map to sorted slice
	result := make([]string, 0, len(imports))
	for imp := range imports {
		result = append(result, imp)
	}
	sort.Strings(result)
	return result
}

// typeReferencesFramework checks if a Go type string contains a reference to a specific framework.
// It handles container types like []pkg.Type, map[string]pkg.Type, etc. by checking if the
// framework package name appears with a dot (pkg.) anywhere in the type string.
// UNUSED: Commented out as unreachable code
func typeReferencesFramework(goType, framework string) bool {
	// Check for "framework." pattern which indicates the framework is used as a package qualifier
	return strings.Contains(goType, framework+".")
}

// OLD IMPLEMENTATION - Replaced by registry-based approach above
// Kept for reference and can be removed in future cleanup
// UNUSED: Commented out as unreachable code
func getClassImportsOld(class *occ2go.ParsedClass, framework, outputModule string) ClassImports {
	imports := ClassImports{}

	if class == nil {
		return imports
	}

	// Determine struct name for self-referential check
	structName := classToStructName(class.Name)

	// Check struct embedding for import needs by using getStructEmbeddedField
	// This ensures we catch all cases where objectivec.Object is embedded
	// Don't import a framework into itself
	embeddedField := getStructEmbeddedField(class, framework)
	if strings.HasPrefix(embeddedField, "objectivec.") {
		imports.NeedsObjectiveC = true
	} else if strings.HasPrefix(embeddedField, "foundation.") && framework != "Foundation" {
		imports.NeedsFoundation = true
	} else if strings.HasPrefix(embeddedField, "quartzcore.") && framework != "QuartzCore" {
		imports.NeedsQuartzCore = true
	} else if strings.HasPrefix(embeddedField, "appkit.") && framework != "AppKit" {
		imports.NeedsAppKit = true
	}

	// Also check superclass for import needs (for interface embedding)
	if framework != "ObjectiveC" && class.SuperClass != "" {
		superStructName := classToStructName(class.SuperClass)
		superResolved := resolveType(framework, superStructName)
		isSelfReferential := (superStructName == structName)

		// If superclass has a framework prefix, we need that import
		// Don't import a framework into itself
		if strings.HasPrefix(superResolved, "foundation.") && framework != "Foundation" {
			imports.NeedsFoundation = true
		} else if strings.HasPrefix(superResolved, "quartzcore.") && framework != "QuartzCore" {
			imports.NeedsQuartzCore = true
		} else if strings.HasPrefix(superResolved, "appkit.") && framework != "AppKit" {
			imports.NeedsAppKit = true
		} else if strings.HasPrefix(superResolved, "objectivec.") {
			imports.NeedsObjectiveC = true
		} else if class.SuperClass == "NSObject" || superStructName == "Object" || isSelfReferential {
			imports.NeedsObjectiveC = true
		}
	} else if framework != "ObjectiveC" {
		// No superclass specified, default to objectivec
		imports.NeedsObjectiveC = true
	}

	// Check methods for CoreGraphics dependencies
	if classDependsOnCoreGraphics(class.Methods, framework) {
		imports.NeedsCoreGraphics = true
	}

	// Check properties for CoreGraphics dependencies
	if !imports.NeedsCoreGraphics {
		for _, prop := range class.Properties {
			goType := mapObjCTypeToGo(prop.Type, framework)
			if strings.HasPrefix(goType, "coregraphics.") {
				imports.NeedsCoreGraphics = true
				break
			}
		}
	}

	// Check method parameters and return types for AppKit, QuartzCore, and CloudKit dependencies
	for _, method := range class.Methods {
		// Check return type
		if method.ReturnType != "" {
			goType := mapObjCTypeToGo(method.ReturnType, framework)
			// Don't import a framework into itself
			if typeReferencesFramework(goType, "appkit") && framework != "AppKit" {
				imports.NeedsAppKit = true
			} else if typeReferencesFramework(goType, "quartzcore") && framework != "QuartzCore" {
				imports.NeedsQuartzCore = true
			} else if typeReferencesFramework(goType, "cloudkit") && framework != "CloudKit" {
				imports.NeedsCloudKit = true
			}
		}
		// Check parameters
		for _, param := range method.Parameters {
			goType := mapObjCTypeToGo(param.Type, framework)
			// Don't import a framework into itself
			if typeReferencesFramework(goType, "foundation") && framework != "Foundation" {
				imports.NeedsFoundation = true
			} else if typeReferencesFramework(goType, "appkit") && framework != "AppKit" {
				imports.NeedsAppKit = true
			} else if typeReferencesFramework(goType, "quartzcore") && framework != "QuartzCore" {
				imports.NeedsQuartzCore = true
			} else if typeReferencesFramework(goType, "cloudkit") && framework != "CloudKit" {
				imports.NeedsCloudKit = true
			}
		}
		if imports.NeedsFoundation || imports.NeedsAppKit || imports.NeedsQuartzCore || imports.NeedsCloudKit {
			break
		}
	}

	// Check method parameters and return types for Foundation dependencies
	if !imports.NeedsFoundation {
		for _, method := range class.Methods {
			// Check return type
			goReturnType := mapObjCTypeToGo(method.ReturnType, framework)
			if typeReferencesFramework(goReturnType, "foundation") {
				imports.NeedsFoundation = true
				break
			}
			// Check parameter types
			for _, param := range method.Parameters {
				goParamType := mapObjCTypeToGo(param.Type, framework)
				if typeReferencesFramework(goParamType, "foundation") {
					imports.NeedsFoundation = true
					break
				}
			}
			if imports.NeedsFoundation {
				break
			}
		}
	}

	// Check properties for Foundation dependencies
	if !imports.NeedsFoundation {
		for _, prop := range class.Properties {
			goType := mapObjCTypeToGo(prop.Type, framework)
			if strings.HasPrefix(goType, "foundation.") {
				imports.NeedsFoundation = true
				break
			}
		}
	}

	// Check method parameters and return types for UserNotifications dependencies
	if !imports.NeedsUserNotifications {
		for _, method := range class.Methods {
			// Check return type (use Contains to handle slices like []usernotifications.NotificationAction)
			goReturnType := mapObjCTypeToGo(method.ReturnType, framework)
			if strings.Contains(goReturnType, "usernotifications.") {
				imports.NeedsUserNotifications = true
				break
			}
			// Check parameter types
			for _, param := range method.Parameters {
				goParamType := mapObjCTypeToGo(param.Type, framework)
				if strings.Contains(goParamType, "usernotifications.") {
					imports.NeedsUserNotifications = true
					break
				}
			}
			if imports.NeedsUserNotifications {
				break
			}
		}
	}

	// Check properties for UserNotifications dependencies
	if !imports.NeedsUserNotifications {
		for _, prop := range class.Properties {
			goType := mapObjCTypeToGo(prop.Type, framework)
			if strings.Contains(goType, "usernotifications.") {
				imports.NeedsUserNotifications = true
				break
			}
		}
	}

	// Check method parameters and return types for UniformTypeIdentifiers dependencies
	if !imports.NeedsUniformTypeIdentifiers {
		for _, method := range class.Methods {
			// Check return type
			goReturnType := mapObjCTypeToGo(method.ReturnType, framework)
			if strings.Contains(goReturnType, "uniformtypeidentifiers.") && framework != "UniformTypeIdentifiers" {
				imports.NeedsUniformTypeIdentifiers = true
				break
			}
			// Check parameter types
			for _, param := range method.Parameters {
				goParamType := mapObjCTypeToGo(param.Type, framework)
				if strings.Contains(goParamType, "uniformtypeidentifiers.") && framework != "UniformTypeIdentifiers" {
					imports.NeedsUniformTypeIdentifiers = true
					break
				}
			}
			if imports.NeedsUniformTypeIdentifiers {
				break
			}
		}
	}

	// Check properties for UniformTypeIdentifiers dependencies
	if !imports.NeedsUniformTypeIdentifiers {
		for _, prop := range class.Properties {
			goType := mapObjCTypeToGo(prop.Type, framework)
			if strings.Contains(goType, "uniformtypeidentifiers.") && framework != "UniformTypeIdentifiers" {
				imports.NeedsUniformTypeIdentifiers = true
				break
			}
		}
	}

	return imports
}

// isTypedefConstant checks if a constant's type matches a known typedef
// UNUSED: Commented out as unreachable code
func isTypedefConstant(g *Generator) func(constType string) bool {
	return func(constType string) bool {
		for _, typedef := range g.Typedefs {
			if typedef.Name == constType {
				return true
			}
		}
		return false
	}
}
*/
