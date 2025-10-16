// occ2go converts Objective-C declarations from Apple documentation to Go code.
//
// Usage:
//   occ2go <json-file>           # Parse and print Go code
//   occ2go -update <json-file>   # Update the JSON file with parsed info (future)
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/tmc/appledocs"
	"github.com/tmc/appledocs/occ2go"
)

func main() {
	update := flag.Bool("update", false, "Update JSON file with parsed information (future)")
	goTest := flag.Bool("go-test", false, "Run go test on generated code to validate syntax")
	debug := flag.Bool("debug", false, "Print debug information about parsed methods")
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

	// Also try to parse as method if it's a method
	var method *occ2go.ParsedMethod
	if strings.Contains(doc.Metadata.ExternalID, "(im)") || strings.Contains(doc.Metadata.ExternalID, "(cm)") {
		method, _ = occ2go.ParseMethod(&doc)
	}

	// Determine framework from path (e.g., CoreGraphics/CGContextMoveToPoint.json -> CoreGraphics)
	framework := extractFramework(jsonFile)

	// Print debug info if requested
	if *debug {
		printDebugInfo(&doc, fn, cls, proto, method)
		return
	}

	// Capture output in buffer
	var buf bytes.Buffer

	if err != nil {
		// Handle parse errors gracefully by printing a comment
		printParseErrorToBuf(&buf, &doc, framework, err)
	} else {
		// Print Go code to buffer
		if fn != nil {
			printFunctionToBuf(&buf, fn, framework)
		}
		if cls != nil {
			printClassToBuf(&buf, cls, framework)
		}
		if proto != nil {
			printProtocolToBuf(&buf, proto, framework)
		}
	}

	// Run goimports on the output
	formatted, err := runGoimports(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "goimports failed: %v\n", err)
		os.Exit(1)
	}

	// Run go test on the generated code (if requested)
	if *goTest {
		if err := runGoTest(formatted); err != nil {
			fmt.Fprintf(os.Stderr, "go test failed: %v\n", err)
			os.Exit(1)
		}
	}

	// Print formatted output
	fmt.Print(string(formatted))

	if *update {
		slog.Info("Update mode not yet implemented")
	}
}

// runGoimports runs goimports on the provided Go code
func runGoimports(code []byte) ([]byte, error) {
	cmd := exec.Command("goimports")
	cmd.Stdin = bytes.NewReader(code)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%w\nStderr: %s", err, stderr.String())
	}

	return stdout.Bytes(), nil
}

// runGoTest runs go test on the provided Go code by writing it to a temporary file
func runGoTest(code []byte) error {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "occ2go-test-*")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	// Write the code to a go file
	goFile := filepath.Join(tmpDir, "generated.go")

	// Wrap the code in a package for testing
	fullCode := fmt.Sprintf(`package occ2gotest

%s
`, code)

	if err := os.WriteFile(goFile, []byte(fullCode), 0644); err != nil {
		return fmt.Errorf("failed to write test file: %w", err)
	}

	// Initialize a go module in the temp directory
	modInit := exec.Command("go", "mod", "init", "occ2gotest")
	modInit.Dir = tmpDir
	if out, err := modInit.CombinedOutput(); err != nil {
		return fmt.Errorf("go mod init failed: %w\nOutput: %s", err, out)
	}

	// Run go build to check syntax
	cmd := exec.Command("go", "build", "-o", "/dev/null", ".")
	cmd.Dir = tmpDir
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w\nStderr: %s", err, stderr.String())
	}

	return nil
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

// printParseErrorToBuf prints a comment indicating a parse error to a buffer
func printParseErrorToBuf(buf *bytes.Buffer, doc *appledocs.Document, framework string, err error) {
	// Print a comment showing the error
	fmt.Fprintf(buf, "// Parse error: %v\n", err)
	fmt.Fprintf(buf, "//\n")

	// Print symbol information if available
	if doc.Metadata.Title != "" {
		fmt.Fprintf(buf, "// Symbol: %s\n", doc.Metadata.Title)
	}
	if doc.Metadata.ExternalID != "" {
		fmt.Fprintf(buf, "// ExternalID: %s\n", doc.Metadata.ExternalID)
	}
	if doc.Identifier.URL != "" {
		fmt.Fprintf(buf, "// URL: %s\n", occ2go.ConvertDocURLToWeb(doc.Identifier.URL))
	}

	// Extract abstract if available
	abstract := occ2go.ExtractAbstract(doc.Abstract)
	if abstract != "" {
		fmt.Fprintf(buf, "//\n// %s\n", abstract)
	}

	fmt.Fprintln(buf)
}

// printFunctionToBuf prints a Go function declaration to a buffer
func printFunctionToBuf(buf *bytes.Buffer, fn *occ2go.ParsedFunction, framework string) {
	// Print documentation
	if fn.Abstract != "" {
		fmt.Fprintf(buf, "// %s\n", fn.Abstract)
	}
	if fn.DocURL != "" {
		fmt.Fprintf(buf, "//\n// [Full Topic]: %s\n", fn.DocURL)
	}

	// Print function signature
	fmt.Fprintf(buf, "func %s(", fn.Name)

	// Parameters
	params := make([]string, len(fn.Parameters))
	for i, p := range fn.Parameters {
		goType := occ2go.MapCTypeToGo(p.Type, framework)
		params[i] = fmt.Sprintf("%s %s", p.Name, goType)
	}
	fmt.Fprint(buf, strings.Join(params, ", "))
	fmt.Fprint(buf, ")")

	// Return type
	if fn.ReturnType != "" && fn.ReturnType != "void" {
		goReturnType := occ2go.MapCTypeToGo(fn.ReturnType, framework)
		if goReturnType != "" {
			fmt.Fprintf(buf, " %s", goReturnType)
		}
	}

	fmt.Fprintln(buf, " {")
	fmt.Fprintln(buf, "\t// TODO: Implementation")
	fmt.Fprintln(buf, "}")
	fmt.Fprintln(buf)
}

// printClassToBuf prints a Go type declaration for a class to a buffer
func printClassToBuf(buf *bytes.Buffer, cls *occ2go.ParsedClass, framework string) {
	// Print documentation
	if cls.Comment != "" {
		fmt.Fprintf(buf, "// %s\n", cls.Comment)
	}

	// Print type declaration
	fmt.Fprintf(buf, "type %s struct {\n", cls.Name)
	if cls.SuperClass != "" {
		fmt.Fprintf(buf, "\t%s\n", cls.SuperClass)
	}
	fmt.Fprintln(buf, "\t// TODO: Add fields")
	fmt.Fprintln(buf, "}")
	fmt.Fprintln(buf)
}

// printProtocolToBuf prints a Go interface declaration to a buffer
func printProtocolToBuf(buf *bytes.Buffer, proto *occ2go.ParsedProtocol, framework string) {
	// Print documentation
	if proto.Comment != "" {
		fmt.Fprintf(buf, "// %s\n", proto.Comment)
	}

	// Print interface declaration
	fmt.Fprintf(buf, "type %s interface {\n", proto.Name)
	fmt.Fprintln(buf, "\t// TODO: Add methods")
	fmt.Fprintln(buf, "}")
	fmt.Fprintln(buf)
}

// printDebugInfo prints debug information about a parsed document
func printDebugInfo(doc *appledocs.Document, fn *occ2go.ParsedFunction, cls *occ2go.ParsedClass, proto *occ2go.ParsedProtocol, method *occ2go.ParsedMethod) {
	fmt.Printf("ExternalID: %s\n", doc.Metadata.ExternalID)
	fmt.Printf("Title: %s\n", doc.Metadata.Title)

	// Print tokens
	tokens := occ2go.GetObjectiveCVariant(doc)
	if tokens == nil {
		for _, section := range doc.PrimaryContentSections {
			if len(section.Declarations) > 0 && len(section.Declarations[0].Tokens) > 0 {
				tokens = section.Declarations[0].Tokens
				break
			}
		}
	}
	if tokens != nil {
		fmt.Printf("\nTokens (%d):\n", len(tokens))
		for i, tok := range tokens {
			fmt.Printf("  [%d] %-15s %q\n", i, tok.Kind, tok.Text)
		}
	}
	fmt.Println()

	if fn != nil {
		fmt.Printf("Function: %s\n", fn.Name)
		fmt.Printf("  ReturnType: %s\n", fn.ReturnType)
		fmt.Printf("  Parameters: %d\n", len(fn.Parameters))
		for i, p := range fn.Parameters {
			fmt.Printf("    [%d] %s: %s\n", i, p.Name, p.Type)
		}
	}

	if cls != nil {
		fmt.Printf("Class: %s\n", cls.Name)
		if cls.SuperClass != "" {
			fmt.Printf("  SuperClass: %s\n", cls.SuperClass)
		}
	}

	if proto != nil {
		fmt.Printf("Protocol: %s\n", proto.Name)
	}

	if method != nil {
		fmt.Printf("Method: %s\n", method.Name)
		fmt.Printf("  Selector: %s\n", method.Selector)
		fmt.Printf("  IsClassMethod: %v\n", method.IsClassMethod)
		fmt.Printf("  ReturnType: %s\n", method.ReturnType)
		fmt.Printf("  Parameters: %d\n", len(method.Parameters))
		for i, p := range method.Parameters {
			fmt.Printf("    [%d] %s: %s\n", i, p.Name, p.Type)
		}
	}
}
