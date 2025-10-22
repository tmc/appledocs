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
	Content() unsafe.Pointer
	SetContent(value unsafe.Pointer)
	GroupName() INSpeakableString
	SetGroupName(value INSpeakableString)
	Title() INSpeakableString
	SetTitle(value INSpeakableString)
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


// The main content of the note.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/increatenoteintent/content
func (i_ INCreateNoteIntent) Content() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("content"))
	return rv
}


// SetContent sets the value of the content property.
// The main content of the note.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/increatenoteintent/content
func (i_ INCreateNoteIntent) SetContent(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContent:"), value)
}

// The group that contains the note.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/increatenoteintent/groupname
func (i_ INCreateNoteIntent) GroupName() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("groupName"))
	return rv
}


// SetGroupName sets the value of the groupName property.
// The group that contains the note.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/increatenoteintent/groupname
func (i_ INCreateNoteIntent) SetGroupName(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGroupName:"), value)
}

// The title text for the note.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/increatenoteintent/title
func (i_ INCreateNoteIntent) Title() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title text for the note.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/increatenoteintent/title
func (i_ INCreateNoteIntent) SetTitle(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTitle:"), value)
}



