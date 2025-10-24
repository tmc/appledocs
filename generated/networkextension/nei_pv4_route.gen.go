// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [NEIPv4Route] class.
type INEIPv4Route interface {
	objectivec.IObject
	// properties:
	DestinationAddress() objc.IObject /* cross-framework: NSString */
	SetDestinationAddress(value objc.IObject /* cross-framework: NSString */)
	DestinationSubnetMask() objc.IObject /* cross-framework: NSString */
	SetDestinationSubnetMask(value objc.IObject /* cross-framework: NSString */)
	GatewayAddress() objc.IObject /* cross-framework: NSString */
	SetGatewayAddress(value objc.IObject /* cross-framework: NSString */)
	ExcludedRoutes() INEIPv4Route
	SetExcludedRoutes(value INEIPv4Route)
	IncludedRoutes() INEIPv4Route
	SetIncludedRoutes(value INEIPv4Route)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (nc _NEIPv4RouteClass) Alloc() NEIPv4Route {
	rv := objc.Send[NEIPv4Route](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Initialize the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route/init(destinationAddress:subnetMask:)
func NewNEIPv4RouteWithDestinationAddressSubnetMask(address objc.IObject /* cross-framework: NSString */, subnetMask objc.IObject /* cross-framework: NSString */) NEIPv4Route {
	instance := getNEIPv4RouteClass().Alloc()
	rv := objc.Send[NEIPv4Route](instance.ID, objc.Sel("initWithDestinationAddress:subnetMask:"), address, subnetMask)
	rv.Autorelease()
	return rv
}



// The destination network address of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4route/destinationaddress
func (n_ NEIPv4Route) DestinationAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("destinationAddress"))
	return rv
}


// The destination network address of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4route/destinationaddress
func (n_ NEIPv4Route) SetDestinationAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDestinationAddress:"), value)
}


// The destination network mask of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4route/destinationsubnetmask
func (n_ NEIPv4Route) DestinationSubnetMask() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("destinationSubnetMask"))
	return rv
}


// The destination network mask of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4route/destinationsubnetmask
func (n_ NEIPv4Route) SetDestinationSubnetMask(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDestinationSubnetMask:"), value)
}


// The address of the next-hop gateway of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4route/gatewayaddress
func (n_ NEIPv4Route) GatewayAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("gatewayAddress"))
	return rv
}


// The address of the next-hop gateway of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4route/gatewayaddress
func (n_ NEIPv4Route) SetGatewayAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGatewayAddress:"), value)
}


// The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/excludedroutes
func (n_ NEIPv4Route) ExcludedRoutes() INEIPv4Route {
	rv := objc.Send[NEIPv4Route](n_.ID, objc.Sel("excludedRoutes"))
	return rv
}


// The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/excludedroutes
func (n_ NEIPv4Route) SetExcludedRoutes(value INEIPv4Route) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedRoutes:"), value)
}


// The IPv4 network traffic that the system routes to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/includedroutes
func (n_ NEIPv4Route) IncludedRoutes() INEIPv4Route {
	rv := objc.Send[NEIPv4Route](n_.ID, objc.Sel("includedRoutes"))
	return rv
}


// The IPv4 network traffic that the system routes to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/includedroutes
func (n_ NEIPv4Route) SetIncludedRoutes(value INEIPv4Route) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludedRoutes:"), value)
}


