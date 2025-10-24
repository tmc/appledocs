// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCControllerButtonInput */


/* debug [class_header]: Header for GCControllerButtonInput */
// The class instance for the [GCControllerButtonInput] class.
var (
	GCControllerButtonInputClass     _GCControllerButtonInputClass
	GCControllerButtonInputClassOnce sync.Once
)

func getGCControllerButtonInputClass() _GCControllerButtonInputClass {
	GCControllerButtonInputClassOnce.Do(func() {
		GCControllerButtonInputClass = _GCControllerButtonInputClass{objc.GetClass("GCControllerButtonInput")}
	})
	return GCControllerButtonInputClass
}

type _GCControllerButtonInputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCControllerButtonInput */
// An interface definition for the [GCControllerButtonInput] class.
type IGCControllerButtonInput interface {
	IGCControllerElement
	
/* debug [class_interface_properties]: Properties for GCControllerButtonInput */
	// properties:
	Pressed() bool
	Touched() bool
	PressedChangedHandler() unsafe.Pointer
	SetPressedChangedHandler(value unsafe.Pointer)
	TouchedChangedHandler() unsafe.Pointer
	SetTouchedChangedHandler(value unsafe.Pointer)
	Value() float32
	ValueChangedHandler() unsafe.Pointer
	SetValueChangedHandler(value unsafe.Pointer)
	IsPressed() bool
	SetIsPressed(value bool)
	IsTouched() bool
	SetIsTouched(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCControllerButtonInput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCControllerButtonInput */
// Alloc allocates a new instance without initialization.
func (gc _GCControllerButtonInputClass) Alloc() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCControllerButtonInputClass) New() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCControllerButtonInput) Init() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCControllerButtonInput) Autorelease() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerButtonInput creates a new GCControllerButtonInput instance.
func NewGCControllerButtonInput() GCControllerButtonInput {
	return getGCControllerButtonInputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCControllerButtonInput */
// A control element that represents a button touch or press.
//
// A object represents a button on a controller that can report either analog or digital values.


// A control element that represents a button touch or press.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput
type GCControllerButtonInput struct {
	GCControllerElement
}

// GCControllerButtonInputFrom constructs a [GCControllerButtonInput] from an unsafe.Pointer.
//
// A control element that represents a button touch or press.
func GCControllerButtonInputFrom(ptr unsafe.Pointer) GCControllerButtonInput {
	return GCControllerButtonInput{
		GCControllerElement: GCControllerElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCControllerButtonInput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCControllerButtonInput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCControllerButtonInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCControllerButtonInput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCControllerButtonInput */

// A Boolean value that indicates whether the user is pressing the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/isPressed
func (g_ GCControllerButtonInput) Pressed() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("pressed"))
	return rv
}/* debug [instance_properties/getter]: pressed */


// A Boolean value that indicates whether the user is touching the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/isTouched
func (g_ GCControllerButtonInput) Touched() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("touched"))
	return rv
}/* debug [instance_properties/getter]: touched */


// The block that the element calls when the user presses or releases the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/pressedChangedHandler
func (g_ GCControllerButtonInput) PressedChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("pressedChangedHandler"))
	return rv
}/* debug [instance_properties/getter]: pressedChangedHandler */


// The block that the element calls when the user presses or releases the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/pressedChangedHandler
func (g_ GCControllerButtonInput) SetPressedChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPressedChangedHandler:"), value)
}/* debug [instance_properties/setter]: pressedChangedHandler */


// The block that the element calls when the user touches the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/touchedChangedHandler
func (g_ GCControllerButtonInput) TouchedChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("touchedChangedHandler"))
	return rv
}/* debug [instance_properties/getter]: touchedChangedHandler */


// The block that the element calls when the user touches the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/touchedChangedHandler
func (g_ GCControllerButtonInput) SetTouchedChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTouchedChangedHandler:"), value)
}/* debug [instance_properties/setter]: touchedChangedHandler */


// The level of pressure the user is applying to the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/value
func (g_ GCControllerButtonInput) Value() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The block that the element calls when the user changes the level of pressure on the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/valueChangedHandler
func (g_ GCControllerButtonInput) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}/* debug [instance_properties/getter]: valueChangedHandler */


// The block that the element calls when the user changes the level of pressure on the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/valueChangedHandler
func (g_ GCControllerButtonInput) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}/* debug [instance_properties/setter]: valueChangedHandler */


// A Boolean value that indicates whether the user is pressing the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerbuttoninput/ispressed
func (g_ GCControllerButtonInput) IsPressed() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isPressed"))
	return rv
}/* debug [instance_properties/getter]: isPressed */


// A Boolean value that indicates whether the user is pressing the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerbuttoninput/ispressed
func (g_ GCControllerButtonInput) SetIsPressed(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsPressed:"), value)
}/* debug [instance_properties/setter]: isPressed */


// A Boolean value that indicates whether the user is touching the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerbuttoninput/istouched
func (g_ GCControllerButtonInput) IsTouched() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isTouched"))
	return rv
}/* debug [instance_properties/getter]: isTouched */


// A Boolean value that indicates whether the user is touching the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerbuttoninput/istouched
func (g_ GCControllerButtonInput) SetIsTouched(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsTouched:"), value)
}/* debug [instance_properties/setter]: isTouched */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCControllerButtonInput */



