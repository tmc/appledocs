
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [WindowController] class.
var WindowControllerClass _WindowControllerClass

func init() {
	WindowControllerClass = _WindowControllerClass{objc.GetClass("NSWindowController")}
}

type _WindowControllerClass struct {
	objc.Class
}

// An interface definition for the [WindowController] class.
type IWindowController interface {
	ID() objc.ID
}

type WindowController struct {
	id objc.ID
}

func WindowControllerFrom(ptr unsafe.Pointer) WindowController {
	return WindowController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ WindowController) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _WindowControllerClass) Alloc() WindowController {
	rv := objc.Send[WindowController](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _WindowControllerClass) New() WindowController {
	rv := objc.Send[WindowController](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewWindowController creates and returns a new initialized instance.
func NewWindowController() WindowController {
	return WindowControllerClass.New()
}

// Init initializes the instance.
func (w_ WindowController) Init() WindowController {
	rv := objc.Send[WindowController](w_.ID(), selInit)
	return rv
}
// The window owned by the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowController/window
func (w_ WindowController) Window() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("window"))
	return rv
}
// SetWindow sets the value of the window property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowController/window
func (w_ WindowController) SetWindow(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setWindow:"), value)
}
