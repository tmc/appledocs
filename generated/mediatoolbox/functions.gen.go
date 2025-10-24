// Code generated from Apple documentation for MediaToolbox. DO NOT EDIT.

package mediatoolbox

/* debug [functions.gen.go]: Generating 7 functions for MediaToolbox */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// MediaToolbox Functions (7 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_MTAudioProcessingTapCreate func(AllocatorRef, unsafe.Pointer, MTAudioProcessingTapCreationFlags, unsafe.Pointer) unsafe.Pointer
	_MTAudioProcessingTapGetSourceAudio func(MTAudioProcessingTapRef, ItemCount, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MTAudioProcessingTapGetStorage func(MTAudioProcessingTapRef) unsafe.Pointer
	_MTAudioProcessingTapGetTypeID func() TypeID
	_MTCopyLocalizedNameForMediaSubType func(MediaType, unsafe.Pointer) StringRef
	_MTCopyLocalizedNameForMediaType func(MediaType) StringRef
	_MTRegisterProfessionalVideoWorkflowFormatReaders func()
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_MTAudioProcessingTapCreate, lib, "MTAudioProcessingTapCreate")
	tryRegister(&_MTAudioProcessingTapGetSourceAudio, lib, "MTAudioProcessingTapGetSourceAudio")
	tryRegister(&_MTAudioProcessingTapGetStorage, lib, "MTAudioProcessingTapGetStorage")
	tryRegister(&_MTAudioProcessingTapGetTypeID, lib, "MTAudioProcessingTapGetTypeID")
	tryRegister(&_MTCopyLocalizedNameForMediaSubType, lib, "MTCopyLocalizedNameForMediaSubType")
	tryRegister(&_MTCopyLocalizedNameForMediaType, lib, "MTCopyLocalizedNameForMediaType")
	tryRegister(&_MTRegisterProfessionalVideoWorkflowFormatReaders, lib, "MTRegisterProfessionalVideoWorkflowFormatReaders")
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



// Creates a new audio processing tap.
//
// Added in macOS 10.9.
// Creates a new audio processing tap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTAudioProcessingTapCreate(_:_:_:_:)
func MTAudioProcessingTapCreate(allocator AllocatorRef, callbacks unsafe.Pointer, flags MTAudioProcessingTapCreationFlags, tapOut unsafe.Pointer) unsafe.Pointer {
	return _MTAudioProcessingTapCreate(allocator, callbacks, flags, tapOut)
}/* debug [functions.gen.go/function]: MTAudioProcessingTapCreate */

// Retrieves source audio for an audio processing tap.
//
// Added in macOS 10.9.
// Retrieves source audio for an audio processing tap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTAudioProcessingTapGetSourceAudio(_:_:_:_:_:_:)
func MTAudioProcessingTapGetSourceAudio(tap MTAudioProcessingTapRef, numberFrames ItemCount, bufferListInOut unsafe.Pointer, flagsOut unsafe.Pointer, timeRangeOut unsafe.Pointer, numberFramesOut unsafe.Pointer) unsafe.Pointer {
	return _MTAudioProcessingTapGetSourceAudio(tap, numberFrames, bufferListInOut, flagsOut, timeRangeOut, numberFramesOut)
}/* debug [functions.gen.go/function]: MTAudioProcessingTapGetSourceAudio */

// Retrieves a custom storage pointer for an audio processing tap.
//
// Added in macOS 10.9.
// Retrieves a custom storage pointer for an audio processing tap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTAudioProcessingTapGetStorage(_:)
func MTAudioProcessingTapGetStorage(tap MTAudioProcessingTapRef) unsafe.Pointer {
	return _MTAudioProcessingTapGetStorage(tap)
}/* debug [functions.gen.go/function]: MTAudioProcessingTapGetStorage */

// Retrieves the type identifier for this audio processing tap.
//
// Added in macOS 10.9.
// Retrieves the type identifier for this audio processing tap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTAudioProcessingTapGetTypeID()
func MTAudioProcessingTapGetTypeID() TypeID {
	return _MTAudioProcessingTapGetTypeID()
}/* debug [functions.gen.go/function]: MTAudioProcessingTapGetTypeID */

// Returns a localized name for the specified media type and subtype.
//
// Added in macOS 10.9.
// Returns a localized name for the specified media type and subtype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTCopyLocalizedNameForMediaSubType(_:_:)
func MTCopyLocalizedNameForMediaSubType(mediaType MediaType, mediaSubType unsafe.Pointer) StringRef {
	return _MTCopyLocalizedNameForMediaSubType(mediaType, mediaSubType)
}/* debug [functions.gen.go/function]: MTCopyLocalizedNameForMediaSubType */

// Returns a localized name for the specified media type.
//
// Added in macOS 10.9.
// Returns a localized name for the specified media type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTCopyLocalizedNameForMediaType(_:)
func MTCopyLocalizedNameForMediaType(mediaType MediaType) StringRef {
	return _MTCopyLocalizedNameForMediaType(mediaType)
}/* debug [functions.gen.go/function]: MTCopyLocalizedNameForMediaType */

// Enables the use of media format readers that support professional video workflows.
//
// Added in macOS 10.10.
// Enables the use of media format readers that support professional video workflows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTRegisterProfessionalVideoWorkflowFormatReaders()
func MTRegisterProfessionalVideoWorkflowFormatReaders() {
	_MTRegisterProfessionalVideoWorkflowFormatReaders()
}/* debug [functions.gen.go/function]: MTRegisterProfessionalVideoWorkflowFormatReaders */




