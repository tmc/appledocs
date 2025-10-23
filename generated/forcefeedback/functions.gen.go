// Code generated from Apple documentation for ForceFeedback. DO NOT EDIT.

package forcefeedback

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// ForceFeedback Functions (10 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_FFCreateDevice func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_FFDeviceCreateEffect func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_FFDeviceGetForceFeedbackCapabilities func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_FFEffectEscape func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_FFEffectGetEffectStatus func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_FFEffectSetParameters func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_FFEffectStart func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_FFEffectUnload func(unsafe.Pointer) unsafe.Pointer
	_FFIsForceFeedback func(unsafe.Pointer) unsafe.Pointer
	_FFReleaseDevice func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_FFCreateDevice, lib, "FFCreateDevice")
	tryRegister(&_FFDeviceCreateEffect, lib, "FFDeviceCreateEffect")
	tryRegister(&_FFDeviceGetForceFeedbackCapabilities, lib, "FFDeviceGetForceFeedbackCapabilities")
	tryRegister(&_FFEffectEscape, lib, "FFEffectEscape")
	tryRegister(&_FFEffectGetEffectStatus, lib, "FFEffectGetEffectStatus")
	tryRegister(&_FFEffectSetParameters, lib, "FFEffectSetParameters")
	tryRegister(&_FFEffectStart, lib, "FFEffectStart")
	tryRegister(&_FFEffectUnload, lib, "FFEffectUnload")
	tryRegister(&_FFIsForceFeedback, lib, "FFIsForceFeedback")
	tryRegister(&_FFReleaseDevice, lib, "FFReleaseDevice")
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



// Creates a new API device object from an OS object in preparation to use the device for force feedback.
//
// Added in macOS 10.2.
// Creates a new API device object from an OS object in preparation to use the device for force feedback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFCreateDevice(_:_:)
func FFCreateDevice(hidDevice unsafe.Pointer, pDeviceReference unsafe.Pointer) unsafe.Pointer {
	return _FFCreateDevice(hidDevice, pDeviceReference)
}

// Creates and initializes an instance of an effect identified by the effect UUID on the device.
//
// Added in macOS 10.2.
// Creates and initializes an instance of an effect identified by the effect UUID on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFDeviceCreateEffect(_:_:_:_:)
func FFDeviceCreateEffect(deviceReference unsafe.Pointer, uuidRef unsafe.Pointer, pEffectDefinition unsafe.Pointer, pEffectReference unsafe.Pointer) unsafe.Pointer {
	return _FFDeviceCreateEffect(deviceReference, uuidRef, pEffectDefinition, pEffectReference)
}

// Retrieves the device’s force feedback capabilities.
//
// Added in macOS 10.2.
// Retrieves the device’s force feedback capabilities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFDeviceGetForceFeedbackCapabilities(_:_:)
func FFDeviceGetForceFeedbackCapabilities(deviceReference unsafe.Pointer, pFFCapabilities unsafe.Pointer) unsafe.Pointer {
	return _FFDeviceGetForceFeedbackCapabilities(deviceReference, pFFCapabilities)
}

// Sends a hardware-specific command to the driver.
//
// Added in macOS 10.2.
// Sends a hardware-specific command to the driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFEffectEscape(_:_:)
func FFEffectEscape(effectReference unsafe.Pointer, pFFEffectEscape unsafe.Pointer) unsafe.Pointer {
	return _FFEffectEscape(effectReference, pFFEffectEscape)
}

// Sends a hardware-specific command to the driver.
//
// Added in macOS 10.2.
// Sends a hardware-specific command to the driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFEffectGetEffectStatus(_:_:)
func FFEffectGetEffectStatus(effectReference unsafe.Pointer, pFlags unsafe.Pointer) unsafe.Pointer {
	return _FFEffectGetEffectStatus(effectReference, pFlags)
}

// Sets the characteristics of an effect.
//
// Added in macOS 10.2.
// Sets the characteristics of an effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFEffectSetParameters(_:_:_:)
func FFEffectSetParameters(effectReference unsafe.Pointer, pFFEffect unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _FFEffectSetParameters(effectReference, pFFEffect, flags)
}

// Begins playing an effect. If the effect is already playing, it is restarted from the beginning. If the effect has not been downloaded or has been modified since its last download, it is downloaded before being started. This default behavior can be suppressed by passing the FFES_NODOWNLOAD flag.
//
// Added in macOS 10.2.
// Begins playing an effect. If the effect is already playing, it is restarted from the beginning. If the effect has not been downloaded or has been modified since its last download, it is downloaded before being started. This default behavior can be suppressed by passing the FFES_NODOWNLOAD flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFEffectStart(_:_:_:)
func FFEffectStart(effectReference unsafe.Pointer, iterations unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _FFEffectStart(effectReference, iterations, flags)
}

// Removes the effect from the device. If the effect is playing, it is automatically stopped before it is unloaded.
//
// Added in macOS 10.2.
// Removes the effect from the device. If the effect is playing, it is automatically stopped before it is unloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFEffectUnload(_:)
func FFEffectUnload(effectReference unsafe.Pointer) unsafe.Pointer {
	return _FFEffectUnload(effectReference)
}

// Used to determine if a particular device provided by HID Manager is a force feedback device.
//
// Added in macOS 10.2.
// Used to determine if a particular device provided by HID Manager is a force feedback device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFIsForceFeedback(_:)
func FFIsForceFeedback(hidDevice unsafe.Pointer) unsafe.Pointer {
	return _FFIsForceFeedback(hidDevice)
}

// Disposes of an API device object created with FFCreateDevice.
//
// Added in macOS 10.2.
// Disposes of an API device object created with FFCreateDevice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFReleaseDevice(_:)
func FFReleaseDevice(deviceReference unsafe.Pointer) unsafe.Pointer {
	return _FFReleaseDevice(deviceReference)
}



