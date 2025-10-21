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
}

// A request to add tasks to an existing task list.
//
// Siri creates an instance of when the user asks to add one or more tasks to a task list. Alternatively, if the user asks to create a note, and your app or Intents extension doesn’t support , SiriKit uses this intent instead. The intent contains the tasks to add and the target task list, and can also include location or time triggers that you assign to each of the new tasks. To process the request, your handler must adopt the protocol. is available to both Siri Intents and Siri Suggestions and doesn’t require an unlocked device before processing.
//
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




