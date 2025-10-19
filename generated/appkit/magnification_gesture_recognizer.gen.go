// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MagnificationGestureRecognizer] class.
var (
	magnificationGestureRecognizerClass     _MagnificationGestureRecognizerClass
	magnificationGestureRecognizerClassOnce sync.Once
)

func getMagnificationGestureRecognizerClass() _MagnificationGestureRecognizerClass {
	magnificationGestureRecognizerClassOnce.Do(func() {
		magnificationGestureRecognizerClass = _MagnificationGestureRecognizerClass{objc.GetClass("NSMagnificationGestureRecognizer")}
	})
	return magnificationGestureRecognizerClass
}

type _MagnificationGestureRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [MagnificationGestureRecognizer] class.
type IMagnificationGestureRecognizer interface {
	IGestureRecognizer
}

// A continuous gesture recognizer that tracks a pinch gesture that magnifies content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMagnificationGestureRecognizer

type MagnificationGestureRecognizer struct {
	GestureRecognizer
}

// MagnificationGestureRecognizerFrom constructs a [MagnificationGestureRecognizer] from an unsafe.Pointer.
//
// A continuous gesture recognizer that tracks a pinch gesture that magnifies content.
func MagnificationGestureRecognizerFrom(ptr unsafe.Pointer) MagnificationGestureRecognizer {
	return MagnificationGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (mc _MagnificationGestureRecognizerClass) Alloc() MagnificationGestureRecognizer {
	rv := objc.Send[MagnificationGestureRecognizer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MagnificationGestureRecognizerClass) New() MagnificationGestureRecognizer {
	rv := objc.Send[MagnificationGestureRecognizer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MagnificationGestureRecognizer) Init() MagnificationGestureRecognizer {
	rv := objc.Send[MagnificationGestureRecognizer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MagnificationGestureRecognizer) Autorelease() MagnificationGestureRecognizer {
	rv := objc.Send[MagnificationGestureRecognizer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMagnificationGestureRecognizer creates a new MagnificationGestureRecognizer instance.
func NewMagnificationGestureRecognizer() MagnificationGestureRecognizer {
	return getMagnificationGestureRecognizerClass().New()
}




