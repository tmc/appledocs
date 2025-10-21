// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
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
	EvaluateScriptWithSourceURL(script appkit.string, sourceURL foundation.IURL) JSValue
	ObjectForKeyedSubscript(key objectivec.IObject) JSValue
}

// A JavaScript execution environment.
//
// You create and use JavaScript contexts to evaluate JavaScript scripts from Objective-C or Swift code; to access values that JavaScript defines or calculates; and to make native objects, methods, or functions accessible to JavaScript.
//
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


// Returns the currently executing JavaScript function.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/currentCallee()
func (jc _JSContextClass) CurrentCallee() JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("currentCallee"))
	return rv
}

// Executes the specified JavaScript code, treating the specified URL as its source location.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/evaluateScript(_:withSourceURL:)
func (j_ JSContext) EvaluateScriptWithSourceURL(script appkit.string, sourceURL foundation.IURL) JSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("evaluateScript:withSourceURL:"), script, sourceURL)
	return rv
}

// Returns the value of the specified JavaScript property in the context’s global object, allowing subscript getter syntax.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/objectForKeyedSubscript(_:)
func (j_ JSContext) ObjectForKeyedSubscript(key objectivec.IObject) JSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}

// A JavaScript exception to be thrown in evaluation of the script.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/exception
func (j_ JSContext) Exception() JSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("exception"))
	return rv
}


// SetException sets the value of the exception property.
// A JavaScript exception to be thrown in evaluation of the script.

//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/exception
func (j_ JSContext) SetException(value IJSValue) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setException:"), value)
}

// The JavaScript virtual machine to which the context belongs.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/virtualMachine
func (j_ JSContext) VirtualMachine() JSVirtualMachine {
	rv := objc.Send[JSVirtualMachine](j_.ID, objc.Sel("virtualMachine"))
	return rv
}

// A block to be invoked should evaluating a script result in a JavaScript exception being thrown.
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/exceptionhandler
func (j_ JSContext) ExceptionHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](j_.ID, objc.Sel("exceptionHandler"))
	return rv
}


// SetExceptionHandler sets the value of the exceptionHandler property.
// A block to be invoked should evaluating a script result in a JavaScript exception being thrown.

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/exceptionhandler
func (j_ JSContext) SetExceptionHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setExceptionHandler:"), value)
}

// The JavaScript global object associated with the context.
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/globalobject
func (j_ JSContext) GlobalObject() JSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("globalObject"))
	return rv
}


// SetGlobalObject sets the value of the globalObject property.
// The JavaScript global object associated with the context.

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/globalobject
func (j_ JSContext) SetGlobalObject(value IJSValue) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setGlobalObject:"), value)
}

// A Boolean value that indicates whether you can inspect the JavaScript context with Safari Web Inspector.
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/isinspectable
func (j_ JSContext) IsInspectable() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isInspectable"))
	return rv
}


// SetIsInspectable sets the value of the isInspectable property.
// A Boolean value that indicates whether you can inspect the JavaScript context with Safari Web Inspector.

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/isinspectable
func (j_ JSContext) SetIsInspectable(value bool) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setIsInspectable:"), value)
}

// Returns the C representation of the JavaScript context.
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/jsglobalcontextref
func (j_ JSContext) JsGlobalContextRef() JSGlobalContextRef {
	rv := objc.Send[JSGlobalContextRef](j_.ID, objc.Sel("jsGlobalContextRef"))
	return rv
}


// SetJsGlobalContextRef sets the value of the jsGlobalContextRef property.
// Returns the C representation of the JavaScript context.

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/jsglobalcontextref
func (j_ JSContext) SetJsGlobalContextRef(value IJSGlobalContextRef) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setJsGlobalContextRef:"), value)
}

// A descriptive name for the context.
//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/name
func (j_ JSContext) Name() appkit.string {
	rv := objc.Send[appkit.string](j_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// A descriptive name for the context.

//
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/name
func (j_ JSContext) SetName(value appkit.string) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setName:"), value)
}



