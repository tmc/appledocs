// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCControllerTouchpad */


/* debug [class_header]: Header for GCControllerTouchpad */
// The class instance for the [GCControllerTouchpad] class.
var (
	GCControllerTouchpadClass     _GCControllerTouchpadClass
	GCControllerTouchpadClassOnce sync.Once
)

func getGCControllerTouchpadClass() _GCControllerTouchpadClass {
	GCControllerTouchpadClassOnce.Do(func() {
		GCControllerTouchpadClass = _GCControllerTouchpadClass{objc.GetClass("GCControllerTouchpad")}
	})
	return GCControllerTouchpadClass
}

type _GCControllerTouchpadClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCControllerTouchpad */
// An interface definition for the [GCControllerTouchpad] class.
type IGCControllerTouchpad interface {
	IGCControllerElement
	
/* debug [class_interface_properties]: Properties for GCControllerTouchpad */
	// properties:
	Button() IGCControllerButtonInput
	ReportsAbsoluteTouchSurfaceValues() bool
	SetReportsAbsoluteTouchSurfaceValues(value bool)
	TouchDown() unsafe.Pointer
	SetTouchDown(value unsafe.Pointer)
	TouchMoved() unsafe.Pointer
	SetTouchMoved(value unsafe.Pointer)
	TouchState() GCTouchState
	TouchSurface() IGCControllerDirectionPad
	TouchUp() unsafe.Pointer
	SetTouchUp(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCControllerTouchpad */
	// methods:
	SetValueForXAxisYAxisTouchDownButtonValue(xAxis float32, yAxis float32, touchDown bool, buttonValue float32)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCControllerTouchpad */
// Alloc allocates a new instance without initialization.
func (gc _GCControllerTouchpadClass) Alloc() GCControllerTouchpad {
	rv := objc.Send[GCControllerTouchpad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCControllerTouchpadClass) New() GCControllerTouchpad {
	rv := objc.Send[GCControllerTouchpad](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCControllerTouchpad) Init() GCControllerTouchpad {
	rv := objc.Send[GCControllerTouchpad](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCControllerTouchpad) Autorelease() GCControllerTouchpad {
	rv := objc.Send[GCControllerTouchpad](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerTouchpad creates a new GCControllerTouchpad instance.
func NewGCControllerTouchpad() GCControllerTouchpad {
	return getGCControllerTouchpadClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCControllerTouchpad */
// A control element that represents a touch event on a touchpad.
//
// A object provides the state of the touches and presses on a touchpad. This is a compound element with button and directional pad subelements.


// A control element that represents a touch event on a touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad
type GCControllerTouchpad struct {
	GCControllerElement
}

// GCControllerTouchpadFrom constructs a [GCControllerTouchpad] from an unsafe.Pointer.
//
// A control element that represents a touch event on a touchpad.
func GCControllerTouchpadFrom(ptr unsafe.Pointer) GCControllerTouchpad {
	return GCControllerTouchpad{
		GCControllerElement: GCControllerElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCControllerTouchpad *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCControllerTouchpad */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCControllerTouchpad */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCControllerTouchpad */

// Sets the input values of a snapshot of a touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/setValueForXAxis(_:yAxis:touchDown:buttonValue:)
func (g_ GCControllerTouchpad) SetValueForXAxisYAxisTouchDownButtonValue(xAxis float32, yAxis float32, touchDown bool, buttonValue float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueForXAxis:yAxis:touchDown:buttonValue:"), xAxis, yAxis, touchDown, buttonValue)
}/* debug [instance_methods/method]: SetValueForXAxisYAxisTouchDownButtonValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCControllerTouchpad */

// The element that represents the button component on the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/button
func (g_ GCControllerTouchpad) Button() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("button"))
	return rv
}/* debug [instance_properties/getter]: button */


// A Boolean value that determines whether the touch values are absolute or relative.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/reportsAbsoluteTouchSurfaceValues
func (g_ GCControllerTouchpad) ReportsAbsoluteTouchSurfaceValues() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("reportsAbsoluteTouchSurfaceValues"))
	return rv
}/* debug [instance_properties/getter]: reportsAbsoluteTouchSurfaceValues */


// A Boolean value that determines whether the touch values are absolute or relative.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/reportsAbsoluteTouchSurfaceValues
func (g_ GCControllerTouchpad) SetReportsAbsoluteTouchSurfaceValues(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReportsAbsoluteTouchSurfaceValues:"), value)
}/* debug [instance_properties/setter]: reportsAbsoluteTouchSurfaceValues */


// The block that the element calls when the user begins touching the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchDown
func (g_ GCControllerTouchpad) TouchDown() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("touchDown"))
	return rv
}/* debug [instance_properties/getter]: touchDown */


// The block that the element calls when the user begins touching the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchDown
func (g_ GCControllerTouchpad) SetTouchDown(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTouchDown:"), value)
}/* debug [instance_properties/setter]: touchDown */


// The block that the element calls when the user continues touching the touchpad, not when the user begins or ends touching the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchMoved
func (g_ GCControllerTouchpad) TouchMoved() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("touchMoved"))
	return rv
}/* debug [instance_properties/getter]: touchMoved */


// The block that the element calls when the user continues touching the touchpad, not when the user begins or ends touching the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchMoved
func (g_ GCControllerTouchpad) SetTouchMoved(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTouchMoved:"), value)
}/* debug [instance_properties/setter]: touchMoved */


// The state of the user’s touch on the surface of the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchState-swift.property
func (g_ GCControllerTouchpad) TouchState() GCTouchState {
	rv := objc.Send[GCTouchState](g_.ID, objc.Sel("touchState"))
	return rv
}/* debug [instance_properties/getter]: touchState */


// The element that represents the state of the user’s touch on the surface of the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchSurface
func (g_ GCControllerTouchpad) TouchSurface() IGCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("touchSurface"))
	return rv
}/* debug [instance_properties/getter]: touchSurface */


// The block that the element calls when the user finishes touching the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchUp
func (g_ GCControllerTouchpad) TouchUp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("touchUp"))
	return rv
}/* debug [instance_properties/getter]: touchUp */


// The block that the element calls when the user finishes touching the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchUp
func (g_ GCControllerTouchpad) SetTouchUp(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTouchUp:"), value)
}/* debug [instance_properties/setter]: touchUp */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCControllerTouchpad */



