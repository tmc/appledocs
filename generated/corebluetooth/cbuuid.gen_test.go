// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth_test

import (
	"github.com/tmc/appledocs/generated/corebluetooth"
)

// Suppress unused import errors
var _ = corebluetooth.NewCBUUID

// ExampleNewCBUUIDWithString demonstrates how to create a CBUUID instance using NewCBUUIDWithString.
// Creates a Core Bluetooth UUID object from a 16-, 32-, or 128-bit UUID string.
func ExampleNewCBUUIDWithString() {
	_ = corebluetooth.NewCBUUIDWithString(
		"theString", // theString string
	)
	// Output:
}

