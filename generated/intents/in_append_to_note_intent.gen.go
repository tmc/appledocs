// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INAppendToNoteIntent] class.
var (
	INAppendToNoteIntentClass     _INAppendToNoteIntentClass
	INAppendToNoteIntentClassOnce sync.Once
)

func getINAppendToNoteIntentClass() _INAppendToNoteIntentClass {
	INAppendToNoteIntentClassOnce.Do(func() {
		INAppendToNoteIntentClass = _INAppendToNoteIntentClass{objc.GetClass("INAppendToNoteIntent")}
	})
	return INAppendToNoteIntentClass
}

type _INAppendToNoteIntentClass struct {
	class objc.Class
}

// An interface definition for the [INAppendToNoteIntent] class.
type IINAppendToNoteIntent interface {
	IINIntent
}

// A request to append content to a note.
//
// Siri creates an object when the user asks to append content to an existing note. The intent object contains the note to modify and the content to add to the note. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the updated note.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INAppendToNoteIntent
type INAppendToNoteIntent struct {
	INIntent
}

// INAppendToNoteIntentFrom constructs a [INAppendToNoteIntent] from an unsafe.Pointer.
//
// A request to append content to a note.
func INAppendToNoteIntentFrom(ptr unsafe.Pointer) INAppendToNoteIntent {
	return INAppendToNoteIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INAppendToNoteIntentClass) Alloc() INAppendToNoteIntent {
	rv := objc.Send[INAppendToNoteIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INAppendToNoteIntentClass) New() INAppendToNoteIntent {
	rv := objc.Send[INAppendToNoteIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INAppendToNoteIntent) Init() INAppendToNoteIntent {
	rv := objc.Send[INAppendToNoteIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INAppendToNoteIntent) Autorelease() INAppendToNoteIntent {
	rv := objc.Send[INAppendToNoteIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINAppendToNoteIntent creates a new INAppendToNoteIntent instance.
func NewINAppendToNoteIntent() INAppendToNoteIntent {
	return getINAppendToNoteIntentClass().New()
}


// The note to receive the additional content.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inappendtonoteintent/targetnote
func (i_ INAppendToNoteIntent) TargetNote() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("targetNote"))
	return rv
}


// SetTargetNote sets the value of the targetNote property.
// The note to receive the additional content.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inappendtonoteintent/targetnote
func (i_ INAppendToNoteIntent) SetTargetNote(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTargetNote:"), value)
}

// The content to append to the note.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inappendtonoteintent/content
func (i_ INAppendToNoteIntent) Content() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("content"))
	return rv
}


// SetContent sets the value of the content property.
// The content to append to the note.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inappendtonoteintent/content
func (i_ INAppendToNoteIntent) SetContent(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContent:"), value)
}



