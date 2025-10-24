// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"unsafe"
)

// PMEFormatReaderExtension is the MEFormatReaderExtension protocol interface.
//
// A protocol that defines a factory to create a new format reader with a byte source.
//
// Availability:
//   - macOS 14.0+
//
// See: doc://com.apple.mediaextension/documentation/MediaExtension/MEFormatReaderExtension
type PMEFormatReaderExtension interface {
	// Required methods
	FormatReaderWithByteSourceOptionsError(primaryByteSource IMEByteSource, options IMEFormatReaderInstantiationOptions, error_ unsafe.Pointer) unsafe.Pointer/* debug [protocol_interface/required_method]: FormatReaderWithByteSourceOptionsError */
	Init() unsafe.Pointer/* debug [protocol_interface/required_method]: Init */
}
