// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [GCDirectionalGamepad] class.
type IGCDirectionalGamepad interface {
	IGCMicroGamepad
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
	IsAnalog() bool
	SetIsAnalog(value bool)
	AllowsRotation() bool
	SetAllowsRotation(value bool)
	ReportsAbsoluteDpadValues() bool
	SetReportsAbsoluteDpadValues(value bool)
}

// A profile that supports only the directional pad, without motion or rotation.
//
// The directional gamepad profile is similar to a micro gamepad profile except it doesn’t support motion or rotation. The controller’s property is and the inherited property is . If you select Micro Gamepad when you add the Game Controllers capability ( ) to your project, and you also support the GCDirectionalGamepad profile, select Directional Gamepad as well. If you support the second-generation Siri Remote and later, set the key to in the information property list in your project. In addition, the directional pad element may report digital or analog values. If the directional pad’s property is , it reports absolute directional pad values (the property is ).
//
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

// Alloc allocates a new instance without initialization.
func (gc _GCDirectionalGamepadClass) Alloc() GCDirectionalGamepad {
	rv := objc.Send[GCDirectionalGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The extended gamepad profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCDirectionalGamepad) ExtendedGamepad() GCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}


// SetExtendedGamepad sets the value of the extendedGamepad property.
// The extended gamepad profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g_ GCDirectionalGamepad) SetExtendedGamepad(value IGCExtendedGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setExtendedGamepad:"), value)
}

// The gamepad profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCDirectionalGamepad) Gamepad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("gamepad"))
	return rv
}


// SetGamepad sets the value of the gamepad property.
// The gamepad profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCDirectionalGamepad) SetGamepad(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}

// The micro gamepad profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCDirectionalGamepad) MicroGamepad() GCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}


// SetMicroGamepad sets the value of the microGamepad property.
// The micro gamepad profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCDirectionalGamepad) SetMicroGamepad(value IGCMicroGamepad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}

// The motion input profile.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCDirectionalGamepad) Motion() GCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}


// SetMotion sets the value of the motion property.
// The motion input profile.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCDirectionalGamepad) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}

// The physical input profile for the controller.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCDirectionalGamepad) PhysicalInputProfile() GCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}


// SetPhysicalInputProfile sets the value of the physicalInputProfile property.
// The physical input profile for the controller.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g_ GCDirectionalGamepad) SetPhysicalInputProfile(value IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPhysicalInputProfile:"), value)
}

// A Boolean value that indicates whether the element provides analog data.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/isanalog
func (g_ GCDirectionalGamepad) IsAnalog() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isAnalog"))
	return rv
}


// SetIsAnalog sets the value of the isAnalog property.
// A Boolean value that indicates whether the element provides analog data.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/isanalog
func (g_ GCDirectionalGamepad) SetIsAnalog(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsAnalog:"), value)
}

// A Boolean value that indicates whether the profile reports the directional pad values relative to its current orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmicrogamepad/allowsrotation
func (g_ GCDirectionalGamepad) AllowsRotation() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("allowsRotation"))
	return rv
}


// SetAllowsRotation sets the value of the allowsRotation property.
// A Boolean value that indicates whether the profile reports the directional pad values relative to its current orientation.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmicrogamepad/allowsrotation
func (g_ GCDirectionalGamepad) SetAllowsRotation(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAllowsRotation:"), value)
}

// A Boolean value that indicates whether the directional pad reports absolute or relative values.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmicrogamepad/reportsabsolutedpadvalues
func (g_ GCDirectionalGamepad) ReportsAbsoluteDpadValues() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("reportsAbsoluteDpadValues"))
	return rv
}


// SetReportsAbsoluteDpadValues sets the value of the reportsAbsoluteDpadValues property.
// A Boolean value that indicates whether the directional pad reports absolute or relative values.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcmicrogamepad/reportsabsolutedpadvalues
func (g_ GCDirectionalGamepad) SetReportsAbsoluteDpadValues(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReportsAbsoluteDpadValues:"), value)
}



