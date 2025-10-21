// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)

// Suppress unused import errors
var _ = coreimage.NewImage

// ExampleNewImageWithColor demonstrates how to create a Image instance using NewImageWithColor.
// Initializes an image of infinite extent whose entire content is the specified color.
func ExampleNewImageWithColor() {
	_ = coreimage.NewImageWithColor(
		coreimage.CIColor{}, // color CIColor
	)
	// Output:
}
// ExampleNewImageWithImage demonstrates how to create a Image instance using NewImageWithImage.
// Initializes an image object with the specified UIKit image object.
func ExampleNewImageWithImage() {
	_ = coreimage.NewImageWithImage(
		coreimage.Image{}, // image Image
	)
	// Output:
}
