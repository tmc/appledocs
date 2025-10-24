// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth_test

import (
	"github.com/tmc/appledocs/generated/corebluetooth"
)

// Suppress unused import errors
var _ = corebluetooth.NewCBCentralManager

// ExampleNewCBCentralManager demonstrates how to create a CBCentralManager instance.
// Initializes the central manager without a delegate.
func ExampleNewCBCentralManager() {
	_ = corebluetooth.NewCBCentralManager()
	// Output:
}
// ExampleCBCentralManager_StopScan demonstrates using StopScan on a CBCentralManager instance.
// Asks the central manager to stop scanning for peripherals.
func ExampleCBCentralManager_StopScan() {
	obj := corebluetooth.NewCBCentralManager()
	obj.StopScan()
	// Output:
	}

