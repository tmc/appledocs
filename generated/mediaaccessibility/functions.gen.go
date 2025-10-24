// Code generated from Apple documentation for MediaAccessibility. DO NOT EDIT.

package mediaaccessibility

/* debug [functions.gen.go]: Generating 27 functions for MediaAccessibility */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// MediaAccessibility Functions (27 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MAAudibleMediaCopyPreferredCharacteristics func() ArrayRef
	_MACaptionAppearanceAddSelectedLanguage func(MACaptionAppearanceDomain, StringRef) bool
	_MACaptionAppearanceCopyActiveProfileID func() StringRef
	_MACaptionAppearanceCopyBackgroundColor func(MACaptionAppearanceDomain, unsafe.Pointer) ColorRef
	_MACaptionAppearanceCopyFontDescriptorForStyle func(MACaptionAppearanceDomain, unsafe.Pointer, MACaptionAppearanceFontStyle) FontDescriptorRef
	_MACaptionAppearanceCopyForegroundColor func(MACaptionAppearanceDomain, unsafe.Pointer) ColorRef
	_MACaptionAppearanceCopyPreferredCaptioningMediaCharacteristics func(MACaptionAppearanceDomain) ArrayRef
	_MACaptionAppearanceCopyProfileIDs func() ArrayRef
	_MACaptionAppearanceCopyProfileName func(StringRef) StringRef
	_MACaptionAppearanceCopySelectedLanguages func(MACaptionAppearanceDomain) ArrayRef
	_MACaptionAppearanceCopyWindowColor func(MACaptionAppearanceDomain, unsafe.Pointer) ColorRef
	_MACaptionAppearanceDidDisplayCaptions func(ArrayRef)
	_MACaptionAppearanceExecuteBlockForProfileID func(StringRef)
	_MACaptionAppearanceGetBackgroundOpacity func(MACaptionAppearanceDomain, unsafe.Pointer) float64
	_MACaptionAppearanceGetDisplayType func(MACaptionAppearanceDomain) MACaptionAppearanceDisplayType
	_MACaptionAppearanceGetForegroundOpacity func(MACaptionAppearanceDomain, unsafe.Pointer) float64
	_MACaptionAppearanceGetRelativeCharacterSize func(MACaptionAppearanceDomain, unsafe.Pointer) float64
	_MACaptionAppearanceGetTextEdgeStyle func(MACaptionAppearanceDomain, unsafe.Pointer) MACaptionAppearanceTextEdgeStyle
	_MACaptionAppearanceGetWindowOpacity func(MACaptionAppearanceDomain, unsafe.Pointer) float64
	_MACaptionAppearanceGetWindowRoundedCornerRadius func(MACaptionAppearanceDomain, unsafe.Pointer) float64
	_MACaptionAppearanceIsCustomized func(MACaptionAppearanceDomain) bool
	_MACaptionAppearanceSetActiveProfileID func(StringRef)
	_MACaptionAppearanceSetDisplayType func(MACaptionAppearanceDomain, MACaptionAppearanceDisplayType)
	_MADimFlashingLightsEnabled func() bool
	_MAImageCaptioningCopyCaption func(URLRef, unsafe.Pointer) StringRef
	_MAImageCaptioningCopyMetadataTagPath func() StringRef
	_MAImageCaptioningSetCaption func(URLRef, StringRef, unsafe.Pointer) bool
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MAAudibleMediaCopyPreferredCharacteristics, lib, "MAAudibleMediaCopyPreferredCharacteristics")
	tryRegister(&_MACaptionAppearanceAddSelectedLanguage, lib, "MACaptionAppearanceAddSelectedLanguage")
	tryRegister(&_MACaptionAppearanceCopyActiveProfileID, lib, "MACaptionAppearanceCopyActiveProfileID")
	tryRegister(&_MACaptionAppearanceCopyBackgroundColor, lib, "MACaptionAppearanceCopyBackgroundColor")
	tryRegister(&_MACaptionAppearanceCopyFontDescriptorForStyle, lib, "MACaptionAppearanceCopyFontDescriptorForStyle")
	tryRegister(&_MACaptionAppearanceCopyForegroundColor, lib, "MACaptionAppearanceCopyForegroundColor")
	tryRegister(&_MACaptionAppearanceCopyPreferredCaptioningMediaCharacteristics, lib, "MACaptionAppearanceCopyPreferredCaptioningMediaCharacteristics")
	tryRegister(&_MACaptionAppearanceCopyProfileIDs, lib, "MACaptionAppearanceCopyProfileIDs")
	tryRegister(&_MACaptionAppearanceCopyProfileName, lib, "MACaptionAppearanceCopyProfileName")
	tryRegister(&_MACaptionAppearanceCopySelectedLanguages, lib, "MACaptionAppearanceCopySelectedLanguages")
	tryRegister(&_MACaptionAppearanceCopyWindowColor, lib, "MACaptionAppearanceCopyWindowColor")
	tryRegister(&_MACaptionAppearanceDidDisplayCaptions, lib, "MACaptionAppearanceDidDisplayCaptions")
	tryRegister(&_MACaptionAppearanceExecuteBlockForProfileID, lib, "MACaptionAppearanceExecuteBlockForProfileID")
	tryRegister(&_MACaptionAppearanceGetBackgroundOpacity, lib, "MACaptionAppearanceGetBackgroundOpacity")
	tryRegister(&_MACaptionAppearanceGetDisplayType, lib, "MACaptionAppearanceGetDisplayType")
	tryRegister(&_MACaptionAppearanceGetForegroundOpacity, lib, "MACaptionAppearanceGetForegroundOpacity")
	tryRegister(&_MACaptionAppearanceGetRelativeCharacterSize, lib, "MACaptionAppearanceGetRelativeCharacterSize")
	tryRegister(&_MACaptionAppearanceGetTextEdgeStyle, lib, "MACaptionAppearanceGetTextEdgeStyle")
	tryRegister(&_MACaptionAppearanceGetWindowOpacity, lib, "MACaptionAppearanceGetWindowOpacity")
	tryRegister(&_MACaptionAppearanceGetWindowRoundedCornerRadius, lib, "MACaptionAppearanceGetWindowRoundedCornerRadius")
	tryRegister(&_MACaptionAppearanceIsCustomized, lib, "MACaptionAppearanceIsCustomized")
	tryRegister(&_MACaptionAppearanceSetActiveProfileID, lib, "MACaptionAppearanceSetActiveProfileID")
	tryRegister(&_MACaptionAppearanceSetDisplayType, lib, "MACaptionAppearanceSetDisplayType")
	tryRegister(&_MADimFlashingLightsEnabled, lib, "MADimFlashingLightsEnabled")
	tryRegister(&_MAImageCaptioningCopyCaption, lib, "MAImageCaptioningCopyCaption")
	tryRegister(&_MAImageCaptioningCopyMetadataTagPath, lib, "MAImageCaptioningCopyMetadataTagPath")
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



// Returns the preference for audible media characteristics.
//
// Added in macOS 10.10.
// Returns the preference for audible media characteristics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAAudibleMediaCopyPreferredCharacteristics()
func MAAudibleMediaCopyPreferredCharacteristics() ArrayRef {
	return _MAAudibleMediaCopyPreferredCharacteristics()
}/* debug [functions.gen.go/function]: MAAudibleMediaCopyPreferredCharacteristics */

// Adds a preference for caption language to the stack of languages.
//
// Added in macOS 10.9.
// Adds a preference for caption language to the stack of languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceAddSelectedLanguage(_:_:)
func MACaptionAppearanceAddSelectedLanguage(domain MACaptionAppearanceDomain, language StringRef) bool {
	return _MACaptionAppearanceAddSelectedLanguage(domain, language)
}/* debug [functions.gen.go/function]: MACaptionAppearanceAddSelectedLanguage */

// MACaptionAppearanceCopyActiveProfileID is a MediaAccessibility function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyActiveProfileID()
func MACaptionAppearanceCopyActiveProfileID() StringRef {
	return _MACaptionAppearanceCopyActiveProfileID()
}/* debug [functions.gen.go/function]: MACaptionAppearanceCopyActiveProfileID */

// Returns the preference for the text highlight color.
//
// Added in macOS 10.9.
// Returns the preference for the text highlight color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyBackgroundColor(_:_:)
func MACaptionAppearanceCopyBackgroundColor(domain MACaptionAppearanceDomain, behavior unsafe.Pointer) ColorRef {
	return _MACaptionAppearanceCopyBackgroundColor(domain, behavior)
}/* debug [functions.gen.go/function]: MACaptionAppearanceCopyBackgroundColor */

// Returns the preferred font for the specified style of type.
//
// Added in macOS 10.9.
// Returns the preferred font for the specified style of type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyFontDescriptorForStyle(_:_:_:)
func MACaptionAppearanceCopyFontDescriptorForStyle(domain MACaptionAppearanceDomain, behavior unsafe.Pointer, fontStyle MACaptionAppearanceFontStyle) FontDescriptorRef {
	return _MACaptionAppearanceCopyFontDescriptorForStyle(domain, behavior, fontStyle)
}/* debug [functions.gen.go/function]: MACaptionAppearanceCopyFontDescriptorForStyle */

// Returns the preference for text color.
//
// Added in macOS 10.9.
// Returns the preference for text color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyForegroundColor(_:_:)
func MACaptionAppearanceCopyForegroundColor(domain MACaptionAppearanceDomain, behavior unsafe.Pointer) ColorRef {
	return _MACaptionAppearanceCopyForegroundColor(domain, behavior)
}/* debug [functions.gen.go/function]: MACaptionAppearanceCopyForegroundColor */

// Returns the preferences for captioning sounds.
//
// Added in macOS 10.9.
// Returns the preferences for captioning sounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyPreferredCaptioningMediaCharacteristics(_:)
func MACaptionAppearanceCopyPreferredCaptioningMediaCharacteristics(domain MACaptionAppearanceDomain) ArrayRef {
	return _MACaptionAppearanceCopyPreferredCaptioningMediaCharacteristics(domain)
}/* debug [functions.gen.go/function]: MACaptionAppearanceCopyPreferredCaptioningMediaCharacteristics */

// MACaptionAppearanceCopyProfileIDs is a MediaAccessibility function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyProfileIDs()
func MACaptionAppearanceCopyProfileIDs() ArrayRef {
	return _MACaptionAppearanceCopyProfileIDs()
}/* debug [functions.gen.go/function]: MACaptionAppearanceCopyProfileIDs */

// MACaptionAppearanceCopyProfileName is a MediaAccessibility function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyProfileName(_:)
func MACaptionAppearanceCopyProfileName(profileID StringRef) StringRef {
	return _MACaptionAppearanceCopyProfileName(profileID)
}/* debug [functions.gen.go/function]: MACaptionAppearanceCopyProfileName */

// Returns the preferred caption languages.
//
// Added in macOS 10.9.
// Returns the preferred caption languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopySelectedLanguages(_:)
func MACaptionAppearanceCopySelectedLanguages(domain MACaptionAppearanceDomain) ArrayRef {
	return _MACaptionAppearanceCopySelectedLanguages(domain)
}/* debug [functions.gen.go/function]: MACaptionAppearanceCopySelectedLanguages */

// Returns the preference for the caption window’s color.
//
// Added in macOS 10.9.
// Returns the preference for the caption window’s color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceCopyWindowColor(_:_:)
func MACaptionAppearanceCopyWindowColor(domain MACaptionAppearanceDomain, behavior unsafe.Pointer) ColorRef {
	return _MACaptionAppearanceCopyWindowColor(domain, behavior)
}/* debug [functions.gen.go/function]: MACaptionAppearanceCopyWindowColor */

// Informs accessibility clients when captions display onscreen.
//
// Added in macOS 10.9.
// Informs accessibility clients when captions display onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceDidDisplayCaptions(_:)
func MACaptionAppearanceDidDisplayCaptions(strings ArrayRef) {
	_MACaptionAppearanceDidDisplayCaptions(strings)
}/* debug [functions.gen.go/function]: MACaptionAppearanceDidDisplayCaptions */

// MACaptionAppearanceExecuteBlockForProfileID is a MediaAccessibility function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceExecuteBlockForProfileID(_:_:)
func MACaptionAppearanceExecuteBlockForProfileID(profileID StringRef) {
	_MACaptionAppearanceExecuteBlockForProfileID(profileID)
}/* debug [functions.gen.go/function]: MACaptionAppearanceExecuteBlockForProfileID */

// Returns the preference for the text highlight opacity.
//
// Added in macOS 10.9.
// Returns the preference for the text highlight opacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceGetBackgroundOpacity(_:_:)
func MACaptionAppearanceGetBackgroundOpacity(domain MACaptionAppearanceDomain, behavior unsafe.Pointer) float64 {
	return _MACaptionAppearanceGetBackgroundOpacity(domain, behavior)
}/* debug [functions.gen.go/function]: MACaptionAppearanceGetBackgroundOpacity */

// Returns the preferred type of captions to display.
//
// Added in macOS 10.9.
// Returns the preferred type of captions to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceGetDisplayType(_:)
func MACaptionAppearanceGetDisplayType(domain MACaptionAppearanceDomain) MACaptionAppearanceDisplayType {
	return _MACaptionAppearanceGetDisplayType(domain)
}/* debug [functions.gen.go/function]: MACaptionAppearanceGetDisplayType */

// Returns the preference for text opacity.
//
// Added in macOS 10.9.
// Returns the preference for text opacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceGetForegroundOpacity(_:_:)
func MACaptionAppearanceGetForegroundOpacity(domain MACaptionAppearanceDomain, behavior unsafe.Pointer) float64 {
	return _MACaptionAppearanceGetForegroundOpacity(domain, behavior)
}/* debug [functions.gen.go/function]: MACaptionAppearanceGetForegroundOpacity */

// Returns the preference for font scaling.
//
// Added in macOS 10.9.
// Returns the preference for font scaling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceGetRelativeCharacterSize(_:_:)
func MACaptionAppearanceGetRelativeCharacterSize(domain MACaptionAppearanceDomain, behavior unsafe.Pointer) float64 {
	return _MACaptionAppearanceGetRelativeCharacterSize(domain, behavior)
}/* debug [functions.gen.go/function]: MACaptionAppearanceGetRelativeCharacterSize */

// Returns the preference for text edge style.
//
// Added in macOS 10.9.
// Returns the preference for text edge style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceGetTextEdgeStyle(_:_:)
func MACaptionAppearanceGetTextEdgeStyle(domain MACaptionAppearanceDomain, behavior unsafe.Pointer) MACaptionAppearanceTextEdgeStyle {
	return _MACaptionAppearanceGetTextEdgeStyle(domain, behavior)
}/* debug [functions.gen.go/function]: MACaptionAppearanceGetTextEdgeStyle */

// Returns the preference for the overlay’s opacity.
//
// Added in macOS 10.9.
// Returns the preference for the overlay’s opacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceGetWindowOpacity(_:_:)
func MACaptionAppearanceGetWindowOpacity(domain MACaptionAppearanceDomain, behavior unsafe.Pointer) float64 {
	return _MACaptionAppearanceGetWindowOpacity(domain, behavior)
}/* debug [functions.gen.go/function]: MACaptionAppearanceGetWindowOpacity */

// Returns the radius of the caption window’s corners.
//
// Added in macOS 10.9.
// Returns the radius of the caption window’s corners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceGetWindowRoundedCornerRadius(_:_:)
func MACaptionAppearanceGetWindowRoundedCornerRadius(domain MACaptionAppearanceDomain, behavior unsafe.Pointer) float64 {
	return _MACaptionAppearanceGetWindowRoundedCornerRadius(domain, behavior)
}/* debug [functions.gen.go/function]: MACaptionAppearanceGetWindowRoundedCornerRadius */

// MACaptionAppearanceIsCustomized is a MediaAccessibility function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceIsCustomized(_:)
func MACaptionAppearanceIsCustomized(domain MACaptionAppearanceDomain) bool {
	return _MACaptionAppearanceIsCustomized(domain)
}/* debug [functions.gen.go/function]: MACaptionAppearanceIsCustomized */

// MACaptionAppearanceSetActiveProfileID is a MediaAccessibility function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceSetActiveProfileID(_:)
func MACaptionAppearanceSetActiveProfileID(profileID StringRef) {
	_MACaptionAppearanceSetActiveProfileID(profileID)
}/* debug [functions.gen.go/function]: MACaptionAppearanceSetActiveProfileID */

// Sets the preference for the type of caption.
//
// Added in macOS 10.9.
// Sets the preference for the type of caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MACaptionAppearanceSetDisplayType(_:_:)
func MACaptionAppearanceSetDisplayType(domain MACaptionAppearanceDomain, displayType MACaptionAppearanceDisplayType) {
	_MACaptionAppearanceSetDisplayType(domain, displayType)
}/* debug [functions.gen.go/function]: MACaptionAppearanceSetDisplayType */

// Returns a Boolean value that indicates whether the flashing lights setting is enabled on the device.
//
// Added in macOS 13.3.
// Returns a Boolean value that indicates whether the flashing lights setting is enabled on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MADimFlashingLightsEnabled()
func MADimFlashingLightsEnabled() bool {
	return _MADimFlashingLightsEnabled()
}/* debug [functions.gen.go/function]: MADimFlashingLightsEnabled */

// Returns an accessibility caption from an image’s metadata.
//
// Added in macOS 10.15.
// Returns an accessibility caption from an image’s metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAImageCaptioningCopyCaption(_:_:)
func MAImageCaptioningCopyCaption(url URLRef, error_ unsafe.Pointer) StringRef {
	return _MAImageCaptioningCopyCaption(url, error_)
}/* debug [functions.gen.go/function]: MAImageCaptioningCopyCaption */

// Returns the metadata tag path.
//
// Added in macOS 10.15.
// Returns the metadata tag path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAImageCaptioningCopyMetadataTagPath()
func MAImageCaptioningCopyMetadataTagPath() StringRef {
	return _MAImageCaptioningCopyMetadataTagPath()
}/* debug [functions.gen.go/function]: MAImageCaptioningCopyMetadataTagPath */

// Sets the accessibility caption for an image’s metadata.
//
// Added in macOS 10.15.
// Sets the accessibility caption for an image’s metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAImageCaptioningSetCaption(_:_:_:)
func MAImageCaptioningSetCaption(url URLRef, string_ StringRef, error_ unsafe.Pointer) bool {
	return _MAImageCaptioningSetCaption(url, string_, error_)
}/* debug [functions.gen.go/function]: MAImageCaptioningSetCaption */




