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
	ComponentsFromDate(unitFlags unsafe.Pointer, date unsafe.Pointer) unsafe.Pointer
	DateFromComponents(comps unsafe.Pointer) unsafe.Pointer
	IsDateInToday(date unsafe.Pointer) bool
	MinimumRangeOfUnit(unit unsafe.Pointer) Range
	NextDateAfterDateMatchingUnitValueOptions(date unsafe.Pointer, unit unsafe.Pointer, value int, options unsafe.Pointer) unsafe.Pointer
	NextDateAfterDateMatchingHourMinuteSecondOptions(date unsafe.Pointer, hourValue int, minuteValue int, secondValue int, options unsafe.Pointer) unsafe.Pointer
	OrdinalityOfUnitInUnitForDate(smaller unsafe.Pointer, larger unsafe.Pointer, date unsafe.Pointer) uint
	RangeOfUnitInUnitForDate(smaller unsafe.Pointer, larger unsafe.Pointer, date unsafe.Pointer) Range
	RangeOfWeekendStartDateIntervalContainingDate(datep unsafe.Pointer, tip TimeInterval, date unsafe.Pointer) bool
}

// A definition of the relationships between calendar units and absolute points in time, providing features for calculation and comparison of dates.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. objects encapsulate information about systems of reckoning time in which the beginning, length, and divisions of a year are defined. They provide information about the calendar and support for calendrical computations such as determining the range of a given calendrical unit and adding units to a given absolute time. is with its Core Foundation counterpart, . See for more information on toll-free bridging.
//
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/init(calendarIdentifier:)
func NewCalendarWithCalendarIdentifier(ident unsafe.Pointer) Calendar {
	instance := getCalendarClass().Alloc()
	rv := objc.Send[Calendar](instance.ID, objc.Sel("initWithCalendarIdentifier:"), ident)
	rv.Autorelease()
	return rv
}



// Creates a new calendar specified by a given identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/init(identifier:)
func NewCalendarWithIdentifier(calendarIdentifierConstant unsafe.Pointer) Calendar {
	rv := objc.Send[Calendar](objc.ID(getCalendarClass().class), objc.Sel("calendarWithIdentifier:"), calendarIdentifierConstant)
	return rv
}


// Creates a new calendar specified by a given identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/init(identifier:)
func (cc _CalendarClass) CalendarWithIdentifier(calendarIdentifierConstant unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("calendarWithIdentifier:"), calendarIdentifierConstant)
	return rv
}

// A calendar that tracks changes to user’s preferred calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/autoupdatingCurrent
func (cc _CalendarClass) AutoupdatingCurrentCalendar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("autoupdatingCurrentCalendar"))
	return rv
}
// The user’s current calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/current
func (cc _CalendarClass) CurrentCalendar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("currentCalendar"))
	return rv
}
// Returns the date components representing a given date.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/components(_:from:)
func (c_ Calendar) ComponentsFromDate(unitFlags unsafe.Pointer, date unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("components:fromDate:"), unitFlags, date)
	return rv
}

// Returns a date representing the absolute time calculated from given components.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(from:)
func (c_ Calendar) DateFromComponents(comps unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("dateFromComponents:"), comps)
	return rv
}

// Indicates whether the given date is in “today.”
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/isDateInToday(_:)
func (c_ Calendar) IsDateInToday(date unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDateInToday:"), date)
	return rv
}

// Returns the minimum range limits of the values that a given unit can take on.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/minimumRange(of:)
func (c_ Calendar) MinimumRangeOfUnit(unit unsafe.Pointer) Range {
	rv := objc.Send[Range](c_.ID, objc.Sel("minimumRangeOfUnit:"), unit)
	return rv
}

// Returns the next date after a given date matching the given calendar unit value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/nextDate(after:matching:value:options:)
func (c_ Calendar) NextDateAfterDateMatchingUnitValueOptions(date unsafe.Pointer, unit unsafe.Pointer, value int, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("nextDateAfterDate:matchingUnit:value:options:"), date, unit, value, options)
	return rv
}

// Returns the next date after a given date that matches the given hour, minute, and second, component values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/nextDate(after:matchingHour:minute:second:options:)
func (c_ Calendar) NextDateAfterDateMatchingHourMinuteSecondOptions(date unsafe.Pointer, hourValue int, minuteValue int, secondValue int, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("nextDateAfterDate:matchingHour:minute:second:options:"), date, hourValue, minuteValue, secondValue, options)
	return rv
}

// Returns, for a given absolute time, the ordinal number of a smaller calendar unit (such as a day) within a specified larger calendar unit (such as a week).
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/ordinality(of:in:for:)
func (c_ Calendar) OrdinalityOfUnitInUnitForDate(smaller unsafe.Pointer, larger unsafe.Pointer, date unsafe.Pointer) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("ordinalityOfUnit:inUnit:forDate:"), smaller, larger, date)
	return rv
}

// Returns the range of absolute time values that a smaller calendar unit (such as a day) can take on in a larger calendar unit (such as a month) that includes a specified absolute time.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/range(of:in:for:)
func (c_ Calendar) RangeOfUnitInUnitForDate(smaller unsafe.Pointer, larger unsafe.Pointer, date unsafe.Pointer) Range {
	rv := objc.Send[Range](c_.ID, objc.Sel("rangeOfUnit:inUnit:forDate:"), smaller, larger, date)
	return rv
}

// Returns whether a given date falls within a weekend period, and if so, returns by reference the start date and time interval of the weekend range.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/range(ofWeekendStart:interval:containing:)
func (c_ Calendar) RangeOfWeekendStartDateIntervalContainingDate(datep unsafe.Pointer, tip TimeInterval, date unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("rangeOfWeekendStartDate:interval:containingDate:"), datep, tip, date)
	return rv
}

// The symbol used to represent “AM” for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/amsymbol
func (c_ Calendar) AmSymbol() string {
	rv := objc.Send[string](c_.ID, objc.Sel("amSymbol"))
	return rv
}


// SetAmSymbol sets the value of the amSymbol property.
// The symbol used to represent “AM” for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/amsymbol
func (c_ Calendar) SetAmSymbol(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAmSymbol:"), objc.String(value))
}

// A list of era symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/erasymbols
func (c_ Calendar) EraSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("eraSymbols"))
	return rv
}


// SetEraSymbols sets the value of the eraSymbols property.
// A list of era symbols for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/erasymbols
func (c_ Calendar) SetEraSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEraSymbols:"), objc.String(value))
}

// A list of long era symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/longerasymbols
func (c_ Calendar) LongEraSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("longEraSymbols"))
	return rv
}


// SetLongEraSymbols sets the value of the longEraSymbols property.
// A list of long era symbols for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/longerasymbols
func (c_ Calendar) SetLongEraSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongEraSymbols:"), objc.String(value))
}

// The minimum number of days in the first week of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/minimumdaysinfirstweek
func (c_ Calendar) MinimumDaysInFirstWeek() int {
	rv := objc.Send[int](c_.ID, objc.Sel("minimumDaysInFirstWeek"))
	return rv
}


// SetMinimumDaysInFirstWeek sets the value of the minimumDaysInFirstWeek property.
// The minimum number of days in the first week of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/minimumdaysinfirstweek
func (c_ Calendar) SetMinimumDaysInFirstWeek(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumDaysInFirstWeek:"), value)
}

// A list of month symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/monthsymbols
func (c_ Calendar) MonthSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("monthSymbols"))
	return rv
}


// SetMonthSymbols sets the value of the monthSymbols property.
// A list of month symbols for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/monthsymbols
func (c_ Calendar) SetMonthSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMonthSymbols:"), objc.String(value))
}

// The symbol used to represent “PM” for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/pmsymbol
func (c_ Calendar) PmSymbol() string {
	rv := objc.Send[string](c_.ID, objc.Sel("pmSymbol"))
	return rv
}


// SetPmSymbol sets the value of the pmSymbol property.
// The symbol used to represent “PM” for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/pmsymbol
func (c_ Calendar) SetPmSymbol(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPmSymbol:"), objc.String(value))
}

// A list of quarter symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/quartersymbols
func (c_ Calendar) QuarterSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("quarterSymbols"))
	return rv
}


// SetQuarterSymbols sets the value of the quarterSymbols property.
// A list of quarter symbols for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/quartersymbols
func (c_ Calendar) SetQuarterSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQuarterSymbols:"), objc.String(value))
}

// A list of short month symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortmonthsymbols
func (c_ Calendar) ShortMonthSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("shortMonthSymbols"))
	return rv
}


// SetShortMonthSymbols sets the value of the shortMonthSymbols property.
// A list of short month symbols for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortmonthsymbols
func (c_ Calendar) SetShortMonthSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShortMonthSymbols:"), objc.String(value))
}

// A list of short quarter symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortquartersymbols
func (c_ Calendar) ShortQuarterSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("shortQuarterSymbols"))
	return rv
}


// SetShortQuarterSymbols sets the value of the shortQuarterSymbols property.
// A list of short quarter symbols for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortquartersymbols
func (c_ Calendar) SetShortQuarterSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShortQuarterSymbols:"), objc.String(value))
}

// A list of short standalone month symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortstandalonemonthsymbols
func (c_ Calendar) ShortStandaloneMonthSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("shortStandaloneMonthSymbols"))
	return rv
}


// SetShortStandaloneMonthSymbols sets the value of the shortStandaloneMonthSymbols property.
// A list of short standalone month symbols for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortstandalonemonthsymbols
func (c_ Calendar) SetShortStandaloneMonthSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShortStandaloneMonthSymbols:"), objc.String(value))
}

// A list of shorter-named weekdays in this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortweekdaysymbols
func (c_ Calendar) ShortWeekdaySymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("shortWeekdaySymbols"))
	return rv
}


// SetShortWeekdaySymbols sets the value of the shortWeekdaySymbols property.
// A list of shorter-named weekdays in this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortweekdaysymbols
func (c_ Calendar) SetShortWeekdaySymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShortWeekdaySymbols:"), objc.String(value))
}

// A list of standalone month symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standalonemonthsymbols
func (c_ Calendar) StandaloneMonthSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("standaloneMonthSymbols"))
	return rv
}


// SetStandaloneMonthSymbols sets the value of the standaloneMonthSymbols property.
// A list of standalone month symbols for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standalonemonthsymbols
func (c_ Calendar) SetStandaloneMonthSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStandaloneMonthSymbols:"), objc.String(value))
}

// A list of standalone quarter symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standalonequartersymbols
func (c_ Calendar) StandaloneQuarterSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("standaloneQuarterSymbols"))
	return rv
}


// SetStandaloneQuarterSymbols sets the value of the standaloneQuarterSymbols property.
// A list of standalone quarter symbols for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standalonequartersymbols
func (c_ Calendar) SetStandaloneQuarterSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStandaloneQuarterSymbols:"), objc.String(value))
}

// A list of standalone weekday symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standaloneweekdaysymbols
func (c_ Calendar) StandaloneWeekdaySymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("standaloneWeekdaySymbols"))
	return rv
}


// SetStandaloneWeekdaySymbols sets the value of the standaloneWeekdaySymbols property.
// A list of standalone weekday symbols for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standaloneweekdaysymbols
func (c_ Calendar) SetStandaloneWeekdaySymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStandaloneWeekdaySymbols:"), objc.String(value))
}

// A list of very short month symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortmonthsymbols
func (c_ Calendar) VeryShortMonthSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("veryShortMonthSymbols"))
	return rv
}


// SetVeryShortMonthSymbols sets the value of the veryShortMonthSymbols property.
// A list of very short month symbols for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortmonthsymbols
func (c_ Calendar) SetVeryShortMonthSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVeryShortMonthSymbols:"), objc.String(value))
}

// A list of very short month symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortstandalonemonthsymbols
func (c_ Calendar) VeryShortStandaloneMonthSymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("veryShortStandaloneMonthSymbols"))
	return rv
}


// SetVeryShortStandaloneMonthSymbols sets the value of the veryShortStandaloneMonthSymbols property.
// A list of very short month symbols for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortstandalonemonthsymbols
func (c_ Calendar) SetVeryShortStandaloneMonthSymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVeryShortStandaloneMonthSymbols:"), objc.String(value))
}

// A list of very short standalone weekday symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortstandaloneweekdaysymbols
func (c_ Calendar) VeryShortStandaloneWeekdaySymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("veryShortStandaloneWeekdaySymbols"))
	return rv
}


// SetVeryShortStandaloneWeekdaySymbols sets the value of the veryShortStandaloneWeekdaySymbols property.
// A list of very short standalone weekday symbols for this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortstandaloneweekdaysymbols
func (c_ Calendar) SetVeryShortStandaloneWeekdaySymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVeryShortStandaloneWeekdaySymbols:"), objc.String(value))
}

// A list of very-shortly-named weekdays in this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortweekdaysymbols
func (c_ Calendar) VeryShortWeekdaySymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("veryShortWeekdaySymbols"))
	return rv
}


// SetVeryShortWeekdaySymbols sets the value of the veryShortWeekdaySymbols property.
// A list of very-shortly-named weekdays in this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortweekdaysymbols
func (c_ Calendar) SetVeryShortWeekdaySymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVeryShortWeekdaySymbols:"), objc.String(value))
}

// A list of weekdays in this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/weekdaysymbols
func (c_ Calendar) WeekdaySymbols() string {
	rv := objc.Send[string](c_.ID, objc.Sel("weekdaySymbols"))
	return rv
}


// SetWeekdaySymbols sets the value of the weekdaySymbols property.
// A list of weekdays in this calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/weekdaysymbols
func (c_ Calendar) SetWeekdaySymbols(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeekdaySymbols:"), objc.String(value))
}

// A calendar that tracks changes to user’s preferred calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/autoupdatingCurrent
func (c_ Calendar) AutoupdatingCurrentCalendar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("autoupdatingCurrentCalendar"))
	return rv
}

// An identifier for the calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/calendarIdentifier
func (c_ Calendar) CalendarIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("calendarIdentifier"))
	return rv
}

// The user’s current calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/current
func (c_ Calendar) CurrentCalendar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("currentCalendar"))
	return rv
}

// The index of the first weekday of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/firstWeekday
func (c_ Calendar) FirstWeekday() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("firstWeekday"))
	return rv
}


// SetFirstWeekday sets the value of the firstWeekday property.
// The index of the first weekday of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/firstWeekday
func (c_ Calendar) SetFirstWeekday(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFirstWeekday:"), value)
}

// The locale of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/locale
func (c_ Calendar) Locale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("locale"))
	return rv
}


// SetLocale sets the value of the locale property.
// The locale of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/locale
func (c_ Calendar) SetLocale(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocale:"), value)
}

// A list of short standalone quarter symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/shortStandaloneQuarterSymbols
func (c_ Calendar) ShortStandaloneQuarterSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("shortStandaloneQuarterSymbols"))
	return rv
}

// A list of short standalone weekday symbols for this calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/shortStandaloneWeekdaySymbols
func (c_ Calendar) ShortStandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("shortStandaloneWeekdaySymbols"))
	return rv
}

// The time zone for the calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/timeZone
func (c_ Calendar) TimeZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("timeZone"))
	return rv
}


// SetTimeZone sets the value of the timeZone property.
// The time zone for the calendar.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/timeZone
func (c_ Calendar) SetTimeZone(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeZone:"), value)
}


