// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSClickGestureRecognizer */


/* debug [class_header]: Header for NSClickGestureRecognizer */
// The class instance for the [ClickGestureRecognizer] class.
var (
	ClickGestureRecognizerClass     _ClickGestureRecognizerClass
	ClickGestureRecognizerClassOnce sync.Once
)

func getClickGestureRecognizerClass() _ClickGestureRecognizerClass {
	ClickGestureRecognizerClassOnce.Do(func() {
		ClickGestureRecognizerClass = _ClickGestureRecognizerClass{objc.GetClass("NSClickGestureRecognizer")}
	})
	return ClickGestureRecognizerClass
}

type _ClickGestureRecognizerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ClickGestureRecognizer */
// An interface definition for the [ClickGestureRecognizer] class.
type IClickGestureRecognizer interface {
	IGestureRecognizer
	
/* debug [class_interface_properties]: Properties for ClickGestureRecognizer */
	// properties:
	ButtonMask() uint
	SetButtonMask(value uint)
	NumberOfClicksRequired() int
	SetNumberOfClicksRequired(value int)
	NumberOfTouchesRequired() int
	SetNumberOfTouchesRequired(value int)
	DelaysOtherMouseButtonEvents() bool
	SetDelaysOtherMouseButtonEvents(value bool)
	DelaysPrimaryMouseButtonEvents() bool
	SetDelaysPrimaryMouseButtonEvents(value bool)
	DelaysSecondaryMouseButtonEvents() bool
	SetDelaysSecondaryMouseButtonEvents(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ClickGestureRecognizer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ClickGestureRecognizer */
// Alloc allocates a new instance without initialization.
func (cc _ClickGestureRecognizerClass) Alloc() ClickGestureRecognizer {
	rv := objc.Send[ClickGestureRecognizer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ClickGestureRecognizerClass) New() ClickGestureRecognizer {
	rv := objc.Send[ClickGestureRecognizer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ClickGestureRecognizer) Init() ClickGestureRecognizer {
	rv := objc.Send[ClickGestureRecognizer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ClickGestureRecognizer) Autorelease() ClickGestureRecognizer {
	rv := objc.Send[ClickGestureRecognizer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewClickGestureRecognizer creates a new ClickGestureRecognizer instance.
func NewClickGestureRecognizer() ClickGestureRecognizer {
	return getClickGestureRecognizerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ClickGestureRecognizer */
// A discrete gesture recognizer that tracks a specified number of mouse clicks.
//
// When configuring this gesture recognizer, you can specify which mouse buttons must be clicked and how many clicks must occur before the action method is called. The user must click the specified mouse button the required number of times without dragging the mouse for the gesture to be recognized. The gesture recognizer automatically sets the values of the , , and properties to for each button in the property.


// A discrete gesture recognizer that tracks a specified number of mouse clicks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClickGestureRecognizer
type ClickGestureRecognizer struct {
	GestureRecognizer
}

// ClickGestureRecognizerFrom constructs a [ClickGestureRecognizer] from an unsafe.Pointer.
//
// A discrete gesture recognizer that tracks a specified number of mouse clicks.
func ClickGestureRecognizerFrom(ptr unsafe.Pointer) ClickGestureRecognizer {
	return ClickGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ClickGestureRecognizer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ClickGestureRecognizer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ClickGestureRecognizer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ClickGestureRecognizer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ClickGestureRecognizer */

// A bit mask of the button (or buttons) required to recognize this click.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClickGestureRecognizer/buttonMask
func (c_ ClickGestureRecognizer) ButtonMask() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("buttonMask"))
	return rv
}/* debug [instance_properties/getter]: buttonMask */


// A bit mask of the button (or buttons) required to recognize this click.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClickGestureRecognizer/buttonMask
func (c_ ClickGestureRecognizer) SetButtonMask(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setButtonMask:"), value)
}/* debug [instance_properties/setter]: buttonMask */


// The number of clicks required to match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClickGestureRecognizer/numberOfClicksRequired
func (c_ ClickGestureRecognizer) NumberOfClicksRequired() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfClicksRequired"))
	return rv
}/* debug [instance_properties/getter]: numberOfClicksRequired */


// The number of clicks required to match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClickGestureRecognizer/numberOfClicksRequired
func (c_ ClickGestureRecognizer) SetNumberOfClicksRequired(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfClicksRequired:"), value)
}/* debug [instance_properties/setter]: numberOfClicksRequired */


// The number of touches required in an object for the gesture recognizer to match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClickGestureRecognizer/numberOfTouchesRequired
func (c_ ClickGestureRecognizer) NumberOfTouchesRequired() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfTouchesRequired"))
	return rv
}/* debug [instance_properties/getter]: numberOfTouchesRequired */


// The number of touches required in an object for the gesture recognizer to match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClickGestureRecognizer/numberOfTouchesRequired
func (c_ ClickGestureRecognizer) SetNumberOfTouchesRequired(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfTouchesRequired:"), value)
}/* debug [instance_properties/setter]: numberOfTouchesRequired */


// A Boolean value that indicates whether other mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysothermousebuttonevents
func (c_ ClickGestureRecognizer) DelaysOtherMouseButtonEvents() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("delaysOtherMouseButtonEvents"))
	return rv
}/* debug [instance_properties/getter]: delaysOtherMouseButtonEvents */


// A Boolean value that indicates whether other mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysothermousebuttonevents
func (c_ ClickGestureRecognizer) SetDelaysOtherMouseButtonEvents(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelaysOtherMouseButtonEvents:"), value)
}/* debug [instance_properties/setter]: delaysOtherMouseButtonEvents */


// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysprimarymousebuttonevents
func (c_ ClickGestureRecognizer) DelaysPrimaryMouseButtonEvents() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("delaysPrimaryMouseButtonEvents"))
	return rv
}/* debug [instance_properties/getter]: delaysPrimaryMouseButtonEvents */


// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysprimarymousebuttonevents
func (c_ ClickGestureRecognizer) SetDelaysPrimaryMouseButtonEvents(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelaysPrimaryMouseButtonEvents:"), value)
}/* debug [instance_properties/setter]: delaysPrimaryMouseButtonEvents */


// A Boolean value that indicates whether secondary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delayssecondarymousebuttonevents
func (c_ ClickGestureRecognizer) DelaysSecondaryMouseButtonEvents() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("delaysSecondaryMouseButtonEvents"))
	return rv
}/* debug [instance_properties/getter]: delaysSecondaryMouseButtonEvents */


// A Boolean value that indicates whether secondary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delayssecondarymousebuttonevents
func (c_ ClickGestureRecognizer) SetDelaysSecondaryMouseButtonEvents(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelaysSecondaryMouseButtonEvents:"), value)
}/* debug [instance_properties/setter]: delaysSecondaryMouseButtonEvents */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSClickGestureRecognizer */



