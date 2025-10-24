// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class JSContext */


/* debug [class_header]: Header for JSContext */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for JSContext */
// An interface definition for the [JSContext] class.
type IJSContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for JSContext */
	// properties:
	Exception() IJSValue
	SetException(value IJSValue)
	ExceptionHandler() unsafe.Pointer
	SetExceptionHandler(value unsafe.Pointer)
	GlobalObject() IJSValue
	Inspectable() bool
	SetInspectable(value bool)
	JSGlobalContextRef() JSGlobalContextRef /* typedef */
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	VirtualMachine() IJSVirtualMachine
	IsInspectable() bool
	SetIsInspectable(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for JSContext */
	// methods:
	EvaluateScript(script objc.IObject /* cross-framework: NSString */) IJSValue
	EvaluateScriptWithSourceURL(script objc.IObject /* cross-framework: NSString */, sourceURL objc.IObject /* cross-framework: NSURL */) IJSValue
	ObjectForKeyedSubscript(key objc.IObject) IJSValue
	SetObjectForKeyedSubscript(object objc.IObject, key unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for JSContext */
// Alloc allocates a new instance without initialization.
func (jc _JSContextClass) Alloc() JSContext {
	rv := objc.Send[JSContext](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for JSContext */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for JSContext */

// Creates a JavaScript context object from the equivalent C representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/init(JSGlobalContextRef:)
func NewJSContextWithJSGlobalContextRef(jsGlobalContextRef JSGlobalContextRef /* typedef */) JSContext {
	rv := objc.Send[JSContext](objc.ID(getJSContextClass().class), objc.Sel("contextWithJSGlobalContextRef:"), jsGlobalContextRef)
	return rv
}/* debug [class_init_methods/constructor]: NewJSContextWithJSGlobalContextRef */


// Creates a new JavaScript context associated with a specific virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/init(virtualMachine:)
func NewJSContextWithVirtualMachine(virtualMachine IJSVirtualMachine) JSContext {
	instance := getJSContextClass().Alloc()
	rv := objc.Send[JSContext](instance.ID, objc.Sel("initWithVirtualMachine:"), virtualMachine)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewJSContextWithVirtualMachine */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for JSContext */

// Returns the context currently executing JavaScript code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/current()
func (jc _JSContextClass) CurrentContext() JSContext {
	rv := objc.Send[JSContext](objc.ID(jc.class), objc.Sel("currentContext"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CurrentContext) */


// Returns the arguments to the current native callback from JavaScript code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/currentArguments()
func (jc _JSContextClass) CurrentArguments() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(jc.class), objc.Sel("currentArguments"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CurrentArguments) */


// Returns the currently executing JavaScript function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/currentCallee()
func (jc _JSContextClass) CurrentCallee() IJSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("currentCallee"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CurrentCallee) */


// Returns the value of the keyword in currently executing JavaScript code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/currentThis()
func (jc _JSContextClass) CurrentThis() IJSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("currentThis"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CurrentThis) */


// Creates a JavaScript context object from the equivalent C representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/init(JSGlobalContextRef:)
func (jc _JSContextClass) ContextWithJSGlobalContextRef(jsGlobalContextRef JSGlobalContextRef /* typedef */) JSContext {
	rv := objc.Send[JSContext](objc.ID(jc.class), objc.Sel("contextWithJSGlobalContextRef:"), jsGlobalContextRef)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContextWithJSGlobalContextRef) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for JSContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for JSContext */

// Executes the specified JavaScript code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/evaluateScript(_:)
func (j_ JSContext) EvaluateScript(script objc.IObject /* cross-framework: NSString */) IJSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("evaluateScript:"), script)
	return rv
}/* debug [instance_methods/method]: EvaluateScript */


// Executes the specified JavaScript code, treating the specified URL as its source location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/evaluateScript(_:withSourceURL:)
func (j_ JSContext) EvaluateScriptWithSourceURL(script objc.IObject /* cross-framework: NSString */, sourceURL objc.IObject /* cross-framework: NSURL */) IJSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("evaluateScript:withSourceURL:"), script, sourceURL)
	return rv
}/* debug [instance_methods/method]: EvaluateScriptWithSourceURL */


// Returns the value of the specified JavaScript property in the context’s global object, allowing subscript getter syntax.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/objectForKeyedSubscript(_:)
func (j_ JSContext) ObjectForKeyedSubscript(key objc.IObject) IJSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}/* debug [instance_methods/method]: ObjectForKeyedSubscript */


// Sets the specified JavaScript property of the context’s global object, allowing subscript setter syntax.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/setObject(_:forKeyedSubscript:)
func (j_ JSContext) SetObjectForKeyedSubscript(object objc.IObject, key unsafe.Pointer) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setObject:forKeyedSubscript:"), object, key)
}/* debug [instance_methods/method]: SetObjectForKeyedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for JSContext */

// A JavaScript exception to be thrown in evaluation of the script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/exception
func (j_ JSContext) Exception() IJSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("exception"))
	return rv
}/* debug [instance_properties/getter]: exception */


// A JavaScript exception to be thrown in evaluation of the script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/exception
func (j_ JSContext) SetException(value IJSValue) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setException:"), value)
}/* debug [instance_properties/setter]: exception */


// A block to be invoked should evaluating a script result in a JavaScript exception being thrown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/exceptionHandler
func (j_ JSContext) ExceptionHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](j_.ID, objc.Sel("exceptionHandler"))
	return rv
}/* debug [instance_properties/getter]: exceptionHandler */


// A block to be invoked should evaluating a script result in a JavaScript exception being thrown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/exceptionHandler
func (j_ JSContext) SetExceptionHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setExceptionHandler:"), value)
}/* debug [instance_properties/setter]: exceptionHandler */


// The JavaScript global object associated with the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/globalObject
func (j_ JSContext) GlobalObject() IJSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("globalObject"))
	return rv
}/* debug [instance_properties/getter]: globalObject */


// A Boolean value that indicates whether you can inspect the JavaScript context with Safari Web Inspector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/isInspectable
func (j_ JSContext) Inspectable() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("inspectable"))
	return rv
}/* debug [instance_properties/getter]: inspectable */


// A Boolean value that indicates whether you can inspect the JavaScript context with Safari Web Inspector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/isInspectable
func (j_ JSContext) SetInspectable(value bool) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setInspectable:"), value)
}/* debug [instance_properties/setter]: inspectable */


// Returns the C representation of the JavaScript context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/jsGlobalContextRef
func (j_ JSContext) JSGlobalContextRef() JSGlobalContextRef /* typedef */ {
	rv := objc.Send[unsafe.Pointer](j_.ID, objc.Sel("JSGlobalContextRef"))
	return rv
}/* debug [instance_properties/getter]: JSGlobalContextRef */


// A descriptive name for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/name
func (j_ JSContext) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](j_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// A descriptive name for the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/name
func (j_ JSContext) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// The JavaScript virtual machine to which the context belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContext/virtualMachine
func (j_ JSContext) VirtualMachine() IJSVirtualMachine {
	rv := objc.Send[JSVirtualMachine](j_.ID, objc.Sel("virtualMachine"))
	return rv
}/* debug [instance_properties/getter]: virtualMachine */


// A Boolean value that indicates whether you can inspect the JavaScript context with Safari Web Inspector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/isinspectable
func (j_ JSContext) IsInspectable() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isInspectable"))
	return rv
}/* debug [instance_properties/getter]: isInspectable */


// A Boolean value that indicates whether you can inspect the JavaScript context with Safari Web Inspector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/javascriptcore/jscontext/isinspectable
func (j_ JSContext) SetIsInspectable(value bool) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setIsInspectable:"), value)
}/* debug [instance_properties/setter]: isInspectable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class JSContext */


