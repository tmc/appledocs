// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [EKReminder] class.
var (
	EKReminderClass     _EKReminderClass
	EKReminderClassOnce sync.Once
)

func getEKReminderClass() _EKReminderClass {
	EKReminderClassOnce.Do(func() {
		EKReminderClass = _EKReminderClass{objc.GetClass("EKReminder")}
	})
	return EKReminderClass
}

type _EKReminderClass struct {
	class objc.Class
}

// An interface definition for the [EKReminder] class.
type IEKReminder interface {
	IEKCalendarItem
	CompletionDate() foundation.NSDate
	SetCompletionDate(value foundation.NSDate)
	DueDateComponents() foundation.DateComponents
	SetDueDateComponents(value foundation.DateComponents)
	Completed() bool
	SetCompleted(value bool)
	Priority() uint
	SetPriority(value uint)
	StartDateComponents() foundation.DateComponents
	SetStartDateComponents(value foundation.DateComponents)
	IsCompleted() bool
	SetIsCompleted(value bool)
}

// A class that represents a reminder in a calendar.
//
// Use the method to create a new reminder. Use the properties in the class to get and modify certain information about a reminder.


// A class that represents a reminder in a calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder
type EKReminder struct {
	EKCalendarItem
}

// EKReminderFrom constructs a [EKReminder] from an unsafe.Pointer.
//
// A class that represents a reminder in a calendar.
func EKReminderFrom(ptr unsafe.Pointer) EKReminder {
	return EKReminder{
		EKCalendarItem: EKCalendarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec _EKReminderClass) Alloc() EKReminder {
	rv := objc.Send[EKReminder](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EKReminderClass) New() EKReminder {
	rv := objc.Send[EKReminder](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKReminder) Init() EKReminder {
	rv := objc.Send[EKReminder](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKReminder) Autorelease() EKReminder {
	rv := objc.Send[EKReminder](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKReminder creates a new EKReminder instance.
func NewEKReminder() EKReminder {
	return getEKReminderClass().New()
}



// Creates and returns a new reminder in the given event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/init(eventStore:)
func NewEKReminderWithEventStore(eventStore IEKEventStore) EKReminder {
	rv := objc.Send[EKReminder](objc.ID(getEKReminderClass().class), objc.Sel("reminderWithEventStore:"), eventStore)
	return rv
}



// Creates and returns a new reminder in the given event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/init(eventStore:)
func (ec _EKReminderClass) ReminderWithEventStore(eventStore IEKEventStore) EKReminder {
	rv := objc.Send[EKReminder](objc.ID(ec.class), objc.Sel("reminderWithEventStore:"), eventStore)
	return rv
}


// The date on which the reminder was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/completionDate
func (e_ EKReminder) CompletionDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](e_.ID, objc.Sel("completionDate"))
	return rv
}


// The date on which the reminder was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/completionDate
func (e_ EKReminder) SetCompletionDate(value foundation.NSDate) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCompletionDate:"), value)
}


// The date by which the reminder should be completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/dueDateComponents
func (e_ EKReminder) DueDateComponents() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](e_.ID, objc.Sel("dueDateComponents"))
	return rv
}


// The date by which the reminder should be completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/dueDateComponents
func (e_ EKReminder) SetDueDateComponents(value foundation.DateComponents) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDueDateComponents:"), value)
}


// A Boolean value determining whether or not the reminder is marked completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/isCompleted
func (e_ EKReminder) Completed() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("completed"))
	return rv
}


// A Boolean value determining whether or not the reminder is marked completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/isCompleted
func (e_ EKReminder) SetCompleted(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCompleted:"), value)
}


// The reminder’s priority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/priority
func (e_ EKReminder) Priority() uint {
	rv := objc.Send[uint](e_.ID, objc.Sel("priority"))
	return rv
}


// The reminder’s priority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/priority
func (e_ EKReminder) SetPriority(value uint) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPriority:"), value)
}


// The start date of the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/startDateComponents
func (e_ EKReminder) StartDateComponents() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](e_.ID, objc.Sel("startDateComponents"))
	return rv
}


// The start date of the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/startDateComponents
func (e_ EKReminder) SetStartDateComponents(value foundation.DateComponents) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStartDateComponents:"), value)
}


// A Boolean value determining whether or not the reminder is marked completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekreminder/iscompleted
func (e_ EKReminder) IsCompleted() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isCompleted"))
	return rv
}


// A Boolean value determining whether or not the reminder is marked completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekreminder/iscompleted
func (e_ EKReminder) SetIsCompleted(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsCompleted:"), value)
}


