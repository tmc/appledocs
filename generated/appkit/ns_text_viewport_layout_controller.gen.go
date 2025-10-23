// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextViewportLayoutController] class.
var (
	TextViewportLayoutControllerClass     _TextViewportLayoutControllerClass
	TextViewportLayoutControllerClassOnce sync.Once
)

func getTextViewportLayoutControllerClass() _TextViewportLayoutControllerClass {
	TextViewportLayoutControllerClassOnce.Do(func() {
		TextViewportLayoutControllerClass = _TextViewportLayoutControllerClass{objc.GetClass("NSTextViewportLayoutController")}
	})
	return TextViewportLayoutControllerClass
}

type _TextViewportLayoutControllerClass struct {
	class objc.Class
}

// An interface definition for the [TextViewportLayoutController] class.
type ITextViewportLayoutController interface {
	objectivec.IObject
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	TextLayoutManager() NSTextLayoutManager
	SetTextLayoutManager(value ITextLayoutManager)
	ViewportBounds() coregraphics.CGRect
	SetViewportBounds(value coregraphics.CGRect)
	ViewportRange() NSTextRange
	SetViewportRange(value ITextRange)
}

// Manages the layout process inside the viewport interacting with its delegate.
//
// A viewport is a rectangular area within a flipped coordinate system expanding along the y-axis. With text contents, lines advance expanding the view in the current writing direction. The viewport defines the active area where the framework lays out text fragments. In most cases, the area corresponds to the user visible area with an additional over-scroll region.


// Manages the layout process inside the viewport interacting with its delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController
type TextViewportLayoutController struct {
	objectivec.Object
}

// TextViewportLayoutControllerFrom constructs a [TextViewportLayoutController] from an unsafe.Pointer.
//
// Manages the layout process inside the viewport interacting with its delegate.
func TextViewportLayoutControllerFrom(ptr unsafe.Pointer) TextViewportLayoutController {
	return TextViewportLayoutController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextViewportLayoutControllerClass) Alloc() TextViewportLayoutController {
	rv := objc.Send[TextViewportLayoutController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextViewportLayoutControllerClass) New() TextViewportLayoutController {
	rv := objc.Send[TextViewportLayoutController](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextViewportLayoutController) Init() TextViewportLayoutController {
	rv := objc.Send[TextViewportLayoutController](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextViewportLayoutController) Autorelease() TextViewportLayoutController {
	rv := objc.Send[TextViewportLayoutController](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextViewportLayoutController creates a new TextViewportLayoutController instance.
func NewTextViewportLayoutController() TextViewportLayoutController {
	return getTextViewportLayoutControllerClass().New()
}



// The delegate for the text layout manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewportlayoutcontroller/delegate
func (t_ TextViewportLayoutController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for the text layout manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewportlayoutcontroller/delegate
func (t_ TextViewportLayoutController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}


// Returns the text layout manager for this viewport layout controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewportlayoutcontroller/textlayoutmanager
func (t_ TextViewportLayoutController) TextLayoutManager() NSTextLayoutManager {
	rv := objc.Send[NSTextLayoutManager](t_.ID, objc.Sel("textLayoutManager"))
	return rv
}


// Returns the text layout manager for this viewport layout controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewportlayoutcontroller/textlayoutmanager
func (t_ TextViewportLayoutController) SetTextLayoutManager(value ITextLayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextLayoutManager:"), value)
}


// Returns the visible bounds of the view, plus the overdraw area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewportlayoutcontroller/viewportbounds
func (t_ TextViewportLayoutController) ViewportBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("viewportBounds"))
	return rv
}


// Returns the visible bounds of the view, plus the overdraw area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewportlayoutcontroller/viewportbounds
func (t_ TextViewportLayoutController) SetViewportBounds(value coregraphics.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setViewportBounds:"), value)
}


// Returns the text range of the current viewport layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewportlayoutcontroller/viewportrange
func (t_ TextViewportLayoutController) ViewportRange() NSTextRange {
	rv := objc.Send[NSTextRange](t_.ID, objc.Sel("viewportRange"))
	return rv
}


// Returns the text range of the current viewport layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextviewportlayoutcontroller/viewportrange
func (t_ TextViewportLayoutController) SetViewportRange(value ITextRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setViewportRange:"), value)
}



