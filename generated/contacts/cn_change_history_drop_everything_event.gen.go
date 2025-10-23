// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNChangeHistoryDropEverythingEvent] class.
var (
	CNChangeHistoryDropEverythingEventClass     _CNChangeHistoryDropEverythingEventClass
	CNChangeHistoryDropEverythingEventClassOnce sync.Once
)

func getCNChangeHistoryDropEverythingEventClass() _CNChangeHistoryDropEverythingEventClass {
	CNChangeHistoryDropEverythingEventClassOnce.Do(func() {
		CNChangeHistoryDropEverythingEventClass = _CNChangeHistoryDropEverythingEventClass{objc.GetClass("CNChangeHistoryDropEverythingEvent")}
	})
	return CNChangeHistoryDropEverythingEventClass
}

type _CNChangeHistoryDropEverythingEventClass struct {
	class objc.Class
}

// An interface definition for the [CNChangeHistoryDropEverythingEvent] class.
type ICNChangeHistoryDropEverythingEvent interface {
	ICNChangeHistoryEvent
	// properties:
	// methods:
}

// An object that indicates the delegate should drop all contacts and groups before handling change events.
//
// The system sends this event to your delegate when the system determines that enough has changed since the last time your app fetched the history changes that an incremental update is no longer possible. Following the drop-everything event, your app receives an add event for each contact and group currently in the database. This allows you to implement full syncs and incremental syncs using the same code.


// An object that indicates the delegate should drop all contacts and groups before handling change events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryDropEverythingEvent
type CNChangeHistoryDropEverythingEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryDropEverythingEventFrom constructs a [CNChangeHistoryDropEverythingEvent] from an unsafe.Pointer.
//
// An object that indicates the delegate should drop all contacts and groups before handling change events.
func CNChangeHistoryDropEverythingEventFrom(ptr unsafe.Pointer) CNChangeHistoryDropEverythingEvent {
	return CNChangeHistoryDropEverythingEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryDropEverythingEventClass) Alloc() CNChangeHistoryDropEverythingEvent {
	rv := objc.Send[CNChangeHistoryDropEverythingEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNChangeHistoryDropEverythingEventClass) New() CNChangeHistoryDropEverythingEvent {
	rv := objc.Send[CNChangeHistoryDropEverythingEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryDropEverythingEvent) Init() CNChangeHistoryDropEverythingEvent {
	rv := objc.Send[CNChangeHistoryDropEverythingEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryDropEverythingEvent) Autorelease() CNChangeHistoryDropEverythingEvent {
	rv := objc.Send[CNChangeHistoryDropEverythingEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryDropEverythingEvent creates a new CNChangeHistoryDropEverythingEvent instance.
func NewCNChangeHistoryDropEverythingEvent() CNChangeHistoryDropEverythingEvent {
	return getCNChangeHistoryDropEverythingEventClass().New()
}




