// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	

	// properties:
	CommonKey() MetadataKey /* typedef */
	DataType() objc.IObject /* cross-framework: NSString */
	DataValue() objc.IObject /* cross-framework: NSData */
	DateValue() objc.IObject /* cross-framework: NSDate */
	Duration() objc.IObject /* cross-framework: Time */
	ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */
	ExtraAttributes() foundation.IDictionary
	Identifier() MetadataIdentifier /* typedef */
	Key() unsafe.Pointer
	KeySpace() MetadataKeySpace /* typedef */
	Locale() foundation.Locale
	NumberValue() objc.IObject /* cross-framework: NSNumber */
	StartDate() objc.IObject /* cross-framework: NSDate */
	StringValue() objc.IObject /* cross-framework: NSString */
	Time() objc.IObject /* cross-framework: Time */
	Value() unsafe.Pointer
	CommonMetadata() IAVMetadataItem
	SetCommonMetadata(value IAVMetadataItem)
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)


	

	// methods:
	LoadValuesAsynchronouslyForKeysCompletionHandler(keys []string, handler unsafe.Pointer)
	StatusOfValueForKeyError(key objc.IObject /* cross-framework: NSString */, outError objectivec.IObject) KeyValueStatus


}





// Alloc allocates a new instance without initialization.
func (mc _MetadataItemClass) Alloc() MetadataItem {
	rv := objc.Send[MetadataItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates a metadata item whose value loads on an on-demand basis only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/init(propertiesOfMetadataItem:valueLoadingHandler:)
func NewMetadataItemWithPropertiesOfMetadataItemValueLoadingHandler(metadataItem IAVMetadataItem, handler unsafe.Pointer) MetadataItem {
	rv := objc.Send[MetadataItem](objc.ID(getMetadataItemClass().class), objc.Sel("metadataItemWithPropertiesOfMetadataItem:valueLoadingHandler:"), metadataItem, handler)
	return rv
}







// Returns a metadata identifier for the specified key and key space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/identifier(forKey:keySpace:)
func (mc _MetadataItemClass) IdentifierForKeyKeySpace(key objc.IObject, keySpace MetadataKeySpace /* typedef */) MetadataIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](objc.ID(mc.class), objc.Sel("identifierForKey:keySpace:"), key, keySpace)
	return rv
}


// Creates a metadata item whose value loads on an on-demand basis only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/init(propertiesOfMetadataItem:valueLoadingHandler:)
func (mc _MetadataItemClass) MetadataItemWithPropertiesOfMetadataItemValueLoadingHandler(metadataItem IAVMetadataItem, handler unsafe.Pointer) IMetadataItem {
	rv := objc.Send[MetadataItem](objc.ID(mc.class), objc.Sel("metadataItemWithPropertiesOfMetadataItem:valueLoadingHandler:"), metadataItem, handler)
	return rv
}


// Returns a metadata key for the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/key(forIdentifier:)
func (mc _MetadataItemClass) KeyForIdentifier(identifier MetadataIdentifier /* typedef */) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("keyForIdentifier:"), identifier)
	return rv
}


// Returns a metadata key space for the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/keySpace(forIdentifier:)
func (mc _MetadataItemClass) KeySpaceForIdentifier(identifier MetadataIdentifier /* typedef */) MetadataKeySpace /* typedef */ {
	rv := objc.Send[foundation.NSString](objc.ID(mc.class), objc.Sel("keySpaceForIdentifier:"), identifier)
	return rv
}


// Returns metadata items whose locales match one of the specified language identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/metadataItems(from:filteredAndSortedAccordingToPreferredLanguages:)
func (mc _MetadataItemClass) MetadataItemsFromArrayFilteredAndSortedAccordingToPreferredLanguages(metadataItems []MetadataItem, preferredLanguages []string) []MetadataItem {
	rv := objc.Send[[]MetadataItem](objc.ID(mc.class), objc.Sel("metadataItemsFromArray:filteredAndSortedAccordingToPreferredLanguages:"), metadataItems, preferredLanguages)
	return rv
}


// Returns filtered metadata items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/metadataItems(from:filteredBy:)
func (mc _MetadataItemClass) MetadataItemsFromArrayFilteredByMetadataItemFilter(metadataItems []MetadataItem, metadataItemFilter IAVMetadataItemFilter) []MetadataItem {
	rv := objc.Send[[]MetadataItem](objc.ID(mc.class), objc.Sel("metadataItemsFromArray:filteredByMetadataItemFilter:"), metadataItems, metadataItemFilter)
	return rv
}


// Returns metadata items for the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/metadataItems(from:filteredByIdentifier:)
func (mc _MetadataItemClass) MetadataItemsFromArrayFilteredByIdentifier(metadataItems []MetadataItem, identifier MetadataIdentifier /* typedef */) []MetadataItem {
	rv := objc.Send[[]MetadataItem](objc.ID(mc.class), objc.Sel("metadataItemsFromArray:filteredByIdentifier:"), metadataItems, identifier)
	return rv
}


// Returns metadata items that match a specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/metadataItems(from:with:)
func (mc _MetadataItemClass) MetadataItemsFromArrayWithLocale(metadataItems []MetadataItem, locale foundation.Locale) []MetadataItem {
	rv := objc.Send[[]MetadataItem](objc.ID(mc.class), objc.Sel("metadataItemsFromArray:withLocale:"), metadataItems, locale)
	return rv
}


// Returns metadata items that match a specified key or key space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/metadataItems(from:withKey:keySpace:)
func (mc _MetadataItemClass) MetadataItemsFromArrayWithKeyKeySpace(metadataItems []MetadataItem, key objc.IObject, keySpace MetadataKeySpace /* typedef */) []MetadataItem {
	rv := objc.Send[[]MetadataItem](objc.ID(mc.class), objc.Sel("metadataItemsFromArray:withKey:keySpace:"), metadataItems, key, keySpace)
	return rv
}












// Tells the object to load the values of any of the specified keys that aren’t already loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/loadValuesAsynchronouslyForKeys:completionHandler:
func (m_ MetadataItem) LoadValuesAsynchronouslyForKeysCompletionHandler(keys []string, handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("loadValuesAsynchronouslyForKeys:completionHandler:"), keys, handler)
}


// Reports whether the value for a given key is immediately available without blocking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/statusOfValueForKey:error:
func (m_ MetadataItem) StatusOfValueForKeyError(key objc.IObject /* cross-framework: NSString */, outError objectivec.IObject) KeyValueStatus {
	rv := objc.Send[KeyValueStatus](m_.ID, objc.Sel("statusOfValueForKey:error:"), key, outError)
	return rv
}







// The common key of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/commonKey
func (m_ MetadataItem) CommonKey() MetadataKey /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("commonKey"))
	return rv
}


// The data type of the metadata item’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/dataType
func (m_ MetadataItem) DataType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("dataType"))
	return rv
}


// The value of the metadata item as a data value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/dataValue
func (m_ MetadataItem) DataValue() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("dataValue"))
	return rv
}


// The value of the metadata item as a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/dateValue
func (m_ MetadataItem) DateValue() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("dateValue"))
	return rv
}


// The duration of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/duration
func (m_ MetadataItem) Duration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](m_.ID, objc.Sel("duration"))
	return rv
}


// The IETF BCP 47 (RFC 4646) language identifier of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/extendedLanguageTag
func (m_ MetadataItem) ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// A dictionary of additional attributes for a metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/extraAttributes
func (m_ MetadataItem) ExtraAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("extraAttributes"))
	return rv
}


// An identifier for a metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/identifier
func (m_ MetadataItem) Identifier() MetadataIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("identifier"))
	return rv
}


// The key of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/key
func (m_ MetadataItem) Key() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("key"))
	return rv
}


// The key space for the metadata item’s key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/keySpace
func (m_ MetadataItem) KeySpace() MetadataKeySpace /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("keySpace"))
	return rv
}


// The locale of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/locale
func (m_ MetadataItem) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](m_.ID, objc.Sel("locale"))
	return rv
}


// The value of the metadata item as a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/numberValue
func (m_ MetadataItem) NumberValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("numberValue"))
	return rv
}


// The start date of the timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/startDate
func (m_ MetadataItem) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("startDate"))
	return rv
}


// The value of the metadata item as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/stringValue
func (m_ MetadataItem) StringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("stringValue"))
	return rv
}


// The timestamp of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/time
func (m_ MetadataItem) Time() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](m_.ID, objc.Sel("time"))
	return rv
}


// The value of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/value
func (m_ MetadataItem) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
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







