// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RelativeDateTimeFormatter] class.
var (
	RelativeDateTimeFormatterClass     _RelativeDateTimeFormatterClass
	RelativeDateTimeFormatterClassOnce sync.Once
)

func getRelativeDateTimeFormatterClass() _RelativeDateTimeFormatterClass {
	RelativeDateTimeFormatterClassOnce.Do(func() {
		RelativeDateTimeFormatterClass = _RelativeDateTimeFormatterClass{objc.GetClass("NSRelativeDateTimeFormatter")}
	})
	return RelativeDateTimeFormatterClass
}

type _RelativeDateTimeFormatterClass struct {
	class objc.Class
}

// An interface definition for the [RelativeDateTimeFormatter] class.
type IRelativeDateTimeFormatter interface {
	IFormatter
	Calendar() Calendar
	SetCalendar(value ICalendar)
	DateTimeStyle() unsafe.Pointer
	SetDateTimeStyle(value unsafe.Pointer)
	FormattingContext() int
	SetFormattingContext(value int)
	Locale() Locale
	SetLocale(value ILocale)
	UnitsStyle() unsafe.Pointer
	SetUnitsStyle(value unsafe.Pointer)
}

// A formatter that creates locale-aware string representations of a relative date or time.
//
// Use the strings that the formatter produces, such as “1 hour ago”, “in 2 weeks”, “yesterday”, and “tomorrow” as standalone strings. Embedding them in other strings may not be grammatically correct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter
type RelativeDateTimeFormatter struct {
	Formatter
}

// RelativeDateTimeFormatterFrom constructs a [RelativeDateTimeFormatter] from an unsafe.Pointer.
//
// A formatter that creates locale-aware string representations of a relative date or time.
func RelativeDateTimeFormatterFrom(ptr unsafe.Pointer) RelativeDateTimeFormatter {
	return RelativeDateTimeFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RelativeDateTimeFormatterClass) Alloc() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RelativeDateTimeFormatterClass) New() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RelativeDateTimeFormatter) Init() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RelativeDateTimeFormatter) Autorelease() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRelativeDateTimeFormatter creates a new RelativeDateTimeFormatter instance.
func NewRelativeDateTimeFormatter() RelativeDateTimeFormatter {
	return getRelativeDateTimeFormatterClass().New()
}


// The calendar to use for formatting values that don’t have an inherent calendar of their own.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/relativedatetimeformatter/calendar
func (r_ RelativeDateTimeFormatter) Calendar() Calendar {
	rv := objc.Send[Calendar](r_.ID, objc.Sel("calendar"))
	return rv
}


// SetCalendar sets the value of the calendar property.
// The calendar to use for formatting values that don’t have an inherent calendar of their own.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/relativedatetimeformatter/calendar
func (r_ RelativeDateTimeFormatter) SetCalendar(value ICalendar) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCalendar:"), value)
}

// The style to use when describing a relative date, for example “yesterday” or “1 day ago”.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/relativedatetimeformatter/datetimestyle-swift.property
func (r_ RelativeDateTimeFormatter) DateTimeStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("dateTimeStyle"))
	return rv
}


// SetDateTimeStyle sets the value of the dateTimeStyle property.
// The style to use when describing a relative date, for example “yesterday” or “1 day ago”.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/relativedatetimeformatter/datetimestyle-swift.property
func (r_ RelativeDateTimeFormatter) SetDateTimeStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDateTimeStyle:"), value)
}

// A description of where the formatted string will appear, allowing the formatter to capitalize the output appropriately.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/relativedatetimeformatter/formattingcontext
func (r_ RelativeDateTimeFormatter) FormattingContext() int {
	rv := objc.Send[int](r_.ID, objc.Sel("formattingContext"))
	return rv
}


// SetFormattingContext sets the value of the formattingContext property.
// A description of where the formatted string will appear, allowing the formatter to capitalize the output appropriately.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/relativedatetimeformatter/formattingcontext
func (r_ RelativeDateTimeFormatter) SetFormattingContext(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFormattingContext:"), value)
}

// The locale to use when formatting the date.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/relativedatetimeformatter/locale
func (r_ RelativeDateTimeFormatter) Locale() Locale {
	rv := objc.Send[Locale](r_.ID, objc.Sel("locale"))
	return rv
}


// SetLocale sets the value of the locale property.
// The locale to use when formatting the date.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/relativedatetimeformatter/locale
func (r_ RelativeDateTimeFormatter) SetLocale(value ILocale) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLocale:"), value)
}

// The style to use when formatting the quantity or the name of the unit, such as “1 day ago” or “one day ago”.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/relativedatetimeformatter/unitsstyle-swift.property
func (r_ RelativeDateTimeFormatter) UnitsStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("unitsStyle"))
	return rv
}


// SetUnitsStyle sets the value of the unitsStyle property.
// The style to use when formatting the quantity or the name of the unit, such as “1 day ago” or “one day ago”.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/relativedatetimeformatter/unitsstyle-swift.property
func (r_ RelativeDateTimeFormatter) SetUnitsStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setUnitsStyle:"), value)
}



