// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PToolPickerDelegate is the PKToolPickerDelegate protocol interface.
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - visionOS 2.0+
//
// See: doc://com.apple.pencilkit/documentation/PencilKit/PKToolPicker/Delegate-swift.protocol
type PToolPickerDelegate interface {
	// Optional methods
	ToolPickerWillDismiss(toolPicker IPKToolPicker) bool
	HasToolPickerWillDismiss() bool
}

// ToolPickerDelegate is a delegate implementation builder for the PToolPickerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ToolPickerDelegate struct {
	_ToolPickerWillDismiss func(toolPicker IPKToolPicker) bool
}

// SetToolPickerWillDismiss sets the handler for the ToolPickerWillDismiss delegate method.
//
// This is called when the user dismisses the tool picker using a built-in control.   This is   called when the tool picker hides from a responder change or other programatic request.   By default, using the dismissal control on the tool picker causes the tool picker to resign the first responder.   The delegate may override that default behavior, taking responsibility for the dismissal of the picker, by returning true from this method.
func (d *ToolPickerDelegate) SetToolPickerWillDismiss(f func(toolPicker IPKToolPicker) bool) {
	d._ToolPickerWillDismiss = f
}

// ToolPickerWillDismiss implements the PToolPickerDelegate interface.
func (d *ToolPickerDelegate) ToolPickerWillDismiss(toolPicker IPKToolPicker) bool {
	if d._ToolPickerWillDismiss != nil {
		return d._ToolPickerWillDismiss(toolPicker)
	}
	var zero bool
	return zero
}

// HasToolPickerWillDismiss returns true if a handler for ToolPickerWillDismiss has been set.
func (d *ToolPickerDelegate) HasToolPickerWillDismiss() bool {
	return d._ToolPickerWillDismiss != nil
}
