// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)


// ExampleNewTouchBarWithCoder demonstrates how to create a TouchBar instance using NewTouchBarWithCoder.
// Creates a Touch Bar object from a coder object provided by a storyboard or NIB file.
func ExampleNewTouchBarWithCoder() {
	_ = appkit.NewTouchBarWithCoder(
		nil, // coder unsafe.Pointer
	)
	// Output:
}

// ExampleNewTouchBar demonstrates how to create a TouchBar instance.
// Creates a Touch Bar object.
func ExampleNewTouchBar() {
	_ = appkit.NewTouchBar()
	// Output:
}


