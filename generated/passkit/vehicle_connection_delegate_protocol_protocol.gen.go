// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PVehicleConnectionDelegate is the PKVehicleConnectionDelegate protocol interface.
//
// Availability:
//   - Mac Catalyst 15.4+
//   - iOS 15.4+
//   - iPadOS 15.4+
//   - macOS +
//   - visionOS 1.0+
//   - watchOS 8.5+
//
// See: doc://com.apple.passkit/documentation/PassKit/PKVehicleConnectionDelegate
type PVehicleConnectionDelegate interface {
	// Required methods
	SessionDidChangeConnectionState(newState VehicleConnectionSessionConnectionState)/* debug [protocol_interface/required_method]: SessionDidChangeConnectionState */
	SessionDidReceiveData(data objc.IObject /* cross-framework: NSData */)/* debug [protocol_interface/required_method]: SessionDidReceiveData */
}

// VehicleConnectionDelegate is a delegate implementation builder for the PVehicleConnectionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type VehicleConnectionDelegate struct {
	_SessionDidChangeConnectionState func(newState VehicleConnectionSessionConnectionState)
	_SessionDidReceiveData func(data objc.IObject /* cross-framework: NSData */)
}

// SetSessionDidChangeConnectionState sets the handler for the SessionDidChangeConnectionState delegate method.
func (d *VehicleConnectionDelegate) SetSessionDidChangeConnectionState(f func(newState VehicleConnectionSessionConnectionState)) {
	d._SessionDidChangeConnectionState = f
}

// SetSessionDidReceiveData sets the handler for the SessionDidReceiveData delegate method.
func (d *VehicleConnectionDelegate) SetSessionDidReceiveData(f func(data objc.IObject /* cross-framework: NSData */)) {
	d._SessionDidReceiveData = f
}

// SessionDidChangeConnectionState implements the PVehicleConnectionDelegate interface.
func (d *VehicleConnectionDelegate) SessionDidChangeConnectionState(newState VehicleConnectionSessionConnectionState) {
	if d._SessionDidChangeConnectionState != nil {
		d._SessionDidChangeConnectionState(newState)
	}
}

// HasSessionDidChangeConnectionState returns true if a handler for SessionDidChangeConnectionState has been set.
func (d *VehicleConnectionDelegate) HasSessionDidChangeConnectionState() bool {
	return d._SessionDidChangeConnectionState != nil
}

// SessionDidReceiveData implements the PVehicleConnectionDelegate interface.
func (d *VehicleConnectionDelegate) SessionDidReceiveData(data objc.IObject /* cross-framework: NSData */) {
	if d._SessionDidReceiveData != nil {
		d._SessionDidReceiveData(data)
	}
}

// HasSessionDidReceiveData returns true if a handler for SessionDidReceiveData has been set.
func (d *VehicleConnectionDelegate) HasSessionDidReceiveData() bool {
	return d._SessionDidReceiveData != nil
}
