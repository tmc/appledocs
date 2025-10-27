// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
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
	AuthorizedToStartStreamForClient(client ExtensionClient) bool
	SetStreamPropertiesError(streamProperties ExtensionStreamProperties, outError foundation.foundation.INSError) bool
	StartStreamAndReturnError(outError foundation.foundation.INSError) bool
	StopStreamAndReturnError(outError foundation.foundation.INSError) bool
	StreamPropertiesForPropertiesError(properties unsafe.Pointer, outError foundation.foundation.INSError) IOExtensionStreamProperties
}
