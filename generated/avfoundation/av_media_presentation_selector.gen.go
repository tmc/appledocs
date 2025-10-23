// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	Settings() IAVMediaPresentationSetting
	SetSettings(value IAVMediaPresentationSetting)
	// methods:
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

// Alloc allocates a new instance without initialization.
func (mc _MediaPresentationSelectorClass) Alloc() MediaPresentationSelector {
	rv := objc.Send[MediaPresentationSelector](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Provides the authored identifier for the selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediapresentationselector/identifier
func (m_ MediaPresentationSelector) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("identifier"))
	return rv
}


// Provides the authored identifier for the selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediapresentationselector/identifier
func (m_ MediaPresentationSelector) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), value)
}


// Provides selectable mutually exclusive settings for the selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediapresentationselector/settings
func (m_ MediaPresentationSelector) Settings() IAVMediaPresentationSetting {
	rv := objc.Send[MediaPresentationSetting](m_.ID, objc.Sel("settings"))
	return rv
}


// Provides selectable mutually exclusive settings for the selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediapresentationselector/settings
func (m_ MediaPresentationSelector) SetSettings(value IAVMediaPresentationSetting) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSettings:"), value)
}



