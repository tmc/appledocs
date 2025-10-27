// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	AvailableLanguages() []string
	Selectors() []MediaPresentationSelector
	ShouldOfferLanguageSelection() bool


	

	// methods:
	MediaPresentationSettingsForSelectorComplementaryToLanguageSettings(selector IAVMediaPresentationSelector, language foundation.foundation.INSString, settings []MediaPresentationSetting) []MediaPresentationSetting


}





// Alloc allocates a new instance without initialization.
func (cc _CustomMediaSelectionSchemeClass) Alloc() CustomMediaSelectionScheme {
	rv := objc.Send[CustomMediaSelectionScheme](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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




















// Provides an array of media presentation settings that can be effective at the same time as the specified language and settings for other selectors of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCustomMediaSelectionScheme/mediaPresentationSettings(for:complementaryToLanguage:settings:)
func (c_ CustomMediaSelectionScheme) MediaPresentationSettingsForSelectorComplementaryToLanguageSettings(selector IAVMediaPresentationSelector, language foundation.foundation.INSString, settings []MediaPresentationSetting) []MediaPresentationSetting {
	rv := objc.Send[[]MediaPresentationSetting](c_.ID, objc.Sel("mediaPresentationSettingsForSelector:complementaryToLanguage:settings:"), selector, language, settings)
	return rv
}







// Provides available language choices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCustomMediaSelectionScheme/availableLanguages
func (c_ CustomMediaSelectionScheme) AvailableLanguages() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableLanguages"))
	return rv
}


// Provides custom settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCustomMediaSelectionScheme/selectors
func (c_ CustomMediaSelectionScheme) Selectors() []MediaPresentationSelector {
	rv := objc.Send[[]MediaPresentationSelector](c_.ID, objc.Sel("selectors"))
	return rv
}


// Indicates whether an alternative selection interface should provide a menu of language choices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCustomMediaSelectionScheme/shouldOfferLanguageSelection
func (c_ CustomMediaSelectionScheme) ShouldOfferLanguageSelection() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldOfferLanguageSelection"))
	return rv
}








