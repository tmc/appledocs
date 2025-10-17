// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [VisualEffectView] class.
var VisualEffectViewClass objc.Class

func init() {
	VisualEffectViewClass = objc.GetClass("NSVisualEffectView")
}

type VisualEffectView struct {
	objc.ID
}

func VisualEffectViewFrom(ptr unsafe.Pointer) VisualEffectView {
	return VisualEffectView{
		ID: objc.ID(ptr),
	}
}


// Notifies the view that it moved to a new window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/viewDidMoveToWindow()
func (v_ VisualEffectView) ViewDidMoveToWindow() {
	sel := objc.RegisterName("viewDidMoveToWindow")
	v_.ID.Send(sel)
}
// Notifies the view immediately before it moves to a new window (which may be  ). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSVisualEffectView/viewWillMove(toWindow:)
func (v_ VisualEffectView) ViewWillMoveToWindow(newWindow unsafe.Pointer) {
	sel := objc.RegisterName("viewWillMoveToWindow:")
	v_.ID.Send(sel, newWindow)
}

