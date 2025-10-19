// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVCaptureAutoExposureBracketedStillImageSettings] class.
var aVCaptureAutoExposureBracketedStillImageSettingsClass = _AVCaptureAutoExposureBracketedStillImageSettingsClass{objc.GetClass("AVCaptureAutoExposureBracketedStillImageSettings")}

type _AVCaptureAutoExposureBracketedStillImageSettingsClass struct {
	class objc.Class
}

// A configuration for defining bracketed photo captures in terms of bias relative to automatic exposure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAutoExposureBracketedStillImageSettings

type AVCaptureAutoExposureBracketedStillImageSettings struct {
	AVCaptureBracketedStillImageSettings
}

// AVCaptureAutoExposureBracketedStillImageSettingsFrom constructs a [AVCaptureAutoExposureBracketedStillImageSettings] from an unsafe.Pointer.
//
// A configuration for defining bracketed photo captures in terms of bias relative to automatic exposure.
func AVCaptureAutoExposureBracketedStillImageSettingsFrom(ptr unsafe.Pointer) AVCaptureAutoExposureBracketedStillImageSettings {
	return AVCaptureAutoExposureBracketedStillImageSettings{
		AVCaptureBracketedStillImageSettings: AVCaptureBracketedStillImageSettingsFrom(ptr),
	}
}



