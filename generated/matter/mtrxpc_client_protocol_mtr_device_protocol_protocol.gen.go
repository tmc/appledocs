// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTRXPCClientProtocol_MTRDevice is the MTRXPCClientProtocol_MTRDevice protocol interface.
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
// See: doc://com.apple.matter/documentation/Matter/MTRXPCClientProtocol_MTRDevice
type PMTRXPCClientProtocol_MTRDevice interface {
	// Required methods
	DeviceInternalStateUpdated(nodeID objc.IObject /* cross-framework: NSNumber */, dictionary objc.IObject /* cross-framework: NSDictionary */)/* debug [protocol_interface/required_method]: DeviceInternalStateUpdated */
}
