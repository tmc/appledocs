// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PressGestureRecognizer] class.
var (
	pressGestureRecognizerClass     _PressGestureRecognizerClass
	pressGestureRecognizerClassOnce sync.Once
)

func getPressGestureRecognizerClass() _PressGestureRecognizerClass {
	pressGestureRecognizerClassOnce.Do(func() {
		pressGestureRecognizerClass = _PressGestureRecognizerClass{objc.GetClass("NSPressGestureRecognizer")}
	})
	return pressGestureRecognizerClass
}

type _PressGestureRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [PressGestureRecognizer] class.
type IPressGestureRecognizer interface {
	IGestureRecognizer
}

// A discrete gesture recognizer that tracks whether the user holds down a mouse button for a minimum amount of time before releasing it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPressGestureRecognizer

type PressGestureRecognizer struct {
	GestureRecognizer
}

// PressGestureRecognizerFrom constructs a [PressGestureRecognizer] from an unsafe.Pointer.
//
// A discrete gesture recognizer that tracks whether the user holds down a mouse button for a minimum amount of time before releasing it.
func PressGestureRecognizerFrom(ptr unsafe.Pointer) PressGestureRecognizer {
	return PressGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (pc _PressGestureRecognizerClass) Alloc() PressGestureRecognizer {
	rv := objc.Send[PressGestureRecognizer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PressGestureRecognizerClass) New() PressGestureRecognizer {
	rv := objc.Send[PressGestureRecognizer](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PressGestureRecognizer) Init() PressGestureRecognizer {
	rv := objc.Send[PressGestureRecognizer](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PressGestureRecognizer) Autorelease() PressGestureRecognizer {
	rv := objc.Send[PressGestureRecognizer](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPressGestureRecognizer creates a new PressGestureRecognizer instance.
func NewPressGestureRecognizer() PressGestureRecognizer {
	return getPressGestureRecognizerClass().New()
}




