// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCDualShockGamepad */


/* debug [class_header]: Header for GCDualShockGamepad */
// The class instance for the [GCDualShockGamepad] class.
var (
	GCDualShockGamepadClass     _GCDualShockGamepadClass
	GCDualShockGamepadClassOnce sync.Once
)

func getGCDualShockGamepadClass() _GCDualShockGamepadClass {
	GCDualShockGamepadClassOnce.Do(func() {
		GCDualShockGamepadClass = _GCDualShockGamepadClass{objc.GetClass("GCDualShockGamepad")}
	})
	return GCDualShockGamepadClass
}

type _GCDualShockGamepadClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCDualShockGamepad */
// An interface definition for the [GCDualShockGamepad] class.
type IGCDualShockGamepad interface {
	IGCExtendedGamepad
	
/* debug [class_interface_properties]: Properties for GCDualShockGamepad */
	// properties:
	TouchpadButton() IGCControllerButtonInput
	TouchpadPrimary() IGCControllerDirectionPad
	TouchpadSecondary() IGCControllerDirectionPad
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

	
/* debug [class_interface_methods]: Methods for GCDualShockGamepad */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCDualShockGamepad */
// Alloc allocates a new instance without initialization.
func (gc _GCDualShockGamepadClass) Alloc() GCDualShockGamepad {
	rv := objc.Send[GCDualShockGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCDualShockGamepadClass) New() GCDualShockGamepad {
	rv := objc.Send[GCDualShockGamepad](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCDualShockGamepad) Init() GCDualShockGamepad {
	rv := objc.Send[GCDualShockGamepad](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCDualShockGamepad) Autorelease() GCDualShockGamepad {
	rv := objc.Send[GCDualShockGamepad](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDualShockGamepad creates a new GCDualShockGamepad instance.
func NewGCDualShockGamepad() GCDualShockGamepad {
	return getGCDualShockGamepadClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCDualShockGamepad */
// A controller profile that supports the DualShock 4 controller.
//
// The DualShock 4 controller profile is similar to an extended gamepad ( ), but has a touchpad with a button and two-finger tracking. This profile also supports motion — that is, the controller’s property is non-nil. If you hold the controller in front of you, the direction of the axes are: The positive x-axis points to your right. The positive y-axis points up. The positive z-axis starts at the touchpad and points to you.


// A controller profile that supports the DualShock 4 controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualShockGamepad
type GCDualShockGamepad struct {
	GCExtendedGamepad
}

// GCDualShockGamepadFrom constructs a [GCDualShockGamepad] from an unsafe.Pointer.
//
// A controller profile that supports the DualShock 4 controller.
func GCDualShockGamepadFrom(ptr unsafe.Pointer) GCDualShockGamepad {
	return GCDualShockGamepad{
		GCExtendedGamepad: GCExtendedGamepadFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCDualShockGamepad *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCDualShockGamepad */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCDualShockGamepad */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCDualShockGamepad */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCDualShockGamepad */

// The button element on the touchpad of the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualShockGamepad/touchpadButton
func (g_ GCDualShockGamepad) TouchpadButton() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("touchpadButton"))
	return rv
}/* debug [instance_properties/getter]: touchpadButton */


// The location of the player’s primary finger on the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualShockGamepad/touchpadPrimary
func (g_ GCDualShockGamepad) TouchpadPrimary() IGCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("touchpadPrimary"))
	return rv
}/* debug [instance_properties/getter]: touchpadPrimary */


// The location of the player’s secondary finger on the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualShockGamepad/touchpadSecondary
func (g_ GCDualShockGamepad) TouchpadSecondary() IGCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("touchpadSecondary"))
	return rv
}/* debug [instance_properties/getter]: touchpadSecondary */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCDualShockGamepad) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}/* debug [instance_properties/getter]: extendedGamepad */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCDualShockGamepad) SetExtendedGamepad(value IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}/* debug [instance_properties/setter]: extendedGamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCDualShockGamepad) Gamepad() IGCGamepad {
	rv := objc.Send[GCGamepad](g_.ID, objc.Sel("gamepad"))
	return rv
}/* debug [instance_properties/getter]: gamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCDualShockGamepad) SetGamepad(value IGCGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}/* debug [instance_properties/setter]: gamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCDualShockGamepad) MicroGamepad() IGCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}/* debug [instance_properties/getter]: microGamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCDualShockGamepad) SetMicroGamepad(value IGCMicroGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}/* debug [instance_properties/setter]: microGamepad */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCDualShockGamepad) Motion() IGCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}/* debug [instance_properties/getter]: motion */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCDualShockGamepad) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}/* debug [instance_properties/setter]: motion */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCDualShockGamepad) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}/* debug [instance_properties/getter]: physicalInputProfile */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCDualShockGamepad) SetPhysicalInputProfile(value IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}/* debug [instance_properties/setter]: physicalInputProfile */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCDualShockGamepad */



