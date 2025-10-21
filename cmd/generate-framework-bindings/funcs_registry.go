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
		typeRegex := regexp.MustCompile(`(?m)^type\s+([A-Z][A-Za-z0-9_]*)\s+(?:struct|interface)(?:\s|{)`)

		for _, genFile := range genFiles {
			// Parse each .gen.go file to extract type names
			data, err := os.ReadFile(genFile)
			if err != nil {
				continue // Skip on error
			}

			matches := typeRegex.FindAllSubmatch(data, -1)

			for _, match := range matches {
			if len(match) > 1 {
				typeName := string(match[1])
				// Add to registry if not already present (first framework wins)
				if _, exists := crossFrameworkTypeRegistry[typeName]; !exists {
					crossFrameworkTypeRegistry[typeName] = frameworkPkg
					if os.Getenv("DEBUG_TYPEMAP") == "1" && (strings.Contains(typeName, "Coder") || strings.Contains(typeName, "Error") || strings.Contains(typeName, "Operation")) {
						fmt.Fprintf(os.Stderr, "DEBUG registry: added %s -> %s (from %s)\n", typeName, frameworkPkg, filepath.Base(genFile))
					}
				}

				// Also register stripped name (NSCellAttribute → CellAttribute)
				// This allows lookups with ObjC names to find the stripped Go type
				strippedName := stripObjCPrefix(typeName)
				if strippedName != typeName {
					// Register stripped name pointing to the STRIPPED type, not the original
					// So NSCellAttribute lookup finds CellAttribute, not NSCellAttribute
					if _, exists := crossFrameworkTypeRegistry[strippedName]; !exists {
						crossFrameworkTypeRegistry[strippedName] = frameworkPkg
						if os.Getenv("DEBUG_TYPEMAP") == "1" && (strings.Contains(typeName, "Coder") || strings.Contains(typeName, "Error") || strings.Contains(typeName, "Operation")) {
							fmt.Fprintf(os.Stderr, "DEBUG registry: added stripped %s (from %s) -> %s (from %s)\n", strippedName, typeName, frameworkPkg, filepath.Base(genFile))
						}
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
	for _, enum := range enums {
		if enum.Name == "" {
			continue
		}
		// Strip NS/CG/CA prefix to get Go type name
		goTypeName := occ2go.StripObjCPrefix(enum.Name)
		// Register both ObjC name and Go name for lookup
		crossFrameworkTypeRegistry[enum.Name] = frameworkLower
		crossFrameworkTypeRegistry[goTypeName] = frameworkLower
	}

	// Register class types: both NSWindow→appkit and Window→appkit
	for _, class := range classes {
		if class.Name == "" {
			continue
		}
		// Strip NS/CG/CA prefix to get Go type name
		goTypeName := occ2go.StripObjCPrefix(class.Name)
		// Register both ObjC name and Go name for lookup
		crossFrameworkTypeRegistry[class.Name] = frameworkLower
		crossFrameworkTypeRegistry[goTypeName] = frameworkLower
	}

	// Register typedef types: both NSTimeInterval→foundation and TimeInterval→foundation
	for _, typedef := range typedefs {
		if typedef.Name == "" {
			continue
		}
		// Strip NS/CG/CA prefix to get Go type name
		goTypeName := occ2go.StripObjCPrefix(typedef.Name)
		// Register both ObjC name and Go name for lookup
		crossFrameworkTypeRegistry[typedef.Name] = frameworkLower
		crossFrameworkTypeRegistry[goTypeName] = frameworkLower
	}
}
