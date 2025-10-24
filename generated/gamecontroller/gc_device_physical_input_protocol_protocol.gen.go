// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"unsafe"
)

// PGCDevicePhysicalInput is the GCDevicePhysicalInput protocol interface.
//
// The common properties and methods for objects that represent the input profile of a device.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.gamecontroller/documentation/GameController/GCDevicePhysicalInput
type PGCDevicePhysicalInput interface {
	// Required methods
	Capture() unsafe.Pointer/* debug [protocol_interface/required_method]: Capture */
	NextInputState() unsafe.Pointer/* debug [protocol_interface/required_method]: NextInputState */
}
