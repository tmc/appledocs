// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewBluetoothDevice


// ExampleNewBluetoothDeviceWithAddressString demonstrates how to create a BluetoothDevice instance using NewBluetoothDeviceWithAddressString.
// Returns the IOBluetoothDevice object for the given BluetoothDeviceAddress
func ExampleNewBluetoothDeviceWithAddressString() {
	_ = iobluetooth.NewBluetoothDeviceWithAddressString(
		"address", // address string
	)
	// Output:
}



