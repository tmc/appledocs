//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureConnection


// iOS-only properties

// The connection’s current stabilization mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/activeVideoStabilizationMode
func (c_ CaptureConnection) ActiveVideoStabilizationMode() CaptureVideoStabilizationMode {
	rv := objc.Send[CaptureVideoStabilizationMode](c_.ID, objc.Sel("activeVideoStabilizationMode"))
	return rv
}

// A Boolean value that indicates whether the system enables video stabilization when it’s available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/enablesVideoStabilizationWhenAvailable
func (c_ CaptureConnection) EnablesVideoStabilizationWhenAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enablesVideoStabilizationWhenAvailable"))
	return rv
}
func (c_ CaptureConnection) SetEnablesVideoStabilizationWhenAvailable(value bool) {
	c_.ID.Send(objc.RegisterName("setEnablesVideoStabilizationWhenAvailable:"), value)
}

// A Boolean value that indicates whether the connection can configure the capture pipeline to deliver camera intrinsics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isCameraIntrinsicMatrixDeliveryEnabled
func (c_ CaptureConnection) CameraIntrinsicMatrixDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cameraIntrinsicMatrixDeliveryEnabled"))
	return rv
}
func (c_ CaptureConnection) SetCameraIntrinsicMatrixDeliveryEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setCameraIntrinsicMatrixDeliveryEnabled:"), value)
}

// A Boolean value that indicates whether the capture connection currently supports delivering camera intrinsics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isCameraIntrinsicMatrixDeliverySupported
func (c_ CaptureConnection) CameraIntrinsicMatrixDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cameraIntrinsicMatrixDeliverySupported"))
	return rv
}

// A Boolean value that indicates whether video stabilization is active for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoStabilizationEnabled
func (c_ CaptureConnection) VideoStabilizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("videoStabilizationEnabled"))
	return rv
}

// A Boolean value that indicates whether this connection supports video stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoStabilizationSupported
func (c_ CaptureConnection) SupportsVideoStabilization() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsVideoStabilization"))
	return rv
}

// The stabilization mode that’s the most appropriate for a video connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/preferredVideoStabilizationMode
func (c_ CaptureConnection) PreferredVideoStabilizationMode() CaptureVideoStabilizationMode {
	rv := objc.Send[CaptureVideoStabilizationMode](c_.ID, objc.Sel("preferredVideoStabilizationMode"))
	return rv
}
func (c_ CaptureConnection) SetPreferredVideoStabilizationMode(value CaptureVideoStabilizationMode) {
	c_.ID.Send(objc.RegisterName("setPreferredVideoStabilizationMode:"), value)
}

// The connection’s maximum video scale and crop factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMaxScaleAndCropFactor
func (c_ CaptureConnection) VideoMaxScaleAndCropFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMaxScaleAndCropFactor"))
	return rv
}

// The current scale and crop factor the video output uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoScaleAndCropFactor
func (c_ CaptureConnection) VideoScaleAndCropFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoScaleAndCropFactor"))
	return rv
}
func (c_ CaptureConnection) SetVideoScaleAndCropFactor(value float64) {
	c_.ID.Send(objc.RegisterName("setVideoScaleAndCropFactor:"), value)
}




