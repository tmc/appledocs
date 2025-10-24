// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ClipView] class.
var (
	ClipViewClass     _ClipViewClass
	ClipViewClassOnce sync.Once
)

func getClipViewClass() _ClipViewClass {
	ClipViewClassOnce.Do(func() {
		ClipViewClass = _ClipViewClass{objc.GetClass("NSClipView")}
	})
	return ClipViewClass
}

type _ClipViewClass struct {
	class objc.Class
}

// An interface definition for the [ClipView] class.
type IClipView interface {
	IView
	// properties:
	AutomaticallyAdjustsContentInsets() bool
	SetAutomaticallyAdjustsContentInsets(value bool)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	ContentInsets() objc.IObject /* cross-framework: EdgeInsets */
	SetContentInsets(value objc.IObject /* cross-framework: EdgeInsets */)
	CopiesOnScroll() bool
	SetCopiesOnScroll(value bool)
	DocumentCursor() ICursor
	SetDocumentCursor(value ICursor)
	DocumentRect() objc.IObject /* cross-framework: Rect */
	DocumentView() IView
	SetDocumentView(value IView)
	DocumentVisibleRect() objc.IObject /* cross-framework: Rect */
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	// methods:
	Autoscroll(event IEvent) bool
	ConstrainBoundsRect(proposedBounds objc.IObject /* cross-framework: Rect */) objc.IObject /* cross-framework: Rect */
	ScrollToPoint(newOrigin objc.IObject /* cross-framework: Point */)
	ViewBoundsChanged(notification objc.IObject /* cross-framework: Notification */)
	ViewFrameChanged(notification objc.IObject /* cross-framework: Notification */)
}

// An object that clips a document view to a scroll view’s frame.
//
// An holds the document view of an , clipping the document view to its frame, handling the details of scrolling in an efficient manner, and updating the when the document view’s size or position changes. You don’t typically use the class directly; it’s provided primarily as the scrolling machinery for the class. However, you might use the class to implement a class similar to .


// An object that clips a document view to a scroll view’s frame.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getClipViewClass().New()
}



// Scrolls the clip view proportionally to ’s distance outside of it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/autoscroll(with:)
func (c_ ClipView) Autoscroll(event IEvent) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoscroll:"), event)
	return rv
}


// Constrains the bounds of the clip view while the user is magnifying and scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/constrainBoundsRect(_:)
func (c_ ClipView) ConstrainBoundsRect(proposedBounds objc.IObject /* cross-framework: Rect */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](c_.ID, objc.Sel("constrainBoundsRect:"), proposedBounds)
	return rv
}


// Changes the origin of the clip view’s bounds rectangle to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/scroll(to:)
func (c_ ClipView) ScrollToPoint(newOrigin objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scrollToPoint:"), newOrigin)
}


// Handles an , passed in the argument, by updating a containing based on the new bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/viewBoundsChanged(_:)
func (c_ ClipView) ViewBoundsChanged(notification objc.IObject /* cross-framework: Notification */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("viewBoundsChanged:"), notification)
}


// Handles an , passed in the argument, by updating a containing based on the new frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/viewFrameChanged(_:)
func (c_ ClipView) ViewFrameChanged(notification objc.IObject /* cross-framework: Notification */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("viewFrameChanged:"), notification)
}


// A Boolean value that indicates if the clip view automatically accounts for other scroll view subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/automaticallyAdjustsContentInsets
func (c_ ClipView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyAdjustsContentInsets"))
	return rv
}


// A Boolean value that indicates if the clip view automatically accounts for other scroll view subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/automaticallyAdjustsContentInsets
func (c_ ClipView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallyAdjustsContentInsets:"), value)
}


// The color of the clip view’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/backgroundColor
func (c_ ClipView) BackgroundColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The color of the clip view’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/backgroundColor
func (c_ ClipView) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The distance that the content view is inset from the enclosing scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/contentInsets
func (c_ ClipView) ContentInsets() objc.IObject /* cross-framework: EdgeInsets */ {
	rv := objc.Send[foundation.EdgeInsets](c_.ID, objc.Sel("contentInsets"))
	return rv
}


// The distance that the content view is inset from the enclosing scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/contentInsets
func (c_ ClipView) SetContentInsets(value objc.IObject /* cross-framework: EdgeInsets */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentInsets:"), value)
}


// A Boolean value that indicates if the clip view copies rendered images while scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/copiesOnScroll
func (c_ ClipView) CopiesOnScroll() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("copiesOnScroll"))
	return rv
}


// A Boolean value that indicates if the clip view copies rendered images while scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/copiesOnScroll
func (c_ ClipView) SetCopiesOnScroll(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCopiesOnScroll:"), value)
}


// The cursor object used when the pointer lies over the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentCursor
func (c_ ClipView) DocumentCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("documentCursor"))
	return rv
}


// The cursor object used when the pointer lies over the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentCursor
func (c_ ClipView) SetDocumentCursor(value ICursor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDocumentCursor:"), value)
}


// The rectangle defining the document view’s frame, adjusted to the size of the clip view if the document view is smaller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentRect
func (c_ ClipView) DocumentRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](c_.ID, objc.Sel("documentRect"))
	return rv
}


// The clip view’s document view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentView
func (c_ ClipView) DocumentView() IView {
	rv := objc.Send[View](c_.ID, objc.Sel("documentView"))
	return rv
}


// The clip view’s document view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentView
func (c_ ClipView) SetDocumentView(value IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDocumentView:"), value)
}


// The exposed rectangle of the clip view’s document view, in the document view’s own coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentVisibleRect
func (c_ ClipView) DocumentVisibleRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](c_.ID, objc.Sel("documentVisibleRect"))
	return rv
}


// A Boolean value that indicates if the clip view draws its background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/drawsBackground
func (c_ ClipView) DrawsBackground() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean value that indicates if the clip view draws its background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/drawsBackground
func (c_ ClipView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDrawsBackground:"), value)
}



