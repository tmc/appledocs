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
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	framework := flag.String("framework", "CoreGraphics", "Framework to generate bindings for")
	inputDir := flag.String("input", "output/tutorials/data/documentation", "Input directory with JSON files")
	outputDir := flag.String("output", "generated", "Output directory for generated bindings")
	style := flag.String("style", "purego", "Binding style: purego, darwinkit, or simple")
	flag.Parse()

	log.Printf("Generating %s bindings for %s framework", *style, *framework)

	// Find framework directory
	frameworkDir := filepath.Join(*inputDir, *framework)
	if _, err := os.Stat(frameworkDir); os.IsNotExist(err) {
		log.Fatalf("Framework directory not found: %s", frameworkDir)
	}

	// Parse all JSON files in the framework directory
	var functions []*ParsedFunction
	var classes []*ParsedClass
	var protocols []*ParsedProtocol

	totalFiles := 0
	processedFiles := 0
	err := filepath.Walk(frameworkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".json") {
			totalFiles++
			fn, cls, proto, err := processJSONFile(path)
			if err == nil {
				processedFiles++
				if fn != nil {
					functions = append(functions, fn)
				}
				if cls != nil {
					classes = append(classes, cls)
				}
				if proto != nil {
					protocols = append(protocols, proto)
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking framework directory: %v", err)
	}

	log.Printf("Processed %d/%d JSON files in %s", processedFiles, totalFiles, *framework)
	log.Printf("Found %d functions, %d classes, %d protocols", len(functions), len(classes), len(protocols))

	// Create output directory
	outDir := filepath.Join(*outputDir, strings.ToLower(*framework))
	if err := os.MkdirAll(outDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Generate bindings based on style
	switch *style {
	case "purego":
		generatePuregoBindings(functions, classes, protocols, outDir, *framework)
	case "darwinkit":
		generateDarkwinKitBindings(functions, classes, protocols, outDir, *framework)
	case "simple":
		generateSimpleBindings(functions, classes, protocols, outDir, *framework)
	default:
		log.Fatalf("Unknown style: %s", *style)
	}

	log.Printf("Generated bindings in %s", outDir)
}

type ParsedFunction struct {
	Name         string
	ReturnType   string
	Parameters   []Parameter
	Comment      string
	Availability Availability
}

type ParsedClass struct {
	Name         string
	SuperClass   string
	Comment      string
	Availability Availability
}

type ParsedProtocol struct {
	Name         string
	Comment      string
	Availability Availability
}

type Parameter struct {
	Name string
	Type string
}

// Availability contains version information for API availability across platforms.
// Uses maps for flexibility - handles new platforms without code changes.
type Availability struct {
	// IntroducedAt maps platform name to version string (e.g., "macOS" -> "10.14")
	IntroducedAt map[string]string

	// DeprecatedAt maps platform name to deprecation version
	DeprecatedAt map[string]string

	Beta bool
}

// IsEmpty returns true if no version information is available
func (a *Availability) IsEmpty() bool {
	return len(a.IntroducedAt) == 0 && len(a.DeprecatedAt) == 0
}

// Platforms returns a sorted list of platforms with availability info
func (a *Availability) Platforms() []string {
	platforms := make([]string, 0, len(a.IntroducedAt))
	for p := range a.IntroducedAt {
		platforms = append(platforms, p)
	}
	sort.Strings(platforms)
	return platforms
}

type AppleDoc struct {
	Metadata struct {
		ExternalID string     `json:"externalID"`
		Title      string     `json:"title"`
		Platforms  []Platform `json:"platforms,omitempty"`
	} `json:"metadata"`
	PrimaryContentSections []struct {
		Kind         string `json:"kind"`
		Declarations []struct {
			Tokens []Token `json:"tokens"`
		} `json:"declarations"`
	} `json:"primaryContentSections"`
	VariantOverrides []struct {
		Traits []struct {
			InterfaceLanguage string `json:"interfaceLanguage"`
		} `json:"traits"`
		Patch []struct {
			Op    string          `json:"op"`
			Path  string          `json:"path"`
			Value json.RawMessage `json:"value"`
		} `json:"patch"`
	} `json:"variantOverrides"`
}

// Platform describes platform availability
type Platform struct {
	Name         string `json:"name"`
	IntroducedAt string `json:"introducedAt,omitempty"`
	DeprecatedAt string `json:"deprecatedAt,omitempty"`
	Beta         bool   `json:"beta"`
	Deprecated   bool   `json:"deprecated,omitempty"`
	Unavailable  bool   `json:"unavailable,omitempty"`
}

type Token struct {
	Kind string `json:"kind"`
	Text string `json:"text"`
}

func processJSONFile(path string) (*ParsedFunction, *ParsedClass, *ParsedProtocol, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, nil, err
	}

	var doc AppleDoc
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, nil, nil, err
	}

	externalID := doc.Metadata.ExternalID
	availability := extractAvailability(doc.Metadata.Platforms)

	// Try to get ObjectiveC variant first
	tokens := getObjectiveCVariant(&doc)
	if tokens == nil {
		// Fall back to primary declarations
		if len(doc.PrimaryContentSections) > 0 {
			// Find a section with declarations
			for _, section := range doc.PrimaryContentSections {
				if len(section.Declarations) > 0 && len(section.Declarations[0].Tokens) > 0 {
					tokens = section.Declarations[0].Tokens
					break
				}
			}
		}
	}

	if tokens == nil || len(tokens) == 0 {
		return nil, nil, nil, fmt.Errorf("no declaration found for %s", externalID)
	}

	// Determine symbol type by external ID prefix
	switch {
	case strings.HasPrefix(externalID, "c:@F@"):
		// C function
		fn, err := parseDeclaration(tokens)
		if err != nil {
			return nil, nil, nil, err
		}
		fn.Availability = availability
		return fn, nil, nil, nil

	case strings.HasPrefix(externalID, "c:objc(cs)"):
		// Objective-C class
		cls := parseClassDeclaration(tokens)
		if cls == nil {
			return nil, nil, nil, fmt.Errorf("failed to parse class declaration")
		}
		cls.Availability = availability
		return nil, cls, nil, nil

	case strings.HasPrefix(externalID, "c:objc(pl)"):
		// Objective-C protocol
		proto := parseProtocolDeclaration(tokens)
		if proto == nil {
			return nil, nil, nil, fmt.Errorf("failed to parse protocol declaration")
		}
		proto.Availability = availability
		return nil, nil, proto, nil

	default:
		return nil, nil, nil, fmt.Errorf("unsupported symbol type: %s", externalID)
	}
}

// extractAvailability converts Platform metadata to Availability.
// Returns availability information for all platforms found in the metadata.
func extractAvailability(platforms []Platform) Availability {
	avail := Availability{
		IntroducedAt: make(map[string]string),
		DeprecatedAt: make(map[string]string),
	}

	for _, p := range platforms {
		if p.Unavailable {
			continue
		}

		if p.IntroducedAt != "" {
			avail.IntroducedAt[p.Name] = p.IntroducedAt
		}

		if p.DeprecatedAt != "" {
			avail.DeprecatedAt[p.Name] = p.DeprecatedAt
		}

		if p.Beta {
			avail.Beta = true
		}
	}

	return avail
}

func getObjectiveCVariant(doc *AppleDoc) []Token {
	for _, variant := range doc.VariantOverrides {
		for _, trait := range variant.Traits {
			if trait.InterfaceLanguage == "occ" {
				for _, patch := range variant.Patch {
					if patch.Op != "replace" {
						continue
					}

					// Check for direct tokens replacement
					if strings.Contains(patch.Path, "declarations/0/tokens") {
						var tokens []Token
						if err := json.Unmarshal(patch.Value, &tokens); err == nil {
							return tokens
						}
					}

					// Check for primaryContentSections replacement (common for functions)
					if patch.Path == "/primaryContentSections/0" {
						var section struct {
							Declarations []struct {
								Languages []string `json:"languages"`
								Platforms []string `json:"platforms"`
								Tokens    []Token  `json:"tokens"`
							} `json:"declarations"`
						}
						if err := json.Unmarshal(patch.Value, &section); err == nil {
							if len(section.Declarations) > 0 && len(section.Declarations[0].Tokens) > 0 {
								return section.Declarations[0].Tokens
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func parseDeclaration(tokens []Token) (*ParsedFunction, error) {
	fn := &ParsedFunction{
		Parameters: []Parameter{},
	}

	// Filter out Swift-only declarations
	// Look for Swift-specific keywords: class func, static func, var, ->
	for _, tok := range tokens {
		if tok.Kind == "keyword" && (tok.Text == "class" || tok.Text == "static" || tok.Text == "var") {
			return nil, fmt.Errorf("skipping Swift declaration with keyword: %s", tok.Text)
		}
		if tok.Text == "->" {
			return nil, fmt.Errorf("skipping Swift declaration with -> syntax")
		}
	}

	// C function format: [extern] <returnType> <functionName>(<params>);
	// Parse: skip "extern" if present, collect return type, get function name, parse parameters

	i := 0
	// Skip "extern" keyword
	if i < len(tokens) && tokens[i].Kind == "keyword" && tokens[i].Text == "extern" {
		i++
		// Skip whitespace after extern
		for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
			i++
		}
	}

	// Collect return type (everything before function name/opening paren)
	// Return type ends when we hit an identifier followed by "("
	returnTypeParts := []string{}
	for i < len(tokens) {
		// Look ahead to see if next non-whitespace token is "("
		if tokens[i].Kind == "identifier" || tokens[i].Kind == "typeIdentifier" {
			// Check if this might be the function name
			j := i + 1
			for j < len(tokens) && tokens[j].Kind == "text" && strings.TrimSpace(tokens[j].Text) == "" {
				j++
			}
			if j < len(tokens) && tokens[j].Text == "(" {
				// This identifier is the function name
				fn.Name = tokens[i].Text
				i = j + 1 // Move past "("
				break
			}
		}

		// Part of return type - collect non-whitespace tokens
		if tokens[i].Text != "" && strings.TrimSpace(tokens[i].Text) != "" {
			returnTypeParts = append(returnTypeParts, tokens[i].Text)
		}
		i++
	}

	fn.ReturnType = strings.Join(returnTypeParts, " ")

	// Parse parameters
	for i < len(tokens) {
		// Skip whitespace
		for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
			i++
		}

		if i >= len(tokens) || tokens[i].Text == ")" || tokens[i].Text == ";" {
			break
		}

		param := Parameter{}
		paramTypeParts := []string{}

		// Collect parameter type (everything up to parameter name or comma/paren)
		for i < len(tokens) {
			if tokens[i].Text == ")" || tokens[i].Text == "," || tokens[i].Text == ";" {
				break
			}

			// Parameter name is typically an internalParam or last identifier
			if tokens[i].Kind == "internalParam" {
				param.Name = tokens[i].Text
				i++
				break
			}

			// Check if this is parameter name (identifier at end of type)
			if tokens[i].Kind == "identifier" {
				// Look ahead - if next is comma or paren, this is the parameter name
				j := i + 1
				for j < len(tokens) && tokens[j].Kind == "text" && strings.TrimSpace(tokens[j].Text) == "" {
					j++
				}
				if j < len(tokens) && (tokens[j].Text == "," || tokens[j].Text == ")" || tokens[j].Text == ";") {
					param.Name = tokens[i].Text
					i = j
					break
				}
			}

			// Part of type - exclude punctuation marks
			text := tokens[i].Text
			if text != "" && strings.TrimSpace(text) != "" && text != ";" && text != ")" && text != "(" {
				paramTypeParts = append(paramTypeParts, text)
			}
			i++
		}

		param.Type = strings.Join(paramTypeParts, " ")
		if param.Type != "" {
			fn.Parameters = append(fn.Parameters, param)
		}

		// Skip comma
		if i < len(tokens) && tokens[i].Text == "," {
			i++
		}
	}

	// Skip functions with colons (ObjC selectors)
	if strings.Contains(fn.Name, ":") {
		return nil, fmt.Errorf("skipping ObjC selector: %s", fn.Name)
	}

	if fn.Name == "" {
		return nil, fmt.Errorf("failed to parse function name")
	}

	// Clean up return type - remove trailing semicolons and parentheses
	fn.ReturnType = strings.TrimRight(fn.ReturnType, ";)")
	fn.ReturnType = strings.TrimSpace(fn.ReturnType)

	return fn, nil
}

// parseClassDeclaration parses an Objective-C class declaration from tokens.
// Example: @interface FSUnaryFileSystem : NSObject
func parseClassDeclaration(tokens []Token) *ParsedClass {
	cls := &ParsedClass{}

	// Look for @interface followed by class name and optional superclass
	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]

		// Found @interface keyword
		if tok.Kind == "keyword" && tok.Text == "@interface" {
			// Next non-whitespace token should be the class name
			i++
			for i < len(tokens) && tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) == "" {
				i++
			}

			if i < len(tokens) && tokens[i].Kind == "identifier" {
				cls.Name = tokens[i].Text
				i++

				// Look for superclass - skip whitespace and check for ":"
				for i < len(tokens) {
					if tokens[i].Kind == "text" {
						if strings.Contains(tokens[i].Text, ":") {
							i++
							break
						}
					}
					i++
				}

				// Find the superclass name (next identifier or typeIdentifier)
				for i < len(tokens) {
					if tokens[i].Kind == "identifier" || tokens[i].Kind == "typeIdentifier" {
						cls.SuperClass = tokens[i].Text
						break
					}
					if tokens[i].Kind == "text" && strings.TrimSpace(tokens[i].Text) != "" {
						break
					}
					i++
				}
				break
			}
		}
	}

	if cls.Name == "" {
		return nil
	}

	return cls
}

// parseProtocolDeclaration parses an Objective-C protocol declaration from tokens.
// Example: @protocol FSModuleExtension
func parseProtocolDeclaration(tokens []Token) *ParsedProtocol {
	proto := &ParsedProtocol{}

	// Look for @protocol followed by protocol name
	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]

		// Found @protocol keyword
		if tok.Kind == "keyword" && tok.Text == "@protocol" {
			// Next token should be the protocol name
			if i+1 < len(tokens) && tokens[i+1].Kind == "identifier" {
				proto.Name = tokens[i+1].Text
				break
			}
		}
	}

	if proto.Name == "" {
		return nil
	}

	return proto
}

func generatePuregoBindings(functions []*ParsedFunction, classes []*ParsedClass, protocols []*ParsedProtocol, outputDir, framework string) {
	pkgName := strings.ToLower(framework)

	// Generate doc.go with package documentation
	generateDocFile(outputDir, pkgName, framework, functions)

	// Generate types.gen.go
	generateTypesFile(outputDir, pkgName, framework)

	// Generate loader.gen.go
	generateLoaderFile(outputDir, pkgName, framework)

	// Generate functions.gen.go
	generateFunctionsFile(outputDir, pkgName, framework, functions)

	// Generate classes.gen.go if there are classes
	if len(classes) > 0 {
		generateClassesFile(outputDir, pkgName, framework, classes)
	}

	// Generate protocols.gen.go if there are protocols
	if len(protocols) > 0 {
		generateProtocolsFile(outputDir, pkgName, framework, protocols)
	}

	fileCount := 4
	if len(classes) > 0 {
		fileCount++
	}
	if len(protocols) > 0 {
		fileCount++
	}

	log.Printf("Generated %d .gen.go files in %s", fileCount, outputDir)
}

// generateDocFile generates package documentation with version information
func generateDocFile(outputDir, pkgName, framework string, functions []*ParsedFunction) {
	filename := filepath.Join(outputDir, "doc.go")
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create doc file: %v", err)
	}
	defer f.Close()

	// Find minimum version across all functions
	minVersion := findMinimumMacOSVersion(functions)

	fmt.Fprintf(f, "// Code generated from Apple documentation for %s. DO NOT EDIT.\n\n", framework)
	fmt.Fprintf(f, "// Package %s provides Go bindings for the %s framework.\n", pkgName, framework)
	fmt.Fprintf(f, "//\n")

	if minVersion != "" {
		fmt.Fprintf(f, "// Minimum macOS version: %s\n", minVersion)
	}

	fmt.Fprintf(f, "// Framework path: /System/Library/Frameworks/%s.framework/%s\n", framework, framework)
	fmt.Fprintf(f, "//\n")
	fmt.Fprintf(f, "// These bindings are generated from Apple's official documentation and\n")
	fmt.Fprintf(f, "// provide purego-based access to %s without requiring cgo.\n", framework)
	fmt.Fprintf(f, "package %s\n\n", pkgName)

	// Add package-level constants
	if minVersion != "" {
		fmt.Fprintf(f, "// MinMacOSVersion is the minimum macOS version required for this framework.\n")
		fmt.Fprintf(f, "const MinMacOSVersion = \"%s\"\n\n", minVersion)
	}

	fmt.Fprintf(f, "// FrameworkPath is the system path to the framework binary.\n")
	fmt.Fprintf(f, "const FrameworkPath = \"/System/Library/Frameworks/%s.framework/%s\"\n", framework, framework)
}

// findMinimumMacOSVersion finds the minimum macOS version across all functions.
// Returns empty string if no macOS versions are found.
func findMinimumMacOSVersion(functions []*ParsedFunction) string {
	var minVersion string

	for _, fn := range functions {
		if ver, ok := fn.Availability.IntroducedAt["macOS"]; ok && ver != "" {
			if minVersion == "" || compareVersionStrings(ver, minVersion) < 0 {
				minVersion = ver
			}
		}
	}

	return minVersion
}

// parseVersion parses a version string like "10.14" into major and minor components.
// Returns major, minor, and ok=true if parsing succeeded.
func parseVersion(s string) (major, minor int, ok bool) {
	parts := strings.SplitN(s, ".", 2)
	if len(parts) == 0 {
		return 0, 0, false
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, false
	}

	minor = 0
	if len(parts) > 1 {
		minor, err = strconv.Atoi(parts[1])
		if err != nil {
			return 0, 0, false
		}
	}

	return major, minor, true
}

// compareVersionStrings compares two version strings semantically (e.g., "10.14" vs "10.9").
// Returns: -1 if a < b, 0 if a == b, 1 if a > b
func compareVersionStrings(a, b string) int {
	if a == b {
		return 0
	}

	maj1, min1, ok1 := parseVersion(a)
	maj2, min2, ok2 := parseVersion(b)

	// Fallback to lexicographic comparison if parsing fails
	if !ok1 || !ok2 {
		return strings.Compare(a, b)
	}

	// Compare major versions first
	if maj1 < maj2 {
		return -1
	}
	if maj1 > maj2 {
		return 1
	}

	// Major versions equal, compare minor versions
	if min1 < min2 {
		return -1
	}
	if min1 > min2 {
		return 1
	}

	return 0
}

func generateTypesFile(outputDir, pkgName, framework string) {
	filename := filepath.Join(outputDir, "types.gen.go")
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create types file: %v", err)
	}
	defer f.Close()

	fmt.Fprintf(f, "// Code generated from Apple documentation for %s. DO NOT EDIT.\n\n", framework)
	fmt.Fprintf(f, "package %s\n\n", pkgName)
	fmt.Fprintf(f, "import \"unsafe\"\n\n")

	fmt.Fprintf(f, "// %s Types\n\n", framework)

	// Common CoreGraphics types
	if framework == "CoreGraphics" {
		fmt.Fprintf(f, "// Fundamental types\n")
		fmt.Fprintf(f, "type CGFloat float64\n\n")

		fmt.Fprintf(f, "// Opaque reference types\n")
		fmt.Fprintf(f, "type CGContextRef unsafe.Pointer\n")
		fmt.Fprintf(f, "type CGColorRef unsafe.Pointer\n")
		fmt.Fprintf(f, "type CGColorSpaceRef unsafe.Pointer\n")
		fmt.Fprintf(f, "type CGPathRef unsafe.Pointer\n")
		fmt.Fprintf(f, "type CGImageRef unsafe.Pointer\n")
		fmt.Fprintf(f, "type CGDataProviderRef unsafe.Pointer\n")
		fmt.Fprintf(f, "type CGFontRef unsafe.Pointer\n")
		fmt.Fprintf(f, "type CGGradientRef unsafe.Pointer\n")
		fmt.Fprintf(f, "type CGLayerRef unsafe.Pointer\n")
		fmt.Fprintf(f, "type CGPDFDocumentRef unsafe.Pointer\n")
		fmt.Fprintf(f, "type CGPDFPageRef unsafe.Pointer\n\n")

		fmt.Fprintf(f, "// Geometric types\n")
		fmt.Fprintf(f, "type CGPoint struct {\n")
		fmt.Fprintf(f, "\tX, Y CGFloat\n")
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "type CGSize struct {\n")
		fmt.Fprintf(f, "\tWidth, Height CGFloat\n")
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "type CGRect struct {\n")
		fmt.Fprintf(f, "\tOrigin CGPoint\n")
		fmt.Fprintf(f, "\tSize   CGSize\n")
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "type CGAffineTransform struct {\n")
		fmt.Fprintf(f, "\tA, B, C, D, Tx, Ty CGFloat\n")
		fmt.Fprintf(f, "}\n\n")
	}
}

func generateLoaderFile(outputDir, pkgName, framework string) {
	filename := filepath.Join(outputDir, "loader.gen.go")
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create loader file: %v", err)
	}
	defer f.Close()

	fmt.Fprintf(f, "// Code generated from Apple documentation for %s. DO NOT EDIT.\n\n", framework)
	fmt.Fprintf(f, "package %s\n\n", pkgName)
	fmt.Fprintf(f, "import \"github.com/ebitengine/purego\"\n\n")

	fmt.Fprintf(f, "// lib holds the framework library handle.\n")
	fmt.Fprintf(f, "// Functions are automatically registered in init().\n")
	fmt.Fprintf(f, "var lib uintptr\n\n")

	fmt.Fprintf(f, "func init() {\n")
	fmt.Fprintf(f, "\tvar err error\n")
	fmt.Fprintf(f, "\tlib, err = purego.Dlopen(FrameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)\n")
	fmt.Fprintf(f, "\tif err != nil {\n")
	fmt.Fprintf(f, "\t\tpanic(err)\n")
	fmt.Fprintf(f, "\t}\n")
	fmt.Fprintf(f, "\tregisterFunctions()\n")
	fmt.Fprintf(f, "}\n")
}

func generateFunctionsFile(outputDir, pkgName, framework string, functions []*ParsedFunction) {
	filename := filepath.Join(outputDir, "functions.gen.go")
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create functions file: %v", err)
	}
	defer f.Close()

	fmt.Fprintf(f, "// Code generated from Apple documentation for %s. DO NOT EDIT.\n\n", framework)
	fmt.Fprintf(f, "package %s\n\n", pkgName)
	fmt.Fprintf(f, "import \"github.com/ebitengine/purego\"\n\n")

	fmt.Fprintf(f, "// %s Functions (%d total)\n", framework, len(functions))
	fmt.Fprintf(f, "//\n")
	fmt.Fprintf(f, "// This file contains executable function bindings automatically registered via purego.\n")
	fmt.Fprintf(f, "// All functions are ready to use after package initialization.\n\n")

	// Generate function variables
	for _, fn := range functions {
		if fn.Name == "" {
			continue
		}
		generateExecutableFunctionDeclaration(f, fn, framework)
	}

	// Generate registerFunctions
	fmt.Fprintf(f, "\n// registerFunctions registers all framework functions with purego\n")
	fmt.Fprintf(f, "func registerFunctions() {\n")
	for _, fn := range functions {
		if fn.Name == "" {
			continue
		}
		fmt.Fprintf(f, "\tpurego.RegisterLibFunc(&%s, lib, \"%s\")\n", fn.Name, fn.Name)
	}
	fmt.Fprintf(f, "}\n")
}

// generateExecutableFunctionDeclaration generates an executable function variable declaration
func generateExecutableFunctionDeclaration(f *os.File, fn *ParsedFunction, framework string) {
	// Write documentation comment
	if !fn.Availability.IsEmpty() {
		for _, platform := range fn.Availability.Platforms() {
			version := fn.Availability.IntroducedAt[platform]
			status := ""
			if fn.Availability.Beta {
				status = " (Beta)"
			} else if deprecatedAt, ok := fn.Availability.DeprecatedAt[platform]; ok {
				status = fmt.Sprintf(" (Deprecated in %s)", deprecatedAt)
			}
			if platform == "macOS" { // Only show macOS for now
				fmt.Fprintf(f, "// %s is available on %s %s+%s\n", fn.Name, platform, version, status)
				break
			}
		}
	}

	// Write function variable declaration
	fmt.Fprintf(f, "var %s func(", fn.Name)

	// Write parameters
	for i, p := range fn.Parameters {
		if i > 0 {
			fmt.Fprintf(f, ", ")
		}
		// Clean parameter type
		paramType := strings.TrimRight(p.Type, ",;)")
		paramType = strings.TrimSpace(paramType)
		paramType = mapCTypeToGo(paramType, framework)

		if p.Name != "" {
			fmt.Fprintf(f, "%s %s", p.Name, paramType)
		} else {
			fmt.Fprintf(f, "%s", paramType)
		}
	}

	fmt.Fprintf(f, ")")

	// Write return type
	if fn.ReturnType != "" && fn.ReturnType != "void" {
		returnType := mapCTypeToGo(fn.ReturnType, framework)
		fmt.Fprintf(f, " %s", returnType)
	}

	fmt.Fprintf(f, "\n\n")
}

// mapCTypeToGo maps C types to Go types for a given framework
func mapCTypeToGo(cType, framework string) string {
	cType = strings.TrimSpace(cType)

	// Framework-specific types
	if framework == "CoreGraphics" {
		switch {
		case strings.HasPrefix(cType, "CG") && strings.HasSuffix(cType, "Ref"):
			return cType // Already a Go type
		case cType == "CGFloat":
			return "CGFloat"
		case cType == "CGPoint":
			return "CGPoint"
		case cType == "CGSize":
			return "CGSize"
		case cType == "CGRect":
			return "CGRect"
		case cType == "CGAffineTransform":
			return "CGAffineTransform"
		}
	}

	// Common C types
	switch {
	case cType == "void":
		return ""
	case cType == "int":
		return "int"
	case cType == "size_t":
		return "uintptr"
	case cType == "uint32_t":
		return "uint32"
	case cType == "uint64_t":
		return "uint64"
	case cType == "float":
		return "float32"
	case cType == "double":
		return "float64"
	case cType == "bool", cType == "BOOL":
		return "bool"
	case strings.Contains(cType, "*"):
		return "unsafe.Pointer"
	default:
		// Default to unsafe.Pointer for unknown types
		return "unsafe.Pointer"
	}
}

// generateFunctionComment generates a comment block for a function including version info
func generateFunctionComment(f *os.File, fn *ParsedFunction) {
	// Function signature
	fmt.Fprintf(f, "// %s", fn.Name)
	if len(fn.Parameters) > 0 {
		fmt.Fprintf(f, "(")
		for j, p := range fn.Parameters {
			if j > 0 {
				fmt.Fprintf(f, ", ")
			}
			// Clean parameter type - remove trailing punctuation
			paramType := strings.TrimRight(p.Type, ",;)")
			paramType = strings.TrimSpace(paramType)

			if p.Name != "" {
				fmt.Fprintf(f, "%s ", p.Name)
			}
			fmt.Fprintf(f, "%s", paramType)
		}
		fmt.Fprintf(f, ")")
	} else {
		fmt.Fprintf(f, "()")
	}
	if fn.ReturnType != "" && fn.ReturnType != "void" {
		fmt.Fprintf(f, " %s", fn.ReturnType)
	}
	fmt.Fprintf(f, "\n")

	// Add availability information if present
	if !fn.Availability.IsEmpty() {
		fmt.Fprintf(f, "//\n")
		fmt.Fprintf(f, "// Availability:\n")

		// Iterate over platforms in sorted order for consistent output
		for _, platform := range fn.Availability.Platforms() {
			version := fn.Availability.IntroducedAt[platform]
			status := ""

			if fn.Availability.Beta {
				status = " (Beta)"
			} else if deprecatedAt, ok := fn.Availability.DeprecatedAt[platform]; ok {
				status = fmt.Sprintf(" (Deprecated in %s)", deprecatedAt)
			}

			fmt.Fprintf(f, "//   - %s %s+%s\n", platform, version, status)
		}
	}

	// Add deprecation warning if any platform is deprecated
	if len(fn.Availability.DeprecatedAt) > 0 {
		fmt.Fprintf(f, "//\n")
		fmt.Fprintf(f, "// Deprecated: This function is deprecated.\n")
	}

	fmt.Fprintf(f, "\n")
}

// generateClassesFile generates a file with Objective-C class declarations
func generateClassesFile(outputDir, pkgName, framework string, classes []*ParsedClass) {
	filename := filepath.Join(outputDir, "classes.gen.go")
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create classes file: %v", err)
	}
	defer f.Close()

	fmt.Fprintf(f, "// Code generated from Apple documentation for %s. DO NOT EDIT.\n\n", framework)
	fmt.Fprintf(f, "package %s\n\n", pkgName)

	fmt.Fprintf(f, "// %s Classes\n", framework)
	fmt.Fprintf(f, "//\n")
	fmt.Fprintf(f, "// This file contains class declarations discovered from Apple's documentation.\n")
	fmt.Fprintf(f, "// These represent Objective-C classes that can be used with objc runtime bindings.\n")
	fmt.Fprintf(f, "\n")

	fmt.Fprintf(f, "// Discovered classes (%d total):\n\n", len(classes))

	for _, cls := range classes {
		if cls.Name == "" {
			continue
		}

		// Generate class comment
		fmt.Fprintf(f, "// %s", cls.Name)
		if cls.SuperClass != "" {
			fmt.Fprintf(f, " : %s", cls.SuperClass)
		}
		fmt.Fprintf(f, "\n")

		// Add availability information if present
		if !cls.Availability.IsEmpty() {
			fmt.Fprintf(f, "//\n")
			fmt.Fprintf(f, "// Availability:\n")

			for _, platform := range cls.Availability.Platforms() {
				version := cls.Availability.IntroducedAt[platform]
				status := ""

				if cls.Availability.Beta {
					status = " (Beta)"
				} else if deprecatedAt, ok := cls.Availability.DeprecatedAt[platform]; ok {
					status = fmt.Sprintf(" (Deprecated in %s)", deprecatedAt)
				}

				fmt.Fprintf(f, "//   - %s %s+%s\n", platform, version, status)
			}
		}

		fmt.Fprintf(f, "\n")
	}
}

// generateProtocolsFile generates a file with Objective-C protocol declarations
func generateProtocolsFile(outputDir, pkgName, framework string, protocols []*ParsedProtocol) {
	filename := filepath.Join(outputDir, "protocols.gen.go")
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create protocols file: %v", err)
	}
	defer f.Close()

	fmt.Fprintf(f, "// Code generated from Apple documentation for %s. DO NOT EDIT.\n\n", framework)
	fmt.Fprintf(f, "package %s\n\n", pkgName)

	fmt.Fprintf(f, "// %s Protocols\n", framework)
	fmt.Fprintf(f, "//\n")
	fmt.Fprintf(f, "// This file contains protocol declarations discovered from Apple's documentation.\n")
	fmt.Fprintf(f, "// These represent Objective-C protocols that can be used with objc runtime bindings.\n")
	fmt.Fprintf(f, "\n")

	fmt.Fprintf(f, "// Discovered protocols (%d total):\n\n", len(protocols))

	for _, proto := range protocols {
		if proto.Name == "" {
			continue
		}

		// Generate protocol comment
		fmt.Fprintf(f, "// @protocol %s\n", proto.Name)

		// Add availability information if present
		if !proto.Availability.IsEmpty() {
			fmt.Fprintf(f, "//\n")
			fmt.Fprintf(f, "// Availability:\n")

			for _, platform := range proto.Availability.Platforms() {
				version := proto.Availability.IntroducedAt[platform]
				status := ""

				if proto.Availability.Beta {
					status = " (Beta)"
				} else if deprecatedAt, ok := proto.Availability.DeprecatedAt[platform]; ok {
					status = fmt.Sprintf(" (Deprecated in %s)", deprecatedAt)
				}

				fmt.Fprintf(f, "//   - %s %s+%s\n", platform, version, status)
			}
		}

		fmt.Fprintf(f, "\n")
	}
}

func generateDarkwinKitBindings(functions []*ParsedFunction, classes []*ParsedClass, protocols []*ParsedProtocol, outputDir, framework string) {
	log.Printf("Darwinkit style not fully implemented yet for %s", framework)
	generatePuregoBindings(functions, classes, protocols, outputDir, framework)
}

func generateSimpleBindings(functions []*ParsedFunction, classes []*ParsedClass, protocols []*ParsedProtocol, outputDir, framework string) {
	log.Printf("Simple style not fully implemented yet for %s", framework)
	generatePuregoBindings(functions, classes, protocols, outputDir, framework)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
