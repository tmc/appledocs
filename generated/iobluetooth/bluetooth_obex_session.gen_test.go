// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewBluetoothOBEXSession

// ExampleNewBluetoothOBEXSessionWithDeviceChannelID demonstrates how to create a BluetoothOBEXSession instance using NewBluetoothOBEXSessionWithDeviceChannelID.
// Initializes a Bluetooth-based OBEX Session using a Bluetooth device.
func ExampleNewBluetoothOBEXSessionWithDeviceChannelID() {
	_ = iobluetooth.NewBluetoothOBEXSessionWithDeviceChannelID(
		iobluetooth.IOBluetoothDevice{},        // inDevice IOBluetoothDevice
		iobluetooth.BluetoothRFCOMMChannelID{}, // inChannelID BluetoothRFCOMMChannelID
	)
	// Output:
}

// ExampleNewBluetoothOBEXSessionWithSDPServiceRecord demonstrates how to create a BluetoothOBEXSession instance using NewBluetoothOBEXSessionWithSDPServiceRecord.
// Initializes a Bluetooth-based OBEX Session using an SDP service record.
func ExampleNewBluetoothOBEXSessionWithSDPServiceRecord() {
	_ = iobluetooth.NewBluetoothOBEXSessionWithSDPServiceRecord(
		iobluetooth.IOBluetoothSDPServiceRecord{}, // inSDPServiceRecord IOBluetoothSDPServiceRecord
	)
	// Output:
}
