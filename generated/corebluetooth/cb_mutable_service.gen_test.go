// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth_test

import (
	"github.com/tmc/appledocs/generated/corebluetooth"
)

// Suppress unused import errors
var _ = corebluetooth.NewCBMutableService

// ExampleNewCBMutableServiceWithTypePrimary demonstrates how to create a CBMutableService instance using NewCBMutableServiceWithTypePrimary.
// Creates a newly initialized mutable service specified by UUID and service type.
func ExampleNewCBMutableServiceWithTypePrimary() {
	_ = corebluetooth.NewCBMutableServiceWithTypePrimary(
		corebluetooth.CBUUID{}, // UUID CBUUID
		false, // isPrimary bool
	)
	// Output:
}
