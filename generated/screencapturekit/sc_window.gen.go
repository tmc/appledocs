// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Window] class.
var (
	WindowClass     _WindowClass
	WindowClassOnce sync.Once
)

func getWindowClass() _WindowClass {
	WindowClassOnce.Do(func() {
		WindowClass = _WindowClass{objc.GetClass("SCWindow")}
	})
	return WindowClass
}

type _WindowClass struct {
	class objc.Class
}

// An interface definition for the [Window] class.
type IWindow interface {
	objectivec.IObject
	// properties:
	OnScreen() bool
	Frame() objc.IObject /* cross-framework: Rect */
	SetFrame(value objc.IObject /* cross-framework: Rect */)
	IsActive() bool
	SetIsActive(value bool)
	IsOnScreen() bool
	SetIsOnScreen(value bool)
	OwningApplication() ISCRunningApplication
	SetOwningApplication(value ISCRunningApplication)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	WindowID() WindowID /* not a class type */
	SetWindowID(value WindowID /* not a class type */)
	WindowLayer() int
	SetWindowLayer(value int)
	// methods:
}

// An instance that represents an onscreen window.
//
// Retrieve the available windows from an instance of . Select one or more windows to capture and use them to create an instance of . Apply the filter to an instance of to limit its output to content matching your criteria.


// An instance that represents an onscreen window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow
type Window struct {
	objectivec.Object
}

// WindowFrom constructs a [Window] from an unsafe.Pointer.
//
// An instance that represents an onscreen window.
func WindowFrom(ptr unsafe.Pointer) Window {
	return Window{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WindowClass) Alloc() Window {
	rv := objc.Send[Window](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WindowClass) New() Window {
	rv := objc.Send[Window](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ Window) Init() Window {
	rv := objc.Send[Window](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ Window) Autorelease() Window {
	rv := objc.Send[Window](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWindow creates a new Window instance.
func NewWindow() Window {
	return getWindowClass().New()
}



// A Boolean value that indicates whether the window is on screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/isOnScreen
func (w_ Window) OnScreen() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("onScreen"))
	return rv
}


// A rectangle the represents the frame of the window within a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/frame
func (w_ Window) Frame() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](w_.ID, objc.Sel("frame"))
	return rv
}


// A rectangle the represents the frame of the window within a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/frame
func (w_ Window) SetFrame(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrame:"), value)
}


// A Boolean value that indicates if the window is currently streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/isactive
func (w_ Window) IsActive() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isActive"))
	return rv
}


// A Boolean value that indicates if the window is currently streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/isactive
func (w_ Window) SetIsActive(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsActive:"), value)
}


// A Boolean value that indicates whether the window is on screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/isonscreen
func (w_ Window) IsOnScreen() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isOnScreen"))
	return rv
}


// A Boolean value that indicates whether the window is on screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/isonscreen
func (w_ Window) SetIsOnScreen(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsOnScreen:"), value)
}


// The app that owns the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/owningapplication
func (w_ Window) OwningApplication() ISCRunningApplication {
	rv := objc.Send[RunningApplication](w_.ID, objc.Sel("owningApplication"))
	return rv
}


// The app that owns the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/owningapplication
func (w_ Window) SetOwningApplication(value ISCRunningApplication) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOwningApplication:"), value)
}


// The string that displays in a window’s title bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/title
func (w_ Window) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("title"))
	return rv
}


// The string that displays in a window’s title bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/title
func (w_ Window) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitle:"), value)
}


// The Core Graphics window identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/windowid
func (w_ Window) WindowID() WindowID /* not a class type */ {
	rv := objc.Send[WindowID](w_.ID, objc.Sel("windowID"))
	return rv
}


// The Core Graphics window identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/windowid
func (w_ Window) SetWindowID(value WindowID /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWindowID:"), value)
}


// The layer of the window relative to other windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/windowlayer
func (w_ Window) WindowLayer() int {
	rv := objc.Send[int](w_.ID, objc.Sel("windowLayer"))
	return rv
}


// The layer of the window relative to other windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/windowlayer
func (w_ Window) SetWindowLayer(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWindowLayer:"), value)
}




