// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PMTRDeviceAttestationDelegate is the MTRDeviceAttestationDelegate protocol interface.
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
// See: doc://com.apple.matter/documentation/Matter/MTRDeviceAttestationDelegate
type PMTRDeviceAttestationDelegate interface {
	// Optional methods
	DeviceAttestationCompletedForDeviceAttestationDeviceInfoError(controller IMTRDeviceController, device unsafe.Pointer, attestationDeviceInfo IMTRDeviceAttestationDeviceInfo, error_ objc.IObject /* cross-framework: Error */)
	HasDeviceAttestationCompletedForDeviceAttestationDeviceInfoError() bool
	DeviceAttestationFailedForDeviceError(controller IMTRDeviceController, device unsafe.Pointer, error_ objc.IObject /* cross-framework: Error */)
	HasDeviceAttestationFailedForDeviceError() bool
	DeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError(controller IMTRDeviceController, opaqueDeviceHandle unsafe.Pointer, attestationDeviceInfo IMTRDeviceAttestationDeviceInfo, error_ objc.IObject /* cross-framework: Error */)
	HasDeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError() bool
	DeviceAttestationFailedForControllerOpaqueDeviceHandleError(controller IMTRDeviceController, opaqueDeviceHandle unsafe.Pointer, error_ objc.IObject /* cross-framework: Error */)
	HasDeviceAttestationFailedForControllerOpaqueDeviceHandleError() bool
}

// MTRDeviceAttestationDelegate is a delegate implementation builder for the PMTRDeviceAttestationDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MTRDeviceAttestationDelegate struct {
	_DeviceAttestationCompletedForDeviceAttestationDeviceInfoError func(controller IMTRDeviceController, device unsafe.Pointer, attestationDeviceInfo IMTRDeviceAttestationDeviceInfo, error_ objc.IObject /* cross-framework: Error */)
	_DeviceAttestationFailedForDeviceError func(controller IMTRDeviceController, device unsafe.Pointer, error_ objc.IObject /* cross-framework: Error */)
	_DeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError func(controller IMTRDeviceController, opaqueDeviceHandle unsafe.Pointer, attestationDeviceInfo IMTRDeviceAttestationDeviceInfo, error_ objc.IObject /* cross-framework: Error */)
	_DeviceAttestationFailedForControllerOpaqueDeviceHandleError func(controller IMTRDeviceController, opaqueDeviceHandle unsafe.Pointer, error_ objc.IObject /* cross-framework: Error */)
}

// SetDeviceAttestationCompletedForDeviceAttestationDeviceInfoError sets the handler for the DeviceAttestationCompletedForDeviceAttestationDeviceInfoError delegate method.
func (d *MTRDeviceAttestationDelegate) SetDeviceAttestationCompletedForDeviceAttestationDeviceInfoError(f func(controller IMTRDeviceController, device unsafe.Pointer, attestationDeviceInfo IMTRDeviceAttestationDeviceInfo, error_ objc.IObject /* cross-framework: Error */)) {
	d._DeviceAttestationCompletedForDeviceAttestationDeviceInfoError = f
}

// SetDeviceAttestationFailedForDeviceError sets the handler for the DeviceAttestationFailedForDeviceError delegate method.
func (d *MTRDeviceAttestationDelegate) SetDeviceAttestationFailedForDeviceError(f func(controller IMTRDeviceController, device unsafe.Pointer, error_ objc.IObject /* cross-framework: Error */)) {
	d._DeviceAttestationFailedForDeviceError = f
}

// SetDeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError sets the handler for the DeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError delegate method.
func (d *MTRDeviceAttestationDelegate) SetDeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError(f func(controller IMTRDeviceController, opaqueDeviceHandle unsafe.Pointer, attestationDeviceInfo IMTRDeviceAttestationDeviceInfo, error_ objc.IObject /* cross-framework: Error */)) {
	d._DeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError = f
}

// SetDeviceAttestationFailedForControllerOpaqueDeviceHandleError sets the handler for the DeviceAttestationFailedForControllerOpaqueDeviceHandleError delegate method.
func (d *MTRDeviceAttestationDelegate) SetDeviceAttestationFailedForControllerOpaqueDeviceHandleError(f func(controller IMTRDeviceController, opaqueDeviceHandle unsafe.Pointer, error_ objc.IObject /* cross-framework: Error */)) {
	d._DeviceAttestationFailedForControllerOpaqueDeviceHandleError = f
}

// DeviceAttestationCompletedForDeviceAttestationDeviceInfoError implements the PMTRDeviceAttestationDelegate interface.
func (d *MTRDeviceAttestationDelegate) DeviceAttestationCompletedForDeviceAttestationDeviceInfoError(controller IMTRDeviceController, device unsafe.Pointer, attestationDeviceInfo IMTRDeviceAttestationDeviceInfo, error_ objc.IObject /* cross-framework: Error */) {
	if d._DeviceAttestationCompletedForDeviceAttestationDeviceInfoError != nil {
		d._DeviceAttestationCompletedForDeviceAttestationDeviceInfoError(controller, device, attestationDeviceInfo, error_)
	}
}

// HasDeviceAttestationCompletedForDeviceAttestationDeviceInfoError returns true if a handler for DeviceAttestationCompletedForDeviceAttestationDeviceInfoError has been set.
func (d *MTRDeviceAttestationDelegate) HasDeviceAttestationCompletedForDeviceAttestationDeviceInfoError() bool {
	return d._DeviceAttestationCompletedForDeviceAttestationDeviceInfoError != nil
}

// DeviceAttestationFailedForDeviceError implements the PMTRDeviceAttestationDelegate interface.
func (d *MTRDeviceAttestationDelegate) DeviceAttestationFailedForDeviceError(controller IMTRDeviceController, device unsafe.Pointer, error_ objc.IObject /* cross-framework: Error */) {
	if d._DeviceAttestationFailedForDeviceError != nil {
		d._DeviceAttestationFailedForDeviceError(controller, device, error_)
	}
}

// HasDeviceAttestationFailedForDeviceError returns true if a handler for DeviceAttestationFailedForDeviceError has been set.
func (d *MTRDeviceAttestationDelegate) HasDeviceAttestationFailedForDeviceError() bool {
	return d._DeviceAttestationFailedForDeviceError != nil
}

// DeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError implements the PMTRDeviceAttestationDelegate interface.
func (d *MTRDeviceAttestationDelegate) DeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError(controller IMTRDeviceController, opaqueDeviceHandle unsafe.Pointer, attestationDeviceInfo IMTRDeviceAttestationDeviceInfo, error_ objc.IObject /* cross-framework: Error */) {
	if d._DeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError != nil {
		d._DeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError(controller, opaqueDeviceHandle, attestationDeviceInfo, error_)
	}
}

// HasDeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError returns true if a handler for DeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError has been set.
func (d *MTRDeviceAttestationDelegate) HasDeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError() bool {
	return d._DeviceAttestationCompletedForControllerOpaqueDeviceHandleAttestationDeviceInfoError != nil
}

// DeviceAttestationFailedForControllerOpaqueDeviceHandleError implements the PMTRDeviceAttestationDelegate interface.
func (d *MTRDeviceAttestationDelegate) DeviceAttestationFailedForControllerOpaqueDeviceHandleError(controller IMTRDeviceController, opaqueDeviceHandle unsafe.Pointer, error_ objc.IObject /* cross-framework: Error */) {
	if d._DeviceAttestationFailedForControllerOpaqueDeviceHandleError != nil {
		d._DeviceAttestationFailedForControllerOpaqueDeviceHandleError(controller, opaqueDeviceHandle, error_)
	}
}

// HasDeviceAttestationFailedForControllerOpaqueDeviceHandleError returns true if a handler for DeviceAttestationFailedForControllerOpaqueDeviceHandleError has been set.
func (d *MTRDeviceAttestationDelegate) HasDeviceAttestationFailedForControllerOpaqueDeviceHandleError() bool {
	return d._DeviceAttestationFailedForControllerOpaqueDeviceHandleError != nil
}
