// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewBitmapImageRep

// ExampleNewBitmapImageRepForIncrementalLoad demonstrates how to create a BitmapImageRep instance using NewBitmapImageRepForIncrementalLoad.
// Initializes a newly allocated bitmap image representation for incremental loading.
func ExampleNewBitmapImageRepForIncrementalLoad() {
	_ = appkit.NewBitmapImageRepForIncrementalLoad()
	// Output:
}
// ExampleNewBitmapImageRepWithCGImage demonstrates how to create a BitmapImageRep instance using NewBitmapImageRepWithCGImage.
// Returns a bitmap image representation from a Core Graphics image object.
func ExampleNewBitmapImageRepWithCGImage() {
	_ = appkit.NewBitmapImageRepWithCGImage(
		appkit.ImageRef /* not a class type */{}, // cgImage ImageRef /* not a class type */
	)
	// Output:
}
