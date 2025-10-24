// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CNChangeHistoryDeleteContactEvent */


/* debug [class_header]: Header for CNChangeHistoryDeleteContactEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNChangeHistoryDeleteContactEvent */
// An interface definition for the [CNChangeHistoryDeleteContactEvent] class.
type ICNChangeHistoryDeleteContactEvent interface {
	ICNChangeHistoryEvent
	
/* debug [class_interface_properties]: Properties for CNChangeHistoryDeleteContactEvent */
	// properties:
	ContactIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNChangeHistoryDeleteContactEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNChangeHistoryDeleteContactEvent */
// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryDeleteContactEventClass) Alloc() CNChangeHistoryDeleteContactEvent {
	rv := objc.Send[CNChangeHistoryDeleteContactEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNChangeHistoryDeleteContactEvent */
// An object that represents a user deleting a contact.


// An object that represents a user deleting a contact.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNChangeHistoryDeleteContactEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNChangeHistoryDeleteContactEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNChangeHistoryDeleteContactEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNChangeHistoryDeleteContactEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNChangeHistoryDeleteContactEvent */

// A string that uniquely identifies the contact that the user deleted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryDeleteContactEvent/contactIdentifier
func (c_ CNChangeHistoryDeleteContactEvent) ContactIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("contactIdentifier"))
	return rv
}/* debug [instance_properties/getter]: contactIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNChangeHistoryDeleteContactEvent */



