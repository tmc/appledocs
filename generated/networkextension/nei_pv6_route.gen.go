// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [NEIPv6Route] class.
type INEIPv6Route interface {
	objectivec.IObject
	

	// properties:
	DestinationAddress() foundation.foundation.INSString
	DestinationNetworkPrefixLength() foundation.foundation.INSNumber
	GatewayAddress() foundation.foundation.INSString
	SetGatewayAddress(value foundation.foundation.INSString)
	ExcludedRoutes() INEIPv6Route
	SetExcludedRoutes(value INEIPv6Route)
	IncludedRoutes() INEIPv6Route
	SetIncludedRoutes(value INEIPv6Route)


	

	// methods:


}





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






// Initialize the NEIPv6Route
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/init(destinationAddress:networkPrefixLength:)
func NewNEIPv6RouteWithDestinationAddressNetworkPrefixLength(address foundation.foundation.INSString, networkPrefixLength foundation.foundation.INSNumber) NEIPv6Route {
	instance := getNEIPv6RouteClass().Alloc()
	rv := objc.Send[NEIPv6Route](instance.ID, objc.Sel("initWithDestinationAddress:networkPrefixLength:"), address, networkPrefixLength)
	rv.Autorelease()
	return rv
}







// A convenience method for creating the default IPv4 route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/default()
func (nc _NEIPv6RouteClass) DefaultRoute() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("defaultRoute"))
	return rv
}

















// The destination network address of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/destinationAddress
func (n_ NEIPv6Route) DestinationAddress() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("destinationAddress"))
	return rv
}


// The destination network prefix length of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/destinationNetworkPrefixLength
func (n_ NEIPv6Route) DestinationNetworkPrefixLength() foundation.foundation.INSNumber {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("destinationNetworkPrefixLength"))
	return rv
}


// The address of the next-hop gateway of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/gatewayAddress
func (n_ NEIPv6Route) GatewayAddress() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("gatewayAddress"))
	return rv
}


// The address of the next-hop gateway of the route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route/gatewayAddress
func (n_ NEIPv6Route) SetGatewayAddress(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGatewayAddress:"), value)
}


// The IPv6 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6settings/excludedroutes
func (n_ NEIPv6Route) ExcludedRoutes() INEIPv6Route {
	rv := objc.Send[NEIPv6Route](n_.ID, objc.Sel("excludedRoutes"))
	return rv
}


// The IPv6 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6settings/excludedroutes
func (n_ NEIPv6Route) SetExcludedRoutes(value INEIPv6Route) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedRoutes:"), value)
}


// The IPv6 network traffic that the system routes to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6settings/includedroutes
func (n_ NEIPv6Route) IncludedRoutes() INEIPv6Route {
	rv := objc.Send[NEIPv6Route](n_.ID, objc.Sel("includedRoutes"))
	return rv
}


// The IPv6 network traffic that the system routes to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6settings/includedroutes
func (n_ NEIPv6Route) SetIncludedRoutes(value INEIPv6Route) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludedRoutes:"), value)
}







