// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class EKRecurrenceDayOfWeek */


/* debug [class_header]: Header for EKRecurrenceDayOfWeek */
// The class instance for the [EKRecurrenceDayOfWeek] class.
var (
	EKRecurrenceDayOfWeekClass     _EKRecurrenceDayOfWeekClass
	EKRecurrenceDayOfWeekClassOnce sync.Once
)

func getEKRecurrenceDayOfWeekClass() _EKRecurrenceDayOfWeekClass {
	EKRecurrenceDayOfWeekClassOnce.Do(func() {
		EKRecurrenceDayOfWeekClass = _EKRecurrenceDayOfWeekClass{objc.GetClass("EKRecurrenceDayOfWeek")}
	})
	return EKRecurrenceDayOfWeekClass
}

type _EKRecurrenceDayOfWeekClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EKRecurrenceDayOfWeek */
// An interface definition for the [EKRecurrenceDayOfWeek] class.
type IEKRecurrenceDayOfWeek interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for EKRecurrenceDayOfWeek */
	// properties:
	DayOfTheWeek() EKWeekday
	WeekNumber() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EKRecurrenceDayOfWeek */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EKRecurrenceDayOfWeek */
// Alloc allocates a new instance without initialization.
func (ec _EKRecurrenceDayOfWeekClass) Alloc() EKRecurrenceDayOfWeek {
	rv := objc.Send[EKRecurrenceDayOfWeek](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EKRecurrenceDayOfWeekClass) New() EKRecurrenceDayOfWeek {
	rv := objc.Send[EKRecurrenceDayOfWeek](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKRecurrenceDayOfWeek) Init() EKRecurrenceDayOfWeek {
	rv := objc.Send[EKRecurrenceDayOfWeek](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKRecurrenceDayOfWeek) Autorelease() EKRecurrenceDayOfWeek {
	rv := objc.Send[EKRecurrenceDayOfWeek](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKRecurrenceDayOfWeek creates a new EKRecurrenceDayOfWeek instance.
func NewEKRecurrenceDayOfWeek() EKRecurrenceDayOfWeek {
	return getEKRecurrenceDayOfWeekClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EKRecurrenceDayOfWeek */
// A class that represents the day of the week.
//
// The class represents a day of the week for use with an object. A day of the week can optionally have a week number, indicating a specific day in the recurrence rule’s frequency. For example, a day of the week with a day value of Tuesday and a week number of 2 would represent the second Tuesday of every month in a monthly recurrence rule, and the second Tuesday of every year in a yearly recurrence rule. A day of the week with a week number of 0 ignores its week number.


// A class that represents the day of the week.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek
type EKRecurrenceDayOfWeek struct {
	objectivec.Object
}

// EKRecurrenceDayOfWeekFrom constructs a [EKRecurrenceDayOfWeek] from an unsafe.Pointer.
//
// A class that represents the day of the week.
func EKRecurrenceDayOfWeekFrom(ptr unsafe.Pointer) EKRecurrenceDayOfWeek {
	return EKRecurrenceDayOfWeek{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EKRecurrenceDayOfWeek */

// Creates and returns a day of the week with a given day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek/init(_:)
func NewEKRecurrenceDayOfWeek(dayOfTheWeek EKWeekday) EKRecurrenceDayOfWeek {
	rv := objc.Send[EKRecurrenceDayOfWeek](objc.ID(getEKRecurrenceDayOfWeekClass().class), objc.Sel("dayOfWeek:"), dayOfTheWeek)
	return rv
}/* debug [class_init_methods/constructor]: NewEKRecurrenceDayOfWeek */


// Creates and returns an autoreleased day of the week with a given day and week number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek/init(_:weekNumber:)
func NewEKRecurrenceDayOfWeekWeekNumber(dayOfTheWeek EKWeekday, weekNumber int) EKRecurrenceDayOfWeek {
	rv := objc.Send[EKRecurrenceDayOfWeek](objc.ID(getEKRecurrenceDayOfWeekClass().class), objc.Sel("dayOfWeek:weekNumber:"), dayOfTheWeek, weekNumber)
	return rv
}/* debug [class_init_methods/constructor]: NewEKRecurrenceDayOfWeekWeekNumber */


// Initializes and returns a day of the week with a given day and week number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek/init(dayOfTheWeek:weekNumber:)
func NewEKRecurrenceDayOfWeekWithDayOfTheWeekWeekNumber(dayOfTheWeek EKWeekday, weekNumber int) EKRecurrenceDayOfWeek {
	instance := getEKRecurrenceDayOfWeekClass().Alloc()
	rv := objc.Send[EKRecurrenceDayOfWeek](instance.ID, objc.Sel("initWithDayOfTheWeek:weekNumber:"), dayOfTheWeek, weekNumber)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewEKRecurrenceDayOfWeekWithDayOfTheWeekWeekNumber */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EKRecurrenceDayOfWeek */

// Creates and returns a day of the week with a given day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek/init(_:)
func (ec _EKRecurrenceDayOfWeekClass) DayOfWeek(dayOfTheWeek EKWeekday) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ec.class), objc.Sel("dayOfWeek:"), dayOfTheWeek)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DayOfWeek) */


// Creates and returns an autoreleased day of the week with a given day and week number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek/init(_:weekNumber:)
func (ec _EKRecurrenceDayOfWeekClass) DayOfWeekWeekNumber(dayOfTheWeek EKWeekday, weekNumber int) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ec.class), objc.Sel("dayOfWeek:weekNumber:"), dayOfTheWeek, weekNumber)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DayOfWeekWeekNumber) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EKRecurrenceDayOfWeek */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EKRecurrenceDayOfWeek */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EKRecurrenceDayOfWeek */

// The day of the week.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek/dayOfTheWeek
func (e_ EKRecurrenceDayOfWeek) DayOfTheWeek() EKWeekday {
	rv := objc.Send[EKWeekday](e_.ID, objc.Sel("dayOfTheWeek"))
	return rv
}/* debug [instance_properties/getter]: dayOfTheWeek */


// The week number of the day of the week.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek/weekNumber
func (e_ EKRecurrenceDayOfWeek) WeekNumber() int {
	rv := objc.Send[int](e_.ID, objc.Sel("weekNumber"))
	return rv
}/* debug [instance_properties/getter]: weekNumber */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EKRecurrenceDayOfWeek */


