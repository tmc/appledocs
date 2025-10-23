package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var extractConstantCmd = &cobra.Command{
	Use:   "extract-constant <framework> <constant_name> <type>",
	Short: "Extract constant values using runtime evaluation",
	Long: `Extract constant values from Apple frameworks by compiling and running
a small C program. This handles constants that can't be extracted statically:

- String constants (NSErrorDomain, NSString constants)
- Bitmask enums with expressions (1UL << 0)
- Float constants with expressions (M_PI * 2)
- Other computed constants

Type can be: string, int, uint, float, double

Examples:
  appledocs extract-constant Foundation NSStreamSocketSSLErrorDomain string
  appledocs extract-constant Foundation NSEnumerationConcurrent uint
  appledocs extract-constant Foundation M_PI double`,
	Args: cobra.ExactArgs(3),
	RunE: runExtractConstant,
}

var (
	extractConstantJSON bool
)

func init() {
	rootCmd.AddCommand(extractConstantCmd)

	extractConstantCmd.Flags().BoolVar(&extractConstantJSON, "json", true,
		"output as JSON (default true)")
}

func runExtractConstant(cmd *cobra.Command, args []string) error {
	framework := args[0]
	constantName := args[1]
	constantType := args[2]

	result, err := extractConstantValue(framework, constantName, constantType)
	if err != nil {
		return fmt.Errorf("extract constant: %w", err)
	}

	if extractConstantJSON {
		output := map[string]interface{}{
			"name":      constantName,
			"value":     result.Value,
			"type":      constantType,
			"framework": framework,
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(output); err != nil {
			return fmt.Errorf("encode JSON: %w", err)
		}
	} else {
		fmt.Printf("%s = %v (%s)\n", constantName, result.Value, constantType)
	}

	return nil
}

type ConstantValue struct {
	Value interface{}
	Type  string
}

func extractConstantValue(framework, constantName, constantType string) (*ConstantValue, error) {
	// Determine printf format and type cast based on type
	var printfFormat, castType string

	switch constantType {
	case "string":
		printfFormat = "%s"
		castType = "[" + constantName + " UTF8String]"
	case "int":
		printfFormat = "%ld"
		castType = "(long)" + constantName
	case "uint":
		printfFormat = "%lu"
		castType = "(unsigned long)" + constantName
	case "float", "double":
		printfFormat = "%f"
		castType = "(double)" + constantName
	default:
		return nil, fmt.Errorf("unsupported type: %s (use: string, int, uint, float, double)", constantType)
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
		return nil, fmt.Errorf("create temp file: %w", err)
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)

	if _, err := tmpFile.WriteString(source); err != nil {
		return nil, fmt.Errorf("write temp file: %w", err)
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
		return nil, fmt.Errorf("compile failed: %w\n%s", err, string(output))
	}

	// Run
	runCmd := exec.Command(outFile)
	output, err := runCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("run failed: %w", err)
	}

	// Parse result based on type
	result := strings.TrimSpace(string(output))

	var value interface{}
	switch constantType {
	case "string":
		value = result
	case "int":
		fmt.Sscanf(result, "%d", &value)
	case "uint":
		var u uint64
		fmt.Sscanf(result, "%d", &u)
		value = u
	case "float", "double":
		var f float64
		fmt.Sscanf(result, "%f", &f)
		value = f
	}

	return &ConstantValue{
		Value: value,
		Type:  constantType,
	}, nil
}
