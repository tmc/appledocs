// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptureManualExposureBracketedStillImageSettings] class.
var (
	CaptureManualExposureBracketedStillImageSettingsClass     _CaptureManualExposureBracketedStillImageSettingsClass
	CaptureManualExposureBracketedStillImageSettingsClassOnce sync.Once
)

func getCaptureManualExposureBracketedStillImageSettingsClass() _CaptureManualExposureBracketedStillImageSettingsClass {
	CaptureManualExposureBracketedStillImageSettingsClassOnce.Do(func() {
		CaptureManualExposureBracketedStillImageSettingsClass = _CaptureManualExposureBracketedStillImageSettingsClass{objc.GetClass("AVCaptureManualExposureBracketedStillImageSettings")}
	})
	return CaptureManualExposureBracketedStillImageSettingsClass
}

type _CaptureManualExposureBracketedStillImageSettingsClass struct {
	class objc.Class
}





// An interface definition for the [CaptureManualExposureBracketedStillImageSettings] class.
type ICaptureManualExposureBracketedStillImageSettings interface {
	ICaptureBracketedStillImageSettings
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureManualExposureBracketedStillImageSettingsClass) Alloc() CaptureManualExposureBracketedStillImageSettings {
	rv := objc.Send[CaptureManualExposureBracketedStillImageSettings](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureManualExposureBracketedStillImageSettingsClass) New() CaptureManualExposureBracketedStillImageSettings {
	rv := objc.Send[CaptureManualExposureBracketedStillImageSettings](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureManualExposureBracketedStillImageSettings) Init() CaptureManualExposureBracketedStillImageSettings {
	rv := objc.Send[CaptureManualExposureBracketedStillImageSettings](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureManualExposureBracketedStillImageSettings) Autorelease() CaptureManualExposureBracketedStillImageSettings {
	rv := objc.Send[CaptureManualExposureBracketedStillImageSettings](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureManualExposureBracketedStillImageSettings creates a new CaptureManualExposureBracketedStillImageSettings instance.
func NewCaptureManualExposureBracketedStillImageSettings() CaptureManualExposureBracketedStillImageSettings {
	return getCaptureManualExposureBracketedStillImageSettingsClass().New()
}





// A configuration for defining bracketed photo captures in terms of specific exposure and ISO values.
//
// The class is a concrete subclass of the class used when bracketing exposure duration and ISO. An instance defines exposure duration and ISO settings that should be applied to one image in a bracket. An array of objects is passed to to specify the bracketing. You can query the minimum and maximum duration and ISO properties of the instance supplying data to an instance. If you wish to leave unchanged for this bracketed still image, you pass the value when creating the instance. To keep the ISO unchanged, you pass when creating the instance.


// A configuration for defining bracketed photo captures in terms of specific exposure and ISO values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureManualExposureBracketedStillImageSettings
type CaptureManualExposureBracketedStillImageSettings struct {
	CaptureBracketedStillImageSettings
}

// CaptureManualExposureBracketedStillImageSettingsFrom constructs a [CaptureManualExposureBracketedStillImageSettings] from an unsafe.Pointer.
//
// A configuration for defining bracketed photo captures in terms of specific exposure and ISO values.
func CaptureManualExposureBracketedStillImageSettingsFrom(ptr unsafe.Pointer) CaptureManualExposureBracketedStillImageSettings {
	return CaptureManualExposureBracketedStillImageSettings{
		CaptureBracketedStillImageSettings: CaptureBracketedStillImageSettingsFrom(ptr),
	}
}










// Creates a configuration of still image settings using the specified exposure duration and ISO.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureManualExposureBracketedStillImageSettings/manualExposureSettings(exposureDuration:iso:)
func (cc _CaptureManualExposureBracketedStillImageSettingsClass) ManualExposureSettingsWithExposureDurationISO(duration objectivec.IObject, ISO float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("manualExposureSettingsWithExposureDuration:ISO:"), duration, ISO)
	return rv
}






















