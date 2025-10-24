// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"unsafe"
)

// PMEFormatReader is the MEFormatReader protocol interface.
//
// A protocol that defines the requirements for a format reader, which represents a single media asset.
//
// Availability:
//   - macOS 14.0+
//
// See: doc://com.apple.mediaextension/documentation/MediaExtension/MEFormatReader
type PMEFormatReader interface {
	// Required methods
	LoadFileInfoWithCompletionHandler(completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: LoadFileInfoWithCompletionHandler */
	LoadMetadataWithCompletionHandler(completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: LoadMetadataWithCompletionHandler */
	LoadTrackReadersWithCompletionHandler(completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: LoadTrackReadersWithCompletionHandler */
	// Optional methods
	ParseAdditionalFragmentsWithCompletionHandler(completionHandler unsafe.Pointer)
	HasParseAdditionalFragmentsWithCompletionHandler() bool
}
