// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewBluetoothDevicePair

// ExampleBluetoothDevicePair_Device demonstrates using Device on a BluetoothDevicePair instance.
// Get the IOBluetoothDevice being used by the object.
func ExampleBluetoothDevicePair_Device() {
	obj := iobluetooth.NewBluetoothDevicePair()
	_ = obj.Device()
	// Output:
	}

// ExampleBluetoothDevicePair_Start demonstrates using Start on a BluetoothDevicePair instance.
// Kicks off the pairing with the device.
func ExampleBluetoothDevicePair_Start() {
	obj := iobluetooth.NewBluetoothDevicePair()
	_ = obj.Start()
	// Output:
	}

// ExampleBluetoothDevicePair_Stop demonstrates using Stop on a BluetoothDevicePair instance.
// Stops the current pairing. Removes the delegate and disconnects if device was connected.
func ExampleBluetoothDevicePair_Stop() {
	obj := iobluetooth.NewBluetoothDevicePair()
	obj.Stop()
	// Output:
	}

