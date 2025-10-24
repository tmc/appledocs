//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ContinuityDevice


// iOS-only properties

// An array of the continuity device’s video-capture devices available to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContinuityDevice/videoDevices
func (c_ ContinuityDevice) VideoDevices() []ICaptureDevice {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("videoDevices"))
	return rv
}





