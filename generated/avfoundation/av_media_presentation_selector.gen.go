// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [MediaPresentationSelector] class.
type IMediaPresentationSelector interface {
	objectivec.IObject
	

	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	Settings() []MediaPresentationSetting


	

	// methods:
	DisplayNameForLocaleIdentifier(localeIdentifier objc.IObject /* cross-framework: NSString */) foundation.String


}





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




















// Returns the display name for the selector that best matches the specified locale identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSelector/displayName(forLocaleIdentifier:)
func (m_ MediaPresentationSelector) DisplayNameForLocaleIdentifier(localeIdentifier objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](m_.ID, objc.Sel("displayNameForLocaleIdentifier:"), localeIdentifier)
	return rv
}







// Provides the authored identifier for the selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSelector/identifier
func (m_ MediaPresentationSelector) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("identifier"))
	return rv
}


// Provides selectable mutually exclusive settings for the selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSelector/settings
func (m_ MediaPresentationSelector) Settings() []MediaPresentationSetting {
	rv := objc.Send[[]MediaPresentationSetting](m_.ID, objc.Sel("settings"))
	return rv
}








