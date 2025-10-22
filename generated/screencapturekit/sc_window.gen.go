// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	Active() bool
	Frame() coregraphics.CGRect
	SetFrame(value coregraphics.CGRect)
	IsActive() bool
	SetIsActive(value bool)
	IsOnScreen() bool
	SetIsOnScreen(value bool)
	OwningApplication() SCRunningApplication
	SetOwningApplication(value ISCRunningApplication)
	Title() string
	SetTitle(value string)
	WindowID() unsafe.Pointer
	SetWindowID(value unsafe.Pointer)
	WindowLayer() int
	SetWindowLayer(value int)
}

// An instance that represents an onscreen window.
//
// Retrieve the available windows from an instance of . Select one or more windows to capture and use them to create an instance of . Apply the filter to an instance of to limit its output to content matching your criteria.
//
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


// A Boolean value that indicates if the window is currently streaming.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/isActive
func (w_ Window) Active() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("active"))
	return rv
}

// A rectangle the represents the frame of the window within a display.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/frame
func (w_ Window) Frame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](w_.ID, objc.Sel("frame"))
	return rv
}


// SetFrame sets the value of the frame property.
// A rectangle the represents the frame of the window within a display.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/frame
func (w_ Window) SetFrame(value coregraphics.CGRect) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrame:"), value)
}

// A Boolean value that indicates if the window is currently streaming.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/isactive
func (w_ Window) IsActive() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isActive"))
	return rv
}


// SetIsActive sets the value of the isActive property.
// A Boolean value that indicates if the window is currently streaming.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/isactive
func (w_ Window) SetIsActive(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsActive:"), value)
}

// A Boolean value that indicates whether the window is on screen.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/isonscreen
func (w_ Window) IsOnScreen() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isOnScreen"))
	return rv
}


// SetIsOnScreen sets the value of the isOnScreen property.
// A Boolean value that indicates whether the window is on screen.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/isonscreen
func (w_ Window) SetIsOnScreen(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsOnScreen:"), value)
}

// The app that owns the window.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/owningapplication
func (w_ Window) OwningApplication() SCRunningApplication {
	rv := objc.Send[SCRunningApplication](w_.ID, objc.Sel("owningApplication"))
	return rv
}


// SetOwningApplication sets the value of the owningApplication property.
// The app that owns the window.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/owningapplication
func (w_ Window) SetOwningApplication(value ISCRunningApplication) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOwningApplication:"), value)
}

// The string that displays in a window’s title bar.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/title
func (w_ Window) Title() string {
	rv := objc.Send[string](w_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The string that displays in a window’s title bar.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/title
func (w_ Window) SetTitle(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// The Core Graphics window identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/windowid
func (w_ Window) WindowID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("windowID"))
	return rv
}


// SetWindowID sets the value of the windowID property.
// The Core Graphics window identifier.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/windowid
func (w_ Window) SetWindowID(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWindowID:"), value)
}

// The layer of the window relative to other windows.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/windowlayer
func (w_ Window) WindowLayer() int {
	rv := objc.Send[int](w_.ID, objc.Sel("windowLayer"))
	return rv
}


// SetWindowLayer sets the value of the windowLayer property.
// The layer of the window relative to other windows.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/windowlayer
func (w_ Window) SetWindowLayer(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWindowLayer:"), value)
}




