// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [EKRecurrenceRule] class.
var (
	EKRecurrenceRuleClass     _EKRecurrenceRuleClass
	EKRecurrenceRuleClassOnce sync.Once
)

func getEKRecurrenceRuleClass() _EKRecurrenceRuleClass {
	EKRecurrenceRuleClassOnce.Do(func() {
		EKRecurrenceRuleClass = _EKRecurrenceRuleClass{objc.GetClass("EKRecurrenceRule")}
	})
	return EKRecurrenceRuleClass
}

type _EKRecurrenceRuleClass struct {
	class objc.Class
}

// An interface definition for the [EKRecurrenceRule] class.
type IEKRecurrenceRule interface {
	IEKObject
	CalendarIdentifier() string
	DaysOfTheMonth() []foundation.Number
	DaysOfTheWeek() []EKRecurrenceDayOfWeek
	DaysOfTheYear() []foundation.Number
	FirstDayOfTheWeek() int
	Frequency() EKRecurrenceFrequency
	Interval() int
	MonthsOfTheYear() []foundation.Number
	RecurrenceEnd() EKRecurrenceEnd
	SetRecurrenceEnd(value IEKRecurrenceEnd)
	SetPositions() []foundation.Number
	WeeksOfTheYear() []foundation.Number
}

// A class that describes the pattern for a recurring event.
//
// After you create a recurrence rule, assign it to an event with the method of . Recurrence rules can have an end, represented by an object. The end can be based on a specific date or a maximum number of occurrences.


// A class that describes the pattern for a recurring event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule
type EKRecurrenceRule struct {
	EKObject
}

// EKRecurrenceRuleFrom constructs a [EKRecurrenceRule] from an unsafe.Pointer.
//
// A class that describes the pattern for a recurring event.
func EKRecurrenceRuleFrom(ptr unsafe.Pointer) EKRecurrenceRule {
	return EKRecurrenceRule{
		EKObject: EKObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec _EKRecurrenceRuleClass) Alloc() EKRecurrenceRule {
	rv := objc.Send[EKRecurrenceRule](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EKRecurrenceRuleClass) New() EKRecurrenceRule {
	rv := objc.Send[EKRecurrenceRule](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKRecurrenceRule) Init() EKRecurrenceRule {
	rv := objc.Send[EKRecurrenceRule](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKRecurrenceRule) Autorelease() EKRecurrenceRule {
	rv := objc.Send[EKRecurrenceRule](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKRecurrenceRule creates a new EKRecurrenceRule instance.
func NewEKRecurrenceRule() EKRecurrenceRule {
	return getEKRecurrenceRuleClass().New()
}



// Initializes and returns a recurrence rule with a given frequency and additional scheduling information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/init(recurrenceWith:interval:daysOfTheWeek:daysOfTheMonth:monthsOfTheYear:weeksOfTheYear:daysOfTheYear:setPositions:end:)
func NewEKRecurrenceRuleRecurrenceWithFrequencyIntervalDaysOfTheWeekDaysOfTheMonthMonthsOfTheYearWeeksOfTheYearDaysOfTheYearSetPositionsEnd(type_ IEKRecurrenceFrequency, interval int, days []EKRecurrenceDayOfWeek, monthDays []foundation.INumber, months []foundation.INumber, weeksOfTheYear []foundation.INumber, daysOfTheYear []foundation.INumber, setPositions []foundation.INumber, end IEKRecurrenceEnd) EKRecurrenceRule {
	instance := getEKRecurrenceRuleClass().Alloc()
	rv := objc.Send[EKRecurrenceRule](instance.ID, objc.Sel("initRecurrenceWithFrequency:interval:daysOfTheWeek:daysOfTheMonth:monthsOfTheYear:weeksOfTheYear:daysOfTheYear:setPositions:end:"), type_, interval, days, monthDays, months, weeksOfTheYear, daysOfTheYear, setPositions, end)
	rv.Autorelease()
	return rv
}


// Initializes and returns a simple recurrence rule with a given frequency, interval, and end.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/init(recurrenceWith:interval:end:)
func NewEKRecurrenceRuleRecurrenceWithFrequencyIntervalEnd(type_ IEKRecurrenceFrequency, interval int, end IEKRecurrenceEnd) EKRecurrenceRule {
	instance := getEKRecurrenceRuleClass().Alloc()
	rv := objc.Send[EKRecurrenceRule](instance.ID, objc.Sel("initRecurrenceWithFrequency:interval:end:"), type_, interval, end)
	rv.Autorelease()
	return rv
}



// The identifier for the recurrence rule’s calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/calendarIdentifier
func (e_ EKRecurrenceRule) CalendarIdentifier() string {
	rv := objc.Send[string](e_.ID, objc.Sel("calendarIdentifier"))
	return rv
}


// The days of the month associated with the recurrence rule, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/daysOfTheMonth
func (e_ EKRecurrenceRule) DaysOfTheMonth() []foundation.Number {
	rv := objc.Send[[]foundation.Number](e_.ID, objc.Sel("daysOfTheMonth"))
	return rv
}


// The days of the week associated with the recurrence rule, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/daysOfTheWeek
func (e_ EKRecurrenceRule) DaysOfTheWeek() []EKRecurrenceDayOfWeek {
	rv := objc.Send[[]EKRecurrenceDayOfWeek](e_.ID, objc.Sel("daysOfTheWeek"))
	return rv
}


// The days of the year associated with the recurrence rule, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/daysOfTheYear
func (e_ EKRecurrenceRule) DaysOfTheYear() []foundation.Number {
	rv := objc.Send[[]foundation.Number](e_.ID, objc.Sel("daysOfTheYear"))
	return rv
}


// Indicates which day of the week the recurrence rule treats as the first day of the week.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/firstDayOfTheWeek
func (e_ EKRecurrenceRule) FirstDayOfTheWeek() int {
	rv := objc.Send[int](e_.ID, objc.Sel("firstDayOfTheWeek"))
	return rv
}


// The frequency of the recurrence rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/frequency
func (e_ EKRecurrenceRule) Frequency() EKRecurrenceFrequency {
	rv := objc.Send[EKRecurrenceFrequency](e_.ID, objc.Sel("frequency"))
	return rv
}


// Specifies how often the recurrence rule repeats over the unit of time indicated by its frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/interval
func (e_ EKRecurrenceRule) Interval() int {
	rv := objc.Send[int](e_.ID, objc.Sel("interval"))
	return rv
}


// The months of the year associated with the recurrence rule, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/monthsOfTheYear
func (e_ EKRecurrenceRule) MonthsOfTheYear() []foundation.Number {
	rv := objc.Send[[]foundation.Number](e_.ID, objc.Sel("monthsOfTheYear"))
	return rv
}


// Indicates when the recurrence rule ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/recurrenceEnd
func (e_ EKRecurrenceRule) RecurrenceEnd() EKRecurrenceEnd {
	rv := objc.Send[EKRecurrenceEnd](e_.ID, objc.Sel("recurrenceEnd"))
	return rv
}


// Indicates when the recurrence rule ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/recurrenceEnd
func (e_ EKRecurrenceRule) SetRecurrenceEnd(value IEKRecurrenceEnd) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRecurrenceEnd:"), value)
}


// An array of ordinal numbers that filters which recurrences to include in the recurrence rule’s frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/setPositions
func (e_ EKRecurrenceRule) SetPositions() []foundation.Number {
	rv := objc.Send[[]foundation.Number](e_.ID, objc.Sel("setPositions"))
	return rv
}


// The weeks of the year associated with the recurrence rule, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/weeksOfTheYear
func (e_ EKRecurrenceRule) WeeksOfTheYear() []foundation.Number {
	rv := objc.Send[[]foundation.Number](e_.ID, objc.Sel("weeksOfTheYear"))
	return rv
}


