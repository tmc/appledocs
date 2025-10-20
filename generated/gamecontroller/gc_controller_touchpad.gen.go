// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GCControllerTouchpad] class.
var (
	GCControllerTouchpadClass     _GCControllerTouchpadClass
	GCControllerTouchpadClassOnce sync.Once
)

func getGCControllerTouchpadClass() _GCControllerTouchpadClass {
	GCControllerTouchpadClassOnce.Do(func() {
		GCControllerTouchpadClass = _GCControllerTouchpadClass{objc.GetClass("GCControllerTouchpad")}
	})
	return GCControllerTouchpadClass
}

type _GCControllerTouchpadClass struct {
	class objc.Class
}

// An interface definition for the [GCControllerTouchpad] class.
type IGCControllerTouchpad interface {
	IGCControllerElement
}

// A control element that represents a touch event on a touchpad.
//
// A object provides the state of the touches and presses on a touchpad. This is a compound element with button and directional pad subelements.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad
type GCControllerTouchpad struct {
	GCControllerElement
}

// GCControllerTouchpadFrom constructs a [GCControllerTouchpad] from an unsafe.Pointer.
//
// A control element that represents a touch event on a touchpad.
func GCControllerTouchpadFrom(ptr unsafe.Pointer) GCControllerTouchpad {
	return GCControllerTouchpad{
		GCControllerElement: GCControllerElementFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GCControllerTouchpadClass) Alloc() GCControllerTouchpad {
	rv := objc.Send[GCControllerTouchpad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCControllerTouchpadClass) New() GCControllerTouchpad {
	rv := objc.Send[GCControllerTouchpad](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCControllerTouchpad) Init() GCControllerTouchpad {
	rv := objc.Send[GCControllerTouchpad](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCControllerTouchpad) Autorelease() GCControllerTouchpad {
	rv := objc.Send[GCControllerTouchpad](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerTouchpad creates a new GCControllerTouchpad instance.
func NewGCControllerTouchpad() GCControllerTouchpad {
	return getGCControllerTouchpadClass().New()
}


// The element that represents the button component on the touchpad.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/button
func (g_ GCControllerTouchpad) Button() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("button"))
	return rv
}

// The block that the element calls when the user continues touching the touchpad, not when the user begins or ends touching the touchpad.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchMoved
func (g_ GCControllerTouchpad) TouchMoved() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("touchMoved"))
	return rv
}


// SetTouchMoved sets the value of the touchMoved property.
// The block that the element calls when the user continues touching the touchpad, not when the user begins or ends touching the touchpad.

//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchMoved
func (g_ GCControllerTouchpad) SetTouchMoved(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTouchMoved:"), value)
}


