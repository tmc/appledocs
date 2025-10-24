// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class JSManagedValue */


/* debug [class_header]: Header for JSManagedValue */
// The class instance for the [JSManagedValue] class.
var (
	JSManagedValueClass     _JSManagedValueClass
	JSManagedValueClassOnce sync.Once
)

func getJSManagedValueClass() _JSManagedValueClass {
	JSManagedValueClassOnce.Do(func() {
		JSManagedValueClass = _JSManagedValueClass{objc.GetClass("JSManagedValue")}
	})
	return JSManagedValueClass
}

type _JSManagedValueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for JSManagedValue */
// An interface definition for the [JSManagedValue] class.
type IJSManagedValue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for JSManagedValue */
	// properties:
	Value() IJSValue
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for JSManagedValue */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for JSManagedValue */
// Alloc allocates a new instance without initialization.
func (jc _JSManagedValueClass) Alloc() JSManagedValue {
	rv := objc.Send[JSManagedValue](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (jc _JSManagedValueClass) New() JSManagedValue {
	rv := objc.Send[JSManagedValue](objc.ID(jc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (j_ JSManagedValue) Init() JSManagedValue {
	rv := objc.Send[JSManagedValue](j_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (j_ JSManagedValue) Autorelease() JSManagedValue {
	rv := objc.Send[JSManagedValue](j_.ID, objc.Sel("autorelease"))
	return rv
}

// NewJSManagedValue creates a new JSManagedValue instance.
func NewJSManagedValue() JSManagedValue {
	return getJSManagedValueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for JSManagedValue */
// A JavaScript value with conditional retain behavior to provide automatic memory management.
//
// The primary use case for a managed value is to store a JavaScript value in an Objective-C or Swift object that exports to JavaScript. A managed value’s behavior ensures retention of its underlying JavaScript value as long as either of the following conditions is true: The JavaScript value is reachable through the JavaScript object graph (that is, not subject to JavaScript garbage collection). The object is reachable through the Objective-C or Swift object graph, as you report to the JavaScriptCore virtual machine using the method. However, if neither of these conditions is true, the managed value sets its property to , releasing the underlying object.


// A JavaScript value with conditional retain behavior to provide automatic memory management.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue
type JSManagedValue struct {
	objectivec.Object
}

// JSManagedValueFrom constructs a [JSManagedValue] from an unsafe.Pointer.
//
// A JavaScript value with conditional retain behavior to provide automatic memory management.
func JSManagedValueFrom(ptr unsafe.Pointer) JSManagedValue {
	return JSManagedValue{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for JSManagedValue */

// Initializes a managed value with the specified JavaScript value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue/init(value:)
func NewJSManagedValueWithValue(value IJSValue) JSManagedValue {
	instance := getJSManagedValueClass().Alloc()
	rv := objc.Send[JSManagedValue](instance.ID, objc.Sel("initWithValue:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewJSManagedValueWithValue */


// Creates a managed value and associates it with an owner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue/init(value:andOwner:)
func NewJSManagedValueWithValueAndOwner(value IJSValue, owner objc.IObject) JSManagedValue {
	rv := objc.Send[JSManagedValue](objc.ID(getJSManagedValueClass().class), objc.Sel("managedValueWithValue:andOwner:"), value, owner)
	return rv
}/* debug [class_init_methods/constructor]: NewJSManagedValueWithValueAndOwner */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for JSManagedValue */

// Creates a managed value and associates it with an owner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue/init(value:andOwner:)
func (jc _JSManagedValueClass) ManagedValueWithValueAndOwner(value IJSValue, owner objc.IObject) JSManagedValue {
	rv := objc.Send[JSManagedValue](objc.ID(jc.class), objc.Sel("managedValueWithValue:andOwner:"), value, owner)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ManagedValueWithValueAndOwner) */


// Creates a managed value with the specified JavaScript value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue/managedValueWithValue:
func (jc _JSManagedValueClass) ManagedValueWithValue(value IJSValue) JSManagedValue {
	rv := objc.Send[JSManagedValue](objc.ID(jc.class), objc.Sel("managedValueWithValue:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ManagedValueWithValue) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for JSManagedValue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for JSManagedValue */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for JSManagedValue */

// The managed value’s underlying JavaScript value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue/value
func (j_ JSManagedValue) Value() IJSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class JSManagedValue */


