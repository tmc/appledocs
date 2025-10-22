// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNChangeHistoryAddGroupEvent] class.
var (
	CNChangeHistoryAddGroupEventClass     _CNChangeHistoryAddGroupEventClass
	CNChangeHistoryAddGroupEventClassOnce sync.Once
)

func getCNChangeHistoryAddGroupEventClass() _CNChangeHistoryAddGroupEventClass {
	CNChangeHistoryAddGroupEventClassOnce.Do(func() {
		CNChangeHistoryAddGroupEventClass = _CNChangeHistoryAddGroupEventClass{objc.GetClass("CNChangeHistoryAddGroupEvent")}
	})
	return CNChangeHistoryAddGroupEventClass
}

type _CNChangeHistoryAddGroupEventClass struct {
	class objc.Class
}

// An interface definition for the [CNChangeHistoryAddGroupEvent] class.
type ICNChangeHistoryAddGroupEvent interface {
	ICNChangeHistoryEvent
	ContainerIdentifier() string
	Group() CNGroup
}

// An object that represents a user adding a group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddGroupEvent
type CNChangeHistoryAddGroupEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryAddGroupEventFrom constructs a [CNChangeHistoryAddGroupEvent] from an unsafe.Pointer.
//
// An object that represents a user adding a group.
func CNChangeHistoryAddGroupEventFrom(ptr unsafe.Pointer) CNChangeHistoryAddGroupEvent {
	return CNChangeHistoryAddGroupEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryAddGroupEventClass) Alloc() CNChangeHistoryAddGroupEvent {
	rv := objc.Send[CNChangeHistoryAddGroupEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNChangeHistoryAddGroupEventClass) New() CNChangeHistoryAddGroupEvent {
	rv := objc.Send[CNChangeHistoryAddGroupEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryAddGroupEvent) Init() CNChangeHistoryAddGroupEvent {
	rv := objc.Send[CNChangeHistoryAddGroupEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryAddGroupEvent) Autorelease() CNChangeHistoryAddGroupEvent {
	rv := objc.Send[CNChangeHistoryAddGroupEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryAddGroupEvent creates a new CNChangeHistoryAddGroupEvent instance.
func NewCNChangeHistoryAddGroupEvent() CNChangeHistoryAddGroupEvent {
	return getCNChangeHistoryAddGroupEventClass().New()
}


// A string that uniquely identifies the container where the user added the group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddGroupEvent/containerIdentifier
func (c_ CNChangeHistoryAddGroupEvent) ContainerIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}

// The group the user added.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddGroupEvent/group
func (c_ CNChangeHistoryAddGroupEvent) Group() CNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("group"))
	return rv
}



