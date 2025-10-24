// Code generated from Apple documentation for MediaAccessibility. DO NOT EDIT.

package mediaaccessibility

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// MediaAccessibility Functions (11 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MACaptionAppearanceCopyActiveProfileID func() unsafe.Pointer
	_MACaptionAppearanceCopyFontDescriptorForStyle func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MACaptionAppearanceCopyProfileIDs func() unsafe.Pointer
	_MACaptionAppearanceCopySelectedLanguages func(unsafe.Pointer) unsafe.Pointer
	_MACaptionAppearanceCopyWindowColor func(unsafe.Pointer, unsafe.Pointer) ColorRef
	_MACaptionAppearanceExecuteBlockForProfileID func(unsafe.Pointer)
	_MACaptionAppearanceGetRelativeCharacterSize func(unsafe.Pointer, unsafe.Pointer) float64
	_MACaptionAppearanceGetWindowOpacity func(unsafe.Pointer, unsafe.Pointer) float64
	_MACaptionAppearanceGetWindowRoundedCornerRadius func(unsafe.Pointer, unsafe.Pointer) float64
	_MACaptionAppearanceSetDisplayType func(unsafe.Pointer, unsafe.Pointer)
	_MAImageCaptioningCopyMetadataTagPath func() unsafe.Pointer
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
	tryRegister(&_MACaptionAppearanceExecuteBlockForProfileID, lib, "MACaptionAppearanceExecuteBlockForProfileID")
	tryRegister(&_MACaptionAppearanceGetRelativeCharacterSize, lib, "MACaptionAppearanceGetRelativeCharacterSize")
	tryRegister(&_MACaptionAppearanceGetWindowOpacity, lib, "MACaptionAppearanceGetWindowOpacity")
	tryRegister(&_MACaptionAppearanceGetWindowRoundedCornerRadius, lib, "MACaptionAppearanceGetWindowRoundedCornerRadius")
	tryRegister(&_MACaptionAppearanceSetDisplayType, lib, "MACaptionAppearanceSetDisplayType")
	tryRegister(&_MAImageCaptioningCopyMetadataTagPath, lib, "MAImageCaptioningCopyMetadataTagPath")
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



// MACaptionAppearanceCopyActiveProfileID is a MediaAccessibility function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyActiveProfileID()
func MACaptionAppearanceCopyActiveProfileID() unsafe.Pointer {
	return _MACaptionAppearanceCopyActiveProfileID()
}

// Returns the preferred font for the specified style of type.
//
// Added in macOS 10.9.
// Returns the preferred font for the specified style of type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyFontDescriptorForStyle(_:_:_:)
func MACaptionAppearanceCopyFontDescriptorForStyle(domain unsafe.Pointer, behavior unsafe.Pointer, fontStyle unsafe.Pointer) unsafe.Pointer {
	return _MACaptionAppearanceCopyFontDescriptorForStyle(domain, behavior, fontStyle)
}

// MACaptionAppearanceCopyProfileIDs is a MediaAccessibility function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyProfileIDs()
func MACaptionAppearanceCopyProfileIDs() unsafe.Pointer {
	return _MACaptionAppearanceCopyProfileIDs()
}

// Returns the preferred caption languages.
//
// Added in macOS 10.9.
// Returns the preferred caption languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopySelectedLanguages(_:)
func MACaptionAppearanceCopySelectedLanguages(domain unsafe.Pointer) unsafe.Pointer {
	return _MACaptionAppearanceCopySelectedLanguages(domain)
}

// Returns the preference for the caption window’s color.
//
// Added in macOS 10.9.
// Returns the preference for the caption window’s color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyWindowColor(_:_:)
func MACaptionAppearanceCopyWindowColor(domain unsafe.Pointer, behavior unsafe.Pointer) ColorRef {
	return _MACaptionAppearanceCopyWindowColor(domain, behavior)
}

// MACaptionAppearanceExecuteBlockForProfileID is a MediaAccessibility function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceExecuteBlockForProfileID(_:_:)
func MACaptionAppearanceExecuteBlockForProfileID(profileID unsafe.Pointer) {
	_MACaptionAppearanceExecuteBlockForProfileID(profileID)
}

// Returns the preference for font scaling.
//
// Added in macOS 10.9.
// Returns the preference for font scaling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceGetRelativeCharacterSize(_:_:)
func MACaptionAppearanceGetRelativeCharacterSize(domain unsafe.Pointer, behavior unsafe.Pointer) float64 {
	return _MACaptionAppearanceGetRelativeCharacterSize(domain, behavior)
}

// Returns the preference for the overlay’s opacity.
//
// Added in macOS 10.9.
// Returns the preference for the overlay’s opacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceGetWindowOpacity(_:_:)
func MACaptionAppearanceGetWindowOpacity(domain unsafe.Pointer, behavior unsafe.Pointer) float64 {
	return _MACaptionAppearanceGetWindowOpacity(domain, behavior)
}

// Returns the radius of the caption window’s corners.
//
// Added in macOS 10.9.
// Returns the radius of the caption window’s corners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceGetWindowRoundedCornerRadius(_:_:)
func MACaptionAppearanceGetWindowRoundedCornerRadius(domain unsafe.Pointer, behavior unsafe.Pointer) float64 {
	return _MACaptionAppearanceGetWindowRoundedCornerRadius(domain, behavior)
}

// Sets the preference for the type of caption.
//
// Added in macOS 10.9.
// Sets the preference for the type of caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceSetDisplayType(_:_:)
func MACaptionAppearanceSetDisplayType(domain unsafe.Pointer, displayType unsafe.Pointer) {
	_MACaptionAppearanceSetDisplayType(domain, displayType)
}

// Returns the metadata tag path.
//
// Added in macOS 10.15.
// Returns the metadata tag path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAImageCaptioningCopyMetadataTagPath()
func MAImageCaptioningCopyMetadataTagPath() unsafe.Pointer {
	return _MAImageCaptioningCopyMetadataTagPath()
}



