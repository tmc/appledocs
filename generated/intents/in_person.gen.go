// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [INPerson] class.
var (
	INPersonClass     _INPersonClass
	INPersonClassOnce sync.Once
)

func getINPersonClass() _INPersonClass {
	INPersonClassOnce.Do(func() {
		INPersonClass = _INPersonClass{objc.GetClass("INPerson")}
	})
	return INPersonClass
}

type _INPersonClass struct {
	class objc.Class
}

// An interface definition for the [INPerson] class.
type IINPerson interface {
	objectivec.IObject
}

// Information about a person participating in a SiriKit interaction.
//
// SiriKit uses objects to represent people with many different roles, including the sender or recipient of calls and messages, the payer or payee of a financial transaction, or the driver of a vehicle. You also use person objects to identify the corresponding contact in your app and to communicate information about that contact back to SiriKit. When resolving the parameters of an intent, use any provided objects to identify the corresponding contacts in your app. A person object contains information provided by the initial request, which could be as little as a single name spoken by the person interacting with Siri. After identifying the contact, create a new object and fill it with the information that you need to identify that contact again later. For example, you might specify a value for property that contains the information about how your app identifies that contact. When resolving the identities of contacts, SiriKit leverages the information in the device owner’s contacts database when that information is available. If the owner denies your app access to their contacts, SiriKit can’t use that information, which might cause many properties of a person object to be . Because the class conforms to the protocol, though, SiriKit still populates the property with what the person interacting with Siri said, and you can use that information to try to identify the contact. For more information about that protocol, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPerson
type INPerson struct {
	objectivec.Object
}

// INPersonFrom constructs a [INPerson] from an unsafe.Pointer.
//
// Information about a person participating in a SiriKit interaction.
func INPersonFrom(ptr unsafe.Pointer) INPerson {
	return INPerson{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INPersonClass) Alloc() INPerson {
	rv := objc.Send[INPerson](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INPersonClass) New() INPerson {
	rv := objc.Send[INPerson](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INPerson) Init() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INPerson) Autorelease() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINPerson creates a new INPerson instance.
func NewINPerson() INPerson {
	return getINPersonClass().New()
}




