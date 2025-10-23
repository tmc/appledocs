// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MetadataItem] class.
var (
	MetadataItemClass     _MetadataItemClass
	MetadataItemClassOnce sync.Once
)

func getMetadataItemClass() _MetadataItemClass {
	MetadataItemClassOnce.Do(func() {
		MetadataItemClass = _MetadataItemClass{objc.GetClass("AVMetadataItem")}
	})
	return MetadataItemClass
}

type _MetadataItemClass struct {
	class objc.Class
}

// An interface definition for the [MetadataItem] class.
type IMetadataItem interface {
	objectivec.IObject
	CommonMetadata() IAVMetadataItem
	SetCommonMetadata(value IAVMetadataItem)
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)
	CommonKey() unsafe.Pointer
	SetCommonKey(value unsafe.Pointer)
	DataType() string
	SetDataType(value string)
	DataValue() foundation.Data
	SetDataValue(value foundation.Data)
	DateValue() foundation.Date
	SetDateValue(value foundation.Date)
	Duration() unsafe.Pointer
	SetDuration(value unsafe.Pointer)
	ExtendedLanguageTag() string
	SetExtendedLanguageTag(value string)
	ExtraAttributes() unsafe.Pointer
	SetExtraAttributes(value unsafe.Pointer)
	Identifier() unsafe.Pointer
	SetIdentifier(value unsafe.Pointer)
	Key() unsafe.Pointer
	SetKey(value unsafe.Pointer)
	KeySpace() unsafe.Pointer
	SetKeySpace(value unsafe.Pointer)
	Locale() foundation.Locale
	SetLocale(value foundation.Locale)
	NumberValue() foundation.Number
	SetNumberValue(value foundation.Number)
	StartDate() foundation.Date
	SetStartDate(value foundation.Date)
	StringValue() string
	SetStringValue(value string)
	Time() unsafe.Pointer
	SetTime(value unsafe.Pointer)
	Value() unsafe.Pointer
	SetValue(value unsafe.Pointer)
}

// A metadata item for an audiovisual asset or one of its tracks.
//
// To effectively use , you need to understand how organizes metadata. To simplify finding and filtering metadata items, the framework groups related metadata into key spaces: The framework defines several format-specific key spaces. They roughly correlate to a particular container or file format, such as QuickTime (QuickTime metadata and user data) or MP3 (ID3). However, a single asset may contain metadata values across multiple key spaces. To retrieve an asset’s complete collection of format-specific metadata, you use its property. There are several common metadata values, such as a movie’s creation date or description, that can exist across multiple key spaces. To help normalize access to this common metadata, the framework provides a common key space that gives access to a limited set of metadata values common to several key spaces. This makes it easy to retrieve commonly used metadata without concern for the specific format. To retrieve an asset’s collection of common metadata, you use its property. Metadata items have keys that accord with the specification of the container format from which they’re drawn. Full details of the metadata formats, metadata keys, and metadata key spaces supported by AVFoundation are available in and . To load values of a metadata item when you access them for the first time, use the methods from the protocol. The class and other classes in turn provide their metadata as needed so that you can obtain objects from those arrays without incurring overhead for items you don’t inspect. To filter arrays of metadata items, you use the methods of this class. For example, you can filter by key and key space, by locale, and by preferred language.


// A metadata item for an audiovisual asset or one of its tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem
type MetadataItem struct {
	objectivec.Object
}

// MetadataItemFrom constructs a [MetadataItem] from an unsafe.Pointer.
//
// A metadata item for an audiovisual asset or one of its tracks.
func MetadataItemFrom(ptr unsafe.Pointer) MetadataItem {
	return MetadataItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MetadataItemClass) Alloc() MetadataItem {
	rv := objc.Send[MetadataItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MetadataItemClass) New() MetadataItem {
	rv := objc.Send[MetadataItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataItem) Init() MetadataItem {
	rv := objc.Send[MetadataItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataItem) Autorelease() MetadataItem {
	rv := objc.Send[MetadataItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataItem creates a new MetadataItem instance.
func NewMetadataItem() MetadataItem {
	return getMetadataItemClass().New()
}



// The metadata items an asset contains for common metadata identifiers that provide a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/commonmetadata
func (m_ MetadataItem) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](m_.ID, objc.Sel("commonMetadata"))
	return rv
}


// The metadata items an asset contains for common metadata identifiers that provide a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/commonmetadata
func (m_ MetadataItem) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommonMetadata:"), value)
}


// An array of metadata items for all metadata identifiers for which a value is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/metadata
func (m_ MetadataItem) Metadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](m_.ID, objc.Sel("metadata"))
	return rv
}


// An array of metadata items for all metadata identifiers for which a value is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/metadata
func (m_ MetadataItem) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadata:"), value)
}


// The common key of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/commonkey
func (m_ MetadataItem) CommonKey() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("commonKey"))
	return rv
}


// The common key of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/commonkey
func (m_ MetadataItem) SetCommonKey(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommonKey:"), value)
}


// The data type of the metadata item’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/datatype
func (m_ MetadataItem) DataType() string {
	rv := objc.Send[string](m_.ID, objc.Sel("dataType"))
	return rv
}


// The data type of the metadata item’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/datatype
func (m_ MetadataItem) SetDataType(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataType:"), objc.String(value))
}


// The value of the metadata item as a data value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/datavalue
func (m_ MetadataItem) DataValue() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("dataValue"))
	return rv
}


// The value of the metadata item as a data value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/datavalue
func (m_ MetadataItem) SetDataValue(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataValue:"), value)
}


// The value of the metadata item as a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/datevalue
func (m_ MetadataItem) DateValue() foundation.Date {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("dateValue"))
	return rv
}


// The value of the metadata item as a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/datevalue
func (m_ MetadataItem) SetDateValue(value foundation.Date) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDateValue:"), value)
}


// The duration of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/duration
func (m_ MetadataItem) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("duration"))
	return rv
}


// The duration of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/duration
func (m_ MetadataItem) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// The IETF BCP 47 (RFC 4646) language identifier of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/extendedlanguagetag
func (m_ MetadataItem) ExtendedLanguageTag() string {
	rv := objc.Send[string](m_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// The IETF BCP 47 (RFC 4646) language identifier of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/extendedlanguagetag
func (m_ MetadataItem) SetExtendedLanguageTag(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedLanguageTag:"), objc.String(value))
}


// A dictionary of additional attributes for a metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/extraattributes
func (m_ MetadataItem) ExtraAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("extraAttributes"))
	return rv
}


// A dictionary of additional attributes for a metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/extraattributes
func (m_ MetadataItem) SetExtraAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtraAttributes:"), value)
}


// An identifier for a metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/identifier
func (m_ MetadataItem) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("identifier"))
	return rv
}


// An identifier for a metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/identifier
func (m_ MetadataItem) SetIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), value)
}


// The key of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/key
func (m_ MetadataItem) Key() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("key"))
	return rv
}


// The key of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/key
func (m_ MetadataItem) SetKey(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKey:"), value)
}


// The key space for the metadata item’s key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/keyspace
func (m_ MetadataItem) KeySpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("keySpace"))
	return rv
}


// The key space for the metadata item’s key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/keyspace
func (m_ MetadataItem) SetKeySpace(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeySpace:"), value)
}


// The locale of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/locale
func (m_ MetadataItem) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](m_.ID, objc.Sel("locale"))
	return rv
}


// The locale of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/locale
func (m_ MetadataItem) SetLocale(value foundation.Locale) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocale:"), value)
}


// The value of the metadata item as a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/numbervalue
func (m_ MetadataItem) NumberValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("numberValue"))
	return rv
}


// The value of the metadata item as a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/numbervalue
func (m_ MetadataItem) SetNumberValue(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberValue:"), value)
}


// The start date of the timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/startdate
func (m_ MetadataItem) StartDate() foundation.Date {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("startDate"))
	return rv
}


// The start date of the timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/startdate
func (m_ MetadataItem) SetStartDate(value foundation.Date) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartDate:"), value)
}


// The value of the metadata item as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/stringvalue
func (m_ MetadataItem) StringValue() string {
	rv := objc.Send[string](m_.ID, objc.Sel("stringValue"))
	return rv
}


// The value of the metadata item as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/stringvalue
func (m_ MetadataItem) SetStringValue(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStringValue:"), objc.String(value))
}


// The timestamp of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/time
func (m_ MetadataItem) Time() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("time"))
	return rv
}


// The timestamp of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/time
func (m_ MetadataItem) SetTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTime:"), value)
}


// The value of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/value
func (m_ MetadataItem) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// The value of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadataitem/value
func (m_ MetadataItem) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



