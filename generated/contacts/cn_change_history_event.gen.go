// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CNChangeHistoryEvent] class.
var (
	CNChangeHistoryEventClass     _CNChangeHistoryEventClass
	CNChangeHistoryEventClassOnce sync.Once
)

func getCNChangeHistoryEventClass() _CNChangeHistoryEventClass {
	CNChangeHistoryEventClassOnce.Do(func() {
		CNChangeHistoryEventClass = _CNChangeHistoryEventClass{objc.GetClass("CNChangeHistoryEvent")}
	})
	return CNChangeHistoryEventClass
}

type _CNChangeHistoryEventClass struct {
	class objc.Class
}

// An interface definition for the [CNChangeHistoryEvent] class.
type ICNChangeHistoryEvent interface {
	objectivec.IObject
	AcceptEventVisitor(visitor objc.ID)
}

// An object that represents the user adding, updating, or deleting a contact or group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryEvent
type CNChangeHistoryEvent struct {
	objectivec.Object
}

// CNChangeHistoryEventFrom constructs a [CNChangeHistoryEvent] from an unsafe.Pointer.
//
// An object that represents the user adding, updating, or deleting a contact or group.
func CNChangeHistoryEventFrom(ptr unsafe.Pointer) CNChangeHistoryEvent {
	return CNChangeHistoryEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryEventClass) Alloc() CNChangeHistoryEvent {
	rv := objc.Send[CNChangeHistoryEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNChangeHistoryEventClass) New() CNChangeHistoryEvent {
	rv := objc.Send[CNChangeHistoryEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryEvent) Init() CNChangeHistoryEvent {
	rv := objc.Send[CNChangeHistoryEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryEvent) Autorelease() CNChangeHistoryEvent {
	rv := objc.Send[CNChangeHistoryEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryEvent creates a new CNChangeHistoryEvent instance.
func NewCNChangeHistoryEvent() CNChangeHistoryEvent {
	return getCNChangeHistoryEventClass().New()
}


// Forwards the event to the delegate you provide to process the change-history event.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryEvent/accept(_:)
func (c_ CNChangeHistoryEvent) AcceptEventVisitor(visitor objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("acceptEventVisitor:"), visitor)
}



