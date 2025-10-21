// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNChangeHistoryAddSubgroupToGroupEvent] class.
var (
	CNChangeHistoryAddSubgroupToGroupEventClass     _CNChangeHistoryAddSubgroupToGroupEventClass
	CNChangeHistoryAddSubgroupToGroupEventClassOnce sync.Once
)

func getCNChangeHistoryAddSubgroupToGroupEventClass() _CNChangeHistoryAddSubgroupToGroupEventClass {
	CNChangeHistoryAddSubgroupToGroupEventClassOnce.Do(func() {
		CNChangeHistoryAddSubgroupToGroupEventClass = _CNChangeHistoryAddSubgroupToGroupEventClass{objc.GetClass("CNChangeHistoryAddSubgroupToGroupEvent")}
	})
	return CNChangeHistoryAddSubgroupToGroupEventClass
}

type _CNChangeHistoryAddSubgroupToGroupEventClass struct {
	class objc.Class
}

// An interface definition for the [CNChangeHistoryAddSubgroupToGroupEvent] class.
type ICNChangeHistoryAddSubgroupToGroupEvent interface {
	ICNChangeHistoryEvent
}

// An object that represents a user adding a subgroup to a group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddSubgroupToGroupEvent
type CNChangeHistoryAddSubgroupToGroupEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryAddSubgroupToGroupEventFrom constructs a [CNChangeHistoryAddSubgroupToGroupEvent] from an unsafe.Pointer.
//
// An object that represents a user adding a subgroup to a group.
func CNChangeHistoryAddSubgroupToGroupEventFrom(ptr unsafe.Pointer) CNChangeHistoryAddSubgroupToGroupEvent {
	return CNChangeHistoryAddSubgroupToGroupEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryAddSubgroupToGroupEventClass) Alloc() CNChangeHistoryAddSubgroupToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddSubgroupToGroupEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNChangeHistoryAddSubgroupToGroupEventClass) New() CNChangeHistoryAddSubgroupToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddSubgroupToGroupEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryAddSubgroupToGroupEvent) Init() CNChangeHistoryAddSubgroupToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddSubgroupToGroupEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryAddSubgroupToGroupEvent) Autorelease() CNChangeHistoryAddSubgroupToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddSubgroupToGroupEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryAddSubgroupToGroupEvent creates a new CNChangeHistoryAddSubgroupToGroupEvent instance.
func NewCNChangeHistoryAddSubgroupToGroupEvent() CNChangeHistoryAddSubgroupToGroupEvent {
	return getCNChangeHistoryAddSubgroupToGroupEventClass().New()
}


// The group where the user added a subgroup.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddSubgroupToGroupEvent/group
func (c_ CNChangeHistoryAddSubgroupToGroupEvent) Group() CNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("group"))
	return rv
}

// The subgroup that the user added to the group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddSubgroupToGroupEvent/subgroup
func (c_ CNChangeHistoryAddSubgroupToGroupEvent) Subgroup() CNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("subgroup"))
	return rv
}



