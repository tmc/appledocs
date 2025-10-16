package appledocs_test

import (
	"fmt"
	"log"

	"github.com/tmc/appledocs"
)

func Example() {
	// Open the docs directory
	fsys, err := appledocs.Open("output/tutorials/data/documentation")
	if err != nil {
		log.Fatal(err)
	}

	// Load Foundation framework
	foundation, _ := appledocs.LoadMap(fsys, "Foundation.json")
	fmt.Println(appledocs.Title(foundation))
	// Output: Foundation
}

func Example_class() {
	fsys, _ := appledocs.Open("output/tutorials/data/documentation")

	// Load NSString class
	nsstring, _ := appledocs.LoadMap(fsys, "Foundation/NSString.json")

	fmt.Println("Title:", appledocs.Title(nsstring))
	fmt.Println("Kind:", appledocs.SymbolKind(nsstring))
	fmt.Println("External ID:", appledocs.ExternalID(nsstring))
	fmt.Println("Platforms:", len(appledocs.Platforms(nsstring)))
	// Output:
	// Title: NSString
	// Kind: class
	// External ID: c:objc(cs)NSString
	// Platforms: 7
}

func Example_dictionaryAccess() {
	fsys, _ := appledocs.Open("output/tutorials/data/documentation")

	doc, _ := appledocs.LoadMap(fsys, "Foundation/NSString.json")

	// Access references directly as a map
	refs := appledocs.References(doc)
	count := 0
	for _, ref := range refs {
		refMap := ref.(map[string]interface{})
		if refMap["role"] == "symbol" {
			count++
		}
	}
	fmt.Printf("Found %d symbol references\n", count)
	// Output varies based on docs
}

func Example_darwinkit() {
	// Example: How DarwinKit might use this package

	fsys, _ := appledocs.Open("output/tutorials/data/documentation")

	// Get Foundation framework info
	framework, _ := appledocs.LoadMap(fsys, "Foundation.json")
	frameworkName := appledocs.Title(framework)

	// Get NSString class
	class, _ := appledocs.LoadMap(fsys, appledocs.SymbolPath(frameworkName, "NSString"))

	// Extract info for code generation
	className := appledocs.Title(class)
	externalID := appledocs.ExternalID(class)
	platforms := appledocs.Platforms(class)

	fmt.Printf("Generating bindings for %s\n", className)
	fmt.Printf("  External ID: %s\n", externalID)
	fmt.Printf("  Platforms: %d\n", len(platforms))

	// Check references for methods/properties
	refs := appledocs.References(class)
	methodCount := 0
	for _, ref := range refs {
		refMap := ref.(map[string]interface{})
		if symbolKind, ok := refMap["fragments"]; ok && symbolKind != nil {
			// Has fragments = likely a method/property
			methodCount++
		}
	}

	fmt.Printf("  Methods/Properties: ~%d\n", methodCount)
	// Output varies
}
