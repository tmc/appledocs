// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [GCDualShockGamepad] class.
type IGCDualShockGamepad interface {
	IGCExtendedGamepad
	ExtendedGamepad() GCExtendedGamepad
	SetExtendedGamepad(value IGCExtendedGamepad)
	Gamepad() unsafe.Pointer
	SetGamepad(value unsafe.Pointer)
	MicroGamepad() GCMicroGamepad
	SetMicroGamepad(value IGCMicroGamepad)
	Motion() GCMotion
	SetMotion(value IGCMotion)
	PhysicalInputProfile() GCPhysicalInputProfile
	SetPhysicalInputProfile(value IGCPhysicalInputProfile)
	TouchpadButton() GCControllerButtonInput
	SetTouchpadButton(value IGCControllerButtonInput)
	TouchpadPrimary() GCControllerDirectionPad
	SetTouchpadPrimary(value IGCControllerDirectionPad)
	TouchpadSecondary() GCControllerDirectionPad
	SetTouchpadSecondary(value IGCControllerDirectionPad)
}

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

// Alloc allocates a new instance without initialization.
func (gc _GCDualShockGamepadClass) Alloc() GCDualShockGamepad {
	rv := objc.Send[GCDualShockGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCDualShockGamepad) ExtendedGamepad() GCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCDualShockGamepad) SetExtendedGamepad(value IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCDualShockGamepad) Gamepad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("gamepad"))
	return rv
}


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCDualShockGamepad) SetGamepad(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCDualShockGamepad) MicroGamepad() GCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCDualShockGamepad) SetMicroGamepad(value IGCMicroGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCDualShockGamepad) Motion() GCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCDualShockGamepad) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCDualShockGamepad) PhysicalInputProfile() GCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCDualShockGamepad) SetPhysicalInputProfile(value IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}


// The button element on the touchpad of the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdualshockgamepad/touchpadbutton
func (g_ GCDualShockGamepad) TouchpadButton() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("touchpadButton"))
	return rv
}


// The button element on the touchpad of the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdualshockgamepad/touchpadbutton
func (g_ GCDualShockGamepad) SetTouchpadButton(value IGCControllerButtonInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTouchpadButton:"), value)
}


// The location of the player’s primary finger on the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdualshockgamepad/touchpadprimary
func (g_ GCDualShockGamepad) TouchpadPrimary() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("touchpadPrimary"))
	return rv
}


// The location of the player’s primary finger on the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdualshockgamepad/touchpadprimary
func (g_ GCDualShockGamepad) SetTouchpadPrimary(value IGCControllerDirectionPad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTouchpadPrimary:"), value)
}


// The location of the player’s secondary finger on the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdualshockgamepad/touchpadsecondary
func (g_ GCDualShockGamepad) TouchpadSecondary() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("touchpadSecondary"))
	return rv
}


// The location of the player’s secondary finger on the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdualshockgamepad/touchpadsecondary
func (g_ GCDualShockGamepad) SetTouchpadSecondary(value IGCControllerDirectionPad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTouchpadSecondary:"), value)
}



