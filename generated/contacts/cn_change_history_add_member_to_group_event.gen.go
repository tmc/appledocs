// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CNChangeHistoryAddMemberToGroupEvent */


/* debug [class_header]: Header for CNChangeHistoryAddMemberToGroupEvent */
// The class instance for the [CNChangeHistoryAddMemberToGroupEvent] class.
var (
	CNChangeHistoryAddMemberToGroupEventClass     _CNChangeHistoryAddMemberToGroupEventClass
	CNChangeHistoryAddMemberToGroupEventClassOnce sync.Once
)

func getCNChangeHistoryAddMemberToGroupEventClass() _CNChangeHistoryAddMemberToGroupEventClass {
	CNChangeHistoryAddMemberToGroupEventClassOnce.Do(func() {
		CNChangeHistoryAddMemberToGroupEventClass = _CNChangeHistoryAddMemberToGroupEventClass{objc.GetClass("CNChangeHistoryAddMemberToGroupEvent")}
	})
	return CNChangeHistoryAddMemberToGroupEventClass
}

type _CNChangeHistoryAddMemberToGroupEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNChangeHistoryAddMemberToGroupEvent */
// An interface definition for the [CNChangeHistoryAddMemberToGroupEvent] class.
type ICNChangeHistoryAddMemberToGroupEvent interface {
	ICNChangeHistoryEvent
	
/* debug [class_interface_properties]: Properties for CNChangeHistoryAddMemberToGroupEvent */
	// properties:
	Group() ICNGroup
	Member() ICNContact
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNChangeHistoryAddMemberToGroupEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNChangeHistoryAddMemberToGroupEvent */
// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryAddMemberToGroupEventClass) Alloc() CNChangeHistoryAddMemberToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddMemberToGroupEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNChangeHistoryAddMemberToGroupEventClass) New() CNChangeHistoryAddMemberToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddMemberToGroupEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryAddMemberToGroupEvent) Init() CNChangeHistoryAddMemberToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddMemberToGroupEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryAddMemberToGroupEvent) Autorelease() CNChangeHistoryAddMemberToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddMemberToGroupEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryAddMemberToGroupEvent creates a new CNChangeHistoryAddMemberToGroupEvent instance.
func NewCNChangeHistoryAddMemberToGroupEvent() CNChangeHistoryAddMemberToGroupEvent {
	return getCNChangeHistoryAddMemberToGroupEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNChangeHistoryAddMemberToGroupEvent */
// An object that represents a user adding a contact to a group.


// An object that represents a user adding a contact to a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddMemberToGroupEvent
type CNChangeHistoryAddMemberToGroupEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryAddMemberToGroupEventFrom constructs a [CNChangeHistoryAddMemberToGroupEvent] from an unsafe.Pointer.
//
// An object that represents a user adding a contact to a group.
func CNChangeHistoryAddMemberToGroupEventFrom(ptr unsafe.Pointer) CNChangeHistoryAddMemberToGroupEvent {
	return CNChangeHistoryAddMemberToGroupEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNChangeHistoryAddMemberToGroupEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNChangeHistoryAddMemberToGroupEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNChangeHistoryAddMemberToGroupEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNChangeHistoryAddMemberToGroupEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNChangeHistoryAddMemberToGroupEvent */

// The group where the user added a contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddMemberToGroupEvent/group
func (c_ CNChangeHistoryAddMemberToGroupEvent) Group() ICNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("group"))
	return rv
}/* debug [instance_properties/getter]: group */


// The contact the user added to the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddMemberToGroupEvent/member
func (c_ CNChangeHistoryAddMemberToGroupEvent) Member() ICNContact {
	rv := objc.Send[CNContact](c_.ID, objc.Sel("member"))
	return rv
}/* debug [instance_properties/getter]: member */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNChangeHistoryAddMemberToGroupEvent */



