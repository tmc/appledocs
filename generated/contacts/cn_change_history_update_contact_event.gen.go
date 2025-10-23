// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNChangeHistoryUpdateContactEvent] class.
var (
	CNChangeHistoryUpdateContactEventClass     _CNChangeHistoryUpdateContactEventClass
	CNChangeHistoryUpdateContactEventClassOnce sync.Once
)

func getCNChangeHistoryUpdateContactEventClass() _CNChangeHistoryUpdateContactEventClass {
	CNChangeHistoryUpdateContactEventClassOnce.Do(func() {
		CNChangeHistoryUpdateContactEventClass = _CNChangeHistoryUpdateContactEventClass{objc.GetClass("CNChangeHistoryUpdateContactEvent")}
	})
	return CNChangeHistoryUpdateContactEventClass
}

type _CNChangeHistoryUpdateContactEventClass struct {
	class objc.Class
}

// An interface definition for the [CNChangeHistoryUpdateContactEvent] class.
type ICNChangeHistoryUpdateContactEvent interface {
	ICNChangeHistoryEvent
	// properties:
	Contact() ICNContact
	// methods:
}

// An object that represents a user updating a contact.


// An object that represents a user updating a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryUpdateContactEvent
type CNChangeHistoryUpdateContactEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryUpdateContactEventFrom constructs a [CNChangeHistoryUpdateContactEvent] from an unsafe.Pointer.
//
// An object that represents a user updating a contact.
func CNChangeHistoryUpdateContactEventFrom(ptr unsafe.Pointer) CNChangeHistoryUpdateContactEvent {
	return CNChangeHistoryUpdateContactEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryUpdateContactEventClass) Alloc() CNChangeHistoryUpdateContactEvent {
	rv := objc.Send[CNChangeHistoryUpdateContactEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNChangeHistoryUpdateContactEventClass) New() CNChangeHistoryUpdateContactEvent {
	rv := objc.Send[CNChangeHistoryUpdateContactEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryUpdateContactEvent) Init() CNChangeHistoryUpdateContactEvent {
	rv := objc.Send[CNChangeHistoryUpdateContactEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryUpdateContactEvent) Autorelease() CNChangeHistoryUpdateContactEvent {
	rv := objc.Send[CNChangeHistoryUpdateContactEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryUpdateContactEvent creates a new CNChangeHistoryUpdateContactEvent instance.
func NewCNChangeHistoryUpdateContactEvent() CNChangeHistoryUpdateContactEvent {
	return getCNChangeHistoryUpdateContactEventClass().New()
}



// The contact that the user updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryUpdateContactEvent/contact
func (c_ CNChangeHistoryUpdateContactEvent) Contact() ICNContact {
	rv := objc.Send[CNContact](c_.ID, objc.Sel("contact"))
	return rv
}



