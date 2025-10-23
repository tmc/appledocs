// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CustomMediaSelectionScheme] class.
var (
	CustomMediaSelectionSchemeClass     _CustomMediaSelectionSchemeClass
	CustomMediaSelectionSchemeClassOnce sync.Once
)

func getCustomMediaSelectionSchemeClass() _CustomMediaSelectionSchemeClass {
	CustomMediaSelectionSchemeClassOnce.Do(func() {
		CustomMediaSelectionSchemeClass = _CustomMediaSelectionSchemeClass{objc.GetClass("AVCustomMediaSelectionScheme")}
	})
	return CustomMediaSelectionSchemeClass
}

type _CustomMediaSelectionSchemeClass struct {
	class objc.Class
}

// An interface definition for the [CustomMediaSelectionScheme] class.
type ICustomMediaSelectionScheme interface {
	objectivec.IObject
	// properties:
	AvailableLanguages() objc.IObject /* cross-framework: NSString */
	SetAvailableLanguages(value objc.IObject /* cross-framework: NSString */)
	Selectors() IAVMediaPresentationSelector
	SetSelectors(value IAVMediaPresentationSelector)
	ShouldOfferLanguageSelection() bool /* primitive/slice/pointer. */
	SetShouldOfferLanguageSelection(value bool /* primitive/slice/pointer. */)
	// methods:
}

// For content that has been authored with the express intent of offering an alternative selection interface for AVMediaSelectionOptions, AVCustomMediaSelectionScheme provides a collection of custom settings for controlling the presentation of the media.
//
// Each selectable setting is associated with a media characteristic that one or more of the AVMediaSelectionOptions in the AVMediaSelectionGroup possesses. By selecting a setting in a user interface based on an AVCustomMediaSelectionScheme, users are essentially indicating a preference for the media characteristic of the selected setting. Selection of a specific AVMediaSelectionOption in the AVMediaSelectionGroup is then derived from the user’s indicated preferences. Subclasses of this type that are used from Swift must fulfill the requirements of a Sendable type.


// For content that has been authored with the express intent of offering an alternative selection interface for AVMediaSelectionOptions, AVCustomMediaSelectionScheme provides a collection of custom settings for controlling the presentation of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCustomMediaSelectionScheme
type CustomMediaSelectionScheme struct {
	objectivec.Object
}

// CustomMediaSelectionSchemeFrom constructs a [CustomMediaSelectionScheme] from an unsafe.Pointer.
//
// For content that has been authored with the express intent of offering an alternative selection interface for AVMediaSelectionOptions, AVCustomMediaSelectionScheme provides a collection of custom settings for controlling the presentation of the media.
func CustomMediaSelectionSchemeFrom(ptr unsafe.Pointer) CustomMediaSelectionScheme {
	return CustomMediaSelectionScheme{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CustomMediaSelectionSchemeClass) Alloc() CustomMediaSelectionScheme {
	rv := objc.Send[CustomMediaSelectionScheme](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CustomMediaSelectionSchemeClass) New() CustomMediaSelectionScheme {
	rv := objc.Send[CustomMediaSelectionScheme](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomMediaSelectionScheme) Init() CustomMediaSelectionScheme {
	rv := objc.Send[CustomMediaSelectionScheme](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomMediaSelectionScheme) Autorelease() CustomMediaSelectionScheme {
	rv := objc.Send[CustomMediaSelectionScheme](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomMediaSelectionScheme creates a new CustomMediaSelectionScheme instance.
func NewCustomMediaSelectionScheme() CustomMediaSelectionScheme {
	return getCustomMediaSelectionSchemeClass().New()
}



// Provides available language choices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcustommediaselectionscheme/availablelanguages
func (c_ CustomMediaSelectionScheme) AvailableLanguages() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("availableLanguages"))
	return rv
}


// Provides available language choices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcustommediaselectionscheme/availablelanguages
func (c_ CustomMediaSelectionScheme) SetAvailableLanguages(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableLanguages:"), value)
}


// Provides custom settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcustommediaselectionscheme/selectors
func (c_ CustomMediaSelectionScheme) Selectors() IAVMediaPresentationSelector {
	rv := objc.Send[MediaPresentationSelector](c_.ID, objc.Sel("selectors"))
	return rv
}


// Provides custom settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcustommediaselectionscheme/selectors
func (c_ CustomMediaSelectionScheme) SetSelectors(value IAVMediaPresentationSelector) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSelectors:"), value)
}


// Indicates whether an alternative selection interface should provide a menu of language choices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcustommediaselectionscheme/shouldofferlanguageselection
func (c_ CustomMediaSelectionScheme) ShouldOfferLanguageSelection() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldOfferLanguageSelection"))
	return rv
}


// Indicates whether an alternative selection interface should provide a menu of language choices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcustommediaselectionscheme/shouldofferlanguageselection
func (c_ CustomMediaSelectionScheme) SetShouldOfferLanguageSelection(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldOfferLanguageSelection:"), value)
}



