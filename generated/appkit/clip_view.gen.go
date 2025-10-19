// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ClipView] class.
var clipViewClass = _ClipViewClass{objc.GetClass("NSClipView")}

type _ClipViewClass struct {
	class objc.Class
}

// An interface definition for the [ClipView] class.
type IClipView interface {
	IView
	Autoscroll(event unsafe.Pointer) bool
	ConstrainBoundsRect(proposedBounds unsafe.Pointer) unsafe.Pointer
	ConstrainScrollPoint(newOrigin unsafe.Pointer) unsafe.Pointer
	ScrollToPoint(newOrigin unsafe.Pointer)
	ViewBoundsChanged(notification unsafe.Pointer)
	ViewFrameChanged(notification unsafe.Pointer)
}

// An object that clips a document view to a scroll view’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView

type ClipView struct {
	View
}

// ClipViewFrom constructs a [ClipView] from an unsafe.Pointer.
//
// An object that clips a document view to a scroll view’s frame.
func ClipViewFrom(ptr unsafe.Pointer) ClipView {
	return ClipView{
		View: ViewFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (cc _ClipViewClass) Alloc() ClipView {
	rv := objc.Send[ClipView](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (cc _ClipViewClass) New() ClipView {
	rv := objc.Send[ClipView](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ClipView) Init() ClipView {
	rv := objc.Send[ClipView](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ClipView) Autorelease() ClipView {
	rv := objc.Send[ClipView](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewClipView creates a new ClipView instance.
func NewClipView() ClipView {
	return clipViewClass.New()
}


// Scrolls the clip view proportionally to ’s distance outside of it. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/autoscroll(with:)
func (c_ ClipView) Autoscroll(event unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoscroll:"), event)
	return rv
}
// Constrains the bounds of the clip view while the user is magnifying and scrolling. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/constrainBoundsRect(_:)
func (c_ ClipView) ConstrainBoundsRect(proposedBounds unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("constrainBoundsRect:"), proposedBounds)
	return rv
}
// Returns a scroll point adjusted from the proposed new origin, if necessary, to guarantee the view will lie within its document view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/constrainScroll(_:)
func (c_ ClipView) ConstrainScrollPoint(newOrigin unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("constrainScrollPoint:"), newOrigin)
	return rv
}
// Changes the origin of the clip view’s bounds rectangle to . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/scroll(to:)
func (c_ ClipView) ScrollToPoint(newOrigin unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scrollToPoint:"), newOrigin)
}
// Handles an , passed in the argument, by updating a containing based on the new bounds. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/viewBoundsChanged(_:)
func (c_ ClipView) ViewBoundsChanged(notification unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("viewBoundsChanged:"), notification)
}
// Handles an , passed in the argument, by updating a containing based on the new frame. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/viewFrameChanged(_:)
func (c_ ClipView) ViewFrameChanged(notification unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("viewFrameChanged:"), notification)
}


