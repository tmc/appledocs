// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// ObjectiveC Functions (160 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NXCompareHashTables func(unsafe.Pointer, unsafe.Pointer) bool
	_NXCopyHashTable func(unsafe.Pointer) unsafe.Pointer
	_NXCountHashTable func(unsafe.Pointer) unsafe.Pointer
	_NXCreateHashTable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXCreateHashTableFromZone func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXEmptyHashTable func(unsafe.Pointer)
	_NXFreeHashTable func(unsafe.Pointer)
	_NXHashGet func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXHashInsert func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXHashInsertIfAbsent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXHashMember func(unsafe.Pointer, unsafe.Pointer) int
	_NXHashRemove func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXInitHashState func(unsafe.Pointer) unsafe.Pointer
	_NXNextHashState func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_NXNoEffectFree func(unsafe.Pointer, unsafe.Pointer)
	_NXPtrHash func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXPtrIsEqual func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_NXReallyFree func(unsafe.Pointer, unsafe.Pointer)
	_NXResetHashTable func(unsafe.Pointer)
	_NXStrHash func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXStrIsEqual func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_class_addIvar func(unsafe.Pointer, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) bool
	_class_addMethod func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_class_addProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_class_addProtocol func(unsafe.Pointer, unsafe.Pointer) bool
	_class_conformsToProtocol func(unsafe.Pointer, unsafe.Pointer) bool
	_class_copyIvarList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_class_copyMethodList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_class_copyPropertyList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_class_copyProtocolList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_class_createInstance func(unsafe.Pointer, uintptr) unsafe.Pointer
	_class_createInstanceFromZone func(unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_class_getClassMethod func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_class_getClassVariable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_class_getImageName func(unsafe.Pointer) unsafe.Pointer
	_class_getInstanceMethod func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_class_getInstanceSize func(unsafe.Pointer) uintptr
	_class_getInstanceVariable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_class_getIvarLayout func(unsafe.Pointer) unsafe.Pointer
	_class_getMethodImplementation func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_class_getMethodImplementation_stret func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_class_getName func(unsafe.Pointer) unsafe.Pointer
	_class_getProperty func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_class_getSuperclass func(unsafe.Pointer) unsafe.Pointer
	_class_getVersion func(unsafe.Pointer) int
	_class_getWeakIvarLayout func(unsafe.Pointer) unsafe.Pointer
	_class_isMetaClass func(unsafe.Pointer) bool
	_class_lookupMethod func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_class_replaceMethod func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_class_replaceProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_class_respondsToMethod func(unsafe.Pointer, unsafe.Pointer) bool
	_class_respondsToSelector func(unsafe.Pointer, unsafe.Pointer) bool
	_class_setIvarLayout func(unsafe.Pointer, unsafe.Pointer)
	_class_setSuperclass func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_class_setVersion func(unsafe.Pointer, int)
	_class_setWeakIvarLayout func(unsafe.Pointer, unsafe.Pointer)
	_imp_getBlock func(unsafe.Pointer) unsafe.Pointer
	_imp_implementationWithBlock func(unsafe.Pointer) unsafe.Pointer
	_imp_removeBlock func(unsafe.Pointer) bool
	_ivar_getName func(unsafe.Pointer) unsafe.Pointer
	_ivar_getOffset func(unsafe.Pointer) unsafe.Pointer
	_ivar_getTypeEncoding func(unsafe.Pointer) unsafe.Pointer
	_method_copyArgumentType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_method_copyReturnType func(unsafe.Pointer) unsafe.Pointer
	_method_exchangeImplementations func(unsafe.Pointer, unsafe.Pointer)
	_method_getArgumentType func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr)
	_method_getDescription func(unsafe.Pointer) unsafe.Pointer
	_method_getImplementation func(unsafe.Pointer) unsafe.Pointer
	_method_getName func(unsafe.Pointer) unsafe.Pointer
	_method_getNumberOfArguments func(unsafe.Pointer) unsafe.Pointer
	_method_getReturnType func(unsafe.Pointer, unsafe.Pointer, uintptr)
	_method_getTypeEncoding func(unsafe.Pointer) unsafe.Pointer
	_method_invoke func()
	_method_invoke_stret func()
	_method_setImplementation func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_objc_addExceptionHandler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_objc_addLoadImageFunc func(unsafe.Pointer)
	_objc_allocateClassPair func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	_objc_allocateProtocol func(unsafe.Pointer) unsafe.Pointer
	_objc_begin_catch func(unsafe.Pointer) unsafe.Pointer
	_objc_constructInstance func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_objc_copyClassList func(unsafe.Pointer) unsafe.Pointer
	_objc_copyClassNamesForImage func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_objc_copyImageNames func(unsafe.Pointer) unsafe.Pointer
	_objc_copyProtocolList func(unsafe.Pointer) unsafe.Pointer
	_objc_destructInstance func(unsafe.Pointer) unsafe.Pointer
	_objc_disposeClassPair func(unsafe.Pointer)
	_objc_duplicateClass func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	_objc_end_catch func()
	_objc_enumerateClasses func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_objc_enumerationMutation func(unsafe.Pointer)
	_objc_exception_rethrow func()
	_objc_exception_throw func(unsafe.Pointer)
	_objc_getAssociatedObject func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_objc_getClass func(unsafe.Pointer) unsafe.Pointer
	_objc_getClassList func(unsafe.Pointer, int) int
	_objc_getFutureClass func(unsafe.Pointer) unsafe.Pointer
	_objc_getMetaClass func(unsafe.Pointer) unsafe.Pointer
	_objc_getProtocol func(unsafe.Pointer) unsafe.Pointer
	_objc_getRequiredClass func(unsafe.Pointer) unsafe.Pointer
	_objc_loadWeak func(unsafe.Pointer) unsafe.Pointer
	_objc_lookUpClass func(unsafe.Pointer) unsafe.Pointer
	_objc_msgSend func()
	_objc_msgSendSuper func()
	_objc_msgSendSuper_stret func()
	_objc_msgSend_fp2ret func()
	_objc_msgSend_fpret func()
	_objc_msgSend_stret func()
	_objc_registerClassPair func(unsafe.Pointer)
	_objc_registerProtocol func(unsafe.Pointer)
	_objc_removeAssociatedObjects func(unsafe.Pointer)
	_objc_removeExceptionHandler func(unsafe.Pointer)
	_objc_setAssociatedObject func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_objc_setEnumerationMutationHandler func()
	_objc_setExceptionMatcher func(unsafe.Pointer) unsafe.Pointer
	_objc_setExceptionPreprocessor func(unsafe.Pointer) unsafe.Pointer
	_objc_setForwardHandler func(unsafe.Pointer, unsafe.Pointer)
	_objc_setHook_getClass func(unsafe.Pointer, unsafe.Pointer)
	_objc_setHook_getImageName func(unsafe.Pointer, unsafe.Pointer)
	_objc_setHook_lazyClassNamer func(unsafe.Pointer, unsafe.Pointer)
	_objc_setUncaughtExceptionHandler func(unsafe.Pointer) unsafe.Pointer
	_objc_storeWeak func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_objc_sync_enter func(unsafe.Pointer) int
	_objc_sync_exit func(unsafe.Pointer) int
	_objc_terminate func()
	_object_copy func(unsafe.Pointer, uintptr) unsafe.Pointer
	_object_copyFromZone func(unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_object_dispose func(unsafe.Pointer) unsafe.Pointer
	_object_getClass func(unsafe.Pointer) unsafe.Pointer
	_object_getClassName func(unsafe.Pointer) unsafe.Pointer
	_object_getIndexedIvars func(unsafe.Pointer) unsafe.Pointer
	_object_getInstanceVariable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_object_getIvar func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_object_isClass func(unsafe.Pointer) bool
	_object_setClass func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_object_setInstanceVariable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_object_setInstanceVariableWithStrongDefault func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_object_setIvar func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_object_setIvarWithStrongDefault func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_property_copyAttributeList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_property_copyAttributeValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_property_getAttributes func(unsafe.Pointer) unsafe.Pointer
	_property_getName func(unsafe.Pointer) unsafe.Pointer
	_protocol_addMethodDescription func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool, bool)
	_protocol_addProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool, bool)
	_protocol_addProtocol func(unsafe.Pointer, unsafe.Pointer)
	_protocol_conformsToProtocol func(unsafe.Pointer, unsafe.Pointer) bool
	_protocol_copyMethodDescriptionList func(unsafe.Pointer, bool, bool, unsafe.Pointer) unsafe.Pointer
	_protocol_copyPropertyList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_protocol_copyPropertyList2 func(unsafe.Pointer, unsafe.Pointer, bool, bool) unsafe.Pointer
	_protocol_copyProtocolList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_protocol_getMethodDescription func(unsafe.Pointer, unsafe.Pointer, bool, bool) unsafe.Pointer
	_protocol_getName func(unsafe.Pointer) unsafe.Pointer
	_protocol_getProperty func(unsafe.Pointer, unsafe.Pointer, bool, bool) unsafe.Pointer
	_protocol_isEqual func(unsafe.Pointer, unsafe.Pointer) bool
	_sel_getName func(unsafe.Pointer) unsafe.Pointer
	_sel_getUid func(unsafe.Pointer) unsafe.Pointer
	_sel_isEqual func(unsafe.Pointer, unsafe.Pointer) bool
	_sel_isMapped func(unsafe.Pointer) bool
	_sel_registerName func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_NXCompareHashTables, lib, "NXCompareHashTables")
	tryRegister(&_NXCopyHashTable, lib, "NXCopyHashTable")
	tryRegister(&_NXCountHashTable, lib, "NXCountHashTable")
	tryRegister(&_NXCreateHashTable, lib, "NXCreateHashTable")
	tryRegister(&_NXCreateHashTableFromZone, lib, "NXCreateHashTableFromZone")
	tryRegister(&_NXEmptyHashTable, lib, "NXEmptyHashTable")
	tryRegister(&_NXFreeHashTable, lib, "NXFreeHashTable")
	tryRegister(&_NXHashGet, lib, "NXHashGet")
	tryRegister(&_NXHashInsert, lib, "NXHashInsert")
	tryRegister(&_NXHashInsertIfAbsent, lib, "NXHashInsertIfAbsent")
	tryRegister(&_NXHashMember, lib, "NXHashMember")
	tryRegister(&_NXHashRemove, lib, "NXHashRemove")
	tryRegister(&_NXInitHashState, lib, "NXInitHashState")
	tryRegister(&_NXNextHashState, lib, "NXNextHashState")
	tryRegister(&_NXNoEffectFree, lib, "NXNoEffectFree")
	tryRegister(&_NXPtrHash, lib, "NXPtrHash")
	tryRegister(&_NXPtrIsEqual, lib, "NXPtrIsEqual")
	tryRegister(&_NXReallyFree, lib, "NXReallyFree")
	tryRegister(&_NXResetHashTable, lib, "NXResetHashTable")
	tryRegister(&_NXStrHash, lib, "NXStrHash")
	tryRegister(&_NXStrIsEqual, lib, "NXStrIsEqual")
	tryRegister(&_class_addIvar, lib, "class_addIvar")
	tryRegister(&_class_addMethod, lib, "class_addMethod")
	tryRegister(&_class_addProperty, lib, "class_addProperty")
	tryRegister(&_class_addProtocol, lib, "class_addProtocol")
	tryRegister(&_class_conformsToProtocol, lib, "class_conformsToProtocol")
	tryRegister(&_class_copyIvarList, lib, "class_copyIvarList")
	tryRegister(&_class_copyMethodList, lib, "class_copyMethodList")
	tryRegister(&_class_copyPropertyList, lib, "class_copyPropertyList")
	tryRegister(&_class_copyProtocolList, lib, "class_copyProtocolList")
	tryRegister(&_class_createInstance, lib, "class_createInstance")
	tryRegister(&_class_createInstanceFromZone, lib, "class_createInstanceFromZone")
	tryRegister(&_class_getClassMethod, lib, "class_getClassMethod")
	tryRegister(&_class_getClassVariable, lib, "class_getClassVariable")
	tryRegister(&_class_getImageName, lib, "class_getImageName")
	tryRegister(&_class_getInstanceMethod, lib, "class_getInstanceMethod")
	tryRegister(&_class_getInstanceSize, lib, "class_getInstanceSize")
	tryRegister(&_class_getInstanceVariable, lib, "class_getInstanceVariable")
	tryRegister(&_class_getIvarLayout, lib, "class_getIvarLayout")
	tryRegister(&_class_getMethodImplementation, lib, "class_getMethodImplementation")
	tryRegister(&_class_getMethodImplementation_stret, lib, "class_getMethodImplementation_stret")
	tryRegister(&_class_getName, lib, "class_getName")
	tryRegister(&_class_getProperty, lib, "class_getProperty")
	tryRegister(&_class_getSuperclass, lib, "class_getSuperclass")
	tryRegister(&_class_getVersion, lib, "class_getVersion")
	tryRegister(&_class_getWeakIvarLayout, lib, "class_getWeakIvarLayout")
	tryRegister(&_class_isMetaClass, lib, "class_isMetaClass")
	tryRegister(&_class_lookupMethod, lib, "class_lookupMethod")
	tryRegister(&_class_replaceMethod, lib, "class_replaceMethod")
	tryRegister(&_class_replaceProperty, lib, "class_replaceProperty")
	tryRegister(&_class_respondsToMethod, lib, "class_respondsToMethod")
	tryRegister(&_class_respondsToSelector, lib, "class_respondsToSelector")
	tryRegister(&_class_setIvarLayout, lib, "class_setIvarLayout")
	tryRegister(&_class_setSuperclass, lib, "class_setSuperclass")
	tryRegister(&_class_setVersion, lib, "class_setVersion")
	tryRegister(&_class_setWeakIvarLayout, lib, "class_setWeakIvarLayout")
	tryRegister(&_imp_getBlock, lib, "imp_getBlock")
	tryRegister(&_imp_implementationWithBlock, lib, "imp_implementationWithBlock")
	tryRegister(&_imp_removeBlock, lib, "imp_removeBlock")
	tryRegister(&_ivar_getName, lib, "ivar_getName")
	tryRegister(&_ivar_getOffset, lib, "ivar_getOffset")
	tryRegister(&_ivar_getTypeEncoding, lib, "ivar_getTypeEncoding")
	tryRegister(&_method_copyArgumentType, lib, "method_copyArgumentType")
	tryRegister(&_method_copyReturnType, lib, "method_copyReturnType")
	tryRegister(&_method_exchangeImplementations, lib, "method_exchangeImplementations")
	tryRegister(&_method_getArgumentType, lib, "method_getArgumentType")
	tryRegister(&_method_getDescription, lib, "method_getDescription")
	tryRegister(&_method_getImplementation, lib, "method_getImplementation")
	tryRegister(&_method_getName, lib, "method_getName")
	tryRegister(&_method_getNumberOfArguments, lib, "method_getNumberOfArguments")
	tryRegister(&_method_getReturnType, lib, "method_getReturnType")
	tryRegister(&_method_getTypeEncoding, lib, "method_getTypeEncoding")
	tryRegister(&_method_invoke, lib, "method_invoke")
	tryRegister(&_method_invoke_stret, lib, "method_invoke_stret")
	tryRegister(&_method_setImplementation, lib, "method_setImplementation")
	tryRegister(&_objc_addExceptionHandler, lib, "objc_addExceptionHandler")
	tryRegister(&_objc_addLoadImageFunc, lib, "objc_addLoadImageFunc")
	tryRegister(&_objc_allocateClassPair, lib, "objc_allocateClassPair")
	tryRegister(&_objc_allocateProtocol, lib, "objc_allocateProtocol")
	tryRegister(&_objc_begin_catch, lib, "objc_begin_catch")
	tryRegister(&_objc_constructInstance, lib, "objc_constructInstance")
	tryRegister(&_objc_copyClassList, lib, "objc_copyClassList")
	tryRegister(&_objc_copyClassNamesForImage, lib, "objc_copyClassNamesForImage")
	tryRegister(&_objc_copyImageNames, lib, "objc_copyImageNames")
	tryRegister(&_objc_copyProtocolList, lib, "objc_copyProtocolList")
	tryRegister(&_objc_destructInstance, lib, "objc_destructInstance")
	tryRegister(&_objc_disposeClassPair, lib, "objc_disposeClassPair")
	tryRegister(&_objc_duplicateClass, lib, "objc_duplicateClass")
	tryRegister(&_objc_end_catch, lib, "objc_end_catch")
	tryRegister(&_objc_enumerateClasses, lib, "objc_enumerateClasses")
	tryRegister(&_objc_enumerationMutation, lib, "objc_enumerationMutation")
	tryRegister(&_objc_exception_rethrow, lib, "objc_exception_rethrow")
	tryRegister(&_objc_exception_throw, lib, "objc_exception_throw")
	tryRegister(&_objc_getAssociatedObject, lib, "objc_getAssociatedObject")
	tryRegister(&_objc_getClass, lib, "objc_getClass")
	tryRegister(&_objc_getClassList, lib, "objc_getClassList")
	tryRegister(&_objc_getFutureClass, lib, "objc_getFutureClass")
	tryRegister(&_objc_getMetaClass, lib, "objc_getMetaClass")
	tryRegister(&_objc_getProtocol, lib, "objc_getProtocol")
	tryRegister(&_objc_getRequiredClass, lib, "objc_getRequiredClass")
	tryRegister(&_objc_loadWeak, lib, "objc_loadWeak")
	tryRegister(&_objc_lookUpClass, lib, "objc_lookUpClass")
	tryRegister(&_objc_msgSend, lib, "objc_msgSend")
	tryRegister(&_objc_msgSendSuper, lib, "objc_msgSendSuper")
	tryRegister(&_objc_msgSendSuper_stret, lib, "objc_msgSendSuper_stret")
	tryRegister(&_objc_msgSend_fp2ret, lib, "objc_msgSend_fp2ret")
	tryRegister(&_objc_msgSend_fpret, lib, "objc_msgSend_fpret")
	tryRegister(&_objc_msgSend_stret, lib, "objc_msgSend_stret")
	tryRegister(&_objc_registerClassPair, lib, "objc_registerClassPair")
	tryRegister(&_objc_registerProtocol, lib, "objc_registerProtocol")
	tryRegister(&_objc_removeAssociatedObjects, lib, "objc_removeAssociatedObjects")
	tryRegister(&_objc_removeExceptionHandler, lib, "objc_removeExceptionHandler")
	tryRegister(&_objc_setAssociatedObject, lib, "objc_setAssociatedObject")
	tryRegister(&_objc_setEnumerationMutationHandler, lib, "objc_setEnumerationMutationHandler")
	tryRegister(&_objc_setExceptionMatcher, lib, "objc_setExceptionMatcher")
	tryRegister(&_objc_setExceptionPreprocessor, lib, "objc_setExceptionPreprocessor")
	tryRegister(&_objc_setForwardHandler, lib, "objc_setForwardHandler")
	tryRegister(&_objc_setHook_getClass, lib, "objc_setHook_getClass")
	tryRegister(&_objc_setHook_getImageName, lib, "objc_setHook_getImageName")
	tryRegister(&_objc_setHook_lazyClassNamer, lib, "objc_setHook_lazyClassNamer")
	tryRegister(&_objc_setUncaughtExceptionHandler, lib, "objc_setUncaughtExceptionHandler")
	tryRegister(&_objc_storeWeak, lib, "objc_storeWeak")
	tryRegister(&_objc_sync_enter, lib, "objc_sync_enter")
	tryRegister(&_objc_sync_exit, lib, "objc_sync_exit")
	tryRegister(&_objc_terminate, lib, "objc_terminate")
	tryRegister(&_object_copy, lib, "object_copy")
	tryRegister(&_object_copyFromZone, lib, "object_copyFromZone")
	tryRegister(&_object_dispose, lib, "object_dispose")
	tryRegister(&_object_getClass, lib, "object_getClass")
	tryRegister(&_object_getClassName, lib, "object_getClassName")
	tryRegister(&_object_getIndexedIvars, lib, "object_getIndexedIvars")
	tryRegister(&_object_getInstanceVariable, lib, "object_getInstanceVariable")
	tryRegister(&_object_getIvar, lib, "object_getIvar")
	tryRegister(&_object_isClass, lib, "object_isClass")
	tryRegister(&_object_setClass, lib, "object_setClass")
	tryRegister(&_object_setInstanceVariable, lib, "object_setInstanceVariable")
	tryRegister(&_object_setInstanceVariableWithStrongDefault, lib, "object_setInstanceVariableWithStrongDefault")
	tryRegister(&_object_setIvar, lib, "object_setIvar")
	tryRegister(&_object_setIvarWithStrongDefault, lib, "object_setIvarWithStrongDefault")
	tryRegister(&_property_copyAttributeList, lib, "property_copyAttributeList")
	tryRegister(&_property_copyAttributeValue, lib, "property_copyAttributeValue")
	tryRegister(&_property_getAttributes, lib, "property_getAttributes")
	tryRegister(&_property_getName, lib, "property_getName")
	tryRegister(&_protocol_addMethodDescription, lib, "protocol_addMethodDescription")
	tryRegister(&_protocol_addProperty, lib, "protocol_addProperty")
	tryRegister(&_protocol_addProtocol, lib, "protocol_addProtocol")
	tryRegister(&_protocol_conformsToProtocol, lib, "protocol_conformsToProtocol")
	tryRegister(&_protocol_copyMethodDescriptionList, lib, "protocol_copyMethodDescriptionList")
	tryRegister(&_protocol_copyPropertyList, lib, "protocol_copyPropertyList")
	tryRegister(&_protocol_copyPropertyList2, lib, "protocol_copyPropertyList2")
	tryRegister(&_protocol_copyProtocolList, lib, "protocol_copyProtocolList")
	tryRegister(&_protocol_getMethodDescription, lib, "protocol_getMethodDescription")
	tryRegister(&_protocol_getName, lib, "protocol_getName")
	tryRegister(&_protocol_getProperty, lib, "protocol_getProperty")
	tryRegister(&_protocol_isEqual, lib, "protocol_isEqual")
	tryRegister(&_sel_getName, lib, "sel_getName")
	tryRegister(&_sel_getUid, lib, "sel_getUid")
	tryRegister(&_sel_isEqual, lib, "sel_isEqual")
	tryRegister(&_sel_isMapped, lib, "sel_isMapped")
	tryRegister(&_sel_registerName, lib, "sel_registerName")
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



// NXCompareHashTables is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXCompareHashTables
func NXCompareHashTables(table1 unsafe.Pointer, table2 unsafe.Pointer) bool {
	return _NXCompareHashTables(table1, table2)
	}


// NXCopyHashTable is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXCopyHashTable
func NXCopyHashTable(table unsafe.Pointer) unsafe.Pointer {
	return _NXCopyHashTable(table)
	}


// NXCountHashTable is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXCountHashTable
func NXCountHashTable(table unsafe.Pointer) unsafe.Pointer {
	return _NXCountHashTable(table)
	}


// NXCreateHashTable is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXCreateHashTable
func NXCreateHashTable(prototype unsafe.Pointer, capacity unsafe.Pointer, info unsafe.Pointer) unsafe.Pointer {
	return _NXCreateHashTable(prototype, capacity, info)
	}


// NXCreateHashTableFromZone is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXCreateHashTableFromZone
func NXCreateHashTableFromZone(prototype unsafe.Pointer, capacity unsafe.Pointer, info unsafe.Pointer, zone unsafe.Pointer) unsafe.Pointer {
	return _NXCreateHashTableFromZone(prototype, capacity, info, zone)
	}


// NXEmptyHashTable is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXEmptyHashTable
func NXEmptyHashTable(table unsafe.Pointer) {
	_NXEmptyHashTable(table)
	}


// NXFreeHashTable is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXFreeHashTable
func NXFreeHashTable(table unsafe.Pointer) {
	_NXFreeHashTable(table)
	}


// NXHashGet is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashGet
func NXHashGet(table unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _NXHashGet(table, data)
	}


// NXHashInsert is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashInsert
func NXHashInsert(table unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _NXHashInsert(table, data)
	}


// NXHashInsertIfAbsent is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashInsertIfAbsent
func NXHashInsertIfAbsent(table unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _NXHashInsertIfAbsent(table, data)
	}


// NXHashMember is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashMember
func NXHashMember(table unsafe.Pointer, data unsafe.Pointer) int {
	return _NXHashMember(table, data)
	}


// NXHashRemove is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashRemove
func NXHashRemove(table unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _NXHashRemove(table, data)
	}


// NXInitHashState is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXInitHashState
func NXInitHashState(table unsafe.Pointer) unsafe.Pointer {
	return _NXInitHashState(table)
	}


// NXNextHashState is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXNextHashState
func NXNextHashState(table unsafe.Pointer, state unsafe.Pointer, data unsafe.Pointer) int {
	return _NXNextHashState(table, state, data)
	}


// NXNoEffectFree is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXNoEffectFree
func NXNoEffectFree(info unsafe.Pointer, data unsafe.Pointer) {
	_NXNoEffectFree(info, data)
	}


// NXPtrHash is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXPtrHash
func NXPtrHash(info unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _NXPtrHash(info, data)
	}


// NXPtrIsEqual is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXPtrIsEqual
func NXPtrIsEqual(info unsafe.Pointer, data1 unsafe.Pointer, data2 unsafe.Pointer) int {
	return _NXPtrIsEqual(info, data1, data2)
	}


// NXReallyFree is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXReallyFree
func NXReallyFree(info unsafe.Pointer, data unsafe.Pointer) {
	_NXReallyFree(info, data)
	}


// NXResetHashTable is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXResetHashTable
func NXResetHashTable(table unsafe.Pointer) {
	_NXResetHashTable(table)
	}


// NXStrHash is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXStrHash
func NXStrHash(info unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _NXStrHash(info, data)
	}


// NXStrIsEqual is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXStrIsEqual
func NXStrIsEqual(info unsafe.Pointer, data1 unsafe.Pointer, data2 unsafe.Pointer) int {
	return _NXStrIsEqual(info, data1, data2)
	}


// Adds a new instance variable to a class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_addIvar(_:_:_:_:_:)
func class_addIvar(cls unsafe.Pointer, name unsafe.Pointer, size uintptr, alignment unsafe.Pointer, types unsafe.Pointer) bool {
	return _class_addIvar(cls, name, size, alignment, types)
	}


// Adds a new method to a class with a given name and implementation. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_addMethod(_:_:_:_:)
func class_addMethod(cls unsafe.Pointer, name unsafe.Pointer, imp unsafe.Pointer, types unsafe.Pointer) bool {
	return _class_addMethod(cls, name, imp, types)
	}


// Adds a property to a class. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_addProperty(_:_:_:_:)
func class_addProperty(cls unsafe.Pointer, name unsafe.Pointer, attributes unsafe.Pointer, attributeCount unsafe.Pointer) bool {
	return _class_addProperty(cls, name, attributes, attributeCount)
	}


// Adds a protocol to a class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_addProtocol(_:_:)
func class_addProtocol(cls unsafe.Pointer, protocol unsafe.Pointer) bool {
	return _class_addProtocol(cls, protocol)
	}


// Returns a Boolean value that indicates whether a class conforms to a given protocol. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_conformsToProtocol(_:_:)
func class_conformsToProtocol(cls unsafe.Pointer, protocol unsafe.Pointer) bool {
	return _class_conformsToProtocol(cls, protocol)
	}


// Describes the instance variables declared by a class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_copyIvarList(_:_:)
func class_copyIvarList(cls unsafe.Pointer, outCount unsafe.Pointer) unsafe.Pointer {
	return _class_copyIvarList(cls, outCount)
	}


// Describes the instance methods implemented by a class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_copyMethodList(_:_:)
func class_copyMethodList(cls unsafe.Pointer, outCount unsafe.Pointer) unsafe.Pointer {
	return _class_copyMethodList(cls, outCount)
	}


// Describes the properties declared by a class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_copyPropertyList(_:_:)
func class_copyPropertyList(cls unsafe.Pointer, outCount unsafe.Pointer) unsafe.Pointer {
	return _class_copyPropertyList(cls, outCount)
	}


// Describes the protocols adopted by a class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_copyProtocolList(_:_:)
func class_copyProtocolList(cls unsafe.Pointer, outCount unsafe.Pointer) unsafe.Pointer {
	return _class_copyProtocolList(cls, outCount)
	}


// Creates an instance of a class, allocating memory for the class in the default malloc memory zone. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_createInstance(_:_:)
func class_createInstance(cls unsafe.Pointer, extraBytes uintptr) unsafe.Pointer {
	return _class_createInstance(cls, extraBytes)
	}


// class_createInstanceFromZone is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_createInstanceFromZone
func class_createInstanceFromZone(p0 unsafe.Pointer, idxIvars uintptr, zone unsafe.Pointer) unsafe.Pointer {
	return _class_createInstanceFromZone(p0, idxIvars, zone)
	}


// Returns a pointer to the data structure describing a given class method for a given class. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getClassMethod(_:_:)
func class_getClassMethod(cls unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _class_getClassMethod(cls, name)
	}


// Returns the for a specified class variable of a given class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getClassVariable(_:_:)
func class_getClassVariable(cls unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _class_getClassVariable(cls, name)
	}


// Returns the name of the dynamic library a class originated from. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getImageName(_:)
func class_getImageName(cls unsafe.Pointer) unsafe.Pointer {
	return _class_getImageName(cls)
	}


// Returns a specified instance method for a given class. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getInstanceMethod(_:_:)
func class_getInstanceMethod(cls unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _class_getInstanceMethod(cls, name)
	}


// Returns the size of instances of a class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getInstanceSize(_:)
func class_getInstanceSize(cls unsafe.Pointer) uintptr {
	return _class_getInstanceSize(cls)
	}


// Returns the for a specified instance variable of a given class. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getInstanceVariable(_:_:)
func class_getInstanceVariable(cls unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _class_getInstanceVariable(cls, name)
	}


// Returns a description of the layout for a given class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getIvarLayout(_:)
func class_getIvarLayout(cls unsafe.Pointer) unsafe.Pointer {
	return _class_getIvarLayout(cls)
	}


// Returns the function pointer that would be called if a particular message were sent to an instance of a class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getMethodImplementation(_:_:)
func class_getMethodImplementation(cls unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _class_getMethodImplementation(cls, name)
	}


// Returns the function pointer that would be called if a particular message were sent to an instance of a class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getMethodImplementation_stret(_:_:)
func class_getMethodImplementation_stret(cls unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _class_getMethodImplementation_stret(cls, name)
	}


// Returns the name of a class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getName(_:)
func class_getName(cls unsafe.Pointer) unsafe.Pointer {
	return _class_getName(cls)
	}


// Returns a property with a given name of a given class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getProperty(_:_:)
func class_getProperty(cls unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _class_getProperty(cls, name)
	}


// Returns the superclass of a class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getSuperclass(_:)
func class_getSuperclass(cls unsafe.Pointer) unsafe.Pointer {
	return _class_getSuperclass(cls)
	}


// Returns the version number of a class definition. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getVersion(_:)
func class_getVersion(cls unsafe.Pointer) int {
	return _class_getVersion(cls)
	}


// Returns a description of the layout of weak s for a given class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getWeakIvarLayout(_:)
func class_getWeakIvarLayout(cls unsafe.Pointer) unsafe.Pointer {
	return _class_getWeakIvarLayout(cls)
	}


// Returns a Boolean value that indicates whether a class object is a metaclass. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_isMetaClass(_:)
func class_isMetaClass(cls unsafe.Pointer) bool {
	return _class_isMetaClass(cls)
	}


// class_lookupMethod is a ObjectiveC function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_lookupMethod(_:_:)
func class_lookupMethod(cls unsafe.Pointer, sel unsafe.Pointer) unsafe.Pointer {
	return _class_lookupMethod(cls, sel)
	}


// Replaces the implementation of a method for a given class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_replaceMethod(_:_:_:_:)
func class_replaceMethod(cls unsafe.Pointer, name unsafe.Pointer, imp unsafe.Pointer, types unsafe.Pointer) unsafe.Pointer {
	return _class_replaceMethod(cls, name, imp, types)
	}


// Replace a property of a class. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_replaceProperty(_:_:_:_:)
func class_replaceProperty(cls unsafe.Pointer, name unsafe.Pointer, attributes unsafe.Pointer, attributeCount unsafe.Pointer) {
	_class_replaceProperty(cls, name, attributes, attributeCount)
	}


// class_respondsToMethod is a ObjectiveC function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_respondsToMethod(_:_:)
func class_respondsToMethod(cls unsafe.Pointer, sel unsafe.Pointer) bool {
	return _class_respondsToMethod(cls, sel)
	}


// Returns a Boolean value that indicates whether instances of a class respond to a particular selector. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_respondsToSelector(_:_:)
func class_respondsToSelector(cls unsafe.Pointer, sel unsafe.Pointer) bool {
	return _class_respondsToSelector(cls, sel)
	}


// Sets the layout for a given class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_setIvarLayout(_:_:)
func class_setIvarLayout(cls unsafe.Pointer, layout unsafe.Pointer) {
	_class_setIvarLayout(cls, layout)
	}


// Sets the superclass of a given class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_setSuperclass(_:_:)
func class_setSuperclass(cls unsafe.Pointer, newSuper unsafe.Pointer) unsafe.Pointer {
	return _class_setSuperclass(cls, newSuper)
	}


// Sets the version number of a class definition. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_setVersion(_:_:)
func class_setVersion(cls unsafe.Pointer, version int) {
	_class_setVersion(cls, version)
	}


// Sets the layout for weak s for a given class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_setWeakIvarLayout(_:_:)
func class_setWeakIvarLayout(cls unsafe.Pointer, layout unsafe.Pointer) {
	_class_setWeakIvarLayout(cls, layout)
	}


// Returns the block associated with an that was created using . [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/imp_getBlock(_:)
func imp_getBlock(anImp unsafe.Pointer) unsafe.Pointer {
	return _imp_getBlock(anImp)
	}


// Creates a pointer to a function that calls the specified block when the method is called. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/imp_implementationWithBlock(_:)
func imp_implementationWithBlock(block unsafe.Pointer) unsafe.Pointer {
	return _imp_implementationWithBlock(block)
	}


// Disassociates a block from an that was created using , and releases the copy of the block that was created. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/imp_removeBlock(_:)
func imp_removeBlock(anImp unsafe.Pointer) bool {
	return _imp_removeBlock(anImp)
	}


// Returns the name of an instance variable. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/ivar_getName(_:)
func ivar_getName(v unsafe.Pointer) unsafe.Pointer {
	return _ivar_getName(v)
	}


// Returns the offset of an instance variable. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/ivar_getOffset(_:)
func ivar_getOffset(v unsafe.Pointer) unsafe.Pointer {
	return _ivar_getOffset(v)
	}


// Returns the type string of an instance variable. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/ivar_getTypeEncoding(_:)
func ivar_getTypeEncoding(v unsafe.Pointer) unsafe.Pointer {
	return _ivar_getTypeEncoding(v)
	}


// Returns a string describing a single parameter type of a method. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_copyArgumentType(_:_:)
func method_copyArgumentType(m unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	return _method_copyArgumentType(m, index)
	}


// Returns a string describing a method’s return type. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_copyReturnType(_:)
func method_copyReturnType(m unsafe.Pointer) unsafe.Pointer {
	return _method_copyReturnType(m)
	}


// Exchanges the implementations of two methods. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_exchangeImplementations(_:_:)
func method_exchangeImplementations(m1 unsafe.Pointer, m2 unsafe.Pointer) {
	_method_exchangeImplementations(m1, m2)
	}


// Returns by reference a string describing a single parameter type of a method. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_getArgumentType(_:_:_:_:)
func method_getArgumentType(m unsafe.Pointer, index unsafe.Pointer, dst unsafe.Pointer, dst_len uintptr) {
	_method_getArgumentType(m, index, dst, dst_len)
	}


// Returns a method description structure for a specified method. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_getDescription(_:)
func method_getDescription(m unsafe.Pointer) unsafe.Pointer {
	return _method_getDescription(m)
	}


// Returns the implementation of a method. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_getImplementation(_:)
func method_getImplementation(m unsafe.Pointer) unsafe.Pointer {
	return _method_getImplementation(m)
	}


// Returns the name of a method. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_getName(_:)
func method_getName(m unsafe.Pointer) unsafe.Pointer {
	return _method_getName(m)
	}


// Returns the number of arguments accepted by a method. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_getNumberOfArguments(_:)
func method_getNumberOfArguments(m unsafe.Pointer) unsafe.Pointer {
	return _method_getNumberOfArguments(m)
	}


// Returns by reference a string describing a method’s return type. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_getReturnType(_:_:_:)
func method_getReturnType(m unsafe.Pointer, dst unsafe.Pointer, dst_len uintptr) {
	_method_getReturnType(m, dst, dst_len)
	}


// Returns a string describing a method’s parameter and return types. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_getTypeEncoding(_:)
func method_getTypeEncoding(m unsafe.Pointer) unsafe.Pointer {
	return _method_getTypeEncoding(m)
	}


// Calls the implementation of a specified method. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_invoke
func method_invoke() {
	_method_invoke()
	}


// Calls the implementation of a specified method that returns a data-structure. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_invoke_stret
func method_invoke_stret() {
	_method_invoke_stret()
	}


// Sets the implementation of a method. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_setImplementation(_:_:)
func method_setImplementation(m unsafe.Pointer, imp unsafe.Pointer) unsafe.Pointer {
	return _method_setImplementation(m, imp)
	}


// objc_addExceptionHandler is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_addExceptionHandler(_:_:)
func objc_addExceptionHandler(fn unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _objc_addExceptionHandler(fn, context)
	}


// objc_addLoadImageFunc is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_addLoadImageFunc(_:)
func objc_addLoadImageFunc(func_ unsafe.Pointer) {
	_objc_addLoadImageFunc(func_)
	}


// Creates a new class and metaclass. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_allocateClassPair(_:_:_:)
func objc_allocateClassPair(superclass unsafe.Pointer, name unsafe.Pointer, extraBytes uintptr) unsafe.Pointer {
	return _objc_allocateClassPair(superclass, name, extraBytes)
	}


// Creates a new protocol instance. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_allocateProtocol(_:)
func objc_allocateProtocol(name unsafe.Pointer) unsafe.Pointer {
	return _objc_allocateProtocol(name)
	}


// objc_begin_catch is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_begin_catch(_:)
func objc_begin_catch(exc_buf unsafe.Pointer) unsafe.Pointer {
	return _objc_begin_catch(exc_buf)
	}


// Creates an instance of a class at the specified location. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_constructInstance
func objc_constructInstance(cls unsafe.Pointer, bytes unsafe.Pointer) unsafe.Pointer {
	return _objc_constructInstance(cls, bytes)
	}


// Creates and returns a list of pointers to all registered class definitions. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_copyClassList(_:)
func objc_copyClassList(outCount unsafe.Pointer) unsafe.Pointer {
	return _objc_copyClassList(outCount)
	}


// Returns the names of all the classes within a specified library or framework. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_copyClassNamesForImage(_:_:)
func objc_copyClassNamesForImage(image unsafe.Pointer, outCount unsafe.Pointer) unsafe.Pointer {
	return _objc_copyClassNamesForImage(image, outCount)
	}


// Returns the names of all the loaded Objective-C frameworks and dynamic libraries. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_copyImageNames(_:)
func objc_copyImageNames(outCount unsafe.Pointer) unsafe.Pointer {
	return _objc_copyImageNames(outCount)
	}


// Returns an array of all the protocols known to the runtime. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_copyProtocolList(_:)
func objc_copyProtocolList(outCount unsafe.Pointer) unsafe.Pointer {
	return _objc_copyProtocolList(outCount)
	}


// Destroys an instance of a class without freeing memory and removes any of its associated references. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_destructInstance
func objc_destructInstance(obj unsafe.Pointer) unsafe.Pointer {
	return _objc_destructInstance(obj)
	}


// Destroys a class and its associated metaclass. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_disposeClassPair(_:)
func objc_disposeClassPair(cls unsafe.Pointer) {
	_objc_disposeClassPair(cls)
	}


// Used by Foundation’s Key-Value Observing. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_duplicateClass(_:_:_:)
func objc_duplicateClass(original unsafe.Pointer, name unsafe.Pointer, extraBytes uintptr) unsafe.Pointer {
	return _objc_duplicateClass(original, name, extraBytes)
	}


// objc_end_catch is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_end_catch()
func objc_end_catch() {
	_objc_end_catch()
	}


// objc_enumerateClasses is a ObjectiveC function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_enumerateClasses
func objc_enumerateClasses(image unsafe.Pointer, namePrefix unsafe.Pointer, conformingTo unsafe.Pointer, subclassing unsafe.Pointer) {
	_objc_enumerateClasses(image, namePrefix, conformingTo, subclassing)
	}


// Inserted by the compiler when a mutation is detected during a foreach iteration. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_enumerationMutation(_:)
func objc_enumerationMutation(obj unsafe.Pointer) {
	_objc_enumerationMutation(obj)
	}


// objc_exception_rethrow is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_exception_rethrow()
func objc_exception_rethrow() {
	_objc_exception_rethrow()
	}


// Throw a runtime exception. This function is inserted by the compiler where \c @throw would otherwise be. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_exception_throw(_:)
func objc_exception_throw(exception unsafe.Pointer) {
	_objc_exception_throw(exception)
	}


// Returns the value associated with a given object for a given key. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_getAssociatedObject(_:_:)
func objc_getAssociatedObject(object unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _objc_getAssociatedObject(object, key)
	}


// Returns the class definition of a specified class. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_getClass(_:)
func objc_getClass(name unsafe.Pointer) unsafe.Pointer {
	return _objc_getClass(name)
	}


// Obtains the list of registered class definitions. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_getClassList(_:_:)
func objc_getClassList(buffer unsafe.Pointer, bufferCount int) int {
	return _objc_getClassList(buffer, bufferCount)
	}


// Used by CoreFoundation’s toll-free bridging. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_getFutureClass
func objc_getFutureClass(name unsafe.Pointer) unsafe.Pointer {
	return _objc_getFutureClass(name)
	}


// Returns the metaclass definition of a specified class. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_getMetaClass(_:)
func objc_getMetaClass(name unsafe.Pointer) unsafe.Pointer {
	return _objc_getMetaClass(name)
	}


// Returns a specified protocol. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_getProtocol(_:)
func objc_getProtocol(name unsafe.Pointer) unsafe.Pointer {
	return _objc_getProtocol(name)
	}


// Returns the class definition of a specified class. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_getRequiredClass(_:)
func objc_getRequiredClass(name unsafe.Pointer) unsafe.Pointer {
	return _objc_getRequiredClass(name)
	}


// Loads the object referenced by a weak pointer and returns it. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_loadWeak(_:)
func objc_loadWeak(location unsafe.Pointer) unsafe.Pointer {
	return _objc_loadWeak(location)
	}


// Returns the class definition of a specified class. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_lookUpClass(_:)
func objc_lookUpClass(name unsafe.Pointer) unsafe.Pointer {
	return _objc_lookUpClass(name)
	}


// Sends a message with a simple return value to an instance of a class. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_msgSend
func objc_msgSend() {
	_objc_msgSend()
	}


// Sends a message with a simple return value to the superclass of an instance of a class. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_msgSendSuper
func objc_msgSendSuper() {
	_objc_msgSendSuper()
	}


// Sends a message with a data-structure return value to the superclass of an instance of a class. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_msgSendSuper_stret
func objc_msgSendSuper_stret() {
	_objc_msgSendSuper_stret()
	}


// objc_msgSend_fp2ret is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_msgSend_fp2ret
func objc_msgSend_fp2ret() {
	_objc_msgSend_fp2ret()
	}


// Sends a message with a floating-point return value to an instance of a class. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_msgSend_fpret
func objc_msgSend_fpret() {
	_objc_msgSend_fpret()
	}


// Sends a message with a data-structure return value to an instance of a class. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_msgSend_stret
func objc_msgSend_stret() {
	_objc_msgSend_stret()
	}


// Registers a class that was allocated using . [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_registerClassPair(_:)
func objc_registerClassPair(cls unsafe.Pointer) {
	_objc_registerClassPair(cls)
	}


// Registers a newly created protocol with the Objective-C runtime. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_registerProtocol(_:)
func objc_registerProtocol(proto unsafe.Pointer) {
	_objc_registerProtocol(proto)
	}


// Removes all associations for a given object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_removeAssociatedObjects(_:)
func objc_removeAssociatedObjects(object unsafe.Pointer) {
	_objc_removeAssociatedObjects(object)
	}


// objc_removeExceptionHandler is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_removeExceptionHandler(_:)
func objc_removeExceptionHandler(token unsafe.Pointer) {
	_objc_removeExceptionHandler(token)
	}


// Sets an associated value for a given object using a given key and association policy. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setAssociatedObject(_:_:_:_:)
func objc_setAssociatedObject(object unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer, policy unsafe.Pointer) {
	_objc_setAssociatedObject(object, key, value, policy)
	}


// Sets the current mutation handler. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setEnumerationMutationHandler(_:)
func objc_setEnumerationMutationHandler() {
	_objc_setEnumerationMutationHandler()
	}


// objc_setExceptionMatcher is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setExceptionMatcher(_:)
func objc_setExceptionMatcher(fn unsafe.Pointer) unsafe.Pointer {
	return _objc_setExceptionMatcher(fn)
	}


// objc_setExceptionPreprocessor is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setExceptionPreprocessor(_:)
func objc_setExceptionPreprocessor(fn unsafe.Pointer) unsafe.Pointer {
	return _objc_setExceptionPreprocessor(fn)
	}


// Set the function to be called by objc_msgForward. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setForwardHandler(_:_:)
func objc_setForwardHandler(fwd unsafe.Pointer, fwd_stret unsafe.Pointer) {
	_objc_setForwardHandler(fwd, fwd_stret)
	}


// objc_setHook_getClass is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.14.4.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setHook_getClass(_:_:)
func objc_setHook_getClass(newValue unsafe.Pointer, outOldValue unsafe.Pointer) {
	_objc_setHook_getClass(newValue, outOldValue)
	}


// objc_setHook_getImageName is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setHook_getImageName(_:_:)
func objc_setHook_getImageName(newValue unsafe.Pointer, outOldValue unsafe.Pointer) {
	_objc_setHook_getImageName(newValue, outOldValue)
	}


// objc_setHook_lazyClassNamer is a ObjectiveC function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setHook_lazyClassNamer(_:_:)
func objc_setHook_lazyClassNamer(newValue unsafe.Pointer, oldOutValue unsafe.Pointer) {
	_objc_setHook_lazyClassNamer(newValue, oldOutValue)
	}


// objc_setUncaughtExceptionHandler is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setUncaughtExceptionHandler(_:)
func objc_setUncaughtExceptionHandler(fn unsafe.Pointer) unsafe.Pointer {
	return _objc_setUncaughtExceptionHandler(fn)
	}


// Stores a new value in a variable. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_storeWeak(_:_:)
func objc_storeWeak(location unsafe.Pointer, obj unsafe.Pointer) unsafe.Pointer {
	return _objc_storeWeak(location, obj)
	}


// Begin synchronizing on ‘obj’. Allocates recursive pthread_mutex associated with ‘obj’ if needed. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_sync_enter
func objc_sync_enter(obj unsafe.Pointer) int {
	return _objc_sync_enter(obj)
	}


// End synchronizing on ‘obj’. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_sync_exit
func objc_sync_exit(obj unsafe.Pointer) int {
	return _objc_sync_exit(obj)
	}


// objc_terminate is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_terminate()
func objc_terminate() {
	_objc_terminate()
	}


// Returns a copy of a given object. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_copy
func object_copy(obj unsafe.Pointer, size uintptr) unsafe.Pointer {
	return _object_copy(obj, size)
	}


// object_copyFromZone is a ObjectiveC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_copyFromZone
func object_copyFromZone(anObject unsafe.Pointer, nBytes uintptr, zone unsafe.Pointer) unsafe.Pointer {
	return _object_copyFromZone(anObject, nBytes, zone)
	}


// Frees the memory occupied by a given object. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_dispose
func object_dispose(obj unsafe.Pointer) unsafe.Pointer {
	return _object_dispose(obj)
	}


// Returns the class of an object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_getClass(_:)
func object_getClass(obj unsafe.Pointer) unsafe.Pointer {
	return _object_getClass(obj)
	}


// Returns the class name of a given object. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_getClassName(_:)
func object_getClassName(obj unsafe.Pointer) unsafe.Pointer {
	return _object_getClassName(obj)
	}


// Returns a pointer to any extra bytes allocated with a instance given object. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_getIndexedIvars(_:)
func object_getIndexedIvars(obj unsafe.Pointer) unsafe.Pointer {
	return _object_getIndexedIvars(obj)
	}


// Obtains the value of an instance variable of a class instance. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_getInstanceVariable
func object_getInstanceVariable(obj unsafe.Pointer, name unsafe.Pointer, outValue unsafe.Pointer) unsafe.Pointer {
	return _object_getInstanceVariable(obj, name, outValue)
	}


// Reads the value of an instance variable in an object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_getIvar(_:_:)
func object_getIvar(obj unsafe.Pointer, ivar unsafe.Pointer) unsafe.Pointer {
	return _object_getIvar(obj, ivar)
	}


// object_isClass is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_isClass(_:)
func object_isClass(obj unsafe.Pointer) bool {
	return _object_isClass(obj)
	}


// Sets the class of an object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_setClass(_:_:)
func object_setClass(obj unsafe.Pointer, cls unsafe.Pointer) unsafe.Pointer {
	return _object_setClass(obj, cls)
	}


// Changes the value of an instance variable of a class instance. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_setInstanceVariable
func object_setInstanceVariable(obj unsafe.Pointer, name unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _object_setInstanceVariable(obj, name, value)
	}


// object_setInstanceVariableWithStrongDefault is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_setInstanceVariableWithStrongDefault
func object_setInstanceVariableWithStrongDefault(obj unsafe.Pointer, name unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _object_setInstanceVariableWithStrongDefault(obj, name, value)
	}


// Sets the value of an instance variable in an object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_setIvar(_:_:_:)
func object_setIvar(obj unsafe.Pointer, ivar unsafe.Pointer, value unsafe.Pointer) {
	_object_setIvar(obj, ivar, value)
	}


// object_setIvarWithStrongDefault is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_setIvarWithStrongDefault(_:_:_:)
func object_setIvarWithStrongDefault(obj unsafe.Pointer, ivar unsafe.Pointer, value unsafe.Pointer) {
	_object_setIvarWithStrongDefault(obj, ivar, value)
	}


// Returns an array of property attributes for a given property. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/property_copyAttributeList(_:_:)
func property_copyAttributeList(property unsafe.Pointer, outCount unsafe.Pointer) unsafe.Pointer {
	return _property_copyAttributeList(property, outCount)
	}


// Returns the value of a property attribute given the attribute name. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/property_copyAttributeValue(_:_:)
func property_copyAttributeValue(property unsafe.Pointer, attributeName unsafe.Pointer) unsafe.Pointer {
	return _property_copyAttributeValue(property, attributeName)
	}


// Returns the attribute string of a property. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/property_getAttributes(_:)
func property_getAttributes(property unsafe.Pointer) unsafe.Pointer {
	return _property_getAttributes(property)
	}


// Returns the name of a property. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/property_getName(_:)
func property_getName(property unsafe.Pointer) unsafe.Pointer {
	return _property_getName(property)
	}


// Adds a method to a protocol. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_addMethodDescription(_:_:_:_:_:)
func protocol_addMethodDescription(proto unsafe.Pointer, name unsafe.Pointer, types unsafe.Pointer, isRequiredMethod bool, isInstanceMethod bool) {
	_protocol_addMethodDescription(proto, name, types, isRequiredMethod, isInstanceMethod)
	}


// Adds a property to a protocol that is under construction. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_addProperty(_:_:_:_:_:_:)
func protocol_addProperty(proto unsafe.Pointer, name unsafe.Pointer, attributes unsafe.Pointer, attributeCount unsafe.Pointer, isRequiredProperty bool, isInstanceProperty bool) {
	_protocol_addProperty(proto, name, attributes, attributeCount, isRequiredProperty, isInstanceProperty)
	}


// Adds a registered protocol to another protocol that is under construction. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_addProtocol(_:_:)
func protocol_addProtocol(proto unsafe.Pointer, addition unsafe.Pointer) {
	_protocol_addProtocol(proto, addition)
	}


// Returns a Boolean value that indicates whether one protocol conforms to another protocol. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_conformsToProtocol(_:_:)
func protocol_conformsToProtocol(proto unsafe.Pointer, other unsafe.Pointer) bool {
	return _protocol_conformsToProtocol(proto, other)
	}


// Returns an array of method descriptions of methods meeting a given specification for a given protocol. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_copyMethodDescriptionList(_:_:_:_:)
func protocol_copyMethodDescriptionList(proto unsafe.Pointer, isRequiredMethod bool, isInstanceMethod bool, outCount unsafe.Pointer) unsafe.Pointer {
	return _protocol_copyMethodDescriptionList(proto, isRequiredMethod, isInstanceMethod, outCount)
	}


// Returns an array of the properties declared by a protocol. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_copyPropertyList(_:_:)
func protocol_copyPropertyList(proto unsafe.Pointer, outCount unsafe.Pointer) unsafe.Pointer {
	return _protocol_copyPropertyList(proto, outCount)
	}


// protocol_copyPropertyList2 is a ObjectiveC function. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_copyPropertyList2(_:_:_:_:)
func protocol_copyPropertyList2(proto unsafe.Pointer, outCount unsafe.Pointer, isRequiredProperty bool, isInstanceProperty bool) unsafe.Pointer {
	return _protocol_copyPropertyList2(proto, outCount, isRequiredProperty, isInstanceProperty)
	}


// Returns an array of the protocols adopted by a protocol. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_copyProtocolList(_:_:)
func protocol_copyProtocolList(proto unsafe.Pointer, outCount unsafe.Pointer) unsafe.Pointer {
	return _protocol_copyProtocolList(proto, outCount)
	}


// Returns a method description structure for a specified method of a given protocol. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_getMethodDescription(_:_:_:_:)
func protocol_getMethodDescription(proto unsafe.Pointer, aSel unsafe.Pointer, isRequiredMethod bool, isInstanceMethod bool) unsafe.Pointer {
	return _protocol_getMethodDescription(proto, aSel, isRequiredMethod, isInstanceMethod)
	}


// Returns the name of a protocol. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_getName(_:)
func protocol_getName(proto unsafe.Pointer) unsafe.Pointer {
	return _protocol_getName(proto)
	}


// Returns the specified property of a given protocol. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_getProperty(_:_:_:_:)
func protocol_getProperty(proto unsafe.Pointer, name unsafe.Pointer, isRequiredProperty bool, isInstanceProperty bool) unsafe.Pointer {
	return _protocol_getProperty(proto, name, isRequiredProperty, isInstanceProperty)
	}


// Returns a Boolean value that indicates whether two protocols are equal. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_isEqual(_:_:)
func protocol_isEqual(proto unsafe.Pointer, other unsafe.Pointer) bool {
	return _protocol_isEqual(proto, other)
	}


// Returns the name of the method specified by a given selector. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/sel_getName(_:)
func sel_getName(sel unsafe.Pointer) unsafe.Pointer {
	return _sel_getName(sel)
	}


// Registers a method name with the Objective-C runtime system. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/sel_getUid(_:)
func sel_getUid(str unsafe.Pointer) unsafe.Pointer {
	return _sel_getUid(str)
	}


// Returns a Boolean value that indicates whether two selectors are equal. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/sel_isEqual(_:_:)
func sel_isEqual(lhs unsafe.Pointer, rhs unsafe.Pointer) bool {
	return _sel_isEqual(lhs, rhs)
	}


// Identifies a selector as being valid or invalid. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/sel_isMapped(_:)
func sel_isMapped(sel unsafe.Pointer) bool {
	return _sel_isMapped(sel)
	}


// Registers a method with the Objective-C runtime system, maps the method name to a selector, and returns the selector value. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/sel_registerName(_:)
func sel_registerName(str unsafe.Pointer) unsafe.Pointer {
	return _sel_registerName(str)
	}




