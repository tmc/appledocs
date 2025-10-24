// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PBluetoothHandsFreeDeviceDelegate is the IOBluetoothHandsFreeDeviceDelegate protocol interface.
//
// A set of optional methods for receiving status change updates and information about a connected Bluetooth hands-free phone or headset.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.iobluetooth/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate
type PBluetoothHandsFreeDeviceDelegate interface {
	// Optional methods
	HandsFreeBatteryCharge(device IOBluetoothHandsFreeDevice, batteryCharge objc.IObject /* cross-framework: NSNumber */)
	HasHandsFreeBatteryCharge() bool
	HandsFreeCallHoldState(device IOBluetoothHandsFreeDevice, callHoldState objc.IObject /* cross-framework: NSNumber */)
	HasHandsFreeCallHoldState() bool
	HandsFreeCallSetupMode(device IOBluetoothHandsFreeDevice, callSetupMode objc.IObject /* cross-framework: NSNumber */)
	HasHandsFreeCallSetupMode() bool
	HandsFreeCurrentCall(device IOBluetoothHandsFreeDevice, currentCall objc.IObject /* cross-framework: NSDictionary */)
	HasHandsFreeCurrentCall() bool
	HandsFreeIncomingCallFrom(device IOBluetoothHandsFreeDevice, number objc.IObject /* cross-framework: NSString */)
	HasHandsFreeIncomingCallFrom() bool
	HandsFreeIncomingSMS(device IOBluetoothHandsFreeDevice, sms objc.IObject /* cross-framework: NSDictionary */)
	HasHandsFreeIncomingSMS() bool
	HandsFreeIsCallActive(device IOBluetoothHandsFreeDevice, isCallActive objc.IObject /* cross-framework: NSNumber */)
	HasHandsFreeIsCallActive() bool
	HandsFreeIsRoaming(device IOBluetoothHandsFreeDevice, isRoaming objc.IObject /* cross-framework: NSNumber */)
	HasHandsFreeIsRoaming() bool
	HandsFreeIsServiceAvailable(device IOBluetoothHandsFreeDevice, isServiceAvailable objc.IObject /* cross-framework: NSNumber */)
	HasHandsFreeIsServiceAvailable() bool
	HandsFreeRingAttempt(device IOBluetoothHandsFreeDevice, ringAttempt objc.IObject /* cross-framework: NSNumber */)
	HasHandsFreeRingAttempt() bool
	HandsFreeSignalStrength(device IOBluetoothHandsFreeDevice, signalStrength objc.IObject /* cross-framework: NSNumber */)
	HasHandsFreeSignalStrength() bool
	HandsFreeSubscriberNumber(device IOBluetoothHandsFreeDevice, subscriberNumber objc.IObject /* cross-framework: NSString */)
	HasHandsFreeSubscriberNumber() bool
	HandsFreeUnhandledResultCode(device IOBluetoothHandsFreeDevice, resultCode objc.IObject /* cross-framework: NSString */)
	HasHandsFreeUnhandledResultCode() bool
}

// BluetoothHandsFreeDeviceDelegate is a delegate implementation builder for the PBluetoothHandsFreeDeviceDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type BluetoothHandsFreeDeviceDelegate struct {
	_HandsFreeBatteryCharge func(device IOBluetoothHandsFreeDevice, batteryCharge objc.IObject /* cross-framework: NSNumber */)
	_HandsFreeCallHoldState func(device IOBluetoothHandsFreeDevice, callHoldState objc.IObject /* cross-framework: NSNumber */)
	_HandsFreeCallSetupMode func(device IOBluetoothHandsFreeDevice, callSetupMode objc.IObject /* cross-framework: NSNumber */)
	_HandsFreeCurrentCall func(device IOBluetoothHandsFreeDevice, currentCall objc.IObject /* cross-framework: NSDictionary */)
	_HandsFreeIncomingCallFrom func(device IOBluetoothHandsFreeDevice, number objc.IObject /* cross-framework: NSString */)
	_HandsFreeIncomingSMS func(device IOBluetoothHandsFreeDevice, sms objc.IObject /* cross-framework: NSDictionary */)
	_HandsFreeIsCallActive func(device IOBluetoothHandsFreeDevice, isCallActive objc.IObject /* cross-framework: NSNumber */)
	_HandsFreeIsRoaming func(device IOBluetoothHandsFreeDevice, isRoaming objc.IObject /* cross-framework: NSNumber */)
	_HandsFreeIsServiceAvailable func(device IOBluetoothHandsFreeDevice, isServiceAvailable objc.IObject /* cross-framework: NSNumber */)
	_HandsFreeRingAttempt func(device IOBluetoothHandsFreeDevice, ringAttempt objc.IObject /* cross-framework: NSNumber */)
	_HandsFreeSignalStrength func(device IOBluetoothHandsFreeDevice, signalStrength objc.IObject /* cross-framework: NSNumber */)
	_HandsFreeSubscriberNumber func(device IOBluetoothHandsFreeDevice, subscriberNumber objc.IObject /* cross-framework: NSString */)
	_HandsFreeUnhandledResultCode func(device IOBluetoothHandsFreeDevice, resultCode objc.IObject /* cross-framework: NSString */)
}

// SetHandsFreeBatteryCharge sets the handler for the HandsFreeBatteryCharge delegate method.
//
// Tells the delegate the battery level indicator of the connected Bluetooth hands-free phone or headset has changed.
func (d *BluetoothHandsFreeDeviceDelegate) SetHandsFreeBatteryCharge(f func(device IOBluetoothHandsFreeDevice, batteryCharge objc.IObject /* cross-framework: NSNumber */)) {
	d._HandsFreeBatteryCharge = f
}

// SetHandsFreeCallHoldState sets the handler for the HandsFreeCallHoldState delegate method.
//
// Tells the delegate the call held indicator of the connected Bluetooth hands-free phone or headset has changed.
func (d *BluetoothHandsFreeDeviceDelegate) SetHandsFreeCallHoldState(f func(device IOBluetoothHandsFreeDevice, callHoldState objc.IObject /* cross-framework: NSNumber */)) {
	d._HandsFreeCallHoldState = f
}

// SetHandsFreeCallSetupMode sets the handler for the HandsFreeCallSetupMode delegate method.
//
// Tells the delegate the call setup indicator of the connected Bluetooth hands-free phone or headset has changed.
func (d *BluetoothHandsFreeDeviceDelegate) SetHandsFreeCallSetupMode(f func(device IOBluetoothHandsFreeDevice, callSetupMode objc.IObject /* cross-framework: NSNumber */)) {
	d._HandsFreeCallSetupMode = f
}

// SetHandsFreeCurrentCall sets the handler for the HandsFreeCurrentCall delegate method.
//
// Sends the delegate information about the current call.
func (d *BluetoothHandsFreeDeviceDelegate) SetHandsFreeCurrentCall(f func(device IOBluetoothHandsFreeDevice, currentCall objc.IObject /* cross-framework: NSDictionary */)) {
	d._HandsFreeCurrentCall = f
}

// SetHandsFreeIncomingCallFrom sets the handler for the HandsFreeIncomingCallFrom delegate method.
//
// Tells the delegate there’s an incoming call on the connected Bluetooth hands-free phone or headset.
func (d *BluetoothHandsFreeDeviceDelegate) SetHandsFreeIncomingCallFrom(f func(device IOBluetoothHandsFreeDevice, number objc.IObject /* cross-framework: NSString */)) {
	d._HandsFreeIncomingCallFrom = f
}

// SetHandsFreeIncomingSMS sets the handler for the HandsFreeIncomingSMS delegate method.
//
// Tells the delegate there’s an incoming text message.
func (d *BluetoothHandsFreeDeviceDelegate) SetHandsFreeIncomingSMS(f func(device IOBluetoothHandsFreeDevice, sms objc.IObject /* cross-framework: NSDictionary */)) {
	d._HandsFreeIncomingSMS = f
}

// SetHandsFreeIsCallActive sets the handler for the HandsFreeIsCallActive delegate method.
//
// Tells the delegate the active call indicator of the connected Bluetooth hands-free phone or headset has changed.
func (d *BluetoothHandsFreeDeviceDelegate) SetHandsFreeIsCallActive(f func(device IOBluetoothHandsFreeDevice, isCallActive objc.IObject /* cross-framework: NSNumber */)) {
	d._HandsFreeIsCallActive = f
}

// SetHandsFreeIsRoaming sets the handler for the HandsFreeIsRoaming delegate method.
//
// Tells the delegate the roaming indicator of the connected Bluetooth hands-free phone or headset has changed.
func (d *BluetoothHandsFreeDeviceDelegate) SetHandsFreeIsRoaming(f func(device IOBluetoothHandsFreeDevice, isRoaming objc.IObject /* cross-framework: NSNumber */)) {
	d._HandsFreeIsRoaming = f
}

// SetHandsFreeIsServiceAvailable sets the handler for the HandsFreeIsServiceAvailable delegate method.
//
// Tells the delegate the service level indicator of the connected Bluetooth hands-free phone or headset has changed.
func (d *BluetoothHandsFreeDeviceDelegate) SetHandsFreeIsServiceAvailable(f func(device IOBluetoothHandsFreeDevice, isServiceAvailable objc.IObject /* cross-framework: NSNumber */)) {
	d._HandsFreeIsServiceAvailable = f
}

// SetHandsFreeRingAttempt sets the handler for the HandsFreeRingAttempt delegate method.
//
// Tells the delegate the phone is ringing.
func (d *BluetoothHandsFreeDeviceDelegate) SetHandsFreeRingAttempt(f func(device IOBluetoothHandsFreeDevice, ringAttempt objc.IObject /* cross-framework: NSNumber */)) {
	d._HandsFreeRingAttempt = f
}

// SetHandsFreeSignalStrength sets the handler for the HandsFreeSignalStrength delegate method.
//
// Tells the delegate the call setup signal strength indicator of the connected Bluetooth hands-free phone or headset has changed.
func (d *BluetoothHandsFreeDeviceDelegate) SetHandsFreeSignalStrength(f func(device IOBluetoothHandsFreeDevice, signalStrength objc.IObject /* cross-framework: NSNumber */)) {
	d._HandsFreeSignalStrength = f
}

// SetHandsFreeSubscriberNumber sets the handler for the HandsFreeSubscriberNumber delegate method.
//
// Tells the delegate the subscriber number of a call.
func (d *BluetoothHandsFreeDeviceDelegate) SetHandsFreeSubscriberNumber(f func(device IOBluetoothHandsFreeDevice, subscriberNumber objc.IObject /* cross-framework: NSString */)) {
	d._HandsFreeSubscriberNumber = f
}

// SetHandsFreeUnhandledResultCode sets the handler for the HandsFreeUnhandledResultCode delegate method.
//
// Tells the delegate the phone sent an unknown code.
func (d *BluetoothHandsFreeDeviceDelegate) SetHandsFreeUnhandledResultCode(f func(device IOBluetoothHandsFreeDevice, resultCode objc.IObject /* cross-framework: NSString */)) {
	d._HandsFreeUnhandledResultCode = f
}

// HandsFreeBatteryCharge implements the PBluetoothHandsFreeDeviceDelegate interface.
func (d *BluetoothHandsFreeDeviceDelegate) HandsFreeBatteryCharge(device IOBluetoothHandsFreeDevice, batteryCharge objc.IObject /* cross-framework: NSNumber */) {
	if d._HandsFreeBatteryCharge != nil {
		d._HandsFreeBatteryCharge(device, batteryCharge)
	}
}

// HasHandsFreeBatteryCharge returns true if a handler for HandsFreeBatteryCharge has been set.
func (d *BluetoothHandsFreeDeviceDelegate) HasHandsFreeBatteryCharge() bool {
	return d._HandsFreeBatteryCharge != nil
}

// HandsFreeCallHoldState implements the PBluetoothHandsFreeDeviceDelegate interface.
func (d *BluetoothHandsFreeDeviceDelegate) HandsFreeCallHoldState(device IOBluetoothHandsFreeDevice, callHoldState objc.IObject /* cross-framework: NSNumber */) {
	if d._HandsFreeCallHoldState != nil {
		d._HandsFreeCallHoldState(device, callHoldState)
	}
}

// HasHandsFreeCallHoldState returns true if a handler for HandsFreeCallHoldState has been set.
func (d *BluetoothHandsFreeDeviceDelegate) HasHandsFreeCallHoldState() bool {
	return d._HandsFreeCallHoldState != nil
}

// HandsFreeCallSetupMode implements the PBluetoothHandsFreeDeviceDelegate interface.
func (d *BluetoothHandsFreeDeviceDelegate) HandsFreeCallSetupMode(device IOBluetoothHandsFreeDevice, callSetupMode objc.IObject /* cross-framework: NSNumber */) {
	if d._HandsFreeCallSetupMode != nil {
		d._HandsFreeCallSetupMode(device, callSetupMode)
	}
}

// HasHandsFreeCallSetupMode returns true if a handler for HandsFreeCallSetupMode has been set.
func (d *BluetoothHandsFreeDeviceDelegate) HasHandsFreeCallSetupMode() bool {
	return d._HandsFreeCallSetupMode != nil
}

// HandsFreeCurrentCall implements the PBluetoothHandsFreeDeviceDelegate interface.
func (d *BluetoothHandsFreeDeviceDelegate) HandsFreeCurrentCall(device IOBluetoothHandsFreeDevice, currentCall objc.IObject /* cross-framework: NSDictionary */) {
	if d._HandsFreeCurrentCall != nil {
		d._HandsFreeCurrentCall(device, currentCall)
	}
}

// HasHandsFreeCurrentCall returns true if a handler for HandsFreeCurrentCall has been set.
func (d *BluetoothHandsFreeDeviceDelegate) HasHandsFreeCurrentCall() bool {
	return d._HandsFreeCurrentCall != nil
}

// HandsFreeIncomingCallFrom implements the PBluetoothHandsFreeDeviceDelegate interface.
func (d *BluetoothHandsFreeDeviceDelegate) HandsFreeIncomingCallFrom(device IOBluetoothHandsFreeDevice, number objc.IObject /* cross-framework: NSString */) {
	if d._HandsFreeIncomingCallFrom != nil {
		d._HandsFreeIncomingCallFrom(device, number)
	}
}

// HasHandsFreeIncomingCallFrom returns true if a handler for HandsFreeIncomingCallFrom has been set.
func (d *BluetoothHandsFreeDeviceDelegate) HasHandsFreeIncomingCallFrom() bool {
	return d._HandsFreeIncomingCallFrom != nil
}

// HandsFreeIncomingSMS implements the PBluetoothHandsFreeDeviceDelegate interface.
func (d *BluetoothHandsFreeDeviceDelegate) HandsFreeIncomingSMS(device IOBluetoothHandsFreeDevice, sms objc.IObject /* cross-framework: NSDictionary */) {
	if d._HandsFreeIncomingSMS != nil {
		d._HandsFreeIncomingSMS(device, sms)
	}
}

// HasHandsFreeIncomingSMS returns true if a handler for HandsFreeIncomingSMS has been set.
func (d *BluetoothHandsFreeDeviceDelegate) HasHandsFreeIncomingSMS() bool {
	return d._HandsFreeIncomingSMS != nil
}

// HandsFreeIsCallActive implements the PBluetoothHandsFreeDeviceDelegate interface.
func (d *BluetoothHandsFreeDeviceDelegate) HandsFreeIsCallActive(device IOBluetoothHandsFreeDevice, isCallActive objc.IObject /* cross-framework: NSNumber */) {
	if d._HandsFreeIsCallActive != nil {
		d._HandsFreeIsCallActive(device, isCallActive)
	}
}

// HasHandsFreeIsCallActive returns true if a handler for HandsFreeIsCallActive has been set.
func (d *BluetoothHandsFreeDeviceDelegate) HasHandsFreeIsCallActive() bool {
	return d._HandsFreeIsCallActive != nil
}

// HandsFreeIsRoaming implements the PBluetoothHandsFreeDeviceDelegate interface.
func (d *BluetoothHandsFreeDeviceDelegate) HandsFreeIsRoaming(device IOBluetoothHandsFreeDevice, isRoaming objc.IObject /* cross-framework: NSNumber */) {
	if d._HandsFreeIsRoaming != nil {
		d._HandsFreeIsRoaming(device, isRoaming)
	}
}

// HasHandsFreeIsRoaming returns true if a handler for HandsFreeIsRoaming has been set.
func (d *BluetoothHandsFreeDeviceDelegate) HasHandsFreeIsRoaming() bool {
	return d._HandsFreeIsRoaming != nil
}

// HandsFreeIsServiceAvailable implements the PBluetoothHandsFreeDeviceDelegate interface.
func (d *BluetoothHandsFreeDeviceDelegate) HandsFreeIsServiceAvailable(device IOBluetoothHandsFreeDevice, isServiceAvailable objc.IObject /* cross-framework: NSNumber */) {
	if d._HandsFreeIsServiceAvailable != nil {
		d._HandsFreeIsServiceAvailable(device, isServiceAvailable)
	}
}

// HasHandsFreeIsServiceAvailable returns true if a handler for HandsFreeIsServiceAvailable has been set.
func (d *BluetoothHandsFreeDeviceDelegate) HasHandsFreeIsServiceAvailable() bool {
	return d._HandsFreeIsServiceAvailable != nil
}

// HandsFreeRingAttempt implements the PBluetoothHandsFreeDeviceDelegate interface.
func (d *BluetoothHandsFreeDeviceDelegate) HandsFreeRingAttempt(device IOBluetoothHandsFreeDevice, ringAttempt objc.IObject /* cross-framework: NSNumber */) {
	if d._HandsFreeRingAttempt != nil {
		d._HandsFreeRingAttempt(device, ringAttempt)
	}
}

// HasHandsFreeRingAttempt returns true if a handler for HandsFreeRingAttempt has been set.
func (d *BluetoothHandsFreeDeviceDelegate) HasHandsFreeRingAttempt() bool {
	return d._HandsFreeRingAttempt != nil
}

// HandsFreeSignalStrength implements the PBluetoothHandsFreeDeviceDelegate interface.
func (d *BluetoothHandsFreeDeviceDelegate) HandsFreeSignalStrength(device IOBluetoothHandsFreeDevice, signalStrength objc.IObject /* cross-framework: NSNumber */) {
	if d._HandsFreeSignalStrength != nil {
		d._HandsFreeSignalStrength(device, signalStrength)
	}
}

// HasHandsFreeSignalStrength returns true if a handler for HandsFreeSignalStrength has been set.
func (d *BluetoothHandsFreeDeviceDelegate) HasHandsFreeSignalStrength() bool {
	return d._HandsFreeSignalStrength != nil
}

// HandsFreeSubscriberNumber implements the PBluetoothHandsFreeDeviceDelegate interface.
func (d *BluetoothHandsFreeDeviceDelegate) HandsFreeSubscriberNumber(device IOBluetoothHandsFreeDevice, subscriberNumber objc.IObject /* cross-framework: NSString */) {
	if d._HandsFreeSubscriberNumber != nil {
		d._HandsFreeSubscriberNumber(device, subscriberNumber)
	}
}

// HasHandsFreeSubscriberNumber returns true if a handler for HandsFreeSubscriberNumber has been set.
func (d *BluetoothHandsFreeDeviceDelegate) HasHandsFreeSubscriberNumber() bool {
	return d._HandsFreeSubscriberNumber != nil
}

// HandsFreeUnhandledResultCode implements the PBluetoothHandsFreeDeviceDelegate interface.
func (d *BluetoothHandsFreeDeviceDelegate) HandsFreeUnhandledResultCode(device IOBluetoothHandsFreeDevice, resultCode objc.IObject /* cross-framework: NSString */) {
	if d._HandsFreeUnhandledResultCode != nil {
		d._HandsFreeUnhandledResultCode(device, resultCode)
	}
}

// HasHandsFreeUnhandledResultCode returns true if a handler for HandsFreeUnhandledResultCode has been set.
func (d *BluetoothHandsFreeDeviceDelegate) HasHandsFreeUnhandledResultCode() bool {
	return d._HandsFreeUnhandledResultCode != nil
}
