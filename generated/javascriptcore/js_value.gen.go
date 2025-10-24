// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class JSValue */


/* debug [class_header]: Header for JSValue */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for JSValue */
// An interface definition for the [JSValue] class.
type IJSValue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for JSValue */
	// properties:
	Context() IJSContext
	IsArray() bool
	IsBigInt() bool
	IsBoolean() bool
	IsDate() bool
	IsNull() bool
	IsNumber() bool
	IsObject() bool
	IsString() bool
	IsSymbol() bool
	IsUndefined() bool
	JSValueRef() JSValueRef /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for JSValue */
	// methods:
	ValueAtIndex(index uint) IJSValue
	CallWithArguments(arguments objc.IObject /* cross-framework: NSArray */) IJSValue
	CompareDouble(other float64) JSRelationCondition
	CompareJSValue(other IJSValue) JSRelationCondition
	CompareUInt64(other uint64) JSRelationCondition
	CompareInt64(other int64) JSRelationCondition
	ConstructWithArguments(arguments objc.IObject /* cross-framework: NSArray */) IJSValue
	DefinePropertyDescriptor(property JSValueProperty /* typedef */, descriptor objc.IObject)
	DeleteProperty(property JSValueProperty /* typedef */) bool
	ValueForProperty(property JSValueProperty /* typedef */) IJSValue
	HasProperty(property JSValueProperty /* typedef */) bool
	InvokeMethodWithArguments(method objc.IObject /* cross-framework: NSString */, arguments objc.IObject /* cross-framework: NSArray */) IJSValue
	IsEqualToObject(value objc.IObject) bool
	IsEqualWithTypeCoercionToObject(value objc.IObject) bool
	IsInstanceOf(value objc.IObject) bool
	ObjectAtIndexedSubscript(index uint) IJSValue
	ObjectForKeyedSubscript(key objc.IObject) IJSValue
	SetObjectAtIndexedSubscript(object objc.IObject, index uint)
	SetObjectForKeyedSubscript(object objc.IObject, key objc.IObject)
	SetValueAtIndex(value objc.IObject, index uint)
	SetValueForProperty(value objc.IObject, property JSValueProperty /* typedef */)
	ToArray() foundation.Array
	ToBool() bool
	ToDate() foundation.Date
	ToDictionary() foundation.Dictionary
	ToDouble() float64
	ToInt32() int32 /* not a class type */
	ToInt64() int64
	ToNumber() foundation.Number
	ToObject() objc.ID
	ToObjectOfClass(expectedClass objc.Class) objc.ID
	ToPoint() corefoundation.CGPoint
	ToRange() corefoundation.Range
	ToRect() corefoundation.CGRect
	ToSize() corefoundation.CGSize
	ToString() foundation.String
	ToUInt32() uint32 /* not a class type */
	ToUInt64() uint64
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for JSValue */
// Alloc allocates a new instance without initialization.
func (jc _JSValueClass) Alloc() JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for JSValue */
// A JavaScript value.
//
// You use the class to convert basic values, such as numbers and strings, between JavaScript and Objective-C or Swift representations to pass data between native code and JavaScript code. You can also use this class to create JavaScript objects that wrap native objects of custom classes or JavaScript functions with implementations that native methods or blocks provide. Each instance originates from a object that represents the JavaScript execution environment containing that value. The value holds a strong reference to its object — as long as it retains any value for a particular instance, that context remains alive. When you invoke an instance method on a object, and that method returns another object, the returned value belongs to the same context as the original value. Each JavaScript value also has an association (indirectly via the property) with a specific object that represents the underlying set of execution resources for its context. You can pass instances only to methods on and instances on the same virtual machine — attempting to pass a value to a different virtual machine raises an Objective-C exception.


// A JavaScript value.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for JSValue */

// Creates a JavaScript representation of the specified Boolean value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(bool:in:)
func NewJSValueWithBoolInContext(value bool, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithBool:inContext:"), value, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithBoolInContext */


// Creates a JavaScript representation of the specified floating-point value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(double:in:)
func NewJSValueWithDoubleInContext(value float64, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithDouble:inContext:"), value, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithDoubleInContext */


// Creates a JavaScript representation of the specified signed integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(int32:in:)
func NewJSValueWithInt32InContext(value int32 /* not a class type */, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithInt32:inContext:"), value, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithInt32InContext */


// Creates a JavaScript value object from the equivalent C representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(JSValueRef:inContext:)
func NewJSValueWithJSValueRefInContext(value JSValueRef /* typedef */, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithJSValueRef:inContext:"), value, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithJSValueRefInContext */


// Creates a new, empty JavaScript array value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newArrayIn:)
func NewJSValueWithNewArrayInContext(context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewArrayInContext:"), context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithNewArrayInContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newBigIntFrom:in:)-r38z
func NewJSValueWithNewBigIntFromDoubleInContext(value float64, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewBigIntFromDouble:inContext:"), value, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithNewBigIntFromDoubleInContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newBigIntFrom:in:)-8l9iv
func NewJSValueWithNewBigIntFromInt64InContext(int64_ int64, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewBigIntFromInt64:inContext:"), int64_, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithNewBigIntFromInt64InContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newBigIntFrom:in:)-1f0xs
func NewJSValueWithNewBigIntFromStringInContext(string_ objc.IObject /* cross-framework: NSString */, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewBigIntFromString:inContext:"), string_, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithNewBigIntFromStringInContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newBigIntFrom:in:)-7worq
func NewJSValueWithNewBigIntFromUInt64InContext(uint64_ uint64, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewBigIntFromUInt64:inContext:"), uint64_, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithNewBigIntFromUInt64InContext */


// Creates a JavaScript error value with the specified error message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newErrorFromMessage:in:)
func NewJSValueWithNewErrorFromMessageInContext(message objc.IObject /* cross-framework: NSString */, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewErrorFromMessage:inContext:"), message, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithNewErrorFromMessageInContext */


// Creates a new, empty JavaScript object value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newObjectIn:)
func NewJSValueWithNewObjectInContext(context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewObjectInContext:"), context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithNewObjectInContext */


// Creates a promise object using the specified executor callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newPromiseIn:fromExecutor:)
func NewJSValueWithNewPromiseInContextFromExecutor(context IJSContext, callback unsafe.Pointer) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewPromiseInContext:fromExecutor:"), context, callback)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithNewPromiseInContextFromExecutor */


// Creates a rejected promise object with the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newPromiseRejectedWithReason:in:)
func NewJSValueWithNewPromiseRejectedWithReasonInContext(reason objc.IObject, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewPromiseRejectedWithReason:inContext:"), reason, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithNewPromiseRejectedWithReasonInContext */


// Creates a resolved promise object with the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newPromiseResolvedWithResult:in:)
func NewJSValueWithNewPromiseResolvedWithResultInContext(result objc.IObject, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewPromiseResolvedWithResult:inContext:"), result, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithNewPromiseResolvedWithResultInContext */


// Creates a JavaScript regular expression value from the specified pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newRegularExpressionFromPattern:flags:in:)
func NewJSValueWithNewRegularExpressionFromPatternFlagsInContext(pattern objc.IObject /* cross-framework: NSString */, flags objc.IObject /* cross-framework: NSString */, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewRegularExpressionFromPattern:flags:inContext:"), pattern, flags, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithNewRegularExpressionFromPatternFlagsInContext */


// Creates a unique symbol object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newSymbolFromDescription:in:)
func NewJSValueWithNewSymbolFromDescriptionInContext(description objc.IObject /* cross-framework: NSString */, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewSymbolFromDescription:inContext:"), description, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithNewSymbolFromDescriptionInContext */


// Creates a JavaScript value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(nullIn:)
func NewJSValueWithNullInContext(context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithNullInContext:"), context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithNullInContext */


// Creates a JavaScript value by converting the specified native object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(object:in:)
func NewJSValueWithObjectInContext(value objc.IObject, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithObject:inContext:"), value, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithObjectInContext */


// Creates a JavaScript representation of the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(point:inContext:)
func NewJSValueWithPointInContext(point corefoundation.CGPoint, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithPoint:inContext:"), point, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithPointInContext */


// Creates a JavaScript representation of the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(range:inContext:)
func NewJSValueWithRangeInContext(range_ corefoundation.Range, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithRange:inContext:"), range_, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithRangeInContext */


// Creates a JavaScript representation of the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(rect:inContext:)
func NewJSValueWithRectInContext(rect corefoundation.CGRect, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithRect:inContext:"), rect, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithRectInContext */


// Creates a JavaScript representation of the specified width and height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(size:inContext:)
func NewJSValueWithSizeInContext(size corefoundation.CGSize, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithSize:inContext:"), size, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithSizeInContext */


// Creates a JavaScript representation of the specified unsigned integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(uInt32:in:)
func NewJSValueWithUInt32InContext(value uint32 /* not a class type */, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithUInt32:inContext:"), value, context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithUInt32InContext */


// Creates a JavaScript value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(undefinedIn:)
func NewJSValueWithUndefinedInContext(context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(getJSValueClass().class), objc.Sel("valueWithUndefinedInContext:"), context)
	return rv
}/* debug [class_init_methods/constructor]: NewJSValueWithUndefinedInContext */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for JSValue */

// Creates a JavaScript representation of the specified Boolean value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(bool:in:)
func (jc _JSValueClass) ValueWithBoolInContext(value bool, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithBool:inContext:"), value, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithBoolInContext) */


// Creates a JavaScript representation of the specified floating-point value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(double:in:)
func (jc _JSValueClass) ValueWithDoubleInContext(value float64, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithDouble:inContext:"), value, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithDoubleInContext) */


// Creates a JavaScript representation of the specified signed integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(int32:in:)
func (jc _JSValueClass) ValueWithInt32InContext(value int32 /* not a class type */, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithInt32:inContext:"), value, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithInt32InContext) */


// Creates a JavaScript value object from the equivalent C representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(JSValueRef:inContext:)
func (jc _JSValueClass) ValueWithJSValueRefInContext(value JSValueRef /* typedef */, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithJSValueRef:inContext:"), value, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithJSValueRefInContext) */


// Creates a new, empty JavaScript array value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newArrayIn:)
func (jc _JSValueClass) ValueWithNewArrayInContext(context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithNewArrayInContext:"), context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithNewArrayInContext) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newBigIntFrom:in:)-1f0xs
func (jc _JSValueClass) ValueWithNewBigIntFromStringInContext(string_ objc.IObject /* cross-framework: NSString */, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithNewBigIntFromString:inContext:"), string_, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithNewBigIntFromStringInContext) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newBigIntFrom:in:)-7worq
func (jc _JSValueClass) ValueWithNewBigIntFromUInt64InContext(uint64_ uint64, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithNewBigIntFromUInt64:inContext:"), uint64_, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithNewBigIntFromUInt64InContext) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newBigIntFrom:in:)-8l9iv
func (jc _JSValueClass) ValueWithNewBigIntFromInt64InContext(int64_ int64, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithNewBigIntFromInt64:inContext:"), int64_, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithNewBigIntFromInt64InContext) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newBigIntFrom:in:)-r38z
func (jc _JSValueClass) ValueWithNewBigIntFromDoubleInContext(value float64, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithNewBigIntFromDouble:inContext:"), value, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithNewBigIntFromDoubleInContext) */


// Creates a JavaScript error value with the specified error message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newErrorFromMessage:in:)
func (jc _JSValueClass) ValueWithNewErrorFromMessageInContext(message objc.IObject /* cross-framework: NSString */, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithNewErrorFromMessage:inContext:"), message, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithNewErrorFromMessageInContext) */


// Creates a new, empty JavaScript object value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newObjectIn:)
func (jc _JSValueClass) ValueWithNewObjectInContext(context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithNewObjectInContext:"), context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithNewObjectInContext) */


// Creates a promise object using the specified executor callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newPromiseIn:fromExecutor:)
func (jc _JSValueClass) ValueWithNewPromiseInContextFromExecutor(context IJSContext, callback unsafe.Pointer) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithNewPromiseInContext:fromExecutor:"), context, callback)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithNewPromiseInContextFromExecutor) */


// Creates a rejected promise object with the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newPromiseRejectedWithReason:in:)
func (jc _JSValueClass) ValueWithNewPromiseRejectedWithReasonInContext(reason objc.IObject, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithNewPromiseRejectedWithReason:inContext:"), reason, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithNewPromiseRejectedWithReasonInContext) */


// Creates a resolved promise object with the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newPromiseResolvedWithResult:in:)
func (jc _JSValueClass) ValueWithNewPromiseResolvedWithResultInContext(result objc.IObject, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithNewPromiseResolvedWithResult:inContext:"), result, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithNewPromiseResolvedWithResultInContext) */


// Creates a JavaScript regular expression value from the specified pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newRegularExpressionFromPattern:flags:in:)
func (jc _JSValueClass) ValueWithNewRegularExpressionFromPatternFlagsInContext(pattern objc.IObject /* cross-framework: NSString */, flags objc.IObject /* cross-framework: NSString */, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithNewRegularExpressionFromPattern:flags:inContext:"), pattern, flags, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithNewRegularExpressionFromPatternFlagsInContext) */


// Creates a unique symbol object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newSymbolFromDescription:in:)
func (jc _JSValueClass) ValueWithNewSymbolFromDescriptionInContext(description objc.IObject /* cross-framework: NSString */, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithNewSymbolFromDescription:inContext:"), description, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithNewSymbolFromDescriptionInContext) */


// Creates a JavaScript value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(nullIn:)
func (jc _JSValueClass) ValueWithNullInContext(context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithNullInContext:"), context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithNullInContext) */


// Creates a JavaScript value by converting the specified native object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(object:in:)
func (jc _JSValueClass) ValueWithObjectInContext(value objc.IObject, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithObject:inContext:"), value, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithObjectInContext) */


// Creates a JavaScript representation of the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(point:inContext:)
func (jc _JSValueClass) ValueWithPointInContext(point corefoundation.CGPoint, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithPoint:inContext:"), point, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithPointInContext) */


// Creates a JavaScript representation of the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(range:inContext:)
func (jc _JSValueClass) ValueWithRangeInContext(range_ corefoundation.Range, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithRange:inContext:"), range_, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithRangeInContext) */


// Creates a JavaScript representation of the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(rect:inContext:)
func (jc _JSValueClass) ValueWithRectInContext(rect corefoundation.CGRect, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithRect:inContext:"), rect, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithRectInContext) */


// Creates a JavaScript representation of the specified width and height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(size:inContext:)
func (jc _JSValueClass) ValueWithSizeInContext(size corefoundation.CGSize, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithSize:inContext:"), size, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithSizeInContext) */


// Creates a JavaScript representation of the specified unsigned integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(uInt32:in:)
func (jc _JSValueClass) ValueWithUInt32InContext(value uint32 /* not a class type */, context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithUInt32:inContext:"), value, context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithUInt32InContext) */


// Creates a JavaScript value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(undefinedIn:)
func (jc _JSValueClass) ValueWithUndefinedInContext(context IJSContext) JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("valueWithUndefinedInContext:"), context)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueWithUndefinedInContext) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for JSValue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for JSValue */

// Returns the value at the specified numeric index in the JavaScript object value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/atIndex(_:)
func (j_ JSValue) ValueAtIndex(index uint) JSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("valueAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ValueAtIndex */


// Invokes the value as a JavaScript function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/call(withArguments:)
func (j_ JSValue) CallWithArguments(arguments objc.IObject /* cross-framework: NSArray */) JSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("callWithArguments:"), arguments)
	return rv
}/* debug [instance_methods/method]: CallWithArguments */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/compare(_:)-35b2t
func (j_ JSValue) CompareDouble(other float64) JSRelationCondition {
	rv := objc.Send[JSRelationCondition](j_.ID, objc.Sel("compareDouble:"), other)
	return rv
}/* debug [instance_methods/method]: CompareDouble */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/compare(_:)-5w184
func (j_ JSValue) CompareJSValue(other IJSValue) JSRelationCondition {
	rv := objc.Send[JSRelationCondition](j_.ID, objc.Sel("compareJSValue:"), other)
	return rv
}/* debug [instance_methods/method]: CompareJSValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/compare(_:)-64n3k
func (j_ JSValue) CompareUInt64(other uint64) JSRelationCondition {
	rv := objc.Send[JSRelationCondition](j_.ID, objc.Sel("compareUInt64:"), other)
	return rv
}/* debug [instance_methods/method]: CompareUInt64 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/compare(_:)-9d4zq
func (j_ JSValue) CompareInt64(other int64) JSRelationCondition {
	rv := objc.Send[JSRelationCondition](j_.ID, objc.Sel("compareInt64:"), other)
	return rv
}/* debug [instance_methods/method]: CompareInt64 */


// Invokes the value as a JavaScript constructor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/construct(withArguments:)
func (j_ JSValue) ConstructWithArguments(arguments objc.IObject /* cross-framework: NSArray */) JSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("constructWithArguments:"), arguments)
	return rv
}/* debug [instance_methods/method]: ConstructWithArguments */


// Defines a property on the JavaScript object value or modifies a property’s definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/defineProperty(_:descriptor:)
func (j_ JSValue) DefinePropertyDescriptor(property JSValueProperty /* typedef */, descriptor objc.IObject) {
	objc.Send[objc.ID](j_.ID, objc.Sel("defineProperty:descriptor:"), property, descriptor)
}/* debug [instance_methods/method]: DefinePropertyDescriptor */


// Deletes the named property from the JavaScript object value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/deleteProperty(_:)
func (j_ JSValue) DeleteProperty(property JSValueProperty /* typedef */) bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("deleteProperty:"), property)
	return rv
}/* debug [instance_methods/method]: DeleteProperty */


// Returns the value of the named property in the JavaScript object value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/forProperty(_:)
func (j_ JSValue) ValueForProperty(property JSValueProperty /* typedef */) JSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("valueForProperty:"), property)
	return rv
}/* debug [instance_methods/method]: ValueForProperty */


// Returns a Boolean value indicating whether the JavaScript value has a defined property with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/hasProperty(_:)
func (j_ JSValue) HasProperty(property JSValueProperty /* typedef */) bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("hasProperty:"), property)
	return rv
}/* debug [instance_methods/method]: HasProperty */


// Calls the named JavaScript method on the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/invokeMethod(_:withArguments:)
func (j_ JSValue) InvokeMethodWithArguments(method objc.IObject /* cross-framework: NSString */, arguments objc.IObject /* cross-framework: NSArray */) JSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("invokeMethod:withArguments:"), method, arguments)
	return rv
}/* debug [instance_methods/method]: InvokeMethodWithArguments */


// Compares the value to another for strict equality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isEqual(to:)
func (j_ JSValue) IsEqualToObject(value objc.IObject) bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isEqualToObject:"), value)
	return rv
}/* debug [instance_methods/method]: IsEqualToObject */


// Compares the value to another for equivalence, allowing type conversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isEqualWithTypeCoercion(to:)
func (j_ JSValue) IsEqualWithTypeCoercionToObject(value objc.IObject) bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isEqualWithTypeCoercionToObject:"), value)
	return rv
}/* debug [instance_methods/method]: IsEqualWithTypeCoercionToObject */


// Returns a Boolean value indicating whether the value is an instance of another JavaScript object value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isInstance(of:)
func (j_ JSValue) IsInstanceOf(value objc.IObject) bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isInstanceOf:"), value)
	return rv
}/* debug [instance_methods/method]: IsInstanceOf */


// Returns the value’s JavaScript property at the specified index, allowing subscript syntax.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/objectAtIndexedSubscript(_:)
func (j_ JSValue) ObjectAtIndexedSubscript(index uint) JSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */


// Returns the value’s JavaScript property named with the specified key, allowing subscript syntax.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/objectForKeyedSubscript(_:)
func (j_ JSValue) ObjectForKeyedSubscript(key objc.IObject) JSValue {
	rv := objc.Send[JSValue](j_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}/* debug [instance_methods/method]: ObjectForKeyedSubscript */


// Sets the value’s JavaScript property at the specified index, allowing subscript syntax.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/setObject(_:atIndexedSubscript:)
func (j_ JSValue) SetObjectAtIndexedSubscript(object objc.IObject, index uint) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setObject:atIndexedSubscript:"), object, index)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */


// Sets the value’s JavaScript property named with the specified key, allowing subscript syntax.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/setObject(_:forKeyedSubscript:)
func (j_ JSValue) SetObjectForKeyedSubscript(object objc.IObject, key objc.IObject) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setObject:forKeyedSubscript:"), object, key)
}/* debug [instance_methods/method]: SetObjectForKeyedSubscript */


// Sets the value at the specified numeric index in the JavaScript object value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/setValue(_:at:)
func (j_ JSValue) SetValueAtIndex(value objc.IObject, index uint) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setValue:atIndex:"), value, index)
}/* debug [instance_methods/method]: SetValueAtIndex */


// Sets the value of the named property in the JavaScript object value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/setValue(_:forProperty:)
func (j_ JSValue) SetValueForProperty(value objc.IObject, property JSValueProperty /* typedef */) {
	objc.Send[objc.ID](j_.ID, objc.Sel("setValue:forProperty:"), value, property)
}/* debug [instance_methods/method]: SetValueForProperty */


// Converts the JavaScript value to an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toArray()
func (j_ JSValue) ToArray() foundation.Array {
	rv := objc.Send[foundation.Array](j_.ID, objc.Sel("toArray"))
	return rv
}/* debug [instance_methods/method]: ToArray */


// Converts the JavaScript value to a native Boolean value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toBool()
func (j_ JSValue) ToBool() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("toBool"))
	return rv
}/* debug [instance_methods/method]: ToBool */


// Converts the JavaScript value to a date object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toDate()
func (j_ JSValue) ToDate() foundation.Date {
	rv := objc.Send[foundation.Date](j_.ID, objc.Sel("toDate"))
	return rv
}/* debug [instance_methods/method]: ToDate */


// Converts the JavaScript value to a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toDictionary()
func (j_ JSValue) ToDictionary() foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](j_.ID, objc.Sel("toDictionary"))
	return rv
}/* debug [instance_methods/method]: ToDictionary */


// Converts the JavaScript value to a native floating-point value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toDouble()
func (j_ JSValue) ToDouble() float64 {
	rv := objc.Send[float64](j_.ID, objc.Sel("toDouble"))
	return rv
}/* debug [instance_methods/method]: ToDouble */


// Converts the JavaScript value to a native signed integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toInt32()
func (j_ JSValue) ToInt32() int32 /* not a class type */ {
	rv := objc.Send[int32](j_.ID, objc.Sel("toInt32"))
	return rv
}/* debug [instance_methods/method]: ToInt32 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toInt64()
func (j_ JSValue) ToInt64() int64 {
	rv := objc.Send[int64](j_.ID, objc.Sel("toInt64"))
	return rv
}/* debug [instance_methods/method]: ToInt64 */


// Converts the JavaScript value to a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toNumber()
func (j_ JSValue) ToNumber() foundation.Number {
	rv := objc.Send[foundation.Number](j_.ID, objc.Sel("toNumber"))
	return rv
}/* debug [instance_methods/method]: ToNumber */


// Converts the JavaScript value to a native object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toObject()
func (j_ JSValue) ToObject() objc.ID {
	rv := objc.Send[objc.ID](j_.ID, objc.Sel("toObject"))
	return rv
}/* debug [instance_methods/method]: ToObject */


// Converts the JavaScript value to a native object of the specified class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toObjectOf(_:)
func (j_ JSValue) ToObjectOfClass(expectedClass objc.Class) objc.ID {
	rv := objc.Send[objc.ID](j_.ID, objc.Sel("toObjectOfClass:"), expectedClass)
	return rv
}/* debug [instance_methods/method]: ToObjectOfClass */


// Converts the value to a point structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toPoint()
func (j_ JSValue) ToPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](j_.ID, objc.Sel("toPoint"))
	return rv
}/* debug [instance_methods/method]: ToPoint */


// Converts the value to a range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toRange()
func (j_ JSValue) ToRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](j_.ID, objc.Sel("toRange"))
	return rv
}/* debug [instance_methods/method]: ToRange */


// Converts the value to a rectangle structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toRect()
func (j_ JSValue) ToRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](j_.ID, objc.Sel("toRect"))
	return rv
}/* debug [instance_methods/method]: ToRect */


// Converts the value to a size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toSize()
func (j_ JSValue) ToSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](j_.ID, objc.Sel("toSize"))
	return rv
}/* debug [instance_methods/method]: ToSize */


// Converts the JavaScript value to a native string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toString()
func (j_ JSValue) ToString() foundation.String {
	rv := objc.Send[foundation.String](j_.ID, objc.Sel("toString"))
	return rv
}/* debug [instance_methods/method]: ToString */


// Converts the JavaScript value to a native unsigned integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toUInt32()
func (j_ JSValue) ToUInt32() uint32 /* not a class type */ {
	rv := objc.Send[uint32](j_.ID, objc.Sel("toUInt32"))
	return rv
}/* debug [instance_methods/method]: ToUInt32 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toUInt64()
func (j_ JSValue) ToUInt64() uint64 {
	rv := objc.Send[uint64](j_.ID, objc.Sel("toUInt64"))
	return rv
}/* debug [instance_methods/method]: ToUInt64 */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for JSValue */

// The JavaScript context hosting this value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/context
func (j_ JSValue) Context() IJSContext {
	rv := objc.Send[JSContext](j_.ID, objc.Sel("context"))
	return rv
}/* debug [instance_properties/getter]: context */


// A Boolean value that indicates whether the instance is a JavaScript array value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isArray
func (j_ JSValue) IsArray() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isArray"))
	return rv
}/* debug [instance_properties/getter]: isArray */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isBigInt
func (j_ JSValue) IsBigInt() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isBigInt"))
	return rv
}/* debug [instance_properties/getter]: isBigInt */


// A Boolean value that indicates whether the instance is a JavaScript Boolean value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isBoolean
func (j_ JSValue) IsBoolean() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isBoolean"))
	return rv
}/* debug [instance_properties/getter]: isBoolean */


// A Boolean value that indicates whether the instance is a JavaScript object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isDate
func (j_ JSValue) IsDate() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isDate"))
	return rv
}/* debug [instance_properties/getter]: isDate */


// A Boolean value that indicates whether the instance corresponds to the JavaScript value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isNull
func (j_ JSValue) IsNull() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isNull"))
	return rv
}/* debug [instance_properties/getter]: isNull */


// A Boolean value that indicates whether the instance is a JavaScript numeric value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isNumber
func (j_ JSValue) IsNumber() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isNumber"))
	return rv
}/* debug [instance_properties/getter]: isNumber */


// A Boolean value that indicates whether the instance is a JavaScript object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isObject
func (j_ JSValue) IsObject() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isObject"))
	return rv
}/* debug [instance_properties/getter]: isObject */


// A Boolean value that indicates whether the instance is a JavaScript object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isString
func (j_ JSValue) IsString() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isString"))
	return rv
}/* debug [instance_properties/getter]: isString */


// A Boolean value that indicates whether the instance is a symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isSymbol
func (j_ JSValue) IsSymbol() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isSymbol"))
	return rv
}/* debug [instance_properties/getter]: isSymbol */


// A Boolean value that indicates whether the instance corresponds to the JavaScript value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isUndefined
func (j_ JSValue) IsUndefined() bool {
	rv := objc.Send[bool](j_.ID, objc.Sel("isUndefined"))
	return rv
}/* debug [instance_properties/getter]: isUndefined */


// Returns the C representation of the JavaScript value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValue/jsValueRef
func (j_ JSValue) JSValueRef() JSValueRef /* typedef */ {
	rv := objc.Send[unsafe.Pointer](j_.ID, objc.Sel("JSValueRef"))
	return rv
}/* debug [instance_properties/getter]: JSValueRef */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class JSValue */


