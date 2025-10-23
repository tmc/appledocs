// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"unsafe"

	"github.com/ebitengine/purego"
	coregraphics "github.com/tmc/appledocs/generated/coregraphics"
)


// Accessibility Functions (11 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_AXAnimatedImagesEnabled func() bool
	_AXMFiHearingDeviceStreamingEar func() unsafe.Pointer
	_AXSupportsBidirectionalAXMFiHearingDeviceStreaming func() bool
	_AXNameFromColor func(coregraphics.ColorRef) unsafe.Pointer
	_AXOpenSettingsFeature func(unsafe.Pointer)
	_AXPrefersActionSliderAlternative func() bool
	_AXPrefersHeadAnchorAlternative func() bool
	_AXPrefersHorizontalTextLayout func() bool
	_AXPrefersNonBlinkingTextInsertionIndicator func() bool
	_AXShowBordersEnabled func() bool
	_AXAssistiveAccessEnabled func() bool
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_AXAnimatedImagesEnabled, lib, "AXAnimatedImagesEnabled")
	tryRegister(&_AXMFiHearingDeviceStreamingEar, lib, "AXMFiHearingDeviceStreamingEar")
	tryRegister(&_AXSupportsBidirectionalAXMFiHearingDeviceStreaming, lib, "AXSupportsBidirectionalAXMFiHearingDeviceStreaming")
	tryRegister(&_AXNameFromColor, lib, "AXNameFromColor")
	tryRegister(&_AXOpenSettingsFeature, lib, "AXOpenSettingsFeature")
	tryRegister(&_AXPrefersActionSliderAlternative, lib, "AXPrefersActionSliderAlternative")
	tryRegister(&_AXPrefersHeadAnchorAlternative, lib, "AXPrefersHeadAnchorAlternative")
	tryRegister(&_AXPrefersHorizontalTextLayout, lib, "AXPrefersHorizontalTextLayout")
	tryRegister(&_AXPrefersNonBlinkingTextInsertionIndicator, lib, "AXPrefersNonBlinkingTextInsertionIndicator")
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



// AXAnimatedImagesEnabled is a Accessibility function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXAnimatedImagesEnabled
func AXAnimatedImagesEnabled() bool {
	return _AXAnimatedImagesEnabled()
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

// Returns a localized description of the color to use in accessibility attributes.
//
// Added in macOS 11.0.
// Returns a localized description of the color to use in accessibility attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNameFromColor(_:)
func AXNameFromColor(color coregraphics.ColorRef) unsafe.Pointer {
	return _AXNameFromColor(color)
}

// AXOpenSettingsFeature is a Accessibility function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXOpenSettingsFeature
func AXOpenSettingsFeature(feature unsafe.Pointer) {
	_AXOpenSettingsFeature(feature)
}

// AXPrefersActionSliderAlternative is a Accessibility function.
//
// Added in macOS 26.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXPrefersActionSliderAlternative
func AXPrefersActionSliderAlternative() bool {
	return _AXPrefersActionSliderAlternative()
}

// AXPrefersHeadAnchorAlternative is a Accessibility function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXPrefersHeadAnchorAlternative
func AXPrefersHeadAnchorAlternative() bool {
	return _AXPrefersHeadAnchorAlternative()
}

// AXPrefersHorizontalTextLayout is a Accessibility function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXPrefersHorizontalTextLayout
func AXPrefersHorizontalTextLayout() bool {
	return _AXPrefersHorizontalTextLayout()
}

// AXPrefersNonBlinkingTextInsertionIndicator is a Accessibility function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXPrefersNonBlinkingTextInsertionIndicator
func AXPrefersNonBlinkingTextInsertionIndicator() bool {
	return _AXPrefersNonBlinkingTextInsertionIndicator()
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



