// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore
import (
"unsafe"
)

// Type aliases and typedefs
// JSChar - A Unicode character.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSChar
type JSChar uint16
// JSClassAttributes - A set of JavaScript class attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSClassAttributes
type JSClassAttributes uint32
// JSClassRef - A JavaScript class.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSClassRef
// JSClassRef has base type: struct OpaqueJSClass *
type JSClassRef uintptr
// JSContextGroupRef - A group that associates JavaScript contexts with one another.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContextGroupRef
// JSContextGroupRef has base type: const struct OpaqueJSContextGroup *
type JSContextGroupRef uintptr
// JSContextRef - A JavaScript execution context.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSContextRef
// JSContextRef has base type: const struct OpaqueJSContext *
type JSContextRef uintptr
// JSGlobalContextRef - A global JavaScript execution context.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextRef
// JSGlobalContextRef has base type: struct OpaqueJSContext *
type JSGlobalContextRef uintptr
// JSObjectCallAsConstructorCallback - The callback type for using an object as a constructor.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectCallAsConstructorCallback
// JSObjectCallAsConstructorCallback is a callback function
// C type: struct OpaqueJSValue *(*)(const struct OpaqueJSContext *, struct OpaqueJSValue *, unsigned long, const struct OpaqueJSValue *const *, const struct OpaqueJSValue **)
type JSObjectCallAsConstructorCallback = func(unsafe.Pointer, unsafe.Pointer, uint, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
// JSObjectCallAsFunctionCallback - The callback type for calling an object as a function.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectCallAsFunctionCallback
// JSObjectCallAsFunctionCallback is a callback function
// C type: const struct OpaqueJSValue *(*)(const struct OpaqueJSContext *, struct OpaqueJSValue *, struct OpaqueJSValue *, unsigned long, const struct OpaqueJSValue *const *, const struct OpaqueJSValue **)
type JSObjectCallAsFunctionCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
// JSObjectConvertToTypeCallback - The callback type for converting an object to a particular JavaScript type.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectConvertToTypeCallback
// JSObjectConvertToTypeCallback is a callback function
// C type: const struct OpaqueJSValue *(*)(const struct OpaqueJSContext *, struct OpaqueJSValue *, JSType, const struct OpaqueJSValue **)
type JSObjectConvertToTypeCallback = func(unsafe.Pointer, unsafe.Pointer, JSType, unsafe.Pointer) unsafe.Pointer
// JSObjectDeletePropertyCallback - The callback type for deleting a property.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectDeletePropertyCallback
// JSObjectDeletePropertyCallback is a callback function
// C type: _Bool (*)(const struct OpaqueJSContext *, struct OpaqueJSValue *, struct OpaqueJSString *, const struct OpaqueJSValue **)
type JSObjectDeletePropertyCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
// JSObjectFinalizeCallback - The callback type for finalizing an object (preparing it for garbage collection).
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectFinalizeCallback
// JSObjectFinalizeCallback is a callback function
// C type: void (*)(struct OpaqueJSValue *)
type JSObjectFinalizeCallback = func(unsafe.Pointer)
// JSObjectGetPropertyCallback - The callback type for getting a property’s value.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPropertyCallback
// JSObjectGetPropertyCallback is a callback function
// C type: const struct OpaqueJSValue *(*)(const struct OpaqueJSContext *, struct OpaqueJSValue *, struct OpaqueJSString *, const struct OpaqueJSValue **)
type JSObjectGetPropertyCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
// JSObjectGetPropertyNamesCallback - The callback type for collecting the names of an object’s properties.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPropertyNamesCallback
// JSObjectGetPropertyNamesCallback is a callback function
// C type: void (*)(const struct OpaqueJSContext *, struct OpaqueJSValue *, struct OpaqueJSPropertyNameAccumulator *)
type JSObjectGetPropertyNamesCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
// JSObjectHasInstanceCallback - The callback type for checking whether an object is an instance of a particular type.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectHasInstanceCallback
// JSObjectHasInstanceCallback is a callback function
// C type: _Bool (*)(const struct OpaqueJSContext *, struct OpaqueJSValue *, const struct OpaqueJSValue *, const struct OpaqueJSValue **)
type JSObjectHasInstanceCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
// JSObjectHasPropertyCallback - The callback type for determining whether an object has a property.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectHasPropertyCallback
// JSObjectHasPropertyCallback is a callback function
// C type: _Bool (*)(const struct OpaqueJSContext *, struct OpaqueJSValue *, struct OpaqueJSString *)
type JSObjectHasPropertyCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
// JSObjectInitializeCallback - The callback type for first creating an object.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectInitializeCallback
// JSObjectInitializeCallback is a callback function
// C type: void (*)(const struct OpaqueJSContext *, struct OpaqueJSValue *)
type JSObjectInitializeCallback = func(unsafe.Pointer, unsafe.Pointer)
// JSObjectRef - A JavaScript object.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectRef
// JSObjectRef has base type: struct OpaqueJSValue *
type JSObjectRef uintptr
// JSObjectSetPropertyCallback - The callback type for setting a property’s value.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetPropertyCallback
// JSObjectSetPropertyCallback is a callback function
// C type: _Bool (*)(const struct OpaqueJSContext *, struct OpaqueJSValue *, struct OpaqueJSString *, const struct OpaqueJSValue *, const struct OpaqueJSValue **)
type JSObjectSetPropertyCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
// JSPropertyAttributes - A set of JavaScript property attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyAttributes
type JSPropertyAttributes uint32
// JSPropertyNameAccumulatorRef - An ordered set of the names of a JavaScript object’s properties.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameAccumulatorRef
// JSPropertyNameAccumulatorRef has base type: struct OpaqueJSPropertyNameAccumulator *
type JSPropertyNameAccumulatorRef uintptr
// JSPropertyNameArrayRef - An array of JavaScript property names.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameArrayRef
// JSPropertyNameArrayRef has base type: struct OpaqueJSPropertyNameArray *
type JSPropertyNameArrayRef uintptr
// JSStringRef - A UTF-16 character buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStringRef
// JSStringRef has base type: struct OpaqueJSString *
type JSStringRef uintptr
// JSTypedArrayBytesDeallocator - A function that deallocates bytes that pass to a typed array constructor.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSTypedArrayBytesDeallocator
// JSTypedArrayBytesDeallocator is a callback function
// C type: void (*)(void *, void *)
type JSTypedArrayBytesDeallocator = func(unsafe.Pointer, unsafe.Pointer)
// JSValueProperty - A type that identifies a property of a JavaScript value.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueProperty
// JSValueProperty has base type: id
type JSValueProperty uintptr
// JSValueRef - A JavaScript value.
//
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSValueRef
// JSValueRef has base type: const struct OpaqueJSValue *
type JSValueRef uintptr

