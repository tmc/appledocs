// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Touch] class.
var touchClass = _TouchClass{objc.GetClass("NSTouch")}

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

// Indicates the previous location of the touch in the view’s coordinates. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/previousLocation(in:)
func (t_ Touch) PreviousLocationInView(view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("previousLocationInView:"), view)
	return rv
}


