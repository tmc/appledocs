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

/* debug [class.gen.go]: Generating class SCWindow */


/* debug [class_header]: Header for SCWindow */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Window */
// An interface definition for the [Window] class.
type IWindow interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Window */
	// properties:
	Frame() corefoundation.CGRect
	Active() bool
	OnScreen() bool
	OwningApplication() ISCRunningApplication
	Title() objc.IObject /* cross-framework: NSString */
	WindowID() WindowID /* not a class type */
	WindowLayer() int
	IsActive() bool
	SetIsActive(value bool)
	IsOnScreen() bool
	SetIsOnScreen(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Window */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Window */
// Alloc allocates a new instance without initialization.
func (wc _WindowClass) Alloc() Window {
	rv := objc.Send[Window](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Window */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Window *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Window */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Window */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Window */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Window */

// A rectangle the represents the frame of the window within a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/frame
func (w_ Window) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_properties/getter]: frame */


// A Boolean value that indicates if the window is currently streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/isActive
func (w_ Window) Active() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("active"))
	return rv
}/* debug [instance_properties/getter]: active */


// A Boolean value that indicates whether the window is on screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/isOnScreen
func (w_ Window) OnScreen() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("onScreen"))
	return rv
}/* debug [instance_properties/getter]: onScreen */


// The app that owns the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/owningApplication
func (w_ Window) OwningApplication() ISCRunningApplication {
	rv := objc.Send[RunningApplication](w_.ID, objc.Sel("owningApplication"))
	return rv
}/* debug [instance_properties/getter]: owningApplication */


// The string that displays in a window’s title bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/title
func (w_ Window) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The Core Graphics window identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/windowID
func (w_ Window) WindowID() WindowID /* not a class type */ {
	rv := objc.Send[WindowID](w_.ID, objc.Sel("windowID"))
	return rv
}/* debug [instance_properties/getter]: windowID */


// The layer of the window relative to other windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow/windowLayer
func (w_ Window) WindowLayer() int {
	rv := objc.Send[int](w_.ID, objc.Sel("windowLayer"))
	return rv
}/* debug [instance_properties/getter]: windowLayer */


// A Boolean value that indicates if the window is currently streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/isactive
func (w_ Window) IsActive() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isActive"))
	return rv
}/* debug [instance_properties/getter]: isActive */


// A Boolean value that indicates if the window is currently streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/isactive
func (w_ Window) SetIsActive(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsActive:"), value)
}/* debug [instance_properties/setter]: isActive */


// A Boolean value that indicates whether the window is on screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/isonscreen
func (w_ Window) IsOnScreen() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isOnScreen"))
	return rv
}/* debug [instance_properties/getter]: isOnScreen */


// A Boolean value that indicates whether the window is on screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scwindow/isonscreen
func (w_ Window) SetIsOnScreen(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsOnScreen:"), value)
}/* debug [instance_properties/setter]: isOnScreen */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SCWindow */



