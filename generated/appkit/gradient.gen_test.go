// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewGradient

// ExampleNewGradientWithColors demonstrates how to create a Gradient instance using NewGradientWithColors.
// Initializes a newly allocated gradient object with an array of colors.
func ExampleNewGradientWithColors() {
	_ = appkit.NewGradientWithColors(
		[]appkit.Color{}, // colorArray []Color
	)
	// Output:
}

// ExampleNewGradientWithColorsAndLocations demonstrates how to create a Gradient instance using NewGradientWithColorsAndLocations.
// Initializes a newly allocated gradient object with a comma-separated list of arguments.
func ExampleNewGradientWithColorsAndLocations() {
	_ = appkit.NewGradientWithColorsAndLocations(
		appkit.NSColor{}, // firstColor NSColor
	)
	// Output:
}

// ExampleNewGradientWithStartingColorEndingColor demonstrates how to create a Gradient instance using NewGradientWithStartingColorEndingColor.
// Initializes a newly allocated gradient object with two colors.
func ExampleNewGradientWithStartingColorEndingColor() {
	_ = appkit.NewGradientWithStartingColorEndingColor(
		appkit.NSColor{}, // startingColor NSColor
		appkit.NSColor{}, // endingColor NSColor
	)
	// Output:
}
