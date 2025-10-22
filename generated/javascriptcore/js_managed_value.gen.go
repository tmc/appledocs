// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [JSManagedValue] class.
type IJSManagedValue interface {
	objectivec.IObject
	Value() JSValue
}

// A JavaScript value with conditional retain behavior to provide automatic memory management.
//
// The primary use case for a managed value is to store a JavaScript value in an Objective-C or Swift object that exports to JavaScript. A managed value’s behavior ensures retention of its underlying JavaScript value as long as either of the following conditions is true: The JavaScript value is reachable through the JavaScript object graph (that is, not subject to JavaScript garbage collection). The object is reachable through the Objective-C or Swift object graph, as you report to the JavaScriptCore virtual machine using the method. However, if neither of these conditions is true, the managed value sets its property to , releasing the underlying object.
//
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

// Alloc allocates a new instance without initialization.
func (jc _JSManagedValueClass) Alloc() JSManagedValue {
	rv := objc.Send[JSManagedValue](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes a managed value with the specified JavaScript value.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue/init(value:)
func NewJSManagedValueWithValue(value IJSValue) JSManagedValue {
	instance := getJSManagedValueClass().Alloc()
	rv := objc.Send[JSManagedValue](instance.ID, objc.Sel("initWithValue:"), value)
	rv.Autorelease()
	return rv
}


// Creates a managed value with the specified JavaScript value.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue/managedValueWithValue:
func (jc _JSManagedValueClass) ManagedValueWithValue(value IJSValue) JSManagedValue {
	rv := objc.Send[JSManagedValue](objc.ID(jc.class), objc.Sel("managedValueWithValue:"), value)
	return rv
}

// The managed value’s underlying JavaScript value.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue/value
func (j_ JSManagedValue) Value() JSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("value"))
	return rv
}


