// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DateFormatter] class.
var (
	DateFormatterClass     _DateFormatterClass
	DateFormatterClassOnce sync.Once
)

func getDateFormatterClass() _DateFormatterClass {
	DateFormatterClassOnce.Do(func() {
		DateFormatterClass = _DateFormatterClass{objc.GetClass("NSDateFormatter")}
	})
	return DateFormatterClass
}

type _DateFormatterClass struct {
	class objc.Class
}

// An interface definition for the [DateFormatter] class.
type IDateFormatter interface {
	IFormatter
	DateFromString(string_ string) unsafe.Pointer
	SetLocalizedDateFormatFromTemplate(dateFormatTemplate string)
}

// A formatter that converts between dates and their textual representations.
//
// Instances of create string representations of objects, and convert textual representations of dates and times into objects. For user-visible representations of dates and times, provides a variety of localized presets and configuration options. For fixed format representations of dates and times, you can specify a custom format string. When working with date representations in ISO 8601 format, use instead. To represent an interval between two objects, use instead. To represent a quantity of time specified by an object, use instead.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter
type DateFormatter struct {
	Formatter
}

// DateFormatterFrom constructs a [DateFormatter] from an unsafe.Pointer.
//
// A formatter that converts between dates and their textual representations.
func DateFormatterFrom(ptr unsafe.Pointer) DateFormatter {
	return DateFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DateFormatterClass) Alloc() DateFormatter {
	rv := objc.Send[DateFormatter](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DateFormatterClass) New() DateFormatter {
	rv := objc.Send[DateFormatter](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DateFormatter) Init() DateFormatter {
	rv := objc.Send[DateFormatter](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DateFormatter) Autorelease() DateFormatter {
	rv := objc.Send[DateFormatter](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDateFormatter creates a new DateFormatter instance.
func NewDateFormatter() DateFormatter {
	return getDateFormatterClass().New()
}

// Initializes and returns an instance that uses the OS X 10.0 formatting behavior and the given date format string in its conversions.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateFormatter/initWithDateFormat:allowNaturalLanguage:
func NewDateFormatterWithDateFormatAllowNaturalLanguage(format string, flag bool) DateFormatter {
	instance := getDateFormatterClass().Alloc()
	rv := objc.Send[DateFormatter](instance.ID, objc.Sel("initWithDateFormat:allowNaturalLanguage:"), objc.String(format), flag)
	rv.Autorelease()
	return rv
}

// Returns the default formatting behavior for instances of the class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultFormatterBehavior
func (dc _DateFormatterClass) DefaultFormatterBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("defaultFormatterBehavior"))
	return rv
}

// Returns a date representation of a specified string that the system interprets using the receiver’s current settings.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/date(from:)
func (d_ DateFormatter) DateFromString(string_ string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("dateFromString:"), objc.String(string_))
	return rv
}

// Sets the date format from a template using the specified locale for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/setLocalizedDateFormatFromTemplate(_:)
func (d_ DateFormatter) SetLocalizedDateFormatFromTemplate(dateFormatTemplate string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocalizedDateFormatFromTemplate:"), objc.String(dateFormatTemplate))
}

// The calendar for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/calendar
func (d_ DateFormatter) Calendar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("calendar"))
	return rv
}

// SetCalendar sets the value of the calendar property.
// The calendar for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/calendar
func (d_ DateFormatter) SetCalendar(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCalendar:"), value)
}

// The date format string used by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateFormat
func (d_ DateFormatter) DateFormat() string {
	rv := objc.Send[string](d_.ID, objc.Sel("dateFormat"))
	return rv
}

// SetDateFormat sets the value of the dateFormat property.
// The date format string used by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateFormat
func (d_ DateFormatter) SetDateFormat(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateFormat:"), objc.String(value))
}

// The date style of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateStyle
func (d_ DateFormatter) DateStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("dateStyle"))
	return rv
}

// SetDateStyle sets the value of the dateStyle property.
// The date style of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateStyle
func (d_ DateFormatter) SetDateStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateStyle:"), value)
}

// Returns the default formatting behavior for instances of the class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultFormatterBehavior
func (d_ DateFormatter) DefaultFormatterBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("defaultFormatterBehavior"))
	return rv
}

// SetDefaultFormatterBehavior sets the value of the defaultFormatterBehavior property.
// Returns the default formatting behavior for instances of the class.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultFormatterBehavior
func (d_ DateFormatter) SetDefaultFormatterBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultFormatterBehavior:"), value)
}

// The era symbols for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/eraSymbols
func (d_ DateFormatter) EraSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("eraSymbols"))
	return rv
}

// SetEraSymbols sets the value of the eraSymbols property.
// The era symbols for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/eraSymbols
func (d_ DateFormatter) SetEraSymbols(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](d_.ID, objc.Sel("setEraSymbols:"), nsArray)
}

// The formatter behavior for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formatterBehavior
func (d_ DateFormatter) FormatterBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("formatterBehavior"))
	return rv
}

// SetFormatterBehavior sets the value of the formatterBehavior property.
// The formatter behavior for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formatterBehavior
func (d_ DateFormatter) SetFormatterBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFormatterBehavior:"), value)
}

// The capitalization formatting context used when formatting a date.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formattingContext
func (d_ DateFormatter) FormattingContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("formattingContext"))
	return rv
}

// SetFormattingContext sets the value of the formattingContext property.
// The capitalization formatting context used when formatting a date.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formattingContext
func (d_ DateFormatter) SetFormattingContext(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFormattingContext:"), value)
}

// The start date of the Gregorian calendar for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/gregorianStartDate
func (d_ DateFormatter) GregorianStartDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("gregorianStartDate"))
	return rv
}

// SetGregorianStartDate sets the value of the gregorianStartDate property.
// The start date of the Gregorian calendar for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/gregorianStartDate
func (d_ DateFormatter) SetGregorianStartDate(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setGregorianStartDate:"), value)
}

// The locale for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/locale
func (d_ DateFormatter) Locale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("locale"))
	return rv
}

// SetLocale sets the value of the locale property.
// The locale for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/locale
func (d_ DateFormatter) SetLocale(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocale:"), value)
}

// The quarter symbols for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/quarterSymbols
func (d_ DateFormatter) QuarterSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("quarterSymbols"))
	return rv
}

// SetQuarterSymbols sets the value of the quarterSymbols property.
// The quarter symbols for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/quarterSymbols
func (d_ DateFormatter) SetQuarterSymbols(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](d_.ID, objc.Sel("setQuarterSymbols:"), nsArray)
}

// The array of standalone weekday symbols for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneWeekdaySymbols
func (d_ DateFormatter) StandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("standaloneWeekdaySymbols"))
	return rv
}

// SetStandaloneWeekdaySymbols sets the value of the standaloneWeekdaySymbols property.
// The array of standalone weekday symbols for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneWeekdaySymbols
func (d_ DateFormatter) SetStandaloneWeekdaySymbols(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](d_.ID, objc.Sel("setStandaloneWeekdaySymbols:"), nsArray)
}

// The time style of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/timeStyle
func (d_ DateFormatter) TimeStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("timeStyle"))
	return rv
}

// SetTimeStyle sets the value of the timeStyle property.
// The time style of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/timeStyle
func (d_ DateFormatter) SetTimeStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeStyle:"), value)
}

// The time zone for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/timeZone
func (d_ DateFormatter) TimeZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("timeZone"))
	return rv
}

// SetTimeZone sets the value of the timeZone property.
// The time zone for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/timeZone
func (d_ DateFormatter) SetTimeZone(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
}

// The very short month symbols for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortMonthSymbols
func (d_ DateFormatter) VeryShortMonthSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("veryShortMonthSymbols"))
	return rv
}

// SetVeryShortMonthSymbols sets the value of the veryShortMonthSymbols property.
// The very short month symbols for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortMonthSymbols
func (d_ DateFormatter) SetVeryShortMonthSymbols(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](d_.ID, objc.Sel("setVeryShortMonthSymbols:"), nsArray)
}
