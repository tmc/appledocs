// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"unsafe"
)

// PMEVideoDecoder is the MEVideoDecoder protocol interface.
//
// A protocol that defines the requirements for a video decoder.
//
// Availability:
//   - macOS 14.0+
//
// See: doc://com.apple.mediaextension/documentation/MediaExtension/MEVideoDecoder
type PMEVideoDecoder interface {
	// Required methods
	DecodeFrameFromSampleBufferOptionsCompletionHandler(sampleBuffer SampleBufferRef /* not a class type */, options IMEDecodeFrameOptions, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: DecodeFrameFromSampleBufferOptionsCompletionHandler */
	// Optional methods
	CanAcceptFormatDescription(formatDescription FormatDescriptionRef /* not a class type */) bool
	HasCanAcceptFormatDescription() bool
}
