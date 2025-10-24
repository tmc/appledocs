// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewWindowController

// ExampleWindowController_Close demonstrates using Close on a WindowController instance.
// Closes the window if it was loaded.
func ExampleWindowController_Close() {
	obj := appkit.NewWindowController()
	obj.Close()
	// Output:
	}

// ExampleWindowController_LoadWindow demonstrates using LoadWindow on a WindowController instance.
// Loads the receiver’s window from the nib file.
func ExampleWindowController_LoadWindow() {
	obj := appkit.NewWindowController()
	obj.LoadWindow()
	// Output:
	}

// ExampleWindowController_SynchronizeWindowTitleWithDocumentName demonstrates using SynchronizeWindowTitleWithDocumentName on a WindowController instance.
// Synchronizes the displayed window title and the represented filename with the information in the associated document.
func ExampleWindowController_SynchronizeWindowTitleWithDocumentName() {
	obj := appkit.NewWindowController()
	obj.SynchronizeWindowTitleWithDocumentName()
	// Output:
	}

// ExampleWindowController_WindowDidLoad demonstrates using WindowDidLoad on a WindowController instance.
// Sent after the window owned by the receiver has been loaded.
func ExampleWindowController_WindowDidLoad() {
	obj := appkit.NewWindowController()
	obj.WindowDidLoad()
	// Output:
	}

// ExampleWindowController_WindowWillLoad demonstrates using WindowWillLoad on a WindowController instance.
// Sent before the window owned by the receiver is loaded.
func ExampleWindowController_WindowWillLoad() {
	obj := appkit.NewWindowController()
	obj.WindowWillLoad()
	// Output:
	}

