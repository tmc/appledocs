//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureManualExposureBracketedStillImageSettings


// iOS-only properties

// The exposure duration for the still image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureManualExposureBracketedStillImageSettings/exposureDuration
func (c_ CaptureManualExposureBracketedStillImageSettings) ExposureDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("exposureDuration"))
	return rv
}

// The ISO for the still image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureManualExposureBracketedStillImageSettings/iso
func (c_ CaptureManualExposureBracketedStillImageSettings) ISO() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("ISO"))
	return rv
}





