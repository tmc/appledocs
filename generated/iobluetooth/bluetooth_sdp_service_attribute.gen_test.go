// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewBluetoothSDPServiceAttribute

// ExampleNewBluetoothSDPServiceAttributeWithIDAttributeElement demonstrates how to create a BluetoothSDPServiceAttribute instance using NewBluetoothSDPServiceAttributeWithIDAttributeElement.
// Initializes a new service attribute with the given ID and data element.
func ExampleNewBluetoothSDPServiceAttributeWithIDAttributeElement() {
	_ = iobluetooth.NewBluetoothSDPServiceAttributeWithIDAttributeElement(
		iobluetooth.BluetoothSDPServiceAttributeID{}, // newAttributeID BluetoothSDPServiceAttributeID
		iobluetooth.IOBluetoothSDPDataElement{},      // attributeElement IOBluetoothSDPDataElement
	)
	// Output:
}
