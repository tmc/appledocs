// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INStartAudioCallIntent] class.
var (
	INStartAudioCallIntentClass     _INStartAudioCallIntentClass
	INStartAudioCallIntentClassOnce sync.Once
)

func getINStartAudioCallIntentClass() _INStartAudioCallIntentClass {
	INStartAudioCallIntentClassOnce.Do(func() {
		INStartAudioCallIntentClass = _INStartAudioCallIntentClass{objc.GetClass("INStartAudioCallIntent")}
	})
	return INStartAudioCallIntentClass
}

type _INStartAudioCallIntentClass struct {
	class objc.Class
}

// An interface definition for the [INStartAudioCallIntent] class.
type IINStartAudioCallIntent interface {
	IINIntent
}

// A request to start an audio-only call with one or more users.
//
// SiriKit creates objects when the user wants to place an audio call using your app. An audio call intent object contains the users to call. Your intent handler matches the information in this object to contacts in your app and to initiate the resulting call. Your Intents extension receives this intent when the user tries to initiate a call from the Siri interface. If your app supports CallKit, you may also receive this intent when the user tries to initiate a call from system interfaces such as the Recents tab of the Phone app. To handle this intent, the handler object in your Intents extension must adopt the protocol. Use your handler to confirm the request and create an object to indicate that it’s possible to begin the call. Don’t try to initiate calls directly from your Intents extension. SiriKit launches your app and passes it an object that your app must then use to initiate the call. SiriKit places an object in the user activity object with this intent. For calls initiated through Siri, the interaction object also includes the response provided by your Intents extension. For a list of other intents in the VoIP calling domain, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartAudioCallIntent
type INStartAudioCallIntent struct {
	INIntent
}

// INStartAudioCallIntentFrom constructs a [INStartAudioCallIntent] from an unsafe.Pointer.
//
// A request to start an audio-only call with one or more users.
func INStartAudioCallIntentFrom(ptr unsafe.Pointer) INStartAudioCallIntent {
	return INStartAudioCallIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INStartAudioCallIntentClass) Alloc() INStartAudioCallIntent {
	rv := objc.Send[INStartAudioCallIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INStartAudioCallIntentClass) New() INStartAudioCallIntent {
	rv := objc.Send[INStartAudioCallIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INStartAudioCallIntent) Init() INStartAudioCallIntent {
	rv := objc.Send[INStartAudioCallIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INStartAudioCallIntent) Autorelease() INStartAudioCallIntent {
	rv := objc.Send[INStartAudioCallIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINStartAudioCallIntent creates a new INStartAudioCallIntent instance.
func NewINStartAudioCallIntent() INStartAudioCallIntent {
	return getINStartAudioCallIntentClass().New()
}


// The type of call to place.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartaudiocallintent/destinationtype
func (i_ INStartAudioCallIntent) DestinationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("destinationType"))
	return rv
}


// SetDestinationType sets the value of the destinationType property.
// The type of call to place.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartaudiocallintent/destinationtype
func (i_ INStartAudioCallIntent) SetDestinationType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDestinationType:"), value)
}

// The users to call.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartaudiocallintent/contacts
func (i_ INStartAudioCallIntent) Contacts() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("contacts"))
	return rv
}


// SetContacts sets the value of the contacts property.
// The users to call.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartaudiocallintent/contacts
func (i_ INStartAudioCallIntent) SetContacts(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContacts:"), value)
}



