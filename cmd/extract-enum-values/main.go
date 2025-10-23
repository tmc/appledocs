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
)

type EnumValue struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Enum struct {
	Name   string      `json:"name"`
	Values []EnumValue `json:"values"`
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <framework> <enum_name>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s AppKit NSBackingStoreType\n", os.Args[0])
		os.Exit(1)
	}

	framework := os.Args[1]
	enumName := os.Args[2]

	enum, err := extractEnumValues(framework, enumName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Output as JSON
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(enum); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
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
	// Match explicit values including negative numbers, hex values, and suffixes like L, UL
	valuePattern := regexp.MustCompile(`^\s*([A-Z][A-Za-z0-9_]+)\s*(?:__attribute__\(\([^)]+\)\)\s*)?=\s*(-?(?:0[xX][0-9A-Fa-f]+|\d+)[UL]*)`)
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
				var val int64
				var err error
				// Check if it's a hex value
				if strings.HasPrefix(valueStr, "0x") || strings.HasPrefix(valueStr, "0X") {
					val, err = strconv.ParseInt(valueStr[2:], 16, 64)
				} else if strings.HasPrefix(valueStr, "-0x") || strings.HasPrefix(valueStr, "-0X") {
					val, err = strconv.ParseInt(valueStr[3:], 16, 64)
					val = -val
				} else {
					val, err = strconv.ParseInt(valueStr, 10, 64)
				}
				if err != nil {
					continue
				}
				result.Values = append(result.Values, EnumValue{
					Name:  matches[1],
					Value: int(val),
				})
				currentValue = int(val) + 1
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

	return result, nil}
