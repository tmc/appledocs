// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	AMSymbol() string /* primitive/slice/pointer */
	SetAMSymbol(value string /* primitive/slice/pointer */)
	Calendar() ICalendar
	SetCalendar(value ICalendar)
	DateFormat() string /* primitive/slice/pointer */
	SetDateFormat(value string /* primitive/slice/pointer */)
	DateStyle() DateFormatterStyle
	SetDateStyle(value DateFormatterStyle)
	DefaultDate() IDate
	SetDefaultDate(value IDate)
	DoesRelativeDateFormatting() bool /* primitive/slice/pointer */
	SetDoesRelativeDateFormatting(value bool /* primitive/slice/pointer */)
	EraSymbols() []string /* primitive/slice/pointer */
	SetEraSymbols(value []string /* primitive/slice/pointer */)
	FormatterBehavior() DateFormatterBehavior
	SetFormatterBehavior(value DateFormatterBehavior)
	FormattingContext() int /* primitive/slice/pointer */
	SetFormattingContext(value int /* primitive/slice/pointer */)
	GeneratesCalendarDates() bool /* primitive/slice/pointer */
	SetGeneratesCalendarDates(value bool /* primitive/slice/pointer */)
	GregorianStartDate() IDate
	SetGregorianStartDate(value IDate)
	Lenient() bool /* primitive/slice/pointer */
	SetLenient(value bool /* primitive/slice/pointer */)
	Locale() ILocale
	SetLocale(value ILocale)
	LongEraSymbols() []string /* primitive/slice/pointer */
	SetLongEraSymbols(value []string /* primitive/slice/pointer */)
	MonthSymbols() []string /* primitive/slice/pointer */
	SetMonthSymbols(value []string /* primitive/slice/pointer */)
	PMSymbol() string /* primitive/slice/pointer */
	SetPMSymbol(value string /* primitive/slice/pointer */)
	QuarterSymbols() []string /* primitive/slice/pointer */
	SetQuarterSymbols(value []string /* primitive/slice/pointer */)
	ShortMonthSymbols() []string /* primitive/slice/pointer */
	SetShortMonthSymbols(value []string /* primitive/slice/pointer */)
	ShortQuarterSymbols() []string /* primitive/slice/pointer */
	SetShortQuarterSymbols(value []string /* primitive/slice/pointer */)
	ShortStandaloneMonthSymbols() []string /* primitive/slice/pointer */
	SetShortStandaloneMonthSymbols(value []string /* primitive/slice/pointer */)
	ShortStandaloneQuarterSymbols() []string /* primitive/slice/pointer */
	SetShortStandaloneQuarterSymbols(value []string /* primitive/slice/pointer */)
	ShortStandaloneWeekdaySymbols() []string /* primitive/slice/pointer */
	SetShortStandaloneWeekdaySymbols(value []string /* primitive/slice/pointer */)
	ShortWeekdaySymbols() []string /* primitive/slice/pointer */
	SetShortWeekdaySymbols(value []string /* primitive/slice/pointer */)
	StandaloneMonthSymbols() []string /* primitive/slice/pointer */
	SetStandaloneMonthSymbols(value []string /* primitive/slice/pointer */)
	StandaloneQuarterSymbols() []string /* primitive/slice/pointer */
	SetStandaloneQuarterSymbols(value []string /* primitive/slice/pointer */)
	StandaloneWeekdaySymbols() []string /* primitive/slice/pointer */
	SetStandaloneWeekdaySymbols(value []string /* primitive/slice/pointer */)
	TimeStyle() DateFormatterStyle
	SetTimeStyle(value DateFormatterStyle)
	TimeZone() ITimeZone
	SetTimeZone(value ITimeZone)
	TwoDigitStartDate() IDate
	SetTwoDigitStartDate(value IDate)
	VeryShortMonthSymbols() []string /* primitive/slice/pointer */
	SetVeryShortMonthSymbols(value []string /* primitive/slice/pointer */)
	VeryShortStandaloneMonthSymbols() []string /* primitive/slice/pointer */
	SetVeryShortStandaloneMonthSymbols(value []string /* primitive/slice/pointer */)
	VeryShortStandaloneWeekdaySymbols() []string /* primitive/slice/pointer */
	SetVeryShortStandaloneWeekdaySymbols(value []string /* primitive/slice/pointer */)
	VeryShortWeekdaySymbols() []string /* primitive/slice/pointer */
	SetVeryShortWeekdaySymbols(value []string /* primitive/slice/pointer */)
	WeekdaySymbols() []string /* primitive/slice/pointer */
	SetWeekdaySymbols(value []string /* primitive/slice/pointer */)
	IsLenient() bool /* primitive/slice/pointer */
	SetIsLenient(value bool /* primitive/slice/pointer */)
	// methods:
	DateFromString(string_ string /* primitive/slice/pointer */) IDate
	GetObjectValueForStringRangeError(obj objectivec.IObject, string_ string /* primitive/slice/pointer */, rangep Range /* foo */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */
	SetLocalizedDateFormatFromTemplate(dateFormatTemplate string /* primitive/slice/pointer */)
	StringFromDate(date IDate) String /* foo */
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



// Returns a localized date format string representing the given date format components arranged appropriately for the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateFormat(fromTemplate:options:locale:)
func (dc _DateFormatterClass) DateFormatFromTemplateOptionsLocale(tmplate string /* primitive/slice/pointer */, opts uint /* primitive/slice/pointer */, locale ILocale) String /* foo */ {
	rv := objc.Send[String](objc.ID(dc.class), objc.Sel("dateFormatFromTemplate:options:locale:"), objc.String(tmplate), opts, locale)
	return rv
}


// Returns a string representation of a specified date, that the system formats for the current locale using the specified date and time styles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/localizedString(from:dateStyle:timeStyle:)
func (dc _DateFormatterClass) LocalizedStringFromDateDateStyleTimeStyle(date IDate, dstyle DateFormatterStyle, tstyle DateFormatterStyle) String /* foo */ {
	rv := objc.Send[String](objc.ID(dc.class), objc.Sel("localizedStringFromDate:dateStyle:timeStyle:"), date, dstyle, tstyle)
	return rv
}


// Returns the default formatting behavior for instances of the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultFormatterBehavior
func (dc _DateFormatterClass) DefaultFormatterBehavior() DateFormatterBehavior {
	rv := objc.Send[DateFormatterBehavior](objc.ID(dc.class), objc.Sel("defaultFormatterBehavior"))
	return rv
}

// Returns a date representation of a specified string that the system interprets using the receiver’s current settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/date(from:)
func (d_ DateFormatter) DateFromString(string_ string /* primitive/slice/pointer */) IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("dateFromString:"), objc.String(string_))
	return rv
}


// Returns by reference a date representation of a specified string and its date range, as well as a Boolean value that indicates whether the system can parse the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/getObjectValue(_:for:range:)
func (d_ DateFormatter) GetObjectValueForStringRangeError(obj objectivec.IObject, string_ string /* primitive/slice/pointer */, rangep Range /* foo */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("getObjectValue:forString:range:error:"), obj, objc.String(string_), rangep, error_)
	return rv
}


// Sets the date format from a template using the specified locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/setLocalizedDateFormatFromTemplate(_:)
func (d_ DateFormatter) SetLocalizedDateFormatFromTemplate(dateFormatTemplate string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocalizedDateFormatFromTemplate:"), objc.String(dateFormatTemplate))
}


// Returns a string representation of a specified date that the system formats using the receiver’s current settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/string(from:)
func (d_ DateFormatter) StringFromDate(date IDate) String /* foo */ {
	rv := objc.Send[String](d_.ID, objc.Sel("stringFromDate:"), date)
	return rv
}


// The AM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/amSymbol
func (d_ DateFormatter) AMSymbol() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](d_.ID, objc.Sel("AMSymbol"))
	return rv
}


// The AM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/amSymbol
func (d_ DateFormatter) SetAMSymbol(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAMSymbol:"), objc.String(value))
}


// The calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/calendar
func (d_ DateFormatter) Calendar() ICalendar {
	rv := objc.Send[Calendar](d_.ID, objc.Sel("calendar"))
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
func (d_ DateFormatter) DateFormat() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](d_.ID, objc.Sel("dateFormat"))
	return rv
}


// The date format string used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateFormat
func (d_ DateFormatter) SetDateFormat(value string /* primitive/slice/pointer */) {
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
func (d_ DateFormatter) SetDateStyle(value DateFormatterStyle) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateStyle:"), value)
}


// The default date for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultDate
func (d_ DateFormatter) DefaultDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("defaultDate"))
	return rv
}


// The default date for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultDate
func (d_ DateFormatter) SetDefaultDate(value IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultDate:"), value)
}


// Returns the default formatting behavior for instances of the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultFormatterBehavior
func (d_ DateFormatter) DefaultFormatterBehavior() DateFormatterBehavior {
	rv := objc.Send[DateFormatterBehavior](d_.ID, objc.Sel("defaultFormatterBehavior"))
	return rv
}


// Returns the default formatting behavior for instances of the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultFormatterBehavior
func (d_ DateFormatter) SetDefaultFormatterBehavior(value DateFormatterBehavior) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultFormatterBehavior:"), value)
}


// A Boolean value that indicates whether the receiver uses phrases such as “today” and “tomorrow” for the date component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/doesRelativeDateFormatting
func (d_ DateFormatter) DoesRelativeDateFormatting() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("doesRelativeDateFormatting"))
	return rv
}


// A Boolean value that indicates whether the receiver uses phrases such as “today” and “tomorrow” for the date component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/doesRelativeDateFormatting
func (d_ DateFormatter) SetDoesRelativeDateFormatting(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDoesRelativeDateFormatting:"), value)
}


// The era symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/eraSymbols
func (d_ DateFormatter) EraSymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("eraSymbols"))
	return rv
}


// The era symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/eraSymbols
func (d_ DateFormatter) SetEraSymbols(value []string /* primitive/slice/pointer */) {
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
func (d_ DateFormatter) FormatterBehavior() DateFormatterBehavior {
	rv := objc.Send[DateFormatterBehavior](d_.ID, objc.Sel("formatterBehavior"))
	return rv
}


// The formatter behavior for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formatterBehavior
func (d_ DateFormatter) SetFormatterBehavior(value DateFormatterBehavior) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFormatterBehavior:"), value)
}


// The capitalization formatting context used when formatting a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formattingContext
func (d_ DateFormatter) FormattingContext() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](d_.ID, objc.Sel("formattingContext"))
	return rv
}


// The capitalization formatting context used when formatting a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formattingContext
func (d_ DateFormatter) SetFormattingContext(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFormattingContext:"), value)
}


// Indicates whether the formatter generates the deprecated calendar date type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/generatesCalendarDates
func (d_ DateFormatter) GeneratesCalendarDates() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("generatesCalendarDates"))
	return rv
}


// Indicates whether the formatter generates the deprecated calendar date type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/generatesCalendarDates
func (d_ DateFormatter) SetGeneratesCalendarDates(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setGeneratesCalendarDates:"), value)
}


// The start date of the Gregorian calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/gregorianStartDate
func (d_ DateFormatter) GregorianStartDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("gregorianStartDate"))
	return rv
}


// The start date of the Gregorian calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/gregorianStartDate
func (d_ DateFormatter) SetGregorianStartDate(value IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setGregorianStartDate:"), value)
}


// A Boolean value that indicates whether the receiver uses heuristics when parsing a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/isLenient
func (d_ DateFormatter) Lenient() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("lenient"))
	return rv
}


// A Boolean value that indicates whether the receiver uses heuristics when parsing a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/isLenient
func (d_ DateFormatter) SetLenient(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLenient:"), value)
}


// The locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/locale
func (d_ DateFormatter) Locale() ILocale {
	rv := objc.Send[Locale](d_.ID, objc.Sel("locale"))
	return rv
}


// The locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/locale
func (d_ DateFormatter) SetLocale(value ILocale) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocale:"), value)
}


// The long era symbols for the receiver
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/longEraSymbols
func (d_ DateFormatter) LongEraSymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("longEraSymbols"))
	return rv
}


// The long era symbols for the receiver
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/longEraSymbols
func (d_ DateFormatter) SetLongEraSymbols(value []string /* primitive/slice/pointer */) {
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
	objc.Send[objc.ID](d_.ID, objc.Sel("setLongEraSymbols:"), nsArray)
}


// The month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/monthSymbols
func (d_ DateFormatter) MonthSymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("monthSymbols"))
	return rv
}


// The month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/monthSymbols
func (d_ DateFormatter) SetMonthSymbols(value []string /* primitive/slice/pointer */) {
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
	objc.Send[objc.ID](d_.ID, objc.Sel("setMonthSymbols:"), nsArray)
}


// The PM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/pmSymbol
func (d_ DateFormatter) PMSymbol() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](d_.ID, objc.Sel("PMSymbol"))
	return rv
}


// The PM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/pmSymbol
func (d_ DateFormatter) SetPMSymbol(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPMSymbol:"), objc.String(value))
}


// The quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/quarterSymbols
func (d_ DateFormatter) QuarterSymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("quarterSymbols"))
	return rv
}


// The quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/quarterSymbols
func (d_ DateFormatter) SetQuarterSymbols(value []string /* primitive/slice/pointer */) {
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


// The array of short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortMonthSymbols
func (d_ DateFormatter) ShortMonthSymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("shortMonthSymbols"))
	return rv
}


// The array of short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortMonthSymbols
func (d_ DateFormatter) SetShortMonthSymbols(value []string /* primitive/slice/pointer */) {
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
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortMonthSymbols:"), nsArray)
}


// The short quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortQuarterSymbols
func (d_ DateFormatter) ShortQuarterSymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("shortQuarterSymbols"))
	return rv
}


// The short quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortQuarterSymbols
func (d_ DateFormatter) SetShortQuarterSymbols(value []string /* primitive/slice/pointer */) {
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
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortQuarterSymbols:"), nsArray)
}


// The short standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneMonthSymbols
func (d_ DateFormatter) ShortStandaloneMonthSymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("shortStandaloneMonthSymbols"))
	return rv
}


// The short standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneMonthSymbols
func (d_ DateFormatter) SetShortStandaloneMonthSymbols(value []string /* primitive/slice/pointer */) {
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
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortStandaloneMonthSymbols:"), nsArray)
}


// The short standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneQuarterSymbols
func (d_ DateFormatter) ShortStandaloneQuarterSymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("shortStandaloneQuarterSymbols"))
	return rv
}


// The short standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneQuarterSymbols
func (d_ DateFormatter) SetShortStandaloneQuarterSymbols(value []string /* primitive/slice/pointer */) {
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
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortStandaloneQuarterSymbols:"), nsArray)
}


// The array of short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneWeekdaySymbols
func (d_ DateFormatter) ShortStandaloneWeekdaySymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("shortStandaloneWeekdaySymbols"))
	return rv
}


// The array of short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneWeekdaySymbols
func (d_ DateFormatter) SetShortStandaloneWeekdaySymbols(value []string /* primitive/slice/pointer */) {
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
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortStandaloneWeekdaySymbols:"), nsArray)
}


// The array of short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortWeekdaySymbols
func (d_ DateFormatter) ShortWeekdaySymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("shortWeekdaySymbols"))
	return rv
}


// The array of short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortWeekdaySymbols
func (d_ DateFormatter) SetShortWeekdaySymbols(value []string /* primitive/slice/pointer */) {
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
	objc.Send[objc.ID](d_.ID, objc.Sel("setShortWeekdaySymbols:"), nsArray)
}


// The standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneMonthSymbols
func (d_ DateFormatter) StandaloneMonthSymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("standaloneMonthSymbols"))
	return rv
}


// The standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneMonthSymbols
func (d_ DateFormatter) SetStandaloneMonthSymbols(value []string /* primitive/slice/pointer */) {
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
	objc.Send[objc.ID](d_.ID, objc.Sel("setStandaloneMonthSymbols:"), nsArray)
}


// The standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneQuarterSymbols
func (d_ DateFormatter) StandaloneQuarterSymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("standaloneQuarterSymbols"))
	return rv
}


// The standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneQuarterSymbols
func (d_ DateFormatter) SetStandaloneQuarterSymbols(value []string /* primitive/slice/pointer */) {
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
	objc.Send[objc.ID](d_.ID, objc.Sel("setStandaloneQuarterSymbols:"), nsArray)
}


// The array of standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneWeekdaySymbols
func (d_ DateFormatter) StandaloneWeekdaySymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("standaloneWeekdaySymbols"))
	return rv
}


// The array of standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneWeekdaySymbols
func (d_ DateFormatter) SetStandaloneWeekdaySymbols(value []string /* primitive/slice/pointer */) {
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
func (d_ DateFormatter) SetTimeStyle(value DateFormatterStyle) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeStyle:"), value)
}


// The time zone for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/timeZone
func (d_ DateFormatter) TimeZone() ITimeZone {
	rv := objc.Send[TimeZone](d_.ID, objc.Sel("timeZone"))
	return rv
}


// The time zone for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/timeZone
func (d_ DateFormatter) SetTimeZone(value ITimeZone) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
}


// The earliest date that can be denoted by a two-digit year specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/twoDigitStartDate
func (d_ DateFormatter) TwoDigitStartDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("twoDigitStartDate"))
	return rv
}


// The earliest date that can be denoted by a two-digit year specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/twoDigitStartDate
func (d_ DateFormatter) SetTwoDigitStartDate(value IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTwoDigitStartDate:"), value)
}


// The very short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortMonthSymbols
func (d_ DateFormatter) VeryShortMonthSymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("veryShortMonthSymbols"))
	return rv
}


// The very short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortMonthSymbols
func (d_ DateFormatter) SetVeryShortMonthSymbols(value []string /* primitive/slice/pointer */) {
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


// The very short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortStandaloneMonthSymbols
func (d_ DateFormatter) VeryShortStandaloneMonthSymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("veryShortStandaloneMonthSymbols"))
	return rv
}


// The very short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortStandaloneMonthSymbols
func (d_ DateFormatter) SetVeryShortStandaloneMonthSymbols(value []string /* primitive/slice/pointer */) {
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
	objc.Send[objc.ID](d_.ID, objc.Sel("setVeryShortStandaloneMonthSymbols:"), nsArray)
}


// The array of very short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortStandaloneWeekdaySymbols
func (d_ DateFormatter) VeryShortStandaloneWeekdaySymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("veryShortStandaloneWeekdaySymbols"))
	return rv
}


// The array of very short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortStandaloneWeekdaySymbols
func (d_ DateFormatter) SetVeryShortStandaloneWeekdaySymbols(value []string /* primitive/slice/pointer */) {
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
	objc.Send[objc.ID](d_.ID, objc.Sel("setVeryShortStandaloneWeekdaySymbols:"), nsArray)
}


// The array of very short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortWeekdaySymbols
func (d_ DateFormatter) VeryShortWeekdaySymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("veryShortWeekdaySymbols"))
	return rv
}


// The array of very short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortWeekdaySymbols
func (d_ DateFormatter) SetVeryShortWeekdaySymbols(value []string /* primitive/slice/pointer */) {
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
	objc.Send[objc.ID](d_.ID, objc.Sel("setVeryShortWeekdaySymbols:"), nsArray)
}


// The array of weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/weekdaySymbols
func (d_ DateFormatter) WeekdaySymbols() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("weekdaySymbols"))
	return rv
}


// The array of weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/weekdaySymbols
func (d_ DateFormatter) SetWeekdaySymbols(value []string /* primitive/slice/pointer */) {
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
	objc.Send[objc.ID](d_.ID, objc.Sel("setWeekdaySymbols:"), nsArray)
}


// A Boolean value that indicates whether the receiver uses heuristics when parsing a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/islenient
func (d_ DateFormatter) IsLenient() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("isLenient"))
	return rv
}


// A Boolean value that indicates whether the receiver uses heuristics when parsing a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/islenient
func (d_ DateFormatter) SetIsLenient(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsLenient:"), value)
}



