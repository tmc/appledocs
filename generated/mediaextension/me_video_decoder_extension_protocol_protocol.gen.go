// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMEVideoDecoderExtension is the MEVideoDecoderExtension protocol interface.
//
// A protocol that defines a factory to create new video decoders for a codec type that the extension implements.
//
// Availability:
//   - macOS 14.0+
//
// See: doc://com.apple.mediaextension/documentation/MediaExtension/MEVideoDecoderExtension
type PMEVideoDecoderExtension interface {
	// Required methods
	Init() unsafe.Pointer/* debug [protocol_interface/required_method]: Init */
	VideoDecoderWithCodecTypeVideoFormatDescriptionVideoDecoderSpecificationsExtensionDecoderPixelBufferManagerError(codecType VideoCodecType /* not a class type */, videoFormatDescription VideoFormatDescriptionRef /* not a class type */, videoDecoderSpecifications foundation.IDictionary, extensionDecoderPixelBufferManager IMEVideoDecoderPixelBufferManager, error_ unsafe.Pointer) unsafe.Pointer/* debug [protocol_interface/required_method]: VideoDecoderWithCodecTypeVideoFormatDescriptionVideoDecoderSpecificationsExtensionDecoderPixelBufferManagerError */
}
