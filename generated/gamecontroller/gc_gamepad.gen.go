// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCGamepad */


/* debug [class_header]: Header for GCGamepad */
// The class instance for the [GCGamepad] class.
var (
	GCGamepadClass     _GCGamepadClass
	GCGamepadClassOnce sync.Once
)

func getGCGamepadClass() _GCGamepadClass {
	GCGamepadClassOnce.Do(func() {
		GCGamepadClass = _GCGamepadClass{objc.GetClass("GCGamepad")}
	})
	return GCGamepadClass
}

type _GCGamepadClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCGamepad */
// An interface definition for the [GCGamepad] class.
type IGCGamepad interface {
	IGCPhysicalInputProfile
	
/* debug [class_interface_properties]: Properties for GCGamepad */
	// properties:
	ButtonA() IGCControllerButtonInput
	ButtonB() IGCControllerButtonInput
	ButtonX() IGCControllerButtonInput
	ButtonY() IGCControllerButtonInput
	Controller() IGCController
	Dpad() IGCControllerDirectionPad
	LeftShoulder() IGCControllerButtonInput
	RightShoulder() IGCControllerButtonInput
	ValueChangedHandler() unsafe.Pointer
	SetValueChangedHandler(value unsafe.Pointer)
	GCCurrentExtendedGamepadSnapshotDataVersion() GCExtendedGamepadSnapshotDataVersion
	GCCurrentMicroGamepadSnapshotDataVersion() GCMicroGamepadSnapshotDataVersion
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCGamepad */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCGamepad */
// Alloc allocates a new instance without initialization.
func (gc _GCGamepadClass) Alloc() GCGamepad {
	rv := objc.Send[GCGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCGamepadClass) New() GCGamepad {
	rv := objc.Send[GCGamepad](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCGamepad) Init() GCGamepad {
	rv := objc.Send[GCGamepad](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCGamepad) Autorelease() GCGamepad {
	rv := objc.Send[GCGamepad](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCGamepad creates a new GCGamepad instance.
func NewGCGamepad() GCGamepad {
	return getGCGamepadClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCGamepad */
// The standard set of gamepad controls.
//
// The controls associated with the gamepad profile include the following: Two shoulder buttons. Four face buttons arranged in a diamond pattern. One directional pad (D-pad).


// The standard set of gamepad controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepad
type GCGamepad struct {
	GCPhysicalInputProfile
}

// GCGamepadFrom constructs a [GCGamepad] from an unsafe.Pointer.
//
// The standard set of gamepad controls.
func GCGamepadFrom(ptr unsafe.Pointer) GCGamepad {
	return GCGamepad{
		GCPhysicalInputProfile: GCPhysicalInputProfileFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCGamepad *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCGamepad */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCGamepad */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCGamepad */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCGamepad */

// The bottom face button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepad/buttonA
func (g_ GCGamepad) ButtonA() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonA"))
	return rv
}/* debug [instance_properties/getter]: buttonA */


// The right face button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepad/buttonB
func (g_ GCGamepad) ButtonB() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonB"))
	return rv
}/* debug [instance_properties/getter]: buttonB */


// The left face button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepad/buttonX
func (g_ GCGamepad) ButtonX() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonX"))
	return rv
}/* debug [instance_properties/getter]: buttonX */


// The top face button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepad/buttonY
func (g_ GCGamepad) ButtonY() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("buttonY"))
	return rv
}/* debug [instance_properties/getter]: buttonY */


// The controller this profile is associated with.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepad/controller
func (g_ GCGamepad) Controller() IGCController {
	rv := objc.Send[GCController](g_.ID, objc.Sel("controller"))
	return rv
}/* debug [instance_properties/getter]: controller */


// The D-pad element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepad/dpad
func (g_ GCGamepad) Dpad() IGCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("dpad"))
	return rv
}/* debug [instance_properties/getter]: dpad */


// The left shoulder button element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepad/leftShoulder
func (g_ GCGamepad) LeftShoulder() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("leftShoulder"))
	return rv
}/* debug [instance_properties/getter]: leftShoulder */


// The right shoulder button element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepad/rightShoulder
func (g_ GCGamepad) RightShoulder() IGCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("rightShoulder"))
	return rv
}/* debug [instance_properties/getter]: rightShoulder */


// A block called when any element in the profile changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepad/valueChangedHandler
func (g_ GCGamepad) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}/* debug [instance_properties/getter]: valueChangedHandler */


// A block called when any element in the profile changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepad/valueChangedHandler
func (g_ GCGamepad) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}/* debug [instance_properties/setter]: valueChangedHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccurrentextendedgamepadsnapshotdataversion
func (g_ GCGamepad) GCCurrentExtendedGamepadSnapshotDataVersion() GCExtendedGamepadSnapshotDataVersion {
	rv := objc.Send[GCExtendedGamepadSnapshotDataVersion](g_.ID, objc.Sel("GCCurrentExtendedGamepadSnapshotDataVersion"))
	return rv
}/* debug [instance_properties/getter]: GCCurrentExtendedGamepadSnapshotDataVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccurrentmicrogamepadsnapshotdataversion
func (g_ GCGamepad) GCCurrentMicroGamepadSnapshotDataVersion() GCMicroGamepadSnapshotDataVersion {
	rv := objc.Send[GCMicroGamepadSnapshotDataVersion](g_.ID, objc.Sel("GCCurrentMicroGamepadSnapshotDataVersion"))
	return rv
}/* debug [instance_properties/getter]: GCCurrentMicroGamepadSnapshotDataVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCGamepad */



