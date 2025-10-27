// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PExternalSyncDeviceDelegate is the AVExternalSyncDeviceDelegate protocol interface.
//
// Defines an interface for delegates of   to respond to events that occur when connecting, calibrating, and disconnecting external sync devices.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVExternalSyncDeviceDelegate
type PExternalSyncDeviceDelegate interface {
	// Optional methods
	ExternalSyncDeviceFailedWithError(device IAVExternalSyncDevice, error_ foundation.foundation.INSError)
	HasExternalSyncDeviceFailedWithError() bool
	ExternalSyncDeviceStatusDidChange(device IAVExternalSyncDevice)
	HasExternalSyncDeviceStatusDidChange() bool
}
