// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CNChangeHistoryDeleteGroupEvent */


/* debug [class_header]: Header for CNChangeHistoryDeleteGroupEvent */
// The class instance for the [CNChangeHistoryDeleteGroupEvent] class.
var (
	CNChangeHistoryDeleteGroupEventClass     _CNChangeHistoryDeleteGroupEventClass
	CNChangeHistoryDeleteGroupEventClassOnce sync.Once
)

func getCNChangeHistoryDeleteGroupEventClass() _CNChangeHistoryDeleteGroupEventClass {
	CNChangeHistoryDeleteGroupEventClassOnce.Do(func() {
		CNChangeHistoryDeleteGroupEventClass = _CNChangeHistoryDeleteGroupEventClass{objc.GetClass("CNChangeHistoryDeleteGroupEvent")}
	})
	return CNChangeHistoryDeleteGroupEventClass
}

type _CNChangeHistoryDeleteGroupEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNChangeHistoryDeleteGroupEvent */
// An interface definition for the [CNChangeHistoryDeleteGroupEvent] class.
type ICNChangeHistoryDeleteGroupEvent interface {
	ICNChangeHistoryEvent
	
/* debug [class_interface_properties]: Properties for CNChangeHistoryDeleteGroupEvent */
	// properties:
	GroupIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNChangeHistoryDeleteGroupEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNChangeHistoryDeleteGroupEvent */
// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryDeleteGroupEventClass) Alloc() CNChangeHistoryDeleteGroupEvent {
	rv := objc.Send[CNChangeHistoryDeleteGroupEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNChangeHistoryDeleteGroupEventClass) New() CNChangeHistoryDeleteGroupEvent {
	rv := objc.Send[CNChangeHistoryDeleteGroupEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryDeleteGroupEvent) Init() CNChangeHistoryDeleteGroupEvent {
	rv := objc.Send[CNChangeHistoryDeleteGroupEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryDeleteGroupEvent) Autorelease() CNChangeHistoryDeleteGroupEvent {
	rv := objc.Send[CNChangeHistoryDeleteGroupEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryDeleteGroupEvent creates a new CNChangeHistoryDeleteGroupEvent instance.
func NewCNChangeHistoryDeleteGroupEvent() CNChangeHistoryDeleteGroupEvent {
	return getCNChangeHistoryDeleteGroupEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNChangeHistoryDeleteGroupEvent */
// An object that represents a user deleting a group.


// An object that represents a user deleting a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryDeleteGroupEvent
type CNChangeHistoryDeleteGroupEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryDeleteGroupEventFrom constructs a [CNChangeHistoryDeleteGroupEvent] from an unsafe.Pointer.
//
// An object that represents a user deleting a group.
func CNChangeHistoryDeleteGroupEventFrom(ptr unsafe.Pointer) CNChangeHistoryDeleteGroupEvent {
	return CNChangeHistoryDeleteGroupEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNChangeHistoryDeleteGroupEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNChangeHistoryDeleteGroupEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNChangeHistoryDeleteGroupEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNChangeHistoryDeleteGroupEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNChangeHistoryDeleteGroupEvent */

// A string that uniquely identifies the group that the user deleted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryDeleteGroupEvent/groupIdentifier
func (c_ CNChangeHistoryDeleteGroupEvent) GroupIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("groupIdentifier"))
	return rv
}/* debug [instance_properties/getter]: groupIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNChangeHistoryDeleteGroupEvent */



