// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

/* debug [functions.gen.go]: Generating 160 functions for ObjectiveC */
import (
	"unsafe"

	"github.com/ebitengine/purego"
	objc "github.com/ebitengine/purego/objc"
)


// ObjectiveC Functions (160 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NXCompareHashTables func(unsafe.Pointer, unsafe.Pointer) bool
	_NXCopyHashTable func(unsafe.Pointer) unsafe.Pointer
	_NXCountHashTable func(unsafe.Pointer) unsafe.Pointer
	_NXCreateHashTable func(NXHashTablePrototype, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXCreateHashTableFromZone func(NXHashTablePrototype, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXEmptyHashTable func(unsafe.Pointer)
	_NXFreeHashTable func(unsafe.Pointer)
	_NXHashGet func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXHashInsert func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXHashInsertIfAbsent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXHashMember func(unsafe.Pointer, unsafe.Pointer) int
	_NXHashRemove func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXInitHashState func(unsafe.Pointer) NXHashState
	_NXNextHashState func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_NXNoEffectFree func(unsafe.Pointer, unsafe.Pointer)
	_NXPtrHash func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXPtrIsEqual func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_NXReallyFree func(unsafe.Pointer, unsafe.Pointer)
	_NXResetHashTable func(unsafe.Pointer)
	_NXStrHash func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NXStrIsEqual func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_class_addIvar func(objc.Class, unsafe.Pointer, uintptr, uint8, unsafe.Pointer) bool
	_class_addMethod func(objc.Class, objc.SEL, IMP, unsafe.Pointer) bool
	_class_addProperty func(objc.Class, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_class_addProtocol func(objc.Class, unsafe.Pointer) bool
	_class_conformsToProtocol func(objc.Class, unsafe.Pointer) bool
	_class_copyIvarList func(objc.Class, unsafe.Pointer) unsafe.Pointer
	_class_copyMethodList func(objc.Class, unsafe.Pointer) unsafe.Pointer
	_class_copyPropertyList func(objc.Class, unsafe.Pointer) unsafe.Pointer
	_class_copyProtocolList func(objc.Class, unsafe.Pointer) unsafe.Pointer
	_class_createInstance func(objc.Class, uintptr) objc.ID
	_class_createInstanceFromZone func(objc.Class, uintptr, unsafe.Pointer) objc.ID
	_class_getClassMethod func(objc.Class, objc.SEL) Method
	_class_getClassVariable func(objc.Class, unsafe.Pointer) Ivar
	_class_getImageName func(objc.Class) unsafe.Pointer
	_class_getInstanceMethod func(objc.Class, objc.SEL) Method
	_class_getInstanceSize func(objc.Class) uintptr
	_class_getInstanceVariable func(objc.Class, unsafe.Pointer) Ivar
	_class_getIvarLayout func(objc.Class) unsafe.Pointer
	_class_getMethodImplementation func(objc.Class, objc.SEL) IMP
	_class_getMethodImplementation_stret func(objc.Class, objc.SEL) IMP
	_class_getName func(objc.Class) unsafe.Pointer
	_class_getProperty func(objc.Class, unsafe.Pointer) objc_property_t
	_class_getSuperclass func(objc.Class) objc.Class
	_class_getVersion func(objc.Class) int
	_class_getWeakIvarLayout func(objc.Class) unsafe.Pointer
	_class_isMetaClass func(objc.Class) bool
	_class_lookupMethod func(objc.Class, objc.SEL) IMP
	_class_replaceMethod func(objc.Class, objc.SEL, IMP, unsafe.Pointer) IMP
	_class_replaceProperty func(objc.Class, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_class_respondsToMethod func(objc.Class, objc.SEL) bool
	_class_respondsToSelector func(objc.Class, objc.SEL) bool
	_class_setIvarLayout func(objc.Class, unsafe.Pointer)
	_class_setSuperclass func(objc.Class, objc.Class) objc.Class
	_class_setVersion func(objc.Class, int)
	_class_setWeakIvarLayout func(objc.Class, unsafe.Pointer)
	_imp_getBlock func(IMP) objc.ID
	_imp_implementationWithBlock func(objc.ID) IMP
	_imp_removeBlock func(IMP) bool
	_ivar_getName func(Ivar) unsafe.Pointer
	_ivar_getOffset func(Ivar) unsafe.Pointer
	_ivar_getTypeEncoding func(Ivar) unsafe.Pointer
	_method_copyArgumentType func(Method, unsafe.Pointer) unsafe.Pointer
	_method_copyReturnType func(Method) unsafe.Pointer
	_method_exchangeImplementations func(Method, Method)
	_method_getArgumentType func(Method, unsafe.Pointer, unsafe.Pointer, uintptr)
	_method_getDescription func(Method) unsafe.Pointer
	_method_getImplementation func(Method) IMP
	_method_getName func(Method) objc.SEL
	_method_getNumberOfArguments func(Method) unsafe.Pointer
	_method_getReturnType func(Method, unsafe.Pointer, uintptr)
	_method_getTypeEncoding func(Method) unsafe.Pointer
	_method_invoke func()
	_method_invoke_stret func()
	_method_setImplementation func(Method, IMP) IMP
	_objc_addExceptionHandler func(objc_exception_handler, unsafe.Pointer) unsafe.Pointer
	_objc_addLoadImageFunc func(objc_func_loadImage)
	_objc_allocateClassPair func(objc.Class, unsafe.Pointer, uintptr) objc.Class
	_objc_allocateProtocol func(unsafe.Pointer) unsafe.Pointer
	_objc_begin_catch func(unsafe.Pointer) objc.ID
	_objc_constructInstance func(objc.Class, unsafe.Pointer) objc.ID
	_objc_copyClassList func(unsafe.Pointer) unsafe.Pointer
	_objc_copyClassNamesForImage func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_objc_copyImageNames func(unsafe.Pointer) unsafe.Pointer
	_objc_copyProtocolList func(unsafe.Pointer) unsafe.Pointer
	_objc_destructInstance func(objc.ID) unsafe.Pointer
	_objc_disposeClassPair func(objc.Class)
	_objc_duplicateClass func(objc.Class, unsafe.Pointer, uintptr) objc.Class
	_objc_end_catch func()
	_objc_enumerateClasses func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, objc.Class)
	_objc_enumerationMutation func(objc.ID)
	_objc_exception_rethrow func()
	_objc_exception_throw func(objc.ID)
	_objc_getAssociatedObject func(objc.ID, unsafe.Pointer) objc.ID
	_objc_getClass func(unsafe.Pointer) objc.ID
	_objc_getClassList func(unsafe.Pointer, int) int
	_objc_getFutureClass func(unsafe.Pointer) objc.Class
	_objc_getMetaClass func(unsafe.Pointer) objc.ID
	_objc_getProtocol func(unsafe.Pointer) unsafe.Pointer
	_objc_getRequiredClass func(unsafe.Pointer) objc.Class
	_objc_loadWeak func(unsafe.Pointer) objc.ID
	_objc_lookUpClass func(unsafe.Pointer) objc.Class
	_objc_msgSend func()
	_objc_msgSendSuper func()
	_objc_msgSendSuper_stret func()
	_objc_msgSend_fp2ret func()
	_objc_msgSend_fpret func()
	_objc_msgSend_stret func()
	_objc_registerClassPair func(objc.Class)
	_objc_registerProtocol func(unsafe.Pointer)
	_objc_removeAssociatedObjects func(objc.ID)
	_objc_removeExceptionHandler func(unsafe.Pointer)
	_objc_setAssociatedObject func(objc.ID, unsafe.Pointer, objc.ID, objc_AssociationPolicy)
	_objc_setEnumerationMutationHandler func()
	_objc_setExceptionMatcher func(objc_exception_matcher) objc_exception_matcher
	_objc_setExceptionPreprocessor func(objc_exception_preprocessor) objc_exception_preprocessor
	_objc_setForwardHandler func(unsafe.Pointer, unsafe.Pointer)
	_objc_setHook_getClass func(objc_hook_getClass, unsafe.Pointer)
	_objc_setHook_getImageName func(objc_hook_getImageName, unsafe.Pointer)
	_objc_setHook_lazyClassNamer func(objc_hook_lazyClassNamer, unsafe.Pointer)
	_objc_setUncaughtExceptionHandler func(objc_uncaught_exception_handler) objc_uncaught_exception_handler
	_objc_storeWeak func(unsafe.Pointer, objc.ID) objc.ID
	_objc_sync_enter func(objc.ID) int
	_objc_sync_exit func(objc.ID) int
	_objc_terminate func()
	_object_copy func(objc.ID, uintptr) objc.ID
	_object_copyFromZone func(objc.ID, uintptr, unsafe.Pointer) objc.ID
	_object_dispose func(objc.ID) objc.ID
	_object_getClass func(objc.ID) objc.Class
	_object_getClassName func(objc.ID) unsafe.Pointer
	_object_getIndexedIvars func(objc.ID) unsafe.Pointer
	_object_getInstanceVariable func(objc.ID, unsafe.Pointer, unsafe.Pointer) Ivar
	_object_getIvar func(objc.ID, Ivar) objc.ID
	_object_isClass func(objc.ID) bool
	_object_setClass func(objc.ID, objc.Class) objc.Class
	_object_setInstanceVariable func(objc.ID, unsafe.Pointer, unsafe.Pointer) Ivar
	_object_setInstanceVariableWithStrongDefault func(objc.ID, unsafe.Pointer, unsafe.Pointer) Ivar
	_object_setIvar func(objc.ID, Ivar, objc.ID)
	_object_setIvarWithStrongDefault func(objc.ID, Ivar, objc.ID)
	_property_copyAttributeList func(objc_property_t, unsafe.Pointer) unsafe.Pointer
	_property_copyAttributeValue func(objc_property_t, unsafe.Pointer) unsafe.Pointer
	_property_getAttributes func(objc_property_t) unsafe.Pointer
	_property_getName func(objc_property_t) unsafe.Pointer
	_protocol_addMethodDescription func(unsafe.Pointer, objc.SEL, unsafe.Pointer, bool, bool)
	_protocol_addProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool, bool)
	_protocol_addProtocol func(unsafe.Pointer, unsafe.Pointer)
	_protocol_conformsToProtocol func(unsafe.Pointer, unsafe.Pointer) bool
	_protocol_copyMethodDescriptionList func(unsafe.Pointer, bool, bool, unsafe.Pointer) unsafe.Pointer
	_protocol_copyPropertyList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_protocol_copyPropertyList2 func(unsafe.Pointer, unsafe.Pointer, bool, bool) unsafe.Pointer
	_protocol_copyProtocolList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_protocol_getMethodDescription func(unsafe.Pointer, objc.SEL, bool, bool) unsafe.Pointer
	_protocol_getName func(unsafe.Pointer) unsafe.Pointer
	_protocol_getProperty func(unsafe.Pointer, unsafe.Pointer, bool, bool) objc_property_t
	_protocol_isEqual func(unsafe.Pointer, unsafe.Pointer) bool
	_sel_getName func(objc.SEL) unsafe.Pointer
	_sel_getUid func(unsafe.Pointer) objc.SEL
	_sel_isEqual func(objc.SEL, objc.SEL) bool
	_sel_isMapped func(objc.SEL) bool
	_sel_registerName func(unsafe.Pointer) objc.SEL
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



// NXCompareHashTables is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXCompareHashTables
func NXCompareHashTables(table1 unsafe.Pointer, table2 unsafe.Pointer) bool {
	return _NXCompareHashTables(table1, table2)
}/* debug [functions.gen.go/function]: NXCompareHashTables */

// NXCopyHashTable is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXCopyHashTable
func NXCopyHashTable(table unsafe.Pointer) unsafe.Pointer {
	return _NXCopyHashTable(table)
}/* debug [functions.gen.go/function]: NXCopyHashTable */

// NXCountHashTable is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXCountHashTable
func NXCountHashTable(table unsafe.Pointer) unsafe.Pointer {
	return _NXCountHashTable(table)
}/* debug [functions.gen.go/function]: NXCountHashTable */

// NXCreateHashTable is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXCreateHashTable
func NXCreateHashTable(prototype NXHashTablePrototype, capacity unsafe.Pointer, info unsafe.Pointer) unsafe.Pointer {
	return _NXCreateHashTable(prototype, capacity, info)
}/* debug [functions.gen.go/function]: NXCreateHashTable */

// NXCreateHashTableFromZone is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXCreateHashTableFromZone
func NXCreateHashTableFromZone(prototype NXHashTablePrototype, capacity unsafe.Pointer, info unsafe.Pointer, zone unsafe.Pointer) unsafe.Pointer {
	return _NXCreateHashTableFromZone(prototype, capacity, info, zone)
}/* debug [functions.gen.go/function]: NXCreateHashTableFromZone */

// NXEmptyHashTable is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXEmptyHashTable
func NXEmptyHashTable(table unsafe.Pointer) {
	_NXEmptyHashTable(table)
}/* debug [functions.gen.go/function]: NXEmptyHashTable */

// NXFreeHashTable is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXFreeHashTable
func NXFreeHashTable(table unsafe.Pointer) {
	_NXFreeHashTable(table)
}/* debug [functions.gen.go/function]: NXFreeHashTable */

// NXHashGet is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashGet
func NXHashGet(table unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _NXHashGet(table, data)
}/* debug [functions.gen.go/function]: NXHashGet */

// NXHashInsert is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashInsert
func NXHashInsert(table unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _NXHashInsert(table, data)
}/* debug [functions.gen.go/function]: NXHashInsert */

// NXHashInsertIfAbsent is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashInsertIfAbsent
func NXHashInsertIfAbsent(table unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _NXHashInsertIfAbsent(table, data)
}/* debug [functions.gen.go/function]: NXHashInsertIfAbsent */

// NXHashMember is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashMember
func NXHashMember(table unsafe.Pointer, data unsafe.Pointer) int {
	return _NXHashMember(table, data)
}/* debug [functions.gen.go/function]: NXHashMember */

// NXHashRemove is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashRemove
func NXHashRemove(table unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _NXHashRemove(table, data)
}/* debug [functions.gen.go/function]: NXHashRemove */

// NXInitHashState is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXInitHashState
func NXInitHashState(table unsafe.Pointer) NXHashState {
	return _NXInitHashState(table)
}/* debug [functions.gen.go/function]: NXInitHashState */

// NXNextHashState is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXNextHashState
func NXNextHashState(table unsafe.Pointer, state unsafe.Pointer, data unsafe.Pointer) int {
	return _NXNextHashState(table, state, data)
}/* debug [functions.gen.go/function]: NXNextHashState */

// NXNoEffectFree is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXNoEffectFree
func NXNoEffectFree(info unsafe.Pointer, data unsafe.Pointer) {
	_NXNoEffectFree(info, data)
}/* debug [functions.gen.go/function]: NXNoEffectFree */

// NXPtrHash is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXPtrHash
func NXPtrHash(info unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _NXPtrHash(info, data)
}/* debug [functions.gen.go/function]: NXPtrHash */

// NXPtrIsEqual is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXPtrIsEqual
func NXPtrIsEqual(info unsafe.Pointer, data1 unsafe.Pointer, data2 unsafe.Pointer) int {
	return _NXPtrIsEqual(info, data1, data2)
}/* debug [functions.gen.go/function]: NXPtrIsEqual */

// NXReallyFree is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXReallyFree
func NXReallyFree(info unsafe.Pointer, data unsafe.Pointer) {
	_NXReallyFree(info, data)
}/* debug [functions.gen.go/function]: NXReallyFree */

// NXResetHashTable is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXResetHashTable
func NXResetHashTable(table unsafe.Pointer) {
	_NXResetHashTable(table)
}/* debug [functions.gen.go/function]: NXResetHashTable */

// NXStrHash is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXStrHash
func NXStrHash(info unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _NXStrHash(info, data)
}/* debug [functions.gen.go/function]: NXStrHash */

// NXStrIsEqual is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.1.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXStrIsEqual
func NXStrIsEqual(info unsafe.Pointer, data1 unsafe.Pointer, data2 unsafe.Pointer) int {
	return _NXStrIsEqual(info, data1, data2)
}/* debug [functions.gen.go/function]: NXStrIsEqual */

// Adds a new instance variable to a class.
//
// Added in macOS 10.5.
// Adds a new instance variable to a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_addIvar(_:_:_:_:_:)
func class_addIvar(cls objc.Class, name unsafe.Pointer, size uintptr, alignment uint8, types unsafe.Pointer) bool {
	return _class_addIvar(cls, name, size, alignment, types)
}/* debug [functions.gen.go/function]: class_addIvar */

// Adds a new method to a class with a given name and implementation.
//
// Added in macOS 10.5.
// Adds a new method to a class with a given name and implementation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_addMethod(_:_:_:_:)
func class_addMethod(cls objc.Class, name objc.SEL, imp IMP, types unsafe.Pointer) bool {
	return _class_addMethod(cls, name, imp, types)
}/* debug [functions.gen.go/function]: class_addMethod */

// Adds a property to a class.
//
// Added in macOS 10.7.
// Adds a property to a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_addProperty(_:_:_:_:)
func class_addProperty(cls objc.Class, name unsafe.Pointer, attributes unsafe.Pointer, attributeCount unsafe.Pointer) bool {
	return _class_addProperty(cls, name, attributes, attributeCount)
}/* debug [functions.gen.go/function]: class_addProperty */

// Adds a protocol to a class.
//
// Added in macOS 10.5.
// Adds a protocol to a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_addProtocol(_:_:)
func class_addProtocol(cls objc.Class, protocol_ unsafe.Pointer) bool {
	return _class_addProtocol(cls, protocol_)
}/* debug [functions.gen.go/function]: class_addProtocol */

// Returns a Boolean value that indicates whether a class conforms to a given protocol.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether a class conforms to a given protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_conformsToProtocol(_:_:)
func class_conformsToProtocol(cls objc.Class, protocol_ unsafe.Pointer) bool {
	return _class_conformsToProtocol(cls, protocol_)
}/* debug [functions.gen.go/function]: class_conformsToProtocol */

// Describes the instance variables declared by a class.
//
// Added in macOS 10.5.
// Describes the instance variables declared by a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_copyIvarList(_:_:)
func class_copyIvarList(cls objc.Class, outCount unsafe.Pointer) unsafe.Pointer {
	return _class_copyIvarList(cls, outCount)
}/* debug [functions.gen.go/function]: class_copyIvarList */

// Describes the instance methods implemented by a class.
//
// Added in macOS 10.5.
// Describes the instance methods implemented by a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_copyMethodList(_:_:)
func class_copyMethodList(cls objc.Class, outCount unsafe.Pointer) unsafe.Pointer {
	return _class_copyMethodList(cls, outCount)
}/* debug [functions.gen.go/function]: class_copyMethodList */

// Describes the properties declared by a class.
//
// Added in macOS 10.5.
// Describes the properties declared by a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_copyPropertyList(_:_:)
func class_copyPropertyList(cls objc.Class, outCount unsafe.Pointer) unsafe.Pointer {
	return _class_copyPropertyList(cls, outCount)
}/* debug [functions.gen.go/function]: class_copyPropertyList */

// Describes the protocols adopted by a class.
//
// Added in macOS 10.5.
// Describes the protocols adopted by a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_copyProtocolList(_:_:)
func class_copyProtocolList(cls objc.Class, outCount unsafe.Pointer) unsafe.Pointer {
	return _class_copyProtocolList(cls, outCount)
}/* debug [functions.gen.go/function]: class_copyProtocolList */

// Creates an instance of a class, allocating memory for the class in the default malloc memory zone.
//
// Added in macOS 10.0.
// Creates an instance of a class, allocating memory for the class in the default malloc memory zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_createInstance(_:_:)
func class_createInstance(cls objc.Class, extraBytes uintptr) objc.ID {
	return _class_createInstance(cls, extraBytes)
}/* debug [functions.gen.go/function]: class_createInstance */

// class_createInstanceFromZone is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_createInstanceFromZone
func class_createInstanceFromZone(p0 objc.Class, idxIvars uintptr, zone unsafe.Pointer) objc.ID {
	return _class_createInstanceFromZone(p0, idxIvars, zone)
}/* debug [functions.gen.go/function]: class_createInstanceFromZone */

// Returns a pointer to the data structure describing a given class method for a given class.
//
// Added in macOS 10.0.
// Returns a pointer to the data structure describing a given class method for a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getClassMethod(_:_:)
func class_getClassMethod(cls objc.Class, name objc.SEL) Method {
	return _class_getClassMethod(cls, name)
}/* debug [functions.gen.go/function]: class_getClassMethod */

// Returns the for a specified class variable of a given class.
//
// Added in macOS 10.5.
// Returns the for a specified class variable of a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getClassVariable(_:_:)
func class_getClassVariable(cls objc.Class, name unsafe.Pointer) Ivar {
	return _class_getClassVariable(cls, name)
}/* debug [functions.gen.go/function]: class_getClassVariable */

// Returns the name of the dynamic library a class originated from.
//
// Added in macOS 10.5.
// Returns the name of the dynamic library a class originated from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getImageName(_:)
func class_getImageName(cls objc.Class) unsafe.Pointer {
	return _class_getImageName(cls)
}/* debug [functions.gen.go/function]: class_getImageName */

// Returns a specified instance method for a given class.
//
// Added in macOS 10.0.
// Returns a specified instance method for a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getInstanceMethod(_:_:)
func class_getInstanceMethod(cls objc.Class, name objc.SEL) Method {
	return _class_getInstanceMethod(cls, name)
}/* debug [functions.gen.go/function]: class_getInstanceMethod */

// Returns the size of instances of a class.
//
// Added in macOS 10.5.
// Returns the size of instances of a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getInstanceSize(_:)
func class_getInstanceSize(cls objc.Class) uintptr {
	return _class_getInstanceSize(cls)
}/* debug [functions.gen.go/function]: class_getInstanceSize */

// Returns the for a specified instance variable of a given class.
//
// Added in macOS 10.0.
// Returns the for a specified instance variable of a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getInstanceVariable(_:_:)
func class_getInstanceVariable(cls objc.Class, name unsafe.Pointer) Ivar {
	return _class_getInstanceVariable(cls, name)
}/* debug [functions.gen.go/function]: class_getInstanceVariable */

// Returns a description of the layout for a given class.
//
// Added in macOS 10.5.
// Returns a description of the layout for a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getIvarLayout(_:)
func class_getIvarLayout(cls objc.Class) unsafe.Pointer {
	return _class_getIvarLayout(cls)
}/* debug [functions.gen.go/function]: class_getIvarLayout */

// Returns the function pointer that would be called if a particular message were sent to an instance of a class.
//
// Added in macOS 10.5.
// Returns the function pointer that would be called if a particular message were sent to an instance of a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getMethodImplementation(_:_:)
func class_getMethodImplementation(cls objc.Class, name objc.SEL) IMP {
	return _class_getMethodImplementation(cls, name)
}/* debug [functions.gen.go/function]: class_getMethodImplementation */

// Returns the function pointer that would be called if a particular message were sent to an instance of a class.
//
// Added in macOS 10.5.
// Returns the function pointer that would be called if a particular message were sent to an instance of a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getMethodImplementation_stret(_:_:)
func class_getMethodImplementation_stret(cls objc.Class, name objc.SEL) IMP {
	return _class_getMethodImplementation_stret(cls, name)
}/* debug [functions.gen.go/function]: class_getMethodImplementation_stret */

// Returns the name of a class.
//
// Added in macOS 10.5.
// Returns the name of a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getName(_:)
func class_getName(cls objc.Class) unsafe.Pointer {
	return _class_getName(cls)
}/* debug [functions.gen.go/function]: class_getName */

// Returns a property with a given name of a given class.
//
// Added in macOS 10.5.
// Returns a property with a given name of a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getProperty(_:_:)
func class_getProperty(cls objc.Class, name unsafe.Pointer) objc_property_t {
	return _class_getProperty(cls, name)
}/* debug [functions.gen.go/function]: class_getProperty */

// Returns the superclass of a class.
//
// Added in macOS 10.5.
// Returns the superclass of a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getSuperclass(_:)
func class_getSuperclass(cls objc.Class) objc.Class {
	return _class_getSuperclass(cls)
}/* debug [functions.gen.go/function]: class_getSuperclass */

// Returns the version number of a class definition.
//
// Added in macOS 10.0.
// Returns the version number of a class definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getVersion(_:)
func class_getVersion(cls objc.Class) int {
	return _class_getVersion(cls)
}/* debug [functions.gen.go/function]: class_getVersion */

// Returns a description of the layout of weak s for a given class.
//
// Added in macOS 10.5.
// Returns a description of the layout of weak s for a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_getWeakIvarLayout(_:)
func class_getWeakIvarLayout(cls objc.Class) unsafe.Pointer {
	return _class_getWeakIvarLayout(cls)
}/* debug [functions.gen.go/function]: class_getWeakIvarLayout */

// Returns a Boolean value that indicates whether a class object is a metaclass.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether a class object is a metaclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_isMetaClass(_:)
func class_isMetaClass(cls objc.Class) bool {
	return _class_isMetaClass(cls)
}/* debug [functions.gen.go/function]: class_isMetaClass */

// class_lookupMethod is a ObjectiveC function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_lookupMethod(_:_:)
func class_lookupMethod(cls objc.Class, sel objc.SEL) IMP {
	return _class_lookupMethod(cls, sel)
}/* debug [functions.gen.go/function]: class_lookupMethod */

// Replaces the implementation of a method for a given class.
//
// Added in macOS 10.5.
// Replaces the implementation of a method for a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_replaceMethod(_:_:_:_:)
func class_replaceMethod(cls objc.Class, name objc.SEL, imp IMP, types unsafe.Pointer) IMP {
	return _class_replaceMethod(cls, name, imp, types)
}/* debug [functions.gen.go/function]: class_replaceMethod */

// Replace a property of a class.
//
// Added in macOS 10.7.
// Replace a property of a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_replaceProperty(_:_:_:_:)
func class_replaceProperty(cls objc.Class, name unsafe.Pointer, attributes unsafe.Pointer, attributeCount unsafe.Pointer) {
	_class_replaceProperty(cls, name, attributes, attributeCount)
}/* debug [functions.gen.go/function]: class_replaceProperty */

// class_respondsToMethod is a ObjectiveC function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_respondsToMethod(_:_:)
func class_respondsToMethod(cls objc.Class, sel objc.SEL) bool {
	return _class_respondsToMethod(cls, sel)
}/* debug [functions.gen.go/function]: class_respondsToMethod */

// Returns a Boolean value that indicates whether instances of a class respond to a particular selector.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether instances of a class respond to a particular selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_respondsToSelector(_:_:)
func class_respondsToSelector(cls objc.Class, sel objc.SEL) bool {
	return _class_respondsToSelector(cls, sel)
}/* debug [functions.gen.go/function]: class_respondsToSelector */

// Sets the layout for a given class.
//
// Added in macOS 10.5.
// Sets the layout for a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_setIvarLayout(_:_:)
func class_setIvarLayout(cls objc.Class, layout unsafe.Pointer) {
	_class_setIvarLayout(cls, layout)
}/* debug [functions.gen.go/function]: class_setIvarLayout */

// Sets the superclass of a given class.

// Sets the superclass of a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_setSuperclass(_:_:)
func class_setSuperclass(cls objc.Class, newSuper objc.Class) objc.Class {
	return _class_setSuperclass(cls, newSuper)
}/* debug [functions.gen.go/function]: class_setSuperclass */

// Sets the version number of a class definition.
//
// Added in macOS 10.0.
// Sets the version number of a class definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_setVersion(_:_:)
func class_setVersion(cls objc.Class, version int) {
	_class_setVersion(cls, version)
}/* debug [functions.gen.go/function]: class_setVersion */

// Sets the layout for weak s for a given class.
//
// Added in macOS 10.5.
// Sets the layout for weak s for a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/class_setWeakIvarLayout(_:_:)
func class_setWeakIvarLayout(cls objc.Class, layout unsafe.Pointer) {
	_class_setWeakIvarLayout(cls, layout)
}/* debug [functions.gen.go/function]: class_setWeakIvarLayout */

// Returns the block associated with an that was created using .
//
// Added in macOS 10.7.
// Returns the block associated with an that was created using .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/imp_getBlock(_:)
func imp_getBlock(anImp IMP) objc.ID {
	return _imp_getBlock(anImp)
}/* debug [functions.gen.go/function]: imp_getBlock */

// Creates a pointer to a function that calls the specified block when the method is called.
//
// Added in macOS 10.7.
// Creates a pointer to a function that calls the specified block when the method is called.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/imp_implementationWithBlock(_:)
func imp_implementationWithBlock(block objc.ID) IMP {
	return _imp_implementationWithBlock(block)
}/* debug [functions.gen.go/function]: imp_implementationWithBlock */

// Disassociates a block from an that was created using , and releases the copy of the block that was created.
//
// Added in macOS 10.7.
// Disassociates a block from an that was created using , and releases the copy of the block that was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/imp_removeBlock(_:)
func imp_removeBlock(anImp IMP) bool {
	return _imp_removeBlock(anImp)
}/* debug [functions.gen.go/function]: imp_removeBlock */

// Returns the name of an instance variable.
//
// Added in macOS 10.5.
// Returns the name of an instance variable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/ivar_getName(_:)
func ivar_getName(v Ivar) unsafe.Pointer {
	return _ivar_getName(v)
}/* debug [functions.gen.go/function]: ivar_getName */

// Returns the offset of an instance variable.
//
// Added in macOS 10.5.
// Returns the offset of an instance variable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/ivar_getOffset(_:)
func ivar_getOffset(v Ivar) unsafe.Pointer {
	return _ivar_getOffset(v)
}/* debug [functions.gen.go/function]: ivar_getOffset */

// Returns the type string of an instance variable.
//
// Added in macOS 10.5.
// Returns the type string of an instance variable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/ivar_getTypeEncoding(_:)
func ivar_getTypeEncoding(v Ivar) unsafe.Pointer {
	return _ivar_getTypeEncoding(v)
}/* debug [functions.gen.go/function]: ivar_getTypeEncoding */

// Returns a string describing a single parameter type of a method.
//
// Added in macOS 10.5.
// Returns a string describing a single parameter type of a method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_copyArgumentType(_:_:)
func method_copyArgumentType(m Method, index unsafe.Pointer) unsafe.Pointer {
	return _method_copyArgumentType(m, index)
}/* debug [functions.gen.go/function]: method_copyArgumentType */

// Returns a string describing a method’s return type.
//
// Added in macOS 10.5.
// Returns a string describing a method’s return type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_copyReturnType(_:)
func method_copyReturnType(m Method) unsafe.Pointer {
	return _method_copyReturnType(m)
}/* debug [functions.gen.go/function]: method_copyReturnType */

// Exchanges the implementations of two methods.
//
// Added in macOS 10.5.
// Exchanges the implementations of two methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_exchangeImplementations(_:_:)
func method_exchangeImplementations(m1 Method, m2 Method) {
	_method_exchangeImplementations(m1, m2)
}/* debug [functions.gen.go/function]: method_exchangeImplementations */

// Returns by reference a string describing a single parameter type of a method.
//
// Added in macOS 10.5.
// Returns by reference a string describing a single parameter type of a method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_getArgumentType(_:_:_:_:)
func method_getArgumentType(m Method, index unsafe.Pointer, dst unsafe.Pointer, dst_len uintptr) {
	_method_getArgumentType(m, index, dst, dst_len)
}/* debug [functions.gen.go/function]: method_getArgumentType */

// Returns a method description structure for a specified method.
//
// Added in macOS 10.5.
// Returns a method description structure for a specified method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_getDescription(_:)
func method_getDescription(m Method) unsafe.Pointer {
	return _method_getDescription(m)
}/* debug [functions.gen.go/function]: method_getDescription */

// Returns the implementation of a method.
//
// Added in macOS 10.5.
// Returns the implementation of a method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_getImplementation(_:)
func method_getImplementation(m Method) IMP {
	return _method_getImplementation(m)
}/* debug [functions.gen.go/function]: method_getImplementation */

// Returns the name of a method.
//
// Added in macOS 10.5.
// Returns the name of a method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_getName(_:)
func method_getName(m Method) objc.SEL {
	return _method_getName(m)
}/* debug [functions.gen.go/function]: method_getName */

// Returns the number of arguments accepted by a method.
//
// Added in macOS 10.0.
// Returns the number of arguments accepted by a method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_getNumberOfArguments(_:)
func method_getNumberOfArguments(m Method) unsafe.Pointer {
	return _method_getNumberOfArguments(m)
}/* debug [functions.gen.go/function]: method_getNumberOfArguments */

// Returns by reference a string describing a method’s return type.
//
// Added in macOS 10.5.
// Returns by reference a string describing a method’s return type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_getReturnType(_:_:_:)
func method_getReturnType(m Method, dst unsafe.Pointer, dst_len uintptr) {
	_method_getReturnType(m, dst, dst_len)
}/* debug [functions.gen.go/function]: method_getReturnType */

// Returns a string describing a method’s parameter and return types.
//
// Added in macOS 10.5.
// Returns a string describing a method’s parameter and return types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_getTypeEncoding(_:)
func method_getTypeEncoding(m Method) unsafe.Pointer {
	return _method_getTypeEncoding(m)
}/* debug [functions.gen.go/function]: method_getTypeEncoding */

// Calls the implementation of a specified method.
//
// Added in macOS 10.5.
// Calls the implementation of a specified method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_invoke
func method_invoke() {
	_method_invoke()
}/* debug [functions.gen.go/function]: method_invoke */

// Calls the implementation of a specified method that returns a data-structure.
//
// Added in macOS 10.5.
// Calls the implementation of a specified method that returns a data-structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_invoke_stret
func method_invoke_stret() {
	_method_invoke_stret()
}/* debug [functions.gen.go/function]: method_invoke_stret */

// Sets the implementation of a method.
//
// Added in macOS 10.5.
// Sets the implementation of a method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/method_setImplementation(_:_:)
func method_setImplementation(m Method, imp IMP) IMP {
	return _method_setImplementation(m, imp)
}/* debug [functions.gen.go/function]: method_setImplementation */

// objc_addExceptionHandler is a ObjectiveC function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_addExceptionHandler(_:_:)
func objc_addExceptionHandler(fn objc_exception_handler, context unsafe.Pointer) unsafe.Pointer {
	return _objc_addExceptionHandler(fn, context)
}/* debug [functions.gen.go/function]: objc_addExceptionHandler */

// objc_addLoadImageFunc is a ObjectiveC function.
//
// Added in macOS 10.15.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_addLoadImageFunc(_:)
func objc_addLoadImageFunc(func_ objc_func_loadImage) {
	_objc_addLoadImageFunc(func_)
}/* debug [functions.gen.go/function]: objc_addLoadImageFunc */

// Creates a new class and metaclass.
//
// Added in macOS 10.5.
// Creates a new class and metaclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_allocateClassPair(_:_:_:)
func objc_allocateClassPair(superclass objc.Class, name unsafe.Pointer, extraBytes uintptr) objc.Class {
	return _objc_allocateClassPair(superclass, name, extraBytes)
}/* debug [functions.gen.go/function]: objc_allocateClassPair */

// Creates a new protocol instance.
//
// Added in macOS 10.7.
// Creates a new protocol instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_allocateProtocol(_:)
func objc_allocateProtocol(name unsafe.Pointer) unsafe.Pointer {
	return _objc_allocateProtocol(name)
}/* debug [functions.gen.go/function]: objc_allocateProtocol */

// objc_begin_catch is a ObjectiveC function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_begin_catch(_:)
func objc_begin_catch(exc_buf unsafe.Pointer) objc.ID {
	return _objc_begin_catch(exc_buf)
}/* debug [functions.gen.go/function]: objc_begin_catch */

// Creates an instance of a class at the specified location.
//
// Added in macOS 10.6.
// Creates an instance of a class at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_constructInstance
func objc_constructInstance(cls objc.Class, bytes unsafe.Pointer) objc.ID {
	return _objc_constructInstance(cls, bytes)
}/* debug [functions.gen.go/function]: objc_constructInstance */

// Creates and returns a list of pointers to all registered class definitions.
//
// Added in macOS 10.7.
// Creates and returns a list of pointers to all registered class definitions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_copyClassList(_:)
func objc_copyClassList(outCount unsafe.Pointer) unsafe.Pointer {
	return _objc_copyClassList(outCount)
}/* debug [functions.gen.go/function]: objc_copyClassList */

// Returns the names of all the classes within a specified library or framework.
//
// Added in macOS 10.5.
// Returns the names of all the classes within a specified library or framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_copyClassNamesForImage(_:_:)
func objc_copyClassNamesForImage(image unsafe.Pointer, outCount unsafe.Pointer) unsafe.Pointer {
	return _objc_copyClassNamesForImage(image, outCount)
}/* debug [functions.gen.go/function]: objc_copyClassNamesForImage */

// Returns the names of all the loaded Objective-C frameworks and dynamic libraries.
//
// Added in macOS 10.5.
// Returns the names of all the loaded Objective-C frameworks and dynamic libraries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_copyImageNames(_:)
func objc_copyImageNames(outCount unsafe.Pointer) unsafe.Pointer {
	return _objc_copyImageNames(outCount)
}/* debug [functions.gen.go/function]: objc_copyImageNames */

// Returns an array of all the protocols known to the runtime.
//
// Added in macOS 10.5.
// Returns an array of all the protocols known to the runtime.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_copyProtocolList(_:)
func objc_copyProtocolList(outCount unsafe.Pointer) unsafe.Pointer {
	return _objc_copyProtocolList(outCount)
}/* debug [functions.gen.go/function]: objc_copyProtocolList */

// Destroys an instance of a class without freeing memory and removes any of its associated references.
//
// Added in macOS 10.6.
// Destroys an instance of a class without freeing memory and removes any of its associated references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_destructInstance
func objc_destructInstance(obj objc.ID) unsafe.Pointer {
	return _objc_destructInstance(obj)
}/* debug [functions.gen.go/function]: objc_destructInstance */

// Destroys a class and its associated metaclass.
//
// Added in macOS 10.5.
// Destroys a class and its associated metaclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_disposeClassPair(_:)
func objc_disposeClassPair(cls objc.Class) {
	_objc_disposeClassPair(cls)
}/* debug [functions.gen.go/function]: objc_disposeClassPair */

// Used by Foundation’s Key-Value Observing.
//
// Added in macOS 10.5.
// Used by Foundation’s Key-Value Observing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_duplicateClass(_:_:_:)
func objc_duplicateClass(original objc.Class, name unsafe.Pointer, extraBytes uintptr) objc.Class {
	return _objc_duplicateClass(original, name, extraBytes)
}/* debug [functions.gen.go/function]: objc_duplicateClass */

// objc_end_catch is a ObjectiveC function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_end_catch()
func objc_end_catch() {
	_objc_end_catch()
}/* debug [functions.gen.go/function]: objc_end_catch */

// objc_enumerateClasses is a ObjectiveC function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_enumerateClasses
func objc_enumerateClasses(image unsafe.Pointer, namePrefix unsafe.Pointer, conformingTo unsafe.Pointer, subclassing objc.Class) {
	_objc_enumerateClasses(image, namePrefix, conformingTo, subclassing)
}/* debug [functions.gen.go/function]: objc_enumerateClasses */

// Inserted by the compiler when a mutation is detected during a foreach iteration.
//
// Added in macOS 10.5.
// Inserted by the compiler when a mutation is detected during a foreach iteration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_enumerationMutation(_:)
func objc_enumerationMutation(obj objc.ID) {
	_objc_enumerationMutation(obj)
}/* debug [functions.gen.go/function]: objc_enumerationMutation */

// objc_exception_rethrow is a ObjectiveC function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_exception_rethrow()
func objc_exception_rethrow() {
	_objc_exception_rethrow()
}/* debug [functions.gen.go/function]: objc_exception_rethrow */

// Throw a runtime exception. This function is inserted by the compiler where \c @throw would otherwise be.
//
// Added in macOS 10.5.
// Throw a runtime exception. This function is inserted by the compiler where \c @throw would otherwise be.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_exception_throw(_:)
func objc_exception_throw(exception objc.ID) {
	_objc_exception_throw(exception)
}/* debug [functions.gen.go/function]: objc_exception_throw */

// Returns the value associated with a given object for a given key.
//
// Added in macOS 10.6.
// Returns the value associated with a given object for a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_getAssociatedObject(_:_:)
func objc_getAssociatedObject(object objc.ID, key unsafe.Pointer) objc.ID {
	return _objc_getAssociatedObject(object, key)
}/* debug [functions.gen.go/function]: objc_getAssociatedObject */

// Returns the class definition of a specified class.
//
// Added in macOS 10.0.
// Returns the class definition of a specified class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_getClass(_:)
func objc_getClass(name unsafe.Pointer) objc.ID {
	return _objc_getClass(name)
}/* debug [functions.gen.go/function]: objc_getClass */

// Obtains the list of registered class definitions.
//
// Added in macOS 10.0.
// Obtains the list of registered class definitions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_getClassList(_:_:)
func objc_getClassList(buffer unsafe.Pointer, bufferCount int) int {
	return _objc_getClassList(buffer, bufferCount)
}/* debug [functions.gen.go/function]: objc_getClassList */

// Used by CoreFoundation’s toll-free bridging.
//
// Added in macOS 10.5.
// Used by CoreFoundation’s toll-free bridging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_getFutureClass
func objc_getFutureClass(name unsafe.Pointer) objc.Class {
	return _objc_getFutureClass(name)
}/* debug [functions.gen.go/function]: objc_getFutureClass */

// Returns the metaclass definition of a specified class.
//
// Added in macOS 10.0.
// Returns the metaclass definition of a specified class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_getMetaClass(_:)
func objc_getMetaClass(name unsafe.Pointer) objc.ID {
	return _objc_getMetaClass(name)
}/* debug [functions.gen.go/function]: objc_getMetaClass */

// Returns a specified protocol.
//
// Added in macOS 10.5.
// Returns a specified protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_getProtocol(_:)
func objc_getProtocol(name unsafe.Pointer) unsafe.Pointer {
	return _objc_getProtocol(name)
}/* debug [functions.gen.go/function]: objc_getProtocol */

// Returns the class definition of a specified class.
//
// Added in macOS 10.0.
// Returns the class definition of a specified class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_getRequiredClass(_:)
func objc_getRequiredClass(name unsafe.Pointer) objc.Class {
	return _objc_getRequiredClass(name)
}/* debug [functions.gen.go/function]: objc_getRequiredClass */

// Loads the object referenced by a weak pointer and returns it.
//
// Added in macOS 10.7.
// Loads the object referenced by a weak pointer and returns it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_loadWeak(_:)
func objc_loadWeak(location unsafe.Pointer) objc.ID {
	return _objc_loadWeak(location)
}/* debug [functions.gen.go/function]: objc_loadWeak */

// Returns the class definition of a specified class.
//
// Added in macOS 10.0.
// Returns the class definition of a specified class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_lookUpClass(_:)
func objc_lookUpClass(name unsafe.Pointer) objc.Class {
	return _objc_lookUpClass(name)
}/* debug [functions.gen.go/function]: objc_lookUpClass */

// Sends a message with a simple return value to an instance of a class.
//
// Added in macOS 10.0.
// Sends a message with a simple return value to an instance of a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_msgSend
func objc_msgSend() {
	_objc_msgSend()
}/* debug [functions.gen.go/function]: objc_msgSend */

// Sends a message with a simple return value to the superclass of an instance of a class.
//
// Added in macOS 10.0.
// Sends a message with a simple return value to the superclass of an instance of a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_msgSendSuper
func objc_msgSendSuper() {
	_objc_msgSendSuper()
}/* debug [functions.gen.go/function]: objc_msgSendSuper */

// Sends a message with a data-structure return value to the superclass of an instance of a class.
//
// Added in macOS 10.0.
// Sends a message with a data-structure return value to the superclass of an instance of a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_msgSendSuper_stret
func objc_msgSendSuper_stret() {
	_objc_msgSendSuper_stret()
}/* debug [functions.gen.go/function]: objc_msgSendSuper_stret */

// objc_msgSend_fp2ret is a ObjectiveC function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_msgSend_fp2ret
func objc_msgSend_fp2ret() {
	_objc_msgSend_fp2ret()
}/* debug [functions.gen.go/function]: objc_msgSend_fp2ret */

// Sends a message with a floating-point return value to an instance of a class.
//
// Added in macOS 10.5.
// Sends a message with a floating-point return value to an instance of a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_msgSend_fpret
func objc_msgSend_fpret() {
	_objc_msgSend_fpret()
}/* debug [functions.gen.go/function]: objc_msgSend_fpret */

// Sends a message with a data-structure return value to an instance of a class.
//
// Added in macOS 10.0.
// Sends a message with a data-structure return value to an instance of a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_msgSend_stret
func objc_msgSend_stret() {
	_objc_msgSend_stret()
}/* debug [functions.gen.go/function]: objc_msgSend_stret */

// Registers a class that was allocated using .
//
// Added in macOS 10.5.
// Registers a class that was allocated using .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_registerClassPair(_:)
func objc_registerClassPair(cls objc.Class) {
	_objc_registerClassPair(cls)
}/* debug [functions.gen.go/function]: objc_registerClassPair */

// Registers a newly created protocol with the Objective-C runtime.
//
// Added in macOS 10.7.
// Registers a newly created protocol with the Objective-C runtime.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_registerProtocol(_:)
func objc_registerProtocol(proto unsafe.Pointer) {
	_objc_registerProtocol(proto)
}/* debug [functions.gen.go/function]: objc_registerProtocol */

// Removes all associations for a given object.
//
// Added in macOS 10.6.
// Removes all associations for a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_removeAssociatedObjects(_:)
func objc_removeAssociatedObjects(object objc.ID) {
	_objc_removeAssociatedObjects(object)
}/* debug [functions.gen.go/function]: objc_removeAssociatedObjects */

// objc_removeExceptionHandler is a ObjectiveC function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_removeExceptionHandler(_:)
func objc_removeExceptionHandler(token unsafe.Pointer) {
	_objc_removeExceptionHandler(token)
}/* debug [functions.gen.go/function]: objc_removeExceptionHandler */

// Sets an associated value for a given object using a given key and association policy.
//
// Added in macOS 10.6.
// Sets an associated value for a given object using a given key and association policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setAssociatedObject(_:_:_:_:)
func objc_setAssociatedObject(object objc.ID, key unsafe.Pointer, value objc.ID, policy objc_AssociationPolicy) {
	_objc_setAssociatedObject(object, key, value, policy)
}/* debug [functions.gen.go/function]: objc_setAssociatedObject */

// Sets the current mutation handler.
//
// Added in macOS 10.5.
// Sets the current mutation handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setEnumerationMutationHandler(_:)
func objc_setEnumerationMutationHandler() {
	_objc_setEnumerationMutationHandler()
}/* debug [functions.gen.go/function]: objc_setEnumerationMutationHandler */

// objc_setExceptionMatcher is a ObjectiveC function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setExceptionMatcher(_:)
func objc_setExceptionMatcher(fn objc_exception_matcher) objc_exception_matcher {
	return _objc_setExceptionMatcher(fn)
}/* debug [functions.gen.go/function]: objc_setExceptionMatcher */

// objc_setExceptionPreprocessor is a ObjectiveC function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setExceptionPreprocessor(_:)
func objc_setExceptionPreprocessor(fn objc_exception_preprocessor) objc_exception_preprocessor {
	return _objc_setExceptionPreprocessor(fn)
}/* debug [functions.gen.go/function]: objc_setExceptionPreprocessor */

// Set the function to be called by objc_msgForward.
//
// Added in macOS 10.5.
// Set the function to be called by objc_msgForward.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setForwardHandler(_:_:)
func objc_setForwardHandler(fwd unsafe.Pointer, fwd_stret unsafe.Pointer) {
	_objc_setForwardHandler(fwd, fwd_stret)
}/* debug [functions.gen.go/function]: objc_setForwardHandler */

// objc_setHook_getClass is a ObjectiveC function.
//
// Added in macOS 10.14.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setHook_getClass(_:_:)
func objc_setHook_getClass(newValue objc_hook_getClass, outOldValue unsafe.Pointer) {
	_objc_setHook_getClass(newValue, outOldValue)
}/* debug [functions.gen.go/function]: objc_setHook_getClass */

// objc_setHook_getImageName is a ObjectiveC function.
//
// Added in macOS 10.14.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setHook_getImageName(_:_:)
func objc_setHook_getImageName(newValue objc_hook_getImageName, outOldValue unsafe.Pointer) {
	_objc_setHook_getImageName(newValue, outOldValue)
}/* debug [functions.gen.go/function]: objc_setHook_getImageName */

// objc_setHook_lazyClassNamer is a ObjectiveC function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setHook_lazyClassNamer(_:_:)
func objc_setHook_lazyClassNamer(newValue objc_hook_lazyClassNamer, oldOutValue unsafe.Pointer) {
	_objc_setHook_lazyClassNamer(newValue, oldOutValue)
}/* debug [functions.gen.go/function]: objc_setHook_lazyClassNamer */

// objc_setUncaughtExceptionHandler is a ObjectiveC function.
//
// Added in macOS 10.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_setUncaughtExceptionHandler(_:)
func objc_setUncaughtExceptionHandler(fn objc_uncaught_exception_handler) objc_uncaught_exception_handler {
	return _objc_setUncaughtExceptionHandler(fn)
}/* debug [functions.gen.go/function]: objc_setUncaughtExceptionHandler */

// Stores a new value in a variable.
//
// Added in macOS 10.7.
// Stores a new value in a variable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_storeWeak(_:_:)
func objc_storeWeak(location unsafe.Pointer, obj objc.ID) objc.ID {
	return _objc_storeWeak(location, obj)
}/* debug [functions.gen.go/function]: objc_storeWeak */

// Begin synchronizing on ‘obj’. Allocates recursive pthread_mutex associated with ‘obj’ if needed.
//
// Added in macOS 10.3.
// Begin synchronizing on ‘obj’. Allocates recursive pthread_mutex associated with ‘obj’ if needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_sync_enter
func objc_sync_enter(obj objc.ID) int {
	return _objc_sync_enter(obj)
}/* debug [functions.gen.go/function]: objc_sync_enter */

// End synchronizing on ‘obj’.
//
// Added in macOS 10.3.
// End synchronizing on ‘obj’.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_sync_exit
func objc_sync_exit(obj objc.ID) int {
	return _objc_sync_exit(obj)
}/* debug [functions.gen.go/function]: objc_sync_exit */

// objc_terminate is a ObjectiveC function.
//
// Added in macOS 10.8.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_terminate()
func objc_terminate() {
	_objc_terminate()
}/* debug [functions.gen.go/function]: objc_terminate */

// Returns a copy of a given object.
//
// Added in macOS 10.0.
// Returns a copy of a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_copy
func object_copy(obj objc.ID, size uintptr) objc.ID {
	return _object_copy(obj, size)
}/* debug [functions.gen.go/function]: object_copy */

// object_copyFromZone is a ObjectiveC function.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_copyFromZone
func object_copyFromZone(anObject objc.ID, nBytes uintptr, zone unsafe.Pointer) objc.ID {
	return _object_copyFromZone(anObject, nBytes, zone)
}/* debug [functions.gen.go/function]: object_copyFromZone */

// Frees the memory occupied by a given object.
//
// Added in macOS 10.0.
// Frees the memory occupied by a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_dispose
func object_dispose(obj objc.ID) objc.ID {
	return _object_dispose(obj)
}/* debug [functions.gen.go/function]: object_dispose */

// Returns the class of an object.
//
// Added in macOS 10.5.
// Returns the class of an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_getClass(_:)
func object_getClass(obj objc.ID) objc.Class {
	return _object_getClass(obj)
}/* debug [functions.gen.go/function]: object_getClass */

// Returns the class name of a given object.
//
// Added in macOS 10.0.
// Returns the class name of a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_getClassName(_:)
func object_getClassName(obj objc.ID) unsafe.Pointer {
	return _object_getClassName(obj)
}/* debug [functions.gen.go/function]: object_getClassName */

// Returns a pointer to any extra bytes allocated with a instance given object.
//
// Added in macOS 10.0.
// Returns a pointer to any extra bytes allocated with a instance given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_getIndexedIvars(_:)
func object_getIndexedIvars(obj objc.ID) unsafe.Pointer {
	return _object_getIndexedIvars(obj)
}/* debug [functions.gen.go/function]: object_getIndexedIvars */

// Obtains the value of an instance variable of a class instance.
//
// Added in macOS 10.0.
// Obtains the value of an instance variable of a class instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_getInstanceVariable
func object_getInstanceVariable(obj objc.ID, name unsafe.Pointer, outValue unsafe.Pointer) Ivar {
	return _object_getInstanceVariable(obj, name, outValue)
}/* debug [functions.gen.go/function]: object_getInstanceVariable */

// Reads the value of an instance variable in an object.
//
// Added in macOS 10.5.
// Reads the value of an instance variable in an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_getIvar(_:_:)
func object_getIvar(obj objc.ID, ivar Ivar) objc.ID {
	return _object_getIvar(obj, ivar)
}/* debug [functions.gen.go/function]: object_getIvar */

// object_isClass is a ObjectiveC function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_isClass(_:)
func object_isClass(obj objc.ID) bool {
	return _object_isClass(obj)
}/* debug [functions.gen.go/function]: object_isClass */

// Sets the class of an object.
//
// Added in macOS 10.5.
// Sets the class of an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_setClass(_:_:)
func object_setClass(obj objc.ID, cls objc.Class) objc.Class {
	return _object_setClass(obj, cls)
}/* debug [functions.gen.go/function]: object_setClass */

// Changes the value of an instance variable of a class instance.
//
// Added in macOS 10.0.
// Changes the value of an instance variable of a class instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_setInstanceVariable
func object_setInstanceVariable(obj objc.ID, name unsafe.Pointer, value unsafe.Pointer) Ivar {
	return _object_setInstanceVariable(obj, name, value)
}/* debug [functions.gen.go/function]: object_setInstanceVariable */

// object_setInstanceVariableWithStrongDefault is a ObjectiveC function.
//
// Added in macOS 10.12.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_setInstanceVariableWithStrongDefault
func object_setInstanceVariableWithStrongDefault(obj objc.ID, name unsafe.Pointer, value unsafe.Pointer) Ivar {
	return _object_setInstanceVariableWithStrongDefault(obj, name, value)
}/* debug [functions.gen.go/function]: object_setInstanceVariableWithStrongDefault */

// Sets the value of an instance variable in an object.
//
// Added in macOS 10.5.
// Sets the value of an instance variable in an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_setIvar(_:_:_:)
func object_setIvar(obj objc.ID, ivar Ivar, value objc.ID) {
	_object_setIvar(obj, ivar, value)
}/* debug [functions.gen.go/function]: object_setIvar */

// object_setIvarWithStrongDefault is a ObjectiveC function.
//
// Added in macOS 10.12.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/object_setIvarWithStrongDefault(_:_:_:)
func object_setIvarWithStrongDefault(obj objc.ID, ivar Ivar, value objc.ID) {
	_object_setIvarWithStrongDefault(obj, ivar, value)
}/* debug [functions.gen.go/function]: object_setIvarWithStrongDefault */

// Returns an array of property attributes for a given property.
//
// Added in macOS 10.7.
// Returns an array of property attributes for a given property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/property_copyAttributeList(_:_:)
func property_copyAttributeList(property objc_property_t, outCount unsafe.Pointer) unsafe.Pointer {
	return _property_copyAttributeList(property, outCount)
}/* debug [functions.gen.go/function]: property_copyAttributeList */

// Returns the value of a property attribute given the attribute name.
//
// Added in macOS 10.7.
// Returns the value of a property attribute given the attribute name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/property_copyAttributeValue(_:_:)
func property_copyAttributeValue(property objc_property_t, attributeName unsafe.Pointer) unsafe.Pointer {
	return _property_copyAttributeValue(property, attributeName)
}/* debug [functions.gen.go/function]: property_copyAttributeValue */

// Returns the attribute string of a property.
//
// Added in macOS 10.5.
// Returns the attribute string of a property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/property_getAttributes(_:)
func property_getAttributes(property objc_property_t) unsafe.Pointer {
	return _property_getAttributes(property)
}/* debug [functions.gen.go/function]: property_getAttributes */

// Returns the name of a property.
//
// Added in macOS 10.5.
// Returns the name of a property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/property_getName(_:)
func property_getName(property objc_property_t) unsafe.Pointer {
	return _property_getName(property)
}/* debug [functions.gen.go/function]: property_getName */

// Adds a method to a protocol.
//
// Added in macOS 10.7.
// Adds a method to a protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_addMethodDescription(_:_:_:_:_:)
func protocol_addMethodDescription(proto unsafe.Pointer, name objc.SEL, types unsafe.Pointer, isRequiredMethod bool, isInstanceMethod bool) {
	_protocol_addMethodDescription(proto, name, types, isRequiredMethod, isInstanceMethod)
}/* debug [functions.gen.go/function]: protocol_addMethodDescription */

// Adds a property to a protocol that is under construction.
//
// Added in macOS 10.7.
// Adds a property to a protocol that is under construction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_addProperty(_:_:_:_:_:_:)
func protocol_addProperty(proto unsafe.Pointer, name unsafe.Pointer, attributes unsafe.Pointer, attributeCount unsafe.Pointer, isRequiredProperty bool, isInstanceProperty bool) {
	_protocol_addProperty(proto, name, attributes, attributeCount, isRequiredProperty, isInstanceProperty)
}/* debug [functions.gen.go/function]: protocol_addProperty */

// Adds a registered protocol to another protocol that is under construction.
//
// Added in macOS 10.7.
// Adds a registered protocol to another protocol that is under construction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_addProtocol(_:_:)
func protocol_addProtocol(proto unsafe.Pointer, addition unsafe.Pointer) {
	_protocol_addProtocol(proto, addition)
}/* debug [functions.gen.go/function]: protocol_addProtocol */

// Returns a Boolean value that indicates whether one protocol conforms to another protocol.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether one protocol conforms to another protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_conformsToProtocol(_:_:)
func protocol_conformsToProtocol(proto unsafe.Pointer, other unsafe.Pointer) bool {
	return _protocol_conformsToProtocol(proto, other)
}/* debug [functions.gen.go/function]: protocol_conformsToProtocol */

// Returns an array of method descriptions of methods meeting a given specification for a given protocol.
//
// Added in macOS 10.5.
// Returns an array of method descriptions of methods meeting a given specification for a given protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_copyMethodDescriptionList(_:_:_:_:)
func protocol_copyMethodDescriptionList(proto unsafe.Pointer, isRequiredMethod bool, isInstanceMethod bool, outCount unsafe.Pointer) unsafe.Pointer {
	return _protocol_copyMethodDescriptionList(proto, isRequiredMethod, isInstanceMethod, outCount)
}/* debug [functions.gen.go/function]: protocol_copyMethodDescriptionList */

// Returns an array of the properties declared by a protocol.
//
// Added in macOS 10.5.
// Returns an array of the properties declared by a protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_copyPropertyList(_:_:)
func protocol_copyPropertyList(proto unsafe.Pointer, outCount unsafe.Pointer) unsafe.Pointer {
	return _protocol_copyPropertyList(proto, outCount)
}/* debug [functions.gen.go/function]: protocol_copyPropertyList */

// protocol_copyPropertyList2 is a ObjectiveC function.
//
// Added in macOS 10.12.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_copyPropertyList2(_:_:_:_:)
func protocol_copyPropertyList2(proto unsafe.Pointer, outCount unsafe.Pointer, isRequiredProperty bool, isInstanceProperty bool) unsafe.Pointer {
	return _protocol_copyPropertyList2(proto, outCount, isRequiredProperty, isInstanceProperty)
}/* debug [functions.gen.go/function]: protocol_copyPropertyList2 */

// Returns an array of the protocols adopted by a protocol.
//
// Added in macOS 10.5.
// Returns an array of the protocols adopted by a protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_copyProtocolList(_:_:)
func protocol_copyProtocolList(proto unsafe.Pointer, outCount unsafe.Pointer) unsafe.Pointer {
	return _protocol_copyProtocolList(proto, outCount)
}/* debug [functions.gen.go/function]: protocol_copyProtocolList */

// Returns a method description structure for a specified method of a given protocol.
//
// Added in macOS 10.5.
// Returns a method description structure for a specified method of a given protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_getMethodDescription(_:_:_:_:)
func protocol_getMethodDescription(proto unsafe.Pointer, aSel objc.SEL, isRequiredMethod bool, isInstanceMethod bool) unsafe.Pointer {
	return _protocol_getMethodDescription(proto, aSel, isRequiredMethod, isInstanceMethod)
}/* debug [functions.gen.go/function]: protocol_getMethodDescription */

// Returns the name of a protocol.
//
// Added in macOS 10.5.
// Returns the name of a protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_getName(_:)
func protocol_getName(proto unsafe.Pointer) unsafe.Pointer {
	return _protocol_getName(proto)
}/* debug [functions.gen.go/function]: protocol_getName */

// Returns the specified property of a given protocol.
//
// Added in macOS 10.5.
// Returns the specified property of a given protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_getProperty(_:_:_:_:)
func protocol_getProperty(proto unsafe.Pointer, name unsafe.Pointer, isRequiredProperty bool, isInstanceProperty bool) objc_property_t {
	return _protocol_getProperty(proto, name, isRequiredProperty, isInstanceProperty)
}/* debug [functions.gen.go/function]: protocol_getProperty */

// Returns a Boolean value that indicates whether two protocols are equal.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether two protocols are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/protocol_isEqual(_:_:)
func protocol_isEqual(proto unsafe.Pointer, other unsafe.Pointer) bool {
	return _protocol_isEqual(proto, other)
}/* debug [functions.gen.go/function]: protocol_isEqual */

// Returns the name of the method specified by a given selector.
//
// Added in macOS 10.0.
// Returns the name of the method specified by a given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/sel_getName(_:)
func sel_getName(sel objc.SEL) unsafe.Pointer {
	return _sel_getName(sel)
}/* debug [functions.gen.go/function]: sel_getName */

// Registers a method name with the Objective-C runtime system.
//
// Added in macOS 10.0.
// Registers a method name with the Objective-C runtime system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/sel_getUid(_:)
func sel_getUid(str unsafe.Pointer) objc.SEL {
	return _sel_getUid(str)
}/* debug [functions.gen.go/function]: sel_getUid */

// Returns a Boolean value that indicates whether two selectors are equal.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether two selectors are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/sel_isEqual(_:_:)
func sel_isEqual(lhs objc.SEL, rhs objc.SEL) bool {
	return _sel_isEqual(lhs, rhs)
}/* debug [functions.gen.go/function]: sel_isEqual */

// Identifies a selector as being valid or invalid.
//
// Added in macOS 10.0.
// Identifies a selector as being valid or invalid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/sel_isMapped(_:)
func sel_isMapped(sel objc.SEL) bool {
	return _sel_isMapped(sel)
}/* debug [functions.gen.go/function]: sel_isMapped */

// Registers a method with the Objective-C runtime system, maps the method name to a selector, and returns the selector value.
//
// Added in macOS 10.0.
// Registers a method with the Objective-C runtime system, maps the method name to a selector, and returns the selector value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/sel_registerName(_:)
func sel_registerName(str unsafe.Pointer) objc.SEL {
	return _sel_registerName(str)
}/* debug [functions.gen.go/function]: sel_registerName */




