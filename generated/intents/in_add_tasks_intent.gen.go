// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INAddTasksIntent] class.
var (
	INAddTasksIntentClass     _INAddTasksIntentClass
	INAddTasksIntentClassOnce sync.Once
)

func getINAddTasksIntentClass() _INAddTasksIntentClass {
	INAddTasksIntentClassOnce.Do(func() {
		INAddTasksIntentClass = _INAddTasksIntentClass{objc.GetClass("INAddTasksIntent")}
	})
	return INAddTasksIntentClass
}

type _INAddTasksIntentClass struct {
	class objc.Class
}

// An interface definition for the [INAddTasksIntent] class.
type IINAddTasksIntent interface {
	IINIntent
	Priority() unsafe.Pointer
	SetPriority(value unsafe.Pointer)
	SpatialEventTrigger() unsafe.Pointer
	SetSpatialEventTrigger(value unsafe.Pointer)
	TargetTaskList() unsafe.Pointer
	SetTargetTaskList(value unsafe.Pointer)
	TaskTitles() INSpeakableString
	SetTaskTitles(value INSpeakableString)
	TemporalEventTrigger() unsafe.Pointer
	SetTemporalEventTrigger(value unsafe.Pointer)
}

// A request to add tasks to an existing task list.
//
// Siri creates an instance of when the user asks to add one or more tasks to a task list. Alternatively, if the user asks to create a note, and your app or Intents extension doesn’t support , SiriKit uses this intent instead. The intent contains the tasks to add and the target task list, and can also include location or time triggers that you assign to each of the new tasks. To process the request, your handler must adopt the protocol. is available to both Siri Intents and Siri Suggestions and doesn’t require an unlocked device before processing.

// A request to add tasks to an existing task list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INAddTasksIntent
type INAddTasksIntent struct {
	INIntent
}

// INAddTasksIntentFrom constructs a [INAddTasksIntent] from an unsafe.Pointer.
//
// A request to add tasks to an existing task list.
func INAddTasksIntentFrom(ptr unsafe.Pointer) INAddTasksIntent {
	return INAddTasksIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INAddTasksIntentClass) Alloc() INAddTasksIntent {
	rv := objc.Send[INAddTasksIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INAddTasksIntentClass) New() INAddTasksIntent {
	rv := objc.Send[INAddTasksIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INAddTasksIntent) Init() INAddTasksIntent {
	rv := objc.Send[INAddTasksIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INAddTasksIntent) Autorelease() INAddTasksIntent {
	rv := objc.Send[INAddTasksIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINAddTasksIntent creates a new INAddTasksIntent instance.
func NewINAddTasksIntent() INAddTasksIntent {
	return getINAddTasksIntentClass().New()
}

// The priority for the new task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddtasksintent/priority
func (i_ INAddTasksIntent) Priority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("priority"))
	return rv
}

// The priority for the new task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddtasksintent/priority
func (i_ INAddTasksIntent) SetPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPriority:"), value)
}

// The location-based trigger to associate with each of the new tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddtasksintent/spatialeventtrigger
func (i_ INAddTasksIntent) SpatialEventTrigger() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("spatialEventTrigger"))
	return rv
}

// The location-based trigger to associate with each of the new tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddtasksintent/spatialeventtrigger
func (i_ INAddTasksIntent) SetSpatialEventTrigger(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSpatialEventTrigger:"), value)
}

// The task list to receive the new tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddtasksintent/targettasklist
func (i_ INAddTasksIntent) TargetTaskList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("targetTaskList"))
	return rv
}

// The task list to receive the new tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddtasksintent/targettasklist
func (i_ INAddTasksIntent) SetTargetTaskList(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTargetTaskList:"), value)
}

// An array of strings containing the titles of the new tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddtasksintent/tasktitles
func (i_ INAddTasksIntent) TaskTitles() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("taskTitles"))
	return rv
}

// An array of strings containing the titles of the new tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddtasksintent/tasktitles
func (i_ INAddTasksIntent) SetTaskTitles(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTaskTitles:"), value)
}

// The time-based trigger to associate with each of the new tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddtasksintent/temporaleventtrigger
func (i_ INAddTasksIntent) TemporalEventTrigger() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("temporalEventTrigger"))
	return rv
}

// The time-based trigger to associate with each of the new tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddtasksintent/temporaleventtrigger
func (i_ INAddTasksIntent) SetTemporalEventTrigger(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTemporalEventTrigger:"), value)
}
