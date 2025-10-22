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
	InvalidateAuthorizationForRoute(route IAVCustomDeviceRoute)
	SetActiveForRoute(active bool, route IAVCustomDeviceRoute)
	CustomActionItems() []CustomRoutingActionItem
	SetCustomActionItems(value []CustomRoutingActionItem)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	AuthorizedRoutes() AVCustomDeviceRoute
	SetAuthorizedRoutes(value IAVCustomDeviceRoute)
	KnownRouteIPs() AVCustomRoutingPartialIP
	SetKnownRouteIPs(value IAVCustomRoutingPartialIP)
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




// Revokes an app’s authorization to connect to a route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/invalidateAuthorization(for:)

func (c_ CustomRoutingController) InvalidateAuthorizationForRoute(route IAVCustomDeviceRoute) {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidateAuthorizationForRoute:"), route)
}



// Sets the active state of a route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/setActive(_:for:)

func (c_ CustomRoutingController) SetActiveForRoute(active bool, route IAVCustomDeviceRoute) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActive:forRoute:"), active, route)
}


// An array of custom action items to add to a route picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/customActionItems

func (c_ CustomRoutingController) CustomActionItems() []CustomRoutingActionItem {
	rv := objc.Send[[]CustomRoutingActionItem](c_.ID, objc.Sel("customActionItems"))
	return rv
}


// An array of custom action items to add to a route picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/customActionItems

func (c_ CustomRoutingController) SetCustomActionItems(value []CustomRoutingActionItem) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setCustomActionItems:"), nsArray)
}


// A delegate object for a routing controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/delegate

func (c_ CustomRoutingController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// A delegate object for a routing controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/delegate

func (c_ CustomRoutingController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}


// A list of authorized routes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingcontroller/authorizedroutes

func (c_ CustomRoutingController) AuthorizedRoutes() AVCustomDeviceRoute {
	rv := objc.Send[AVCustomDeviceRoute](c_.ID, objc.Sel("authorizedRoutes"))
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

func (c_ CustomRoutingController) KnownRouteIPs() AVCustomRoutingPartialIP {
	rv := objc.Send[AVCustomRoutingPartialIP](c_.ID, objc.Sel("knownRouteIPs"))
	return rv
}


// An array of route addresses known to be on the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avrouting/avcustomroutingcontroller/knownrouteips

func (c_ CustomRoutingController) SetKnownRouteIPs(value IAVCustomRoutingPartialIP) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKnownRouteIPs:"), value)
}



