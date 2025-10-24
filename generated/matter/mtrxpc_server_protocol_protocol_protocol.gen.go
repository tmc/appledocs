// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTRXPCServerProtocol is the MTRXPCServerProtocol protocol interface.
//
// Availability:
//   - Mac Catalyst 18.3+
//   - iOS 18.3+
//   - iPadOS 18.3+
//   - macOS 15.3+
//   - tvOS 18.3+
//   - visionOS 2.3+
//   - watchOS 11.3+
//
// See: doc://com.apple.matter/documentation/Matter/MTRXPCServerProtocol
type PMTRXPCServerProtocol interface {
	// Optional methods
	DeviceControllerCheckInWithContext(controller foundation.UUID, context objc.IObject /* cross-framework: NSDictionary */)
	HasDeviceControllerCheckInWithContext() bool
}
