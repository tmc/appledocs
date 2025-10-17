// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Touch] class.
var TouchClass objc.Class

func init() {
	TouchClass = objc.GetClass("NSTouch")
}

type Touch struct {
	objc.ID
}

func TouchFrom(ptr unsafe.Pointer) Touch {
	return Touch{
		ID: objc.ID(ptr),
	}
}


// Indicates the previous location of the touch in the view’s coordinates. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouch/previousLocation(in:)
func (t_ Touch) PreviousLocationInView(view unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("previousLocationInView:")
	ret := t_.ID.Send(sel, view)
	return unsafe.Pointer(ret)
}

