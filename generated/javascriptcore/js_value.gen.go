// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [JSValue] class.
var (
	JSValueClass     _JSValueClass
	JSValueClassOnce sync.Once
)

func getJSValueClass() _JSValueClass {
	JSValueClassOnce.Do(func() {
		JSValueClass = _JSValueClass{objc.GetClass("JSValue")}
	})
	return JSValueClass
}

type _JSValueClass struct {
	class objc.Class
}

// An interface definition for the [JSValue] class.
type IJSValue interface {
	objectivec.IObject
}

// A JavaScript value.
//
// You use the class to convert basic values, such as numbers and strings, between JavaScript and Objective-C or Swift representations to pass data between native code and JavaScript code. You can also use this class to create JavaScript objects that wrap native objects of custom classes or JavaScript functions with implementations that native methods or blocks provide. Each instance originates from a object that represents the JavaScript execution environment containing that value. The value holds a strong reference to its object — as long as it retains any value for a particular instance, that context remains alive. When you invoke an instance method on a object, and that method returns another object, the returned value belongs to the same context as the original value. Each JavaScript value also has an association (indirectly via the property) with a specific object that represents the underlying set of execution resources for its context. You can pass instances only to methods on and instances on the same virtual machine — attempting to pass a value to a different virtual machine raises an Objective-C exception.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue
type JSValue struct {
	objectivec.Object
}

// JSValueFrom constructs a [JSValue] from an unsafe.Pointer.
//
// A JavaScript value.
func JSValueFrom(ptr unsafe.Pointer) JSValue {
	return JSValue{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (jc _JSValueClass) Alloc() JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (jc _JSValueClass) New() JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (j_ JSValue) Init() JSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (j_ JSValue) Autorelease() JSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("autorelease"))
	return rv
}

// NewJSValue creates a new JSValue instance.
func NewJSValue() JSValue {
	return getJSValueClass().New()
}


// The JavaScript context hosting this value.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/context
func (j_ JSValue) Context() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](j_.ID, objc.Sel("context"))
	return rv
}



