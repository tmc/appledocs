// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [JSContext] class.
var (
	JSContextClass     _JSContextClass
	JSContextClassOnce sync.Once
)

func getJSContextClass() _JSContextClass {
	JSContextClassOnce.Do(func() {
		JSContextClass = _JSContextClass{objc.GetClass("JSContext")}
	})
	return JSContextClass
}

type _JSContextClass struct {
	class objc.Class
}

// An interface definition for the [JSContext] class.
type IJSContext interface {
	objectivec.IObject
	// properties:
	Exception() IJSValue
	SetException(value IJSValue)
	ExceptionHandler() unsafe.Pointer
	SetExceptionHandler(value unsafe.Pointer)
	GlobalObject() IJSValue
	SetGlobalObject(value IJSValue)
	IsInspectable() bool
	SetIsInspectable(value bool)
	JsGlobalContextRef() unsafe.Pointer
	SetJsGlobalContextRef(value unsafe.Pointer)
	Name() string
	SetName(value string)
	VirtualMachine() IJSVirtualMachine
	SetVirtualMachine(value IJSVirtualMachine)
	// methods:
	ObjectForKeyedSubscript(key objectivec.IObject) IJSValue
}

// A JavaScript execution environment.
//
// You create and use JavaScript contexts to evaluate JavaScript scripts from Objective-C or Swift code; to access values that JavaScript defines or calculates; and to make native objects, methods, or functions accessible to JavaScript.


// A JavaScript execution environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext
type JSContext struct {
	objectivec.Object
}

// JSContextFrom constructs a [JSContext] from an unsafe.Pointer.
//
// A JavaScript execution environment.
func JSContextFrom(ptr unsafe.Pointer) JSContext {
	return JSContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (jc _JSContextClass) Alloc() JSContext {
	rv := objc.Send[JSContext](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (jc _JSContextClass) New() JSContext {
	rv := objc.Send[JSContext](objc.ID(jc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (j_ JSContext) Init() JSContext {
	rv := objc.Send[JSContext](j_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (j_ JSContext) Autorelease() JSContext {
	rv := objc.Send[JSContext](j_.ID, objc.Sel("autorelease"))
	return rv
}

// NewJSContext creates a new JSContext instance.
func NewJSContext() JSContext {
	return getJSContextClass().New()
}



// Returns the value of the specified JavaScript property in the context’s global object, allowing subscript getter syntax.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/objectForKeyedSubscript(_:)
func (j_ JSContext) ObjectForKeyedSubscript(key objectivec.IObject) IJSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}


// A JavaScript exception to be thrown in evaluation of the script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/exception
func (j_ JSContext) Exception() IJSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("exception"))
	return rv
}


// A JavaScript exception to be thrown in evaluation of the script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/exception
func (j_ JSContext) SetException(value IJSValue) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setException:"), value)
}


// A block to be invoked should evaluating a script result in a JavaScript exception being thrown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/exceptionhandler
func (j_ JSContext) ExceptionHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](j_.ID, objc.Sel("exceptionHandler"))
	return rv
}


// A block to be invoked should evaluating a script result in a JavaScript exception being thrown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/exceptionhandler
func (j_ JSContext) SetExceptionHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setExceptionHandler:"), value)
}


// The JavaScript global object associated with the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/globalobject
func (j_ JSContext) GlobalObject() IJSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("globalObject"))
	return rv
}


// The JavaScript global object associated with the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/globalobject
func (j_ JSContext) SetGlobalObject(value IJSValue) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setGlobalObject:"), value)
}


// A Boolean value that indicates whether you can inspect the JavaScript context with Safari Web Inspector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/isinspectable
func (j_ JSContext) IsInspectable() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isInspectable"))
	return rv
}


// A Boolean value that indicates whether you can inspect the JavaScript context with Safari Web Inspector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/isinspectable
func (j_ JSContext) SetIsInspectable(value bool) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setIsInspectable:"), value)
}


// Returns the C representation of the JavaScript context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/jsglobalcontextref
func (j_ JSContext) JsGlobalContextRef() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](j_.ID, objc.Sel("jsGlobalContextRef"))
	return rv
}


// Returns the C representation of the JavaScript context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/jsglobalcontextref
func (j_ JSContext) SetJsGlobalContextRef(value unsafe.Pointer) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setJsGlobalContextRef:"), value)
}


// A descriptive name for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/name
func (j_ JSContext) Name() string {
	rv := objc.Send[string](j_.ID, objc.Sel("name"))
	return rv
}


// A descriptive name for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/name
func (j_ JSContext) SetName(value string) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setName:"), objc.String(value))
}


// The JavaScript virtual machine to which the context belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/virtualmachine
func (j_ JSContext) VirtualMachine() IJSVirtualMachine {
	rv := objc.Send[JSVirtualMachine](j_.ID, objc.Sel("virtualMachine"))
	return rv
}


// The JavaScript virtual machine to which the context belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/virtualmachine
func (j_ JSContext) SetVirtualMachine(value IJSVirtualMachine) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setVirtualMachine:"), value)
}



