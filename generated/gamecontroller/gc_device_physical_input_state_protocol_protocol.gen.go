// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PGCDevicePhysicalInputState is the GCDevicePhysicalInputState protocol interface.
//
// The common properties for physical devices with elements.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.gamecontroller/documentation/GameController/GCDevicePhysicalInputState
type PGCDevicePhysicalInputState interface {
	// Required methods
	ObjectForKeyedSubscript(key objc.IObject /* cross-framework: NSString */) unsafe.Pointer/* debug [protocol_interface/required_method]: ObjectForKeyedSubscript */
}
