// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)

// Suppress unused import errors
var _ = coreimage.NewRenderDestination

// ExampleNewRenderDestinationWithIOSurface demonstrates how to create a RenderDestination instance using NewRenderDestinationWithIOSurface.
// Creates a render destination based on an   object.
func ExampleNewRenderDestinationWithIOSurface() {
	_ = coreimage.NewRenderDestinationWithIOSurface(
		coreimage.Surface{}, // surface Surface
	)
	// Output:
}

// ExampleNewRenderDestinationWithPixelBuffer demonstrates how to create a RenderDestination instance using NewRenderDestinationWithPixelBuffer.
// Creates a render destination based on a Core Video pixel buffer.
func ExampleNewRenderDestinationWithPixelBuffer() {
	_ = coreimage.NewRenderDestinationWithPixelBuffer(
		coreimage.PixelBufferRef{}, // pixelBuffer PixelBufferRef
	)
	// Output:
}
