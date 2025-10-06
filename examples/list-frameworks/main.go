// List all available frameworks in the Apple documentation.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tmc/appledocs"
)

func main() {
	docsPath := flag.String("docs", "output/tutorials/data/documentation", "Path to documentation directory")
	verbose := flag.Bool("v", false, "Show framework details")
	flag.Parse()

	// Open the documentation
	fsys, err := appledocs.Open(*docsPath)
	if err != nil {
		log.Fatalf("Failed to open docs: %v", err)
	}

	// List all frameworks
	frameworks, err := appledocs.ListFrameworks(fsys)
	if err != nil {
		log.Fatalf("Failed to list frameworks: %v", err)
	}

	fmt.Printf("Found %d frameworks:\n\n", len(frameworks))

	if *verbose {
		// Show details for each framework
		for i, name := range frameworks {
			info, err := appledocs.GetFrameworkInfo(fsys, name)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Warning: couldn't get info for %s: %v\n", name, err)
				fmt.Printf("%4d. %s\n", i+1, name)
				continue
			}

			fmt.Printf("%4d. %s\n", i+1, info.Title)
			if info.Abstract != "" {
				fmt.Printf("      %s\n", truncate(info.Abstract, 80))
			}
			if len(info.Platforms) > 0 {
				fmt.Printf("      Platforms: %d\n", len(info.Platforms))
			}
			fmt.Println()
		}
	} else {
		// Just show names in columns
		for i, name := range frameworks {
			if i > 0 && i%4 == 0 {
				fmt.Println()
			}
			fmt.Printf("  %-20s", name)
		}
		fmt.Println()
	}
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
