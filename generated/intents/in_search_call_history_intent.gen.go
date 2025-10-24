// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSearchCallHistoryIntent] class.
var (
	INSearchCallHistoryIntentClass     _INSearchCallHistoryIntentClass
	INSearchCallHistoryIntentClassOnce sync.Once
)

func getINSearchCallHistoryIntentClass() _INSearchCallHistoryIntentClass {
	INSearchCallHistoryIntentClassOnce.Do(func() {
		INSearchCallHistoryIntentClass = _INSearchCallHistoryIntentClass{objc.GetClass("INSearchCallHistoryIntent")}
	})
	return INSearchCallHistoryIntentClass
}

type _INSearchCallHistoryIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSearchCallHistoryIntent] class.
type IINSearchCallHistoryIntent interface {
	IINIntent
	CallCapabilities() unsafe.Pointer
	SetCallCapabilities(value unsafe.Pointer)
	CallType() INCallRecordType
	SetCallType(value INCallRecordType)
	CallTypes() unsafe.Pointer
	SetCallTypes(value unsafe.Pointer)
	DateCreated() INDateComponentsRange
	SetDateCreated(value INDateComponentsRange)
	Recipient() INPerson
	SetRecipient(value INPerson)
	Unseen() bool
	SetUnseen(value bool)
}

// A request to list the calls matching the specified criteria.
//
// SiriKit creates objects when the user asks to see previous calls from their call history. This intent object contains the values for you to match when searching the user’s call history. Users can search for calls involving a specific person, calls that occurred on specific dates, or calls that are of a specific type such as missed calls. When performing the search, use only the parameters provided and ignore any that have no values. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler confirms the request and creates an object with the results of the search. For successful searches, Siri offers the user a way to launch your app and see the results.

// A request to list the calls matching the specified criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchCallHistoryIntent
type INSearchCallHistoryIntent struct {
	INIntent
}

// INSearchCallHistoryIntentFrom constructs a [INSearchCallHistoryIntent] from an unsafe.Pointer.
//
// A request to list the calls matching the specified criteria.
func INSearchCallHistoryIntentFrom(ptr unsafe.Pointer) INSearchCallHistoryIntent {
	return INSearchCallHistoryIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSearchCallHistoryIntentClass) Alloc() INSearchCallHistoryIntent {
	rv := objc.Send[INSearchCallHistoryIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSearchCallHistoryIntentClass) New() INSearchCallHistoryIntent {
	rv := objc.Send[INSearchCallHistoryIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSearchCallHistoryIntent) Init() INSearchCallHistoryIntent {
	rv := objc.Send[INSearchCallHistoryIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSearchCallHistoryIntent) Autorelease() INSearchCallHistoryIntent {
	rv := objc.Send[INSearchCallHistoryIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSearchCallHistoryIntent creates a new INSearchCallHistoryIntent instance.
func NewINSearchCallHistoryIntent() INSearchCallHistoryIntent {
	return getINSearchCallHistoryIntentClass().New()
}

// The audio-video capabilities of the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchcallhistoryintent/callcapabilities
func (i_ INSearchCallHistoryIntent) CallCapabilities() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("callCapabilities"))
	return rv
}

// The audio-video capabilities of the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchcallhistoryintent/callcapabilities
func (i_ INSearchCallHistoryIntent) SetCallCapabilities(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCallCapabilities:"), value)
}

// The call type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchcallhistoryintent/calltype
func (i_ INSearchCallHistoryIntent) CallType() INCallRecordType {
	rv := objc.Send[INCallRecordType](i_.ID, objc.Sel("callType"))
	return rv
}

// The call type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchcallhistoryintent/calltype
func (i_ INSearchCallHistoryIntent) SetCallType(value INCallRecordType) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCallType:"), value)
}

// The types of calls to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchcallhistoryintent/calltypes
func (i_ INSearchCallHistoryIntent) CallTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("callTypes"))
	return rv
}

// The types of calls to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchcallhistoryintent/calltypes
func (i_ INSearchCallHistoryIntent) SetCallTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCallTypes:"), value)
}

// The range of dates associated with the call records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchcallhistoryintent/datecreated
func (i_ INSearchCallHistoryIntent) DateCreated() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("dateCreated"))
	return rv
}

// The range of dates associated with the call records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchcallhistoryintent/datecreated
func (i_ INSearchCallHistoryIntent) SetDateCreated(value INDateComponentsRange) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDateCreated:"), value)
}

// The person involved in the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchcallhistoryintent/recipient
func (i_ INSearchCallHistoryIntent) Recipient() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("recipient"))
	return rv
}

// The person involved in the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchcallhistoryintent/recipient
func (i_ INSearchCallHistoryIntent) SetRecipient(value INPerson) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRecipient:"), value)
}

// A Boolean value that indicates whether the user has seen the call yet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchcallhistoryintent/unseen-9ua7o
func (i_ INSearchCallHistoryIntent) Unseen() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("unseen"))
	return rv
}

// A Boolean value that indicates whether the user has seen the call yet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchcallhistoryintent/unseen-9ua7o
func (i_ INSearchCallHistoryIntent) SetUnseen(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUnseen:"), value)
}
