// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [WindowController] class.
var (
	WindowControllerClass     _WindowControllerClass
	WindowControllerClassOnce sync.Once
)

func getWindowControllerClass() _WindowControllerClass {
	WindowControllerClassOnce.Do(func() {
		WindowControllerClass = _WindowControllerClass{objc.GetClass("NSWindowController")}
	})
	return WindowControllerClass
}

type _WindowControllerClass struct {
	class objc.Class
}

// An interface definition for the [WindowController] class.
type IWindowController interface {
	IResponder
}

// A controller that manages a window, usually a window stored in a nib file.
//
// Managing a window entails: Loading and displaying the window Closing the window when appropriate Customizing the window’s title Storing the window’s frame (size and location) in the defaults database Cascading the window in relation to other document windows of the app A window controller can manage a window by itself or as a role player in AppKit’s document-based architecture, which also includes and objects. In this architecture, a window controller is created and managed by a “document” (an instance of an subclass) and, in turn, keeps a reference to the document. The relationship between a window controller and a nib file is important. Although a window controller can manage a programmatically created window, it usually manages a window in a nib file. The nib file can contain other top-level objects, including other windows, but the window controller’s responsibility is this primary window. The window controller is usually the owner of the nib file, even when it is part of a document-based app. Regardless of who is the file’s owner, the window controller is responsible for freeing all top-level objects in the nib file it loads. For simple documents—that is, documents with only one nib file containing a window—you need to do little directly with ; AppKit creates one for you. However, if the default window controller is not sufficient, you can create a custom subclass of . For documents with multiple windows or panels, your document must create separate instances of (or of custom subclasses of ), one for each window or panel. An example is a CAD app that has different windows for side, top, and front views of drawn objects. What you do in your subclass determines whether the default or separately created and configured objects are used.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getWindowControllerClass().New()
}

// Returns a window controller initialized with a given window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/init(window:)
func NewWindowControllerWithWindow(window unsafe.Pointer) WindowController {
	instance := getWindowControllerClass().Alloc()
	rv := objc.Send[WindowController](instance.ID, objc.Sel("initWithWindow:"), window)
	rv.Autorelease()
	return rv
}

// The window owned by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/window
func (w_ WindowController) Window() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("window"))
	return rv
}

// SetWindow sets the value of the window property.
// The window owned by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowController/window
func (w_ WindowController) SetWindow(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWindow:"), value)
}
