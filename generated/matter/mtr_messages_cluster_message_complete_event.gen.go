// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	FutureMessagesPreference() objc.IObject /* cross-framework: NSNumber */
	SetFutureMessagesPreference(value objc.IObject /* cross-framework: NSNumber */)
	MessageID() objc.IObject /* cross-framework: NSData */
	SetMessageID(value objc.IObject /* cross-framework: NSData */)
	Reply() objc.IObject /* cross-framework: NSString */
	SetReply(value objc.IObject /* cross-framework: NSString */)
	ResponseID() objc.IObject /* cross-framework: NSNumber */
	SetResponseID(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/futureMessagesPreference
func (m_ MTRMessagesClusterMessageCompleteEvent) FutureMessagesPreference() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("futureMessagesPreference"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/futureMessagesPreference
func (m_ MTRMessagesClusterMessageCompleteEvent) SetFutureMessagesPreference(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFutureMessagesPreference:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/messageID
func (m_ MTRMessagesClusterMessageCompleteEvent) MessageID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("messageID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/messageID
func (m_ MTRMessagesClusterMessageCompleteEvent) SetMessageID(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/reply
func (m_ MTRMessagesClusterMessageCompleteEvent) Reply() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("reply"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/reply
func (m_ MTRMessagesClusterMessageCompleteEvent) SetReply(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReply:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/responseID
func (m_ MTRMessagesClusterMessageCompleteEvent) ResponseID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("responseID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/responseID
func (m_ MTRMessagesClusterMessageCompleteEvent) SetResponseID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResponseID:"), value)
}



