//go:build darwin && ios

// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CustomRoutingController


// Sets the active state of a route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/setActive(_:for:)
func (c_ CustomRoutingController) SetActiveForRoute(active bool, route IAVCustomDeviceRoute) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActive:forRoute:"), active, route)
}

// iOS-only properties

// An array of custom action items to add to a route picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/customActionItems
func (c_ CustomRoutingController) CustomActionItems() []ICustomRoutingActionItem {
	rv := objc.Send[[]CustomRoutingActionItem](c_.ID, objc.Sel("customActionItems"))
	return rv
}
func (c_ CustomRoutingController) SetCustomActionItems(value []ICustomRoutingActionItem) {
	c_.ID.Send(objc.RegisterName("setCustomActionItems:"), value)
}

// A delegate object for a routing controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController/delegate
func (c_ CustomRoutingController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}
func (c_ CustomRoutingController) SetDelegate(value objc.ID) {
	c_.ID.Send(objc.RegisterName("setDelegate:"), value)
}





