// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"unsafe"
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
	ConnectClientError(client ExtensionClient, outError unsafe.Pointer) bool/* debug [protocol_interface/required_method]: ConnectClientError */
	DisconnectClient(client ExtensionClient)/* debug [protocol_interface/required_method]: DisconnectClient */
	ProviderPropertiesForPropertiesError(properties unsafe.Pointer, outError unsafe.Pointer) IOExtensionProviderProperties/* debug [protocol_interface/required_method]: ProviderPropertiesForPropertiesError */
	SetProviderPropertiesError(providerProperties ExtensionProviderProperties, outError unsafe.Pointer) bool/* debug [protocol_interface/required_method]: SetProviderPropertiesError */
}
