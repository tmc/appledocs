// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCalendar */


/* debug [class_header]: Header for NSCalendar */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Calendar */
// An interface definition for the [Calendar] class.
type ICalendar interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Calendar */
	// properties:
	AMSymbol() IString
	CalendarIdentifier() CalendarIdentifier
	EraSymbols() []string
	FirstWeekday() uint
	SetFirstWeekday(value uint)
	Locale() ILocale
	SetLocale(value ILocale)
	LongEraSymbols() []string
	MinimumDaysInFirstWeek() uint
	SetMinimumDaysInFirstWeek(value uint)
	MonthSymbols() []string
	PMSymbol() IString
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
	TimeZone() ITimeZone
	SetTimeZone(value ITimeZone)
	VeryShortMonthSymbols() []string
	VeryShortStandaloneMonthSymbols() []string
	VeryShortStandaloneWeekdaySymbols() []string
	VeryShortWeekdaySymbols() []string
	WeekdaySymbols() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Calendar */
	// methods:
	CompareDateToDateToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) ComparisonResult
	ComponentFromDate(unit CalendarUnit, date IDate) int
	ComponentsFromDate(unitFlags CalendarUnit, date IDate) IDateComponents
	ComponentsFromDateComponentsToDateComponentsOptions(unitFlags CalendarUnit, startingDateComp IDateComponents, resultDateComp IDateComponents, options CalendarOptions) IDateComponents
	ComponentsFromDateToDateOptions(unitFlags CalendarUnit, startingDate IDate, resultDate IDate, opts CalendarOptions) IDateComponents
	ComponentsInTimeZoneFromDate(timezone ITimeZone, date IDate) IDateComponents
	DateMatchesComponents(date IDate, components IDateComponents) bool
	DateByAddingComponentsToDateOptions(comps IDateComponents, date IDate, opts CalendarOptions) IDate
	DateByAddingUnitValueToDateOptions(unit CalendarUnit, value int, date IDate, options CalendarOptions) IDate
	DateBySettingHourMinuteSecondOfDateOptions(h int, m int, s int, date IDate, opts CalendarOptions) IDate
	DateBySettingUnitValueOfDateOptions(unit CalendarUnit, v int, date IDate, opts CalendarOptions) IDate
	DateWithEraYearMonthDayHourMinuteSecondNanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) IDate
	DateWithEraYearForWeekOfYearWeekOfYearWeekdayHourMinuteSecondNanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) IDate
	DateFromComponents(comps IDateComponents) IDate
	EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock(start IDate, comps IDateComponents, opts CalendarOptions, block unsafe.Pointer)
	GetEraYearMonthDayFromDate(eraValuePointer int, yearValuePointer int, monthValuePointer int, dayValuePointer int, date IDate)
	GetEraYearForWeekOfYearWeekOfYearWeekdayFromDate(eraValuePointer int, yearValuePointer int, weekValuePointer int, weekdayValuePointer int, date IDate)
	GetHourMinuteSecondNanosecondFromDate(hourValuePointer int, minuteValuePointer int, secondValuePointer int, nanosecondValuePointer int, date IDate)
	IsDateEqualToDateToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) bool
	IsDateInSameDayAsDate(date1 IDate, date2 IDate) bool
	IsDateInToday(date IDate) bool
	IsDateInTomorrow(date IDate) bool
	IsDateInWeekend(date IDate) bool
	IsDateInYesterday(date IDate) bool
	MaximumRangeOfUnit(unit CalendarUnit) objc.IObject /* cross-framework: Range */
	MinimumRangeOfUnit(unit CalendarUnit) objc.IObject /* cross-framework: Range */
	NextDateAfterDateMatchingComponentsOptions(date IDate, comps IDateComponents, options CalendarOptions) IDate
	NextDateAfterDateMatchingUnitValueOptions(date IDate, unit CalendarUnit, value int, options CalendarOptions) IDate
	NextDateAfterDateMatchingHourMinuteSecondOptions(date IDate, hourValue int, minuteValue int, secondValue int, options CalendarOptions) IDate
	NextWeekendStartDateIntervalOptionsAfterDate(datep IDate, tip float64, options CalendarOptions, date IDate) bool
	OrdinalityOfUnitInUnitForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) uint
	RangeOfUnitInUnitForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) objc.IObject /* cross-framework: Range */
	RangeOfUnitStartDateIntervalForDate(unit CalendarUnit, datep IDate, tip float64, date IDate) bool
	RangeOfWeekendStartDateIntervalContainingDate(datep IDate, tip float64, date IDate) bool
	StartOfDayForDate(date IDate) IDate
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Calendar */
// Alloc allocates a new instance without initialization.
func (cc _CalendarClass) Alloc() Calendar {
	rv := objc.Send[Calendar](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Calendar */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Calendar */

// Initializes a calendar according to a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/init(calendarIdentifier:)
func NewCalendarWithCalendarIdentifier(ident CalendarIdentifier) Calendar {
	instance := getCalendarClass().Alloc()
	rv := objc.Send[Calendar](instance.ID, objc.Sel("initWithCalendarIdentifier:"), ident)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCalendarWithCalendarIdentifier */


// Creates a new calendar specified by a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/init(identifier:)
func NewCalendarWithIdentifier(calendarIdentifierConstant CalendarIdentifier) Calendar {
	rv := objc.Send[Calendar](objc.ID(getCalendarClass().class), objc.Sel("calendarWithIdentifier:"), calendarIdentifierConstant)
	return rv
}/* debug [class_init_methods/constructor]: NewCalendarWithIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Calendar */

// Creates a new calendar specified by a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/init(identifier:)
func (cc _CalendarClass) CalendarWithIdentifier(calendarIdentifierConstant CalendarIdentifier) ICalendar {
	rv := objc.Send[Calendar](objc.ID(cc.class), objc.Sel("calendarWithIdentifier:"), calendarIdentifierConstant)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CalendarWithIdentifier) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Calendar */

// A calendar that tracks changes to user’s preferred calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/autoupdatingCurrent
func (cc _CalendarClass) AutoupdatingCurrentCalendar() Calendar {
	rv := objc.Send[Calendar](objc.ID(cc.class), objc.Sel("autoupdatingCurrentCalendar"))
	return rv
}/* debug [class_properties_class/property]: autoupdatingCurrentCalendar */

// The user’s current calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/current
func (cc _CalendarClass) CurrentCalendar() Calendar {
	rv := objc.Send[Calendar](objc.ID(cc.class), objc.Sel("currentCalendar"))
	return rv
}/* debug [class_properties_class/property]: currentCalendar */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Calendar */

// Indicates the ordering of two given dates based on their components down to a given unit granularity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/compare(_:to:toUnitGranularity:)
func (c_ Calendar) CompareDateToDateToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) ComparisonResult {
	rv := objc.Send[ComparisonResult](c_.ID, objc.Sel("compareDate:toDate:toUnitGranularity:"), date1, date2, unit)
	return rv
}/* debug [instance_methods/method]: CompareDateToDateToUnitGranularity */


// Returns the specified date component from a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/component(_:from:)
func (c_ Calendar) ComponentFromDate(unit CalendarUnit, date IDate) int {
	rv := objc.Send[int](c_.ID, objc.Sel("component:fromDate:"), unit, date)
	return rv
}/* debug [instance_methods/method]: ComponentFromDate */


// Returns the date components representing a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/components(_:from:)
func (c_ Calendar) ComponentsFromDate(unitFlags CalendarUnit, date IDate) IDateComponents {
	rv := objc.Send[DateComponents](c_.ID, objc.Sel("components:fromDate:"), unitFlags, date)
	return rv
}/* debug [instance_methods/method]: ComponentsFromDate */


// Returns the difference between start and end dates given as date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/components(_:from:to:options:)-49lo8
func (c_ Calendar) ComponentsFromDateComponentsToDateComponentsOptions(unitFlags CalendarUnit, startingDateComp IDateComponents, resultDateComp IDateComponents, options CalendarOptions) IDateComponents {
	rv := objc.Send[DateComponents](c_.ID, objc.Sel("components:fromDateComponents:toDateComponents:options:"), unitFlags, startingDateComp, resultDateComp, options)
	return rv
}/* debug [instance_methods/method]: ComponentsFromDateComponentsToDateComponentsOptions */


// Returns the difference between two supplied dates as date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/components(_:from:to:options:)-84y5w
func (c_ Calendar) ComponentsFromDateToDateOptions(unitFlags CalendarUnit, startingDate IDate, resultDate IDate, opts CalendarOptions) IDateComponents {
	rv := objc.Send[DateComponents](c_.ID, objc.Sel("components:fromDate:toDate:options:"), unitFlags, startingDate, resultDate, opts)
	return rv
}/* debug [instance_methods/method]: ComponentsFromDateToDateOptions */


// Returns all the date components of a date, as if in a given time zone (instead of the receiving calendar’s time zone).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/components(in:from:)
func (c_ Calendar) ComponentsInTimeZoneFromDate(timezone ITimeZone, date IDate) IDateComponents {
	rv := objc.Send[DateComponents](c_.ID, objc.Sel("componentsInTimeZone:fromDate:"), timezone, date)
	return rv
}/* debug [instance_methods/method]: ComponentsInTimeZoneFromDate */


// Returns whether a given date matches all of the given date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(_:matchesComponents:)
func (c_ Calendar) DateMatchesComponents(date IDate, components IDateComponents) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("date:matchesComponents:"), date, components)
	return rv
}/* debug [instance_methods/method]: DateMatchesComponents */


// Returns a date representing the absolute time calculated by adding given components to a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(byAdding:to:options:)
func (c_ Calendar) DateByAddingComponentsToDateOptions(comps IDateComponents, date IDate, opts CalendarOptions) IDate {
	rv := objc.Send[Date](c_.ID, objc.Sel("dateByAddingComponents:toDate:options:"), comps, date, opts)
	return rv
}/* debug [instance_methods/method]: DateByAddingComponentsToDateOptions */


// Returns a date representing the absolute time calculated by adding the value of a given component to a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(byAdding:value:to:options:)
func (c_ Calendar) DateByAddingUnitValueToDateOptions(unit CalendarUnit, value int, date IDate, options CalendarOptions) IDate {
	rv := objc.Send[Date](c_.ID, objc.Sel("dateByAddingUnit:value:toDate:options:"), unit, value, date, options)
	return rv
}/* debug [instance_methods/method]: DateByAddingUnitValueToDateOptions */


// Creates a new date calculated with the given time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(bySettingHour:minute:second:of:options:)
func (c_ Calendar) DateBySettingHourMinuteSecondOfDateOptions(h int, m int, s int, date IDate, opts CalendarOptions) IDate {
	rv := objc.Send[Date](c_.ID, objc.Sel("dateBySettingHour:minute:second:ofDate:options:"), h, m, s, date, opts)
	return rv
}/* debug [instance_methods/method]: DateBySettingHourMinuteSecondOfDateOptions */


// Returns a new date representing the date calculated by setting a specific component of a given date to a given value, while trying to keep lower components the same.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(bySettingUnit:value:of:options:)
func (c_ Calendar) DateBySettingUnitValueOfDateOptions(unit CalendarUnit, v int, date IDate, opts CalendarOptions) IDate {
	rv := objc.Send[Date](c_.ID, objc.Sel("dateBySettingUnit:value:ofDate:options:"), unit, v, date, opts)
	return rv
}/* debug [instance_methods/method]: DateBySettingUnitValueOfDateOptions */


// Returns a date created with the given components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(era:year:month:day:hour:minute:second:nanosecond:)
func (c_ Calendar) DateWithEraYearMonthDayHourMinuteSecondNanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) IDate {
	rv := objc.Send[Date](c_.ID, objc.Sel("dateWithEra:year:month:day:hour:minute:second:nanosecond:"), eraValue, yearValue, monthValue, dayValue, hourValue, minuteValue, secondValue, nanosecondValue)
	return rv
}/* debug [instance_methods/method]: DateWithEraYearMonthDayHourMinuteSecondNanosecond */


// Returns a new date created with the given components base on a week-of-year value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(era:yearForWeekOfYear:weekOfYear:weekday:hour:minute:second:nanosecond:)
func (c_ Calendar) DateWithEraYearForWeekOfYearWeekOfYearWeekdayHourMinuteSecondNanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) IDate {
	rv := objc.Send[Date](c_.ID, objc.Sel("dateWithEra:yearForWeekOfYear:weekOfYear:weekday:hour:minute:second:nanosecond:"), eraValue, yearValue, weekValue, weekdayValue, hourValue, minuteValue, secondValue, nanosecondValue)
	return rv
}/* debug [instance_methods/method]: DateWithEraYearForWeekOfYearWeekOfYearWeekdayHourMinuteSecondNanosecond */


// Returns a date representing the absolute time calculated from given components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/date(from:)
func (c_ Calendar) DateFromComponents(comps IDateComponents) IDate {
	rv := objc.Send[Date](c_.ID, objc.Sel("dateFromComponents:"), comps)
	return rv
}/* debug [instance_methods/method]: DateFromComponents */


// Computes the dates that match (or most closely match) a given set of components, and calls the block once for each of them, until the enumeration is stopped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/enumerateDates(startingAfter:matching:options:using:)
func (c_ Calendar) EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock(start IDate, comps IDateComponents, opts CalendarOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("enumerateDatesStartingAfterDate:matchingComponents:options:usingBlock:"), start, comps, opts, block)
}/* debug [instance_methods/method]: EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock */


// Returns by reference the era, year, week of year, and weekday component values for a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/getEra(_:year:month:day:from:)
func (c_ Calendar) GetEraYearMonthDayFromDate(eraValuePointer int, yearValuePointer int, monthValuePointer int, dayValuePointer int, date IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getEra:year:month:day:fromDate:"), eraValuePointer, yearValuePointer, monthValuePointer, dayValuePointer, date)
}/* debug [instance_methods/method]: GetEraYearMonthDayFromDate */


// Returns by reference the era, year, week of year, and weekday component values for a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/getEra(_:yearForWeekOfYear:weekOfYear:weekday:from:)
func (c_ Calendar) GetEraYearForWeekOfYearWeekOfYearWeekdayFromDate(eraValuePointer int, yearValuePointer int, weekValuePointer int, weekdayValuePointer int, date IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getEra:yearForWeekOfYear:weekOfYear:weekday:fromDate:"), eraValuePointer, yearValuePointer, weekValuePointer, weekdayValuePointer, date)
}/* debug [instance_methods/method]: GetEraYearForWeekOfYearWeekOfYearWeekdayFromDate */


// Returns by reference the hour, minute, second, and nanosecond component values for a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/getHour(_:minute:second:nanosecond:from:)
func (c_ Calendar) GetHourMinuteSecondNanosecondFromDate(hourValuePointer int, minuteValuePointer int, secondValuePointer int, nanosecondValuePointer int, date IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getHour:minute:second:nanosecond:fromDate:"), hourValuePointer, minuteValuePointer, secondValuePointer, nanosecondValuePointer, date)
}/* debug [instance_methods/method]: GetHourMinuteSecondNanosecondFromDate */


// Indicates whether two dates are equal to a given unit of granularity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/isDate(_:equalTo:toUnitGranularity:)
func (c_ Calendar) IsDateEqualToDateToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDate:equalToDate:toUnitGranularity:"), date1, date2, unit)
	return rv
}/* debug [instance_methods/method]: IsDateEqualToDateToUnitGranularity */


// Indicates whether two dates are in the same day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/isDate(_:inSameDayAs:)
func (c_ Calendar) IsDateInSameDayAsDate(date1 IDate, date2 IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDate:inSameDayAsDate:"), date1, date2)
	return rv
}/* debug [instance_methods/method]: IsDateInSameDayAsDate */


// Indicates whether the given date is in “today.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/isDateInToday(_:)
func (c_ Calendar) IsDateInToday(date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDateInToday:"), date)
	return rv
}/* debug [instance_methods/method]: IsDateInToday */


// Indicates whether the given date is in “tomorrow.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/isDateInTomorrow(_:)
func (c_ Calendar) IsDateInTomorrow(date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDateInTomorrow:"), date)
	return rv
}/* debug [instance_methods/method]: IsDateInTomorrow */


// Indicates whether a given date falls within a weekend period, as defined by the calendar and the calendar’s locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/isDateInWeekend(_:)
func (c_ Calendar) IsDateInWeekend(date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDateInWeekend:"), date)
	return rv
}/* debug [instance_methods/method]: IsDateInWeekend */


// Indicates whether the given date is in “yesterday.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/isDateInYesterday(_:)
func (c_ Calendar) IsDateInYesterday(date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDateInYesterday:"), date)
	return rv
}/* debug [instance_methods/method]: IsDateInYesterday */


// Returns the maximum range limits of the values that a given unit can take on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/maximumRange(of:)
func (c_ Calendar) MaximumRangeOfUnit(unit CalendarUnit) objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("maximumRangeOfUnit:"), unit)
	return rv
}/* debug [instance_methods/method]: MaximumRangeOfUnit */


// Returns the minimum range limits of the values that a given unit can take on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/minimumRange(of:)
func (c_ Calendar) MinimumRangeOfUnit(unit CalendarUnit) objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("minimumRangeOfUnit:"), unit)
	return rv
}/* debug [instance_methods/method]: MinimumRangeOfUnit */


// Returns the next date after a given date matching the given components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/nextDate(after:matching:options:)
func (c_ Calendar) NextDateAfterDateMatchingComponentsOptions(date IDate, comps IDateComponents, options CalendarOptions) IDate {
	rv := objc.Send[Date](c_.ID, objc.Sel("nextDateAfterDate:matchingComponents:options:"), date, comps, options)
	return rv
}/* debug [instance_methods/method]: NextDateAfterDateMatchingComponentsOptions */


// Returns the next date after a given date matching the given calendar unit value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/nextDate(after:matching:value:options:)
func (c_ Calendar) NextDateAfterDateMatchingUnitValueOptions(date IDate, unit CalendarUnit, value int, options CalendarOptions) IDate {
	rv := objc.Send[Date](c_.ID, objc.Sel("nextDateAfterDate:matchingUnit:value:options:"), date, unit, value, options)
	return rv
}/* debug [instance_methods/method]: NextDateAfterDateMatchingUnitValueOptions */


// Returns the next date after a given date that matches the given hour, minute, and second, component values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/nextDate(after:matchingHour:minute:second:options:)
func (c_ Calendar) NextDateAfterDateMatchingHourMinuteSecondOptions(date IDate, hourValue int, minuteValue int, secondValue int, options CalendarOptions) IDate {
	rv := objc.Send[Date](c_.ID, objc.Sel("nextDateAfterDate:matchingHour:minute:second:options:"), date, hourValue, minuteValue, secondValue, options)
	return rv
}/* debug [instance_methods/method]: NextDateAfterDateMatchingHourMinuteSecondOptions */


// Returns by reference the starting date and time interval range of the next weekend period after a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/nextWeekendStart(_:interval:options:after:)
func (c_ Calendar) NextWeekendStartDateIntervalOptionsAfterDate(datep IDate, tip float64, options CalendarOptions, date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("nextWeekendStartDate:interval:options:afterDate:"), datep, tip, options, date)
	return rv
}/* debug [instance_methods/method]: NextWeekendStartDateIntervalOptionsAfterDate */


// Returns, for a given absolute time, the ordinal number of a smaller calendar unit (such as a day) within a specified larger calendar unit (such as a week).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/ordinality(of:in:for:)
func (c_ Calendar) OrdinalityOfUnitInUnitForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("ordinalityOfUnit:inUnit:forDate:"), smaller, larger, date)
	return rv
}/* debug [instance_methods/method]: OrdinalityOfUnitInUnitForDate */


// Returns the range of absolute time values that a smaller calendar unit (such as a day) can take on in a larger calendar unit (such as a month) that includes a specified absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/range(of:in:for:)
func (c_ Calendar) RangeOfUnitInUnitForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("rangeOfUnit:inUnit:forDate:"), smaller, larger, date)
	return rv
}/* debug [instance_methods/method]: RangeOfUnitInUnitForDate */


// Returns by reference the starting time and duration of a given calendar unit that contains a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/range(of:start:interval:for:)
func (c_ Calendar) RangeOfUnitStartDateIntervalForDate(unit CalendarUnit, datep IDate, tip float64, date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("rangeOfUnit:startDate:interval:forDate:"), unit, datep, tip, date)
	return rv
}/* debug [instance_methods/method]: RangeOfUnitStartDateIntervalForDate */


// Returns whether a given date falls within a weekend period, and if so, returns by reference the start date and time interval of the weekend range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/range(ofWeekendStart:interval:containing:)
func (c_ Calendar) RangeOfWeekendStartDateIntervalContainingDate(datep IDate, tip float64, date IDate) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("rangeOfWeekendStartDate:interval:containingDate:"), datep, tip, date)
	return rv
}/* debug [instance_methods/method]: RangeOfWeekendStartDateIntervalContainingDate */


// Returns the first moment of a given date as a date instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/startOfDay(for:)
func (c_ Calendar) StartOfDayForDate(date IDate) IDate {
	rv := objc.Send[Date](c_.ID, objc.Sel("startOfDayForDate:"), date)
	return rv
}/* debug [instance_methods/method]: StartOfDayForDate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Calendar */

// The symbol used to represent “AM” for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/amSymbol
func (c_ Calendar) AMSymbol() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("AMSymbol"))
	return rv
}/* debug [instance_properties/getter]: AMSymbol */


// A calendar that tracks changes to user’s preferred calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/autoupdatingCurrent
func (c_ Calendar) AutoupdatingCurrentCalendar() ICalendar {
	rv := objc.Send[Calendar](c_.ID, objc.Sel("autoupdatingCurrentCalendar"))
	return rv
}/* debug [instance_properties/getter]: autoupdatingCurrentCalendar */


// An identifier for the calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/calendarIdentifier
func (c_ Calendar) CalendarIdentifier() CalendarIdentifier {
	rv := objc.Send[CalendarIdentifier](c_.ID, objc.Sel("calendarIdentifier"))
	return rv
}/* debug [instance_properties/getter]: calendarIdentifier */


// The user’s current calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/current
func (c_ Calendar) CurrentCalendar() ICalendar {
	rv := objc.Send[Calendar](c_.ID, objc.Sel("currentCalendar"))
	return rv
}/* debug [instance_properties/getter]: currentCalendar */


// A list of era symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/eraSymbols
func (c_ Calendar) EraSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("eraSymbols"))
	return rv
}/* debug [instance_properties/getter]: eraSymbols */


// The index of the first weekday of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/firstWeekday
func (c_ Calendar) FirstWeekday() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("firstWeekday"))
	return rv
}/* debug [instance_properties/getter]: firstWeekday */


// The index of the first weekday of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/firstWeekday
func (c_ Calendar) SetFirstWeekday(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFirstWeekday:"), value)
}/* debug [instance_properties/setter]: firstWeekday */


// The locale of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/locale
func (c_ Calendar) Locale() ILocale {
	rv := objc.Send[Locale](c_.ID, objc.Sel("locale"))
	return rv
}/* debug [instance_properties/getter]: locale */


// The locale of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/locale
func (c_ Calendar) SetLocale(value ILocale) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocale:"), value)
}/* debug [instance_properties/setter]: locale */


// A list of long era symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/longEraSymbols
func (c_ Calendar) LongEraSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("longEraSymbols"))
	return rv
}/* debug [instance_properties/getter]: longEraSymbols */


// The minimum number of days in the first week of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/minimumDaysInFirstWeek
func (c_ Calendar) MinimumDaysInFirstWeek() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("minimumDaysInFirstWeek"))
	return rv
}/* debug [instance_properties/getter]: minimumDaysInFirstWeek */


// The minimum number of days in the first week of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/minimumDaysInFirstWeek
func (c_ Calendar) SetMinimumDaysInFirstWeek(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumDaysInFirstWeek:"), value)
}/* debug [instance_properties/setter]: minimumDaysInFirstWeek */


// A list of month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/monthSymbols
func (c_ Calendar) MonthSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("monthSymbols"))
	return rv
}/* debug [instance_properties/getter]: monthSymbols */


// The symbol used to represent “PM” for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/pmSymbol
func (c_ Calendar) PMSymbol() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("PMSymbol"))
	return rv
}/* debug [instance_properties/getter]: PMSymbol */


// A list of quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/quarterSymbols
func (c_ Calendar) QuarterSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("quarterSymbols"))
	return rv
}/* debug [instance_properties/getter]: quarterSymbols */


// A list of short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/shortMonthSymbols
func (c_ Calendar) ShortMonthSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("shortMonthSymbols"))
	return rv
}/* debug [instance_properties/getter]: shortMonthSymbols */


// A list of short quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/shortQuarterSymbols
func (c_ Calendar) ShortQuarterSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("shortQuarterSymbols"))
	return rv
}/* debug [instance_properties/getter]: shortQuarterSymbols */


// A list of short standalone month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/shortStandaloneMonthSymbols
func (c_ Calendar) ShortStandaloneMonthSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("shortStandaloneMonthSymbols"))
	return rv
}/* debug [instance_properties/getter]: shortStandaloneMonthSymbols */


// A list of short standalone quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/shortStandaloneQuarterSymbols
func (c_ Calendar) ShortStandaloneQuarterSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("shortStandaloneQuarterSymbols"))
	return rv
}/* debug [instance_properties/getter]: shortStandaloneQuarterSymbols */


// A list of short standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/shortStandaloneWeekdaySymbols
func (c_ Calendar) ShortStandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("shortStandaloneWeekdaySymbols"))
	return rv
}/* debug [instance_properties/getter]: shortStandaloneWeekdaySymbols */


// A list of shorter-named weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/shortWeekdaySymbols
func (c_ Calendar) ShortWeekdaySymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("shortWeekdaySymbols"))
	return rv
}/* debug [instance_properties/getter]: shortWeekdaySymbols */


// A list of standalone month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/standaloneMonthSymbols
func (c_ Calendar) StandaloneMonthSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("standaloneMonthSymbols"))
	return rv
}/* debug [instance_properties/getter]: standaloneMonthSymbols */


// A list of standalone quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/standaloneQuarterSymbols
func (c_ Calendar) StandaloneQuarterSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("standaloneQuarterSymbols"))
	return rv
}/* debug [instance_properties/getter]: standaloneQuarterSymbols */


// A list of standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/standaloneWeekdaySymbols
func (c_ Calendar) StandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("standaloneWeekdaySymbols"))
	return rv
}/* debug [instance_properties/getter]: standaloneWeekdaySymbols */


// The time zone for the calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/timeZone
func (c_ Calendar) TimeZone() ITimeZone {
	rv := objc.Send[TimeZone](c_.ID, objc.Sel("timeZone"))
	return rv
}/* debug [instance_properties/getter]: timeZone */


// The time zone for the calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/timeZone
func (c_ Calendar) SetTimeZone(value ITimeZone) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeZone:"), value)
}/* debug [instance_properties/setter]: timeZone */


// A list of very short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/veryShortMonthSymbols
func (c_ Calendar) VeryShortMonthSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("veryShortMonthSymbols"))
	return rv
}/* debug [instance_properties/getter]: veryShortMonthSymbols */


// A list of very short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/veryShortStandaloneMonthSymbols
func (c_ Calendar) VeryShortStandaloneMonthSymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("veryShortStandaloneMonthSymbols"))
	return rv
}/* debug [instance_properties/getter]: veryShortStandaloneMonthSymbols */


// A list of very short standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/veryShortStandaloneWeekdaySymbols
func (c_ Calendar) VeryShortStandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("veryShortStandaloneWeekdaySymbols"))
	return rv
}/* debug [instance_properties/getter]: veryShortStandaloneWeekdaySymbols */


// A list of very-shortly-named weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/veryShortWeekdaySymbols
func (c_ Calendar) VeryShortWeekdaySymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("veryShortWeekdaySymbols"))
	return rv
}/* debug [instance_properties/getter]: veryShortWeekdaySymbols */


// A list of weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/weekdaySymbols
func (c_ Calendar) WeekdaySymbols() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("weekdaySymbols"))
	return rv
}/* debug [instance_properties/getter]: weekdaySymbols */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCalendar */


