// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)

// Suppress unused import errors
var _ = coreimage.NewColor



// ExampleNewColorWithRedGreenBlue demonstrates how to create a Color instance using NewColorWithRedGreenBlue.
// Initialize a Core Image color object in the sRGB color space   with the specified red, green, and blue component values.
func ExampleNewColorWithRedGreenBlue() {
	_ = coreimage.NewColorWithRedGreenBlue(
		0.0, // red float64
		0.0, // green float64
		0.0, // blue float64
	)
	// Output:
}

// ExampleNewColorWithRedGreenBlueAlpha demonstrates how to create a Color instance using NewColorWithRedGreenBlueAlpha.
// Initialize a Core Image color object in the sRGB color space   with the specified red, green, blue, and alpha component values.
func ExampleNewColorWithRedGreenBlueAlpha() {
	_ = coreimage.NewColorWithRedGreenBlueAlpha(
		0.0, // red float64
		0.0, // green float64
		0.0, // blue float64
		0.0, // alpha float64
	)
	// Output:
}



// ExampleNewColorWithString demonstrates how to create a Color instance using NewColorWithString.
// Create a Core Image color object in the sRGB color space using a string containing the RGBA color component values.
func ExampleNewColorWithString() {
	_ = coreimage.NewColorWithString(
		"representation", // representation string
	)
	// Output:
}



