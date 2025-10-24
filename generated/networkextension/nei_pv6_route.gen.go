// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEIPv6Route */


/* debug [class_header]: Header for NEIPv6Route */
// The class instance for the [NEIPv6Route] class.
var (
	NEIPv6RouteClass     _NEIPv6RouteClass
	NEIPv6RouteClassOnce sync.Once
)

func getNEIPv6RouteClass() _NEIPv6RouteClass {
	NEIPv6RouteClassOnce.Do(func() {
		NEIPv6RouteClass = _NEIPv6RouteClass{objc.GetClass("NEIPv6Route")}
	})
	return NEIPv6RouteClass
}

type _NEIPv6RouteClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEIPv6Route */
// An interface definition for the [NEIPv6Route] class.
type INEIPv6Route interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEIPv6Route */
	// properties:
	DestinationAddress() objc.IObject /* cross-framework: NSString */
	DestinationNetworkPrefixLength() objc.IObject /* cross-framework: NSNumber */
	GatewayAddress() objc.IObject /* cross-framework: NSString */
	SetGatewayAddress(value objc.IObject /* cross-framework: NSString */)
	ExcludedRoutes() INEIPv6Route
	SetExcludedRoutes(value INEIPv6Route)
	IncludedRoutes() INEIPv6Route
	SetIncludedRoutes(value INEIPv6Route)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEIPv6Route */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEIPv6Route */
// Alloc allocates a new instance without initialization.
func (nc _NEIPv6RouteClass) Alloc() NEIPv6Route {
	rv := objc.Send[NEIPv6Route](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEIPv6RouteClass) New() NEIPv6Route {
	rv := objc.Send[NEIPv6Route](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEIPv6Route) Init() NEIPv6Route {
	rv := objc.Send[NEIPv6Route](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEIPv6Route) Autorelease() NEIPv6Route {
	rv := objc.Send[NEIPv6Route](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEIPv6Route creates a new NEIPv6Route instance.
func NewNEIPv6Route() NEIPv6Route {
	return getNEIPv6RouteClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEIPv6Route */
// The settings for an IPv6 route.


// The settings for an IPv6 route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route
type NEIPv6Route struct {
	objectivec.Object
}

// NEIPv6RouteFrom constructs a [NEIPv6Route] from an unsafe.Pointer.
//
// The settings for an IPv6 route.
func NEIPv6RouteFrom(ptr unsafe.Pointer) NEIPv6Route {
	return NEIPv6Route{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEIPv6Route */

// Initialize the NEIPv6Route
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/init(destinationAddress:networkPrefixLength:)
func NewNEIPv6RouteWithDestinationAddressNetworkPrefixLength(address objc.IObject /* cross-framework: NSString */, networkPrefixLength objc.IObject /* cross-framework: NSNumber */) NEIPv6Route {
	instance := getNEIPv6RouteClass().Alloc()
	rv := objc.Send[NEIPv6Route](instance.ID, objc.Sel("initWithDestinationAddress:networkPrefixLength:"), address, networkPrefixLength)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEIPv6RouteWithDestinationAddressNetworkPrefixLength */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEIPv6Route */

// A convenience method for creating the default IPv4 route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/default()
func (nc _NEIPv6RouteClass) DefaultRoute() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("defaultRoute"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultRoute) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEIPv6Route */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEIPv6Route */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEIPv6Route */

// The destination network address of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/destinationAddress
func (n_ NEIPv6Route) DestinationAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("destinationAddress"))
	return rv
}/* debug [instance_properties/getter]: destinationAddress */


// The destination network prefix length of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/destinationNetworkPrefixLength
func (n_ NEIPv6Route) DestinationNetworkPrefixLength() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("destinationNetworkPrefixLength"))
	return rv
}/* debug [instance_properties/getter]: destinationNetworkPrefixLength */


// The address of the next-hop gateway of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/gatewayAddress
func (n_ NEIPv6Route) GatewayAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("gatewayAddress"))
	return rv
}/* debug [instance_properties/getter]: gatewayAddress */


// The address of the next-hop gateway of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/gatewayAddress
func (n_ NEIPv6Route) SetGatewayAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGatewayAddress:"), value)
}/* debug [instance_properties/setter]: gatewayAddress */


// The IPv6 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6settings/excludedroutes
func (n_ NEIPv6Route) ExcludedRoutes() INEIPv6Route {
	rv := objc.Send[NEIPv6Route](n_.ID, objc.Sel("excludedRoutes"))
	return rv
}/* debug [instance_properties/getter]: excludedRoutes */


// The IPv6 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6settings/excludedroutes
func (n_ NEIPv6Route) SetExcludedRoutes(value INEIPv6Route) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedRoutes:"), value)
}/* debug [instance_properties/setter]: excludedRoutes */


// The IPv6 network traffic that the system routes to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6settings/includedroutes
func (n_ NEIPv6Route) IncludedRoutes() INEIPv6Route {
	rv := objc.Send[NEIPv6Route](n_.ID, objc.Sel("includedRoutes"))
	return rv
}/* debug [instance_properties/getter]: includedRoutes */


// The IPv6 network traffic that the system routes to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6settings/includedroutes
func (n_ NEIPv6Route) SetIncludedRoutes(value INEIPv6Route) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludedRoutes:"), value)
}/* debug [instance_properties/setter]: includedRoutes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEIPv6Route */


