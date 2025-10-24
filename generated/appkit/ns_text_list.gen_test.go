// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTextList

// ExampleNewTextListWithMarkerFormatOptions demonstrates how to create a TextList instance using NewTextListWithMarkerFormatOptions.
// Returns an initialized text list.
func ExampleNewTextListWithMarkerFormatOptions() {
	_ = appkit.NewTextListWithMarkerFormatOptions(
		appkit.TextListMarkerFormat /* typedef */{}, // markerFormat TextListMarkerFormat /* typedef */
		0, // options uint
	)
	// Output:
}
// ExampleNewTextListWithMarkerFormatOptionsStartingItemNumber demonstrates how to create a TextList instance using NewTextListWithMarkerFormatOptionsStartingItemNumber.
// Returns a new text list with the format, options, and starting item number you provide.
func ExampleNewTextListWithMarkerFormatOptionsStartingItemNumber() {
	_ = appkit.NewTextListWithMarkerFormatOptionsStartingItemNumber(
		appkit.TextListMarkerFormat /* typedef */{}, // markerFormat TextListMarkerFormat /* typedef */
		appkit.TextListOptions{}, // options TextListOptions
		0, // startingItemNumber int
	)
	// Output:
}
