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
	// properties:
	AmSymbol() IString
	SetAmSymbol(value IString)
	Calendar() ICalendar
	SetCalendar(value ICalendar)
	DateFormat() IString
	SetDateFormat(value IString)
	DateStyle() unsafe.Pointer
	SetDateStyle(value unsafe.Pointer)
	DefaultDate() IDate
	SetDefaultDate(value IDate)
	DoesRelativeDateFormatting() bool
	SetDoesRelativeDateFormatting(value bool)
	EraSymbols() IString
	SetEraSymbols(value IString)
	GeneratesCalendarDates() bool
	SetGeneratesCalendarDates(value bool)
	GregorianStartDate() IDate
	SetGregorianStartDate(value IDate)
	IsLenient() bool
	SetIsLenient(value bool)
	Locale() ILocale
	SetLocale(value ILocale)
	LongEraSymbols() IString
	SetLongEraSymbols(value IString)
	MonthSymbols() IString
	SetMonthSymbols(value IString)
	PmSymbol() IString
	SetPmSymbol(value IString)
	QuarterSymbols() IString
	SetQuarterSymbols(value IString)
	ShortMonthSymbols() IString
	SetShortMonthSymbols(value IString)
	ShortQuarterSymbols() IString
	SetShortQuarterSymbols(value IString)
	ShortStandaloneMonthSymbols() IString
	SetShortStandaloneMonthSymbols(value IString)
	ShortStandaloneQuarterSymbols() IString
	SetShortStandaloneQuarterSymbols(value IString)
	ShortStandaloneWeekdaySymbols() IString
	SetShortStandaloneWeekdaySymbols(value IString)
	ShortWeekdaySymbols() IString
	SetShortWeekdaySymbols(value IString)
	StandaloneMonthSymbols() IString
	SetStandaloneMonthSymbols(value IString)
	StandaloneQuarterSymbols() IString
	SetStandaloneQuarterSymbols(value IString)
	StandaloneWeekdaySymbols() IString
	SetStandaloneWeekdaySymbols(value IString)
	TimeStyle() unsafe.Pointer
	SetTimeStyle(value unsafe.Pointer)
	TimeZone() ITimeZone
	SetTimeZone(value ITimeZone)
	TwoDigitStartDate() IDate
	SetTwoDigitStartDate(value IDate)
	VeryShortMonthSymbols() IString
	SetVeryShortMonthSymbols(value IString)
	VeryShortStandaloneMonthSymbols() IString
	SetVeryShortStandaloneMonthSymbols(value IString)
	VeryShortStandaloneWeekdaySymbols() IString
	SetVeryShortStandaloneWeekdaySymbols(value IString)
	VeryShortWeekdaySymbols() IString
	SetVeryShortWeekdaySymbols(value IString)
	WeekdaySymbols() IString
	SetWeekdaySymbols(value IString)
	// methods:
}

// A formatter that converts between dates and their textual representations.
//
// Instances of create string representations of objects, and convert textual representations of dates and times into objects. For user-visible representations of dates and times, provides a variety of localized presets and configuration options. For fixed format representations of dates and times, you can specify a custom format string. When working with date representations in ISO 8601 format, use instead. To represent an interval between two objects, use instead. To represent a quantity of time specified by an object, use instead.


// A formatter that converts between dates and their textual representations.
//
// [Full Topic]
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



// The AM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/amsymbol
func (d_ DateFormatter) AmSymbol() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("amSymbol"))
	return rv
}


// The AM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/amsymbol
func (d_ DateFormatter) SetAmSymbol(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAmSymbol:"), value)
}


// The calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/calendar
func (d_ DateFormatter) Calendar() ICalendar {
	rv := objc.Send[Calendar](d_.ID, objc.Sel("calendar"))
	return rv
}


// The calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/calendar
func (d_ DateFormatter) SetCalendar(value ICalendar) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCalendar:"), value)
}


// The date format string used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/dateformat
func (d_ DateFormatter) DateFormat() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("dateFormat"))
	return rv
}


// The date format string used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/dateformat
func (d_ DateFormatter) SetDateFormat(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateFormat:"), value)
}


// The date style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/datestyle
func (d_ DateFormatter) DateStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("dateStyle"))
	return rv
}


// The date style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/datestyle
func (d_ DateFormatter) SetDateStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateStyle:"), value)
}


// The default date for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/defaultdate
func (d_ DateFormatter) DefaultDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("defaultDate"))
	return rv
}


// The default date for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/defaultdate
func (d_ DateFormatter) SetDefaultDate(value IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultDate:"), value)
}


// A Boolean value that indicates whether the receiver uses phrases such as “today” and “tomorrow” for the date component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/doesrelativedateformatting
func (d_ DateFormatter) DoesRelativeDateFormatting() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("doesRelativeDateFormatting"))
	return rv
}


// A Boolean value that indicates whether the receiver uses phrases such as “today” and “tomorrow” for the date component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/doesrelativedateformatting
func (d_ DateFormatter) SetDoesRelativeDateFormatting(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDoesRelativeDateFormatting:"), value)
}


// The era symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/erasymbols
func (d_ DateFormatter) EraSymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("eraSymbols"))
	return rv
}


// The era symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/erasymbols
func (d_ DateFormatter) SetEraSymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setEraSymbols:"), value)
}


// Indicates whether the formatter generates the deprecated calendar date type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/generatescalendardates
func (d_ DateFormatter) GeneratesCalendarDates() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("generatesCalendarDates"))
	return rv
}


// Indicates whether the formatter generates the deprecated calendar date type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/generatescalendardates
func (d_ DateFormatter) SetGeneratesCalendarDates(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setGeneratesCalendarDates:"), value)
}


// The start date of the Gregorian calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/gregorianstartdate
func (d_ DateFormatter) GregorianStartDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("gregorianStartDate"))
	return rv
}


// The start date of the Gregorian calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/gregorianstartdate
func (d_ DateFormatter) SetGregorianStartDate(value IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setGregorianStartDate:"), value)
}


// A Boolean value that indicates whether the receiver uses heuristics when parsing a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/islenient
func (d_ DateFormatter) IsLenient() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isLenient"))
	return rv
}


// A Boolean value that indicates whether the receiver uses heuristics when parsing a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/islenient
func (d_ DateFormatter) SetIsLenient(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsLenient:"), value)
}


// The locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/locale
func (d_ DateFormatter) Locale() ILocale {
	rv := objc.Send[Locale](d_.ID, objc.Sel("locale"))
	return rv
}


// The locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/locale
func (d_ DateFormatter) SetLocale(value ILocale) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocale:"), value)
}


// The long era symbols for the receiver
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/longerasymbols
func (d_ DateFormatter) LongEraSymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("longEraSymbols"))
	return rv
}


// The long era symbols for the receiver
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/longerasymbols
func (d_ DateFormatter) SetLongEraSymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLongEraSymbols:"), value)
}


// The month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/monthsymbols
func (d_ DateFormatter) MonthSymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("monthSymbols"))
	return rv
}


// The month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/monthsymbols
func (d_ DateFormatter) SetMonthSymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMonthSymbols:"), value)
}


// The PM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/pmsymbol
func (d_ DateFormatter) PmSymbol() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("pmSymbol"))
	return rv
}


// The PM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/pmsymbol
func (d_ DateFormatter) SetPmSymbol(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPmSymbol:"), value)
}


// The quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/quartersymbols
func (d_ DateFormatter) QuarterSymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("quarterSymbols"))
	return rv
}


// The quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/quartersymbols
func (d_ DateFormatter) SetQuarterSymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setQuarterSymbols:"), value)
}


// The array of short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortmonthsymbols
func (d_ DateFormatter) ShortMonthSymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("shortMonthSymbols"))
	return rv
}


// The array of short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortmonthsymbols
func (d_ DateFormatter) SetShortMonthSymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortMonthSymbols:"), value)
}


// The short quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortquartersymbols
func (d_ DateFormatter) ShortQuarterSymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("shortQuarterSymbols"))
	return rv
}


// The short quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortquartersymbols
func (d_ DateFormatter) SetShortQuarterSymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortQuarterSymbols:"), value)
}


// The short standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortstandalonemonthsymbols
func (d_ DateFormatter) ShortStandaloneMonthSymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("shortStandaloneMonthSymbols"))
	return rv
}


// The short standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortstandalonemonthsymbols
func (d_ DateFormatter) SetShortStandaloneMonthSymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortStandaloneMonthSymbols:"), value)
}


// The short standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortstandalonequartersymbols
func (d_ DateFormatter) ShortStandaloneQuarterSymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("shortStandaloneQuarterSymbols"))
	return rv
}


// The short standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortstandalonequartersymbols
func (d_ DateFormatter) SetShortStandaloneQuarterSymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortStandaloneQuarterSymbols:"), value)
}


// The array of short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortstandaloneweekdaysymbols
func (d_ DateFormatter) ShortStandaloneWeekdaySymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("shortStandaloneWeekdaySymbols"))
	return rv
}


// The array of short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortstandaloneweekdaysymbols
func (d_ DateFormatter) SetShortStandaloneWeekdaySymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortStandaloneWeekdaySymbols:"), value)
}


// The array of short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortweekdaysymbols
func (d_ DateFormatter) ShortWeekdaySymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("shortWeekdaySymbols"))
	return rv
}


// The array of short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortweekdaysymbols
func (d_ DateFormatter) SetShortWeekdaySymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortWeekdaySymbols:"), value)
}


// The standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/standalonemonthsymbols
func (d_ DateFormatter) StandaloneMonthSymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("standaloneMonthSymbols"))
	return rv
}


// The standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/standalonemonthsymbols
func (d_ DateFormatter) SetStandaloneMonthSymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setStandaloneMonthSymbols:"), value)
}


// The standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/standalonequartersymbols
func (d_ DateFormatter) StandaloneQuarterSymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("standaloneQuarterSymbols"))
	return rv
}


// The standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/standalonequartersymbols
func (d_ DateFormatter) SetStandaloneQuarterSymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setStandaloneQuarterSymbols:"), value)
}


// The array of standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/standaloneweekdaysymbols
func (d_ DateFormatter) StandaloneWeekdaySymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("standaloneWeekdaySymbols"))
	return rv
}


// The array of standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/standaloneweekdaysymbols
func (d_ DateFormatter) SetStandaloneWeekdaySymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setStandaloneWeekdaySymbols:"), value)
}


// The time style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/timestyle
func (d_ DateFormatter) TimeStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("timeStyle"))
	return rv
}


// The time style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/timestyle
func (d_ DateFormatter) SetTimeStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeStyle:"), value)
}


// The time zone for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/timezone
func (d_ DateFormatter) TimeZone() ITimeZone {
	rv := objc.Send[TimeZone](d_.ID, objc.Sel("timeZone"))
	return rv
}


// The time zone for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/timezone
func (d_ DateFormatter) SetTimeZone(value ITimeZone) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
}


// The earliest date that can be denoted by a two-digit year specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/twodigitstartdate
func (d_ DateFormatter) TwoDigitStartDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("twoDigitStartDate"))
	return rv
}


// The earliest date that can be denoted by a two-digit year specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/twodigitstartdate
func (d_ DateFormatter) SetTwoDigitStartDate(value IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTwoDigitStartDate:"), value)
}


// The very short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/veryshortmonthsymbols
func (d_ DateFormatter) VeryShortMonthSymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("veryShortMonthSymbols"))
	return rv
}


// The very short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/veryshortmonthsymbols
func (d_ DateFormatter) SetVeryShortMonthSymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVeryShortMonthSymbols:"), value)
}


// The very short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/veryshortstandalonemonthsymbols
func (d_ DateFormatter) VeryShortStandaloneMonthSymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("veryShortStandaloneMonthSymbols"))
	return rv
}


// The very short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/veryshortstandalonemonthsymbols
func (d_ DateFormatter) SetVeryShortStandaloneMonthSymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVeryShortStandaloneMonthSymbols:"), value)
}


// The array of very short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/veryshortstandaloneweekdaysymbols
func (d_ DateFormatter) VeryShortStandaloneWeekdaySymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("veryShortStandaloneWeekdaySymbols"))
	return rv
}


// The array of very short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/veryshortstandaloneweekdaysymbols
func (d_ DateFormatter) SetVeryShortStandaloneWeekdaySymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVeryShortStandaloneWeekdaySymbols:"), value)
}


// The array of very short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/veryshortweekdaysymbols
func (d_ DateFormatter) VeryShortWeekdaySymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("veryShortWeekdaySymbols"))
	return rv
}


// The array of very short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/veryshortweekdaysymbols
func (d_ DateFormatter) SetVeryShortWeekdaySymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVeryShortWeekdaySymbols:"), value)
}


// The array of weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/weekdaysymbols
func (d_ DateFormatter) WeekdaySymbols() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("weekdaySymbols"))
	return rv
}


// The array of weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/weekdaysymbols
func (d_ DateFormatter) SetWeekdaySymbols(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWeekdaySymbols:"), value)
}



