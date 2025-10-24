// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewBluetoothSDPServiceAttribute

// ExampleBluetoothSDPServiceAttribute_GetDataElement demonstrates using GetDataElement on a BluetoothSDPServiceAttribute instance.
// Returns the data element for the target service attribute.
func ExampleBluetoothSDPServiceAttribute_GetDataElement() {
	obj := iobluetooth.NewBluetoothSDPServiceAttribute()
	_ = obj.GetDataElement()
	// Output:
	}

// ExampleBluetoothSDPServiceAttribute_GetAttributeID demonstrates using GetAttributeID on a BluetoothSDPServiceAttribute instance.
// Returns the attribute ID for the target service attribute.
func ExampleBluetoothSDPServiceAttribute_GetAttributeID() {
	obj := iobluetooth.NewBluetoothSDPServiceAttribute()
	_ = obj.GetAttributeID()
	// Output:
	}

// ExampleBluetoothSDPServiceAttribute_GetIDDataElement demonstrates using GetIDDataElement on a BluetoothSDPServiceAttribute instance.
// Returns the data element representing the attribute ID for the target service attribute.
func ExampleBluetoothSDPServiceAttribute_GetIDDataElement() {
	obj := iobluetooth.NewBluetoothSDPServiceAttribute()
	_ = obj.GetIDDataElement()
	// Output:
	}

