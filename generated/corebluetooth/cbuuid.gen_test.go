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
		corebluetooth.UUIDRef /* not a class type */{}, // theUUID UUIDRef /* not a class type */
	)
	// Output:
}

