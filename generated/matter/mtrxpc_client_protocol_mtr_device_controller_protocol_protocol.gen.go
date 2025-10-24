// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTRXPCClientProtocol_MTRDeviceController is the MTRXPCClientProtocol_MTRDeviceController protocol interface.
//
// Availability:
//   - Mac Catalyst 18.2+
//   - iOS 18.2+
//   - iPadOS 18.2+
//   - macOS 15.2+
//   - tvOS 18.2+
//   - visionOS 2.2+
//   - watchOS 11.2+
//
// See: doc://com.apple.matter/documentation/Matter/MTRXPCClientProtocol_MTRDeviceController
type PMTRXPCClientProtocol_MTRDeviceController interface {
	// Optional methods
	ControllerControllerConfigurationUpdated(controller foundation.UUID, configuration objc.IObject /* cross-framework: NSDictionary */)
	HasControllerControllerConfigurationUpdated() bool
}
