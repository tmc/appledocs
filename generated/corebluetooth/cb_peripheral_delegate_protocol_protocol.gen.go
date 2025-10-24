// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PCBPeripheralDelegate is the CBPeripheralDelegate protocol interface.
//
// A protocol that provides updates on the use of a peripheral’s services.
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
// See: doc://com.apple.corebluetooth/documentation/CoreBluetooth/CBPeripheralDelegate
type PCBPeripheralDelegate interface {
	// Optional methods
	PeripheralDidDiscoverCharacteristicsForServiceError(peripheral ICBPeripheral, service ICBService, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralDidDiscoverCharacteristicsForServiceError() bool
	PeripheralDidDiscoverDescriptorsForCharacteristicError(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralDidDiscoverDescriptorsForCharacteristicError() bool
	PeripheralDidDiscoverIncludedServicesForServiceError(peripheral ICBPeripheral, service ICBService, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralDidDiscoverIncludedServicesForServiceError() bool
	PeripheralDidDiscoverServices(peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralDidDiscoverServices() bool
	PeripheralDidModifyServices(peripheral ICBPeripheral, invalidatedServices []CBService)
	HasPeripheralDidModifyServices() bool
	PeripheralDidOpenL2CAPChannelError(peripheral ICBPeripheral, channel ICBL2CAPChannel, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralDidOpenL2CAPChannelError() bool
	PeripheralDidReadRSSIError(peripheral ICBPeripheral, RSSI objc.IObject /* cross-framework: NSNumber */, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralDidReadRSSIError() bool
	PeripheralDidUpdateNotificationStateForCharacteristicError(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralDidUpdateNotificationStateForCharacteristicError() bool
	PeripheralDidUpdateValueForDescriptorError(peripheral ICBPeripheral, descriptor ICBDescriptor, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralDidUpdateValueForDescriptorError() bool
	PeripheralDidUpdateValueForCharacteristicError(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralDidUpdateValueForCharacteristicError() bool
	PeripheralDidWriteValueForDescriptorError(peripheral ICBPeripheral, descriptor ICBDescriptor, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralDidWriteValueForDescriptorError() bool
	PeripheralDidWriteValueForCharacteristicError(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralDidWriteValueForCharacteristicError() bool
	PeripheralDidUpdateName(peripheral ICBPeripheral)
	HasPeripheralDidUpdateName() bool
	PeripheralDidUpdateRSSIError(peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */)
	HasPeripheralDidUpdateRSSIError() bool
	PeripheralIsReadyToSendWriteWithoutResponse(peripheral ICBPeripheral)
	HasPeripheralIsReadyToSendWriteWithoutResponse() bool
}

// CBPeripheralDelegate is a delegate implementation builder for the PCBPeripheralDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CBPeripheralDelegate struct {
	_PeripheralDidDiscoverCharacteristicsForServiceError func(peripheral ICBPeripheral, service ICBService, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralDidDiscoverDescriptorsForCharacteristicError func(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralDidDiscoverIncludedServicesForServiceError func(peripheral ICBPeripheral, service ICBService, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralDidDiscoverServices func(peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralDidModifyServices func(peripheral ICBPeripheral, invalidatedServices []CBService)
	_PeripheralDidOpenL2CAPChannelError func(peripheral ICBPeripheral, channel ICBL2CAPChannel, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralDidReadRSSIError func(peripheral ICBPeripheral, RSSI objc.IObject /* cross-framework: NSNumber */, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralDidUpdateNotificationStateForCharacteristicError func(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralDidUpdateValueForDescriptorError func(peripheral ICBPeripheral, descriptor ICBDescriptor, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralDidUpdateValueForCharacteristicError func(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralDidWriteValueForDescriptorError func(peripheral ICBPeripheral, descriptor ICBDescriptor, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralDidWriteValueForCharacteristicError func(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralDidUpdateName func(peripheral ICBPeripheral)
	_PeripheralDidUpdateRSSIError func(peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */)
	_PeripheralIsReadyToSendWriteWithoutResponse func(peripheral ICBPeripheral)
}

// SetPeripheralDidDiscoverCharacteristicsForServiceError sets the handler for the PeripheralDidDiscoverCharacteristicsForServiceError delegate method.
//
// Tells the delegate that the peripheral found characteristics for a service.
func (d *CBPeripheralDelegate) SetPeripheralDidDiscoverCharacteristicsForServiceError(f func(peripheral ICBPeripheral, service ICBService, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralDidDiscoverCharacteristicsForServiceError = f
}

// SetPeripheralDidDiscoverDescriptorsForCharacteristicError sets the handler for the PeripheralDidDiscoverDescriptorsForCharacteristicError delegate method.
//
// Tells the delegate that the peripheral found descriptors for a characteristic.
func (d *CBPeripheralDelegate) SetPeripheralDidDiscoverDescriptorsForCharacteristicError(f func(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralDidDiscoverDescriptorsForCharacteristicError = f
}

// SetPeripheralDidDiscoverIncludedServicesForServiceError sets the handler for the PeripheralDidDiscoverIncludedServicesForServiceError delegate method.
//
// Tells the delegate that discovering included services within the indicated service completed.
func (d *CBPeripheralDelegate) SetPeripheralDidDiscoverIncludedServicesForServiceError(f func(peripheral ICBPeripheral, service ICBService, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralDidDiscoverIncludedServicesForServiceError = f
}

// SetPeripheralDidDiscoverServices sets the handler for the PeripheralDidDiscoverServices delegate method.
//
// Tells the delegate that peripheral service discovery succeeded.
func (d *CBPeripheralDelegate) SetPeripheralDidDiscoverServices(f func(peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralDidDiscoverServices = f
}

// SetPeripheralDidModifyServices sets the handler for the PeripheralDidModifyServices delegate method.
//
// Tells the delegate that a peripheral’s services changed.
func (d *CBPeripheralDelegate) SetPeripheralDidModifyServices(f func(peripheral ICBPeripheral, invalidatedServices []CBService)) {
	d._PeripheralDidModifyServices = f
}

// SetPeripheralDidOpenL2CAPChannelError sets the handler for the PeripheralDidOpenL2CAPChannelError delegate method.
//
// Delivers the result of an attempt to open an L2CAP channel.
func (d *CBPeripheralDelegate) SetPeripheralDidOpenL2CAPChannelError(f func(peripheral ICBPeripheral, channel ICBL2CAPChannel, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralDidOpenL2CAPChannelError = f
}

// SetPeripheralDidReadRSSIError sets the handler for the PeripheralDidReadRSSIError delegate method.
//
// Tells the delegate that retrieving the value of the peripheral’s current Received Signal Strength Indicator (RSSI) succeeded.
func (d *CBPeripheralDelegate) SetPeripheralDidReadRSSIError(f func(peripheral ICBPeripheral, RSSI objc.IObject /* cross-framework: NSNumber */, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralDidReadRSSIError = f
}

// SetPeripheralDidUpdateNotificationStateForCharacteristicError sets the handler for the PeripheralDidUpdateNotificationStateForCharacteristicError delegate method.
//
// Tells the delegate that the peripheral received a request to start or stop providing notifications for a specified characteristic’s value.
func (d *CBPeripheralDelegate) SetPeripheralDidUpdateNotificationStateForCharacteristicError(f func(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralDidUpdateNotificationStateForCharacteristicError = f
}

// SetPeripheralDidUpdateValueForDescriptorError sets the handler for the PeripheralDidUpdateValueForDescriptorError delegate method.
//
// Tells the delegate that retrieving a specified characteristic descriptor’s value succeeded.
func (d *CBPeripheralDelegate) SetPeripheralDidUpdateValueForDescriptorError(f func(peripheral ICBPeripheral, descriptor ICBDescriptor, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralDidUpdateValueForDescriptorError = f
}

// SetPeripheralDidUpdateValueForCharacteristicError sets the handler for the PeripheralDidUpdateValueForCharacteristicError delegate method.
//
// Tells the delegate that retrieving the specified characteristic’s value succeeded, or that the characteristic’s value changed.
func (d *CBPeripheralDelegate) SetPeripheralDidUpdateValueForCharacteristicError(f func(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralDidUpdateValueForCharacteristicError = f
}

// SetPeripheralDidWriteValueForDescriptorError sets the handler for the PeripheralDidWriteValueForDescriptorError delegate method.
//
// Tells the delegate that the peripheral successfully set a value for the descriptor.
func (d *CBPeripheralDelegate) SetPeripheralDidWriteValueForDescriptorError(f func(peripheral ICBPeripheral, descriptor ICBDescriptor, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralDidWriteValueForDescriptorError = f
}

// SetPeripheralDidWriteValueForCharacteristicError sets the handler for the PeripheralDidWriteValueForCharacteristicError delegate method.
//
// Tells the delegate that the peripheral successfully set a value for the characteristic.
func (d *CBPeripheralDelegate) SetPeripheralDidWriteValueForCharacteristicError(f func(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralDidWriteValueForCharacteristicError = f
}

// SetPeripheralDidUpdateName sets the handler for the PeripheralDidUpdateName delegate method.
//
// Tells the delegate that a peripheral’s name changed.
func (d *CBPeripheralDelegate) SetPeripheralDidUpdateName(f func(peripheral ICBPeripheral)) {
	d._PeripheralDidUpdateName = f
}

// SetPeripheralDidUpdateRSSIError sets the handler for the PeripheralDidUpdateRSSIError delegate method.
//
// Tells the delegate that retrieving the value of the peripheral’s current Received Signal Strength Indicator (RSSI) succeeded.
func (d *CBPeripheralDelegate) SetPeripheralDidUpdateRSSIError(f func(peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */)) {
	d._PeripheralDidUpdateRSSIError = f
}

// SetPeripheralIsReadyToSendWriteWithoutResponse sets the handler for the PeripheralIsReadyToSendWriteWithoutResponse delegate method.
//
// Tells the delegate that a peripheral is again ready to send characteristic updates.
func (d *CBPeripheralDelegate) SetPeripheralIsReadyToSendWriteWithoutResponse(f func(peripheral ICBPeripheral)) {
	d._PeripheralIsReadyToSendWriteWithoutResponse = f
}

// PeripheralDidDiscoverCharacteristicsForServiceError implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralDidDiscoverCharacteristicsForServiceError(peripheral ICBPeripheral, service ICBService, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralDidDiscoverCharacteristicsForServiceError != nil {
		d._PeripheralDidDiscoverCharacteristicsForServiceError(peripheral, service, error_)
	}
}

// HasPeripheralDidDiscoverCharacteristicsForServiceError returns true if a handler for PeripheralDidDiscoverCharacteristicsForServiceError has been set.
func (d *CBPeripheralDelegate) HasPeripheralDidDiscoverCharacteristicsForServiceError() bool {
	return d._PeripheralDidDiscoverCharacteristicsForServiceError != nil
}

// PeripheralDidDiscoverDescriptorsForCharacteristicError implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralDidDiscoverDescriptorsForCharacteristicError(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralDidDiscoverDescriptorsForCharacteristicError != nil {
		d._PeripheralDidDiscoverDescriptorsForCharacteristicError(peripheral, characteristic, error_)
	}
}

// HasPeripheralDidDiscoverDescriptorsForCharacteristicError returns true if a handler for PeripheralDidDiscoverDescriptorsForCharacteristicError has been set.
func (d *CBPeripheralDelegate) HasPeripheralDidDiscoverDescriptorsForCharacteristicError() bool {
	return d._PeripheralDidDiscoverDescriptorsForCharacteristicError != nil
}

// PeripheralDidDiscoverIncludedServicesForServiceError implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralDidDiscoverIncludedServicesForServiceError(peripheral ICBPeripheral, service ICBService, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralDidDiscoverIncludedServicesForServiceError != nil {
		d._PeripheralDidDiscoverIncludedServicesForServiceError(peripheral, service, error_)
	}
}

// HasPeripheralDidDiscoverIncludedServicesForServiceError returns true if a handler for PeripheralDidDiscoverIncludedServicesForServiceError has been set.
func (d *CBPeripheralDelegate) HasPeripheralDidDiscoverIncludedServicesForServiceError() bool {
	return d._PeripheralDidDiscoverIncludedServicesForServiceError != nil
}

// PeripheralDidDiscoverServices implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralDidDiscoverServices(peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralDidDiscoverServices != nil {
		d._PeripheralDidDiscoverServices(peripheral, error_)
	}
}

// HasPeripheralDidDiscoverServices returns true if a handler for PeripheralDidDiscoverServices has been set.
func (d *CBPeripheralDelegate) HasPeripheralDidDiscoverServices() bool {
	return d._PeripheralDidDiscoverServices != nil
}

// PeripheralDidModifyServices implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralDidModifyServices(peripheral ICBPeripheral, invalidatedServices []CBService) {
	if d._PeripheralDidModifyServices != nil {
		d._PeripheralDidModifyServices(peripheral, invalidatedServices)
	}
}

// HasPeripheralDidModifyServices returns true if a handler for PeripheralDidModifyServices has been set.
func (d *CBPeripheralDelegate) HasPeripheralDidModifyServices() bool {
	return d._PeripheralDidModifyServices != nil
}

// PeripheralDidOpenL2CAPChannelError implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralDidOpenL2CAPChannelError(peripheral ICBPeripheral, channel ICBL2CAPChannel, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralDidOpenL2CAPChannelError != nil {
		d._PeripheralDidOpenL2CAPChannelError(peripheral, channel, error_)
	}
}

// HasPeripheralDidOpenL2CAPChannelError returns true if a handler for PeripheralDidOpenL2CAPChannelError has been set.
func (d *CBPeripheralDelegate) HasPeripheralDidOpenL2CAPChannelError() bool {
	return d._PeripheralDidOpenL2CAPChannelError != nil
}

// PeripheralDidReadRSSIError implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralDidReadRSSIError(peripheral ICBPeripheral, RSSI objc.IObject /* cross-framework: NSNumber */, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralDidReadRSSIError != nil {
		d._PeripheralDidReadRSSIError(peripheral, RSSI, error_)
	}
}

// HasPeripheralDidReadRSSIError returns true if a handler for PeripheralDidReadRSSIError has been set.
func (d *CBPeripheralDelegate) HasPeripheralDidReadRSSIError() bool {
	return d._PeripheralDidReadRSSIError != nil
}

// PeripheralDidUpdateNotificationStateForCharacteristicError implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralDidUpdateNotificationStateForCharacteristicError(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralDidUpdateNotificationStateForCharacteristicError != nil {
		d._PeripheralDidUpdateNotificationStateForCharacteristicError(peripheral, characteristic, error_)
	}
}

// HasPeripheralDidUpdateNotificationStateForCharacteristicError returns true if a handler for PeripheralDidUpdateNotificationStateForCharacteristicError has been set.
func (d *CBPeripheralDelegate) HasPeripheralDidUpdateNotificationStateForCharacteristicError() bool {
	return d._PeripheralDidUpdateNotificationStateForCharacteristicError != nil
}

// PeripheralDidUpdateValueForDescriptorError implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralDidUpdateValueForDescriptorError(peripheral ICBPeripheral, descriptor ICBDescriptor, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralDidUpdateValueForDescriptorError != nil {
		d._PeripheralDidUpdateValueForDescriptorError(peripheral, descriptor, error_)
	}
}

// HasPeripheralDidUpdateValueForDescriptorError returns true if a handler for PeripheralDidUpdateValueForDescriptorError has been set.
func (d *CBPeripheralDelegate) HasPeripheralDidUpdateValueForDescriptorError() bool {
	return d._PeripheralDidUpdateValueForDescriptorError != nil
}

// PeripheralDidUpdateValueForCharacteristicError implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralDidUpdateValueForCharacteristicError(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralDidUpdateValueForCharacteristicError != nil {
		d._PeripheralDidUpdateValueForCharacteristicError(peripheral, characteristic, error_)
	}
}

// HasPeripheralDidUpdateValueForCharacteristicError returns true if a handler for PeripheralDidUpdateValueForCharacteristicError has been set.
func (d *CBPeripheralDelegate) HasPeripheralDidUpdateValueForCharacteristicError() bool {
	return d._PeripheralDidUpdateValueForCharacteristicError != nil
}

// PeripheralDidWriteValueForDescriptorError implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralDidWriteValueForDescriptorError(peripheral ICBPeripheral, descriptor ICBDescriptor, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralDidWriteValueForDescriptorError != nil {
		d._PeripheralDidWriteValueForDescriptorError(peripheral, descriptor, error_)
	}
}

// HasPeripheralDidWriteValueForDescriptorError returns true if a handler for PeripheralDidWriteValueForDescriptorError has been set.
func (d *CBPeripheralDelegate) HasPeripheralDidWriteValueForDescriptorError() bool {
	return d._PeripheralDidWriteValueForDescriptorError != nil
}

// PeripheralDidWriteValueForCharacteristicError implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralDidWriteValueForCharacteristicError(peripheral ICBPeripheral, characteristic ICBCharacteristic, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralDidWriteValueForCharacteristicError != nil {
		d._PeripheralDidWriteValueForCharacteristicError(peripheral, characteristic, error_)
	}
}

// HasPeripheralDidWriteValueForCharacteristicError returns true if a handler for PeripheralDidWriteValueForCharacteristicError has been set.
func (d *CBPeripheralDelegate) HasPeripheralDidWriteValueForCharacteristicError() bool {
	return d._PeripheralDidWriteValueForCharacteristicError != nil
}

// PeripheralDidUpdateName implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralDidUpdateName(peripheral ICBPeripheral) {
	if d._PeripheralDidUpdateName != nil {
		d._PeripheralDidUpdateName(peripheral)
	}
}

// HasPeripheralDidUpdateName returns true if a handler for PeripheralDidUpdateName has been set.
func (d *CBPeripheralDelegate) HasPeripheralDidUpdateName() bool {
	return d._PeripheralDidUpdateName != nil
}

// PeripheralDidUpdateRSSIError implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralDidUpdateRSSIError(peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */) {
	if d._PeripheralDidUpdateRSSIError != nil {
		d._PeripheralDidUpdateRSSIError(peripheral, error_)
	}
}

// HasPeripheralDidUpdateRSSIError returns true if a handler for PeripheralDidUpdateRSSIError has been set.
func (d *CBPeripheralDelegate) HasPeripheralDidUpdateRSSIError() bool {
	return d._PeripheralDidUpdateRSSIError != nil
}

// PeripheralIsReadyToSendWriteWithoutResponse implements the PCBPeripheralDelegate interface.
func (d *CBPeripheralDelegate) PeripheralIsReadyToSendWriteWithoutResponse(peripheral ICBPeripheral) {
	if d._PeripheralIsReadyToSendWriteWithoutResponse != nil {
		d._PeripheralIsReadyToSendWriteWithoutResponse(peripheral)
	}
}

// HasPeripheralIsReadyToSendWriteWithoutResponse returns true if a handler for PeripheralIsReadyToSendWriteWithoutResponse has been set.
func (d *CBPeripheralDelegate) HasPeripheralIsReadyToSendWriteWithoutResponse() bool {
	return d._PeripheralIsReadyToSendWriteWithoutResponse != nil
}
