package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/tmc/appledocs/occ2go"
	"golang.org/x/tools/txtar"
)

func generateFiles(outDir, framework, packageName, inputDir string, functions []*occ2go.ParsedFunction, classes []*occ2go.ParsedClass, protocols []*occ2go.ParsedProtocol, enums []*occ2go.ParsedEnum, typedefs []*occ2go.ParsedTypedef, constants []*occ2go.ParsedConstant, withRefMethods, generateTests, generateExamples bool, variant string) error {
	// Determine output module - default to github.com/tmc/appledocs/generated for now
	outputModule := "github.com/tmc/appledocs/generated"

	// Create generator to get gen.go data
	gen := NewGenerator(framework, packageName, inputDir, outputModule, variant, withRefMethods, generateTests, generateExamples)
	gen.Functions = functions
	gen.Classes = classes
	gen.Protocols = protocols
	gen.Enums = enums
	gen.Typedefs = typedefs
	gen.Constants = constants

	// DEBUG: Check if enums have cases after assignment
	if verbose {
		// for _, enum := range gen.Enums {
		// 	if len(enum.Cases) > 0 {
		// 		fmt.Fprintf(os.Stderr, "DEBUG generateFiles: enum %s has %d cases\n", enum.Name, len(enum.Cases))
		// 	}
		// }
	}

	// Apply property overrides for undocumented properties
	for _, cls := range gen.Classes {
		MergePropertyOverrides(framework, cls.Name, cls)
	}

	gen.prepare()

	// Generate stubs for missing parent classes
	stubs := gen.GenerateMissingParentStubs()
	if len(stubs) > 0 {
		gen.Classes = append(stubs, gen.Classes...)
	}

	// Generate using module template and parse txtar output
	var buf bytes.Buffer
	if err := gen.GenerateTxtarFromModule(&buf); err != nil {
		return err
	}

	// Parse txtar output
	archive := txtar.Parse(buf.Bytes())

	// Write each file from the archive
	for _, file := range archive.Files {
		if file.Name == "" {
			continue
		}
		filePath := filepath.Join(outDir, file.Name)
		// Create parent directory if needed
		if dir := filepath.Dir(filePath); dir != "." {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", dir, err)
			}
		}
		if err := os.WriteFile(filePath, file.Data, 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", file.Name, err)
		}
	}

	// Log any collected errors/warnings
	if verbose {
		for _, e := range gen.Errors {
			fmt.Fprintf(os.Stderr, "Warning: %v\n", e)
		}
	}

	return nil
}

// GenerateTxtarFromModule generates the entire txtar output using the module template
// GenerateTxtarFromModule is defined in generator.go

// generateTxtar generates all files as txtar format
func generateTxtar(w io.Writer, framework, packageName, inputDir string, functions []*occ2go.ParsedFunction, classes []*occ2go.ParsedClass, protocols []*occ2go.ParsedProtocol, enums []*occ2go.ParsedEnum, typedefs []*occ2go.ParsedTypedef, constants []*occ2go.ParsedConstant, withRefMethods, generateTests, generateExamples bool, variant string) error {
	// Determine output module - default to github.com/tmc/appledocs/generated for now
	outputModule := "github.com/tmc/appledocs/generated"

	// Create generator instance
	gen := NewGenerator(framework, packageName, inputDir, outputModule, variant, withRefMethods, generateTests, generateExamples)
	gen.Functions = functions
	gen.Classes = classes
	gen.Protocols = protocols
	gen.Enums = enums
	gen.Typedefs = typedefs
	gen.Constants = constants

	// Try to use the module template if it exists
	if _, err := getTemplateVariant("module", variant); err == nil {
		if err := gen.GenerateTxtarFromModule(w); err != nil {
			return err
		}
		// Log any collected errors/warnings
		if verbose {
			for _, e := range gen.Errors {
				fmt.Fprintf(os.Stderr, "Warning: %v\n", e)
			}
		}
		return nil
	}

	// Fallback to individual file generation
	files := make(map[string][]byte)

	genFile := func(filename string, generator func(io.Writer) error) error {
		var buf bytes.Buffer
		if err := generator(&buf); err != nil {
			return err
		}
		files[filename] = buf.Bytes()
		return nil
	}

	if err := genFile("doc.gen.go", func(w io.Writer) error { return generateDoc(w, framework, packageName, inputDir, functions, variant) }); err != nil {
		return err
	}
	if err := genFile("types.gen.go", func(w io.Writer) error {
		return generateTypes(w, framework, packageName, functions, typedefs, withRefMethods, variant)
	}); err != nil {
		return err
	}
	if err := genFile("functions.gen.go", func(w io.Writer) error {
		return generateFunctions(w, framework, packageName, functions, withRefMethods, variant)
	}); err != nil {
		return err
	}
	if withRefMethods {
		if err := genFile("methods.gen.go", func(w io.Writer) error {
			return generateMethods(w, framework, packageName, functions, typedefs, variant)
		}); err != nil {
			return err
		}
	}
	if len(classes) > 0 {
		if err := genFile("classes.gen.go", func(w io.Writer) error { return generateClasses(w, framework, packageName, classes, variant) }); err != nil {
			return err
		}
	}
	if len(protocols) > 0 {
		if err := genFile("protocols.gen.go", func(w io.Writer) error { return generateProtocols(w, framework, packageName, protocols, variant) }); err != nil {
			return err
		}
	}

	// Write txtar format
	fmt.Fprintf(w, "# Generated bindings for %s framework\n", framework)
	fmt.Fprintf(w, "# Package: %s\n#\n", packageName)
	fmt.Fprintf(w, "# Functions: %d\n", len(functions))
	fmt.Fprintf(w, "# Classes: %d\n", len(classes))
	fmt.Fprintf(w, "# Protocols: %d\n", len(protocols))
	fmt.Fprintf(w, "# Enums: %d\n\n", len(enums))

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
func generateDoc(w io.Writer, framework, packageName, inputDir string, functions []*occ2go.ParsedFunction, variant string) error {
	frameworkAbstract, frameworkURL, _ := loadFrameworkMetadata(inputDir, framework)

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

	// Load template with variant support
	templateContent, err := getTemplateVariant("doc.gen.go", variant)
	if err != nil {
		return err
	}
	tmpl, err := template.New("doc.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// generateTypes generates framework-specific type definitions
func generateTypes(w io.Writer, framework, packageName string, functions []*occ2go.ParsedFunction, typedefs []*occ2go.ParsedTypedef, withRefMethods bool, variant string) error {
	// Build typedef names map to exclude from refTypes
	typedefNames := make(map[string]bool)
	for _, typedef := range typedefs {
		if typedef.Name != "" {
			typedefNames[typedef.Name] = true
		}
	}
	refTypes := extractRefTypes(functions, getFrameworkPrefix(framework), typedefNames)

	// Create a minimal generator for template execution
	gen := &Generator{
		Framework:      framework,
		PackageName:    packageName,
		Functions:      functions,
		Typedefs:       typedefs,
		WithRefMethods: withRefMethods,
	}
	gen.refTypes = refTypes

	data := gen

	// Load template with variant support
	templateContent, err := getTemplateVariant("types.gen.go", variant)
	if err != nil {
		// Fallback: write empty types file if template not found
		fmt.Fprintf(w, "// Code generated from Apple documentation for %s. DO NOT EDIT.\n\n", framework)
		fmt.Fprintf(w, "package %s\n", packageName)
		return nil
	}

	tmpl, err := template.New("types.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// generateFunctions generates function bindings
func generateFunctions(w io.Writer, framework, packageName string, functions []*occ2go.ParsedFunction, withRefMethods bool, variant string) error {
	data := struct {
		Framework      string
		PackageName    string
		Count          int
		Functions      []*occ2go.ParsedFunction
		WithRefMethods bool
	}{framework, packageName, len(functions), functions, withRefMethods}

	// Load template with variant support
	templateContent, err := getTemplateVariant("functions.gen.go", variant)
	if err != nil {
		return err
	}
	tmpl, err := template.New("functions.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// generateClasses generates class declarations
func generateClasses(w io.Writer, framework, packageName string, classes []*occ2go.ParsedClass, variant string) error {
	data := struct {
		Framework   string
		PackageName string
		Count       int
		Classes     []*occ2go.ParsedClass
	}{framework, packageName, len(classes), classes}

	// Load template with variant support
	templateContent, err := getTemplateVariant("classes.gen.go", variant)
	if err != nil {
		return err
	}
	tmpl, err := template.New("classes.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// generateProtocols generates protocol declarations
func generateProtocols(w io.Writer, framework, packageName string, protocols []*occ2go.ParsedProtocol, variant string) error {
	data := struct {
		Framework   string
		PackageName string
		Count       int
		Protocols   []*occ2go.ParsedProtocol
	}{framework, packageName, len(protocols), protocols}

	// Load template with variant support
	templateContent, err := getTemplateVariant("protocols.gen.go", variant)
	if err != nil {
		return err
	}
	tmpl, err := template.New("protocols.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// generateMethods generates method-style wrappers
func generateMethods(w io.Writer, framework, packageName string, functions []*occ2go.ParsedFunction, typedefs []*occ2go.ParsedTypedef, variant string) error {
	// Group functions by type
	typeMethods := groupFunctionsByType(functions, framework)

	// Build a map from type names to their underlying ref type
	typeToRef := make(map[string]string)
	// Build typedef names map to exclude from refTypes
	typedefNames := make(map[string]bool)
	for _, typedef := range typedefs {
		if typedef.Name != "" {
			typedefNames[typedef.Name] = true
		}
	}
	refTypes := extractRefTypes(functions, getFrameworkPrefix(framework), typedefNames)

	for _, refType := range refTypes {
		// Extract type name from ref type (e.g., CGContextRef -> Context)
		prefix := getFrameworkPrefix(framework)
		if strings.HasPrefix(refType, prefix) && strings.HasSuffix(refType, "Ref") {
			typeName := strings.TrimSuffix(strings.TrimPrefix(refType, prefix), "Ref")
			typeToRef[typeName] = refType
		}
	}

	data := struct {
		Framework   string
		PackageName string
		TypeMethods map[string][]*occ2go.ParsedFunction
		TypeToRef   map[string]string
	}{framework, packageName, typeMethods, typeToRef}

	// Load template with variant support
	templateContent, err := getTemplateVariant("methods.gen.go", variant)
	if err != nil {
		return err
	}
	tmpl, err := template.New("methods.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// getFrameworkPrefix returns the common type prefix for a framework
// Deprecated: Use GetFrameworkPrefix from registry instead
func getFrameworkPrefix(framework string) string {
	return GetFrameworkPrefix(framework)
}

// extractRefTypes extracts all Ref types from function signatures
func extractRefTypes(functions []*occ2go.ParsedFunction, prefix string, typedefNames map[string]bool) []string {
	refTypesMap := make(map[string]bool)

	for _, fn := range functions {
		if strings.HasPrefix(fn.ReturnType, prefix) && strings.HasSuffix(fn.ReturnType, "Ref") {
			// Skip if this type is defined as a typedef
			if typedefNames != nil && typedefNames[fn.ReturnType] {
				continue
			}
			refTypesMap[fn.ReturnType] = true
		}
		for _, param := range fn.Parameters {
			typ := strings.TrimSpace(param.Type)
			if strings.HasPrefix(typ, prefix) && strings.HasSuffix(typ, "Ref") {
				// Skip if this type is defined as a typedef
				if typedefNames != nil && typedefNames[typ] {
					continue
				}
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

// loadFrameworkMetadata loads the framework-level JSON to extract abstract, URL, and platform info
func loadFrameworkMetadata(inputDir, framework string) (abstract string, docURL string, iOSOnly bool) {
	frameworkPath := filepath.Join(inputDir, framework+".json")
	data, err := os.ReadFile(frameworkPath)
	if err != nil {
		// Try lowercase
		frameworkPath = filepath.Join(inputDir, strings.ToLower(framework)+".json")
		data, err = os.ReadFile(frameworkPath)
		if err != nil {
			return "", "", false
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
		Metadata struct {
			Platforms []struct {
				Name         string `json:"name"`
				IntroducedAt string `json:"introducedAt"`
			} `json:"platforms"`
		} `json:"metadata"`
		PrimaryContentSections []struct {
			Kind    string `json:"kind"`
			Content []struct {
				Type    string `json:"type"`
				Style   string `json:"style"`
				Name    string `json:"name"`
				Content []struct {
					Type          string `json:"type"`
					InlineContent []struct {
						Type string `json:"type"`
						Text string `json:"text"`
					} `json:"inlineContent"`
				} `json:"content"`
			} `json:"content"`
		} `json:"primaryContentSections"`
	}

	if err := json.Unmarshal(data, &doc); err != nil {
		return "", "", false
	}

	// Extract abstract text
	for _, item := range doc.Abstract {
		if item.Type == "text" && item.Text != "" {
			abstract = item.Text
			break
		}
	}

	// Check if this is an iOS-only framework (no macOS support)
	hasMacOS := false
	for _, platform := range doc.Metadata.Platforms {
		if platform.Name == "macOS" {
			hasMacOS = true
			break
		}
	}

	if !hasMacOS && len(doc.Metadata.Platforms) > 0 {
		// iOS-only framework - skip generation for macOS bindings
		platformNames := make([]string, 0, len(doc.Metadata.Platforms))
		for _, p := range doc.Metadata.Platforms {
			platformNames = append(platformNames, p.Name)
		}

		if verbose {
			fmt.Fprintf(os.Stderr, "Framework %s is iOS-only (platforms: %s). Skipping generation for macOS.\n",
				framework, strings.Join(platformNames, ", "))
		}

		docURL = occ2go.ConvertDocURLToWeb(doc.Identifier.URL)
		return abstract, docURL, true // iOS-only
	}

	// Extract deprecation information from Important asides
	deprecationText := extractDeprecationNotice(doc.PrimaryContentSections)
	if deprecationText != "" {
		if verbose {
			fmt.Fprintf(os.Stderr, "Framework %s is deprecated: %s\n", framework, deprecationText)
		}
		// Update config with deprecation information
		if config != nil {
			fwConfig := config.Frameworks[framework]
			fwConfig.Deprecated = true
			fwConfig.DeprecationReason = deprecationText
			config.Frameworks[framework] = fwConfig
		}
	}

	docURL = occ2go.ConvertDocURLToWeb(doc.Identifier.URL)
	return abstract, docURL, false // Not iOS-only
}

// extractDeprecationNotice extracts deprecation text from Important asides in content sections
func extractDeprecationNotice(sections []struct {
	Kind    string `json:"kind"`
	Content []struct {
		Type    string `json:"type"`
		Style   string `json:"style"`
		Name    string `json:"name"`
		Content []struct {
			Type          string `json:"type"`
			InlineContent []struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"inlineContent"`
		} `json:"content"`
	} `json:"content"`
}) string {
	for _, section := range sections {
		if section.Kind != "content" {
			continue
		}
		for _, content := range section.Content {
			if content.Type == "aside" && content.Style == "important" {
				// Extract text from inline content
				var texts []string
				for _, para := range content.Content {
					if para.Type == "paragraph" {
						for _, inline := range para.InlineContent {
							if inline.Type == "text" && inline.Text != "" {
								texts = append(texts, inline.Text)
							}
						}
					}
				}
				deprecationText := strings.Join(texts, "")
				// Check if it mentions "Do not use" or deprecation
				if strings.Contains(deprecationText, "Do not use") || strings.Contains(deprecationText, "deprecated") {
					// Clean up the text
					deprecationText = strings.TrimSpace(deprecationText)
					deprecationText = strings.ReplaceAll(deprecationText, "  ", " ")
					return deprecationText
				}
			}
		}
	}
	return ""
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

// generateTestsFile generates test file
func generateTestsFile(w io.Writer, framework, packageName string, functions []*occ2go.ParsedFunction, variant string) error {
	data := struct {
		Framework   string
		PackageName string
		Functions   []*occ2go.ParsedFunction
	}{framework, packageName, functions}

	// Load template with variant support
	templateContent, err := getTemplateVariant("functions_test.gen.go", variant)
	if err != nil {
		return err
	}
	tmpl, err := template.New("functions_test.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// generateExamplesFile generates examples file
func generateExamplesFile(w io.Writer, framework, packageName string, functions []*occ2go.ParsedFunction, variant string) error {
	data := struct {
		Framework   string
		PackageName string
		Functions   []*occ2go.ParsedFunction
	}{framework, packageName, functions}

	// Load template with variant support
	templateContent, err := getTemplateVariant("examples_test.gen.go", variant)
	if err != nil {
		return err
	}
	tmpl, err := template.New("examples_test.gen.go").Funcs(templateFuncs).Parse(templateContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// generateObjcRuntimePackage generates the framework-independent objc runtime package
// by extracting all templates with the objc/ prefix from the template archive.
func generateObjcRuntimePackage(outputDir string) error {
	// Collect all templates that start with "objc/"
	objcFiles := make(map[string][]byte)

	for _, file := range templateArchive.Files {
		if strings.HasPrefix(file.Name, "objc/") {
			// Strip the "objc/" prefix to get the filename
			filename := strings.TrimPrefix(file.Name, "objc/")
			objcFiles[filename] = file.Data
		}
	}

	if len(objcFiles) == 0 {
		return fmt.Errorf("no objc/ templates found in template archive")
	}

	// Create output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Write each file
	for name, data := range objcFiles {
		filePath := filepath.Join(outputDir, name)
		if err := os.WriteFile(filePath, data, 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", name, err)
		}
		if verbose {
			fmt.Fprintf(os.Stderr, "Generated %s\n", filePath)
		}
	}

	return nil
}

// enrichEnumValues calls the extract-enum-values tool to get actual enum values from macOS SDK headers.
// It populates the IntValue field of each ParsedEnumCase with the resolved integer value.
func enrichEnumValues(framework string, enums []*occ2go.ParsedEnum, verbose bool) error {
	// Use the appledocs extract-enums subcommand
	extractTool := "appledocs"

	enrichedCount := 0
	failedCount := 0

	// Process each enum
	for _, enum := range enums {
		// Call extract-enum-values to get actual values from SDK headers
		// This works even for enums with no cases from documentation
		// Try with NS prefix if enum name doesn't already have it
		enumNameForSDK := enum.Name
		if !strings.HasPrefix(enumNameForSDK, "NS") {
			enumNameForSDK = "NS" + enumNameForSDK
		}
		cmd := exec.Command(extractTool, "extract-enums", framework, enumNameForSDK)
		output, err := cmd.Output()
		if err != nil {
			failedCount++
			// Enum not found in SDK - this is expected for many enums
			continue
		}

		// Parse JSON output
		var result struct {
			Name   string `json:"name"`
			Values []struct {
				Name  string `json:"name"`
				Value int    `json:"value"`
			} `json:"values"`
		}
		if err := json.Unmarshal(output, &result); err != nil {
			failedCount++
			if verbose {
				fmt.Fprintf(os.Stderr, "Warning: failed to parse JSON for enum %s: %v\n", enum.Name, err)
			}
			continue
		}

		// If enum has no cases from documentation, create them from SDK extraction
		if len(enum.Cases) == 0 && len(result.Values) > 0 {
			if verbose {
				fmt.Fprintf(os.Stderr, "Creating %d cases for enum %s from SDK headers\n", len(result.Values), enum.Name)
			}
			for _, v := range result.Values {
				enum.Cases = append(enum.Cases, &occ2go.ParsedEnumCase{
					Name:     v.Name,
					IntValue: v.Value,
				})
			}
			if verbose {
				fmt.Fprintf(os.Stderr, "DEBUG: After appending, enum %s now has %d cases\n", enum.Name, len(enum.Cases))
			}
			enrichedCount++
			continue
		}

		// Build a map of case names to values for existing cases
		valueMap := make(map[string]int)
		for _, v := range result.Values {
			valueMap[v.Name] = v.Value
		}

		// Populate IntValue field for each case
		matchedCases := 0
		for _, enumCase := range enum.Cases {
			// Try exact match first
			if val, ok := valueMap[enumCase.Name]; ok {
				enumCase.IntValue = val
				matchedCases++
				continue
			}
			// Try with NS prefix (SDK uses NSEnumCase but docs might use EnumCase)
			if val, ok := valueMap["NS"+enumCase.Name]; ok {
				enumCase.IntValue = val
				matchedCases++
				continue
			}
			// Try without NS prefix (in case enum case has NS but value map doesn't)
			nameWithoutNS := strings.TrimPrefix(enumCase.Name, "NS")
			if nameWithoutNS != enumCase.Name {
				if val, ok := valueMap[nameWithoutNS]; ok {
					enumCase.IntValue = val
					matchedCases++
				}
			}
		}

		if matchedCases > 0 {
			enrichedCount++
			if verbose {
				fmt.Fprintf(os.Stderr, "Enriched %d/%d values for enum %s\n", matchedCases, len(enum.Cases), enum.Name)
			}
		}
	}

	if verbose && enrichedCount > 0 {
		fmt.Fprintf(os.Stderr, "Successfully enriched %d enums (failed: %d)\n", enrichedCount, failedCount)
	}

	return nil
}
