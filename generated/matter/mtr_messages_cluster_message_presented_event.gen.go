// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMessagesClusterMessagePresentedEvent] class.
var (
	MTRMessagesClusterMessagePresentedEventClass     _MTRMessagesClusterMessagePresentedEventClass
	MTRMessagesClusterMessagePresentedEventClassOnce sync.Once
)

func getMTRMessagesClusterMessagePresentedEventClass() _MTRMessagesClusterMessagePresentedEventClass {
	MTRMessagesClusterMessagePresentedEventClassOnce.Do(func() {
		MTRMessagesClusterMessagePresentedEventClass = _MTRMessagesClusterMessagePresentedEventClass{objc.GetClass("MTRMessagesClusterMessagePresentedEvent")}
	})
	return MTRMessagesClusterMessagePresentedEventClass
}

type _MTRMessagesClusterMessagePresentedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRMessagesClusterMessagePresentedEvent] class.
type IMTRMessagesClusterMessagePresentedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessagePresentedEvent
type MTRMessagesClusterMessagePresentedEvent struct {
	objectivec.Object
}

// MTRMessagesClusterMessagePresentedEventFrom constructs a [MTRMessagesClusterMessagePresentedEvent] from an unsafe.Pointer.
func MTRMessagesClusterMessagePresentedEventFrom(ptr unsafe.Pointer) MTRMessagesClusterMessagePresentedEvent {
	return MTRMessagesClusterMessagePresentedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMessagesClusterMessagePresentedEventClass) Alloc() MTRMessagesClusterMessagePresentedEvent {
	rv := objc.Send[MTRMessagesClusterMessagePresentedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMessagesClusterMessagePresentedEventClass) New() MTRMessagesClusterMessagePresentedEvent {
	rv := objc.Send[MTRMessagesClusterMessagePresentedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMessagesClusterMessagePresentedEvent) Init() MTRMessagesClusterMessagePresentedEvent {
	rv := objc.Send[MTRMessagesClusterMessagePresentedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMessagesClusterMessagePresentedEvent) Autorelease() MTRMessagesClusterMessagePresentedEvent {
	rv := objc.Send[MTRMessagesClusterMessagePresentedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMessagesClusterMessagePresentedEvent creates a new MTRMessagesClusterMessagePresentedEvent instance.
func NewMTRMessagesClusterMessagePresentedEvent() MTRMessagesClusterMessagePresentedEvent {
	return getMTRMessagesClusterMessagePresentedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessagePresentedEvent/messageID
func (m_ MTRMessagesClusterMessagePresentedEvent) MessageID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("messageID"))
	return rv
}


// SetMessageID sets the value of the messageID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessagePresentedEvent/messageID
func (m_ MTRMessagesClusterMessagePresentedEvent) SetMessageID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageID:"), value)
}


