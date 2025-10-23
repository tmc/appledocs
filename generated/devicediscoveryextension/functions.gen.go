// Code generated from Apple documentation for DeviceDiscoveryExtension. DO NOT EDIT.

package devicediscoveryextension

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// DeviceDiscoveryExtension Functions (5 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_DDDeviceCategoryToString func(unsafe.Pointer) unsafe.Pointer
	_DDDeviceMediaPlaybackStateToString func(unsafe.Pointer) unsafe.Pointer
	_DDDeviceProtocolToString func(unsafe.Pointer) unsafe.Pointer
	_DDDeviceStateToString func(unsafe.Pointer) unsafe.Pointer
	_DDEventTypeToString func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_DDDeviceCategoryToString, lib, "DDDeviceCategoryToString")
	tryRegister(&_DDDeviceMediaPlaybackStateToString, lib, "DDDeviceMediaPlaybackStateToString")
	tryRegister(&_DDDeviceProtocolToString, lib, "DDDeviceProtocolToString")
	tryRegister(&_DDDeviceStateToString, lib, "DDDeviceStateToString")
	tryRegister(&_DDEventTypeToString, lib, "DDEventTypeToString")
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



// Returns human-readable text for the specified identifier that describes a device’s category.

// Returns human-readable text for the specified identifier that describes a device’s category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceCategoryToString(_:)
func DDDeviceCategoryToString(inValue unsafe.Pointer) unsafe.Pointer {
	return _DDDeviceCategoryToString(inValue)
	}


// Returns human-readable text for the specified media playback state.

// Returns human-readable text for the specified media playback state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceMediaPlaybackStateToString(_:)
func DDDeviceMediaPlaybackStateToString(inValue unsafe.Pointer) unsafe.Pointer {
	return _DDDeviceMediaPlaybackStateToString(inValue)
	}


// Returns human-readable text for the specified protocol identifier.

// Returns human-readable text for the specified protocol identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceProtocolToString(_:)
func DDDeviceProtocolToString(inValue unsafe.Pointer) unsafe.Pointer {
	return _DDDeviceProtocolToString(inValue)
	}


// Returns human-readable text for the specified identifier that describes a device’s status.

// Returns human-readable text for the specified identifier that describes a device’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceStateToString(_:)
func DDDeviceStateToString(inValue unsafe.Pointer) unsafe.Pointer {
	return _DDDeviceStateToString(inValue)
	}


// Returns human-readable text for the specified event identifier.

// Returns human-readable text for the specified event identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDEventTypeToString(_:)
func DDEventTypeToString(inValue unsafe.Pointer) unsafe.Pointer {
	return _DDEventTypeToString(inValue)
	}




