// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz_test

import (
	"github.com/tmc/appledocs/generated/quartz"
)

// Suppress unused import errors
var _ = quartz.NewQCPlugInViewController

// ExampleNewQCPlugInViewControllerWithPlugInViewNibName demonstrates how to create a QCPlugInViewController instance using NewQCPlugInViewControllerWithPlugInViewNibName.
// Creates and initializes a controller for the specified   object and nib file.
func ExampleNewQCPlugInViewControllerWithPlugInViewNibName() {
	_ = quartz.NewQCPlugInViewControllerWithPlugInViewNibName(
		quartz.QCPlugIn{}, // plugIn QCPlugIn
		"name", // name string
	)
	// Output:
}
