// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNChangeHistoryUpdateGroupEvent] class.
var (
	CNChangeHistoryUpdateGroupEventClass     _CNChangeHistoryUpdateGroupEventClass
	CNChangeHistoryUpdateGroupEventClassOnce sync.Once
)

func getCNChangeHistoryUpdateGroupEventClass() _CNChangeHistoryUpdateGroupEventClass {
	CNChangeHistoryUpdateGroupEventClassOnce.Do(func() {
		CNChangeHistoryUpdateGroupEventClass = _CNChangeHistoryUpdateGroupEventClass{objc.GetClass("CNChangeHistoryUpdateGroupEvent")}
	})
	return CNChangeHistoryUpdateGroupEventClass
}

type _CNChangeHistoryUpdateGroupEventClass struct {
	class objc.Class
}

// An interface definition for the [CNChangeHistoryUpdateGroupEvent] class.
type ICNChangeHistoryUpdateGroupEvent interface {
	ICNChangeHistoryEvent
	// properties:
	Group() ICNGroup
	// methods:
}

// An object that represents an updated group event.


// An object that represents an updated group event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryUpdateGroupEvent
type CNChangeHistoryUpdateGroupEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryUpdateGroupEventFrom constructs a [CNChangeHistoryUpdateGroupEvent] from an unsafe.Pointer.
//
// An object that represents an updated group event.
func CNChangeHistoryUpdateGroupEventFrom(ptr unsafe.Pointer) CNChangeHistoryUpdateGroupEvent {
	return CNChangeHistoryUpdateGroupEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryUpdateGroupEventClass) Alloc() CNChangeHistoryUpdateGroupEvent {
	rv := objc.Send[CNChangeHistoryUpdateGroupEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNChangeHistoryUpdateGroupEventClass) New() CNChangeHistoryUpdateGroupEvent {
	rv := objc.Send[CNChangeHistoryUpdateGroupEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryUpdateGroupEvent) Init() CNChangeHistoryUpdateGroupEvent {
	rv := objc.Send[CNChangeHistoryUpdateGroupEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryUpdateGroupEvent) Autorelease() CNChangeHistoryUpdateGroupEvent {
	rv := objc.Send[CNChangeHistoryUpdateGroupEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryUpdateGroupEvent creates a new CNChangeHistoryUpdateGroupEvent instance.
func NewCNChangeHistoryUpdateGroupEvent() CNChangeHistoryUpdateGroupEvent {
	return getCNChangeHistoryUpdateGroupEventClass().New()
}



// The group that the user updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryUpdateGroupEvent/group
func (c_ CNChangeHistoryUpdateGroupEvent) Group() ICNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("group"))
	return rv
}



