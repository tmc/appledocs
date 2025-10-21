// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewRulerView

// ExampleNewRulerViewWithScrollViewOrientation demonstrates how to create a RulerView instance using NewRulerViewWithScrollViewOrientation.
// Initializes a newly allocated NSRulerView to have   (  or  ) within  .
func ExampleNewRulerViewWithScrollViewOrientation() {
	_ = appkit.NewRulerViewWithScrollViewOrientation(
		appkit.NSScrollView{}, // scrollView NSScrollView
		appkit.RulerOrientation{}, // orientation RulerOrientation
	)
	// Output:
}
