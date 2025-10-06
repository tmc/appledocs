// Demonstrate code generation from Apple documentation.
// This shows how tools like DarwinKit can generate Go bindings.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/tmc/appledocs"
)

func main() {
	docsPath := flag.String("docs", "output/tutorials/data/documentation", "Path to documentation directory")
	framework := flag.String("framework", "Foundation", "Framework to generate from")
	class := flag.String("class", "NSString", "Class to generate (optional)")
	output := flag.String("output", "", "Output file (default: stdout)")
	flag.Parse()

	// Open the documentation
	fsys, err := appledocs.Open(*docsPath)
	if err != nil {
		log.Fatalf("Failed to open docs: %v", err)
	}

	// Load framework
	fwDoc, err := appledocs.GetFramework(fsys, *framework)
	if err != nil {
		log.Fatalf("Failed to load framework: %v", err)
	}

	fmt.Fprintf(os.Stderr, "Generating Go bindings for %s...\n", fwDoc.Metadata.Title)

	// Find classes to generate
	var classes []classInfo
	for _, ref := range fwDoc.References {
		if ref.Role != "symbol" && ref.SymbolKind != "class" {
			continue
		}

		// If specific class requested, filter
		if *class != "" && ref.Title != *class {
			continue
		}

		// Load full class details
		classPath := *framework + "/" + ref.Title
		classDoc, err := appledocs.GetSymbol(fsys, classPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: couldn't load %s: %v\n", ref.Title, err)
			continue
		}

		// Extract class info
		info := classInfo{
			Name:       classDoc.Metadata.Title,
			ExternalID: classDoc.Metadata.ExternalID,
			Abstract:   extractAbstract(classDoc.Abstract),
		}

		// Extract methods
		for _, methodRef := range classDoc.References {
			if methodRef.Role != "symbol" {
				continue
			}

			if methodRef.SymbolKind == "method" {
				info.Methods = append(info.Methods, methodInfo{
					Name:      methodRef.Title,
					Signature: buildSignature(methodRef.Fragments),
				})
			} else if methodRef.SymbolKind == "property" {
				info.Properties = append(info.Properties, methodRef.Title)
			}
		}

		classes = append(classes, info)
	}

	if len(classes) == 0 {
		log.Fatal("No classes found")
	}

	// Generate code
	out := os.Stdout
	if *output != "" {
		f, err := os.Create(*output)
		if err != nil {
			log.Fatalf("Failed to create output file: %v", err)
		}
		defer f.Close()
		out = f
	}

	tmpl := template.Must(template.New("bindings").Parse(bindingTemplate))
	data := struct {
		Framework string
		Classes   []classInfo
	}{
		Framework: *framework,
		Classes:   classes,
	}

	if err := tmpl.Execute(out, data); err != nil {
		log.Fatalf("Template execution failed: %v", err)
	}

	fmt.Fprintf(os.Stderr, "Generated bindings for %d classes\n", len(classes))
}

type classInfo struct {
	Name       string
	ExternalID string
	Abstract   string
	Methods    []methodInfo
	Properties []string
}

type methodInfo struct {
	Name      string
	Signature string
}

func extractAbstract(abstract []appledocs.InlineContent) string {
	var parts []string
	for _, content := range abstract {
		if content.Text != "" {
			parts = append(parts, content.Text)
		}
	}
	return strings.Join(parts, " ")
}

func buildSignature(fragments []appledocs.Fragment) string {
	var parts []string
	for _, frag := range fragments {
		parts = append(parts, frag.Text)
	}
	return strings.Join(parts, "")
}

const bindingTemplate = `// Code generated from Apple Documentation for {{.Framework}}
// DO NOT EDIT

package {{.Framework | lower}}

{{range .Classes}}
// {{.Name}} - {{.Abstract}}
//
// External ID: {{.ExternalID}}
type {{.Name}} struct {
	// Internal pointer to Objective-C object
	ptr unsafe.Pointer
}

{{if .Properties}}
// Properties:
{{range .Properties}}//   - {{.}}
{{end}}{{end}}

{{if .Methods}}
// Methods:
{{range .Methods}}//   {{.Signature}}
{{end}}{{end}}
{{end}}
`
