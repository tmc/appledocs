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
	Context() JSContext
	IsArray() bool
	SetIsArray(value bool)
	IsBigInt() bool
	SetIsBigInt(value bool)
	IsBoolean() bool
	SetIsBoolean(value bool)
	IsDate() bool
	SetIsDate(value bool)
	IsNull() bool
	SetIsNull(value bool)
	IsNumber() bool
	SetIsNumber(value bool)
	IsObject() bool
	SetIsObject(value bool)
	IsString() bool
	SetIsString(value bool)
	IsSymbol() bool
	SetIsSymbol(value bool)
	IsUndefined() bool
	SetIsUndefined(value bool)
	JsValueRef() JSValueRef
	SetJsValueRef(value IJSValueRef)
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
func (j_ JSValue) Context() JSContext {
	rv := objc.Send[JSContext](j_.ID, objc.Sel("context"))
	return rv
}

// A Boolean value that indicates whether the instance is a JavaScript array value.
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isarray
func (j_ JSValue) IsArray() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isArray"))
	return rv
}


// SetIsArray sets the value of the isArray property.
// A Boolean value that indicates whether the instance is a JavaScript array value.

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isarray
func (j_ JSValue) SetIsArray(value bool) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setIsArray:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isbigint
func (j_ JSValue) IsBigInt() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isBigInt"))
	return rv
}


// SetIsBigInt sets the value of the isBigInt property.
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isbigint
func (j_ JSValue) SetIsBigInt(value bool) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setIsBigInt:"), value)
}

// A Boolean value that indicates whether the instance is a JavaScript Boolean value.
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isboolean
func (j_ JSValue) IsBoolean() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isBoolean"))
	return rv
}


// SetIsBoolean sets the value of the isBoolean property.
// A Boolean value that indicates whether the instance is a JavaScript Boolean value.

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isboolean
func (j_ JSValue) SetIsBoolean(value bool) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setIsBoolean:"), value)
}

// A Boolean value that indicates whether the instance is a JavaScript
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isdate
func (j_ JSValue) IsDate() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isDate"))
	return rv
}


// SetIsDate sets the value of the isDate property.
// A Boolean value that indicates whether the instance is a JavaScript

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isdate
func (j_ JSValue) SetIsDate(value bool) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setIsDate:"), value)
}

// A Boolean value that indicates whether the instance corresponds to the JavaScript
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isnull
func (j_ JSValue) IsNull() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isNull"))
	return rv
}


// SetIsNull sets the value of the isNull property.
// A Boolean value that indicates whether the instance corresponds to the JavaScript

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isnull
func (j_ JSValue) SetIsNull(value bool) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setIsNull:"), value)
}

// A Boolean value that indicates whether the instance is a JavaScript numeric value.
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isnumber
func (j_ JSValue) IsNumber() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isNumber"))
	return rv
}


// SetIsNumber sets the value of the isNumber property.
// A Boolean value that indicates whether the instance is a JavaScript numeric value.

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isnumber
func (j_ JSValue) SetIsNumber(value bool) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setIsNumber:"), value)
}

// A Boolean value that indicates whether the instance is a JavaScript object.
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isobject
func (j_ JSValue) IsObject() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isObject"))
	return rv
}


// SetIsObject sets the value of the isObject property.
// A Boolean value that indicates whether the instance is a JavaScript object.

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isobject
func (j_ JSValue) SetIsObject(value bool) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setIsObject:"), value)
}

// A Boolean value that indicates whether the instance is a JavaScript
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isstring
func (j_ JSValue) IsString() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isString"))
	return rv
}


// SetIsString sets the value of the isString property.
// A Boolean value that indicates whether the instance is a JavaScript

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isstring
func (j_ JSValue) SetIsString(value bool) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setIsString:"), value)
}

// A Boolean value that indicates whether the instance is a symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/issymbol
func (j_ JSValue) IsSymbol() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isSymbol"))
	return rv
}


// SetIsSymbol sets the value of the isSymbol property.
// A Boolean value that indicates whether the instance is a symbol.

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/issymbol
func (j_ JSValue) SetIsSymbol(value bool) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setIsSymbol:"), value)
}

// A Boolean value that indicates whether the instance corresponds to the JavaScript
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isundefined
func (j_ JSValue) IsUndefined() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isUndefined"))
	return rv
}


// SetIsUndefined sets the value of the isUndefined property.
// A Boolean value that indicates whether the instance corresponds to the JavaScript

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/isundefined
func (j_ JSValue) SetIsUndefined(value bool) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setIsUndefined:"), value)
}

// Returns the C representation of the JavaScript value.
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/jsvalueref
func (j_ JSValue) JsValueRef() JSValueRef {
	rv := objc.Send[JSValueRef](j_.ID, objc.Sel("jsValueRef"))
	return rv
}


// SetJsValueRef sets the value of the jsValueRef property.
// Returns the C representation of the JavaScript value.

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jsvalue/jsvalueref
func (j_ JSValue) SetJsValueRef(value IJSValueRef) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setJsValueRef:"), value)
}



