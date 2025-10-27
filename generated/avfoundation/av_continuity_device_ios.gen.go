//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ContinuityDevice


// iOS-only properties

// An array of the continuity device’s audio session port descriptions that’s available to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContinuityDevice/audioSessionInputs
func (c_ ContinuityDevice) AudioSessionInputs() []AudioSessionPortDescription /* not a class type */ {
	rv := objc.Send[[]AudioSessionPortDescription](c_.ID, objc.Sel("audioSessionInputs"))
	return rv
}

// A universally unique value that identifies a specific continuity device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContinuityDevice/connectionID
func (c_ ContinuityDevice) ConnectionID() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("connectionID"))
	return rv
}

// A Boolean value that indicates whether you can use the continuity device because it’s connected to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContinuityDevice/isConnected
func (c_ ContinuityDevice) Connected() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("connected"))
	return rv
}

// An array of the continuity device’s video-capture devices available to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContinuityDevice/videoDevices
func (c_ ContinuityDevice) VideoDevices() []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("videoDevices"))
	return rv
}





