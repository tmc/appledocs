// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSMagnificationGestureRecognizer */


/* debug [class_header]: Header for NSMagnificationGestureRecognizer */
// The class instance for the [MagnificationGestureRecognizer] class.
var (
	MagnificationGestureRecognizerClass     _MagnificationGestureRecognizerClass
	MagnificationGestureRecognizerClassOnce sync.Once
)

func getMagnificationGestureRecognizerClass() _MagnificationGestureRecognizerClass {
	MagnificationGestureRecognizerClassOnce.Do(func() {
		MagnificationGestureRecognizerClass = _MagnificationGestureRecognizerClass{objc.GetClass("NSMagnificationGestureRecognizer")}
	})
	return MagnificationGestureRecognizerClass
}

type _MagnificationGestureRecognizerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MagnificationGestureRecognizer */
// An interface definition for the [MagnificationGestureRecognizer] class.
type IMagnificationGestureRecognizer interface {
	IGestureRecognizer
	
/* debug [class_interface_properties]: Properties for MagnificationGestureRecognizer */
	// properties:
	Magnification() float64
	SetMagnification(value float64)
	DelaysMagnificationEvents() bool
	SetDelaysMagnificationEvents(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MagnificationGestureRecognizer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MagnificationGestureRecognizer */
// Alloc allocates a new instance without initialization.
func (mc _MagnificationGestureRecognizerClass) Alloc() MagnificationGestureRecognizer {
	rv := objc.Send[MagnificationGestureRecognizer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MagnificationGestureRecognizer */
// A continuous gesture recognizer that tracks a pinch gesture that magnifies content.
//
// This object tracks pinch gestures on a track pad or other input device and stores the resulting magnification value for you to use in your code. This gesture recognizer automatically sets the value of the property to .


// A continuous gesture recognizer that tracks a pinch gesture that magnifies content.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MagnificationGestureRecognizer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MagnificationGestureRecognizer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MagnificationGestureRecognizer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MagnificationGestureRecognizer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MagnificationGestureRecognizer */

// The amount of magnification to apply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMagnificationGestureRecognizer/magnification
func (m_ MagnificationGestureRecognizer) Magnification() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("magnification"))
	return rv
}/* debug [instance_properties/getter]: magnification */


// The amount of magnification to apply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMagnificationGestureRecognizer/magnification
func (m_ MagnificationGestureRecognizer) SetMagnification(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMagnification:"), value)
}/* debug [instance_properties/setter]: magnification */


// A Boolean value that indicates whether magnification events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysmagnificationevents
func (m_ MagnificationGestureRecognizer) DelaysMagnificationEvents() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("delaysMagnificationEvents"))
	return rv
}/* debug [instance_properties/getter]: delaysMagnificationEvents */


// A Boolean value that indicates whether magnification events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysmagnificationevents
func (m_ MagnificationGestureRecognizer) SetDelaysMagnificationEvents(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelaysMagnificationEvents:"), value)
}/* debug [instance_properties/setter]: delaysMagnificationEvents */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMagnificationGestureRecognizer */



