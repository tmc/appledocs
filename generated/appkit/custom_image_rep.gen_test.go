// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewCustomImageRep


// ExampleNewCustomImageRepWithDrawSelectorDelegate demonstrates how to create a CustomImageRep instance using NewCustomImageRepWithDrawSelectorDelegate.
// Returns a representation of an image initialized with the specified delegate information.
func ExampleNewCustomImageRepWithDrawSelectorDelegate() {
	_ = appkit.NewCustomImageRepWithDrawSelectorDelegate(
		0, // selector objc.SEL
		0, // delegate objc.ID
	)
	// Output:
}


