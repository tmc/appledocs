// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// JavaScriptCore Functions (15 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_JSBigIntCreateWithDouble func(JSContextRef, float64, unsafe.Pointer) JSValueRef
	_JSBigIntCreateWithInt64 func(JSContextRef, int64, unsafe.Pointer) JSValueRef
	_JSBigIntCreateWithString func(JSContextRef, unsafe.Pointer, unsafe.Pointer) JSValueRef
	_JSBigIntCreateWithUInt64 func(JSContextRef, uint64, unsafe.Pointer) JSValueRef
	_JSEvaluateScript func(JSContextRef, unsafe.Pointer, JSObjectRef, unsafe.Pointer, int, unsafe.Pointer) JSValueRef
	_JSValueCompare func(JSContextRef, JSValueRef, JSValueRef, unsafe.Pointer) unsafe.Pointer
	_JSValueCompareDouble func(JSContextRef, JSValueRef, float64, unsafe.Pointer) unsafe.Pointer
	_JSValueCompareInt64 func(JSContextRef, JSValueRef, int64, unsafe.Pointer) unsafe.Pointer
	_JSValueCompareUInt64 func(JSContextRef, JSValueRef, uint64, unsafe.Pointer) unsafe.Pointer
	_JSValueIsBigInt func(JSContextRef, JSValueRef) bool
	_JSValueToInt32 func(JSContextRef, JSValueRef, unsafe.Pointer) int32
	_JSValueToInt64 func(JSContextRef, JSValueRef, unsafe.Pointer) int64
	_JSValueToUInt32 func(JSContextRef, JSValueRef, unsafe.Pointer) uint32
	_JSValueToUInt64 func(JSContextRef, JSValueRef, unsafe.Pointer) uint64
	_JSGarbageCollect func(JSContextRef) unsafe.Pointer
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
	tryRegister(&_JSEvaluateScript, lib, "JSEvaluateScript")
	tryRegister(&_JSValueCompare, lib, "JSValueCompare")
	tryRegister(&_JSValueCompareDouble, lib, "JSValueCompareDouble")
	tryRegister(&_JSValueCompareInt64, lib, "JSValueCompareInt64")
	tryRegister(&_JSValueCompareUInt64, lib, "JSValueCompareUInt64")
	tryRegister(&_JSValueIsBigInt, lib, "JSValueIsBigInt")
	tryRegister(&_JSValueToInt32, lib, "JSValueToInt32")
	tryRegister(&_JSValueToInt64, lib, "JSValueToInt64")
	tryRegister(&_JSValueToUInt32, lib, "JSValueToUInt32")
	tryRegister(&_JSValueToUInt64, lib, "JSValueToUInt64")
	tryRegister(&_JSGarbageCollect, lib, "JSGarbageCollect")
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
}

// JSBigIntCreateWithInt64 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithInt64(_:_:_:)
func JSBigIntCreateWithInt64(ctx JSContextRef, integer int64, exception unsafe.Pointer) JSValueRef {
	return _JSBigIntCreateWithInt64(ctx, integer, exception)
}

// JSBigIntCreateWithString is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithString(_:_:_:)
func JSBigIntCreateWithString(ctx JSContextRef, string_ unsafe.Pointer, exception unsafe.Pointer) JSValueRef {
	return _JSBigIntCreateWithString(ctx, string_, exception)
}

// JSBigIntCreateWithUInt64 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithUInt64(_:_:_:)
func JSBigIntCreateWithUInt64(ctx JSContextRef, integer uint64, exception unsafe.Pointer) JSValueRef {
	return _JSBigIntCreateWithUInt64(ctx, integer, exception)
}

// Evaluates a string of JavaScript.
//
// Added in macOS 10.5.
// Evaluates a string of JavaScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSEvaluateScript(_:_:_:_:_:_:)
func JSEvaluateScript(ctx JSContextRef, script unsafe.Pointer, thisObject JSObjectRef, sourceURL unsafe.Pointer, startingLineNumber int, exception unsafe.Pointer) JSValueRef {
	return _JSEvaluateScript(ctx, script, thisObject, sourceURL, startingLineNumber, exception)
}

// JSValueCompare is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompare(_:_:_:_:)
func JSValueCompare(ctx JSContextRef, left JSValueRef, right JSValueRef, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueCompare(ctx, left, right, exception)
}

// JSValueCompareDouble is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompareDouble(_:_:_:_:)
func JSValueCompareDouble(ctx JSContextRef, left JSValueRef, right float64, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueCompareDouble(ctx, left, right, exception)
}

// JSValueCompareInt64 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompareInt64(_:_:_:_:)
func JSValueCompareInt64(ctx JSContextRef, left JSValueRef, right int64, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueCompareInt64(ctx, left, right, exception)
}

// JSValueCompareUInt64 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompareUInt64(_:_:_:_:)
func JSValueCompareUInt64(ctx JSContextRef, left JSValueRef, right uint64, exception unsafe.Pointer) unsafe.Pointer {
	return _JSValueCompareUInt64(ctx, left, right, exception)
}

// JSValueIsBigInt is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsBigInt(_:_:)
func JSValueIsBigInt(ctx JSContextRef, value JSValueRef) bool {
	return _JSValueIsBigInt(ctx, value)
}

// JSValueToInt32 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToInt32(_:_:_:)
func JSValueToInt32(ctx JSContextRef, value JSValueRef, exception unsafe.Pointer) int32 {
	return _JSValueToInt32(ctx, value, exception)
}

// JSValueToInt64 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToInt64(_:_:_:)
func JSValueToInt64(ctx JSContextRef, value JSValueRef, exception unsafe.Pointer) int64 {
	return _JSValueToInt64(ctx, value, exception)
}

// JSValueToUInt32 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToUInt32(_:_:_:)
func JSValueToUInt32(ctx JSContextRef, value JSValueRef, exception unsafe.Pointer) uint32 {
	return _JSValueToUInt32(ctx, value, exception)
}

// JSValueToUInt64 is a JavaScriptCore function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueToUInt64(_:_:_:)
func JSValueToUInt64(ctx JSContextRef, value JSValueRef, exception unsafe.Pointer) uint64 {
	return _JSValueToUInt64(ctx, value, exception)
}

// Performs a JavaScript garbage collection.

// Performs a JavaScript garbage collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGarbageCollect(_:)
func JSGarbageCollect(p0 JSContextRef) unsafe.Pointer {
	return _JSGarbageCollect(p0)
}



