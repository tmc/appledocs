// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewPasteboard

// ExampleNewPasteboardWithName demonstrates how to create a Pasteboard instance using NewPasteboardWithName.
// Returns the pasteboard with the specified name.
func ExampleNewPasteboardWithName() {
	_ = appkit.NewPasteboardWithName(
		appkit.PasteboardName /* typedef */{}, // name PasteboardName /* typedef */
	)
	// Output:
}
// ExamplePasteboard_ClearContents demonstrates using ClearContents on a Pasteboard instance.
// Clears the existing contents of the pasteboard.
func ExamplePasteboard_ClearContents() {
	obj := appkit.NewPasteboard()
	_ = obj.ClearContents()
	// Output:
	}

// ExamplePasteboard_ReadFileWrapper demonstrates using ReadFileWrapper on a Pasteboard instance.
// Reads data representing a file’s contents from the receiver and returns it as a file wrapper.
func ExamplePasteboard_ReadFileWrapper() {
	obj := appkit.NewPasteboard()
	_ = obj.ReadFileWrapper()
	// Output:
	}

// ExamplePasteboard_ReleaseGlobally demonstrates using ReleaseGlobally on a Pasteboard instance.
// Releases the receiver’s resources in the pasteboard server.
func ExamplePasteboard_ReleaseGlobally() {
	obj := appkit.NewPasteboard()
	obj.ReleaseGlobally()
	// Output:
	}

