// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// Manages the layout process inside the viewport interacting with its delegate.
//
// A viewport is a rectangular area within a flipped coordinate system expanding along the y-axis. With text contents, lines advance expanding the view in the current writing direction. The viewport defines the active area where the framework lays out text fragments. In most cases, the area corresponds to the user visible area with an additional over-scroll region.
//
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




