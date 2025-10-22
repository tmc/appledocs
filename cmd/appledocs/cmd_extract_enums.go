package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var extractEnumsCmd = &cobra.Command{
	Use:   "extract-enums <framework> <enum_name>",
	Short: "Extract enum values from a framework using clang preprocessor",
	Long: `Extract enum values from an Apple framework by preprocessing headers.

This command uses the clang preprocessor to extract enum definitions and their
values from framework headers. It's useful for discovering the actual integer
values assigned to enum cases.

Examples:
  appledocs extract-enums AppKit NSBackingStoreType
  appledocs extract-enums Foundation NSComparisonResult
  appledocs extract-enums CoreGraphics CGImageAlphaInfo`,
	Args: cobra.ExactArgs(2),
	RunE: runExtractEnums,
}

var (
	extractEnumsJSON bool
)

func init() {
	rootCmd.AddCommand(extractEnumsCmd)

	extractEnumsCmd.Flags().BoolVar(&extractEnumsJSON, "json", true,
		"output as JSON (default true)")
}

func runExtractEnums(cmd *cobra.Command, args []string) error {
	framework := args[0]
	enumName := args[1]

	enum, err := extractEnumValues(framework, enumName)
	if err != nil {
		return fmt.Errorf("extract enum values: %w", err)
	}

	if extractEnumsJSON {
		// Output as JSON
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(enum); err != nil {
			return fmt.Errorf("encode JSON: %w", err)
		}
	} else {
		// Output as human-readable format
		fmt.Printf("Enum: %s\n", enum.Name)
		fmt.Println("Values:")
		for _, v := range enum.Values {
			fmt.Printf("  %s = %d\n", v.Name, v.Value)
		}
	}

	return nil
}

type EnumValue struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Enum struct {
	Name   string      `json:"name"`
	Values []EnumValue `json:"values"`
}

func extractEnumValues(framework, enumName string) (*Enum, error) {
	// Run clang preprocessor
	cmd := exec.Command("clang", "-x", "objective-c", "-E", "-")
	cmd.Stdin = strings.NewReader(fmt.Sprintf("#import <%s/%s.h>", framework, framework))
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("running clang: %w", err)
	}

	// Parse the output to find the enum
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	inEnum := false
	// Match both Swift-style (enum Name : Type {) and C-style (typedef ... Name {)
	enumPattern := regexp.MustCompile(`(?:enum\s+` + regexp.QuoteMeta(enumName) + `\s*:\s*\w+|typedef.*?` + regexp.QuoteMeta(enumName) + `)\s*\{`)
	// Match explicit values including negative numbers and suffixes like L, UL
	valuePattern := regexp.MustCompile(`^\s*([A-Z][A-Za-z0-9_]+)\s*(?:__attribute__\(\([^)]+\)\)\s*)?=\s*(-?\d+[UL]*)`)
	// Match implicit values (no = assignment) - allow __attribute__, comma, or nothing (last enum case)
	namePattern := regexp.MustCompile(`^\s*([A-Z][A-Za-z0-9_]+)\s*(?:__attribute__|,|$)`)

	result := &Enum{
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
			if strings.Contains(line, "};") {
				break
			}

			// Try explicit value first
			matches := valuePattern.FindStringSubmatch(line)
			if len(matches) >= 3 {
				// Strip L, UL suffixes from the value string
				valueStr := strings.TrimSuffix(strings.TrimSuffix(matches[2], "UL"), "L")
				val, _ := strconv.Atoi(valueStr)
				result.Values = append(result.Values, EnumValue{
					Name:  matches[1],
					Value: val,
				})
				currentValue = val + 1
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
		return nil, fmt.Errorf("enum %s not found in %s framework", enumName, framework)
	}

	return result, nil
}
