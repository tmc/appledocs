// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSPanGestureRecognizer */


/* debug [class_header]: Header for NSPanGestureRecognizer */
// The class instance for the [PanGestureRecognizer] class.
var (
	PanGestureRecognizerClass     _PanGestureRecognizerClass
	PanGestureRecognizerClassOnce sync.Once
)

func getPanGestureRecognizerClass() _PanGestureRecognizerClass {
	PanGestureRecognizerClassOnce.Do(func() {
		PanGestureRecognizerClass = _PanGestureRecognizerClass{objc.GetClass("NSPanGestureRecognizer")}
	})
	return PanGestureRecognizerClass
}

type _PanGestureRecognizerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PanGestureRecognizer */
// An interface definition for the [PanGestureRecognizer] class.
type IPanGestureRecognizer interface {
	IGestureRecognizer
	
/* debug [class_interface_properties]: Properties for PanGestureRecognizer */
	// properties:
	ButtonMask() uint
	SetButtonMask(value uint)
	NumberOfTouchesRequired() int
	SetNumberOfTouchesRequired(value int)
	DelaysPrimaryMouseButtonEvents() bool
	SetDelaysPrimaryMouseButtonEvents(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PanGestureRecognizer */
	// methods:
	SetTranslationInView(translation vision.Point, view IView)
	TranslationInView(view IView) vision.Point
	VelocityInView(view IView) vision.Point
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PanGestureRecognizer */
// Alloc allocates a new instance without initialization.
func (pc _PanGestureRecognizerClass) Alloc() PanGestureRecognizer {
	rv := objc.Send[PanGestureRecognizer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PanGestureRecognizer */
// A continuous gesture recognizer for panning gestures.
//
// The gesture is recognized when the user clicks all of specified buttons, drags the mouse, and releases one or more of the buttons. Use the pan gesture recognizer object to retrieve the distance traveled during the pan and the location of the mouse as it pans. Upon creation, the gesture recognizer is configured to recognize pan gestures involving only the primary button. It also delays sending primary button events to the view by setting the property to . To change the set of buttons to track, modify the property. In this gesture recognizer, the method always reports the current mouse point, which changes as the user drags the mouse.


// A continuous gesture recognizer for panning gestures.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PanGestureRecognizer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PanGestureRecognizer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PanGestureRecognizer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PanGestureRecognizer */

// Changes the current translation value of the gesture recognizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer/setTranslation(_:in:)
func (p_ PanGestureRecognizer) SetTranslationInView(translation vision.Point, view IView) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTranslation:inView:"), translation, view)
}/* debug [instance_methods/method]: SetTranslationInView */


// The distance traveled by the mouse during the gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer/translation(in:)
func (p_ PanGestureRecognizer) TranslationInView(view IView) vision.Point {
	rv := objc.Send[vision.Point](p_.ID, objc.Sel("translationInView:"), view)
	return rv
}/* debug [instance_methods/method]: TranslationInView */


// The velocity of the pan, measured in points per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer/velocity(in:)
func (p_ PanGestureRecognizer) VelocityInView(view IView) vision.Point {
	rv := objc.Send[vision.Point](p_.ID, objc.Sel("velocityInView:"), view)
	return rv
}/* debug [instance_methods/method]: VelocityInView */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PanGestureRecognizer */

// A bit mask of the button (or buttons) required to recognize this gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer/buttonMask
func (p_ PanGestureRecognizer) ButtonMask() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("buttonMask"))
	return rv
}/* debug [instance_properties/getter]: buttonMask */


// A bit mask of the button (or buttons) required to recognize this gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer/buttonMask
func (p_ PanGestureRecognizer) SetButtonMask(value uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setButtonMask:"), value)
}/* debug [instance_properties/setter]: buttonMask */


// The number of necessary touches on a Touch Bar for the gesture recognizer to match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer/numberOfTouchesRequired
func (p_ PanGestureRecognizer) NumberOfTouchesRequired() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfTouchesRequired"))
	return rv
}/* debug [instance_properties/getter]: numberOfTouchesRequired */


// The number of necessary touches on a Touch Bar for the gesture recognizer to match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer/numberOfTouchesRequired
func (p_ PanGestureRecognizer) SetNumberOfTouchesRequired(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNumberOfTouchesRequired:"), value)
}/* debug [instance_properties/setter]: numberOfTouchesRequired */


// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysprimarymousebuttonevents
func (p_ PanGestureRecognizer) DelaysPrimaryMouseButtonEvents() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("delaysPrimaryMouseButtonEvents"))
	return rv
}/* debug [instance_properties/getter]: delaysPrimaryMouseButtonEvents */


// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysprimarymousebuttonevents
func (p_ PanGestureRecognizer) SetDelaysPrimaryMouseButtonEvents(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelaysPrimaryMouseButtonEvents:"), value)
}/* debug [instance_properties/setter]: delaysPrimaryMouseButtonEvents */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPanGestureRecognizer */



