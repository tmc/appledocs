// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PIOExtensionProviderSource is the CMIOExtensionProviderSource protocol interface.
//
// A protocol for objects that act as provider sources.
//
// Availability:
//   - Mac Catalyst 15.4+
//   - macOS 12.3+
//
// See: doc://com.apple.coremediaio/documentation/CoreMediaIO/CMIOExtensionProviderSource
type PIOExtensionProviderSource interface {
	// Required methods
	ConnectClientError(client ExtensionClient, outError foundation.foundation.INSError) bool
	DisconnectClient(client ExtensionClient)
	ProviderPropertiesForPropertiesError(properties unsafe.Pointer, outError foundation.foundation.INSError) IOExtensionProviderProperties
	SetProviderPropertiesError(providerProperties ExtensionProviderProperties, outError foundation.foundation.INSError) bool
}
