// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CNChangeHistoryAddContactEvent */


/* debug [class_header]: Header for CNChangeHistoryAddContactEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNChangeHistoryAddContactEvent */
// An interface definition for the [CNChangeHistoryAddContactEvent] class.
type ICNChangeHistoryAddContactEvent interface {
	ICNChangeHistoryEvent
	
/* debug [class_interface_properties]: Properties for CNChangeHistoryAddContactEvent */
	// properties:
	Contact() ICNContact
	ContainerIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNChangeHistoryAddContactEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNChangeHistoryAddContactEvent */
// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryAddContactEventClass) Alloc() CNChangeHistoryAddContactEvent {
	rv := objc.Send[CNChangeHistoryAddContactEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNChangeHistoryAddContactEvent */
// An object that represents a user adding a contact.


// An object that represents a user adding a contact.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNChangeHistoryAddContactEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNChangeHistoryAddContactEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNChangeHistoryAddContactEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNChangeHistoryAddContactEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNChangeHistoryAddContactEvent */

// The contact the user added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddContactEvent/contact
func (c_ CNChangeHistoryAddContactEvent) Contact() ICNContact {
	rv := objc.Send[CNContact](c_.ID, objc.Sel("contact"))
	return rv
}/* debug [instance_properties/getter]: contact */


// A string that uniquely identifies the container where the user added the contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddContactEvent/containerIdentifier
func (c_ CNChangeHistoryAddContactEvent) ContainerIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}/* debug [instance_properties/getter]: containerIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNChangeHistoryAddContactEvent */



