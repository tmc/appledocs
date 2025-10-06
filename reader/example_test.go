package reader_test

import (
	"fmt"
	"log"

	"github.com/tmc/appledocs/reader"
)

func ExampleOpen() {
	// Open the documentation filesystem
	fsys, err := reader.Open("../output/tutorials/data/documentation")
	if err != nil {
		log.Fatal(err)
	}

	// List available frameworks
	frameworks, err := reader.ListFrameworks(fsys)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d frameworks\n", len(frameworks))
	// Output is variable, so we can't check it exactly
}

func ExampleGetFramework() {
	fsys, err := reader.Open("../output/tutorials/data/documentation")
	if err != nil {
		log.Fatal(err)
	}

	// Get framework documentation
	doc, err := reader.GetFramework(fsys, "Foundation")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc.Metadata.Title)
	// Output: Foundation
}

func ExampleGetFrameworkInfo() {
	fsys, err := reader.Open("../output/tutorials/data/documentation")
	if err != nil {
		log.Fatal(err)
	}

	// Get high-level framework info
	info, err := reader.GetFrameworkInfo(fsys, "Foundation")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Framework: %s\n", info.Name)
	fmt.Printf("Platforms: %d\n", len(info.Platforms))
	// Note: Output will vary based on downloaded docs
}

func ExampleListSymbols() {
	fsys, err := reader.Open("../output/tutorials/data/documentation")
	if err != nil {
		log.Fatal(err)
	}

	// List all symbols in Foundation
	symbols, err := reader.ListSymbols(fsys, "Foundation")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Foundation has %d symbols\n", len(symbols))
	// Output is variable
}

func ExampleGetSymbol() {
	fsys, err := reader.Open("../output/tutorials/data/documentation")
	if err != nil {
		log.Fatal(err)
	}

	// Get a specific symbol
	doc, err := reader.GetSymbol(fsys, "Foundation/NSString")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc.Metadata.Title)
	// Output: NSString
}

func ExampleSearchSymbols() {
	fsys, err := reader.Open("../output/tutorials/data/documentation")
	if err != nil {
		log.Fatal(err)
	}

	// Search for symbols containing "Array"
	matches, err := reader.SearchSymbols(fsys, "Foundation", "Array")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d symbols matching 'Array'\n", len(matches))
	// Output is variable
}
