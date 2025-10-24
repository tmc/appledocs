// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTRXPCServerProtocol_MTRDevice is the MTRXPCServerProtocol_MTRDevice protocol interface.
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
// See: doc://com.apple.matter/documentation/Matter/MTRXPCServerProtocol_MTRDevice
type PMTRXPCServerProtocol_MTRDevice interface {
	// Required methods
	DeviceControllerNodeIDOpenCommissioningWindowWithSetupPasscodeDiscriminatorDurationCompletion(controller foundation.UUID, nodeID objc.IObject /* cross-framework: NSNumber */, setupPasscode objc.IObject /* cross-framework: NSNumber */, discriminator objc.IObject /* cross-framework: NSNumber */, duration objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: DeviceControllerNodeIDOpenCommissioningWindowWithSetupPasscodeDiscriminatorDurationCompletion */
}
