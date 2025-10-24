// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PBluetoothDeviceInquiryDelegate is the IOBluetoothDeviceInquiryDelegate protocol interface.
//
// This category on NSObject describes the delegate methods for the IOBluetoothDeviceInquiry object. All methods are optional, but it is highly recommended you implement them all. Do NOT invoke remote name requests on found IOBluetoothDevice objects unless the inquiry object has been stopped. Doing so may deadlock your process.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.iobluetooth/documentation/IOBluetooth/IOBluetoothDeviceInquiryDelegate
type PBluetoothDeviceInquiryDelegate interface {
	// Optional methods
	DeviceInquiryCompleteErrorAborted(sender IOBluetoothDeviceInquiry, error_ int, aborted bool)
	HasDeviceInquiryCompleteErrorAborted() bool
	DeviceInquiryDeviceFoundDevice(sender IOBluetoothDeviceInquiry, device IOBluetoothDevice)
	HasDeviceInquiryDeviceFoundDevice() bool
	DeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining(sender IOBluetoothDeviceInquiry, device IOBluetoothDevice, devicesRemaining uint32 /* not a class type */)
	HasDeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining() bool
	DeviceInquiryStarted(sender IOBluetoothDeviceInquiry)
	HasDeviceInquiryStarted() bool
	DeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining(sender IOBluetoothDeviceInquiry, devicesRemaining uint32 /* not a class type */)
	HasDeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining() bool
}

// BluetoothDeviceInquiryDelegate is a delegate implementation builder for the PBluetoothDeviceInquiryDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type BluetoothDeviceInquiryDelegate struct {
	_DeviceInquiryCompleteErrorAborted func(sender IOBluetoothDeviceInquiry, error_ int, aborted bool)
	_DeviceInquiryDeviceFoundDevice func(sender IOBluetoothDeviceInquiry, device IOBluetoothDevice)
	_DeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining func(sender IOBluetoothDeviceInquiry, device IOBluetoothDevice, devicesRemaining uint32 /* not a class type */)
	_DeviceInquiryStarted func(sender IOBluetoothDeviceInquiry)
	_DeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining func(sender IOBluetoothDeviceInquiry, devicesRemaining uint32 /* not a class type */)
}

// SetDeviceInquiryCompleteErrorAborted sets the handler for the DeviceInquiryCompleteErrorAborted delegate method.
func (d *BluetoothDeviceInquiryDelegate) SetDeviceInquiryCompleteErrorAborted(f func(sender IOBluetoothDeviceInquiry, error_ int, aborted bool)) {
	d._DeviceInquiryCompleteErrorAborted = f
}

// SetDeviceInquiryDeviceFoundDevice sets the handler for the DeviceInquiryDeviceFoundDevice delegate method.
func (d *BluetoothDeviceInquiryDelegate) SetDeviceInquiryDeviceFoundDevice(f func(sender IOBluetoothDeviceInquiry, device IOBluetoothDevice)) {
	d._DeviceInquiryDeviceFoundDevice = f
}

// SetDeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining sets the handler for the DeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining delegate method.
func (d *BluetoothDeviceInquiryDelegate) SetDeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining(f func(sender IOBluetoothDeviceInquiry, device IOBluetoothDevice, devicesRemaining uint32 /* not a class type */)) {
	d._DeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining = f
}

// SetDeviceInquiryStarted sets the handler for the DeviceInquiryStarted delegate method.
func (d *BluetoothDeviceInquiryDelegate) SetDeviceInquiryStarted(f func(sender IOBluetoothDeviceInquiry)) {
	d._DeviceInquiryStarted = f
}

// SetDeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining sets the handler for the DeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining delegate method.
func (d *BluetoothDeviceInquiryDelegate) SetDeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining(f func(sender IOBluetoothDeviceInquiry, devicesRemaining uint32 /* not a class type */)) {
	d._DeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining = f
}

// DeviceInquiryCompleteErrorAborted implements the PBluetoothDeviceInquiryDelegate interface.
func (d *BluetoothDeviceInquiryDelegate) DeviceInquiryCompleteErrorAborted(sender IOBluetoothDeviceInquiry, error_ int, aborted bool) {
	if d._DeviceInquiryCompleteErrorAborted != nil {
		d._DeviceInquiryCompleteErrorAborted(sender, error_, aborted)
	}
}

// HasDeviceInquiryCompleteErrorAborted returns true if a handler for DeviceInquiryCompleteErrorAborted has been set.
func (d *BluetoothDeviceInquiryDelegate) HasDeviceInquiryCompleteErrorAborted() bool {
	return d._DeviceInquiryCompleteErrorAborted != nil
}

// DeviceInquiryDeviceFoundDevice implements the PBluetoothDeviceInquiryDelegate interface.
func (d *BluetoothDeviceInquiryDelegate) DeviceInquiryDeviceFoundDevice(sender IOBluetoothDeviceInquiry, device IOBluetoothDevice) {
	if d._DeviceInquiryDeviceFoundDevice != nil {
		d._DeviceInquiryDeviceFoundDevice(sender, device)
	}
}

// HasDeviceInquiryDeviceFoundDevice returns true if a handler for DeviceInquiryDeviceFoundDevice has been set.
func (d *BluetoothDeviceInquiryDelegate) HasDeviceInquiryDeviceFoundDevice() bool {
	return d._DeviceInquiryDeviceFoundDevice != nil
}

// DeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining implements the PBluetoothDeviceInquiryDelegate interface.
func (d *BluetoothDeviceInquiryDelegate) DeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining(sender IOBluetoothDeviceInquiry, device IOBluetoothDevice, devicesRemaining uint32 /* not a class type */) {
	if d._DeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining != nil {
		d._DeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining(sender, device, devicesRemaining)
	}
}

// HasDeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining returns true if a handler for DeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining has been set.
func (d *BluetoothDeviceInquiryDelegate) HasDeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining() bool {
	return d._DeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining != nil
}

// DeviceInquiryStarted implements the PBluetoothDeviceInquiryDelegate interface.
func (d *BluetoothDeviceInquiryDelegate) DeviceInquiryStarted(sender IOBluetoothDeviceInquiry) {
	if d._DeviceInquiryStarted != nil {
		d._DeviceInquiryStarted(sender)
	}
}

// HasDeviceInquiryStarted returns true if a handler for DeviceInquiryStarted has been set.
func (d *BluetoothDeviceInquiryDelegate) HasDeviceInquiryStarted() bool {
	return d._DeviceInquiryStarted != nil
}

// DeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining implements the PBluetoothDeviceInquiryDelegate interface.
func (d *BluetoothDeviceInquiryDelegate) DeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining(sender IOBluetoothDeviceInquiry, devicesRemaining uint32 /* not a class type */) {
	if d._DeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining != nil {
		d._DeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining(sender, devicesRemaining)
	}
}

// HasDeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining returns true if a handler for DeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining has been set.
func (d *BluetoothDeviceInquiryDelegate) HasDeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining() bool {
	return d._DeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining != nil
}
