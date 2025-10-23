// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth_test

import (
	"github.com/tmc/appledocs/generated/corebluetooth"
)

// Suppress unused import errors
var _ = corebluetooth.NewCBUUID

// ExampleNewCBUUIDWithCFUUID demonstrates how to create a CBUUID instance using NewCBUUIDWithCFUUID.
// Creates a Core Bluetooth UUID object from a Core Foundation UUID object.
func ExampleNewCBUUIDWithCFUUID() {
	_ = corebluetooth.NewCBUUIDWithCFUUID(
		corebluetooth.UUIDRef{}, // theUUID UUIDRef
	)
	// Output:
}
// ExampleNewCBUUIDWithNSUUID demonstrates how to create a CBUUID instance using NewCBUUIDWithNSUUID.
// Creates a Core Bluetooth UUID object from a Foundation UUID object.
func ExampleNewCBUUIDWithNSUUID() {
	_ = corebluetooth.NewCBUUIDWithNSUUID(
		corebluetooth.UUID{}, // theUUID UUID
	)
	// Output:
}
// ExampleNewCBUUIDWithString demonstrates how to create a CBUUID instance using NewCBUUIDWithString.
// Creates a Core Bluetooth UUID object from a 16-, 32-, or 128-bit UUID string.
func ExampleNewCBUUIDWithString() {
	_ = corebluetooth.NewCBUUIDWithString(
		"theString", // theString string
	)
	// Output:
}

