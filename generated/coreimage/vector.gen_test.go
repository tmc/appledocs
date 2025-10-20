// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)

// Suppress unused import errors
var _ = coreimage.NewVector


// ExampleNewVectorWithXYZ demonstrates how to create a Vector instance using NewVectorWithXYZ.
// Initialize a Core Image vector object with three values.
func ExampleNewVectorWithXYZ() {
	_ = coreimage.NewVectorWithXYZ(
		0.0, // x float64
		0.0, // y float64
		0.0, // z float64
	)
	// Output:
}


// ExampleNewVectorWithX demonstrates how to create a Vector instance using NewVectorWithX.
// Initialize a Core Image vector object with one value.
func ExampleNewVectorWithX() {
	_ = coreimage.NewVectorWithX(
		0.0, // x float64
	)
	// Output:
}

// ExampleNewVectorWithXY demonstrates how to create a Vector instance using NewVectorWithXY.
// Initialize a Core Image vector object with two values.
func ExampleNewVectorWithXY() {
	_ = coreimage.NewVectorWithXY(
		0.0, // x float64
		0.0, // y float64
	)
	// Output:
}

// ExampleNewVectorWithXYZW demonstrates how to create a Vector instance using NewVectorWithXYZW.
// Initialize a Core Image vector object with four values.
func ExampleNewVectorWithXYZW() {
	_ = coreimage.NewVectorWithXYZW(
		0.0, // x float64
		0.0, // y float64
		0.0, // z float64
		0.0, // w float64
	)
	// Output:
}



// ExampleNewVectorWithString demonstrates how to create a Vector instance using NewVectorWithString.
// Initialize a Core Image vector object with values provided in a string representation.
func ExampleNewVectorWithString() {
	_ = coreimage.NewVectorWithString(
		"representation", // representation string
	)
	// Output:
}



