
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GestureRecognizer] class.
var GestureRecognizerClass _GestureRecognizerClass

func init() {
	GestureRecognizerClass = _GestureRecognizerClass{objc.GetClass("NSGestureRecognizer")}
}

type _GestureRecognizerClass struct {
	objc.Class
}

// An interface definition for the [GestureRecognizer] class.
type IGestureRecognizer interface {
	ID() objc.ID
	LocationInView(view unsafe.Pointer) unsafe.Pointer
	TouchesBeganWithEvent(event unsafe.Pointer)
	TouchesCancelledWithEvent(event unsafe.Pointer)
	TouchesEndedWithEvent(event unsafe.Pointer)
	TouchesMovedWithEvent(event unsafe.Pointer)
}

type GestureRecognizer struct {
	id objc.ID
}

func GestureRecognizerFrom(ptr unsafe.Pointer) GestureRecognizer {
	return GestureRecognizer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (g_ GestureRecognizer) ID() objc.ID {
	return g_.id
}

// Alloc allocates a new instance without initialization.
func (gc _GestureRecognizerClass) Alloc() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](objc.ID(gc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (gc _GestureRecognizerClass) New() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](objc.ID(gc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewGestureRecognizer creates and returns a new initialized instance.
func NewGestureRecognizer() GestureRecognizer {
	return GestureRecognizerClass.New()
}

// Init initializes the instance.
func (g_ GestureRecognizer) Init() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](g_.ID(), selInit)
	return rv
}
// Returns the point computed as the location of the gesture. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGestureRecognizer/location(in:)
func (g_ GestureRecognizer) LocationInView(view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("locationInView:"), view)
	return rv
}
// Called when one or more fingers first make contact with an   instance on the Touch Bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGestureRecognizer/touchesBegan(with:)
func (g_ GestureRecognizer) TouchesBeganWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("touchesBeganWithEvent:"), event)
}
// Called when a system event, such as a low-memory warning, cancels an in-progress touch event in an   object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGestureRecognizer/touchesCancelled(with:)
func (g_ GestureRecognizer) TouchesCancelledWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("touchesCancelledWithEvent:"), event)
}
// Called when one or more fingers are removed from contact with an   instance on the Touch Bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGestureRecognizer/touchesEnded(with:)
func (g_ GestureRecognizer) TouchesEndedWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("touchesEndedWithEvent:"), event)
}
// Called when one or more fingers, associated with an in-progress event, move within an   instance on the Touch Bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGestureRecognizer/touchesMoved(with:)
func (g_ GestureRecognizer) TouchesMovedWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("touchesMovedWithEvent:"), event)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGestureRecognizer/allowedTouchTypes
func (g_ GestureRecognizer) AllowedTouchTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("allowedTouchTypes"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGestureRecognizer/allowedTouchTypes
func (g_ GestureRecognizer) SetAllowedTouchTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("setAllowedTouchTypes:"), value)
}
