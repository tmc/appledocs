// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewBluetoothDevicePair

// ExampleNewBluetoothDevicePairWithDevice demonstrates how to create a BluetoothDevicePair instance using NewBluetoothDevicePairWithDevice.
// Creates an autorelease IOBluetoothDevicePair object with a device as the pairing target.
func ExampleNewBluetoothDevicePairWithDevice() {
	_ = iobluetooth.NewBluetoothDevicePairWithDevice(
		iobluetooth.IOBluetoothDevice{}, // device IOBluetoothDevice
	)
	// Output:
}
