// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewScrubber

// ExampleNewScrubberWithFrame demonstrates how to create a Scrubber instance using NewScrubberWithFrame.
// Initializes and returns a newly allocated scrubber object with the specified frame rectangle.
func ExampleNewScrubberWithFrame() {
	_ = appkit.NewScrubberWithFrame(
		appkit.Rect /* not a class type */{}, // frameRect Rect /* not a class type */
	)
	// Output:
}
// ExampleScrubber_ReloadData demonstrates using ReloadData on a Scrubber instance.
// Reloads the content of the entire scrubber, and deselects the currently selected item.
func ExampleScrubber_ReloadData() {
	obj := appkit.NewScrubber()
	obj.ReloadData()
	// Output:
	}

