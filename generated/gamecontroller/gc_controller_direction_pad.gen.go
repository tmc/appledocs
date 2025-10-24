// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCControllerDirectionPad */


/* debug [class_header]: Header for GCControllerDirectionPad */
// The class instance for the [GCControllerDirectionPad] class.
var (
	GCControllerDirectionPadClass     _GCControllerDirectionPadClass
	GCControllerDirectionPadClassOnce sync.Once
)

func getGCControllerDirectionPadClass() _GCControllerDirectionPadClass {
	GCControllerDirectionPadClassOnce.Do(func() {
		GCControllerDirectionPadClass = _GCControllerDirectionPadClass{objc.GetClass("GCControllerDirectionPad")}
	})
	return GCControllerDirectionPadClass
}

type _GCControllerDirectionPadClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCControllerDirectionPad */
// An interface definition for the [GCControllerDirectionPad] class.
type IGCControllerDirectionPad interface {
	IGCControllerElement
	
/* debug [class_interface_properties]: Properties for GCControllerDirectionPad */
	// properties:
	Down() IGCControllerButtonInput
	Left() IGCControllerButtonInput
	Right() IGCControllerButtonInput
	Up() IGCControllerButtonInput
	ValueChangedHandler() unsafe.Pointer
	SetValueChangedHandler(value unsafe.Pointer)
	XAxis() IGCControllerAxisInput
	YAxis() IGCControllerAxisInput
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCControllerDirectionPad */
	// methods:
	SetValueForXAxisYAxis(xAxis float32, yAxis float32)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCControllerDirectionPad */
// Alloc allocates a new instance without initialization.
func (gc _GCControllerDirectionPadClass) Alloc() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCControllerDirectionPadClass) New() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCControllerDirectionPad) Init() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCControllerDirectionPad) Autorelease() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerDirectionPad creates a new GCControllerDirectionPad instance.
func NewGCControllerDirectionPad() GCControllerDirectionPad {
	return getGCControllerDirectionPadClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCControllerDirectionPad */
// A control element associated with a directional pad or a thumbstick.
//
// You get the input values for this element from its subelements. You can use either the and properties to get coordinates, or the , , , and buttons that simulate directional pad buttons.


// A control element associated with a directional pad or a thumbstick.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad
type GCControllerDirectionPad struct {
	GCControllerElement
}

// GCControllerDirectionPadFrom constructs a [GCControllerDirectionPad] from an unsafe.Pointer.
//
// A control element associated with a directional pad or a thumbstick.
func GCControllerDirectionPadFrom(ptr unsafe.Pointer) GCControllerDirectionPad {
	return GCControllerDirectionPad{
		GCControllerElement: GCControllerElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCControllerDirectionPad *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCControllerDirectionPad */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCControllerDirectionPad */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCControllerDirectionPad */

// Sets the input values of a snapshot of a directional pad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/setValueForXAxis(_:yAxis:)
func (g_ GCControllerDirectionPad) SetValueForXAxisYAxis(xAxis float32, yAxis float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueForXAxis:yAxis:"), xAxis, yAxis)
}/* debug [instance_methods/method]: SetValueForXAxisYAxis */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCControllerDirectionPad */

// The button element used for the negative y-axis direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/down
func (g_ GCControllerDirectionPad) Down() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("down"))
	return rv
}/* debug [instance_properties/getter]: down */


// The button element that changes the negative x-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/left
func (g_ GCControllerDirectionPad) Left() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("left"))
	return rv
}/* debug [instance_properties/getter]: left */


// The button element that changes the positive x-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/right
func (g_ GCControllerDirectionPad) Right() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("right"))
	return rv
}/* debug [instance_properties/getter]: right */


// The button element that changes the positive y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/up
func (g_ GCControllerDirectionPad) Up() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("up"))
	return rv
}/* debug [instance_properties/getter]: up */


// The block that the directional pad calls when the user changes its values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/valueChangedHandler
func (g_ GCControllerDirectionPad) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}/* debug [instance_properties/getter]: valueChangedHandler */


// The block that the directional pad calls when the user changes its values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/valueChangedHandler
func (g_ GCControllerDirectionPad) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}/* debug [instance_properties/setter]: valueChangedHandler */


// The x-axis element of the directional pad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/xAxis
func (g_ GCControllerDirectionPad) XAxis() IGCControllerAxisInput {
	rv := objc.Send[GCControllerAxisInput](g_.ID, objc.Sel("xAxis"))
	return rv
}/* debug [instance_properties/getter]: xAxis */


// The y-axis element of the directional pad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/yAxis
func (g_ GCControllerDirectionPad) YAxis() IGCControllerAxisInput {
	rv := objc.Send[GCControllerAxisInput](g_.ID, objc.Sel("yAxis"))
	return rv
}/* debug [instance_properties/getter]: yAxis */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCControllerDirectionPad */



