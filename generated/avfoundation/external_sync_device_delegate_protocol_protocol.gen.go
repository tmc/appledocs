// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"
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
	ExternalSyncDeviceFailedWithError(device IAVExternalSyncDevice, error_ Error)
	HasExternalSyncDeviceFailedWithError() bool
	ExternalSyncDeviceStatusDidChange(device IAVExternalSyncDevice)
	HasExternalSyncDeviceStatusDidChange() bool
}

// ExternalSyncDeviceDelegate is a delegate implementation builder for the PExternalSyncDeviceDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ExternalSyncDeviceDelegate struct {
	_ExternalSyncDeviceFailedWithError func(device IAVExternalSyncDevice, error_ Error)
	_ExternalSyncDeviceStatusDidChange func(device IAVExternalSyncDevice)
}

// SetExternalSyncDeviceFailedWithError sets the handler for the ExternalSyncDeviceFailedWithError delegate method.
func (d *ExternalSyncDeviceDelegate) SetExternalSyncDeviceFailedWithError(f func(device IAVExternalSyncDevice, error_ Error)) {
	d._ExternalSyncDeviceFailedWithError = f
}

// SetExternalSyncDeviceStatusDidChange sets the handler for the ExternalSyncDeviceStatusDidChange delegate method.
//
// Informs your delegate when the external sync device status has changed.
func (d *ExternalSyncDeviceDelegate) SetExternalSyncDeviceStatusDidChange(f func(device IAVExternalSyncDevice)) {
	d._ExternalSyncDeviceStatusDidChange = f
}

// ExternalSyncDeviceFailedWithError implements the PExternalSyncDeviceDelegate interface.
func (d *ExternalSyncDeviceDelegate) ExternalSyncDeviceFailedWithError(device IAVExternalSyncDevice, error_ Error) {
	if d._ExternalSyncDeviceFailedWithError != nil {
		d._ExternalSyncDeviceFailedWithError(device, error_)
	}
}

// HasExternalSyncDeviceFailedWithError returns true if a handler for ExternalSyncDeviceFailedWithError has been set.
func (d *ExternalSyncDeviceDelegate) HasExternalSyncDeviceFailedWithError() bool {
	return d._ExternalSyncDeviceFailedWithError != nil
}

// ExternalSyncDeviceStatusDidChange implements the PExternalSyncDeviceDelegate interface.
func (d *ExternalSyncDeviceDelegate) ExternalSyncDeviceStatusDidChange(device IAVExternalSyncDevice) {
	if d._ExternalSyncDeviceStatusDidChange != nil {
		d._ExternalSyncDeviceStatusDidChange(device)
	}
}

// HasExternalSyncDeviceStatusDidChange returns true if a handler for ExternalSyncDeviceStatusDidChange has been set.
func (d *ExternalSyncDeviceDelegate) HasExternalSyncDeviceStatusDidChange() bool {
	return d._ExternalSyncDeviceStatusDidChange != nil
}
