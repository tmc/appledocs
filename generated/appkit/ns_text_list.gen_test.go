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
		appkit.TextListMarkerFormat /* not a class type */{}, // markerFormat TextListMarkerFormat /* not a class type */
		0, // options uint
	)
	// Output:
}
// ExampleNewTextListWithMarkerFormatOptionsStartingItemNumber demonstrates how to create a TextList instance using NewTextListWithMarkerFormatOptionsStartingItemNumber.
// Returns a new text list with the format, options, and starting item number you provide.
func ExampleNewTextListWithMarkerFormatOptionsStartingItemNumber() {
	_ = appkit.NewTextListWithMarkerFormatOptionsStartingItemNumber(
		appkit.TextListMarkerFormat /* not a class type */{}, // markerFormat TextListMarkerFormat /* not a class type */
		appkit.TextListOptions /* not a class type */{}, // options TextListOptions /* not a class type */
		0, // startingItemNumber int
	)
	// Output:
}
