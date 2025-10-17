// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VisualEffectView] class.
var visualEffectViewClass = _VisualEffectViewClass{objc.GetClass("NSVisualEffectView")}

type _VisualEffectViewClass struct {
	class objc.Class
}

// A view that adds translucency and vibrancy effects to the views in your interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView

type VisualEffectView struct {
	View
}

// VisualEffectViewFrom constructs a [VisualEffectView] from an unsafe.Pointer.
//
// A view that adds translucency and vibrancy effects to the views in your interface.
func VisualEffectViewFrom(ptr unsafe.Pointer) VisualEffectView {
	return VisualEffectView{
		View: ViewFrom(ptr),
	}
}

// Notifies the view that it moved to a new window. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/viewDidMoveToWindow()
func (v_ VisualEffectView) ViewDidMoveToWindow() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidMoveToWindow"))
}
// Notifies the view immediately before it moves to a new window (which may be ). [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/viewWillMove(toWindow:)
func (v_ VisualEffectView) ViewWillMoveToWindow(newWindow unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillMoveToWindow:"), newWindow)
}


