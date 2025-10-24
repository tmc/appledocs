// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewStackView

// ExampleNewStackViewWithViews demonstrates how to create a StackView instance using NewStackViewWithViews.
// Creates and returns a stack view with a specified array of views.
func ExampleNewStackViewWithViews() {
	_ = appkit.NewStackViewWithViews(
		[]appkit.IView{}, // views []IView
	)
	// Output:
}
