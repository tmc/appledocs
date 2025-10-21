// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INEditMessageIntent] class.
var (
	INEditMessageIntentClass     _INEditMessageIntentClass
	INEditMessageIntentClassOnce sync.Once
)

func getINEditMessageIntentClass() _INEditMessageIntentClass {
	INEditMessageIntentClassOnce.Do(func() {
		INEditMessageIntentClass = _INEditMessageIntentClass{objc.GetClass("INEditMessageIntent")}
	})
	return INEditMessageIntentClass
}

type _INEditMessageIntentClass struct {
	class objc.Class
}

// An interface definition for the [INEditMessageIntent] class.
type IINEditMessageIntent interface {
	IINIntent
}

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INEditMessageIntent
type INEditMessageIntent struct {
	INIntent
}

// INEditMessageIntentFrom constructs a [INEditMessageIntent] from an unsafe.Pointer.
func INEditMessageIntentFrom(ptr unsafe.Pointer) INEditMessageIntent {
	return INEditMessageIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INEditMessageIntentClass) Alloc() INEditMessageIntent {
	rv := objc.Send[INEditMessageIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INEditMessageIntentClass) New() INEditMessageIntent {
	rv := objc.Send[INEditMessageIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INEditMessageIntent) Init() INEditMessageIntent {
	rv := objc.Send[INEditMessageIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INEditMessageIntent) Autorelease() INEditMessageIntent {
	rv := objc.Send[INEditMessageIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINEditMessageIntent creates a new INEditMessageIntent instance.
func NewINEditMessageIntent() INEditMessageIntent {
	return getINEditMessageIntentClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/intents/ineditmessageintent/editedcontent
func (i_ INEditMessageIntent) EditedContent() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("editedContent"))
	return rv
}


// SetEditedContent sets the value of the editedContent property.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/ineditmessageintent/editedcontent
func (i_ INEditMessageIntent) SetEditedContent(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEditedContent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/intents/ineditmessageintent/messageidentifier
func (i_ INEditMessageIntent) MessageIdentifier() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("messageIdentifier"))
	return rv
}


// SetMessageIdentifier sets the value of the messageIdentifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/ineditmessageintent/messageidentifier
func (i_ INEditMessageIntent) SetMessageIdentifier(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMessageIdentifier:"), value)
}



