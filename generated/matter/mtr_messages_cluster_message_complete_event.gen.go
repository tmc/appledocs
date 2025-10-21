// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMessagesClusterMessageCompleteEvent] class.
var (
	MTRMessagesClusterMessageCompleteEventClass     _MTRMessagesClusterMessageCompleteEventClass
	MTRMessagesClusterMessageCompleteEventClassOnce sync.Once
)

func getMTRMessagesClusterMessageCompleteEventClass() _MTRMessagesClusterMessageCompleteEventClass {
	MTRMessagesClusterMessageCompleteEventClassOnce.Do(func() {
		MTRMessagesClusterMessageCompleteEventClass = _MTRMessagesClusterMessageCompleteEventClass{objc.GetClass("MTRMessagesClusterMessageCompleteEvent")}
	})
	return MTRMessagesClusterMessageCompleteEventClass
}

type _MTRMessagesClusterMessageCompleteEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRMessagesClusterMessageCompleteEvent] class.
type IMTRMessagesClusterMessageCompleteEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent
type MTRMessagesClusterMessageCompleteEvent struct {
	objectivec.Object
}

// MTRMessagesClusterMessageCompleteEventFrom constructs a [MTRMessagesClusterMessageCompleteEvent] from an unsafe.Pointer.
func MTRMessagesClusterMessageCompleteEventFrom(ptr unsafe.Pointer) MTRMessagesClusterMessageCompleteEvent {
	return MTRMessagesClusterMessageCompleteEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMessagesClusterMessageCompleteEventClass) Alloc() MTRMessagesClusterMessageCompleteEvent {
	rv := objc.Send[MTRMessagesClusterMessageCompleteEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMessagesClusterMessageCompleteEventClass) New() MTRMessagesClusterMessageCompleteEvent {
	rv := objc.Send[MTRMessagesClusterMessageCompleteEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMessagesClusterMessageCompleteEvent) Init() MTRMessagesClusterMessageCompleteEvent {
	rv := objc.Send[MTRMessagesClusterMessageCompleteEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMessagesClusterMessageCompleteEvent) Autorelease() MTRMessagesClusterMessageCompleteEvent {
	rv := objc.Send[MTRMessagesClusterMessageCompleteEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMessagesClusterMessageCompleteEvent creates a new MTRMessagesClusterMessageCompleteEvent instance.
func NewMTRMessagesClusterMessageCompleteEvent() MTRMessagesClusterMessageCompleteEvent {
	return getMTRMessagesClusterMessageCompleteEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/futureMessagesPreference
func (m_ MTRMessagesClusterMessageCompleteEvent) FutureMessagesPreference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("futureMessagesPreference"))
	return rv
}


// SetFutureMessagesPreference sets the value of the futureMessagesPreference property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/futureMessagesPreference
func (m_ MTRMessagesClusterMessageCompleteEvent) SetFutureMessagesPreference(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFutureMessagesPreference:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/messageID
func (m_ MTRMessagesClusterMessageCompleteEvent) MessageID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("messageID"))
	return rv
}


// SetMessageID sets the value of the messageID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/messageID
func (m_ MTRMessagesClusterMessageCompleteEvent) SetMessageID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageID:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/reply
func (m_ MTRMessagesClusterMessageCompleteEvent) Reply() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("reply"))
	return rv
}


// SetReply sets the value of the reply property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/reply
func (m_ MTRMessagesClusterMessageCompleteEvent) SetReply(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReply:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/responseID
func (m_ MTRMessagesClusterMessageCompleteEvent) ResponseID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("responseID"))
	return rv
}


// SetResponseID sets the value of the responseID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/responseID
func (m_ MTRMessagesClusterMessageCompleteEvent) SetResponseID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResponseID:"), value)
}


