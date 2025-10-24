// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCExtendedGamepad */


/* debug [class_header]: Header for GCExtendedGamepad */
// The class instance for the [GCExtendedGamepad] class.
var (
	GCExtendedGamepadClass     _GCExtendedGamepadClass
	GCExtendedGamepadClassOnce sync.Once
)

func getGCExtendedGamepadClass() _GCExtendedGamepadClass {
	GCExtendedGamepadClassOnce.Do(func() {
		GCExtendedGamepadClass = _GCExtendedGamepadClass{objc.GetClass("GCExtendedGamepad")}
	})
	return GCExtendedGamepadClass
}

type _GCExtendedGamepadClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCExtendedGamepad */
// An interface definition for the [GCExtendedGamepad] class.
type IGCExtendedGamepad interface {
	IGCPhysicalInputProfile
	
/* debug [class_interface_properties]: Properties for GCExtendedGamepad */
	// properties:
	ButtonA() IGCControllerButtonInput
	ButtonB() IGCControllerButtonInput
	ButtonHome() IGCControllerButtonInput
	ButtonMenu() IGCControllerButtonInput
	ButtonOptions() IGCControllerButtonInput
	ButtonX() IGCControllerButtonInput
	ButtonY() IGCControllerButtonInput
	Controller() IGCController
	Dpad() IGCControllerDirectionPad
	LeftShoulder() IGCControllerButtonInput
	LeftThumbstick() IGCControllerDirectionPad
	LeftThumbstickButton() IGCControllerButtonInput
	LeftTrigger() IGCControllerButtonInput
	RightShoulder() IGCControllerButtonInput
	RightThumbstick() IGCControllerDirectionPad
	RightThumbstickButton() IGCControllerButtonInput
	RightTrigger() IGCControllerButtonInput
	ValueChangedHandler() unsafe.Pointer
	SetValueChangedHandler(value unsafe.Pointer)
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

	
/* debug [class_interface_methods]: Methods for GCExtendedGamepad */
	// methods:
	SetStateFromExtendedGamepad(extendedGamepad IGCExtendedGamepad)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCExtendedGamepad */
// Alloc allocates a new instance without initialization.
func (gc _GCExtendedGamepadClass) Alloc() GCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCExtendedGamepadClass) New() GCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCExtendedGamepad) Init() GCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCExtendedGamepad) Autorelease() GCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCExtendedGamepad creates a new GCExtendedGamepad instance.
func NewGCExtendedGamepad() GCExtendedGamepad {
	return getGCExtendedGamepadClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCExtendedGamepad */
// A controller profile that supports the extended set of gamepad controls.
//
// The extended gamepad controller profile represents a physical or virtual controller with the following input elements: Two shoulder buttons Two trigger buttons Four face buttons in a diamond pattern One directional pad Two thumbsticks with optional thumbstick buttons Optional Home and Options buttons A Menu button If a object supports this type of profile, get the input values of the elements from the controller’s property or use the profile’s method to receive a callback when the input values change. Alternatively, use the method to capture the input values at a moment in time. If the controller’s property is , the controller doesn’t support this type of profile. See for other profiles you can use.


// A controller profile that supports the extended set of gamepad controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad
type GCExtendedGamepad struct {
	GCPhysicalInputProfile
}

// GCExtendedGamepadFrom constructs a [GCExtendedGamepad] from an unsafe.Pointer.
//
// A controller profile that supports the extended set of gamepad controls.
func GCExtendedGamepadFrom(ptr unsafe.Pointer) GCExtendedGamepad {
	return GCExtendedGamepad{
		GCPhysicalInputProfile: GCPhysicalInputProfileFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCExtendedGamepad *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCExtendedGamepad */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCExtendedGamepad */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCExtendedGamepad */

// Copies the input values from a specified extended gamepad to a snapshot of an extended gamepad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/setStateFrom(_:)
func (g_ GCExtendedGamepad) SetStateFromExtendedGamepad(extendedGamepad IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStateFromExtendedGamepad:"), extendedGamepad)
}/* debug [instance_methods/method]: SetStateFromExtendedGamepad */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCExtendedGamepad */

// The bottom face button that uses or another indicator as its label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonA
func (g_ GCExtendedGamepad) ButtonA() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonA"))
	return rv
}/* debug [instance_properties/getter]: buttonA */


// The right face button that uses or another indicator as its label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonB
func (g_ GCExtendedGamepad) ButtonB() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonB"))
	return rv
}/* debug [instance_properties/getter]: buttonB */


// The main menu button element that players use to enter the secondary menu and pause the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonHome
func (g_ GCExtendedGamepad) ButtonHome() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonHome"))
	return rv
}/* debug [instance_properties/getter]: buttonHome */


// The primary menu button element that players use to enter the main menu and pause the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonMenu
func (g_ GCExtendedGamepad) ButtonMenu() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonMenu"))
	return rv
}/* debug [instance_properties/getter]: buttonMenu */


// The controller’s secondary menu button element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonOptions
func (g_ GCExtendedGamepad) ButtonOptions() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonOptions"))
	return rv
}/* debug [instance_properties/getter]: buttonOptions */


// The left face button that uses or another indicator as its label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonX
func (g_ GCExtendedGamepad) ButtonX() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonX"))
	return rv
}/* debug [instance_properties/getter]: buttonX */


// The top face button that uses or another indicator as its label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonY
func (g_ GCExtendedGamepad) ButtonY() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonY"))
	return rv
}/* debug [instance_properties/getter]: buttonY */


// The controller for the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/controller
func (g_ GCExtendedGamepad) Controller() IGCController {
	rv := objc.Send[GCController](g_.ID, objc.Sel("controller"))
	return rv
}/* debug [instance_properties/getter]: controller */


// The controller’s directional pad element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/dpad
func (g_ GCExtendedGamepad) Dpad() IGCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("dpad"))
	return rv
}/* debug [instance_properties/getter]: dpad */


// The controller’s left shoulder button element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/leftShoulder
func (g_ GCExtendedGamepad) LeftShoulder() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("leftShoulder"))
	return rv
}/* debug [instance_properties/getter]: leftShoulder */


// The controller’s left thumbstick element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/leftThumbstick
func (g_ GCExtendedGamepad) LeftThumbstick() IGCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("leftThumbstick"))
	return rv
}/* debug [instance_properties/getter]: leftThumbstick */


// The button on the left thumbstick of the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/leftThumbstickButton
func (g_ GCExtendedGamepad) LeftThumbstickButton() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("leftThumbstickButton"))
	return rv
}/* debug [instance_properties/getter]: leftThumbstickButton */


// The controller’s left trigger element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/leftTrigger
func (g_ GCExtendedGamepad) LeftTrigger() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("leftTrigger"))
	return rv
}/* debug [instance_properties/getter]: leftTrigger */


// The controller’s right shoulder button element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/rightShoulder
func (g_ GCExtendedGamepad) RightShoulder() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("rightShoulder"))
	return rv
}/* debug [instance_properties/getter]: rightShoulder */


// The controller’s right thumbstick element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/rightThumbstick
func (g_ GCExtendedGamepad) RightThumbstick() IGCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("rightThumbstick"))
	return rv
}/* debug [instance_properties/getter]: rightThumbstick */


// The button on the right thumbstick of the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/rightThumbstickButton
func (g_ GCExtendedGamepad) RightThumbstickButton() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("rightThumbstickButton"))
	return rv
}/* debug [instance_properties/getter]: rightThumbstickButton */


// The controller’s right trigger element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/rightTrigger
func (g_ GCExtendedGamepad) RightTrigger() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("rightTrigger"))
	return rv
}/* debug [instance_properties/getter]: rightTrigger */


// The block that the profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/valueChangedHandler
func (g_ GCExtendedGamepad) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}/* debug [instance_properties/getter]: valueChangedHandler */


// The block that the profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/valueChangedHandler
func (g_ GCExtendedGamepad) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}/* debug [instance_properties/setter]: valueChangedHandler */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCExtendedGamepad) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}/* debug [instance_properties/getter]: extendedGamepad */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCExtendedGamepad) SetExtendedGamepad(value IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}/* debug [instance_properties/setter]: extendedGamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCExtendedGamepad) Gamepad() IGCGamepad {
	rv := objc.Send[GCGamepad](g_.ID, objc.Sel("gamepad"))
	return rv
}/* debug [instance_properties/getter]: gamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCExtendedGamepad) SetGamepad(value IGCGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}/* debug [instance_properties/setter]: gamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCExtendedGamepad) MicroGamepad() IGCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}/* debug [instance_properties/getter]: microGamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCExtendedGamepad) SetMicroGamepad(value IGCMicroGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}/* debug [instance_properties/setter]: microGamepad */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCExtendedGamepad) Motion() IGCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}/* debug [instance_properties/getter]: motion */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCExtendedGamepad) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}/* debug [instance_properties/setter]: motion */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCExtendedGamepad) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}/* debug [instance_properties/getter]: physicalInputProfile */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCExtendedGamepad) SetPhysicalInputProfile(value IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}/* debug [instance_properties/setter]: physicalInputProfile */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCExtendedGamepad */



