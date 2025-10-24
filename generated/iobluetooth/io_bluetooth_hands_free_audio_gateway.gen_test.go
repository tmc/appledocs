// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewBluetoothHandsFreeAudioGateway

// ExampleBluetoothHandsFreeAudioGateway_SendOKResponse demonstrates using SendOKResponse on a BluetoothHandsFreeAudioGateway instance.
// Sends a success message to a connected Bluetooth hands-free phone or headset.
func ExampleBluetoothHandsFreeAudioGateway_SendOKResponse() {
	obj := iobluetooth.NewBluetoothHandsFreeAudioGateway()
	obj.SendOKResponse()
	// Output:
	}

