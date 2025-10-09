// Analyze platform availability across a framework.
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

type platformStats struct {
	name       string
	total      int
	deprecated int
	beta       int
}

func main() {
	docsPath := flag.String("docs", "output/tutorials/data/documentation", "Path to documentation directory")
	framework := flag.String("framework", "Foundation", "Framework to analyze")
	showDeprecated := flag.Bool("deprecated", false, "Show deprecated symbols")
	flag.Parse()

	// Open the documentation
	fsys, err := appledocs.Open(*docsPath)
	if err != nil {
		log.Fatalf("Failed to open docs: %v", err)
	}

	// Load framework
	doc, err := appledocs.GetFramework(fsys, *framework)
	if err != nil {
		log.Fatalf("Failed to load framework: %v", err)
	}

	fmt.Printf("Platform Analysis: %s\n\n", doc.Metadata.Title)

	// Framework-level platform info
	if len(doc.Metadata.Platforms) > 0 {
		fmt.Println("Framework available on:")
		for _, p := range doc.Metadata.Platforms {
			status := ""
			if p.Beta {
				status = " (beta)"
			} else if p.Deprecated {
				status = fmt.Sprintf(" (deprecated in %s)", p.DeprecatedAt)
			}
			fmt.Printf("  - %s: since %s%s\n", p.Name, p.IntroducedAt, status)
		}
		fmt.Println()
	}

	// Analyze symbols
	platformCounts := make(map[string]*platformStats)
	deprecatedSymbols := []string{}

	// Count symbols per platform
	for id, ref := range doc.References {
		if ref.Role != "symbol" {
			continue
		}

		// Load full symbol info to get platform details
		symbolDoc, err := appledocs.GetSymbol(fsys, *framework+"/"+id)
		if err != nil {
			continue
		}

		// Track deprecated symbols
		for _, p := range symbolDoc.Metadata.Platforms {
			if p.Deprecated {
				deprecatedSymbols = append(deprecatedSymbols, ref.Title)
				break
			}
		}

		// Count by platform
		for _, p := range symbolDoc.Metadata.Platforms {
			stats, ok := platformCounts[p.Name]
			if !ok {
				stats = &platformStats{name: p.Name}
				platformCounts[p.Name] = stats
			}
			stats.total++
			if p.Deprecated {
				stats.deprecated++
			}
			if p.Beta {
				stats.beta++
			}
		}
	}

	// Sort platforms by name
	var platforms []string
	for name := range platformCounts {
		platforms = append(platforms, name)
	}
	sort.Strings(platforms)

	// Display statistics
	fmt.Println("Symbol Availability by Platform:")
	fmt.Printf("%-20s %8s %12s %8s\n", "Platform", "Total", "Deprecated", "Beta")
	fmt.Println(strings.Repeat("-", 52))
	for _, name := range platforms {
		stats := platformCounts[name]
		fmt.Printf("%-20s %8d %12d %8d\n",
			stats.name, stats.total, stats.deprecated, stats.beta)
	}

	// Show deprecated symbols if requested
	if *showDeprecated && len(deprecatedSymbols) > 0 {
		fmt.Printf("\nDeprecated Symbols (%d):\n", len(deprecatedSymbols))
		sort.Strings(deprecatedSymbols)
		for _, name := range deprecatedSymbols {
			fmt.Printf("  - %s\n", name)
		}
	}
}
