//go:build darwin && ios

// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// iOS-only methods for RoutePickerView


// iOS-only properties

// The view’s tint color when AirPlay is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/activeTintColor
func (r_ RoutePickerView) ActiveTintColor() appkit.Color {
	rv := objc.Send[appkit.Color](r_.ID, objc.Sel("activeTintColor"))
	return rv
}
func (r_ RoutePickerView) SetActiveTintColor(value appkit.Color) {
	r_.ID.Send(objc.RegisterName("setActiveTintColor:"), value)
}

// A routing controller that enables connections to non-AirPlay devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/customRoutingController
func (r_ RoutePickerView) CustomRoutingController() avrouting.CustomRoutingController {
	rv := objc.Send[avrouting.CustomRoutingController](r_.ID, objc.Sel("customRoutingController"))
	return rv
}
func (r_ RoutePickerView) SetCustomRoutingController(value avrouting.CustomRoutingController) {
	r_.ID.Send(objc.RegisterName("setCustomRoutingController:"), value)
}

// A Boolean value that indicates whether the route picker sorts video output devices to the top of the list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/prioritizesVideoDevices
func (r_ RoutePickerView) PrioritizesVideoDevices() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("prioritizesVideoDevices"))
	return rv
}
func (r_ RoutePickerView) SetPrioritizesVideoDevices(value bool) {
	r_.ID.Send(objc.RegisterName("setPrioritizesVideoDevices:"), value)
}

// The button style for the route picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/routePickerButtonStyle
func (r_ RoutePickerView) RoutePickerButtonStyle() RoutePickerViewButtonStyle {
	rv := objc.Send[RoutePickerViewButtonStyle](r_.ID, objc.Sel("routePickerButtonStyle"))
	return rv
}
func (r_ RoutePickerView) SetRoutePickerButtonStyle(value RoutePickerViewButtonStyle) {
	r_.ID.Send(objc.RegisterName("setRoutePickerButtonStyle:"), value)
}





