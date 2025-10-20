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
}

// A controller profile that supports the DualShock 4 controller.
//
// The DualShock 4 controller profile is similar to an extended gamepad ( ), but has a touchpad with a button and two-finger tracking. This profile also supports motion — that is, the controller’s property is non-nil. If you hold the controller in front of you, the direction of the axes are: The positive x-axis points to your right. The positive y-axis points up. The positive z-axis starts at the touchpad and points to you.
//
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




