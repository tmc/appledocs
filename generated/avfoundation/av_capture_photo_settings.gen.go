// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CapturePhotoSettings] class.
var (
	CapturePhotoSettingsClass     _CapturePhotoSettingsClass
	CapturePhotoSettingsClassOnce sync.Once
)

func getCapturePhotoSettingsClass() _CapturePhotoSettingsClass {
	CapturePhotoSettingsClassOnce.Do(func() {
		CapturePhotoSettingsClass = _CapturePhotoSettingsClass{objc.GetClass("AVCapturePhotoSettings")}
	})
	return CapturePhotoSettingsClass
}

type _CapturePhotoSettingsClass struct {
	class objc.Class
}

// An interface definition for the [CapturePhotoSettings] class.
type ICapturePhotoSettings interface {
	objectivec.IObject
}

// A specification of the features and settings to use for a single photo capture request.
//
// To take a photo, you create and configure a object, then pass it to the method. A instance can include any combination of settings, regardless of whether that combination is valid for a given capture session. When you initiate a capture by passing a photo settings object to the method, the photo capture output validates your settings to ensure deterministic behavior. For example, the setting must specify a value that’s present in the photo output’s array. For detailed validation rules, see each property description below.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings
type CapturePhotoSettings struct {
	objectivec.Object
}

// CapturePhotoSettingsFrom constructs a [CapturePhotoSettings] from an unsafe.Pointer.
//
// A specification of the features and settings to use for a single photo capture request.
func CapturePhotoSettingsFrom(ptr unsafe.Pointer) CapturePhotoSettings {
	return CapturePhotoSettings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CapturePhotoSettingsClass) Alloc() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CapturePhotoSettingsClass) New() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CapturePhotoSettings) Init() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CapturePhotoSettings) Autorelease() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCapturePhotoSettings creates a new CapturePhotoSettings instance.
func NewCapturePhotoSettings() CapturePhotoSettings {
	return getCapturePhotoSettingsClass().New()
}


// Specifies whether a portrait effects matte should be captured along with the photo.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isPortraitEffectsMatteDeliveryEnabled
func (c_ CapturePhotoSettings) PortraitEffectsMatteDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("portraitEffectsMatteDeliveryEnabled"))
	return rv
}


// SetPortraitEffectsMatteDeliveryEnabled sets the value of the portraitEffectsMatteDeliveryEnabled property.
// Specifies whether a portrait effects matte should be captured along with the photo.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isPortraitEffectsMatteDeliveryEnabled
func (c_ CapturePhotoSettings) SetPortraitEffectsMatteDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPortraitEffectsMatteDeliveryEnabled:"), value)
}



