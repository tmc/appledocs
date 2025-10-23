// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INCallRecord] class.
var (
	INCallRecordClass     _INCallRecordClass
	INCallRecordClassOnce sync.Once
)

func getINCallRecordClass() _INCallRecordClass {
	INCallRecordClassOnce.Do(func() {
		INCallRecordClass = _INCallRecordClass{objc.GetClass("INCallRecord")}
	})
	return INCallRecordClass
}

type _INCallRecordClass struct {
	class objc.Class
}

// An interface definition for the [INCallRecord] class.
type IINCallRecord interface {
	objectivec.IObject
	CallCapability() INCallCapability
	SetCallCapability(value INCallCapability)
	CallDuration() float64
	SetCallDuration(value float64)
	CallRecordType() INCallRecordType
	SetCallRecordType(value INCallRecordType)
	Caller() INPerson
	SetCaller(value INPerson)
	DateCreated() foundation.Date
	SetDateCreated(value foundation.IDate)
	Identifier() string
	SetIdentifier(value string)
	NumberOfCalls() int
	SetNumberOfCalls(value int)
	Participants() INPerson
	SetParticipants(value INPerson)
	Unseen() bool
	SetUnseen(value bool)
}

// The details about a call handled by your app.
//
// An object stores details about calls made by the user through your app. You use call record objects to communicate basic information about calls to SiriKit. A call record identifies the type of call, the duration of the call, the date and time of the call, and the person on the other end of the call. You create call record objects when reporting search results back to SiriKit and when identifying voicemails to play.


// The details about a call handled by your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallRecord
type INCallRecord struct {
	objectivec.Object
}

// INCallRecordFrom constructs a [INCallRecord] from an unsafe.Pointer.
//
// The details about a call handled by your app.
func INCallRecordFrom(ptr unsafe.Pointer) INCallRecord {
	return INCallRecord{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INCallRecordClass) Alloc() INCallRecord {
	rv := objc.Send[INCallRecord](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INCallRecordClass) New() INCallRecord {
	rv := objc.Send[INCallRecord](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INCallRecord) Init() INCallRecord {
	rv := objc.Send[INCallRecord](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INCallRecord) Autorelease() INCallRecord {
	rv := objc.Send[INCallRecord](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINCallRecord creates a new INCallRecord instance.
func NewINCallRecord() INCallRecord {
	return getINCallRecordClass().New()
}



// The audio and video capabilities of the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/callcapability
func (i_ INCallRecord) CallCapability() INCallCapability {
	rv := objc.Send[INCallCapability](i_.ID, objc.Sel("callCapability"))
	return rv
}


// The audio and video capabilities of the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/callcapability
func (i_ INCallRecord) SetCallCapability(value INCallCapability) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCallCapability:"), value)
}


// The duration (measured in seconds) of the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/callduration-47iud
func (i_ INCallRecord) CallDuration() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("callDuration"))
	return rv
}


// The duration (measured in seconds) of the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/callduration-47iud
func (i_ INCallRecord) SetCallDuration(value float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCallDuration:"), value)
}


// The type of call that resulted from the attempt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/callrecordtype
func (i_ INCallRecord) CallRecordType() INCallRecordType {
	rv := objc.Send[INCallRecordType](i_.ID, objc.Sel("callRecordType"))
	return rv
}


// The type of call that resulted from the attempt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/callrecordtype
func (i_ INCallRecord) SetCallRecordType(value INCallRecordType) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCallRecordType:"), value)
}


// The person who participated in the call with the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/caller
func (i_ INCallRecord) Caller() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("caller"))
	return rv
}


// The person who participated in the call with the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/caller
func (i_ INCallRecord) SetCaller(value INPerson) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCaller:"), value)
}


// The date and time at which the call was initiated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/datecreated
func (i_ INCallRecord) DateCreated() foundation.Date {
	rv := objc.Send[foundation.Date](i_.ID, objc.Sel("dateCreated"))
	return rv
}


// The date and time at which the call was initiated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/datecreated
func (i_ INCallRecord) SetDateCreated(value foundation.IDate) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDateCreated:"), value)
}


// A unique string that you can use to locate the call in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/identifier
func (i_ INCallRecord) Identifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("identifier"))
	return rv
}


// A unique string that you can use to locate the call in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/identifier
func (i_ INCallRecord) SetIdentifier(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The number of calls in the call record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/numberofcalls-r8kl
func (i_ INCallRecord) NumberOfCalls() int {
	rv := objc.Send[int](i_.ID, objc.Sel("numberOfCalls"))
	return rv
}


// The number of calls in the call record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/numberofcalls-r8kl
func (i_ INCallRecord) SetNumberOfCalls(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNumberOfCalls:"), value)
}


// The recipient of the user’s call request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/participants
func (i_ INCallRecord) Participants() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("participants"))
	return rv
}


// The recipient of the user’s call request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/participants
func (i_ INCallRecord) SetParticipants(value INPerson) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setParticipants:"), value)
}


// A Boolean value indicating whether the user has seen the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/unseen-2m7sz
func (i_ INCallRecord) Unseen() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("unseen"))
	return rv
}


// A Boolean value indicating whether the user has seen the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecord/unseen-2m7sz
func (i_ INCallRecord) SetUnseen(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUnseen:"), value)
}



