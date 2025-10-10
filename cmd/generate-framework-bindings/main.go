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
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/tmc/appledocs"
	"github.com/tmc/appledocs/occ2go"
	"golang.org/x/tools/txtar"
)

//go:embed funcs.go
var _ string // Force funcs.go to be included in binary for template compilation

//go:embed templates.txtar
var templatesData []byte

var (
	docTemplate               *template.Template
	coreGraphicsTypesTemplate *template.Template
	classesTemplate           *template.Template
	protocolsTemplate         *template.Template
	functionsGenTemplate      *template.Template
)

func init() {
	// Parse txtar archive
	archive := txtar.Parse(templatesData)
	templates := make(map[string]string)
	for _, file := range archive.Files {
		templates[file.Name] = string(file.Data)
	}

	var err error
	docTemplate, err = template.New("doc.gen.go").Funcs(templateFuncs).Parse(templates["doc.gen.go"])
	if err != nil {
		panic(fmt.Errorf("failed to parse doc.gen.go: %w", err))
	}

	coreGraphicsTypesTemplate, err = template.New("types.gen.go").Funcs(templateFuncs).Parse(templates["types.gen.go"])
	if err != nil {
		panic(fmt.Errorf("failed to parse types.gen.go: %w", err))
	}

	classesTemplate, err = template.New("classes.gen.go").Funcs(templateFuncs).Parse(templates["classes.gen.go"])
	if err != nil {
		panic(fmt.Errorf("failed to parse classes.gen.go: %w", err))
	}

	protocolsTemplate, err = template.New("protocols.gen.go").Funcs(templateFuncs).Parse(templates["protocols.gen.go"])
	if err != nil {
		panic(fmt.Errorf("failed to parse protocols.gen.go: %w", err))
	}

	functionsGenTemplate, err = template.New("functions.gen.go").Funcs(templateFuncs).Parse(templates["functions.gen.go"])
	if err != nil {
		panic(fmt.Errorf("failed to parse functions.gen.go: %w", err))
	}
}

func main() {
	framework := flag.String("framework", "CoreGraphics", "Framework to generate bindings for")
	inputDir := flag.String("input", "", "Input directory with JSON files (defaults to ~/.appledocs/cache/developer.apple.com/tutorials/data/documentation)")
	outputDir := flag.String("output", "generated", "Output directory for generated bindings")
	style := flag.String("style", "purego", "Binding style: purego, darwinkit, or simple")
	filterRegexp := flag.String("filter", "", "Only generate symbols matching this regexp (e.g., '^CGRect' or '^NS(Window|View)')")
	txtarOutput := flag.Bool("txtar", false, "Output as txtar format to stdout instead of files")
	flag.Parse()

	// Default to cache directory if not specified
	if *inputDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Failed to get home directory: %v", err)
		}
		*inputDir = filepath.Join(homeDir, ".appledocs/cache/developer.apple.com/tutorials/data/documentation")
	}

	log.Printf("Generating %s bindings for %s framework", *style, *framework)
	log.Printf("Input directory: %s", *inputDir)
	log.Printf("Output directory: %s", *outputDir)

	// Open the appledocs filesystem
	fsys, err := appledocs.Open(*inputDir)
	if err != nil {
		log.Fatalf("Failed to open appledocs filesystem: %v", err)
	}

	// Parse all symbols in the framework using the appledocs iterator
	var functions []*occ2go.ParsedFunction
	var classes []*occ2go.ParsedClass
	var protocols []*occ2go.ParsedProtocol

	processedFiles := 0
	for path, doc := range appledocs.Symbols(fsys, *framework) {
		processedFiles++
		fn, cls, proto, err := occ2go.ParseDocument(doc)
		if err == nil {
			if fn != nil {
				functions = append(functions, fn)
			}
			if cls != nil {
				classes = append(classes, cls)
			}
			if proto != nil {
				protocols = append(protocols, proto)
			}
		} else {
			log.Printf("Warning: failed to parse %s: %v", path, err)
		}
	}

	log.Printf("Processed %d symbols in %s", processedFiles, *framework)
	log.Printf("Found %d functions, %d classes, %d protocols", len(functions), len(classes), len(protocols))

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
		log.Printf("Deduplicated %d functions down to %d unique functions", len(functions), len(uniqueFunctions))
		functions = uniqueFunctions
	}

	// Apply regexp filter if specified
	if *filterRegexp != "" {
		re, err := regexp.Compile(*filterRegexp)
		if err != nil {
			log.Fatalf("Invalid filter regexp: %v", err)
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

		log.Printf("Filter '%s' matched %d/%d functions, %d/%d classes, %d/%d protocols",
			*filterRegexp,
			len(filteredFunctions), len(functions),
			len(filteredClasses), len(classes),
			len(filteredProtocols), len(protocols))

		functions = filteredFunctions
		classes = filteredClasses
		protocols = filteredProtocols
	}

	// Create output directory
	packageName := strings.ToLower(*framework)
	outDir := filepath.Join(*outputDir, packageName)
	if !*txtarOutput {
		if err := os.MkdirAll(outDir, 0755); err != nil {
			log.Fatalf("Failed to create output directory: %v", err)
		}
	}

	// Generate bindings
	if *txtarOutput {
		if err := generateTxtar(os.Stdout, *framework, packageName, *inputDir, functions, classes, protocols); err != nil {
			log.Fatalf("Failed to generate bindings: %v", err)
		}
	} else {
		if err := generateFiles(outDir, *framework, packageName, *inputDir, functions, classes, protocols); err != nil {
			log.Fatalf("Failed to generate bindings: %v", err)
		}
		log.Printf("Generated bindings in %s", outDir)
	}
}

// generateFiles generates all files to disk
func generateFiles(outDir, framework, packageName, inputDir string, functions []*occ2go.ParsedFunction, classes []*occ2go.ParsedClass, protocols []*occ2go.ParsedProtocol) error {
	generators := []struct {
		filename string
		generate func(io.Writer) error
	}{
		{"doc.gen.go", func(w io.Writer) error { return generateDoc(w, framework, packageName, inputDir, functions) }},
		{"types.gen.go", func(w io.Writer) error { return generateTypes(w, framework, packageName, functions) }},
		{"functions.gen.go", func(w io.Writer) error { return generateFunctions(w, framework, packageName, functions) }},
	}

	if len(classes) > 0 {
		generators = append(generators, struct {
			filename string
			generate func(io.Writer) error
		}{"classes.gen.go", func(w io.Writer) error { return generateClasses(w, framework, packageName, classes) }})
	}

	if len(protocols) > 0 {
		generators = append(generators, struct {
			filename string
			generate func(io.Writer) error
		}{"protocols.gen.go", func(w io.Writer) error { return generateProtocols(w, framework, packageName, protocols) }})
	}

	for _, gen := range generators {
		f, err := os.Create(filepath.Join(outDir, gen.filename))
		if err != nil {
			return err
		}
		err = gen.generate(f)
		f.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// generateTxtar generates all files as txtar format
func generateTxtar(w io.Writer, framework, packageName, inputDir string, functions []*occ2go.ParsedFunction, classes []*occ2go.ParsedClass, protocols []*occ2go.ParsedProtocol) error {
	files := make(map[string][]byte)

	genFile := func(filename string, generator func(io.Writer) error) error {
		var buf bytes.Buffer
		if err := generator(&buf); err != nil {
			return err
		}
		files[filename] = buf.Bytes()
		return nil
	}

	if err := genFile("doc.gen.go", func(w io.Writer) error { return generateDoc(w, framework, packageName, inputDir, functions) }); err != nil {
		return err
	}
	if err := genFile("types.gen.go", func(w io.Writer) error { return generateTypes(w, framework, packageName, functions) }); err != nil {
		return err
	}
	if err := genFile("functions.gen.go", func(w io.Writer) error { return generateFunctions(w, framework, packageName, functions) }); err != nil {
		return err
	}
	if len(classes) > 0 {
		if err := genFile("classes.gen.go", func(w io.Writer) error { return generateClasses(w, framework, packageName, classes) }); err != nil {
			return err
		}
	}
	if len(protocols) > 0 {
		if err := genFile("protocols.gen.go", func(w io.Writer) error { return generateProtocols(w, framework, packageName, protocols) }); err != nil {
			return err
		}
	}

	// Write txtar format
	fmt.Fprintf(w, "# Generated bindings for %s framework\n", framework)
	fmt.Fprintf(w, "# Package: %s\n#\n", packageName)
	fmt.Fprintf(w, "# Functions: %d\n", len(functions))
	fmt.Fprintf(w, "# Classes: %d\n", len(classes))
	fmt.Fprintf(w, "# Protocols: %d\n\n", len(protocols))

	fileNames := make([]string, 0, len(files))
	for name := range files {
		fileNames = append(fileNames, name)
	}
	sort.Strings(fileNames)

	for _, name := range fileNames {
		content := files[name]
		fmt.Fprintf(w, "-- %s --\n", name)
		w.Write(content)
		if len(content) > 0 && content[len(content)-1] != '\n' {
			fmt.Fprintf(w, "\n")
		}
	}

	return nil
}

// generateDoc generates package documentation
func generateDoc(w io.Writer, framework, packageName, inputDir string, functions []*occ2go.ParsedFunction) error {
	frameworkAbstract, frameworkURL := loadFrameworkMetadata(inputDir, framework)

	data := struct {
		Framework   string
		PackageName string
		MinVersion  string
		Abstract    string
		DocURL      string
	}{
		Framework:   framework,
		PackageName: packageName,
		MinVersion:  findMinimumMacOSVersion(functions),
		Abstract:    frameworkAbstract,
		DocURL:      frameworkURL,
	}

	return docTemplate.Execute(w, data)
}

// generateTypes generates framework-specific type definitions
func generateTypes(w io.Writer, framework, packageName string, functions []*occ2go.ParsedFunction) error {
	if framework == "CoreGraphics" {
		refTypes := extractRefTypes(functions, "CG")
		data := struct {
			Framework   string
			PackageName string
			RefTypes    []string
		}{framework, packageName, refTypes}
		return coreGraphicsTypesTemplate.Execute(w, data)
	}

	// Empty types file for other frameworks
	fmt.Fprintf(w, "// Code generated from Apple documentation for %s. DO NOT EDIT.\n\n", framework)
	fmt.Fprintf(w, "package %s\n", packageName)
	return nil
}

// generateFunctions generates function bindings
func generateFunctions(w io.Writer, framework, packageName string, functions []*occ2go.ParsedFunction) error {
	data := struct {
		Framework   string
		PackageName string
		Count       int
		Functions   []*occ2go.ParsedFunction
	}{framework, packageName, len(functions), functions}

	return functionsGenTemplate.Execute(w, data)
}

// generateClasses generates class declarations
func generateClasses(w io.Writer, framework, packageName string, classes []*occ2go.ParsedClass) error {
	data := struct {
		Framework   string
		PackageName string
		Count       int
		Classes     []*occ2go.ParsedClass
	}{framework, packageName, len(classes), classes}

	return classesTemplate.Execute(w, data)
}

// generateProtocols generates protocol declarations
func generateProtocols(w io.Writer, framework, packageName string, protocols []*occ2go.ParsedProtocol) error {
	data := struct {
		Framework   string
		PackageName string
		Count       int
		Protocols   []*occ2go.ParsedProtocol
	}{framework, packageName, len(protocols), protocols}

	return protocolsTemplate.Execute(w, data)
}

// extractRefTypes extracts all Ref types from function signatures
func extractRefTypes(functions []*occ2go.ParsedFunction, prefix string) []string {
	refTypesMap := make(map[string]bool)

	for _, fn := range functions {
		if strings.HasPrefix(fn.ReturnType, prefix) && strings.HasSuffix(fn.ReturnType, "Ref") {
			refTypesMap[fn.ReturnType] = true
		}
		for _, param := range fn.Parameters {
			typ := strings.TrimSpace(param.Type)
			if strings.HasPrefix(typ, prefix) && strings.HasSuffix(typ, "Ref") {
				refTypesMap[typ] = true
			}
		}
	}

	refTypes := make([]string, 0, len(refTypesMap))
	for typ := range refTypesMap {
		refTypes = append(refTypes, typ)
	}
	sort.Strings(refTypes)
	return refTypes
}

// loadFrameworkMetadata loads the framework-level JSON to extract abstract and URL
func loadFrameworkMetadata(inputDir, framework string) (abstract string, docURL string) {
	frameworkPath := filepath.Join(inputDir, framework+".json")
	data, err := os.ReadFile(frameworkPath)
	if err != nil {
		// Try lowercase
		frameworkPath = filepath.Join(inputDir, strings.ToLower(framework)+".json")
		data, err = os.ReadFile(frameworkPath)
		if err != nil {
			return "", ""
		}
	}

	var doc struct {
		Abstract []struct {
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"abstract"`
		Identifier struct {
			URL string `json:"url"`
		} `json:"identifier"`
	}

	if err := json.Unmarshal(data, &doc); err != nil {
		return "", ""
	}

	// Extract abstract text
	for _, item := range doc.Abstract {
		if item.Type == "text" && item.Text != "" {
			abstract = item.Text
			break
		}
	}

	docURL = occ2go.ConvertDocURLToWeb(doc.Identifier.URL)
	return abstract, docURL
}

// findMinimumMacOSVersion finds the minimum macOS version across all functions
func findMinimumMacOSVersion(functions []*occ2go.ParsedFunction) string {
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

// compareVersionStrings compares two version strings semantically
func compareVersionStrings(a, b string) int {
	if a == b {
		return 0
	}

	maj1, min1, ok1 := parseVersion(a)
	maj2, min2, ok2 := parseVersion(b)

	if !ok1 || !ok2 {
		return strings.Compare(a, b)
	}

	if maj1 < maj2 {
		return -1
	}
	if maj1 > maj2 {
		return 1
	}

	if min1 < min2 {
		return -1
	}
	if min1 > min2 {
		return 1
	}

	return 0
}

// parseVersion parses a version string like "10.14" into major and minor components
func parseVersion(s string) (major, minor int, ok bool) {
	parts := strings.SplitN(s, ".", 2)
	if len(parts) == 0 {
		return 0, 0, false
	}

	if _, err := fmt.Sscanf(parts[0], "%d", &major); err != nil {
		return 0, 0, false
	}

	minor = 0
	if len(parts) > 1 {
		fmt.Sscanf(parts[1], "%d", &minor)
	}

	return major, minor, true
}
