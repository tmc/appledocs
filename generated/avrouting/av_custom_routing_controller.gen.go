// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CustomRoutingController] class.
var (
	CustomRoutingControllerClass     _CustomRoutingControllerClass
	CustomRoutingControllerClassOnce sync.Once
)

func getCustomRoutingControllerClass() _CustomRoutingControllerClass {
	CustomRoutingControllerClassOnce.Do(func() {
		CustomRoutingControllerClass = _CustomRoutingControllerClass{objc.GetClass("AVCustomRoutingController")}
	})
	return CustomRoutingControllerClass
}

type _CustomRoutingControllerClass struct {
	class objc.Class
}

// An interface definition for the [CustomRoutingController] class.
type ICustomRoutingController interface {
	objectivec.IObject
	// properties:
	AuthorizedRoutes() IAVCustomDeviceRoute
	SetAuthorizedRoutes(value IAVCustomDeviceRoute)
	KnownRouteIPs() objc.IObject /* cross-framework: CustomRoutingPartialIP */
	SetKnownRouteIPs(value objc.IObject /* cross-framework: CustomRoutingPartialIP */)
	// methods:
}

// An object that manages the connection from a device to a destination.
//
// A routing controller also informs its object about which routes the user previously authorized, so it can reconnect, if appropriate.


// An object that manages the connection from a device to a destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController
type CustomRoutingController struct {
	objectivec.Object
}

// CustomRoutingControllerFrom constructs a [CustomRoutingController] from an unsafe.Pointer.
//
// An object that manages the connection from a device to a destination.
func CustomRoutingControllerFrom(ptr unsafe.Pointer) CustomRoutingController {
	return CustomRoutingController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CustomRoutingControllerClass) Alloc() CustomRoutingController {
	rv := objc.Send[CustomRoutingController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CustomRoutingControllerClass) New() CustomRoutingController {
	rv := objc.Send[CustomRoutingController](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomRoutingController) Init() CustomRoutingController {
	rv := objc.Send[CustomRoutingController](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomRoutingController) Autorelease() CustomRoutingController {
	rv := objc.Send[CustomRoutingController](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomRoutingController creates a new CustomRoutingController instance.
func NewCustomRoutingController() CustomRoutingController {
	return getCustomRoutingControllerClass().New()
}



// A list of authorized routes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingcontroller/authorizedroutes
func (c_ CustomRoutingController) AuthorizedRoutes() IAVCustomDeviceRoute {
	rv := objc.Send[CustomDeviceRoute](c_.ID, objc.Sel("authorizedRoutes"))
	return rv
}


// A list of authorized routes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingcontroller/authorizedroutes
func (c_ CustomRoutingController) SetAuthorizedRoutes(value IAVCustomDeviceRoute) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthorizedRoutes:"), value)
}


// An array of route addresses known to be on the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingcontroller/knownrouteips
func (c_ CustomRoutingController) KnownRouteIPs() objc.IObject /* cross-framework: CustomRoutingPartialIP */ {
	rv := objc.Send[CustomRoutingPartialIP](c_.ID, objc.Sel("knownRouteIPs"))
	return rv
}


// An array of route addresses known to be on the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingcontroller/knownrouteips
func (c_ CustomRoutingController) SetKnownRouteIPs(value objc.IObject /* cross-framework: CustomRoutingPartialIP */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKnownRouteIPs:"), value)
}


