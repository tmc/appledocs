// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMutableMetadataItem */


/* debug [class_header]: Header for AVMutableMetadataItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableMetadataItem */
// An interface definition for the [MutableMetadataItem] class.
type IMutableMetadataItem interface {
	IMetadataItem
	
/* debug [class_interface_properties]: Properties for MutableMetadataItem */
	// properties:
	DataType() objc.IObject /* cross-framework: NSString */
	SetDataType(value objc.IObject /* cross-framework: NSString */)
	DataValue() objectivec.IObject
	SetDataValue(value objectivec.IObject)
	DateValue() objectivec.IObject
	SetDateValue(value objectivec.IObject)
	Duration() objc.IObject /* cross-framework: Time */
	SetDuration(value objc.IObject /* cross-framework: Time */)
	ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */
	SetExtendedLanguageTag(value objc.IObject /* cross-framework: NSString */)
	ExtraAttributes() foundation.IDictionary
	SetExtraAttributes(value foundation.IDictionary)
	Identifier() MetadataIdentifier /* typedef */
	SetIdentifier(value MetadataIdentifier /* typedef */)
	Key() unsafe.Pointer
	SetKey(value unsafe.Pointer)
	KeySpace() MetadataKeySpace /* typedef */
	SetKeySpace(value MetadataKeySpace /* typedef */)
	Locale() foundation.Locale
	SetLocale(value foundation.Locale)
	NumberValue() objc.IObject /* cross-framework: NSNumber */
	SetNumberValue(value objc.IObject /* cross-framework: NSNumber */)
	StartDate() objc.IObject /* cross-framework: NSDate */
	SetStartDate(value objc.IObject /* cross-framework: NSDate */)
	StringValue() objectivec.IObject
	SetStringValue(value objectivec.IObject)
	Time() objc.IObject /* cross-framework: Time */
	SetTime(value objc.IObject /* cross-framework: Time */)
	Value() unsafe.Pointer
	SetValue(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableMetadataItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableMetadataItem */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableMetadataItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableMetadataItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableMetadataItem */

// Returns a new mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/metadataItem
func (mc _MutableMetadataItemClass) MetadataItem() IMutableMetadataItem {
	rv := objc.Send[MutableMetadataItem](objc.ID(mc.class), objc.Sel("metadataItem"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MetadataItem) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableMetadataItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableMetadataItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableMetadataItem */

// The data type of the metadata item’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/dataType
func (m_ MutableMetadataItem) DataType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// The data type of the metadata item’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/dataType
func (m_ MutableMetadataItem) SetDataType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataType:"), value)
}/* debug [instance_properties/setter]: dataType */


// The value of the metadata item as a data value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/dataValue
func (m_ MutableMetadataItem) DataValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("dataValue"))
	return rv
}/* debug [instance_properties/getter]: dataValue */


// The value of the metadata item as a data value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/dataValue
func (m_ MutableMetadataItem) SetDataValue(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataValue:"), value)
}/* debug [instance_properties/setter]: dataValue */


// The value of the metadata item as a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/dateValue
func (m_ MutableMetadataItem) DateValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("dateValue"))
	return rv
}/* debug [instance_properties/getter]: dateValue */


// The value of the metadata item as a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/dateValue
func (m_ MutableMetadataItem) SetDateValue(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDateValue:"), value)
}/* debug [instance_properties/setter]: dateValue */


// The duration of a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/duration
func (m_ MutableMetadataItem) Duration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The duration of a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/duration
func (m_ MutableMetadataItem) SetDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// The IETF BCP 47 (RFC 4646) language identifier of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/extendedLanguageTag
func (m_ MutableMetadataItem) ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}/* debug [instance_properties/getter]: extendedLanguageTag */


// The IETF BCP 47 (RFC 4646) language identifier of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/extendedLanguageTag
func (m_ MutableMetadataItem) SetExtendedLanguageTag(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedLanguageTag:"), value)
}/* debug [instance_properties/setter]: extendedLanguageTag */


// A dictionary of additional attributes for a metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/extraAttributes
func (m_ MutableMetadataItem) ExtraAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("extraAttributes"))
	return rv
}/* debug [instance_properties/getter]: extraAttributes */


// A dictionary of additional attributes for a metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/extraAttributes
func (m_ MutableMetadataItem) SetExtraAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtraAttributes:"), value)
}/* debug [instance_properties/setter]: extraAttributes */


// Indicates the identifier of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/identifier
func (m_ MutableMetadataItem) Identifier() MetadataIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// Indicates the identifier of the metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/identifier
func (m_ MutableMetadataItem) SetIdentifier(value MetadataIdentifier /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// The key for a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/key
func (m_ MutableMetadataItem) Key() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("key"))
	return rv
}/* debug [instance_properties/getter]: key */


// The key for a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/key
func (m_ MutableMetadataItem) SetKey(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKey:"), value)
}/* debug [instance_properties/setter]: key */


// The key space of the metadata item’s key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/keySpace
func (m_ MutableMetadataItem) KeySpace() MetadataKeySpace /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("keySpace"))
	return rv
}/* debug [instance_properties/getter]: keySpace */


// The key space of the metadata item’s key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/keySpace
func (m_ MutableMetadataItem) SetKeySpace(value MetadataKeySpace /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeySpace:"), value)
}/* debug [instance_properties/setter]: keySpace */


// The locale for a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/locale
func (m_ MutableMetadataItem) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](m_.ID, objc.Sel("locale"))
	return rv
}/* debug [instance_properties/getter]: locale */


// The locale for a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/locale
func (m_ MutableMetadataItem) SetLocale(value foundation.Locale) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocale:"), value)
}/* debug [instance_properties/setter]: locale */


// The value of the metadata item as a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/numberValue
func (m_ MutableMetadataItem) NumberValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("numberValue"))
	return rv
}/* debug [instance_properties/getter]: numberValue */


// The value of the metadata item as a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/numberValue
func (m_ MutableMetadataItem) SetNumberValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberValue:"), value)
}/* debug [instance_properties/setter]: numberValue */


// The start date of the timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/startDate
func (m_ MutableMetadataItem) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */


// The start date of the timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/startDate
func (m_ MutableMetadataItem) SetStartDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartDate:"), value)
}/* debug [instance_properties/setter]: startDate */


// The value of the metadata item as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/stringValue
func (m_ MutableMetadataItem) StringValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("stringValue"))
	return rv
}/* debug [instance_properties/getter]: stringValue */


// The value of the metadata item as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/stringValue
func (m_ MutableMetadataItem) SetStringValue(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStringValue:"), value)
}/* debug [instance_properties/setter]: stringValue */


// The timestamp for a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/time
func (m_ MutableMetadataItem) Time() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](m_.ID, objc.Sel("time"))
	return rv
}/* debug [instance_properties/getter]: time */


// The timestamp for a mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/time
func (m_ MutableMetadataItem) SetTime(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTime:"), value)
}/* debug [instance_properties/setter]: time */


// The value for the mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/value
func (m_ MutableMetadataItem) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The value for the mutable metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableMetadataItem/value
func (m_ MutableMetadataItem) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMutableMetadataItem */



