// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [DateComponents] class.
var (
	DateComponentsClass     _DateComponentsClass
	DateComponentsClassOnce sync.Once
)

func getDateComponentsClass() _DateComponentsClass {
	DateComponentsClassOnce.Do(func() {
		DateComponentsClass = _DateComponentsClass{objc.GetClass("NSDateComponents")}
	})
	return DateComponentsClass
}

type _DateComponentsClass struct {
	class objc.Class
}





// An interface definition for the [DateComponents] class.
type IDateComponents interface {
	objectivec.IObject
	

	// properties:
	Calendar() ICalendar
	SetCalendar(value ICalendar)
	Date() IDate
	Day() int
	SetDay(value int)
	DayOfYear() int
	SetDayOfYear(value int)
	Era() int
	SetEra(value int)
	Hour() int
	SetHour(value int)
	LeapMonth() bool
	SetLeapMonth(value bool)
	RepeatedDay() bool
	SetRepeatedDay(value bool)
	ValidDate() bool
	Minute() int
	SetMinute(value int)
	Month() int
	SetMonth(value int)
	Nanosecond() int
	SetNanosecond(value int)
	Quarter() int
	SetQuarter(value int)
	Second() int
	SetSecond(value int)
	TimeZone() ITimeZone
	SetTimeZone(value ITimeZone)
	WeekOfMonth() int
	SetWeekOfMonth(value int)
	WeekOfYear() int
	SetWeekOfYear(value int)
	Weekday() int
	SetWeekday(value int)
	WeekdayOrdinal() int
	SetWeekdayOrdinal(value int)
	Year() int
	SetYear(value int)
	YearForWeekOfYear() int
	SetYearForWeekOfYear(value int)
	NSDateComponentUndefined() int
	SetNSDateComponentUndefined(value int)
	IsLeapMonth() bool
	SetIsLeapMonth(value bool)
	IsRepeatedDay() bool
	SetIsRepeatedDay(value bool)
	IsValidDate() bool
	SetIsValidDate(value bool)


	

	// methods:
	IsValidDateInCalendar(calendar ICalendar) bool
	SetValueForComponent(value int, unit CalendarUnit)
	ValueForComponent(unit CalendarUnit) int


}





// Alloc allocates a new instance without initialization.
func (dc _DateComponentsClass) Alloc() DateComponents {
	rv := objc.Send[DateComponents](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DateComponentsClass) New() DateComponents {
	rv := objc.Send[DateComponents](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DateComponents) Init() DateComponents {
	rv := objc.Send[DateComponents](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DateComponents) Autorelease() DateComponents {
	rv := objc.Send[DateComponents](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDateComponents creates a new DateComponents instance.
func NewDateComponents() DateComponents {
	return getDateComponentsClass().New()
}





// An object that specifies a date or time in terms of units (such as year, month, day, hour, and minute) to be evaluated in a calendar system and time zone.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. encapsulates the components of a date in an extendable, object-oriented manner. It’s used to specify a date by providing the temporal components that make up a date and time: hour, minutes, seconds, day, month, year, and so on. You can also use it to specify a duration of time, for example, 5 hours and 16 minutes. An object is not required to define all the component fields. When a new instance of is created, the date components are set to . An instance of is not responsible for answering questions about a date beyond the information with which it was initialized. For example, if you initialize one with May 4, 2017, its weekday is , not Thursday. To get the correct day of the week, you must create a suitable instance of , create an object using and then use to retrieve the weekday—as illustrated in the following example. For more details, see in .


// An object that specifies a date or time in terms of units (such as year, month, day, hour, and minute) to be evaluated in a calendar system and time zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents
type DateComponents struct {
	objectivec.Object
}

// DateComponentsFrom constructs a [DateComponents] from an unsafe.Pointer.
//
// An object that specifies a date or time in terms of units (such as year, month, day, hour, and minute) to be evaluated in a calendar system and time zone.
func DateComponentsFrom(ptr unsafe.Pointer) DateComponents {
	return DateComponents{objectivec.Object{objc.ID(ptr)}}
}




















// Returns a Boolean value that indicates whether the current combination of properties represents a date which exists in the specified calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/isValidDate(in:)
func (d_ DateComponents) IsValidDateInCalendar(calendar ICalendar) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isValidDateInCalendar:"), calendar)
	return rv
}


// Sets a value for a given calendar unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/setValue(_:forComponent:)
func (d_ DateComponents) SetValueForComponent(value int, unit CalendarUnit) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setValue:forComponent:"), value, unit)
}


// Returns the value for a given calendar unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/value(forComponent:)
func (d_ DateComponents) ValueForComponent(unit CalendarUnit) int {
	rv := objc.Send[int](d_.ID, objc.Sel("valueForComponent:"), unit)
	return rv
}







// The calendar used to interpret the date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/calendar
func (d_ DateComponents) Calendar() ICalendar {
	rv := objc.Send[Calendar](d_.ID, objc.Sel("calendar"))
	return rv
}


// The calendar used to interpret the date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/calendar
func (d_ DateComponents) SetCalendar(value ICalendar) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCalendar:"), value)
}


// The date calculated from the current components using the stored calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/date
func (d_ DateComponents) Date() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("date"))
	return rv
}


// The number of days.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/day
func (d_ DateComponents) Day() int {
	rv := objc.Send[int](d_.ID, objc.Sel("day"))
	return rv
}


// The number of days.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/day
func (d_ DateComponents) SetDay(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDay:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/dayOfYear
func (d_ DateComponents) DayOfYear() int {
	rv := objc.Send[int](d_.ID, objc.Sel("dayOfYear"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/dayOfYear
func (d_ DateComponents) SetDayOfYear(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDayOfYear:"), value)
}


// The number of eras.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/era
func (d_ DateComponents) Era() int {
	rv := objc.Send[int](d_.ID, objc.Sel("era"))
	return rv
}


// The number of eras.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/era
func (d_ DateComponents) SetEra(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setEra:"), value)
}


// The number of hour units for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/hour
func (d_ DateComponents) Hour() int {
	rv := objc.Send[int](d_.ID, objc.Sel("hour"))
	return rv
}


// The number of hour units for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/hour
func (d_ DateComponents) SetHour(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHour:"), value)
}


// A Boolean value that indicates whether the month is a leap month.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/isLeapMonth
func (d_ DateComponents) LeapMonth() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("leapMonth"))
	return rv
}


// A Boolean value that indicates whether the month is a leap month.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/isLeapMonth
func (d_ DateComponents) SetLeapMonth(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLeapMonth:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/isRepeatedDay
func (d_ DateComponents) RepeatedDay() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("repeatedDay"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/isRepeatedDay
func (d_ DateComponents) SetRepeatedDay(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRepeatedDay:"), value)
}


// A Boolean value that indicates whether the current combination of properties represents a date which exists in the current calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/isValidDate
func (d_ DateComponents) ValidDate() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("validDate"))
	return rv
}


// The number of minute units for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/minute
func (d_ DateComponents) Minute() int {
	rv := objc.Send[int](d_.ID, objc.Sel("minute"))
	return rv
}


// The number of minute units for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/minute
func (d_ DateComponents) SetMinute(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinute:"), value)
}


// The number of months.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/month
func (d_ DateComponents) Month() int {
	rv := objc.Send[int](d_.ID, objc.Sel("month"))
	return rv
}


// The number of months.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/month
func (d_ DateComponents) SetMonth(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMonth:"), value)
}


// The number of nanosecond units for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/nanosecond
func (d_ DateComponents) Nanosecond() int {
	rv := objc.Send[int](d_.ID, objc.Sel("nanosecond"))
	return rv
}


// The number of nanosecond units for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/nanosecond
func (d_ DateComponents) SetNanosecond(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNanosecond:"), value)
}


// The number of quarters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/quarter
func (d_ DateComponents) Quarter() int {
	rv := objc.Send[int](d_.ID, objc.Sel("quarter"))
	return rv
}


// The number of quarters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/quarter
func (d_ DateComponents) SetQuarter(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setQuarter:"), value)
}


// The number of second units for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/second
func (d_ DateComponents) Second() int {
	rv := objc.Send[int](d_.ID, objc.Sel("second"))
	return rv
}


// The number of second units for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/second
func (d_ DateComponents) SetSecond(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSecond:"), value)
}


// The time zone used to interpret the date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/timeZone
func (d_ DateComponents) TimeZone() ITimeZone {
	rv := objc.Send[TimeZone](d_.ID, objc.Sel("timeZone"))
	return rv
}


// The time zone used to interpret the date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/timeZone
func (d_ DateComponents) SetTimeZone(value ITimeZone) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
}


// The week number of the months.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/weekOfMonth
func (d_ DateComponents) WeekOfMonth() int {
	rv := objc.Send[int](d_.ID, objc.Sel("weekOfMonth"))
	return rv
}


// The week number of the months.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/weekOfMonth
func (d_ DateComponents) SetWeekOfMonth(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWeekOfMonth:"), value)
}


// The ISO 8601 week date of the year.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/weekOfYear
func (d_ DateComponents) WeekOfYear() int {
	rv := objc.Send[int](d_.ID, objc.Sel("weekOfYear"))
	return rv
}


// The ISO 8601 week date of the year.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/weekOfYear
func (d_ DateComponents) SetWeekOfYear(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWeekOfYear:"), value)
}


// The number of the weekdays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/weekday
func (d_ DateComponents) Weekday() int {
	rv := objc.Send[int](d_.ID, objc.Sel("weekday"))
	return rv
}


// The number of the weekdays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/weekday
func (d_ DateComponents) SetWeekday(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWeekday:"), value)
}


// The ordinal number of weekdays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/weekdayOrdinal
func (d_ DateComponents) WeekdayOrdinal() int {
	rv := objc.Send[int](d_.ID, objc.Sel("weekdayOrdinal"))
	return rv
}


// The ordinal number of weekdays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/weekdayOrdinal
func (d_ DateComponents) SetWeekdayOrdinal(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWeekdayOrdinal:"), value)
}


// The number of years.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/year
func (d_ DateComponents) Year() int {
	rv := objc.Send[int](d_.ID, objc.Sel("year"))
	return rv
}


// The number of years.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/year
func (d_ DateComponents) SetYear(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setYear:"), value)
}


// The ISO 8601 week-numbering year.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/yearForWeekOfYear
func (d_ DateComponents) YearForWeekOfYear() int {
	rv := objc.Send[int](d_.ID, objc.Sel("yearForWeekOfYear"))
	return rv
}


// The ISO 8601 week-numbering year.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents/yearForWeekOfYear
func (d_ DateComponents) SetYearForWeekOfYear(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setYearForWeekOfYear:"), value)
}


// Specifies a date component without a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentundefined
func (d_ DateComponents) NSDateComponentUndefined() int {
	rv := objc.Send[int](d_.ID, objc.Sel("NSDateComponentUndefined"))
	return rv
}


// Specifies a date component without a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponentundefined
func (d_ DateComponents) SetNSDateComponentUndefined(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNSDateComponentUndefined:"), value)
}


// A Boolean value that indicates whether the month is a leap month.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/isleapmonth
func (d_ DateComponents) IsLeapMonth() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isLeapMonth"))
	return rv
}


// A Boolean value that indicates whether the month is a leap month.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/isleapmonth
func (d_ DateComponents) SetIsLeapMonth(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsLeapMonth:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/isrepeatedday
func (d_ DateComponents) IsRepeatedDay() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isRepeatedDay"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/isrepeatedday
func (d_ DateComponents) SetIsRepeatedDay(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsRepeatedDay:"), value)
}


// A Boolean value that indicates whether the current combination of properties represents a date which exists in the current calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/isvaliddate
func (d_ DateComponents) IsValidDate() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isValidDate"))
	return rv
}


// A Boolean value that indicates whether the current combination of properties represents a date which exists in the current calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatecomponents/isvaliddate
func (d_ DateComponents) SetIsValidDate(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsValidDate:"), value)
}







