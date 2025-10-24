// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCustomDeviceRoute */


/* debug [class_header]: Header for AVCustomDeviceRoute */
// The class instance for the [CustomDeviceRoute] class.
var (
	CustomDeviceRouteClass     _CustomDeviceRouteClass
	CustomDeviceRouteClassOnce sync.Once
)

func getCustomDeviceRouteClass() _CustomDeviceRouteClass {
	CustomDeviceRouteClassOnce.Do(func() {
		CustomDeviceRouteClass = _CustomDeviceRouteClass{objc.GetClass("AVCustomDeviceRoute")}
	})
	return CustomDeviceRouteClass
}

type _CustomDeviceRouteClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CustomDeviceRoute */
// An interface definition for the [CustomDeviceRoute] class.
type ICustomDeviceRoute interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CustomDeviceRoute */
	// properties:
	Reason() CustomRoutingEventReason
	SetReason(value CustomRoutingEventReason)
	Route() IAVCustomDeviceRoute
	SetRoute(value IAVCustomDeviceRoute)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CustomDeviceRoute */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CustomDeviceRoute */
// Alloc allocates a new instance without initialization.
func (cc _CustomDeviceRouteClass) Alloc() CustomDeviceRoute {
	rv := objc.Send[CustomDeviceRoute](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CustomDeviceRouteClass) New() CustomDeviceRoute {
	rv := objc.Send[CustomDeviceRoute](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomDeviceRoute) Init() CustomDeviceRoute {
	rv := objc.Send[CustomDeviceRoute](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomDeviceRoute) Autorelease() CustomDeviceRoute {
	rv := objc.Send[CustomDeviceRoute](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomDeviceRoute creates a new CustomDeviceRoute instance.
func NewCustomDeviceRoute() CustomDeviceRoute {
	return getCustomDeviceRouteClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CustomDeviceRoute */
// An object that represents a custom device route.
//
// Use the value of a route’s or property to establish a connection to a device. Typically, only one of the properties provides a valid value, depending on the type of device. In certain cases, both properties may provide valid values, in which case your app determines which one to use.


// An object that represents a custom device route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomDeviceRoute
type CustomDeviceRoute struct {
	objectivec.Object
}

// CustomDeviceRouteFrom constructs a [CustomDeviceRoute] from an unsafe.Pointer.
//
// An object that represents a custom device route.
func CustomDeviceRouteFrom(ptr unsafe.Pointer) CustomDeviceRoute {
	return CustomDeviceRoute{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CustomDeviceRoute *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CustomDeviceRoute */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CustomDeviceRoute */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CustomDeviceRoute */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CustomDeviceRoute */

// A reason for an event, such as a user request to activate or deactivate a route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingevent/reason
func (c_ CustomDeviceRoute) Reason() CustomRoutingEventReason {
	rv := objc.Send[CustomRoutingEventReason](c_.ID, objc.Sel("reason"))
	return rv
}/* debug [instance_properties/getter]: reason */


// A reason for an event, such as a user request to activate or deactivate a route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingevent/reason
func (c_ CustomDeviceRoute) SetReason(value CustomRoutingEventReason) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReason:"), value)
}/* debug [instance_properties/setter]: reason */


// A route for the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingevent/route
func (c_ CustomDeviceRoute) Route() IAVCustomDeviceRoute {
	rv := objc.Send[CustomDeviceRoute](c_.ID, objc.Sel("route"))
	return rv
}/* debug [instance_properties/getter]: route */


// A route for the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingevent/route
func (c_ CustomDeviceRoute) SetRoute(value IAVCustomDeviceRoute) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRoute:"), value)
}/* debug [instance_properties/setter]: route */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCustomDeviceRoute */


