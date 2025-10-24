// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INCallRecordFilter] class.
var (
	INCallRecordFilterClass     _INCallRecordFilterClass
	INCallRecordFilterClassOnce sync.Once
)

func getINCallRecordFilterClass() _INCallRecordFilterClass {
	INCallRecordFilterClassOnce.Do(func() {
		INCallRecordFilterClass = _INCallRecordFilterClass{objc.GetClass("INCallRecordFilter")}
	})
	return INCallRecordFilterClass
}

type _INCallRecordFilterClass struct {
	class objc.Class
}

// An interface definition for the [INCallRecordFilter] class.
type IINCallRecordFilter interface {
	objectivec.IObject
	CallCapability() INCallCapability
	SetCallCapability(value INCallCapability)
	CallTypes() unsafe.Pointer
	SetCallTypes(value unsafe.Pointer)
	Participants() INPerson
	SetParticipants(value INPerson)
}

// Filters a user specifies to redial a call.
//
// Use this method to create filters contributed by the user to redial a call. The object identifies the person, type of call, and ability to make the call to initiate the user’s request.

// Filters a user specifies to redial a call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INCallRecordFilter
type INCallRecordFilter struct {
	objectivec.Object
}

// INCallRecordFilterFrom constructs a [INCallRecordFilter] from an unsafe.Pointer.
//
// Filters a user specifies to redial a call.
func INCallRecordFilterFrom(ptr unsafe.Pointer) INCallRecordFilter {
	return INCallRecordFilter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INCallRecordFilterClass) Alloc() INCallRecordFilter {
	rv := objc.Send[INCallRecordFilter](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INCallRecordFilterClass) New() INCallRecordFilter {
	rv := objc.Send[INCallRecordFilter](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INCallRecordFilter) Init() INCallRecordFilter {
	rv := objc.Send[INCallRecordFilter](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INCallRecordFilter) Autorelease() INCallRecordFilter {
	rv := objc.Send[INCallRecordFilter](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINCallRecordFilter creates a new INCallRecordFilter instance.
func NewINCallRecordFilter() INCallRecordFilter {
	return getINCallRecordFilterClass().New()
}

// An indicator of whether the call supports audio or video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecordfilter/callcapability
func (i_ INCallRecordFilter) CallCapability() INCallCapability {
	rv := objc.Send[INCallCapability](i_.ID, objc.Sel("callCapability"))
	return rv
}

// An indicator of whether the call supports audio or video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecordfilter/callcapability
func (i_ INCallRecordFilter) SetCallCapability(value INCallCapability) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCallCapability:"), value)
}

// The various call options that the user can requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecordfilter/calltypes
func (i_ INCallRecordFilter) CallTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("callTypes"))
	return rv
}

// The various call options that the user can requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecordfilter/calltypes
func (i_ INCallRecordFilter) SetCallTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCallTypes:"), value)
}

// The recipient of the user’s call request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecordfilter/participants
func (i_ INCallRecordFilter) Participants() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("participants"))
	return rv
}

// The recipient of the user’s call request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/incallrecordfilter/participants
func (i_ INCallRecordFilter) SetParticipants(value INPerson) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setParticipants:"), value)
}
