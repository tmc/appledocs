// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)


// ExampleNewSamplerWithImage demonstrates how to create a Sampler instance using NewSamplerWithImage.
// Initializes a sampler with an image object.
func ExampleNewSamplerWithImage() {
	_ = coreimage.NewSamplerWithImage(
		nil, // im unsafe.Pointer
	)
	// Output:
}

// ExampleNewSamplerWithImageOptions demonstrates how to create a Sampler instance using NewSamplerWithImageOptions.
// Initializes the sampler with an image object using options specified in a dictionary.
func ExampleNewSamplerWithImageOptions() {
	_ = coreimage.NewSamplerWithImageOptions(
		nil, // im unsafe.Pointer
		nil, // dict unsafe.Pointer
	)
	// Output:
}

// ExampleNewSamplerWithImageKeysAndValues demonstrates how to create a Sampler instance using NewSamplerWithImageKeysAndValues.
// Initializes the sampler with an image object using options specified as key-value pairs.
func ExampleNewSamplerWithImageKeysAndValues() {
	_ = coreimage.NewSamplerWithImageKeysAndValues(
		nil, // im unsafe.Pointer
		0, // key0 objc.ID
	)
	// Output:
}


