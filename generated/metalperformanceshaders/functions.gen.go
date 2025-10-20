// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// MetalPerformanceShaders Functions (2 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MPSGetImageType      func(unsafe.Pointer) unsafe.Pointer
	_MPSSupportsMTLDevice func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MPSGetImageType, lib, "MPSGetImageType")
	tryRegister(&_MPSSupportsMTLDevice, lib, "MPSSupportsMTLDevice")
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

// MPSGetImageType is a MetalPerformanceShaders function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSGetImageType(_:)
func MPSGetImageType(image unsafe.Pointer) unsafe.Pointer {
	return _MPSGetImageType(image)
}

// Determines whether the Metal Performance Shaders framework supports a Metal device. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSupportsMTLDevice(_:)
func MPSSupportsMTLDevice(device unsafe.Pointer) unsafe.Pointer {
	return _MPSSupportsMTLDevice(device)
}
