// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PEAAccessoryDelegate is the EAAccessoryDelegate protocol interface.
//
// A protocol that defines an optional method for receiving notifications when the associated accessory object is disconnected.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.13+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.externalaccessory/documentation/ExternalAccessory/EAAccessoryDelegate
type PEAAccessoryDelegate interface {
	// Optional methods
	AccessoryDidDisconnect(accessory IEAAccessory)
	HasAccessoryDidDisconnect() bool
}

// EAAccessoryDelegate is a delegate implementation builder for the PEAAccessoryDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type EAAccessoryDelegate struct {
	_AccessoryDidDisconnect func(accessory IEAAccessory)
}

// SetAccessoryDidDisconnect sets the handler for the AccessoryDidDisconnect delegate method.
//
// Tells the delegate that the specified accessory was disconnected from the device.
func (d *EAAccessoryDelegate) SetAccessoryDidDisconnect(f func(accessory IEAAccessory)) {
	d._AccessoryDidDisconnect = f
}

// AccessoryDidDisconnect implements the PEAAccessoryDelegate interface.
func (d *EAAccessoryDelegate) AccessoryDidDisconnect(accessory IEAAccessory) {
	if d._AccessoryDidDisconnect != nil {
		d._AccessoryDidDisconnect(accessory)
	}
}

// HasAccessoryDidDisconnect returns true if a handler for AccessoryDidDisconnect has been set.
func (d *EAAccessoryDelegate) HasAccessoryDidDisconnect() bool {
	return d._AccessoryDidDisconnect != nil
}
