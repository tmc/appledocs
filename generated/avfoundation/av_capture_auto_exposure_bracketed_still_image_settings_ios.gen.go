//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureAutoExposureBracketedStillImageSettings


// iOS-only properties

// The exposure bias for the auto exposure bracketed settings
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAutoExposureBracketedStillImageSettings/exposureTargetBias
func (c_ CaptureAutoExposureBracketedStillImageSettings) ExposureTargetBias() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("exposureTargetBias"))
	return rv
}





