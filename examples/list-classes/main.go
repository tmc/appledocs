// Find all classes, protocols, and structs in a framework.
package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/tmc/appledocs"
)

func main() {
	docsPath := flag.String("docs", "output/tutorials/data/documentation", "Path to documentation directory")
	framework := flag.String("framework", "Foundation", "Framework to analyze")
	kind := flag.String("kind", "", "Filter by kind (class, protocol, struct, enum)")
	flag.Parse()

	// Open the documentation
	fsys, err := appledocs.Open(*docsPath)
	if err != nil {
		log.Fatalf("Failed to open docs: %v", err)
	}

	// Load the framework
	doc, err := appledocs.GetFramework(fsys, *framework)
	if err != nil {
		log.Fatalf("Failed to load framework %s: %v", *framework, err)
	}

	// Collect symbols by kind
	symbols := make(map[string][]string)

	for _, ref := range doc.References {
		if ref.Role != "symbol" {
			continue
		}

		symbolKind := ref.SymbolKind
		if symbolKind == "" {
			continue
		}

		// Filter by kind if specified
		if *kind != "" && !strings.EqualFold(symbolKind, *kind) {
			continue
		}

		symbols[symbolKind] = append(symbols[symbolKind], ref.Title)
	}

	// Sort and display
	fmt.Printf("Symbols in %s:\n\n", doc.Metadata.Title)

	// Get sorted list of kinds
	var kinds []string
	for k := range symbols {
		kinds = append(kinds, k)
	}
	sort.Strings(kinds)

	for _, k := range kinds {
		list := symbols[k]
		sort.Strings(list)

		fmt.Printf("%s (%d):\n", strings.Title(k), len(list))
		for _, name := range list {
			fmt.Printf("  - %s\n", name)
		}
		fmt.Println()
	}

	// Summary
	total := 0
	for _, list := range symbols {
		total += len(list)
	}
	fmt.Printf("Total: %d symbols\n", total)
}
