//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureFraming


// iOS-only properties

// An aspect ratio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFraming/aspectRatio
func (c_ CaptureFraming) AspectRatio() CaptureAspectRatio /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("aspectRatio"))
	return rv
}

// A zoom factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFraming/zoomFactor
func (c_ CaptureFraming) ZoomFactor() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("zoomFactor"))
	return rv
}





