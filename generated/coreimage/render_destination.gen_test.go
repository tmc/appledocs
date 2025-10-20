// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)

// Suppress unused import errors
var _ = coreimage.NewRenderDestination


// ExampleNewRenderDestinationWithMTLTextureCommandBuffer demonstrates how to create a RenderDestination instance using NewRenderDestinationWithMTLTextureCommandBuffer.
// Creates a render destination based on a Metal texture.
func ExampleNewRenderDestinationWithMTLTextureCommandBuffer() {
	_ = coreimage.NewRenderDestinationWithMTLTextureCommandBuffer(
		0, // texture objc.ID
		0, // commandBuffer objc.ID
	)
	// Output:
}







