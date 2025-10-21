// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewView

// ExampleNewViewWithCoder demonstrates how to create a View instance using NewViewWithCoder.
// Initializes a view using from data in the specified coder object.
func ExampleNewViewWithCoder() {
	_ = appkit.NewViewWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
