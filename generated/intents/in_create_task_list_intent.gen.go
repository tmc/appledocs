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
	GroupName() INSpeakableString
	SetGroupName(value INSpeakableString)
	TaskTitles() INSpeakableString
	SetTaskTitles(value INSpeakableString)
	Title() INSpeakableString
	SetTitle(value INSpeakableString)
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


// The group that contains the task list.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/increatetasklistintent/groupname
func (i_ INCreateTaskListIntent) GroupName() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("groupName"))
	return rv
}


// SetGroupName sets the value of the groupName property.
// The group that contains the task list.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/increatetasklistintent/groupname
func (i_ INCreateTaskListIntent) SetGroupName(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGroupName:"), value)
}

// An array of strings containing the titles for individual tasks to add to the new task list.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/increatetasklistintent/tasktitles
func (i_ INCreateTaskListIntent) TaskTitles() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("taskTitles"))
	return rv
}


// SetTaskTitles sets the value of the taskTitles property.
// An array of strings containing the titles for individual tasks to add to the new task list.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/increatetasklistintent/tasktitles
func (i_ INCreateTaskListIntent) SetTaskTitles(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTaskTitles:"), value)
}

// The title of the task list.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/increatetasklistintent/title
func (i_ INCreateTaskListIntent) Title() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title of the task list.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/increatetasklistintent/title
func (i_ INCreateTaskListIntent) SetTitle(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTitle:"), value)
}



