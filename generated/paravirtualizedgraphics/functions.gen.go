// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// ParavirtualizedGraphics Functions (4 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_PGCopyOptionROMURL func() unsafe.Pointer
	_PGCreateDeviceWithDescriptor func(unsafe.Pointer) unsafe.Pointer
	_PGMaxDisplayPortCount func() uint32
	_PGNewDeviceWithDescriptor func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_PGCopyOptionROMURL, lib, "PGCopyOptionROMURL")
	tryRegister(&_PGCreateDeviceWithDescriptor, lib, "PGCreateDeviceWithDescriptor")
	tryRegister(&_PGMaxDisplayPortCount, lib, "PGMaxDisplayPortCount")
	tryRegister(&_PGNewDeviceWithDescriptor, lib, "PGNewDeviceWithDescriptor")
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



// Copies the URL of the ROM image to use on the guest graphics device.
//
// Added in macOS 11.0.
// Copies the URL of the ROM image to use on the guest graphics device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGCopyOptionROMURL()
func PGCopyOptionROMURL() unsafe.Pointer {
	return _PGCopyOptionROMURL()
}

// PGCreateDeviceWithDescriptor is a ParavirtualizedGraphics function.
//
// Added in macOS 15.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGCreateDeviceWithDescriptor(_:)
func PGCreateDeviceWithDescriptor(descriptor unsafe.Pointer) unsafe.Pointer {
	return _PGCreateDeviceWithDescriptor(descriptor)
}

// PGMaxDisplayPortCount is a ParavirtualizedGraphics function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGMaxDisplayPortCount()
func PGMaxDisplayPortCount() uint32 {
	return _PGMaxDisplayPortCount()
}

// Creates a new paravirtualized graphics device.
//
// Added in macOS 11.0.
// Creates a new paravirtualized graphics device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGNewDeviceWithDescriptor(_:)
func PGNewDeviceWithDescriptor(descriptor unsafe.Pointer) unsafe.Pointer {
	return _PGNewDeviceWithDescriptor(descriptor)
}



