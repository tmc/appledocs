// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PanGestureRecognizer] class.
var (
	panGestureRecognizerClass     _PanGestureRecognizerClass
	panGestureRecognizerClassOnce sync.Once
)

func getPanGestureRecognizerClass() _PanGestureRecognizerClass {
	panGestureRecognizerClassOnce.Do(func() {
		panGestureRecognizerClass = _PanGestureRecognizerClass{objc.GetClass("NSPanGestureRecognizer")}
	})
	return panGestureRecognizerClass
}

type _PanGestureRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [PanGestureRecognizer] class.
type IPanGestureRecognizer interface {
	IGestureRecognizer
}

// A continuous gesture recognizer for panning gestures.
//
// The gesture is recognized when the user clicks all of specified buttons, drags the mouse, and releases one or more of the buttons. Use the pan gesture recognizer object to retrieve the distance traveled during the pan and the location of the mouse as it pans. Upon creation, the gesture recognizer is configured to recognize pan gestures involving only the primary button. It also delays sending primary button events to the view by setting the property to . To change the set of buttons to track, modify the property. In this gesture recognizer, the method always reports the current mouse point, which changes as the user drags the mouse.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer
type PanGestureRecognizer struct {
	GestureRecognizer
}

// PanGestureRecognizerFrom constructs a [PanGestureRecognizer] from an unsafe.Pointer.
//
// A continuous gesture recognizer for panning gestures.
func PanGestureRecognizerFrom(ptr unsafe.Pointer) PanGestureRecognizer {
	return PanGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PanGestureRecognizerClass) Alloc() PanGestureRecognizer {
	rv := objc.Send[PanGestureRecognizer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PanGestureRecognizerClass) New() PanGestureRecognizer {
	rv := objc.Send[PanGestureRecognizer](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PanGestureRecognizer) Init() PanGestureRecognizer {
	rv := objc.Send[PanGestureRecognizer](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PanGestureRecognizer) Autorelease() PanGestureRecognizer {
	rv := objc.Send[PanGestureRecognizer](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPanGestureRecognizer creates a new PanGestureRecognizer instance.
func NewPanGestureRecognizer() PanGestureRecognizer {
	return getPanGestureRecognizerClass().New()
}




