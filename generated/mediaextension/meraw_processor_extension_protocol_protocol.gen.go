// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"unsafe"
)

// PMERAWProcessorExtension is the MERAWProcessorExtension protocol interface.
//
// A protocol that defines a factory to create RAW processors for a codec type that the extension implements.
//
// Availability:
//   - macOS 15.0+
//
// See: doc://com.apple.mediaextension/documentation/MediaExtension/MERAWProcessorExtension
type PMERAWProcessorExtension interface {
	// Required methods
	Init() unsafe.Pointer/* debug [protocol_interface/required_method]: Init */
	ProcessorWithFormatDescriptionExtensionPixelBufferManagerError(formatDescription VideoFormatDescriptionRef /* not a class type */, extensionPixelBufferManager IMERAWProcessorPixelBufferManager, error_ unsafe.Pointer) unsafe.Pointer/* debug [protocol_interface/required_method]: ProcessorWithFormatDescriptionExtensionPixelBufferManagerError */
}
