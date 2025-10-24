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
		iobluetooth.BluetoothSDPUUID16 /* typedef */{}, // uuid16 BluetoothSDPUUID16 /* typedef */
	)
	// Output:
}
// ExampleNewBluetoothSDPUUIDWithUUID32 demonstrates how to create a BluetoothSDPUUID instance using NewBluetoothSDPUUIDWithUUID32.
// Creates a new 32-bit IOBluetoothSDPUUID with the given UUID32
func ExampleNewBluetoothSDPUUIDWithUUID32() {
	_ = iobluetooth.NewBluetoothSDPUUIDWithUUID32(
		iobluetooth.BluetoothSDPUUID32 /* typedef */{}, // uuid32 BluetoothSDPUUID32 /* typedef */
	)
	// Output:
}
// ExampleBluetoothSDPUUID_ClassForArchiver demonstrates using ClassForArchiver on a BluetoothSDPUUID instance.
func ExampleBluetoothSDPUUID_ClassForArchiver() {
	obj := iobluetooth.NewBluetoothSDPUUID()
	_ = obj.ClassForArchiver()
	// Output:
	}

// ExampleBluetoothSDPUUID_ClassForCoder demonstrates using ClassForCoder on a BluetoothSDPUUID instance.
func ExampleBluetoothSDPUUID_ClassForCoder() {
	obj := iobluetooth.NewBluetoothSDPUUID()
	_ = obj.ClassForCoder()
	// Output:
	}

// ExampleBluetoothSDPUUID_ClassForPortCoder demonstrates using ClassForPortCoder on a BluetoothSDPUUID instance.
func ExampleBluetoothSDPUUID_ClassForPortCoder() {
	obj := iobluetooth.NewBluetoothSDPUUID()
	_ = obj.ClassForPortCoder()
	// Output:
	}

