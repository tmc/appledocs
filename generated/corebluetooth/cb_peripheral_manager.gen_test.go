// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth_test

import (
	"github.com/tmc/appledocs/generated/corebluetooth"
)

// Suppress unused import errors
var _ = corebluetooth.NewCBPeripheralManager

// ExampleNewCBPeripheralManager demonstrates how to create a CBPeripheralManager instance.
// Initializes the peripheral manager without a delegate.
func ExampleNewCBPeripheralManager() {
	_ = corebluetooth.NewCBPeripheralManager()
	// Output:
}
// ExampleCBPeripheralManager_RemoveAllServices demonstrates using RemoveAllServices on a CBPeripheralManager instance.
// Removes all published services from the local GATT database.
func ExampleCBPeripheralManager_RemoveAllServices() {
	obj := corebluetooth.NewCBPeripheralManager()
	obj.RemoveAllServices()
	// Output:
	}

// ExampleCBPeripheralManager_StopAdvertising demonstrates using StopAdvertising on a CBPeripheralManager instance.
// Stops advertising peripheral manager data.
func ExampleCBPeripheralManager_StopAdvertising() {
	obj := corebluetooth.NewCBPeripheralManager()
	obj.StopAdvertising()
	// Output:
	}

