// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewImage

// ExampleNewImage demonstrates how to create a Image instance.
func ExampleNewImage() {
	_ = appkit.NewImage()
	// Output:
}
// ExampleImage_Name demonstrates using Name on a Image instance.
// Returns the name associated with the image, if any.
func ExampleImage_Name() {
	obj := appkit.NewImage()
	_ = obj.Name()
	// Output:
	}

// ExampleImage_Recache demonstrates using Recache on a Image instance.
// Invalidates and frees offscreen caches of all image representations.
func ExampleImage_Recache() {
	obj := appkit.NewImage()
	obj.Recache()
	// Output:
	}

