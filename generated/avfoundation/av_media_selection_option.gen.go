// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MediaSelectionOption] class.
var (
	MediaSelectionOptionClass     _MediaSelectionOptionClass
	MediaSelectionOptionClassOnce sync.Once
)

func getMediaSelectionOptionClass() _MediaSelectionOptionClass {
	MediaSelectionOptionClassOnce.Do(func() {
		MediaSelectionOptionClass = _MediaSelectionOptionClass{objc.GetClass("AVMediaSelectionOption")}
	})
	return MediaSelectionOptionClass
}

type _MediaSelectionOptionClass struct {
	class objc.Class
}





// An interface definition for the [MediaSelectionOption] class.
type IMediaSelectionOption interface {
	objectivec.IObject
	

	// properties:
	AvailableMetadataFormats() []string
	CommonMetadata() []MetadataItem
	DisplayName() objc.IObject /* cross-framework: NSString */
	ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */
	Playable() bool
	Locale() foundation.Locale
	MediaSubTypes() []foundation.Number
	MediaType() MediaType /* typedef */
	IsPlayable() bool
	SetIsPlayable(value bool)


	

	// methods:
	AssociatedMediaSelectionOptionInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) IMediaSelectionOption
	DisplayNameWithLocale(locale foundation.Locale) foundation.String
	HasMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) bool
	MakeNowPlayingInfoLanguageOption() objectivec.IObject
	MetadataForFormat(format objc.IObject /* cross-framework: NSString */) []MetadataItem
	PropertyList() objc.ID


}





// Alloc allocates a new instance without initialization.
func (mc _MediaSelectionOptionClass) Alloc() MediaSelectionOption {
	rv := objc.Send[MediaSelectionOption](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaSelectionOptionClass) New() MediaSelectionOption {
	rv := objc.Send[MediaSelectionOption](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaSelectionOption) Init() MediaSelectionOption {
	rv := objc.Send[MediaSelectionOption](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaSelectionOption) Autorelease() MediaSelectionOption {
	rv := objc.Send[MediaSelectionOption](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaSelectionOption creates a new MediaSelectionOption instance.
func NewMediaSelectionOption() MediaSelectionOption {
	return getMediaSelectionOptionClass().New()
}





// An object that represents a specific option for the presentation of media within a group of options.


// An object that represents a specific option for the presentation of media within a group of options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption
type MediaSelectionOption struct {
	objectivec.Object
}

// MediaSelectionOptionFrom constructs a [MediaSelectionOption] from an unsafe.Pointer.
//
// An object that represents a specific option for the presentation of media within a group of options.
func MediaSelectionOptionFrom(ptr unsafe.Pointer) MediaSelectionOption {
	return MediaSelectionOption{objectivec.Object{objc.ID(ptr)}}
}




















// Returns a media selection option associated with the receiver in a given group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/associatedMediaSelectionOption(in:)
func (m_ MediaSelectionOption) AssociatedMediaSelectionOptionInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) IMediaSelectionOption {
	rv := objc.Send[MediaSelectionOption](m_.ID, objc.Sel("associatedMediaSelectionOptionInMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}


// Returns a string suitable for display using the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/displayName(with:)
func (m_ MediaSelectionOption) DisplayNameWithLocale(locale foundation.Locale) foundation.String {
	rv := objc.Send[foundation.String](m_.ID, objc.Sel("displayNameWithLocale:"), locale)
	return rv
}


// Returns a Boolean value that indicates whether the receiver has media with the given media characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/hasMediaCharacteristic(_:)
func (m_ MediaSelectionOption) HasMediaCharacteristic(mediaCharacteristic MediaCharacteristic /* typedef */) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasMediaCharacteristic:"), mediaCharacteristic)
	return rv
}


// Creates a language option for a media selection option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/makeNowPlayingInfoLanguageOption()
func (m_ MediaSelectionOption) MakeNowPlayingInfoLanguageOption() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("makeNowPlayingInfoLanguageOption"))
	return rv
}


// Returns an array of metadata items—one for each metadata item in the container of a given format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/metadata(forFormat:)
func (m_ MediaSelectionOption) MetadataForFormat(format objc.IObject /* cross-framework: NSString */) []MetadataItem {
	rv := objc.Send[[]MetadataItem](m_.ID, objc.Sel("metadataForFormat:"), format)
	return rv
}


// Returns a serializable property list that’s sufficient to identify the option within its group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/propertyList()
func (m_ MediaSelectionOption) PropertyList() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("propertyList"))
	return rv
}







// The metadata formats that contain metadata associated with the option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/availableMetadataFormats
func (m_ MediaSelectionOption) AvailableMetadataFormats() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// An array of metadata items for each common metadata key for which a value is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/commonMetadata
func (m_ MediaSelectionOption) CommonMetadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](m_.ID, objc.Sel("commonMetadata"))
	return rv
}


// A string suitable for display using the current system locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/displayName
func (m_ MediaSelectionOption) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("displayName"))
	return rv
}


// The IETF BCP 47 language tag associated with the option
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/extendedLanguageTag
func (m_ MediaSelectionOption) ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// A Boolean value that indicates whether the media selection option is playable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/isPlayable
func (m_ MediaSelectionOption) Playable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("playable"))
	return rv
}


// The locale for which the media option was authored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/locale
func (m_ MediaSelectionOption) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](m_.ID, objc.Sel("locale"))
	return rv
}


// The media sub-types of the media data associated with the option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/mediaSubTypes
func (m_ MediaSelectionOption) MediaSubTypes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("mediaSubTypes"))
	return rv
}


// The media type of the media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/mediaType
func (m_ MediaSelectionOption) MediaType() MediaType /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("mediaType"))
	return rv
}


// A Boolean value that indicates whether the media selection option is playable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/isplayable
func (m_ MediaSelectionOption) IsPlayable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isPlayable"))
	return rv
}


// A Boolean value that indicates whether the media selection option is playable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/isplayable
func (m_ MediaSelectionOption) SetIsPlayable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsPlayable:"), value)
}








