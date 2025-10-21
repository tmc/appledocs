// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INCreateNoteIntent] class.
var (
	INCreateNoteIntentClass     _INCreateNoteIntentClass
	INCreateNoteIntentClassOnce sync.Once
)

func getINCreateNoteIntentClass() _INCreateNoteIntentClass {
	INCreateNoteIntentClassOnce.Do(func() {
		INCreateNoteIntentClass = _INCreateNoteIntentClass{objc.GetClass("INCreateNoteIntent")}
	})
	return INCreateNoteIntentClass
}

type _INCreateNoteIntentClass struct {
	class objc.Class
}

// An interface definition for the [INCreateNoteIntent] class.
type IINCreateNoteIntent interface {
	IINIntent
}

// A request to create a new note.
//
// Siri creates an object when the user asks to create a new note. (If the user asks to add a task and none of your Intents app extensions declare support for the class, SiriKit uses this intent as a substitute). The intent object contains the content of the note and possibly additional information, such as the name of the folder in which to create the note. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the updated note.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCreateNoteIntent
type INCreateNoteIntent struct {
	INIntent
}

// INCreateNoteIntentFrom constructs a [INCreateNoteIntent] from an unsafe.Pointer.
//
// A request to create a new note.
func INCreateNoteIntentFrom(ptr unsafe.Pointer) INCreateNoteIntent {
	return INCreateNoteIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INCreateNoteIntentClass) Alloc() INCreateNoteIntent {
	rv := objc.Send[INCreateNoteIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INCreateNoteIntentClass) New() INCreateNoteIntent {
	rv := objc.Send[INCreateNoteIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INCreateNoteIntent) Init() INCreateNoteIntent {
	rv := objc.Send[INCreateNoteIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INCreateNoteIntent) Autorelease() INCreateNoteIntent {
	rv := objc.Send[INCreateNoteIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINCreateNoteIntent creates a new INCreateNoteIntent instance.
func NewINCreateNoteIntent() INCreateNoteIntent {
	return getINCreateNoteIntentClass().New()
}




