// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WebScriptObject */

/* debug [class_header]: Header for WebScriptObject */
// The class instance for the [WebScriptObject] class.
var (
	WebScriptObjectClass     _WebScriptObjectClass
	WebScriptObjectClassOnce sync.Once
)

func getWebScriptObjectClass() _WebScriptObjectClass {
	WebScriptObjectClassOnce.Do(func() {
		WebScriptObjectClass = _WebScriptObjectClass{objc.GetClass("WebScriptObject")}
	})
	return WebScriptObjectClass
}

type _WebScriptObjectClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for WebScriptObject */
// An interface definition for the [WebScriptObject] class.
type IWebScriptObject interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for WebScriptObject */
	// properties:
	WindowScriptObject() IWebScriptObject
	SetWindowScriptObject(value IWebScriptObject)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for WebScriptObject */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for WebScriptObject */
// Alloc allocates a new instance without initialization.
func (wc _WebScriptObjectClass) Alloc() WebScriptObject {
	rv := objc.Send[WebScriptObject](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebScriptObjectClass) New() WebScriptObject {
	rv := objc.Send[WebScriptObject](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebScriptObject) Init() WebScriptObject {
	rv := objc.Send[WebScriptObject](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebScriptObject) Autorelease() WebScriptObject {
	rv := objc.Send[WebScriptObject](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebScriptObject creates a new WebScriptObject instance.
func NewWebScriptObject() WebScriptObject {
	return getWebScriptObjectClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for WebScriptObject */
// A object is an Objective-C wrapper for a scripting object passed to your application from the scripting environment.
//
// You can not create a object directly. You get a window object by sending to your object. You can use key-value coding methods—for example, and —to get and set properties of a object. You can also access properties by index using the and methods. Use the method to remove a scripting object property. Not all properties and methods of a class are exported. Use the and methods to intercept access to properties that are not exported. Similarly, use the method to intercept method invocations that are not exported. If you want access to properties and methods defined in your own classes, use the methods in the WebScripting informal protocol to specify the properties and methods the class should export to WebKit’s JavaScript environment. Use the and methods to execute scripts in the scripting environment.

// A object is an Objective-C wrapper for a scripting object passed to your application from the scripting environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebScriptObject
type WebScriptObject struct {
	objectivec.Object
}

// WebScriptObjectFrom constructs a [WebScriptObject] from an unsafe.Pointer.
//
// A object is an Objective-C wrapper for a scripting object passed to your application from the scripting environment.
func WebScriptObjectFrom(ptr unsafe.Pointer) WebScriptObject {
	return WebScriptObject{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for WebScriptObject */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for WebScriptObject */

// Raises an exception in the current script execution context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebScriptObject/throwException(_:)
func (wc _WebScriptObjectClass) ThrowException(exceptionMessage objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](objc.ID(wc.class), objc.Sel("throwException:"), exceptionMessage)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=ThrowException) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for WebScriptObject */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for WebScriptObject */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for WebScriptObject */

// The receiver’s window object from the scripting environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/windowscriptobject
func (w_ WebScriptObject) WindowScriptObject() IWebScriptObject {
	rv := objc.Send[WebScriptObject](w_.ID, objc.Sel("windowScriptObject"))
	return rv
} /* debug [instance_properties/getter]: windowScriptObject */

// The receiver’s window object from the scripting environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webview-swift.class/windowscriptobject
func (w_ WebScriptObject) SetWindowScriptObject(value IWebScriptObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWindowScriptObject:"), value)
} /* debug [instance_properties/setter]: windowScriptObject */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WebScriptObject */
