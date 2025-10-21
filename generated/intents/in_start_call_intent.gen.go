// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INStartCallIntent] class.
var (
	INStartCallIntentClass     _INStartCallIntentClass
	INStartCallIntentClassOnce sync.Once
)

func getINStartCallIntentClass() _INStartCallIntentClass {
	INStartCallIntentClassOnce.Do(func() {
		INStartCallIntentClass = _INStartCallIntentClass{objc.GetClass("INStartCallIntent")}
	})
	return INStartCallIntentClass
}

type _INStartCallIntentClass struct {
	class objc.Class
}

// An interface definition for the [INStartCallIntent] class.
type IINStartCallIntent interface {
	IINIntent
}

// A request to start an audio or video call with one or more users.
//
// SiriKit creates objects when the user wants to place a call using your app. A call intent object contains either the users to call or redialing information. It’s up to you to match the information in this object to contacts in your app and initiate the resulting call. Your Intents extension receives this intent when the user tries to initiate a call from the Siri interface. If your app supports CallKit, you may also receive this intent when the user tries to initiate a call from system interfaces, such as the Recents tab of the Phone app. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object that indicates it’s possible to begin the call. Don’t try to initiate calls directly from your Intents extension. Instead, SiriKit launches your app and passes it an object that your app must then use to initiate the call. SiriKit places an object in the user activity object with this intent. For calls initiated through Siri, the interaction object also includes the response provided by your Intents extension.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntent
type INStartCallIntent struct {
	INIntent
}

// INStartCallIntentFrom constructs a [INStartCallIntent] from an unsafe.Pointer.
//
// A request to start an audio or video call with one or more users.
func INStartCallIntentFrom(ptr unsafe.Pointer) INStartCallIntent {
	return INStartCallIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INStartCallIntentClass) Alloc() INStartCallIntent {
	rv := objc.Send[INStartCallIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INStartCallIntentClass) New() INStartCallIntent {
	rv := objc.Send[INStartCallIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INStartCallIntent) Init() INStartCallIntent {
	rv := objc.Send[INStartCallIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INStartCallIntent) Autorelease() INStartCallIntent {
	rv := objc.Send[INStartCallIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINStartCallIntent creates a new INStartCallIntent instance.
func NewINStartCallIntent() INStartCallIntent {
	return getINStartCallIntentClass().New()
}


// Creates a start call intent object with the specified parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntent/init(audioRoute:destinationType:contacts:recordTypeForRedialing:callCapability:)
func NewINStartCallIntentWithAudioRouteDestinationTypeContactsRecordTypeForRedialingCallCapability(audioRoute unsafe.Pointer, destinationType unsafe.Pointer, contacts unsafe.Pointer, recordTypeForRedialing unsafe.Pointer, callCapability unsafe.Pointer) INStartCallIntent {
	instance := getINStartCallIntentClass().Alloc()
	rv := objc.Send[INStartCallIntent](instance.ID, objc.Sel("initWithAudioRoute:destinationType:contacts:recordTypeForRedialing:callCapability:"), audioRoute, destinationType, contacts, recordTypeForRedialing, callCapability)
	rv.Autorelease()
	return rv
}

// Creates a start call intent object with the specified parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntent/init(callRecordFilter:callRecordToCallBack:audioRoute:destinationType:contacts:callCapability:)
func NewINStartCallIntentWithCallRecordFilterCallRecordToCallBackAudioRouteDestinationTypeContactsCallCapability(callRecordFilter unsafe.Pointer, callRecordToCallBack unsafe.Pointer, audioRoute unsafe.Pointer, destinationType unsafe.Pointer, contacts unsafe.Pointer, callCapability unsafe.Pointer) INStartCallIntent {
	instance := getINStartCallIntentClass().Alloc()
	rv := objc.Send[INStartCallIntent](instance.ID, objc.Sel("initWithCallRecordFilter:callRecordToCallBack:audioRoute:destinationType:contacts:callCapability:"), callRecordFilter, callRecordToCallBack, audioRoute, destinationType, contacts, callCapability)
	rv.Autorelease()
	return rv
}


// The audio route the call is using.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntent/audioRoute
func (i_ INStartCallIntent) AudioRoute() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("audioRoute"))
	return rv
}

// The type of call the user initiated.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntent/callCapability
func (i_ INStartCallIntent) CallCapability() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("callCapability"))
	return rv
}

// Filters specified by the user to redial a call.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntent/callRecordFilter
func (i_ INStartCallIntent) CallRecordFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("callRecordFilter"))
	return rv
}

// Details about a call to redial a missed call.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntent/callRecordToCallBack
func (i_ INStartCallIntent) CallRecordToCallBack() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("callRecordToCallBack"))
	return rv
}

// The users to call.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntent/contacts
func (i_ INStartCallIntent) Contacts() []INPerson {
	rv := objc.Send[[]INPerson](i_.ID, objc.Sel("contacts"))
	return rv
}

// The type of call to place.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntent/destinationType
func (i_ INStartCallIntent) DestinationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("destinationType"))
	return rv
}

// The category of past call record contact information used for redialing.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntent/recordTypeForRedialing
func (i_ INStartCallIntent) RecordTypeForRedialing() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("recordTypeForRedialing"))
	return rv
}


