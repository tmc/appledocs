// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CNChangeHistoryAddSubgroupToGroupEvent */


/* debug [class_header]: Header for CNChangeHistoryAddSubgroupToGroupEvent */
// The class instance for the [CNChangeHistoryAddSubgroupToGroupEvent] class.
var (
	CNChangeHistoryAddSubgroupToGroupEventClass     _CNChangeHistoryAddSubgroupToGroupEventClass
	CNChangeHistoryAddSubgroupToGroupEventClassOnce sync.Once
)

func getCNChangeHistoryAddSubgroupToGroupEventClass() _CNChangeHistoryAddSubgroupToGroupEventClass {
	CNChangeHistoryAddSubgroupToGroupEventClassOnce.Do(func() {
		CNChangeHistoryAddSubgroupToGroupEventClass = _CNChangeHistoryAddSubgroupToGroupEventClass{objc.GetClass("CNChangeHistoryAddSubgroupToGroupEvent")}
	})
	return CNChangeHistoryAddSubgroupToGroupEventClass
}

type _CNChangeHistoryAddSubgroupToGroupEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNChangeHistoryAddSubgroupToGroupEvent */
// An interface definition for the [CNChangeHistoryAddSubgroupToGroupEvent] class.
type ICNChangeHistoryAddSubgroupToGroupEvent interface {
	ICNChangeHistoryEvent
	
/* debug [class_interface_properties]: Properties for CNChangeHistoryAddSubgroupToGroupEvent */
	// properties:
	Group() ICNGroup
	Subgroup() ICNGroup
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNChangeHistoryAddSubgroupToGroupEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNChangeHistoryAddSubgroupToGroupEvent */
// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryAddSubgroupToGroupEventClass) Alloc() CNChangeHistoryAddSubgroupToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddSubgroupToGroupEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNChangeHistoryAddSubgroupToGroupEventClass) New() CNChangeHistoryAddSubgroupToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddSubgroupToGroupEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryAddSubgroupToGroupEvent) Init() CNChangeHistoryAddSubgroupToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddSubgroupToGroupEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryAddSubgroupToGroupEvent) Autorelease() CNChangeHistoryAddSubgroupToGroupEvent {
	rv := objc.Send[CNChangeHistoryAddSubgroupToGroupEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryAddSubgroupToGroupEvent creates a new CNChangeHistoryAddSubgroupToGroupEvent instance.
func NewCNChangeHistoryAddSubgroupToGroupEvent() CNChangeHistoryAddSubgroupToGroupEvent {
	return getCNChangeHistoryAddSubgroupToGroupEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNChangeHistoryAddSubgroupToGroupEvent */
// An object that represents a user adding a subgroup to a group.


// An object that represents a user adding a subgroup to a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddSubgroupToGroupEvent
type CNChangeHistoryAddSubgroupToGroupEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryAddSubgroupToGroupEventFrom constructs a [CNChangeHistoryAddSubgroupToGroupEvent] from an unsafe.Pointer.
//
// An object that represents a user adding a subgroup to a group.
func CNChangeHistoryAddSubgroupToGroupEventFrom(ptr unsafe.Pointer) CNChangeHistoryAddSubgroupToGroupEvent {
	return CNChangeHistoryAddSubgroupToGroupEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNChangeHistoryAddSubgroupToGroupEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNChangeHistoryAddSubgroupToGroupEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNChangeHistoryAddSubgroupToGroupEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNChangeHistoryAddSubgroupToGroupEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNChangeHistoryAddSubgroupToGroupEvent */

// The group where the user added a subgroup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddSubgroupToGroupEvent/group
func (c_ CNChangeHistoryAddSubgroupToGroupEvent) Group() ICNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("group"))
	return rv
}/* debug [instance_properties/getter]: group */


// The subgroup that the user added to the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddSubgroupToGroupEvent/subgroup
func (c_ CNChangeHistoryAddSubgroupToGroupEvent) Subgroup() ICNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("subgroup"))
	return rv
}/* debug [instance_properties/getter]: subgroup */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNChangeHistoryAddSubgroupToGroupEvent */



