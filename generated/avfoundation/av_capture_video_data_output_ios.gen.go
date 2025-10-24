//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coremedia"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureVideoDataOutput


// iOS-only properties

// The minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/minFrameDuration
func (c_ CaptureVideoDataOutput) MinFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](c_.ID, objc.Sel("minFrameDuration"))
	return rv
}
func (c_ CaptureVideoDataOutput) SetMinFrameDuration(value objc.IObject /* cross-framework: Time */) {
	c_.ID.Send(objc.RegisterName("setMinFrameDuration:"), value)
}





