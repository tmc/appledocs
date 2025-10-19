// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ClickGestureRecognizer] class.
var (
	clickGestureRecognizerClass     _ClickGestureRecognizerClass
	clickGestureRecognizerClassOnce sync.Once
)

func getClickGestureRecognizerClass() _ClickGestureRecognizerClass {
	clickGestureRecognizerClassOnce.Do(func() {
		clickGestureRecognizerClass = _ClickGestureRecognizerClass{objc.GetClass("NSClickGestureRecognizer")}
	})
	return clickGestureRecognizerClass
}

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




