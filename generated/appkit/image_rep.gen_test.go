// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewImageRep

// ExampleNewImageRepWithPasteboard demonstrates how to create a ImageRep instance using NewImageRepWithPasteboard.
// Creates and returns an image representation object using the contents of the specified pasteboard.
func ExampleNewImageRepWithPasteboard() {
	_ = appkit.NewImageRepWithPasteboard(
		appkit.NSPasteboard{}, // pasteboard NSPasteboard
	)
	// Output:
}
