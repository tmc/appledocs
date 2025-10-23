// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// IOUSBHost Functions (6 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_IOUSBGetEndpointMaxStreams func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_IOUSBHostCIControllerStateToString func(unsafe.Pointer) unsafe.Pointer
	_IOUSBHostCIDeviceStateToString func(unsafe.Pointer) unsafe.Pointer
	_IOUSBHostCILinkStateToString func(unsafe.Pointer) unsafe.Pointer
	_IOUSBHostCIMessageStatusFromIOReturn func(unsafe.Pointer) unsafe.Pointer
	_IOUSBHostCIPortStateToString func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_IOUSBGetEndpointMaxStreams, lib, "IOUSBGetEndpointMaxStreams")
	tryRegister(&_IOUSBHostCIControllerStateToString, lib, "IOUSBHostCIControllerStateToString")
	tryRegister(&_IOUSBHostCIDeviceStateToString, lib, "IOUSBHostCIDeviceStateToString")
	tryRegister(&_IOUSBHostCILinkStateToString, lib, "IOUSBHostCILinkStateToString")
	tryRegister(&_IOUSBHostCIMessageStatusFromIOReturn, lib, "IOUSBHostCIMessageStatusFromIOReturn")
	tryRegister(&_IOUSBHostCIPortStateToString, lib, "IOUSBHostCIPortStateToString")
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



// Obtains the number of supported streams.
//
// Added in macOS 10.15.

// Obtains the number of supported streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBGetEndpointMaxStreams(_:_:_:)
func IOUSBGetEndpointMaxStreams(usbDeviceSpeed unsafe.Pointer, descriptor unsafe.Pointer, companionDescriptor unsafe.Pointer) unsafe.Pointer {
	return _IOUSBGetEndpointMaxStreams(usbDeviceSpeed, descriptor, companionDescriptor)
	}


// IOUSBHostCIControllerStateToString is a IOUSBHost function.
//
// Added in macOS 10.15.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateToString(_:)
func IOUSBHostCIControllerStateToString(controllerState unsafe.Pointer) unsafe.Pointer {
	return _IOUSBHostCIControllerStateToString(controllerState)
	}


// IOUSBHostCIDeviceStateToString is a IOUSBHost function.
//
// Added in macOS 10.15.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateToString(_:)
func IOUSBHostCIDeviceStateToString(deviceState unsafe.Pointer) unsafe.Pointer {
	return _IOUSBHostCIDeviceStateToString(deviceState)
	}


// IOUSBHostCILinkStateToString is a IOUSBHost function.
//
// Added in macOS 10.15.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCILinkStateToString(_:)
func IOUSBHostCILinkStateToString(linkState unsafe.Pointer) unsafe.Pointer {
	return _IOUSBHostCILinkStateToString(linkState)
	}


// IOUSBHostCIMessageStatusFromIOReturn is a IOUSBHost function.
//
// Added in macOS 10.15.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIMessageStatusFromIOReturn(_:)
func IOUSBHostCIMessageStatusFromIOReturn(status unsafe.Pointer) unsafe.Pointer {
	return _IOUSBHostCIMessageStatusFromIOReturn(status)
	}


// IOUSBHostCIPortStateToString is a IOUSBHost function.
//
// Added in macOS 10.15.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateToString(_:)
func IOUSBHostCIPortStateToString(portState unsafe.Pointer) unsafe.Pointer {
	return _IOUSBHostCIPortStateToString(portState)
	}




