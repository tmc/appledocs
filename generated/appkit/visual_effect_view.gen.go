// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VisualEffectView] class.
var (
	visualEffectViewClass     _VisualEffectViewClass
	visualEffectViewClassOnce sync.Once
)

func getVisualEffectViewClass() _VisualEffectViewClass {
	visualEffectViewClassOnce.Do(func() {
		visualEffectViewClass = _VisualEffectViewClass{objc.GetClass("NSVisualEffectView")}
	})
	return visualEffectViewClass
}

type _VisualEffectViewClass struct {
	class objc.Class
}

// An interface definition for the [VisualEffectView] class.
type IVisualEffectView interface {
	IView
	ViewDidMoveToWindow()
	ViewWillMoveToWindow(newWindow unsafe.Pointer)
}

// A view that adds translucency and vibrancy effects to the views in your interface.
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

// Alloc allocates a new instance without initialization.
func (vc _VisualEffectViewClass) Alloc() VisualEffectView {
	rv := objc.Send[VisualEffectView](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VisualEffectViewClass) New() VisualEffectView {
	rv := objc.Send[VisualEffectView](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VisualEffectView) Init() VisualEffectView {
	rv := objc.Send[VisualEffectView](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VisualEffectView) Autorelease() VisualEffectView {
	rv := objc.Send[VisualEffectView](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVisualEffectView creates a new VisualEffectView instance.
func NewVisualEffectView() VisualEffectView {
	return getVisualEffectViewClass().New()
}


// Notifies the view that it moved to a new window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/viewDidMoveToWindow()
func (v_ VisualEffectView) ViewDidMoveToWindow() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidMoveToWindow"))
}

// Notifies the view immediately before it moves to a new window (which may be ).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/viewWillMove(toWindow:)
func (v_ VisualEffectView) ViewWillMoveToWindow(newWindow unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillMoveToWindow:"), newWindow)
}



