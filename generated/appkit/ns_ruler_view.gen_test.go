// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewRulerView

// ExampleRulerView_InvalidateHashMarks demonstrates using InvalidateHashMarks on a RulerView instance.
// Forces recalculation of the hash mark spacing for the next time the receiver is displayed.
func ExampleRulerView_InvalidateHashMarks() {
	obj := appkit.NewRulerView()
	obj.InvalidateHashMarks()
	// Output:
	}

