// Search for symbols across frameworks.
package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/tmc/appledocs"
)

func main() {
	docsPath := flag.String("docs", "output/tutorials/data/documentation", "Path to documentation directory")
	query := flag.String("query", "", "Search query (required)")
	framework := flag.String("framework", "", "Limit to specific framework (optional)")
	maxResults := flag.Int("max", 50, "Maximum results to show")
	flag.Parse()

	if *query == "" {
		log.Fatal("Please provide a search query with -query")
	}

	// Open the documentation
	fsys, err := appledocs.Open(*docsPath)
	if err != nil {
		log.Fatalf("Failed to open docs: %v", err)
	}

	// Determine which frameworks to search
	var frameworks []string
	if *framework != "" {
		frameworks = []string{*framework}
	} else {
		frameworks, err = appledocs.ListFrameworks(fsys)
		if err != nil {
			log.Fatalf("Failed to list frameworks: %v", err)
		}
	}

	fmt.Printf("Searching for '%s'...\n\n", *query)

	results := 0
	queryLower := strings.ToLower(*query)

	// Search each framework
	for _, fw := range frameworks {
		matches, err := appledocs.SearchSymbols(fsys, fw, queryLower)
		if err != nil {
			continue
		}

		if len(matches) == 0 {
			continue
		}

		fmt.Printf("%s (%d matches):\n", fw, len(matches))

		// Show up to maxResults per framework
		shown := 0
		for _, symbolName := range matches {
			if shown >= *maxResults {
				fmt.Printf("  ... and %d more\n", len(matches)-shown)
				break
			}

			// Try to load symbol info
			info, err := appledocs.GetSymbolInfo(fsys, fw+"/"+symbolName)
			if err != nil {
				fmt.Printf("  - %s\n", symbolName)
			} else {
				kind := info.SymbolKind
				if kind == "" {
					kind = info.Role
				}
				abstract := ""
				if info.Abstract != "" {
					abstract = " - " + truncate(info.Abstract, 60)
				}
				fmt.Printf("  - %s (%s)%s\n", info.Title, kind, abstract)
			}

			shown++
			results++
		}
		fmt.Println()
	}

	if results == 0 {
		fmt.Printf("No results found for '%s'\n", *query)
	} else {
		fmt.Printf("Total: %d results\n", results)
	}
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
