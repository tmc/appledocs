// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCKeyboardInput */


/* debug [class_header]: Header for GCKeyboardInput */
// The class instance for the [GCKeyboardInput] class.
var (
	GCKeyboardInputClass     _GCKeyboardInputClass
	GCKeyboardInputClassOnce sync.Once
)

func getGCKeyboardInputClass() _GCKeyboardInputClass {
	GCKeyboardInputClassOnce.Do(func() {
		GCKeyboardInputClass = _GCKeyboardInputClass{objc.GetClass("GCKeyboardInput")}
	})
	return GCKeyboardInputClass
}

type _GCKeyboardInputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCKeyboardInput */
// An interface definition for the [GCKeyboardInput] class.
type IGCKeyboardInput interface {
	IGCPhysicalInputProfile
	
/* debug [class_interface_properties]: Properties for GCKeyboardInput */
	// properties:
	AnyKeyPressed() bool
	KeyChangedHandler() unsafe.Pointer
	SetKeyChangedHandler(value unsafe.Pointer)
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
	IsAnyKeyPressed() bool
	SetIsAnyKeyPressed(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCKeyboardInput */
	// methods:
	ButtonForKeyCode(code GCKeyCode /* typedef */) IGCControllerButtonInput
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCKeyboardInput */
// Alloc allocates a new instance without initialization.
func (gc _GCKeyboardInputClass) Alloc() GCKeyboardInput {
	rv := objc.Send[GCKeyboardInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCKeyboardInputClass) New() GCKeyboardInput {
	rv := objc.Send[GCKeyboardInput](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCKeyboardInput) Init() GCKeyboardInput {
	rv := objc.Send[GCKeyboardInput](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCKeyboardInput) Autorelease() GCKeyboardInput {
	rv := objc.Send[GCKeyboardInput](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCKeyboardInput creates a new GCKeyboardInput instance.
func NewGCKeyboardInput() GCKeyboardInput {
	return getGCKeyboardInputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCKeyboardInput */
// A controller profile that uses the keyboard as the input device.
//
// Use this profile to get the state of the keyboard buttons that the structure defines.


// A controller profile that uses the keyboard as the input device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCKeyboardInput
type GCKeyboardInput struct {
	GCPhysicalInputProfile
}

// GCKeyboardInputFrom constructs a [GCKeyboardInput] from an unsafe.Pointer.
//
// A controller profile that uses the keyboard as the input device.
func GCKeyboardInputFrom(ptr unsafe.Pointer) GCKeyboardInput {
	return GCKeyboardInput{
		GCPhysicalInputProfile: GCPhysicalInputProfileFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCKeyboardInput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCKeyboardInput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCKeyboardInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCKeyboardInput */

// Returns the button element for the specified key code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCKeyboardInput/button(forKeyCode:)
func (g_ GCKeyboardInput) ButtonForKeyCode(code GCKeyCode /* typedef */) IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonForKeyCode:"), code)
	return rv
}/* debug [instance_methods/method]: ButtonForKeyCode */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCKeyboardInput */

// A Boolean value that indicates whether the user is pressing any of the keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCKeyboardInput/isAnyKeyPressed
func (g_ GCKeyboardInput) AnyKeyPressed() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("anyKeyPressed"))
	return rv
}/* debug [instance_properties/getter]: anyKeyPressed */


// The block that the profile calls when the user presses a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCKeyboardInput/keyChangedHandler
func (g_ GCKeyboardInput) KeyChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("keyChangedHandler"))
	return rv
}/* debug [instance_properties/getter]: keyChangedHandler */


// The block that the profile calls when the user presses a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCKeyboardInput/keyChangedHandler
func (g_ GCKeyboardInput) SetKeyChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKeyChangedHandler:"), value)
}/* debug [instance_properties/setter]: keyChangedHandler */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCKeyboardInput) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}/* debug [instance_properties/getter]: extendedGamepad */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCKeyboardInput) SetExtendedGamepad(value IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}/* debug [instance_properties/setter]: extendedGamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCKeyboardInput) Gamepad() IGCGamepad {
	rv := objc.Send[GCGamepad](g_.ID, objc.Sel("gamepad"))
	return rv
}/* debug [instance_properties/getter]: gamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCKeyboardInput) SetGamepad(value IGCGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}/* debug [instance_properties/setter]: gamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCKeyboardInput) MicroGamepad() IGCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}/* debug [instance_properties/getter]: microGamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCKeyboardInput) SetMicroGamepad(value IGCMicroGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}/* debug [instance_properties/setter]: microGamepad */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCKeyboardInput) Motion() IGCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}/* debug [instance_properties/getter]: motion */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCKeyboardInput) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}/* debug [instance_properties/setter]: motion */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCKeyboardInput) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}/* debug [instance_properties/getter]: physicalInputProfile */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCKeyboardInput) SetPhysicalInputProfile(value IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}/* debug [instance_properties/setter]: physicalInputProfile */


// A Boolean value that indicates whether the user is pressing any of the keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gckeyboardinput/isanykeypressed
func (g_ GCKeyboardInput) IsAnyKeyPressed() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isAnyKeyPressed"))
	return rv
}/* debug [instance_properties/getter]: isAnyKeyPressed */


// A Boolean value that indicates whether the user is pressing any of the keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gckeyboardinput/isanykeypressed
func (g_ GCKeyboardInput) SetIsAnyKeyPressed(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsAnyKeyPressed:"), value)
}/* debug [instance_properties/setter]: isAnyKeyPressed */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCKeyboardInput */



