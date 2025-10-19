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

// A continuous gesture recognizer for panning gestures. [Full Topic]
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




