// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Accessibility Functions (5 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_AXMFiHearingDeviceStreamingEar func() unsafe.Pointer
	_AXSupportsBidirectionalAXMFiHearingDeviceStreaming func() bool
	_AXPrefersActionSliderAlternative func() bool
	_AXShowBordersEnabled func() bool
	_AXAssistiveAccessEnabled func() bool
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_AXMFiHearingDeviceStreamingEar, lib, "AXMFiHearingDeviceStreamingEar")
	tryRegister(&_AXSupportsBidirectionalAXMFiHearingDeviceStreaming, lib, "AXSupportsBidirectionalAXMFiHearingDeviceStreaming")
	tryRegister(&_AXPrefersActionSliderAlternative, lib, "AXPrefersActionSliderAlternative")
	tryRegister(&_AXShowBordersEnabled, lib, "AXShowBordersEnabled")
	tryRegister(&_AXAssistiveAccessEnabled, lib, "AXAssistiveAccessEnabled")
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



// Returns which ears enable streaming.

// Returns which ears enable streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMFiHearingDevice/streamingEar()
func AXMFiHearingDeviceStreamingEar() unsafe.Pointer {
	return _AXMFiHearingDeviceStreamingEar()
}

// Returns a Boolean value that indicates whether the iOS device supports bidirectional streaming.

// Returns a Boolean value that indicates whether the iOS device supports bidirectional streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMFiHearingDevice/supportsBidirectionalStreaming()
func AXSupportsBidirectionalAXMFiHearingDeviceStreaming() bool {
	return _AXSupportsBidirectionalAXMFiHearingDeviceStreaming()
}

// AXPrefersActionSliderAlternative is a Accessibility function.
//
// Added in macOS 26.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXPrefersActionSliderAlternative
func AXPrefersActionSliderAlternative() bool {
	return _AXPrefersActionSliderAlternative()
}

// AXShowBordersEnabled is a Accessibility function.
//
// Added in macOS 26.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXShowBordersEnabled
func AXShowBordersEnabled() bool {
	return _AXShowBordersEnabled()
}

// A Boolean value that indicates whether Assistive Access is running.
//
// Added in macOS 15.0.
// A Boolean value that indicates whether Assistive Access is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/isAssistiveAccessEnabled
func AXAssistiveAccessEnabled() bool {
	return _AXAssistiveAccessEnabled()
}



