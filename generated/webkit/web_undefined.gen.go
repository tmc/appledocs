// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WebUndefined */

/* debug [class_header]: Header for WebUndefined */
// The class instance for the [WebUndefined] class.
var (
	WebUndefinedClass     _WebUndefinedClass
	WebUndefinedClassOnce sync.Once
)

func getWebUndefinedClass() _WebUndefinedClass {
	WebUndefinedClassOnce.Do(func() {
		WebUndefinedClass = _WebUndefinedClass{objc.GetClass("WebUndefined")}
	})
	return WebUndefinedClass
}

type _WebUndefinedClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for WebUndefined */
// An interface definition for the [WebUndefined] class.
type IWebUndefined interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for WebUndefined */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for WebUndefined */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for WebUndefined */
// Alloc allocates a new instance without initialization.
func (wc _WebUndefinedClass) Alloc() WebUndefined {
	rv := objc.Send[WebUndefined](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebUndefinedClass) New() WebUndefined {
	rv := objc.Send[WebUndefined](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebUndefined) Init() WebUndefined {
	rv := objc.Send[WebUndefined](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebUndefined) Autorelease() WebUndefined {
	rv := objc.Send[WebUndefined](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebUndefined creates a new WebUndefined instance.
func NewWebUndefined() WebUndefined {
	return getWebUndefinedClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for WebUndefined */
// objects are simply used to represent the JavaScript “undefined” value in methods when bridging between JavaScript and Objective-C. For example, if you invoke a JavaScript function that returns the JavaScript “undefined” value, then a object is returned to the Objective-C calling context.

// objects are simply used to represent the JavaScript “undefined” value in methods when bridging between JavaScript and Objective-C. For example, if you invoke a JavaScript function that returns the JavaScript “undefined” value, then a object is returned to the Objective-C calling context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebUndefined
type WebUndefined struct {
	objectivec.Object
}

// WebUndefinedFrom constructs a [WebUndefined] from an unsafe.Pointer.
//
// objects are simply used to represent the JavaScript “undefined” value in methods when bridging between JavaScript and Objective-C. For example, if you invoke a JavaScript function that returns the JavaScript “undefined” value, then a object is returned to the Objective-C calling context.
func WebUndefinedFrom(ptr unsafe.Pointer) WebUndefined {
	return WebUndefined{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for WebUndefined */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for WebUndefined */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for WebUndefined */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for WebUndefined */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for WebUndefined */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WebUndefined */
