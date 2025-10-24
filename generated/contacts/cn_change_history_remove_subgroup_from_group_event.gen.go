// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CNChangeHistoryRemoveSubgroupFromGroupEvent */


/* debug [class_header]: Header for CNChangeHistoryRemoveSubgroupFromGroupEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNChangeHistoryRemoveSubgroupFromGroupEvent */
// An interface definition for the [CNChangeHistoryRemoveSubgroupFromGroupEvent] class.
type ICNChangeHistoryRemoveSubgroupFromGroupEvent interface {
	ICNChangeHistoryEvent
	
/* debug [class_interface_properties]: Properties for CNChangeHistoryRemoveSubgroupFromGroupEvent */
	// properties:
	Group() ICNGroup
	Subgroup() ICNGroup
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNChangeHistoryRemoveSubgroupFromGroupEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNChangeHistoryRemoveSubgroupFromGroupEvent */
// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryRemoveSubgroupFromGroupEventClass) Alloc() CNChangeHistoryRemoveSubgroupFromGroupEvent {
	rv := objc.Send[CNChangeHistoryRemoveSubgroupFromGroupEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNChangeHistoryRemoveSubgroupFromGroupEvent */
// An object that represents a user removing a subgroup from a group.


// An object that represents a user removing a subgroup from a group.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNChangeHistoryRemoveSubgroupFromGroupEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNChangeHistoryRemoveSubgroupFromGroupEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNChangeHistoryRemoveSubgroupFromGroupEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNChangeHistoryRemoveSubgroupFromGroupEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNChangeHistoryRemoveSubgroupFromGroupEvent */

// The group where the user removed a subgroup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryRemoveSubgroupFromGroupEvent/group
func (c_ CNChangeHistoryRemoveSubgroupFromGroupEvent) Group() ICNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("group"))
	return rv
}/* debug [instance_properties/getter]: group */


// The subgroup that the user removed from the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryRemoveSubgroupFromGroupEvent/subgroup
func (c_ CNChangeHistoryRemoveSubgroupFromGroupEvent) Subgroup() ICNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("subgroup"))
	return rv
}/* debug [instance_properties/getter]: subgroup */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNChangeHistoryRemoveSubgroupFromGroupEvent */



