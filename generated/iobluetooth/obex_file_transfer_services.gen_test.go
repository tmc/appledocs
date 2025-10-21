// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewOBEXFileTransferServices

// ExampleNewOBEXFileTransferServicesWithOBEXSession demonstrates how to create a OBEXFileTransferServices instance using NewOBEXFileTransferServicesWithOBEXSession.
// Create a new OBEXFileTransferServices object
func ExampleNewOBEXFileTransferServicesWithOBEXSession() {
	_ = iobluetooth.NewOBEXFileTransferServicesWithOBEXSession(
		iobluetooth.IOBluetoothOBEXSession{}, // inOBEXSession IOBluetoothOBEXSession
	)
	// Output:
}
