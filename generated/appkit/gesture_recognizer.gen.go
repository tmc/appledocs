// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GestureRecognizer] class.
var GestureRecognizerClass objc.Class

func init() {
	GestureRecognizerClass = objc.GetClass("NSGestureRecognizer")
}

type GestureRecognizer struct {
	objc.ID
}

func GestureRecognizerFrom(ptr unsafe.Pointer) GestureRecognizer {
	return GestureRecognizer{
		ID: objc.ID(ptr),
	}
}


// Returns the point computed as the location of the gesture. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGestureRecognizer/location(in:)
func (g_ GestureRecognizer) LocationInView(view unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("locationInView:")
	ret := g_.ID.Send(sel, view)
	return unsafe.Pointer(ret)
}
// Called when one or more fingers first make contact with an   instance on the Touch Bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGestureRecognizer/touchesBegan(with:)
func (g_ GestureRecognizer) TouchesBeganWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("touchesBeganWithEvent:")
	g_.ID.Send(sel, event)
}
// Called when a system event, such as a low-memory warning, cancels an in-progress touch event in an   object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGestureRecognizer/touchesCancelled(with:)
func (g_ GestureRecognizer) TouchesCancelledWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("touchesCancelledWithEvent:")
	g_.ID.Send(sel, event)
}
// Called when one or more fingers are removed from contact with an   instance on the Touch Bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGestureRecognizer/touchesEnded(with:)
func (g_ GestureRecognizer) TouchesEndedWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("touchesEndedWithEvent:")
	g_.ID.Send(sel, event)
}
// Called when one or more fingers, associated with an in-progress event, move within an   instance on the Touch Bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGestureRecognizer/touchesMoved(with:)
func (g_ GestureRecognizer) TouchesMovedWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("touchesMovedWithEvent:")
	g_.ID.Send(sel, event)
}

