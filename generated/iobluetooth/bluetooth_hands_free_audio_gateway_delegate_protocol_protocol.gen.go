// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PBluetoothHandsFreeAudioGatewayDelegate is the IOBluetoothHandsFreeAudioGatewayDelegate protocol interface.
//
// A set of optional methods for receiving information about status changes for a connected Bluetooth hands-free phone or headset.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.iobluetooth/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGatewayDelegate
type PBluetoothHandsFreeAudioGatewayDelegate interface {
	// Optional methods
	HandsFreeHangup(device IOBluetoothHandsFreeAudioGateway, hangup objc.IObject /* cross-framework: NSNumber */)
	HasHandsFreeHangup() bool
	HandsFreeRedial(device IOBluetoothHandsFreeAudioGateway, redial objc.IObject /* cross-framework: NSNumber */)
	HasHandsFreeRedial() bool
}

// BluetoothHandsFreeAudioGatewayDelegate is a delegate implementation builder for the PBluetoothHandsFreeAudioGatewayDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type BluetoothHandsFreeAudioGatewayDelegate struct {
	_HandsFreeHangup func(device IOBluetoothHandsFreeAudioGateway, hangup objc.IObject /* cross-framework: NSNumber */)
	_HandsFreeRedial func(device IOBluetoothHandsFreeAudioGateway, redial objc.IObject /* cross-framework: NSNumber */)
}

// SetHandsFreeHangup sets the handler for the HandsFreeHangup delegate method.
//
// Tells the delegate the connected Bluetooth hands-free phone or headset is sending a hang-up signal.
func (d *BluetoothHandsFreeAudioGatewayDelegate) SetHandsFreeHangup(f func(device IOBluetoothHandsFreeAudioGateway, hangup objc.IObject /* cross-framework: NSNumber */)) {
	d._HandsFreeHangup = f
}

// SetHandsFreeRedial sets the handler for the HandsFreeRedial delegate method.
//
// Tells the delegate the connected Bluetooth hands-free phone or headset is redialing the last phone number.
func (d *BluetoothHandsFreeAudioGatewayDelegate) SetHandsFreeRedial(f func(device IOBluetoothHandsFreeAudioGateway, redial objc.IObject /* cross-framework: NSNumber */)) {
	d._HandsFreeRedial = f
}

// HandsFreeHangup implements the PBluetoothHandsFreeAudioGatewayDelegate interface.
func (d *BluetoothHandsFreeAudioGatewayDelegate) HandsFreeHangup(device IOBluetoothHandsFreeAudioGateway, hangup objc.IObject /* cross-framework: NSNumber */) {
	if d._HandsFreeHangup != nil {
		d._HandsFreeHangup(device, hangup)
	}
}

// HasHandsFreeHangup returns true if a handler for HandsFreeHangup has been set.
func (d *BluetoothHandsFreeAudioGatewayDelegate) HasHandsFreeHangup() bool {
	return d._HandsFreeHangup != nil
}

// HandsFreeRedial implements the PBluetoothHandsFreeAudioGatewayDelegate interface.
func (d *BluetoothHandsFreeAudioGatewayDelegate) HandsFreeRedial(device IOBluetoothHandsFreeAudioGateway, redial objc.IObject /* cross-framework: NSNumber */) {
	if d._HandsFreeRedial != nil {
		d._HandsFreeRedial(device, redial)
	}
}

// HasHandsFreeRedial returns true if a handler for HandsFreeRedial has been set.
func (d *BluetoothHandsFreeAudioGatewayDelegate) HasHandsFreeRedial() bool {
	return d._HandsFreeRedial != nil
}
