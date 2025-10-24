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
	// properties:
	CallRecordToCallBack() INCallRecord
	AudioRoute() unsafe.Pointer
	SetAudioRoute(value unsafe.Pointer)
	CallCapability() unsafe.Pointer
	SetCallCapability(value unsafe.Pointer)
	CallRecordFilter() INCallRecordFilter
	SetCallRecordFilter(value INCallRecordFilter)
	Contacts() INPerson
	SetContacts(value INPerson)
	DestinationType() unsafe.Pointer
	SetDestinationType(value unsafe.Pointer)
	RecordTypeForRedialing() unsafe.Pointer
	SetRecordTypeForRedialing(value unsafe.Pointer)
	// methods:
}

// A request to start an audio or video call with one or more users.
//
// SiriKit creates objects when the user wants to place a call using your app. A call intent object contains either the users to call or redialing information. It’s up to you to match the information in this object to contacts in your app and initiate the resulting call. Your Intents extension receives this intent when the user tries to initiate a call from the Siri interface. If your app supports CallKit, you may also receive this intent when the user tries to initiate a call from system interfaces, such as the Recents tab of the Phone app. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object that indicates it’s possible to begin the call. Don’t try to initiate calls directly from your Intents extension. Instead, SiriKit launches your app and passes it an object that your app must then use to initiate the call. SiriKit places an object in the user activity object with this intent. For calls initiated through Siri, the interaction object also includes the response provided by your Intents extension.


// A request to start an audio or video call with one or more users.
//
// [Full Topic]
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



// Details about a call to redial a missed call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartCallIntent/callRecordToCallBack
func (i_ INStartCallIntent) CallRecordToCallBack() INCallRecord {
	rv := objc.Send[INCallRecord](i_.ID, objc.Sel("callRecordToCallBack"))
	return rv
}


// The audio route the call is using.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartcallintent/audioroute
func (i_ INStartCallIntent) AudioRoute() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("audioRoute"))
	return rv
}


// The audio route the call is using.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartcallintent/audioroute
func (i_ INStartCallIntent) SetAudioRoute(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAudioRoute:"), value)
}


// The type of call the user initiated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartcallintent/callcapability
func (i_ INStartCallIntent) CallCapability() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("callCapability"))
	return rv
}


// The type of call the user initiated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartcallintent/callcapability
func (i_ INStartCallIntent) SetCallCapability(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCallCapability:"), value)
}


// Filters specified by the user to redial a call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartcallintent/callrecordfilter
func (i_ INStartCallIntent) CallRecordFilter() INCallRecordFilter {
	rv := objc.Send[INCallRecordFilter](i_.ID, objc.Sel("callRecordFilter"))
	return rv
}


// Filters specified by the user to redial a call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartcallintent/callrecordfilter
func (i_ INStartCallIntent) SetCallRecordFilter(value INCallRecordFilter) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCallRecordFilter:"), value)
}


// The users to call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartcallintent/contacts
func (i_ INStartCallIntent) Contacts() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("contacts"))
	return rv
}


// The users to call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartcallintent/contacts
func (i_ INStartCallIntent) SetContacts(value INPerson) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContacts:"), value)
}


// The type of call to place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartcallintent/destinationtype
func (i_ INStartCallIntent) DestinationType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("destinationType"))
	return rv
}


// The type of call to place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartcallintent/destinationtype
func (i_ INStartCallIntent) SetDestinationType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDestinationType:"), value)
}


// The category of past call record contact information used for redialing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartcallintent/recordtypeforredialing
func (i_ INStartCallIntent) RecordTypeForRedialing() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("recordTypeForRedialing"))
	return rv
}


// The category of past call record contact information used for redialing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartcallintent/recordtypeforredialing
func (i_ INStartCallIntent) SetRecordTypeForRedialing(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRecordTypeForRedialing:"), value)
}



