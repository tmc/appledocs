// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [RoutePickerView] class.
var (
	RoutePickerViewClass     _RoutePickerViewClass
	RoutePickerViewClassOnce sync.Once
)

func getRoutePickerViewClass() _RoutePickerViewClass {
	RoutePickerViewClassOnce.Do(func() {
		RoutePickerViewClass = _RoutePickerViewClass{objc.GetClass("AVRoutePickerView")}
	})
	return RoutePickerViewClass
}

type _RoutePickerViewClass struct {
	class objc.Class
}

// An interface definition for the [RoutePickerView] class.
type IRoutePickerView interface {
	appkit.IView
	RoutePickerButtonColorForState(state unsafe.Pointer) unsafe.Pointer
	SetRoutePickerButtonColorForState(color unsafe.Pointer, state unsafe.Pointer)
}

// A view that presents a list of nearby media receivers.
//
// This view represents a button that users tap to stream audio/video content to a media receiver, such as a Mac or Apple TV. When the user taps the button, the system presents a popover that displays all of the nearby AirPlay devices that can receive and play back media. If your app prefers video content, the system displays video-capable devices higher in the list. In iOS 16 and later, you can add devices to the list that implement custom protocols. For more information about displaying third-party routes, see .
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView
type RoutePickerView struct {
	appkit.View
}

// RoutePickerViewFrom constructs a [RoutePickerView] from an unsafe.Pointer.
//
// A view that presents a list of nearby media receivers.
func RoutePickerViewFrom(ptr unsafe.Pointer) RoutePickerView {
	return RoutePickerView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RoutePickerViewClass) Alloc() RoutePickerView {
	rv := objc.Send[RoutePickerView](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RoutePickerViewClass) New() RoutePickerView {
	rv := objc.Send[RoutePickerView](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RoutePickerView) Init() RoutePickerView {
	rv := objc.Send[RoutePickerView](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RoutePickerView) Autorelease() RoutePickerView {
	rv := objc.Send[RoutePickerView](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRoutePickerView creates a new RoutePickerView instance.
func NewRoutePickerView() RoutePickerView {
	return getRoutePickerViewClass().New()
}


// Returns the color of the picker button for the specified state.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/routePickerButtonColor(for:)
func (r_ RoutePickerView) RoutePickerButtonColorForState(state unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("routePickerButtonColorForState:"), state)
	return rv
}

// Sets the route picker button color for the specified state.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/setRoutePickerButtonColor(_:for:)
func (r_ RoutePickerView) SetRoutePickerButtonColorForState(color unsafe.Pointer, state unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRoutePickerButtonColor:forState:"), color, state)
}

// The view’s tint color when AirPlay is active.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/activeTintColor
func (r_ RoutePickerView) ActiveTintColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("activeTintColor"))
	return rv
}


// SetActiveTintColor sets the value of the activeTintColor property.
// The view’s tint color when AirPlay is active.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/activeTintColor
func (r_ RoutePickerView) SetActiveTintColor(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setActiveTintColor:"), value)
}

// A routing controller that enables connections to non-AirPlay devices.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/customRoutingController
func (r_ RoutePickerView) CustomRoutingController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("customRoutingController"))
	return rv
}


// SetCustomRoutingController sets the value of the customRoutingController property.
// A routing controller that enables connections to non-AirPlay devices.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/customRoutingController
func (r_ RoutePickerView) SetCustomRoutingController(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCustomRoutingController:"), value)
}

// The delegate object for the route picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/delegate
func (r_ RoutePickerView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object for the route picker.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/delegate
func (r_ RoutePickerView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the route picker button has a border.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/isRoutePickerButtonBordered
func (r_ RoutePickerView) RoutePickerButtonBordered() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("routePickerButtonBordered"))
	return rv
}


// SetRoutePickerButtonBordered sets the value of the routePickerButtonBordered property.
// A Boolean value that indicates whether the route picker button has a border.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/isRoutePickerButtonBordered
func (r_ RoutePickerView) SetRoutePickerButtonBordered(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRoutePickerButtonBordered:"), value)
}

// The player object to perform routing operations for.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/player
func (r_ RoutePickerView) Player() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("player"))
	return rv
}


// SetPlayer sets the value of the player property.
// The player object to perform routing operations for.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/player
func (r_ RoutePickerView) SetPlayer(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPlayer:"), value)
}

// A Boolean value that indicates whether the route picker sorts video output devices to the top of the list.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/prioritizesVideoDevices
func (r_ RoutePickerView) PrioritizesVideoDevices() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("prioritizesVideoDevices"))
	return rv
}


// SetPrioritizesVideoDevices sets the value of the prioritizesVideoDevices property.
// A Boolean value that indicates whether the route picker sorts video output devices to the top of the list.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/prioritizesVideoDevices
func (r_ RoutePickerView) SetPrioritizesVideoDevices(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPrioritizesVideoDevices:"), value)
}

// The button style for the route picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/routePickerButtonStyle
func (r_ RoutePickerView) RoutePickerButtonStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("routePickerButtonStyle"))
	return rv
}


// SetRoutePickerButtonStyle sets the value of the routePickerButtonStyle property.
// The button style for the route picker.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/routePickerButtonStyle
func (r_ RoutePickerView) SetRoutePickerButtonStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRoutePickerButtonStyle:"), value)
}

// A Boolean value that indicates whether the route picker button has a border.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/isroutepickerbuttonbordered
func (r_ RoutePickerView) IsRoutePickerButtonBordered() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isRoutePickerButtonBordered"))
	return rv
}


// SetIsRoutePickerButtonBordered sets the value of the isRoutePickerButtonBordered property.
// A Boolean value that indicates whether the route picker button has a border.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/isroutepickerbuttonbordered
func (r_ RoutePickerView) SetIsRoutePickerButtonBordered(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsRoutePickerButtonBordered:"), value)
}




