// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSetTaskAttributeIntent] class.
var (
	INSetTaskAttributeIntentClass     _INSetTaskAttributeIntentClass
	INSetTaskAttributeIntentClassOnce sync.Once
)

func getINSetTaskAttributeIntentClass() _INSetTaskAttributeIntentClass {
	INSetTaskAttributeIntentClassOnce.Do(func() {
		INSetTaskAttributeIntentClass = _INSetTaskAttributeIntentClass{objc.GetClass("INSetTaskAttributeIntent")}
	})
	return INSetTaskAttributeIntentClass
}

type _INSetTaskAttributeIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSetTaskAttributeIntent] class.
type IINSetTaskAttributeIntent interface {
	IINIntent
	Priority() unsafe.Pointer
	SetPriority(value unsafe.Pointer)
	SpatialEventTrigger() unsafe.Pointer
	SetSpatialEventTrigger(value unsafe.Pointer)
	Status() unsafe.Pointer
	SetStatus(value unsafe.Pointer)
	TargetTask() unsafe.Pointer
	SetTargetTask(value unsafe.Pointer)
	TaskTitle() INSpeakableString
	SetTaskTitle(value INSpeakableString)
	TemporalEventTrigger() unsafe.Pointer
	SetTemporalEventTrigger(value unsafe.Pointer)
}

// A request to modify the attributes of a task.
//
// Siri creates an object when the user marks a task as complete or changes the triggers used to generate reminders for the task. The intent object can contain the task information and values for any changed attributes. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the updated task information.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetTaskAttributeIntent
type INSetTaskAttributeIntent struct {
	INIntent
}

// INSetTaskAttributeIntentFrom constructs a [INSetTaskAttributeIntent] from an unsafe.Pointer.
//
// A request to modify the attributes of a task.
func INSetTaskAttributeIntentFrom(ptr unsafe.Pointer) INSetTaskAttributeIntent {
	return INSetTaskAttributeIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSetTaskAttributeIntentClass) Alloc() INSetTaskAttributeIntent {
	rv := objc.Send[INSetTaskAttributeIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSetTaskAttributeIntentClass) New() INSetTaskAttributeIntent {
	rv := objc.Send[INSetTaskAttributeIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSetTaskAttributeIntent) Init() INSetTaskAttributeIntent {
	rv := objc.Send[INSetTaskAttributeIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSetTaskAttributeIntent) Autorelease() INSetTaskAttributeIntent {
	rv := objc.Send[INSetTaskAttributeIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSetTaskAttributeIntent creates a new INSetTaskAttributeIntent instance.
func NewINSetTaskAttributeIntent() INSetTaskAttributeIntent {
	return getINSetTaskAttributeIntentClass().New()
}


// The priority for the task.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insettaskattributeintent/priority
func (i_ INSetTaskAttributeIntent) Priority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("priority"))
	return rv
}


// SetPriority sets the value of the priority property.
// The priority for the task.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insettaskattributeintent/priority
func (i_ INSetTaskAttributeIntent) SetPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPriority:"), value)
}

// The location-based trigger to apply to the task.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insettaskattributeintent/spatialeventtrigger
func (i_ INSetTaskAttributeIntent) SpatialEventTrigger() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("spatialEventTrigger"))
	return rv
}


// SetSpatialEventTrigger sets the value of the spatialEventTrigger property.
// The location-based trigger to apply to the task.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insettaskattributeintent/spatialeventtrigger
func (i_ INSetTaskAttributeIntent) SetSpatialEventTrigger(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSpatialEventTrigger:"), value)
}

// The status to apply to the task.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insettaskattributeintent/status
func (i_ INSetTaskAttributeIntent) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
// The status to apply to the task.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insettaskattributeintent/status
func (i_ INSetTaskAttributeIntent) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setStatus:"), value)
}

// The task targetted for modification.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insettaskattributeintent/targettask
func (i_ INSetTaskAttributeIntent) TargetTask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("targetTask"))
	return rv
}


// SetTargetTask sets the value of the targetTask property.
// The task targetted for modification.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insettaskattributeintent/targettask
func (i_ INSetTaskAttributeIntent) SetTargetTask(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTargetTask:"), value)
}

// The title for the task.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insettaskattributeintent/tasktitle
func (i_ INSetTaskAttributeIntent) TaskTitle() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("taskTitle"))
	return rv
}


// SetTaskTitle sets the value of the taskTitle property.
// The title for the task.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insettaskattributeintent/tasktitle
func (i_ INSetTaskAttributeIntent) SetTaskTitle(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTaskTitle:"), value)
}

// The time-based trigger to apply to the task.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insettaskattributeintent/temporaleventtrigger
func (i_ INSetTaskAttributeIntent) TemporalEventTrigger() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("temporalEventTrigger"))
	return rv
}


// SetTemporalEventTrigger sets the value of the temporalEventTrigger property.
// The time-based trigger to apply to the task.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insettaskattributeintent/temporaleventtrigger
func (i_ INSetTaskAttributeIntent) SetTemporalEventTrigger(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTemporalEventTrigger:"), value)
}



