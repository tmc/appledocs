// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)


// ExampleNewFilterWithNameWithInputParameters demonstrates how to create a Filter instance using NewFilterWithNameWithInputParameters.
// Creates a   object for a specific kind of filter and initializes the input values.
func ExampleNewFilterWithNameWithInputParameters() {
	_ = coreimage.NewFilterWithNameWithInputParameters(
		"name", // name string
		nil, // params unsafe.Pointer
	)
	// Output:
}

// ExampleNewFilterWithCVPixelBufferPropertiesOptions demonstrates how to create a Filter instance using NewFilterWithCVPixelBufferPropertiesOptions.
// Creates a filter from a Core Video pixel buffer.
func ExampleNewFilterWithCVPixelBufferPropertiesOptions() {
	_ = coreimage.NewFilterWithCVPixelBufferPropertiesOptions(
		nil, // pixelBuffer unsafe.Pointer
		nil, // properties unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewFilterWithImageDataOptions demonstrates how to create a Filter instance using NewFilterWithImageDataOptions.
// Creates a filter that allows the processing of RAW images.
func ExampleNewFilterWithImageDataOptions() {
	_ = coreimage.NewFilterWithImageDataOptions(
		nil, // data unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewFilterWithImageURLOptions demonstrates how to create a Filter instance using NewFilterWithImageURLOptions.
// Creates a filter that allows the processing of RAW images.
func ExampleNewFilterWithImageURLOptions() {
	_ = coreimage.NewFilterWithImageURLOptions(
		nil, // url unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewFilterWithName demonstrates how to create a Filter instance using NewFilterWithName.
// Creates a   object for a specific kind of filter.
func ExampleNewFilterWithName() {
	_ = coreimage.NewFilterWithName(
		"name", // name string
	)
	// Output:
}


