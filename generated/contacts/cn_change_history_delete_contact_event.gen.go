// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNChangeHistoryDeleteContactEvent] class.
var (
	CNChangeHistoryDeleteContactEventClass     _CNChangeHistoryDeleteContactEventClass
	CNChangeHistoryDeleteContactEventClassOnce sync.Once
)

func getCNChangeHistoryDeleteContactEventClass() _CNChangeHistoryDeleteContactEventClass {
	CNChangeHistoryDeleteContactEventClassOnce.Do(func() {
		CNChangeHistoryDeleteContactEventClass = _CNChangeHistoryDeleteContactEventClass{objc.GetClass("CNChangeHistoryDeleteContactEvent")}
	})
	return CNChangeHistoryDeleteContactEventClass
}

type _CNChangeHistoryDeleteContactEventClass struct {
	class objc.Class
}

// An interface definition for the [CNChangeHistoryDeleteContactEvent] class.
type ICNChangeHistoryDeleteContactEvent interface {
	ICNChangeHistoryEvent
	ContactIdentifier() string
}

// An object that represents a user deleting a contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryDeleteContactEvent
type CNChangeHistoryDeleteContactEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryDeleteContactEventFrom constructs a [CNChangeHistoryDeleteContactEvent] from an unsafe.Pointer.
//
// An object that represents a user deleting a contact.
func CNChangeHistoryDeleteContactEventFrom(ptr unsafe.Pointer) CNChangeHistoryDeleteContactEvent {
	return CNChangeHistoryDeleteContactEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryDeleteContactEventClass) Alloc() CNChangeHistoryDeleteContactEvent {
	rv := objc.Send[CNChangeHistoryDeleteContactEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNChangeHistoryDeleteContactEventClass) New() CNChangeHistoryDeleteContactEvent {
	rv := objc.Send[CNChangeHistoryDeleteContactEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryDeleteContactEvent) Init() CNChangeHistoryDeleteContactEvent {
	rv := objc.Send[CNChangeHistoryDeleteContactEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryDeleteContactEvent) Autorelease() CNChangeHistoryDeleteContactEvent {
	rv := objc.Send[CNChangeHistoryDeleteContactEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryDeleteContactEvent creates a new CNChangeHistoryDeleteContactEvent instance.
func NewCNChangeHistoryDeleteContactEvent() CNChangeHistoryDeleteContactEvent {
	return getCNChangeHistoryDeleteContactEventClass().New()
}


// A string that uniquely identifies the contact that the user deleted.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryDeleteContactEvent/contactIdentifier
func (c_ CNChangeHistoryDeleteContactEvent) ContactIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("contactIdentifier"))
	return rv
}



