// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEIPv4Route */


/* debug [class_header]: Header for NEIPv4Route */
// The class instance for the [NEIPv4Route] class.
var (
	NEIPv4RouteClass     _NEIPv4RouteClass
	NEIPv4RouteClassOnce sync.Once
)

func getNEIPv4RouteClass() _NEIPv4RouteClass {
	NEIPv4RouteClassOnce.Do(func() {
		NEIPv4RouteClass = _NEIPv4RouteClass{objc.GetClass("NEIPv4Route")}
	})
	return NEIPv4RouteClass
}

type _NEIPv4RouteClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEIPv4Route */
// An interface definition for the [NEIPv4Route] class.
type INEIPv4Route interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEIPv4Route */
	// properties:
	DestinationAddress() objc.IObject /* cross-framework: NSString */
	DestinationSubnetMask() objc.IObject /* cross-framework: NSString */
	GatewayAddress() objc.IObject /* cross-framework: NSString */
	SetGatewayAddress(value objc.IObject /* cross-framework: NSString */)
	ExcludedRoutes() INEIPv4Route
	SetExcludedRoutes(value INEIPv4Route)
	IncludedRoutes() INEIPv4Route
	SetIncludedRoutes(value INEIPv4Route)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEIPv4Route */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEIPv4Route */
// Alloc allocates a new instance without initialization.
func (nc _NEIPv4RouteClass) Alloc() NEIPv4Route {
	rv := objc.Send[NEIPv4Route](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEIPv4RouteClass) New() NEIPv4Route {
	rv := objc.Send[NEIPv4Route](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEIPv4Route) Init() NEIPv4Route {
	rv := objc.Send[NEIPv4Route](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEIPv4Route) Autorelease() NEIPv4Route {
	rv := objc.Send[NEIPv4Route](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEIPv4Route creates a new NEIPv4Route instance.
func NewNEIPv4Route() NEIPv4Route {
	return getNEIPv4RouteClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEIPv4Route */
// The settings for an IPv4 route.


// The settings for an IPv4 route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route
type NEIPv4Route struct {
	objectivec.Object
}

// NEIPv4RouteFrom constructs a [NEIPv4Route] from an unsafe.Pointer.
//
// The settings for an IPv4 route.
func NEIPv4RouteFrom(ptr unsafe.Pointer) NEIPv4Route {
	return NEIPv4Route{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEIPv4Route */

// Initialize the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route/init(destinationAddress:subnetMask:)
func NewNEIPv4RouteWithDestinationAddressSubnetMask(address objc.IObject /* cross-framework: NSString */, subnetMask objc.IObject /* cross-framework: NSString */) NEIPv4Route {
	instance := getNEIPv4RouteClass().Alloc()
	rv := objc.Send[NEIPv4Route](instance.ID, objc.Sel("initWithDestinationAddress:subnetMask:"), address, subnetMask)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEIPv4RouteWithDestinationAddressSubnetMask */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEIPv4Route */

// A convenience method for creating the default IPv4 route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route/default()
func (nc _NEIPv4RouteClass) DefaultRoute() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("defaultRoute"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultRoute) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEIPv4Route */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEIPv4Route */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEIPv4Route */

// The destination network address of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route/destinationAddress
func (n_ NEIPv4Route) DestinationAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("destinationAddress"))
	return rv
}/* debug [instance_properties/getter]: destinationAddress */


// The destination network mask of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route/destinationSubnetMask
func (n_ NEIPv4Route) DestinationSubnetMask() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("destinationSubnetMask"))
	return rv
}/* debug [instance_properties/getter]: destinationSubnetMask */


// The address of the next-hop gateway of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route/gatewayAddress
func (n_ NEIPv4Route) GatewayAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("gatewayAddress"))
	return rv
}/* debug [instance_properties/getter]: gatewayAddress */


// The address of the next-hop gateway of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route/gatewayAddress
func (n_ NEIPv4Route) SetGatewayAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGatewayAddress:"), value)
}/* debug [instance_properties/setter]: gatewayAddress */


// The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/excludedroutes
func (n_ NEIPv4Route) ExcludedRoutes() INEIPv4Route {
	rv := objc.Send[NEIPv4Route](n_.ID, objc.Sel("excludedRoutes"))
	return rv
}/* debug [instance_properties/getter]: excludedRoutes */


// The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/excludedroutes
func (n_ NEIPv4Route) SetExcludedRoutes(value INEIPv4Route) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedRoutes:"), value)
}/* debug [instance_properties/setter]: excludedRoutes */


// The IPv4 network traffic that the system routes to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/includedroutes
func (n_ NEIPv4Route) IncludedRoutes() INEIPv4Route {
	rv := objc.Send[NEIPv4Route](n_.ID, objc.Sel("includedRoutes"))
	return rv
}/* debug [instance_properties/getter]: includedRoutes */


// The IPv4 network traffic that the system routes to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/includedroutes
func (n_ NEIPv4Route) SetIncludedRoutes(value INEIPv4Route) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludedRoutes:"), value)
}/* debug [instance_properties/setter]: includedRoutes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEIPv4Route */


