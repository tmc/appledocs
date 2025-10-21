// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// The settings for an IPv4 route.
//
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


// The destination network address of the route.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4route/destinationaddress
func (n_ NEIPv4Route) DestinationAddress() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("destinationAddress"))
	return rv
}


// SetDestinationAddress sets the value of the destinationAddress property.
// The destination network address of the route.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4route/destinationaddress
func (n_ NEIPv4Route) SetDestinationAddress(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDestinationAddress:"), value)
}

// The destination network mask of the route.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4route/destinationsubnetmask
func (n_ NEIPv4Route) DestinationSubnetMask() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("destinationSubnetMask"))
	return rv
}


// SetDestinationSubnetMask sets the value of the destinationSubnetMask property.
// The destination network mask of the route.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4route/destinationsubnetmask
func (n_ NEIPv4Route) SetDestinationSubnetMask(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDestinationSubnetMask:"), value)
}

// The address of the next-hop gateway of the route.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4route/gatewayaddress
func (n_ NEIPv4Route) GatewayAddress() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("gatewayAddress"))
	return rv
}


// SetGatewayAddress sets the value of the gatewayAddress property.
// The address of the next-hop gateway of the route.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4route/gatewayaddress
func (n_ NEIPv4Route) SetGatewayAddress(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGatewayAddress:"), value)
}

// The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/excludedroutes
func (n_ NEIPv4Route) ExcludedRoutes() NEIPv4Route {
	rv := objc.Send[NEIPv4Route](n_.ID, objc.Sel("excludedRoutes"))
	return rv
}


// SetExcludedRoutes sets the value of the excludedRoutes property.
// The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/excludedroutes
func (n_ NEIPv4Route) SetExcludedRoutes(value INEIPv4Route) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedRoutes:"), value)
}

// The IPv4 network traffic that the system routes to the TUN interface.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/includedroutes
func (n_ NEIPv4Route) IncludedRoutes() NEIPv4Route {
	rv := objc.Send[NEIPv4Route](n_.ID, objc.Sel("includedRoutes"))
	return rv
}


// SetIncludedRoutes sets the value of the includedRoutes property.
// The IPv4 network traffic that the system routes to the TUN interface.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/includedroutes
func (n_ NEIPv4Route) SetIncludedRoutes(value INEIPv4Route) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludedRoutes:"), value)
}



