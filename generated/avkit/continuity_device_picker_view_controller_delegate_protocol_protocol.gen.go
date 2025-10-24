// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/avfoundation"
)

// PContinuityDevicePickerViewControllerDelegate is the AVContinuityDevicePickerViewControllerDelegate protocol interface.
//
// An interface that responds to events from a continuity device picker view controller.
//
// Availability:
//   - tvOS 9.0+
//
// See: doc://com.apple.avkit/documentation/AVKit/AVContinuityDevicePickerViewControllerDelegate
type PContinuityDevicePickerViewControllerDelegate interface {
	// Optional methods
	ContinuityDevicePickerDidConnectDevice(pickerViewController IAVContinuityDevicePickerViewController, device avfoundation.ContinuityDevice)
	HasContinuityDevicePickerDidConnectDevice() bool
	ContinuityDevicePickerDidCancel(pickerViewController IAVContinuityDevicePickerViewController)
	HasContinuityDevicePickerDidCancel() bool
	ContinuityDevicePickerDidEndPresenting(pickerViewController IAVContinuityDevicePickerViewController)
	HasContinuityDevicePickerDidEndPresenting() bool
	ContinuityDevicePickerWillBeginPresenting(pickerViewController IAVContinuityDevicePickerViewController)
	HasContinuityDevicePickerWillBeginPresenting() bool
}

// ContinuityDevicePickerViewControllerDelegate is a delegate implementation builder for the PContinuityDevicePickerViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ContinuityDevicePickerViewControllerDelegate struct {
	_ContinuityDevicePickerDidConnectDevice func(pickerViewController IAVContinuityDevicePickerViewController, device avfoundation.ContinuityDevice)
	_ContinuityDevicePickerDidCancel func(pickerViewController IAVContinuityDevicePickerViewController)
	_ContinuityDevicePickerDidEndPresenting func(pickerViewController IAVContinuityDevicePickerViewController)
	_ContinuityDevicePickerWillBeginPresenting func(pickerViewController IAVContinuityDevicePickerViewController)
}

// SetContinuityDevicePickerDidConnectDevice sets the handler for the ContinuityDevicePickerDidConnectDevice delegate method.
//
// Informs the delegate when a person selects and connects a continuity device to the system with a continuity device picker.
func (d *ContinuityDevicePickerViewControllerDelegate) SetContinuityDevicePickerDidConnectDevice(f func(pickerViewController IAVContinuityDevicePickerViewController, device avfoundation.ContinuityDevice)) {
	d._ContinuityDevicePickerDidConnectDevice = f
}

// SetContinuityDevicePickerDidCancel sets the handler for the ContinuityDevicePickerDidCancel delegate method.
//
// Informs the delegate when a person declines to select a continuity device by dismissing an app’s continuity device picker.
func (d *ContinuityDevicePickerViewControllerDelegate) SetContinuityDevicePickerDidCancel(f func(pickerViewController IAVContinuityDevicePickerViewController)) {
	d._ContinuityDevicePickerDidCancel = f
}

// SetContinuityDevicePickerDidEndPresenting sets the handler for the ContinuityDevicePickerDidEndPresenting delegate method.
//
// Informs the delegate that a continuity device picker is no longer presenting its UI to a person.
func (d *ContinuityDevicePickerViewControllerDelegate) SetContinuityDevicePickerDidEndPresenting(f func(pickerViewController IAVContinuityDevicePickerViewController)) {
	d._ContinuityDevicePickerDidEndPresenting = f
}

// SetContinuityDevicePickerWillBeginPresenting sets the handler for the ContinuityDevicePickerWillBeginPresenting delegate method.
//
// Informs the delegate that a continuity device picker is about to present its UI so that a person can select and connect a continuity device.
func (d *ContinuityDevicePickerViewControllerDelegate) SetContinuityDevicePickerWillBeginPresenting(f func(pickerViewController IAVContinuityDevicePickerViewController)) {
	d._ContinuityDevicePickerWillBeginPresenting = f
}

// ContinuityDevicePickerDidConnectDevice implements the PContinuityDevicePickerViewControllerDelegate interface.
func (d *ContinuityDevicePickerViewControllerDelegate) ContinuityDevicePickerDidConnectDevice(pickerViewController IAVContinuityDevicePickerViewController, device avfoundation.ContinuityDevice) {
	if d._ContinuityDevicePickerDidConnectDevice != nil {
		d._ContinuityDevicePickerDidConnectDevice(pickerViewController, device)
	}
}

// HasContinuityDevicePickerDidConnectDevice returns true if a handler for ContinuityDevicePickerDidConnectDevice has been set.
func (d *ContinuityDevicePickerViewControllerDelegate) HasContinuityDevicePickerDidConnectDevice() bool {
	return d._ContinuityDevicePickerDidConnectDevice != nil
}

// ContinuityDevicePickerDidCancel implements the PContinuityDevicePickerViewControllerDelegate interface.
func (d *ContinuityDevicePickerViewControllerDelegate) ContinuityDevicePickerDidCancel(pickerViewController IAVContinuityDevicePickerViewController) {
	if d._ContinuityDevicePickerDidCancel != nil {
		d._ContinuityDevicePickerDidCancel(pickerViewController)
	}
}

// HasContinuityDevicePickerDidCancel returns true if a handler for ContinuityDevicePickerDidCancel has been set.
func (d *ContinuityDevicePickerViewControllerDelegate) HasContinuityDevicePickerDidCancel() bool {
	return d._ContinuityDevicePickerDidCancel != nil
}

// ContinuityDevicePickerDidEndPresenting implements the PContinuityDevicePickerViewControllerDelegate interface.
func (d *ContinuityDevicePickerViewControllerDelegate) ContinuityDevicePickerDidEndPresenting(pickerViewController IAVContinuityDevicePickerViewController) {
	if d._ContinuityDevicePickerDidEndPresenting != nil {
		d._ContinuityDevicePickerDidEndPresenting(pickerViewController)
	}
}

// HasContinuityDevicePickerDidEndPresenting returns true if a handler for ContinuityDevicePickerDidEndPresenting has been set.
func (d *ContinuityDevicePickerViewControllerDelegate) HasContinuityDevicePickerDidEndPresenting() bool {
	return d._ContinuityDevicePickerDidEndPresenting != nil
}

// ContinuityDevicePickerWillBeginPresenting implements the PContinuityDevicePickerViewControllerDelegate interface.
func (d *ContinuityDevicePickerViewControllerDelegate) ContinuityDevicePickerWillBeginPresenting(pickerViewController IAVContinuityDevicePickerViewController) {
	if d._ContinuityDevicePickerWillBeginPresenting != nil {
		d._ContinuityDevicePickerWillBeginPresenting(pickerViewController)
	}
}

// HasContinuityDevicePickerWillBeginPresenting returns true if a handler for ContinuityDevicePickerWillBeginPresenting has been set.
func (d *ContinuityDevicePickerViewControllerDelegate) HasContinuityDevicePickerWillBeginPresenting() bool {
	return d._ContinuityDevicePickerWillBeginPresenting != nil
}
