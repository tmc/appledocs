// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMediaPresentationSetting */


/* debug [class_header]: Header for AVMediaPresentationSetting */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaPresentationSetting */
// An interface definition for the [MediaPresentationSetting] class.
type IMediaPresentationSetting interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaPresentationSetting */
	// properties:
	MediaCharacteristic() MediaCharacteristic /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaPresentationSetting */
	// methods:
	DisplayNameForLocaleIdentifier(localeIdentifier objc.IObject /* cross-framework: NSString */) foundation.String
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaPresentationSetting */
// Alloc allocates a new instance without initialization.
func (mc _MediaPresentationSettingClass) Alloc() MediaPresentationSetting {
	rv := objc.Send[MediaPresentationSetting](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaPresentationSetting */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaPresentationSetting *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaPresentationSetting */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaPresentationSetting */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaPresentationSetting */

// Returns the display name for the selectable setting that best matches the specified locale identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSetting/displayName(forLocaleIdentifier:)
func (m_ MediaPresentationSetting) DisplayNameForLocaleIdentifier(localeIdentifier objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](m_.ID, objc.Sel("displayNameForLocaleIdentifier:"), localeIdentifier)
	return rv
}/* debug [instance_methods/method]: DisplayNameForLocaleIdentifier */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaPresentationSetting */

// Provides the media characteristic that corresponds to the selectable setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSetting/mediaCharacteristic
func (m_ MediaPresentationSetting) MediaCharacteristic() MediaCharacteristic /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("mediaCharacteristic"))
	return rv
}/* debug [instance_properties/getter]: mediaCharacteristic */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMediaPresentationSetting */



