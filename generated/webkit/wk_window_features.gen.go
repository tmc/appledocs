// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WindowFeatures] class.
var (
	WindowFeaturesClass     _WindowFeaturesClass
	WindowFeaturesClassOnce sync.Once
)

func getWindowFeaturesClass() _WindowFeaturesClass {
	WindowFeaturesClassOnce.Do(func() {
		WindowFeaturesClass = _WindowFeaturesClass{objc.GetClass("WKWindowFeatures")}
	})
	return WindowFeaturesClass
}

type _WindowFeaturesClass struct {
	class objc.Class
}

// An interface definition for the [WindowFeatures] class.
type IWindowFeatures interface {
	objectivec.IObject
}

// Display-related attributes that a webpage requests for its window.
//
// A object contains the attributes that a webpage requests from its containing web view. You don’t create a object directly. When a navigation action results in the display of a new web view, creates this object and passes it to the method of its UI delegate object. The delegate uses the information in this object to configure and return the new web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWindowFeatures
type WindowFeatures struct {
	objectivec.Object
}

// WindowFeaturesFrom constructs a [WindowFeatures] from an unsafe.Pointer.
//
// Display-related attributes that a webpage requests for its window.
func WindowFeaturesFrom(ptr unsafe.Pointer) WindowFeatures {
	return WindowFeatures{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WindowFeaturesClass) Alloc() WindowFeatures {
	rv := objc.Send[WindowFeatures](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WindowFeaturesClass) New() WindowFeatures {
	rv := objc.Send[WindowFeatures](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WindowFeatures) Init() WindowFeatures {
	rv := objc.Send[WindowFeatures](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WindowFeatures) Autorelease() WindowFeatures {
	rv := objc.Send[WindowFeatures](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWindowFeatures creates a new WindowFeatures instance.
func NewWindowFeatures() WindowFeatures {
	return getWindowFeaturesClass().New()
}


// The requested height of the containing window.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/height
func (w_ WindowFeatures) Height() foundation.Number {
	rv := objc.Send[foundation.Number](w_.ID, objc.Sel("height"))
	return rv
}

// A Boolean value that indicates whether to make the containing window window resizable.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/allowsresizing
func (w_ WindowFeatures) AllowsResizing() foundation.Number {
	rv := objc.Send[foundation.Number](w_.ID, objc.Sel("allowsResizing"))
	return rv
}


// SetAllowsResizing sets the value of the allowsResizing property.
// A Boolean value that indicates whether to make the containing window window resizable.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/allowsresizing
func (w_ WindowFeatures) SetAllowsResizing(value foundation.INumber) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsResizing:"), value)
}

// A Boolean value that indicates whether the webpage requests a visible menu bar.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/menubarvisibility
func (w_ WindowFeatures) MenuBarVisibility() foundation.Number {
	rv := objc.Send[foundation.Number](w_.ID, objc.Sel("menuBarVisibility"))
	return rv
}


// SetMenuBarVisibility sets the value of the menuBarVisibility property.
// A Boolean value that indicates whether the webpage requests a visible menu bar.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/menubarvisibility
func (w_ WindowFeatures) SetMenuBarVisibility(value foundation.INumber) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMenuBarVisibility:"), value)
}

// A Boolean value that indicates whether the webpage requested a visible status bar.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/statusbarvisibility
func (w_ WindowFeatures) StatusBarVisibility() foundation.Number {
	rv := objc.Send[foundation.Number](w_.ID, objc.Sel("statusBarVisibility"))
	return rv
}


// SetStatusBarVisibility sets the value of the statusBarVisibility property.
// A Boolean value that indicates whether the webpage requested a visible status bar.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/statusbarvisibility
func (w_ WindowFeatures) SetStatusBarVisibility(value foundation.INumber) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setStatusBarVisibility:"), value)
}

// A Boolean value that indicates whether the webpage requested a visible toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/toolbarsvisibility
func (w_ WindowFeatures) ToolbarsVisibility() foundation.Number {
	rv := objc.Send[foundation.Number](w_.ID, objc.Sel("toolbarsVisibility"))
	return rv
}


// SetToolbarsVisibility sets the value of the toolbarsVisibility property.
// A Boolean value that indicates whether the webpage requested a visible toolbar.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/toolbarsvisibility
func (w_ WindowFeatures) SetToolbarsVisibility(value foundation.INumber) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setToolbarsVisibility:"), value)
}

// The requested width of the containing window.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/width
func (w_ WindowFeatures) Width() foundation.Number {
	rv := objc.Send[foundation.Number](w_.ID, objc.Sel("width"))
	return rv
}


// SetWidth sets the value of the width property.
// The requested width of the containing window.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/width
func (w_ WindowFeatures) SetWidth(value foundation.INumber) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWidth:"), value)
}

// The requested x-coordinate of the containing window.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/x
func (w_ WindowFeatures) X() foundation.Number {
	rv := objc.Send[foundation.Number](w_.ID, objc.Sel("x"))
	return rv
}


// SetX sets the value of the x property.
// The requested x-coordinate of the containing window.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/x
func (w_ WindowFeatures) SetX(value foundation.INumber) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setX:"), value)
}

// The requested y-coordinate of the containing window.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/y
func (w_ WindowFeatures) Y() foundation.Number {
	rv := objc.Send[foundation.Number](w_.ID, objc.Sel("y"))
	return rv
}


// SetY sets the value of the y property.
// The requested y-coordinate of the containing window.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwindowfeatures/y
func (w_ WindowFeatures) SetY(value foundation.INumber) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setY:"), value)
}



