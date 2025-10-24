// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"unsafe"
)

// PGCDevicePhysicalInputStateDiff is the GCDevicePhysicalInputStateDiff protocol interface.
//
// The common functions for objects that contain the differences between a current and previous input state object.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.gamecontroller/documentation/GameController/GCDevicePhysicalInputStateDiff
type PGCDevicePhysicalInputStateDiff interface {
	// Required methods
	ChangeForElement(element unsafe.Pointer) GCDevicePhysicalInputElementChange/* debug [protocol_interface/required_method]: ChangeForElement */
	ChangedElements() unsafe.Pointer/* debug [protocol_interface/required_method]: ChangedElements */
}
