// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MediaSelectionGroup] class.
var (
	MediaSelectionGroupClass     _MediaSelectionGroupClass
	MediaSelectionGroupClassOnce sync.Once
)

func getMediaSelectionGroupClass() _MediaSelectionGroupClass {
	MediaSelectionGroupClassOnce.Do(func() {
		MediaSelectionGroupClass = _MediaSelectionGroupClass{objc.GetClass("AVMediaSelectionGroup")}
	})
	return MediaSelectionGroupClass
}

type _MediaSelectionGroupClass struct {
	class objc.Class
}





// An interface definition for the [MediaSelectionGroup] class.
type IMediaSelectionGroup interface {
	objectivec.IObject
	

	// properties:
	AllowsEmptySelection() bool
	CustomMediaSelectionScheme() IAVCustomMediaSelectionScheme
	DefaultOption() IAVMediaSelectionOption
	Options() []MediaSelectionOption


	

	// methods:
	MakeNowPlayingInfoLanguageOptionGroup() NowPlayingInfoLanguageOptionGroup /* not a class type */
	MediaSelectionOptionWithPropertyList(plist objectivec.IObject) IMediaSelectionOption


}





// Alloc allocates a new instance without initialization.
func (mc _MediaSelectionGroupClass) Alloc() MediaSelectionGroup {
	rv := objc.Send[MediaSelectionGroup](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaSelectionGroupClass) New() MediaSelectionGroup {
	rv := objc.Send[MediaSelectionGroup](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaSelectionGroup) Init() MediaSelectionGroup {
	rv := objc.Send[MediaSelectionGroup](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaSelectionGroup) Autorelease() MediaSelectionGroup {
	rv := objc.Send[MediaSelectionGroup](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaSelectionGroup creates a new MediaSelectionGroup instance.
func NewMediaSelectionGroup() MediaSelectionGroup {
	return getMediaSelectionGroupClass().New()
}





// An object that represents a collection of mutually exclusive options for the presentation of media within an asset.


// An object that represents a collection of mutually exclusive options for the presentation of media within an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup
type MediaSelectionGroup struct {
	objectivec.Object
}

// MediaSelectionGroupFrom constructs a [MediaSelectionGroup] from an unsafe.Pointer.
//
// An object that represents a collection of mutually exclusive options for the presentation of media within an asset.
func MediaSelectionGroupFrom(ptr unsafe.Pointer) MediaSelectionGroup {
	return MediaSelectionGroup{objectivec.Object{objc.ID(ptr)}}
}










// Returns an array of media selection options, filtering them according to whether their locales match one of the specified languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/mediaSelectionOptions(from:filteredAndSortedAccordingToPreferredLanguages:)
func (mc _MediaSelectionGroupClass) MediaSelectionOptionsFromArrayFilteredAndSortedAccordingToPreferredLanguages(mediaSelectionOptions []MediaSelectionOption, preferredLanguages []string) []MediaSelectionOption {
	rv := objc.Send[[]MediaSelectionOption](objc.ID(mc.class), objc.Sel("mediaSelectionOptionsFromArray:filteredAndSortedAccordingToPreferredLanguages:"), mediaSelectionOptions, preferredLanguages)
	return rv
}


// Returns an array containing the media selection options from a given array that match the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/mediaSelectionOptions(from:with:)
func (mc _MediaSelectionGroupClass) MediaSelectionOptionsFromArrayWithLocale(mediaSelectionOptions []MediaSelectionOption, locale foundation.Locale) []MediaSelectionOption {
	rv := objc.Send[[]MediaSelectionOption](objc.ID(mc.class), objc.Sel("mediaSelectionOptionsFromArray:withLocale:"), mediaSelectionOptions, locale)
	return rv
}


// Returns an array containing the media selection options from a given array that match given media characteristics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/mediaSelectionOptions(from:withMediaCharacteristics:)
func (mc _MediaSelectionGroupClass) MediaSelectionOptionsFromArrayWithMediaCharacteristics(mediaSelectionOptions []MediaSelectionOption, mediaCharacteristics []string) []MediaSelectionOption {
	rv := objc.Send[[]MediaSelectionOption](objc.ID(mc.class), objc.Sel("mediaSelectionOptionsFromArray:withMediaCharacteristics:"), mediaSelectionOptions, mediaCharacteristics)
	return rv
}


// Returns an array containing the media selection options from a given array that do not match given media characteristics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/mediaSelectionOptions(from:withoutMediaCharacteristics:)
func (mc _MediaSelectionGroupClass) MediaSelectionOptionsFromArrayWithoutMediaCharacteristics(mediaSelectionOptions []MediaSelectionOption, mediaCharacteristics []string) []MediaSelectionOption {
	rv := objc.Send[[]MediaSelectionOption](objc.ID(mc.class), objc.Sel("mediaSelectionOptionsFromArray:withoutMediaCharacteristics:"), mediaSelectionOptions, mediaCharacteristics)
	return rv
}


// Returns an array containing the media selection options from a given array that are playable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/playableMediaSelectionOptions(from:)
func (mc _MediaSelectionGroupClass) PlayableMediaSelectionOptionsFromArray(mediaSelectionOptions []MediaSelectionOption) []MediaSelectionOption {
	rv := objc.Send[[]MediaSelectionOption](objc.ID(mc.class), objc.Sel("playableMediaSelectionOptionsFromArray:"), mediaSelectionOptions)
	return rv
}












// Creates a language option group from the media selection group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/makeNowPlayingInfoLanguageOptionGroup()
func (m_ MediaSelectionGroup) MakeNowPlayingInfoLanguageOptionGroup() NowPlayingInfoLanguageOptionGroup /* not a class type */ {
	rv := objc.Send[NowPlayingInfoLanguageOptionGroup](m_.ID, objc.Sel("makeNowPlayingInfoLanguageOptionGroup"))
	return rv
}


// Returns the media selection options that match the given property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/mediaSelectionOption(withPropertyList:)
func (m_ MediaSelectionGroup) MediaSelectionOptionWithPropertyList(plist objectivec.IObject) IMediaSelectionOption {
	rv := objc.Send[MediaSelectionOption](m_.ID, objc.Sel("mediaSelectionOptionWithPropertyList:"), plist)
	return rv
}







// A Boolean value that indicates whether it’s possible to present none of the options in the group when an associated player item is played.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/allowsEmptySelection
func (m_ MediaSelectionGroup) AllowsEmptySelection() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsEmptySelection"))
	return rv
}


// For content that has been authored with the express intent of offering an alternative selection interface for AVMediaSelectionOptions, AVCustomMediaSelectionScheme provides a collection of custom settings for controlling the presentation of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/customMediaSelectionScheme
func (m_ MediaSelectionGroup) CustomMediaSelectionScheme() IAVCustomMediaSelectionScheme {
	rv := objc.Send[CustomMediaSelectionScheme](m_.ID, objc.Sel("customMediaSelectionScheme"))
	return rv
}


// The default option in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/defaultOption
func (m_ MediaSelectionGroup) DefaultOption() IAVMediaSelectionOption {
	rv := objc.Send[MediaSelectionOption](m_.ID, objc.Sel("defaultOption"))
	return rv
}


// A collection of mutually exclusive media selection options
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/options
func (m_ MediaSelectionGroup) Options() []MediaSelectionOption {
	rv := objc.Send[[]MediaSelectionOption](m_.ID, objc.Sel("options"))
	return rv
}








