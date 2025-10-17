// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ClipView] class.
var ClipViewClass objc.Class

func init() {
	ClipViewClass = objc.GetClass("NSClipView")
}

type ClipView struct {
	objc.ID
}

func ClipViewFrom(ptr unsafe.Pointer) ClipView {
	return ClipView{
		ID: objc.ID(ptr),
	}
}


// Scrolls the clip view proportionally to  ’s distance outside of it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/autoscroll(with:)
func (c_ ClipView) Autoscroll(event unsafe.Pointer) bool {
	sel := objc.RegisterName("autoscroll:")
	ret := c_.ID.Send(sel, event)
	return ret != 0
}
// Constrains the bounds of the clip view while the user is magnifying and scrolling. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/constrainBoundsRect(_:)
func (c_ ClipView) ConstrainBoundsRect(proposedBounds foundation.Rect) foundation.Rect {
	sel := objc.RegisterName("constrainBoundsRect:")
	ret := c_.ID.Send(sel, proposedBounds)
	return foundation.Rect(ret)
}
// Returns a scroll point adjusted from the proposed new origin, if necessary, to guarantee the view will lie within its document view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/constrainScroll(_:)
func (c_ ClipView) ConstrainScrollPoint(newOrigin foundation.Point) foundation.Point {
	sel := objc.RegisterName("constrainScrollPoint:")
	ret := c_.ID.Send(sel, newOrigin)
	return foundation.Point(ret)
}
// Changes the origin of the clip view’s bounds rectangle to  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/scroll(to:)
func (c_ ClipView) ScrollToPoint(newOrigin foundation.Point) {
	sel := objc.RegisterName("scrollToPoint:")
	c_.ID.Send(sel, newOrigin)
}
// Handles an  , passed in the   argument, by updating a containing   based on the new bounds. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/viewBoundsChanged(_:)
func (c_ ClipView) ViewBoundsChanged(notification unsafe.Pointer) {
	sel := objc.RegisterName("viewBoundsChanged:")
	c_.ID.Send(sel, notification)
}
// Handles an  , passed in the   argument, by updating a containing   based on the new frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/viewFrameChanged(_:)
func (c_ ClipView) ViewFrameChanged(notification unsafe.Pointer) {
	sel := objc.RegisterName("viewFrameChanged:")
	c_.ID.Send(sel, notification)
}

