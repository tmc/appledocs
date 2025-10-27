// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PIOExtensionDeviceSource is the CMIOExtensionDeviceSource protocol interface.
//
// A protocol for objects that act as device sources.
//
// Availability:
//   - Mac Catalyst 15.4+
//   - macOS 12.3+
//
// See: doc://com.apple.coremediaio/documentation/CoreMediaIO/CMIOExtensionDeviceSource
type PIOExtensionDeviceSource interface {
	// Required methods
	DevicePropertiesForPropertiesError(properties unsafe.Pointer, outError foundation.foundation.INSError) IOExtensionDeviceProperties
	SetDevicePropertiesError(deviceProperties ExtensionDeviceProperties, outError foundation.foundation.INSError) bool
}
