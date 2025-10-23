// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewViewController

// ExampleNewViewControllerWithCoder demonstrates how to create a ViewController instance using NewViewControllerWithCoder.
func ExampleNewViewControllerWithCoder() {
	_ = appkit.NewViewControllerWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
// ExampleNewViewControllerWithNibNameBundle demonstrates how to create a ViewController instance using NewViewControllerWithNibNameBundle.
// Returns a view controller object initialized to the nib file in the specified bundle.
func ExampleNewViewControllerWithNibNameBundle() {
	_ = appkit.NewViewControllerWithNibNameBundle(
		appkit.NibName{}, // nibNameOrNil NibName
		appkit.Bundle{}, // nibBundleOrNil Bundle
	)
	// Output:
}
