// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewResponder

// ExampleNewResponder demonstrates how to create a Responder instance.
// Creates a new responder object.
func ExampleNewResponder() {
	_ = appkit.NewResponder()
	// Output:
}
// ExampleResponder_BecomeFirstResponder demonstrates using BecomeFirstResponder on a Responder instance.
// Notifies the receiver that it’s about to become first responder in its  .
func ExampleResponder_BecomeFirstResponder() {
	obj := appkit.NewResponder()
	_ = obj.BecomeFirstResponder()
	// Output:
	}

// ExampleResponder_FlushBufferedKeyEvents demonstrates using FlushBufferedKeyEvents on a Responder instance.
// Clears any unprocessed key events when overridden by subclasses.
func ExampleResponder_FlushBufferedKeyEvents() {
	obj := appkit.NewResponder()
	obj.FlushBufferedKeyEvents()
	// Output:
	}

// ExampleResponder_InvalidateRestorableState demonstrates using InvalidateRestorableState on a Responder instance.
// Marks the responder’s interface-related state as dirty.
func ExampleResponder_InvalidateRestorableState() {
	obj := appkit.NewResponder()
	obj.InvalidateRestorableState()
	// Output:
	}

// ExampleResponder_MakeTouchBar demonstrates using MakeTouchBar on a Responder instance.
// Your custom subclass of the   class should override this method to create and configure your subclass’s default   object.
func ExampleResponder_MakeTouchBar() {
	obj := appkit.NewResponder()
	_ = obj.MakeTouchBar()
	// Output:
	}

// ExampleResponder_ResignFirstResponder demonstrates using ResignFirstResponder on a Responder instance.
// Notifies the receiver that it’s been asked to relinquish its status as first responder in its window.
func ExampleResponder_ResignFirstResponder() {
	obj := appkit.NewResponder()
	_ = obj.ResignFirstResponder()
	// Output:
	}

