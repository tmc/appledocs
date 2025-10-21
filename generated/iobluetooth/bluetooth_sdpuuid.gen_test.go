// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewBluetoothSDPUUID

// ExampleNewBluetoothSDPUUIDWithUUID16 demonstrates how to create a BluetoothSDPUUID instance using NewBluetoothSDPUUIDWithUUID16.
// Initializes a new 16-bit IOBluetoothSDPUUID with the given UUID16
func ExampleNewBluetoothSDPUUIDWithUUID16() {
	_ = iobluetooth.NewBluetoothSDPUUIDWithUUID16(
		iobluetooth.BluetoothSDPUUID16{}, // uuid16 BluetoothSDPUUID16
	)
	// Output:
}
// ExampleNewBluetoothSDPUUIDWithUUID32 demonstrates how to create a BluetoothSDPUUID instance using NewBluetoothSDPUUIDWithUUID32.
// Creates a new 32-bit IOBluetoothSDPUUID with the given UUID32
func ExampleNewBluetoothSDPUUIDWithUUID32() {
	_ = iobluetooth.NewBluetoothSDPUUIDWithUUID32(
		iobluetooth.BluetoothSDPUUID32{}, // uuid32 BluetoothSDPUUID32
	)
	// Output:
}
