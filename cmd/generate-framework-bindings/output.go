package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

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

	// Apply property overrides
	for _, cls := range gen.Classes {
		MergePropertyOverrides(framework, cls.Name, cls)
	}

	gen.prepare()

	// Generate stubs for missing parent classes
	stubs := gen.GenerateMissingParentStubs()
	if len(stubs) > 0 {
		gen.Classes = append(stubs, gen.Classes...)
	}

	// Use the unified module template
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

// EnumValue represents a single enum case with its integer value
type EnumValue struct {
	Name  string
	Value int
}

// EnumResult represents an enum with all its cases
type EnumResult struct {
	Name   string
	Values []EnumValue
}

// parseEnumFromPreprocessed extracts enum values from clang-preprocessed source
func parseEnumFromPreprocessed(preprocessedSource, enumName string) (*EnumResult, error) {
	// First, normalize the preprocessed source by splitting on common delimiters
	// This handles the case where the entire enum is on one line
	normalized := strings.ReplaceAll(preprocessedSource, ",", ",\n")
	normalized = strings.ReplaceAll(normalized, "{", "{\n")
	normalized = strings.ReplaceAll(normalized, "}", "\n}")

	scanner := bufio.NewScanner(strings.NewReader(normalized))
	inEnum := false
	// Match both Swift-style (enum Name : Type {) and C-style (typedef ... Name {)
	enumPattern := regexp.MustCompile(`(?:enum\s+` + regexp.QuoteMeta(enumName) + `\s*:\s*\w+|typedef.*?` + regexp.QuoteMeta(enumName) + `)`)
	// Match bit-shift pattern FIRST (before explicit values): Name = 1UL << 0
	bitShiftPattern := regexp.MustCompile(`^\s*([A-Z][A-Za-z0-9_]+)\s*(?:__attribute__\(\([^)]+\)\)\s*)?=\s*\(?\s*(\d+[UL]*)\s*<<\s*(\d+)\s*\)?`)
	// Match explicit values including negative numbers, hex values, and suffixes like L, UL
	// NOTE: We check bit-shift BEFORE this pattern in the loop, so order matters!
	valuePattern := regexp.MustCompile(`^\s*([A-Z][A-Za-z0-9_]+)\s*(?:__attribute__\(\([^)]+\)\)\s*)?=\s*(-?(?:0[xX][0-9A-Fa-f]+|\d+)[UL]*)`)
	// Match implicit values (no = assignment) - allow __attribute__, comma, or nothing (last enum case)
	namePattern := regexp.MustCompile(`^\s*([A-Z][A-Za-z0-9_]+)\s*(?:__attribute__|,|$)`)

	result := &EnumResult{
		Name:   enumName,
		Values: []EnumValue{},
	}
	currentValue := 0

	for scanner.Scan() {
		line := scanner.Text()

		if enumPattern.MatchString(line) {
			inEnum = true
			continue
		}

		if inEnum {
			// Check for closing brace (may have attributes between } and ;)
			if strings.Contains(line, "}") {
				break
			}

			// Try bit-shift pattern FIRST: Name = 1UL << 0
			matches := bitShiftPattern.FindStringSubmatch(line)
			if len(matches) >= 4 {
				// Strip UL/L suffixes from base value
				baseStr := strings.TrimSuffix(strings.TrimSuffix(matches[2], "UL"), "L")
				base := 1
				shift := 0
				fmt.Sscanf(baseStr, "%d", &base)
				fmt.Sscanf(matches[3], "%d", &shift)
				value := base << uint(shift)
				result.Values = append(result.Values, EnumValue{
					Name:  matches[1],
					Value: value,
				})
				currentValue = value + 1
				continue
			}

			// Try explicit value (hex or decimal)
			matches = valuePattern.FindStringSubmatch(line)
			if len(matches) >= 3 {
				// Strip L, UL suffixes from the value string
				valueStr := strings.TrimSuffix(strings.TrimSuffix(matches[2], "UL"), "L")
				var val int64
				var err error
				// Parse hex or decimal
				if strings.HasPrefix(valueStr, "0x") || strings.HasPrefix(valueStr, "0X") {
					val, err = strconv.ParseInt(valueStr[2:], 16, 64)
				} else if strings.HasPrefix(valueStr, "-0x") || strings.HasPrefix(valueStr, "-0X") {
					val, err = strconv.ParseInt(valueStr[3:], 16, 64)
					val = -val
				} else {
					val, err = strconv.ParseInt(valueStr, 10, 64)
				}
				if err == nil {
					result.Values = append(result.Values, EnumValue{
						Name:  matches[1],
						Value: int(val),
					})
					currentValue = int(val) + 1
				}
				continue
			}

			// Try implicit value (no = assignment)
			matches = namePattern.FindStringSubmatch(line)
			if len(matches) >= 2 {
				result.Values = append(result.Values, EnumValue{
					Name:  matches[1],
					Value: currentValue,
				})
				currentValue++
			}
		}
	}

	if len(result.Values) == 0 {
		return nil, fmt.Errorf("enum %s not found", enumName)
	}

	return result, nil
}

// Global cache for preprocessed framework headers (in-memory for this run)
var preprocessedCache = make(map[string]string)

// enrichEnumValues calls the extract-enum-values tool to get actual enum values from macOS SDK headers.
// It populates the IntValue field of each ParsedEnumCase with the resolved integer value.
func enrichEnumValues(framework string, enums []*occ2go.ParsedEnum, verbose bool) error {
	// Batch extract all enums with a single clang invocation for performance
	enrichedCount := 0
	failedCount := 0

	// Check cache first
	preprocessedSource, cached := preprocessedCache[framework]
	if !cached {
		// Run clang preprocessor once for all enums
		cmd := exec.Command("clang", "-x", "objective-c", "-E", "-")
		cmd.Stdin = strings.NewReader(fmt.Sprintf("#import <%s/%s.h>", framework, framework))
		cmd.Stderr = os.Stderr
		output, err := cmd.Output()
		if err != nil {
			if verbose {
				fmt.Fprintf(os.Stderr, "Warning: clang preprocessing failed for %s: %v\n", framework, err)
			}
			return err
		}
		preprocessedSource = string(output)
		preprocessedCache[framework] = preprocessedSource
	}

	// Process each enum by parsing the preprocessed output
	for _, enum := range enums {
		// Try with NS prefix if enum name doesn't already have it
		enumNameForSDK := enum.Name
		if !strings.HasPrefix(enumNameForSDK, "NS") {
			enumNameForSDK = "NS" + enumNameForSDK
		}

		// Extract enum values from preprocessed source
		result, err := parseEnumFromPreprocessed(preprocessedSource, enumNameForSDK)
		if err != nil {
			failedCount++
			// Enum not found in SDK - this is expected for many enums
			continue
		}

		Debug.EnumCases("parseEnumFromPreprocessed results", enum.Name, enumNameForSDK,
			"enumName", enum.Name,
			"enumNameForSDK", enumNameForSDK,
			"valueCount", len(result.Values))
		for _, v := range result.Values {
			Debug.EnumCases("extracted enum value", enum.Name, v.Name,
				"enumName", enum.Name,
				"valueName", v.Name,
				"value", v.Value)
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
			Debug.EnumCases("enum cases created from SDK", enum.Name, "",
				"enumName", enum.Name,
				"caseCount", len(enum.Cases))
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
				Debug.EnumCases("matched enum case", enum.Name, enumCase.Name,
					"enumName", enum.Name,
					"caseName", enumCase.Name,
					"value", val,
					"matchType", "exact")
				enumCase.IntValue = val
				matchedCases++
				continue
			}
			// Try with NS prefix (SDK uses NSEnumCase but docs might use EnumCase)
			if val, ok := valueMap["NS"+enumCase.Name]; ok {
				Debug.EnumCases("matched enum case", enum.Name, enumCase.Name,
					"enumName", enum.Name,
					"caseName", enumCase.Name,
					"value", val,
					"matchType", "with NS prefix")
				enumCase.IntValue = val
				matchedCases++
				continue
			}
			// Try without NS prefix (in case enum case has NS but value map doesn't)
			nameWithoutNS := strings.TrimPrefix(enumCase.Name, "NS")
			if nameWithoutNS != enumCase.Name {
				if val, ok := valueMap[nameWithoutNS]; ok {
					Debug.EnumCases("matched enum case", enum.Name, enumCase.Name,
						"enumName", enum.Name,
						"caseName", enumCase.Name,
						"value", val,
						"matchType", "without NS prefix")
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
