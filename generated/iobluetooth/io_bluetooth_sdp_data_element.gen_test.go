// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth_test

import (
	"github.com/tmc/appledocs/generated/iobluetooth"
)

// Suppress unused import errors
var _ = iobluetooth.NewBluetoothSDPDataElement

// ExampleBluetoothSDPDataElement_GetArrayValue demonstrates using GetArrayValue on a BluetoothSDPDataElement instance.
// If the data element is represented by an array object, it returns the value as an NSArray.
func ExampleBluetoothSDPDataElement_GetArrayValue() {
	obj := iobluetooth.NewBluetoothSDPDataElement()
	_ = obj.GetArrayValue()
	// Output:
	}

// ExampleBluetoothSDPDataElement_GetDataValue demonstrates using GetDataValue on a BluetoothSDPDataElement instance.
// If the data element is represented by a data object, it returns the value as an NSData.
func ExampleBluetoothSDPDataElement_GetDataValue() {
	obj := iobluetooth.NewBluetoothSDPDataElement()
	_ = obj.GetDataValue()
	// Output:
	}

// ExampleBluetoothSDPDataElement_GetNumberValue demonstrates using GetNumberValue on a BluetoothSDPDataElement instance.
// If the data element is represented by a number, it returns the value as an NSNumber.
func ExampleBluetoothSDPDataElement_GetNumberValue() {
	obj := iobluetooth.NewBluetoothSDPDataElement()
	_ = obj.GetNumberValue()
	// Output:
	}

// ExampleBluetoothSDPDataElement_GetSDPDataElementRef demonstrates using GetSDPDataElementRef on a BluetoothSDPDataElement instance.
// Returns an IOBluetoothSDPDataElementRef representation of the target IOBluetoothSDPDataElement object.
func ExampleBluetoothSDPDataElement_GetSDPDataElementRef() {
	obj := iobluetooth.NewBluetoothSDPDataElement()
	_ = obj.GetSDPDataElementRef()
	// Output:
	}

// ExampleBluetoothSDPDataElement_GetSize demonstrates using GetSize on a BluetoothSDPDataElement instance.
// Returns the size in bytes of the target data element.
func ExampleBluetoothSDPDataElement_GetSize() {
	obj := iobluetooth.NewBluetoothSDPDataElement()
	_ = obj.GetSize()
	// Output:
	}

// ExampleBluetoothSDPDataElement_GetSizeDescriptor demonstrates using GetSizeDescriptor on a BluetoothSDPDataElement instance.
// Returns the SDP spec defined data element size descriptor for the target data element.
func ExampleBluetoothSDPDataElement_GetSizeDescriptor() {
	obj := iobluetooth.NewBluetoothSDPDataElement()
	_ = obj.GetSizeDescriptor()
	// Output:
	}

// ExampleBluetoothSDPDataElement_GetStringValue demonstrates using GetStringValue on a BluetoothSDPDataElement instance.
// If the data element is represented by a string object, it returns the value as an NSString.
func ExampleBluetoothSDPDataElement_GetStringValue() {
	obj := iobluetooth.NewBluetoothSDPDataElement()
	_ = obj.GetStringValue()
	// Output:
	}

// ExampleBluetoothSDPDataElement_GetTypeDescriptor demonstrates using GetTypeDescriptor on a BluetoothSDPDataElement instance.
// Returns the SDP spec defined data element type descriptor for the target data element.
func ExampleBluetoothSDPDataElement_GetTypeDescriptor() {
	obj := iobluetooth.NewBluetoothSDPDataElement()
	_ = obj.GetTypeDescriptor()
	// Output:
	}

// ExampleBluetoothSDPDataElement_GetUUIDValue demonstrates using GetUUIDValue on a BluetoothSDPDataElement instance.
// If the data element is a UUID (type 3), it returns the value as an IOBluetoothSDPUUID.
func ExampleBluetoothSDPDataElement_GetUUIDValue() {
	obj := iobluetooth.NewBluetoothSDPDataElement()
	_ = obj.GetUUIDValue()
	// Output:
	}

// ExampleBluetoothSDPDataElement_GetValue demonstrates using GetValue on a BluetoothSDPDataElement instance.
// Returns the object value of the data element.
func ExampleBluetoothSDPDataElement_GetValue() {
	obj := iobluetooth.NewBluetoothSDPDataElement()
	_ = obj.GetValue()
	// Output:
	}

