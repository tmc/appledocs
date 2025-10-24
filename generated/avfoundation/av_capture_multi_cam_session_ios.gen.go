//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for CaptureMultiCamSession


// iOS-only properties

// A value that indicates the percentage of the session’s available hardware budget currently in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession/hardwareCost
func (c_ CaptureMultiCamSession) HardwareCost() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("hardwareCost"))
	return rv
}

// A value that indicates the system pressure cost of the current session configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession/systemPressureCost
func (c_ CaptureMultiCamSession) SystemPressureCost() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("systemPressureCost"))
	return rv
}





