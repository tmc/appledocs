// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PICDeviceDelegate is the ICDeviceDelegate protocol interface.
//
// Methods for responding to device events and changes.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - visionOS +
//
// See: doc://com.apple.imagecapturecore/documentation/ImageCaptureCore/ICDeviceDelegate
type PICDeviceDelegate interface {
	// Required methods
	DidRemove()/* debug [protocol_interface/required_method]: DidRemove */
	DidRemoveDevice(device ICDevice)/* debug [protocol_interface/required_method]: DidRemoveDevice */
	DeviceDidCloseSessionWithError(device ICDevice, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: DeviceDidCloseSessionWithError */
	DeviceDidOpenSessionWithError(device ICDevice, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: DeviceDidOpenSessionWithError */
	// Optional methods
	Device()
	HasDevice() bool
	DeviceDidBecomeReady()
	HasDeviceDidBecomeReady() bool
	DeviceDidChangeSharingState()
	HasDeviceDidChangeSharingState() bool
	DeviceDidChangeName()
	HasDeviceDidChangeName() bool
	DeviceDidEjectWithError(device ICDevice, error_ objc.IObject /* cross-framework: Error */)
	HasDeviceDidEjectWithError() bool
	DeviceDidEncounterError(device ICDevice, error_ objc.IObject /* cross-framework: Error */)
	HasDeviceDidEncounterError() bool
	DeviceDidReceiveStatusInformation(device ICDevice, status foundation.IDictionary)
	HasDeviceDidReceiveStatusInformation() bool
}

// ICDeviceDelegate is a delegate implementation builder for the PICDeviceDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ICDeviceDelegate struct {
	_Device func()
	_DeviceDidBecomeReady func()
	_DeviceDidChangeSharingState func()
	_DeviceDidChangeName func()
	_DeviceDidEjectWithError func(device ICDevice, error_ objc.IObject /* cross-framework: Error */)
	_DeviceDidEncounterError func(device ICDevice, error_ objc.IObject /* cross-framework: Error */)
	_DeviceDidReceiveStatusInformation func(device ICDevice, status foundation.IDictionary)
	_DidRemove func()
	_DidRemoveDevice func(device ICDevice)
	_DeviceDidCloseSessionWithError func(device ICDevice, error_ objc.IObject /* cross-framework: Error */)
	_DeviceDidOpenSessionWithError func(device ICDevice, error_ objc.IObject /* cross-framework: Error */)
}

// SetDevice sets the handler for the Device delegate method.
//
// Tells the delegate when a device encounters an error.
func (d *ICDeviceDelegate) SetDevice(f func()) {
	d._Device = f
}

// SetDeviceDidBecomeReady sets the handler for the DeviceDidBecomeReady delegate method.
//
// Tells the delegate when the device is ready to receive requests.
func (d *ICDeviceDelegate) SetDeviceDidBecomeReady(f func()) {
	d._DeviceDidBecomeReady = f
}

// SetDeviceDidChangeSharingState sets the handler for the DeviceDidChangeSharingState delegate method.
//
// Tells the delegate when the sharing state of a device changes.
func (d *ICDeviceDelegate) SetDeviceDidChangeSharingState(f func()) {
	d._DeviceDidChangeSharingState = f
}

// SetDeviceDidChangeName sets the handler for the DeviceDidChangeName delegate method.
//
// Tells the delegate when the name of a device changes.
func (d *ICDeviceDelegate) SetDeviceDidChangeName(f func()) {
	d._DeviceDidChangeName = f
}

// SetDeviceDidEjectWithError sets the handler for the DeviceDidEjectWithError delegate method.
//
// Tells the delegate when the ejection is complete.
func (d *ICDeviceDelegate) SetDeviceDidEjectWithError(f func(device ICDevice, error_ objc.IObject /* cross-framework: Error */)) {
	d._DeviceDidEjectWithError = f
}

// SetDeviceDidEncounterError sets the handler for the DeviceDidEncounterError delegate method.
//
// Tells the delegate when a device encounters an error.
func (d *ICDeviceDelegate) SetDeviceDidEncounterError(f func(device ICDevice, error_ objc.IObject /* cross-framework: Error */)) {
	d._DeviceDidEncounterError = f
}

// SetDeviceDidReceiveStatusInformation sets the handler for the DeviceDidReceiveStatusInformation delegate method.
//
// Tells the delegate when status information is received from a device.
func (d *ICDeviceDelegate) SetDeviceDidReceiveStatusInformation(f func(device ICDevice, status foundation.IDictionary)) {
	d._DeviceDidReceiveStatusInformation = f
}

// SetDidRemove sets the handler for the DidRemove delegate method.
//
// Tells the delegate that a device has been removed.
func (d *ICDeviceDelegate) SetDidRemove(f func()) {
	d._DidRemove = f
}

// SetDidRemoveDevice sets the handler for the DidRemoveDevice delegate method.
//
// Tells the delegate that a device has been removed.
func (d *ICDeviceDelegate) SetDidRemoveDevice(f func(device ICDevice)) {
	d._DidRemoveDevice = f
}

// SetDeviceDidCloseSessionWithError sets the handler for the DeviceDidCloseSessionWithError delegate method.
//
// Tells the delegate when a session is closed on a device.
func (d *ICDeviceDelegate) SetDeviceDidCloseSessionWithError(f func(device ICDevice, error_ objc.IObject /* cross-framework: Error */)) {
	d._DeviceDidCloseSessionWithError = f
}

// SetDeviceDidOpenSessionWithError sets the handler for the DeviceDidOpenSessionWithError delegate method.
//
// Tells the delegate when a session is opened on a device.
func (d *ICDeviceDelegate) SetDeviceDidOpenSessionWithError(f func(device ICDevice, error_ objc.IObject /* cross-framework: Error */)) {
	d._DeviceDidOpenSessionWithError = f
}

// Device implements the PICDeviceDelegate interface.
func (d *ICDeviceDelegate) Device() {
	if d._Device != nil {
		d._Device()
	}
}

// HasDevice returns true if a handler for Device has been set.
func (d *ICDeviceDelegate) HasDevice() bool {
	return d._Device != nil
}

// DeviceDidBecomeReady implements the PICDeviceDelegate interface.
func (d *ICDeviceDelegate) DeviceDidBecomeReady() {
	if d._DeviceDidBecomeReady != nil {
		d._DeviceDidBecomeReady()
	}
}

// HasDeviceDidBecomeReady returns true if a handler for DeviceDidBecomeReady has been set.
func (d *ICDeviceDelegate) HasDeviceDidBecomeReady() bool {
	return d._DeviceDidBecomeReady != nil
}

// DeviceDidChangeSharingState implements the PICDeviceDelegate interface.
func (d *ICDeviceDelegate) DeviceDidChangeSharingState() {
	if d._DeviceDidChangeSharingState != nil {
		d._DeviceDidChangeSharingState()
	}
}

// HasDeviceDidChangeSharingState returns true if a handler for DeviceDidChangeSharingState has been set.
func (d *ICDeviceDelegate) HasDeviceDidChangeSharingState() bool {
	return d._DeviceDidChangeSharingState != nil
}

// DeviceDidChangeName implements the PICDeviceDelegate interface.
func (d *ICDeviceDelegate) DeviceDidChangeName() {
	if d._DeviceDidChangeName != nil {
		d._DeviceDidChangeName()
	}
}

// HasDeviceDidChangeName returns true if a handler for DeviceDidChangeName has been set.
func (d *ICDeviceDelegate) HasDeviceDidChangeName() bool {
	return d._DeviceDidChangeName != nil
}

// DeviceDidEjectWithError implements the PICDeviceDelegate interface.
func (d *ICDeviceDelegate) DeviceDidEjectWithError(device ICDevice, error_ objc.IObject /* cross-framework: Error */) {
	if d._DeviceDidEjectWithError != nil {
		d._DeviceDidEjectWithError(device, error_)
	}
}

// HasDeviceDidEjectWithError returns true if a handler for DeviceDidEjectWithError has been set.
func (d *ICDeviceDelegate) HasDeviceDidEjectWithError() bool {
	return d._DeviceDidEjectWithError != nil
}

// DeviceDidEncounterError implements the PICDeviceDelegate interface.
func (d *ICDeviceDelegate) DeviceDidEncounterError(device ICDevice, error_ objc.IObject /* cross-framework: Error */) {
	if d._DeviceDidEncounterError != nil {
		d._DeviceDidEncounterError(device, error_)
	}
}

// HasDeviceDidEncounterError returns true if a handler for DeviceDidEncounterError has been set.
func (d *ICDeviceDelegate) HasDeviceDidEncounterError() bool {
	return d._DeviceDidEncounterError != nil
}

// DeviceDidReceiveStatusInformation implements the PICDeviceDelegate interface.
func (d *ICDeviceDelegate) DeviceDidReceiveStatusInformation(device ICDevice, status foundation.IDictionary) {
	if d._DeviceDidReceiveStatusInformation != nil {
		d._DeviceDidReceiveStatusInformation(device, status)
	}
}

// HasDeviceDidReceiveStatusInformation returns true if a handler for DeviceDidReceiveStatusInformation has been set.
func (d *ICDeviceDelegate) HasDeviceDidReceiveStatusInformation() bool {
	return d._DeviceDidReceiveStatusInformation != nil
}

// DidRemove implements the PICDeviceDelegate interface.
func (d *ICDeviceDelegate) DidRemove() {
	if d._DidRemove != nil {
		d._DidRemove()
	}
}

// HasDidRemove returns true if a handler for DidRemove has been set.
func (d *ICDeviceDelegate) HasDidRemove() bool {
	return d._DidRemove != nil
}

// DidRemoveDevice implements the PICDeviceDelegate interface.
func (d *ICDeviceDelegate) DidRemoveDevice(device ICDevice) {
	if d._DidRemoveDevice != nil {
		d._DidRemoveDevice(device)
	}
}

// HasDidRemoveDevice returns true if a handler for DidRemoveDevice has been set.
func (d *ICDeviceDelegate) HasDidRemoveDevice() bool {
	return d._DidRemoveDevice != nil
}

// DeviceDidCloseSessionWithError implements the PICDeviceDelegate interface.
func (d *ICDeviceDelegate) DeviceDidCloseSessionWithError(device ICDevice, error_ objc.IObject /* cross-framework: Error */) {
	if d._DeviceDidCloseSessionWithError != nil {
		d._DeviceDidCloseSessionWithError(device, error_)
	}
}

// HasDeviceDidCloseSessionWithError returns true if a handler for DeviceDidCloseSessionWithError has been set.
func (d *ICDeviceDelegate) HasDeviceDidCloseSessionWithError() bool {
	return d._DeviceDidCloseSessionWithError != nil
}

// DeviceDidOpenSessionWithError implements the PICDeviceDelegate interface.
func (d *ICDeviceDelegate) DeviceDidOpenSessionWithError(device ICDevice, error_ objc.IObject /* cross-framework: Error */) {
	if d._DeviceDidOpenSessionWithError != nil {
		d._DeviceDidOpenSessionWithError(device, error_)
	}
}

// HasDeviceDidOpenSessionWithError returns true if a handler for DeviceDidOpenSessionWithError has been set.
func (d *ICDeviceDelegate) HasDeviceDidOpenSessionWithError() bool {
	return d._DeviceDidOpenSessionWithError != nil
}
