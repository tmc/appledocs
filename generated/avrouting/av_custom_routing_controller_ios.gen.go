//go:build darwin && ios

// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CustomRoutingController


// Revokes an app’s authorization to connect to a route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/invalidateAuthorization(for:)
func (c_ CustomRoutingController) InvalidateAuthorizationForRoute(route IAVCustomDeviceRoute) {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidateAuthorizationForRoute:"), route)
}

// Returns a Boolean value that indicates whether a route is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/isRouteActive(_:)
func (c_ CustomRoutingController) IsRouteActive(route IAVCustomDeviceRoute) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isRouteActive:"), route)
	return rv
}

// Sets the active state of a route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/setActive(_:for:)
func (c_ CustomRoutingController) SetActiveForRoute(active bool, route IAVCustomDeviceRoute) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActive:forRoute:"), active, route)
}

// iOS-only properties

// A list of authorized routes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/authorizedRoutes
func (c_ CustomRoutingController) AuthorizedRoutes() []CustomDeviceRoute {
	rv := objc.Send[[]CustomDeviceRoute](c_.ID, objc.Sel("authorizedRoutes"))
	return rv
}

// An array of custom action items to add to a route picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/customActionItems
func (c_ CustomRoutingController) CustomActionItems() []CustomRoutingActionItem {
	rv := objc.Send[[]CustomRoutingActionItem](c_.ID, objc.Sel("customActionItems"))
	return rv
}
func (c_ CustomRoutingController) SetCustomActionItems(value []CustomRoutingActionItem) {
	c_.ID.Send(objc.RegisterName("setCustomActionItems:"), value)
}

// A delegate object for a routing controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/delegate
func (c_ CustomRoutingController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}
func (c_ CustomRoutingController) SetDelegate(value unsafe.Pointer) {
	c_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// An array of route addresses known to be on the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/knownRouteIPs
func (c_ CustomRoutingController) KnownRouteIPs() []CustomRoutingPartialIP {
	rv := objc.Send[[]CustomRoutingPartialIP](c_.ID, objc.Sel("knownRouteIPs"))
	return rv
}
func (c_ CustomRoutingController) SetKnownRouteIPs(value []CustomRoutingPartialIP) {
	c_.ID.Send(objc.RegisterName("setKnownRouteIPs:"), value)
}





