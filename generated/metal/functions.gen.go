// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Metal Functions (2 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MTLCopyAllDevices func() unsafe.Pointer
	_MTLCreateSystemDefaultDevice func() unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MTLCopyAllDevices, lib, "MTLCopyAllDevices")
	tryRegister(&_MTLCreateSystemDefaultDevice, lib, "MTLCreateSystemDefaultDevice")
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



// Returns an array of all the Metal device instances in the system. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCopyAllDevices()
func MTLCopyAllDevices() unsafe.Pointer {
	return _MTLCopyAllDevices()
	}


// Returns the device instance Metal selects as the default. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCreateSystemDefaultDevice()
func MTLCreateSystemDefaultDevice() unsafe.Pointer {
	return _MTLCreateSystemDefaultDevice()
	}




