// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore_test

import (
	"github.com/tmc/appledocs/generated/quartzcore"
)

// Suppress unused import errors
var _ = quartzcore.NewRenderer



// ExampleNewRendererWithMTLTextureOptions demonstrates how to create a Renderer instance using NewRendererWithMTLTextureOptions.
// Creates a layer renderer from a Metal texture.
func ExampleNewRendererWithMTLTextureOptions() {
	_ = quartzcore.NewRendererWithMTLTextureOptions(
		0, // tex objc.ID
		0, // dict objc.ID
	)
	// Output:
}


