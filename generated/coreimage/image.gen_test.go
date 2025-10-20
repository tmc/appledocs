// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// Suppress unused import errors
var _ = coreimage.NewImage



















// ExampleNewImageWithCGLayer demonstrates how to create a Image instance using NewImageWithCGLayer.
// Initializes an image object  from the contents supplied by a CGLayer object.
func ExampleNewImageWithCGLayer() {
	_ = coreimage.NewImageWithCGLayer(
		coreimage.CGLayerRef{}, // layer CGLayerRef
	)
	// Output:
}



// ExampleNewImageWithCGImage demonstrates how to create a Image instance using NewImageWithCGImage.
// Initializes an image object with a Quartz 2D image.
func ExampleNewImageWithCGImage() {
	_ = coreimage.NewImageWithCGImage(
		coreimage.CGImageRef{}, // image CGImageRef
	)
	// Output:
}












