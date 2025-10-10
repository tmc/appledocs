// main_improved.go - Enhanced version of the generator with better error handling,
// support for blocks/methods/properties, and improved documentation generation.
//
// To use: rename this file to main.go (after backing up the original)

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Generator configuration
type GeneratorConfig struct {
	Framework       string
	InputDir        string
	OutputDir       string
	Style           string
	Verbose         bool
	StrictMode      bool // Fail on any errors
	SkipValidation  bool
	SwiftInterop    bool // EXPERIMENTAL: Generate Swift interop bindings (purego vs cgo)
}

// GeneratorResult contains the results of generation
type GeneratorResult struct {
	Functions   []*ParsedFunction
	Classes     []*ParsedClass
	Protocols   []*ParsedProtocol
	Methods     []*ParsedMethod
	Properties  []*ParsedProperty
	Blocks      []*ParsedBlock
	Errors      *ErrorCollector
	FilesTotal  int
	FilesParsed int
}

func mainImproved() {
	config := parseFlags()

	log.Printf("Generating %s bindings for %s framework", config.Style, config.Framework)

	// Find framework directory
	frameworkDir := filepath.Join(config.InputDir, config.Framework)
	if _, err := os.Stat(frameworkDir); os.IsNotExist(err) {
		log.Fatalf("Framework directory not found: %s", frameworkDir)
	}

	// Parse all JSON files with error collection
	result := parseFrameworkDirectory(frameworkDir, config)

	// Report parsing results
	log.Printf("Parsed %d/%d JSON files in %s", result.FilesParsed, result.FilesTotal, config.Framework)
	log.Printf("Found: %d functions, %d classes, %d protocols, %d methods, %d properties, %d blocks",
		len(result.Functions), len(result.Classes), len(result.Protocols),
		len(result.Methods), len(result.Properties), len(result.Blocks))

	// Report errors if any
	if result.Errors.HasErrors() {
		log.Printf("Encountered %d errors during parsing:", result.Errors.Count())
		if config.Verbose {
			fmt.Println(result.Errors.Summary())
		} else {
			stats := result.Errors.Stats()
			for stage, count := range stats.ByStage {
				log.Printf("  %s: %d errors", stage, count)
			}
			log.Println("Run with -verbose to see detailed error messages")
		}

		if config.StrictMode {
			log.Fatalf("Aborting due to errors in strict mode")
		}
	}

	// Create output directory
	outDir := filepath.Join(config.OutputDir, "frameworks", strings.ToLower(config.Framework))
	if err := os.MkdirAll(outDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Generate bindings based on style
	switch config.Style {
	case "purego":
		generateEnhancedPuregoBindings(result, outDir, config.Framework)
	case "darwinkit":
		log.Printf("Darwinkit style not yet implemented, falling back to purego")
		generateEnhancedPuregoBindings(result, outDir, config.Framework)
	case "simple":
		log.Printf("Simple style not yet implemented, falling back to purego")
		generateEnhancedPuregoBindings(result, outDir, config.Framework)
	default:
		log.Fatalf("Unknown style: %s", config.Style)
	}

	log.Printf("Generated bindings in %s", outDir)
}

func parseFlags() GeneratorConfig {
	config := GeneratorConfig{}

	flag.StringVar(&config.Framework, "framework", "CoreGraphics", "Framework to generate bindings for")
	flag.StringVar(&config.InputDir, "input", "output/tutorials/data/documentation", "Input directory with JSON files")
	flag.StringVar(&config.OutputDir, "output", "generated", "Output directory for generated bindings")
	flag.StringVar(&config.Style, "style", "purego", "Binding style: purego, darwinkit, or simple")
	flag.BoolVar(&config.Verbose, "verbose", false, "Enable verbose output including all errors")
	flag.BoolVar(&config.StrictMode, "strict", false, "Fail on any parsing errors")
	flag.BoolVar(&config.SkipValidation, "skip-validation", false, "Skip validation of generated code")
	flag.BoolVar(&config.SwiftInterop, "swift-interop", false, "EXPERIMENTAL: Generate Swift interop bindings using purego")

	flag.Parse()

	return config
}

// parseFrameworkDirectory walks the framework directory and parses all JSON files
func parseFrameworkDirectory(frameworkDir string, config GeneratorConfig) *GeneratorResult {
	result := &GeneratorResult{
		Functions:  []*ParsedFunction{},
		Classes:    []*ParsedClass{},
		Protocols:  []*ParsedProtocol{},
		Methods:    []*ParsedMethod{},
		Properties: []*ParsedProperty{},
		Blocks:     []*ParsedBlock{},
		Errors:     &ErrorCollector{},
	}

	err := filepath.Walk(frameworkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			result.Errors.Add(path, "", "filesystem", err)
			return nil // Continue walking
		}

		if info.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}

		result.FilesTotal++

		// Parse the JSON file
		symbols, parseErr := processJSONFileEnhanced(path)
		if parseErr != nil {
			result.Errors.Add(path, "", "parse", parseErr)
			return nil // Continue despite errors
		}

		result.FilesParsed++

		// Collect symbols
		if symbols.Function != nil {
			result.Functions = append(result.Functions, symbols.Function)
		}
		if symbols.Class != nil {
			result.Classes = append(result.Classes, symbols.Class)
		}
		if symbols.Protocol != nil {
			result.Protocols = append(result.Protocols, symbols.Protocol)
		}
		if symbols.Method != nil {
			result.Methods = append(result.Methods, symbols.Method)
		}
		if symbols.Property != nil {
			result.Properties = append(result.Properties, symbols.Property)
		}
		if symbols.Block != nil {
			result.Blocks = append(result.Blocks, symbols.Block)
		}

		return nil
	})

	if err != nil {
		result.Errors.Add(frameworkDir, "", "walk", err)
	}

	return result
}

// ParsedSymbols contains all symbol types that can be parsed from a JSON file
type ParsedSymbols struct {
	Function *ParsedFunction
	Class    *ParsedClass
	Protocol *ParsedProtocol
	Method   *ParsedMethod
	Property *ParsedProperty
	Block    *ParsedBlock
}

// processJSONFileEnhanced processes a single JSON file with enhanced parsing
func processJSONFileEnhanced(path string) (*ParsedSymbols, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var doc AppleDoc
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	symbols := &ParsedSymbols{}

	externalID := doc.Metadata.ExternalID
	availability := extractAvailability(doc.Metadata.Platforms)

	// Try to get ObjectiveC variant first
	tokens := getObjectiveCVariant(&doc)
	if tokens == nil {
		// Fall back to primary declarations
		if len(doc.PrimaryContentSections) > 0 {
			for _, section := range doc.PrimaryContentSections {
				if len(section.Declarations) > 0 && len(section.Declarations[0].Tokens) > 0 {
					tokens = section.Declarations[0].Tokens
					break
				}
			}
		}
	}

	if tokens == nil || len(tokens) == 0 {
		return nil, fmt.Errorf("no declaration found for %s", externalID)
	}

	// Determine symbol type by external ID prefix
	switch {
	case strings.HasPrefix(externalID, "c:@F@"):
		// C function
		fn, err := parseDeclaration(tokens)
		if err != nil {
			return nil, fmt.Errorf("failed to parse function: %w", err)
		}
		fn.Availability = availability
		symbols.Function = fn

	case strings.HasPrefix(externalID, "c:objc(cs)"):
		// Objective-C class
		cls := parseClassDeclaration(tokens)
		if cls == nil {
			return nil, fmt.Errorf("failed to parse class declaration")
		}
		cls.Availability = availability
		symbols.Class = cls

	case strings.HasPrefix(externalID, "c:objc(pl)"):
		// Objective-C protocol
		proto := parseProtocolDeclaration(tokens)
		if proto == nil {
			return nil, fmt.Errorf("failed to parse protocol declaration")
		}
		proto.Availability = availability
		symbols.Protocol = proto

	case strings.HasPrefix(externalID, "c:objc(cm)") || strings.HasPrefix(externalID, "c:objc(im)"):
		// Objective-C method (class or instance)
		className := extractClassNameFromExternalID(externalID)
		method := parseMethodDeclaration(tokens, className)
		if method != nil {
			method.Availability = availability
			symbols.Method = method
		}

	case strings.HasPrefix(externalID, "c:objc(py)"):
		// Objective-C property
		className := extractClassNameFromExternalID(externalID)
		prop := parsePropertyDeclaration(tokens, className)
		if prop != nil {
			prop.Availability = availability
			symbols.Property = prop
		}

	case strings.Contains(externalID, "^"):
		// Block type (heuristic - may need refinement)
		block := parseBlockTypedef(tokens)
		if block != nil {
			block.Availability = availability
			symbols.Block = block
		}

	default:
		return nil, fmt.Errorf("unsupported symbol type: %s", externalID)
	}

	return symbols, nil
}

// extractClassNameFromExternalID extracts the class name from an external ID
// Example: "c:objc(cs)NSString(im)init" -> "NSString"
func extractClassNameFromExternalID(externalID string) string {
	// Remove the prefix
	id := externalID
	if strings.HasPrefix(id, "c:objc(cs)") {
		id = strings.TrimPrefix(id, "c:objc(cs)")
	} else if strings.HasPrefix(id, "c:objc(cm)") {
		id = strings.TrimPrefix(id, "c:objc(cm)")
	} else if strings.HasPrefix(id, "c:objc(im)") {
		id = strings.TrimPrefix(id, "c:objc(im)")
	} else if strings.HasPrefix(id, "c:objc(py)") {
		id = strings.TrimPrefix(id, "c:objc(py)")
	}

	// Extract class name (before the next parenthesis or end)
	if idx := strings.Index(id, "("); idx != -1 {
		return id[:idx]
	}

	return id
}

// generateEnhancedPuregoBindings generates improved purego bindings
func generateEnhancedPuregoBindings(result *GeneratorResult, outputDir, framework string) {
	pkgName := strings.ToLower(framework)

	// Generate doc.go with enhanced documentation
	generateEnhancedDocFile(outputDir, pkgName, framework, result)

	// Generate types.gen.go with discovered types
	generateEnhancedTypesFile(outputDir, pkgName, framework, result)

	// Generate loader.gen.go (same as before)
	generateLoaderFile(outputDir, pkgName, framework)

	// Generate functions.gen.go with enhanced documentation
	generateEnhancedFunctionsFile(outputDir, pkgName, framework, result.Functions)

	// Generate classes.gen.go
	if len(result.Classes) > 0 {
		generateEnhancedClassesFile(outputDir, pkgName, framework, result.Classes)
	}

	// Generate protocols.gen.go
	if len(result.Protocols) > 0 {
		generateEnhancedProtocolsFile(outputDir, pkgName, framework, result.Protocols)
	}

	// Generate methods.gen.go (NEW)
	if len(result.Methods) > 0 {
		generateMethodsFile(outputDir, pkgName, framework, result.Methods)
	}

	// Generate properties.gen.go (NEW)
	if len(result.Properties) > 0 {
		generatePropertiesFile(outputDir, pkgName, framework, result.Properties)
	}

	// Generate blocks.gen.go (NEW)
	if len(result.Blocks) > 0 {
		generateBlocksFile(outputDir, pkgName, framework, result.Blocks)
	}

	fileCount := 4 // doc, types, loader, functions
	if len(result.Classes) > 0 {
		fileCount++
	}
	if len(result.Protocols) > 0 {
		fileCount++
	}
	if len(result.Methods) > 0 {
		fileCount++
	}
	if len(result.Properties) > 0 {
		fileCount++
	}
	if len(result.Blocks) > 0 {
		fileCount++
	}

	log.Printf("Generated %d .gen.go files in %s", fileCount, outputDir)
}

// Placeholder for enhanced generation functions
// These would use the Documentation and formatting utilities from docgen.go

func generateEnhancedDocFile(outputDir, pkgName, framework string, result *GeneratorResult) {
	// Enhanced version would include statistics about all symbol types
	generateDocFile(outputDir, pkgName, framework, result.Functions)
}

func generateEnhancedTypesFile(outputDir, pkgName, framework string, result *GeneratorResult) {
	// Enhanced version would discover types from all parsed symbols
	generateTypesFile(outputDir, pkgName, framework)
}

func generateEnhancedFunctionsFile(outputDir, pkgName, framework string, functions []*ParsedFunction) {
	// Enhanced version would use Documentation extraction
	generateFunctionsFile(outputDir, pkgName, framework, functions)
}

func generateEnhancedClassesFile(outputDir, pkgName, framework string, classes []*ParsedClass) {
	// Enhanced version with better documentation
	generateClassesFile(outputDir, pkgName, framework, classes)
}

func generateEnhancedProtocolsFile(outputDir, pkgName, framework string, protocols []*ParsedProtocol) {
	// Enhanced version with protocol requirements
	generateProtocolsFile(outputDir, pkgName, framework, protocols)
}

// NEW: Generate methods file
func generateMethodsFile(outputDir, pkgName, framework string, methods []*ParsedMethod) {
	filename := filepath.Join(outputDir, "methods.gen.go")
	f, err := os.Create(filename)
	if err != nil {
		log.Printf("Warning: Failed to create methods file: %v", err)
		return
	}
	defer f.Close()

	fmt.Fprintf(f, "// Code generated from Apple documentation for %s. DO NOT EDIT.\n\n", framework)
	fmt.Fprintf(f, "package %s\n\n", pkgName)

	fmt.Fprintf(f, "// %s Methods\n", framework)
	fmt.Fprintf(f, "//\n")
	fmt.Fprintf(f, "// This file contains Objective-C method declarations.\n")
	fmt.Fprintf(f, "// To use these methods, you need to use the Objective-C runtime (objc package).\n")
	fmt.Fprintf(f, "\n")

	fmt.Fprintf(f, "// Discovered methods (%d total):\n\n", len(methods))

	for _, method := range methods {
		if method.Selector == "" {
			continue
		}

		fmt.Fprintf(f, "// %s\n", FormatMethodSignature(method))
		fmt.Fprintf(f, "// Class: %s\n", method.ClassName)
		fmt.Fprintf(f, "// Go name: %s\n", SelectorToGoName(method.Selector))

		// Add availability if present
		if !method.Availability.IsEmpty() {
			fmt.Fprintf(f, "//\n// Availability:\n")
			for _, platform := range method.Availability.Platforms() {
				version := method.Availability.IntroducedAt[platform]
				fmt.Fprintf(f, "//   - %s %s+\n", platform, version)
			}
		}

		fmt.Fprintf(f, "\n")
	}
}

// NEW: Generate properties file
func generatePropertiesFile(outputDir, pkgName, framework string, properties []*ParsedProperty) {
	filename := filepath.Join(outputDir, "properties.gen.go")
	f, err := os.Create(filename)
	if err != nil {
		log.Printf("Warning: Failed to create properties file: %v", err)
		return
	}
	defer f.Close()

	fmt.Fprintf(f, "// Code generated from Apple documentation for %s. DO NOT EDIT.\n\n", framework)
	fmt.Fprintf(f, "package %s\n\n", pkgName)

	fmt.Fprintf(f, "// %s Properties\n", framework)
	fmt.Fprintf(f, "//\n")
	fmt.Fprintf(f, "// This file contains Objective-C property declarations.\n")
	fmt.Fprintf(f, "\n")

	fmt.Fprintf(f, "// Discovered properties (%d total):\n\n", len(properties))

	for _, prop := range properties {
		if prop.Name == "" {
			continue
		}

		fmt.Fprintf(f, "// @property")
		if len(prop.Attributes) > 0 {
			fmt.Fprintf(f, " (%s)", strings.Join(prop.Attributes, ", "))
		}
		fmt.Fprintf(f, " %s %s\n", prop.Type, prop.Name)
		fmt.Fprintf(f, "// Class: %s\n", prop.ClassName)

		// Add availability if present
		if !prop.Availability.IsEmpty() {
			fmt.Fprintf(f, "//\n// Availability:\n")
			for _, platform := range prop.Availability.Platforms() {
				version := prop.Availability.IntroducedAt[platform]
				fmt.Fprintf(f, "//   - %s %s+\n", platform, version)
			}
		}

		fmt.Fprintf(f, "\n")
	}
}

// NEW: Generate blocks file
func generateBlocksFile(outputDir, pkgName, framework string, blocks []*ParsedBlock) {
	filename := filepath.Join(outputDir, "blocks.gen.go")
	f, err := os.Create(filename)
	if err != nil {
		log.Printf("Warning: Failed to create blocks file: %v", err)
		return
	}
	defer f.Close()

	fmt.Fprintf(f, "// Code generated from Apple documentation for %s. DO NOT EDIT.\n\n", framework)
	fmt.Fprintf(f, "package %s\n\n", pkgName)

	fmt.Fprintf(f, "// %s Block Types\n", framework)
	fmt.Fprintf(f, "//\n")
	fmt.Fprintf(f, "// This file contains Objective-C block type declarations.\n")
	fmt.Fprintf(f, "// Blocks are Objective-C closures - use carefully with purego.\n")
	fmt.Fprintf(f, "\n")

	fmt.Fprintf(f, "// Discovered block types (%d total):\n\n", len(blocks))

	for _, block := range blocks {
		if block.Name == "" {
			continue
		}

		fmt.Fprintf(f, "// %s\n", block.Name)
		fmt.Fprintf(f, "// Return type: %s\n", block.ReturnType)

		if len(block.Parameters) > 0 {
			fmt.Fprintf(f, "// Parameters:\n")
			for _, param := range block.Parameters {
				fmt.Fprintf(f, "//   %s: %s\n", param.Name, param.Type)
			}
		}

		// Suggest Go function signature
		fmt.Fprintf(f, "// Go type: func(")
		for i, param := range block.Parameters {
			if i > 0 {
				fmt.Fprintf(f, ", ")
			}
			goType := TypeToGo(param.Type)
			fmt.Fprintf(f, "%s", goType)
		}
		fmt.Fprintf(f, ")")

		if block.ReturnType != "" && block.ReturnType != "void" {
			goReturnType := TypeToGo(block.ReturnType)
			fmt.Fprintf(f, " %s", goReturnType)
		}
		fmt.Fprintf(f, "\n")

		// Add availability if present
		if !block.Availability.IsEmpty() {
			fmt.Fprintf(f, "//\n// Availability:\n")
			for _, platform := range block.Availability.Platforms() {
				version := block.Availability.IntroducedAt[platform]
				fmt.Fprintf(f, "//   - %s %s+\n", platform, version)
			}
		}

		fmt.Fprintf(f, "\n")
	}
}
