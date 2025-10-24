// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCDirectionalGamepad */


/* debug [class_header]: Header for GCDirectionalGamepad */
// The class instance for the [GCDirectionalGamepad] class.
var (
	GCDirectionalGamepadClass     _GCDirectionalGamepadClass
	GCDirectionalGamepadClassOnce sync.Once
)

func getGCDirectionalGamepadClass() _GCDirectionalGamepadClass {
	GCDirectionalGamepadClassOnce.Do(func() {
		GCDirectionalGamepadClass = _GCDirectionalGamepadClass{objc.GetClass("GCDirectionalGamepad")}
	})
	return GCDirectionalGamepadClass
}

type _GCDirectionalGamepadClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCDirectionalGamepad */
// An interface definition for the [GCDirectionalGamepad] class.
type IGCDirectionalGamepad interface {
	IGCMicroGamepad
	
/* debug [class_interface_properties]: Properties for GCDirectionalGamepad */
	// properties:
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
	IsAnalog() bool
	SetIsAnalog(value bool)
	AllowsRotation() bool
	SetAllowsRotation(value bool)
	ReportsAbsoluteDpadValues() bool
	SetReportsAbsoluteDpadValues(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCDirectionalGamepad */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCDirectionalGamepad */
// Alloc allocates a new instance without initialization.
func (gc _GCDirectionalGamepadClass) Alloc() GCDirectionalGamepad {
	rv := objc.Send[GCDirectionalGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCDirectionalGamepadClass) New() GCDirectionalGamepad {
	rv := objc.Send[GCDirectionalGamepad](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCDirectionalGamepad) Init() GCDirectionalGamepad {
	rv := objc.Send[GCDirectionalGamepad](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCDirectionalGamepad) Autorelease() GCDirectionalGamepad {
	rv := objc.Send[GCDirectionalGamepad](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDirectionalGamepad creates a new GCDirectionalGamepad instance.
func NewGCDirectionalGamepad() GCDirectionalGamepad {
	return getGCDirectionalGamepadClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCDirectionalGamepad */
// A profile that supports only the directional pad, without motion or rotation.
//
// The directional gamepad profile is similar to a micro gamepad profile except it doesn’t support motion or rotation. The controller’s property is and the inherited property is . If you select Micro Gamepad when you add the Game Controllers capability ( ) to your project, and you also support the GCDirectionalGamepad profile, select Directional Gamepad as well. If you support the second-generation Siri Remote and later, set the key to in the information property list in your project. In addition, the directional pad element may report digital or analog values. If the directional pad’s property is , it reports absolute directional pad values (the property is ).


// A profile that supports only the directional pad, without motion or rotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDirectionalGamepad
type GCDirectionalGamepad struct {
	GCMicroGamepad
}

// GCDirectionalGamepadFrom constructs a [GCDirectionalGamepad] from an unsafe.Pointer.
//
// A profile that supports only the directional pad, without motion or rotation.
func GCDirectionalGamepadFrom(ptr unsafe.Pointer) GCDirectionalGamepad {
	return GCDirectionalGamepad{
		GCMicroGamepad: GCMicroGamepadFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCDirectionalGamepad *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCDirectionalGamepad */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCDirectionalGamepad */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCDirectionalGamepad */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCDirectionalGamepad */

// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCDirectionalGamepad) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}/* debug [instance_properties/getter]: extendedGamepad */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCDirectionalGamepad) SetExtendedGamepad(value IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}/* debug [instance_properties/setter]: extendedGamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCDirectionalGamepad) Gamepad() IGCGamepad {
	rv := objc.Send[GCGamepad](g_.ID, objc.Sel("gamepad"))
	return rv
}/* debug [instance_properties/getter]: gamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCDirectionalGamepad) SetGamepad(value IGCGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}/* debug [instance_properties/setter]: gamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCDirectionalGamepad) MicroGamepad() IGCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}/* debug [instance_properties/getter]: microGamepad */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCDirectionalGamepad) SetMicroGamepad(value IGCMicroGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}/* debug [instance_properties/setter]: microGamepad */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCDirectionalGamepad) Motion() IGCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}/* debug [instance_properties/getter]: motion */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCDirectionalGamepad) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}/* debug [instance_properties/setter]: motion */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCDirectionalGamepad) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}/* debug [instance_properties/getter]: physicalInputProfile */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCDirectionalGamepad) SetPhysicalInputProfile(value IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}/* debug [instance_properties/setter]: physicalInputProfile */


// A Boolean value that indicates whether the element provides analog data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/isanalog
func (g_ GCDirectionalGamepad) IsAnalog() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isAnalog"))
	return rv
}/* debug [instance_properties/getter]: isAnalog */


// A Boolean value that indicates whether the element provides analog data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/isanalog
func (g_ GCDirectionalGamepad) SetIsAnalog(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsAnalog:"), value)
}/* debug [instance_properties/setter]: isAnalog */


// A Boolean value that indicates whether the profile reports the directional pad values relative to its current orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmicrogamepad/allowsrotation
func (g_ GCDirectionalGamepad) AllowsRotation() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("allowsRotation"))
	return rv
}/* debug [instance_properties/getter]: allowsRotation */


// A Boolean value that indicates whether the profile reports the directional pad values relative to its current orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmicrogamepad/allowsrotation
func (g_ GCDirectionalGamepad) SetAllowsRotation(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAllowsRotation:"), value)
}/* debug [instance_properties/setter]: allowsRotation */


// A Boolean value that indicates whether the directional pad reports absolute or relative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmicrogamepad/reportsabsolutedpadvalues
func (g_ GCDirectionalGamepad) ReportsAbsoluteDpadValues() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("reportsAbsoluteDpadValues"))
	return rv
}/* debug [instance_properties/getter]: reportsAbsoluteDpadValues */


// A Boolean value that indicates whether the directional pad reports absolute or relative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmicrogamepad/reportsabsolutedpadvalues
func (g_ GCDirectionalGamepad) SetReportsAbsoluteDpadValues(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReportsAbsoluteDpadValues:"), value)
}/* debug [instance_properties/setter]: reportsAbsoluteDpadValues */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCDirectionalGamepad */



