// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextViewportLayoutController] class.
var textViewportLayoutControllerClass = _TextViewportLayoutControllerClass{objc.GetClass("NSTextViewportLayoutController")}

type _TextViewportLayoutControllerClass struct {
	class objc.Class
}

// An interface definition for the [TextViewportLayoutController] class.
type ITextViewportLayoutController interface {
	objectivec.IObject
}

// Manages the layout process inside the viewport interacting with its delegate. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return textViewportLayoutControllerClass.New()
}




