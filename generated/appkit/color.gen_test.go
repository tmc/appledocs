// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewColor

// ExampleNewColorWithPatternImage demonstrates how to create a Color instance using NewColorWithPatternImage.
// Creates a color object that uses the specified image pattern to paint the target area.
func ExampleNewColorWithPatternImage() {
	_ = appkit.NewColorWithPatternImage(
		appkit.Image{}, // image Image
	)
	// Output:
}
