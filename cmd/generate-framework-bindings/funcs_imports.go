package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tmc/appledocs/occ2go"
)

// needsFoundationImport checks if any types require foundation import
func needsFoundationImport(methods []*occ2go.ParsedMethod) bool {
	for _, m := range methods {
		// Check return type
		if strings.Contains(m.ReturnType, "NS") &&
			!strings.Contains(m.ReturnType, "NSInteger") &&
			!strings.Contains(m.ReturnType, "NSUInteger") {
			return true
		}
		// Check parameters
		for _, p := range m.Parameters {
			if strings.Contains(p.Type, "NS") &&
				!strings.Contains(p.Type, "NSInteger") &&
				!strings.Contains(p.Type, "NSUInteger") {
				return true
			}
		}
	}
	return false
}

// needsQuartzCoreImport checks if any types require quartzcore import
func needsQuartzCoreImport(methods []*occ2go.ParsedMethod) bool {
	for _, m := range methods {
		// Check return type
		if strings.Contains(m.ReturnType, "CA") || strings.Contains(m.ReturnType, "CI") {
			return true
		}
		// Check parameters
		for _, p := range m.Parameters {
			if strings.Contains(p.Type, "CA") || strings.Contains(p.Type, "CI") {
				return true
			}
		}
	}
	return false
}

// needsCustomImports checks if any methods use types that require custom imports
func needsCustomImports(methods []*occ2go.ParsedMethod, framework string) bool {
	return getRequiredImports(methods, framework) != nil
}

// getRequiredImports returns a map of import paths needed for methods.
// It maps Objective-C types to Go types first, then checks if those Go types need imports.
// This prevents adding imports for types that get mapped to unsafe.Pointer or built-in types.
// For example: {"github.com/tmc/appledocs/generated/coregraphics": true}
func getRequiredImports(methods []*occ2go.ParsedMethod, framework string) map[string]bool {
	imports := make(map[string]bool)

	// Check all methods for types that need custom imports
	for _, m := range methods {
		// Check return type - map to Go first, then check if it needs an import
		if m.ReturnType != "" && m.ReturnType != "void" {
			goType := mapObjCTypeToGo(m.ReturnType, framework)
			// Check if this Go type needs an import (e.g., coregraphics.CGAffineTransform)
			if importPath := getGoTypeImportPath(goType); importPath != "" {
				imports[importPath] = true
			}
		}

		// Check parameters - map to Go first, then check if they need imports
		for _, p := range m.Parameters {
			goType := mapObjCTypeToGo(p.Type, framework)
			// Check if this Go type needs an import
			if importPath := getGoTypeImportPath(goType); importPath != "" {
				imports[importPath] = true
			}
		}
	}

	if len(imports) == 0 {
		return nil
	}
	return imports
}

// getClassRequiredImports was moved to funcs_unused.go - see that file for the implementation

// getFunctionRequiredImports analyzes standalone functions to collect required import paths.
// It examines each function's return type and parameters, applying type mappings and extracting
// import paths from qualified type names (e.g., "coregraphics.CGAffineTransform").
func getFunctionRequiredImports(functions []*occ2go.ParsedFunction, framework string) map[string]bool {
	imports := make(map[string]bool)

	// Build the current framework's import path to filter it out
	currentFrameworkImportPath := "github.com/tmc/appledocs/generated/" + strings.ToLower(framework)

	for _, fn := range functions {
		// Check return type
		if fn.ReturnType != "" && fn.ReturnType != "void" {
			// Use mapCTypeToGoWithFramework to get the same result as prepareFunctionData
			goType := mapCTypeToGoWithFramework(fn.ReturnType, framework)
			if importPath := getGoTypeImportPath(goType); importPath != "" {
				// Don't import the current framework itself
				if importPath != currentFrameworkImportPath {
					imports[importPath] = true
				}
			}
		}

		// Check all parameters
		for _, param := range fn.Parameters {
			// Use mapCTypeToGoWithFramework to get the same result as prepareFunctionData
			goType := mapCTypeToGoWithFramework(param.Type, framework)
			if importPath := getGoTypeImportPath(goType); importPath != "" {
				// Don't import the current framework itself
				if importPath != currentFrameworkImportPath {
					imports[importPath] = true
				}
			}
		}
	}

	return imports
}

// sortedImportPaths returns a sorted slice of import paths for template iteration.
// This makes it easy for templates to range over imports in a consistent order.
func sortedImportPaths(imports map[string]bool) []ImportInfo {
	if imports == nil || len(imports) == 0 {
		return []ImportInfo{}
	}

	// Convert map keys to slice
	paths := make([]string, 0, len(imports))
	for path := range imports {
		paths = append(paths, path)
	}

	// Simple sort by string value for consistency
	for i := 0; i < len(paths)-1; i++ {
		for j := i + 1; j < len(paths); j++ {
			if paths[i] > paths[j] {
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}

	// Convert paths to ImportInfo structs
	result := make([]ImportInfo, 0, len(paths))
	for _, path := range paths {
		packageName := extractPackageNameFromImportPath(path)
		result = append(result, ImportInfo{
			PackageName: packageName,
			ImportPath:  path,
		})
	}

	return result
}

// getGoTypeImportPath takes a Go type string (after mapping from Objective-C) and returns
// the import path if it requires one, or empty string if it doesn't.
// Examples:
//
//	"coregraphics.CGAffineTransform" -> "github.com/tmc/appledocs/generated/coregraphics"
//	"foundation.Rect" -> "github.com/tmc/appledocs/generated/foundation"
//	"unsafe.Pointer" -> ""
//	"int" -> ""
//	"bool" -> ""
func getGoTypeImportPath(goType string) string {
	// Built-in types and types from std library don't need custom imports
	if goType == "" || goType == "unsafe.Pointer" {
		return ""
	}

	// Use registry helper to extract import path from type
	return GetImportPathFromType(goType)
}

// extractPackageNameFromImportPath extracts the package name from an import path.
// For example: "github.com/tmc/appledocs/generated/coregraphics" -> "coregraphics"
func extractPackageNameFromImportPath(importPath string) string {
	parts := strings.Split(importPath, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return importPath
}

// mergeImports merges two import maps into a single deduplicated map.
// Useful for combining imports from class methods and instance methods.
func mergeImports(map1, map2 map[string]bool) map[string]bool {
	result := make(map[string]bool)

	// Add all imports from map1
	for path := range map1 {
		result[path] = true
	}

	// Add all imports from map2
	for path := range map2 {
		result[path] = true
	}

	return result
}

// getClassImportPaths analyzes a class and returns a set of import paths needed.
// This is a registry-based helper that automatically detects all framework dependencies.
func getClassImportPaths(class *occ2go.ParsedClass, framework, outputModule string) map[string]bool {
	imports := make(map[string]bool)

	if class == nil {
		return imports
	}

	// Determine current framework's import path so we don't import ourselves
	currentFrameworkImportPath := outputModule + "/" + strings.ToLower(framework)

	// Check embedded field for import needs
	// Use getStructEmbeddedField which has the correct logic for NSObject handling
	embeddedField := getStructEmbeddedField(class, framework)
	if importPath := GetImportPathFromType(embeddedField); importPath != "" {
		if importPath != currentFrameworkImportPath {
			if class.Name == "CKQueryCursor" && strings.Contains(importPath, "appkit") {
				fmt.Fprintf(os.Stderr, ">>> CKQueryCursor embedded field %q -> import %q\n", embeddedField, importPath)
			}
			Debug.Imports("embedded field import", embeddedField, importPath,
				"class", class.Name,
				"embeddedField", embeddedField,
				"importPath", importPath)
			imports[importPath] = true
		}
	}

	// Check all method return types and parameters
	for _, method := range class.Methods {
		// Check return type
		if method.ReturnType != "" {
			goType := mapObjCTypeToGo(method.ReturnType, framework)
			if importPath := GetImportPathFromType(goType); importPath != "" {
				if importPath != currentFrameworkImportPath {
					// Check for framework hierarchy violations before adding import
					// If the target framework is at a higher level than current, skip the import
					// since the template will use objectivec.IObject instead
					if strings.Contains(goType, ".") {
						parts := strings.Split(goType, ".")
						if len(parts) >= 2 {
							targetFramework := parts[0]
							currentLevel := getFrameworkLevel(strings.ToLower(framework))
							targetLevel := getFrameworkLevel(targetFramework)
							if currentLevel >= 0 && targetLevel > currentLevel {
								// Skip import for hierarchy violation
								Debug.Imports("skipping return type import due to hierarchy violation", method.Name, goType,
									"class", class.Name,
									"method", method.Name,
									"currentLevel", currentLevel,
									"targetLevel", targetLevel)
								continue
							}
						}
					}

					if class.Name == "CKQueryCursor" && strings.Contains(importPath, "appkit") {
						fmt.Fprintf(os.Stderr, ">>> CKQueryCursor method %s return type %q -> go type %q -> import %q\n",
							method.Name, method.ReturnType, goType, importPath)
					}
					Debug.Imports("method return type import", method.Name, goType,
						"class", class.Name,
						"method", method.Name,
						"returnType", method.ReturnType,
						"goType", goType,
						"importPath", importPath)
					imports[importPath] = true
				}
			}
		}

		// Check all parameters
		for _, param := range method.Parameters {
			goType := mapObjCTypeToGo(param.Type, framework)

			// When we use objc.ID as a parameter, we convert it to objectivec.IObject in formatMethodParams
			// So we need to import the objectivec package
			if goType == "objc.ID" && framework != "ObjectiveC" {
				// Add objectivec package import
				if globalRegistry != nil {
					if objcImport := globalRegistry.GetImportPathByPackage("objectivec"); objcImport != "" {
						imports[objcImport] = true
					}
				}
			}

			if importPath := GetImportPathFromType(goType); importPath != "" {
				if importPath != currentFrameworkImportPath {
					// Check for framework hierarchy violations before adding import
					// If the target framework is at a higher level than current, skip the import
					// since the template will use objectivec.IObject instead (fixes appledocs-496)
					if strings.Contains(goType, ".") {
						parts := strings.Split(goType, ".")
						if len(parts) >= 2 {
							targetFramework := parts[0]
							currentLevel := getFrameworkLevel(strings.ToLower(framework))
							targetLevel := getFrameworkLevel(targetFramework)
							if currentLevel >= 0 && targetLevel > currentLevel {
								// Skip import for hierarchy violation - template will use objectivec.IObject
								Debug.Imports("skipping import due to hierarchy violation", method.Name, goType,
									"class", class.Name,
									"method", method.Name,
									"param", param.Name,
									"currentLevel", currentLevel,
									"targetLevel", targetLevel)
								continue
							}
						}
					}

					if class.Name == "CKQueryCursor" && strings.Contains(importPath, "appkit") {
						fmt.Fprintf(os.Stderr, ">>> CKQueryCursor method %s param %q type %q -> go type %q -> import %q\n",
							method.Name, param.Name, param.Type, goType, importPath)
					}
					Debug.Imports("method parameter import", method.Name, goType,
						"class", class.Name,
						"method", method.Name,
						"param", param.Name,
						"paramType", param.Type,
						"goType", goType,
						"importPath", importPath)
					imports[importPath] = true
				}
			}
		}
	}

	// Check all properties
	for _, prop := range class.Properties {
		// Use ObjCType instead of Type - Type contains Swift syntax, ObjCType is the proper Objective-C type
		objcType := prop.Type
		if prop.ObjCType != "" {
			objcType = prop.ObjCType
		}
		goType := mapObjCTypeToGo(objcType, framework)
		if importPath := GetImportPathFromType(goType); importPath != "" {
			if importPath != currentFrameworkImportPath {
				// Check for framework hierarchy violations before adding import
				// If the target framework is at a higher level than current, skip the import
				// since the template will use objectivec.IObject instead
				if strings.Contains(goType, ".") {
					parts := strings.Split(goType, ".")
					if len(parts) >= 2 {
						targetFramework := parts[0]
						currentLevel := getFrameworkLevel(strings.ToLower(framework))
						targetLevel := getFrameworkLevel(targetFramework)
						if currentLevel >= 0 && targetLevel > currentLevel {
							// Skip import for hierarchy violation
							Debug.Imports("skipping property import due to hierarchy violation", prop.Name, goType,
								"class", class.Name,
								"property", prop.Name,
								"currentLevel", currentLevel,
								"targetLevel", targetLevel)
							continue
						}
					}
				}

				if class.Name == "CKQueryCursor" && strings.Contains(importPath, "appkit") {
					fmt.Fprintf(os.Stderr, ">>> CKQueryCursor property %q objcType %q -> go type %q -> import %q\n",
						prop.Name, objcType, goType, importPath)
				}
				Debug.Imports("property type import", prop.Name, goType,
					"class", class.Name,
					"property", prop.Name,
					"objcType", objcType,
					"goType", goType,
					"importPath", importPath)
				imports[importPath] = true
			}
		}
	}

	return imports
}

// typeReferencesFramework was moved to funcs_unused.go - see that file for the implementation

// getClassImports analyzes a class and its methods to determine which framework imports are needed.
// This uses the registry-based system to automatically detect all framework dependencies.
// Returns a ClassImports struct with:
//   - ImportPaths: Dynamic map of all imports (package name -> import path) - RECOMMENDED
//   - Boolean fields: Deprecated legacy fields for backward compatibility - will be removed
func getClassImports(class *occ2go.ParsedClass, framework, outputModule string) ClassImports {
	// Use the registry-based function to detect all imports automatically
	importPathsSet := getClassImportPaths(class, framework, outputModule)

	// Initialize struct with dynamic map
	imports := ClassImports{
		ImportPaths: make(map[string]string),
	}

	// Populate the dynamic ImportPaths map (this is the source of truth)
	for importPath := range importPathsSet {
		// Extract package name from import path
		// e.g., "github.com/tmc/appledocs/generated/foundation" -> "foundation"
		parts := strings.Split(importPath, "/")
		if len(parts) == 0 {
			continue
		}
		pkgName := parts[len(parts)-1]

		// Add to dynamic map (works for ALL frameworks automatically)
		imports.ImportPaths[pkgName] = importPath
	}

	// Set legacy boolean fields for backward compatibility
	// These are derived from ImportPaths map - the map is the source of truth
	imports.NeedsObjectiveC = imports.ImportPaths["objectivec"] != ""
	imports.NeedsFoundation = imports.ImportPaths["foundation"] != ""
	imports.NeedsQuartzCore = imports.ImportPaths["quartzcore"] != ""
	imports.NeedsCoreGraphics = imports.ImportPaths["coregraphics"] != ""
	imports.NeedsCloudKit = imports.ImportPaths["cloudkit"] != ""
	imports.NeedsAppKit = imports.ImportPaths["appkit"] != ""
	imports.NeedsUserNotifications = imports.ImportPaths["usernotifications"] != ""
	imports.NeedsUniformTypeIdentifiers = imports.ImportPaths["uniformtypeidentifiers"] != ""

	return imports
}

// getSortedClassImports returns a sorted slice of import paths from ClassImports.
// This is a template helper that makes it easy to iterate over all imports dynamically.
// Returns []ImportInfo with PackageName and ImportPath fields.
func getSortedClassImports(imports ClassImports) []ImportInfo {
	if imports.ImportPaths == nil || len(imports.ImportPaths) == 0 {
		return []ImportInfo{}
	}

	// Convert map to slice of ImportInfo
	result := make([]ImportInfo, 0, len(imports.ImportPaths))
	for pkgName, importPath := range imports.ImportPaths {
		result = append(result, ImportInfo{
			PackageName: pkgName,
			ImportPath:  importPath,
		})
	}

	// Sort by import path for consistency
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i].ImportPath > result[j].ImportPath {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result
}
