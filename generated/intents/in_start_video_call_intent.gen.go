// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INStartVideoCallIntent] class.
var (
	INStartVideoCallIntentClass     _INStartVideoCallIntentClass
	INStartVideoCallIntentClassOnce sync.Once
)

func getINStartVideoCallIntentClass() _INStartVideoCallIntentClass {
	INStartVideoCallIntentClassOnce.Do(func() {
		INStartVideoCallIntentClass = _INStartVideoCallIntentClass{objc.GetClass("INStartVideoCallIntent")}
	})
	return INStartVideoCallIntentClass
}

type _INStartVideoCallIntentClass struct {
	class objc.Class
}

// An interface definition for the [INStartVideoCallIntent] class.
type IINStartVideoCallIntent interface {
	IINIntent
	Contacts() INPerson
	SetContacts(value INPerson)
}

// A request to start a video call with one or more users.
//
// The system creates objects to let you know when the user wants to place a video call using your app. A video call intent object contains the users to include in the call. It’s up to you to match the information in this object to contacts in your app and to initiate the resulting call. Your Intents extension receives this intent when the user tries to initiate a call from the Siri interface. If your app supports CallKit, you may also receive this intent when the user tries to initiate a call from system interfaces such as the Recents tab of the Phone app. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler confirms the request and creates an object to indicate that it’s possible to begin the call. Don’t try to initiate calls directly from your Intents extension. SiriKit launches your app and passes it an object that your app must then use to initiate the call. SiriKit places an object in the user activity object with this intent. For calls initiated through Siri, the interaction object also includes the response provided by your Intents extension. For a list of other intents in the VoIP calling domain, see .


// A request to start a video call with one or more users.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartVideoCallIntent
type INStartVideoCallIntent struct {
	INIntent
}

// INStartVideoCallIntentFrom constructs a [INStartVideoCallIntent] from an unsafe.Pointer.
//
// A request to start a video call with one or more users.
func INStartVideoCallIntentFrom(ptr unsafe.Pointer) INStartVideoCallIntent {
	return INStartVideoCallIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INStartVideoCallIntentClass) Alloc() INStartVideoCallIntent {
	rv := objc.Send[INStartVideoCallIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INStartVideoCallIntentClass) New() INStartVideoCallIntent {
	rv := objc.Send[INStartVideoCallIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INStartVideoCallIntent) Init() INStartVideoCallIntent {
	rv := objc.Send[INStartVideoCallIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INStartVideoCallIntent) Autorelease() INStartVideoCallIntent {
	rv := objc.Send[INStartVideoCallIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINStartVideoCallIntent creates a new INStartVideoCallIntent instance.
func NewINStartVideoCallIntent() INStartVideoCallIntent {
	return getINStartVideoCallIntentClass().New()
}



// The users to call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartvideocallintent/contacts
func (i_ INStartVideoCallIntent) Contacts() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("contacts"))
	return rv
}


// The users to call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartvideocallintent/contacts
func (i_ INStartVideoCallIntent) SetContacts(value INPerson) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContacts:"), value)
}



