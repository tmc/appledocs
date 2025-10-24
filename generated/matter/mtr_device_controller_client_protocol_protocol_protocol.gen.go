// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PMTRDeviceControllerClientProtocol is the MTRDeviceControllerClientProtocol protocol interface.
//
// Availability:
//   - Mac Catalyst 16.1+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+
//
// See: doc://com.apple.matter/documentation/Matter/MTRDeviceControllerClientProtocol
type PMTRDeviceControllerClientProtocol interface {
	// Required methods
	HandleReportWithControllerNodeIdValuesError(controller objc.IObject, nodeId uint64, values objc.IObject, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: HandleReportWithControllerNodeIdValuesError */
}
