// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class EKRecurrenceRule */


/* debug [class_header]: Header for EKRecurrenceRule */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EKRecurrenceRule */
// An interface definition for the [EKRecurrenceRule] class.
type IEKRecurrenceRule interface {
	IEKObject
	
/* debug [class_interface_properties]: Properties for EKRecurrenceRule */
	// properties:
	CalendarIdentifier() objc.IObject /* cross-framework: NSString */
	DaysOfTheMonth() []foundation.Number
	DaysOfTheWeek() []EKRecurrenceDayOfWeek
	DaysOfTheYear() []foundation.Number
	FirstDayOfTheWeek() int
	Frequency() EKRecurrenceFrequency
	Interval() int
	MonthsOfTheYear() []foundation.Number
	RecurrenceEnd() IEKRecurrenceEnd
	SetRecurrenceEnd(value IEKRecurrenceEnd)
	SetPositions() []foundation.Number
	WeeksOfTheYear() []foundation.Number
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EKRecurrenceRule */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EKRecurrenceRule */
// Alloc allocates a new instance without initialization.
func (ec _EKRecurrenceRuleClass) Alloc() EKRecurrenceRule {
	rv := objc.Send[EKRecurrenceRule](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EKRecurrenceRule */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EKRecurrenceRule */

// Initializes and returns a recurrence rule with a given frequency and additional scheduling information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/init(recurrenceWith:interval:daysOfTheWeek:daysOfTheMonth:monthsOfTheYear:weeksOfTheYear:daysOfTheYear:setPositions:end:)
func NewEKRecurrenceRuleRecurrenceWithFrequencyIntervalDaysOfTheWeekDaysOfTheMonthMonthsOfTheYearWeeksOfTheYearDaysOfTheYearSetPositionsEnd(type_ EKRecurrenceFrequency, interval int, days []EKRecurrenceDayOfWeek, monthDays []foundation.Number, months []foundation.Number, weeksOfTheYear []foundation.Number, daysOfTheYear []foundation.Number, setPositions []foundation.Number, end IEKRecurrenceEnd) EKRecurrenceRule {
	instance := getEKRecurrenceRuleClass().Alloc()
	rv := objc.Send[EKRecurrenceRule](instance.ID, objc.Sel("initRecurrenceWithFrequency:interval:daysOfTheWeek:daysOfTheMonth:monthsOfTheYear:weeksOfTheYear:daysOfTheYear:setPositions:end:"), type_, interval, days, monthDays, months, weeksOfTheYear, daysOfTheYear, setPositions, end)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewEKRecurrenceRuleRecurrenceWithFrequencyIntervalDaysOfTheWeekDaysOfTheMonthMonthsOfTheYearWeeksOfTheYearDaysOfTheYearSetPositionsEnd */


// Initializes and returns a simple recurrence rule with a given frequency, interval, and end.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/init(recurrenceWith:interval:end:)
func NewEKRecurrenceRuleRecurrenceWithFrequencyIntervalEnd(type_ EKRecurrenceFrequency, interval int, end IEKRecurrenceEnd) EKRecurrenceRule {
	instance := getEKRecurrenceRuleClass().Alloc()
	rv := objc.Send[EKRecurrenceRule](instance.ID, objc.Sel("initRecurrenceWithFrequency:interval:end:"), type_, interval, end)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewEKRecurrenceRuleRecurrenceWithFrequencyIntervalEnd */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EKRecurrenceRule */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EKRecurrenceRule */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EKRecurrenceRule */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EKRecurrenceRule */

// The identifier for the recurrence rule’s calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/calendarIdentifier
func (e_ EKRecurrenceRule) CalendarIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("calendarIdentifier"))
	return rv
}/* debug [instance_properties/getter]: calendarIdentifier */


// The days of the month associated with the recurrence rule, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/daysOfTheMonth
func (e_ EKRecurrenceRule) DaysOfTheMonth() []foundation.Number {
	rv := objc.Send[[]foundation.Number](e_.ID, objc.Sel("daysOfTheMonth"))
	return rv
}/* debug [instance_properties/getter]: daysOfTheMonth */


// The days of the week associated with the recurrence rule, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/daysOfTheWeek
func (e_ EKRecurrenceRule) DaysOfTheWeek() []EKRecurrenceDayOfWeek {
	rv := objc.Send[[]EKRecurrenceDayOfWeek](e_.ID, objc.Sel("daysOfTheWeek"))
	return rv
}/* debug [instance_properties/getter]: daysOfTheWeek */


// The days of the year associated with the recurrence rule, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/daysOfTheYear
func (e_ EKRecurrenceRule) DaysOfTheYear() []foundation.Number {
	rv := objc.Send[[]foundation.Number](e_.ID, objc.Sel("daysOfTheYear"))
	return rv
}/* debug [instance_properties/getter]: daysOfTheYear */


// Indicates which day of the week the recurrence rule treats as the first day of the week.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/firstDayOfTheWeek
func (e_ EKRecurrenceRule) FirstDayOfTheWeek() int {
	rv := objc.Send[int](e_.ID, objc.Sel("firstDayOfTheWeek"))
	return rv
}/* debug [instance_properties/getter]: firstDayOfTheWeek */


// The frequency of the recurrence rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/frequency
func (e_ EKRecurrenceRule) Frequency() EKRecurrenceFrequency {
	rv := objc.Send[EKRecurrenceFrequency](e_.ID, objc.Sel("frequency"))
	return rv
}/* debug [instance_properties/getter]: frequency */


// Specifies how often the recurrence rule repeats over the unit of time indicated by its frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/interval
func (e_ EKRecurrenceRule) Interval() int {
	rv := objc.Send[int](e_.ID, objc.Sel("interval"))
	return rv
}/* debug [instance_properties/getter]: interval */


// The months of the year associated with the recurrence rule, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/monthsOfTheYear
func (e_ EKRecurrenceRule) MonthsOfTheYear() []foundation.Number {
	rv := objc.Send[[]foundation.Number](e_.ID, objc.Sel("monthsOfTheYear"))
	return rv
}/* debug [instance_properties/getter]: monthsOfTheYear */


// Indicates when the recurrence rule ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/recurrenceEnd
func (e_ EKRecurrenceRule) RecurrenceEnd() IEKRecurrenceEnd {
	rv := objc.Send[EKRecurrenceEnd](e_.ID, objc.Sel("recurrenceEnd"))
	return rv
}/* debug [instance_properties/getter]: recurrenceEnd */


// Indicates when the recurrence rule ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/recurrenceEnd
func (e_ EKRecurrenceRule) SetRecurrenceEnd(value IEKRecurrenceEnd) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRecurrenceEnd:"), value)
}/* debug [instance_properties/setter]: recurrenceEnd */


// An array of ordinal numbers that filters which recurrences to include in the recurrence rule’s frequency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/setPositions
func (e_ EKRecurrenceRule) SetPositions() []foundation.Number {
	rv := objc.Send[[]foundation.Number](e_.ID, objc.Sel("setPositions"))
	return rv
}/* debug [instance_properties/getter]: setPositions */


// The weeks of the year associated with the recurrence rule, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceRule/weeksOfTheYear
func (e_ EKRecurrenceRule) WeeksOfTheYear() []foundation.Number {
	rv := objc.Send[[]foundation.Number](e_.ID, objc.Sel("weeksOfTheYear"))
	return rv
}/* debug [instance_properties/getter]: weeksOfTheYear */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EKRecurrenceRule */


