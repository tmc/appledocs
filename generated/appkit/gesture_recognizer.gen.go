// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GestureRecognizer] class.
var (
	gestureRecognizerClass     _GestureRecognizerClass
	gestureRecognizerClassOnce sync.Once
)

func getGestureRecognizerClass() _GestureRecognizerClass {
	gestureRecognizerClassOnce.Do(func() {
		gestureRecognizerClass = _GestureRecognizerClass{objc.GetClass("NSGestureRecognizer")}
	})
	return gestureRecognizerClass
}

type _GestureRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [GestureRecognizer] class.
type IGestureRecognizer interface {
	objectivec.IObject
	LocationInView(view unsafe.Pointer) unsafe.Pointer
	TouchesBeganWithEvent(event unsafe.Pointer)
	TouchesCancelledWithEvent(event unsafe.Pointer)
	TouchesEndedWithEvent(event unsafe.Pointer)
	TouchesMovedWithEvent(event unsafe.Pointer)
}

// An object that monitors events and calls its action method when a predefined sequence of events occur. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer

type GestureRecognizer struct {
	objectivec.Object
}

// GestureRecognizerFrom constructs a [GestureRecognizer] from an unsafe.Pointer.
//
// An object that monitors events and calls its action method when a predefined sequence of events occur.
func GestureRecognizerFrom(ptr unsafe.Pointer) GestureRecognizer {
	return GestureRecognizer{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (gc _GestureRecognizerClass) Alloc() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (gc _GestureRecognizerClass) New() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GestureRecognizer) Init() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GestureRecognizer) Autorelease() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGestureRecognizer creates a new GestureRecognizer instance.
func NewGestureRecognizer() GestureRecognizer {
	return getGestureRecognizerClass().New()
}


// Returns the point computed as the location of the gesture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/location(in:)
func (g_ GestureRecognizer) LocationInView(view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("locationInView:"), view)
	return rv
}
// Called when one or more fingers first make contact with an instance on the Touch Bar. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/touchesBegan(with:)
func (g_ GestureRecognizer) TouchesBeganWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("touchesBeganWithEvent:"), event)
}
// Called when a system event, such as a low-memory warning, cancels an in-progress touch event in an object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/touchesCancelled(with:)
func (g_ GestureRecognizer) TouchesCancelledWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("touchesCancelledWithEvent:"), event)
}
// Called when one or more fingers are removed from contact with an instance on the Touch Bar. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/touchesEnded(with:)
func (g_ GestureRecognizer) TouchesEndedWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("touchesEndedWithEvent:"), event)
}
// Called when one or more fingers, associated with an in-progress event, move within an instance on the Touch Bar. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/touchesMoved(with:)
func (g_ GestureRecognizer) TouchesMovedWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("touchesMovedWithEvent:"), event)
}


