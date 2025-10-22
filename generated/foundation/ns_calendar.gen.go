// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Calendar] class.
var (
	CalendarClass     _CalendarClass
	CalendarClassOnce sync.Once
)

func getCalendarClass() _CalendarClass {
	CalendarClassOnce.Do(func() {
		CalendarClass = _CalendarClass{objc.GetClass("NSCalendar")}
	})
	return CalendarClass
}

type _CalendarClass struct {
	class objc.Class
}

// An interface definition for the [Calendar] class.
type ICalendar interface {
	objectivec.IObject
	ComponentsFromDate(unitFlags ICalendarUnit, date IDate) DateComponents
	DateFromComponents(comps IDateComponents) Date
	IsDateInToday(date IDate) bool
	MinimumRangeOfUnit(unit ICalendarUnit) Range
	NextDateAfterDateMatchingUnitValueOptions(date IDate, unit ICalendarUnit, value int, options unsafe.Pointer) Date
	NextDateAfterDateMatchingHourMinuteSecondOptions(date IDate, hourValue int, minuteValue int, secondValue int, options unsafe.Pointer) Date
	OrdinalityOfUnitInUnitForDate(smaller ICalendarUnit, larger ICalendarUnit, date IDate) uint
	RangeOfUnitInUnitForDate(smaller ICalendarUnit, larger ICalendarUnit, date IDate) Range
	RangeOfWeekendStartDateIntervalContainingDate(datep IDate, tip ITimeInterval, date IDate) bool
	CalendarIdentifier() CalendarIdentifier
	FirstWeekday() uint
	SetFirstWeekday(value uint)
	Locale() NSLocale
	SetLocale(value ILocale)
	ShortStandaloneQuarterSymbols() []string
	ShortStandaloneWeekdaySymbols() []string
	TimeZone() NSTimeZone
	SetTimeZone(value ITimeZone)
	AmSymbol() string
	SetAmSymbol(value string)
	EraSymbols() string
	SetEraSymbols(value string)
	LongEraSymbols() string
	SetLongEraSymbols(value string)
	MinimumDaysInFirstWeek() int
	SetMinimumDaysInFirstWeek(value int)
	MonthSymbols() string
	SetMonthSymbols(value string)
	PmSymbol() string
	SetPmSymbol(value string)
	QuarterSymbols() string
	SetQuarterSymbols(value string)
	ShortMonthSymbols() string
	SetShortMonthSymbols(value string)
	ShortQuarterSymbols() string
	SetShortQuarterSymbols(value string)
	ShortStandaloneMonthSymbols() string
	SetShortStandaloneMonthSymbols(value string)
	ShortWeekdaySymbols() string
	SetShortWeekdaySymbols(value string)
	StandaloneMonthSymbols() string
	SetStandaloneMonthSymbols(value string)
	StandaloneQuarterSymbols() string
	SetStandaloneQuarterSymbols(value string)
	StandaloneWeekdaySymbols() string
	SetStandaloneWeekdaySymbols(value string)
	VeryShortMonthSymbols() string
	SetVeryShortMonthSymbols(value string)
	VeryShortStandaloneMonthSymbols() string
	SetVeryShortStandaloneMonthSymbols(value string)
	VeryShortStandaloneWeekdaySymbols() string
	SetVeryShortStandaloneWeekdaySymbols(value string)
	VeryShortWeekdaySymbols() string
	SetVeryShortWeekdaySymbols(value string)
	WeekdaySymbols() string
	SetWeekdaySymbols(value string)
}

// A definition of the relationships between calendar units and absolute points in time, providing features for calculation and comparison of dates.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. objects encapsulate information about systems of reckoning time in which the beginning, length, and divisions of a year are defined. They provide information about the calendar and support for calendrical computations such as determining the range of a given calendrical unit and adding units to a given absolute time. is with its Core Foundation counterpart, . See for more information on toll-free bridging.


// A definition of the relationships between calendar units and absolute points in time, providing features for calculation and comparison of dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar

type Calendar struct {
	objectivec.Object
}

// CalendarFrom constructs a [Calendar] from an unsafe.Pointer.
//
// A definition of the relationships between calendar units and absolute points in time, providing features for calculation and comparison of dates.
func CalendarFrom(ptr unsafe.Pointer) Calendar {
	return Calendar{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CalendarClass) Alloc() Calendar {
	rv := objc.Send[Calendar](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CalendarClass) New() Calendar {
	rv := objc.Send[Calendar](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Calendar) Init() Calendar {
	rv := objc.Send[Calendar](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Calendar) Autorelease() Calendar {
	rv := objc.Send[Calendar](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCalendar creates a new Calendar instance.
func NewCalendar() Calendar {
	return getCalendarClass().New()
}




// Initializes a calendar according to a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/init(calendarIdentifier:)

func NewCalendarWithCalendarIdentifier(ident ICalendarIdentifier) Calendar {
	instance := getCalendarClass().Alloc()
	rv := objc.Send[Calendar](instance.ID, objc.Sel("initWithCalendarIdentifier:"), ident)
	rv.Autorelease()
	return rv
}



// Creates a new calendar specified by a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/init(identifier:)

func NewCalendarWithIdentifier(calendarIdentifierConstant ICalendarIdentifier) Calendar {
	rv := objc.Send[Calendar](objc.ID(getCalendarClass().class), objc.Sel("calendarWithIdentifier:"), calendarIdentifierConstant)
	return rv
}



// Creates a new calendar specified by a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/init(identifier:)

func (cc _CalendarClass) CalendarWithIdentifier(calendarIdentifierConstant ICalendarIdentifier) Calendar {
	rv := objc.Send[Calendar](objc.ID(cc.class), objc.Sel("calendarWithIdentifier:"), calendarIdentifierConstant)
	return rv
}


// A calendar that tracks changes to user’s preferred calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/autoupdatingCurrent

func (cc _CalendarClass) AutoupdatingCurrentCalendar() Calendar {
	rv := objc.Send[NSCalendar](objc.ID(cc.class), objc.Sel("autoupdatingCurrentCalendar"))
	return rv
}

// The user’s current calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/current

func (cc _CalendarClass) CurrentCalendar() Calendar {
	rv := objc.Send[NSCalendar](objc.ID(cc.class), objc.Sel("currentCalendar"))
	return rv
}


// Returns the date components representing a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/components(_:from:)

func (c_ Calendar) ComponentsFromDate(unitFlags ICalendarUnit, date IDate) DateComponents {
	rv := objc.Send[DateComponents](c_.ID, objc.Sel("components:fromDate:"), unitFlags, date)
	return rv
}



// Returns a date representing the absolute time calculated from given components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(from:)

func (c_ Calendar) DateFromComponents(comps IDateComponents) Date {
	rv := objc.Send[Date](c_.ID, objc.Sel("dateFromComponents:"), comps)
	return rv
}



// Indicates whether the given date is in “today.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/isDateInToday(_:)

func (c_ Calendar) IsDateInToday(date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDateInToday:"), date)
	return rv
}



// Returns the minimum range limits of the values that a given unit can take on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/minimumRange(of:)

func (c_ Calendar) MinimumRangeOfUnit(unit ICalendarUnit) Range {
	rv := objc.Send[Range](c_.ID, objc.Sel("minimumRangeOfUnit:"), unit)
	return rv
}



// Returns the next date after a given date matching the given calendar unit value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/nextDate(after:matching:value:options:)

func (c_ Calendar) NextDateAfterDateMatchingUnitValueOptions(date IDate, unit ICalendarUnit, value int, options unsafe.Pointer) Date {
	rv := objc.Send[Date](c_.ID, objc.Sel("nextDateAfterDate:matchingUnit:value:options:"), date, unit, value, options)
	return rv
}



// Returns the next date after a given date that matches the given hour, minute, and second, component values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/nextDate(after:matchingHour:minute:second:options:)

func (c_ Calendar) NextDateAfterDateMatchingHourMinuteSecondOptions(date IDate, hourValue int, minuteValue int, secondValue int, options unsafe.Pointer) Date {
	rv := objc.Send[Date](c_.ID, objc.Sel("nextDateAfterDate:matchingHour:minute:second:options:"), date, hourValue, minuteValue, secondValue, options)
	return rv
}



// Returns, for a given absolute time, the ordinal number of a smaller calendar unit (such as a day) within a specified larger calendar unit (such as a week).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/ordinality(of:in:for:)

func (c_ Calendar) OrdinalityOfUnitInUnitForDate(smaller ICalendarUnit, larger ICalendarUnit, date IDate) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("ordinalityOfUnit:inUnit:forDate:"), smaller, larger, date)
	return rv
}



// Returns the range of absolute time values that a smaller calendar unit (such as a day) can take on in a larger calendar unit (such as a month) that includes a specified absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/range(of:in:for:)

func (c_ Calendar) RangeOfUnitInUnitForDate(smaller ICalendarUnit, larger ICalendarUnit, date IDate) Range {
	rv := objc.Send[Range](c_.ID, objc.Sel("rangeOfUnit:inUnit:forDate:"), smaller, larger, date)
	return rv
}



// Returns whether a given date falls within a weekend period, and if so, returns by reference the start date and time interval of the weekend range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/range(ofWeekendStart:interval:containing:)

func (c_ Calendar) RangeOfWeekendStartDateIntervalContainingDate(datep IDate, tip ITimeInterval, date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("rangeOfWeekendStartDate:interval:containingDate:"), datep, tip, date)
	return rv
}


// A calendar that tracks changes to user’s preferred calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/autoupdatingCurrent

func (c_ Calendar) AutoupdatingCurrentCalendar() NSCalendar {
	rv := objc.Send[NSCalendar](c_.ID, objc.Sel("autoupdatingCurrentCalendar"))
	return rv
}


// An identifier for the calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/calendarIdentifier

func (c_ Calendar) CalendarIdentifier() CalendarIdentifier {
	rv := objc.Send[CalendarIdentifier](c_.ID, objc.Sel("calendarIdentifier"))
	return rv
}


// The user’s current calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/current

func (c_ Calendar) CurrentCalendar() NSCalendar {
	rv := objc.Send[NSCalendar](c_.ID, objc.Sel("currentCalendar"))
	return rv
}


// The index of the first weekday of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/firstWeekday

func (c_ Calendar) FirstWeekday() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("firstWeekday"))
	return rv
}


// The index of the first weekday of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/firstWeekday

func (c_ Calendar) SetFirstWeekday(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFirstWeekday:"), value)
}


// The locale of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/locale

func (c_ Calendar) Locale() NSLocale {
	rv := objc.Send[NSLocale](c_.ID, objc.Sel("locale"))
	return rv
}


// The locale of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/locale

func (c_ Calendar) SetLocale(value ILocale) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocale:"), value)
}


// A list of short standalone quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/shortStandaloneQuarterSymbols

func (c_ Calendar) ShortStandaloneQuarterSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("shortStandaloneQuarterSymbols"))
	return rv
}


// A list of short standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/shortStandaloneWeekdaySymbols

func (c_ Calendar) ShortStandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("shortStandaloneWeekdaySymbols"))
	return rv
}


// The time zone for the calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/timeZone

func (c_ Calendar) TimeZone() NSTimeZone {
	rv := objc.Send[NSTimeZone](c_.ID, objc.Sel("timeZone"))
	return rv
}


// The time zone for the calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/timeZone

func (c_ Calendar) SetTimeZone(value ITimeZone) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeZone:"), value)
}


// The symbol used to represent “AM” for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/amsymbol

func (c_ Calendar) AmSymbol() string {
	rv := objc.Send[string](c_.ID, objc.Sel("amSymbol"))
	return rv
}


// The symbol used to represent “AM” for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/amsymbol

func (c_ Calendar) SetAmSymbol(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAmSymbol:"), objc.String(value))
}


// A list of era symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/erasymbols

func (c_ Calendar) EraSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("eraSymbols"))
	return rv
}


// A list of era symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/erasymbols

func (c_ Calendar) SetEraSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEraSymbols:"), objc.String(value))
}


// A list of long era symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/longerasymbols

func (c_ Calendar) LongEraSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("longEraSymbols"))
	return rv
}


// A list of long era symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/longerasymbols

func (c_ Calendar) SetLongEraSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongEraSymbols:"), objc.String(value))
}


// The minimum number of days in the first week of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/minimumdaysinfirstweek

func (c_ Calendar) MinimumDaysInFirstWeek() int {
	rv := objc.Send[int](c_.ID, objc.Sel("minimumDaysInFirstWeek"))
	return rv
}


// The minimum number of days in the first week of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/minimumdaysinfirstweek

func (c_ Calendar) SetMinimumDaysInFirstWeek(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumDaysInFirstWeek:"), value)
}


// A list of month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/monthsymbols

func (c_ Calendar) MonthSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("monthSymbols"))
	return rv
}


// A list of month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/monthsymbols

func (c_ Calendar) SetMonthSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMonthSymbols:"), objc.String(value))
}


// The symbol used to represent “PM” for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/pmsymbol

func (c_ Calendar) PmSymbol() string {
	rv := objc.Send[string](c_.ID, objc.Sel("pmSymbol"))
	return rv
}


// The symbol used to represent “PM” for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/pmsymbol

func (c_ Calendar) SetPmSymbol(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPmSymbol:"), objc.String(value))
}


// A list of quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/quartersymbols

func (c_ Calendar) QuarterSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("quarterSymbols"))
	return rv
}


// A list of quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/quartersymbols

func (c_ Calendar) SetQuarterSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQuarterSymbols:"), objc.String(value))
}


// A list of short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortmonthsymbols

func (c_ Calendar) ShortMonthSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("shortMonthSymbols"))
	return rv
}


// A list of short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortmonthsymbols

func (c_ Calendar) SetShortMonthSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShortMonthSymbols:"), objc.String(value))
}


// A list of short quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortquartersymbols

func (c_ Calendar) ShortQuarterSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("shortQuarterSymbols"))
	return rv
}


// A list of short quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortquartersymbols

func (c_ Calendar) SetShortQuarterSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShortQuarterSymbols:"), objc.String(value))
}


// A list of short standalone month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortstandalonemonthsymbols

func (c_ Calendar) ShortStandaloneMonthSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("shortStandaloneMonthSymbols"))
	return rv
}


// A list of short standalone month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortstandalonemonthsymbols

func (c_ Calendar) SetShortStandaloneMonthSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShortStandaloneMonthSymbols:"), objc.String(value))
}


// A list of shorter-named weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortweekdaysymbols

func (c_ Calendar) ShortWeekdaySymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("shortWeekdaySymbols"))
	return rv
}


// A list of shorter-named weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortweekdaysymbols

func (c_ Calendar) SetShortWeekdaySymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShortWeekdaySymbols:"), objc.String(value))
}


// A list of standalone month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standalonemonthsymbols

func (c_ Calendar) StandaloneMonthSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("standaloneMonthSymbols"))
	return rv
}


// A list of standalone month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standalonemonthsymbols

func (c_ Calendar) SetStandaloneMonthSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStandaloneMonthSymbols:"), objc.String(value))
}


// A list of standalone quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standalonequartersymbols

func (c_ Calendar) StandaloneQuarterSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("standaloneQuarterSymbols"))
	return rv
}


// A list of standalone quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standalonequartersymbols

func (c_ Calendar) SetStandaloneQuarterSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStandaloneQuarterSymbols:"), objc.String(value))
}


// A list of standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standaloneweekdaysymbols

func (c_ Calendar) StandaloneWeekdaySymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("standaloneWeekdaySymbols"))
	return rv
}


// A list of standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standaloneweekdaysymbols

func (c_ Calendar) SetStandaloneWeekdaySymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStandaloneWeekdaySymbols:"), objc.String(value))
}


// A list of very short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortmonthsymbols

func (c_ Calendar) VeryShortMonthSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("veryShortMonthSymbols"))
	return rv
}


// A list of very short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortmonthsymbols

func (c_ Calendar) SetVeryShortMonthSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVeryShortMonthSymbols:"), objc.String(value))
}


// A list of very short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortstandalonemonthsymbols

func (c_ Calendar) VeryShortStandaloneMonthSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("veryShortStandaloneMonthSymbols"))
	return rv
}


// A list of very short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortstandalonemonthsymbols

func (c_ Calendar) SetVeryShortStandaloneMonthSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVeryShortStandaloneMonthSymbols:"), objc.String(value))
}


// A list of very short standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortstandaloneweekdaysymbols

func (c_ Calendar) VeryShortStandaloneWeekdaySymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("veryShortStandaloneWeekdaySymbols"))
	return rv
}


// A list of very short standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortstandaloneweekdaysymbols

func (c_ Calendar) SetVeryShortStandaloneWeekdaySymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVeryShortStandaloneWeekdaySymbols:"), objc.String(value))
}


// A list of very-shortly-named weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortweekdaysymbols

func (c_ Calendar) VeryShortWeekdaySymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("veryShortWeekdaySymbols"))
	return rv
}


// A list of very-shortly-named weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortweekdaysymbols

func (c_ Calendar) SetVeryShortWeekdaySymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVeryShortWeekdaySymbols:"), objc.String(value))
}


// A list of weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/weekdaysymbols

func (c_ Calendar) WeekdaySymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("weekdaySymbols"))
	return rv
}


// A list of weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/weekdaysymbols

func (c_ Calendar) SetWeekdaySymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeekdaySymbols:"), objc.String(value))
}


