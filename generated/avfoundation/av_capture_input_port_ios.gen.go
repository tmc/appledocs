//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureInputPort


// iOS-only properties

// The position of the source device providing input through this port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port/sourceDevicePosition
func (c_ CaptureInputPort) SourceDevicePosition() CaptureDevicePosition {
	rv := objc.Send[CaptureDevicePosition](c_.ID, objc.Sel("sourceDevicePosition"))
	return rv
}

// The device type of the source camera that provides data to the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port/sourceDeviceType
func (c_ CaptureInputPort) SourceDeviceType() CaptureDeviceType /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("sourceDeviceType"))
	return rv
}





