package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/tmc/appledocs/occ2go"
)

// buildCrossFrameworkTypeRegistry scans generated frameworks and populates the type registry.
// This allows proper type resolution instead of falling back to unsafe.Pointer.
//
// It scans the output directory for generated frameworks and extracts class names,
// building a map of type name -> framework package name.
//
// Example registry entries:
//
//	"Window" -> "appkit"
//	"String" -> "foundation"
//	"Layer" -> "quartzcore"
func buildCrossFrameworkTypeRegistry(outputDir string) error {
	// Check if output directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		// Output directory doesn't exist yet, registry will be empty
		return nil
	}

	// Scan all subdirectories (frameworks)
	entries, err := os.ReadDir(outputDir)
	if err != nil {
		return fmt.Errorf("failed to read output directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		frameworkPkg := strings.ToLower(entry.Name())
		frameworkDir := filepath.Join(outputDir, entry.Name())

		// Scan all .gen.go files in the framework directory to find type definitions
		// This includes types.gen.go as well as individual class files like ns_error.gen.go
		genFiles, err := filepath.Glob(filepath.Join(frameworkDir, "*.gen.go"))
		if err != nil || len(genFiles) == 0 {
			continue
		}

		// Extract type definitions (e.g., "type Window struct")
		// Match struct and interface types, but NOT unsafe.Pointer aliases
		// unsafe.Pointer aliases are forward declarations for types in other frameworks
		// Capture group 1 is the type name, group 2 is whether it's "struct" or "interface"
		typeRegex := regexp.MustCompile(`(?m)^type\s+([A-Z][A-Za-z0-9_]*)\s+(struct|interface)(?:\s|{)`)

		for _, genFile := range genFiles {
			// Parse each .gen.go file to extract type names
			data, err := os.ReadFile(genFile)
			if err != nil {
				continue // Skip on error
			}

			matches := typeRegex.FindAllSubmatch(data, -1)

			for _, match := range matches {
				if len(match) > 2 {
					typeName := string(match[1])
					typeKind := string(match[2]) // "struct" or "interface"

					// Add to registry if not already present (first framework wins)
					if _, exists := crossFrameworkTypeRegistry[typeName]; !exists {
						crossFrameworkTypeRegistry[typeName] = frameworkPkg
						Debug.TypeMap("registry: added type", typeName, frameworkPkg,
							"typeName", typeName,
							"framework", frameworkPkg,
							"typeKind", typeKind,
							"file", filepath.Base(genFile))
					}

					// If it's a struct, also add to the struct registry
					if typeKind == "struct" {
						qualifiedType := frameworkPkg + "." + typeName
						crossFrameworkStructRegistry[qualifiedType] = true
						Debug.TypeMap("registry: added struct", qualifiedType, "true",
							"qualifiedType", qualifiedType,
							"framework", frameworkPkg,
							"file", filepath.Base(genFile))
					}

					// Also register stripped name (NSCellAttribute → CellAttribute)
					// This allows lookups with ObjC names to find the stripped Go type
					strippedName := stripObjCPrefix(typeName)
					if strippedName != typeName {
						// Register stripped name pointing to the STRIPPED type, not the original
						// So NSCellAttribute lookup finds CellAttribute, not NSCellAttribute
						//
						// EXCEPTION: Don't register stripped names for CoreGraphics geometry types
						// (Point, Size, Rect, AffineTransform) because these conflict with the actual
						// Apple type names (CGPoint, CGSize, CGRect, CGAffineTransform).
						// NSPoint/NSSize/NSRect are Foundation typedefs for the CG types, so we need
						// the full "CG" prefix to be preserved.
						geometryTypes := map[string]bool{
							"Point":           true,
							"Size":            true,
							"Rect":            true,
							"AffineTransform": true,
						}
						if !geometryTypes[strippedName] {
							if _, exists := crossFrameworkTypeRegistry[strippedName]; !exists {
								crossFrameworkTypeRegistry[strippedName] = frameworkPkg
								Debug.TypeMap("registry: added stripped type", strippedName, typeName,
									"strippedName", strippedName,
									"originalName", typeName,
									"framework", frameworkPkg,
									"file", filepath.Base(genFile))
							}
						} else {
							Debug.TypeMap("registry: SKIPPED stripped geometry type", strippedName, typeName,
								"strippedName", strippedName,
								"originalName", typeName,
								"framework", frameworkPkg,
								"reason", "geometry type")
						}
					}
				}
			}
		}
	}

	// After scanning all frameworks, override common Foundation types to ensure they're always mapped to foundation
	// This prevents other frameworks (accessibility, authenticationservices, etc.) from claiming Foundation types
	// just because they come first alphabetically
	foundationCoreTypes := []string{
		"NSObject", "Object",
		"NSURL", "URL",
		"NSNumber", "Number",
		"NSString", "String",
		"NSArray", "Array",
		"NSDictionary", "Dictionary",
		"NSData", "Data",
		"NSDate", "Date",
		"NSSet", "Set",
		"NSErrorDomain",      // typedef to String - used across all frameworks for error domains
		"NSExtensionContext", // Foundation class, not CallKit/FileProviderUI/PhotosUI
		"ExtensionContext",   // stripped name
	}
	for _, typeName := range foundationCoreTypes {
		crossFrameworkTypeRegistry[typeName] = "foundation"
	}

	// UniformTypeIdentifiers core types
	uniformTypeIdentifiersCoreTypes := []string{
		"UTType",
	}
	for _, typeName := range uniformTypeIdentifiersCoreTypes {
		crossFrameworkTypeRegistry[typeName] = "uniformtypeidentifiers"
	}

	// CoreFoundation geometry types (used by CoreGraphics but defined in CoreFoundation)
	// See https://developer.apple.com/documentation/CoreGraphics - "Geometric Data Types" section
	// explicitly shows these types are defined in CoreFoundation package
	// IMPORTANT: These overrides must happen UNCONDITIONALLY to fix incorrect registrations
	coreFoundationGeometryTypes := []string{
		"CGPoint",
		"CGSize",
		"CGRect",
		"CGVector",
		"CGAffineTransform",
		"CGFloat", // Also from CoreFoundation
	}
	for _, typeName := range coreFoundationGeometryTypes {
		// Unconditionally set to corefoundation (overwrite any previous registration)
		crossFrameworkTypeRegistry[typeName] = "corefoundation"
		Debug.TypeMap("registry override: geometry type", typeName, "corefoundation",
			"typeName", typeName,
			"framework", "corefoundation")
	}

	return nil
}

// buildTypeRegistryFromParsedData populates the type registry from parsed documentation data.
// This is the preferred approach - using the source of truth (parsed docs) rather than
// scanning generated code artifacts.
//
// It extracts type mappings from:
//   - Enums: NSImageScaling → ImageScaling
//   - Classes: NSWindow → Window
//   - Typedefs: NSTimeInterval → TimeInterval
//
// The registry is used during type resolution to avoid unsafe.Pointer fallbacks.
func buildTypeRegistryFromParsedData(framework string, classes []*occ2go.ParsedClass, enums []*occ2go.ParsedEnum, typedefs []*occ2go.ParsedTypedef) {
	frameworkLower := strings.ToLower(framework)

	// Register enum types: both NSImageScaling→appkit and ImageScaling→appkit
	// Only register if not already present (preserves overrides from buildCrossFrameworkTypeRegistry)
	for _, enum := range enums {
		if enum.Name == "" {
			continue
		}
		// Strip NS/CG/CA prefix to get Go type name
		goTypeName := occ2go.StripObjCPrefix(enum.Name)
		// Register both ObjC name and Go name for lookup, but don't overwrite existing entries
		if _, exists := crossFrameworkTypeRegistry[enum.Name]; !exists {
			crossFrameworkTypeRegistry[enum.Name] = frameworkLower
		}
		if _, exists := crossFrameworkTypeRegistry[goTypeName]; !exists {
			crossFrameworkTypeRegistry[goTypeName] = frameworkLower
		}
	}

	// Register class types: both NSWindow→appkit and Window→appkit
	// Only register if not already present (preserves overrides from buildCrossFrameworkTypeRegistry)
	for _, class := range classes {
		if class.Name == "" {
			continue
		}
		// Strip NS/CG/CA prefix to get Go type name
		goTypeName := occ2go.StripObjCPrefix(class.Name)
		// Register both ObjC name and Go name for lookup, but don't overwrite existing entries
		if _, exists := crossFrameworkTypeRegistry[class.Name]; !exists {
			crossFrameworkTypeRegistry[class.Name] = frameworkLower
		}
		if _, exists := crossFrameworkTypeRegistry[goTypeName]; !exists {
			crossFrameworkTypeRegistry[goTypeName] = frameworkLower
		}
	}

	// Register typedef types: both NSTimeInterval→foundation and TimeInterval→foundation
	// Only register if not already present (preserves overrides from buildCrossFrameworkTypeRegistry)
	for _, typedef := range typedefs {
		if typedef.Name == "" {
			continue
		}
		// Strip NS/CG/CA prefix to get Go type name, then title-case it
		goTypeName := occ2go.StripObjCPrefix(typedef.Name)
		// Title-case the typedef name (e.g., objc_property_t → Objc_property_t)
		if goTypeName != "" {
			goTypeName = strings.ToUpper(goTypeName[:1]) + goTypeName[1:]
		}
		// Register both ObjC name and Go name for lookup, but don't overwrite existing entries
		if _, exists := crossFrameworkTypeRegistry[typedef.Name]; !exists {
			crossFrameworkTypeRegistry[typedef.Name] = frameworkLower
		}
		if _, exists := crossFrameworkTypeRegistry[goTypeName]; !exists {
			crossFrameworkTypeRegistry[goTypeName] = frameworkLower
		}
	}
}

// buildStructRegistryFromParsedData populates a registry tracking which types are structs vs classes.
// This allows TypeToInterfaceType to make data-driven decisions instead of hardcoding package names.
//
// Structs should NOT be converted to interfaces (AffineTransform stays AffineTransform).
// Classes should be converted to interfaces (Window becomes IWindow).
func buildStructRegistryFromParsedData(framework string, structs []*occ2go.ParsedStruct) {
	frameworkLower := strings.ToLower(framework)

	// Register struct types so TypeToInterfaceType can detect them
	for _, strct := range structs {
		if strct.Name == "" {
			continue
		}
		// Strip NS/CG/CA prefix to get Go type name
		goTypeName := occ2go.StripObjCPrefix(strct.Name)

		// Store in a separate struct-specific registry
		// Format: "coregraphics.AffineTransform" → true (is a struct)
		qualifiedType := frameworkLower + "." + goTypeName
		crossFrameworkStructRegistry[qualifiedType] = true

		// Also register the unqualified name for current-framework lookups
		// This allows intra-framework lookups (e.g., AffineTransform in coregraphics → coregraphics.AffineTransform)
		if _, exists := crossFrameworkTypeRegistry[goTypeName]; !exists {
			crossFrameworkTypeRegistry[goTypeName] = frameworkLower
		}
	}
}
