// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MutableMetadataItem] class.
var (
	MutableMetadataItemClass     _MutableMetadataItemClass
	MutableMetadataItemClassOnce sync.Once
)

func getMutableMetadataItemClass() _MutableMetadataItemClass {
	MutableMetadataItemClassOnce.Do(func() {
		MutableMetadataItemClass = _MutableMetadataItemClass{objc.GetClass("AVMutableMetadataItem")}
	})
	return MutableMetadataItemClass
}

type _MutableMetadataItemClass struct {
	class objc.Class
}





// An interface definition for the [MutableMetadataItem] class.
type IMutableMetadataItem interface {
	IMetadataItem
	

	// properties:
	DataType() foundation.foundation.INSString
	SetDataType(value foundation.foundation.INSString)
	DataValue() objectivec.IObject
	SetDataValue(value objectivec.IObject)
	DateValue() objectivec.IObject
	SetDateValue(value objectivec.IObject)
	Duration() objectivec.IObject
	SetDuration(value objectivec.IObject)
	ExtendedLanguageTag() foundation.foundation.INSString
	SetExtendedLanguageTag(value foundation.foundation.INSString)
	ExtraAttributes() foundation.IDictionary
	SetExtraAttributes(value foundation.IDictionary)
	Identifier() MetadataIdentifier
	SetIdentifier(value MetadataIdentifier)
	Key() unsafe.Pointer
	SetKey(value unsafe.Pointer)
	KeySpace() MetadataKeySpace
	SetKeySpace(value MetadataKeySpace)
	Locale() foundation.Locale
	SetLocale(value foundation.Locale)
	NumberValue() foundation.foundation.INSNumber
	SetNumberValue(value foundation.foundation.INSNumber)
	StartDate() foundation.foundation.INSDate
	SetStartDate(value foundation.foundation.INSDate)
	StringValue() objectivec.IObject
	SetStringValue(value objectivec.IObject)
	Time() objectivec.IObject
	SetTime(value objectivec.IObject)
	Value() unsafe.Pointer
	SetValue(value unsafe.Pointer)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MutableMetadataItemClass) Alloc() MutableMetadataItem {
	rv := objc.Send[MutableMetadataItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableMetadataItemClass) New() MutableMetadataItem {
	rv := objc.Send[MutableMetadataItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableMetadataItem) Init() MutableMetadataItem {
	rv := objc.Send[MutableMetadataItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableMetadataItem) Autorelease() MutableMetadataItem {
	rv := objc.Send[MutableMetadataItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableMetadataItem creates a new MutableMetadataItem instance.
func NewMutableMetadataItem() MutableMetadataItem {
	return getMutableMetadataItemClass().New()
}





// A mutable metadata item for an audiovisual asset or for one of its tracks.
//
// You can initialize a mutable metadata item from an existing object or with a one or more of the basic properties of a metadata item: a key, a key space, a locale, and a value.


// A mutable metadata item for an audiovisual asset or for one of its tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem
type MutableMetadataItem struct {
	MetadataItem
}

// MutableMetadataItemFrom constructs a [MutableMetadataItem] from an unsafe.Pointer.
//
// A mutable metadata item for an audiovisual asset or for one of its tracks.
func MutableMetadataItemFrom(ptr unsafe.Pointer) MutableMetadataItem {
	return MutableMetadataItem{
		MetadataItem: MetadataItemFrom(ptr),
	}
}










// Returns a new mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/metadataItem
func (mc _MutableMetadataItemClass) MetadataItem() IMutableMetadataItem {
	rv := objc.Send[MutableMetadataItem](objc.ID(mc.class), objc.Sel("metadataItem"))
	return rv
}

















// The data type of the metadata item’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/dataType
func (m_ MutableMetadataItem) DataType() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("dataType"))
	return rv
}


// The data type of the metadata item’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/dataType
func (m_ MutableMetadataItem) SetDataType(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataType:"), value)
}


// The value of the metadata item as a data value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/dataValue
func (m_ MutableMetadataItem) DataValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("dataValue"))
	return rv
}


// The value of the metadata item as a data value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/dataValue
func (m_ MutableMetadataItem) SetDataValue(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataValue:"), value)
}


// The value of the metadata item as a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/dateValue
func (m_ MutableMetadataItem) DateValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("dateValue"))
	return rv
}


// The value of the metadata item as a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/dateValue
func (m_ MutableMetadataItem) SetDateValue(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDateValue:"), value)
}


// The duration of a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/duration
func (m_ MutableMetadataItem) Duration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("duration"))
	return rv
}


// The duration of a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/duration
func (m_ MutableMetadataItem) SetDuration(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// The IETF BCP 47 (RFC 4646) language identifier of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/extendedLanguageTag
func (m_ MutableMetadataItem) ExtendedLanguageTag() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// The IETF BCP 47 (RFC 4646) language identifier of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/extendedLanguageTag
func (m_ MutableMetadataItem) SetExtendedLanguageTag(value foundation.foundation.INSString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedLanguageTag:"), value)
}


// A dictionary of additional attributes for a metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/extraAttributes
func (m_ MutableMetadataItem) ExtraAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("extraAttributes"))
	return rv
}


// A dictionary of additional attributes for a metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/extraAttributes
func (m_ MutableMetadataItem) SetExtraAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtraAttributes:"), value)
}


// Indicates the identifier of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/identifier
func (m_ MutableMetadataItem) Identifier() MetadataIdentifier {
	rv := objc.Send[MetadataIdentifier](m_.ID, objc.Sel("identifier"))
	return rv
}


// Indicates the identifier of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/identifier
func (m_ MutableMetadataItem) SetIdentifier(value MetadataIdentifier) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), value)
}


// The key for a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/key
func (m_ MutableMetadataItem) Key() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("key"))
	return rv
}


// The key for a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/key
func (m_ MutableMetadataItem) SetKey(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKey:"), value)
}


// The key space of the metadata item’s key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/keySpace
func (m_ MutableMetadataItem) KeySpace() MetadataKeySpace {
	rv := objc.Send[MetadataKeySpace](m_.ID, objc.Sel("keySpace"))
	return rv
}


// The key space of the metadata item’s key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/keySpace
func (m_ MutableMetadataItem) SetKeySpace(value MetadataKeySpace) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeySpace:"), value)
}


// The locale for a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/locale
func (m_ MutableMetadataItem) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](m_.ID, objc.Sel("locale"))
	return rv
}


// The locale for a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/locale
func (m_ MutableMetadataItem) SetLocale(value foundation.Locale) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocale:"), value)
}


// The value of the metadata item as a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/numberValue
func (m_ MutableMetadataItem) NumberValue() foundation.foundation.INSNumber {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("numberValue"))
	return rv
}


// The value of the metadata item as a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/numberValue
func (m_ MutableMetadataItem) SetNumberValue(value foundation.foundation.INSNumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberValue:"), value)
}


// The start date of the timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/startDate
func (m_ MutableMetadataItem) StartDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("startDate"))
	return rv
}


// The start date of the timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/startDate
func (m_ MutableMetadataItem) SetStartDate(value foundation.foundation.INSDate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartDate:"), value)
}


// The value of the metadata item as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/stringValue
func (m_ MutableMetadataItem) StringValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("stringValue"))
	return rv
}


// The value of the metadata item as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/stringValue
func (m_ MutableMetadataItem) SetStringValue(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStringValue:"), value)
}


// The timestamp for a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/time
func (m_ MutableMetadataItem) Time() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("time"))
	return rv
}


// The timestamp for a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/time
func (m_ MutableMetadataItem) SetTime(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTime:"), value)
}


// The value for the mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/value
func (m_ MutableMetadataItem) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// The value for the mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/value
func (m_ MutableMetadataItem) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}








