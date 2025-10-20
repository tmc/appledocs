// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RotationGestureRecognizer] class.
var (
	RotationGestureRecognizerClass     _RotationGestureRecognizerClass
	RotationGestureRecognizerClassOnce sync.Once
)

func getRotationGestureRecognizerClass() _RotationGestureRecognizerClass {
	RotationGestureRecognizerClassOnce.Do(func() {
		RotationGestureRecognizerClass = _RotationGestureRecognizerClass{objc.GetClass("NSRotationGestureRecognizer")}
	})
	return RotationGestureRecognizerClass
}

type _RotationGestureRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [RotationGestureRecognizer] class.
type IRotationGestureRecognizer interface {
	IGestureRecognizer
}

// A continuous gesture recognizer that tracks two trackpad touches moving opposite each other in a circular motion.
//
// This rotation gesture implies that the underlying view should rotate in a matching direction. The gesture is recognized when the trackpad touches end. Upon creation, the gesture recognizer sets the value of the property to .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRotationGestureRecognizer
type RotationGestureRecognizer struct {
	GestureRecognizer
}

// RotationGestureRecognizerFrom constructs a [RotationGestureRecognizer] from an unsafe.Pointer.
//
// A continuous gesture recognizer that tracks two trackpad touches moving opposite each other in a circular motion.
func RotationGestureRecognizerFrom(ptr unsafe.Pointer) RotationGestureRecognizer {
	return RotationGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RotationGestureRecognizerClass) Alloc() RotationGestureRecognizer {
	rv := objc.Send[RotationGestureRecognizer](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RotationGestureRecognizerClass) New() RotationGestureRecognizer {
	rv := objc.Send[RotationGestureRecognizer](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RotationGestureRecognizer) Init() RotationGestureRecognizer {
	rv := objc.Send[RotationGestureRecognizer](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RotationGestureRecognizer) Autorelease() RotationGestureRecognizer {
	rv := objc.Send[RotationGestureRecognizer](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRotationGestureRecognizer creates a new RotationGestureRecognizer instance.
func NewRotationGestureRecognizer() RotationGestureRecognizer {
	return getRotationGestureRecognizerClass().New()
}




