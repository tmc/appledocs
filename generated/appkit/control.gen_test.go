// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewControl

// ExampleNewControlWithCoder demonstrates how to create a Control instance using NewControlWithCoder.
// Initializes a control with data in an unarchiver.
func ExampleNewControlWithCoder() {
	_ = appkit.NewControlWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
