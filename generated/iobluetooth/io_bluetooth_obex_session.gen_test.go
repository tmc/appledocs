// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewBluetoothOBEXSession

// ExampleBluetoothOBEXSession_CloseTransportConnection demonstrates using CloseTransportConnection on a BluetoothOBEXSession instance.
// An OBEXSession override. When this is called by the session baseclass, we will close the transport connection if it is opened. In our case, it will be the RFCOMM channel that needs closing.
func ExampleBluetoothOBEXSession_CloseTransportConnection() {
	obj := iobluetooth.NewBluetoothOBEXSession()
	_ = obj.CloseTransportConnection()
	// Output:
	}

// ExampleBluetoothOBEXSession_GetDevice demonstrates using GetDevice on a BluetoothOBEXSession instance.
// Get the Bluetooth Device being used by the session object.
func ExampleBluetoothOBEXSession_GetDevice() {
	obj := iobluetooth.NewBluetoothOBEXSession()
	_ = obj.GetDevice()
	// Output:
	}

// ExampleBluetoothOBEXSession_GetRFCOMMChannel demonstrates using GetRFCOMMChannel on a BluetoothOBEXSession instance.
// Get the Bluetooth RFCOMM channel being used by the session object.
func ExampleBluetoothOBEXSession_GetRFCOMMChannel() {
	obj := iobluetooth.NewBluetoothOBEXSession()
	_ = obj.GetRFCOMMChannel()
	// Output:
	}

// ExampleBluetoothOBEXSession_HasOpenTransportConnection demonstrates using HasOpenTransportConnection on a BluetoothOBEXSession instance.
// An OBEXSession override. When this is called by the session baseclass, we will return whether or not we have a transport connection established to another OBEX server/client. In our case we will tell whether or not the RFCOMM channel to a remote device is still open.
func ExampleBluetoothOBEXSession_HasOpenTransportConnection() {
	obj := iobluetooth.NewBluetoothOBEXSession()
	_ = obj.HasOpenTransportConnection()
	// Output:
	}

// ExampleBluetoothOBEXSession_IsSessionTargetAMac demonstrates using IsSessionTargetAMac on a BluetoothOBEXSession instance.
// Tells whether the target device is a Mac by checking its service record.
func ExampleBluetoothOBEXSession_IsSessionTargetAMac() {
	obj := iobluetooth.NewBluetoothOBEXSession()
	_ = obj.IsSessionTargetAMac()
	// Output:
	}

// ExampleBluetoothOBEXSession_RestartTransmission demonstrates using RestartTransmission on a BluetoothOBEXSession instance.
// If the transmission was stopped due to the lack of buffers this call restarts it.
func ExampleBluetoothOBEXSession_RestartTransmission() {
	obj := iobluetooth.NewBluetoothOBEXSession()
	obj.RestartTransmission()
	// Output:
	}

// ExampleBluetoothOBEXSession_SendBufferTroughChannel demonstrates using SendBufferTroughChannel on a BluetoothOBEXSession instance.
// Sends the next block of data through the rfcomm channel.
func ExampleBluetoothOBEXSession_SendBufferTroughChannel() {
	obj := iobluetooth.NewBluetoothOBEXSession()
	_ = obj.SendBufferTroughChannel()
	// Output:
	}

