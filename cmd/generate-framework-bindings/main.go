// generate-framework-bindings generates Go bindings for macOS/iOS frameworks
// from Apple's documentation JSON files.
//
// This generator supports framework-by-framework generation for:
// - ObjectiveC runtime
// - CoreGraphics
// - CoreFoundation
// - AppKit
// - Foundation
// - And other Apple frameworks
package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/tmc/appledocs"
	"github.com/tmc/appledocs/occ2go"
)

//go:embed funcs.go templates.txtar templates_*.txtar config.yaml
var embeddedFS embed.FS

var verbose bool

func generateFramework(framework, inputDir, outputDir, filterRegexp string, txtarOutput bool, variant string, withRefMethods, generateTests, generateExamples bool) error {
	startTime := time.Now()
	if verbose {
		fmt.Fprintf(os.Stderr, "[%s] Starting generation\n", framework)
	}

	// Check if this is an iOS-only framework and skip if so
	_, _, iOSOnly := loadFrameworkMetadata(inputDir, framework)
	if iOSOnly {
		// Already logged in loadFrameworkMetadata with verbose flag
		return nil // Skip generation silently
	}

	// Open the appledocs filesystem
	phaseStart := time.Now()
	fsys, err := appledocs.Open(inputDir)
	if err != nil {
		return fmt.Errorf("failed to open appledocs filesystem: %w", err)
	}
	if verbose {
		fmt.Fprintf(os.Stderr, "[%s] Opened filesystem (%.2fs)\n", framework, time.Since(phaseStart).Seconds())
	}

	// Parse all symbols in the framework using the appledocs iterator
	var functions []*occ2go.ParsedFunction
	var classes []*occ2go.ParsedClass
	var protocols []*occ2go.ParsedProtocol
	var enums []*occ2go.ParsedEnum
	var typedefs []*occ2go.ParsedTypedef
	var constants []*occ2go.ParsedConstant

	processedFiles := 0
	parseErrors := 0

	// First, extract synthetic documents from API collection pages
	phaseStart = time.Now()
	syntheticDocs := extractSymbolsFromAPICollections(fsys, framework, verbose)
	if verbose {
		fmt.Fprintf(os.Stderr, "[%s] Extracted %d synthetic docs (%.2fs)\n", framework, len(syntheticDocs), time.Since(phaseStart).Seconds())
	}

	// Build a set of properties that have separate JSON files
	// to avoid duplicates when extracting from class references
	propertyFiles := make(map[string]bool)
	for _, doc := range appledocs.Symbols(fsys, framework) {
		if strings.Contains(doc.Metadata.ExternalID, "(py)") || strings.Contains(doc.Metadata.ExternalID, "(cpy)") {
			// This is a property file
			// Extract class and property name from external ID
			// Format: c:objc(cs)ClassName(py)propertyName or c:objc(cs)ClassName(cpy)propertyName
			externalID := doc.Metadata.ExternalID

			// Extract class name: c:objc(cs)ClassName(py)... -> ClassName
			if idx := strings.Index(externalID, "(cs)"); idx != -1 {
				rest := externalID[idx+4:] // Skip "(cs)"
				// Find the next '(' which marks the start of (py) or (cpy)
				if endIdx := strings.Index(rest, "("); endIdx != -1 {
					className := rest[:endIdx]
					// Extract property name after (py) or (cpy)
					if pyIdx := strings.Index(rest, "(py)"); pyIdx != -1 {
						propertyName := rest[pyIdx+4:]
						// Use lowercase for case-insensitive comparison
						propertyKey := className + "." + strings.ToLower(propertyName)
						propertyFiles[propertyKey] = true
					} else if cpyIdx := strings.Index(rest, "(cpy)"); cpyIdx != -1 {
						propertyName := rest[cpyIdx+5:]
						// Use lowercase for case-insensitive comparison
						propertyKey := className + "." + strings.ToLower(propertyName)
						propertyFiles[propertyKey] = true
					}
				}
			}
		}
	}
	if verbose && len(propertyFiles) > 0 {
		fmt.Fprintf(os.Stderr, "Found %d properties with separate JSON files\n", len(propertyFiles))
		// Debug: show first few CIVector properties
		for key := range propertyFiles {
			if strings.HasPrefix(key, "CIVector.") {
				fmt.Fprintf(os.Stderr, "  Property file: %s\n", key)
			}
		}
	}

	// Extract properties from class references (e.g., NSButton.title)
	// These properties don't have separate JSON files and are only in class references
	refProperties := extractPropertiesFromClassReferences(fsys, framework, propertyFiles, verbose)

	// Keep track of properties separately so we can attach them to classes later
	classPropertiesMap := make(map[string][]*occ2go.ParsedProperty)
	// Track seen property names per class to prevent duplicates from documentation
	classSeenProperties := make(map[string]map[string]bool)

	// Note: We'll add reference-based properties AFTER parsing separate files
	// to prefer the detailed property information from dedicated files

	// Keep track of enums and their cases
	enumCasesMap := make(map[string][]*occ2go.ParsedEnumCase)

	// Process regular symbols
	phaseStart = time.Now()
	for _, doc := range appledocs.Symbols(fsys, framework) {
		processedFiles++

		fn, cls, proto, err := occ2go.ParseDocument(doc)
		if err == nil {
			if fn != nil {
				functions = append(functions, fn)
			}
			if cls != nil {
				// Only include classes that actually belong to this framework
				if classbelongsToFramework(doc, framework) {
					classes = append(classes, cls)
				}
			}
			if proto != nil {
				protocols = append(protocols, proto)
			}
		} else if strings.Contains(err.Error(), "property (use ParseProperty)") {
			// This is a property file, parse it as a property
			property, propErr := occ2go.ParseProperty(doc)
			if propErr != nil {
				parseErrors++
			} else if property != nil {
				// Extract class name from external ID
				externalID := doc.Metadata.ExternalID
				if strings.Contains(externalID, "(py)") || strings.Contains(externalID, "(cpy)") {
					parts := strings.Split(externalID, "(")
					if len(parts) >= 2 {
						className := strings.TrimPrefix(parts[1], "cs)")

						// Initialize property tracking for this class if needed
						if classSeenProperties[className] == nil {
							classSeenProperties[className] = make(map[string]bool)
						}

						// Only add if we haven't seen this property name before
						if !classSeenProperties[className][property.Name] {
							classPropertiesMap[className] = append(classPropertiesMap[className], property)
							classSeenProperties[className][property.Name] = true
						}
					}
				}
			}
		} else if strings.HasPrefix(doc.Metadata.ExternalID, "c:@E@") {
			// This is an Objective-C enum (may be represented as Swift struct/enum)
			// Swift structs with ObjC enum externalIDs are option sets (like NSWindowStyleMask)
			// We should process these as enums, not skip them
			// Try parsing as enum type first (enum type has only 3 parts: c:@E@EnumName)
			parts := strings.Split(doc.Metadata.ExternalID, "@")
			if len(parts) == 3 {
				// This is an enum type declaration
				// Extract tokens from primary content sections
				tokens := []appledocs.Token{}
				for _, section := range doc.PrimaryContentSections {
					if section.Kind == "declarations" && len(section.Declarations) > 0 {
						tokens = section.Declarations[0].Tokens
						break
					}
				}
				// Try ObjC enum parsing first (enum NSBackingStoreType : NSUInteger)
				enum := occ2go.ParseEnumDeclaration(tokens)
				// If that fails, try Swift struct/enum parsing (struct StyleMask)
				if enum == nil {
					enum = occ2go.ParseSwiftEnumDeclaration(tokens)
				}
				if enum != nil {
					enum.Name = parts[2]
					enum.DocURL = doc.Identifier.URL
					if len(doc.Abstract) > 0 {
						enum.Abstract = doc.Abstract[0].Text
					}
					if os.Getenv("DEBUG_ENUM_CREATE") == "1" && (strings.Contains(enum.Name, "Base64") || strings.Contains(enum.Name, "Compression")) {
						fmt.Fprintf(os.Stderr, "DEBUG_CREATE: Creating enum %s with %d initial cases\n", enum.Name, len(enum.Cases))
					}
					enums = append(enums, enum)
				}
			} else if len(parts) >= 4 {
				// This is an enum case (c:@E@EnumName@CaseName)
				enumCase, enumCaseErr := occ2go.ParseEnumCase(doc)
				if enumCaseErr == nil && enumCase != nil {
					enumName := parts[2]
					if os.Getenv("DEBUG_ENUM_CASES") == "1" && (strings.Contains(doc.Metadata.ExternalID, "Base64") || strings.Contains(doc.Metadata.ExternalID, "Compression")) {
						fmt.Fprintf(os.Stderr, "DEBUG_CASES: Adding case %s to enum %s (from externalID: %s)\n",
							enumCase.Name, enumName, doc.Metadata.ExternalID)
					}
					enumCasesMap[enumName] = append(enumCasesMap[enumName], enumCase)
				}
			}
		} else if strings.HasPrefix(doc.Metadata.ExternalID, "c:@T@") {
			// This is a C typedef (e.g., typedef int CIFormat)
			typedef, typedefErr := occ2go.ParseTypedef(doc)
			if typedefErr == nil && typedef != nil {
				typedefs = append(typedefs, typedef)
			}
		} else if strings.Contains(doc.Metadata.ExternalID, "@k") && (strings.HasPrefix(doc.Metadata.ExternalID, "c:@k") || strings.HasPrefix(doc.Metadata.ExternalID, "c:@E@")) {
			// This is an extern const declaration
			constant, constErr := occ2go.ParseConstant(doc)
			if constErr == nil && constant != nil {
				constants = append(constants, constant)
			}
		} else {
			parseErrors++
		}
	}

	// Also process synthetic documents from API collection references
	for _, doc := range syntheticDocs {
		processedFiles++
		fn, cls, proto, err := occ2go.ParseDocument(doc)
		if err == nil {
			if fn != nil {
				functions = append(functions, fn)
			}
			if cls != nil {
				// Only include classes that actually belong to this framework
				if classbelongsToFramework(doc, framework) {
					classes = append(classes, cls)
				}
			}
			if proto != nil {
				protocols = append(protocols, proto)
			}
		} else {
			parseErrors++
		}
	}

	if verbose {
		fmt.Fprintf(os.Stderr, "[%s] Parsed %d files in %.2fs (%d errors)\n", framework, processedFiles, time.Since(phaseStart).Seconds(), parseErrors)
		fmt.Fprintf(os.Stderr, "[%s] Found: %d functions, %d classes, %d protocols, %d enums, %d typedefs, %d constants\n", framework, len(functions), len(classes), len(protocols), len(enums), len(typedefs), len(constants))
	}

	// Populate currentFrameworkClasses map for cross-framework type detection
	phaseStart = time.Now()
	// This helps resolve whether a type like "Number" or "AccessibilityCustomAction"
	// exists in the current framework or is from another framework
	currentFrameworkClasses = make(map[string]bool)
	for _, cls := range classes {
		// Store both the original name and the stripped name
		strippedName := stripObjCPrefix(cls.Name)
		currentFrameworkClasses[strippedName] = true
	}

	// Populate enum names for current framework type resolution
	currentFrameworkEnums = make(map[string]bool)
	for _, enum := range enums {
		if enum.Name != "" {
			strippedName := stripObjCPrefix(enum.Name)
			currentFrameworkEnums[strippedName] = true
		}
	}

	// Populate typedef names for current framework type resolution
	currentFrameworkTypedefs = make(map[string]bool)
	for _, typedef := range typedefs {
		if typedef.Name != "" {
			strippedName := stripObjCPrefix(typedef.Name)
			currentFrameworkTypedefs[strippedName] = true
		}
	}

	// Second pass: collect methods for each class
	// (Properties are already collected in the first pass)
	if len(classes) > 0 {
		classMethodsMap := make(map[string][]*occ2go.ParsedMethod)
		// Track seen selectors per class to prevent duplicates from documentation
		classSeenSelectors := make(map[string]map[string]bool)

		methodCount := 0
		for _, doc := range appledocs.Symbols(fsys, framework) {
			externalID := doc.Metadata.ExternalID

			// Check for methods: c:objc(cs)ClassName(im)methodName or c:objc(cs)ClassName(cm)methodName
			if strings.Contains(externalID, "(im)") || strings.Contains(externalID, "(cm)") {
				method, err := occ2go.ParseMethod(doc)
				if err == nil && method != nil {
					// Extract class name from external ID
					// Format: c:objc(cs)NSButton(im)initWithFrame:
					parts := strings.Split(externalID, "(")
					if len(parts) >= 2 {
						className := strings.TrimPrefix(parts[1], "cs)")

						// Initialize selector tracking for this class if needed
						if classSeenSelectors[className] == nil {
							classSeenSelectors[className] = make(map[string]bool)
						}

						// Create a unique key combining selector and method type (class vs instance)
						methodKey := method.Selector
						if method.IsClassMethod {
							methodKey = "class:" + methodKey
						} else {
							methodKey = "instance:" + methodKey
						}

						// Only add if we haven't seen this exact method before
						if !classSeenSelectors[className][methodKey] {
							classMethodsMap[className] = append(classMethodsMap[className], method)
							classSeenSelectors[className][methodKey] = true
							methodCount++
						}
					}
				}
			}
		}

		// Now add properties from references (for properties without separate files)
		// This runs AFTER parsing separate property files, so detailed property info is preferred
		var refPropertyClasses []string
		for className := range refProperties {
			refPropertyClasses = append(refPropertyClasses, className)
		}
		sort.Strings(refPropertyClasses)

		for _, className := range refPropertyClasses {
			props := refProperties[className]
			if classSeenProperties[className] == nil {
				classSeenProperties[className] = make(map[string]bool)
			}
			for _, prop := range props {
				// Only add if we haven't seen this property name before
				// Properties from separate files (added earlier) take precedence
				if !classSeenProperties[className][prop.Name] {
					classPropertiesMap[className] = append(classPropertiesMap[className], prop)
					classSeenProperties[className][prop.Name] = true
				}
			}
		}

		// Attach methods and properties to classes
		propertyCount := 0
		skippedMethodCount := 0
		skippedPropertyCount := 0
		for i := range classes {
			if methods, ok := classMethodsMap[classes[i].Name]; ok {
				// Debug: log methods for NSExtensionContext
				if classes[i].Name == "NSExtensionContext" && os.Getenv("DEBUG_HIERARCHY") == "1" {
					for _, m := range methods {
						if strings.Contains(m.Name, "Broadcast") {
							fmt.Fprintf(os.Stderr, "DEBUG main: NSExtensionContext method %s has %d params\n", m.Name, len(m.Parameters))
							for j, p := range m.Parameters {
								fmt.Fprintf(os.Stderr, "DEBUG main:   param[%d] name=%s type=%s\n", j, p.Name, p.Type)
							}
						}
					}
				}
				// Relax parameters that violate hierarchy (map to objectivec.IObject)
				RelaxMethodParameters(methods, framework)
				// Filter out methods that would create upward dependency violations
				originalCount := len(methods)
				classes[i].Methods = FilterMethodsByHierarchy(methods, framework)
				skippedMethodCount += (originalCount - len(classes[i].Methods))
			}
			if properties, ok := classPropertiesMap[classes[i].Name]; ok {
				// Filter out properties that would create upward dependency violations
				originalCount := len(properties)
				classes[i].Properties = FilterPropertiesByHierarchy(properties, framework)
				skippedPropertyCount += (originalCount - len(classes[i].Properties))
				propertyCount += len(classes[i].Properties)
			}
		}

		if verbose {
			fmt.Fprintf(os.Stderr, "Found %d methods and %d properties for %d classes\n", methodCount, propertyCount, len(classes))
			if skippedMethodCount > 0 || skippedPropertyCount > 0 {
				fmt.Fprintf(os.Stderr, "Skipped %d methods and %d properties due to framework hierarchy violations\n",
					skippedMethodCount, skippedPropertyCount)
			}
		}
	}

	// Attach enum cases to enums
	if len(enums) > 0 {
		caseCount := 0
		for i := range enums {
			if cases, ok := enumCasesMap[enums[i].Name]; ok {
				// Make a copy of the cases slice to avoid shared slice references
				// that would cause issues during enum deduplication
				enumCases := make([]*occ2go.ParsedEnumCase, len(cases))
				copy(enumCases, cases)
				enums[i].Cases = enumCases
				caseCount += len(cases)

				if os.Getenv("DEBUG_ENUM_ATTACH") == "1" && (strings.Contains(enums[i].Name, "Base64") || strings.Contains(enums[i].Name, "Compression")) {
					fmt.Fprintf(os.Stderr, "DEBUG_ATTACH: Attaching %d cases to enum %s\n", len(cases), enums[i].Name)
					for _, c := range cases {
						fmt.Fprintf(os.Stderr, "DEBUG_ATTACH:   - %s\n", c.Name)
					}
				}
			}
		}
		if verbose {
			fmt.Fprintf(os.Stderr, "Found %d enum cases for %d enums\n", caseCount, len(enums))
		}

		// Enrich enum values from macOS SDK headers using extract-enum-values tool
		enrichStart := time.Now()
		_ = enrichEnumValues(framework, enums, verbose)
		if verbose {
			fmt.Fprintf(os.Stderr, "[%s] Enriched enums in %.2fs\n", framework, time.Since(enrichStart).Seconds())
		}
	}

	// Build type registry from parsed data (source of truth)
	// This populates crossFrameworkTypeRegistry with mappings like:
	//   NSImageScaling → appkit.ImageScaling
	//   NSWindow → appkit.Window
	buildTypeRegistryFromParsedData(framework, classes, enums, typedefs)
	if verbose {
		fmt.Fprintf(os.Stderr, "Built type registry from parsed data: %d classes, %d enums, %d typedefs\n",
			len(classes), len(enums), len(typedefs))
	}

	// Fail if no symbols were found and no filter was applied
	if len(functions) == 0 && len(classes) == 0 && len(protocols) == 0 && filterRegexp == "" && processedFiles == 0 {
		return fmt.Errorf("no symbols found for framework %s", framework)
	}

	// Deduplicate functions by name (keep first occurrence)
	seenFunctions := make(map[string]bool)
	uniqueFunctions := make([]*occ2go.ParsedFunction, 0, len(functions))
	for _, fn := range functions {
		if !seenFunctions[fn.Name] {
			seenFunctions[fn.Name] = true
			uniqueFunctions = append(uniqueFunctions, fn)
		}
	}
	if len(uniqueFunctions) < len(functions) {
		if verbose {
			fmt.Fprintf(os.Stderr, "Deduplicated functions: %d -> %d\n", len(functions), len(uniqueFunctions))
		}
		functions = uniqueFunctions
	}

	// Apply regexp filter if specified
	if filterRegexp != "" {
		re, err := regexp.Compile(filterRegexp)
		if err != nil {
			return fmt.Errorf("invalid filter regexp: %w", err)
		}

		filteredFunctions := make([]*occ2go.ParsedFunction, 0)
		for _, fn := range functions {
			if re.MatchString(fn.Name) {
				filteredFunctions = append(filteredFunctions, fn)
			}
		}

		filteredClasses := make([]*occ2go.ParsedClass, 0)
		for _, cls := range classes {
			if re.MatchString(cls.Name) {
				filteredClasses = append(filteredClasses, cls)
			}
		}

		filteredProtocols := make([]*occ2go.ParsedProtocol, 0)
		for _, proto := range protocols {
			if re.MatchString(proto.Name) {
				filteredProtocols = append(filteredProtocols, proto)
			}
		}

		if verbose {
			fmt.Fprintf(os.Stderr, "Filter '%s' matched: %d/%d functions, %d/%d classes, %d/%d protocols\n",
				filterRegexp,
				len(filteredFunctions), len(functions),
				len(filteredClasses), len(classes),
				len(filteredProtocols), len(protocols))
		}

		functions = filteredFunctions
		classes = filteredClasses
		protocols = filteredProtocols
	}

	// Create output directory
	packageName := strings.ToLower(framework)

	// Check if outputDir already ends with the package name to avoid double-nesting
	// (e.g., when running from generated/appkit/ with -output .)
	outDir := outputDir
	if filepath.Base(outputDir) != packageName {
		outDir = filepath.Join(outputDir, packageName)
	}

	if !txtarOutput {
		if err := os.MkdirAll(outDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	// Generate bindings
	phaseStart = time.Now()
	if verbose {
		fmt.Fprintf(os.Stderr, "[%s] Starting code generation\n", framework)
	}
	if txtarOutput {
		if err := generateTxtar(os.Stdout, framework, packageName, inputDir, functions, classes, protocols, enums, typedefs, constants, withRefMethods, generateTests, generateExamples, variant); err != nil {
			return fmt.Errorf("failed to generate bindings: %w", err)
		}
	} else {
		if err := generateFiles(outDir, framework, packageName, inputDir, functions, classes, protocols, enums, typedefs, constants, withRefMethods, generateTests, generateExamples, variant); err != nil {
			return fmt.Errorf("failed to generate bindings: %w", err)
		}
		if verbose {
			fmt.Fprintf(os.Stderr, "[%s] Generated code in %.2fs\n", framework, time.Since(phaseStart).Seconds())
		}
		fmt.Printf("Generated %s bindings in %s\n", framework, outDir)
	}

	if verbose {
		fmt.Fprintf(os.Stderr, "[%s] Total time: %.2fs\n", framework, time.Since(startTime).Seconds())
	}

	return nil
}

func main() {
	framework := flag.String("framework", "CoreGraphics", "Framework to generate bindings for (supports regexp patterns like 'Core.*' or '^(AppKit|Foundation)$')")
	inputDir := flag.String("input", "", "Input directory with JSON files (defaults to ~/.appledocs/cache/developer.apple.com/tutorials/data/documentation)")
	outputDir := flag.String("output", "generated", "Output directory for generated bindings")
	filterRegexp := flag.String("filter", "", "Only generate symbols matching this regexp (e.g., '^CGRect' or '^NS(Window|View)')")
	txtarOutput := flag.Bool("txtar", false, "Output as txtar format to stdout instead of files")
	variant := flag.String("variant", "", "Comma-separated template variants (e.g., 'darwinkit' or 'base,ref-methods'). Later variants override earlier ones.")
	withRefMethods := flag.Bool("with-ref-methods", false, "Generate struct-wrapped Ref types with methods (enables method-style API)")
	generateTests := flag.Bool("generate-tests", false, "Generate test files for the bindings")
	generateExamples := flag.Bool("generate-examples", false, "Generate example code demonstrating API usage")
	generateObjcRuntime := flag.Bool("generate-objc-runtime", false, "Generate only the objc runtime package (framework-independent)")
	verboseFlag := flag.Bool("v", false, "Enable verbose output")
	flag.Parse()

	verbose = *verboseFlag

	// Handle objc runtime generation
	if *generateObjcRuntime {
		objcDir := filepath.Join(*outputDir, "objc")
		if err := generateObjcRuntimePackage(objcDir); err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to generate objc runtime: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Generated objc runtime package in %s\n", objcDir)
		return
	}

	// Default to cache directory if not specified
	if *inputDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to get home directory: %v\n", err)
			os.Exit(1)
		}
		*inputDir = filepath.Join(homeDir, ".appledocs/cache/developer.apple.com/tutorials/data/documentation")
	}

	// Discover available frameworks by looking for .json files in the input directory
	frameworks, err := discoverFrameworks(*inputDir, *framework)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to discover frameworks: %v\n", err)
		os.Exit(1)
	}

	if len(frameworks) == 0 {
		fmt.Fprintf(os.Stderr, "Error: no frameworks found matching pattern '%s'\n", *framework)
		os.Exit(1)
	}

	if verbose {
		if len(frameworks) == 1 {
			fmt.Fprintf(os.Stderr, "Generating bindings for %s\n", frameworks[0])
		} else {
			fmt.Fprintf(os.Stderr, "Generating bindings for %d frameworks matching '%s': %v\n", len(frameworks), *framework, frameworks)
		}
		fmt.Fprintf(os.Stderr, "Input: %s\n", *inputDir)
		fmt.Fprintf(os.Stderr, "Output: %s\n", *outputDir)
	}

	// Initialize framework registry
	baseModule := "github.com/tmc/appledocs/generated"
	if err := initializeFrameworkRegistry(baseModule, *outputDir); err == nil && verbose && globalRegistry != nil {
		allFrameworks := globalRegistry.All()
		fmt.Fprintf(os.Stderr, "Initialized framework registry with %d frameworks\n", len(allFrameworks))
	}

	// Build cross-framework type registry for proper type resolution
	if err := buildCrossFrameworkTypeRegistry(*outputDir); err == nil && verbose && len(crossFrameworkTypeRegistry) > 0 {
		fmt.Fprintf(os.Stderr, "Built cross-framework type registry with %d types\n", len(crossFrameworkTypeRegistry))
	}

	// Generate bindings for each matching framework
	for _, fw := range frameworks {
		if err := generateFramework(fw, *inputDir, *outputDir, *filterRegexp, *txtarOutput, *variant, *withRefMethods, *generateTests, *generateExamples); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating %s: %v\n", fw, err)
			os.Exit(1)
		}
	}
}

// generateFiles generates all files to disk
