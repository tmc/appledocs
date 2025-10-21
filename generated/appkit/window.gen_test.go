// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewWindow

// ExampleNewWindowWithContentViewController demonstrates how to create a Window instance using NewWindowWithContentViewController.
// Creates a titled window that contains the specified content view controller.
func ExampleNewWindowWithContentViewController() {
	_ = appkit.NewWindowWithContentViewController(
		appkit.NSViewController{}, // contentViewController NSViewController
	)
	// Output:
}
