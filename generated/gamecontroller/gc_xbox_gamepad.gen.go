// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCXboxGamepad */


/* debug [class_header]: Header for GCXboxGamepad */
// The class instance for the [GCXboxGamepad] class.
var (
	GCXboxGamepadClass     _GCXboxGamepadClass
	GCXboxGamepadClassOnce sync.Once
)

func getGCXboxGamepadClass() _GCXboxGamepadClass {
	GCXboxGamepadClassOnce.Do(func() {
		GCXboxGamepadClass = _GCXboxGamepadClass{objc.GetClass("GCXboxGamepad")}
	})
	return GCXboxGamepadClass
}

type _GCXboxGamepadClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCXboxGamepad */
// An interface definition for the [GCXboxGamepad] class.
type IGCXboxGamepad interface {
	IGCExtendedGamepad
	
/* debug [class_interface_properties]: Properties for GCXboxGamepad */
	// properties:
	ButtonShare() IGCControllerButtonInput
	PaddleButton1() IGCControllerButtonInput
	PaddleButton2() IGCControllerButtonInput
	PaddleButton3() IGCControllerButtonInput
	PaddleButton4() IGCControllerButtonInput
	ExtendedGamepad() IGCExtendedGamepad
	SetExtendedGamepad(value IGCExtendedGamepad)
	Gamepad() IGCGamepad
	SetGamepad(value IGCGamepad)
	MicroGamepad() IGCMicroGamepad
	SetMicroGamepad(value IGCMicroGamepad)
	Motion() IGCMotion
	SetMotion(value IGCMotion)
	PhysicalInputProfile() IGCPhysicalInputProfile
	SetPhysicalInputProfile(value IGCPhysicalInputProfile)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCXboxGamepad */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCXboxGamepad */
// Alloc allocates a new instance without initialization.
func (gc _GCXboxGamepadClass) Alloc() GCXboxGamepad {
	rv := objc.Send[GCXboxGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCXboxGamepadClass) New() GCXboxGamepad {
	rv := objc.Send[GCXboxGamepad](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCXboxGamepad) Init() GCXboxGamepad {
	rv := objc.Send[GCXboxGamepad](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCXboxGamepad) Autorelease() GCXboxGamepad {
	rv := objc.Send[GCXboxGamepad](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCXboxGamepad creates a new GCXboxGamepad instance.
func NewGCXboxGamepad() GCXboxGamepad {
	return getGCXboxGamepadClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCXboxGamepad */
// A controller profile that supports the Xbox controller.
//
// The Xbox controller profile is similar to an extended game pad ( ), but has four paddle button elements.


// A controller profile that supports the Xbox controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCXboxGamepad
type GCXboxGamepad struct {
	GCExtendedGamepad
}

// GCXboxGamepadFrom constructs a [GCXboxGamepad] from an unsafe.Pointer.
//
// A controller profile that supports the Xbox controller.
func GCXboxGamepadFrom(ptr unsafe.Pointer) GCXboxGamepad {
	return GCXboxGamepad{
		GCExtendedGamepad: GCExtendedGamepadFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCXboxGamepad *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCXboxGamepad */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCXboxGamepad */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCXboxGamepad */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCXboxGamepad */

// The share button on an Xbox Series X|S controller or later.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCXboxGamepad/buttonShare
func (g_ GCXboxGamepad) ButtonShare() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonShare"))
	return rv
}/* debug [instance_properties/getter]: buttonShare */


// The controller’s paddle 1 button element, which has a P1 label on the back of the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCXboxGamepad/paddleButton1
func (g_ GCXboxGamepad) PaddleButton1() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("paddleButton1"))
	return rv
}/* debug [instance_properties/getter]: paddleButton1 */


// The paddle 2 button element, which has a P2 label on the back of the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCXboxGamepad/paddleButton2
func (g_ GCXboxGamepad) PaddleButton2() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("paddleButton2"))
	return rv
}/* debug [instance_properties/getter]: paddleButton2 */


// The paddle 3 button element, which has a P3 label on the back of the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCXboxGamepad/paddleButton3
func (g_ GCXboxGamepad) PaddleButton3() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("paddleButton3"))
	return rv
}/* debug [instance_properties/getter]: paddleButton3 */


// The paddle 4 button element, which has a P4 label on the back of the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCXboxGamepad/paddleButton4
func (g_ GCXboxGamepad) PaddleButton4() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("paddleButton4"))
	return rv
}/* debug [instance_properties/getter]: paddleButton4 */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCXboxGamepad) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}/* debug [instance_properties/getter]: extendedGamepad */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCXboxGamepad) SetExtendedGamepad(value IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}/* debug [instance_properties/setter]: extendedGamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCXboxGamepad) Gamepad() IGCGamepad {
	rv := objc.Send[GCGamepad](g_.ID, objc.Sel("gamepad"))
	return rv
}/* debug [instance_properties/getter]: gamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCXboxGamepad) SetGamepad(value IGCGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}/* debug [instance_properties/setter]: gamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCXboxGamepad) MicroGamepad() IGCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}/* debug [instance_properties/getter]: microGamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCXboxGamepad) SetMicroGamepad(value IGCMicroGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}/* debug [instance_properties/setter]: microGamepad */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCXboxGamepad) Motion() IGCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}/* debug [instance_properties/getter]: motion */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCXboxGamepad) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}/* debug [instance_properties/setter]: motion */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCXboxGamepad) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}/* debug [instance_properties/getter]: physicalInputProfile */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCXboxGamepad) SetPhysicalInputProfile(value IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}/* debug [instance_properties/setter]: physicalInputProfile */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCXboxGamepad */





