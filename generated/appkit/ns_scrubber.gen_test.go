// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewScrubber

// ExampleScrubber_ReloadData demonstrates using ReloadData on a Scrubber instance.
// Reloads the content of the entire scrubber, and deselects the currently selected item.
func ExampleScrubber_ReloadData() {
	obj := appkit.NewScrubber()
	obj.ReloadData()
	// Output:
	}

