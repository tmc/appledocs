// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec
import (
	"unsafe"
)

// Type aliases and typedefs
// Category - An opaque type that represents a category.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/Category
// Category has base type: struct objc_category *
type Category uintptr
// IMP - A pointer to the start of a method implementation.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/IMP
// IMP is a callback function
// C type: void (*)(void)
type IMP = func()
// Ivar - An opaque type that represents an instance variable.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/Ivar
// Ivar has base type: struct objc_ivar *
type Ivar uintptr
// Method - An opaque type that represents a method in a class definition.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/Method
// Method has base type: struct objc_method *
type Method uintptr
// Integer - Describes an integer.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSInteger
// NSInteger has base type: long
type Integer uintptr
// objc_exception_handler type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_exception_handler
// objc_exception_handler is a callback function
// C type: void (*)(id, void *)
type objc_exception_handler = func(id, unsafe.Pointer)
// objc_exception_matcher type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_exception_matcher
// objc_exception_matcher is a callback function
// C type: int (*)(Class, id)
type objc_exception_matcher = func(Class, id) int32
// objc_exception_preprocessor type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_exception_preprocessor
// objc_exception_preprocessor is a callback function
// C type: id (*)(id)
type objc_exception_preprocessor = func(id) id
// objc_func_loadImage type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_func_loadImage
// objc_func_loadImage is a callback function
// C type: void (*)(const struct mach_header *)
type objc_func_loadImage = func(unsafe.Pointer)
// objc_hook_getClass type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_hook_getClass
// objc_hook_getClass is a callback function
// C type: _Bool (*)(const char *, Class *)
type objc_hook_getClass = func(string, unsafe.Pointer) _Bool
// objc_hook_getImageName type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_hook_getImageName
// objc_hook_getImageName is a callback function
// C type: _Bool (*)(Class, const char **)
type objc_hook_getImageName = func(Class, unsafe.Pointer) _Bool
// objc_hook_lazyClassNamer type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_hook_lazyClassNamer
// objc_hook_lazyClassNamer is a callback function
// C type: const char *(*)(Class)
type objc_hook_lazyClassNamer = func(Class) string
// objc_objectptr_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_objectptr_t
// objc_objectptr_t has base type: const void *
type objc_objectptr_t uintptr
// objc_property_t - An opaque type that represents an Objective-C declared property.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_property_t
// objc_property_t has base type: struct objc_property *
type objc_property_t uintptr
// objc_uncaught_exception_handler type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_uncaught_exception_handler
// objc_uncaught_exception_handler is a callback function
// C type: void (*)(id)
type objc_uncaught_exception_handler = func(id)
// objc_zone_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_zone_t
// objc_zone_t has base type: struct _malloc_zone_t *
type objc_zone_t uintptr

