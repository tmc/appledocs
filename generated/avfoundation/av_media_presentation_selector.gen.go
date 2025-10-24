// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMediaPresentationSelector */


/* debug [class_header]: Header for AVMediaPresentationSelector */
// The class instance for the [MediaPresentationSelector] class.
var (
	MediaPresentationSelectorClass     _MediaPresentationSelectorClass
	MediaPresentationSelectorClassOnce sync.Once
)

func getMediaPresentationSelectorClass() _MediaPresentationSelectorClass {
	MediaPresentationSelectorClassOnce.Do(func() {
		MediaPresentationSelectorClass = _MediaPresentationSelectorClass{objc.GetClass("AVMediaPresentationSelector")}
	})
	return MediaPresentationSelectorClass
}

type _MediaPresentationSelectorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaPresentationSelector */
// An interface definition for the [MediaPresentationSelector] class.
type IMediaPresentationSelector interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaPresentationSelector */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	Settings() []MediaPresentationSetting
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaPresentationSelector */
	// methods:
	DisplayNameForLocaleIdentifier(localeIdentifier objc.IObject /* cross-framework: NSString */) foundation.String
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaPresentationSelector */
// Alloc allocates a new instance without initialization.
func (mc _MediaPresentationSelectorClass) Alloc() MediaPresentationSelector {
	rv := objc.Send[MediaPresentationSelector](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaPresentationSelectorClass) New() MediaPresentationSelector {
	rv := objc.Send[MediaPresentationSelector](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaPresentationSelector) Init() MediaPresentationSelector {
	rv := objc.Send[MediaPresentationSelector](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaPresentationSelector) Autorelease() MediaPresentationSelector {
	rv := objc.Send[MediaPresentationSelector](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaPresentationSelector creates a new MediaPresentationSelector instance.
func NewMediaPresentationSelector() MediaPresentationSelector {
	return getMediaPresentationSelectorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaPresentationSelector */
// For content that has been authored with the express intent of offering an alternative selection interface for AVMediaSelectionOptions, AVMediaPresentationSelector represents a collection of mutually exclusive settings.
//
// Subclasses of this type that are used from Swift must fulfill the requirements of a Sendable type.


// For content that has been authored with the express intent of offering an alternative selection interface for AVMediaSelectionOptions, AVMediaPresentationSelector represents a collection of mutually exclusive settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSelector
type MediaPresentationSelector struct {
	objectivec.Object
}

// MediaPresentationSelectorFrom constructs a [MediaPresentationSelector] from an unsafe.Pointer.
//
// For content that has been authored with the express intent of offering an alternative selection interface for AVMediaSelectionOptions, AVMediaPresentationSelector represents a collection of mutually exclusive settings.
func MediaPresentationSelectorFrom(ptr unsafe.Pointer) MediaPresentationSelector {
	return MediaPresentationSelector{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaPresentationSelector *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaPresentationSelector */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaPresentationSelector */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaPresentationSelector */

// Returns the display name for the selector that best matches the specified locale identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSelector/displayName(forLocaleIdentifier:)
func (m_ MediaPresentationSelector) DisplayNameForLocaleIdentifier(localeIdentifier objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](m_.ID, objc.Sel("displayNameForLocaleIdentifier:"), localeIdentifier)
	return rv
}/* debug [instance_methods/method]: DisplayNameForLocaleIdentifier */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaPresentationSelector */

// Provides the authored identifier for the selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSelector/identifier
func (m_ MediaPresentationSelector) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// Provides selectable mutually exclusive settings for the selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSelector/settings
func (m_ MediaPresentationSelector) Settings() []MediaPresentationSetting {
	rv := objc.Send[[]MediaPresentationSetting](m_.ID, objc.Sel("settings"))
	return rv
}/* debug [instance_properties/getter]: settings */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMediaPresentationSelector */



