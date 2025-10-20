// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	Autoscroll(event unsafe.Pointer) bool
	ConstrainBoundsRect(proposedBounds coregraphics.CGRect) coregraphics.CGRect
	ConstrainScrollPoint(newOrigin coregraphics.CGPoint) coregraphics.CGPoint
	ScrollToPoint(newOrigin coregraphics.CGPoint)
	ViewBoundsChanged(notification unsafe.Pointer)
	ViewFrameChanged(notification unsafe.Pointer)
}

// An object that clips a document view to a scroll view’s frame.
//
// An holds the document view of an , clipping the document view to its frame, handling the details of scrolling in an efficient manner, and updating the when the document view’s size or position changes. You don’t typically use the class directly; it’s provided primarily as the scrolling machinery for the class. However, you might use the class to implement a class similar to .
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
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/autoscroll(with:)
func (c_ ClipView) Autoscroll(event unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoscroll:"), event)
	return rv
}

// Constrains the bounds of the clip view while the user is magnifying and scrolling.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/constrainBoundsRect(_:)
func (c_ ClipView) ConstrainBoundsRect(proposedBounds coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("constrainBoundsRect:"), proposedBounds)
	return rv
}

// Returns a scroll point adjusted from the proposed new origin, if necessary, to guarantee the view will lie within its document view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/constrainScroll(_:)
func (c_ ClipView) ConstrainScrollPoint(newOrigin coregraphics.CGPoint) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](c_.ID, objc.Sel("constrainScrollPoint:"), newOrigin)
	return rv
}

// Changes the origin of the clip view’s bounds rectangle to .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/scroll(to:)
func (c_ ClipView) ScrollToPoint(newOrigin coregraphics.CGPoint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scrollToPoint:"), newOrigin)
}

// Handles an , passed in the argument, by updating a containing based on the new bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/viewBoundsChanged(_:)
func (c_ ClipView) ViewBoundsChanged(notification unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("viewBoundsChanged:"), notification)
}

// Handles an , passed in the argument, by updating a containing based on the new frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/viewFrameChanged(_:)
func (c_ ClipView) ViewFrameChanged(notification unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("viewFrameChanged:"), notification)
}

// A Boolean value that indicates if the clip view automatically accounts for other scroll view subviews.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/automaticallyAdjustsContentInsets
func (c_ ClipView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyAdjustsContentInsets"))
	return rv
}


// SetAutomaticallyAdjustsContentInsets sets the value of the automaticallyAdjustsContentInsets property.
// A Boolean value that indicates if the clip view automatically accounts for other scroll view subviews.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/automaticallyAdjustsContentInsets
func (c_ ClipView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallyAdjustsContentInsets:"), value)
}
// The color of the clip view’s background.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/backgroundColor
func (c_ ClipView) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The color of the clip view’s background.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/backgroundColor
func (c_ ClipView) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBackgroundColor:"), value)
}
// The distance that the content view is inset from the enclosing scroll view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/contentInsets
func (c_ ClipView) ContentInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("contentInsets"))
	return rv
}


// SetContentInsets sets the value of the contentInsets property.
// The distance that the content view is inset from the enclosing scroll view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/contentInsets
func (c_ ClipView) SetContentInsets(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentInsets:"), value)
}
// A Boolean value that indicates if the clip view copies rendered images while scrolling.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/copiesOnScroll
func (c_ ClipView) CopiesOnScroll() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("copiesOnScroll"))
	return rv
}


// SetCopiesOnScroll sets the value of the copiesOnScroll property.
// A Boolean value that indicates if the clip view copies rendered images while scrolling.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/copiesOnScroll
func (c_ ClipView) SetCopiesOnScroll(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCopiesOnScroll:"), value)
}
// The cursor object used when the pointer lies over the view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentCursor
func (c_ ClipView) DocumentCursor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("documentCursor"))
	return rv
}


// SetDocumentCursor sets the value of the documentCursor property.
// The cursor object used when the pointer lies over the view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentCursor
func (c_ ClipView) SetDocumentCursor(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDocumentCursor:"), value)
}
// The rectangle defining the document view’s frame, adjusted to the size of the clip view if the document view is smaller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentRect
func (c_ ClipView) DocumentRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("documentRect"))
	return rv
}

// The clip view’s document view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentView
func (c_ ClipView) DocumentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("documentView"))
	return rv
}


// SetDocumentView sets the value of the documentView property.
// The clip view’s document view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentView
func (c_ ClipView) SetDocumentView(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDocumentView:"), value)
}
// The exposed rectangle of the clip view’s document view, in the document view’s own coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentVisibleRect
func (c_ ClipView) DocumentVisibleRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("documentVisibleRect"))
	return rv
}

// A Boolean value that indicates if the clip view draws its background color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/drawsBackground
func (c_ ClipView) DrawsBackground() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("drawsBackground"))
	return rv
}


// SetDrawsBackground sets the value of the drawsBackground property.
// A Boolean value that indicates if the clip view draws its background color.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/drawsBackground
func (c_ ClipView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDrawsBackground:"), value)
}


