// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PressGestureRecognizer] class.
var (
	PressGestureRecognizerClass     _PressGestureRecognizerClass
	PressGestureRecognizerClassOnce sync.Once
)

func getPressGestureRecognizerClass() _PressGestureRecognizerClass {
	PressGestureRecognizerClassOnce.Do(func() {
		PressGestureRecognizerClass = _PressGestureRecognizerClass{objc.GetClass("NSPressGestureRecognizer")}
	})
	return PressGestureRecognizerClass
}

type _PressGestureRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [PressGestureRecognizer] class.
type IPressGestureRecognizer interface {
	IGestureRecognizer
	MinimumPressDuration() float64
	SetMinimumPressDuration(value float64)
	DelaysPrimaryMouseButtonEvents() bool
	SetDelaysPrimaryMouseButtonEvents(value bool)
	AllowableMovement() float64
	SetAllowableMovement(value float64)
	ButtonMask() int
	SetButtonMask(value int)
	NumberOfTouchesRequired() int
	SetNumberOfTouchesRequired(value int)
}

// A discrete gesture recognizer that tracks whether the user holds down a mouse button for a minimum amount of time before releasing it.
//
// Use a press gesture recognizer to configure which button the user must hold and the length of time they must hold it. You can also specify how far the mouse can move for a valid gesture. Upon creation, the gesture recognizer recognizes press gestures involving only the primary button. It also delays sending primary button events to the view by setting the property to . To change the set of buttons to track, modify the property.


// A discrete gesture recognizer that tracks whether the user holds down a mouse button for a minimum amount of time before releasing it.
//
// [Full Topic]
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



// The minimum time (in seconds) that the user must hold the mouse button in the view for a valid gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPressGestureRecognizer/minimumPressDuration

func (p_ PressGestureRecognizer) MinimumPressDuration() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("minimumPressDuration"))
	return rv
}


// The minimum time (in seconds) that the user must hold the mouse button in the view for a valid gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPressGestureRecognizer/minimumPressDuration

func (p_ PressGestureRecognizer) SetMinimumPressDuration(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMinimumPressDuration:"), value)
}


// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysprimarymousebuttonevents

func (p_ PressGestureRecognizer) DelaysPrimaryMouseButtonEvents() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("delaysPrimaryMouseButtonEvents"))
	return rv
}


// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysprimarymousebuttonevents

func (p_ PressGestureRecognizer) SetDelaysPrimaryMouseButtonEvents(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelaysPrimaryMouseButtonEvents:"), value)
}


// The maximum movement of the mouse in the view before the gesture fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer/allowablemovement

func (p_ PressGestureRecognizer) AllowableMovement() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("allowableMovement"))
	return rv
}


// The maximum movement of the mouse in the view before the gesture fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer/allowablemovement

func (p_ PressGestureRecognizer) SetAllowableMovement(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowableMovement:"), value)
}


// A bit mask of the buttons required to recognize this press.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer/buttonmask

func (p_ PressGestureRecognizer) ButtonMask() int {
	rv := objc.Send[int](p_.ID, objc.Sel("buttonMask"))
	return rv
}


// A bit mask of the buttons required to recognize this press.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer/buttonmask

func (p_ PressGestureRecognizer) SetButtonMask(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setButtonMask:"), value)
}


// The number of necessary touches on a Touch Bar for the gesture recognizer to match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer/numberoftouchesrequired

func (p_ PressGestureRecognizer) NumberOfTouchesRequired() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfTouchesRequired"))
	return rv
}


// The number of necessary touches on a Touch Bar for the gesture recognizer to match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressgesturerecognizer/numberoftouchesrequired

func (p_ PressGestureRecognizer) SetNumberOfTouchesRequired(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNumberOfTouchesRequired:"), value)
}



