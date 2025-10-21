// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNChangeHistoryAddContactEvent] class.
var (
	CNChangeHistoryAddContactEventClass     _CNChangeHistoryAddContactEventClass
	CNChangeHistoryAddContactEventClassOnce sync.Once
)

func getCNChangeHistoryAddContactEventClass() _CNChangeHistoryAddContactEventClass {
	CNChangeHistoryAddContactEventClassOnce.Do(func() {
		CNChangeHistoryAddContactEventClass = _CNChangeHistoryAddContactEventClass{objc.GetClass("CNChangeHistoryAddContactEvent")}
	})
	return CNChangeHistoryAddContactEventClass
}

type _CNChangeHistoryAddContactEventClass struct {
	class objc.Class
}

// An interface definition for the [CNChangeHistoryAddContactEvent] class.
type ICNChangeHistoryAddContactEvent interface {
	ICNChangeHistoryEvent
}

// An object that represents a user adding a contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddContactEvent
type CNChangeHistoryAddContactEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryAddContactEventFrom constructs a [CNChangeHistoryAddContactEvent] from an unsafe.Pointer.
//
// An object that represents a user adding a contact.
func CNChangeHistoryAddContactEventFrom(ptr unsafe.Pointer) CNChangeHistoryAddContactEvent {
	return CNChangeHistoryAddContactEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryAddContactEventClass) Alloc() CNChangeHistoryAddContactEvent {
	rv := objc.Send[CNChangeHistoryAddContactEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNChangeHistoryAddContactEventClass) New() CNChangeHistoryAddContactEvent {
	rv := objc.Send[CNChangeHistoryAddContactEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryAddContactEvent) Init() CNChangeHistoryAddContactEvent {
	rv := objc.Send[CNChangeHistoryAddContactEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryAddContactEvent) Autorelease() CNChangeHistoryAddContactEvent {
	rv := objc.Send[CNChangeHistoryAddContactEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryAddContactEvent creates a new CNChangeHistoryAddContactEvent instance.
func NewCNChangeHistoryAddContactEvent() CNChangeHistoryAddContactEvent {
	return getCNChangeHistoryAddContactEventClass().New()
}


// The contact the user added.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddContactEvent/contact
func (c_ CNChangeHistoryAddContactEvent) Contact() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("contact"))
	return rv
}

// A string that uniquely identifies the container where the user added the contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddContactEvent/containerIdentifier
func (c_ CNChangeHistoryAddContactEvent) ContainerIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}



