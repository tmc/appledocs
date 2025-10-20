// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// FSKit Functions (3 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_fs_errorForCocoaError func(unsafe.Pointer) unsafe.Pointer
	_fs_errorForMachError func(unsafe.Pointer) unsafe.Pointer
	_fs_errorForPOSIXError func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_fs_errorForCocoaError, lib, "fs_errorForCocoaError")
	tryRegister(&_fs_errorForMachError, lib, "fs_errorForMachError")
	tryRegister(&_fs_errorForPOSIXError, lib, "fs_errorForPOSIXError")
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



// Creates an error object for the given Cocoa error code. [Full Topic]
//
// Added in macOS 15.4.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/fs_errorForCocoaError(_:)
func fs_errorForCocoaError(errorCode unsafe.Pointer) unsafe.Pointer {
	return _fs_errorForCocoaError(errorCode)
	}


// Creates an error object for the given Mach error code. [Full Topic]
//
// Added in macOS 15.4.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/fs_errorForMachError(_:)
func fs_errorForMachError(errorCode unsafe.Pointer) unsafe.Pointer {
	return _fs_errorForMachError(errorCode)
	}


// Creates an error object for the given POSIX error code. [Full Topic]
//
// Added in macOS 15.4.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/fs_errorForPOSIXError(_:)
func fs_errorForPOSIXError(p0 unsafe.Pointer) unsafe.Pointer {
	return _fs_errorForPOSIXError(p0)
	}




