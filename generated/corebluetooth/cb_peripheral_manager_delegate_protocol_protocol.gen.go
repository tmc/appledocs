// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PCBPeripheralManagerDelegate is the CBPeripheralManagerDelegate protocol interface.
//
// A protocol that provides updates for local peripheral state and interactions with remote central devices.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 4.0+
//
// See: doc://com.apple.corebluetooth/documentation/CoreBluetooth/CBPeripheralManagerDelegate
type PCBPeripheralManagerDelegate interface {
	// Required methods
	PeripheralManagerDidUpdateState(peripheral ICBPeripheralManager)/* debug [protocol_interface/required_method]: PeripheralManagerDidUpdateState */
	// Optional methods
	PeripheralManagerCentralDidSubscribeToCharacteristic(peripheral ICBPeripheralManager, central ICBCentral, characteristic ICBCharacteristic)
	HasPeripheralManagerCentralDidSubscribeToCharacteristic() bool
	PeripheralManagerCentralDidUnsubscribeFromCharacteristic(peripheral ICBPeripheralManager, central ICBCentral, characteristic ICBCharacteristic)
	HasPeripheralManagerCentralDidUnsubscribeFromCharacteristic() bool
	PeripheralManagerDidAddServiceError(peripheral ICBPeripheralManager, service ICBService, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralManagerDidAddServiceError() bool
	PeripheralManagerDidOpenL2CAPChannelError(peripheral ICBPeripheralManager, channel ICBL2CAPChannel, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralManagerDidOpenL2CAPChannelError() bool
	PeripheralManagerDidPublishL2CAPChannelError(peripheral ICBPeripheralManager, PSM CBL2CAPPSM /* typedef */, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralManagerDidPublishL2CAPChannelError() bool
	PeripheralManagerDidReceiveReadRequest(peripheral ICBPeripheralManager, request ICBATTRequest)
	HasPeripheralManagerDidReceiveReadRequest() bool
	PeripheralManagerDidReceiveWriteRequests(peripheral ICBPeripheralManager, requests []CBATTRequest)
	HasPeripheralManagerDidReceiveWriteRequests() bool
	PeripheralManagerDidUnpublishL2CAPChannelError(peripheral ICBPeripheralManager, PSM CBL2CAPPSM /* typedef */, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralManagerDidUnpublishL2CAPChannelError() bool
	PeripheralManagerWillRestoreState(peripheral ICBPeripheralManager, dict foundation.IDictionary)
	HasPeripheralManagerWillRestoreState() bool
	PeripheralManagerDidStartAdvertisingError(peripheral ICBPeripheralManager, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralManagerDidStartAdvertisingError() bool
	PeripheralManagerIsReadyToUpdateSubscribers(peripheral ICBPeripheralManager)
	HasPeripheralManagerIsReadyToUpdateSubscribers() bool
}

// CBPeripheralManagerDelegate is a delegate implementation builder for the PCBPeripheralManagerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CBPeripheralManagerDelegate struct {
	_PeripheralManagerCentralDidSubscribeToCharacteristic func(peripheral ICBPeripheralManager, central ICBCentral, characteristic ICBCharacteristic)
	_PeripheralManagerCentralDidUnsubscribeFromCharacteristic func(peripheral ICBPeripheralManager, central ICBCentral, characteristic ICBCharacteristic)
	_PeripheralManagerDidAddServiceError func(peripheral ICBPeripheralManager, service ICBService, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralManagerDidOpenL2CAPChannelError func(peripheral ICBPeripheralManager, channel ICBL2CAPChannel, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralManagerDidPublishL2CAPChannelError func(peripheral ICBPeripheralManager, PSM CBL2CAPPSM /* typedef */, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralManagerDidReceiveReadRequest func(peripheral ICBPeripheralManager, request ICBATTRequest)
	_PeripheralManagerDidReceiveWriteRequests func(peripheral ICBPeripheralManager, requests []CBATTRequest)
	_PeripheralManagerDidUnpublishL2CAPChannelError func(peripheral ICBPeripheralManager, PSM CBL2CAPPSM /* typedef */, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralManagerWillRestoreState func(peripheral ICBPeripheralManager, dict foundation.IDictionary)
	_PeripheralManagerDidStartAdvertisingError func(peripheral ICBPeripheralManager, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralManagerIsReadyToUpdateSubscribers func(peripheral ICBPeripheralManager)
	_PeripheralManagerDidUpdateState func(peripheral ICBPeripheralManager)
}

// SetPeripheralManagerCentralDidSubscribeToCharacteristic sets the handler for the PeripheralManagerCentralDidSubscribeToCharacteristic delegate method.
//
// Tells the delegate that a remote central device subscribed to a characteristic’s value.
func (d *CBPeripheralManagerDelegate) SetPeripheralManagerCentralDidSubscribeToCharacteristic(f func(peripheral ICBPeripheralManager, central ICBCentral, characteristic ICBCharacteristic)) {
	d._PeripheralManagerCentralDidSubscribeToCharacteristic = f
}

// SetPeripheralManagerCentralDidUnsubscribeFromCharacteristic sets the handler for the PeripheralManagerCentralDidUnsubscribeFromCharacteristic delegate method.
//
// Tells the delegate that a remote central device unsubscribed from a characteristic’s value.
func (d *CBPeripheralManagerDelegate) SetPeripheralManagerCentralDidUnsubscribeFromCharacteristic(f func(peripheral ICBPeripheralManager, central ICBCentral, characteristic ICBCharacteristic)) {
	d._PeripheralManagerCentralDidUnsubscribeFromCharacteristic = f
}

// SetPeripheralManagerDidAddServiceError sets the handler for the PeripheralManagerDidAddServiceError delegate method.
//
// Tells the delegate the peripheral manager published a service to the local GATT database.
func (d *CBPeripheralManagerDelegate) SetPeripheralManagerDidAddServiceError(f func(peripheral ICBPeripheralManager, service ICBService, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralManagerDidAddServiceError = f
}

// SetPeripheralManagerDidOpenL2CAPChannelError sets the handler for the PeripheralManagerDidOpenL2CAPChannelError delegate method.
//
// Tells the delegate that the peripheral manager opened an L2CAP channel.
func (d *CBPeripheralManagerDelegate) SetPeripheralManagerDidOpenL2CAPChannelError(f func(peripheral ICBPeripheralManager, channel ICBL2CAPChannel, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralManagerDidOpenL2CAPChannelError = f
}

// SetPeripheralManagerDidPublishL2CAPChannelError sets the handler for the PeripheralManagerDidPublishL2CAPChannelError delegate method.
//
// Tells the delegate that the peripheral manager created a listener for incoming L2CAP channel connections.
func (d *CBPeripheralManagerDelegate) SetPeripheralManagerDidPublishL2CAPChannelError(f func(peripheral ICBPeripheralManager, PSM CBL2CAPPSM /* typedef */, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralManagerDidPublishL2CAPChannelError = f
}

// SetPeripheralManagerDidReceiveReadRequest sets the handler for the PeripheralManagerDidReceiveReadRequest delegate method.
//
// Tells the delegate that a local peripheral received an Attribute Protocol (ATT) read request for a characteristic with a dynamic value.
func (d *CBPeripheralManagerDelegate) SetPeripheralManagerDidReceiveReadRequest(f func(peripheral ICBPeripheralManager, request ICBATTRequest)) {
	d._PeripheralManagerDidReceiveReadRequest = f
}

// SetPeripheralManagerDidReceiveWriteRequests sets the handler for the PeripheralManagerDidReceiveWriteRequests delegate method.
//
// Tells the delegate that a local peripheral device received an Attribute Protocol (ATT) write request for a characteristic with a dynamic value.
func (d *CBPeripheralManagerDelegate) SetPeripheralManagerDidReceiveWriteRequests(f func(peripheral ICBPeripheralManager, requests []CBATTRequest)) {
	d._PeripheralManagerDidReceiveWriteRequests = f
}

// SetPeripheralManagerDidUnpublishL2CAPChannelError sets the handler for the PeripheralManagerDidUnpublishL2CAPChannelError delegate method.
//
// Tells the delegate that the peripheral manager removed a published service from the local system.
func (d *CBPeripheralManagerDelegate) SetPeripheralManagerDidUnpublishL2CAPChannelError(f func(peripheral ICBPeripheralManager, PSM CBL2CAPPSM /* typedef */, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralManagerDidUnpublishL2CAPChannelError = f
}

// SetPeripheralManagerWillRestoreState sets the handler for the PeripheralManagerWillRestoreState delegate method.
//
// Tells the delegate the system is about to restore the peripheral manager.
func (d *CBPeripheralManagerDelegate) SetPeripheralManagerWillRestoreState(f func(peripheral ICBPeripheralManager, dict foundation.IDictionary)) {
	d._PeripheralManagerWillRestoreState = f
}

// SetPeripheralManagerDidStartAdvertisingError sets the handler for the PeripheralManagerDidStartAdvertisingError delegate method.
//
// Tells the delegate the peripheral manager started advertising the local peripheral device’s data.
func (d *CBPeripheralManagerDelegate) SetPeripheralManagerDidStartAdvertisingError(f func(peripheral ICBPeripheralManager, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralManagerDidStartAdvertisingError = f
}

// SetPeripheralManagerIsReadyToUpdateSubscribers sets the handler for the PeripheralManagerIsReadyToUpdateSubscribers delegate method.
//
// Tells the delegate that a local peripheral device is ready to send characteristic value updates.
func (d *CBPeripheralManagerDelegate) SetPeripheralManagerIsReadyToUpdateSubscribers(f func(peripheral ICBPeripheralManager)) {
	d._PeripheralManagerIsReadyToUpdateSubscribers = f
}

// SetPeripheralManagerDidUpdateState sets the handler for the PeripheralManagerDidUpdateState delegate method.
//
// Tells the delegate the peripheral manager’s state updated.
func (d *CBPeripheralManagerDelegate) SetPeripheralManagerDidUpdateState(f func(peripheral ICBPeripheralManager)) {
	d._PeripheralManagerDidUpdateState = f
}

// PeripheralManagerCentralDidSubscribeToCharacteristic implements the PCBPeripheralManagerDelegate interface.
func (d *CBPeripheralManagerDelegate) PeripheralManagerCentralDidSubscribeToCharacteristic(peripheral ICBPeripheralManager, central ICBCentral, characteristic ICBCharacteristic) {
	if d._PeripheralManagerCentralDidSubscribeToCharacteristic != nil {
		d._PeripheralManagerCentralDidSubscribeToCharacteristic(peripheral, central, characteristic)
	}
}

// HasPeripheralManagerCentralDidSubscribeToCharacteristic returns true if a handler for PeripheralManagerCentralDidSubscribeToCharacteristic has been set.
func (d *CBPeripheralManagerDelegate) HasPeripheralManagerCentralDidSubscribeToCharacteristic() bool {
	return d._PeripheralManagerCentralDidSubscribeToCharacteristic != nil
}

// PeripheralManagerCentralDidUnsubscribeFromCharacteristic implements the PCBPeripheralManagerDelegate interface.
func (d *CBPeripheralManagerDelegate) PeripheralManagerCentralDidUnsubscribeFromCharacteristic(peripheral ICBPeripheralManager, central ICBCentral, characteristic ICBCharacteristic) {
	if d._PeripheralManagerCentralDidUnsubscribeFromCharacteristic != nil {
		d._PeripheralManagerCentralDidUnsubscribeFromCharacteristic(peripheral, central, characteristic)
	}
}

// HasPeripheralManagerCentralDidUnsubscribeFromCharacteristic returns true if a handler for PeripheralManagerCentralDidUnsubscribeFromCharacteristic has been set.
func (d *CBPeripheralManagerDelegate) HasPeripheralManagerCentralDidUnsubscribeFromCharacteristic() bool {
	return d._PeripheralManagerCentralDidUnsubscribeFromCharacteristic != nil
}

// PeripheralManagerDidAddServiceError implements the PCBPeripheralManagerDelegate interface.
func (d *CBPeripheralManagerDelegate) PeripheralManagerDidAddServiceError(peripheral ICBPeripheralManager, service ICBService, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralManagerDidAddServiceError != nil {
		d._PeripheralManagerDidAddServiceError(peripheral, service, error_)
	}
}

// HasPeripheralManagerDidAddServiceError returns true if a handler for PeripheralManagerDidAddServiceError has been set.
func (d *CBPeripheralManagerDelegate) HasPeripheralManagerDidAddServiceError() bool {
	return d._PeripheralManagerDidAddServiceError != nil
}

// PeripheralManagerDidOpenL2CAPChannelError implements the PCBPeripheralManagerDelegate interface.
func (d *CBPeripheralManagerDelegate) PeripheralManagerDidOpenL2CAPChannelError(peripheral ICBPeripheralManager, channel ICBL2CAPChannel, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralManagerDidOpenL2CAPChannelError != nil {
		d._PeripheralManagerDidOpenL2CAPChannelError(peripheral, channel, error_)
	}
}

// HasPeripheralManagerDidOpenL2CAPChannelError returns true if a handler for PeripheralManagerDidOpenL2CAPChannelError has been set.
func (d *CBPeripheralManagerDelegate) HasPeripheralManagerDidOpenL2CAPChannelError() bool {
	return d._PeripheralManagerDidOpenL2CAPChannelError != nil
}

// PeripheralManagerDidPublishL2CAPChannelError implements the PCBPeripheralManagerDelegate interface.
func (d *CBPeripheralManagerDelegate) PeripheralManagerDidPublishL2CAPChannelError(peripheral ICBPeripheralManager, PSM CBL2CAPPSM /* typedef */, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralManagerDidPublishL2CAPChannelError != nil {
		d._PeripheralManagerDidPublishL2CAPChannelError(peripheral, PSM, error_)
	}
}

// HasPeripheralManagerDidPublishL2CAPChannelError returns true if a handler for PeripheralManagerDidPublishL2CAPChannelError has been set.
func (d *CBPeripheralManagerDelegate) HasPeripheralManagerDidPublishL2CAPChannelError() bool {
	return d._PeripheralManagerDidPublishL2CAPChannelError != nil
}

// PeripheralManagerDidReceiveReadRequest implements the PCBPeripheralManagerDelegate interface.
func (d *CBPeripheralManagerDelegate) PeripheralManagerDidReceiveReadRequest(peripheral ICBPeripheralManager, request ICBATTRequest) {
	if d._PeripheralManagerDidReceiveReadRequest != nil {
		d._PeripheralManagerDidReceiveReadRequest(peripheral, request)
	}
}

// HasPeripheralManagerDidReceiveReadRequest returns true if a handler for PeripheralManagerDidReceiveReadRequest has been set.
func (d *CBPeripheralManagerDelegate) HasPeripheralManagerDidReceiveReadRequest() bool {
	return d._PeripheralManagerDidReceiveReadRequest != nil
}

// PeripheralManagerDidReceiveWriteRequests implements the PCBPeripheralManagerDelegate interface.
func (d *CBPeripheralManagerDelegate) PeripheralManagerDidReceiveWriteRequests(peripheral ICBPeripheralManager, requests []CBATTRequest) {
	if d._PeripheralManagerDidReceiveWriteRequests != nil {
		d._PeripheralManagerDidReceiveWriteRequests(peripheral, requests)
	}
}

// HasPeripheralManagerDidReceiveWriteRequests returns true if a handler for PeripheralManagerDidReceiveWriteRequests has been set.
func (d *CBPeripheralManagerDelegate) HasPeripheralManagerDidReceiveWriteRequests() bool {
	return d._PeripheralManagerDidReceiveWriteRequests != nil
}

// PeripheralManagerDidUnpublishL2CAPChannelError implements the PCBPeripheralManagerDelegate interface.
func (d *CBPeripheralManagerDelegate) PeripheralManagerDidUnpublishL2CAPChannelError(peripheral ICBPeripheralManager, PSM CBL2CAPPSM /* typedef */, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralManagerDidUnpublishL2CAPChannelError != nil {
		d._PeripheralManagerDidUnpublishL2CAPChannelError(peripheral, PSM, error_)
	}
}

// HasPeripheralManagerDidUnpublishL2CAPChannelError returns true if a handler for PeripheralManagerDidUnpublishL2CAPChannelError has been set.
func (d *CBPeripheralManagerDelegate) HasPeripheralManagerDidUnpublishL2CAPChannelError() bool {
	return d._PeripheralManagerDidUnpublishL2CAPChannelError != nil
}

// PeripheralManagerWillRestoreState implements the PCBPeripheralManagerDelegate interface.
func (d *CBPeripheralManagerDelegate) PeripheralManagerWillRestoreState(peripheral ICBPeripheralManager, dict foundation.IDictionary) {
	if d._PeripheralManagerWillRestoreState != nil {
		d._PeripheralManagerWillRestoreState(peripheral, dict)
	}
}

// HasPeripheralManagerWillRestoreState returns true if a handler for PeripheralManagerWillRestoreState has been set.
func (d *CBPeripheralManagerDelegate) HasPeripheralManagerWillRestoreState() bool {
	return d._PeripheralManagerWillRestoreState != nil
}

// PeripheralManagerDidStartAdvertisingError implements the PCBPeripheralManagerDelegate interface.
func (d *CBPeripheralManagerDelegate) PeripheralManagerDidStartAdvertisingError(peripheral ICBPeripheralManager, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralManagerDidStartAdvertisingError != nil {
		d._PeripheralManagerDidStartAdvertisingError(peripheral, error_)
	}
}

// HasPeripheralManagerDidStartAdvertisingError returns true if a handler for PeripheralManagerDidStartAdvertisingError has been set.
func (d *CBPeripheralManagerDelegate) HasPeripheralManagerDidStartAdvertisingError() bool {
	return d._PeripheralManagerDidStartAdvertisingError != nil
}

// PeripheralManagerIsReadyToUpdateSubscribers implements the PCBPeripheralManagerDelegate interface.
func (d *CBPeripheralManagerDelegate) PeripheralManagerIsReadyToUpdateSubscribers(peripheral ICBPeripheralManager) {
	if d._PeripheralManagerIsReadyToUpdateSubscribers != nil {
		d._PeripheralManagerIsReadyToUpdateSubscribers(peripheral)
	}
}

// HasPeripheralManagerIsReadyToUpdateSubscribers returns true if a handler for PeripheralManagerIsReadyToUpdateSubscribers has been set.
func (d *CBPeripheralManagerDelegate) HasPeripheralManagerIsReadyToUpdateSubscribers() bool {
	return d._PeripheralManagerIsReadyToUpdateSubscribers != nil
}

// PeripheralManagerDidUpdateState implements the PCBPeripheralManagerDelegate interface.
func (d *CBPeripheralManagerDelegate) PeripheralManagerDidUpdateState(peripheral ICBPeripheralManager) {
	if d._PeripheralManagerDidUpdateState != nil {
		d._PeripheralManagerDidUpdateState(peripheral)
	}
}

// HasPeripheralManagerDidUpdateState returns true if a handler for PeripheralManagerDidUpdateState has been set.
func (d *CBPeripheralManagerDelegate) HasPeripheralManagerDidUpdateState() bool {
	return d._PeripheralManagerDidUpdateState != nil
}
