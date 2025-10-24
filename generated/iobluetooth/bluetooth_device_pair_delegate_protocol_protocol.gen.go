// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PBluetoothDevicePairDelegate is the IOBluetoothDevicePairDelegate protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.iobluetooth/documentation/IOBluetooth/IOBluetoothDevicePairDelegate
type PBluetoothDevicePairDelegate interface {
	// Optional methods
	DevicePairingConnected(sender objc.IObject)
	HasDevicePairingConnected() bool
	DevicePairingConnecting(sender objc.IObject)
	HasDevicePairingConnecting() bool
	DevicePairingFinishedError(sender objc.IObject, error_ int)
	HasDevicePairingFinishedError() bool
	DevicePairingPINCodeRequest(sender objc.IObject)
	HasDevicePairingPINCodeRequest() bool
	DevicePairingStarted(sender objc.IObject)
	HasDevicePairingStarted() bool
	DevicePairingUserConfirmationRequestNumericValue(sender objc.IObject, numericValue BluetoothNumericValue /* typedef */)
	HasDevicePairingUserConfirmationRequestNumericValue() bool
	DevicePairingUserPasskeyNotificationPasskey(sender objc.IObject, passkey BluetoothPasskey /* typedef */)
	HasDevicePairingUserPasskeyNotificationPasskey() bool
	DeviceSimplePairingCompleteStatus(sender objc.IObject, status BluetoothHCIEventStatus /* typedef */)
	HasDeviceSimplePairingCompleteStatus() bool
}

// BluetoothDevicePairDelegate is a delegate implementation builder for the PBluetoothDevicePairDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type BluetoothDevicePairDelegate struct {
	_DevicePairingConnected func(sender objc.IObject)
	_DevicePairingConnecting func(sender objc.IObject)
	_DevicePairingFinishedError func(sender objc.IObject, error_ int)
	_DevicePairingPINCodeRequest func(sender objc.IObject)
	_DevicePairingStarted func(sender objc.IObject)
	_DevicePairingUserConfirmationRequestNumericValue func(sender objc.IObject, numericValue BluetoothNumericValue /* typedef */)
	_DevicePairingUserPasskeyNotificationPasskey func(sender objc.IObject, passkey BluetoothPasskey /* typedef */)
	_DeviceSimplePairingCompleteStatus func(sender objc.IObject, status BluetoothHCIEventStatus /* typedef */)
}

// SetDevicePairingConnected sets the handler for the DevicePairingConnected delegate method.
func (d *BluetoothDevicePairDelegate) SetDevicePairingConnected(f func(sender objc.IObject)) {
	d._DevicePairingConnected = f
}

// SetDevicePairingConnecting sets the handler for the DevicePairingConnecting delegate method.
func (d *BluetoothDevicePairDelegate) SetDevicePairingConnecting(f func(sender objc.IObject)) {
	d._DevicePairingConnecting = f
}

// SetDevicePairingFinishedError sets the handler for the DevicePairingFinishedError delegate method.
func (d *BluetoothDevicePairDelegate) SetDevicePairingFinishedError(f func(sender objc.IObject, error_ int)) {
	d._DevicePairingFinishedError = f
}

// SetDevicePairingPINCodeRequest sets the handler for the DevicePairingPINCodeRequest delegate method.
func (d *BluetoothDevicePairDelegate) SetDevicePairingPINCodeRequest(f func(sender objc.IObject)) {
	d._DevicePairingPINCodeRequest = f
}

// SetDevicePairingStarted sets the handler for the DevicePairingStarted delegate method.
func (d *BluetoothDevicePairDelegate) SetDevicePairingStarted(f func(sender objc.IObject)) {
	d._DevicePairingStarted = f
}

// SetDevicePairingUserConfirmationRequestNumericValue sets the handler for the DevicePairingUserConfirmationRequestNumericValue delegate method.
func (d *BluetoothDevicePairDelegate) SetDevicePairingUserConfirmationRequestNumericValue(f func(sender objc.IObject, numericValue BluetoothNumericValue /* typedef */)) {
	d._DevicePairingUserConfirmationRequestNumericValue = f
}

// SetDevicePairingUserPasskeyNotificationPasskey sets the handler for the DevicePairingUserPasskeyNotificationPasskey delegate method.
func (d *BluetoothDevicePairDelegate) SetDevicePairingUserPasskeyNotificationPasskey(f func(sender objc.IObject, passkey BluetoothPasskey /* typedef */)) {
	d._DevicePairingUserPasskeyNotificationPasskey = f
}

// SetDeviceSimplePairingCompleteStatus sets the handler for the DeviceSimplePairingCompleteStatus delegate method.
func (d *BluetoothDevicePairDelegate) SetDeviceSimplePairingCompleteStatus(f func(sender objc.IObject, status BluetoothHCIEventStatus /* typedef */)) {
	d._DeviceSimplePairingCompleteStatus = f
}

// DevicePairingConnected implements the PBluetoothDevicePairDelegate interface.
func (d *BluetoothDevicePairDelegate) DevicePairingConnected(sender objc.IObject) {
	if d._DevicePairingConnected != nil {
		d._DevicePairingConnected(sender)
	}
}

// HasDevicePairingConnected returns true if a handler for DevicePairingConnected has been set.
func (d *BluetoothDevicePairDelegate) HasDevicePairingConnected() bool {
	return d._DevicePairingConnected != nil
}

// DevicePairingConnecting implements the PBluetoothDevicePairDelegate interface.
func (d *BluetoothDevicePairDelegate) DevicePairingConnecting(sender objc.IObject) {
	if d._DevicePairingConnecting != nil {
		d._DevicePairingConnecting(sender)
	}
}

// HasDevicePairingConnecting returns true if a handler for DevicePairingConnecting has been set.
func (d *BluetoothDevicePairDelegate) HasDevicePairingConnecting() bool {
	return d._DevicePairingConnecting != nil
}

// DevicePairingFinishedError implements the PBluetoothDevicePairDelegate interface.
func (d *BluetoothDevicePairDelegate) DevicePairingFinishedError(sender objc.IObject, error_ int) {
	if d._DevicePairingFinishedError != nil {
		d._DevicePairingFinishedError(sender, error_)
	}
}

// HasDevicePairingFinishedError returns true if a handler for DevicePairingFinishedError has been set.
func (d *BluetoothDevicePairDelegate) HasDevicePairingFinishedError() bool {
	return d._DevicePairingFinishedError != nil
}

// DevicePairingPINCodeRequest implements the PBluetoothDevicePairDelegate interface.
func (d *BluetoothDevicePairDelegate) DevicePairingPINCodeRequest(sender objc.IObject) {
	if d._DevicePairingPINCodeRequest != nil {
		d._DevicePairingPINCodeRequest(sender)
	}
}

// HasDevicePairingPINCodeRequest returns true if a handler for DevicePairingPINCodeRequest has been set.
func (d *BluetoothDevicePairDelegate) HasDevicePairingPINCodeRequest() bool {
	return d._DevicePairingPINCodeRequest != nil
}

// DevicePairingStarted implements the PBluetoothDevicePairDelegate interface.
func (d *BluetoothDevicePairDelegate) DevicePairingStarted(sender objc.IObject) {
	if d._DevicePairingStarted != nil {
		d._DevicePairingStarted(sender)
	}
}

// HasDevicePairingStarted returns true if a handler for DevicePairingStarted has been set.
func (d *BluetoothDevicePairDelegate) HasDevicePairingStarted() bool {
	return d._DevicePairingStarted != nil
}

// DevicePairingUserConfirmationRequestNumericValue implements the PBluetoothDevicePairDelegate interface.
func (d *BluetoothDevicePairDelegate) DevicePairingUserConfirmationRequestNumericValue(sender objc.IObject, numericValue BluetoothNumericValue /* typedef */) {
	if d._DevicePairingUserConfirmationRequestNumericValue != nil {
		d._DevicePairingUserConfirmationRequestNumericValue(sender, numericValue)
	}
}

// HasDevicePairingUserConfirmationRequestNumericValue returns true if a handler for DevicePairingUserConfirmationRequestNumericValue has been set.
func (d *BluetoothDevicePairDelegate) HasDevicePairingUserConfirmationRequestNumericValue() bool {
	return d._DevicePairingUserConfirmationRequestNumericValue != nil
}

// DevicePairingUserPasskeyNotificationPasskey implements the PBluetoothDevicePairDelegate interface.
func (d *BluetoothDevicePairDelegate) DevicePairingUserPasskeyNotificationPasskey(sender objc.IObject, passkey BluetoothPasskey /* typedef */) {
	if d._DevicePairingUserPasskeyNotificationPasskey != nil {
		d._DevicePairingUserPasskeyNotificationPasskey(sender, passkey)
	}
}

// HasDevicePairingUserPasskeyNotificationPasskey returns true if a handler for DevicePairingUserPasskeyNotificationPasskey has been set.
func (d *BluetoothDevicePairDelegate) HasDevicePairingUserPasskeyNotificationPasskey() bool {
	return d._DevicePairingUserPasskeyNotificationPasskey != nil
}

// DeviceSimplePairingCompleteStatus implements the PBluetoothDevicePairDelegate interface.
func (d *BluetoothDevicePairDelegate) DeviceSimplePairingCompleteStatus(sender objc.IObject, status BluetoothHCIEventStatus /* typedef */) {
	if d._DeviceSimplePairingCompleteStatus != nil {
		d._DeviceSimplePairingCompleteStatus(sender, status)
	}
}

// HasDeviceSimplePairingCompleteStatus returns true if a handler for DeviceSimplePairingCompleteStatus has been set.
func (d *BluetoothDevicePairDelegate) HasDeviceSimplePairingCompleteStatus() bool {
	return d._DeviceSimplePairingCompleteStatus != nil
}
