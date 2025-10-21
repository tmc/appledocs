// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewPasteboard

// ExampleNewPasteboardByFilteringFile demonstrates how to create a Pasteboard instance using NewPasteboardByFilteringFile.
// Creates a new pasteboard object that supplies the specified file in as many types as possible based on the available filter services.
func ExampleNewPasteboardByFilteringFile() {
	_ = appkit.NewPasteboardByFilteringFile(
		"filename", // filename string
	)
	// Output:
}
// ExampleNewPasteboardByFilteringTypesInPasteboard demonstrates how to create a Pasteboard instance using NewPasteboardByFilteringTypesInPasteboard.
// Creates a new pasteboard object that supplies the specified pasteboard data in as many types as possible based on the available filter services.
func ExampleNewPasteboardByFilteringTypesInPasteboard() {
	_ = appkit.NewPasteboardByFilteringTypesInPasteboard(
		appkit.NSPasteboard{}, // pboard NSPasteboard
	)
	// Output:
}
// ExampleNewPasteboardWithName demonstrates how to create a Pasteboard instance using NewPasteboardWithName.
// Returns the pasteboard with the specified name.
func ExampleNewPasteboardWithName() {
	_ = appkit.NewPasteboardWithName(
		appkit.PasteboardName{}, // name PasteboardName
	)
	// Output:
}
