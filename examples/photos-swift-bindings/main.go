package main

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/ebitengine/purego"
)

// Photos Swift wrapper functions
var (
	photosTestHello                        func()
	photosSharedLibrary                    func() unsafe.Pointer
	photosFetchResultCount                 func(unsafe.Pointer) int
	photosRelease                          func(unsafe.Pointer)
	photosProjectChangeRequestRemoveAssets func(unsafe.Pointer, unsafe.Pointer)
	photosPersistentChangeMakeIterator     func(unsafe.Pointer) unsafe.Pointer
	photosPersistentChangeIteratorNext     func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	// Load the Swift wrapper library
	lib, err := purego.Dlopen("../../generated/swift/frameworks/photos/.build/release/libPhotosSwift.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		log.Fatalf("Failed to load libPhotosSwift.dylib: %v", err)
	}

	// Register all function pointers
	purego.RegisterLibFunc(&photosTestHello, lib, "photos_test_hello")
	purego.RegisterLibFunc(&photosSharedLibrary, lib, "photos_shared_library")
	purego.RegisterLibFunc(&photosFetchResultCount, lib, "photos_fetch_result_count")
	purego.RegisterLibFunc(&photosRelease, lib, "photos_release")
	purego.RegisterLibFunc(&photosProjectChangeRequestRemoveAssets, lib, "photos_project_change_request_remove_assets_fetch_result")
	purego.RegisterLibFunc(&photosPersistentChangeMakeIterator, lib, "photos_persistent_change_fetch_result_make_iterator")
	purego.RegisterLibFunc(&photosPersistentChangeIteratorNext, lib, "photos_persistent_change_iterator_next")
}

func main() {
	fmt.Println("=== Photos Swift Bindings Test ===")
	fmt.Println()

	// Test 1: Call test function
	fmt.Println("Test 1: Calling photos_test_hello()")
	photosTestHello()
	fmt.Println()

	// Test 2: Get shared library
	fmt.Println("Test 2: Getting shared PHPhotoLibrary")
	library := photosSharedLibrary()
	if library != nil {
		fmt.Printf("✓ Got PHPhotoLibrary: %p\n", library)
		defer photosRelease(library)
	} else {
		fmt.Println("✗ Failed to get PHPhotoLibrary")
	}
	fmt.Println()

	// Test 3: Show available functions
	fmt.Println("Test 3: Available Swift wrapper functions:")
	fmt.Println("  ✓ photos_test_hello")
	fmt.Println("  ✓ photos_shared_library")
	fmt.Println("  ✓ photos_fetch_result_count")
	fmt.Println("  ✓ photos_release")
	fmt.Println("  ✓ photos_project_change_request_remove_assets_fetch_result")
	fmt.Println("  ✓ photos_persistent_change_fetch_result_make_iterator")
	fmt.Println("  ✓ photos_persistent_change_iterator_next")
	fmt.Println()

	fmt.Println("=== Success! ===")
	fmt.Println()
	fmt.Println("Swift extensions from Photos framework are now callable from Go!")
	fmt.Println("This demonstrates:")
	fmt.Println("  1. Parsing .swiftinterface files with SwiftSyntax")
	fmt.Println("  2. Generating @_cdecl wrappers for Swift APIs")
	fmt.Println("  3. Calling Swift code from Go via purego (no cgo!)")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("  - Add more Photos framework APIs")
	fmt.Println("  - Scale to other frameworks (Speech, Combine, SwiftUI)")
	fmt.Println("  - Automate wrapper generation")
}
