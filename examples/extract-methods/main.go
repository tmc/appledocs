// Extract and display method signatures for a class.
package main

import (
	"flag"
	"fmt"
	"strings"
	"log"
	"sort"
	"strings"

	"github.com/tmc/appledocs"
)

func main() {
	docsPath := flag.String("docs", "output/tutorials/data/documentation", "Path to documentation directory")
	class := flag.String("class", "Foundation/NSString", "Class to analyze (framework/ClassName)")
	methodType := flag.String("type", "", "Filter by type (instance, class, property)")
	flag.Parse()

	// Open the documentation
	fsys, err := appledocs.Open(*docsPath)
	if err != nil {
		log.Fatalf("Failed to open docs: %v", err)
	}

	// Load the class
	doc, err := appledocs.GetSymbol(fsys, *class)
	if err != nil {
		log.Fatalf("Failed to load class: %v", err)
	}

	fmt.Printf("Methods for %s\n\n", doc.Metadata.Title)

	// Categorize methods
	instanceMethods := []methodInfo{}
	classMethods := []methodInfo{}
	properties := []methodInfo{}
	other := []methodInfo{}

	for _, ref := range doc.References {
		if ref.Role != "symbol" {
			continue
		}

		method := methodInfo{
			name:       ref.Title,
			kind:       ref.SymbolKind,
			signature:  buildSignature(ref.Fragments),
			deprecated: isDeprecated(ref),
		}

		switch ref.SymbolKind {
		case "method":
			// Determine if instance or class method from signature
			if strings.HasPrefix(method.signature, "+") {
				classMethods = append(classMethods, method)
			} else {
				instanceMethods = append(instanceMethods, method)
			}
		case "property":
			properties = append(properties, method)
		default:
			other = append(other, method)
		}
	}

	// Display based on filter
	if *methodType == "" || *methodType == "instance" {
		displayMethods("Instance Methods", instanceMethods)
	}
	if *methodType == "" || *methodType == "class" {
		displayMethods("Class Methods", classMethods)
	}
	if *methodType == "" || *methodType == "property" {
		displayMethods("Properties", properties)
	}
	if *methodType == "" && len(other) > 0 {
		displayMethods("Other Symbols", other)
	}

	// Summary
	total := len(instanceMethods) + len(classMethods) + len(properties) + len(other)
	fmt.Printf("\nTotal: %d symbols\n", total)
}

type methodInfo struct {
	name       string
	kind       string
	signature  string
	deprecated bool
}

func displayMethods(title string, methods []methodInfo) {
	if len(methods) == 0 {
		return
	}

	// Sort by name
	sort.Slice(methods, func(i, j int) bool {
		return methods[i].name < methods[j].name
	})

	fmt.Printf("%s (%d):\n", title, len(methods))
	for _, m := range methods {
		deprecated := ""
		if m.deprecated {
			deprecated = " [DEPRECATED]"
		}

		if m.signature != "" {
			fmt.Printf("  %s%s\n", m.signature, deprecated)
		} else {
			fmt.Printf("  %s%s\n", m.name, deprecated)
		}
	}
	fmt.Println()
}

func buildSignature(fragments []appledocs.Fragment) string {
	if len(fragments) == 0 {
		return ""
	}

	var parts []string
	for _, frag := range fragments {
		parts = append(parts, frag.Text)
	}
	return strings.Join(parts, "")
}

func isDeprecated(ref appledocs.Reference) bool {
	// Check if any abstract mentions deprecation
	for _, content := range ref.Abstract {
		if strings.Contains(strings.ToLower(content.Text), "deprecat") {
			return true
		}
	}
	return false
}
