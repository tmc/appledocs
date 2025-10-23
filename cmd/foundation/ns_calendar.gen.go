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
	CompareDateToDateToUnitGranularity(date1 IDate, date2 IDate, unit NSCalendarUnit) ComparisonResult
	ComponentFromDate(unit NSCalendarUnit, date IDate) int
	ComponentsFromDate(unitFlags NSCalendarUnit, date IDate) DateComponents
	ComponentsFromDateComponentsToDateComponentsOptions(unitFlags NSCalendarUnit, startingDateComp IDateComponents, resultDateComp IDateComponents, options NSCalendarOptions) DateComponents
	ComponentsFromDateToDateOptions(unitFlags NSCalendarUnit, startingDate IDate, resultDate IDate, opts NSCalendarOptions) DateComponents
	ComponentsInTimeZoneFromDate(timezone ITimeZone, date IDate) DateComponents
	DateMatchesComponents(date IDate, components IDateComponents) bool
	DateByAddingComponentsToDateOptions(comps IDateComponents, date IDate, opts NSCalendarOptions) Date
	DateByAddingUnitValueToDateOptions(unit NSCalendarUnit, value int, date IDate, options NSCalendarOptions) Date
	DateBySettingHourMinuteSecondOfDateOptions(h int, m int, s int, date IDate, opts NSCalendarOptions) Date
	DateBySettingUnitValueOfDateOptions(unit NSCalendarUnit, v int, date IDate, opts NSCalendarOptions) Date
	DateWithEraYearMonthDayHourMinuteSecondNanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date
	DateWithEraYearForWeekOfYearWeekOfYearWeekdayHourMinuteSecondNanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date
	DateFromComponents(comps IDateComponents) Date
	EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock(start IDate, comps IDateComponents, opts NSCalendarOptions, block unsafe.Pointer)
	GetEraYearMonthDayFromDate(eraValuePointer unsafe.Pointer, yearValuePointer unsafe.Pointer, monthValuePointer unsafe.Pointer, dayValuePointer unsafe.Pointer, date IDate)
	GetEraYearForWeekOfYearWeekOfYearWeekdayFromDate(eraValuePointer unsafe.Pointer, yearValuePointer unsafe.Pointer, weekValuePointer unsafe.Pointer, weekdayValuePointer unsafe.Pointer, date IDate)
	GetHourMinuteSecondNanosecondFromDate(hourValuePointer unsafe.Pointer, minuteValuePointer unsafe.Pointer, secondValuePointer unsafe.Pointer, nanosecondValuePointer unsafe.Pointer, date IDate)
	IsDateEqualToDateToUnitGranularity(date1 IDate, date2 IDate, unit NSCalendarUnit) bool
	IsDateInSameDayAsDate(date1 IDate, date2 IDate) bool
	IsDateInToday(date IDate) bool
	IsDateInTomorrow(date IDate) bool
	IsDateInWeekend(date IDate) bool
	IsDateInYesterday(date IDate) bool
	MaximumRangeOfUnit(unit NSCalendarUnit) Range
	MinimumRangeOfUnit(unit NSCalendarUnit) Range
	NextDateAfterDateMatchingComponentsOptions(date IDate, comps IDateComponents, options NSCalendarOptions) Date
	NextDateAfterDateMatchingUnitValueOptions(date IDate, unit NSCalendarUnit, value int, options NSCalendarOptions) Date
	NextDateAfterDateMatchingHourMinuteSecondOptions(date IDate, hourValue int, minuteValue int, secondValue int, options NSCalendarOptions) Date
	NextWeekendStartDateIntervalOptionsAfterDate(datep IDate, tip TimeInterval, options NSCalendarOptions, date IDate) bool
	OrdinalityOfUnitInUnitForDate(smaller NSCalendarUnit, larger NSCalendarUnit, date IDate) uint
	RangeOfUnitInUnitForDate(smaller NSCalendarUnit, larger NSCalendarUnit, date IDate) Range
	RangeOfUnitStartDateIntervalForDate(unit NSCalendarUnit, datep IDate, tip TimeInterval, date IDate) bool
	RangeOfWeekendStartDateIntervalContainingDate(datep IDate, tip TimeInterval, date IDate) bool
	StartOfDayForDate(date IDate) Date
	AMSymbol() string
	CalendarIdentifier() CalendarIdentifier
	EraSymbols() []string
	FirstWeekday() uint
	SetFirstWeekday(value uint)
	Locale() NSLocale
	SetLocale(value ILocale)
	LongEraSymbols() []string
	MinimumDaysInFirstWeek() uint
	SetMinimumDaysInFirstWeek(value uint)
	MonthSymbols() []string
	PMSymbol() string
	QuarterSymbols() []string
	ShortMonthSymbols() []string
	ShortQuarterSymbols() []string
	ShortStandaloneMonthSymbols() []string
	ShortStandaloneQuarterSymbols() []string
	ShortStandaloneWeekdaySymbols() []string
	ShortWeekdaySymbols() []string
	StandaloneMonthSymbols() []string
	StandaloneQuarterSymbols() []string
	StandaloneWeekdaySymbols() []string
	TimeZone() NSTimeZone
	SetTimeZone(value ITimeZone)
	VeryShortMonthSymbols() []string
	VeryShortStandaloneMonthSymbols() []string
	VeryShortStandaloneWeekdaySymbols() []string
	VeryShortWeekdaySymbols() []string
	WeekdaySymbols() []string
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
func NewCalendarWithCalendarIdentifier(ident CalendarIdentifier) Calendar {
	instance := getCalendarClass().Alloc()
	rv := objc.Send[Calendar](instance.ID, objc.Sel("initWithCalendarIdentifier:"), ident)
	rv.Autorelease()
	return rv
}


// Creates a new calendar specified by a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/init(identifier:)
func NewCalendarWithIdentifier(calendarIdentifierConstant CalendarIdentifier) Calendar {
	rv := objc.Send[Calendar](objc.ID(getCalendarClass().class), objc.Sel("calendarWithIdentifier:"), calendarIdentifierConstant)
	return rv
}



// Creates a new calendar specified by a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/init(identifier:)
func (cc _CalendarClass) CalendarWithIdentifier(calendarIdentifierConstant CalendarIdentifier) Calendar {
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

// Indicates the ordering of two given dates based on their components down to a given unit granularity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/compare(_:to:toUnitGranularity:)
func (c_ Calendar) CompareDateToDateToUnitGranularity(date1 IDate, date2 IDate, unit NSCalendarUnit) ComparisonResult {
	rv := objc.Send[ComparisonResult](c_.ID, objc.Sel("compareDate:toDate:toUnitGranularity:"), date1, date2, unit)
	return rv
}


// Returns the specified date component from a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/component(_:from:)
func (c_ Calendar) ComponentFromDate(unit NSCalendarUnit, date IDate) int {
	rv := objc.Send[int](c_.ID, objc.Sel("component:fromDate:"), unit, date)
	return rv
}


// Returns the date components representing a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/components(_:from:)
func (c_ Calendar) ComponentsFromDate(unitFlags NSCalendarUnit, date IDate) DateComponents {
	rv := objc.Send[DateComponents](c_.ID, objc.Sel("components:fromDate:"), unitFlags, date)
	return rv
}


// Returns the difference between start and end dates given as date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/components(_:from:to:options:)-49lo8
func (c_ Calendar) ComponentsFromDateComponentsToDateComponentsOptions(unitFlags NSCalendarUnit, startingDateComp IDateComponents, resultDateComp IDateComponents, options NSCalendarOptions) DateComponents {
	rv := objc.Send[DateComponents](c_.ID, objc.Sel("components:fromDateComponents:toDateComponents:options:"), unitFlags, startingDateComp, resultDateComp, options)
	return rv
}


// Returns the difference between two supplied dates as date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/components(_:from:to:options:)-84y5w
func (c_ Calendar) ComponentsFromDateToDateOptions(unitFlags NSCalendarUnit, startingDate IDate, resultDate IDate, opts NSCalendarOptions) DateComponents {
	rv := objc.Send[DateComponents](c_.ID, objc.Sel("components:fromDate:toDate:options:"), unitFlags, startingDate, resultDate, opts)
	return rv
}


// Returns all the date components of a date, as if in a given time zone (instead of the receiving calendar’s time zone).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/components(in:from:)
func (c_ Calendar) ComponentsInTimeZoneFromDate(timezone ITimeZone, date IDate) DateComponents {
	rv := objc.Send[DateComponents](c_.ID, objc.Sel("componentsInTimeZone:fromDate:"), timezone, date)
	return rv
}


// Returns whether a given date matches all of the given date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(_:matchesComponents:)
func (c_ Calendar) DateMatchesComponents(date IDate, components IDateComponents) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("date:matchesComponents:"), date, components)
	return rv
}


// Returns a date representing the absolute time calculated by adding given components to a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(byAdding:to:options:)
func (c_ Calendar) DateByAddingComponentsToDateOptions(comps IDateComponents, date IDate, opts NSCalendarOptions) Date {
	rv := objc.Send[Date](c_.ID, objc.Sel("dateByAddingComponents:toDate:options:"), comps, date, opts)
	return rv
}


// Returns a date representing the absolute time calculated by adding the value of a given component to a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(byAdding:value:to:options:)
func (c_ Calendar) DateByAddingUnitValueToDateOptions(unit NSCalendarUnit, value int, date IDate, options NSCalendarOptions) Date {
	rv := objc.Send[Date](c_.ID, objc.Sel("dateByAddingUnit:value:toDate:options:"), unit, value, date, options)
	return rv
}


// Creates a new date calculated with the given time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(bySettingHour:minute:second:of:options:)
func (c_ Calendar) DateBySettingHourMinuteSecondOfDateOptions(h int, m int, s int, date IDate, opts NSCalendarOptions) Date {
	rv := objc.Send[Date](c_.ID, objc.Sel("dateBySettingHour:minute:second:ofDate:options:"), h, m, s, date, opts)
	return rv
}


// Returns a new date representing the date calculated by setting a specific component of a given date to a given value, while trying to keep lower components the same.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(bySettingUnit:value:of:options:)
func (c_ Calendar) DateBySettingUnitValueOfDateOptions(unit NSCalendarUnit, v int, date IDate, opts NSCalendarOptions) Date {
	rv := objc.Send[Date](c_.ID, objc.Sel("dateBySettingUnit:value:ofDate:options:"), unit, v, date, opts)
	return rv
}


// Returns a date created with the given components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(era:year:month:day:hour:minute:second:nanosecond:)
func (c_ Calendar) DateWithEraYearMonthDayHourMinuteSecondNanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	rv := objc.Send[Date](c_.ID, objc.Sel("dateWithEra:year:month:day:hour:minute:second:nanosecond:"), eraValue, yearValue, monthValue, dayValue, hourValue, minuteValue, secondValue, nanosecondValue)
	return rv
}


// Returns a new date created with the given components base on a week-of-year value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(era:yearForWeekOfYear:weekOfYear:weekday:hour:minute:second:nanosecond:)
func (c_ Calendar) DateWithEraYearForWeekOfYearWeekOfYearWeekdayHourMinuteSecondNanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	rv := objc.Send[Date](c_.ID, objc.Sel("dateWithEra:yearForWeekOfYear:weekOfYear:weekday:hour:minute:second:nanosecond:"), eraValue, yearValue, weekValue, weekdayValue, hourValue, minuteValue, secondValue, nanosecondValue)
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


// Computes the dates that match (or most closely match) a given set of components, and calls the block once for each of them, until the enumeration is stopped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/enumerateDates(startingAfter:matching:options:using:)
func (c_ Calendar) EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock(start IDate, comps IDateComponents, opts NSCalendarOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("enumerateDatesStartingAfterDate:matchingComponents:options:usingBlock:"), start, comps, opts, block)
}


// Returns by reference the era, year, week of year, and weekday component values for a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/getEra(_:year:month:day:from:)
func (c_ Calendar) GetEraYearMonthDayFromDate(eraValuePointer unsafe.Pointer, yearValuePointer unsafe.Pointer, monthValuePointer unsafe.Pointer, dayValuePointer unsafe.Pointer, date IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getEra:year:month:day:fromDate:"), eraValuePointer, yearValuePointer, monthValuePointer, dayValuePointer, date)
}


// Returns by reference the era, year, week of year, and weekday component values for a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/getEra(_:yearForWeekOfYear:weekOfYear:weekday:from:)
func (c_ Calendar) GetEraYearForWeekOfYearWeekOfYearWeekdayFromDate(eraValuePointer unsafe.Pointer, yearValuePointer unsafe.Pointer, weekValuePointer unsafe.Pointer, weekdayValuePointer unsafe.Pointer, date IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getEra:yearForWeekOfYear:weekOfYear:weekday:fromDate:"), eraValuePointer, yearValuePointer, weekValuePointer, weekdayValuePointer, date)
}


// Returns by reference the hour, minute, second, and nanosecond component values for a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/getHour(_:minute:second:nanosecond:from:)
func (c_ Calendar) GetHourMinuteSecondNanosecondFromDate(hourValuePointer unsafe.Pointer, minuteValuePointer unsafe.Pointer, secondValuePointer unsafe.Pointer, nanosecondValuePointer unsafe.Pointer, date IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getHour:minute:second:nanosecond:fromDate:"), hourValuePointer, minuteValuePointer, secondValuePointer, nanosecondValuePointer, date)
}


// Indicates whether two dates are equal to a given unit of granularity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/isDate(_:equalTo:toUnitGranularity:)
func (c_ Calendar) IsDateEqualToDateToUnitGranularity(date1 IDate, date2 IDate, unit NSCalendarUnit) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDate:equalToDate:toUnitGranularity:"), date1, date2, unit)
	return rv
}


// Indicates whether two dates are in the same day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/isDate(_:inSameDayAs:)
func (c_ Calendar) IsDateInSameDayAsDate(date1 IDate, date2 IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDate:inSameDayAsDate:"), date1, date2)
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


// Indicates whether the given date is in “tomorrow.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/isDateInTomorrow(_:)
func (c_ Calendar) IsDateInTomorrow(date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDateInTomorrow:"), date)
	return rv
}


// Indicates whether a given date falls within a weekend period, as defined by the calendar and the calendar’s locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/isDateInWeekend(_:)
func (c_ Calendar) IsDateInWeekend(date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDateInWeekend:"), date)
	return rv
}


// Indicates whether the given date is in “yesterday.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/isDateInYesterday(_:)
func (c_ Calendar) IsDateInYesterday(date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDateInYesterday:"), date)
	return rv
}


// Returns the maximum range limits of the values that a given unit can take on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/maximumRange(of:)
func (c_ Calendar) MaximumRangeOfUnit(unit NSCalendarUnit) Range {
	rv := objc.Send[Range](c_.ID, objc.Sel("maximumRangeOfUnit:"), unit)
	return rv
}


// Returns the minimum range limits of the values that a given unit can take on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/minimumRange(of:)
func (c_ Calendar) MinimumRangeOfUnit(unit NSCalendarUnit) Range {
	rv := objc.Send[Range](c_.ID, objc.Sel("minimumRangeOfUnit:"), unit)
	return rv
}


// Returns the next date after a given date matching the given components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/nextDate(after:matching:options:)
func (c_ Calendar) NextDateAfterDateMatchingComponentsOptions(date IDate, comps IDateComponents, options NSCalendarOptions) Date {
	rv := objc.Send[Date](c_.ID, objc.Sel("nextDateAfterDate:matchingComponents:options:"), date, comps, options)
	return rv
}


// Returns the next date after a given date matching the given calendar unit value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/nextDate(after:matching:value:options:)
func (c_ Calendar) NextDateAfterDateMatchingUnitValueOptions(date IDate, unit NSCalendarUnit, value int, options NSCalendarOptions) Date {
	rv := objc.Send[Date](c_.ID, objc.Sel("nextDateAfterDate:matchingUnit:value:options:"), date, unit, value, options)
	return rv
}


// Returns the next date after a given date that matches the given hour, minute, and second, component values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/nextDate(after:matchingHour:minute:second:options:)
func (c_ Calendar) NextDateAfterDateMatchingHourMinuteSecondOptions(date IDate, hourValue int, minuteValue int, secondValue int, options NSCalendarOptions) Date {
	rv := objc.Send[Date](c_.ID, objc.Sel("nextDateAfterDate:matchingHour:minute:second:options:"), date, hourValue, minuteValue, secondValue, options)
	return rv
}


// Returns by reference the starting date and time interval range of the next weekend period after a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/nextWeekendStart(_:interval:options:after:)
func (c_ Calendar) NextWeekendStartDateIntervalOptionsAfterDate(datep IDate, tip TimeInterval, options NSCalendarOptions, date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("nextWeekendStartDate:interval:options:afterDate:"), datep, tip, options, date)
	return rv
}


// Returns, for a given absolute time, the ordinal number of a smaller calendar unit (such as a day) within a specified larger calendar unit (such as a week).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/ordinality(of:in:for:)
func (c_ Calendar) OrdinalityOfUnitInUnitForDate(smaller NSCalendarUnit, larger NSCalendarUnit, date IDate) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("ordinalityOfUnit:inUnit:forDate:"), smaller, larger, date)
	return rv
}


// Returns the range of absolute time values that a smaller calendar unit (such as a day) can take on in a larger calendar unit (such as a month) that includes a specified absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/range(of:in:for:)
func (c_ Calendar) RangeOfUnitInUnitForDate(smaller NSCalendarUnit, larger NSCalendarUnit, date IDate) Range {
	rv := objc.Send[Range](c_.ID, objc.Sel("rangeOfUnit:inUnit:forDate:"), smaller, larger, date)
	return rv
}


// Returns by reference the starting time and duration of a given calendar unit that contains a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/range(of:start:interval:for:)
func (c_ Calendar) RangeOfUnitStartDateIntervalForDate(unit NSCalendarUnit, datep IDate, tip TimeInterval, date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("rangeOfUnit:startDate:interval:forDate:"), unit, datep, tip, date)
	return rv
}


// Returns whether a given date falls within a weekend period, and if so, returns by reference the start date and time interval of the weekend range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/range(ofWeekendStart:interval:containing:)
func (c_ Calendar) RangeOfWeekendStartDateIntervalContainingDate(datep IDate, tip TimeInterval, date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("rangeOfWeekendStartDate:interval:containingDate:"), datep, tip, date)
	return rv
}


// Returns the first moment of a given date as a date instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/startOfDay(for:)
func (c_ Calendar) StartOfDayForDate(date IDate) Date {
	rv := objc.Send[Date](c_.ID, objc.Sel("startOfDayForDate:"), date)
	return rv
}


// The symbol used to represent “AM” for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/amSymbol
func (c_ Calendar) AMSymbol() string {
	rv := objc.Send[string](c_.ID, objc.Sel("AMSymbol"))
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


// A list of era symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/eraSymbols
func (c_ Calendar) EraSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("eraSymbols"))
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


// A list of long era symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/longEraSymbols
func (c_ Calendar) LongEraSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("longEraSymbols"))
	return rv
}


// The minimum number of days in the first week of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/minimumDaysInFirstWeek
func (c_ Calendar) MinimumDaysInFirstWeek() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("minimumDaysInFirstWeek"))
	return rv
}


// The minimum number of days in the first week of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/minimumDaysInFirstWeek
func (c_ Calendar) SetMinimumDaysInFirstWeek(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumDaysInFirstWeek:"), value)
}


// A list of month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/monthSymbols
func (c_ Calendar) MonthSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("monthSymbols"))
	return rv
}


// The symbol used to represent “PM” for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/pmSymbol
func (c_ Calendar) PMSymbol() string {
	rv := objc.Send[string](c_.ID, objc.Sel("PMSymbol"))
	return rv
}


// A list of quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/quarterSymbols
func (c_ Calendar) QuarterSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("quarterSymbols"))
	return rv
}


// A list of short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/shortMonthSymbols
func (c_ Calendar) ShortMonthSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("shortMonthSymbols"))
	return rv
}


// A list of short quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/shortQuarterSymbols
func (c_ Calendar) ShortQuarterSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("shortQuarterSymbols"))
	return rv
}


// A list of short standalone month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/shortStandaloneMonthSymbols
func (c_ Calendar) ShortStandaloneMonthSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("shortStandaloneMonthSymbols"))
	return rv
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


// A list of shorter-named weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/shortWeekdaySymbols
func (c_ Calendar) ShortWeekdaySymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("shortWeekdaySymbols"))
	return rv
}


// A list of standalone month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/standaloneMonthSymbols
func (c_ Calendar) StandaloneMonthSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("standaloneMonthSymbols"))
	return rv
}


// A list of standalone quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/standaloneQuarterSymbols
func (c_ Calendar) StandaloneQuarterSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("standaloneQuarterSymbols"))
	return rv
}


// A list of standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/standaloneWeekdaySymbols
func (c_ Calendar) StandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("standaloneWeekdaySymbols"))
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


// A list of very short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/veryShortMonthSymbols
func (c_ Calendar) VeryShortMonthSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("veryShortMonthSymbols"))
	return rv
}


// A list of very short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/veryShortStandaloneMonthSymbols
func (c_ Calendar) VeryShortStandaloneMonthSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("veryShortStandaloneMonthSymbols"))
	return rv
}


// A list of very short standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/veryShortStandaloneWeekdaySymbols
func (c_ Calendar) VeryShortStandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("veryShortStandaloneWeekdaySymbols"))
	return rv
}


// A list of very-shortly-named weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/veryShortWeekdaySymbols
func (c_ Calendar) VeryShortWeekdaySymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("veryShortWeekdaySymbols"))
	return rv
}


// A list of weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/weekdaySymbols
func (c_ Calendar) WeekdaySymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("weekdaySymbols"))
	return rv
}


