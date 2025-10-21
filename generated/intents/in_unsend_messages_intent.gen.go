// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INUnsendMessagesIntent] class.
var (
	INUnsendMessagesIntentClass     _INUnsendMessagesIntentClass
	INUnsendMessagesIntentClassOnce sync.Once
)

func getINUnsendMessagesIntentClass() _INUnsendMessagesIntentClass {
	INUnsendMessagesIntentClassOnce.Do(func() {
		INUnsendMessagesIntentClass = _INUnsendMessagesIntentClass{objc.GetClass("INUnsendMessagesIntent")}
	})
	return INUnsendMessagesIntentClass
}

type _INUnsendMessagesIntentClass struct {
	class objc.Class
}

// An interface definition for the [INUnsendMessagesIntent] class.
type IINUnsendMessagesIntent interface {
	IINIntent
}

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INUnsendMessagesIntent
type INUnsendMessagesIntent struct {
	INIntent
}

// INUnsendMessagesIntentFrom constructs a [INUnsendMessagesIntent] from an unsafe.Pointer.
func INUnsendMessagesIntentFrom(ptr unsafe.Pointer) INUnsendMessagesIntent {
	return INUnsendMessagesIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INUnsendMessagesIntentClass) Alloc() INUnsendMessagesIntent {
	rv := objc.Send[INUnsendMessagesIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INUnsendMessagesIntentClass) New() INUnsendMessagesIntent {
	rv := objc.Send[INUnsendMessagesIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INUnsendMessagesIntent) Init() INUnsendMessagesIntent {
	rv := objc.Send[INUnsendMessagesIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INUnsendMessagesIntent) Autorelease() INUnsendMessagesIntent {
	rv := objc.Send[INUnsendMessagesIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINUnsendMessagesIntent creates a new INUnsendMessagesIntent instance.
func NewINUnsendMessagesIntent() INUnsendMessagesIntent {
	return getINUnsendMessagesIntentClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/intents/inunsendmessagesintent/messageidentifiers
func (i_ INUnsendMessagesIntent) MessageIdentifiers() string {
	rv := objc.Send[string](i_.ID, objc.Sel("messageIdentifiers"))
	return rv
}


// SetMessageIdentifiers sets the value of the messageIdentifiers property.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inunsendmessagesintent/messageidentifiers
func (i_ INUnsendMessagesIntent) SetMessageIdentifiers(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMessageIdentifiers:"), objc.String(value))
}



