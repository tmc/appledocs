// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// JavaScriptCore Functions (112 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_JSBigIntCreateWithDouble func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSBigIntCreateWithInt64 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSBigIntCreateWithString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSBigIntCreateWithUInt64 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSClassCreate func(unsafe.Pointer) unsafe.Pointer
	_JSClassRelease func(unsafe.Pointer) unsafe.Pointer
	_JSClassRetain func(unsafe.Pointer) unsafe.Pointer
	_JSContextGetGlobalContext func(unsafe.Pointer) unsafe.Pointer
	_JSContextGetGlobalObject func(unsafe.Pointer) unsafe.Pointer
	_JSContextGetGroup func(unsafe.Pointer) unsafe.Pointer
	_JSContextGroupCreate func() unsafe.Pointer
	_JSContextGroupRelease func(unsafe.Pointer) unsafe.Pointer
	_JSContextGroupRetain func(unsafe.Pointer) unsafe.Pointer
	_JSEvaluateScript func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSGarbageCollect func(unsafe.Pointer) unsafe.Pointer
	_JSGlobalContextCopyName func(unsafe.Pointer) unsafe.Pointer
	_JSGlobalContextCreate func(unsafe.Pointer) unsafe.Pointer
	_JSGlobalContextCreateInGroup func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSGlobalContextIsInspectable func(unsafe.Pointer) unsafe.Pointer
	_JSGlobalContextRelease func(unsafe.Pointer) unsafe.Pointer
	_JSGlobalContextRetain func(unsafe.Pointer) unsafe.Pointer
	_JSGlobalContextSetInspectable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSGlobalContextSetName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectCallAsConstructor func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectCallAsFunction func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectCopyPropertyNames func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectDeleteProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectDeletePropertyForKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectGetArrayBufferByteLength func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectGetArrayBufferBytesPtr func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectGetPrivate func(unsafe.Pointer) unsafe.Pointer
	_JSObjectGetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectGetPropertyAtIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectGetPropertyForKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectGetPrototype func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectGetTypedArrayBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectGetTypedArrayByteLength func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectGetTypedArrayByteOffset func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectGetTypedArrayBytesPtr func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectGetTypedArrayLength func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectHasProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectHasPropertyForKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectIsConstructor func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectIsFunction func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectMake func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectMakeArray func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectMakeArrayBufferWithBytesNoCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectMakeConstructor func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectMakeDate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectMakeDeferredPromise func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectMakeError func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectMakeFunction func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectMakeFunctionWithCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectMakeRegExp func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectMakeTypedArray func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectMakeTypedArrayWithArrayBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectMakeTypedArrayWithArrayBufferAndOffset func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectMakeTypedArrayWithBytesNoCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectSetPrivate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectSetPropertyAtIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectSetPropertyForKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSObjectSetPrototype func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSPropertyNameAccumulatorAddName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSPropertyNameArrayGetCount func(unsafe.Pointer) unsafe.Pointer
	_JSPropertyNameArrayGetNameAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSPropertyNameArrayRelease func(unsafe.Pointer) unsafe.Pointer
	_JSPropertyNameArrayRetain func(unsafe.Pointer) unsafe.Pointer
	_JSStringCopyCFString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSStringCreateWithCFString func(unsafe.Pointer) unsafe.Pointer
	_JSStringCreateWithCharacters func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSStringCreateWithUTF8CString func(unsafe.Pointer) unsafe.Pointer
	_JSStringGetCharactersPtr func(unsafe.Pointer) unsafe.Pointer
	_JSStringGetLength func(unsafe.Pointer) unsafe.Pointer
	_JSStringGetMaximumUTF8CStringSize func(unsafe.Pointer) unsafe.Pointer
	_JSStringGetUTF8CString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSStringIsEqual func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSStringIsEqualToUTF8CString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSStringRelease func(unsafe.Pointer) unsafe.Pointer
	_JSStringRetain func(unsafe.Pointer) unsafe.Pointer
	_JSValueCompare func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueCompareDouble func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueCompareInt64 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueCompareUInt64 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueCreateJSONString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueGetType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueGetTypedArrayType func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueIsBigInt func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueIsBoolean func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueIsDate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueIsInstanceOfConstructor func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueIsNull func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueIsNumber func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueIsObject func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueIsStrictEqual func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueIsString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueMakeBoolean func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueMakeFromJSONString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueMakeNull func(unsafe.Pointer) unsafe.Pointer
	_JSValueMakeNumber func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueMakeString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueMakeSymbol func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueMakeUndefined func(unsafe.Pointer) unsafe.Pointer
	_JSValueProtect func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueToInt32 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueToInt64 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueToNumber func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueToObject func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueToStringCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueToUInt32 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueToUInt64 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_JSValueUnprotect func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
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
	tryRegister(&_JSValueCompare, lib, "JSValueCompare")
	tryRegister(&_JSValueCompareDouble, lib, "JSValueCompareDouble")
	tryRegister(&_JSValueCompareInt64, lib, "JSValueCompareInt64")
	tryRegister(&_JSValueCompareUInt64, lib, "JSValueCompareUInt64")
	tryRegister(&_JSValueCreateJSONString, lib, "JSValueCreateJSONString")
	tryRegister(&_JSValueGetType, lib, "JSValueGetType")
	tryRegister(&_JSValueGetTypedArrayType, lib, "JSValueGetTypedArrayType")
	tryRegister(&_JSValueIsBigInt, lib, "JSValueIsBigInt")
	tryRegister(&_JSValueIsBoolean, lib, "JSValueIsBoolean")
	tryRegister(&_JSValueIsDate, lib, "JSValueIsDate")
	tryRegister(&_JSValueIsInstanceOfConstructor, lib, "JSValueIsInstanceOfConstructor")
	tryRegister(&_JSValueIsNull, lib, "JSValueIsNull")
	tryRegister(&_JSValueIsNumber, lib, "JSValueIsNumber")
	tryRegister(&_JSValueIsObject, lib, "JSValueIsObject")
	tryRegister(&_JSValueIsStrictEqual, lib, "JSValueIsStrictEqual")
	tryRegister(&_JSValueIsString, lib, "JSValueIsString")
	tryRegister(&_JSValueMakeBoolean, lib, "JSValueMakeBoolean")
	tryRegister(&_JSValueMakeFromJSONString, lib, "JSValueMakeFromJSONString")
	tryRegister(&_JSValueMakeNull, lib, "JSValueMakeNull")
	tryRegister(&_JSValueMakeNumber, lib, "JSValueMakeNumber")
	tryRegister(&_JSValueMakeString, lib, "JSValueMakeString")
	tryRegister(&_JSValueMakeSymbol, lib, "JSValueMakeSymbol")
	tryRegister(&_JSValueMakeUndefined, lib, "JSValueMakeUndefined")
	tryRegister(&_JSValueProtect, lib, "JSValueProtect")
	tryRegister(&_JSValueToInt32, lib, "JSValueToInt32")
	tryRegister(&_JSValueToInt64, lib, "JSValueToInt64")
	tryRegister(&_JSValueToNumber, lib, "JSValueToNumber")
	tryRegister(&_JSValueToObject, lib, "JSValueToObject")
	tryRegister(&_JSValueToStringCopy, lib, "JSValueToStringCopy")
	tryRegister(&_JSValueToUInt32, lib, "JSValueToUInt32")
	tryRegister(&_JSValueToUInt64, lib, "JSValueToUInt64")
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



// JSBigIntCreateWithDouble is a JavaScriptCore function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithDouble(_:_:_:)
func JSBigIntCreateWithDouble(ctx unsafe.Pointer, value unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSBigIntCreateWithDouble(ctx, value, exception)
	}


// JSBigIntCreateWithInt64 is a JavaScriptCore function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithInt64(_:_:_:)
func JSBigIntCreateWithInt64(ctx unsafe.Pointer, integer unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSBigIntCreateWithInt64(ctx, integer, exception)
	}


// JSBigIntCreateWithString is a JavaScriptCore function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithString(_:_:_:)
func JSBigIntCreateWithString(ctx unsafe.Pointer, string_ unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSBigIntCreateWithString(ctx, string_, exception)
	}


// JSBigIntCreateWithUInt64 is a JavaScriptCore function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithUInt64(_:_:_:)
func JSBigIntCreateWithUInt64(ctx unsafe.Pointer, integer unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSBigIntCreateWithUInt64(ctx, integer, exception)
	}


// Creates a JavaScript class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSClassCreate(_:)
func JSClassCreate(definition unsafe.Pointer) unsafe.Pointer {
	return _JSClassCreate(definition)
	}


// Releases a JavaScript class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSClassRelease(_:)
func JSClassRelease(jsClass unsafe.Pointer) {
	_JSClassRelease(jsClass)
	}


// Retains a JavaScript class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSClassRetain(_:)
func JSClassRetain(jsClass unsafe.Pointer) unsafe.Pointer {
	return _JSClassRetain(jsClass)
	}


// Gets the global context of a JavaScript execution context. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContextGetGlobalContext(_:)
func JSContextGetGlobalContext(ctx unsafe.Pointer) unsafe.Pointer {
	return _JSContextGetGlobalContext(ctx)
	}


// Gets the global object of a JavaScript execution context. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContextGetGlobalObject(_:)
func JSContextGetGlobalObject(ctx unsafe.Pointer) unsafe.Pointer {
	return _JSContextGetGlobalObject(ctx)
	}


// Gets the context group that a JavaScript execution context belongs to. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContextGetGroup(_:)
func JSContextGetGroup(ctx unsafe.Pointer) unsafe.Pointer {
	return _JSContextGetGroup(ctx)
	}


// Creates a JavaScript context group. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContextGroupCreate()
func JSContextGroupCreate() unsafe.Pointer {
	return _JSContextGroupCreate()
	}


// Releases a JavaScript context group. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContextGroupRelease(_:)
func JSContextGroupRelease(group unsafe.Pointer) {
	_JSContextGroupRelease(group)
	}


// Retains a JavaScript context group. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContextGroupRetain(_:)
func JSContextGroupRetain(group unsafe.Pointer) unsafe.Pointer {
	return _JSContextGroupRetain(group)
	}


// Evaluates a string of JavaScript. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSEvaluateScript(_:_:_:_:_:_:)
func JSEvaluateScript(ctx unsafe.Pointer, script unsafe.Pointer, thisObject unsafe.Pointer, sourceURL unsafe.Pointer, startingLineNumber unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSEvaluateScript(ctx, script, thisObject, sourceURL, startingLineNumber, exception)
	}


// Performs a JavaScript garbage collection. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGarbageCollect(_:)
func JSGarbageCollect(ctx unsafe.Pointer) {
	_JSGarbageCollect(ctx)
	}


// Gets a copy of the name of a context. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextCopyName(_:)
func JSGlobalContextCopyName(ctx unsafe.Pointer) unsafe.Pointer {
	return _JSGlobalContextCopyName(ctx)
	}


// Creates a global JavaScript execution context. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextCreate(_:)
func JSGlobalContextCreate(globalObjectClass unsafe.Pointer) unsafe.Pointer {
	return _JSGlobalContextCreate(globalObjectClass)
	}


// Creates a global JavaScript execution context in the provided context group. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextCreateInGroup(_:_:)
func JSGlobalContextCreateInGroup(group unsafe.Pointer, globalObjectClass unsafe.Pointer) unsafe.Pointer {
	return _JSGlobalContextCreateInGroup(group, globalObjectClass)
	}


// Returns a Boolean value that indicates whether the JavaScript context is inspectable. [Full Topic]
//
// Added in macOS 13.3.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextIsInspectable(_:)
func JSGlobalContextIsInspectable(ctx unsafe.Pointer) unsafe.Pointer {
	return _JSGlobalContextIsInspectable(ctx)
	}


// Releases a global JavaScript execution context. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextRelease(_:)
func JSGlobalContextRelease(ctx unsafe.Pointer) {
	_JSGlobalContextRelease(ctx)
	}


// Retains a global JavaScript execution context. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextRetain(_:)
func JSGlobalContextRetain(ctx unsafe.Pointer) unsafe.Pointer {
	return _JSGlobalContextRetain(ctx)
	}


// Sets a JavaScript context to be either inspectable or not inspectable. [Full Topic]
//
// Added in macOS 13.3.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextSetInspectable(_:_:)
func JSGlobalContextSetInspectable(ctx unsafe.Pointer, inspectable unsafe.Pointer) {
	_JSGlobalContextSetInspectable(ctx, inspectable)
	}


// Sets the remote debugging name for a context. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextSetName(_:_:)
func JSGlobalContextSetName(ctx unsafe.Pointer, name unsafe.Pointer) {
	_JSGlobalContextSetName(ctx, name)
	}


// Calls an object as a constructor. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectCallAsConstructor(_:_:_:_:_:)
func JSObjectCallAsConstructor(ctx unsafe.Pointer, object unsafe.Pointer, argumentCount unsafe.Pointer, arguments unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectCallAsConstructor(ctx, object, argumentCount, arguments, exception)
	}


// Calls an object as a function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectCallAsFunction(_:_:_:_:_:_:)
func JSObjectCallAsFunction(ctx unsafe.Pointer, object unsafe.Pointer, thisObject unsafe.Pointer, argumentCount unsafe.Pointer, arguments unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectCallAsFunction(ctx, object, thisObject, argumentCount, arguments, exception)
	}


// Gets the names of an object’s enumerable properties. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectCopyPropertyNames(_:_:)
func JSObjectCopyPropertyNames(ctx unsafe.Pointer, object unsafe.Pointer) unsafe.Pointer {
	return _JSObjectCopyPropertyNames(ctx, object)
	}


// Deletes a property from an object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectDeleteProperty(_:_:_:_:)
func JSObjectDeleteProperty(ctx unsafe.Pointer, object unsafe.Pointer, propertyName unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectDeleteProperty(ctx, object, propertyName, exception)
	}


// Deletes a property from an object using a JavaScript value as the property key. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectDeletePropertyForKey(_:_:_:_:)
func JSObjectDeletePropertyForKey(ctx unsafe.Pointer, object unsafe.Pointer, propertyKey unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectDeletePropertyForKey(ctx, object, propertyKey, exception)
	}


// Returns the number of bytes in a JavaScript data object. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetArrayBufferByteLength(_:_:_:)
func JSObjectGetArrayBufferByteLength(ctx unsafe.Pointer, object unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectGetArrayBufferByteLength(ctx, object, exception)
	}


// Returns a pointer to the data buffer that serves as the backing store for a JavaScript typed array object. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetArrayBufferBytesPtr(_:_:_:)
func JSObjectGetArrayBufferBytesPtr(ctx unsafe.Pointer, object unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectGetArrayBufferBytesPtr(ctx, object, exception)
	}


// Gets an object’s private data. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPrivate(_:)
func JSObjectGetPrivate(object unsafe.Pointer) unsafe.Pointer {
	return _JSObjectGetPrivate(object)
	}


// Gets a property from an object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetProperty(_:_:_:_:)
func JSObjectGetProperty(ctx unsafe.Pointer, object unsafe.Pointer, propertyName unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectGetProperty(ctx, object, propertyName, exception)
	}


// Gets a property from an object by numeric index. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPropertyAtIndex(_:_:_:_:)
func JSObjectGetPropertyAtIndex(ctx unsafe.Pointer, object unsafe.Pointer, propertyIndex unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectGetPropertyAtIndex(ctx, object, propertyIndex, exception)
	}


// Gets a property from an object using a JavaScript value as the property key. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPropertyForKey(_:_:_:_:)
func JSObjectGetPropertyForKey(ctx unsafe.Pointer, object unsafe.Pointer, propertyKey unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectGetPropertyForKey(ctx, object, propertyKey, exception)
	}


// Gets an object’s prototype. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPrototype(_:_:)
func JSObjectGetPrototype(ctx unsafe.Pointer, object unsafe.Pointer) unsafe.Pointer {
	return _JSObjectGetPrototype(ctx, object)
	}


// Returns the JavaScript array buffer object to use as the backing of a JavaScript typed array object. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayBuffer(_:_:_:)
func JSObjectGetTypedArrayBuffer(ctx unsafe.Pointer, object unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectGetTypedArrayBuffer(ctx, object, exception)
	}


// Returns the byte length of a JavaScript typed array object. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayByteLength(_:_:_:)
func JSObjectGetTypedArrayByteLength(ctx unsafe.Pointer, object unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectGetTypedArrayByteLength(ctx, object, exception)
	}


// Returns the byte offset of a JavaScript typed array object. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayByteOffset(_:_:_:)
func JSObjectGetTypedArrayByteOffset(ctx unsafe.Pointer, object unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectGetTypedArrayByteOffset(ctx, object, exception)
	}


// Returns a temporary pointer to the backing store of a JavaScript typed array object. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayBytesPtr(_:_:_:)
func JSObjectGetTypedArrayBytesPtr(ctx unsafe.Pointer, object unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectGetTypedArrayBytesPtr(ctx, object, exception)
	}


// Returns the length of a JavaScript typed array object. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayLength(_:_:_:)
func JSObjectGetTypedArrayLength(ctx unsafe.Pointer, object unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectGetTypedArrayLength(ctx, object, exception)
	}


// Tests whether an object has a specified property. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectHasProperty(_:_:_:)
func JSObjectHasProperty(ctx unsafe.Pointer, object unsafe.Pointer, propertyName unsafe.Pointer) unsafe.Pointer {
	return _JSObjectHasProperty(ctx, object, propertyName)
	}


// Tests whether an object has the specified property using a JavaScript value as the property key. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectHasPropertyForKey(_:_:_:_:)
func JSObjectHasPropertyForKey(ctx unsafe.Pointer, object unsafe.Pointer, propertyKey unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectHasPropertyForKey(ctx, object, propertyKey, exception)
	}


// Tests whether you can call an object as a constructor. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectIsConstructor(_:_:)
func JSObjectIsConstructor(ctx unsafe.Pointer, object unsafe.Pointer) unsafe.Pointer {
	return _JSObjectIsConstructor(ctx, object)
	}


// Tests whether you can call an object as a function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectIsFunction(_:_:)
func JSObjectIsFunction(ctx unsafe.Pointer, object unsafe.Pointer) unsafe.Pointer {
	return _JSObjectIsFunction(ctx, object)
	}


// Creates a JavaScript object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMake(_:_:_:)
func JSObjectMake(ctx unsafe.Pointer, jsClass unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _JSObjectMake(ctx, jsClass, data)
	}


// Creates a JavaScript array object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeArray(_:_:_:_:)
func JSObjectMakeArray(ctx unsafe.Pointer, argumentCount unsafe.Pointer, arguments unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectMakeArray(ctx, argumentCount, arguments, exception)
	}


// Creates a JavaScript array buffer object from an existing pointer. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeArrayBufferWithBytesNoCopy(_:_:_:_:_:_:)
func JSObjectMakeArrayBufferWithBytesNoCopy(ctx unsafe.Pointer, bytes unsafe.Pointer, byteLength unsafe.Pointer, bytesDeallocator unsafe.Pointer, deallocatorContext unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectMakeArrayBufferWithBytesNoCopy(ctx, bytes, byteLength, bytesDeallocator, deallocatorContext, exception)
	}


// Creates a JavaScript constructor. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeConstructor(_:_:_:)
func JSObjectMakeConstructor(ctx unsafe.Pointer, jsClass unsafe.Pointer, callAsConstructor unsafe.Pointer) unsafe.Pointer {
	return _JSObjectMakeConstructor(ctx, jsClass, callAsConstructor)
	}


// Creates a JavaScript date object as though invoking the built-in date constructor. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeDate(_:_:_:_:)
func JSObjectMakeDate(ctx unsafe.Pointer, argumentCount unsafe.Pointer, arguments unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectMakeDate(ctx, argumentCount, arguments, exception)
	}


// Creates a JavaScript promise object by invoking the provided executor. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeDeferredPromise(_:_:_:_:)
func JSObjectMakeDeferredPromise(ctx unsafe.Pointer, resolve unsafe.Pointer, reject unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectMakeDeferredPromise(ctx, resolve, reject, exception)
	}


// Creates a JavaScript error object as though invoking the built-in error constructor. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeError(_:_:_:_:)
func JSObjectMakeError(ctx unsafe.Pointer, argumentCount unsafe.Pointer, arguments unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectMakeError(ctx, argumentCount, arguments, exception)
	}


// Creates a function with a specified script as its body. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeFunction(_:_:_:_:_:_:_:_:)
func JSObjectMakeFunction(ctx unsafe.Pointer, name unsafe.Pointer, parameterCount unsafe.Pointer, parameterNames unsafe.Pointer, body unsafe.Pointer, sourceURL unsafe.Pointer, startingLineNumber unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectMakeFunction(ctx, name, parameterCount, parameterNames, body, sourceURL, startingLineNumber, exception)
	}


// Creates a JavaScript function with a specified callback as its implementation. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeFunctionWithCallback(_:_:_:)
func JSObjectMakeFunctionWithCallback(ctx unsafe.Pointer, name unsafe.Pointer, callAsFunction unsafe.Pointer) unsafe.Pointer {
	return _JSObjectMakeFunctionWithCallback(ctx, name, callAsFunction)
	}


// Creates a JavaScript regular expression object as though invoking the built-in regular expression constructor. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeRegExp(_:_:_:_:)
func JSObjectMakeRegExp(ctx unsafe.Pointer, argumentCount unsafe.Pointer, arguments unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectMakeRegExp(ctx, argumentCount, arguments, exception)
	}


// Creates a JavaScript typed array object with the specified number of elements. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeTypedArray(_:_:_:_:)
func JSObjectMakeTypedArray(ctx unsafe.Pointer, arrayType unsafe.Pointer, length unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectMakeTypedArray(ctx, arrayType, length, exception)
	}


// Creates a JavaScript typed array object from an existing JavaScript array buffer object. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeTypedArrayWithArrayBuffer(_:_:_:_:)
func JSObjectMakeTypedArrayWithArrayBuffer(ctx unsafe.Pointer, arrayType unsafe.Pointer, buffer unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectMakeTypedArrayWithArrayBuffer(ctx, arrayType, buffer, exception)
	}


// Creates a JavaScript typed array object from an existing JavaScript array buffer object with the specified offset and length. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeTypedArrayWithArrayBufferAndOffset(_:_:_:_:_:_:)
func JSObjectMakeTypedArrayWithArrayBufferAndOffset(ctx unsafe.Pointer, arrayType unsafe.Pointer, buffer unsafe.Pointer, byteOffset unsafe.Pointer, length unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectMakeTypedArrayWithArrayBufferAndOffset(ctx, arrayType, buffer, byteOffset, length, exception)
	}


// Creates a JavaScript typed array object from an existing pointer. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeTypedArrayWithBytesNoCopy(_:_:_:_:_:_:_:)
func JSObjectMakeTypedArrayWithBytesNoCopy(ctx unsafe.Pointer, arrayType unsafe.Pointer, bytes unsafe.Pointer, byteLength unsafe.Pointer, bytesDeallocator unsafe.Pointer, deallocatorContext unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSObjectMakeTypedArrayWithBytesNoCopy(ctx, arrayType, bytes, byteLength, bytesDeallocator, deallocatorContext, exception)
	}


// Sets a pointer to private data on an object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetPrivate(_:_:)
func JSObjectSetPrivate(object unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _JSObjectSetPrivate(object, data)
	}


// Sets a property on an object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetProperty(_:_:_:_:_:_:)
func JSObjectSetProperty(ctx unsafe.Pointer, object unsafe.Pointer, propertyName unsafe.Pointer, value unsafe.Pointer, attributes unsafe.Pointer, exception unsafe.Pointer) {
	_JSObjectSetProperty(ctx, object, propertyName, value, attributes, exception)
	}


// Sets a property on an object by numeric index. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetPropertyAtIndex(_:_:_:_:_:)
func JSObjectSetPropertyAtIndex(ctx unsafe.Pointer, object unsafe.Pointer, propertyIndex unsafe.Pointer, value unsafe.Pointer, exception unsafe.Pointer) {
	_JSObjectSetPropertyAtIndex(ctx, object, propertyIndex, value, exception)
	}


// Sets a property on an object using a JavaScript value as the property key. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetPropertyForKey(_:_:_:_:_:_:)
func JSObjectSetPropertyForKey(ctx unsafe.Pointer, object unsafe.Pointer, propertyKey unsafe.Pointer, value unsafe.Pointer, attributes unsafe.Pointer, exception unsafe.Pointer) {
	_JSObjectSetPropertyForKey(ctx, object, propertyKey, value, attributes, exception)
	}


// Sets an object’s prototype. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetPrototype(_:_:_:)
func JSObjectSetPrototype(ctx unsafe.Pointer, object unsafe.Pointer, value unsafe.Pointer) {
	_JSObjectSetPrototype(ctx, object, value)
	}


// Adds a property name to a JavaScript property name accumulator. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameAccumulatorAddName(_:_:)
func JSPropertyNameAccumulatorAddName(accumulator unsafe.Pointer, propertyName unsafe.Pointer) {
	_JSPropertyNameAccumulatorAddName(accumulator, propertyName)
	}


// Gets a count of the number of items in a JavaScript property name array. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameArrayGetCount(_:)
func JSPropertyNameArrayGetCount(array unsafe.Pointer) unsafe.Pointer {
	return _JSPropertyNameArrayGetCount(array)
	}


// Gets a property name at a specified index in a JavaScript property name array. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameArrayGetNameAtIndex(_:_:)
func JSPropertyNameArrayGetNameAtIndex(array unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	return _JSPropertyNameArrayGetNameAtIndex(array, index)
	}


// Releases a JavaScript property name array. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameArrayRelease(_:)
func JSPropertyNameArrayRelease(array unsafe.Pointer) {
	_JSPropertyNameArrayRelease(array)
	}


// Retains a JavaScript property name array. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameArrayRetain(_:)
func JSPropertyNameArrayRetain(array unsafe.Pointer) unsafe.Pointer {
	return _JSPropertyNameArrayRetain(array)
	}


// Creates a Core Foundation string from a JavaScript string. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringCopyCFString(_:_:)
func JSStringCopyCFString(alloc unsafe.Pointer, string_ unsafe.Pointer) unsafe.Pointer {
	return _JSStringCopyCFString(alloc, string_)
	}


// Creates a JavaScript string from a Core Foundation string. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringCreateWithCFString(_:)
func JSStringCreateWithCFString(string_ unsafe.Pointer) unsafe.Pointer {
	return _JSStringCreateWithCFString(string_)
	}


// Creates a JavaScript string from a buffer of Unicode characters. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringCreateWithCharacters(_:_:)
func JSStringCreateWithCharacters(chars unsafe.Pointer, numChars unsafe.Pointer) unsafe.Pointer {
	return _JSStringCreateWithCharacters(chars, numChars)
	}


// Creates a JavaScript string from a null-terminated UTF-8 string. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringCreateWithUTF8CString(_:)
func JSStringCreateWithUTF8CString(string_ unsafe.Pointer) unsafe.Pointer {
	return _JSStringCreateWithUTF8CString(string_)
	}


// Returns a pointer to the Unicode character buffer that serves as the backing store for a JavaScript string. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringGetCharactersPtr(_:)
func JSStringGetCharactersPtr(string_ unsafe.Pointer) unsafe.Pointer {
	return _JSStringGetCharactersPtr(string_)
	}


// Returns the number of Unicode characters in a JavaScript string. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringGetLength(_:)
func JSStringGetLength(string_ unsafe.Pointer) unsafe.Pointer {
	return _JSStringGetLength(string_)
	}


// Returns the maximum number of bytes a JavaScript string uses when you convert it into a null-terminated UTF-8 string. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringGetMaximumUTF8CStringSize(_:)
func JSStringGetMaximumUTF8CStringSize(string_ unsafe.Pointer) unsafe.Pointer {
	return _JSStringGetMaximumUTF8CStringSize(string_)
	}


// Converts a JavaScript string into a null-terminated UTF-8 string, and copies the result into an external byte buffer. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringGetUTF8CString(_:_:_:)
func JSStringGetUTF8CString(string_ unsafe.Pointer, buffer unsafe.Pointer, bufferSize unsafe.Pointer) unsafe.Pointer {
	return _JSStringGetUTF8CString(string_, buffer, bufferSize)
	}


// Tests whether two JavaScript strings match. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringIsEqual(_:_:)
func JSStringIsEqual(a unsafe.Pointer, b unsafe.Pointer) unsafe.Pointer {
	return _JSStringIsEqual(a, b)
	}


// Tests whether a JavaScript string matches a null-terminated UTF-8 string. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringIsEqualToUTF8CString(_:_:)
func JSStringIsEqualToUTF8CString(a unsafe.Pointer, b unsafe.Pointer) unsafe.Pointer {
	return _JSStringIsEqualToUTF8CString(a, b)
	}


// Releases a JavaScript string. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringRelease(_:)
func JSStringRelease(string_ unsafe.Pointer) {
	_JSStringRelease(string_)
	}


// Retains a JavaScript string. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringRetain(_:)
func JSStringRetain(string_ unsafe.Pointer) unsafe.Pointer {
	return _JSStringRetain(string_)
	}


// JSValueCompare is a JavaScriptCore function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompare(_:_:_:_:)
func JSValueCompare(ctx unsafe.Pointer, left unsafe.Pointer, right unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueCompare(ctx, left, right, exception)
	}


// JSValueCompareDouble is a JavaScriptCore function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompareDouble(_:_:_:_:)
func JSValueCompareDouble(ctx unsafe.Pointer, left unsafe.Pointer, right unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueCompareDouble(ctx, left, right, exception)
	}


// JSValueCompareInt64 is a JavaScriptCore function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompareInt64(_:_:_:_:)
func JSValueCompareInt64(ctx unsafe.Pointer, left unsafe.Pointer, right unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueCompareInt64(ctx, left, right, exception)
	}


// JSValueCompareUInt64 is a JavaScriptCore function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompareUInt64(_:_:_:_:)
func JSValueCompareUInt64(ctx unsafe.Pointer, left unsafe.Pointer, right unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueCompareUInt64(ctx, left, right, exception)
	}


// Creates a JavaScript string that contains the JSON-serialized representation of a JavaScript value. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueCreateJSONString(_:_:_:_:)
func JSValueCreateJSONString(ctx unsafe.Pointer, value unsafe.Pointer, indent unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueCreateJSONString(ctx, value, indent, exception)
	}


// Returns a JavaScript value’s type. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueGetType(_:_:)
func JSValueGetType(ctx unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _JSValueGetType(ctx, value)
	}


// Returns a JavaScript value’s typed array type. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueGetTypedArrayType(_:_:_:)
func JSValueGetTypedArrayType(ctx unsafe.Pointer, value unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueGetTypedArrayType(ctx, value, exception)
	}


// JSValueIsBigInt is a JavaScriptCore function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsBigInt(_:_:)
func JSValueIsBigInt(ctx unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _JSValueIsBigInt(ctx, value)
	}


// Tests whether a JavaScript value is Boolean. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsBoolean(_:_:)
func JSValueIsBoolean(ctx unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _JSValueIsBoolean(ctx, value)
	}


// Tests whether a JavaScript value is a date. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsDate(_:_:)
func JSValueIsDate(ctx unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _JSValueIsDate(ctx, value)
	}


// Tests whether a JavaScript value is an object that the specified constructor creates. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsInstanceOfConstructor(_:_:_:_:)
func JSValueIsInstanceOfConstructor(ctx unsafe.Pointer, value unsafe.Pointer, constructor unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueIsInstanceOfConstructor(ctx, value, constructor, exception)
	}


// Tests whether a JavaScript value’s type is the null type. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsNull(_:_:)
func JSValueIsNull(ctx unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _JSValueIsNull(ctx, value)
	}


// Tests whether a JavaScript value’s type is the number type. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsNumber(_:_:)
func JSValueIsNumber(ctx unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _JSValueIsNumber(ctx, value)
	}


// Tests whether a JavaScript value’s type is the object type. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsObject(_:_:)
func JSValueIsObject(ctx unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _JSValueIsObject(ctx, value)
	}


// Tests whether two JavaScript values are strict equal. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsStrictEqual(_:_:_:)
func JSValueIsStrictEqual(ctx unsafe.Pointer, a unsafe.Pointer, b unsafe.Pointer) unsafe.Pointer {
	return _JSValueIsStrictEqual(ctx, a, b)
	}


// Tests whether a JavaScript value’s type is the string type. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsString(_:_:)
func JSValueIsString(ctx unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _JSValueIsString(ctx, value)
	}


// Creates a JavaScript Boolean value. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeBoolean(_:_:)
func JSValueMakeBoolean(ctx unsafe.Pointer, boolean unsafe.Pointer) unsafe.Pointer {
	return _JSValueMakeBoolean(ctx, boolean)
	}


// Creates a JavaScript value from a JSON-formatted string. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeFromJSONString(_:_:)
func JSValueMakeFromJSONString(ctx unsafe.Pointer, string_ unsafe.Pointer) unsafe.Pointer {
	return _JSValueMakeFromJSONString(ctx, string_)
	}


// Creates a JavaScript value of the null type. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeNull(_:)
func JSValueMakeNull(ctx unsafe.Pointer) unsafe.Pointer {
	return _JSValueMakeNull(ctx)
	}


// Creates a JavaScript value of the number type. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeNumber(_:_:)
func JSValueMakeNumber(ctx unsafe.Pointer, number unsafe.Pointer) unsafe.Pointer {
	return _JSValueMakeNumber(ctx, number)
	}


// Creates a JavaScript value of the string type. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeString(_:_:)
func JSValueMakeString(ctx unsafe.Pointer, string_ unsafe.Pointer) unsafe.Pointer {
	return _JSValueMakeString(ctx, string_)
	}


// Creates a JavaScript value of the symbol type. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeSymbol(_:_:)
func JSValueMakeSymbol(ctx unsafe.Pointer, description unsafe.Pointer) unsafe.Pointer {
	return _JSValueMakeSymbol(ctx, description)
	}


// Creates a JavaScript value of the undefined type. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeUndefined(_:)
func JSValueMakeUndefined(ctx unsafe.Pointer) unsafe.Pointer {
	return _JSValueMakeUndefined(ctx)
	}


// Protects a JavaScript value from garbage collection. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueProtect(_:_:)
func JSValueProtect(ctx unsafe.Pointer, value unsafe.Pointer) {
	_JSValueProtect(ctx, value)
	}


// JSValueToInt32 is a JavaScriptCore function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToInt32(_:_:_:)
func JSValueToInt32(ctx unsafe.Pointer, value unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueToInt32(ctx, value, exception)
	}


// JSValueToInt64 is a JavaScriptCore function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToInt64(_:_:_:)
func JSValueToInt64(ctx unsafe.Pointer, value unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueToInt64(ctx, value, exception)
	}


// Converts a JavaScript value to a number and returns the resulting number. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToNumber(_:_:_:)
func JSValueToNumber(ctx unsafe.Pointer, value unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueToNumber(ctx, value, exception)
	}


// Converts a JavaScript value to an object and returns the resulting object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToObject(_:_:_:)
func JSValueToObject(ctx unsafe.Pointer, value unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueToObject(ctx, value, exception)
	}


// Converts a JavaScript value to a string and copies the result into a JavaScript string. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToStringCopy(_:_:_:)
func JSValueToStringCopy(ctx unsafe.Pointer, value unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueToStringCopy(ctx, value, exception)
	}


// JSValueToUInt32 is a JavaScriptCore function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToUInt32(_:_:_:)
func JSValueToUInt32(ctx unsafe.Pointer, value unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueToUInt32(ctx, value, exception)
	}


// JSValueToUInt64 is a JavaScriptCore function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToUInt64(_:_:_:)
func JSValueToUInt64(ctx unsafe.Pointer, value unsafe.Pointer, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueToUInt64(ctx, value, exception)
	}


// Unprotects a JavaScript value from garbage collection. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueUnprotect(_:_:)
func JSValueUnprotect(ctx unsafe.Pointer, value unsafe.Pointer) {
	_JSValueUnprotect(ctx, value)
	}




