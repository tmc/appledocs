// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewColor

// ExampleNewColorWithCGColor demonstrates how to create a Color instance using NewColorWithCGColor.
// Creates a color object using the specified Core Graphics color.
func ExampleNewColorWithCGColor() {
	_ = appkit.NewColorWithCGColor(
		appkit.ColorRef /* not a class type */ {}, // cgColor ColorRef /* not a class type */
	)
	// Output:
}
