// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ByteCountFormatter] class.
type IByteCountFormatter interface {
	IFormatter
	AllowedUnits() NSByteCountFormatterUnits
	SetAllowedUnits(value NSByteCountFormatterUnits)
	AllowsNonnumericFormatting() bool
	SetAllowsNonnumericFormatting(value bool)
	CountStyle() NSByteCountFormatterCountStyle
	SetCountStyle(value NSByteCountFormatterCountStyle)
	FormattingContext() int
	SetFormattingContext(value int)
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
	StringForObjectValue(obj objectivec.IObject) IString
	StringFromMeasurement(measurement unsafe.Pointer) IString
	StringFromByteCount(byteCount unsafe.Pointer) IString
}

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

// Alloc allocates a new instance without initialization.
func (bc _ByteCountFormatterClass) Alloc() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/string(from:countStyle:)
func (bc _ByteCountFormatterClass) StringFromMeasurementCountStyle(measurement unsafe.Pointer, countStyle NSByteCountFormatterCountStyle) IString {
	rv := objc.Send[String](objc.ID(bc.class), objc.Sel("stringFromMeasurement:countStyle:"), measurement, countStyle)
	return rv
}


// Converts a byte count into the specified string format without creating an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/string(fromByteCount:countStyle:)
func (bc _ByteCountFormatterClass) StringFromByteCountCountStyle(byteCount unsafe.Pointer, countStyle NSByteCountFormatterCountStyle) IString {
	rv := objc.Send[String](objc.ID(bc.class), objc.Sel("stringFromByteCount:countStyle:"), byteCount, countStyle)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/string(for:)
func (b_ ByteCountFormatter) StringForObjectValue(obj objectivec.IObject) IString {
	rv := objc.Send[String](b_.ID, objc.Sel("stringForObjectValue:"), obj)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/string(from:)
func (b_ ByteCountFormatter) StringFromMeasurement(measurement unsafe.Pointer) IString {
	rv := objc.Send[String](b_.ID, objc.Sel("stringFromMeasurement:"), measurement)
	return rv
}


// Converts a byte count into a string without creating an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/string(fromByteCount:)
func (b_ ByteCountFormatter) StringFromByteCount(byteCount unsafe.Pointer) IString {
	rv := objc.Send[String](b_.ID, objc.Sel("stringFromByteCount:"), byteCount)
	return rv
}


// Specify the units that can be used in the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/allowedUnits
func (b_ ByteCountFormatter) AllowedUnits() NSByteCountFormatterUnits {
	rv := objc.Send[ByteCountFormatterUnits](b_.ID, objc.Sel("allowedUnits"))
	return rv
}


// Specify the units that can be used in the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/allowedUnits
func (b_ ByteCountFormatter) SetAllowedUnits(value NSByteCountFormatterUnits) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowedUnits:"), value)
}


// Determines whether to allow more natural display of some values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/allowsNonnumericFormatting
func (b_ ByteCountFormatter) AllowsNonnumericFormatting() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("allowsNonnumericFormatting"))
	return rv
}


// Determines whether to allow more natural display of some values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/allowsNonnumericFormatting
func (b_ ByteCountFormatter) SetAllowsNonnumericFormatting(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAllowsNonnumericFormatting:"), value)
}


// Specify the number of bytes to be used for kilobytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/countStyle-swift.property
func (b_ ByteCountFormatter) CountStyle() NSByteCountFormatterCountStyle {
	rv := objc.Send[ByteCountFormatterCountStyle](b_.ID, objc.Sel("countStyle"))
	return rv
}


// Specify the number of bytes to be used for kilobytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/countStyle-swift.property
func (b_ ByteCountFormatter) SetCountStyle(value NSByteCountFormatterCountStyle) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCountStyle:"), value)
}


// Specify the formatting context for the formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/formattingContext
func (b_ ByteCountFormatter) FormattingContext() int {
	rv := objc.Send[int](b_.ID, objc.Sel("formattingContext"))
	return rv
}


// Specify the formatting context for the formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/formattingContext
func (b_ ByteCountFormatter) SetFormattingContext(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFormattingContext:"), value)
}


// Determines whether to include the number of bytes after the formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesActualByteCount
func (b_ ByteCountFormatter) IncludesActualByteCount() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("includesActualByteCount"))
	return rv
}


// Determines whether to include the number of bytes after the formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesActualByteCount
func (b_ ByteCountFormatter) SetIncludesActualByteCount(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIncludesActualByteCount:"), value)
}


// Determines whether to include the count in the resulting formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesCount
func (b_ ByteCountFormatter) IncludesCount() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("includesCount"))
	return rv
}


// Determines whether to include the count in the resulting formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesCount
func (b_ ByteCountFormatter) SetIncludesCount(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIncludesCount:"), value)
}


// Determines whether to include the units in the resulting formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesUnit
func (b_ ByteCountFormatter) IncludesUnit() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("includesUnit"))
	return rv
}


// Determines whether to include the units in the resulting formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/includesUnit
func (b_ ByteCountFormatter) SetIncludesUnit(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIncludesUnit:"), value)
}


// Determines the display style of the size representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/isAdaptive
func (b_ ByteCountFormatter) Adaptive() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("adaptive"))
	return rv
}


// Determines the display style of the size representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/isAdaptive
func (b_ ByteCountFormatter) SetAdaptive(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAdaptive:"), value)
}


// Determines whether to zero pad fraction digits so a consistent number of characters is displayed in a representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/zeroPadsFractionDigits
func (b_ ByteCountFormatter) ZeroPadsFractionDigits() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("zeroPadsFractionDigits"))
	return rv
}


// Determines whether to zero pad fraction digits so a consistent number of characters is displayed in a representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/zeroPadsFractionDigits
func (b_ ByteCountFormatter) SetZeroPadsFractionDigits(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setZeroPadsFractionDigits:"), value)
}


// Determines the display style of the size representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/isadaptive
func (b_ ByteCountFormatter) IsAdaptive() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isAdaptive"))
	return rv
}


// Determines the display style of the size representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bytecountformatter/isadaptive
func (b_ ByteCountFormatter) SetIsAdaptive(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsAdaptive:"), value)
}



