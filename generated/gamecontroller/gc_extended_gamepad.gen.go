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
	SaveSnapshot() unsafe.Pointer
}

// A controller profile that supports the extended set of gamepad controls.
//
// The extended gamepad controller profile represents a physical or virtual controller with the following input elements: Two shoulder buttons Two trigger buttons Four face buttons in a diamond pattern One directional pad Two thumbsticks with optional thumbstick buttons Optional Home and Options buttons A Menu button If a object supports this type of profile, get the input values of the elements from the controller’s property or use the profile’s method to receive a callback when the input values change. Alternatively, use the method to capture the input values at a moment in time. If the controller’s property is , the controller doesn’t support this type of profile. See for other profiles you can use.
//
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


// Saves a snapshot of all of the profile’s elements.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/saveSnapshot()
func (g_ GCExtendedGamepad) SaveSnapshot() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("saveSnapshot"))
	return rv
}

// The main menu button element that players use to enter the secondary menu and pause the game.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonHome
func (g_ GCExtendedGamepad) ButtonHome() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("buttonHome"))
	return rv
}

// The top face button that uses or another indicator as its label.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonY
func (g_ GCExtendedGamepad) ButtonY() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("buttonY"))
	return rv
}

// The controller’s directional pad element.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/dpad
func (g_ GCExtendedGamepad) Dpad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("dpad"))
	return rv
}

// The block that the profile calls when an element’s value changes.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/valueChangedHandler
func (g_ GCExtendedGamepad) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}


// SetValueChangedHandler sets the value of the valueChangedHandler property.
// The block that the profile calls when an element’s value changes.

//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/valueChangedHandler
func (g_ GCExtendedGamepad) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}

// The extended gamepad profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCExtendedGamepad) ExtendedGamepad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}


// SetExtendedGamepad sets the value of the extendedGamepad property.
// The extended gamepad profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCExtendedGamepad) SetExtendedGamepad(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}

// The gamepad profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCExtendedGamepad) Gamepad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("gamepad"))
	return rv
}


// SetGamepad sets the value of the gamepad property.
// The gamepad profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCExtendedGamepad) SetGamepad(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}

// The micro gamepad profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCExtendedGamepad) MicroGamepad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("microGamepad"))
	return rv
}


// SetMicroGamepad sets the value of the microGamepad property.
// The micro gamepad profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCExtendedGamepad) SetMicroGamepad(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}

// The motion input profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCExtendedGamepad) Motion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("motion"))
	return rv
}


// SetMotion sets the value of the motion property.
// The motion input profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCExtendedGamepad) SetMotion(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}

// The physical input profile for the controller.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCExtendedGamepad) PhysicalInputProfile() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}


// SetPhysicalInputProfile sets the value of the physicalInputProfile property.
// The physical input profile for the controller.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCExtendedGamepad) SetPhysicalInputProfile(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}

// The bottom face button that uses
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttona
func (g_ GCExtendedGamepad) ButtonA() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("buttonA"))
	return rv
}


// SetButtonA sets the value of the buttonA property.
// The bottom face button that uses

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttona
func (g_ GCExtendedGamepad) SetButtonA(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setButtonA:"), value)
}

// The right face button that uses
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttonb
func (g_ GCExtendedGamepad) ButtonB() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("buttonB"))
	return rv
}


// SetButtonB sets the value of the buttonB property.
// The right face button that uses

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttonb
func (g_ GCExtendedGamepad) SetButtonB(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setButtonB:"), value)
}

// The primary menu button element that players use to enter the main menu and pause the game.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttonmenu
func (g_ GCExtendedGamepad) ButtonMenu() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("buttonMenu"))
	return rv
}


// SetButtonMenu sets the value of the buttonMenu property.
// The primary menu button element that players use to enter the main menu and pause the game.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttonmenu
func (g_ GCExtendedGamepad) SetButtonMenu(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setButtonMenu:"), value)
}

// The controller’s secondary menu button element.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttonoptions
func (g_ GCExtendedGamepad) ButtonOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("buttonOptions"))
	return rv
}


// SetButtonOptions sets the value of the buttonOptions property.
// The controller’s secondary menu button element.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttonoptions
func (g_ GCExtendedGamepad) SetButtonOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setButtonOptions:"), value)
}

// The left face button that uses
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttonx
func (g_ GCExtendedGamepad) ButtonX() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("buttonX"))
	return rv
}


// SetButtonX sets the value of the buttonX property.
// The left face button that uses

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/buttonx
func (g_ GCExtendedGamepad) SetButtonX(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setButtonX:"), value)
}

// The controller for the profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/controller
func (g_ GCExtendedGamepad) Controller() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("controller"))
	return rv
}


// SetController sets the value of the controller property.
// The controller for the profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/controller
func (g_ GCExtendedGamepad) SetController(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setController:"), value)
}

// The controller’s left shoulder button element.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/leftshoulder
func (g_ GCExtendedGamepad) LeftShoulder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("leftShoulder"))
	return rv
}


// SetLeftShoulder sets the value of the leftShoulder property.
// The controller’s left shoulder button element.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/leftshoulder
func (g_ GCExtendedGamepad) SetLeftShoulder(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeftShoulder:"), value)
}

// The controller’s left thumbstick element.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/leftthumbstick
func (g_ GCExtendedGamepad) LeftThumbstick() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("leftThumbstick"))
	return rv
}


// SetLeftThumbstick sets the value of the leftThumbstick property.
// The controller’s left thumbstick element.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/leftthumbstick
func (g_ GCExtendedGamepad) SetLeftThumbstick(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeftThumbstick:"), value)
}

// The button on the left thumbstick of the controller.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/leftthumbstickbutton
func (g_ GCExtendedGamepad) LeftThumbstickButton() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("leftThumbstickButton"))
	return rv
}


// SetLeftThumbstickButton sets the value of the leftThumbstickButton property.
// The button on the left thumbstick of the controller.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/leftthumbstickbutton
func (g_ GCExtendedGamepad) SetLeftThumbstickButton(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeftThumbstickButton:"), value)
}

// The controller’s left trigger element.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/lefttrigger
func (g_ GCExtendedGamepad) LeftTrigger() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("leftTrigger"))
	return rv
}


// SetLeftTrigger sets the value of the leftTrigger property.
// The controller’s left trigger element.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/lefttrigger
func (g_ GCExtendedGamepad) SetLeftTrigger(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeftTrigger:"), value)
}

// The controller’s right shoulder button element.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/rightshoulder
func (g_ GCExtendedGamepad) RightShoulder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("rightShoulder"))
	return rv
}


// SetRightShoulder sets the value of the rightShoulder property.
// The controller’s right shoulder button element.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/rightshoulder
func (g_ GCExtendedGamepad) SetRightShoulder(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRightShoulder:"), value)
}

// The controller’s right thumbstick element.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/rightthumbstick
func (g_ GCExtendedGamepad) RightThumbstick() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("rightThumbstick"))
	return rv
}


// SetRightThumbstick sets the value of the rightThumbstick property.
// The controller’s right thumbstick element.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/rightthumbstick
func (g_ GCExtendedGamepad) SetRightThumbstick(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRightThumbstick:"), value)
}

// The button on the right thumbstick of the controller.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/rightthumbstickbutton
func (g_ GCExtendedGamepad) RightThumbstickButton() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("rightThumbstickButton"))
	return rv
}


// SetRightThumbstickButton sets the value of the rightThumbstickButton property.
// The button on the right thumbstick of the controller.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/rightthumbstickbutton
func (g_ GCExtendedGamepad) SetRightThumbstickButton(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRightThumbstickButton:"), value)
}

// The controller’s right trigger element.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/righttrigger
func (g_ GCExtendedGamepad) RightTrigger() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("rightTrigger"))
	return rv
}


// SetRightTrigger sets the value of the rightTrigger property.
// The controller’s right trigger element.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/righttrigger
func (g_ GCExtendedGamepad) SetRightTrigger(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRightTrigger:"), value)
}



