// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INDeleteTasksIntent] class.
var (
	INDeleteTasksIntentClass     _INDeleteTasksIntentClass
	INDeleteTasksIntentClassOnce sync.Once
)

func getINDeleteTasksIntentClass() _INDeleteTasksIntentClass {
	INDeleteTasksIntentClassOnce.Do(func() {
		INDeleteTasksIntentClass = _INDeleteTasksIntentClass{objc.GetClass("INDeleteTasksIntent")}
	})
	return INDeleteTasksIntentClass
}

type _INDeleteTasksIntentClass struct {
	class objc.Class
}

// An interface definition for the [INDeleteTasksIntent] class.
type IINDeleteTasksIntent interface {
	IINIntent
}

// A request to delete one or more tasks.
//
// Siri creates an object when the user marks one or more tasks for deletion. The intent object can contain the task information. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the updated task information.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INDeleteTasksIntent
type INDeleteTasksIntent struct {
	INIntent
}

// INDeleteTasksIntentFrom constructs a [INDeleteTasksIntent] from an unsafe.Pointer.
//
// A request to delete one or more tasks.
func INDeleteTasksIntentFrom(ptr unsafe.Pointer) INDeleteTasksIntent {
	return INDeleteTasksIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INDeleteTasksIntentClass) Alloc() INDeleteTasksIntent {
	rv := objc.Send[INDeleteTasksIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INDeleteTasksIntentClass) New() INDeleteTasksIntent {
	rv := objc.Send[INDeleteTasksIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INDeleteTasksIntent) Init() INDeleteTasksIntent {
	rv := objc.Send[INDeleteTasksIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INDeleteTasksIntent) Autorelease() INDeleteTasksIntent {
	rv := objc.Send[INDeleteTasksIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINDeleteTasksIntent creates a new INDeleteTasksIntent instance.
func NewINDeleteTasksIntent() INDeleteTasksIntent {
	return getINDeleteTasksIntentClass().New()
}




