// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCustomRoutingEvent */


/* debug [class_header]: Header for AVCustomRoutingEvent */
// The class instance for the [CustomRoutingEvent] class.
var (
	CustomRoutingEventClass     _CustomRoutingEventClass
	CustomRoutingEventClassOnce sync.Once
)

func getCustomRoutingEventClass() _CustomRoutingEventClass {
	CustomRoutingEventClassOnce.Do(func() {
		CustomRoutingEventClass = _CustomRoutingEventClass{objc.GetClass("AVCustomRoutingEvent")}
	})
	return CustomRoutingEventClass
}

type _CustomRoutingEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CustomRoutingEvent */
// An interface definition for the [CustomRoutingEvent] class.
type ICustomRoutingEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CustomRoutingEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CustomRoutingEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CustomRoutingEvent */
// Alloc allocates a new instance without initialization.
func (cc _CustomRoutingEventClass) Alloc() CustomRoutingEvent {
	rv := objc.Send[CustomRoutingEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CustomRoutingEventClass) New() CustomRoutingEvent {
	rv := objc.Send[CustomRoutingEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomRoutingEvent) Init() CustomRoutingEvent {
	rv := objc.Send[CustomRoutingEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomRoutingEvent) Autorelease() CustomRoutingEvent {
	rv := objc.Send[CustomRoutingEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomRoutingEvent creates a new CustomRoutingEvent instance.
func NewCustomRoutingEvent() CustomRoutingEvent {
	return getCustomRoutingEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CustomRoutingEvent */
// An object that represents an event that occurs on a route.
//
// Depending on the route’s reason, apps establish or tear down a connection to a specified route.


// An object that represents an event that occurs on a route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingEvent
type CustomRoutingEvent struct {
	objectivec.Object
}

// CustomRoutingEventFrom constructs a [CustomRoutingEvent] from an unsafe.Pointer.
//
// An object that represents an event that occurs on a route.
func CustomRoutingEventFrom(ptr unsafe.Pointer) CustomRoutingEvent {
	return CustomRoutingEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CustomRoutingEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CustomRoutingEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CustomRoutingEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CustomRoutingEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CustomRoutingEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCustomRoutingEvent */


