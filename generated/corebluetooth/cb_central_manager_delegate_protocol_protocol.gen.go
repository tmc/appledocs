// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PCBCentralManagerDelegate is the CBCentralManagerDelegate protocol interface.
//
// A protocol that provides updates for the discovery and management of peripheral devices.
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
// See: doc://com.apple.corebluetooth/documentation/CoreBluetooth/CBCentralManagerDelegate
type PCBCentralManagerDelegate interface {
	// Required methods
	CentralManagerDidUpdateState(central ICBCentralManager)/* debug [protocol_interface/required_method]: CentralManagerDidUpdateState */
	// Optional methods
	CentralManagerConnectionEventDidOccurForPeripheral(central ICBCentralManager, event CBConnectionEvent, peripheral ICBPeripheral)
	HasCentralManagerConnectionEventDidOccurForPeripheral() bool
	CentralManagerDidConnectPeripheral(central ICBCentralManager, peripheral ICBPeripheral)
	HasCentralManagerDidConnectPeripheral() bool
	CentralManagerDidDisconnectPeripheralError(central ICBCentralManager, peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */)
	HasCentralManagerDidDisconnectPeripheralError() bool
	CentralManagerDidDisconnectPeripheralTimestampIsReconnectingError(central ICBCentralManager, peripheral ICBPeripheral, timestamp AbsoluteTime /* not a class type */, isReconnecting bool, error_ objc.IObject /* cross-framework: Error */)
	HasCentralManagerDidDisconnectPeripheralTimestampIsReconnectingError() bool
	CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI(central ICBCentralManager, peripheral ICBPeripheral, advertisementData foundation.IDictionary, RSSI objc.IObject /* cross-framework: NSNumber */)
	HasCentralManagerDidDiscoverPeripheralAdvertisementDataRSSI() bool
	CentralManagerDidFailToConnectPeripheralError(central ICBCentralManager, peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */)
	HasCentralManagerDidFailToConnectPeripheralError() bool
	CentralManagerDidUpdateANCSAuthorizationForPeripheral(central ICBCentralManager, peripheral ICBPeripheral)
	HasCentralManagerDidUpdateANCSAuthorizationForPeripheral() bool
	CentralManagerWillRestoreState(central ICBCentralManager, dict foundation.IDictionary)
	HasCentralManagerWillRestoreState() bool
}

// CBCentralManagerDelegate is a delegate implementation builder for the PCBCentralManagerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CBCentralManagerDelegate struct {
	_CentralManagerConnectionEventDidOccurForPeripheral func(central ICBCentralManager, event CBConnectionEvent, peripheral ICBPeripheral)
	_CentralManagerDidConnectPeripheral func(central ICBCentralManager, peripheral ICBPeripheral)
	_CentralManagerDidDisconnectPeripheralError func(central ICBCentralManager, peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */)
	_CentralManagerDidDisconnectPeripheralTimestampIsReconnectingError func(central ICBCentralManager, peripheral ICBPeripheral, timestamp AbsoluteTime /* not a class type */, isReconnecting bool, error_ objc.IObject /* cross-framework: Error */)
	_CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI func(central ICBCentralManager, peripheral ICBPeripheral, advertisementData foundation.IDictionary, RSSI objc.IObject /* cross-framework: NSNumber */)
	_CentralManagerDidFailToConnectPeripheralError func(central ICBCentralManager, peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */)
	_CentralManagerDidUpdateANCSAuthorizationForPeripheral func(central ICBCentralManager, peripheral ICBPeripheral)
	_CentralManagerWillRestoreState func(central ICBCentralManager, dict foundation.IDictionary)
	_CentralManagerDidUpdateState func(central ICBCentralManager)
}

// SetCentralManagerConnectionEventDidOccurForPeripheral sets the handler for the CentralManagerConnectionEventDidOccurForPeripheral delegate method.
//
// Tells the delegate that a connection event occurred which matches the registered options.
func (d *CBCentralManagerDelegate) SetCentralManagerConnectionEventDidOccurForPeripheral(f func(central ICBCentralManager, event CBConnectionEvent, peripheral ICBPeripheral)) {
	d._CentralManagerConnectionEventDidOccurForPeripheral = f
}

// SetCentralManagerDidConnectPeripheral sets the handler for the CentralManagerDidConnectPeripheral delegate method.
//
// Tells the delegate that the central manager connected to a peripheral.
func (d *CBCentralManagerDelegate) SetCentralManagerDidConnectPeripheral(f func(central ICBCentralManager, peripheral ICBPeripheral)) {
	d._CentralManagerDidConnectPeripheral = f
}

// SetCentralManagerDidDisconnectPeripheralError sets the handler for the CentralManagerDidDisconnectPeripheralError delegate method.
//
// Tells the delegate that the central manager disconnected from a peripheral.
func (d *CBCentralManagerDelegate) SetCentralManagerDidDisconnectPeripheralError(f func(central ICBCentralManager, peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */)) {
	d._CentralManagerDidDisconnectPeripheralError = f
}

// SetCentralManagerDidDisconnectPeripheralTimestampIsReconnectingError sets the handler for the CentralManagerDidDisconnectPeripheralTimestampIsReconnectingError delegate method.
func (d *CBCentralManagerDelegate) SetCentralManagerDidDisconnectPeripheralTimestampIsReconnectingError(f func(central ICBCentralManager, peripheral ICBPeripheral, timestamp AbsoluteTime /* not a class type */, isReconnecting bool, error_ objc.IObject /* cross-framework: Error */)) {
	d._CentralManagerDidDisconnectPeripheralTimestampIsReconnectingError = f
}

// SetCentralManagerDidDiscoverPeripheralAdvertisementDataRSSI sets the handler for the CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI delegate method.
//
// Tells the delegate the central manager discovered a peripheral while scanning for devices.
func (d *CBCentralManagerDelegate) SetCentralManagerDidDiscoverPeripheralAdvertisementDataRSSI(f func(central ICBCentralManager, peripheral ICBPeripheral, advertisementData foundation.IDictionary, RSSI objc.IObject /* cross-framework: NSNumber */)) {
	d._CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI = f
}

// SetCentralManagerDidFailToConnectPeripheralError sets the handler for the CentralManagerDidFailToConnectPeripheralError delegate method.
//
// Tells the delegate the central manager failed to create a connection with a peripheral.
func (d *CBCentralManagerDelegate) SetCentralManagerDidFailToConnectPeripheralError(f func(central ICBCentralManager, peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */)) {
	d._CentralManagerDidFailToConnectPeripheralError = f
}

// SetCentralManagerDidUpdateANCSAuthorizationForPeripheral sets the handler for the CentralManagerDidUpdateANCSAuthorizationForPeripheral delegate method.
//
// Tells the delegate the authorization status changed for a ANCS-requiring connected peripheral.
func (d *CBCentralManagerDelegate) SetCentralManagerDidUpdateANCSAuthorizationForPeripheral(f func(central ICBCentralManager, peripheral ICBPeripheral)) {
	d._CentralManagerDidUpdateANCSAuthorizationForPeripheral = f
}

// SetCentralManagerWillRestoreState sets the handler for the CentralManagerWillRestoreState delegate method.
//
// Tells the delegate the system is about to restore the central manager, as part of relaunching the app into the background.
func (d *CBCentralManagerDelegate) SetCentralManagerWillRestoreState(f func(central ICBCentralManager, dict foundation.IDictionary)) {
	d._CentralManagerWillRestoreState = f
}

// SetCentralManagerDidUpdateState sets the handler for the CentralManagerDidUpdateState delegate method.
//
// Tells the delegate the central manager’s state updated.
func (d *CBCentralManagerDelegate) SetCentralManagerDidUpdateState(f func(central ICBCentralManager)) {
	d._CentralManagerDidUpdateState = f
}

// CentralManagerConnectionEventDidOccurForPeripheral implements the PCBCentralManagerDelegate interface.
func (d *CBCentralManagerDelegate) CentralManagerConnectionEventDidOccurForPeripheral(central ICBCentralManager, event CBConnectionEvent, peripheral ICBPeripheral) {
	if d._CentralManagerConnectionEventDidOccurForPeripheral != nil {
		d._CentralManagerConnectionEventDidOccurForPeripheral(central, event, peripheral)
	}
}

// HasCentralManagerConnectionEventDidOccurForPeripheral returns true if a handler for CentralManagerConnectionEventDidOccurForPeripheral has been set.
func (d *CBCentralManagerDelegate) HasCentralManagerConnectionEventDidOccurForPeripheral() bool {
	return d._CentralManagerConnectionEventDidOccurForPeripheral != nil
}

// CentralManagerDidConnectPeripheral implements the PCBCentralManagerDelegate interface.
func (d *CBCentralManagerDelegate) CentralManagerDidConnectPeripheral(central ICBCentralManager, peripheral ICBPeripheral) {
	if d._CentralManagerDidConnectPeripheral != nil {
		d._CentralManagerDidConnectPeripheral(central, peripheral)
	}
}

// HasCentralManagerDidConnectPeripheral returns true if a handler for CentralManagerDidConnectPeripheral has been set.
func (d *CBCentralManagerDelegate) HasCentralManagerDidConnectPeripheral() bool {
	return d._CentralManagerDidConnectPeripheral != nil
}

// CentralManagerDidDisconnectPeripheralError implements the PCBCentralManagerDelegate interface.
func (d *CBCentralManagerDelegate) CentralManagerDidDisconnectPeripheralError(central ICBCentralManager, peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */) {
	if d._CentralManagerDidDisconnectPeripheralError != nil {
		d._CentralManagerDidDisconnectPeripheralError(central, peripheral, error_)
	}
}

// HasCentralManagerDidDisconnectPeripheralError returns true if a handler for CentralManagerDidDisconnectPeripheralError has been set.
func (d *CBCentralManagerDelegate) HasCentralManagerDidDisconnectPeripheralError() bool {
	return d._CentralManagerDidDisconnectPeripheralError != nil
}

// CentralManagerDidDisconnectPeripheralTimestampIsReconnectingError implements the PCBCentralManagerDelegate interface.
func (d *CBCentralManagerDelegate) CentralManagerDidDisconnectPeripheralTimestampIsReconnectingError(central ICBCentralManager, peripheral ICBPeripheral, timestamp AbsoluteTime /* not a class type */, isReconnecting bool, error_ objc.IObject /* cross-framework: Error */) {
	if d._CentralManagerDidDisconnectPeripheralTimestampIsReconnectingError != nil {
		d._CentralManagerDidDisconnectPeripheralTimestampIsReconnectingError(central, peripheral, timestamp, isReconnecting, error_)
	}
}

// HasCentralManagerDidDisconnectPeripheralTimestampIsReconnectingError returns true if a handler for CentralManagerDidDisconnectPeripheralTimestampIsReconnectingError has been set.
func (d *CBCentralManagerDelegate) HasCentralManagerDidDisconnectPeripheralTimestampIsReconnectingError() bool {
	return d._CentralManagerDidDisconnectPeripheralTimestampIsReconnectingError != nil
}

// CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI implements the PCBCentralManagerDelegate interface.
func (d *CBCentralManagerDelegate) CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI(central ICBCentralManager, peripheral ICBPeripheral, advertisementData foundation.IDictionary, RSSI objc.IObject /* cross-framework: NSNumber */) {
	if d._CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI != nil {
		d._CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI(central, peripheral, advertisementData, RSSI)
	}
}

// HasCentralManagerDidDiscoverPeripheralAdvertisementDataRSSI returns true if a handler for CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI has been set.
func (d *CBCentralManagerDelegate) HasCentralManagerDidDiscoverPeripheralAdvertisementDataRSSI() bool {
	return d._CentralManagerDidDiscoverPeripheralAdvertisementDataRSSI != nil
}

// CentralManagerDidFailToConnectPeripheralError implements the PCBCentralManagerDelegate interface.
func (d *CBCentralManagerDelegate) CentralManagerDidFailToConnectPeripheralError(central ICBCentralManager, peripheral ICBPeripheral, error_ objc.IObject /* cross-framework: Error */) {
	if d._CentralManagerDidFailToConnectPeripheralError != nil {
		d._CentralManagerDidFailToConnectPeripheralError(central, peripheral, error_)
	}
}

// HasCentralManagerDidFailToConnectPeripheralError returns true if a handler for CentralManagerDidFailToConnectPeripheralError has been set.
func (d *CBCentralManagerDelegate) HasCentralManagerDidFailToConnectPeripheralError() bool {
	return d._CentralManagerDidFailToConnectPeripheralError != nil
}

// CentralManagerDidUpdateANCSAuthorizationForPeripheral implements the PCBCentralManagerDelegate interface.
func (d *CBCentralManagerDelegate) CentralManagerDidUpdateANCSAuthorizationForPeripheral(central ICBCentralManager, peripheral ICBPeripheral) {
	if d._CentralManagerDidUpdateANCSAuthorizationForPeripheral != nil {
		d._CentralManagerDidUpdateANCSAuthorizationForPeripheral(central, peripheral)
	}
}

// HasCentralManagerDidUpdateANCSAuthorizationForPeripheral returns true if a handler for CentralManagerDidUpdateANCSAuthorizationForPeripheral has been set.
func (d *CBCentralManagerDelegate) HasCentralManagerDidUpdateANCSAuthorizationForPeripheral() bool {
	return d._CentralManagerDidUpdateANCSAuthorizationForPeripheral != nil
}

// CentralManagerWillRestoreState implements the PCBCentralManagerDelegate interface.
func (d *CBCentralManagerDelegate) CentralManagerWillRestoreState(central ICBCentralManager, dict foundation.IDictionary) {
	if d._CentralManagerWillRestoreState != nil {
		d._CentralManagerWillRestoreState(central, dict)
	}
}

// HasCentralManagerWillRestoreState returns true if a handler for CentralManagerWillRestoreState has been set.
func (d *CBCentralManagerDelegate) HasCentralManagerWillRestoreState() bool {
	return d._CentralManagerWillRestoreState != nil
}

// CentralManagerDidUpdateState implements the PCBCentralManagerDelegate interface.
func (d *CBCentralManagerDelegate) CentralManagerDidUpdateState(central ICBCentralManager) {
	if d._CentralManagerDidUpdateState != nil {
		d._CentralManagerDidUpdateState(central)
	}
}

// HasCentralManagerDidUpdateState returns true if a handler for CentralManagerDidUpdateState has been set.
func (d *CBCentralManagerDelegate) HasCentralManagerDidUpdateState() bool {
	return d._CentralManagerDidUpdateState != nil
}
