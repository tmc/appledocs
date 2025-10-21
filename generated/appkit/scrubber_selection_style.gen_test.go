// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewScrubberSelectionStyle

// ExampleNewScrubberSelectionStyleWithCoder demonstrates how to create a ScrubberSelectionStyle instance using NewScrubberSelectionStyleWithCoder.
// Initializes a scrubber selection style when included from a nib or Storyboard.
func ExampleNewScrubberSelectionStyleWithCoder() {
	_ = appkit.NewScrubberSelectionStyleWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
