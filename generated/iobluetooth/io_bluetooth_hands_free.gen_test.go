// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewBluetoothHandsFree

// ExampleBluetoothHandsFree_Connect demonstrates using Connect on a BluetoothHandsFree instance.
// Connect to the device
func ExampleBluetoothHandsFree_Connect() {
	obj := iobluetooth.NewBluetoothHandsFree()
	obj.Connect()
	// Output:
	}

// ExampleBluetoothHandsFree_ConnectSCO demonstrates using ConnectSCO on a BluetoothHandsFree instance.
// Open a SCO connection with the device
func ExampleBluetoothHandsFree_ConnectSCO() {
	obj := iobluetooth.NewBluetoothHandsFree()
	obj.ConnectSCO()
	// Output:
	}

// ExampleBluetoothHandsFree_Disconnect demonstrates using Disconnect on a BluetoothHandsFree instance.
// Disconnect from the device
func ExampleBluetoothHandsFree_Disconnect() {
	obj := iobluetooth.NewBluetoothHandsFree()
	obj.Disconnect()
	// Output:
	}

// ExampleBluetoothHandsFree_DisconnectSCO demonstrates using DisconnectSCO on a BluetoothHandsFree instance.
// Disconnect the SCO connection with the device
func ExampleBluetoothHandsFree_DisconnectSCO() {
	obj := iobluetooth.NewBluetoothHandsFree()
	obj.DisconnectSCO()
	// Output:
	}

// ExampleBluetoothHandsFree_IsSCOConnected demonstrates using IsSCOConnected on a BluetoothHandsFree instance.
// Determine if there is a SCO connection to the device
func ExampleBluetoothHandsFree_IsSCOConnected() {
	obj := iobluetooth.NewBluetoothHandsFree()
	_ = obj.IsSCOConnected()
	// Output:
	}

