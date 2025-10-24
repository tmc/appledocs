//go:build darwin && ios

// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ResponderState


// iOS-only properties

// The current tool picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKResponderState/activeToolPicker
func (r_ ResponderState) ActiveToolPicker() IPKToolPicker {
	rv := objc.Send[ToolPicker](r_.ID, objc.Sel("activeToolPicker"))
	return rv
}
func (r_ ResponderState) SetActiveToolPicker(value IPKToolPicker) {
	r_.ID.Send(objc.RegisterName("setActiveToolPicker:"), value)
}

// The visibility state of the tool picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKResponderState/toolPickerVisibility-7hikj
func (r_ ResponderState) ToolPickerVisibility() ToolPickerVisibility {
	rv := objc.Send[ToolPickerVisibility](r_.ID, objc.Sel("toolPickerVisibility"))
	return rv
}
func (r_ ResponderState) SetToolPickerVisibility(value ToolPickerVisibility) {
	r_.ID.Send(objc.RegisterName("setToolPickerVisibility:"), value)
}





