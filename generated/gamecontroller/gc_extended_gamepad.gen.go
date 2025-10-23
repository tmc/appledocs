// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [GCExtendedGamepad] class.
type IGCExtendedGamepad interface {
	IGCPhysicalInputProfile
	ButtonB() GCControllerButtonInput
	ButtonHome() GCControllerButtonInput
	ButtonOptions() GCControllerButtonInput
	Dpad() GCControllerDirectionPad
	LeftThumbstick() GCControllerDirectionPad
	LeftThumbstickButton() GCControllerButtonInput
	LeftTrigger() GCControllerButtonInput
	RightShoulder() GCControllerButtonInput
	RightThumbstick() GCControllerDirectionPad
	RightThumbstickButton() GCControllerButtonInput
	RightTrigger() GCControllerButtonInput
	ValueChangedHandler() unsafe.Pointer
	SetValueChangedHandler(value unsafe.Pointer)
	ExtendedGamepad() IGCExtendedGamepad
	SetExtendedGamepad(value IGCExtendedGamepad)
	Gamepad() unsafe.Pointer
	SetGamepad(value unsafe.Pointer)
	MicroGamepad() GCMicroGamepad
	SetMicroGamepad(value GCMicroGamepad)
	Motion() IGCMotion
	SetMotion(value IGCMotion)
	PhysicalInputProfile() IGCPhysicalInputProfile
	SetPhysicalInputProfile(value IGCPhysicalInputProfile)
	ButtonA() GCControllerButtonInput
	SetButtonA(value GCControllerButtonInput)
	ButtonMenu() GCControllerButtonInput
	SetButtonMenu(value GCControllerButtonInput)
	ButtonX() GCControllerButtonInput
	SetButtonX(value GCControllerButtonInput)
	ButtonY() GCControllerButtonInput
	SetButtonY(value GCControllerButtonInput)
	Controller() IGCController
	SetController(value IGCController)
	LeftShoulder() GCControllerButtonInput
	SetLeftShoulder(value GCControllerButtonInput)
	SetStateFromExtendedGamepad(extendedGamepad IGCExtendedGamepad)
}

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

// Alloc allocates a new instance without initialization.
func (gc _GCExtendedGamepadClass) Alloc() GCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Copies the input values from a specified extended gamepad to a snapshot of an extended gamepad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/setStateFrom(_:)
func (g_ GCExtendedGamepad) SetStateFromExtendedGamepad(extendedGamepad IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStateFromExtendedGamepad:"), extendedGamepad)
}


// The right face button that uses or another indicator as its label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonB
func (g_ GCExtendedGamepad) ButtonB() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonB"))
	return rv
}


// The main menu button element that players use to enter the secondary menu and pause the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonHome
func (g_ GCExtendedGamepad) ButtonHome() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonHome"))
	return rv
}


// The controller’s secondary menu button element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonOptions
func (g_ GCExtendedGamepad) ButtonOptions() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonOptions"))
	return rv
}


// The controller’s directional pad element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/dpad
func (g_ GCExtendedGamepad) Dpad() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("dpad"))
	return rv
}


// The controller’s left thumbstick element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/leftThumbstick
func (g_ GCExtendedGamepad) LeftThumbstick() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("leftThumbstick"))
	return rv
}


// The button on the left thumbstick of the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/leftThumbstickButton
func (g_ GCExtendedGamepad) LeftThumbstickButton() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("leftThumbstickButton"))
	return rv
}


// The controller’s left trigger element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/leftTrigger
func (g_ GCExtendedGamepad) LeftTrigger() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("leftTrigger"))
	return rv
}


// The controller’s right shoulder button element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/rightShoulder
func (g_ GCExtendedGamepad) RightShoulder() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("rightShoulder"))
	return rv
}


// The controller’s right thumbstick element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/rightThumbstick
func (g_ GCExtendedGamepad) RightThumbstick() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("rightThumbstick"))
	return rv
}


// The button on the right thumbstick of the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/rightThumbstickButton
func (g_ GCExtendedGamepad) RightThumbstickButton() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("rightThumbstickButton"))
	return rv
}


// The controller’s right trigger element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/rightTrigger
func (g_ GCExtendedGamepad) RightTrigger() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("rightTrigger"))
	return rv
}


// The block that the profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/valueChangedHandler
func (g_ GCExtendedGamepad) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}


// The block that the profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/valueChangedHandler
func (g_ GCExtendedGamepad) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCExtendedGamepad) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCExtendedGamepad) SetExtendedGamepad(value IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCExtendedGamepad) Gamepad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("gamepad"))
	return rv
}


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCExtendedGamepad) SetGamepad(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCExtendedGamepad) MicroGamepad() GCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCExtendedGamepad) SetMicroGamepad(value GCMicroGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCExtendedGamepad) Motion() IGCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCExtendedGamepad) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCExtendedGamepad) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCExtendedGamepad) SetPhysicalInputProfile(value IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}


// The bottom face button that uses
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttona
func (g_ GCExtendedGamepad) ButtonA() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonA"))
	return rv
}


// The bottom face button that uses
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttona
func (g_ GCExtendedGamepad) SetButtonA(value GCControllerButtonInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setButtonA:"), value)
}


// The primary menu button element that players use to enter the main menu and pause the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttonmenu
func (g_ GCExtendedGamepad) ButtonMenu() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonMenu"))
	return rv
}


// The primary menu button element that players use to enter the main menu and pause the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttonmenu
func (g_ GCExtendedGamepad) SetButtonMenu(value GCControllerButtonInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setButtonMenu:"), value)
}


// The left face button that uses
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttonx
func (g_ GCExtendedGamepad) ButtonX() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonX"))
	return rv
}


// The left face button that uses
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttonx
func (g_ GCExtendedGamepad) SetButtonX(value GCControllerButtonInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setButtonX:"), value)
}


// The top face button that uses
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttony
func (g_ GCExtendedGamepad) ButtonY() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonY"))
	return rv
}


// The top face button that uses
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttony
func (g_ GCExtendedGamepad) SetButtonY(value GCControllerButtonInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setButtonY:"), value)
}


// The controller for the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/controller
func (g_ GCExtendedGamepad) Controller() IGCController {
	rv := objc.Send[GCController](g_.ID, objc.Sel("controller"))
	return rv
}


// The controller for the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/controller
func (g_ GCExtendedGamepad) SetController(value IGCController) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setController:"), value)
}


// The controller’s left shoulder button element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/leftshoulder
func (g_ GCExtendedGamepad) LeftShoulder() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("leftShoulder"))
	return rv
}


// The controller’s left shoulder button element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/leftshoulder
func (g_ GCExtendedGamepad) SetLeftShoulder(value GCControllerButtonInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeftShoulder:"), value)
}



