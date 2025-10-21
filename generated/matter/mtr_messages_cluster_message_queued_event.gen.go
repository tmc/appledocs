// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMessagesClusterMessageQueuedEvent] class.
var (
	MTRMessagesClusterMessageQueuedEventClass     _MTRMessagesClusterMessageQueuedEventClass
	MTRMessagesClusterMessageQueuedEventClassOnce sync.Once
)

func getMTRMessagesClusterMessageQueuedEventClass() _MTRMessagesClusterMessageQueuedEventClass {
	MTRMessagesClusterMessageQueuedEventClassOnce.Do(func() {
		MTRMessagesClusterMessageQueuedEventClass = _MTRMessagesClusterMessageQueuedEventClass{objc.GetClass("MTRMessagesClusterMessageQueuedEvent")}
	})
	return MTRMessagesClusterMessageQueuedEventClass
}

type _MTRMessagesClusterMessageQueuedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRMessagesClusterMessageQueuedEvent] class.
type IMTRMessagesClusterMessageQueuedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageQueuedEvent
type MTRMessagesClusterMessageQueuedEvent struct {
	objectivec.Object
}

// MTRMessagesClusterMessageQueuedEventFrom constructs a [MTRMessagesClusterMessageQueuedEvent] from an unsafe.Pointer.
func MTRMessagesClusterMessageQueuedEventFrom(ptr unsafe.Pointer) MTRMessagesClusterMessageQueuedEvent {
	return MTRMessagesClusterMessageQueuedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMessagesClusterMessageQueuedEventClass) Alloc() MTRMessagesClusterMessageQueuedEvent {
	rv := objc.Send[MTRMessagesClusterMessageQueuedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMessagesClusterMessageQueuedEventClass) New() MTRMessagesClusterMessageQueuedEvent {
	rv := objc.Send[MTRMessagesClusterMessageQueuedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMessagesClusterMessageQueuedEvent) Init() MTRMessagesClusterMessageQueuedEvent {
	rv := objc.Send[MTRMessagesClusterMessageQueuedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMessagesClusterMessageQueuedEvent) Autorelease() MTRMessagesClusterMessageQueuedEvent {
	rv := objc.Send[MTRMessagesClusterMessageQueuedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMessagesClusterMessageQueuedEvent creates a new MTRMessagesClusterMessageQueuedEvent instance.
func NewMTRMessagesClusterMessageQueuedEvent() MTRMessagesClusterMessageQueuedEvent {
	return getMTRMessagesClusterMessageQueuedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageQueuedEvent/messageID
func (m_ MTRMessagesClusterMessageQueuedEvent) MessageID() foundation.NSData {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("messageID"))
	return rv
}


// SetMessageID sets the value of the messageID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageQueuedEvent/messageID
func (m_ MTRMessagesClusterMessageQueuedEvent) SetMessageID(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageID:"), value)
}



