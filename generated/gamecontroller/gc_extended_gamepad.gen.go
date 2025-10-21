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



