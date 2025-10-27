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
	

	// properties:
	Rotation() float64
	SetRotation(value float64)
	RotationInDegrees() float64
	SetRotationInDegrees(value float64)
	DelaysRotationEvents() bool
	SetDelaysRotationEvents(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _RotationGestureRecognizerClass) Alloc() RotationGestureRecognizer {
	rv := objc.Send[RotationGestureRecognizer](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A continuous gesture recognizer that tracks two trackpad touches moving opposite each other in a circular motion.
//
// This rotation gesture implies that the underlying view should rotate in a matching direction. The gesture is recognized when the trackpad touches end. Upon creation, the gesture recognizer sets the value of the property to .


// A continuous gesture recognizer that tracks two trackpad touches moving opposite each other in a circular motion.
//
// [Full Topic]
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

























// The rotation of the gesture in radians.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRotationGestureRecognizer/rotation
func (r_ RotationGestureRecognizer) Rotation() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("rotation"))
	return rv
}


// The rotation of the gesture in radians.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRotationGestureRecognizer/rotation
func (r_ RotationGestureRecognizer) SetRotation(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRotation:"), value)
}


// The rotation of the gesture in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRotationGestureRecognizer/rotationInDegrees
func (r_ RotationGestureRecognizer) RotationInDegrees() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("rotationInDegrees"))
	return rv
}


// The rotation of the gesture in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRotationGestureRecognizer/rotationInDegrees
func (r_ RotationGestureRecognizer) SetRotationInDegrees(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRotationInDegrees:"), value)
}


// A Boolean value that indicates whether rotation events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysrotationevents
func (r_ RotationGestureRecognizer) DelaysRotationEvents() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("delaysRotationEvents"))
	return rv
}


// A Boolean value that indicates whether rotation events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysrotationevents
func (r_ RotationGestureRecognizer) SetDelaysRotationEvents(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelaysRotationEvents:"), value)
}








