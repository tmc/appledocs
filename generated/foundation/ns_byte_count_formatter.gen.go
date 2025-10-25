// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSByteCountFormatter */


/* debug [class_header]: Header for NSByteCountFormatter */
// The class instance for the [ByteCountFormatter] class.
var (
	ByteCountFormatterClass     _ByteCountFormatterClass
	ByteCountFormatterClassOnce sync.Once
)

func getByteCountFormatterClass() _ByteCountFormatterClass {
	ByteCountFormatterClassOnce.Do(func() {
		ByteCountFormatterClass = _ByteCountFormatterClass{objc.GetClass("NSByteCountFormatter")}
	})
	return ByteCountFormatterClass
}

type _ByteCountFormatterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ByteCountFormatter */
// An interface definition for the [ByteCountFormatter] class.
type IByteCountFormatter interface {
	IFormatter
	
/* debug [class_interface_properties]: Properties for ByteCountFormatter */
	// properties:
	AllowedUnits() ByteCountFormatterUnits
	SetAllowedUnits(value ByteCountFormatterUnits)
	AllowsNonnumericFormatting() bool
	SetAllowsNonnumericFormatting(value bool)
	CountStyle() ByteCountFormatterCountStyle
	SetCountStyle(value ByteCountFormatterCountStyle)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	IncludesActualByteCount() bool
	SetIncludesActualByteCount(value bool)
	IncludesCount() bool
	SetIncludesCount(value bool)
	IncludesUnit() bool
	SetIncludesUnit(value bool)
	Adaptive() bool
	SetAdaptive(value bool)
	ZeroPadsFractionDigits() bool
	SetZeroPadsFractionDigits(value bool)
	IsAdaptive() bool
	SetIsAdaptive(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ByteCountFormatter */
	// methods:
	StringForObjectValue(obj objc.IObject) IString
	StringFromMeasurement(measurement unsafe.Pointer) IString
	StringFromByteCount(byteCount objectivec.IObject) IString
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ByteCountFormatter */
// Alloc allocates a new instance without initialization.
func (bc _ByteCountFormatterClass) Alloc() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _ByteCountFormatterClass) New() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ ByteCountFormatter) Init() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ ByteCountFormatter) Autorelease() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewByteCountFormatter creates a new ByteCountFormatter instance.
func NewByteCountFormatter() ByteCountFormatter {
	return getByteCountFormatterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ByteCountFormatter */
// A formatter that converts a byte count value into a localized description that is formatted with the appropriate byte modifier (KB, MB, GB and so on).


// A formatter that converts a byte count value into a localized description that is formatted with the appropriate byte modifier (KB, MB, GB and so on).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter
type ByteCountFormatter struct {
	Formatter
}

// ByteCountFormatterFrom constructs a [ByteCountFormatter] from an unsafe.Pointer.
//
// A formatter that converts a byte count value into a localized description that is formatted with the appropriate byte modifier (KB, MB, GB and so on).
func ByteCountFormatterFrom(ptr unsafe.Pointer) ByteCountFormatter {
	return ByteCountFormatter{
		Formatter: FormatterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ByteCountFormatter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ByteCountFormatter */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/string(from:countStyle:)
func (bc _ByteCountFormatterClass) StringFromMeasurementCountStyle(measurement unsafe.Pointer, countStyle ByteCountFormatterCountStyle) IString {
	rv := objc.Send[String](objc.ID(bc.class), objc.Sel("stringFromMeasurement:countStyle:"), measurement, countStyle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StringFromMeasurementCountStyle) */


// Converts a byte count into the specified string format without creating an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/string(fromByteCount:countStyle:)
func (bc _ByteCountFormatterClass) StringFromByteCountCountStyle(byteCount objectivec.IObject, countStyle ByteCountFormatterCountStyle) IString {
	rv := objc.Send[String](objc.ID(bc.class), objc.Sel("stringFromByteCount:countStyle:"), byteCount, countStyle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StringFromByteCountCountStyle) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ByteCountFormatter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ByteCountFormatter */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/string(for:)
func (b_ ByteCountFormatter) StringForObjectValue(obj objc.IObject) IString {
	rv := objc.Send[String](b_.ID, objc.Sel("stringForObjectValue:"), obj)
	return rv
}/* debug [instance_methods/method]: StringForObjectValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/string(from:)
func (b_ ByteCountFormatter) StringFromMeasurement(measurement unsafe.Pointer) IString {
	rv := objc.Send[String](b_.ID, objc.Sel("stringFromMeasurement:"), measurement)
	return rv
}/* debug [instance_methods/method]: StringFromMeasurement */


// Converts a byte count into a string without creating an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/string(fromByteCount:)
func (b_ ByteCountFormatter) StringFromByteCount(byteCount objectivec.IObject) IString {
	rv := objc.Send[String](b_.ID, objc.Sel("stringFromByteCount:"), byteCount)
	return rv
}/* debug [instance_methods/method]: StringFromByteCount */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ByteCountFormatter */

// Specify the units that can be used in the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/allowedUnits
func (b_ ByteCountFormatter) AllowedUnits() ByteCountFormatterUnits {
	rv := objc.Send[ByteCountFormatterUnits](b_.ID, objc.Sel("allowedUnits"))
	return rv
}/* debug [instance_properties/getter]: allowedUnits */


// Specify the units that can be used in the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/allowedUnits
func (b_ ByteCountFormatter) SetAllowedUnits(value ByteCountFormatterUnits) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowedUnits:"), value)
}/* debug [instance_properties/setter]: allowedUnits */


// Determines whether to allow more natural display of some values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/allowsNonnumericFormatting
func (b_ ByteCountFormatter) AllowsNonnumericFormatting() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("allowsNonnumericFormatting"))
	return rv
}/* debug [instance_properties/getter]: allowsNonnumericFormatting */


// Determines whether to allow more natural display of some values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/allowsNonnumericFormatting
func (b_ ByteCountFormatter) SetAllowsNonnumericFormatting(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowsNonnumericFormatting:"), value)
}/* debug [instance_properties/setter]: allowsNonnumericFormatting */


// Specify the number of bytes to be used for kilobytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/countStyle-swift.property
func (b_ ByteCountFormatter) CountStyle() ByteCountFormatterCountStyle {
	rv := objc.Send[ByteCountFormatterCountStyle](b_.ID, objc.Sel("countStyle"))
	return rv
}/* debug [instance_properties/getter]: countStyle */


// Specify the number of bytes to be used for kilobytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/countStyle-swift.property
func (b_ ByteCountFormatter) SetCountStyle(value ByteCountFormatterCountStyle) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCountStyle:"), value)
}/* debug [instance_properties/setter]: countStyle */


// Specify the formatting context for the formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/formattingContext
func (b_ ByteCountFormatter) FormattingContext() FormattingContext {
	rv := objc.Send[FormattingContext](b_.ID, objc.Sel("formattingContext"))
	return rv
}/* debug [instance_properties/getter]: formattingContext */


// Specify the formatting context for the formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/formattingContext
func (b_ ByteCountFormatter) SetFormattingContext(value FormattingContext) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFormattingContext:"), value)
}/* debug [instance_properties/setter]: formattingContext */


// Determines whether to include the number of bytes after the formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesActualByteCount
func (b_ ByteCountFormatter) IncludesActualByteCount() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("includesActualByteCount"))
	return rv
}/* debug [instance_properties/getter]: includesActualByteCount */


// Determines whether to include the number of bytes after the formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesActualByteCount
func (b_ ByteCountFormatter) SetIncludesActualByteCount(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIncludesActualByteCount:"), value)
}/* debug [instance_properties/setter]: includesActualByteCount */


// Determines whether to include the count in the resulting formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesCount
func (b_ ByteCountFormatter) IncludesCount() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("includesCount"))
	return rv
}/* debug [instance_properties/getter]: includesCount */


// Determines whether to include the count in the resulting formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesCount
func (b_ ByteCountFormatter) SetIncludesCount(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIncludesCount:"), value)
}/* debug [instance_properties/setter]: includesCount */


// Determines whether to include the units in the resulting formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesUnit
func (b_ ByteCountFormatter) IncludesUnit() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("includesUnit"))
	return rv
}/* debug [instance_properties/getter]: includesUnit */


// Determines whether to include the units in the resulting formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesUnit
func (b_ ByteCountFormatter) SetIncludesUnit(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIncludesUnit:"), value)
}/* debug [instance_properties/setter]: includesUnit */


// Determines the display style of the size representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/isAdaptive
func (b_ ByteCountFormatter) Adaptive() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("adaptive"))
	return rv
}/* debug [instance_properties/getter]: adaptive */


// Determines the display style of the size representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/isAdaptive
func (b_ ByteCountFormatter) SetAdaptive(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAdaptive:"), value)
}/* debug [instance_properties/setter]: adaptive */


// Determines whether to zero pad fraction digits so a consistent number of characters is displayed in a representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/zeroPadsFractionDigits
func (b_ ByteCountFormatter) ZeroPadsFractionDigits() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("zeroPadsFractionDigits"))
	return rv
}/* debug [instance_properties/getter]: zeroPadsFractionDigits */


// Determines whether to zero pad fraction digits so a consistent number of characters is displayed in a representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/zeroPadsFractionDigits
func (b_ ByteCountFormatter) SetZeroPadsFractionDigits(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setZeroPadsFractionDigits:"), value)
}/* debug [instance_properties/setter]: zeroPadsFractionDigits */


// Determines the display style of the size representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/isadaptive
func (b_ ByteCountFormatter) IsAdaptive() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isAdaptive"))
	return rv
}/* debug [instance_properties/getter]: isAdaptive */


// Determines the display style of the size representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/isadaptive
func (b_ ByteCountFormatter) SetIsAdaptive(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsAdaptive:"), value)
}/* debug [instance_properties/setter]: isAdaptive */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSByteCountFormatter */



