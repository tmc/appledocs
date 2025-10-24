// Code generated from Apple documentation for ForceFeedback. DO NOT EDIT.

package forcefeedback

/* debug [functions.gen.go]: Generating 20 functions for ForceFeedback */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// ForceFeedback Functions (20 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_FFCreateDevice func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_FFDeviceCreateEffect func(FFDeviceObjectReference, UUIDRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_FFDeviceEscape func(FFDeviceObjectReference, unsafe.Pointer) unsafe.Pointer
	_FFDeviceGetForceFeedbackCapabilities func(FFDeviceObjectReference, unsafe.Pointer) unsafe.Pointer
	_FFDeviceGetForceFeedbackProperty func(FFDeviceObjectReference, FFProperty, unsafe.Pointer, ByteCount) unsafe.Pointer
	_FFDeviceGetForceFeedbackState func(FFDeviceObjectReference, unsafe.Pointer) unsafe.Pointer
	_FFDeviceReleaseEffect func(FFDeviceObjectReference, FFEffectObjectReference) unsafe.Pointer
	_FFDeviceSendForceFeedbackCommand func(FFDeviceObjectReference, FFCommandFlag) unsafe.Pointer
	_FFDeviceSetCooperativeLevel func(FFDeviceObjectReference, unsafe.Pointer, FFCooperativeLevelFlag) unsafe.Pointer
	_FFDeviceSetForceFeedbackProperty func(FFDeviceObjectReference, FFProperty, unsafe.Pointer) unsafe.Pointer
	_FFEffectDownload func(FFEffectObjectReference) unsafe.Pointer
	_FFEffectEscape func(FFEffectObjectReference, unsafe.Pointer) unsafe.Pointer
	_FFEffectGetEffectStatus func(FFEffectObjectReference, unsafe.Pointer) unsafe.Pointer
	_FFEffectGetParameters func(FFEffectObjectReference, unsafe.Pointer, FFEffectParameterFlag) unsafe.Pointer
	_FFEffectSetParameters func(FFEffectObjectReference, unsafe.Pointer, FFEffectParameterFlag) unsafe.Pointer
	_FFEffectStart func(FFEffectObjectReference, unsafe.Pointer, FFEffectStartFlag) unsafe.Pointer
	_FFEffectStop func(FFEffectObjectReference) unsafe.Pointer
	_FFEffectUnload func(FFEffectObjectReference) unsafe.Pointer
	_FFIsForceFeedback func(unsafe.Pointer) unsafe.Pointer
	_FFReleaseDevice func(FFDeviceObjectReference) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_FFCreateDevice, lib, "FFCreateDevice")
	tryRegister(&_FFDeviceCreateEffect, lib, "FFDeviceCreateEffect")
	tryRegister(&_FFDeviceEscape, lib, "FFDeviceEscape")
	tryRegister(&_FFDeviceGetForceFeedbackCapabilities, lib, "FFDeviceGetForceFeedbackCapabilities")
	tryRegister(&_FFDeviceGetForceFeedbackProperty, lib, "FFDeviceGetForceFeedbackProperty")
	tryRegister(&_FFDeviceGetForceFeedbackState, lib, "FFDeviceGetForceFeedbackState")
	tryRegister(&_FFDeviceReleaseEffect, lib, "FFDeviceReleaseEffect")
	tryRegister(&_FFDeviceSendForceFeedbackCommand, lib, "FFDeviceSendForceFeedbackCommand")
	tryRegister(&_FFDeviceSetCooperativeLevel, lib, "FFDeviceSetCooperativeLevel")
	tryRegister(&_FFDeviceSetForceFeedbackProperty, lib, "FFDeviceSetForceFeedbackProperty")
	tryRegister(&_FFEffectDownload, lib, "FFEffectDownload")
	tryRegister(&_FFEffectEscape, lib, "FFEffectEscape")
	tryRegister(&_FFEffectGetEffectStatus, lib, "FFEffectGetEffectStatus")
	tryRegister(&_FFEffectGetParameters, lib, "FFEffectGetParameters")
	tryRegister(&_FFEffectSetParameters, lib, "FFEffectSetParameters")
	tryRegister(&_FFEffectStart, lib, "FFEffectStart")
	tryRegister(&_FFEffectStop, lib, "FFEffectStop")
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
}/* debug [functions.gen.go/function]: FFCreateDevice */

// Creates and initializes an instance of an effect identified by the effect UUID on the device.
//
// Added in macOS 10.2.
// Creates and initializes an instance of an effect identified by the effect UUID on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFDeviceCreateEffect(_:_:_:_:)
func FFDeviceCreateEffect(deviceReference FFDeviceObjectReference, uuidRef UUIDRef, pEffectDefinition unsafe.Pointer, pEffectReference unsafe.Pointer) unsafe.Pointer {
	return _FFDeviceCreateEffect(deviceReference, uuidRef, pEffectDefinition, pEffectReference)
}/* debug [functions.gen.go/function]: FFDeviceCreateEffect */

// Sends a hardware-specific command to the device.
//
// Added in macOS 10.2.
// Sends a hardware-specific command to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFDeviceEscape(_:_:)
func FFDeviceEscape(deviceReference FFDeviceObjectReference, pFFEffectEscape unsafe.Pointer) unsafe.Pointer {
	return _FFDeviceEscape(deviceReference, pFFEffectEscape)
}/* debug [functions.gen.go/function]: FFDeviceEscape */

// Retrieves the device’s force feedback capabilities.
//
// Added in macOS 10.2.
// Retrieves the device’s force feedback capabilities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFDeviceGetForceFeedbackCapabilities(_:_:)
func FFDeviceGetForceFeedbackCapabilities(deviceReference FFDeviceObjectReference, pFFCapabilities unsafe.Pointer) unsafe.Pointer {
	return _FFDeviceGetForceFeedbackCapabilities(deviceReference, pFFCapabilities)
}/* debug [functions.gen.go/function]: FFDeviceGetForceFeedbackCapabilities */

// Gets properties that define the device behavior.
//
// Added in macOS 10.2.
// Gets properties that define the device behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFDeviceGetForceFeedbackProperty(_:_:_:_:)
func FFDeviceGetForceFeedbackProperty(deviceReference FFDeviceObjectReference, property FFProperty, pValue unsafe.Pointer, valueSize ByteCount) unsafe.Pointer {
	return _FFDeviceGetForceFeedbackProperty(deviceReference, property, pValue, valueSize)
}/* debug [functions.gen.go/function]: FFDeviceGetForceFeedbackProperty */

// Retrieves the state of the device’s force feedback system.
//
// Added in macOS 10.2.
// Retrieves the state of the device’s force feedback system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFDeviceGetForceFeedbackState(_:_:)
func FFDeviceGetForceFeedbackState(deviceReference FFDeviceObjectReference, pFFState unsafe.Pointer) unsafe.Pointer {
	return _FFDeviceGetForceFeedbackState(deviceReference, pFFState)
}/* debug [functions.gen.go/function]: FFDeviceGetForceFeedbackState */

// Disposes of an API effect object created with FFDeviceCreateEffect.
//
// Added in macOS 10.2.
// Disposes of an API effect object created with FFDeviceCreateEffect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFDeviceReleaseEffect(_:_:)
func FFDeviceReleaseEffect(deviceReference FFDeviceObjectReference, effectReference FFEffectObjectReference) unsafe.Pointer {
	return _FFDeviceReleaseEffect(deviceReference, effectReference)
}/* debug [functions.gen.go/function]: FFDeviceReleaseEffect */

// Sends a command to the device’s force feedback system.
//
// Added in macOS 10.2.
// Sends a command to the device’s force feedback system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFDeviceSendForceFeedbackCommand(_:_:)
func FFDeviceSendForceFeedbackCommand(deviceReference FFDeviceObjectReference, flags FFCommandFlag) unsafe.Pointer {
	return _FFDeviceSendForceFeedbackCommand(deviceReference, flags)
}/* debug [functions.gen.go/function]: FFDeviceSendForceFeedbackCommand */

// Function is unimplemented in version 1.0 of this API
//
// Added in macOS 10.2.
// Function is unimplemented in version 1.0 of this API
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFDeviceSetCooperativeLevel(_:_:_:)
func FFDeviceSetCooperativeLevel(deviceReference FFDeviceObjectReference, taskIdentifier unsafe.Pointer, flags FFCooperativeLevelFlag) unsafe.Pointer {
	return _FFDeviceSetCooperativeLevel(deviceReference, taskIdentifier, flags)
}/* debug [functions.gen.go/function]: FFDeviceSetCooperativeLevel */

// Retrieves the device’s force feedback capabilities.
//
// Added in macOS 10.2.
// Retrieves the device’s force feedback capabilities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFDeviceSetForceFeedbackProperty(_:_:_:)
func FFDeviceSetForceFeedbackProperty(deviceReference FFDeviceObjectReference, property FFProperty, pValue unsafe.Pointer) unsafe.Pointer {
	return _FFDeviceSetForceFeedbackProperty(deviceReference, property, pValue)
}/* debug [functions.gen.go/function]: FFDeviceSetForceFeedbackProperty */

// Places the effect on the device. If the effect is already on the device, the existing effect is updated to match the values set by the FFEffectSetParameters method.
//
// Added in macOS 10.2.
// Places the effect on the device. If the effect is already on the device, the existing effect is updated to match the values set by the FFEffectSetParameters method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFEffectDownload(_:)
func FFEffectDownload(effectReference FFEffectObjectReference) unsafe.Pointer {
	return _FFEffectDownload(effectReference)
}/* debug [functions.gen.go/function]: FFEffectDownload */

// Sends a hardware-specific command to the driver.
//
// Added in macOS 10.2.
// Sends a hardware-specific command to the driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFEffectEscape(_:_:)
func FFEffectEscape(effectReference FFEffectObjectReference, pFFEffectEscape unsafe.Pointer) unsafe.Pointer {
	return _FFEffectEscape(effectReference, pFFEffectEscape)
}/* debug [functions.gen.go/function]: FFEffectEscape */

// Sends a hardware-specific command to the driver.
//
// Added in macOS 10.2.
// Sends a hardware-specific command to the driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFEffectGetEffectStatus(_:_:)
func FFEffectGetEffectStatus(effectReference FFEffectObjectReference, pFlags unsafe.Pointer) unsafe.Pointer {
	return _FFEffectGetEffectStatus(effectReference, pFlags)
}/* debug [functions.gen.go/function]: FFEffectGetEffectStatus */

// Retrieves information about an effect.
//
// Added in macOS 10.2.
// Retrieves information about an effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFEffectGetParameters(_:_:_:)
func FFEffectGetParameters(effectReference FFEffectObjectReference, pFFEffect unsafe.Pointer, flags FFEffectParameterFlag) unsafe.Pointer {
	return _FFEffectGetParameters(effectReference, pFFEffect, flags)
}/* debug [functions.gen.go/function]: FFEffectGetParameters */

// Sets the characteristics of an effect.
//
// Added in macOS 10.2.
// Sets the characteristics of an effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFEffectSetParameters(_:_:_:)
func FFEffectSetParameters(effectReference FFEffectObjectReference, pFFEffect unsafe.Pointer, flags FFEffectParameterFlag) unsafe.Pointer {
	return _FFEffectSetParameters(effectReference, pFFEffect, flags)
}/* debug [functions.gen.go/function]: FFEffectSetParameters */

// Begins playing an effect. If the effect is already playing, it is restarted from the beginning. If the effect has not been downloaded or has been modified since its last download, it is downloaded before being started. This default behavior can be suppressed by passing the FFES_NODOWNLOAD flag.
//
// Added in macOS 10.2.
// Begins playing an effect. If the effect is already playing, it is restarted from the beginning. If the effect has not been downloaded or has been modified since its last download, it is downloaded before being started. This default behavior can be suppressed by passing the FFES_NODOWNLOAD flag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFEffectStart(_:_:_:)
func FFEffectStart(effectReference FFEffectObjectReference, iterations unsafe.Pointer, flags FFEffectStartFlag) unsafe.Pointer {
	return _FFEffectStart(effectReference, iterations, flags)
}/* debug [functions.gen.go/function]: FFEffectStart */

// Stops playing an effect.
//
// Added in macOS 10.2.
// Stops playing an effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFEffectStop(_:)
func FFEffectStop(effectReference FFEffectObjectReference) unsafe.Pointer {
	return _FFEffectStop(effectReference)
}/* debug [functions.gen.go/function]: FFEffectStop */

// Removes the effect from the device. If the effect is playing, it is automatically stopped before it is unloaded.
//
// Added in macOS 10.2.
// Removes the effect from the device. If the effect is playing, it is automatically stopped before it is unloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFEffectUnload(_:)
func FFEffectUnload(effectReference FFEffectObjectReference) unsafe.Pointer {
	return _FFEffectUnload(effectReference)
}/* debug [functions.gen.go/function]: FFEffectUnload */

// Used to determine if a particular device provided by HID Manager is a force feedback device.
//
// Added in macOS 10.2.
// Used to determine if a particular device provided by HID Manager is a force feedback device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFIsForceFeedback(_:)
func FFIsForceFeedback(hidDevice unsafe.Pointer) unsafe.Pointer {
	return _FFIsForceFeedback(hidDevice)
}/* debug [functions.gen.go/function]: FFIsForceFeedback */

// Disposes of an API device object created with FFCreateDevice.
//
// Added in macOS 10.2.
// Disposes of an API device object created with FFCreateDevice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ForceFeedback/FFReleaseDevice(_:)
func FFReleaseDevice(deviceReference FFDeviceObjectReference) unsafe.Pointer {
	return _FFReleaseDevice(deviceReference)
}/* debug [functions.gen.go/function]: FFReleaseDevice */




