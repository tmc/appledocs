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

// An identifier for the calendar.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/calendarIdentifier
func (c_ Calendar) CalendarIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("calendarIdentifier"))
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

