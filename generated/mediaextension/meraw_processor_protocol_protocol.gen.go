// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"unsafe"
)

// PMERAWProcessor is the MERAWProcessor protocol interface.
//
// A protocol that defines the requirements for a RAW processor.
//
// Availability:
//   - macOS 15.0+
//
// See: doc://com.apple.mediaextension/documentation/MediaExtension/MERAWProcessor
type PMERAWProcessor interface {
	// Required methods
	ProcessFrameFromImageBufferCompletionHandler(inputFrame PixelBufferRef /* not a class type */, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: ProcessFrameFromImageBufferCompletionHandler */
}
