// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [ClickGestureRecognizer] class.
type IClickGestureRecognizer interface {
	IGestureRecognizer
}

// A discrete gesture recognizer that tracks a specified number of mouse clicks.
//
// When configuring this gesture recognizer, you can specify which mouse buttons must be clicked and how many clicks must occur before the action method is called. The user must click the specified mouse button the required number of times without dragging the mouse for the gesture to be recognized. The gesture recognizer automatically sets the values of the , , and properties to for each button in the property.
//
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

// Alloc allocates a new instance without initialization.
func (cc _ClickGestureRecognizerClass) Alloc() ClickGestureRecognizer {
	rv := objc.Send[ClickGestureRecognizer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A bit mask of the button (or buttons) required to recognize this click.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclickgesturerecognizer/buttonmask
func (c_ ClickGestureRecognizer) ButtonMask() int {
	rv := objc.Send[int](c_.ID, objc.Sel("buttonMask"))
	return rv
}


// SetButtonMask sets the value of the buttonMask property.
// A bit mask of the button (or buttons) required to recognize this click.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclickgesturerecognizer/buttonmask
func (c_ ClickGestureRecognizer) SetButtonMask(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setButtonMask:"), value)
}

// The number of clicks required to match.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclickgesturerecognizer/numberofclicksrequired
func (c_ ClickGestureRecognizer) NumberOfClicksRequired() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfClicksRequired"))
	return rv
}


// SetNumberOfClicksRequired sets the value of the numberOfClicksRequired property.
// The number of clicks required to match.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclickgesturerecognizer/numberofclicksrequired
func (c_ ClickGestureRecognizer) SetNumberOfClicksRequired(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfClicksRequired:"), value)
}

// The number of touches required in an
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclickgesturerecognizer/numberoftouchesrequired
func (c_ ClickGestureRecognizer) NumberOfTouchesRequired() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfTouchesRequired"))
	return rv
}


// SetNumberOfTouchesRequired sets the value of the numberOfTouchesRequired property.
// The number of touches required in an

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclickgesturerecognizer/numberoftouchesrequired
func (c_ ClickGestureRecognizer) SetNumberOfTouchesRequired(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfTouchesRequired:"), value)
}

// A Boolean value that indicates whether other mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysothermousebuttonevents
func (c_ ClickGestureRecognizer) DelaysOtherMouseButtonEvents() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("delaysOtherMouseButtonEvents"))
	return rv
}


// SetDelaysOtherMouseButtonEvents sets the value of the delaysOtherMouseButtonEvents property.
// A Boolean value that indicates whether other mouse button events are delivered only after gesture recognition fails.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysothermousebuttonevents
func (c_ ClickGestureRecognizer) SetDelaysOtherMouseButtonEvents(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelaysOtherMouseButtonEvents:"), value)
}

// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysprimarymousebuttonevents
func (c_ ClickGestureRecognizer) DelaysPrimaryMouseButtonEvents() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("delaysPrimaryMouseButtonEvents"))
	return rv
}


// SetDelaysPrimaryMouseButtonEvents sets the value of the delaysPrimaryMouseButtonEvents property.
// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysprimarymousebuttonevents
func (c_ ClickGestureRecognizer) SetDelaysPrimaryMouseButtonEvents(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelaysPrimaryMouseButtonEvents:"), value)
}

// A Boolean value that indicates whether secondary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delayssecondarymousebuttonevents
func (c_ ClickGestureRecognizer) DelaysSecondaryMouseButtonEvents() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("delaysSecondaryMouseButtonEvents"))
	return rv
}


// SetDelaysSecondaryMouseButtonEvents sets the value of the delaysSecondaryMouseButtonEvents property.
// A Boolean value that indicates whether secondary mouse button events are delivered only after gesture recognition fails.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delayssecondarymousebuttonevents
func (c_ ClickGestureRecognizer) SetDelaysSecondaryMouseButtonEvents(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelaysSecondaryMouseButtonEvents:"), value)
}



