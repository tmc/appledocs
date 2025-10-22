// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)

// The class instance for the [ISO8601DateFormatter] class.
var (
	ISO8601DateFormatterClass     _ISO8601DateFormatterClass
	ISO8601DateFormatterClassOnce sync.Once
)

func getISO8601DateFormatterClass() _ISO8601DateFormatterClass {
	ISO8601DateFormatterClassOnce.Do(func() {
		ISO8601DateFormatterClass = _ISO8601DateFormatterClass{objc.GetClass("NSISO8601DateFormatter")}
	})
	return ISO8601DateFormatterClass
}

type _ISO8601DateFormatterClass struct {
	class objc.Class
}

// An interface definition for the [ISO8601DateFormatter] class.
type IISO8601DateFormatter interface {
	IFormatter
}

// A formatter that converts between dates and their ISO 8601 string representations.
//
// The class generates and parses string representations of dates following the standard. Use this class to create ISO 8601 representations of dates and create dates from text strings in ISO 8601 format.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter
type ISO8601DateFormatter struct {
	Formatter
}

// ISO8601DateFormatterFrom constructs a [ISO8601DateFormatter] from an unsafe.Pointer.
//
// A formatter that converts between dates and their ISO 8601 string representations.
func ISO8601DateFormatterFrom(ptr unsafe.Pointer) ISO8601DateFormatter {
	return ISO8601DateFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ISO8601DateFormatterClass) Alloc() ISO8601DateFormatter {
	rv := objc.Send[ISO8601DateFormatter](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ISO8601DateFormatterClass) New() ISO8601DateFormatter {
	rv := objc.Send[ISO8601DateFormatter](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ISO8601DateFormatter) Init() ISO8601DateFormatter {
	rv := objc.Send[ISO8601DateFormatter](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ISO8601DateFormatter) Autorelease() ISO8601DateFormatter {
	rv := objc.Send[ISO8601DateFormatter](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewISO8601DateFormatter creates a new ISO8601DateFormatter instance.
func NewISO8601DateFormatter() ISO8601DateFormatter {
	return getISO8601DateFormatterClass().New()
}



// Creates a representation of the specified date with a given time zone and format options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/string(from:timeZone:formatOptions:)
func (ic _ISO8601DateFormatterClass) StringFromDateTimeZoneFormatOptions(date IDate, timeZone ITimeZone, formatOptions corefoundation.ISO8601DateFormatOptions) String {
	rv := objc.Send[String](objc.ID(ic.class), objc.Sel("stringFromDate:timeZone:formatOptions:"), date, timeZone, formatOptions)
	return rv
}

// Options for generating and parsing ISO 8601 date representations. See for possible values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/formatOptions
func (i_ ISO8601DateFormatter) FormatOptions() corefoundation.ISO8601DateFormatOptions {
	rv := objc.Send[corefoundation.ISO8601DateFormatOptions](i_.ID, objc.Sel("formatOptions"))
	return rv
}


// SetFormatOptions sets the value of the formatOptions property.
// Options for generating and parsing ISO 8601 date representations. See for possible values.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/formatOptions
func (i_ ISO8601DateFormatter) SetFormatOptions(value corefoundation.ISO8601DateFormatOptions) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFormatOptions:"), value)
}

// The time zone used to create and parse date representations. When unspecified, GMT is used.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/iso8601dateformatter/timezone
func (i_ ISO8601DateFormatter) TimeZone() TimeZone {
	rv := objc.Send[TimeZone](i_.ID, objc.Sel("timeZone"))
	return rv
}


// SetTimeZone sets the value of the timeZone property.
// The time zone used to create and parse date representations. When unspecified, GMT is used.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/iso8601dateformatter/timezone
func (i_ ISO8601DateFormatter) SetTimeZone(value ITimeZone) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTimeZone:"), value)
}


