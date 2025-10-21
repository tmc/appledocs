// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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



