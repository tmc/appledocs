// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewBluetoothDeviceInquiry

// ExampleBluetoothDeviceInquiry_ClearFoundDevices demonstrates using ClearFoundDevices on a BluetoothDeviceInquiry instance.
// Removes all found devices from the inquiry object.
func ExampleBluetoothDeviceInquiry_ClearFoundDevices() {
	obj := iobluetooth.NewBluetoothDeviceInquiry()
	obj.ClearFoundDevices()
	// Output:
	}

// ExampleBluetoothDeviceInquiry_FoundDevices demonstrates using FoundDevices on a BluetoothDeviceInquiry instance.
// Returns found IOBluetoothDevice objects as an array.
func ExampleBluetoothDeviceInquiry_FoundDevices() {
	obj := iobluetooth.NewBluetoothDeviceInquiry()
	_ = obj.FoundDevices()
	// Output:
	}

// ExampleBluetoothDeviceInquiry_Start demonstrates using Start on a BluetoothDeviceInquiry instance.
// Tells inquiry object to begin the inquiry and name updating process, if specified.
func ExampleBluetoothDeviceInquiry_Start() {
	obj := iobluetooth.NewBluetoothDeviceInquiry()
	_ = obj.Start()
	// Output:
	}

// ExampleBluetoothDeviceInquiry_Stop demonstrates using Stop on a BluetoothDeviceInquiry instance.
// Halts the inquiry object. Could either stop the search for new devices, or the updating of found device names.
func ExampleBluetoothDeviceInquiry_Stop() {
	obj := iobluetooth.NewBluetoothDeviceInquiry()
	_ = obj.Stop()
	// Output:
	}

