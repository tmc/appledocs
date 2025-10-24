// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTRDeviceDelegate is the MTRDeviceDelegate protocol interface.
//
// Availability:
//   - Mac Catalyst 16.1+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+
//
// See: doc://com.apple.matter/documentation/Matter/MTRDeviceDelegate
type PMTRDeviceDelegate interface {
	// Required methods
	DeviceReceivedAttributeReport(device IMTRDevice, attributeReport foundation.IDictionary)/* debug [protocol_interface/required_method]: DeviceReceivedAttributeReport */
	DeviceReceivedEventReport(device IMTRDevice, eventReport foundation.IDictionary)/* debug [protocol_interface/required_method]: DeviceReceivedEventReport */
	DeviceStateChanged(device IMTRDevice, state unsafe.Pointer)/* debug [protocol_interface/required_method]: DeviceStateChanged */
	// Optional methods
	DeviceBecameActive(device IMTRDevice)
	HasDeviceBecameActive() bool
	DeviceCachePrimed(device IMTRDevice)
	HasDeviceCachePrimed() bool
	DeviceConfigurationChanged(device IMTRDevice)
	HasDeviceConfigurationChanged() bool
}

// MTRDeviceDelegate is a delegate implementation builder for the PMTRDeviceDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MTRDeviceDelegate struct {
	_DeviceBecameActive func(device IMTRDevice)
	_DeviceCachePrimed func(device IMTRDevice)
	_DeviceConfigurationChanged func(device IMTRDevice)
	_DeviceReceivedAttributeReport func(device IMTRDevice, attributeReport foundation.IDictionary)
	_DeviceReceivedEventReport func(device IMTRDevice, eventReport foundation.IDictionary)
	_DeviceStateChanged func(device IMTRDevice, state unsafe.Pointer)
}

// SetDeviceBecameActive sets the handler for the DeviceBecameActive delegate method.
func (d *MTRDeviceDelegate) SetDeviceBecameActive(f func(device IMTRDevice)) {
	d._DeviceBecameActive = f
}

// SetDeviceCachePrimed sets the handler for the DeviceCachePrimed delegate method.
func (d *MTRDeviceDelegate) SetDeviceCachePrimed(f func(device IMTRDevice)) {
	d._DeviceCachePrimed = f
}

// SetDeviceConfigurationChanged sets the handler for the DeviceConfigurationChanged delegate method.
func (d *MTRDeviceDelegate) SetDeviceConfigurationChanged(f func(device IMTRDevice)) {
	d._DeviceConfigurationChanged = f
}

// SetDeviceReceivedAttributeReport sets the handler for the DeviceReceivedAttributeReport delegate method.
func (d *MTRDeviceDelegate) SetDeviceReceivedAttributeReport(f func(device IMTRDevice, attributeReport foundation.IDictionary)) {
	d._DeviceReceivedAttributeReport = f
}

// SetDeviceReceivedEventReport sets the handler for the DeviceReceivedEventReport delegate method.
func (d *MTRDeviceDelegate) SetDeviceReceivedEventReport(f func(device IMTRDevice, eventReport foundation.IDictionary)) {
	d._DeviceReceivedEventReport = f
}

// SetDeviceStateChanged sets the handler for the DeviceStateChanged delegate method.
func (d *MTRDeviceDelegate) SetDeviceStateChanged(f func(device IMTRDevice, state unsafe.Pointer)) {
	d._DeviceStateChanged = f
}

// DeviceBecameActive implements the PMTRDeviceDelegate interface.
func (d *MTRDeviceDelegate) DeviceBecameActive(device IMTRDevice) {
	if d._DeviceBecameActive != nil {
		d._DeviceBecameActive(device)
	}
}

// HasDeviceBecameActive returns true if a handler for DeviceBecameActive has been set.
func (d *MTRDeviceDelegate) HasDeviceBecameActive() bool {
	return d._DeviceBecameActive != nil
}

// DeviceCachePrimed implements the PMTRDeviceDelegate interface.
func (d *MTRDeviceDelegate) DeviceCachePrimed(device IMTRDevice) {
	if d._DeviceCachePrimed != nil {
		d._DeviceCachePrimed(device)
	}
}

// HasDeviceCachePrimed returns true if a handler for DeviceCachePrimed has been set.
func (d *MTRDeviceDelegate) HasDeviceCachePrimed() bool {
	return d._DeviceCachePrimed != nil
}

// DeviceConfigurationChanged implements the PMTRDeviceDelegate interface.
func (d *MTRDeviceDelegate) DeviceConfigurationChanged(device IMTRDevice) {
	if d._DeviceConfigurationChanged != nil {
		d._DeviceConfigurationChanged(device)
	}
}

// HasDeviceConfigurationChanged returns true if a handler for DeviceConfigurationChanged has been set.
func (d *MTRDeviceDelegate) HasDeviceConfigurationChanged() bool {
	return d._DeviceConfigurationChanged != nil
}

// DeviceReceivedAttributeReport implements the PMTRDeviceDelegate interface.
func (d *MTRDeviceDelegate) DeviceReceivedAttributeReport(device IMTRDevice, attributeReport foundation.IDictionary) {
	if d._DeviceReceivedAttributeReport != nil {
		d._DeviceReceivedAttributeReport(device, attributeReport)
	}
}

// HasDeviceReceivedAttributeReport returns true if a handler for DeviceReceivedAttributeReport has been set.
func (d *MTRDeviceDelegate) HasDeviceReceivedAttributeReport() bool {
	return d._DeviceReceivedAttributeReport != nil
}

// DeviceReceivedEventReport implements the PMTRDeviceDelegate interface.
func (d *MTRDeviceDelegate) DeviceReceivedEventReport(device IMTRDevice, eventReport foundation.IDictionary) {
	if d._DeviceReceivedEventReport != nil {
		d._DeviceReceivedEventReport(device, eventReport)
	}
}

// HasDeviceReceivedEventReport returns true if a handler for DeviceReceivedEventReport has been set.
func (d *MTRDeviceDelegate) HasDeviceReceivedEventReport() bool {
	return d._DeviceReceivedEventReport != nil
}

// DeviceStateChanged implements the PMTRDeviceDelegate interface.
func (d *MTRDeviceDelegate) DeviceStateChanged(device IMTRDevice, state unsafe.Pointer) {
	if d._DeviceStateChanged != nil {
		d._DeviceStateChanged(device, state)
	}
}

// HasDeviceStateChanged returns true if a handler for DeviceStateChanged has been set.
func (d *MTRDeviceDelegate) HasDeviceStateChanged() bool {
	return d._DeviceStateChanged != nil
}
