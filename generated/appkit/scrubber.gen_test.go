// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewScrubber

// ExampleNewScrubberWithCoder demonstrates how to create a Scrubber instance using NewScrubberWithCoder.
// Initializes and returns a newly allocated scrubber object from a storyboard or nib file.
func ExampleNewScrubberWithCoder() {
	_ = appkit.NewScrubberWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
