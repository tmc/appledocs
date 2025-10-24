// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CNChangeHistoryUpdateContactEvent */


/* debug [class_header]: Header for CNChangeHistoryUpdateContactEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNChangeHistoryUpdateContactEvent */
// An interface definition for the [CNChangeHistoryUpdateContactEvent] class.
type ICNChangeHistoryUpdateContactEvent interface {
	ICNChangeHistoryEvent
	
/* debug [class_interface_properties]: Properties for CNChangeHistoryUpdateContactEvent */
	// properties:
	Contact() ICNContact
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNChangeHistoryUpdateContactEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNChangeHistoryUpdateContactEvent */
// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryUpdateContactEventClass) Alloc() CNChangeHistoryUpdateContactEvent {
	rv := objc.Send[CNChangeHistoryUpdateContactEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNChangeHistoryUpdateContactEvent */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNChangeHistoryUpdateContactEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNChangeHistoryUpdateContactEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNChangeHistoryUpdateContactEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNChangeHistoryUpdateContactEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNChangeHistoryUpdateContactEvent */

// The contact that the user updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryUpdateContactEvent/contact
func (c_ CNChangeHistoryUpdateContactEvent) Contact() ICNContact {
	rv := objc.Send[CNContact](c_.ID, objc.Sel("contact"))
	return rv
}/* debug [instance_properties/getter]: contact */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNChangeHistoryUpdateContactEvent */



