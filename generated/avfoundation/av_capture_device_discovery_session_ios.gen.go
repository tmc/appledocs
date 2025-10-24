//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureDeviceDiscoverySession


// iOS-only properties

// Sets of capture devices that you can use simultaneously in a multi-camera session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession/supportedMultiCamDeviceSets
func (c_ CaptureDeviceDiscoverySession) SupportedMultiCamDeviceSets() []objc.IObject /* cross-framework: Set */ {
	rv := objc.Send[[]foundation.Set](c_.ID, objc.Sel("supportedMultiCamDeviceSets"))
	return rv
}




