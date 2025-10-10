// occ2go converts Objective-C declarations from Apple documentation to Go code.
//
// Usage:
//   occ2go <json-file>           # Parse and print Go code
//   occ2go -update <json-file>   # Update the JSON file with parsed info (future)
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/tmc/appledocs"
	"github.com/tmc/appledocs/occ2go"
)

func main() {
	update := flag.Bool("update", false, "Update JSON file with parsed information (future)")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "Usage: occ2go [-update] <json-file>\n")
		os.Exit(1)
	}

	// Configure slog
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}))
	slog.SetDefault(logger)

	jsonFile := flag.Arg(0)

	// Read JSON file
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		slog.Error("Failed to read JSON file", "file", jsonFile, "error", err)
		os.Exit(1)
	}

	// Parse as appledocs.Document
	var doc appledocs.Document
	if err := json.Unmarshal(data, &doc); err != nil {
		slog.Error("Failed to parse JSON", "file", jsonFile, "error", err)
		os.Exit(1)
	}

	// Parse document
	fn, cls, proto, err := occ2go.ParseDocument(&doc)

	// Determine framework from path (e.g., CoreGraphics/CGContextMoveToPoint.json -> CoreGraphics)
	framework := extractFramework(jsonFile)

	if err != nil {
		// Handle parse errors gracefully by printing a comment
		printParseError(&doc, framework, err)
		return
	}

	// Print Go code
	if fn != nil {
		printFunction(fn, framework)
	}
	if cls != nil {
		printClass(cls, framework)
	}
	if proto != nil {
		printProtocol(proto, framework)
	}

	if *update {
		slog.Info("Update mode not yet implemented")
	}
}

// extractFramework extracts the framework name from a file path
// e.g., "CoreGraphics/CGContextMoveToPoint.json" -> "CoreGraphics"
func extractFramework(path string) string {
	dir := filepath.Dir(path)
	framework := filepath.Base(dir)
	// Handle case where path is just a filename
	if framework == "." || framework == "/" {
		return "Unknown"
	}
	return framework
}

// printParseError prints a comment indicating a parse error
func printParseError(doc *appledocs.Document, framework string, err error) {
	// Print a comment showing the error
	fmt.Printf("// Parse error: %v\n", err)
	fmt.Printf("//\n")

	// Print symbol information if available
	if doc.Metadata.Title != "" {
		fmt.Printf("// Symbol: %s\n", doc.Metadata.Title)
	}
	if doc.Metadata.ExternalID != "" {
		fmt.Printf("// ExternalID: %s\n", doc.Metadata.ExternalID)
	}
	if doc.Identifier.URL != "" {
		fmt.Printf("// URL: %s\n", occ2go.ConvertDocURLToWeb(doc.Identifier.URL))
	}

	// Extract abstract if available
	abstract := occ2go.ExtractAbstract(doc.Abstract)
	if abstract != "" {
		fmt.Printf("//\n// %s\n", abstract)
	}

	fmt.Println()
}

// printFunction prints a Go function declaration
func printFunction(fn *occ2go.ParsedFunction, framework string) {
	// Print documentation
	if fn.Abstract != "" {
		fmt.Printf("// %s\n", fn.Abstract)
	}
	if fn.DocURL != "" {
		fmt.Printf("//\n// [Full Topic]: %s\n", fn.DocURL)
	}

	// Print function signature
	fmt.Printf("func %s(", fn.Name)

	// Parameters
	params := make([]string, len(fn.Parameters))
	for i, p := range fn.Parameters {
		goType := occ2go.MapCTypeToGo(p.Type, framework)
		params[i] = fmt.Sprintf("%s %s", p.Name, goType)
	}
	fmt.Print(strings.Join(params, ", "))
	fmt.Print(")")

	// Return type
	if fn.ReturnType != "" && fn.ReturnType != "void" {
		goReturnType := occ2go.MapCTypeToGo(fn.ReturnType, framework)
		if goReturnType != "" {
			fmt.Printf(" %s", goReturnType)
		}
	}

	fmt.Println(" {")
	fmt.Println("\t// TODO: Implementation")
	fmt.Println("}")
	fmt.Println()
}

// printClass prints a Go type declaration for a class
func printClass(cls *occ2go.ParsedClass, framework string) {
	// Print documentation
	if cls.Comment != "" {
		fmt.Printf("// %s\n", cls.Comment)
	}

	// Print type declaration
	fmt.Printf("type %s struct {\n", cls.Name)
	if cls.SuperClass != "" {
		fmt.Printf("\t%s\n", cls.SuperClass)
	}
	fmt.Println("\t// TODO: Add fields")
	fmt.Println("}")
	fmt.Println()
}

// printProtocol prints a Go interface declaration
func printProtocol(proto *occ2go.ParsedProtocol, framework string) {
	// Print documentation
	if proto.Comment != "" {
		fmt.Printf("// %s\n", proto.Comment)
	}

	// Print interface declaration
	fmt.Printf("type %s interface {\n", proto.Name)
	fmt.Println("\t// TODO: Add methods")
	fmt.Println("}")
	fmt.Println()
}
