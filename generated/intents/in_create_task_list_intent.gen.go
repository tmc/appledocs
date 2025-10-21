// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INCreateTaskListIntent] class.
var (
	INCreateTaskListIntentClass     _INCreateTaskListIntentClass
	INCreateTaskListIntentClassOnce sync.Once
)

func getINCreateTaskListIntentClass() _INCreateTaskListIntentClass {
	INCreateTaskListIntentClassOnce.Do(func() {
		INCreateTaskListIntentClass = _INCreateTaskListIntentClass{objc.GetClass("INCreateTaskListIntent")}
	})
	return INCreateTaskListIntentClass
}

type _INCreateTaskListIntentClass struct {
	class objc.Class
}

// An interface definition for the [INCreateTaskListIntent] class.
type IINCreateTaskListIntent interface {
	IINIntent
}

// A request to create a new task list.
//
// Siri creates an object when the user asks to create a new task list. The intent object can contain the title of the task list, a set of initial tasks, and possibly the group in which to create the task list. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the updated task list.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCreateTaskListIntent
type INCreateTaskListIntent struct {
	INIntent
}

// INCreateTaskListIntentFrom constructs a [INCreateTaskListIntent] from an unsafe.Pointer.
//
// A request to create a new task list.
func INCreateTaskListIntentFrom(ptr unsafe.Pointer) INCreateTaskListIntent {
	return INCreateTaskListIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INCreateTaskListIntentClass) Alloc() INCreateTaskListIntent {
	rv := objc.Send[INCreateTaskListIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INCreateTaskListIntentClass) New() INCreateTaskListIntent {
	rv := objc.Send[INCreateTaskListIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INCreateTaskListIntent) Init() INCreateTaskListIntent {
	rv := objc.Send[INCreateTaskListIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INCreateTaskListIntent) Autorelease() INCreateTaskListIntent {
	rv := objc.Send[INCreateTaskListIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINCreateTaskListIntent creates a new INCreateTaskListIntent instance.
func NewINCreateTaskListIntent() INCreateTaskListIntent {
	return getINCreateTaskListIntentClass().New()
}




