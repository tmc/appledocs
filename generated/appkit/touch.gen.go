// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Touch] class.
var (
	touchClass     _TouchClass
	touchClassOnce sync.Once
)

func getTouchClass() _TouchClass {
	touchClassOnce.Do(func() {
		touchClass = _TouchClass{objc.GetClass("NSTouch")}
	})
	return touchClass
}

type _TouchClass struct {
	class objc.Class
}

// An interface definition for the [Touch] class.
type ITouch interface {
	objectivec.IObject
	PreviousLocationInView(view unsafe.Pointer) unsafe.Pointer
}

// A snapshot of a particular touch at an instant in time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch
type Touch struct {
	objectivec.Object
}

// TouchFrom constructs a [Touch] from an unsafe.Pointer.
//
// A snapshot of a particular touch at an instant in time.
func TouchFrom(ptr unsafe.Pointer) Touch {
	return Touch{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TouchClass) Alloc() Touch {
	rv := objc.Send[Touch](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TouchClass) New() Touch {
	rv := objc.Send[Touch](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Touch) Init() Touch {
	rv := objc.Send[Touch](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Touch) Autorelease() Touch {
	rv := objc.Send[Touch](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTouch creates a new Touch instance.
func NewTouch() Touch {
	return getTouchClass().New()
}


// Indicates the previous location of the touch in the view’s coordinates. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/previousLocation(in:)
func (t_ Touch) PreviousLocationInView(view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("previousLocationInView:"), view)
	return rv
}


