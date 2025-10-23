package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/tmc/appledocs/occ2go"
)

// ExtractMissingEnums attempts to extract enum values for undefined types
// that look like they might be enums (e.g., NSEnumerationOptions, NSSortOptions)
func (g *Generator) ExtractMissingEnums() ([]*occ2go.ParsedEnum, error) {
	undefined := g.CollectUndefinedTypes()

	if os.Getenv("VERBOSE") == "1" {
		fmt.Fprintf(os.Stderr, "ExtractMissingEnums: Found %d undefined types\n", len(undefined))
		for name := range undefined {
			fmt.Fprintf(os.Stderr, "  - %s\n", name)
		}
	}

	extractedEnums := make([]*occ2go.ParsedEnum, 0)

	// Patterns that indicate a type is likely an enum
	enumPatterns := []*regexp.Regexp{
		regexp.MustCompile(`Options$`),      // NSEnumerationOptions, NSSortOptions
		regexp.MustCompile(`Flags$`),        // NSMatchingFlags
		regexp.MustCompile(`Mask$`),         // NSEventMask
		regexp.MustCompile(`^NS\w+Status$`), // NSURLSessionStatus
	}

	for _, ut := range undefined {
		// Check if this looks like an enum
		isLikelyEnum := false
		for _, pattern := range enumPatterns {
			if pattern.MatchString(ut.Name) {
				isLikelyEnum = true
				break
			}
		}

		if !isLikelyEnum {
			continue
		}

		// Try to extract enum values using clang preprocessor first
		enum, err := extractEnumWithPreprocessor(g.Framework, ut.Name)
		if err == nil && len(enum.Cases) > 0 {
			if os.Getenv("VERBOSE") == "1" {
				fmt.Fprintf(os.Stderr, "Extracted enum %s with %d cases using preprocessor\n", ut.Name, len(enum.Cases))
			}
			extractedEnums = append(extractedEnums, enum)
			continue
		}

		// If preprocessor failed, try runtime extraction for known constants
		// This requires knowing the constant names, which we can try to guess
		constants := guessEnumConstants(ut.Name)
		if len(constants) > 0 {
			enum, err := extractEnumWithRuntime(g.Framework, ut.Name, constants)
			if err == nil && len(enum.Cases) > 0 {
				if os.Getenv("VERBOSE") == "1" {
					fmt.Fprintf(os.Stderr, "Extracted enum %s with %d cases using runtime\n", ut.Name, len(enum.Cases))
				}
				extractedEnums = append(extractedEnums, enum)
			}
		}
	}

	return extractedEnums, nil
}

// extractEnumWithPreprocessor uses clang preprocessor to extract enum values
func extractEnumWithPreprocessor(framework, enumName string) (*occ2go.ParsedEnum, error) {
	// Run clang preprocessor
	cmd := exec.Command("clang", "-x", "objective-c", "-E", "-")
	cmd.Stdin = strings.NewReader(fmt.Sprintf("#import <%s/%s.h>\n", framework, framework))
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("clang preprocessor failed: %w", err)
	}

	// Parse the output to find the enum
	lines := strings.Split(string(output), "\n")
	inEnum := false
	enumPattern := regexp.MustCompile(`(?:enum\s+` + regexp.QuoteMeta(enumName) + `\s*:\s*\w+|typedef.*?` + regexp.QuoteMeta(enumName) + `)`)
	valuePattern := regexp.MustCompile(`^\s*([A-Z][A-Za-z0-9_]+)\s*(?:__attribute__\(\([^)]+\)\)\s*)?=\s*\(?\s*(\d+[UL]*)\s*<<\s*(\d+)\s*\)?`)
	simpleValuePattern := regexp.MustCompile(`^\s*([A-Z][A-Za-z0-9_]+)\s*(?:__attribute__\(\([^)]+\)\)\s*)?=\s*(-?(?:0[xX][0-9A-Fa-f]+|\d+)[UL]*)`)
	namePattern := regexp.MustCompile(`^\s*([A-Z][A-Za-z0-9_]+)`)

	enum := &occ2go.ParsedEnum{
		Name:  enumName,
		Cases: []*occ2go.ParsedEnumCase{},
	}
	currentValue := 0

	for _, line := range lines {
		if enumPattern.MatchString(line) {
			inEnum = true
			continue
		}

		if inEnum {
			if strings.Contains(line, "}") {
				break
			}

			// Try bit shift pattern (e.g., 1 << 0)
			matches := valuePattern.FindStringSubmatch(line)
			if len(matches) >= 4 {
				base := 1
				shift := 0
				fmt.Sscanf(matches[2], "%d", &base)
				fmt.Sscanf(matches[3], "%d", &shift)
				value := base << uint(shift)
				enum.Cases = append(enum.Cases, &occ2go.ParsedEnumCase{
					Name:  matches[1],
					Value: fmt.Sprintf("%d", value),
				})
				currentValue = value + 1
				continue
			}

			// Try simple value pattern
			matches = simpleValuePattern.FindStringSubmatch(line)
			if len(matches) >= 3 {
				valueStr := strings.TrimSuffix(strings.TrimSuffix(matches[2], "UL"), "L")
				value := 0
				// Parse hex or decimal
				if strings.HasPrefix(valueStr, "0x") || strings.HasPrefix(valueStr, "0X") {
					fmt.Sscanf(valueStr, "%x", &value)
				} else if strings.HasPrefix(valueStr, "-0x") || strings.HasPrefix(valueStr, "-0X") {
					fmt.Sscanf(valueStr[1:], "%x", &value)
					value = -value
				} else {
					fmt.Sscanf(valueStr, "%d", &value)
				}
				enum.Cases = append(enum.Cases, &occ2go.ParsedEnumCase{
					Name:  matches[1],
					Value: fmt.Sprintf("%d", value),
				})
				currentValue = value + 1
				continue
			}

			// Try implicit value
			matches = namePattern.FindStringSubmatch(line)
			if len(matches) >= 2 {
				enum.Cases = append(enum.Cases, &occ2go.ParsedEnumCase{
					Name:  matches[1],
					Value: fmt.Sprintf("%d", currentValue),
				})
				currentValue++
			}
		}
	}

	if len(enum.Cases) == 0 {
		return nil, fmt.Errorf("no enum cases found for %s", enumName)
	}

	return enum, nil
}

// extractEnumWithRuntime uses runtime evaluation to extract enum constant values
func extractEnumWithRuntime(framework, enumName string, constants []string) (*occ2go.ParsedEnum, error) {
	enum := &occ2go.ParsedEnum{
		Name:  enumName,
		Cases: make([]*occ2go.ParsedEnumCase, 0, len(constants)),
	}

	for _, constName := range constants {
		value, err := extractConstantValue(framework, constName, "uint")
		if err != nil {
			continue // Skip constants that fail
		}

		enum.Cases = append(enum.Cases, &occ2go.ParsedEnumCase{
			Name:  constName,
			Value: fmt.Sprintf("%d", value),
		})
	}

	return enum, nil
}

// extractConstantValue extracts a single constant value using runtime evaluation
func extractConstantValue(framework, constantName, constantType string) (uint64, error) {
	var printfFormat, castType string

	switch constantType {
	case "uint":
		printfFormat = "%lu"
		castType = "(unsigned long)" + constantName
	default:
		return 0, fmt.Errorf("unsupported type: %s", constantType)
	}

	// Generate C program to evaluate the constant
	source := fmt.Sprintf(`#import <Foundation/Foundation.h>
#import <%s/%s.h>
#include <stdio.h>

int main() {
    @autoreleasepool {
        printf("%s", %s);
    }
    return 0;
}
`, framework, framework, printfFormat, castType)

	// Write to temporary file
	tmpFile, err := os.CreateTemp("", "const-*.m")
	if err != nil {
		return 0, fmt.Errorf("create temp file: %w", err)
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)

	if _, err := tmpFile.WriteString(source); err != nil {
		return 0, fmt.Errorf("write temp file: %w", err)
	}
	tmpFile.Close()

	// Compile
	outFile := filepath.Join(os.TempDir(), "const-"+constantName)
	defer os.Remove(outFile)

	compileCmd := exec.Command("clang",
		"-framework", "Foundation",
		"-framework", framework,
		"-o", outFile,
		tmpPath)
	if output, err := compileCmd.CombinedOutput(); err != nil {
		return 0, fmt.Errorf("compile failed: %w\n%s", err, string(output))
	}

	// Run
	runCmd := exec.Command(outFile)
	output, err := runCmd.Output()
	if err != nil {
		return 0, fmt.Errorf("run failed: %w", err)
	}

	// Parse result
	result := strings.TrimSpace(string(output))
	var value uint64
	fmt.Sscanf(result, "%d", &value)

	return value, nil
}

// guessEnumConstants tries to guess the constant names for an enum type
func guessEnumConstants(enumName string) []string {
	// Known enum constants for common types
	knownConstants := map[string][]string{
		"NSEnumerationOptions": {
			"NSEnumerationConcurrent",
			"NSEnumerationReverse",
		},
		"NSSortOptions": {
			"NSSortConcurrent",
			"NSSortStable",
		},
	}

	if constants, ok := knownConstants[enumName]; ok {
		return constants
	}

	return nil
}
