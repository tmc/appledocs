// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)


// ExampleNewFilterWithCVPixelBufferProperties demonstrates how to create a RAWFilter instance using NewFilterWithCVPixelBufferProperties.
// Creates a RAW filter from the pixel buffer and its properties that you specify.
func ExampleNewFilterWithCVPixelBufferProperties() {
	_ = coreimage.NewFilterWithCVPixelBufferProperties(
		nil, // buffer unsafe.Pointer
		nil, // properties unsafe.Pointer
	)
	// Output:
}

// ExampleNewFilterWithImageDataIdentifierHint demonstrates how to create a RAWFilter instance using NewFilterWithImageDataIdentifierHint.
// Creates a RAW filter from the image data and type hint that you specify.
func ExampleNewFilterWithImageDataIdentifierHint() {
	_ = coreimage.NewFilterWithImageDataIdentifierHint(
		nil, // data unsafe.Pointer
		"identifierHint", // identifierHint string
	)
	// Output:
}

// ExampleNewFilterWithImageURL demonstrates how to create a RAWFilter instance using NewFilterWithImageURL.
// Creates a RAW filter from the image at the URL location that you specify.
func ExampleNewFilterWithImageURL() {
	_ = coreimage.NewFilterWithImageURL(
		nil, // url unsafe.Pointer
	)
	// Output:
}


