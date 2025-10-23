// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// IOBluetoothUI Functions (2 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_IOBluetoothGetDeviceSelectorController func() unsafe.Pointer
	_IOBluetoothGetPairingController func() unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_IOBluetoothGetDeviceSelectorController, lib, "IOBluetoothGetDeviceSelectorController")
	tryRegister(&_IOBluetoothGetPairingController, lib, "IOBluetoothGetPairingController")
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



// IOBluetoothGetDeviceSelectorController is a IOBluetoothUI function.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothGetDeviceSelectorController()
func IOBluetoothGetDeviceSelectorController() unsafe.Pointer {
	return _IOBluetoothGetDeviceSelectorController()
}

// IOBluetoothGetPairingController is a IOBluetoothUI function.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothGetPairingController()
func IOBluetoothGetPairingController() unsafe.Pointer {
	return _IOBluetoothGetPairingController()
}



