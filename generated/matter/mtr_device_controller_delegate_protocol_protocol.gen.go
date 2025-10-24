// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTRDeviceControllerDelegate is the MTRDeviceControllerDelegate protocol interface.
//
// Availability:
//   - Mac Catalyst 16.4+
//   - iOS 16.4+
//   - iPadOS 16.4+
//   - macOS 13.3+
//   - tvOS 16.4+
//   - visionOS 1.0+
//   - watchOS 9.4+
//
// See: doc://com.apple.matter/documentation/Matter/MTRDeviceControllerDelegate
type PMTRDeviceControllerDelegate interface {
	// Optional methods
	ControllerCommissioneeHasReceivedNetworkCredentials(controller IMTRDeviceController, nodeID objc.IObject /* cross-framework: NSNumber */)
	HasControllerCommissioneeHasReceivedNetworkCredentials() bool
	ControllerCommissioningComplete(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */)
	HasControllerCommissioningComplete() bool
	ControllerCommissioningCompleteNodeID(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */, nodeID objc.IObject /* cross-framework: NSNumber */)
	HasControllerCommissioningCompleteNodeID() bool
	ControllerCommissioningCompleteNodeIDMetrics(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */, nodeID objc.IObject /* cross-framework: NSNumber */, metrics objc.IObject /* cross-framework: MTRMetrics */)
	HasControllerCommissioningCompleteNodeIDMetrics() bool
	ControllerCommissioningSessionEstablishmentDone(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */)
	HasControllerCommissioningSessionEstablishmentDone() bool
	ControllerReadCommissioneeInfo(controller IMTRDeviceController, info IMTRCommissioneeInfo)
	HasControllerReadCommissioneeInfo() bool
	ControllerReadCommissioningInfo(controller IMTRDeviceController, info IMTRProductIdentity)
	HasControllerReadCommissioningInfo() bool
	ControllerStatusUpdate(controller IMTRDeviceController, status unsafe.Pointer)
	HasControllerStatusUpdate() bool
	ControllerSuspendedChangedTo(controller IMTRDeviceController, suspended bool)
	HasControllerSuspendedChangedTo() bool
	DevicesChangedForController(controller IMTRDeviceController)
	HasDevicesChangedForController() bool
}

// MTRDeviceControllerDelegate is a delegate implementation builder for the PMTRDeviceControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MTRDeviceControllerDelegate struct {
	_ControllerCommissioneeHasReceivedNetworkCredentials func(controller IMTRDeviceController, nodeID objc.IObject /* cross-framework: NSNumber */)
	_ControllerCommissioningComplete func(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */)
	_ControllerCommissioningCompleteNodeID func(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */, nodeID objc.IObject /* cross-framework: NSNumber */)
	_ControllerCommissioningCompleteNodeIDMetrics func(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */, nodeID objc.IObject /* cross-framework: NSNumber */, metrics objc.IObject /* cross-framework: MTRMetrics */)
	_ControllerCommissioningSessionEstablishmentDone func(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */)
	_ControllerReadCommissioneeInfo func(controller IMTRDeviceController, info IMTRCommissioneeInfo)
	_ControllerReadCommissioningInfo func(controller IMTRDeviceController, info IMTRProductIdentity)
	_ControllerStatusUpdate func(controller IMTRDeviceController, status unsafe.Pointer)
	_ControllerSuspendedChangedTo func(controller IMTRDeviceController, suspended bool)
	_DevicesChangedForController func(controller IMTRDeviceController)
}

// SetControllerCommissioneeHasReceivedNetworkCredentials sets the handler for the ControllerCommissioneeHasReceivedNetworkCredentials delegate method.
//
// Notify the delegate that we have successfully communicated the network   credentials to the device being commissioned and are about to tell it to join   that network.  Note that for devices that are already on-network this   notification will not happen.
func (d *MTRDeviceControllerDelegate) SetControllerCommissioneeHasReceivedNetworkCredentials(f func(controller IMTRDeviceController, nodeID objc.IObject /* cross-framework: NSNumber */)) {
	d._ControllerCommissioneeHasReceivedNetworkCredentials = f
}

// SetControllerCommissioningComplete sets the handler for the ControllerCommissioningComplete delegate method.
func (d *MTRDeviceControllerDelegate) SetControllerCommissioningComplete(f func(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */)) {
	d._ControllerCommissioningComplete = f
}

// SetControllerCommissioningCompleteNodeID sets the handler for the ControllerCommissioningCompleteNodeID delegate method.
func (d *MTRDeviceControllerDelegate) SetControllerCommissioningCompleteNodeID(f func(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */, nodeID objc.IObject /* cross-framework: NSNumber */)) {
	d._ControllerCommissioningCompleteNodeID = f
}

// SetControllerCommissioningCompleteNodeIDMetrics sets the handler for the ControllerCommissioningCompleteNodeIDMetrics delegate method.
func (d *MTRDeviceControllerDelegate) SetControllerCommissioningCompleteNodeIDMetrics(f func(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */, nodeID objc.IObject /* cross-framework: NSNumber */, metrics objc.IObject /* cross-framework: MTRMetrics */)) {
	d._ControllerCommissioningCompleteNodeIDMetrics = f
}

// SetControllerCommissioningSessionEstablishmentDone sets the handler for the ControllerCommissioningSessionEstablishmentDone delegate method.
func (d *MTRDeviceControllerDelegate) SetControllerCommissioningSessionEstablishmentDone(f func(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */)) {
	d._ControllerCommissioningSessionEstablishmentDone = f
}

// SetControllerReadCommissioneeInfo sets the handler for the ControllerReadCommissioneeInfo delegate method.
//
// Notify the delegate when commissioning infomation has been read from the commissionee.
func (d *MTRDeviceControllerDelegate) SetControllerReadCommissioneeInfo(f func(controller IMTRDeviceController, info IMTRCommissioneeInfo)) {
	d._ControllerReadCommissioneeInfo = f
}

// SetControllerReadCommissioningInfo sets the handler for the ControllerReadCommissioningInfo delegate method.
func (d *MTRDeviceControllerDelegate) SetControllerReadCommissioningInfo(f func(controller IMTRDeviceController, info IMTRProductIdentity)) {
	d._ControllerReadCommissioningInfo = f
}

// SetControllerStatusUpdate sets the handler for the ControllerStatusUpdate delegate method.
func (d *MTRDeviceControllerDelegate) SetControllerStatusUpdate(f func(controller IMTRDeviceController, status unsafe.Pointer)) {
	d._ControllerStatusUpdate = f
}

// SetControllerSuspendedChangedTo sets the handler for the ControllerSuspendedChangedTo delegate method.
//
// Notify the delegate when the suspended state changed of the controller, after this happens   the controller will be in the specified state.
func (d *MTRDeviceControllerDelegate) SetControllerSuspendedChangedTo(f func(controller IMTRDeviceController, suspended bool)) {
	d._ControllerSuspendedChangedTo = f
}

// SetDevicesChangedForController sets the handler for the DevicesChangedForController delegate method.
//
// Notify the delegate when the list of MTRDevice objects in memory has changed.
func (d *MTRDeviceControllerDelegate) SetDevicesChangedForController(f func(controller IMTRDeviceController)) {
	d._DevicesChangedForController = f
}

// ControllerCommissioneeHasReceivedNetworkCredentials implements the PMTRDeviceControllerDelegate interface.
func (d *MTRDeviceControllerDelegate) ControllerCommissioneeHasReceivedNetworkCredentials(controller IMTRDeviceController, nodeID objc.IObject /* cross-framework: NSNumber */) {
	if d._ControllerCommissioneeHasReceivedNetworkCredentials != nil {
		d._ControllerCommissioneeHasReceivedNetworkCredentials(controller, nodeID)
	}
}

// HasControllerCommissioneeHasReceivedNetworkCredentials returns true if a handler for ControllerCommissioneeHasReceivedNetworkCredentials has been set.
func (d *MTRDeviceControllerDelegate) HasControllerCommissioneeHasReceivedNetworkCredentials() bool {
	return d._ControllerCommissioneeHasReceivedNetworkCredentials != nil
}

// ControllerCommissioningComplete implements the PMTRDeviceControllerDelegate interface.
func (d *MTRDeviceControllerDelegate) ControllerCommissioningComplete(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */) {
	if d._ControllerCommissioningComplete != nil {
		d._ControllerCommissioningComplete(controller, error_)
	}
}

// HasControllerCommissioningComplete returns true if a handler for ControllerCommissioningComplete has been set.
func (d *MTRDeviceControllerDelegate) HasControllerCommissioningComplete() bool {
	return d._ControllerCommissioningComplete != nil
}

// ControllerCommissioningCompleteNodeID implements the PMTRDeviceControllerDelegate interface.
func (d *MTRDeviceControllerDelegate) ControllerCommissioningCompleteNodeID(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */, nodeID objc.IObject /* cross-framework: NSNumber */) {
	if d._ControllerCommissioningCompleteNodeID != nil {
		d._ControllerCommissioningCompleteNodeID(controller, error_, nodeID)
	}
}

// HasControllerCommissioningCompleteNodeID returns true if a handler for ControllerCommissioningCompleteNodeID has been set.
func (d *MTRDeviceControllerDelegate) HasControllerCommissioningCompleteNodeID() bool {
	return d._ControllerCommissioningCompleteNodeID != nil
}

// ControllerCommissioningCompleteNodeIDMetrics implements the PMTRDeviceControllerDelegate interface.
func (d *MTRDeviceControllerDelegate) ControllerCommissioningCompleteNodeIDMetrics(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */, nodeID objc.IObject /* cross-framework: NSNumber */, metrics objc.IObject /* cross-framework: MTRMetrics */) {
	if d._ControllerCommissioningCompleteNodeIDMetrics != nil {
		d._ControllerCommissioningCompleteNodeIDMetrics(controller, error_, nodeID, metrics)
	}
}

// HasControllerCommissioningCompleteNodeIDMetrics returns true if a handler for ControllerCommissioningCompleteNodeIDMetrics has been set.
func (d *MTRDeviceControllerDelegate) HasControllerCommissioningCompleteNodeIDMetrics() bool {
	return d._ControllerCommissioningCompleteNodeIDMetrics != nil
}

// ControllerCommissioningSessionEstablishmentDone implements the PMTRDeviceControllerDelegate interface.
func (d *MTRDeviceControllerDelegate) ControllerCommissioningSessionEstablishmentDone(controller IMTRDeviceController, error_ objc.IObject /* cross-framework: Error */) {
	if d._ControllerCommissioningSessionEstablishmentDone != nil {
		d._ControllerCommissioningSessionEstablishmentDone(controller, error_)
	}
}

// HasControllerCommissioningSessionEstablishmentDone returns true if a handler for ControllerCommissioningSessionEstablishmentDone has been set.
func (d *MTRDeviceControllerDelegate) HasControllerCommissioningSessionEstablishmentDone() bool {
	return d._ControllerCommissioningSessionEstablishmentDone != nil
}

// ControllerReadCommissioneeInfo implements the PMTRDeviceControllerDelegate interface.
func (d *MTRDeviceControllerDelegate) ControllerReadCommissioneeInfo(controller IMTRDeviceController, info IMTRCommissioneeInfo) {
	if d._ControllerReadCommissioneeInfo != nil {
		d._ControllerReadCommissioneeInfo(controller, info)
	}
}

// HasControllerReadCommissioneeInfo returns true if a handler for ControllerReadCommissioneeInfo has been set.
func (d *MTRDeviceControllerDelegate) HasControllerReadCommissioneeInfo() bool {
	return d._ControllerReadCommissioneeInfo != nil
}

// ControllerReadCommissioningInfo implements the PMTRDeviceControllerDelegate interface.
func (d *MTRDeviceControllerDelegate) ControllerReadCommissioningInfo(controller IMTRDeviceController, info IMTRProductIdentity) {
	if d._ControllerReadCommissioningInfo != nil {
		d._ControllerReadCommissioningInfo(controller, info)
	}
}

// HasControllerReadCommissioningInfo returns true if a handler for ControllerReadCommissioningInfo has been set.
func (d *MTRDeviceControllerDelegate) HasControllerReadCommissioningInfo() bool {
	return d._ControllerReadCommissioningInfo != nil
}

// ControllerStatusUpdate implements the PMTRDeviceControllerDelegate interface.
func (d *MTRDeviceControllerDelegate) ControllerStatusUpdate(controller IMTRDeviceController, status unsafe.Pointer) {
	if d._ControllerStatusUpdate != nil {
		d._ControllerStatusUpdate(controller, status)
	}
}

// HasControllerStatusUpdate returns true if a handler for ControllerStatusUpdate has been set.
func (d *MTRDeviceControllerDelegate) HasControllerStatusUpdate() bool {
	return d._ControllerStatusUpdate != nil
}

// ControllerSuspendedChangedTo implements the PMTRDeviceControllerDelegate interface.
func (d *MTRDeviceControllerDelegate) ControllerSuspendedChangedTo(controller IMTRDeviceController, suspended bool) {
	if d._ControllerSuspendedChangedTo != nil {
		d._ControllerSuspendedChangedTo(controller, suspended)
	}
}

// HasControllerSuspendedChangedTo returns true if a handler for ControllerSuspendedChangedTo has been set.
func (d *MTRDeviceControllerDelegate) HasControllerSuspendedChangedTo() bool {
	return d._ControllerSuspendedChangedTo != nil
}

// DevicesChangedForController implements the PMTRDeviceControllerDelegate interface.
func (d *MTRDeviceControllerDelegate) DevicesChangedForController(controller IMTRDeviceController) {
	if d._DevicesChangedForController != nil {
		d._DevicesChangedForController(controller)
	}
}

// HasDevicesChangedForController returns true if a handler for DevicesChangedForController has been set.
func (d *MTRDeviceControllerDelegate) HasDevicesChangedForController() bool {
	return d._DevicesChangedForController != nil
}
