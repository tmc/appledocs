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
	// properties:
	ActiveTintColor() objc.IObject /* cross-framework: Color */
	SetActiveTintColor(value objc.IObject /* cross-framework: Color */)
	CustomRoutingController() objc.IObject /* cross-framework: CustomRoutingController */
	SetCustomRoutingController(value objc.IObject /* cross-framework: CustomRoutingController */)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	RoutePickerButtonBordered() bool /* primitive/slice/pointer. */
	SetRoutePickerButtonBordered(value bool /* primitive/slice/pointer. */)
	Player() objc.IObject /* cross-framework: Player */
	SetPlayer(value objc.IObject /* cross-framework: Player */)
	PrioritizesVideoDevices() bool /* primitive/slice/pointer. */
	SetPrioritizesVideoDevices(value bool /* primitive/slice/pointer. */)
	RoutePickerButtonStyle() RoutePickerViewButtonStyle
	SetRoutePickerButtonStyle(value RoutePickerViewButtonStyle)
	IsRoutePickerButtonBordered() bool /* primitive/slice/pointer. */
	SetIsRoutePickerButtonBordered(value bool /* primitive/slice/pointer. */)
	// methods:
	RoutePickerButtonColorForState(state RoutePickerViewButtonState) objc.IObject /* cross-framework: Color */
	SetRoutePickerButtonColorForState(color objc.IObject /* cross-framework Color */, state RoutePickerViewButtonState)
}

// A view that presents a list of nearby media receivers.
//
// This view represents a button that users tap to stream audio/video content to a media receiver, such as a Mac or Apple TV. When the user taps the button, the system presents a popover that displays all of the nearby AirPlay devices that can receive and play back media. If your app prefers video content, the system displays video-capable devices higher in the list. In iOS 16 and later, you can add devices to the list that implement custom protocols. For more information about displaying third-party routes, see .


// A view that presents a list of nearby media receivers.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/routePickerButtonColor(for:)
func (r_ RoutePickerView) RoutePickerButtonColorForState(state RoutePickerViewButtonState) objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[Color](r_.ID, objc.Sel("routePickerButtonColorForState:"), state)
	return rv
}


// Sets the route picker button color for the specified state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/setRoutePickerButtonColor(_:for:)
func (r_ RoutePickerView) SetRoutePickerButtonColorForState(color objc.IObject /* cross-framework Color */, state RoutePickerViewButtonState) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRoutePickerButtonColor:forState:"), color, state)
}


// The view’s tint color when AirPlay is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/activeTintColor
func (r_ RoutePickerView) ActiveTintColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[Color](r_.ID, objc.Sel("activeTintColor"))
	return rv
}


// The view’s tint color when AirPlay is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/activeTintColor
func (r_ RoutePickerView) SetActiveTintColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setActiveTintColor:"), value)
}


// A routing controller that enables connections to non-AirPlay devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/customRoutingController
func (r_ RoutePickerView) CustomRoutingController() objc.IObject /* cross-framework: CustomRoutingController */ {
	rv := objc.Send[CustomRoutingController](r_.ID, objc.Sel("customRoutingController"))
	return rv
}


// A routing controller that enables connections to non-AirPlay devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/customRoutingController
func (r_ RoutePickerView) SetCustomRoutingController(value objc.IObject /* cross-framework: CustomRoutingController */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCustomRoutingController:"), value)
}


// The delegate object for the route picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/delegate
func (r_ RoutePickerView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate object for the route picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/delegate
func (r_ RoutePickerView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that indicates whether the route picker button has a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/isRoutePickerButtonBordered
func (r_ RoutePickerView) RoutePickerButtonBordered() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("routePickerButtonBordered"))
	return rv
}


// A Boolean value that indicates whether the route picker button has a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/isRoutePickerButtonBordered
func (r_ RoutePickerView) SetRoutePickerButtonBordered(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRoutePickerButtonBordered:"), value)
}


// The player object to perform routing operations for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/player
func (r_ RoutePickerView) Player() objc.IObject /* cross-framework: Player */ {
	rv := objc.Send[Player](r_.ID, objc.Sel("player"))
	return rv
}


// The player object to perform routing operations for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/player
func (r_ RoutePickerView) SetPlayer(value objc.IObject /* cross-framework: Player */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPlayer:"), value)
}


// A Boolean value that indicates whether the route picker sorts video output devices to the top of the list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/prioritizesVideoDevices
func (r_ RoutePickerView) PrioritizesVideoDevices() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("prioritizesVideoDevices"))
	return rv
}


// A Boolean value that indicates whether the route picker sorts video output devices to the top of the list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/prioritizesVideoDevices
func (r_ RoutePickerView) SetPrioritizesVideoDevices(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPrioritizesVideoDevices:"), value)
}


// The button style for the route picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/routePickerButtonStyle
func (r_ RoutePickerView) RoutePickerButtonStyle() RoutePickerViewButtonStyle {
	rv := objc.Send[RoutePickerViewButtonStyle](r_.ID, objc.Sel("routePickerButtonStyle"))
	return rv
}


// The button style for the route picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/routePickerButtonStyle
func (r_ RoutePickerView) SetRoutePickerButtonStyle(value RoutePickerViewButtonStyle) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRoutePickerButtonStyle:"), value)
}


// A Boolean value that indicates whether the route picker button has a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/isroutepickerbuttonbordered
func (r_ RoutePickerView) IsRoutePickerButtonBordered() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isRoutePickerButtonBordered"))
	return rv
}


// A Boolean value that indicates whether the route picker button has a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/isroutepickerbuttonbordered
func (r_ RoutePickerView) SetIsRoutePickerButtonBordered(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsRoutePickerButtonBordered:"), value)
}




