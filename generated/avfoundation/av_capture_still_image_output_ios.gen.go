//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureStillImageOutput


// iOS-only properties

// A Boolean value that indicates whether still image stabilization should be automatically enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/automaticallyEnablesStillImageStabilizationWhenAvailable
func (c_ CaptureStillImageOutput) AutomaticallyEnablesStillImageStabilizationWhenAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyEnablesStillImageStabilizationWhenAvailable"))
	return rv
}
func (c_ CaptureStillImageOutput) SetAutomaticallyEnablesStillImageStabilizationWhenAvailable(value bool) {
	c_.ID.Send(objc.RegisterName("setAutomaticallyEnablesStillImageStabilizationWhenAvailable:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/isCameraSensorOrientationCompensationEnabled
func (c_ CaptureStillImageOutput) CameraSensorOrientationCompensationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cameraSensorOrientationCompensationEnabled"))
	return rv
}
func (c_ CaptureStillImageOutput) SetCameraSensorOrientationCompensationEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setCameraSensorOrientationCompensationEnabled:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/isCameraSensorOrientationCompensationSupported
func (c_ CaptureStillImageOutput) CameraSensorOrientationCompensationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cameraSensorOrientationCompensationSupported"))
	return rv
}

// A Boolean value that specifies whether to stabilize the lens across the duration of a bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/isLensStabilizationDuringBracketedCaptureEnabled
func (c_ CaptureStillImageOutput) LensStabilizationDuringBracketedCaptureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("lensStabilizationDuringBracketedCaptureEnabled"))
	return rv
}
func (c_ CaptureStillImageOutput) SetLensStabilizationDuringBracketedCaptureEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setLensStabilizationDuringBracketedCaptureEnabled:"), value)
}

// A Boolean value that indicates whether the capture output supports lens stabilization across the duration of a bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/isLensStabilizationDuringBracketedCaptureSupported
func (c_ CaptureStillImageOutput) LensStabilizationDuringBracketedCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("lensStabilizationDuringBracketedCaptureSupported"))
	return rv
}

// Indicates whether still image stabilization is in use for the current capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/isStillImageStabilizationActive
func (c_ CaptureStillImageOutput) StillImageStabilizationActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("stillImageStabilizationActive"))
	return rv
}

// A Boolean value that indicates whether the still image currently being captured supports still image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/isStillImageStabilizationSupported
func (c_ CaptureStillImageOutput) StillImageStabilizationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("stillImageStabilizationSupported"))
	return rv
}

// Specifies the maximum number of still images that may be taken in a single bracket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/maxBracketedCaptureStillImageCount
func (c_ CaptureStillImageOutput) MaxBracketedCaptureStillImageCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxBracketedCaptureStillImageCount"))
	return rv
}




