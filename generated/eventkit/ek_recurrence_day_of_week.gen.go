// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [EKRecurrenceDayOfWeek] class.
type IEKRecurrenceDayOfWeek interface {
	objectivec.IObject
}

// A class that represents the day of the week.
//
// The class represents a day of the week for use with an object. A day of the week can optionally have a week number, indicating a specific day in the recurrence rule’s frequency. For example, a day of the week with a day value of Tuesday and a week number of 2 would represent the second Tuesday of every month in a monthly recurrence rule, and the second Tuesday of every year in a yearly recurrence rule. A day of the week with a week number of 0 ignores its week number.
//
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

// Alloc allocates a new instance without initialization.
func (ec _EKRecurrenceDayOfWeekClass) Alloc() EKRecurrenceDayOfWeek {
	rv := objc.Send[EKRecurrenceDayOfWeek](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates and returns a day of the week with a given day.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek/init(_:)
func NewEKRecurrenceDayOfWeekWithDay(dayOfTheWeek unsafe.Pointer) EKRecurrenceDayOfWeek {
	rv := objc.Send[EKRecurrenceDayOfWeek](objc.ID(getEKRecurrenceDayOfWeekClass().class), objc.Sel("dayOfWeek:"), dayOfTheWeek)
	return rv
}

// Creates and returns an autoreleased day of the week with a given day and week number.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek/init(_:weekNumber:)
func NewEKRecurrenceDayOfWeekWeekNumber(dayOfTheWeek unsafe.Pointer, weekNumber int) EKRecurrenceDayOfWeek {
	rv := objc.Send[EKRecurrenceDayOfWeek](objc.ID(getEKRecurrenceDayOfWeekClass().class), objc.Sel("dayOfWeek:weekNumber:"), dayOfTheWeek, weekNumber)
	return rv
}

// Initializes and returns a day of the week with a given day and week number.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek/init(dayOfTheWeek:weekNumber:)
func NewEKRecurrenceDayOfWeekWithDayOfTheWeekWeekNumber(dayOfTheWeek unsafe.Pointer, weekNumber int) EKRecurrenceDayOfWeek {
	instance := getEKRecurrenceDayOfWeekClass().Alloc()
	rv := objc.Send[EKRecurrenceDayOfWeek](instance.ID, objc.Sel("initWithDayOfTheWeek:weekNumber:"), dayOfTheWeek, weekNumber)
	rv.Autorelease()
	return rv
}


// Creates and returns a day of the week with a given day.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek/init(_:)
func (ec _EKRecurrenceDayOfWeekClass) DayOfWeek(dayOfTheWeek unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("dayOfWeek:"), dayOfTheWeek)
	return rv
}

// Creates and returns an autoreleased day of the week with a given day and week number.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek/init(_:weekNumber:)
func (ec _EKRecurrenceDayOfWeekClass) DayOfWeekWeekNumber(dayOfTheWeek unsafe.Pointer, weekNumber int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("dayOfWeek:weekNumber:"), dayOfTheWeek, weekNumber)
	return rv
}

// The day of the week.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek/dayOfTheWeek
func (e_ EKRecurrenceDayOfWeek) DayOfTheWeek() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("dayOfTheWeek"))
	return rv
}

// The week number of the day of the week.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKRecurrenceDayOfWeek/weekNumber
func (e_ EKRecurrenceDayOfWeek) WeekNumber() int {
	rv := objc.Send[int](e_.ID, objc.Sel("weekNumber"))
	return rv
}


