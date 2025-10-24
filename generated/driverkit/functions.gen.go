// Code generated from Apple documentation for DriverKit. DO NOT EDIT.

package driverkit

/* debug [functions.gen.go]: Generating 92 functions for DriverKit */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// DriverKit Functions (92 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_crc32 func(uint32, unsafe.Pointer, uintptr) uint32
	_IOCallOnce func(unsafe.Pointer, CallOnceBlock)
	_IODelay func(uint64)
	_IOFree func(unsafe.Pointer, uintptr)
	_IOLockAlloc func() unsafe.Pointer
	_IOLockAssert func(unsafe.Pointer, LockAssertState)
	_IOLockFree func(unsafe.Pointer)
	_IOLockLock func(unsafe.Pointer)
	_IOLockTryLock func(unsafe.Pointer) bool
	_IOLockUnlock func(unsafe.Pointer)
	_IOLog func(unsafe.Pointer) int
	_IOLogBuffer func(unsafe.Pointer, unsafe.Pointer, uintptr)
	_IOLogv func(unsafe.Pointer, unsafe.Pointer) int
	_IOMalloc func(uintptr) unsafe.Pointer
	_IOMallocTyped func(uintptr, unsafe.Pointer) unsafe.Pointer
	_IOMallocZero func(uintptr) unsafe.Pointer
	_IOMallocZeroTyped func(uintptr, unsafe.Pointer) unsafe.Pointer
	_IOParseBootArgNumber func(unsafe.Pointer, unsafe.Pointer, int) bool
	_IOParseBootArgString func(unsafe.Pointer, unsafe.Pointer, int) bool
	_IORecursiveConditionLockAlloc func() unsafe.Pointer
	_IORecursiveConditionLockFree func(unsafe.Pointer)
	_IORecursiveConditionLockHaveLock func(unsafe.Pointer) bool
	_IORecursiveConditionLockLock func(unsafe.Pointer)
	_IORecursiveConditionLockTryLock func(unsafe.Pointer) bool
	_IORecursiveConditionLockUnlock func(unsafe.Pointer)
	_IORecursiveLockAlloc func() unsafe.Pointer
	_IORecursiveLockFree func(unsafe.Pointer)
	_IORecursiveLockHaveLock func(unsafe.Pointer) bool
	_IORecursiveLockLock func(unsafe.Pointer)
	_IORecursiveLockTryLock func(unsafe.Pointer) bool
	_IORecursiveLockUnlock func(unsafe.Pointer)
	_IORPCMessageFromMach func(unsafe.Pointer, bool) unsafe.Pointer
	_IORWLockAlloc func() unsafe.Pointer
	_IORWLockFree func(unsafe.Pointer)
	_IORWLockRead func(unsafe.Pointer)
	_IORWLockUnlock func(unsafe.Pointer)
	_IORWLockWrite func(unsafe.Pointer)
	_IOSleep func(uint64)
	_IOThreadLocalStorageGet func(uint64) unsafe.Pointer
	_IOThreadLocalStorageKeyCreate func([]uint64) unsafe.Pointer
	_IOThreadLocalStorageKeyDelete func(uint64) unsafe.Pointer
	_IOThreadLocalStorageSet func(uint64, unsafe.Pointer) unsafe.Pointer
	_mach_absolute_time func() uint64
	_mach_continuous_time func() uint64
	_mach_timebase_info func(unsafe.Pointer) unsafe.Pointer
	_operator new func(uintptr, unsafe.Pointer) unsafe.Pointer
	_OSArrayAppendValue func(unsafe.Pointer, unsafe.Pointer) bool
	_OSArrayApply func(unsafe.Pointer, unsafe.Pointer) bool
	_OSArrayCreate func() unsafe.Pointer
	_OSArrayGetCount func(unsafe.Pointer) uint32
	_OSArrayGetStringValue func(unsafe.Pointer, uintptr) unsafe.Pointer
	_OSArrayGetUInt64Value func(unsafe.Pointer, uintptr) uint64
	_OSArrayGetValue func(unsafe.Pointer, uintptr) unsafe.Pointer
	_OSArrayReplaceValue func(unsafe.Pointer, uintptr, unsafe.Pointer) bool
	_OSArraySetStringValue func(unsafe.Pointer, uintptr, unsafe.Pointer)
	_OSArraySetUInt64Value func(unsafe.Pointer, uintptr, uint64)
	_OSArraySetValue func(unsafe.Pointer, uintptr, unsafe.Pointer) bool
	_OSCollectionsInitialize func()
	_OSCollectionTypeID func(unsafe.Pointer) unsafe.Pointer
	_OSCollectionTypeName func(unsafe.Pointer) unsafe.Pointer
	_OSCreateObjectFromSerialization func(unsafe.Pointer) unsafe.Pointer
	_OSCreateSerializationFromBytes func(unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_OSCreateSerializationFromObject func(unsafe.Pointer) unsafe.Pointer
	_OSDataAppendBytes func(unsafe.Pointer, unsafe.Pointer, uintptr) bool
	_OSDataCreate func(unsafe.Pointer, uintptr) unsafe.Pointer
	_OSDataGetBytes func(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr) uintptr
	_OSDataGetBytesPtr func(unsafe.Pointer, uintptr, uintptr) unsafe.Pointer
	_OSDataGetLength func(unsafe.Pointer) uintptr
	_OSDictionaryApply func(unsafe.Pointer, unsafe.Pointer) bool
	_OSDictionaryCreate func() unsafe.Pointer
	_OSDictionaryGetCount func(unsafe.Pointer) uint32
	_OSDictionaryGetStringValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OSDictionaryGetUInt64Value func(unsafe.Pointer, unsafe.Pointer) uint64
	_OSDictionaryGetValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OSDictionarySetStringValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_OSDictionarySetUInt64Value func(unsafe.Pointer, unsafe.Pointer, uint64)
	_OSDictionarySetValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_OSNumberCreateWithUInt64Value func(uint64) unsafe.Pointer
	_OSNumberGetUInt64Value func(unsafe.Pointer) uint64
	_OSObjectAllocate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OSObjectLog func(unsafe.Pointer)
	_OSObjectRelease func(unsafe.Pointer)
	_OSObjectRetain func(unsafe.Pointer)
	_OSReportWithBacktrace func(unsafe.Pointer)
	_OSSerializationGetBytes func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OSStringCreate func(unsafe.Pointer, uintptr) unsafe.Pointer
	_OSStringGetLength func(unsafe.Pointer) uintptr
	_OSStringGetStringPtr func(unsafe.Pointer) unsafe.Pointer
	_OSSynchronizeIO func()
	_panic func(unsafe.Pointer)
	_read_random func(unsafe.Pointer, uintptr)
	_swap func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_crc32, lib, "crc32")
	tryRegister(&_IOCallOnce, lib, "IOCallOnce")
	tryRegister(&_IODelay, lib, "IODelay")
	tryRegister(&_IOFree, lib, "IOFree")
	tryRegister(&_IOLockAlloc, lib, "IOLockAlloc")
	tryRegister(&_IOLockAssert, lib, "IOLockAssert")
	tryRegister(&_IOLockFree, lib, "IOLockFree")
	tryRegister(&_IOLockLock, lib, "IOLockLock")
	tryRegister(&_IOLockTryLock, lib, "IOLockTryLock")
	tryRegister(&_IOLockUnlock, lib, "IOLockUnlock")
	tryRegister(&_IOLog, lib, "IOLog")
	tryRegister(&_IOLogBuffer, lib, "IOLogBuffer")
	tryRegister(&_IOLogv, lib, "IOLogv")
	tryRegister(&_IOMalloc, lib, "IOMalloc")
	tryRegister(&_IOMallocTyped, lib, "IOMallocTyped")
	tryRegister(&_IOMallocZero, lib, "IOMallocZero")
	tryRegister(&_IOMallocZeroTyped, lib, "IOMallocZeroTyped")
	tryRegister(&_IOParseBootArgNumber, lib, "IOParseBootArgNumber")
	tryRegister(&_IOParseBootArgString, lib, "IOParseBootArgString")
	tryRegister(&_IORecursiveConditionLockAlloc, lib, "IORecursiveConditionLockAlloc")
	tryRegister(&_IORecursiveConditionLockFree, lib, "IORecursiveConditionLockFree")
	tryRegister(&_IORecursiveConditionLockHaveLock, lib, "IORecursiveConditionLockHaveLock")
	tryRegister(&_IORecursiveConditionLockLock, lib, "IORecursiveConditionLockLock")
	tryRegister(&_IORecursiveConditionLockTryLock, lib, "IORecursiveConditionLockTryLock")
	tryRegister(&_IORecursiveConditionLockUnlock, lib, "IORecursiveConditionLockUnlock")
	tryRegister(&_IORecursiveLockAlloc, lib, "IORecursiveLockAlloc")
	tryRegister(&_IORecursiveLockFree, lib, "IORecursiveLockFree")
	tryRegister(&_IORecursiveLockHaveLock, lib, "IORecursiveLockHaveLock")
	tryRegister(&_IORecursiveLockLock, lib, "IORecursiveLockLock")
	tryRegister(&_IORecursiveLockTryLock, lib, "IORecursiveLockTryLock")
	tryRegister(&_IORecursiveLockUnlock, lib, "IORecursiveLockUnlock")
	tryRegister(&_IORPCMessageFromMach, lib, "IORPCMessageFromMach")
	tryRegister(&_IORWLockAlloc, lib, "IORWLockAlloc")
	tryRegister(&_IORWLockFree, lib, "IORWLockFree")
	tryRegister(&_IORWLockRead, lib, "IORWLockRead")
	tryRegister(&_IORWLockUnlock, lib, "IORWLockUnlock")
	tryRegister(&_IORWLockWrite, lib, "IORWLockWrite")
	tryRegister(&_IOSleep, lib, "IOSleep")
	tryRegister(&_IOThreadLocalStorageGet, lib, "IOThreadLocalStorageGet")
	tryRegister(&_IOThreadLocalStorageKeyCreate, lib, "IOThreadLocalStorageKeyCreate")
	tryRegister(&_IOThreadLocalStorageKeyDelete, lib, "IOThreadLocalStorageKeyDelete")
	tryRegister(&_IOThreadLocalStorageSet, lib, "IOThreadLocalStorageSet")
	tryRegister(&_mach_absolute_time, lib, "mach_absolute_time")
	tryRegister(&_mach_continuous_time, lib, "mach_continuous_time")
	tryRegister(&_mach_timebase_info, lib, "mach_timebase_info")
	tryRegister(&_operator new, lib, "operator new")
	tryRegister(&_OSArrayAppendValue, lib, "OSArrayAppendValue")
	tryRegister(&_OSArrayApply, lib, "OSArrayApply")
	tryRegister(&_OSArrayCreate, lib, "OSArrayCreate")
	tryRegister(&_OSArrayGetCount, lib, "OSArrayGetCount")
	tryRegister(&_OSArrayGetStringValue, lib, "OSArrayGetStringValue")
	tryRegister(&_OSArrayGetUInt64Value, lib, "OSArrayGetUInt64Value")
	tryRegister(&_OSArrayGetValue, lib, "OSArrayGetValue")
	tryRegister(&_OSArrayReplaceValue, lib, "OSArrayReplaceValue")
	tryRegister(&_OSArraySetStringValue, lib, "OSArraySetStringValue")
	tryRegister(&_OSArraySetUInt64Value, lib, "OSArraySetUInt64Value")
	tryRegister(&_OSArraySetValue, lib, "OSArraySetValue")
	tryRegister(&_OSCollectionsInitialize, lib, "OSCollectionsInitialize")
	tryRegister(&_OSCollectionTypeID, lib, "OSCollectionTypeID")
	tryRegister(&_OSCollectionTypeName, lib, "OSCollectionTypeName")
	tryRegister(&_OSCreateObjectFromSerialization, lib, "OSCreateObjectFromSerialization")
	tryRegister(&_OSCreateSerializationFromBytes, lib, "OSCreateSerializationFromBytes")
	tryRegister(&_OSCreateSerializationFromObject, lib, "OSCreateSerializationFromObject")
	tryRegister(&_OSDataAppendBytes, lib, "OSDataAppendBytes")
	tryRegister(&_OSDataCreate, lib, "OSDataCreate")
	tryRegister(&_OSDataGetBytes, lib, "OSDataGetBytes")
	tryRegister(&_OSDataGetBytesPtr, lib, "OSDataGetBytesPtr")
	tryRegister(&_OSDataGetLength, lib, "OSDataGetLength")
	tryRegister(&_OSDictionaryApply, lib, "OSDictionaryApply")
	tryRegister(&_OSDictionaryCreate, lib, "OSDictionaryCreate")
	tryRegister(&_OSDictionaryGetCount, lib, "OSDictionaryGetCount")
	tryRegister(&_OSDictionaryGetStringValue, lib, "OSDictionaryGetStringValue")
	tryRegister(&_OSDictionaryGetUInt64Value, lib, "OSDictionaryGetUInt64Value")
	tryRegister(&_OSDictionaryGetValue, lib, "OSDictionaryGetValue")
	tryRegister(&_OSDictionarySetStringValue, lib, "OSDictionarySetStringValue")
	tryRegister(&_OSDictionarySetUInt64Value, lib, "OSDictionarySetUInt64Value")
	tryRegister(&_OSDictionarySetValue, lib, "OSDictionarySetValue")
	tryRegister(&_OSNumberCreateWithUInt64Value, lib, "OSNumberCreateWithUInt64Value")
	tryRegister(&_OSNumberGetUInt64Value, lib, "OSNumberGetUInt64Value")
	tryRegister(&_OSObjectAllocate, lib, "OSObjectAllocate")
	tryRegister(&_OSObjectLog, lib, "OSObjectLog")
	tryRegister(&_OSObjectRelease, lib, "OSObjectRelease")
	tryRegister(&_OSObjectRetain, lib, "OSObjectRetain")
	tryRegister(&_OSReportWithBacktrace, lib, "OSReportWithBacktrace")
	tryRegister(&_OSSerializationGetBytes, lib, "OSSerializationGetBytes")
	tryRegister(&_OSStringCreate, lib, "OSStringCreate")
	tryRegister(&_OSStringGetLength, lib, "OSStringGetLength")
	tryRegister(&_OSStringGetStringPtr, lib, "OSStringGetStringPtr")
	tryRegister(&_OSSynchronizeIO, lib, "OSSynchronizeIO")
	tryRegister(&_panic, lib, "panic")
	tryRegister(&_read_random, lib, "read_random")
	tryRegister(&_swap, lib, "swap")
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



// crc32 is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/crc32
func crc32(crc uint32, buf unsafe.Pointer, size uintptr) uint32 {
	return _crc32(crc, buf, size)
}/* debug [functions.gen.go/function]: crc32 */

// IOCallOnce is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOCallOnce
func IOCallOnce(flag unsafe.Pointer, block CallOnceBlock) {
	_IOCallOnce(flag, block)
}/* debug [functions.gen.go/function]: IOCallOnce */

// Sleep the calling thread for a number of microseconds.
//
// Added in macOS .
// Sleep the calling thread for a number of microseconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODelay
func IODelay(us uint64) {
	_IODelay(us)
}/* debug [functions.gen.go/function]: IODelay */

// Frees a memory block that contains general-purpose memory.
//
// Added in macOS .
// Frees a memory block that contains general-purpose memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOFree
func IOFree(address unsafe.Pointer, length uintptr) {
	_IOFree(address, length)
}/* debug [functions.gen.go/function]: IOFree */

// IOLockAlloc is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOLockAlloc
func IOLockAlloc() unsafe.Pointer {
	return _IOLockAlloc()
}/* debug [functions.gen.go/function]: IOLockAlloc */

// IOLockAssert is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOLockAssert
func IOLockAssert(lock unsafe.Pointer, type_ LockAssertState) {
	_IOLockAssert(lock, type_)
}/* debug [functions.gen.go/function]: IOLockAssert */

// IOLockFree is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOLockFree
func IOLockFree(lock unsafe.Pointer) {
	_IOLockFree(lock)
}/* debug [functions.gen.go/function]: IOLockFree */

// IOLockLock is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOLockLock
func IOLockLock(lock unsafe.Pointer) {
	_IOLockLock(lock)
}/* debug [functions.gen.go/function]: IOLockLock */

// IOLockTryLock is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOLockTryLock
func IOLockTryLock(lock unsafe.Pointer) bool {
	return _IOLockTryLock(lock)
}/* debug [functions.gen.go/function]: IOLockTryLock */

// IOLockUnlock is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOLockUnlock
func IOLockUnlock(lock unsafe.Pointer) {
	_IOLockUnlock(lock)
}/* debug [functions.gen.go/function]: IOLockUnlock */

// IOLog is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOLog
func IOLog(format unsafe.Pointer) int {
	return _IOLog(format)
}/* debug [functions.gen.go/function]: IOLog */

// IOLogBuffer is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOLogBuffer
func IOLogBuffer(title unsafe.Pointer, buffer unsafe.Pointer, size uintptr) {
	_IOLogBuffer(title, buffer, size)
}/* debug [functions.gen.go/function]: IOLogBuffer */

// IOLogv is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOLogv
func IOLogv(format unsafe.Pointer, ap unsafe.Pointer) int {
	return _IOLogv(format, ap)
}/* debug [functions.gen.go/function]: IOLogv */

// Allocates the specified amount of general-purpose memory.
//
// Added in macOS .
// Allocates the specified amount of general-purpose memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMalloc
func IOMalloc(length uintptr) unsafe.Pointer {
	return _IOMalloc(length)
}/* debug [functions.gen.go/function]: IOMalloc */

// IOMallocTyped is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMallocTyped
func IOMallocTyped(length uintptr, type_id unsafe.Pointer) unsafe.Pointer {
	return _IOMallocTyped(length, type_id)
}/* debug [functions.gen.go/function]: IOMallocTyped */

// Allocates the specified amount of general-purpose memory and zero-initializes it.
//
// Added in macOS .
// Allocates the specified amount of general-purpose memory and zero-initializes it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMallocZero
func IOMallocZero(length uintptr) unsafe.Pointer {
	return _IOMallocZero(length)
}/* debug [functions.gen.go/function]: IOMallocZero */

// IOMallocZeroTyped is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMallocZeroTyped
func IOMallocZeroTyped(length uintptr, type_id unsafe.Pointer) unsafe.Pointer {
	return _IOMallocZeroTyped(length, type_id)
}/* debug [functions.gen.go/function]: IOMallocZeroTyped */

// Parses any boot arguments in the macOS kernel boot-args.
//
// Added in macOS .
// Parses any boot arguments in the macOS kernel boot-args.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOParseBootArgNumber
func IOParseBootArgNumber(arg_string unsafe.Pointer, arg_ptr unsafe.Pointer, max_len int) bool {
	return _IOParseBootArgNumber(arg_string, arg_ptr, max_len)
}/* debug [functions.gen.go/function]: IOParseBootArgNumber */

// Parses any boot arguments in the macOS kernel boot-args.
//
// Added in macOS .
// Parses any boot arguments in the macOS kernel boot-args.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOParseBootArgString
func IOParseBootArgString(arg_string unsafe.Pointer, arg_ptr unsafe.Pointer, strlen int) bool {
	return _IOParseBootArgString(arg_string, arg_ptr, strlen)
}/* debug [functions.gen.go/function]: IOParseBootArgString */

// IORecursiveConditionLockAlloc is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORecursiveConditionLockAlloc
func IORecursiveConditionLockAlloc() unsafe.Pointer {
	return _IORecursiveConditionLockAlloc()
}/* debug [functions.gen.go/function]: IORecursiveConditionLockAlloc */

// IORecursiveConditionLockFree is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORecursiveConditionLockFree
func IORecursiveConditionLockFree(lock unsafe.Pointer) {
	_IORecursiveConditionLockFree(lock)
}/* debug [functions.gen.go/function]: IORecursiveConditionLockFree */

// IORecursiveConditionLockHaveLock is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORecursiveConditionLockHaveLock
func IORecursiveConditionLockHaveLock(lock unsafe.Pointer) bool {
	return _IORecursiveConditionLockHaveLock(lock)
}/* debug [functions.gen.go/function]: IORecursiveConditionLockHaveLock */

// IORecursiveConditionLockLock is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORecursiveConditionLockLock
func IORecursiveConditionLockLock(lock unsafe.Pointer) {
	_IORecursiveConditionLockLock(lock)
}/* debug [functions.gen.go/function]: IORecursiveConditionLockLock */

// IORecursiveConditionLockTryLock is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORecursiveConditionLockTryLock
func IORecursiveConditionLockTryLock(lock unsafe.Pointer) bool {
	return _IORecursiveConditionLockTryLock(lock)
}/* debug [functions.gen.go/function]: IORecursiveConditionLockTryLock */

// IORecursiveConditionLockUnlock is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORecursiveConditionLockUnlock
func IORecursiveConditionLockUnlock(lock unsafe.Pointer) {
	_IORecursiveConditionLockUnlock(lock)
}/* debug [functions.gen.go/function]: IORecursiveConditionLockUnlock */

// IORecursiveLockAlloc is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORecursiveLockAlloc
func IORecursiveLockAlloc() unsafe.Pointer {
	return _IORecursiveLockAlloc()
}/* debug [functions.gen.go/function]: IORecursiveLockAlloc */

// IORecursiveLockFree is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORecursiveLockFree
func IORecursiveLockFree(lock unsafe.Pointer) {
	_IORecursiveLockFree(lock)
}/* debug [functions.gen.go/function]: IORecursiveLockFree */

// IORecursiveLockHaveLock is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORecursiveLockHaveLock
func IORecursiveLockHaveLock(lock unsafe.Pointer) bool {
	return _IORecursiveLockHaveLock(lock)
}/* debug [functions.gen.go/function]: IORecursiveLockHaveLock */

// IORecursiveLockLock is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORecursiveLockLock
func IORecursiveLockLock(lock unsafe.Pointer) {
	_IORecursiveLockLock(lock)
}/* debug [functions.gen.go/function]: IORecursiveLockLock */

// IORecursiveLockTryLock is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORecursiveLockTryLock
func IORecursiveLockTryLock(lock unsafe.Pointer) bool {
	return _IORecursiveLockTryLock(lock)
}/* debug [functions.gen.go/function]: IORecursiveLockTryLock */

// IORecursiveLockUnlock is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORecursiveLockUnlock
func IORecursiveLockUnlock(lock unsafe.Pointer) {
	_IORecursiveLockUnlock(lock)
}/* debug [functions.gen.go/function]: IORecursiveLockUnlock */

// IORPCMessageFromMach is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORPCMessageFromMach
func IORPCMessageFromMach(msg unsafe.Pointer, reply bool) unsafe.Pointer {
	return _IORPCMessageFromMach(msg, reply)
}/* debug [functions.gen.go/function]: IORPCMessageFromMach */

// IORWLockAlloc is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORWLockAlloc
func IORWLockAlloc() unsafe.Pointer {
	return _IORWLockAlloc()
}/* debug [functions.gen.go/function]: IORWLockAlloc */

// IORWLockFree is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORWLockFree
func IORWLockFree(lock unsafe.Pointer) {
	_IORWLockFree(lock)
}/* debug [functions.gen.go/function]: IORWLockFree */

// IORWLockRead is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORWLockRead
func IORWLockRead(lock unsafe.Pointer) {
	_IORWLockRead(lock)
}/* debug [functions.gen.go/function]: IORWLockRead */

// IORWLockUnlock is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORWLockUnlock
func IORWLockUnlock(lock unsafe.Pointer) {
	_IORWLockUnlock(lock)
}/* debug [functions.gen.go/function]: IORWLockUnlock */

// IORWLockWrite is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORWLockWrite
func IORWLockWrite(lock unsafe.Pointer) {
	_IORWLockWrite(lock)
}/* debug [functions.gen.go/function]: IORWLockWrite */

// Sleep the calling thread for a number of milliseconds.
//
// Added in macOS .
// Sleep the calling thread for a number of milliseconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOSleep
func IOSleep(ms uint64) {
	_IOSleep(ms)
}/* debug [functions.gen.go/function]: IOSleep */

// IOThreadLocalStorageGet is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOThreadLocalStorageGet
func IOThreadLocalStorageGet(key uint64) unsafe.Pointer {
	return _IOThreadLocalStorageGet(key)
}/* debug [functions.gen.go/function]: IOThreadLocalStorageGet */

// IOThreadLocalStorageKeyCreate is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOThreadLocalStorageKeyCreate
func IOThreadLocalStorageKeyCreate(key []uint64) unsafe.Pointer {
	return _IOThreadLocalStorageKeyCreate(key)
}/* debug [functions.gen.go/function]: IOThreadLocalStorageKeyCreate */

// IOThreadLocalStorageKeyDelete is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOThreadLocalStorageKeyDelete
func IOThreadLocalStorageKeyDelete(key uint64) unsafe.Pointer {
	return _IOThreadLocalStorageKeyDelete(key)
}/* debug [functions.gen.go/function]: IOThreadLocalStorageKeyDelete */

// IOThreadLocalStorageSet is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOThreadLocalStorageSet
func IOThreadLocalStorageSet(key uint64, value unsafe.Pointer) unsafe.Pointer {
	return _IOThreadLocalStorageSet(key, value)
}/* debug [functions.gen.go/function]: IOThreadLocalStorageSet */

// Returns current value of a clock that increments monotonically in tick units (starting at an arbitrary point), this clock does not increment while the system is asleep.
//
// Added in macOS .
// Returns current value of a clock that increments monotonically in tick units (starting at an arbitrary point), this clock does not increment while the system is asleep.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/mach_absolute_time
func mach_absolute_time() uint64 {
	return _mach_absolute_time()
}/* debug [functions.gen.go/function]: mach_absolute_time */

// Returns current value of a clock that increments monotonically in tick units (starting at an arbitrary point), including while the system is asleep.
//
// Added in macOS .
// Returns current value of a clock that increments monotonically in tick units (starting at an arbitrary point), including while the system is asleep.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/mach_continuous_time
func mach_continuous_time() uint64 {
	return _mach_continuous_time()
}/* debug [functions.gen.go/function]: mach_continuous_time */

// Returns fraction to multiply a value in mach tick units with to convert it to nanoseconds.
//
// Added in macOS .
// Returns fraction to multiply a value in mach tick units with to convert it to nanoseconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/mach_timebase_info-c.func
func mach_timebase_info(info unsafe.Pointer) unsafe.Pointer {
	return _mach_timebase_info(info)
}/* debug [functions.gen.go/function]: mach_timebase_info */

// operator new is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/operator_new
func operator new(p0 uintptr, p1 unsafe.Pointer) unsafe.Pointer {
	return _operator new(p0, p1)
}/* debug [functions.gen.go/function]: operator new */

// OSArrayAppendValue is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArrayAppendValue
func OSArrayAppendValue(obj unsafe.Pointer, value unsafe.Pointer) bool {
	return _OSArrayAppendValue(obj, value)
}/* debug [functions.gen.go/function]: OSArrayAppendValue */

// OSArrayApply is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArrayApply
func OSArrayApply(obj unsafe.Pointer, applier unsafe.Pointer) bool {
	return _OSArrayApply(obj, applier)
}/* debug [functions.gen.go/function]: OSArrayApply */

// OSArrayCreate is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArrayCreate
func OSArrayCreate() unsafe.Pointer {
	return _OSArrayCreate()
}/* debug [functions.gen.go/function]: OSArrayCreate */

// OSArrayGetCount is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArrayGetCount
func OSArrayGetCount(obj unsafe.Pointer) uint32 {
	return _OSArrayGetCount(obj)
}/* debug [functions.gen.go/function]: OSArrayGetCount */

// OSArrayGetStringValue is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArrayGetStringValue
func OSArrayGetStringValue(obj unsafe.Pointer, index uintptr) unsafe.Pointer {
	return _OSArrayGetStringValue(obj, index)
}/* debug [functions.gen.go/function]: OSArrayGetStringValue */

// OSArrayGetUInt64Value is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArrayGetUInt64Value
func OSArrayGetUInt64Value(obj unsafe.Pointer, index uintptr) uint64 {
	return _OSArrayGetUInt64Value(obj, index)
}/* debug [functions.gen.go/function]: OSArrayGetUInt64Value */

// OSArrayGetValue is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArrayGetValue
func OSArrayGetValue(obj unsafe.Pointer, index uintptr) unsafe.Pointer {
	return _OSArrayGetValue(obj, index)
}/* debug [functions.gen.go/function]: OSArrayGetValue */

// OSArrayReplaceValue is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArrayReplaceValue
func OSArrayReplaceValue(obj unsafe.Pointer, index uintptr, value unsafe.Pointer) bool {
	return _OSArrayReplaceValue(obj, index, value)
}/* debug [functions.gen.go/function]: OSArrayReplaceValue */

// OSArraySetStringValue is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArraySetStringValue
func OSArraySetStringValue(obj unsafe.Pointer, index uintptr, value unsafe.Pointer) {
	_OSArraySetStringValue(obj, index, value)
}/* debug [functions.gen.go/function]: OSArraySetStringValue */

// OSArraySetUInt64Value is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArraySetUInt64Value
func OSArraySetUInt64Value(obj unsafe.Pointer, index uintptr, value uint64) {
	_OSArraySetUInt64Value(obj, index, value)
}/* debug [functions.gen.go/function]: OSArraySetUInt64Value */

// OSArraySetValue is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArraySetValue
func OSArraySetValue(obj unsafe.Pointer, index uintptr, value unsafe.Pointer) bool {
	return _OSArraySetValue(obj, index, value)
}/* debug [functions.gen.go/function]: OSArraySetValue */

// OSCollectionsInitialize is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSCollectionsInitialize
func OSCollectionsInitialize() {
	_OSCollectionsInitialize()
}/* debug [functions.gen.go/function]: OSCollectionsInitialize */

// OSCollectionTypeID is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSCollectionTypeID
func OSCollectionTypeID(obj unsafe.Pointer) unsafe.Pointer {
	return _OSCollectionTypeID(obj)
}/* debug [functions.gen.go/function]: OSCollectionTypeID */

// OSCollectionTypeName is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSCollectionTypeName
func OSCollectionTypeName(t unsafe.Pointer) unsafe.Pointer {
	return _OSCollectionTypeName(t)
}/* debug [functions.gen.go/function]: OSCollectionTypeName */

// OSCreateObjectFromSerialization is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSCreateObjectFromSerialization
func OSCreateObjectFromSerialization(serial unsafe.Pointer) unsafe.Pointer {
	return _OSCreateObjectFromSerialization(serial)
}/* debug [functions.gen.go/function]: OSCreateObjectFromSerialization */

// OSCreateSerializationFromBytes is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSCreateSerializationFromBytes
func OSCreateSerializationFromBytes(bytes unsafe.Pointer, length uintptr, freeBuffer unsafe.Pointer) unsafe.Pointer {
	return _OSCreateSerializationFromBytes(bytes, length, freeBuffer)
}/* debug [functions.gen.go/function]: OSCreateSerializationFromBytes */

// OSCreateSerializationFromObject is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSCreateSerializationFromObject
func OSCreateSerializationFromObject(obj unsafe.Pointer) unsafe.Pointer {
	return _OSCreateSerializationFromObject(obj)
}/* debug [functions.gen.go/function]: OSCreateSerializationFromObject */

// OSDataAppendBytes is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDataAppendBytes
func OSDataAppendBytes(data unsafe.Pointer, bytes unsafe.Pointer, length uintptr) bool {
	return _OSDataAppendBytes(data, bytes, length)
}/* debug [functions.gen.go/function]: OSDataAppendBytes */

// OSDataCreate is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDataCreate
func OSDataCreate(bytes unsafe.Pointer, length uintptr) unsafe.Pointer {
	return _OSDataCreate(bytes, length)
}/* debug [functions.gen.go/function]: OSDataCreate */

// OSDataGetBytes is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDataGetBytes
func OSDataGetBytes(obj unsafe.Pointer, buffer unsafe.Pointer, offset uintptr, length uintptr) uintptr {
	return _OSDataGetBytes(obj, buffer, offset, length)
}/* debug [functions.gen.go/function]: OSDataGetBytes */

// OSDataGetBytesPtr is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDataGetBytesPtr
func OSDataGetBytesPtr(obj unsafe.Pointer, offset uintptr, length uintptr) unsafe.Pointer {
	return _OSDataGetBytesPtr(obj, offset, length)
}/* debug [functions.gen.go/function]: OSDataGetBytesPtr */

// OSDataGetLength is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDataGetLength
func OSDataGetLength(obj unsafe.Pointer) uintptr {
	return _OSDataGetLength(obj)
}/* debug [functions.gen.go/function]: OSDataGetLength */

// OSDictionaryApply is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionaryApply
func OSDictionaryApply(obj unsafe.Pointer, applier unsafe.Pointer) bool {
	return _OSDictionaryApply(obj, applier)
}/* debug [functions.gen.go/function]: OSDictionaryApply */

// OSDictionaryCreate is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionaryCreate
func OSDictionaryCreate() unsafe.Pointer {
	return _OSDictionaryCreate()
}/* debug [functions.gen.go/function]: OSDictionaryCreate */

// OSDictionaryGetCount is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionaryGetCount
func OSDictionaryGetCount(obj unsafe.Pointer) uint32 {
	return _OSDictionaryGetCount(obj)
}/* debug [functions.gen.go/function]: OSDictionaryGetCount */

// OSDictionaryGetStringValue is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionaryGetStringValue
func OSDictionaryGetStringValue(obj unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _OSDictionaryGetStringValue(obj, key)
}/* debug [functions.gen.go/function]: OSDictionaryGetStringValue */

// OSDictionaryGetUInt64Value is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionaryGetUInt64Value
func OSDictionaryGetUInt64Value(obj unsafe.Pointer, key unsafe.Pointer) uint64 {
	return _OSDictionaryGetUInt64Value(obj, key)
}/* debug [functions.gen.go/function]: OSDictionaryGetUInt64Value */

// OSDictionaryGetValue is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionaryGetValue
func OSDictionaryGetValue(obj unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _OSDictionaryGetValue(obj, key)
}/* debug [functions.gen.go/function]: OSDictionaryGetValue */

// OSDictionarySetStringValue is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionarySetStringValue
func OSDictionarySetStringValue(obj unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_OSDictionarySetStringValue(obj, key, value)
}/* debug [functions.gen.go/function]: OSDictionarySetStringValue */

// OSDictionarySetUInt64Value is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionarySetUInt64Value
func OSDictionarySetUInt64Value(obj unsafe.Pointer, key unsafe.Pointer, value uint64) {
	_OSDictionarySetUInt64Value(obj, key, value)
}/* debug [functions.gen.go/function]: OSDictionarySetUInt64Value */

// OSDictionarySetValue is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionarySetValue
func OSDictionarySetValue(obj unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _OSDictionarySetValue(obj, key, value)
}/* debug [functions.gen.go/function]: OSDictionarySetValue */

// OSNumberCreateWithUInt64Value is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSNumberCreateWithUInt64Value
func OSNumberCreateWithUInt64Value(value uint64) unsafe.Pointer {
	return _OSNumberCreateWithUInt64Value(value)
}/* debug [functions.gen.go/function]: OSNumberCreateWithUInt64Value */

// OSNumberGetUInt64Value is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSNumberGetUInt64Value
func OSNumberGetUInt64Value(obj unsafe.Pointer) uint64 {
	return _OSNumberGetUInt64Value(obj)
}/* debug [functions.gen.go/function]: OSNumberGetUInt64Value */

// Helper function for OSTypeAlloc(). Not to be called directly.
//
// Added in macOS .
// Helper function for OSTypeAlloc(). Not to be called directly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSObjectAllocate
func OSObjectAllocate(meta unsafe.Pointer, pObject unsafe.Pointer) unsafe.Pointer {
	return _OSObjectAllocate(meta, pObject)
}/* debug [functions.gen.go/function]: OSObjectAllocate */

// OSObjectLog is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSObjectLog
func OSObjectLog(obj unsafe.Pointer) {
	_OSObjectLog(obj)
}/* debug [functions.gen.go/function]: OSObjectLog */

// OSObjectRelease is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSObjectRelease
func OSObjectRelease(container unsafe.Pointer) {
	_OSObjectRelease(container)
}/* debug [functions.gen.go/function]: OSObjectRelease */

// OSObjectRetain is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSObjectRetain
func OSObjectRetain(container unsafe.Pointer) {
	_OSObjectRetain(container)
}/* debug [functions.gen.go/function]: OSObjectRetain */

// Generates a backtrace and message for debugging.
//
// Added in macOS .
// Generates a backtrace and message for debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSReportWithBacktrace
func OSReportWithBacktrace(str unsafe.Pointer) {
	_OSReportWithBacktrace(str)
}/* debug [functions.gen.go/function]: OSReportWithBacktrace */

// OSSerializationGetBytes is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSerializationGetBytes
func OSSerializationGetBytes(serial unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer {
	return _OSSerializationGetBytes(serial, length)
}/* debug [functions.gen.go/function]: OSSerializationGetBytes */

// OSStringCreate is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSStringCreate
func OSStringCreate(cstring unsafe.Pointer, length uintptr) unsafe.Pointer {
	return _OSStringCreate(cstring, length)
}/* debug [functions.gen.go/function]: OSStringCreate */

// OSStringGetLength is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSStringGetLength
func OSStringGetLength(string_ unsafe.Pointer) uintptr {
	return _OSStringGetLength(string_)
}/* debug [functions.gen.go/function]: OSStringGetLength */

// OSStringGetStringPtr is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSStringGetStringPtr
func OSStringGetStringPtr(obj unsafe.Pointer) unsafe.Pointer {
	return _OSStringGetStringPtr(obj)
}/* debug [functions.gen.go/function]: OSStringGetStringPtr */

// Performs an instruction on Intel-based Mac computers.
//
// Added in macOS .
// Performs an instruction on Intel-based Mac computers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSSynchronizeIO
func OSSynchronizeIO() {
	_OSSynchronizeIO()
}/* debug [functions.gen.go/function]: OSSynchronizeIO */

// panic is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/panic
func panic(string_ unsafe.Pointer) {
	_panic(string_)
}/* debug [functions.gen.go/function]: panic */

// read_random is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/read_random
func read_random(buffer unsafe.Pointer, numBytes uintptr) {
	_read_random(buffer, numBytes)
}/* debug [functions.gen.go/function]: read_random */

// swap is a DriverKit function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/swap
func swap(a unsafe.Pointer, b unsafe.Pointer) unsafe.Pointer {
	return _swap(a, b)
}/* debug [functions.gen.go/function]: swap */




