// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"unsafe"
)

// PIOExtensionStreamSource is the CMIOExtensionStreamSource protocol interface.
//
// A protocol for objects that act as stream sources.
//
// Availability:
//   - Mac Catalyst 15.4+
//   - macOS 12.3+
//
// See: doc://com.apple.coremediaio/documentation/CoreMediaIO/CMIOExtensionStreamSource
type PIOExtensionStreamSource interface {
	// Required methods
	AuthorizedToStartStreamForClient(client ExtensionClient) bool/* debug [protocol_interface/required_method]: AuthorizedToStartStreamForClient */
	SetStreamPropertiesError(streamProperties ExtensionStreamProperties, outError unsafe.Pointer) bool/* debug [protocol_interface/required_method]: SetStreamPropertiesError */
	StartStreamAndReturnError(outError unsafe.Pointer) bool/* debug [protocol_interface/required_method]: StartStreamAndReturnError */
	StopStreamAndReturnError(outError unsafe.Pointer) bool/* debug [protocol_interface/required_method]: StopStreamAndReturnError */
	StreamPropertiesForPropertiesError(properties unsafe.Pointer, outError unsafe.Pointer) IOExtensionStreamProperties/* debug [protocol_interface/required_method]: StreamPropertiesForPropertiesError */
}
