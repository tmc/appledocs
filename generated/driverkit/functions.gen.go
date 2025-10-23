// Code generated from Apple documentation for DriverKit. DO NOT EDIT.

package driverkit

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// DriverKit Functions (8 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_IODelay func(uint64)
	_IOParseBootArgString func(unsafe.Pointer, unsafe.Pointer, int) bool
	_OSArrayAppendValue func(unsafe.Pointer, unsafe.Pointer) bool
	_OSCollectionTypeName func(unsafe.Pointer) unsafe.Pointer
	_OSDataAppendBytes func(unsafe.Pointer, unsafe.Pointer, uintptr) bool
	_OSDataGetBytes func(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr) uintptr
	_OSDictionaryApply func(unsafe.Pointer, unsafe.Pointer) bool
	_mach_absolute_time func() uint64
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_IODelay, lib, "IODelay")
	tryRegister(&_IOParseBootArgString, lib, "IOParseBootArgString")
	tryRegister(&_OSArrayAppendValue, lib, "OSArrayAppendValue")
	tryRegister(&_OSCollectionTypeName, lib, "OSCollectionTypeName")
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



// Sleep the calling thread for a number of microseconds.

// Sleep the calling thread for a number of microseconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IODelay
func IODelay(us uint64) {
	_IODelay(us)
}

// Parses any boot arguments in the macOS kernel boot-args.

// Parses any boot arguments in the macOS kernel boot-args.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOParseBootArgString
func IOParseBootArgString(arg_string unsafe.Pointer, arg_ptr unsafe.Pointer, strlen int) bool {
	return _IOParseBootArgString(arg_string, arg_ptr, strlen)
}

// OSArrayAppendValue is a DriverKit function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSArrayAppendValue
func OSArrayAppendValue(obj unsafe.Pointer, value unsafe.Pointer) bool {
	return _OSArrayAppendValue(obj, value)
}

// OSCollectionTypeName is a DriverKit function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSCollectionTypeName
func OSCollectionTypeName(t unsafe.Pointer) unsafe.Pointer {
	return _OSCollectionTypeName(t)
}

// OSDataAppendBytes is a DriverKit function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDataAppendBytes
func OSDataAppendBytes(data unsafe.Pointer, bytes unsafe.Pointer, length uintptr) bool {
	return _OSDataAppendBytes(data, bytes, length)
}

// OSDataGetBytes is a DriverKit function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDataGetBytes
func OSDataGetBytes(obj unsafe.Pointer, buffer unsafe.Pointer, offset uintptr, length uintptr) uintptr {
	return _OSDataGetBytes(obj, buffer, offset, length)
}

// OSDictionaryApply is a DriverKit function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/OSDictionaryApply
func OSDictionaryApply(obj unsafe.Pointer, applier unsafe.Pointer) bool {
	return _OSDictionaryApply(obj, applier)
}

// Returns current value of a clock that increments monotonically in tick units (starting at an arbitrary point), this clock does not increment while the system is asleep.

// Returns current value of a clock that increments monotonically in tick units (starting at an arbitrary point), this clock does not increment while the system is asleep.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/mach_absolute_time
func mach_absolute_time() uint64 {
	return _mach_absolute_time()
}



