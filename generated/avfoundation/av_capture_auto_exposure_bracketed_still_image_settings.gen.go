// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVCaptureAutoExposureBracketedStillImageSettings] class.
var (
	aVCaptureAutoExposureBracketedStillImageSettingsClass     _AVCaptureAutoExposureBracketedStillImageSettingsClass
	aVCaptureAutoExposureBracketedStillImageSettingsClassOnce sync.Once
)

func getAVCaptureAutoExposureBracketedStillImageSettingsClass() _AVCaptureAutoExposureBracketedStillImageSettingsClass {
	aVCaptureAutoExposureBracketedStillImageSettingsClassOnce.Do(func() {
		aVCaptureAutoExposureBracketedStillImageSettingsClass = _AVCaptureAutoExposureBracketedStillImageSettingsClass{objc.GetClass("AVCaptureAutoExposureBracketedStillImageSettings")}
	})
	return aVCaptureAutoExposureBracketedStillImageSettingsClass
}

type _AVCaptureAutoExposureBracketedStillImageSettingsClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureAutoExposureBracketedStillImageSettings] class.
type IAVCaptureAutoExposureBracketedStillImageSettings interface {
	IAVCaptureBracketedStillImageSettings
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

// Alloc allocates a new instance without initialization.
func (ac _AVCaptureAutoExposureBracketedStillImageSettingsClass) Alloc() AVCaptureAutoExposureBracketedStillImageSettings {
	rv := objc.Send[AVCaptureAutoExposureBracketedStillImageSettings](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVCaptureAutoExposureBracketedStillImageSettingsClass) New() AVCaptureAutoExposureBracketedStillImageSettings {
	rv := objc.Send[AVCaptureAutoExposureBracketedStillImageSettings](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureAutoExposureBracketedStillImageSettings) Init() AVCaptureAutoExposureBracketedStillImageSettings {
	rv := objc.Send[AVCaptureAutoExposureBracketedStillImageSettings](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureAutoExposureBracketedStillImageSettings) Autorelease() AVCaptureAutoExposureBracketedStillImageSettings {
	rv := objc.Send[AVCaptureAutoExposureBracketedStillImageSettings](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureAutoExposureBracketedStillImageSettings creates a new AVCaptureAutoExposureBracketedStillImageSettings instance.
func NewAVCaptureAutoExposureBracketedStillImageSettings() AVCaptureAutoExposureBracketedStillImageSettings {
	return getAVCaptureAutoExposureBracketedStillImageSettingsClass().New()
}




