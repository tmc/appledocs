//go:build darwin && ios

// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// iOS-only methods for ContinuityDevicePickerViewController


// iOS-only properties

// The delegate that responds to events from the continuity device picker view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContinuityDevicePickerViewController/delegate
func (c_ ContinuityDevicePickerViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}
func (c_ ContinuityDevicePickerViewController) SetDelegate(value objc.ID) {
	c_.ID.Send(objc.RegisterName("setDelegate:"), value)
}





