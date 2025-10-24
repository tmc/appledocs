// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewControl

// ExampleNewControlWithFrame demonstrates how to create a Control instance using NewControlWithFrame.
// Initializes a control with the specified frame rectangle.
func ExampleNewControlWithFrame() {
	_ = appkit.NewControlWithFrame(
		appkit.Rect /* not a class type */{}, // frameRect Rect /* not a class type */
	)
	// Output:
}
// ExampleControl_AbortEditing demonstrates using AbortEditing on a Control instance.
// Terminates the current editing operation and discards any edited text.
func ExampleControl_AbortEditing() {
	obj := appkit.NewControl()
	_ = obj.AbortEditing()
	// Output:
	}

// ExampleControl_CurrentEditor demonstrates using CurrentEditor on a Control instance.
// Returns the current field editor for the control.
func ExampleControl_CurrentEditor() {
	obj := appkit.NewControl()
	_ = obj.CurrentEditor()
	// Output:
	}

// ExampleControl_SelectedCell demonstrates using SelectedCell on a Control instance.
// Returns the receiver’s selected cell.
func ExampleControl_SelectedCell() {
	obj := appkit.NewControl()
	_ = obj.SelectedCell()
	// Output:
	}

// ExampleControl_SelectedTag demonstrates using SelectedTag on a Control instance.
// Returns the tag of the receiver’s selected cell.
func ExampleControl_SelectedTag() {
	obj := appkit.NewControl()
	_ = obj.SelectedTag()
	// Output:
	}

// ExampleControl_SizeToFit demonstrates using SizeToFit on a Control instance.
// Resizes the receiver’s frame so that it’s the minimum size needed to contain its cell.
func ExampleControl_SizeToFit() {
	obj := appkit.NewControl()
	obj.SizeToFit()
	// Output:
	}

// ExampleControl_ValidateEditing demonstrates using ValidateEditing on a Control instance.
// Validates changes to any user-typed text.
func ExampleControl_ValidateEditing() {
	obj := appkit.NewControl()
	obj.ValidateEditing()
	// Output:
	}

