// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

// PToolPickerObserver is the PKToolPickerObserver protocol interface.
//
// An interface you use to detect when the user changes the selected tools and drawing characteristics of a tool picker object.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.pencilkit/documentation/PencilKit/PKToolPickerObserver
type PToolPickerObserver interface {
	// Optional methods
	ToolPickerFramesObscuredDidChange(toolPicker IPKToolPicker)
	HasToolPickerFramesObscuredDidChange() bool
	ToolPickerIsRulerActiveDidChange(toolPicker IPKToolPicker)
	HasToolPickerIsRulerActiveDidChange() bool
	ToolPickerSelectedToolDidChange(toolPicker IPKToolPicker)
	HasToolPickerSelectedToolDidChange() bool
	ToolPickerSelectedToolItemDidChange(toolPicker IPKToolPicker)
	HasToolPickerSelectedToolItemDidChange() bool
	ToolPickerVisibilityDidChange(toolPicker IPKToolPicker)
	HasToolPickerVisibilityDidChange() bool
}
