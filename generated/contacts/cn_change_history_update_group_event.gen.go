// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CNChangeHistoryUpdateGroupEvent */


/* debug [class_header]: Header for CNChangeHistoryUpdateGroupEvent */
// The class instance for the [CNChangeHistoryUpdateGroupEvent] class.
var (
	CNChangeHistoryUpdateGroupEventClass     _CNChangeHistoryUpdateGroupEventClass
	CNChangeHistoryUpdateGroupEventClassOnce sync.Once
)

func getCNChangeHistoryUpdateGroupEventClass() _CNChangeHistoryUpdateGroupEventClass {
	CNChangeHistoryUpdateGroupEventClassOnce.Do(func() {
		CNChangeHistoryUpdateGroupEventClass = _CNChangeHistoryUpdateGroupEventClass{objc.GetClass("CNChangeHistoryUpdateGroupEvent")}
	})
	return CNChangeHistoryUpdateGroupEventClass
}

type _CNChangeHistoryUpdateGroupEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNChangeHistoryUpdateGroupEvent */
// An interface definition for the [CNChangeHistoryUpdateGroupEvent] class.
type ICNChangeHistoryUpdateGroupEvent interface {
	ICNChangeHistoryEvent
	
/* debug [class_interface_properties]: Properties for CNChangeHistoryUpdateGroupEvent */
	// properties:
	Group() ICNGroup
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNChangeHistoryUpdateGroupEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNChangeHistoryUpdateGroupEvent */
// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryUpdateGroupEventClass) Alloc() CNChangeHistoryUpdateGroupEvent {
	rv := objc.Send[CNChangeHistoryUpdateGroupEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNChangeHistoryUpdateGroupEventClass) New() CNChangeHistoryUpdateGroupEvent {
	rv := objc.Send[CNChangeHistoryUpdateGroupEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryUpdateGroupEvent) Init() CNChangeHistoryUpdateGroupEvent {
	rv := objc.Send[CNChangeHistoryUpdateGroupEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryUpdateGroupEvent) Autorelease() CNChangeHistoryUpdateGroupEvent {
	rv := objc.Send[CNChangeHistoryUpdateGroupEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryUpdateGroupEvent creates a new CNChangeHistoryUpdateGroupEvent instance.
func NewCNChangeHistoryUpdateGroupEvent() CNChangeHistoryUpdateGroupEvent {
	return getCNChangeHistoryUpdateGroupEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNChangeHistoryUpdateGroupEvent */
// An object that represents an updated group event.


// An object that represents an updated group event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryUpdateGroupEvent
type CNChangeHistoryUpdateGroupEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryUpdateGroupEventFrom constructs a [CNChangeHistoryUpdateGroupEvent] from an unsafe.Pointer.
//
// An object that represents an updated group event.
func CNChangeHistoryUpdateGroupEventFrom(ptr unsafe.Pointer) CNChangeHistoryUpdateGroupEvent {
	return CNChangeHistoryUpdateGroupEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNChangeHistoryUpdateGroupEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNChangeHistoryUpdateGroupEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNChangeHistoryUpdateGroupEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNChangeHistoryUpdateGroupEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNChangeHistoryUpdateGroupEvent */

// The group that the user updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryUpdateGroupEvent/group
func (c_ CNChangeHistoryUpdateGroupEvent) Group() ICNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("group"))
	return rv
}/* debug [instance_properties/getter]: group */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNChangeHistoryUpdateGroupEvent */



