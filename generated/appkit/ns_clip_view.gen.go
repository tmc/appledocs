// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSClipView */


/* debug [class_header]: Header for NSClipView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ClipView */
// An interface definition for the [ClipView] class.
type IClipView interface {
	IView
	
/* debug [class_interface_properties]: Properties for ClipView */
	// properties:
	AutomaticallyAdjustsContentInsets() bool
	SetAutomaticallyAdjustsContentInsets(value bool)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	ContentInsets() foundation.EdgeInsets
	SetContentInsets(value foundation.EdgeInsets)
	CopiesOnScroll() bool
	SetCopiesOnScroll(value bool)
	DocumentCursor() ICursor
	SetDocumentCursor(value ICursor)
	DocumentRect() Rect /* not a class type */
	DocumentView() IView
	SetDocumentView(value IView)
	DocumentVisibleRect() Rect /* not a class type */
	DrawsBackground() bool
	SetDrawsBackground(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ClipView */
	// methods:
	Autoscroll(event IEvent) bool
	ConstrainBoundsRect(proposedBounds Rect /* not a class type */) Rect /* not a class type */
	ScrollToPoint(newOrigin vision.Point)
	ViewBoundsChanged(notification foundation.Notification)
	ViewFrameChanged(notification foundation.Notification)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ClipView */
// Alloc allocates a new instance without initialization.
func (cc _ClipViewClass) Alloc() ClipView {
	rv := objc.Send[ClipView](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ClipView */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ClipView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ClipView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ClipView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ClipView */

// Scrolls the clip view proportionally to ’s distance outside of it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/autoscroll(with:)
func (c_ ClipView) Autoscroll(event IEvent) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoscroll:"), event)
	return rv
}/* debug [instance_methods/method]: Autoscroll */


// Constrains the bounds of the clip view while the user is magnifying and scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/constrainBoundsRect(_:)
func (c_ ClipView) ConstrainBoundsRect(proposedBounds Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("constrainBoundsRect:"), proposedBounds)
	return rv
}/* debug [instance_methods/method]: ConstrainBoundsRect */


// Changes the origin of the clip view’s bounds rectangle to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/scroll(to:)
func (c_ ClipView) ScrollToPoint(newOrigin vision.Point) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scrollToPoint:"), newOrigin)
}/* debug [instance_methods/method]: ScrollToPoint */


// Handles an , passed in the argument, by updating a containing based on the new bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/viewBoundsChanged(_:)
func (c_ ClipView) ViewBoundsChanged(notification foundation.Notification) {
	objc.Send[objc.ID](c_.ID, objc.Sel("viewBoundsChanged:"), notification)
}/* debug [instance_methods/method]: ViewBoundsChanged */


// Handles an , passed in the argument, by updating a containing based on the new frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/viewFrameChanged(_:)
func (c_ ClipView) ViewFrameChanged(notification foundation.Notification) {
	objc.Send[objc.ID](c_.ID, objc.Sel("viewFrameChanged:"), notification)
}/* debug [instance_methods/method]: ViewFrameChanged */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ClipView */

// A Boolean value that indicates if the clip view automatically accounts for other scroll view subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/automaticallyAdjustsContentInsets
func (c_ ClipView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyAdjustsContentInsets"))
	return rv
}/* debug [instance_properties/getter]: automaticallyAdjustsContentInsets */


// A Boolean value that indicates if the clip view automatically accounts for other scroll view subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/automaticallyAdjustsContentInsets
func (c_ ClipView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallyAdjustsContentInsets:"), value)
}/* debug [instance_properties/setter]: automaticallyAdjustsContentInsets */


// The color of the clip view’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/backgroundColor
func (c_ ClipView) BackgroundColor() IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The color of the clip view’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/backgroundColor
func (c_ ClipView) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// The distance that the content view is inset from the enclosing scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/contentInsets
func (c_ ClipView) ContentInsets() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](c_.ID, objc.Sel("contentInsets"))
	return rv
}/* debug [instance_properties/getter]: contentInsets */


// The distance that the content view is inset from the enclosing scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/contentInsets
func (c_ ClipView) SetContentInsets(value foundation.EdgeInsets) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentInsets:"), value)
}/* debug [instance_properties/setter]: contentInsets */


// A Boolean value that indicates if the clip view copies rendered images while scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/copiesOnScroll
func (c_ ClipView) CopiesOnScroll() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("copiesOnScroll"))
	return rv
}/* debug [instance_properties/getter]: copiesOnScroll */


// A Boolean value that indicates if the clip view copies rendered images while scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/copiesOnScroll
func (c_ ClipView) SetCopiesOnScroll(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCopiesOnScroll:"), value)
}/* debug [instance_properties/setter]: copiesOnScroll */


// The cursor object used when the pointer lies over the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentCursor
func (c_ ClipView) DocumentCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("documentCursor"))
	return rv
}/* debug [instance_properties/getter]: documentCursor */


// The cursor object used when the pointer lies over the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentCursor
func (c_ ClipView) SetDocumentCursor(value ICursor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDocumentCursor:"), value)
}/* debug [instance_properties/setter]: documentCursor */


// The rectangle defining the document view’s frame, adjusted to the size of the clip view if the document view is smaller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentRect
func (c_ ClipView) DocumentRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("documentRect"))
	return rv
}/* debug [instance_properties/getter]: documentRect */


// The clip view’s document view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentView
func (c_ ClipView) DocumentView() IView {
	rv := objc.Send[View](c_.ID, objc.Sel("documentView"))
	return rv
}/* debug [instance_properties/getter]: documentView */


// The clip view’s document view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentView
func (c_ ClipView) SetDocumentView(value IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDocumentView:"), value)
}/* debug [instance_properties/setter]: documentView */


// The exposed rectangle of the clip view’s document view, in the document view’s own coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/documentVisibleRect
func (c_ ClipView) DocumentVisibleRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("documentVisibleRect"))
	return rv
}/* debug [instance_properties/getter]: documentVisibleRect */


// A Boolean value that indicates if the clip view draws its background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/drawsBackground
func (c_ ClipView) DrawsBackground() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("drawsBackground"))
	return rv
}/* debug [instance_properties/getter]: drawsBackground */


// A Boolean value that indicates if the clip view draws its background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClipView/drawsBackground
func (c_ ClipView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDrawsBackground:"), value)
}/* debug [instance_properties/setter]: drawsBackground */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSClipView */



