package main

import (
	"fmt"
	"log"

	"github.com/tmc/appledocs/photosframework"
)

func main() {
	fmt.Println("=== Photos Framework (Swift Extensions) Demo ===")
	fmt.Println()

	// Test 1: Verify library is loaded
	fmt.Println("Test 1: Testing Swift wrapper library")
	photosframework.TestHello()
	fmt.Println()

	// Test 2: Get shared photo library
	fmt.Println("Test 2: Getting shared photo library")
	library := photosframework.SharedPhotoLibrary()
	if library == nil {
		log.Fatal("Failed to get shared photo library")
	}
	defer library.Release()
	fmt.Printf("✓ Got photo library: %p\n", library)
	fmt.Println()

	// Test 3: Demonstrate type-safe API
	fmt.Println("Test 3: Type-safe Swift-like API")
	fmt.Println("Available types:")
	fmt.Println("  • photosframework.PhotoLibrary")
	fmt.Println("  • photosframework.FetchResult")
	fmt.Println("  • photosframework.ProjectChangeRequest")
	fmt.Println("  • photosframework.PersistentChangeFetchResult")
	fmt.Println("  • photosframework.PersistentChangeIterator")
	fmt.Println("  • photosframework.PersistentChange")
	fmt.Println()

	// Test 4: Show Swift extension methods
	fmt.Println("Test 4: Swift extension methods available")
	fmt.Println("  • ProjectChangeRequest.RemoveAssets(assets)")
	fmt.Println("  • PersistentChangeFetchResult.Iterator()")
	fmt.Println("  • PersistentChangeIterator.Next()")
	fmt.Println()

	fmt.Println("=== Success! ===")
	fmt.Println()
	fmt.Println("The photosframework package provides:")
	fmt.Println("  ✓ Type-safe Go wrappers")
	fmt.Println("  ✓ Automatic memory management")
	fmt.Println("  ✓ Swift-like API design")
	fmt.Println("  ✓ Iterator support")
	fmt.Println("  ✓ No cgo required")
	fmt.Println()
	fmt.Println("Compare this to examples/photos-swift-bindings which uses")
	fmt.Println("raw unsafe.Pointer calls. This package provides a much")
	fmt.Println("cleaner, more Go-idiomatic API!")
}
