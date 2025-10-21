// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewCIImageRep

// ExampleNewCIImageRepWithCIImage demonstrates how to create a CIImageRep instance using NewCIImageRepWithCIImage.
// Returns a representation of an image initialized to the specified Core Image instance.
func ExampleNewCIImageRepWithCIImage() {
	_ = appkit.NewCIImageRepWithCIImage(
		appkit.Image{}, // image Image
	)
	// Output:
}
