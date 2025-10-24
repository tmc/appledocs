//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CapturePhotoBracketSettings


// iOS-only properties

// An array describing the number of and settings for images to produce in a bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings/bracketedSettings
func (c_ CapturePhotoBracketSettings) BracketedSettings() []CaptureBracketedStillImageSettings {
	rv := objc.Send[[]CaptureBracketedStillImageSettings](c_.ID, objc.Sel("bracketedSettings"))
	return rv
}

// A Boolean value that specifies whether to stabilize the lens for the duration of the bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings/isLensStabilizationEnabled
func (c_ CapturePhotoBracketSettings) LensStabilizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("lensStabilizationEnabled"))
	return rv
}
func (c_ CapturePhotoBracketSettings) SetLensStabilizationEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setLensStabilizationEnabled:"), value)
}




