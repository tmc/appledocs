// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNChangeHistoryDeleteGroupEvent] class.
var (
	CNChangeHistoryDeleteGroupEventClass     _CNChangeHistoryDeleteGroupEventClass
	CNChangeHistoryDeleteGroupEventClassOnce sync.Once
)

func getCNChangeHistoryDeleteGroupEventClass() _CNChangeHistoryDeleteGroupEventClass {
	CNChangeHistoryDeleteGroupEventClassOnce.Do(func() {
		CNChangeHistoryDeleteGroupEventClass = _CNChangeHistoryDeleteGroupEventClass{objc.GetClass("CNChangeHistoryDeleteGroupEvent")}
	})
	return CNChangeHistoryDeleteGroupEventClass
}

type _CNChangeHistoryDeleteGroupEventClass struct {
	class objc.Class
}

// An interface definition for the [CNChangeHistoryDeleteGroupEvent] class.
type ICNChangeHistoryDeleteGroupEvent interface {
	ICNChangeHistoryEvent
	// properties:
	GroupIdentifier() string /* primitive/slice/pointer. */
	// methods:
}

// An object that represents a user deleting a group.


// An object that represents a user deleting a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryDeleteGroupEvent
type CNChangeHistoryDeleteGroupEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryDeleteGroupEventFrom constructs a [CNChangeHistoryDeleteGroupEvent] from an unsafe.Pointer.
//
// An object that represents a user deleting a group.
func CNChangeHistoryDeleteGroupEventFrom(ptr unsafe.Pointer) CNChangeHistoryDeleteGroupEvent {
	return CNChangeHistoryDeleteGroupEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryDeleteGroupEventClass) Alloc() CNChangeHistoryDeleteGroupEvent {
	rv := objc.Send[CNChangeHistoryDeleteGroupEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNChangeHistoryDeleteGroupEventClass) New() CNChangeHistoryDeleteGroupEvent {
	rv := objc.Send[CNChangeHistoryDeleteGroupEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryDeleteGroupEvent) Init() CNChangeHistoryDeleteGroupEvent {
	rv := objc.Send[CNChangeHistoryDeleteGroupEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryDeleteGroupEvent) Autorelease() CNChangeHistoryDeleteGroupEvent {
	rv := objc.Send[CNChangeHistoryDeleteGroupEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryDeleteGroupEvent creates a new CNChangeHistoryDeleteGroupEvent instance.
func NewCNChangeHistoryDeleteGroupEvent() CNChangeHistoryDeleteGroupEvent {
	return getCNChangeHistoryDeleteGroupEventClass().New()
}



// A string that uniquely identifies the group that the user deleted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryDeleteGroupEvent/groupIdentifier
func (c_ CNChangeHistoryDeleteGroupEvent) GroupIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("groupIdentifier"))
	return rv
}



