// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [GCMicroGamepad] class.
type IGCMicroGamepad interface {
	IGCPhysicalInputProfile
	SaveSnapshot() unsafe.Pointer
	SetStateFromMicroGamepad(microGamepad unsafe.Pointer)
}

// A controller profile that supports the Siri Remote.
//
// The micro gamepad controller profile supports the following input elements: Two digital face buttons (A and X). One analog directional pad (D-pad) that functions as a touchpad. Users can rotate game controllers that support the micro gamepad profile, switching them between landscape and portrait orientation. If you want to get directional values according to the orientation, set the property to .
//
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

// Alloc allocates a new instance without initialization.
func (gc _GCMicroGamepadClass) Alloc() GCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Saves a snapshot of all of the profile’s elements.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/saveSnapshot()
func (g_ GCMicroGamepad) SaveSnapshot() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("saveSnapshot"))
	return rv
}

// Copies the input values from a specified micro gamepad to a snapshot of a micro gamepad.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/setStateFrom(_:)
func (g_ GCMicroGamepad) SetStateFromMicroGamepad(microGamepad unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStateFromMicroGamepad:"), microGamepad)
}

// A Boolean value that indicates whether the profile reports the directional pad values relative to its current orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/allowsRotation
func (g_ GCMicroGamepad) AllowsRotation() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("allowsRotation"))
	return rv
}


// SetAllowsRotation sets the value of the allowsRotation property.
// A Boolean value that indicates whether the profile reports the directional pad values relative to its current orientation.

//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/allowsRotation
func (g_ GCMicroGamepad) SetAllowsRotation(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAllowsRotation:"), value)
}

// The button that the user activates by pressing harder on the touchpad.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/buttonA
func (g_ GCMicroGamepad) ButtonA() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("buttonA"))
	return rv
}

// The menu face button that players use to enter the main menu and pause the game.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/buttonMenu
func (g_ GCMicroGamepad) ButtonMenu() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("buttonMenu"))
	return rv
}

// The second face button element.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/buttonX
func (g_ GCMicroGamepad) ButtonX() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("buttonX"))
	return rv
}

// The controller associated with this profile.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/controller
func (g_ GCMicroGamepad) Controller() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("controller"))
	return rv
}

// The controller’s directional pad element.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/dpad
func (g_ GCMicroGamepad) Dpad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("dpad"))
	return rv
}

// A Boolean value that indicates whether the directional pad reports absolute or relative values.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/reportsAbsoluteDpadValues
func (g_ GCMicroGamepad) ReportsAbsoluteDpadValues() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("reportsAbsoluteDpadValues"))
	return rv
}


// SetReportsAbsoluteDpadValues sets the value of the reportsAbsoluteDpadValues property.
// A Boolean value that indicates whether the directional pad reports absolute or relative values.

//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/reportsAbsoluteDpadValues
func (g_ GCMicroGamepad) SetReportsAbsoluteDpadValues(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReportsAbsoluteDpadValues:"), value)
}

// The block that this profile calls when an element’s value changes.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/valueChangedHandler
func (g_ GCMicroGamepad) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}


// SetValueChangedHandler sets the value of the valueChangedHandler property.
// The block that this profile calls when an element’s value changes.

//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepad/valueChangedHandler
func (g_ GCMicroGamepad) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}



