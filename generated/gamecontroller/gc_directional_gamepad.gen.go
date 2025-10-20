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




