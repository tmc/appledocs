// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PICDeviceBrowserDelegate is the ICDeviceBrowserDelegate protocol interface.
//
// Methods for managing the addition and removal of devices and responding to device changes.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - visionOS +
//
// See: doc://com.apple.imagecapturecore/documentation/ImageCaptureCore/ICDeviceBrowserDelegate
type PICDeviceBrowserDelegate interface {
	// Required methods
	DeviceBrowserDidAddDeviceMoreComing(browser ICDeviceBrowser, device ICDevice, moreComing bool)/* debug [protocol_interface/required_method]: DeviceBrowserDidAddDeviceMoreComing */
	DeviceBrowserDidRemoveDeviceMoreGoing(browser ICDeviceBrowser, device ICDevice, moreGoing bool)/* debug [protocol_interface/required_method]: DeviceBrowserDidRemoveDeviceMoreGoing */
	// Optional methods
	DeviceBrowser()
	HasDeviceBrowser() bool
	DeviceBrowserDidEnumerateLocalDevices()
	HasDeviceBrowserDidEnumerateLocalDevices() bool
	DeviceBrowserDidCancelSuspendOperations()
	HasDeviceBrowserDidCancelSuspendOperations() bool
	DeviceBrowserDidResumeOperations()
	HasDeviceBrowserDidResumeOperations() bool
	DeviceBrowserDidSuspendOperations()
	HasDeviceBrowserDidSuspendOperations() bool
	DeviceBrowserWillSuspendOperations()
	HasDeviceBrowserWillSuspendOperations() bool
	DeviceBrowserDeviceDidChangeName(browser ICDeviceBrowser, device ICDevice)
	HasDeviceBrowserDeviceDidChangeName() bool
	DeviceBrowserDeviceDidChangeSharingState(browser ICDeviceBrowser, device ICDevice)
	HasDeviceBrowserDeviceDidChangeSharingState() bool
	DeviceBrowserRequestsSelectDevice(browser ICDeviceBrowser, device ICDevice)
	HasDeviceBrowserRequestsSelectDevice() bool
}

// ICDeviceBrowserDelegate is a delegate implementation builder for the PICDeviceBrowserDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ICDeviceBrowserDelegate struct {
	_DeviceBrowser func()
	_DeviceBrowserDidEnumerateLocalDevices func()
	_DeviceBrowserDidCancelSuspendOperations func()
	_DeviceBrowserDidResumeOperations func()
	_DeviceBrowserDidSuspendOperations func()
	_DeviceBrowserWillSuspendOperations func()
	_DeviceBrowserDeviceDidChangeName func(browser ICDeviceBrowser, device ICDevice)
	_DeviceBrowserDeviceDidChangeSharingState func(browser ICDeviceBrowser, device ICDevice)
	_DeviceBrowserRequestsSelectDevice func(browser ICDeviceBrowser, device ICDevice)
	_DeviceBrowserDidAddDeviceMoreComing func(browser ICDeviceBrowser, device ICDevice, moreComing bool)
	_DeviceBrowserDidRemoveDeviceMoreGoing func(browser ICDeviceBrowser, device ICDevice, moreGoing bool)
}

// SetDeviceBrowser sets the handler for the DeviceBrowser delegate method.
//
// Tells the delegate when an event occurs on the device that may be of interest to the client application.
func (d *ICDeviceBrowserDelegate) SetDeviceBrowser(f func()) {
	d._DeviceBrowser = f
}

// SetDeviceBrowserDidEnumerateLocalDevices sets the handler for the DeviceBrowserDidEnumerateLocalDevices delegate method.
//
// Tells the delegate that the device browser has completed sending   for all local devices.
func (d *ICDeviceBrowserDelegate) SetDeviceBrowserDidEnumerateLocalDevices(f func()) {
	d._DeviceBrowserDidEnumerateLocalDevices = f
}

// SetDeviceBrowserDidCancelSuspendOperations sets the handler for the DeviceBrowserDidCancelSuspendOperations delegate method.
func (d *ICDeviceBrowserDelegate) SetDeviceBrowserDidCancelSuspendOperations(f func()) {
	d._DeviceBrowserDidCancelSuspendOperations = f
}

// SetDeviceBrowserDidResumeOperations sets the handler for the DeviceBrowserDidResumeOperations delegate method.
func (d *ICDeviceBrowserDelegate) SetDeviceBrowserDidResumeOperations(f func()) {
	d._DeviceBrowserDidResumeOperations = f
}

// SetDeviceBrowserDidSuspendOperations sets the handler for the DeviceBrowserDidSuspendOperations delegate method.
func (d *ICDeviceBrowserDelegate) SetDeviceBrowserDidSuspendOperations(f func()) {
	d._DeviceBrowserDidSuspendOperations = f
}

// SetDeviceBrowserWillSuspendOperations sets the handler for the DeviceBrowserWillSuspendOperations delegate method.
func (d *ICDeviceBrowserDelegate) SetDeviceBrowserWillSuspendOperations(f func()) {
	d._DeviceBrowserWillSuspendOperations = f
}

// SetDeviceBrowserDeviceDidChangeName sets the handler for the DeviceBrowserDeviceDidChangeName delegate method.
//
// Tells the delegate when the name of a device changes.
func (d *ICDeviceBrowserDelegate) SetDeviceBrowserDeviceDidChangeName(f func(browser ICDeviceBrowser, device ICDevice)) {
	d._DeviceBrowserDeviceDidChangeName = f
}

// SetDeviceBrowserDeviceDidChangeSharingState sets the handler for the DeviceBrowserDeviceDidChangeSharingState delegate method.
//
// Tells the delegate when the sharing state of a device changes.
func (d *ICDeviceBrowserDelegate) SetDeviceBrowserDeviceDidChangeSharingState(f func(browser ICDeviceBrowser, device ICDevice)) {
	d._DeviceBrowserDeviceDidChangeSharingState = f
}

// SetDeviceBrowserRequestsSelectDevice sets the handler for the DeviceBrowserRequestsSelectDevice delegate method.
//
// Tells the delegate when an event occurs on the device that may be of interest to the client application.
func (d *ICDeviceBrowserDelegate) SetDeviceBrowserRequestsSelectDevice(f func(browser ICDeviceBrowser, device ICDevice)) {
	d._DeviceBrowserRequestsSelectDevice = f
}

// SetDeviceBrowserDidAddDeviceMoreComing sets the handler for the DeviceBrowserDidAddDeviceMoreComing delegate method.
//
// Tells the delegate that a device has been added.
func (d *ICDeviceBrowserDelegate) SetDeviceBrowserDidAddDeviceMoreComing(f func(browser ICDeviceBrowser, device ICDevice, moreComing bool)) {
	d._DeviceBrowserDidAddDeviceMoreComing = f
}

// SetDeviceBrowserDidRemoveDeviceMoreGoing sets the handler for the DeviceBrowserDidRemoveDeviceMoreGoing delegate method.
//
// Tells the delegate that a device has been removed.
func (d *ICDeviceBrowserDelegate) SetDeviceBrowserDidRemoveDeviceMoreGoing(f func(browser ICDeviceBrowser, device ICDevice, moreGoing bool)) {
	d._DeviceBrowserDidRemoveDeviceMoreGoing = f
}

// DeviceBrowser implements the PICDeviceBrowserDelegate interface.
func (d *ICDeviceBrowserDelegate) DeviceBrowser() {
	if d._DeviceBrowser != nil {
		d._DeviceBrowser()
	}
}

// HasDeviceBrowser returns true if a handler for DeviceBrowser has been set.
func (d *ICDeviceBrowserDelegate) HasDeviceBrowser() bool {
	return d._DeviceBrowser != nil
}

// DeviceBrowserDidEnumerateLocalDevices implements the PICDeviceBrowserDelegate interface.
func (d *ICDeviceBrowserDelegate) DeviceBrowserDidEnumerateLocalDevices() {
	if d._DeviceBrowserDidEnumerateLocalDevices != nil {
		d._DeviceBrowserDidEnumerateLocalDevices()
	}
}

// HasDeviceBrowserDidEnumerateLocalDevices returns true if a handler for DeviceBrowserDidEnumerateLocalDevices has been set.
func (d *ICDeviceBrowserDelegate) HasDeviceBrowserDidEnumerateLocalDevices() bool {
	return d._DeviceBrowserDidEnumerateLocalDevices != nil
}

// DeviceBrowserDidCancelSuspendOperations implements the PICDeviceBrowserDelegate interface.
func (d *ICDeviceBrowserDelegate) DeviceBrowserDidCancelSuspendOperations() {
	if d._DeviceBrowserDidCancelSuspendOperations != nil {
		d._DeviceBrowserDidCancelSuspendOperations()
	}
}

// HasDeviceBrowserDidCancelSuspendOperations returns true if a handler for DeviceBrowserDidCancelSuspendOperations has been set.
func (d *ICDeviceBrowserDelegate) HasDeviceBrowserDidCancelSuspendOperations() bool {
	return d._DeviceBrowserDidCancelSuspendOperations != nil
}

// DeviceBrowserDidResumeOperations implements the PICDeviceBrowserDelegate interface.
func (d *ICDeviceBrowserDelegate) DeviceBrowserDidResumeOperations() {
	if d._DeviceBrowserDidResumeOperations != nil {
		d._DeviceBrowserDidResumeOperations()
	}
}

// HasDeviceBrowserDidResumeOperations returns true if a handler for DeviceBrowserDidResumeOperations has been set.
func (d *ICDeviceBrowserDelegate) HasDeviceBrowserDidResumeOperations() bool {
	return d._DeviceBrowserDidResumeOperations != nil
}

// DeviceBrowserDidSuspendOperations implements the PICDeviceBrowserDelegate interface.
func (d *ICDeviceBrowserDelegate) DeviceBrowserDidSuspendOperations() {
	if d._DeviceBrowserDidSuspendOperations != nil {
		d._DeviceBrowserDidSuspendOperations()
	}
}

// HasDeviceBrowserDidSuspendOperations returns true if a handler for DeviceBrowserDidSuspendOperations has been set.
func (d *ICDeviceBrowserDelegate) HasDeviceBrowserDidSuspendOperations() bool {
	return d._DeviceBrowserDidSuspendOperations != nil
}

// DeviceBrowserWillSuspendOperations implements the PICDeviceBrowserDelegate interface.
func (d *ICDeviceBrowserDelegate) DeviceBrowserWillSuspendOperations() {
	if d._DeviceBrowserWillSuspendOperations != nil {
		d._DeviceBrowserWillSuspendOperations()
	}
}

// HasDeviceBrowserWillSuspendOperations returns true if a handler for DeviceBrowserWillSuspendOperations has been set.
func (d *ICDeviceBrowserDelegate) HasDeviceBrowserWillSuspendOperations() bool {
	return d._DeviceBrowserWillSuspendOperations != nil
}

// DeviceBrowserDeviceDidChangeName implements the PICDeviceBrowserDelegate interface.
func (d *ICDeviceBrowserDelegate) DeviceBrowserDeviceDidChangeName(browser ICDeviceBrowser, device ICDevice) {
	if d._DeviceBrowserDeviceDidChangeName != nil {
		d._DeviceBrowserDeviceDidChangeName(browser, device)
	}
}

// HasDeviceBrowserDeviceDidChangeName returns true if a handler for DeviceBrowserDeviceDidChangeName has been set.
func (d *ICDeviceBrowserDelegate) HasDeviceBrowserDeviceDidChangeName() bool {
	return d._DeviceBrowserDeviceDidChangeName != nil
}

// DeviceBrowserDeviceDidChangeSharingState implements the PICDeviceBrowserDelegate interface.
func (d *ICDeviceBrowserDelegate) DeviceBrowserDeviceDidChangeSharingState(browser ICDeviceBrowser, device ICDevice) {
	if d._DeviceBrowserDeviceDidChangeSharingState != nil {
		d._DeviceBrowserDeviceDidChangeSharingState(browser, device)
	}
}

// HasDeviceBrowserDeviceDidChangeSharingState returns true if a handler for DeviceBrowserDeviceDidChangeSharingState has been set.
func (d *ICDeviceBrowserDelegate) HasDeviceBrowserDeviceDidChangeSharingState() bool {
	return d._DeviceBrowserDeviceDidChangeSharingState != nil
}

// DeviceBrowserRequestsSelectDevice implements the PICDeviceBrowserDelegate interface.
func (d *ICDeviceBrowserDelegate) DeviceBrowserRequestsSelectDevice(browser ICDeviceBrowser, device ICDevice) {
	if d._DeviceBrowserRequestsSelectDevice != nil {
		d._DeviceBrowserRequestsSelectDevice(browser, device)
	}
}

// HasDeviceBrowserRequestsSelectDevice returns true if a handler for DeviceBrowserRequestsSelectDevice has been set.
func (d *ICDeviceBrowserDelegate) HasDeviceBrowserRequestsSelectDevice() bool {
	return d._DeviceBrowserRequestsSelectDevice != nil
}

// DeviceBrowserDidAddDeviceMoreComing implements the PICDeviceBrowserDelegate interface.
func (d *ICDeviceBrowserDelegate) DeviceBrowserDidAddDeviceMoreComing(browser ICDeviceBrowser, device ICDevice, moreComing bool) {
	if d._DeviceBrowserDidAddDeviceMoreComing != nil {
		d._DeviceBrowserDidAddDeviceMoreComing(browser, device, moreComing)
	}
}

// HasDeviceBrowserDidAddDeviceMoreComing returns true if a handler for DeviceBrowserDidAddDeviceMoreComing has been set.
func (d *ICDeviceBrowserDelegate) HasDeviceBrowserDidAddDeviceMoreComing() bool {
	return d._DeviceBrowserDidAddDeviceMoreComing != nil
}

// DeviceBrowserDidRemoveDeviceMoreGoing implements the PICDeviceBrowserDelegate interface.
func (d *ICDeviceBrowserDelegate) DeviceBrowserDidRemoveDeviceMoreGoing(browser ICDeviceBrowser, device ICDevice, moreGoing bool) {
	if d._DeviceBrowserDidRemoveDeviceMoreGoing != nil {
		d._DeviceBrowserDidRemoveDeviceMoreGoing(browser, device, moreGoing)
	}
}

// HasDeviceBrowserDidRemoveDeviceMoreGoing returns true if a handler for DeviceBrowserDidRemoveDeviceMoreGoing has been set.
func (d *ICDeviceBrowserDelegate) HasDeviceBrowserDidRemoveDeviceMoreGoing() bool {
	return d._DeviceBrowserDidRemoveDeviceMoreGoing != nil
}
