// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CNChangeHistoryRemoveMemberFromGroupEvent */


/* debug [class_header]: Header for CNChangeHistoryRemoveMemberFromGroupEvent */
// The class instance for the [CNChangeHistoryRemoveMemberFromGroupEvent] class.
var (
	CNChangeHistoryRemoveMemberFromGroupEventClass     _CNChangeHistoryRemoveMemberFromGroupEventClass
	CNChangeHistoryRemoveMemberFromGroupEventClassOnce sync.Once
)

func getCNChangeHistoryRemoveMemberFromGroupEventClass() _CNChangeHistoryRemoveMemberFromGroupEventClass {
	CNChangeHistoryRemoveMemberFromGroupEventClassOnce.Do(func() {
		CNChangeHistoryRemoveMemberFromGroupEventClass = _CNChangeHistoryRemoveMemberFromGroupEventClass{objc.GetClass("CNChangeHistoryRemoveMemberFromGroupEvent")}
	})
	return CNChangeHistoryRemoveMemberFromGroupEventClass
}

type _CNChangeHistoryRemoveMemberFromGroupEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNChangeHistoryRemoveMemberFromGroupEvent */
// An interface definition for the [CNChangeHistoryRemoveMemberFromGroupEvent] class.
type ICNChangeHistoryRemoveMemberFromGroupEvent interface {
	ICNChangeHistoryEvent
	
/* debug [class_interface_properties]: Properties for CNChangeHistoryRemoveMemberFromGroupEvent */
	// properties:
	Group() ICNGroup
	Member() ICNContact
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNChangeHistoryRemoveMemberFromGroupEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNChangeHistoryRemoveMemberFromGroupEvent */
// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryRemoveMemberFromGroupEventClass) Alloc() CNChangeHistoryRemoveMemberFromGroupEvent {
	rv := objc.Send[CNChangeHistoryRemoveMemberFromGroupEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNChangeHistoryRemoveMemberFromGroupEventClass) New() CNChangeHistoryRemoveMemberFromGroupEvent {
	rv := objc.Send[CNChangeHistoryRemoveMemberFromGroupEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryRemoveMemberFromGroupEvent) Init() CNChangeHistoryRemoveMemberFromGroupEvent {
	rv := objc.Send[CNChangeHistoryRemoveMemberFromGroupEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryRemoveMemberFromGroupEvent) Autorelease() CNChangeHistoryRemoveMemberFromGroupEvent {
	rv := objc.Send[CNChangeHistoryRemoveMemberFromGroupEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryRemoveMemberFromGroupEvent creates a new CNChangeHistoryRemoveMemberFromGroupEvent instance.
func NewCNChangeHistoryRemoveMemberFromGroupEvent() CNChangeHistoryRemoveMemberFromGroupEvent {
	return getCNChangeHistoryRemoveMemberFromGroupEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNChangeHistoryRemoveMemberFromGroupEvent */
// An object that represents a user removing a contact from a group.


// An object that represents a user removing a contact from a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryRemoveMemberFromGroupEvent
type CNChangeHistoryRemoveMemberFromGroupEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryRemoveMemberFromGroupEventFrom constructs a [CNChangeHistoryRemoveMemberFromGroupEvent] from an unsafe.Pointer.
//
// An object that represents a user removing a contact from a group.
func CNChangeHistoryRemoveMemberFromGroupEventFrom(ptr unsafe.Pointer) CNChangeHistoryRemoveMemberFromGroupEvent {
	return CNChangeHistoryRemoveMemberFromGroupEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNChangeHistoryRemoveMemberFromGroupEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNChangeHistoryRemoveMemberFromGroupEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNChangeHistoryRemoveMemberFromGroupEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNChangeHistoryRemoveMemberFromGroupEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNChangeHistoryRemoveMemberFromGroupEvent */

// The group where the user removed a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryRemoveMemberFromGroupEvent/group
func (c_ CNChangeHistoryRemoveMemberFromGroupEvent) Group() ICNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("group"))
	return rv
}/* debug [instance_properties/getter]: group */


// The contact that the user removed from the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryRemoveMemberFromGroupEvent/member
func (c_ CNChangeHistoryRemoveMemberFromGroupEvent) Member() ICNContact {
	rv := objc.Send[CNContact](c_.ID, objc.Sel("member"))
	return rv
}/* debug [instance_properties/getter]: member */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNChangeHistoryRemoveMemberFromGroupEvent */



