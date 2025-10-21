// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	EvaluateScriptWithSourceURL(script string, sourceURL unsafe.Pointer) unsafe.Pointer
	ObjectForKeyedSubscript(key objc.ID) unsafe.Pointer
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
func (jc _JSContextClass) CurrentCallee() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(jc.class), objc.Sel("currentCallee"))
	return rv
}

// Executes the specified JavaScript code, treating the specified URL as its source location.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/evaluateScript(_:withSourceURL:)
func (j_ JSContext) EvaluateScriptWithSourceURL(script string, sourceURL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](j_.ID, objc.Sel("evaluateScript:withSourceURL:"), objc.String(script), sourceURL)
	return rv
}

// Returns the value of the specified JavaScript property in the context’s global object, allowing subscript getter syntax.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/objectForKeyedSubscript(_:)
func (j_ JSContext) ObjectForKeyedSubscript(key objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](j_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}

// A JavaScript exception to be thrown in evaluation of the script.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/exception
func (j_ JSContext) Exception() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](j_.ID, objc.Sel("exception"))
	return rv
}


// SetException sets the value of the exception property.
// A JavaScript exception to be thrown in evaluation of the script.

//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/exception
func (j_ JSContext) SetException(value unsafe.Pointer) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setException:"), value)
}

// The JavaScript virtual machine to which the context belongs.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/virtualMachine
func (j_ JSContext) VirtualMachine() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](j_.ID, objc.Sel("virtualMachine"))
	return rv
}



