// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"unsafe"
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
	DevicePropertiesForPropertiesError(properties unsafe.Pointer, outError unsafe.Pointer) IOExtensionDeviceProperties/* debug [protocol_interface/required_method]: DevicePropertiesForPropertiesError */
	SetDevicePropertiesError(deviceProperties ExtensionDeviceProperties, outError unsafe.Pointer) bool/* debug [protocol_interface/required_method]: SetDevicePropertiesError */
}
