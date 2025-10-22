// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	DestinationAddress() string
	SetDestinationAddress(value string)
	DestinationNetworkPrefixLength() foundation.Number
	SetDestinationNetworkPrefixLength(value foundation.INumber)
	GatewayAddress() string
	SetGatewayAddress(value string)
	ExcludedRoutes() NEIPv6Route
	SetExcludedRoutes(value INEIPv6Route)
	IncludedRoutes() NEIPv6Route
	SetIncludedRoutes(value INEIPv6Route)
}

// The settings for an IPv6 route.
//
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

// Alloc allocates a new instance without initialization.
func (nc _NEIPv6RouteClass) Alloc() NEIPv6Route {
	rv := objc.Send[NEIPv6Route](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The destination network address of the route.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6route/destinationaddress
func (n_ NEIPv6Route) DestinationAddress() string {
	rv := objc.Send[string](n_.ID, objc.Sel("destinationAddress"))
	return rv
}


// SetDestinationAddress sets the value of the destinationAddress property.
// The destination network address of the route.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6route/destinationaddress
func (n_ NEIPv6Route) SetDestinationAddress(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDestinationAddress:"), objc.String(value))
}

// The destination network prefix length of the route.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6route/destinationnetworkprefixlength
func (n_ NEIPv6Route) DestinationNetworkPrefixLength() foundation.Number {
	rv := objc.Send[foundation.Number](n_.ID, objc.Sel("destinationNetworkPrefixLength"))
	return rv
}


// SetDestinationNetworkPrefixLength sets the value of the destinationNetworkPrefixLength property.
// The destination network prefix length of the route.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6route/destinationnetworkprefixlength
func (n_ NEIPv6Route) SetDestinationNetworkPrefixLength(value foundation.INumber) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDestinationNetworkPrefixLength:"), value)
}

// The address of the next-hop gateway of the route.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6route/gatewayaddress
func (n_ NEIPv6Route) GatewayAddress() string {
	rv := objc.Send[string](n_.ID, objc.Sel("gatewayAddress"))
	return rv
}


// SetGatewayAddress sets the value of the gatewayAddress property.
// The address of the next-hop gateway of the route.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6route/gatewayaddress
func (n_ NEIPv6Route) SetGatewayAddress(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGatewayAddress:"), objc.String(value))
}

// The IPv6 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6settings/excludedroutes
func (n_ NEIPv6Route) ExcludedRoutes() NEIPv6Route {
	rv := objc.Send[NEIPv6Route](n_.ID, objc.Sel("excludedRoutes"))
	return rv
}


// SetExcludedRoutes sets the value of the excludedRoutes property.
// The IPv6 network traffic that the system routes to the primary physical interface, not the TUN interface.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6settings/excludedroutes
func (n_ NEIPv6Route) SetExcludedRoutes(value INEIPv6Route) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedRoutes:"), value)
}

// The IPv6 network traffic that the system routes to the TUN interface.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6settings/includedroutes
func (n_ NEIPv6Route) IncludedRoutes() NEIPv6Route {
	rv := objc.Send[NEIPv6Route](n_.ID, objc.Sel("includedRoutes"))
	return rv
}


// SetIncludedRoutes sets the value of the includedRoutes property.
// The IPv6 network traffic that the system routes to the TUN interface.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv6settings/includedroutes
func (n_ NEIPv6Route) SetIncludedRoutes(value INEIPv6Route) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludedRoutes:"), value)
}



