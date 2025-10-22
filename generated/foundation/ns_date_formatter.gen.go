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
	DateFromString(string_ string) Date
	SetLocalizedDateFormatFromTemplate(dateFormatTemplate string)
	Calendar() NSCalendar
	SetCalendar(value ICalendar)
	DateFormat() string
	SetDateFormat(value string)
	DateStyle() DateFormatterStyle
	SetDateStyle(value IDateFormatterStyle)
	EraSymbols() []string
	SetEraSymbols(value []string)
	FormatterBehavior() unsafe.Pointer
	SetFormatterBehavior(value unsafe.Pointer)
	FormattingContext() int
	SetFormattingContext(value int)
	GregorianStartDate() NSDate
	SetGregorianStartDate(value IDate)
	Locale() NSLocale
	SetLocale(value ILocale)
	QuarterSymbols() []string
	SetQuarterSymbols(value []string)
	StandaloneWeekdaySymbols() []string
	SetStandaloneWeekdaySymbols(value []string)
	TimeStyle() DateFormatterStyle
	SetTimeStyle(value IDateFormatterStyle)
	TimeZone() NSTimeZone
	SetTimeZone(value ITimeZone)
	VeryShortMonthSymbols() []string
	SetVeryShortMonthSymbols(value []string)
	AmSymbol() string
	SetAmSymbol(value string)
	DefaultDate() Date
	SetDefaultDate(value IDate)
	DoesRelativeDateFormatting() bool
	SetDoesRelativeDateFormatting(value bool)
	GeneratesCalendarDates() bool
	SetGeneratesCalendarDates(value bool)
	IsLenient() bool
	SetIsLenient(value bool)
	LongEraSymbols() string
	SetLongEraSymbols(value string)
	MonthSymbols() string
	SetMonthSymbols(value string)
	PmSymbol() string
	SetPmSymbol(value string)
	ShortMonthSymbols() string
	SetShortMonthSymbols(value string)
	ShortQuarterSymbols() string
	SetShortQuarterSymbols(value string)
	ShortStandaloneMonthSymbols() string
	SetShortStandaloneMonthSymbols(value string)
	ShortStandaloneQuarterSymbols() string
	SetShortStandaloneQuarterSymbols(value string)
	ShortStandaloneWeekdaySymbols() string
	SetShortStandaloneWeekdaySymbols(value string)
	ShortWeekdaySymbols() string
	SetShortWeekdaySymbols(value string)
	StandaloneMonthSymbols() string
	SetStandaloneMonthSymbols(value string)
	StandaloneQuarterSymbols() string
	SetStandaloneQuarterSymbols(value string)
	TwoDigitStartDate() Date
	SetTwoDigitStartDate(value IDate)
	VeryShortStandaloneMonthSymbols() string
	SetVeryShortStandaloneMonthSymbols(value string)
	VeryShortStandaloneWeekdaySymbols() string
	SetVeryShortStandaloneWeekdaySymbols(value string)
	VeryShortWeekdaySymbols() string
	SetVeryShortWeekdaySymbols(value string)
	WeekdaySymbols() string
	SetWeekdaySymbols(value string)
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




// Initializes and returns an instance that uses the OS X 10.0 formatting behavior and the given date format string in its conversions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateFormatter/initWithDateFormat:allowNaturalLanguage:

func NewDateFormatterWithDateFormatAllowNaturalLanguage(format string, flag bool) DateFormatter {
	instance := getDateFormatterClass().Alloc()
	rv := objc.Send[DateFormatter](instance.ID, objc.Sel("initWithDateFormat:allowNaturalLanguage:"), objc.String(format), flag)
	rv.Autorelease()
	return rv
}



// Returns the default formatting behavior for instances of the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultFormatterBehavior

func (dc _DateFormatterClass) DefaultFormatterBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("defaultFormatterBehavior"))
	return rv
}


// Returns a date representation of a specified string that the system interprets using the receiver’s current settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/date(from:)

func (d_ DateFormatter) DateFromString(string_ string) Date {
	rv := objc.Send[Date](d_.ID, objc.Sel("dateFromString:"), objc.String(string_))
	return rv
}



// Sets the date format from a template using the specified locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/setLocalizedDateFormatFromTemplate(_:)

func (d_ DateFormatter) SetLocalizedDateFormatFromTemplate(dateFormatTemplate string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocalizedDateFormatFromTemplate:"), objc.String(dateFormatTemplate))
}


// The calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/calendar

func (d_ DateFormatter) Calendar() NSCalendar {
	rv := objc.Send[NSCalendar](d_.ID, objc.Sel("calendar"))
	return rv
}


// The calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/calendar

func (d_ DateFormatter) SetCalendar(value ICalendar) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCalendar:"), value)
}


// The date format string used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateFormat

func (d_ DateFormatter) DateFormat() string {
	rv := objc.Send[string](d_.ID, objc.Sel("dateFormat"))
	return rv
}


// The date format string used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateFormat

func (d_ DateFormatter) SetDateFormat(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateFormat:"), objc.String(value))
}


// The date style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateStyle

func (d_ DateFormatter) DateStyle() DateFormatterStyle {
	rv := objc.Send[DateFormatterStyle](d_.ID, objc.Sel("dateStyle"))
	return rv
}


// The date style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateStyle

func (d_ DateFormatter) SetDateStyle(value IDateFormatterStyle) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateStyle:"), value)
}


// Returns the default formatting behavior for instances of the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultFormatterBehavior

func (d_ DateFormatter) DefaultFormatterBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("defaultFormatterBehavior"))
	return rv
}


// Returns the default formatting behavior for instances of the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultFormatterBehavior

func (d_ DateFormatter) SetDefaultFormatterBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultFormatterBehavior:"), value)
}


// The era symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/eraSymbols

func (d_ DateFormatter) EraSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("eraSymbols"))
	return rv
}


// The era symbols for the receiver.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formatterBehavior

func (d_ DateFormatter) FormatterBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("formatterBehavior"))
	return rv
}


// The formatter behavior for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formatterBehavior

func (d_ DateFormatter) SetFormatterBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFormatterBehavior:"), value)
}


// The capitalization formatting context used when formatting a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formattingContext

func (d_ DateFormatter) FormattingContext() int {
	rv := objc.Send[int](d_.ID, objc.Sel("formattingContext"))
	return rv
}


// The capitalization formatting context used when formatting a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formattingContext

func (d_ DateFormatter) SetFormattingContext(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFormattingContext:"), value)
}


// The start date of the Gregorian calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/gregorianStartDate

func (d_ DateFormatter) GregorianStartDate() NSDate {
	rv := objc.Send[NSDate](d_.ID, objc.Sel("gregorianStartDate"))
	return rv
}


// The start date of the Gregorian calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/gregorianStartDate

func (d_ DateFormatter) SetGregorianStartDate(value IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setGregorianStartDate:"), value)
}


// The locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/locale

func (d_ DateFormatter) Locale() NSLocale {
	rv := objc.Send[NSLocale](d_.ID, objc.Sel("locale"))
	return rv
}


// The locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/locale

func (d_ DateFormatter) SetLocale(value ILocale) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocale:"), value)
}


// The quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/quarterSymbols

func (d_ DateFormatter) QuarterSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("quarterSymbols"))
	return rv
}


// The quarter symbols for the receiver.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneWeekdaySymbols

func (d_ DateFormatter) StandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("standaloneWeekdaySymbols"))
	return rv
}


// The array of standalone weekday symbols for the receiver.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/timeStyle

func (d_ DateFormatter) TimeStyle() DateFormatterStyle {
	rv := objc.Send[DateFormatterStyle](d_.ID, objc.Sel("timeStyle"))
	return rv
}


// The time style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/timeStyle

func (d_ DateFormatter) SetTimeStyle(value IDateFormatterStyle) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeStyle:"), value)
}


// The time zone for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/timeZone

func (d_ DateFormatter) TimeZone() NSTimeZone {
	rv := objc.Send[NSTimeZone](d_.ID, objc.Sel("timeZone"))
	return rv
}


// The time zone for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/timeZone

func (d_ DateFormatter) SetTimeZone(value ITimeZone) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
}


// The very short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortMonthSymbols

func (d_ DateFormatter) VeryShortMonthSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("veryShortMonthSymbols"))
	return rv
}


// The very short month symbols for the receiver.
//
// [Full Topic]
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


// The AM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/amsymbol

func (d_ DateFormatter) AmSymbol() string {
	rv := objc.Send[string](d_.ID, objc.Sel("amSymbol"))
	return rv
}


// The AM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/amsymbol

func (d_ DateFormatter) SetAmSymbol(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAmSymbol:"), objc.String(value))
}


// The default date for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/defaultdate

func (d_ DateFormatter) DefaultDate() Date {
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


// The long era symbols for the receiver
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/longerasymbols

func (d_ DateFormatter) LongEraSymbols() string {
	rv := objc.Send[string](d_.ID, objc.Sel("longEraSymbols"))
	return rv
}


// The long era symbols for the receiver
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/longerasymbols

func (d_ DateFormatter) SetLongEraSymbols(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLongEraSymbols:"), objc.String(value))
}


// The month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/monthsymbols

func (d_ DateFormatter) MonthSymbols() string {
	rv := objc.Send[string](d_.ID, objc.Sel("monthSymbols"))
	return rv
}


// The month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/monthsymbols

func (d_ DateFormatter) SetMonthSymbols(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMonthSymbols:"), objc.String(value))
}


// The PM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/pmsymbol

func (d_ DateFormatter) PmSymbol() string {
	rv := objc.Send[string](d_.ID, objc.Sel("pmSymbol"))
	return rv
}


// The PM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/pmsymbol

func (d_ DateFormatter) SetPmSymbol(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPmSymbol:"), objc.String(value))
}


// The array of short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortmonthsymbols

func (d_ DateFormatter) ShortMonthSymbols() string {
	rv := objc.Send[string](d_.ID, objc.Sel("shortMonthSymbols"))
	return rv
}


// The array of short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortmonthsymbols

func (d_ DateFormatter) SetShortMonthSymbols(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortMonthSymbols:"), objc.String(value))
}


// The short quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortquartersymbols

func (d_ DateFormatter) ShortQuarterSymbols() string {
	rv := objc.Send[string](d_.ID, objc.Sel("shortQuarterSymbols"))
	return rv
}


// The short quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortquartersymbols

func (d_ DateFormatter) SetShortQuarterSymbols(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortQuarterSymbols:"), objc.String(value))
}


// The short standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortstandalonemonthsymbols

func (d_ DateFormatter) ShortStandaloneMonthSymbols() string {
	rv := objc.Send[string](d_.ID, objc.Sel("shortStandaloneMonthSymbols"))
	return rv
}


// The short standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortstandalonemonthsymbols

func (d_ DateFormatter) SetShortStandaloneMonthSymbols(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortStandaloneMonthSymbols:"), objc.String(value))
}


// The short standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortstandalonequartersymbols

func (d_ DateFormatter) ShortStandaloneQuarterSymbols() string {
	rv := objc.Send[string](d_.ID, objc.Sel("shortStandaloneQuarterSymbols"))
	return rv
}


// The short standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortstandalonequartersymbols

func (d_ DateFormatter) SetShortStandaloneQuarterSymbols(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortStandaloneQuarterSymbols:"), objc.String(value))
}


// The array of short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortstandaloneweekdaysymbols

func (d_ DateFormatter) ShortStandaloneWeekdaySymbols() string {
	rv := objc.Send[string](d_.ID, objc.Sel("shortStandaloneWeekdaySymbols"))
	return rv
}


// The array of short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortstandaloneweekdaysymbols

func (d_ DateFormatter) SetShortStandaloneWeekdaySymbols(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortStandaloneWeekdaySymbols:"), objc.String(value))
}


// The array of short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortweekdaysymbols

func (d_ DateFormatter) ShortWeekdaySymbols() string {
	rv := objc.Send[string](d_.ID, objc.Sel("shortWeekdaySymbols"))
	return rv
}


// The array of short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/shortweekdaysymbols

func (d_ DateFormatter) SetShortWeekdaySymbols(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortWeekdaySymbols:"), objc.String(value))
}


// The standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/standalonemonthsymbols

func (d_ DateFormatter) StandaloneMonthSymbols() string {
	rv := objc.Send[string](d_.ID, objc.Sel("standaloneMonthSymbols"))
	return rv
}


// The standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/standalonemonthsymbols

func (d_ DateFormatter) SetStandaloneMonthSymbols(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setStandaloneMonthSymbols:"), objc.String(value))
}


// The standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/standalonequartersymbols

func (d_ DateFormatter) StandaloneQuarterSymbols() string {
	rv := objc.Send[string](d_.ID, objc.Sel("standaloneQuarterSymbols"))
	return rv
}


// The standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/standalonequartersymbols

func (d_ DateFormatter) SetStandaloneQuarterSymbols(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setStandaloneQuarterSymbols:"), objc.String(value))
}


// The earliest date that can be denoted by a two-digit year specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/twodigitstartdate

func (d_ DateFormatter) TwoDigitStartDate() Date {
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
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/veryshortstandalonemonthsymbols

func (d_ DateFormatter) VeryShortStandaloneMonthSymbols() string {
	rv := objc.Send[string](d_.ID, objc.Sel("veryShortStandaloneMonthSymbols"))
	return rv
}


// The very short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/veryshortstandalonemonthsymbols

func (d_ DateFormatter) SetVeryShortStandaloneMonthSymbols(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVeryShortStandaloneMonthSymbols:"), objc.String(value))
}


// The array of very short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/veryshortstandaloneweekdaysymbols

func (d_ DateFormatter) VeryShortStandaloneWeekdaySymbols() string {
	rv := objc.Send[string](d_.ID, objc.Sel("veryShortStandaloneWeekdaySymbols"))
	return rv
}


// The array of very short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/veryshortstandaloneweekdaysymbols

func (d_ DateFormatter) SetVeryShortStandaloneWeekdaySymbols(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVeryShortStandaloneWeekdaySymbols:"), objc.String(value))
}


// The array of very short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/veryshortweekdaysymbols

func (d_ DateFormatter) VeryShortWeekdaySymbols() string {
	rv := objc.Send[string](d_.ID, objc.Sel("veryShortWeekdaySymbols"))
	return rv
}


// The array of very short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/veryshortweekdaysymbols

func (d_ DateFormatter) SetVeryShortWeekdaySymbols(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVeryShortWeekdaySymbols:"), objc.String(value))
}


// The array of weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/weekdaysymbols

func (d_ DateFormatter) WeekdaySymbols() string {
	rv := objc.Send[string](d_.ID, objc.Sel("weekdaySymbols"))
	return rv
}


// The array of weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/weekdaysymbols

func (d_ DateFormatter) SetWeekdaySymbols(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWeekdaySymbols:"), objc.String(value))
}


