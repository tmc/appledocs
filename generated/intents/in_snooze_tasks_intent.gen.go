// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSnoozeTasksIntent] class.
var (
	INSnoozeTasksIntentClass     _INSnoozeTasksIntentClass
	INSnoozeTasksIntentClassOnce sync.Once
)

func getINSnoozeTasksIntentClass() _INSnoozeTasksIntentClass {
	INSnoozeTasksIntentClassOnce.Do(func() {
		INSnoozeTasksIntentClass = _INSnoozeTasksIntentClass{objc.GetClass("INSnoozeTasksIntent")}
	})
	return INSnoozeTasksIntentClass
}

type _INSnoozeTasksIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSnoozeTasksIntent] class.
type IINSnoozeTasksIntent interface {
	IINIntent
	All() bool
	SetAll(value bool)
	NextTriggerTime() INDateComponentsRange
	SetNextTriggerTime(value INDateComponentsRange)
	Tasks() unsafe.Pointer
	SetTasks(value unsafe.Pointer)
}

// A request to snooze one or more tasks.
//
// Siri creates an object when the user marks one or more tasks for snoozing. The intent object can contain the task information. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the updated task information.


// A request to snooze one or more tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSnoozeTasksIntent
type INSnoozeTasksIntent struct {
	INIntent
}

// INSnoozeTasksIntentFrom constructs a [INSnoozeTasksIntent] from an unsafe.Pointer.
//
// A request to snooze one or more tasks.
func INSnoozeTasksIntentFrom(ptr unsafe.Pointer) INSnoozeTasksIntent {
	return INSnoozeTasksIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSnoozeTasksIntentClass) Alloc() INSnoozeTasksIntent {
	rv := objc.Send[INSnoozeTasksIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSnoozeTasksIntentClass) New() INSnoozeTasksIntent {
	rv := objc.Send[INSnoozeTasksIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSnoozeTasksIntent) Init() INSnoozeTasksIntent {
	rv := objc.Send[INSnoozeTasksIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSnoozeTasksIntent) Autorelease() INSnoozeTasksIntent {
	rv := objc.Send[INSnoozeTasksIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSnoozeTasksIntent creates a new INSnoozeTasksIntent instance.
func NewINSnoozeTasksIntent() INSnoozeTasksIntent {
	return getINSnoozeTasksIntentClass().New()
}



// A Boolean value that indicates whether to snooze all of the tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insnoozetasksintent/all-spsb
func (i_ INSnoozeTasksIntent) All() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("all"))
	return rv
}


// A Boolean value that indicates whether to snooze all of the tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insnoozetasksintent/all-spsb
func (i_ INSnoozeTasksIntent) SetAll(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAll:"), value)
}


// The next time after the current time that triggers a task to snooze.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insnoozetasksintent/nexttriggertime
func (i_ INSnoozeTasksIntent) NextTriggerTime() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("nextTriggerTime"))
	return rv
}


// The next time after the current time that triggers a task to snooze.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insnoozetasksintent/nexttriggertime
func (i_ INSnoozeTasksIntent) SetNextTriggerTime(value INDateComponentsRange) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNextTriggerTime:"), value)
}


// An array of tasks to snooze.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insnoozetasksintent/tasks
func (i_ INSnoozeTasksIntent) Tasks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("tasks"))
	return rv
}


// An array of tasks to snooze.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insnoozetasksintent/tasks
func (i_ INSnoozeTasksIntent) SetTasks(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTasks:"), value)
}



