// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MediaPresentationSetting] class.
var (
	MediaPresentationSettingClass     _MediaPresentationSettingClass
	MediaPresentationSettingClassOnce sync.Once
)

func getMediaPresentationSettingClass() _MediaPresentationSettingClass {
	MediaPresentationSettingClassOnce.Do(func() {
		MediaPresentationSettingClass = _MediaPresentationSettingClass{objc.GetClass("AVMediaPresentationSetting")}
	})
	return MediaPresentationSettingClass
}

type _MediaPresentationSettingClass struct {
	class objc.Class
}

// An interface definition for the [MediaPresentationSetting] class.
type IMediaPresentationSetting interface {
	objectivec.IObject
	// properties:
	MediaCharacteristic() MediaCharacteristic /* not a class type */
	SetMediaCharacteristic(value MediaCharacteristic /* not a class type */)
	// methods:
}

// For content that has been authored with the express intent of offering an alternative selection interface for AVMediaSelectionOptions, AVMediaPresentationSetting represents a selectable setting for controlling the presentation of the media.
//
// Each selectable setting is associated with a media characteristic that one or more of the AVMediaSelectionOptions in the AVMediaSelectionGroup possesses. By selecting a setting in a user interface that offers AVMediaPresentationSettings, users are essentially indicating a preference for the media characteristic of the selected setting. Subclasses of this type that are used from Swift must fulfill the requirements of a Sendable type.


// For content that has been authored with the express intent of offering an alternative selection interface for AVMediaSelectionOptions, AVMediaPresentationSetting represents a selectable setting for controlling the presentation of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSetting
type MediaPresentationSetting struct {
	objectivec.Object
}

// MediaPresentationSettingFrom constructs a [MediaPresentationSetting] from an unsafe.Pointer.
//
// For content that has been authored with the express intent of offering an alternative selection interface for AVMediaSelectionOptions, AVMediaPresentationSetting represents a selectable setting for controlling the presentation of the media.
func MediaPresentationSettingFrom(ptr unsafe.Pointer) MediaPresentationSetting {
	return MediaPresentationSetting{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaPresentationSettingClass) Alloc() MediaPresentationSetting {
	rv := objc.Send[MediaPresentationSetting](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaPresentationSettingClass) New() MediaPresentationSetting {
	rv := objc.Send[MediaPresentationSetting](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaPresentationSetting) Init() MediaPresentationSetting {
	rv := objc.Send[MediaPresentationSetting](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaPresentationSetting) Autorelease() MediaPresentationSetting {
	rv := objc.Send[MediaPresentationSetting](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaPresentationSetting creates a new MediaPresentationSetting instance.
func NewMediaPresentationSetting() MediaPresentationSetting {
	return getMediaPresentationSettingClass().New()
}



// Provides the media characteristic that corresponds to the selectable setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediapresentationsetting/mediacharacteristic
func (m_ MediaPresentationSetting) MediaCharacteristic() MediaCharacteristic /* not a class type */ {
	rv := objc.Send[MediaCharacteristic](m_.ID, objc.Sel("mediaCharacteristic"))
	return rv
}


// Provides the media characteristic that corresponds to the selectable setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediapresentationsetting/mediacharacteristic
func (m_ MediaPresentationSetting) SetMediaCharacteristic(value MediaCharacteristic /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaCharacteristic:"), value)
}



