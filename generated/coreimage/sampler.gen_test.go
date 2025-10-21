// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)

// Suppress unused import errors
var _ = coreimage.NewSampler

// ExampleNewSamplerWithImage demonstrates how to create a Sampler instance using NewSamplerWithImage.
// Initializes a sampler with an image object.
func ExampleNewSamplerWithImage() {
	_ = coreimage.NewSamplerWithImage(
		coreimage.CIImage{}, // im CIImage
	)
	// Output:
}
