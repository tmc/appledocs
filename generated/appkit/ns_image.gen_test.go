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
// ExampleNewImageNamed demonstrates how to create a Image instance using NewImageNamed.
// Returns the image object associated with the specified name.
func ExampleNewImageNamed() {
	_ = appkit.NewImageNamed(
		appkit.ImageName /* typedef */{}, // name ImageName /* typedef */
	)
	// Output:
}
// ExampleNewImageWithCGImageSize demonstrates how to create a Image instance using NewImageWithCGImageSize.
// Creates a new image using the contents of the provided image.
func ExampleNewImageWithCGImageSize() {
	_ = appkit.NewImageWithCGImageSize(
		appkit.ImageRef /* not a class type */{}, // cgImage ImageRef /* not a class type */
		appkit.Size /* not a class type */{}, // size Size /* not a class type */
	)
	// Output:
}
// ExampleNewImageWithSize demonstrates how to create a Image instance using NewImageWithSize.
// Initializes and returns an image object with the specified dimensions.
func ExampleNewImageWithSize() {
	_ = appkit.NewImageWithSize(
		appkit.Size /* not a class type */{}, // size Size /* not a class type */
	)
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

