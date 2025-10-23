// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AvailableMetadataFormats() objc.IObject /* cross-framework: NSString */
	SetAvailableMetadataFormats(value objc.IObject /* cross-framework: NSString */)
	CommonMetadata() IAVMetadataItem
	SetCommonMetadata(value IAVMetadataItem)
	DisplayName() objc.IObject /* cross-framework: NSString */
	SetDisplayName(value objc.IObject /* cross-framework: NSString */)
	ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */
	SetExtendedLanguageTag(value objc.IObject /* cross-framework: NSString */)
	IsPlayable() bool /* primitive/slice/pointer. */
	SetIsPlayable(value bool /* primitive/slice/pointer. */)
	Locale() unsafe.Pointer
	SetLocale(value unsafe.Pointer)
	MediaSubTypes() objc.IObject /* cross-framework: NSNumber */
	SetMediaSubTypes(value objc.IObject /* cross-framework: NSNumber */)
	MediaType() MediaType /* not a class type */
	SetMediaType(value MediaType /* not a class type */)
	// methods:
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

// Alloc allocates a new instance without initialization.
func (mc _MediaSelectionOptionClass) Alloc() MediaSelectionOption {
	rv := objc.Send[MediaSelectionOption](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The metadata formats that contain metadata associated with the option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/availablemetadataformats
func (m_ MediaSelectionOption) AvailableMetadataFormats() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("availableMetadataFormats"))
	return rv
}


// The metadata formats that contain metadata associated with the option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/availablemetadataformats
func (m_ MediaSelectionOption) SetAvailableMetadataFormats(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAvailableMetadataFormats:"), value)
}


// An array of metadata items for each common metadata key for which a value is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/commonmetadata
func (m_ MediaSelectionOption) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](m_.ID, objc.Sel("commonMetadata"))
	return rv
}


// An array of metadata items for each common metadata key for which a value is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/commonmetadata
func (m_ MediaSelectionOption) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommonMetadata:"), value)
}


// A string suitable for display using the current system locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/displayname
func (m_ MediaSelectionOption) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("displayName"))
	return rv
}


// A string suitable for display using the current system locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/displayname
func (m_ MediaSelectionOption) SetDisplayName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDisplayName:"), value)
}


// The IETF BCP 47 language tag associated with the option
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/extendedlanguagetag
func (m_ MediaSelectionOption) ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// The IETF BCP 47 language tag associated with the option
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/extendedlanguagetag
func (m_ MediaSelectionOption) SetExtendedLanguageTag(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedLanguageTag:"), value)
}


// A Boolean value that indicates whether the media selection option is playable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/isplayable
func (m_ MediaSelectionOption) IsPlayable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isPlayable"))
	return rv
}


// A Boolean value that indicates whether the media selection option is playable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/isplayable
func (m_ MediaSelectionOption) SetIsPlayable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsPlayable:"), value)
}


// The locale for which the media option was authored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/locale
func (m_ MediaSelectionOption) Locale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("locale"))
	return rv
}


// The locale for which the media option was authored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/locale
func (m_ MediaSelectionOption) SetLocale(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocale:"), value)
}


// The media sub-types of the media data associated with the option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/mediasubtypes
func (m_ MediaSelectionOption) MediaSubTypes() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mediaSubTypes"))
	return rv
}


// The media sub-types of the media data associated with the option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/mediasubtypes
func (m_ MediaSelectionOption) SetMediaSubTypes(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaSubTypes:"), value)
}


// The media type of the media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/mediatype
func (m_ MediaSelectionOption) MediaType() MediaType /* not a class type */ {
	rv := objc.Send[MediaType](m_.ID, objc.Sel("mediaType"))
	return rv
}


// The media type of the media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaselectionoption/mediatype
func (m_ MediaSelectionOption) SetMediaType(value MediaType /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaType:"), value)
}



