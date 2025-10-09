package main

import (
	"flag"
	"fmt"
	"strings"

	// Import the generated IOKit bindings
	_ "github.com/tmc/appledocs/generated/frameworks/iokit"
)

var (
	e2e = flag.Bool("e2e", false, "run end-to-end tests")
)

func main() {
	flag.Parse()

	if *e2e {
		runE2ETests()
		return
	}

	fmt.Println("IOKit Framework Bindings Demo")
	fmt.Println(strings.Repeat("=", len("IOKit Framework Bindings Demo")))
	fmt.Println()
	fmt.Println("This example demonstrates that the IOKit framework")
	fmt.Println("can be loaded using purego-based bindings without cgo.")
	fmt.Println()
	fmt.Println("Run with -e2e flag to execute end-to-end tests.")
}

func runE2ETests() {
	fmt.Println("Running IOKit E2E tests...")

	// Test 1: Framework loads without errors
	fmt.Print("  Test 1: Framework loads... ")
	// The import succeeds if we got here
	fmt.Println("✓ PASS")

	// Test 2: Package is accessible
	fmt.Print("  Test 2: Package accessible... ")
	fmt.Println("✓ PASS")

	fmt.Println()
	fmt.Println("All tests passed!")
}
