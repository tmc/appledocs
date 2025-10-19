// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [WindowController] class.
var windowControllerClass = _WindowControllerClass{objc.GetClass("NSWindowController")}

type _WindowControllerClass struct {
	class objc.Class
}

// An interface definition for the [WindowController] class.
type IWindowController interface {
	IResponder
}

// A controller that manages a window, usually a window stored in a nib file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController

type WindowController struct {
	Responder
}

// WindowControllerFrom constructs a [WindowController] from an unsafe.Pointer.
//
// A controller that manages a window, usually a window stored in a nib file.
func WindowControllerFrom(ptr unsafe.Pointer) WindowController {
	return WindowController{
		Responder: ResponderFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (wc _WindowControllerClass) Alloc() WindowController {
	rv := objc.Send[WindowController](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (wc _WindowControllerClass) New() WindowController {
	rv := objc.Send[WindowController](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WindowController) Init() WindowController {
	rv := objc.Send[WindowController](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WindowController) Autorelease() WindowController {
	rv := objc.Send[WindowController](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWindowController creates a new WindowController instance.
func NewWindowController() WindowController {
	return windowControllerClass.New()
}




