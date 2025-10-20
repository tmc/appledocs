// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNChangeHistoryRemoveMemberFromGroupEvent] class.
var (
	CNChangeHistoryRemoveMemberFromGroupEventClass     _CNChangeHistoryRemoveMemberFromGroupEventClass
	CNChangeHistoryRemoveMemberFromGroupEventClassOnce sync.Once
)

func getCNChangeHistoryRemoveMemberFromGroupEventClass() _CNChangeHistoryRemoveMemberFromGroupEventClass {
	CNChangeHistoryRemoveMemberFromGroupEventClassOnce.Do(func() {
		CNChangeHistoryRemoveMemberFromGroupEventClass = _CNChangeHistoryRemoveMemberFromGroupEventClass{objc.GetClass("CNChangeHistoryRemoveMemberFromGroupEvent")}
	})
	return CNChangeHistoryRemoveMemberFromGroupEventClass
}

type _CNChangeHistoryRemoveMemberFromGroupEventClass struct {
	class objc.Class
}

// An interface definition for the [CNChangeHistoryRemoveMemberFromGroupEvent] class.
type ICNChangeHistoryRemoveMemberFromGroupEvent interface {
	ICNChangeHistoryEvent
}

// An object that represents a user removing a contact from a group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryRemoveMemberFromGroupEvent
type CNChangeHistoryRemoveMemberFromGroupEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryRemoveMemberFromGroupEventFrom constructs a [CNChangeHistoryRemoveMemberFromGroupEvent] from an unsafe.Pointer.
//
// An object that represents a user removing a contact from a group.
func CNChangeHistoryRemoveMemberFromGroupEventFrom(ptr unsafe.Pointer) CNChangeHistoryRemoveMemberFromGroupEvent {
	return CNChangeHistoryRemoveMemberFromGroupEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryRemoveMemberFromGroupEventClass) Alloc() CNChangeHistoryRemoveMemberFromGroupEvent {
	rv := objc.Send[CNChangeHistoryRemoveMemberFromGroupEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNChangeHistoryRemoveMemberFromGroupEventClass) New() CNChangeHistoryRemoveMemberFromGroupEvent {
	rv := objc.Send[CNChangeHistoryRemoveMemberFromGroupEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryRemoveMemberFromGroupEvent) Init() CNChangeHistoryRemoveMemberFromGroupEvent {
	rv := objc.Send[CNChangeHistoryRemoveMemberFromGroupEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryRemoveMemberFromGroupEvent) Autorelease() CNChangeHistoryRemoveMemberFromGroupEvent {
	rv := objc.Send[CNChangeHistoryRemoveMemberFromGroupEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryRemoveMemberFromGroupEvent creates a new CNChangeHistoryRemoveMemberFromGroupEvent instance.
func NewCNChangeHistoryRemoveMemberFromGroupEvent() CNChangeHistoryRemoveMemberFromGroupEvent {
	return getCNChangeHistoryRemoveMemberFromGroupEventClass().New()
}


// The group where the user removed a contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryRemoveMemberFromGroupEvent/group
func (c_ CNChangeHistoryRemoveMemberFromGroupEvent) Group() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("group"))
	return rv
}

// The contact that the user removed from the group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryRemoveMemberFromGroupEvent/member
func (c_ CNChangeHistoryRemoveMemberFromGroupEvent) Member() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("member"))
	return rv
}



