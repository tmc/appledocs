package main

import (
	"fmt"

	// Import the generated FSKit bindings
	_ "github.com/tmc/appledocs/generated/frameworks/fskit"
)

func main() {
	fmt.Println("FSKit Framework Bindings Demo")
	fmt.Println("==============================")
	fmt.Println()
	fmt.Println("Minimum macOS version: 15.4")
	fmt.Println("Framework loaded successfully!")
	fmt.Println()
	fmt.Println("This example demonstrates that the FSKit framework")
	fmt.Println("can be loaded using purego-based bindings without cgo.")
	fmt.Println()
	fmt.Println("Available functions:")
	fmt.Println("  - fs_errorForCocoaError")
	fmt.Println("  - fs_errorForMachError")
	fmt.Println("  - fs_errorForPOSIXError")
}
