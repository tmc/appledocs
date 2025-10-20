// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// CoreML Functions (1 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MLAllComputeDevices func() unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MLAllComputeDevices, lib, "MLAllComputeDevices")
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



// Returns an array that contains all of the compute devices that are accessible. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLAllComputeDevices
func MLAllComputeDevices() unsafe.Pointer {
	return _MLAllComputeDevices()
	}




