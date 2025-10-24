// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCMicroGamepad */


/* debug [class_header]: Header for GCMicroGamepad */
// The class instance for the [GCMicroGamepad] class.
var (
	GCMicroGamepadClass     _GCMicroGamepadClass
	GCMicroGamepadClassOnce sync.Once
)

func getGCMicroGamepadClass() _GCMicroGamepadClass {
	GCMicroGamepadClassOnce.Do(func() {
		GCMicroGamepadClass = _GCMicroGamepadClass{objc.GetClass("GCMicroGamepad")}
	})
	return GCMicroGamepadClass
}

type _GCMicroGamepadClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCMicroGamepad */
// An interface definition for the [GCMicroGamepad] class.
type IGCMicroGamepad interface {
	IGCPhysicalInputProfile
	
/* debug [class_interface_properties]: Properties for GCMicroGamepad */
	// properties:
	AllowsRotation() bool
	SetAllowsRotation(value bool)
	ButtonA() IGCControllerButtonInput
	ButtonMenu() IGCControllerButtonInput
	ButtonX() IGCControllerButtonInput
	Controller() IGCController
	Dpad() IGCControllerDirectionPad
	ReportsAbsoluteDpadValues() bool
	SetReportsAbsoluteDpadValues(value bool)
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

	
/* debug [class_interface_methods]: Methods for GCMicroGamepad */
	// methods:
	SetStateFromMicroGamepad(microGamepad IGCMicroGamepad)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCMicroGamepad */
// Alloc allocates a new instance without initialization.
func (gc _GCMicroGamepadClass) Alloc() GCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCMicroGamepadClass) New() GCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCMicroGamepad) Init() GCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCMicroGamepad) Autorelease() GCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCMicroGamepad creates a new GCMicroGamepad instance.
func NewGCMicroGamepad() GCMicroGamepad {
	return getGCMicroGamepadClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCMicroGamepad */
// A controller profile that supports the Siri Remote.
//
// The micro gamepad controller profile supports the following input elements: Two digital face buttons (A and X). One analog directional pad (D-pad) that functions as a touchpad. Users can rotate game controllers that support the micro gamepad profile, switching them between landscape and portrait orientation. If you want to get directional values according to the orientation, set the property to .


// A controller profile that supports the Siri Remote.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad
type GCMicroGamepad struct {
	GCPhysicalInputProfile
}

// GCMicroGamepadFrom constructs a [GCMicroGamepad] from an unsafe.Pointer.
//
// A controller profile that supports the Siri Remote.
func GCMicroGamepadFrom(ptr unsafe.Pointer) GCMicroGamepad {
	return GCMicroGamepad{
		GCPhysicalInputProfile: GCPhysicalInputProfileFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCMicroGamepad *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCMicroGamepad */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCMicroGamepad */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCMicroGamepad */

// Copies the input values from a specified micro gamepad to a snapshot of a micro gamepad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/setStateFrom(_:)
func (g_ GCMicroGamepad) SetStateFromMicroGamepad(microGamepad IGCMicroGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStateFromMicroGamepad:"), microGamepad)
}/* debug [instance_methods/method]: SetStateFromMicroGamepad */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCMicroGamepad */

// A Boolean value that indicates whether the profile reports the directional pad values relative to its current orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/allowsRotation
func (g_ GCMicroGamepad) AllowsRotation() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("allowsRotation"))
	return rv
}/* debug [instance_properties/getter]: allowsRotation */


// A Boolean value that indicates whether the profile reports the directional pad values relative to its current orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/allowsRotation
func (g_ GCMicroGamepad) SetAllowsRotation(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAllowsRotation:"), value)
}/* debug [instance_properties/setter]: allowsRotation */


// The button that the user activates by pressing harder on the touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/buttonA
func (g_ GCMicroGamepad) ButtonA() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonA"))
	return rv
}/* debug [instance_properties/getter]: buttonA */


// The menu face button that players use to enter the main menu and pause the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/buttonMenu
func (g_ GCMicroGamepad) ButtonMenu() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonMenu"))
	return rv
}/* debug [instance_properties/getter]: buttonMenu */


// The second face button element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/buttonX
func (g_ GCMicroGamepad) ButtonX() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonX"))
	return rv
}/* debug [instance_properties/getter]: buttonX */


// The controller associated with this profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/controller
func (g_ GCMicroGamepad) Controller() IGCController {
	rv := objc.Send[GCController](g_.ID, objc.Sel("controller"))
	return rv
}/* debug [instance_properties/getter]: controller */


// The controller’s directional pad element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/dpad
func (g_ GCMicroGamepad) Dpad() IGCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("dpad"))
	return rv
}/* debug [instance_properties/getter]: dpad */


// A Boolean value that indicates whether the directional pad reports absolute or relative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/reportsAbsoluteDpadValues
func (g_ GCMicroGamepad) ReportsAbsoluteDpadValues() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("reportsAbsoluteDpadValues"))
	return rv
}/* debug [instance_properties/getter]: reportsAbsoluteDpadValues */


// A Boolean value that indicates whether the directional pad reports absolute or relative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/reportsAbsoluteDpadValues
func (g_ GCMicroGamepad) SetReportsAbsoluteDpadValues(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReportsAbsoluteDpadValues:"), value)
}/* debug [instance_properties/setter]: reportsAbsoluteDpadValues */


// The block that this profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/valueChangedHandler
func (g_ GCMicroGamepad) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}/* debug [instance_properties/getter]: valueChangedHandler */


// The block that this profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/valueChangedHandler
func (g_ GCMicroGamepad) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}/* debug [instance_properties/setter]: valueChangedHandler */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCMicroGamepad) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}/* debug [instance_properties/getter]: extendedGamepad */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCMicroGamepad) SetExtendedGamepad(value IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}/* debug [instance_properties/setter]: extendedGamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCMicroGamepad) Gamepad() IGCGamepad {
	rv := objc.Send[GCGamepad](g_.ID, objc.Sel("gamepad"))
	return rv
}/* debug [instance_properties/getter]: gamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCMicroGamepad) SetGamepad(value IGCGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}/* debug [instance_properties/setter]: gamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCMicroGamepad) MicroGamepad() IGCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}/* debug [instance_properties/getter]: microGamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCMicroGamepad) SetMicroGamepad(value IGCMicroGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}/* debug [instance_properties/setter]: microGamepad */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCMicroGamepad) Motion() IGCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}/* debug [instance_properties/getter]: motion */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCMicroGamepad) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}/* debug [instance_properties/setter]: motion */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCMicroGamepad) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}/* debug [instance_properties/getter]: physicalInputProfile */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCMicroGamepad) SetPhysicalInputProfile(value IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}/* debug [instance_properties/setter]: physicalInputProfile */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCMicroGamepad */



