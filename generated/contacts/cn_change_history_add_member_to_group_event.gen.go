// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNChangeHistoryAddMemberToGroupEvent] class.
var (
	CNChangeHistoryAddMemberToGroupEventClass     _CNChangeHistoryAddMemberToGroupEventClass
	CNChangeHistoryAddMemberToGroupEventClassOnce sync.Once
)

func getCNChangeHistoryAddMemberToGroupEventClass() _CNChangeHistoryAddMemberToGroupEventClass {
	CNChangeHistoryAddMemberToGroupEventClassOnce.Do(func() {
		CNChangeHistoryAddMemberToGroupEventClass = _CNChangeHistoryAddMemberToGroupEventClass{objc.GetClass("CNChangeHistoryAddMemberToGroupEvent")}
	})
	return CNChangeHistoryAddMemberToGroupEventClass
}

type _CNChangeHistoryAddMemberToGroupEventClass struct {
	class objc.Class
}

// An interface definition for the [CNChangeHistoryAddMemberToGroupEvent] class.
type ICNChangeHistoryAddMemberToGroupEvent interface {
	ICNChangeHistoryEvent
	Group() ICNGroup
	Member() ICNContact
}

// An object that represents a user adding a contact to a group.


// An object that represents a user adding a contact to a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddMemberToGroupEvent
type CNChangeHistoryAddMemberToGroupEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryAddMemberToGroupEventFrom constructs a [CNChangeHistoryAddMemberToGroupEvent] from an unsafe.Pointer.
//
// An object that represents a user adding a contact to a group.
func CNChangeHistoryAddMemberToGroupEventFrom(ptr unsafe.Pointer) CNChangeHistoryAddMemberToGroupEvent {
	return CNChangeHistoryAddMemberToGroupEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryAddMemberToGroupEventClass) Alloc() CNChangeHistoryAddMemberToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddMemberToGroupEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNChangeHistoryAddMemberToGroupEventClass) New() CNChangeHistoryAddMemberToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddMemberToGroupEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryAddMemberToGroupEvent) Init() CNChangeHistoryAddMemberToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddMemberToGroupEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryAddMemberToGroupEvent) Autorelease() CNChangeHistoryAddMemberToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddMemberToGroupEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryAddMemberToGroupEvent creates a new CNChangeHistoryAddMemberToGroupEvent instance.
func NewCNChangeHistoryAddMemberToGroupEvent() CNChangeHistoryAddMemberToGroupEvent {
	return getCNChangeHistoryAddMemberToGroupEventClass().New()
}



// The group where the user added a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddMemberToGroupEvent/group
func (c_ CNChangeHistoryAddMemberToGroupEvent) Group() ICNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("group"))
	return rv
}


// The contact the user added to the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddMemberToGroupEvent/member
func (c_ CNChangeHistoryAddMemberToGroupEvent) Member() ICNContact {
	rv := objc.Send[CNContact](c_.ID, objc.Sel("member"))
	return rv
}



