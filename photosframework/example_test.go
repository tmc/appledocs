package photosframework_test

import (
	"fmt"

	"github.com/tmc/appledocs/photosframework"
)

// Example_basic demonstrates basic usage of the photosframework package.
func Example_basic() {
	// Test the library is loaded
	photosframework.TestHello()

	// Get shared photo library
	library := photosframework.SharedPhotoLibrary()
	if library != nil {
		defer library.Release()
		fmt.Println("Got photo library")
	}

	// Output:
	// Hello from Photos Swift wrapper!
	// PHPhotoLibrary is available: PHPhotoLibrary
	// Got photo library
}

// Example_iterator demonstrates using the iterator pattern with persistent changes.
func Example_iterator() {
	// This example shows the API, but doesn't have real data
	var fetchResult *photosframework.PersistentChangeFetchResult

	if fetchResult != nil {
		// Create iterator for changes
		iter := fetchResult.Iterator()
		defer iter.Release()

		// Iterate through all changes
		for {
			change := iter.Next()
			if change == nil {
				break // No more changes
			}
			defer change.Release()

			// Process the change
			fmt.Println("Processing change:", change)
		}
	}
}

// Example_removeAssets demonstrates removing assets from a project.
func Example_removeAssets() {
	// This example shows the API, but doesn't have real data
	var request *photosframework.ProjectChangeRequest
	var assets *photosframework.FetchResult

	if request != nil && assets != nil {
		// Remove assets using Swift extension method
		// This method is not available in the Objective-C API!
		request.RemoveAssets(assets)
		fmt.Println("Assets removed from project")
	}
}

// Example_memoryManagement demonstrates automatic and manual memory management.
func Example_memoryManagement() {
	// Automatic memory management with defer
	library := photosframework.SharedPhotoLibrary()
	if library != nil {
		defer library.Release() // Automatically called when function returns

		// Use the library...
	}

	// Manual memory management
	library2 := photosframework.SharedPhotoLibrary()
	if library2 != nil {
		// Use the library...
		library2.Release() // Release immediately when done
	}

	// Note: Release is safe to call multiple times
	// and will be called by finalizer if you forget
}
