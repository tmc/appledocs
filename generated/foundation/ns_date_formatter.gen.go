// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDateFormatter */


/* debug [class_header]: Header for NSDateFormatter */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DateFormatter */
// An interface definition for the [DateFormatter] class.
type IDateFormatter interface {
	IFormatter
	
/* debug [class_interface_properties]: Properties for DateFormatter */
	// properties:
	AMSymbol() IString
	SetAMSymbol(value IString)
	Calendar() ICalendar
	SetCalendar(value ICalendar)
	DateFormat() IString
	SetDateFormat(value IString)
	DateStyle() DateFormatterStyle
	SetDateStyle(value DateFormatterStyle)
	DefaultDate() IDate
	SetDefaultDate(value IDate)
	DoesRelativeDateFormatting() bool
	SetDoesRelativeDateFormatting(value bool)
	EraSymbols() []string
	SetEraSymbols(value []string)
	FormatterBehavior() DateFormatterBehavior
	SetFormatterBehavior(value DateFormatterBehavior)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	GeneratesCalendarDates() bool
	SetGeneratesCalendarDates(value bool)
	GregorianStartDate() IDate
	SetGregorianStartDate(value IDate)
	Lenient() bool
	SetLenient(value bool)
	Locale() ILocale
	SetLocale(value ILocale)
	LongEraSymbols() []string
	SetLongEraSymbols(value []string)
	MonthSymbols() []string
	SetMonthSymbols(value []string)
	PMSymbol() IString
	SetPMSymbol(value IString)
	QuarterSymbols() []string
	SetQuarterSymbols(value []string)
	ShortMonthSymbols() []string
	SetShortMonthSymbols(value []string)
	ShortQuarterSymbols() []string
	SetShortQuarterSymbols(value []string)
	ShortStandaloneMonthSymbols() []string
	SetShortStandaloneMonthSymbols(value []string)
	ShortStandaloneQuarterSymbols() []string
	SetShortStandaloneQuarterSymbols(value []string)
	ShortStandaloneWeekdaySymbols() []string
	SetShortStandaloneWeekdaySymbols(value []string)
	ShortWeekdaySymbols() []string
	SetShortWeekdaySymbols(value []string)
	StandaloneMonthSymbols() []string
	SetStandaloneMonthSymbols(value []string)
	StandaloneQuarterSymbols() []string
	SetStandaloneQuarterSymbols(value []string)
	StandaloneWeekdaySymbols() []string
	SetStandaloneWeekdaySymbols(value []string)
	TimeStyle() DateFormatterStyle
	SetTimeStyle(value DateFormatterStyle)
	TimeZone() ITimeZone
	SetTimeZone(value ITimeZone)
	TwoDigitStartDate() IDate
	SetTwoDigitStartDate(value IDate)
	VeryShortMonthSymbols() []string
	SetVeryShortMonthSymbols(value []string)
	VeryShortStandaloneMonthSymbols() []string
	SetVeryShortStandaloneMonthSymbols(value []string)
	VeryShortStandaloneWeekdaySymbols() []string
	SetVeryShortStandaloneWeekdaySymbols(value []string)
	VeryShortWeekdaySymbols() []string
	SetVeryShortWeekdaySymbols(value []string)
	WeekdaySymbols() []string
	SetWeekdaySymbols(value []string)
	IsLenient() bool
	SetIsLenient(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DateFormatter */
	// methods:
	DateFromString(string_ IString) IDate
	GetObjectValueForStringRangeError(obj objectivec.IObject, string_ IString, rangep objc.IObject /* cross-framework: Range */, error_ IError) bool
	SetLocalizedDateFormatFromTemplate(dateFormatTemplate IString)
	StringFromDate(date IDate) IString
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DateFormatter */
// Alloc allocates a new instance without initialization.
func (dc _DateFormatterClass) Alloc() DateFormatter {
	rv := objc.Send[DateFormatter](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DateFormatter */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DateFormatter */

// Initializes and returns an instance that uses the OS X 10.0 formatting behavior and the given date format string in its conversions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateFormatter/initWithDateFormat:allowNaturalLanguage:
func NewDateFormatterWithDateFormatAllowNaturalLanguage(format IString, flag bool) DateFormatter {
	instance := getDateFormatterClass().Alloc()
	rv := objc.Send[DateFormatter](instance.ID, objc.Sel("initWithDateFormat:allowNaturalLanguage:"), format, flag)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDateFormatterWithDateFormatAllowNaturalLanguage */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DateFormatter */

// Returns a localized date format string representing the given date format components arranged appropriately for the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateFormat(fromTemplate:options:locale:)
func (dc _DateFormatterClass) DateFormatFromTemplateOptionsLocale(tmplate IString, opts uint, locale ILocale) IString {
	rv := objc.Send[String](objc.ID(dc.class), objc.Sel("dateFormatFromTemplate:options:locale:"), tmplate, opts, locale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DateFormatFromTemplateOptionsLocale) */


// Returns a string representation of a specified date, that the system formats for the current locale using the specified date and time styles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/localizedString(from:dateStyle:timeStyle:)
func (dc _DateFormatterClass) LocalizedStringFromDateDateStyleTimeStyle(date IDate, dstyle DateFormatterStyle, tstyle DateFormatterStyle) IString {
	rv := objc.Send[String](objc.ID(dc.class), objc.Sel("localizedStringFromDate:dateStyle:timeStyle:"), date, dstyle, tstyle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocalizedStringFromDateDateStyleTimeStyle) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DateFormatter */

// Returns the default formatting behavior for instances of the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultFormatterBehavior
func (dc _DateFormatterClass) DefaultFormatterBehavior() DateFormatterBehavior {
	rv := objc.Send[DateFormatterBehavior](objc.ID(dc.class), objc.Sel("defaultFormatterBehavior"))
	return rv
}/* debug [class_properties_class/property]: defaultFormatterBehavior */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DateFormatter */

// Returns a date representation of a specified string that the system interprets using the receiver’s current settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/date(from:)
func (d_ DateFormatter) DateFromString(string_ IString) IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("dateFromString:"), string_)
	return rv
}/* debug [instance_methods/method]: DateFromString */


// Returns by reference a date representation of a specified string and its date range, as well as a Boolean value that indicates whether the system can parse the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/getObjectValue(_:for:range:)
func (d_ DateFormatter) GetObjectValueForStringRangeError(obj objectivec.IObject, string_ IString, rangep objc.IObject /* cross-framework: Range */, error_ IError) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("getObjectValue:forString:range:error:"), obj, string_, rangep, error_)
	return rv
}/* debug [instance_methods/method]: GetObjectValueForStringRangeError */


// Sets the date format from a template using the specified locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/setLocalizedDateFormatFromTemplate(_:)
func (d_ DateFormatter) SetLocalizedDateFormatFromTemplate(dateFormatTemplate IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocalizedDateFormatFromTemplate:"), dateFormatTemplate)
}/* debug [instance_methods/method]: SetLocalizedDateFormatFromTemplate */


// Returns a string representation of a specified date that the system formats using the receiver’s current settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/string(from:)
func (d_ DateFormatter) StringFromDate(date IDate) IString {
	rv := objc.Send[String](d_.ID, objc.Sel("stringFromDate:"), date)
	return rv
}/* debug [instance_methods/method]: StringFromDate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DateFormatter */

// The AM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/amSymbol
func (d_ DateFormatter) AMSymbol() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("AMSymbol"))
	return rv
}/* debug [instance_properties/getter]: AMSymbol */


// The AM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/amSymbol
func (d_ DateFormatter) SetAMSymbol(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAMSymbol:"), value)
}/* debug [instance_properties/setter]: AMSymbol */


// The calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/calendar
func (d_ DateFormatter) Calendar() ICalendar {
	rv := objc.Send[Calendar](d_.ID, objc.Sel("calendar"))
	return rv
}/* debug [instance_properties/getter]: calendar */


// The calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/calendar
func (d_ DateFormatter) SetCalendar(value ICalendar) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCalendar:"), value)
}/* debug [instance_properties/setter]: calendar */


// The date format string used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateFormat
func (d_ DateFormatter) DateFormat() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("dateFormat"))
	return rv
}/* debug [instance_properties/getter]: dateFormat */


// The date format string used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateFormat
func (d_ DateFormatter) SetDateFormat(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateFormat:"), value)
}/* debug [instance_properties/setter]: dateFormat */


// The date style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateStyle
func (d_ DateFormatter) DateStyle() DateFormatterStyle {
	rv := objc.Send[DateFormatterStyle](d_.ID, objc.Sel("dateStyle"))
	return rv
}/* debug [instance_properties/getter]: dateStyle */


// The date style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/dateStyle
func (d_ DateFormatter) SetDateStyle(value DateFormatterStyle) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateStyle:"), value)
}/* debug [instance_properties/setter]: dateStyle */


// The default date for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultDate
func (d_ DateFormatter) DefaultDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("defaultDate"))
	return rv
}/* debug [instance_properties/getter]: defaultDate */


// The default date for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultDate
func (d_ DateFormatter) SetDefaultDate(value IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultDate:"), value)
}/* debug [instance_properties/setter]: defaultDate */


// Returns the default formatting behavior for instances of the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultFormatterBehavior
func (d_ DateFormatter) DefaultFormatterBehavior() DateFormatterBehavior {
	rv := objc.Send[DateFormatterBehavior](d_.ID, objc.Sel("defaultFormatterBehavior"))
	return rv
}/* debug [instance_properties/getter]: defaultFormatterBehavior */


// Returns the default formatting behavior for instances of the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultFormatterBehavior
func (d_ DateFormatter) SetDefaultFormatterBehavior(value DateFormatterBehavior) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultFormatterBehavior:"), value)
}/* debug [instance_properties/setter]: defaultFormatterBehavior */


// A Boolean value that indicates whether the receiver uses phrases such as “today” and “tomorrow” for the date component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/doesRelativeDateFormatting
func (d_ DateFormatter) DoesRelativeDateFormatting() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("doesRelativeDateFormatting"))
	return rv
}/* debug [instance_properties/getter]: doesRelativeDateFormatting */


// A Boolean value that indicates whether the receiver uses phrases such as “today” and “tomorrow” for the date component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/doesRelativeDateFormatting
func (d_ DateFormatter) SetDoesRelativeDateFormatting(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDoesRelativeDateFormatting:"), value)
}/* debug [instance_properties/setter]: doesRelativeDateFormatting */


// The era symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/eraSymbols
func (d_ DateFormatter) EraSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("eraSymbols"))
	return rv
}/* debug [instance_properties/getter]: eraSymbols */


// The era symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/eraSymbols
func (d_ DateFormatter) SetEraSymbols(value []string) {
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
}/* debug [instance_properties/setter]: eraSymbols */


// The formatter behavior for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formatterBehavior
func (d_ DateFormatter) FormatterBehavior() DateFormatterBehavior {
	rv := objc.Send[DateFormatterBehavior](d_.ID, objc.Sel("formatterBehavior"))
	return rv
}/* debug [instance_properties/getter]: formatterBehavior */


// The formatter behavior for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formatterBehavior
func (d_ DateFormatter) SetFormatterBehavior(value DateFormatterBehavior) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFormatterBehavior:"), value)
}/* debug [instance_properties/setter]: formatterBehavior */


// The capitalization formatting context used when formatting a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formattingContext
func (d_ DateFormatter) FormattingContext() FormattingContext {
	rv := objc.Send[FormattingContext](d_.ID, objc.Sel("formattingContext"))
	return rv
}/* debug [instance_properties/getter]: formattingContext */


// The capitalization formatting context used when formatting a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/formattingContext
func (d_ DateFormatter) SetFormattingContext(value FormattingContext) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFormattingContext:"), value)
}/* debug [instance_properties/setter]: formattingContext */


// Indicates whether the formatter generates the deprecated calendar date type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/generatesCalendarDates
func (d_ DateFormatter) GeneratesCalendarDates() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("generatesCalendarDates"))
	return rv
}/* debug [instance_properties/getter]: generatesCalendarDates */


// Indicates whether the formatter generates the deprecated calendar date type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/generatesCalendarDates
func (d_ DateFormatter) SetGeneratesCalendarDates(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setGeneratesCalendarDates:"), value)
}/* debug [instance_properties/setter]: generatesCalendarDates */


// The start date of the Gregorian calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/gregorianStartDate
func (d_ DateFormatter) GregorianStartDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("gregorianStartDate"))
	return rv
}/* debug [instance_properties/getter]: gregorianStartDate */


// The start date of the Gregorian calendar for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/gregorianStartDate
func (d_ DateFormatter) SetGregorianStartDate(value IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setGregorianStartDate:"), value)
}/* debug [instance_properties/setter]: gregorianStartDate */


// A Boolean value that indicates whether the receiver uses heuristics when parsing a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/isLenient
func (d_ DateFormatter) Lenient() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("lenient"))
	return rv
}/* debug [instance_properties/getter]: lenient */


// A Boolean value that indicates whether the receiver uses heuristics when parsing a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/isLenient
func (d_ DateFormatter) SetLenient(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLenient:"), value)
}/* debug [instance_properties/setter]: lenient */


// The locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/locale
func (d_ DateFormatter) Locale() ILocale {
	rv := objc.Send[Locale](d_.ID, objc.Sel("locale"))
	return rv
}/* debug [instance_properties/getter]: locale */


// The locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/locale
func (d_ DateFormatter) SetLocale(value ILocale) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocale:"), value)
}/* debug [instance_properties/setter]: locale */


// The long era symbols for the receiver
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/longEraSymbols
func (d_ DateFormatter) LongEraSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("longEraSymbols"))
	return rv
}/* debug [instance_properties/getter]: longEraSymbols */


// The long era symbols for the receiver
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/longEraSymbols
func (d_ DateFormatter) SetLongEraSymbols(value []string) {
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
}/* debug [instance_properties/setter]: longEraSymbols */


// The month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/monthSymbols
func (d_ DateFormatter) MonthSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("monthSymbols"))
	return rv
}/* debug [instance_properties/getter]: monthSymbols */


// The month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/monthSymbols
func (d_ DateFormatter) SetMonthSymbols(value []string) {
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
}/* debug [instance_properties/setter]: monthSymbols */


// The PM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/pmSymbol
func (d_ DateFormatter) PMSymbol() IString {
	rv := objc.Send[String](d_.ID, objc.Sel("PMSymbol"))
	return rv
}/* debug [instance_properties/getter]: PMSymbol */


// The PM symbol for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/pmSymbol
func (d_ DateFormatter) SetPMSymbol(value IString) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPMSymbol:"), value)
}/* debug [instance_properties/setter]: PMSymbol */


// The quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/quarterSymbols
func (d_ DateFormatter) QuarterSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("quarterSymbols"))
	return rv
}/* debug [instance_properties/getter]: quarterSymbols */


// The quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/quarterSymbols
func (d_ DateFormatter) SetQuarterSymbols(value []string) {
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
}/* debug [instance_properties/setter]: quarterSymbols */


// The array of short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortMonthSymbols
func (d_ DateFormatter) ShortMonthSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("shortMonthSymbols"))
	return rv
}/* debug [instance_properties/getter]: shortMonthSymbols */


// The array of short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortMonthSymbols
func (d_ DateFormatter) SetShortMonthSymbols(value []string) {
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
}/* debug [instance_properties/setter]: shortMonthSymbols */


// The short quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortQuarterSymbols
func (d_ DateFormatter) ShortQuarterSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("shortQuarterSymbols"))
	return rv
}/* debug [instance_properties/getter]: shortQuarterSymbols */


// The short quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortQuarterSymbols
func (d_ DateFormatter) SetShortQuarterSymbols(value []string) {
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
}/* debug [instance_properties/setter]: shortQuarterSymbols */


// The short standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneMonthSymbols
func (d_ DateFormatter) ShortStandaloneMonthSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("shortStandaloneMonthSymbols"))
	return rv
}/* debug [instance_properties/getter]: shortStandaloneMonthSymbols */


// The short standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneMonthSymbols
func (d_ DateFormatter) SetShortStandaloneMonthSymbols(value []string) {
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
}/* debug [instance_properties/setter]: shortStandaloneMonthSymbols */


// The short standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneQuarterSymbols
func (d_ DateFormatter) ShortStandaloneQuarterSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("shortStandaloneQuarterSymbols"))
	return rv
}/* debug [instance_properties/getter]: shortStandaloneQuarterSymbols */


// The short standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneQuarterSymbols
func (d_ DateFormatter) SetShortStandaloneQuarterSymbols(value []string) {
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
}/* debug [instance_properties/setter]: shortStandaloneQuarterSymbols */


// The array of short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneWeekdaySymbols
func (d_ DateFormatter) ShortStandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("shortStandaloneWeekdaySymbols"))
	return rv
}/* debug [instance_properties/getter]: shortStandaloneWeekdaySymbols */


// The array of short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneWeekdaySymbols
func (d_ DateFormatter) SetShortStandaloneWeekdaySymbols(value []string) {
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
}/* debug [instance_properties/setter]: shortStandaloneWeekdaySymbols */


// The array of short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortWeekdaySymbols
func (d_ DateFormatter) ShortWeekdaySymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("shortWeekdaySymbols"))
	return rv
}/* debug [instance_properties/getter]: shortWeekdaySymbols */


// The array of short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/shortWeekdaySymbols
func (d_ DateFormatter) SetShortWeekdaySymbols(value []string) {
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
}/* debug [instance_properties/setter]: shortWeekdaySymbols */


// The standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneMonthSymbols
func (d_ DateFormatter) StandaloneMonthSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("standaloneMonthSymbols"))
	return rv
}/* debug [instance_properties/getter]: standaloneMonthSymbols */


// The standalone month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneMonthSymbols
func (d_ DateFormatter) SetStandaloneMonthSymbols(value []string) {
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
}/* debug [instance_properties/setter]: standaloneMonthSymbols */


// The standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneQuarterSymbols
func (d_ DateFormatter) StandaloneQuarterSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("standaloneQuarterSymbols"))
	return rv
}/* debug [instance_properties/getter]: standaloneQuarterSymbols */


// The standalone quarter symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneQuarterSymbols
func (d_ DateFormatter) SetStandaloneQuarterSymbols(value []string) {
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
}/* debug [instance_properties/setter]: standaloneQuarterSymbols */


// The array of standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneWeekdaySymbols
func (d_ DateFormatter) StandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("standaloneWeekdaySymbols"))
	return rv
}/* debug [instance_properties/getter]: standaloneWeekdaySymbols */


// The array of standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneWeekdaySymbols
func (d_ DateFormatter) SetStandaloneWeekdaySymbols(value []string) {
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
}/* debug [instance_properties/setter]: standaloneWeekdaySymbols */


// The time style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/timeStyle
func (d_ DateFormatter) TimeStyle() DateFormatterStyle {
	rv := objc.Send[DateFormatterStyle](d_.ID, objc.Sel("timeStyle"))
	return rv
}/* debug [instance_properties/getter]: timeStyle */


// The time style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/timeStyle
func (d_ DateFormatter) SetTimeStyle(value DateFormatterStyle) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeStyle:"), value)
}/* debug [instance_properties/setter]: timeStyle */


// The time zone for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/timeZone
func (d_ DateFormatter) TimeZone() ITimeZone {
	rv := objc.Send[TimeZone](d_.ID, objc.Sel("timeZone"))
	return rv
}/* debug [instance_properties/getter]: timeZone */


// The time zone for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/timeZone
func (d_ DateFormatter) SetTimeZone(value ITimeZone) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
}/* debug [instance_properties/setter]: timeZone */


// The earliest date that can be denoted by a two-digit year specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/twoDigitStartDate
func (d_ DateFormatter) TwoDigitStartDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("twoDigitStartDate"))
	return rv
}/* debug [instance_properties/getter]: twoDigitStartDate */


// The earliest date that can be denoted by a two-digit year specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/twoDigitStartDate
func (d_ DateFormatter) SetTwoDigitStartDate(value IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTwoDigitStartDate:"), value)
}/* debug [instance_properties/setter]: twoDigitStartDate */


// The very short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortMonthSymbols
func (d_ DateFormatter) VeryShortMonthSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("veryShortMonthSymbols"))
	return rv
}/* debug [instance_properties/getter]: veryShortMonthSymbols */


// The very short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortMonthSymbols
func (d_ DateFormatter) SetVeryShortMonthSymbols(value []string) {
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
}/* debug [instance_properties/setter]: veryShortMonthSymbols */


// The very short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortStandaloneMonthSymbols
func (d_ DateFormatter) VeryShortStandaloneMonthSymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("veryShortStandaloneMonthSymbols"))
	return rv
}/* debug [instance_properties/getter]: veryShortStandaloneMonthSymbols */


// The very short month symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortStandaloneMonthSymbols
func (d_ DateFormatter) SetVeryShortStandaloneMonthSymbols(value []string) {
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
}/* debug [instance_properties/setter]: veryShortStandaloneMonthSymbols */


// The array of very short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortStandaloneWeekdaySymbols
func (d_ DateFormatter) VeryShortStandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("veryShortStandaloneWeekdaySymbols"))
	return rv
}/* debug [instance_properties/getter]: veryShortStandaloneWeekdaySymbols */


// The array of very short standalone weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortStandaloneWeekdaySymbols
func (d_ DateFormatter) SetVeryShortStandaloneWeekdaySymbols(value []string) {
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
}/* debug [instance_properties/setter]: veryShortStandaloneWeekdaySymbols */


// The array of very short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortWeekdaySymbols
func (d_ DateFormatter) VeryShortWeekdaySymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("veryShortWeekdaySymbols"))
	return rv
}/* debug [instance_properties/getter]: veryShortWeekdaySymbols */


// The array of very short weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortWeekdaySymbols
func (d_ DateFormatter) SetVeryShortWeekdaySymbols(value []string) {
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
}/* debug [instance_properties/setter]: veryShortWeekdaySymbols */


// The array of weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/weekdaySymbols
func (d_ DateFormatter) WeekdaySymbols() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("weekdaySymbols"))
	return rv
}/* debug [instance_properties/getter]: weekdaySymbols */


// The array of weekday symbols for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateFormatter/weekdaySymbols
func (d_ DateFormatter) SetWeekdaySymbols(value []string) {
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
}/* debug [instance_properties/setter]: weekdaySymbols */


// A Boolean value that indicates whether the receiver uses heuristics when parsing a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/islenient
func (d_ DateFormatter) IsLenient() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isLenient"))
	return rv
}/* debug [instance_properties/getter]: isLenient */


// A Boolean value that indicates whether the receiver uses heuristics when parsing a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/islenient
func (d_ DateFormatter) SetIsLenient(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsLenient:"), value)
}/* debug [instance_properties/setter]: isLenient */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDateFormatter */


