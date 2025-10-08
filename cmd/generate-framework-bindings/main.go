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
	err := filepath.Walk(frameworkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".json") {
			fn, err := processJSONFile(path)
			if err == nil && fn != nil {
				functions = append(functions, fn)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking framework directory: %v", err)
	}

	log.Printf("Found %d functions in %s", len(functions), *framework)

	// Create output directory
	outDir := filepath.Join(*outputDir, strings.ToLower(*framework))
	if err := os.MkdirAll(outDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Generate bindings based on style
	switch *style {
	case "purego":
		generatePuregoBindings(functions, outDir, *framework)
	case "darwinkit":
		generateDarkwinKitBindings(functions, outDir, *framework)
	case "simple":
		generateSimpleBindings(functions, outDir, *framework)
	default:
		log.Fatalf("Unknown style: %s", *style)
	}

	log.Printf("Generated bindings in %s", outDir)
}

type ParsedFunction struct {
	Name       string
	ReturnType string
	Parameters []Parameter
	Comment    string
}

type Parameter struct {
	Name string
	Type string
}

type AppleDoc struct {
	Metadata struct {
		ExternalID string `json:"externalID"`
		Title      string `json:"title"`
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

type Token struct {
	Kind string `json:"kind"`
	Text string `json:"text"`
}

func processJSONFile(path string) (*ParsedFunction, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var doc AppleDoc
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, err
	}

	// Skip if not a C function (has external ID starting with c:@F@)
	if !strings.HasPrefix(doc.Metadata.ExternalID, "c:@F@") {
		return nil, fmt.Errorf("not a C function")
	}

	// Try to get ObjectiveC variant first
	tokens := getObjectiveCVariant(&doc)
	if tokens == nil {
		// Fall back to primary declarations
		if len(doc.PrimaryContentSections) > 0 && len(doc.PrimaryContentSections[0].Declarations) > 0 {
			tokens = doc.PrimaryContentSections[0].Declarations[0].Tokens
		}
	}

	if tokens == nil || len(tokens) == 0 {
		return nil, fmt.Errorf("no declaration found")
	}

	fn, err := parseDeclaration(tokens)
	if err != nil {
		return nil, err
	}

	return fn, nil
}

func getObjectiveCVariant(doc *AppleDoc) []Token {
	for _, variant := range doc.VariantOverrides {
		for _, trait := range variant.Traits {
			if trait.InterfaceLanguage == "occ" {
				for _, patch := range variant.Patch {
					if patch.Op == "replace" && strings.Contains(patch.Path, "declarations/0/tokens") {
						var tokens []Token
						if err := json.Unmarshal(patch.Value, &tokens); err == nil {
							return tokens
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

	// Find function name and parse signature
	var i int
	for i < len(tokens) {
		tok := tokens[i]

		if tok.Kind == "identifier" && fn.Name == "" {
			fn.Name = tok.Text
			i++
			continue
		}

		if tok.Kind == "text" && tok.Text == "(" && fn.Name != "" {
			// Parse parameters
			i++
			break
		}

		// Collect return type
		if fn.Name == "" {
			if fn.ReturnType != "" {
				fn.ReturnType += " "
			}
			fn.ReturnType += tok.Text
		}

		i++
	}

	// Parse parameters
	for i < len(tokens) {
		if tokens[i].Kind == "text" && tokens[i].Text == ")" {
			break
		}

		param := Parameter{}
		paramType := ""

		// Collect type tokens until we hit internalParam or identifier
		for i < len(tokens) && !(tokens[i].Kind == "internalParam" || (tokens[i].Kind == "identifier" && paramType != "")) {
			if tokens[i].Kind == "text" && (tokens[i].Text == "," || tokens[i].Text == ")") {
				break
			}
			if paramType != "" && tokens[i].Kind != "text" {
				paramType += " "
			}
			if tokens[i].Text != "" && tokens[i].Text != "(" {
				paramType += tokens[i].Text
			}
			i++
		}

		// Get parameter name
		if i < len(tokens) && tokens[i].Kind == "internalParam" {
			param.Name = tokens[i].Text
			i++
		} else if i < len(tokens) && tokens[i].Kind == "identifier" {
			param.Name = tokens[i].Text
			i++
		}

		param.Type = strings.TrimSpace(paramType)
		if param.Type != "" {
			fn.Parameters = append(fn.Parameters, param)
		}

		// Skip comma
		if i < len(tokens) && tokens[i].Text == "," {
			i++
		}
	}

	fn.ReturnType = strings.TrimSpace(fn.ReturnType)

	// Skip functions with colons (ObjC selectors)
	if strings.Contains(fn.Name, ":") {
		return nil, fmt.Errorf("skipping ObjC selector: %s", fn.Name)
	}

	return fn, nil
}

func generatePuregoBindings(functions []*ParsedFunction, outputDir, framework string) {
	// Generate a simple file with function declarations
	filename := filepath.Join(outputDir, "bindings.go")
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer f.Close()

	fmt.Fprintf(f, "// Code generated from Apple documentation for %s. DO NOT EDIT.\n\n", framework)
	fmt.Fprintf(f, "package %s\n\n", strings.ToLower(framework))
	fmt.Fprintf(f, "import (\n")
	fmt.Fprintf(f, "\t\"unsafe\"\n")
	fmt.Fprintf(f, "\t\"github.com/ebitengine/purego\"\n")
	fmt.Fprintf(f, ")\n\n")

	// Generate type definitions for common CG types
	fmt.Fprintf(f, "// Common types\n")
	fmt.Fprintf(f, "type CGFloat float64\n")
	fmt.Fprintf(f, "type CGContextRef unsafe.Pointer\n")
	fmt.Fprintf(f, "type CGColorRef unsafe.Pointer\n")
	fmt.Fprintf(f, "type CGColorSpaceRef unsafe.Pointer\n")
	fmt.Fprintf(f, "type CGPathRef unsafe.Pointer\n")
	fmt.Fprintf(f, "type CGImageRef unsafe.Pointer\n\n")

	fmt.Fprintf(f, "var lib uintptr\n\n")
	fmt.Fprintf(f, "func init() {\n")
	fmt.Fprintf(f, "\tvar err error\n")
	fmt.Fprintf(f, "\tlib, err = purego.Dlopen(\"/System/Library/Frameworks/%s.framework/%s\", purego.RTLD_LAZY|purego.RTLD_GLOBAL)\n", framework, framework)
	fmt.Fprintf(f, "\tif err != nil {\n\t\tpanic(err)\n\t}\n")
	fmt.Fprintf(f, "}\n\n")

	fmt.Fprintf(f, "// %s Functions\n\n", framework)

	// Generate function list as comments
	for i, fn := range functions {
		if i >= 20 { // Limit output for now
			break
		}
		if fn.Name == "" {
			continue
		}

		fmt.Fprintf(f, "// %s", fn.Name)
		if len(fn.Parameters) > 0 {
			fmt.Fprintf(f, "(")
			for j, p := range fn.Parameters {
				if j > 0 {
					fmt.Fprintf(f, ", ")
				}
				fmt.Fprintf(f, "%s %s", p.Name, p.Type)
			}
			fmt.Fprintf(f, ")")
		}
		if fn.ReturnType != "" {
			fmt.Fprintf(f, " -> %s", fn.ReturnType)
		}
		fmt.Fprintf(f, "\n")
	}

	fmt.Fprintf(f, "\n// Add bindings here using purego.RegisterLibFunc\n")
}

func generateDarkwinKitBindings(functions []*ParsedFunction, outputDir, framework string) {
	log.Printf("Darwinkit style not fully implemented yet for %s", framework)
	generatePuregoBindings(functions, outputDir, framework)
}

func generateSimpleBindings(functions []*ParsedFunction, outputDir, framework string) {
	log.Printf("Simple style not fully implemented yet for %s", framework)
	generatePuregoBindings(functions, outputDir, framework)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
