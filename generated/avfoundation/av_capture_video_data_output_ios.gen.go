//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureVideoDataOutput


// iOS-only properties

// A Boolean value that indicates whether the output automatically configures the size of output buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/automaticallyConfiguresOutputBufferDimensions
func (c_ CaptureVideoDataOutput) AutomaticallyConfiguresOutputBufferDimensions() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyConfiguresOutputBufferDimensions"))
	return rv
}
func (c_ CaptureVideoDataOutput) SetAutomaticallyConfiguresOutputBufferDimensions(value bool) {
	c_.ID.Send(objc.RegisterName("setAutomaticallyConfiguresOutputBufferDimensions:"), value)
}

// A Boolean value that indicates whether the output is configured to deliver preview-sized buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/deliversPreviewSizedOutputBuffers
func (c_ CaptureVideoDataOutput) DeliversPreviewSizedOutputBuffers() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("deliversPreviewSizedOutputBuffers"))
	return rv
}
func (c_ CaptureVideoDataOutput) SetDeliversPreviewSizedOutputBuffers(value bool) {
	c_.ID.Send(objc.RegisterName("setDeliversPreviewSizedOutputBuffers:"), value)
}

// The minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/minFrameDuration
func (c_ CaptureVideoDataOutput) MinFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("minFrameDuration"))
	return rv
}
func (c_ CaptureVideoDataOutput) SetMinFrameDuration(value objc.IObject /* cross-framework: Time */) {
	c_.ID.Send(objc.RegisterName("setMinFrameDuration:"), value)
}

// Indicates whether the receiver should prepare the cellular radio for imminent network activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/preparesCellularRadioForNetworkConnection
func (c_ CaptureVideoDataOutput) PreparesCellularRadioForNetworkConnection() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("preparesCellularRadioForNetworkConnection"))
	return rv
}
func (c_ CaptureVideoDataOutput) SetPreparesCellularRadioForNetworkConnection(value bool) {
	c_.ID.Send(objc.RegisterName("setPreparesCellularRadioForNetworkConnection:"), value)
}




