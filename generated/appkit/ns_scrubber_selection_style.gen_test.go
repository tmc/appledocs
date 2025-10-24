// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewScrubberSelectionStyle

// ExampleNewScrubberSelectionStyle demonstrates how to create a ScrubberSelectionStyle instance.
// Initializes a new scrubber selection style.
func ExampleNewScrubberSelectionStyle() {
	_ = appkit.NewScrubberSelectionStyle()
	// Output:
}
// ExampleScrubberSelectionStyle_MakeSelectionView demonstrates using MakeSelectionView on a ScrubberSelectionStyle instance.
// Provides an opportunity to create a customized scrubber selection style.
func ExampleScrubberSelectionStyle_MakeSelectionView() {
	obj := appkit.NewScrubberSelectionStyle()
	_ = obj.MakeSelectionView()
	// Output:
	}

