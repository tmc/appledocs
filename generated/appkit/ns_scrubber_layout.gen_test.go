// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewScrubberLayout

// ExampleNewScrubberLayout demonstrates how to create a ScrubberLayout instance.
// Initializes and returns a newly allocated scrubber layout object from code.
func ExampleNewScrubberLayout() {
	_ = appkit.NewScrubberLayout()
	// Output:
}
// ExampleScrubberLayout_InvalidateLayout demonstrates using InvalidateLayout on a ScrubberLayout instance.
// Signals that the layout has been invalidated, and that the scrubber control should perform a new layout pass.
func ExampleScrubberLayout_InvalidateLayout() {
	obj := appkit.NewScrubberLayout()
	obj.InvalidateLayout()
	// Output:
	}

// ExampleScrubberLayout_PrepareLayout demonstrates using PrepareLayout on a ScrubberLayout instance.
// Gives you an opportunity to perform layout calculations when the scrubber’s layout is invalidated.
//
// Note: This example is not executed because PrepareLayout crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleScrubberLayout_PrepareLayout() {
	obj := appkit.NewScrubberLayout()
	obj.PrepareLayout()
	}

