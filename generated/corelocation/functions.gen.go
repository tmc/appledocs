// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// CoreLocation Functions (2 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CLLocationCoordinate2DIsValid func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CLLocationCoordinate2DMake func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	tryRegister(&_CLLocationCoordinate2DIsValid, lib, "CLLocationCoordinate2DIsValid")
	tryRegister(&_CLLocationCoordinate2DMake, lib, "CLLocationCoordinate2DMake")
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


// Returns a Boolean value indicating whether the specified coordinate is valid. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationCoordinate2DIsValid(_:)
func CLLocationCoordinate2DIsValid(coord unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _CLLocationCoordinate2DIsValid(coord, p1)
	}


// Formats a latitude and longitude value into a coordinate data structure format. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationCoordinate2DMake(_:_:)
func CLLocationCoordinate2DMake(latitude unsafe.Pointer, longitude unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _CLLocationCoordinate2DMake(latitude, longitude, p2)
	}




