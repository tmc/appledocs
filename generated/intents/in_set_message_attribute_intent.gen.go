// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSetMessageAttributeIntent] class.
var (
	INSetMessageAttributeIntentClass     _INSetMessageAttributeIntentClass
	INSetMessageAttributeIntentClassOnce sync.Once
)

func getINSetMessageAttributeIntentClass() _INSetMessageAttributeIntentClass {
	INSetMessageAttributeIntentClassOnce.Do(func() {
		INSetMessageAttributeIntentClass = _INSetMessageAttributeIntentClass{objc.GetClass("INSetMessageAttributeIntent")}
	})
	return INSetMessageAttributeIntentClass
}

type _INSetMessageAttributeIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSetMessageAttributeIntent] class.
type IINSetMessageAttributeIntent interface {
	IINIntent
}

// A request to modify the attributes of a message.
//
// Siri creates an object when the user asks to modify the attributes of one or more messages. Attributes represent information about the message such as whether the user read or flagged the message. This intent object includes the messages to modify and which attributes to change. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the results of modifying the messages.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetMessageAttributeIntent
type INSetMessageAttributeIntent struct {
	INIntent
}

// INSetMessageAttributeIntentFrom constructs a [INSetMessageAttributeIntent] from an unsafe.Pointer.
//
// A request to modify the attributes of a message.
func INSetMessageAttributeIntentFrom(ptr unsafe.Pointer) INSetMessageAttributeIntent {
	return INSetMessageAttributeIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSetMessageAttributeIntentClass) Alloc() INSetMessageAttributeIntent {
	rv := objc.Send[INSetMessageAttributeIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSetMessageAttributeIntentClass) New() INSetMessageAttributeIntent {
	rv := objc.Send[INSetMessageAttributeIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSetMessageAttributeIntent) Init() INSetMessageAttributeIntent {
	rv := objc.Send[INSetMessageAttributeIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSetMessageAttributeIntent) Autorelease() INSetMessageAttributeIntent {
	rv := objc.Send[INSetMessageAttributeIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSetMessageAttributeIntent creates a new INSetMessageAttributeIntent instance.
func NewINSetMessageAttributeIntent() INSetMessageAttributeIntent {
	return getINSetMessageAttributeIntentClass().New()
}




// Initializes the set message attribute intent object with the specified identifiers and attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetMessageAttributeIntent/init(identifiers:attribute:)
func NewINSetMessageAttributeIntentWithIdentifiersAttribute(identifiers unsafe.Pointer, attribute unsafe.Pointer) INSetMessageAttributeIntent {
	instance := getINSetMessageAttributeIntentClass().Alloc()
	rv := objc.Send[INSetMessageAttributeIntent](instance.ID, objc.Sel("initWithIdentifiers:attribute:"), identifiers, attribute)
	rv.Autorelease()
	return rv
}


// The attribute to apply to the messages.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetMessageAttributeIntent/attribute
func (i_ INSetMessageAttributeIntent) Attribute() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("attribute"))
	return rv
}

// The array of message identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetMessageAttributeIntent/identifiers
func (i_ INSetMessageAttributeIntent) Identifiers() []string {
	rv := objc.Send[[]string](i_.ID, objc.Sel("identifiers"))
	return rv
}


