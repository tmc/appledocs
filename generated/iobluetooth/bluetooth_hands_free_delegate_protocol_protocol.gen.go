// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PBluetoothHandsFreeDelegate is the IOBluetoothHandsFreeDelegate protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.iobluetooth/documentation/IOBluetooth/IOBluetoothHandsFreeDelegate
type PBluetoothHandsFreeDelegate interface {
	// Optional methods
	HandsFreeConnected(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */)
	HasHandsFreeConnected() bool
	HandsFreeDisconnected(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */)
	HasHandsFreeDisconnected() bool
	HandsFreeScoConnectionClosed(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */)
	HasHandsFreeScoConnectionClosed() bool
	HandsFreeScoConnectionOpened(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */)
	HasHandsFreeScoConnectionOpened() bool
}

// BluetoothHandsFreeDelegate is a delegate implementation builder for the PBluetoothHandsFreeDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type BluetoothHandsFreeDelegate struct {
	_HandsFreeConnected func(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */)
	_HandsFreeDisconnected func(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */)
	_HandsFreeScoConnectionClosed func(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */)
	_HandsFreeScoConnectionOpened func(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */)
}

// SetHandsFreeConnected sets the handler for the HandsFreeConnected delegate method.
func (d *BluetoothHandsFreeDelegate) SetHandsFreeConnected(f func(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */)) {
	d._HandsFreeConnected = f
}

// SetHandsFreeDisconnected sets the handler for the HandsFreeDisconnected delegate method.
func (d *BluetoothHandsFreeDelegate) SetHandsFreeDisconnected(f func(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */)) {
	d._HandsFreeDisconnected = f
}

// SetHandsFreeScoConnectionClosed sets the handler for the HandsFreeScoConnectionClosed delegate method.
func (d *BluetoothHandsFreeDelegate) SetHandsFreeScoConnectionClosed(f func(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */)) {
	d._HandsFreeScoConnectionClosed = f
}

// SetHandsFreeScoConnectionOpened sets the handler for the HandsFreeScoConnectionOpened delegate method.
func (d *BluetoothHandsFreeDelegate) SetHandsFreeScoConnectionOpened(f func(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */)) {
	d._HandsFreeScoConnectionOpened = f
}

// HandsFreeConnected implements the PBluetoothHandsFreeDelegate interface.
func (d *BluetoothHandsFreeDelegate) HandsFreeConnected(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */) {
	if d._HandsFreeConnected != nil {
		d._HandsFreeConnected(device, status)
	}
}

// HasHandsFreeConnected returns true if a handler for HandsFreeConnected has been set.
func (d *BluetoothHandsFreeDelegate) HasHandsFreeConnected() bool {
	return d._HandsFreeConnected != nil
}

// HandsFreeDisconnected implements the PBluetoothHandsFreeDelegate interface.
func (d *BluetoothHandsFreeDelegate) HandsFreeDisconnected(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */) {
	if d._HandsFreeDisconnected != nil {
		d._HandsFreeDisconnected(device, status)
	}
}

// HasHandsFreeDisconnected returns true if a handler for HandsFreeDisconnected has been set.
func (d *BluetoothHandsFreeDelegate) HasHandsFreeDisconnected() bool {
	return d._HandsFreeDisconnected != nil
}

// HandsFreeScoConnectionClosed implements the PBluetoothHandsFreeDelegate interface.
func (d *BluetoothHandsFreeDelegate) HandsFreeScoConnectionClosed(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */) {
	if d._HandsFreeScoConnectionClosed != nil {
		d._HandsFreeScoConnectionClosed(device, status)
	}
}

// HasHandsFreeScoConnectionClosed returns true if a handler for HandsFreeScoConnectionClosed has been set.
func (d *BluetoothHandsFreeDelegate) HasHandsFreeScoConnectionClosed() bool {
	return d._HandsFreeScoConnectionClosed != nil
}

// HandsFreeScoConnectionOpened implements the PBluetoothHandsFreeDelegate interface.
func (d *BluetoothHandsFreeDelegate) HandsFreeScoConnectionOpened(device IOBluetoothHandsFree, status objc.IObject /* cross-framework: NSNumber */) {
	if d._HandsFreeScoConnectionOpened != nil {
		d._HandsFreeScoConnectionOpened(device, status)
	}
}

// HasHandsFreeScoConnectionOpened returns true if a handler for HandsFreeScoConnectionOpened has been set.
func (d *BluetoothHandsFreeDelegate) HasHandsFreeScoConnectionOpened() bool {
	return d._HandsFreeScoConnectionOpened != nil
}
