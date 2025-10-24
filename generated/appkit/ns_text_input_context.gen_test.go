// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTextInputContext

// ExampleTextInputContext_Deactivate demonstrates using Deactivate on a TextInputContext instance.
// Deactivates the receiver.
func ExampleTextInputContext_Deactivate() {
	obj := appkit.NewTextInputContext()
	obj.Deactivate()
	// Output:
	}

// ExampleTextInputContext_DiscardMarkedText demonstrates using DiscardMarkedText on a TextInputContext instance.
// Tells the Cocoa text input system to discard the current conversion session.
//
// Note: This example is not executed because DiscardMarkedText crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleTextInputContext_DiscardMarkedText() {
	obj := appkit.NewTextInputContext()
	obj.DiscardMarkedText()
	}

// ExampleTextInputContext_InvalidateCharacterCoordinates demonstrates using InvalidateCharacterCoordinates on a TextInputContext instance.
// Notifies the Cocoa text input system that the position information previously queried via methods like   needs to be updated.
func ExampleTextInputContext_InvalidateCharacterCoordinates() {
	obj := appkit.NewTextInputContext()
	obj.InvalidateCharacterCoordinates()
	// Output:
	}

// ExampleTextInputContext_TextInputClientDidScroll demonstrates using TextInputClientDidScroll on a TextInputContext instance.
func ExampleTextInputContext_TextInputClientDidScroll() {
	obj := appkit.NewTextInputContext()
	obj.TextInputClientDidScroll()
	// Output:
	}

// ExampleTextInputContext_TextInputClientDidUpdateSelection demonstrates using TextInputClientDidUpdateSelection on a TextInputContext instance.
func ExampleTextInputContext_TextInputClientDidUpdateSelection() {
	obj := appkit.NewTextInputContext()
	obj.TextInputClientDidUpdateSelection()
	// Output:
	}

// ExampleTextInputContext_TextInputClientWillStartScrollingOrZooming demonstrates using TextInputClientWillStartScrollingOrZooming on a TextInputContext instance.
func ExampleTextInputContext_TextInputClientWillStartScrollingOrZooming() {
	obj := appkit.NewTextInputContext()
	obj.TextInputClientWillStartScrollingOrZooming()
	// Output:
	}

