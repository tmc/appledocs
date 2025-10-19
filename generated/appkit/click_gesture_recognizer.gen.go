// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ClickGestureRecognizer] class.
var clickGestureRecognizerClass = _ClickGestureRecognizerClass{objc.GetClass("NSClickGestureRecognizer")}

type _ClickGestureRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [ClickGestureRecognizer] class.
type IClickGestureRecognizer interface {
	IGestureRecognizer
}

// A discrete gesture recognizer that tracks a specified number of mouse clicks. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return clickGestureRecognizerClass.New()
}




