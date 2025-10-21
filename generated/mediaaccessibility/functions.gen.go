// Code generated from Apple documentation for MediaAccessibility. DO NOT EDIT.

package mediaaccessibility

import (
	"unsafe"

	"github.com/ebitengine/purego"
	coregraphics "github.com/tmc/appledocs/generated/coregraphics"
)


// MediaAccessibility Functions (15 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MACaptionAppearanceCopyActiveProfileID func() unsafe.Pointer
	_MACaptionAppearanceCopyFontDescriptorForStyle func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MACaptionAppearanceCopyProfileIDs func() unsafe.Pointer
	_MACaptionAppearanceCopySelectedLanguages func(unsafe.Pointer) unsafe.Pointer
	_MACaptionAppearanceCopyWindowColor func(unsafe.Pointer, unsafe.Pointer) coregraphics.CGColorRef
	_MACaptionAppearanceDidDisplayCaptions func(unsafe.Pointer) unsafe.Pointer
	_MACaptionAppearanceExecuteBlockForProfileID func(unsafe.Pointer) unsafe.Pointer
	_MACaptionAppearanceGetDisplayType func(unsafe.Pointer) unsafe.Pointer
	_MACaptionAppearanceGetRelativeCharacterSize func(unsafe.Pointer, unsafe.Pointer) float64
	_MACaptionAppearanceGetTextEdgeStyle func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MACaptionAppearanceGetWindowOpacity func(unsafe.Pointer, unsafe.Pointer) float64
	_MACaptionAppearanceIsCustomized func(unsafe.Pointer) unsafe.Pointer
	_MACaptionAppearanceSetDisplayType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MADimFlashingLightsEnabled func() unsafe.Pointer
	_MAImageCaptioningSetCaption func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MACaptionAppearanceCopyActiveProfileID, lib, "MACaptionAppearanceCopyActiveProfileID")
	tryRegister(&_MACaptionAppearanceCopyFontDescriptorForStyle, lib, "MACaptionAppearanceCopyFontDescriptorForStyle")
	tryRegister(&_MACaptionAppearanceCopyProfileIDs, lib, "MACaptionAppearanceCopyProfileIDs")
	tryRegister(&_MACaptionAppearanceCopySelectedLanguages, lib, "MACaptionAppearanceCopySelectedLanguages")
	tryRegister(&_MACaptionAppearanceCopyWindowColor, lib, "MACaptionAppearanceCopyWindowColor")
	tryRegister(&_MACaptionAppearanceDidDisplayCaptions, lib, "MACaptionAppearanceDidDisplayCaptions")
	tryRegister(&_MACaptionAppearanceExecuteBlockForProfileID, lib, "MACaptionAppearanceExecuteBlockForProfileID")
	tryRegister(&_MACaptionAppearanceGetDisplayType, lib, "MACaptionAppearanceGetDisplayType")
	tryRegister(&_MACaptionAppearanceGetRelativeCharacterSize, lib, "MACaptionAppearanceGetRelativeCharacterSize")
	tryRegister(&_MACaptionAppearanceGetTextEdgeStyle, lib, "MACaptionAppearanceGetTextEdgeStyle")
	tryRegister(&_MACaptionAppearanceGetWindowOpacity, lib, "MACaptionAppearanceGetWindowOpacity")
	tryRegister(&_MACaptionAppearanceIsCustomized, lib, "MACaptionAppearanceIsCustomized")
	tryRegister(&_MACaptionAppearanceSetDisplayType, lib, "MACaptionAppearanceSetDisplayType")
	tryRegister(&_MADimFlashingLightsEnabled, lib, "MADimFlashingLightsEnabled")
	tryRegister(&_MAImageCaptioningSetCaption, lib, "MAImageCaptioningSetCaption")
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



// MACaptionAppearanceCopyActiveProfileID is a MediaAccessibility function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyActiveProfileID()
func MACaptionAppearanceCopyActiveProfileID() unsafe.Pointer {
	return _MACaptionAppearanceCopyActiveProfileID()
	}


// Returns the preferred font for the specified style of type. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyFontDescriptorForStyle(_:_:_:)
func MACaptionAppearanceCopyFontDescriptorForStyle(domain unsafe.Pointer, behavior unsafe.Pointer, fontStyle unsafe.Pointer) unsafe.Pointer {
	return _MACaptionAppearanceCopyFontDescriptorForStyle(domain, behavior, fontStyle)
	}


// MACaptionAppearanceCopyProfileIDs is a MediaAccessibility function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyProfileIDs()
func MACaptionAppearanceCopyProfileIDs() unsafe.Pointer {
	return _MACaptionAppearanceCopyProfileIDs()
	}


// Returns the preferred caption languages. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopySelectedLanguages(_:)
func MACaptionAppearanceCopySelectedLanguages(domain unsafe.Pointer) unsafe.Pointer {
	return _MACaptionAppearanceCopySelectedLanguages(domain)
	}


// Returns the preference for the caption window’s color. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyWindowColor(_:_:)
func MACaptionAppearanceCopyWindowColor(domain unsafe.Pointer, behavior unsafe.Pointer) coregraphics.CGColorRef {
	return _MACaptionAppearanceCopyWindowColor(domain, behavior)
	}


// Informs accessibility clients when captions display onscreen. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceDidDisplayCaptions(_:)
func MACaptionAppearanceDidDisplayCaptions(strings unsafe.Pointer) {
	_MACaptionAppearanceDidDisplayCaptions(strings)
	}


// MACaptionAppearanceExecuteBlockForProfileID is a MediaAccessibility function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceExecuteBlockForProfileID(_:_:)
func MACaptionAppearanceExecuteBlockForProfileID(profileID unsafe.Pointer) {
	_MACaptionAppearanceExecuteBlockForProfileID(profileID)
	}


// Returns the preferred type of captions to display. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceGetDisplayType(_:)
func MACaptionAppearanceGetDisplayType(domain unsafe.Pointer) unsafe.Pointer {
	return _MACaptionAppearanceGetDisplayType(domain)
	}


// Returns the preference for font scaling. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceGetRelativeCharacterSize(_:_:)
func MACaptionAppearanceGetRelativeCharacterSize(domain unsafe.Pointer, behavior unsafe.Pointer) float64 {
	return _MACaptionAppearanceGetRelativeCharacterSize(domain, behavior)
	}


// Returns the preference for text edge style. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceGetTextEdgeStyle(_:_:)
func MACaptionAppearanceGetTextEdgeStyle(domain unsafe.Pointer, behavior unsafe.Pointer) unsafe.Pointer {
	return _MACaptionAppearanceGetTextEdgeStyle(domain, behavior)
	}


// Returns the preference for the overlay’s opacity. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceGetWindowOpacity(_:_:)
func MACaptionAppearanceGetWindowOpacity(domain unsafe.Pointer, behavior unsafe.Pointer) float64 {
	return _MACaptionAppearanceGetWindowOpacity(domain, behavior)
	}


// MACaptionAppearanceIsCustomized is a MediaAccessibility function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceIsCustomized(_:)
func MACaptionAppearanceIsCustomized(domain unsafe.Pointer) unsafe.Pointer {
	return _MACaptionAppearanceIsCustomized(domain)
	}


// Sets the preference for the type of caption. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceSetDisplayType(_:_:)
func MACaptionAppearanceSetDisplayType(domain unsafe.Pointer, displayType unsafe.Pointer) {
	_MACaptionAppearanceSetDisplayType(domain, displayType)
	}


// Returns a Boolean value that indicates whether the flashing lights setting is enabled on the device. [Full Topic]
//
// Added in macOS 13.3.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MADimFlashingLightsEnabled()
func MADimFlashingLightsEnabled() unsafe.Pointer {
	return _MADimFlashingLightsEnabled()
	}


// Sets the accessibility caption for an image’s metadata. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAImageCaptioningSetCaption(_:_:_:)
func MAImageCaptioningSetCaption(url unsafe.Pointer, string_ unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _MAImageCaptioningSetCaption(url, string_, error_)
	}




