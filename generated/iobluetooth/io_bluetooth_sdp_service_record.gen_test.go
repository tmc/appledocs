// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewBluetoothSDPServiceRecord

// ExampleBluetoothSDPServiceRecord_GetSDPServiceRecordRef demonstrates using GetSDPServiceRecordRef on a BluetoothSDPServiceRecord instance.
// Returns an IOBluetoothSDPServiceRecordRef representation of the target IOBluetoothSDPServiceRecord object.
func ExampleBluetoothSDPServiceRecord_GetSDPServiceRecordRef() {
	obj := iobluetooth.NewBluetoothSDPServiceRecord()
	_ = obj.GetSDPServiceRecordRef()
	// Output:
	}

// ExampleBluetoothSDPServiceRecord_GetServiceName demonstrates using GetServiceName on a BluetoothSDPServiceRecord instance.
// Returns the name of the service.
func ExampleBluetoothSDPServiceRecord_GetServiceName() {
	obj := iobluetooth.NewBluetoothSDPServiceRecord()
	_ = obj.GetServiceName()
	// Output:
	}

// ExampleBluetoothSDPServiceRecord_HandsFreeSupportedFeatures demonstrates using HandsFreeSupportedFeatures on a BluetoothSDPServiceRecord instance.
func ExampleBluetoothSDPServiceRecord_HandsFreeSupportedFeatures() {
	obj := iobluetooth.NewBluetoothSDPServiceRecord()
	_ = obj.HandsFreeSupportedFeatures()
	// Output:
	}

// ExampleBluetoothSDPServiceRecord_RemoveServiceRecord demonstrates using RemoveServiceRecord on a BluetoothSDPServiceRecord instance.
// Removes the service from the local SDP server.
func ExampleBluetoothSDPServiceRecord_RemoveServiceRecord() {
	obj := iobluetooth.NewBluetoothSDPServiceRecord()
	_ = obj.RemoveServiceRecord()
	// Output:
	}

