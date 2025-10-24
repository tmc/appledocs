// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCMouseInput */


/* debug [class_header]: Header for GCMouseInput */
// The class instance for the [GCMouseInput] class.
var (
	GCMouseInputClass     _GCMouseInputClass
	GCMouseInputClassOnce sync.Once
)

func getGCMouseInputClass() _GCMouseInputClass {
	GCMouseInputClassOnce.Do(func() {
		GCMouseInputClass = _GCMouseInputClass{objc.GetClass("GCMouseInput")}
	})
	return GCMouseInputClass
}

type _GCMouseInputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCMouseInput */
// An interface definition for the [GCMouseInput] class.
type IGCMouseInput interface {
	IGCPhysicalInputProfile
	
/* debug [class_interface_properties]: Properties for GCMouseInput */
	// properties:
	AuxiliaryButtons() []GCControllerButtonInput
	LeftButton() IGCControllerButtonInput
	MiddleButton() IGCControllerButtonInput
	MouseMovedHandler() unsafe.Pointer
	SetMouseMovedHandler(value unsafe.Pointer)
	RightButton() IGCControllerButtonInput
	Scroll() IGCDeviceCursor
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

	
/* debug [class_interface_methods]: Methods for GCMouseInput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCMouseInput */
// Alloc allocates a new instance without initialization.
func (gc _GCMouseInputClass) Alloc() GCMouseInput {
	rv := objc.Send[GCMouseInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCMouseInputClass) New() GCMouseInput {
	rv := objc.Send[GCMouseInput](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCMouseInput) Init() GCMouseInput {
	rv := objc.Send[GCMouseInput](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCMouseInput) Autorelease() GCMouseInput {
	rv := objc.Send[GCMouseInput](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCMouseInput creates a new GCMouseInput instance.
func NewGCMouseInput() GCMouseInput {
	return getGCMouseInputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCMouseInput */
// A controller profile that tracks input from a mouse.
//
// This profile supports a mouse with the following features: A two-axis cursor and scroll A left button An optional right button An optional middle button An optional set of auxiliary buttons This profile provides only raw mouse movement delta values. For the cursor position at a specific time, use the class and the method.


// A controller profile that tracks input from a mouse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouseInput
type GCMouseInput struct {
	GCPhysicalInputProfile
}

// GCMouseInputFrom constructs a [GCMouseInput] from an unsafe.Pointer.
//
// A controller profile that tracks input from a mouse.
func GCMouseInputFrom(ptr unsafe.Pointer) GCMouseInput {
	return GCMouseInput{
		GCPhysicalInputProfile: GCPhysicalInputProfileFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCMouseInput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCMouseInput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCMouseInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCMouseInput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCMouseInput */

// The optional additional buttons on the mouse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouseInput/auxiliaryButtons
func (g_ GCMouseInput) AuxiliaryButtons() []GCControllerButtonInput {
	rv := objc.Send[[]GCControllerButtonInput](g_.ID, objc.Sel("auxiliaryButtons"))
	return rv
}/* debug [instance_properties/getter]: auxiliaryButtons */


// The left button on the mouse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouseInput/leftButton
func (g_ GCMouseInput) LeftButton() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("leftButton"))
	return rv
}/* debug [instance_properties/getter]: leftButton */


// The optional middle button on the mouse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouseInput/middleButton
func (g_ GCMouseInput) MiddleButton() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("middleButton"))
	return rv
}/* debug [instance_properties/getter]: middleButton */


// The block that the profile calls when the mouse moves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouseInput/mouseMovedHandler
func (g_ GCMouseInput) MouseMovedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("mouseMovedHandler"))
	return rv
}/* debug [instance_properties/getter]: mouseMovedHandler */


// The block that the profile calls when the mouse moves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouseInput/mouseMovedHandler
func (g_ GCMouseInput) SetMouseMovedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMouseMovedHandler:"), value)
}/* debug [instance_properties/setter]: mouseMovedHandler */


// The optional right button on the mouse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouseInput/rightButton
func (g_ GCMouseInput) RightButton() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("rightButton"))
	return rv
}/* debug [instance_properties/getter]: rightButton */


// The location of the directional pad cursor with an undefined range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouseInput/scroll
func (g_ GCMouseInput) Scroll() IGCDeviceCursor {
	rv := objc.Send[GCDeviceCursor](g_.ID, objc.Sel("scroll"))
	return rv
}/* debug [instance_properties/getter]: scroll */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCMouseInput) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}/* debug [instance_properties/getter]: extendedGamepad */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCMouseInput) SetExtendedGamepad(value IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}/* debug [instance_properties/setter]: extendedGamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCMouseInput) Gamepad() IGCGamepad {
	rv := objc.Send[GCGamepad](g_.ID, objc.Sel("gamepad"))
	return rv
}/* debug [instance_properties/getter]: gamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCMouseInput) SetGamepad(value IGCGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}/* debug [instance_properties/setter]: gamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCMouseInput) MicroGamepad() IGCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}/* debug [instance_properties/getter]: microGamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCMouseInput) SetMicroGamepad(value IGCMicroGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}/* debug [instance_properties/setter]: microGamepad */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCMouseInput) Motion() IGCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}/* debug [instance_properties/getter]: motion */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCMouseInput) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}/* debug [instance_properties/setter]: motion */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCMouseInput) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}/* debug [instance_properties/getter]: physicalInputProfile */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCMouseInput) SetPhysicalInputProfile(value IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}/* debug [instance_properties/setter]: physicalInputProfile */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCMouseInput */



