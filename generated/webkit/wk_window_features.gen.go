// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWindowFeatures */

/* debug [class_header]: Header for WKWindowFeatures */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for WindowFeatures */
// An interface definition for the [WindowFeatures] class.
type IWindowFeatures interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for WindowFeatures */
	// properties:
	AllowsResizing() objc.IObject      /* cross-framework: NSNumber */
	Height() objc.IObject              /* cross-framework: NSNumber */
	MenuBarVisibility() objc.IObject   /* cross-framework: NSNumber */
	StatusBarVisibility() objc.IObject /* cross-framework: NSNumber */
	ToolbarsVisibility() objc.IObject  /* cross-framework: NSNumber */
	Width() objc.IObject               /* cross-framework: NSNumber */
	X() objc.IObject                   /* cross-framework: NSNumber */
	Y() objc.IObject                   /* cross-framework: NSNumber */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for WindowFeatures */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for WindowFeatures */
// Alloc allocates a new instance without initialization.
func (wc _WindowFeaturesClass) Alloc() WindowFeatures {
	rv := objc.Send[WindowFeatures](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for WindowFeatures */
// Display-related attributes that a webpage requests for its window.
//
// A object contains the attributes that a webpage requests from its containing web view. You don’t create a object directly. When a navigation action results in the display of a new web view, creates this object and passes it to the method of its UI delegate object. The delegate uses the information in this object to configure and return the new web view.

// Display-related attributes that a webpage requests for its window.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for WindowFeatures */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for WindowFeatures */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for WindowFeatures */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for WindowFeatures */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for WindowFeatures */

// A Boolean value that indicates whether to make the containing window window resizable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/allowsResizing
func (w_ WindowFeatures) AllowsResizing() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](w_.ID, objc.Sel("allowsResizing"))
	return rv
} /* debug [instance_properties/getter]: allowsResizing */

// The requested height of the containing window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/height
func (w_ WindowFeatures) Height() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](w_.ID, objc.Sel("height"))
	return rv
} /* debug [instance_properties/getter]: height */

// A Boolean value that indicates whether the webpage requests a visible menu bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/menuBarVisibility
func (w_ WindowFeatures) MenuBarVisibility() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](w_.ID, objc.Sel("menuBarVisibility"))
	return rv
} /* debug [instance_properties/getter]: menuBarVisibility */

// A Boolean value that indicates whether the webpage requested a visible status bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/statusBarVisibility
func (w_ WindowFeatures) StatusBarVisibility() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](w_.ID, objc.Sel("statusBarVisibility"))
	return rv
} /* debug [instance_properties/getter]: statusBarVisibility */

// A Boolean value that indicates whether the webpage requested a visible toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/toolbarsVisibility
func (w_ WindowFeatures) ToolbarsVisibility() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](w_.ID, objc.Sel("toolbarsVisibility"))
	return rv
} /* debug [instance_properties/getter]: toolbarsVisibility */

// The requested width of the containing window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/width
func (w_ WindowFeatures) Width() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](w_.ID, objc.Sel("width"))
	return rv
} /* debug [instance_properties/getter]: width */

// The requested x-coordinate of the containing window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/x
func (w_ WindowFeatures) X() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](w_.ID, objc.Sel("x"))
	return rv
} /* debug [instance_properties/getter]: x */

// The requested y-coordinate of the containing window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/y
func (w_ WindowFeatures) Y() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](w_.ID, objc.Sel("y"))
	return rv
} /* debug [instance_properties/getter]: y */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKWindowFeatures */
