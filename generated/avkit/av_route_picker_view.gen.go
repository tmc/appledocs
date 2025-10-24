// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class AVRoutePickerView */


/* debug [class_header]: Header for AVRoutePickerView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RoutePickerView */
// An interface definition for the [RoutePickerView] class.
type IRoutePickerView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for RoutePickerView */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	RoutePickerButtonBordered() bool
	SetRoutePickerButtonBordered(value bool)
	Player() objc.IObject /* cross-framework: Player */
	SetPlayer(value objc.IObject /* cross-framework: Player */)
	IsRoutePickerButtonBordered() bool
	SetIsRoutePickerButtonBordered(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RoutePickerView */
	// methods:
	RoutePickerButtonColorForState(state RoutePickerViewButtonState) appkit.Color
	SetRoutePickerButtonColorForState(color appkit.Color, state RoutePickerViewButtonState)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RoutePickerView */
// Alloc allocates a new instance without initialization.
func (rc _RoutePickerViewClass) Alloc() RoutePickerView {
	rv := objc.Send[RoutePickerView](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RoutePickerView */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RoutePickerView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RoutePickerView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RoutePickerView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RoutePickerView */

// Returns the color of the picker button for the specified state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/routePickerButtonColor(for:)
func (r_ RoutePickerView) RoutePickerButtonColorForState(state RoutePickerViewButtonState) appkit.Color {
	rv := objc.Send[appkit.Color](r_.ID, objc.Sel("routePickerButtonColorForState:"), state)
	return rv
}/* debug [instance_methods/method]: RoutePickerButtonColorForState */


// Sets the route picker button color for the specified state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/setRoutePickerButtonColor(_:for:)
func (r_ RoutePickerView) SetRoutePickerButtonColorForState(color appkit.Color, state RoutePickerViewButtonState) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRoutePickerButtonColor:forState:"), color, state)
}/* debug [instance_methods/method]: SetRoutePickerButtonColorForState */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RoutePickerView */

// The delegate object for the route picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/delegate
func (r_ RoutePickerView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object for the route picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/delegate
func (r_ RoutePickerView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that indicates whether the route picker button has a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/isRoutePickerButtonBordered
func (r_ RoutePickerView) RoutePickerButtonBordered() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("routePickerButtonBordered"))
	return rv
}/* debug [instance_properties/getter]: routePickerButtonBordered */


// A Boolean value that indicates whether the route picker button has a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/isRoutePickerButtonBordered
func (r_ RoutePickerView) SetRoutePickerButtonBordered(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRoutePickerButtonBordered:"), value)
}/* debug [instance_properties/setter]: routePickerButtonBordered */


// The player object to perform routing operations for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/player
func (r_ RoutePickerView) Player() objc.IObject /* cross-framework: Player */ {
	rv := objc.Send[appkit.Player](r_.ID, objc.Sel("player"))
	return rv
}/* debug [instance_properties/getter]: player */


// The player object to perform routing operations for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVRoutePickerView/player
func (r_ RoutePickerView) SetPlayer(value objc.IObject /* cross-framework: Player */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPlayer:"), value)
}/* debug [instance_properties/setter]: player */


// A Boolean value that indicates whether the route picker button has a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/isroutepickerbuttonbordered
func (r_ RoutePickerView) IsRoutePickerButtonBordered() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isRoutePickerButtonBordered"))
	return rv
}/* debug [instance_properties/getter]: isRoutePickerButtonBordered */


// A Boolean value that indicates whether the route picker button has a border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerview/isroutepickerbuttonbordered
func (r_ RoutePickerView) SetIsRoutePickerButtonBordered(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsRoutePickerButtonBordered:"), value)
}/* debug [instance_properties/setter]: isRoutePickerButtonBordered */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVRoutePickerView */


