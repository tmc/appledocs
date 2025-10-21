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




