// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewWindowController

// ExampleNewWindowControllerWithWindow demonstrates how to create a WindowController instance using NewWindowControllerWithWindow.
// Returns a window controller initialized with a given window.
func ExampleNewWindowControllerWithWindow() {
	_ = appkit.NewWindowControllerWithWindow(
		appkit.NSWindow{}, // window NSWindow
	)
	// Output:
}
// ExampleNewWindowControllerWithWindowNibName demonstrates how to create a WindowController instance using NewWindowControllerWithWindowNibName.
// Returns a window controller initialized with a nib file.
func ExampleNewWindowControllerWithWindowNibName() {
	_ = appkit.NewWindowControllerWithWindowNibName(
		appkit.NibName{}, // windowNibName NibName
	)
	// Output:
}
