// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CNChangeHistoryAddGroupEvent */


/* debug [class_header]: Header for CNChangeHistoryAddGroupEvent */
// The class instance for the [CNChangeHistoryAddGroupEvent] class.
var (
	CNChangeHistoryAddGroupEventClass     _CNChangeHistoryAddGroupEventClass
	CNChangeHistoryAddGroupEventClassOnce sync.Once
)

func getCNChangeHistoryAddGroupEventClass() _CNChangeHistoryAddGroupEventClass {
	CNChangeHistoryAddGroupEventClassOnce.Do(func() {
		CNChangeHistoryAddGroupEventClass = _CNChangeHistoryAddGroupEventClass{objc.GetClass("CNChangeHistoryAddGroupEvent")}
	})
	return CNChangeHistoryAddGroupEventClass
}

type _CNChangeHistoryAddGroupEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNChangeHistoryAddGroupEvent */
// An interface definition for the [CNChangeHistoryAddGroupEvent] class.
type ICNChangeHistoryAddGroupEvent interface {
	ICNChangeHistoryEvent
	
/* debug [class_interface_properties]: Properties for CNChangeHistoryAddGroupEvent */
	// properties:
	ContainerIdentifier() objc.IObject /* cross-framework: NSString */
	Group() ICNGroup
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNChangeHistoryAddGroupEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNChangeHistoryAddGroupEvent */
// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryAddGroupEventClass) Alloc() CNChangeHistoryAddGroupEvent {
	rv := objc.Send[CNChangeHistoryAddGroupEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNChangeHistoryAddGroupEventClass) New() CNChangeHistoryAddGroupEvent {
	rv := objc.Send[CNChangeHistoryAddGroupEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryAddGroupEvent) Init() CNChangeHistoryAddGroupEvent {
	rv := objc.Send[CNChangeHistoryAddGroupEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryAddGroupEvent) Autorelease() CNChangeHistoryAddGroupEvent {
	rv := objc.Send[CNChangeHistoryAddGroupEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryAddGroupEvent creates a new CNChangeHistoryAddGroupEvent instance.
func NewCNChangeHistoryAddGroupEvent() CNChangeHistoryAddGroupEvent {
	return getCNChangeHistoryAddGroupEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNChangeHistoryAddGroupEvent */
// An object that represents a user adding a group.


// An object that represents a user adding a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddGroupEvent
type CNChangeHistoryAddGroupEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryAddGroupEventFrom constructs a [CNChangeHistoryAddGroupEvent] from an unsafe.Pointer.
//
// An object that represents a user adding a group.
func CNChangeHistoryAddGroupEventFrom(ptr unsafe.Pointer) CNChangeHistoryAddGroupEvent {
	return CNChangeHistoryAddGroupEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNChangeHistoryAddGroupEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNChangeHistoryAddGroupEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNChangeHistoryAddGroupEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNChangeHistoryAddGroupEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNChangeHistoryAddGroupEvent */

// A string that uniquely identifies the container where the user added the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddGroupEvent/containerIdentifier
func (c_ CNChangeHistoryAddGroupEvent) ContainerIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}/* debug [instance_properties/getter]: containerIdentifier */


// The group the user added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryAddGroupEvent/group
func (c_ CNChangeHistoryAddGroupEvent) Group() ICNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("group"))
	return rv
}/* debug [instance_properties/getter]: group */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNChangeHistoryAddGroupEvent */



