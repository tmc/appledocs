// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewColorSpace

// ExampleNewColorSpaceWithCGColorSpace demonstrates how to create a ColorSpace instance using NewColorSpaceWithCGColorSpace.
// Initializes and returns a color space object initialized from a Core Graphics color-space object.
func ExampleNewColorSpaceWithCGColorSpace() {
	_ = appkit.NewColorSpaceWithCGColorSpace(
		appkit.ColorSpaceRef /* not a class type */{}, // cgColorSpace ColorSpaceRef /* not a class type */
	)
	// Output:
}
