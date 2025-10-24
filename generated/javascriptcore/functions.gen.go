// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

/* debug [functions.gen.go]: Generating 119 functions for JavaScriptCore */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// JavaScriptCore Functions (119 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_JSBigIntCreateWithDouble func(JSContextRef, float64, unsafe.Pointer) JSValueRef
	_JSBigIntCreateWithInt64 func(JSContextRef, int64, unsafe.Pointer) JSValueRef
	_JSBigIntCreateWithString func(JSContextRef, JSStringRef, unsafe.Pointer) JSValueRef
	_JSBigIntCreateWithUInt64 func(JSContextRef, uint64, unsafe.Pointer) JSValueRef
	_JSValueCompare func(JSContextRef, JSValueRef, JSValueRef, unsafe.Pointer) JSRelationCondition
	_JSValueCompareDouble func(JSContextRef, JSValueRef, float64, unsafe.Pointer) JSRelationCondition
	_JSValueCompareInt64 func(JSContextRef, JSValueRef, int64, unsafe.Pointer) JSRelationCondition
	_JSValueCompareUInt64 func(JSContextRef, JSValueRef, uint64, unsafe.Pointer) JSRelationCondition
	_JSValueIsBigInt func(JSContextRef, JSValueRef) bool
	_JSValueToInt32 func(JSContextRef, JSValueRef, unsafe.Pointer) int32
	_JSValueToInt64 func(JSContextRef, JSValueRef, unsafe.Pointer) int64
	_JSValueToUInt32 func(JSContextRef, JSValueRef, unsafe.Pointer) uint32
	_JSValueToUInt64 func(JSContextRef, JSValueRef, unsafe.Pointer) uint64
	_JSCheckScriptSyntax func(JSContextRef, JSStringRef, JSStringRef, int, unsafe.Pointer) bool
	_JSClassCreate func(unsafe.Pointer) JSClassRef
	_JSClassRelease func(JSClassRef)
	_JSClassRetain func(JSClassRef) JSClassRef
	_JSContextGetGlobalContext func(JSContextRef) JSGlobalContextRef
	_JSContextGetGlobalObject func(JSContextRef) JSObjectRef
	_JSContextGetGroup func(JSContextRef) JSContextGroupRef
	_JSContextGroupCreate func() JSContextGroupRef
	_JSContextGroupRelease func(JSContextGroupRef)
	_JSContextGroupRetain func(JSContextGroupRef) JSContextGroupRef
	_JSEvaluateScript func(JSContextRef, JSStringRef, JSObjectRef, JSStringRef, int, unsafe.Pointer) JSValueRef
	_JSGarbageCollect func(JSContextRef)
	_JSGlobalContextCopyName func(JSGlobalContextRef) JSStringRef
	_JSGlobalContextCreate func(JSClassRef) JSGlobalContextRef
	_JSGlobalContextCreateInGroup func(JSContextGroupRef, JSClassRef) JSGlobalContextRef
	_JSGlobalContextIsInspectable func(JSGlobalContextRef) bool
	_JSGlobalContextRelease func(JSGlobalContextRef)
	_JSGlobalContextRetain func(JSGlobalContextRef) JSGlobalContextRef
	_JSGlobalContextSetInspectable func(JSGlobalContextRef, bool)
	_JSGlobalContextSetName func(JSGlobalContextRef, JSStringRef)
	_JSObjectCallAsConstructor func(JSContextRef, JSObjectRef, uintptr, unsafe.Pointer, unsafe.Pointer) JSObjectRef
	_JSObjectCallAsFunction func(JSContextRef, JSObjectRef, JSObjectRef, uintptr, unsafe.Pointer, unsafe.Pointer) JSValueRef
	_JSObjectCopyPropertyNames func(JSContextRef, JSObjectRef) JSPropertyNameArrayRef
	_JSObjectDeleteProperty func(JSContextRef, JSObjectRef, JSStringRef, unsafe.Pointer) bool
	_JSObjectDeletePropertyForKey func(JSContextRef, JSObjectRef, JSValueRef, unsafe.Pointer) bool
	_JSObjectGetArrayBufferByteLength func(JSContextRef, JSObjectRef, unsafe.Pointer) uintptr
	_JSObjectGetArrayBufferBytesPtr func(JSContextRef, JSObjectRef, unsafe.Pointer) unsafe.Pointer
	_JSObjectGetPrivate func(JSObjectRef) unsafe.Pointer
	_JSObjectGetProperty func(JSContextRef, JSObjectRef, JSStringRef, unsafe.Pointer) JSValueRef
	_JSObjectGetPropertyAtIndex func(JSContextRef, JSObjectRef, unsafe.Pointer, unsafe.Pointer) JSValueRef
	_JSObjectGetPropertyForKey func(JSContextRef, JSObjectRef, JSValueRef, unsafe.Pointer) JSValueRef
	_JSObjectGetPrototype func(JSContextRef, JSObjectRef) JSValueRef
	_JSObjectGetTypedArrayBuffer func(JSContextRef, JSObjectRef, unsafe.Pointer) JSObjectRef
	_JSObjectGetTypedArrayByteLength func(JSContextRef, JSObjectRef, unsafe.Pointer) uintptr
	_JSObjectGetTypedArrayByteOffset func(JSContextRef, JSObjectRef, unsafe.Pointer) uintptr
	_JSObjectGetTypedArrayBytesPtr func(JSContextRef, JSObjectRef, unsafe.Pointer) unsafe.Pointer
	_JSObjectGetTypedArrayLength func(JSContextRef, JSObjectRef, unsafe.Pointer) uintptr
	_JSObjectHasProperty func(JSContextRef, JSObjectRef, JSStringRef) bool
	_JSObjectHasPropertyForKey func(JSContextRef, JSObjectRef, JSValueRef, unsafe.Pointer) bool
	_JSObjectIsConstructor func(JSContextRef, JSObjectRef) bool
	_JSObjectIsFunction func(JSContextRef, JSObjectRef) bool
	_JSObjectMake func(JSContextRef, JSClassRef, unsafe.Pointer) JSObjectRef
	_JSObjectMakeArray func(JSContextRef, uintptr, unsafe.Pointer, unsafe.Pointer) JSObjectRef
	_JSObjectMakeArrayBufferWithBytesNoCopy func(JSContextRef, unsafe.Pointer, uintptr, JSTypedArrayBytesDeallocator, unsafe.Pointer, unsafe.Pointer) JSObjectRef
	_JSObjectMakeConstructor func(JSContextRef, JSClassRef, JSObjectCallAsConstructorCallback) JSObjectRef
	_JSObjectMakeDate func(JSContextRef, uintptr, unsafe.Pointer, unsafe.Pointer) JSObjectRef
	_JSObjectMakeDeferredPromise func(JSContextRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) JSObjectRef
	_JSObjectMakeError func(JSContextRef, uintptr, unsafe.Pointer, unsafe.Pointer) JSObjectRef
	_JSObjectMakeFunction func(JSContextRef, JSStringRef, unsafe.Pointer, unsafe.Pointer, JSStringRef, JSStringRef, int, unsafe.Pointer) JSObjectRef
	_JSObjectMakeFunctionWithCallback func(JSContextRef, JSStringRef, JSObjectCallAsFunctionCallback) JSObjectRef
	_JSObjectMakeRegExp func(JSContextRef, uintptr, unsafe.Pointer, unsafe.Pointer) JSObjectRef
	_JSObjectMakeTypedArray func(JSContextRef, unsafe.Pointer, uintptr, unsafe.Pointer) JSObjectRef
	_JSObjectMakeTypedArrayWithArrayBuffer func(JSContextRef, unsafe.Pointer, JSObjectRef, unsafe.Pointer) JSObjectRef
	_JSObjectMakeTypedArrayWithArrayBufferAndOffset func(JSContextRef, unsafe.Pointer, JSObjectRef, uintptr, uintptr, unsafe.Pointer) JSObjectRef
	_JSObjectMakeTypedArrayWithBytesNoCopy func(JSContextRef, unsafe.Pointer, unsafe.Pointer, uintptr, JSTypedArrayBytesDeallocator, unsafe.Pointer, unsafe.Pointer) JSObjectRef
	_JSObjectSetPrivate func(JSObjectRef, unsafe.Pointer) bool
	_JSObjectSetProperty func(JSContextRef, JSObjectRef, JSStringRef, JSValueRef, JSPropertyAttributes, unsafe.Pointer)
	_JSObjectSetPropertyAtIndex func(JSContextRef, JSObjectRef, unsafe.Pointer, JSValueRef, unsafe.Pointer)
	_JSObjectSetPropertyForKey func(JSContextRef, JSObjectRef, JSValueRef, JSValueRef, JSPropertyAttributes, unsafe.Pointer)
	_JSObjectSetPrototype func(JSContextRef, JSObjectRef, JSValueRef)
	_JSPropertyNameAccumulatorAddName func(JSPropertyNameAccumulatorRef, JSStringRef)
	_JSPropertyNameArrayGetCount func(JSPropertyNameArrayRef) uintptr
	_JSPropertyNameArrayGetNameAtIndex func(JSPropertyNameArrayRef, uintptr) JSStringRef
	_JSPropertyNameArrayRelease func(JSPropertyNameArrayRef)
	_JSPropertyNameArrayRetain func(JSPropertyNameArrayRef) JSPropertyNameArrayRef
	_JSStringCopyCFString func(AllocatorRef, JSStringRef) StringRef
	_JSStringCreateWithCFString func(StringRef) JSStringRef
	_JSStringCreateWithCharacters func(unsafe.Pointer, uintptr) JSStringRef
	_JSStringCreateWithUTF8CString func(unsafe.Pointer) JSStringRef
	_JSStringGetCharactersPtr func(JSStringRef) unsafe.Pointer
	_JSStringGetLength func(JSStringRef) uintptr
	_JSStringGetMaximumUTF8CStringSize func(JSStringRef) uintptr
	_JSStringGetUTF8CString func(JSStringRef, unsafe.Pointer, uintptr) uintptr
	_JSStringIsEqual func(JSStringRef, JSStringRef) bool
	_JSStringIsEqualToUTF8CString func(JSStringRef, unsafe.Pointer) bool
	_JSStringRelease func(JSStringRef)
	_JSStringRetain func(JSStringRef) JSStringRef
	_JSValueCreateJSONString func(JSContextRef, JSValueRef, unsafe.Pointer, unsafe.Pointer) JSStringRef
	_JSValueGetType func(JSContextRef, JSValueRef) unsafe.Pointer
	_JSValueGetTypedArrayType func(JSContextRef, JSValueRef, unsafe.Pointer) unsafe.Pointer
	_JSValueIsArray func(JSContextRef, JSValueRef) bool
	_JSValueIsBoolean func(JSContextRef, JSValueRef) bool
	_JSValueIsDate func(JSContextRef, JSValueRef) bool
	_JSValueIsEqual func(JSContextRef, JSValueRef, JSValueRef, unsafe.Pointer) bool
	_JSValueIsInstanceOfConstructor func(JSContextRef, JSValueRef, JSObjectRef, unsafe.Pointer) bool
	_JSValueIsNull func(JSContextRef, JSValueRef) bool
	_JSValueIsNumber func(JSContextRef, JSValueRef) bool
	_JSValueIsObject func(JSContextRef, JSValueRef) bool
	_JSValueIsObjectOfClass func(JSContextRef, JSValueRef, JSClassRef) bool
	_JSValueIsStrictEqual func(JSContextRef, JSValueRef, JSValueRef) bool
	_JSValueIsString func(JSContextRef, JSValueRef) bool
	_JSValueIsSymbol func(JSContextRef, JSValueRef) bool
	_JSValueIsUndefined func(JSContextRef, JSValueRef) bool
	_JSValueMakeBoolean func(JSContextRef, bool) JSValueRef
	_JSValueMakeFromJSONString func(JSContextRef, JSStringRef) JSValueRef
	_JSValueMakeNull func(JSContextRef) JSValueRef
	_JSValueMakeNumber func(JSContextRef, float64) JSValueRef
	_JSValueMakeString func(JSContextRef, JSStringRef) JSValueRef
	_JSValueMakeSymbol func(JSContextRef, JSStringRef) JSValueRef
	_JSValueMakeUndefined func(JSContextRef) JSValueRef
	_JSValueProtect func(JSContextRef, JSValueRef)
	_JSValueToBoolean func(JSContextRef, JSValueRef) bool
	_JSValueToNumber func(JSContextRef, JSValueRef, unsafe.Pointer) float64
	_JSValueToObject func(JSContextRef, JSValueRef, unsafe.Pointer) JSObjectRef
	_JSValueToStringCopy func(JSContextRef, JSValueRef, unsafe.Pointer) JSStringRef
	_JSValueUnprotect func(JSContextRef, JSValueRef)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_JSBigIntCreateWithDouble, lib, "JSBigIntCreateWithDouble")
	tryRegister(&_JSBigIntCreateWithInt64, lib, "JSBigIntCreateWithInt64")
	tryRegister(&_JSBigIntCreateWithString, lib, "JSBigIntCreateWithString")
	tryRegister(&_JSBigIntCreateWithUInt64, lib, "JSBigIntCreateWithUInt64")
	tryRegister(&_JSValueCompare, lib, "JSValueCompare")
	tryRegister(&_JSValueCompareDouble, lib, "JSValueCompareDouble")
	tryRegister(&_JSValueCompareInt64, lib, "JSValueCompareInt64")
	tryRegister(&_JSValueCompareUInt64, lib, "JSValueCompareUInt64")
	tryRegister(&_JSValueIsBigInt, lib, "JSValueIsBigInt")
	tryRegister(&_JSValueToInt32, lib, "JSValueToInt32")
	tryRegister(&_JSValueToInt64, lib, "JSValueToInt64")
	tryRegister(&_JSValueToUInt32, lib, "JSValueToUInt32")
	tryRegister(&_JSValueToUInt64, lib, "JSValueToUInt64")
	tryRegister(&_JSCheckScriptSyntax, lib, "JSCheckScriptSyntax")
	tryRegister(&_JSClassCreate, lib, "JSClassCreate")
	tryRegister(&_JSClassRelease, lib, "JSClassRelease")
	tryRegister(&_JSClassRetain, lib, "JSClassRetain")
	tryRegister(&_JSContextGetGlobalContext, lib, "JSContextGetGlobalContext")
	tryRegister(&_JSContextGetGlobalObject, lib, "JSContextGetGlobalObject")
	tryRegister(&_JSContextGetGroup, lib, "JSContextGetGroup")
	tryRegister(&_JSContextGroupCreate, lib, "JSContextGroupCreate")
	tryRegister(&_JSContextGroupRelease, lib, "JSContextGroupRelease")
	tryRegister(&_JSContextGroupRetain, lib, "JSContextGroupRetain")
	tryRegister(&_JSEvaluateScript, lib, "JSEvaluateScript")
	tryRegister(&_JSGarbageCollect, lib, "JSGarbageCollect")
	tryRegister(&_JSGlobalContextCopyName, lib, "JSGlobalContextCopyName")
	tryRegister(&_JSGlobalContextCreate, lib, "JSGlobalContextCreate")
	tryRegister(&_JSGlobalContextCreateInGroup, lib, "JSGlobalContextCreateInGroup")
	tryRegister(&_JSGlobalContextIsInspectable, lib, "JSGlobalContextIsInspectable")
	tryRegister(&_JSGlobalContextRelease, lib, "JSGlobalContextRelease")
	tryRegister(&_JSGlobalContextRetain, lib, "JSGlobalContextRetain")
	tryRegister(&_JSGlobalContextSetInspectable, lib, "JSGlobalContextSetInspectable")
	tryRegister(&_JSGlobalContextSetName, lib, "JSGlobalContextSetName")
	tryRegister(&_JSObjectCallAsConstructor, lib, "JSObjectCallAsConstructor")
	tryRegister(&_JSObjectCallAsFunction, lib, "JSObjectCallAsFunction")
	tryRegister(&_JSObjectCopyPropertyNames, lib, "JSObjectCopyPropertyNames")
	tryRegister(&_JSObjectDeleteProperty, lib, "JSObjectDeleteProperty")
	tryRegister(&_JSObjectDeletePropertyForKey, lib, "JSObjectDeletePropertyForKey")
	tryRegister(&_JSObjectGetArrayBufferByteLength, lib, "JSObjectGetArrayBufferByteLength")
	tryRegister(&_JSObjectGetArrayBufferBytesPtr, lib, "JSObjectGetArrayBufferBytesPtr")
	tryRegister(&_JSObjectGetPrivate, lib, "JSObjectGetPrivate")
	tryRegister(&_JSObjectGetProperty, lib, "JSObjectGetProperty")
	tryRegister(&_JSObjectGetPropertyAtIndex, lib, "JSObjectGetPropertyAtIndex")
	tryRegister(&_JSObjectGetPropertyForKey, lib, "JSObjectGetPropertyForKey")
	tryRegister(&_JSObjectGetPrototype, lib, "JSObjectGetPrototype")
	tryRegister(&_JSObjectGetTypedArrayBuffer, lib, "JSObjectGetTypedArrayBuffer")
	tryRegister(&_JSObjectGetTypedArrayByteLength, lib, "JSObjectGetTypedArrayByteLength")
	tryRegister(&_JSObjectGetTypedArrayByteOffset, lib, "JSObjectGetTypedArrayByteOffset")
	tryRegister(&_JSObjectGetTypedArrayBytesPtr, lib, "JSObjectGetTypedArrayBytesPtr")
	tryRegister(&_JSObjectGetTypedArrayLength, lib, "JSObjectGetTypedArrayLength")
	tryRegister(&_JSObjectHasProperty, lib, "JSObjectHasProperty")
	tryRegister(&_JSObjectHasPropertyForKey, lib, "JSObjectHasPropertyForKey")
	tryRegister(&_JSObjectIsConstructor, lib, "JSObjectIsConstructor")
	tryRegister(&_JSObjectIsFunction, lib, "JSObjectIsFunction")
	tryRegister(&_JSObjectMake, lib, "JSObjectMake")
	tryRegister(&_JSObjectMakeArray, lib, "JSObjectMakeArray")
	tryRegister(&_JSObjectMakeArrayBufferWithBytesNoCopy, lib, "JSObjectMakeArrayBufferWithBytesNoCopy")
	tryRegister(&_JSObjectMakeConstructor, lib, "JSObjectMakeConstructor")
	tryRegister(&_JSObjectMakeDate, lib, "JSObjectMakeDate")
	tryRegister(&_JSObjectMakeDeferredPromise, lib, "JSObjectMakeDeferredPromise")
	tryRegister(&_JSObjectMakeError, lib, "JSObjectMakeError")
	tryRegister(&_JSObjectMakeFunction, lib, "JSObjectMakeFunction")
	tryRegister(&_JSObjectMakeFunctionWithCallback, lib, "JSObjectMakeFunctionWithCallback")
	tryRegister(&_JSObjectMakeRegExp, lib, "JSObjectMakeRegExp")
	tryRegister(&_JSObjectMakeTypedArray, lib, "JSObjectMakeTypedArray")
	tryRegister(&_JSObjectMakeTypedArrayWithArrayBuffer, lib, "JSObjectMakeTypedArrayWithArrayBuffer")
	tryRegister(&_JSObjectMakeTypedArrayWithArrayBufferAndOffset, lib, "JSObjectMakeTypedArrayWithArrayBufferAndOffset")
	tryRegister(&_JSObjectMakeTypedArrayWithBytesNoCopy, lib, "JSObjectMakeTypedArrayWithBytesNoCopy")
	tryRegister(&_JSObjectSetPrivate, lib, "JSObjectSetPrivate")
	tryRegister(&_JSObjectSetProperty, lib, "JSObjectSetProperty")
	tryRegister(&_JSObjectSetPropertyAtIndex, lib, "JSObjectSetPropertyAtIndex")
	tryRegister(&_JSObjectSetPropertyForKey, lib, "JSObjectSetPropertyForKey")
	tryRegister(&_JSObjectSetPrototype, lib, "JSObjectSetPrototype")
	tryRegister(&_JSPropertyNameAccumulatorAddName, lib, "JSPropertyNameAccumulatorAddName")
	tryRegister(&_JSPropertyNameArrayGetCount, lib, "JSPropertyNameArrayGetCount")
	tryRegister(&_JSPropertyNameArrayGetNameAtIndex, lib, "JSPropertyNameArrayGetNameAtIndex")
	tryRegister(&_JSPropertyNameArrayRelease, lib, "JSPropertyNameArrayRelease")
	tryRegister(&_JSPropertyNameArrayRetain, lib, "JSPropertyNameArrayRetain")
	tryRegister(&_JSStringCopyCFString, lib, "JSStringCopyCFString")
	tryRegister(&_JSStringCreateWithCFString, lib, "JSStringCreateWithCFString")
	tryRegister(&_JSStringCreateWithCharacters, lib, "JSStringCreateWithCharacters")
	tryRegister(&_JSStringCreateWithUTF8CString, lib, "JSStringCreateWithUTF8CString")
	tryRegister(&_JSStringGetCharactersPtr, lib, "JSStringGetCharactersPtr")
	tryRegister(&_JSStringGetLength, lib, "JSStringGetLength")
	tryRegister(&_JSStringGetMaximumUTF8CStringSize, lib, "JSStringGetMaximumUTF8CStringSize")
	tryRegister(&_JSStringGetUTF8CString, lib, "JSStringGetUTF8CString")
	tryRegister(&_JSStringIsEqual, lib, "JSStringIsEqual")
	tryRegister(&_JSStringIsEqualToUTF8CString, lib, "JSStringIsEqualToUTF8CString")
	tryRegister(&_JSStringRelease, lib, "JSStringRelease")
	tryRegister(&_JSStringRetain, lib, "JSStringRetain")
	tryRegister(&_JSValueCreateJSONString, lib, "JSValueCreateJSONString")
	tryRegister(&_JSValueGetType, lib, "JSValueGetType")
	tryRegister(&_JSValueGetTypedArrayType, lib, "JSValueGetTypedArrayType")
	tryRegister(&_JSValueIsArray, lib, "JSValueIsArray")
	tryRegister(&_JSValueIsBoolean, lib, "JSValueIsBoolean")
	tryRegister(&_JSValueIsDate, lib, "JSValueIsDate")
	tryRegister(&_JSValueIsEqual, lib, "JSValueIsEqual")
	tryRegister(&_JSValueIsInstanceOfConstructor, lib, "JSValueIsInstanceOfConstructor")
	tryRegister(&_JSValueIsNull, lib, "JSValueIsNull")
	tryRegister(&_JSValueIsNumber, lib, "JSValueIsNumber")
	tryRegister(&_JSValueIsObject, lib, "JSValueIsObject")
	tryRegister(&_JSValueIsObjectOfClass, lib, "JSValueIsObjectOfClass")
	tryRegister(&_JSValueIsStrictEqual, lib, "JSValueIsStrictEqual")
	tryRegister(&_JSValueIsString, lib, "JSValueIsString")
	tryRegister(&_JSValueIsSymbol, lib, "JSValueIsSymbol")
	tryRegister(&_JSValueIsUndefined, lib, "JSValueIsUndefined")
	tryRegister(&_JSValueMakeBoolean, lib, "JSValueMakeBoolean")
	tryRegister(&_JSValueMakeFromJSONString, lib, "JSValueMakeFromJSONString")
	tryRegister(&_JSValueMakeNull, lib, "JSValueMakeNull")
	tryRegister(&_JSValueMakeNumber, lib, "JSValueMakeNumber")
	tryRegister(&_JSValueMakeString, lib, "JSValueMakeString")
	tryRegister(&_JSValueMakeSymbol, lib, "JSValueMakeSymbol")
	tryRegister(&_JSValueMakeUndefined, lib, "JSValueMakeUndefined")
	tryRegister(&_JSValueProtect, lib, "JSValueProtect")
	tryRegister(&_JSValueToBoolean, lib, "JSValueToBoolean")
	tryRegister(&_JSValueToNumber, lib, "JSValueToNumber")
	tryRegister(&_JSValueToObject, lib, "JSValueToObject")
	tryRegister(&_JSValueToStringCopy, lib, "JSValueToStringCopy")
	tryRegister(&_JSValueUnprotect, lib, "JSValueUnprotect")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// JSBigIntCreateWithDouble is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithDouble(_:_:_:)
func JSBigIntCreateWithDouble(ctx JSContextRef, value float64, exception unsafe.Pointer) JSValueRef {
	return _JSBigIntCreateWithDouble(ctx, value, exception)
}/* debug [functions.gen.go/function]: JSBigIntCreateWithDouble */

// JSBigIntCreateWithInt64 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithInt64(_:_:_:)
func JSBigIntCreateWithInt64(ctx JSContextRef, integer int64, exception unsafe.Pointer) JSValueRef {
	return _JSBigIntCreateWithInt64(ctx, integer, exception)
}/* debug [functions.gen.go/function]: JSBigIntCreateWithInt64 */

// JSBigIntCreateWithString is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithString(_:_:_:)
func JSBigIntCreateWithString(ctx JSContextRef, string_ JSStringRef, exception unsafe.Pointer) JSValueRef {
	return _JSBigIntCreateWithString(ctx, string_, exception)
}/* debug [functions.gen.go/function]: JSBigIntCreateWithString */

// JSBigIntCreateWithUInt64 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithUInt64(_:_:_:)
func JSBigIntCreateWithUInt64(ctx JSContextRef, integer uint64, exception unsafe.Pointer) JSValueRef {
	return _JSBigIntCreateWithUInt64(ctx, integer, exception)
}/* debug [functions.gen.go/function]: JSBigIntCreateWithUInt64 */

// JSValueCompare is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompare(_:_:_:_:)
func JSValueCompare(ctx JSContextRef, left JSValueRef, right JSValueRef, exception unsafe.Pointer) JSRelationCondition {
	return _JSValueCompare(ctx, left, right, exception)
}/* debug [functions.gen.go/function]: JSValueCompare */

// JSValueCompareDouble is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompareDouble(_:_:_:_:)
func JSValueCompareDouble(ctx JSContextRef, left JSValueRef, right float64, exception unsafe.Pointer) JSRelationCondition {
	return _JSValueCompareDouble(ctx, left, right, exception)
}/* debug [functions.gen.go/function]: JSValueCompareDouble */

// JSValueCompareInt64 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompareInt64(_:_:_:_:)
func JSValueCompareInt64(ctx JSContextRef, left JSValueRef, right int64, exception unsafe.Pointer) JSRelationCondition {
	return _JSValueCompareInt64(ctx, left, right, exception)
}/* debug [functions.gen.go/function]: JSValueCompareInt64 */

// JSValueCompareUInt64 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompareUInt64(_:_:_:_:)
func JSValueCompareUInt64(ctx JSContextRef, left JSValueRef, right uint64, exception unsafe.Pointer) JSRelationCondition {
	return _JSValueCompareUInt64(ctx, left, right, exception)
}/* debug [functions.gen.go/function]: JSValueCompareUInt64 */

// JSValueIsBigInt is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsBigInt(_:_:)
func JSValueIsBigInt(ctx JSContextRef, value JSValueRef) bool {
	return _JSValueIsBigInt(ctx, value)
}/* debug [functions.gen.go/function]: JSValueIsBigInt */

// JSValueToInt32 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToInt32(_:_:_:)
func JSValueToInt32(ctx JSContextRef, value JSValueRef, exception unsafe.Pointer) int32 {
	return _JSValueToInt32(ctx, value, exception)
}/* debug [functions.gen.go/function]: JSValueToInt32 */

// JSValueToInt64 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToInt64(_:_:_:)
func JSValueToInt64(ctx JSContextRef, value JSValueRef, exception unsafe.Pointer) int64 {
	return _JSValueToInt64(ctx, value, exception)
}/* debug [functions.gen.go/function]: JSValueToInt64 */

// JSValueToUInt32 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToUInt32(_:_:_:)
func JSValueToUInt32(ctx JSContextRef, value JSValueRef, exception unsafe.Pointer) uint32 {
	return _JSValueToUInt32(ctx, value, exception)
}/* debug [functions.gen.go/function]: JSValueToUInt32 */

// JSValueToUInt64 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToUInt64(_:_:_:)
func JSValueToUInt64(ctx JSContextRef, value JSValueRef, exception unsafe.Pointer) uint64 {
	return _JSValueToUInt64(ctx, value, exception)
}/* debug [functions.gen.go/function]: JSValueToUInt64 */

// Checks for syntax errors in a string of JavaScript.
//
// Added in macOS 10.5.
// Checks for syntax errors in a string of JavaScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSCheckScriptSyntax(_:_:_:_:_:)
func JSCheckScriptSyntax(ctx JSContextRef, script JSStringRef, sourceURL JSStringRef, startingLineNumber int, exception unsafe.Pointer) bool {
	return _JSCheckScriptSyntax(ctx, script, sourceURL, startingLineNumber, exception)
}/* debug [functions.gen.go/function]: JSCheckScriptSyntax */

// Creates a JavaScript class.
//
// Added in macOS 10.5.
// Creates a JavaScript class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSClassCreate(_:)
func JSClassCreate(definition unsafe.Pointer) JSClassRef {
	return _JSClassCreate(definition)
}/* debug [functions.gen.go/function]: JSClassCreate */

// Releases a JavaScript class.
//
// Added in macOS 10.5.
// Releases a JavaScript class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSClassRelease(_:)
func JSClassRelease(jsClass JSClassRef) {
	_JSClassRelease(jsClass)
}/* debug [functions.gen.go/function]: JSClassRelease */

// Retains a JavaScript class.
//
// Added in macOS 10.5.
// Retains a JavaScript class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSClassRetain(_:)
func JSClassRetain(jsClass JSClassRef) JSClassRef {
	return _JSClassRetain(jsClass)
}/* debug [functions.gen.go/function]: JSClassRetain */

// Gets the global context of a JavaScript execution context.
//
// Added in macOS 10.7.
// Gets the global context of a JavaScript execution context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContextGetGlobalContext(_:)
func JSContextGetGlobalContext(ctx JSContextRef) JSGlobalContextRef {
	return _JSContextGetGlobalContext(ctx)
}/* debug [functions.gen.go/function]: JSContextGetGlobalContext */

// Gets the global object of a JavaScript execution context.
//
// Added in macOS 10.5.
// Gets the global object of a JavaScript execution context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContextGetGlobalObject(_:)
func JSContextGetGlobalObject(ctx JSContextRef) JSObjectRef {
	return _JSContextGetGlobalObject(ctx)
}/* debug [functions.gen.go/function]: JSContextGetGlobalObject */

// Gets the context group that a JavaScript execution context belongs to.
//
// Added in macOS 10.6.
// Gets the context group that a JavaScript execution context belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContextGetGroup(_:)
func JSContextGetGroup(ctx JSContextRef) JSContextGroupRef {
	return _JSContextGetGroup(ctx)
}/* debug [functions.gen.go/function]: JSContextGetGroup */

// Creates a JavaScript context group.
//
// Added in macOS 10.6.
// Creates a JavaScript context group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContextGroupCreate()
func JSContextGroupCreate() JSContextGroupRef {
	return _JSContextGroupCreate()
}/* debug [functions.gen.go/function]: JSContextGroupCreate */

// Releases a JavaScript context group.
//
// Added in macOS 10.6.
// Releases a JavaScript context group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContextGroupRelease(_:)
func JSContextGroupRelease(group JSContextGroupRef) {
	_JSContextGroupRelease(group)
}/* debug [functions.gen.go/function]: JSContextGroupRelease */

// Retains a JavaScript context group.
//
// Added in macOS 10.6.
// Retains a JavaScript context group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContextGroupRetain(_:)
func JSContextGroupRetain(group JSContextGroupRef) JSContextGroupRef {
	return _JSContextGroupRetain(group)
}/* debug [functions.gen.go/function]: JSContextGroupRetain */

// Evaluates a string of JavaScript.
//
// Added in macOS 10.5.
// Evaluates a string of JavaScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSEvaluateScript(_:_:_:_:_:_:)
func JSEvaluateScript(ctx JSContextRef, script JSStringRef, thisObject JSObjectRef, sourceURL JSStringRef, startingLineNumber int, exception unsafe.Pointer) JSValueRef {
	return _JSEvaluateScript(ctx, script, thisObject, sourceURL, startingLineNumber, exception)
}/* debug [functions.gen.go/function]: JSEvaluateScript */

// Performs a JavaScript garbage collection.
//
// Added in macOS 10.5.
// Performs a JavaScript garbage collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGarbageCollect(_:)
func JSGarbageCollect(ctx JSContextRef) {
	_JSGarbageCollect(ctx)
}/* debug [functions.gen.go/function]: JSGarbageCollect */

// Gets a copy of the name of a context.
//
// Added in macOS 10.10.
// Gets a copy of the name of a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextCopyName(_:)
func JSGlobalContextCopyName(ctx JSGlobalContextRef) JSStringRef {
	return _JSGlobalContextCopyName(ctx)
}/* debug [functions.gen.go/function]: JSGlobalContextCopyName */

// Creates a global JavaScript execution context.
//
// Added in macOS 10.5.
// Creates a global JavaScript execution context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextCreate(_:)
func JSGlobalContextCreate(globalObjectClass JSClassRef) JSGlobalContextRef {
	return _JSGlobalContextCreate(globalObjectClass)
}/* debug [functions.gen.go/function]: JSGlobalContextCreate */

// Creates a global JavaScript execution context in the provided context group.
//
// Added in macOS 10.6.
// Creates a global JavaScript execution context in the provided context group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextCreateInGroup(_:_:)
func JSGlobalContextCreateInGroup(group JSContextGroupRef, globalObjectClass JSClassRef) JSGlobalContextRef {
	return _JSGlobalContextCreateInGroup(group, globalObjectClass)
}/* debug [functions.gen.go/function]: JSGlobalContextCreateInGroup */

// Returns a Boolean value that indicates whether the JavaScript context is inspectable.
//
// Added in macOS 13.3.
// Returns a Boolean value that indicates whether the JavaScript context is inspectable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextIsInspectable(_:)
func JSGlobalContextIsInspectable(ctx JSGlobalContextRef) bool {
	return _JSGlobalContextIsInspectable(ctx)
}/* debug [functions.gen.go/function]: JSGlobalContextIsInspectable */

// Releases a global JavaScript execution context.
//
// Added in macOS 10.5.
// Releases a global JavaScript execution context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextRelease(_:)
func JSGlobalContextRelease(ctx JSGlobalContextRef) {
	_JSGlobalContextRelease(ctx)
}/* debug [functions.gen.go/function]: JSGlobalContextRelease */

// Retains a global JavaScript execution context.
//
// Added in macOS 10.5.
// Retains a global JavaScript execution context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextRetain(_:)
func JSGlobalContextRetain(ctx JSGlobalContextRef) JSGlobalContextRef {
	return _JSGlobalContextRetain(ctx)
}/* debug [functions.gen.go/function]: JSGlobalContextRetain */

// Sets a JavaScript context to be either inspectable or not inspectable.
//
// Added in macOS 13.3.
// Sets a JavaScript context to be either inspectable or not inspectable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextSetInspectable(_:_:)
func JSGlobalContextSetInspectable(ctx JSGlobalContextRef, inspectable bool) {
	_JSGlobalContextSetInspectable(ctx, inspectable)
}/* debug [functions.gen.go/function]: JSGlobalContextSetInspectable */

// Sets the remote debugging name for a context.
//
// Added in macOS 10.10.
// Sets the remote debugging name for a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextSetName(_:_:)
func JSGlobalContextSetName(ctx JSGlobalContextRef, name JSStringRef) {
	_JSGlobalContextSetName(ctx, name)
}/* debug [functions.gen.go/function]: JSGlobalContextSetName */

// Calls an object as a constructor.
//
// Added in macOS 10.5.
// Calls an object as a constructor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectCallAsConstructor(_:_:_:_:_:)
func JSObjectCallAsConstructor(ctx JSContextRef, object JSObjectRef, argumentCount uintptr, arguments unsafe.Pointer, exception unsafe.Pointer) JSObjectRef {
	return _JSObjectCallAsConstructor(ctx, object, argumentCount, arguments, exception)
}/* debug [functions.gen.go/function]: JSObjectCallAsConstructor */

// Calls an object as a function.
//
// Added in macOS 10.5.
// Calls an object as a function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectCallAsFunction(_:_:_:_:_:_:)
func JSObjectCallAsFunction(ctx JSContextRef, object JSObjectRef, thisObject JSObjectRef, argumentCount uintptr, arguments unsafe.Pointer, exception unsafe.Pointer) JSValueRef {
	return _JSObjectCallAsFunction(ctx, object, thisObject, argumentCount, arguments, exception)
}/* debug [functions.gen.go/function]: JSObjectCallAsFunction */

// Gets the names of an object’s enumerable properties.
//
// Added in macOS 10.5.
// Gets the names of an object’s enumerable properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectCopyPropertyNames(_:_:)
func JSObjectCopyPropertyNames(ctx JSContextRef, object JSObjectRef) JSPropertyNameArrayRef {
	return _JSObjectCopyPropertyNames(ctx, object)
}/* debug [functions.gen.go/function]: JSObjectCopyPropertyNames */

// Deletes a property from an object.
//
// Added in macOS 10.5.
// Deletes a property from an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectDeleteProperty(_:_:_:_:)
func JSObjectDeleteProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, exception unsafe.Pointer) bool {
	return _JSObjectDeleteProperty(ctx, object, propertyName, exception)
}/* debug [functions.gen.go/function]: JSObjectDeleteProperty */

// Deletes a property from an object using a JavaScript value as the property key.
//
// Added in macOS 10.15.
// Deletes a property from an object using a JavaScript value as the property key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectDeletePropertyForKey(_:_:_:_:)
func JSObjectDeletePropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception unsafe.Pointer) bool {
	return _JSObjectDeletePropertyForKey(ctx, object, propertyKey, exception)
}/* debug [functions.gen.go/function]: JSObjectDeletePropertyForKey */

// Returns the number of bytes in a JavaScript data object.
//
// Added in macOS 10.12.
// Returns the number of bytes in a JavaScript data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetArrayBufferByteLength(_:_:_:)
func JSObjectGetArrayBufferByteLength(ctx JSContextRef, object JSObjectRef, exception unsafe.Pointer) uintptr {
	return _JSObjectGetArrayBufferByteLength(ctx, object, exception)
}/* debug [functions.gen.go/function]: JSObjectGetArrayBufferByteLength */

// Returns a pointer to the data buffer that serves as the backing store for a JavaScript typed array object.
//
// Added in macOS 10.12.
// Returns a pointer to the data buffer that serves as the backing store for a JavaScript typed array object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetArrayBufferBytesPtr(_:_:_:)
func JSObjectGetArrayBufferBytesPtr(ctx JSContextRef, object JSObjectRef, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectGetArrayBufferBytesPtr(ctx, object, exception)
}/* debug [functions.gen.go/function]: JSObjectGetArrayBufferBytesPtr */

// Gets an object’s private data.
//
// Added in macOS 10.5.
// Gets an object’s private data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPrivate(_:)
func JSObjectGetPrivate(object JSObjectRef) unsafe.Pointer {
	return _JSObjectGetPrivate(object)
}/* debug [functions.gen.go/function]: JSObjectGetPrivate */

// Gets a property from an object.
//
// Added in macOS 10.5.
// Gets a property from an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetProperty(_:_:_:_:)
func JSObjectGetProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, exception unsafe.Pointer) JSValueRef {
	return _JSObjectGetProperty(ctx, object, propertyName, exception)
}/* debug [functions.gen.go/function]: JSObjectGetProperty */

// Gets a property from an object by numeric index.
//
// Added in macOS 10.5.
// Gets a property from an object by numeric index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPropertyAtIndex(_:_:_:_:)
func JSObjectGetPropertyAtIndex(ctx JSContextRef, object JSObjectRef, propertyIndex unsafe.Pointer, exception unsafe.Pointer) JSValueRef {
	return _JSObjectGetPropertyAtIndex(ctx, object, propertyIndex, exception)
}/* debug [functions.gen.go/function]: JSObjectGetPropertyAtIndex */

// Gets a property from an object using a JavaScript value as the property key.
//
// Added in macOS 10.15.
// Gets a property from an object using a JavaScript value as the property key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPropertyForKey(_:_:_:_:)
func JSObjectGetPropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception unsafe.Pointer) JSValueRef {
	return _JSObjectGetPropertyForKey(ctx, object, propertyKey, exception)
}/* debug [functions.gen.go/function]: JSObjectGetPropertyForKey */

// Gets an object’s prototype.
//
// Added in macOS 10.5.
// Gets an object’s prototype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPrototype(_:_:)
func JSObjectGetPrototype(ctx JSContextRef, object JSObjectRef) JSValueRef {
	return _JSObjectGetPrototype(ctx, object)
}/* debug [functions.gen.go/function]: JSObjectGetPrototype */

// Returns the JavaScript array buffer object to use as the backing of a JavaScript typed array object.
//
// Added in macOS 10.12.
// Returns the JavaScript array buffer object to use as the backing of a JavaScript typed array object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayBuffer(_:_:_:)
func JSObjectGetTypedArrayBuffer(ctx JSContextRef, object JSObjectRef, exception unsafe.Pointer) JSObjectRef {
	return _JSObjectGetTypedArrayBuffer(ctx, object, exception)
}/* debug [functions.gen.go/function]: JSObjectGetTypedArrayBuffer */

// Returns the byte length of a JavaScript typed array object.
//
// Added in macOS 10.12.
// Returns the byte length of a JavaScript typed array object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayByteLength(_:_:_:)
func JSObjectGetTypedArrayByteLength(ctx JSContextRef, object JSObjectRef, exception unsafe.Pointer) uintptr {
	return _JSObjectGetTypedArrayByteLength(ctx, object, exception)
}/* debug [functions.gen.go/function]: JSObjectGetTypedArrayByteLength */

// Returns the byte offset of a JavaScript typed array object.
//
// Added in macOS 10.12.
// Returns the byte offset of a JavaScript typed array object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayByteOffset(_:_:_:)
func JSObjectGetTypedArrayByteOffset(ctx JSContextRef, object JSObjectRef, exception unsafe.Pointer) uintptr {
	return _JSObjectGetTypedArrayByteOffset(ctx, object, exception)
}/* debug [functions.gen.go/function]: JSObjectGetTypedArrayByteOffset */

// Returns a temporary pointer to the backing store of a JavaScript typed array object.
//
// Added in macOS 10.12.
// Returns a temporary pointer to the backing store of a JavaScript typed array object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayBytesPtr(_:_:_:)
func JSObjectGetTypedArrayBytesPtr(ctx JSContextRef, object JSObjectRef, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectGetTypedArrayBytesPtr(ctx, object, exception)
}/* debug [functions.gen.go/function]: JSObjectGetTypedArrayBytesPtr */

// Returns the length of a JavaScript typed array object.
//
// Added in macOS 10.12.
// Returns the length of a JavaScript typed array object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayLength(_:_:_:)
func JSObjectGetTypedArrayLength(ctx JSContextRef, object JSObjectRef, exception unsafe.Pointer) uintptr {
	return _JSObjectGetTypedArrayLength(ctx, object, exception)
}/* debug [functions.gen.go/function]: JSObjectGetTypedArrayLength */

// Tests whether an object has a specified property.
//
// Added in macOS 10.5.
// Tests whether an object has a specified property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectHasProperty(_:_:_:)
func JSObjectHasProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef) bool {
	return _JSObjectHasProperty(ctx, object, propertyName)
}/* debug [functions.gen.go/function]: JSObjectHasProperty */

// Tests whether an object has the specified property using a JavaScript value as the property key.
//
// Added in macOS 10.15.
// Tests whether an object has the specified property using a JavaScript value as the property key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectHasPropertyForKey(_:_:_:_:)
func JSObjectHasPropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception unsafe.Pointer) bool {
	return _JSObjectHasPropertyForKey(ctx, object, propertyKey, exception)
}/* debug [functions.gen.go/function]: JSObjectHasPropertyForKey */

// Tests whether you can call an object as a constructor.
//
// Added in macOS 10.5.
// Tests whether you can call an object as a constructor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectIsConstructor(_:_:)
func JSObjectIsConstructor(ctx JSContextRef, object JSObjectRef) bool {
	return _JSObjectIsConstructor(ctx, object)
}/* debug [functions.gen.go/function]: JSObjectIsConstructor */

// Tests whether you can call an object as a function.
//
// Added in macOS 10.5.
// Tests whether you can call an object as a function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectIsFunction(_:_:)
func JSObjectIsFunction(ctx JSContextRef, object JSObjectRef) bool {
	return _JSObjectIsFunction(ctx, object)
}/* debug [functions.gen.go/function]: JSObjectIsFunction */

// Creates a JavaScript object.
//
// Added in macOS 10.5.
// Creates a JavaScript object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMake(_:_:_:)
func JSObjectMake(ctx JSContextRef, jsClass JSClassRef, data unsafe.Pointer) JSObjectRef {
	return _JSObjectMake(ctx, jsClass, data)
}/* debug [functions.gen.go/function]: JSObjectMake */

// Creates a JavaScript array object.
//
// Added in macOS 10.6.
// Creates a JavaScript array object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeArray(_:_:_:_:)
func JSObjectMakeArray(ctx JSContextRef, argumentCount uintptr, arguments unsafe.Pointer, exception unsafe.Pointer) JSObjectRef {
	return _JSObjectMakeArray(ctx, argumentCount, arguments, exception)
}/* debug [functions.gen.go/function]: JSObjectMakeArray */

// Creates a JavaScript array buffer object from an existing pointer.
//
// Added in macOS 10.12.
// Creates a JavaScript array buffer object from an existing pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeArrayBufferWithBytesNoCopy(_:_:_:_:_:_:)
func JSObjectMakeArrayBufferWithBytesNoCopy(ctx JSContextRef, bytes unsafe.Pointer, byteLength uintptr, bytesDeallocator JSTypedArrayBytesDeallocator, deallocatorContext unsafe.Pointer, exception unsafe.Pointer) JSObjectRef {
	return _JSObjectMakeArrayBufferWithBytesNoCopy(ctx, bytes, byteLength, bytesDeallocator, deallocatorContext, exception)
}/* debug [functions.gen.go/function]: JSObjectMakeArrayBufferWithBytesNoCopy */

// Creates a JavaScript constructor.
//
// Added in macOS 10.5.
// Creates a JavaScript constructor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeConstructor(_:_:_:)
func JSObjectMakeConstructor(ctx JSContextRef, jsClass JSClassRef, callAsConstructor JSObjectCallAsConstructorCallback) JSObjectRef {
	return _JSObjectMakeConstructor(ctx, jsClass, callAsConstructor)
}/* debug [functions.gen.go/function]: JSObjectMakeConstructor */

// Creates a JavaScript date object as though invoking the built-in date constructor.
//
// Added in macOS 10.6.
// Creates a JavaScript date object as though invoking the built-in date constructor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeDate(_:_:_:_:)
func JSObjectMakeDate(ctx JSContextRef, argumentCount uintptr, arguments unsafe.Pointer, exception unsafe.Pointer) JSObjectRef {
	return _JSObjectMakeDate(ctx, argumentCount, arguments, exception)
}/* debug [functions.gen.go/function]: JSObjectMakeDate */

// Creates a JavaScript promise object by invoking the provided executor.
//
// Added in macOS 10.15.
// Creates a JavaScript promise object by invoking the provided executor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeDeferredPromise(_:_:_:_:)
func JSObjectMakeDeferredPromise(ctx JSContextRef, resolve unsafe.Pointer, reject unsafe.Pointer, exception unsafe.Pointer) JSObjectRef {
	return _JSObjectMakeDeferredPromise(ctx, resolve, reject, exception)
}/* debug [functions.gen.go/function]: JSObjectMakeDeferredPromise */

// Creates a JavaScript error object as though invoking the built-in error constructor.
//
// Added in macOS 10.6.
// Creates a JavaScript error object as though invoking the built-in error constructor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeError(_:_:_:_:)
func JSObjectMakeError(ctx JSContextRef, argumentCount uintptr, arguments unsafe.Pointer, exception unsafe.Pointer) JSObjectRef {
	return _JSObjectMakeError(ctx, argumentCount, arguments, exception)
}/* debug [functions.gen.go/function]: JSObjectMakeError */

// Creates a function with a specified script as its body.
//
// Added in macOS 10.5.
// Creates a function with a specified script as its body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeFunction(_:_:_:_:_:_:_:_:)
func JSObjectMakeFunction(ctx JSContextRef, name JSStringRef, parameterCount unsafe.Pointer, parameterNames unsafe.Pointer, body JSStringRef, sourceURL JSStringRef, startingLineNumber int, exception unsafe.Pointer) JSObjectRef {
	return _JSObjectMakeFunction(ctx, name, parameterCount, parameterNames, body, sourceURL, startingLineNumber, exception)
}/* debug [functions.gen.go/function]: JSObjectMakeFunction */

// Creates a JavaScript function with a specified callback as its implementation.
//
// Added in macOS 10.5.
// Creates a JavaScript function with a specified callback as its implementation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeFunctionWithCallback(_:_:_:)
func JSObjectMakeFunctionWithCallback(ctx JSContextRef, name JSStringRef, callAsFunction JSObjectCallAsFunctionCallback) JSObjectRef {
	return _JSObjectMakeFunctionWithCallback(ctx, name, callAsFunction)
}/* debug [functions.gen.go/function]: JSObjectMakeFunctionWithCallback */

// Creates a JavaScript regular expression object as though invoking the built-in regular expression constructor.
//
// Added in macOS 10.6.
// Creates a JavaScript regular expression object as though invoking the built-in regular expression constructor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeRegExp(_:_:_:_:)
func JSObjectMakeRegExp(ctx JSContextRef, argumentCount uintptr, arguments unsafe.Pointer, exception unsafe.Pointer) JSObjectRef {
	return _JSObjectMakeRegExp(ctx, argumentCount, arguments, exception)
}/* debug [functions.gen.go/function]: JSObjectMakeRegExp */

// Creates a JavaScript typed array object with the specified number of elements.
//
// Added in macOS 10.12.
// Creates a JavaScript typed array object with the specified number of elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeTypedArray(_:_:_:_:)
func JSObjectMakeTypedArray(ctx JSContextRef, arrayType unsafe.Pointer, length uintptr, exception unsafe.Pointer) JSObjectRef {
	return _JSObjectMakeTypedArray(ctx, arrayType, length, exception)
}/* debug [functions.gen.go/function]: JSObjectMakeTypedArray */

// Creates a JavaScript typed array object from an existing JavaScript array buffer object.
//
// Added in macOS 10.12.
// Creates a JavaScript typed array object from an existing JavaScript array buffer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeTypedArrayWithArrayBuffer(_:_:_:_:)
func JSObjectMakeTypedArrayWithArrayBuffer(ctx JSContextRef, arrayType unsafe.Pointer, buffer JSObjectRef, exception unsafe.Pointer) JSObjectRef {
	return _JSObjectMakeTypedArrayWithArrayBuffer(ctx, arrayType, buffer, exception)
}/* debug [functions.gen.go/function]: JSObjectMakeTypedArrayWithArrayBuffer */

// Creates a JavaScript typed array object from an existing JavaScript array buffer object with the specified offset and length.
//
// Added in macOS 10.12.
// Creates a JavaScript typed array object from an existing JavaScript array buffer object with the specified offset and length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeTypedArrayWithArrayBufferAndOffset(_:_:_:_:_:_:)
func JSObjectMakeTypedArrayWithArrayBufferAndOffset(ctx JSContextRef, arrayType unsafe.Pointer, buffer JSObjectRef, byteOffset uintptr, length uintptr, exception unsafe.Pointer) JSObjectRef {
	return _JSObjectMakeTypedArrayWithArrayBufferAndOffset(ctx, arrayType, buffer, byteOffset, length, exception)
}/* debug [functions.gen.go/function]: JSObjectMakeTypedArrayWithArrayBufferAndOffset */

// Creates a JavaScript typed array object from an existing pointer.
//
// Added in macOS 10.12.
// Creates a JavaScript typed array object from an existing pointer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeTypedArrayWithBytesNoCopy(_:_:_:_:_:_:_:)
func JSObjectMakeTypedArrayWithBytesNoCopy(ctx JSContextRef, arrayType unsafe.Pointer, bytes unsafe.Pointer, byteLength uintptr, bytesDeallocator JSTypedArrayBytesDeallocator, deallocatorContext unsafe.Pointer, exception unsafe.Pointer) JSObjectRef {
	return _JSObjectMakeTypedArrayWithBytesNoCopy(ctx, arrayType, bytes, byteLength, bytesDeallocator, deallocatorContext, exception)
}/* debug [functions.gen.go/function]: JSObjectMakeTypedArrayWithBytesNoCopy */

// Sets a pointer to private data on an object.
//
// Added in macOS 10.5.
// Sets a pointer to private data on an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetPrivate(_:_:)
func JSObjectSetPrivate(object JSObjectRef, data unsafe.Pointer) bool {
	return _JSObjectSetPrivate(object, data)
}/* debug [functions.gen.go/function]: JSObjectSetPrivate */

// Sets a property on an object.
//
// Added in macOS 10.5.
// Sets a property on an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetProperty(_:_:_:_:_:_:)
func JSObjectSetProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, value JSValueRef, attributes JSPropertyAttributes, exception unsafe.Pointer) {
	_JSObjectSetProperty(ctx, object, propertyName, value, attributes, exception)
}/* debug [functions.gen.go/function]: JSObjectSetProperty */

// Sets a property on an object by numeric index.
//
// Added in macOS 10.5.
// Sets a property on an object by numeric index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetPropertyAtIndex(_:_:_:_:_:)
func JSObjectSetPropertyAtIndex(ctx JSContextRef, object JSObjectRef, propertyIndex unsafe.Pointer, value JSValueRef, exception unsafe.Pointer) {
	_JSObjectSetPropertyAtIndex(ctx, object, propertyIndex, value, exception)
}/* debug [functions.gen.go/function]: JSObjectSetPropertyAtIndex */

// Sets a property on an object using a JavaScript value as the property key.
//
// Added in macOS 10.15.
// Sets a property on an object using a JavaScript value as the property key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetPropertyForKey(_:_:_:_:_:_:)
func JSObjectSetPropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, value JSValueRef, attributes JSPropertyAttributes, exception unsafe.Pointer) {
	_JSObjectSetPropertyForKey(ctx, object, propertyKey, value, attributes, exception)
}/* debug [functions.gen.go/function]: JSObjectSetPropertyForKey */

// Sets an object’s prototype.
//
// Added in macOS 10.5.
// Sets an object’s prototype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetPrototype(_:_:_:)
func JSObjectSetPrototype(ctx JSContextRef, object JSObjectRef, value JSValueRef) {
	_JSObjectSetPrototype(ctx, object, value)
}/* debug [functions.gen.go/function]: JSObjectSetPrototype */

// Adds a property name to a JavaScript property name accumulator.
//
// Added in macOS 10.5.
// Adds a property name to a JavaScript property name accumulator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameAccumulatorAddName(_:_:)
func JSPropertyNameAccumulatorAddName(accumulator JSPropertyNameAccumulatorRef, propertyName JSStringRef) {
	_JSPropertyNameAccumulatorAddName(accumulator, propertyName)
}/* debug [functions.gen.go/function]: JSPropertyNameAccumulatorAddName */

// Gets a count of the number of items in a JavaScript property name array.
//
// Added in macOS 10.5.
// Gets a count of the number of items in a JavaScript property name array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameArrayGetCount(_:)
func JSPropertyNameArrayGetCount(array JSPropertyNameArrayRef) uintptr {
	return _JSPropertyNameArrayGetCount(array)
}/* debug [functions.gen.go/function]: JSPropertyNameArrayGetCount */

// Gets a property name at a specified index in a JavaScript property name array.
//
// Added in macOS 10.5.
// Gets a property name at a specified index in a JavaScript property name array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameArrayGetNameAtIndex(_:_:)
func JSPropertyNameArrayGetNameAtIndex(array JSPropertyNameArrayRef, index uintptr) JSStringRef {
	return _JSPropertyNameArrayGetNameAtIndex(array, index)
}/* debug [functions.gen.go/function]: JSPropertyNameArrayGetNameAtIndex */

// Releases a JavaScript property name array.
//
// Added in macOS 10.5.
// Releases a JavaScript property name array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameArrayRelease(_:)
func JSPropertyNameArrayRelease(array JSPropertyNameArrayRef) {
	_JSPropertyNameArrayRelease(array)
}/* debug [functions.gen.go/function]: JSPropertyNameArrayRelease */

// Retains a JavaScript property name array.
//
// Added in macOS 10.5.
// Retains a JavaScript property name array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameArrayRetain(_:)
func JSPropertyNameArrayRetain(array JSPropertyNameArrayRef) JSPropertyNameArrayRef {
	return _JSPropertyNameArrayRetain(array)
}/* debug [functions.gen.go/function]: JSPropertyNameArrayRetain */

// Creates a Core Foundation string from a JavaScript string.
//
// Added in macOS 10.5.
// Creates a Core Foundation string from a JavaScript string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringCopyCFString(_:_:)
func JSStringCopyCFString(alloc AllocatorRef, string_ JSStringRef) StringRef {
	return _JSStringCopyCFString(alloc, string_)
}/* debug [functions.gen.go/function]: JSStringCopyCFString */

// Creates a JavaScript string from a Core Foundation string.
//
// Added in macOS 10.5.
// Creates a JavaScript string from a Core Foundation string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringCreateWithCFString(_:)
func JSStringCreateWithCFString(string_ StringRef) JSStringRef {
	return _JSStringCreateWithCFString(string_)
}/* debug [functions.gen.go/function]: JSStringCreateWithCFString */

// Creates a JavaScript string from a buffer of Unicode characters.
//
// Added in macOS 10.5.
// Creates a JavaScript string from a buffer of Unicode characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringCreateWithCharacters(_:_:)
func JSStringCreateWithCharacters(chars unsafe.Pointer, numChars uintptr) JSStringRef {
	return _JSStringCreateWithCharacters(chars, numChars)
}/* debug [functions.gen.go/function]: JSStringCreateWithCharacters */

// Creates a JavaScript string from a null-terminated UTF-8 string.
//
// Added in macOS 10.5.
// Creates a JavaScript string from a null-terminated UTF-8 string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringCreateWithUTF8CString(_:)
func JSStringCreateWithUTF8CString(string_ unsafe.Pointer) JSStringRef {
	return _JSStringCreateWithUTF8CString(string_)
}/* debug [functions.gen.go/function]: JSStringCreateWithUTF8CString */

// Returns a pointer to the Unicode character buffer that serves as the backing store for a JavaScript string.
//
// Added in macOS 10.5.
// Returns a pointer to the Unicode character buffer that serves as the backing store for a JavaScript string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringGetCharactersPtr(_:)
func JSStringGetCharactersPtr(string_ JSStringRef) unsafe.Pointer {
	return _JSStringGetCharactersPtr(string_)
}/* debug [functions.gen.go/function]: JSStringGetCharactersPtr */

// Returns the number of Unicode characters in a JavaScript string.
//
// Added in macOS 10.5.
// Returns the number of Unicode characters in a JavaScript string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringGetLength(_:)
func JSStringGetLength(string_ JSStringRef) uintptr {
	return _JSStringGetLength(string_)
}/* debug [functions.gen.go/function]: JSStringGetLength */

// Returns the maximum number of bytes a JavaScript string uses when you convert it into a null-terminated UTF-8 string.
//
// Added in macOS 10.5.
// Returns the maximum number of bytes a JavaScript string uses when you convert it into a null-terminated UTF-8 string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringGetMaximumUTF8CStringSize(_:)
func JSStringGetMaximumUTF8CStringSize(string_ JSStringRef) uintptr {
	return _JSStringGetMaximumUTF8CStringSize(string_)
}/* debug [functions.gen.go/function]: JSStringGetMaximumUTF8CStringSize */

// Converts a JavaScript string into a null-terminated UTF-8 string, and copies the result into an external byte buffer.
//
// Added in macOS 10.5.
// Converts a JavaScript string into a null-terminated UTF-8 string, and copies the result into an external byte buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringGetUTF8CString(_:_:_:)
func JSStringGetUTF8CString(string_ JSStringRef, buffer unsafe.Pointer, bufferSize uintptr) uintptr {
	return _JSStringGetUTF8CString(string_, buffer, bufferSize)
}/* debug [functions.gen.go/function]: JSStringGetUTF8CString */

// Tests whether two JavaScript strings match.
//
// Added in macOS 10.5.
// Tests whether two JavaScript strings match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringIsEqual(_:_:)
func JSStringIsEqual(a JSStringRef, b JSStringRef) bool {
	return _JSStringIsEqual(a, b)
}/* debug [functions.gen.go/function]: JSStringIsEqual */

// Tests whether a JavaScript string matches a null-terminated UTF-8 string.
//
// Added in macOS 10.5.
// Tests whether a JavaScript string matches a null-terminated UTF-8 string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringIsEqualToUTF8CString(_:_:)
func JSStringIsEqualToUTF8CString(a JSStringRef, b unsafe.Pointer) bool {
	return _JSStringIsEqualToUTF8CString(a, b)
}/* debug [functions.gen.go/function]: JSStringIsEqualToUTF8CString */

// Releases a JavaScript string.
//
// Added in macOS 10.5.
// Releases a JavaScript string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringRelease(_:)
func JSStringRelease(string_ JSStringRef) {
	_JSStringRelease(string_)
}/* debug [functions.gen.go/function]: JSStringRelease */

// Retains a JavaScript string.
//
// Added in macOS 10.5.
// Retains a JavaScript string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringRetain(_:)
func JSStringRetain(string_ JSStringRef) JSStringRef {
	return _JSStringRetain(string_)
}/* debug [functions.gen.go/function]: JSStringRetain */

// Creates a JavaScript string that contains the JSON-serialized representation of a JavaScript value.
//
// Added in macOS 10.7.
// Creates a JavaScript string that contains the JSON-serialized representation of a JavaScript value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueCreateJSONString(_:_:_:_:)
func JSValueCreateJSONString(ctx JSContextRef, value JSValueRef, indent unsafe.Pointer, exception unsafe.Pointer) JSStringRef {
	return _JSValueCreateJSONString(ctx, value, indent, exception)
}/* debug [functions.gen.go/function]: JSValueCreateJSONString */

// Returns a JavaScript value’s type.
//
// Added in macOS 10.5.
// Returns a JavaScript value’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueGetType(_:_:)
func JSValueGetType(ctx JSContextRef, value JSValueRef) unsafe.Pointer {
	return _JSValueGetType(ctx, value)
}/* debug [functions.gen.go/function]: JSValueGetType */

// Returns a JavaScript value’s typed array type.
//
// Added in macOS 10.12.
// Returns a JavaScript value’s typed array type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueGetTypedArrayType(_:_:_:)
func JSValueGetTypedArrayType(ctx JSContextRef, value JSValueRef, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueGetTypedArrayType(ctx, value, exception)
}/* debug [functions.gen.go/function]: JSValueGetTypedArrayType */

// Tests whether a JavaScript value is an array.
//
// Added in macOS 10.11.
// Tests whether a JavaScript value is an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsArray(_:_:)
func JSValueIsArray(ctx JSContextRef, value JSValueRef) bool {
	return _JSValueIsArray(ctx, value)
}/* debug [functions.gen.go/function]: JSValueIsArray */

// Tests whether a JavaScript value is Boolean.
//
// Added in macOS 10.5.
// Tests whether a JavaScript value is Boolean.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsBoolean(_:_:)
func JSValueIsBoolean(ctx JSContextRef, value JSValueRef) bool {
	return _JSValueIsBoolean(ctx, value)
}/* debug [functions.gen.go/function]: JSValueIsBoolean */

// Tests whether a JavaScript value is a date.
//
// Added in macOS 10.11.
// Tests whether a JavaScript value is a date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsDate(_:_:)
func JSValueIsDate(ctx JSContextRef, value JSValueRef) bool {
	return _JSValueIsDate(ctx, value)
}/* debug [functions.gen.go/function]: JSValueIsDate */

// Tests whether two JavaScript values are equal.
//
// Added in macOS 10.5.
// Tests whether two JavaScript values are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsEqual(_:_:_:_:)
func JSValueIsEqual(ctx JSContextRef, a JSValueRef, b JSValueRef, exception unsafe.Pointer) bool {
	return _JSValueIsEqual(ctx, a, b, exception)
}/* debug [functions.gen.go/function]: JSValueIsEqual */

// Tests whether a JavaScript value is an object that the specified constructor creates.
//
// Added in macOS 10.5.
// Tests whether a JavaScript value is an object that the specified constructor creates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsInstanceOfConstructor(_:_:_:_:)
func JSValueIsInstanceOfConstructor(ctx JSContextRef, value JSValueRef, constructor JSObjectRef, exception unsafe.Pointer) bool {
	return _JSValueIsInstanceOfConstructor(ctx, value, constructor, exception)
}/* debug [functions.gen.go/function]: JSValueIsInstanceOfConstructor */

// Tests whether a JavaScript value’s type is the null type.
//
// Added in macOS 10.5.
// Tests whether a JavaScript value’s type is the null type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsNull(_:_:)
func JSValueIsNull(ctx JSContextRef, value JSValueRef) bool {
	return _JSValueIsNull(ctx, value)
}/* debug [functions.gen.go/function]: JSValueIsNull */

// Tests whether a JavaScript value’s type is the number type.
//
// Added in macOS 10.5.
// Tests whether a JavaScript value’s type is the number type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsNumber(_:_:)
func JSValueIsNumber(ctx JSContextRef, value JSValueRef) bool {
	return _JSValueIsNumber(ctx, value)
}/* debug [functions.gen.go/function]: JSValueIsNumber */

// Tests whether a JavaScript value’s type is the object type.
//
// Added in macOS 10.5.
// Tests whether a JavaScript value’s type is the object type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsObject(_:_:)
func JSValueIsObject(ctx JSContextRef, value JSValueRef) bool {
	return _JSValueIsObject(ctx, value)
}/* debug [functions.gen.go/function]: JSValueIsObject */

// Tests whether a JavaScript value is an object with a specified class in its class chain.
//
// Added in macOS 10.5.
// Tests whether a JavaScript value is an object with a specified class in its class chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsObjectOfClass(_:_:_:)
func JSValueIsObjectOfClass(ctx JSContextRef, value JSValueRef, jsClass JSClassRef) bool {
	return _JSValueIsObjectOfClass(ctx, value, jsClass)
}/* debug [functions.gen.go/function]: JSValueIsObjectOfClass */

// Tests whether two JavaScript values are strict equal.
//
// Added in macOS 10.5.
// Tests whether two JavaScript values are strict equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsStrictEqual(_:_:_:)
func JSValueIsStrictEqual(ctx JSContextRef, a JSValueRef, b JSValueRef) bool {
	return _JSValueIsStrictEqual(ctx, a, b)
}/* debug [functions.gen.go/function]: JSValueIsStrictEqual */

// Tests whether a JavaScript value’s type is the string type.
//
// Added in macOS 10.5.
// Tests whether a JavaScript value’s type is the string type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsString(_:_:)
func JSValueIsString(ctx JSContextRef, value JSValueRef) bool {
	return _JSValueIsString(ctx, value)
}/* debug [functions.gen.go/function]: JSValueIsString */

// Tests whether a JavaScript value’s type is the symbol type.
//
// Added in macOS 10.15.
// Tests whether a JavaScript value’s type is the symbol type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsSymbol(_:_:)
func JSValueIsSymbol(ctx JSContextRef, value JSValueRef) bool {
	return _JSValueIsSymbol(ctx, value)
}/* debug [functions.gen.go/function]: JSValueIsSymbol */

// Tests whether a JavaScript value’s type is the undefined type.
//
// Added in macOS 10.5.
// Tests whether a JavaScript value’s type is the undefined type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsUndefined(_:_:)
func JSValueIsUndefined(ctx JSContextRef, value JSValueRef) bool {
	return _JSValueIsUndefined(ctx, value)
}/* debug [functions.gen.go/function]: JSValueIsUndefined */

// Creates a JavaScript Boolean value.
//
// Added in macOS 10.5.
// Creates a JavaScript Boolean value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeBoolean(_:_:)
func JSValueMakeBoolean(ctx JSContextRef, boolean bool) JSValueRef {
	return _JSValueMakeBoolean(ctx, boolean)
}/* debug [functions.gen.go/function]: JSValueMakeBoolean */

// Creates a JavaScript value from a JSON-formatted string.
//
// Added in macOS 10.7.
// Creates a JavaScript value from a JSON-formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeFromJSONString(_:_:)
func JSValueMakeFromJSONString(ctx JSContextRef, string_ JSStringRef) JSValueRef {
	return _JSValueMakeFromJSONString(ctx, string_)
}/* debug [functions.gen.go/function]: JSValueMakeFromJSONString */

// Creates a JavaScript value of the null type.
//
// Added in macOS 10.5.
// Creates a JavaScript value of the null type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeNull(_:)
func JSValueMakeNull(ctx JSContextRef) JSValueRef {
	return _JSValueMakeNull(ctx)
}/* debug [functions.gen.go/function]: JSValueMakeNull */

// Creates a JavaScript value of the number type.
//
// Added in macOS 10.5.
// Creates a JavaScript value of the number type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeNumber(_:_:)
func JSValueMakeNumber(ctx JSContextRef, number float64) JSValueRef {
	return _JSValueMakeNumber(ctx, number)
}/* debug [functions.gen.go/function]: JSValueMakeNumber */

// Creates a JavaScript value of the string type.
//
// Added in macOS 10.5.
// Creates a JavaScript value of the string type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeString(_:_:)
func JSValueMakeString(ctx JSContextRef, string_ JSStringRef) JSValueRef {
	return _JSValueMakeString(ctx, string_)
}/* debug [functions.gen.go/function]: JSValueMakeString */

// Creates a JavaScript value of the symbol type.
//
// Added in macOS 10.15.
// Creates a JavaScript value of the symbol type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeSymbol(_:_:)
func JSValueMakeSymbol(ctx JSContextRef, description JSStringRef) JSValueRef {
	return _JSValueMakeSymbol(ctx, description)
}/* debug [functions.gen.go/function]: JSValueMakeSymbol */

// Creates a JavaScript value of the undefined type.
//
// Added in macOS 10.5.
// Creates a JavaScript value of the undefined type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeUndefined(_:)
func JSValueMakeUndefined(ctx JSContextRef) JSValueRef {
	return _JSValueMakeUndefined(ctx)
}/* debug [functions.gen.go/function]: JSValueMakeUndefined */

// Protects a JavaScript value from garbage collection.
//
// Added in macOS 10.5.
// Protects a JavaScript value from garbage collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueProtect(_:_:)
func JSValueProtect(ctx JSContextRef, value JSValueRef) {
	_JSValueProtect(ctx, value)
}/* debug [functions.gen.go/function]: JSValueProtect */

// Converts a JavaScript value to a Boolean and returns the resulting Boolean.
//
// Added in macOS 10.5.
// Converts a JavaScript value to a Boolean and returns the resulting Boolean.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToBoolean(_:_:)
func JSValueToBoolean(ctx JSContextRef, value JSValueRef) bool {
	return _JSValueToBoolean(ctx, value)
}/* debug [functions.gen.go/function]: JSValueToBoolean */

// Converts a JavaScript value to a number and returns the resulting number.
//
// Added in macOS 10.5.
// Converts a JavaScript value to a number and returns the resulting number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToNumber(_:_:_:)
func JSValueToNumber(ctx JSContextRef, value JSValueRef, exception unsafe.Pointer) float64 {
	return _JSValueToNumber(ctx, value, exception)
}/* debug [functions.gen.go/function]: JSValueToNumber */

// Converts a JavaScript value to an object and returns the resulting object.
//
// Added in macOS 10.5.
// Converts a JavaScript value to an object and returns the resulting object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToObject(_:_:_:)
func JSValueToObject(ctx JSContextRef, value JSValueRef, exception unsafe.Pointer) JSObjectRef {
	return _JSValueToObject(ctx, value, exception)
}/* debug [functions.gen.go/function]: JSValueToObject */

// Converts a JavaScript value to a string and copies the result into a JavaScript string.
//
// Added in macOS 10.5.
// Converts a JavaScript value to a string and copies the result into a JavaScript string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToStringCopy(_:_:_:)
func JSValueToStringCopy(ctx JSContextRef, value JSValueRef, exception unsafe.Pointer) JSStringRef {
	return _JSValueToStringCopy(ctx, value, exception)
}/* debug [functions.gen.go/function]: JSValueToStringCopy */

// Unprotects a JavaScript value from garbage collection.
//
// Added in macOS 10.5.
// Unprotects a JavaScript value from garbage collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueUnprotect(_:_:)
func JSValueUnprotect(ctx JSContextRef, value JSValueRef) {
	_JSValueUnprotect(ctx, value)
}/* debug [functions.gen.go/function]: JSValueUnprotect */




