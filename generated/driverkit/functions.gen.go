// Code generated from Apple documentation for DriverKit. DO NOT EDIT.

package driverkit

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// DriverKit Functions (10 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_IOCallOnce func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IODelay func(unsafe.Pointer) unsafe.Pointer
	_IOLogv func(unsafe.Pointer, unsafe.Pointer) int
	_IOMallocZero func(unsafe.Pointer) unsafe.Pointer
	_IOParseBootArgString func(unsafe.Pointer, unsafe.Pointer, int) bool
	_IORWLockUnlock func(unsafe.Pointer) unsafe.Pointer
	_OSDataAppendBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_OSDataGetBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_OSDictionaryApply func(unsafe.Pointer, unsafe.Pointer) bool
	_mach_absolute_time func() unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_IOCallOnce, lib, "IOCallOnce")
	tryRegister(&_IODelay, lib, "IODelay")
	tryRegister(&_IOLogv, lib, "IOLogv")
	tryRegister(&_IOMallocZero, lib, "IOMallocZero")
	tryRegister(&_IOParseBootArgString, lib, "IOParseBootArgString")
	tryRegister(&_IORWLockUnlock, lib, "IORWLockUnlock")
	tryRegister(&_OSDataAppendBytes, lib, "OSDataAppendBytes")
	tryRegister(&_OSDataGetBytes, lib, "OSDataGetBytes")
	tryRegister(&_OSDictionaryApply, lib, "OSDictionaryApply")
	tryRegister(&_mach_absolute_time, lib, "mach_absolute_time")
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



// IOCallOnce is a DriverKit function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOCallOnce
func IOCallOnce(flag unsafe.Pointer, block unsafe.Pointer) {
	_IOCallOnce(flag, block)
	}


// Sleep the calling thread for a number of microseconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODelay
func IODelay(us unsafe.Pointer) {
	_IODelay(us)
	}


// IOLogv is a DriverKit function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOLogv
func IOLogv(format unsafe.Pointer, ap unsafe.Pointer) int {
	return _IOLogv(format, ap)
	}


// Allocates the specified amount of general-purpose memory and zero-initializes it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOMallocZero
func IOMallocZero(length unsafe.Pointer) unsafe.Pointer {
	return _IOMallocZero(length)
	}


// Parses any boot arguments in the macOS kernel boot-args. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOParseBootArgString
func IOParseBootArgString(arg_string unsafe.Pointer, arg_ptr unsafe.Pointer, strlen int) bool {
	return _IOParseBootArgString(arg_string, arg_ptr, strlen)
	}


// IORWLockUnlock is a DriverKit function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IORWLockUnlock
func IORWLockUnlock(lock unsafe.Pointer) {
	_IORWLockUnlock(lock)
	}


// OSDataAppendBytes is a DriverKit function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDataAppendBytes
func OSDataAppendBytes(data unsafe.Pointer, bytes unsafe.Pointer, length unsafe.Pointer) bool {
	return _OSDataAppendBytes(data, bytes, length)
	}


// OSDataGetBytes is a DriverKit function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDataGetBytes
func OSDataGetBytes(obj unsafe.Pointer, buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer {
	return _OSDataGetBytes(obj, buffer, offset, length)
	}


// OSDictionaryApply is a DriverKit function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionaryApply
func OSDictionaryApply(obj unsafe.Pointer, applier unsafe.Pointer) bool {
	return _OSDictionaryApply(obj, applier)
	}


// Returns current value of a clock that increments monotonically in tick units (starting at an arbitrary point), this clock does not increment while the system is asleep. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/mach_absolute_time
func mach_absolute_time() unsafe.Pointer {
	return _mach_absolute_time()
	}




