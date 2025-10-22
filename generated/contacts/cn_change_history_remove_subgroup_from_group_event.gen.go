// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNChangeHistoryRemoveSubgroupFromGroupEvent] class.
var (
	CNChangeHistoryRemoveSubgroupFromGroupEventClass     _CNChangeHistoryRemoveSubgroupFromGroupEventClass
	CNChangeHistoryRemoveSubgroupFromGroupEventClassOnce sync.Once
)

func getCNChangeHistoryRemoveSubgroupFromGroupEventClass() _CNChangeHistoryRemoveSubgroupFromGroupEventClass {
	CNChangeHistoryRemoveSubgroupFromGroupEventClassOnce.Do(func() {
		CNChangeHistoryRemoveSubgroupFromGroupEventClass = _CNChangeHistoryRemoveSubgroupFromGroupEventClass{objc.GetClass("CNChangeHistoryRemoveSubgroupFromGroupEvent")}
	})
	return CNChangeHistoryRemoveSubgroupFromGroupEventClass
}

type _CNChangeHistoryRemoveSubgroupFromGroupEventClass struct {
	class objc.Class
}

// An interface definition for the [CNChangeHistoryRemoveSubgroupFromGroupEvent] class.
type ICNChangeHistoryRemoveSubgroupFromGroupEvent interface {
	ICNChangeHistoryEvent
	Group() CNGroup
	Subgroup() CNGroup
}

// An object that represents a user removing a subgroup from a group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryRemoveSubgroupFromGroupEvent
type CNChangeHistoryRemoveSubgroupFromGroupEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryRemoveSubgroupFromGroupEventFrom constructs a [CNChangeHistoryRemoveSubgroupFromGroupEvent] from an unsafe.Pointer.
//
// An object that represents a user removing a subgroup from a group.
func CNChangeHistoryRemoveSubgroupFromGroupEventFrom(ptr unsafe.Pointer) CNChangeHistoryRemoveSubgroupFromGroupEvent {
	return CNChangeHistoryRemoveSubgroupFromGroupEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryRemoveSubgroupFromGroupEventClass) Alloc() CNChangeHistoryRemoveSubgroupFromGroupEvent {
	rv := objc.Send[CNChangeHistoryRemoveSubgroupFromGroupEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNChangeHistoryRemoveSubgroupFromGroupEventClass) New() CNChangeHistoryRemoveSubgroupFromGroupEvent {
	rv := objc.Send[CNChangeHistoryRemoveSubgroupFromGroupEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryRemoveSubgroupFromGroupEvent) Init() CNChangeHistoryRemoveSubgroupFromGroupEvent {
	rv := objc.Send[CNChangeHistoryRemoveSubgroupFromGroupEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryRemoveSubgroupFromGroupEvent) Autorelease() CNChangeHistoryRemoveSubgroupFromGroupEvent {
	rv := objc.Send[CNChangeHistoryRemoveSubgroupFromGroupEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryRemoveSubgroupFromGroupEvent creates a new CNChangeHistoryRemoveSubgroupFromGroupEvent instance.
func NewCNChangeHistoryRemoveSubgroupFromGroupEvent() CNChangeHistoryRemoveSubgroupFromGroupEvent {
	return getCNChangeHistoryRemoveSubgroupFromGroupEventClass().New()
}


// The group where the user removed a subgroup.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryRemoveSubgroupFromGroupEvent/group
func (c_ CNChangeHistoryRemoveSubgroupFromGroupEvent) Group() CNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("group"))
	return rv
}

// The subgroup that the user removed from the group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryRemoveSubgroupFromGroupEvent/subgroup
func (c_ CNChangeHistoryRemoveSubgroupFromGroupEvent) Subgroup() CNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("subgroup"))
	return rv
}



