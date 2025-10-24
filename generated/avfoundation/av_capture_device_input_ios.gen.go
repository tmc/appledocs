//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureDeviceInput


// Retrieves a virtual device’s constituent device ports for use in a multi-camera session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/ports(for:sourceDeviceType:sourceDevicePosition:)
func (c_ CaptureDeviceInput) PortsWithMediaTypeSourceDeviceTypeSourceDevicePosition(mediaType MediaType /* typedef */, sourceDeviceType CaptureDeviceType /* typedef */, sourceDevicePosition CaptureDevicePosition) []CaptureInputPort {
	rv := objc.Send[[]CaptureInputPort](c_.ID, objc.Sel("portsWithMediaType:sourceDeviceType:sourceDevicePosition:"), mediaType, sourceDeviceType, sourceDevicePosition)
	return rv
}

// iOS-only properties

// A Boolean value that indicates whether the input enables unified auto-exposure defaults.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/unifiedAutoExposureDefaultsEnabled
func (c_ CaptureDeviceInput) UnifiedAutoExposureDefaultsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("unifiedAutoExposureDefaultsEnabled"))
	return rv
}
func (c_ CaptureDeviceInput) SetUnifiedAutoExposureDefaultsEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setUnifiedAutoExposureDefaultsEnabled:"), value)
}

// A time value that acts as a modifier to a capture device’s active video minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/videoMinFrameDurationOverride
func (c_ CaptureDeviceInput) VideoMinFrameDurationOverride() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("videoMinFrameDurationOverride"))
	return rv
}
func (c_ CaptureDeviceInput) SetVideoMinFrameDurationOverride(value objc.IObject /* cross-framework: Time */) {
	c_.ID.Send(objc.RegisterName("setVideoMinFrameDurationOverride:"), value)
}




