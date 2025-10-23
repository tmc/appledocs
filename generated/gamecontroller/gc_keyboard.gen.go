// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCKeyboard] class.
var (
	GCKeyboardClass     _GCKeyboardClass
	GCKeyboardClassOnce sync.Once
)

func getGCKeyboardClass() _GCKeyboardClass {
	GCKeyboardClassOnce.Do(func() {
		GCKeyboardClass = _GCKeyboardClass{objc.GetClass("GCKeyboard")}
	})
	return GCKeyboardClass
}

type _GCKeyboardClass struct {
	class objc.Class
}

// An interface definition for the [GCKeyboard] class.
type IGCKeyboard interface {
	objectivec.IObject
	// properties:
	KeyboardInput() unsafe.Pointer
	SetKeyboardInput(value unsafe.Pointer)
	// methods:
}

// An object that represents a physical keyboard connected to a device.
//
// To get the keyboard object and its input values, register for the (Swift) or (Objective-C) notification for when a keyboard connects to the device, or use the class property. Then get the input values from the keyboard object’s controller profile.


// An object that represents a physical keyboard connected to a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCKeyboard
type GCKeyboard struct {
	objectivec.Object
}

// GCKeyboardFrom constructs a [GCKeyboard] from an unsafe.Pointer.
//
// An object that represents a physical keyboard connected to a device.
func GCKeyboardFrom(ptr unsafe.Pointer) GCKeyboard {
	return GCKeyboard{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCKeyboardClass) Alloc() GCKeyboard {
	rv := objc.Send[GCKeyboard](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCKeyboardClass) New() GCKeyboard {
	rv := objc.Send[GCKeyboard](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCKeyboard) Init() GCKeyboard {
	rv := objc.Send[GCKeyboard](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCKeyboard) Autorelease() GCKeyboard {
	rv := objc.Send[GCKeyboard](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCKeyboard creates a new GCKeyboard instance.
func NewGCKeyboard() GCKeyboard {
	return getGCKeyboardClass().New()
}



// The keyboard currently connected to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCKeyboard/coalesced
func (gc _GCKeyboardClass) CoalescedKeyboard() GCKeyboard {
	rv := objc.Send[GCKeyboard](objc.ID(gc.class), objc.Sel("coalescedKeyboard"))
	return rv
}

// The keyboard currently connected to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCKeyboard/coalesced
func (g_ GCKeyboard) CoalescedKeyboard() IGCKeyboard {
	rv := objc.Send[GCKeyboard](g_.ID, objc.Sel("coalescedKeyboard"))
	return rv
}


// The controller profile for the keyboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gckeyboard/keyboardinput
func (g_ GCKeyboard) KeyboardInput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("keyboardInput"))
	return rv
}


// The controller profile for the keyboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gckeyboard/keyboardinput
func (g_ GCKeyboard) SetKeyboardInput(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKeyboardInput:"), value)
}



