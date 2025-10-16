
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ClipView] class.
var ClipViewClass _ClipViewClass

func init() {
	ClipViewClass = _ClipViewClass{objc.GetClass("NSClipView")}
}

type _ClipViewClass struct {
	objc.Class
}

// An interface definition for the [ClipView] class.
type IClipView interface {
	ID() objc.ID
	Autoscroll(event unsafe.Pointer) bool
	ConstrainBoundsRect(proposedBounds unsafe.Pointer) unsafe.Pointer
	ConstrainScrollPoint(newOrigin unsafe.Pointer) unsafe.Pointer
	ScrollToPoint(newOrigin unsafe.Pointer)
	ViewBoundsChanged(notification unsafe.Pointer)
	ViewFrameChanged(notification unsafe.Pointer)
}

type ClipView struct {
	id objc.ID
}

func ClipViewFrom(ptr unsafe.Pointer) ClipView {
	return ClipView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ ClipView) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ClipViewClass) Alloc() ClipView {
	rv := objc.Send[ClipView](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ClipViewClass) New() ClipView {
	rv := objc.Send[ClipView](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewClipView creates and returns a new initialized instance.
func NewClipView() ClipView {
	return ClipViewClass.New()
}

// Init initializes the instance.
func (c_ ClipView) Init() ClipView {
	rv := objc.Send[ClipView](c_.ID(), selInit)
	return rv
}
// Scrolls the clip view proportionally to  ’s distance outside of it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/autoscroll(with:)
func (c_ ClipView) Autoscroll(event unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("autoscroll:"), event)
	return rv
}
// Constrains the bounds of the clip view while the user is magnifying and scrolling. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/constrainBoundsRect(_:)
func (c_ ClipView) ConstrainBoundsRect(proposedBounds unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("constrainBoundsRect:"), proposedBounds)
	return rv
}
// Returns a scroll point adjusted from the proposed new origin, if necessary, to guarantee the view will lie within its document view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/constrainScroll(_:)
func (c_ ClipView) ConstrainScrollPoint(newOrigin unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("constrainScrollPoint:"), newOrigin)
	return rv
}
// Changes the origin of the clip view’s bounds rectangle to  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/scroll(to:)
func (c_ ClipView) ScrollToPoint(newOrigin unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("scrollToPoint:"), newOrigin)
}
// Handles an  , passed in the   argument, by updating a containing   based on the new bounds. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/viewBoundsChanged(_:)
func (c_ ClipView) ViewBoundsChanged(notification unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("viewBoundsChanged:"), notification)
}
// Handles an  , passed in the   argument, by updating a containing   based on the new frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/viewFrameChanged(_:)
func (c_ ClipView) ViewFrameChanged(notification unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("viewFrameChanged:"), notification)
}
// A Boolean value that indicates if the clip view automatically accounts for other scroll view subviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/automaticallyAdjustsContentInsets
func (c_ ClipView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("automaticallyAdjustsContentInsets"))
	return rv
}
// SetAutomaticallyAdjustsContentInsets sets the value of the automaticallyAdjustsContentInsets property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/automaticallyAdjustsContentInsets
func (c_ ClipView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setAutomaticallyAdjustsContentInsets:"), value)
}
// The color of the clip view’s background. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/backgroundColor
func (c_ ClipView) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("backgroundColor"))
	return rv
}
// SetBackgroundColor sets the value of the backgroundColor property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/backgroundColor
func (c_ ClipView) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setBackgroundColor:"), value)
}
// The distance that the content view is inset from the enclosing scroll view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/contentInsets
func (c_ ClipView) ContentInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("contentInsets"))
	return rv
}
// SetContentInsets sets the value of the contentInsets property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/contentInsets
func (c_ ClipView) SetContentInsets(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setContentInsets:"), value)
}
// A Boolean value that indicates if the clip view copies rendered images while scrolling. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/copiesOnScroll
func (c_ ClipView) CopiesOnScroll() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("copiesOnScroll"))
	return rv
}
// SetCopiesOnScroll sets the value of the copiesOnScroll property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/copiesOnScroll
func (c_ ClipView) SetCopiesOnScroll(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setCopiesOnScroll:"), value)
}
// The cursor object used when the pointer lies over the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/documentCursor
func (c_ ClipView) DocumentCursor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("documentCursor"))
	return rv
}
// SetDocumentCursor sets the value of the documentCursor property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/documentCursor
func (c_ ClipView) SetDocumentCursor(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setDocumentCursor:"), value)
}
// The rectangle defining the document view’s frame, adjusted to the size of the clip view if the document view is smaller. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/documentRect
func (c_ ClipView) DocumentRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("documentRect"))
	return rv
}
// The clip view’s document view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/documentView
func (c_ ClipView) DocumentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("documentView"))
	return rv
}
// SetDocumentView sets the value of the documentView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/documentView
func (c_ ClipView) SetDocumentView(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setDocumentView:"), value)
}
// The exposed rectangle of the clip view’s document view, in the document view’s own coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/documentVisibleRect
func (c_ ClipView) DocumentVisibleRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("documentVisibleRect"))
	return rv
}
// A Boolean value that indicates if the clip view draws its background color. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/drawsBackground
func (c_ ClipView) DrawsBackground() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("drawsBackground"))
	return rv
}
// SetDrawsBackground sets the value of the drawsBackground property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSClipView/drawsBackground
func (c_ ClipView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setDrawsBackground:"), value)
}
