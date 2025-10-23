// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// CoreMediaIO Functions (1 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CMIOStreamClockConvertHostTimeToDeviceTime func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CMIOStreamClockConvertHostTimeToDeviceTime, lib, "CMIOStreamClockConvertHostTimeToDeviceTime")
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



// CMIOStreamClockConvertHostTimeToDeviceTime is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamClockConvertHostTimeToDeviceTime(_:_:)
func CMIOStreamClockConvertHostTimeToDeviceTime(hostTime unsafe.Pointer, clock unsafe.Pointer) unsafe.Pointer {
	return _CMIOStreamClockConvertHostTimeToDeviceTime(hostTime, clock)
}



