// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

/* debug [functions.gen.go]: Generating 3 functions for IOBluetoothUI */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// IOBluetoothUI Functions (3 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_IOBluetoothGetDeviceSelectorController func() BluetoothDeviceSelectorControllerRef
	_IOBluetoothGetPairingController func() BluetoothPairingControllerRef
	_IOBluetoothValidateHardwareWithDescription func(StringRef, StringRef) int
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_IOBluetoothGetDeviceSelectorController, lib, "IOBluetoothGetDeviceSelectorController")
	tryRegister(&_IOBluetoothGetPairingController, lib, "IOBluetoothGetPairingController")
	tryRegister(&_IOBluetoothValidateHardwareWithDescription, lib, "IOBluetoothValidateHardwareWithDescription")
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
func IOBluetoothGetDeviceSelectorController() BluetoothDeviceSelectorControllerRef {
	return _IOBluetoothGetDeviceSelectorController()
}/* debug [functions.gen.go/function]: IOBluetoothGetDeviceSelectorController */

// IOBluetoothGetPairingController is a IOBluetoothUI function.
//
// Added in macOS 10.2.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothGetPairingController()
func IOBluetoothGetPairingController() BluetoothPairingControllerRef {
	return _IOBluetoothGetPairingController()
}/* debug [functions.gen.go/function]: IOBluetoothGetPairingController */

// IOBluetoothValidateHardwareWithDescription is a IOBluetoothUI function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothValidateHardwareWithDescription(_:_:)
func IOBluetoothValidateHardwareWithDescription(cancelButtonTitle StringRef, descriptionText StringRef) int {
	return _IOBluetoothValidateHardwareWithDescription(cancelButtonTitle, descriptionText)
}/* debug [functions.gen.go/function]: IOBluetoothValidateHardwareWithDescription */




