// Code generated from Apple documentation for MediaToolbox. DO NOT EDIT.

package mediatoolbox
import (
"unsafe"
)

// Type aliases and typedefs
// MTAudioProcessingTapRef - An audio processing tap object.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTAudioProcessingTap
// MTAudioProcessingTapRef has base type: const struct opaqueMTAudioProcessingTap *
type MTAudioProcessingTapRef uintptr
// MTAudioProcessingTapFlags - Flags that indicate where to tap the audio.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTAudioProcessingTapFlags
// MTAudioProcessingTapFlags has base type: uint32_t
type MTAudioProcessingTapFlags uintptr
// MTAudioProcessingTapCreationFlags - Flags to use when creating audio processing taps.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTAudioProcessingTapCreationFlags
// MTAudioProcessingTapCreationFlags has base type: uint32_t
type MTAudioProcessingTapCreationFlags uintptr
// MTAudioProcessingTapFinalizeCallback - A finalization callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTAudioProcessingTapFinalizeCallback
// MTAudioProcessingTapFinalizeCallback is a callback function
// C type: void (*)(const struct opaqueMTAudioProcessingTap *)
type MTAudioProcessingTapFinalizeCallback = func(unsafe.Pointer)
// MTAudioProcessingTapInitCallback - An initialization callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTAudioProcessingTapInitCallback
// MTAudioProcessingTapInitCallback is a callback function
// C type: void (*)(const struct opaqueMTAudioProcessingTap *, void *, void **)
type MTAudioProcessingTapInitCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
// MTAudioProcessingTapPrepareCallback - An audio processing preparation callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTAudioProcessingTapPrepareCallback
// MTAudioProcessingTapPrepareCallback is a callback function
// C type: void (*)(const struct opaqueMTAudioProcessingTap *, long, const struct AudioStreamBasicDescription *)
type MTAudioProcessingTapPrepareCallback = func(unsafe.Pointer, int, unsafe.Pointer)
// MTAudioProcessingTapProcessCallback - An audio processing callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTAudioProcessingTapProcessCallback
// MTAudioProcessingTapProcessCallback is a callback function
// C type: void (*)(const struct opaqueMTAudioProcessingTap *, long, unsigned int, struct AudioBufferList *, long *, unsigned int *)
type MTAudioProcessingTapProcessCallback = func(unsafe.Pointer, int, uint32, unsafe.Pointer, int, uint32)
// MTAudioProcessingTapUnprepareCallback - An audio processing unpreparation callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaToolbox/MTAudioProcessingTapUnprepareCallback
// MTAudioProcessingTapUnprepareCallback is a callback function
// C type: void (*)(const struct opaqueMTAudioProcessingTap *)
type MTAudioProcessingTapUnprepareCallback = func(unsafe.Pointer)

