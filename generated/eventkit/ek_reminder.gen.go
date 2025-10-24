// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class EKReminder */


/* debug [class_header]: Header for EKReminder */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EKReminder */
// An interface definition for the [EKReminder] class.
type IEKReminder interface {
	IEKCalendarItem
	
/* debug [class_interface_properties]: Properties for EKReminder */
	// properties:
	CompletionDate() objc.IObject /* cross-framework: NSDate */
	SetCompletionDate(value objc.IObject /* cross-framework: NSDate */)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EKReminder */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EKReminder */
// Alloc allocates a new instance without initialization.
func (ec _EKReminderClass) Alloc() EKReminder {
	rv := objc.Send[EKReminder](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EKReminder */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EKReminder */

// Creates and returns a new reminder in the given event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/init(eventStore:)
func NewEKReminderWithEventStore(eventStore IEKEventStore) EKReminder {
	rv := objc.Send[EKReminder](objc.ID(getEKReminderClass().class), objc.Sel("reminderWithEventStore:"), eventStore)
	return rv
}/* debug [class_init_methods/constructor]: NewEKReminderWithEventStore */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EKReminder */

// Creates and returns a new reminder in the given event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/init(eventStore:)
func (ec _EKReminderClass) ReminderWithEventStore(eventStore IEKEventStore) EKReminder {
	rv := objc.Send[EKReminder](objc.ID(ec.class), objc.Sel("reminderWithEventStore:"), eventStore)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReminderWithEventStore) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EKReminder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EKReminder */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EKReminder */

// The date on which the reminder was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/completionDate
func (e_ EKReminder) CompletionDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](e_.ID, objc.Sel("completionDate"))
	return rv
}/* debug [instance_properties/getter]: completionDate */


// The date on which the reminder was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/completionDate
func (e_ EKReminder) SetCompletionDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCompletionDate:"), value)
}/* debug [instance_properties/setter]: completionDate */


// The date by which the reminder should be completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/dueDateComponents
func (e_ EKReminder) DueDateComponents() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](e_.ID, objc.Sel("dueDateComponents"))
	return rv
}/* debug [instance_properties/getter]: dueDateComponents */


// The date by which the reminder should be completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/dueDateComponents
func (e_ EKReminder) SetDueDateComponents(value foundation.DateComponents) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDueDateComponents:"), value)
}/* debug [instance_properties/setter]: dueDateComponents */


// A Boolean value determining whether or not the reminder is marked completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/isCompleted
func (e_ EKReminder) Completed() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("completed"))
	return rv
}/* debug [instance_properties/getter]: completed */


// A Boolean value determining whether or not the reminder is marked completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/isCompleted
func (e_ EKReminder) SetCompleted(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCompleted:"), value)
}/* debug [instance_properties/setter]: completed */


// The reminder’s priority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/priority
func (e_ EKReminder) Priority() uint {
	rv := objc.Send[uint](e_.ID, objc.Sel("priority"))
	return rv
}/* debug [instance_properties/getter]: priority */


// The reminder’s priority.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/priority
func (e_ EKReminder) SetPriority(value uint) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setPriority:"), value)
}/* debug [instance_properties/setter]: priority */


// The start date of the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/startDateComponents
func (e_ EKReminder) StartDateComponents() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](e_.ID, objc.Sel("startDateComponents"))
	return rv
}/* debug [instance_properties/getter]: startDateComponents */


// The start date of the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKReminder/startDateComponents
func (e_ EKReminder) SetStartDateComponents(value foundation.DateComponents) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStartDateComponents:"), value)
}/* debug [instance_properties/setter]: startDateComponents */


// A Boolean value determining whether or not the reminder is marked completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekreminder/iscompleted
func (e_ EKReminder) IsCompleted() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isCompleted"))
	return rv
}/* debug [instance_properties/getter]: isCompleted */


// A Boolean value determining whether or not the reminder is marked completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekreminder/iscompleted
func (e_ EKReminder) SetIsCompleted(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsCompleted:"), value)
}/* debug [instance_properties/setter]: isCompleted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EKReminder */


