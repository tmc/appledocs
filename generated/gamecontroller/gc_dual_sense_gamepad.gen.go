// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCDualSenseGamepad */


/* debug [class_header]: Header for GCDualSenseGamepad */
// The class instance for the [GCDualSenseGamepad] class.
var (
	GCDualSenseGamepadClass     _GCDualSenseGamepadClass
	GCDualSenseGamepadClassOnce sync.Once
)

func getGCDualSenseGamepadClass() _GCDualSenseGamepadClass {
	GCDualSenseGamepadClassOnce.Do(func() {
		GCDualSenseGamepadClass = _GCDualSenseGamepadClass{objc.GetClass("GCDualSenseGamepad")}
	})
	return GCDualSenseGamepadClass
}

type _GCDualSenseGamepadClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCDualSenseGamepad */
// An interface definition for the [GCDualSenseGamepad] class.
type IGCDualSenseGamepad interface {
	IGCExtendedGamepad
	
/* debug [class_interface_properties]: Properties for GCDualSenseGamepad */
	// properties:
	LeftTrigger() IGCDualSenseAdaptiveTrigger
	RightTrigger() IGCDualSenseAdaptiveTrigger
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

	
/* debug [class_interface_methods]: Methods for GCDualSenseGamepad */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCDualSenseGamepad */
// Alloc allocates a new instance without initialization.
func (gc _GCDualSenseGamepadClass) Alloc() GCDualSenseGamepad {
	rv := objc.Send[GCDualSenseGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCDualSenseGamepadClass) New() GCDualSenseGamepad {
	rv := objc.Send[GCDualSenseGamepad](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCDualSenseGamepad) Init() GCDualSenseGamepad {
	rv := objc.Send[GCDualSenseGamepad](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCDualSenseGamepad) Autorelease() GCDualSenseGamepad {
	rv := objc.Send[GCDualSenseGamepad](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDualSenseGamepad creates a new GCDualSenseGamepad instance.
func NewGCDualSenseGamepad() GCDualSenseGamepad {
	return getGCDualSenseGamepadClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCDualSenseGamepad */
// A controller profile that supported the DualSense controller.
//
// The DualSense controller profile is similar to a DualShock profile ( ), but has adaptive triggers that allow you to specify a dynamic resistance force when the user pulls the trigger. For example, you can emulate the feeling of pulling back a bow string, firing a weapon, or pulling a lever. This profile also supports motion — that is, the controller’s property is non-nil. If you hold the controller in front of you, the direction of the axes are: The positive x-axis points to your right. The positive y-axis points up out of the USB-C port. The positive z-axis starts at the touchpad and points to you.


// A controller profile that supported the DualSense controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseGamepad
type GCDualSenseGamepad struct {
	GCExtendedGamepad
}

// GCDualSenseGamepadFrom constructs a [GCDualSenseGamepad] from an unsafe.Pointer.
//
// A controller profile that supported the DualSense controller.
func GCDualSenseGamepadFrom(ptr unsafe.Pointer) GCDualSenseGamepad {
	return GCDualSenseGamepad{
		GCExtendedGamepad: GCExtendedGamepadFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCDualSenseGamepad *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCDualSenseGamepad */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCDualSenseGamepad */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCDualSenseGamepad */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCDualSenseGamepad */

// The controller’s left adaptive trigger element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseGamepad/leftTrigger
func (g_ GCDualSenseGamepad) LeftTrigger() IGCDualSenseAdaptiveTrigger {
	rv := objc.Send[GCDualSenseAdaptiveTrigger](g_.ID, objc.Sel("leftTrigger"))
	return rv
}/* debug [instance_properties/getter]: leftTrigger */


// The controller’s right adaptive trigger element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseGamepad/rightTrigger
func (g_ GCDualSenseGamepad) RightTrigger() IGCDualSenseAdaptiveTrigger {
	rv := objc.Send[GCDualSenseAdaptiveTrigger](g_.ID, objc.Sel("rightTrigger"))
	return rv
}/* debug [instance_properties/getter]: rightTrigger */


// The button element on the touchpad of the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseGamepad/touchpadButton
func (g_ GCDualSenseGamepad) TouchpadButton() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("touchpadButton"))
	return rv
}/* debug [instance_properties/getter]: touchpadButton */


// The location of the player’s primary finger on the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseGamepad/touchpadPrimary
func (g_ GCDualSenseGamepad) TouchpadPrimary() IGCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("touchpadPrimary"))
	return rv
}/* debug [instance_properties/getter]: touchpadPrimary */


// The location of the player’s secondary finger on the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseGamepad/touchpadSecondary
func (g_ GCDualSenseGamepad) TouchpadSecondary() IGCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("touchpadSecondary"))
	return rv
}/* debug [instance_properties/getter]: touchpadSecondary */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCDualSenseGamepad) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}/* debug [instance_properties/getter]: extendedGamepad */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCDualSenseGamepad) SetExtendedGamepad(value IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}/* debug [instance_properties/setter]: extendedGamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCDualSenseGamepad) Gamepad() IGCGamepad {
	rv := objc.Send[GCGamepad](g_.ID, objc.Sel("gamepad"))
	return rv
}/* debug [instance_properties/getter]: gamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCDualSenseGamepad) SetGamepad(value IGCGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}/* debug [instance_properties/setter]: gamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCDualSenseGamepad) MicroGamepad() IGCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}/* debug [instance_properties/getter]: microGamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCDualSenseGamepad) SetMicroGamepad(value IGCMicroGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}/* debug [instance_properties/setter]: microGamepad */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCDualSenseGamepad) Motion() IGCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}/* debug [instance_properties/getter]: motion */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCDualSenseGamepad) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}/* debug [instance_properties/setter]: motion */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCDualSenseGamepad) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}/* debug [instance_properties/getter]: physicalInputProfile */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCDualSenseGamepad) SetPhysicalInputProfile(value IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}/* debug [instance_properties/setter]: physicalInputProfile */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCDualSenseGamepad */



