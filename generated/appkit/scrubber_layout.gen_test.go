// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewScrubberLayout

// ExampleNewScrubberLayoutWithCoder demonstrates how to create a ScrubberLayout instance using NewScrubberLayoutWithCoder.
// Initializes and returns a newly allocated scrubber layout object from a storyboard or nib file.
func ExampleNewScrubberLayoutWithCoder() {
	_ = appkit.NewScrubberLayoutWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
