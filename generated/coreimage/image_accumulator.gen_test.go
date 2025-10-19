// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)


// ExampleNewImageAccumulatorWithExtentFormat demonstrates how to create a ImageAccumulator instance using NewImageAccumulatorWithExtentFormat.
// Initializes an image accumulator with the specified extent and pixel format.
func ExampleNewImageAccumulatorWithExtentFormat() {
	_ = coreimage.NewImageAccumulatorWithExtentFormat(
		nil, // extent unsafe.Pointer
		nil, // format unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageAccumulatorWithExtentFormatColorSpace demonstrates how to create a ImageAccumulator instance using NewImageAccumulatorWithExtentFormatColorSpace.
// Initializes an image accumulator with the specified extent, pixel format, and color space.
func ExampleNewImageAccumulatorWithExtentFormatColorSpace() {
	_ = coreimage.NewImageAccumulatorWithExtentFormatColorSpace(
		nil, // extent unsafe.Pointer
		nil, // format unsafe.Pointer
		nil, // colorSpace unsafe.Pointer
	)
	// Output:
}


