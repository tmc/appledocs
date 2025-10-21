// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTouchBar

// ExampleNewTouchBar demonstrates how to create a TouchBar instance.
// Creates a Touch Bar object.
func ExampleNewTouchBar() {
	_ = appkit.NewTouchBar()
	// Output:
}
// ExampleNewTouchBarWithCoder demonstrates how to create a TouchBar instance using NewTouchBarWithCoder.
// Creates a Touch Bar object from a coder object provided by a storyboard or NIB file.
func ExampleNewTouchBarWithCoder() {
	_ = appkit.NewTouchBarWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
