// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
func (w_ WindowFeatures) Height() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("height"))
	return rv
}



