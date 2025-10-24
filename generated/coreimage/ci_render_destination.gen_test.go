// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)

// Suppress unused import errors
var _ = coreimage.NewRenderDestination

// ExampleNewRenderDestinationWithPixelBuffer demonstrates how to create a RenderDestination instance using NewRenderDestinationWithPixelBuffer.
// Creates a render destination based on a Core Video pixel buffer.
func ExampleNewRenderDestinationWithPixelBuffer() {
	_ = coreimage.NewRenderDestinationWithPixelBuffer(
		coreimage.PixelBufferRef /* not a class type */{}, // pixelBuffer PixelBufferRef /* not a class type */
	)
	// Output:
}
