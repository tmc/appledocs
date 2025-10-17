// debug-types prints the actual ObjC types found in init methods
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tmc/appledocs"
	"github.com/tmc/appledocs/occ2go"
)

func main() {
	framework := flag.String("framework", "AppKit", "Framework to analyze")
	inputDir := flag.String("input", "", "Input directory with JSON files")
	className := flag.String("class", "NSWindow", "Class to analyze")
	flag.Parse()

	// Default to cache directory if not specified
	if *inputDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to get home directory: %v\n", err)
			os.Exit(1)
		}
		*inputDir = filepath.Join(homeDir, ".appledocs/cache/developer.apple.com/tutorials/data/documentation")
	}

	// Open the appledocs filesystem
	fsys, err := appledocs.Open(*inputDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to open appledocs filesystem: %v\n", err)
		os.Exit(1)
	}

	// Collect all methods for this class
	classMethodsMap := make(map[string][]*occ2go.ParsedMethod)
	for _, doc := range appledocs.Symbols(fsys, *framework) {
		externalID := doc.Metadata.ExternalID

		// Check for methods: c:objc(cs)ClassName(im)methodName or c:objc(cs)ClassName(cm)methodName
		if strings.Contains(externalID, "(im)") || strings.Contains(externalID, "(cm)") {
			method, err := occ2go.ParseMethod(doc)
			if err == nil && method != nil {
				// Extract class name from external ID
				parts := strings.Split(externalID, "(")
				if len(parts) >= 2 {
					parsedClassName := strings.TrimPrefix(parts[1], "cs)")
					if parsedClassName == *className {
						classMethodsMap[parsedClassName] = append(classMethodsMap[parsedClassName], method)
					}
				}
			}
		}
	}

	// Filter for init methods
	var initMethods []*occ2go.ParsedMethod
	for _, method := range classMethodsMap[*className] {
		if strings.HasPrefix(method.Selector, "init") {
			initMethods = append(initMethods, method)
		}
	}

	if len(initMethods) == 0 {
		fmt.Printf("No init methods found for class %s\n", *className)
		return
	}

	fmt.Printf("Found %d init methods for %s:\n\n", len(initMethods), *className)
	for _, method := range initMethods {
		fmt.Printf("Selector: %s\n", method.Selector)
		fmt.Printf("  Return Type: '%s'\n", method.ReturnType)
		if len(method.Parameters) == 0 {
			fmt.Printf("  Parameters: (none)\n")
		} else {
			fmt.Printf("  Parameters:\n")
			for i, param := range method.Parameters {
				fmt.Printf("    [%d] Name: '%s', Type: '%s'\n", i, param.Name, param.Type)
			}
		}
		fmt.Printf("\n")
	}
}
