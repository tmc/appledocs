// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMIDICIProfileResponderDelegate is the MIDICIProfileResponderDelegate protocol interface.
//
// A protocol that defines the methods to respond to MIDI-CI responder life-cycle events.
//
// Availability:
//   - Mac Catalyst 14.0+ (Deprecated in 18.0)
//   - iOS 14.0+ (Deprecated in 18.0)
//   - iPadOS 14.0+ (Deprecated in 18.0)
//   - macOS 11.0+ (Deprecated in 15.0)
//   - visionOS 1.0+ (Deprecated in 2.0)
//
// See: doc://com.apple.coremidi/documentation/CoreMIDI/MIDICIProfileResponderDelegate
type PMIDICIProfileResponderDelegate interface {
	// Required methods
	ConnectInitiatorWithDeviceInfo(initiatorMUID MIDICIInitiatiorMUID /* typedef */, deviceInfo IMIDICIDeviceInfo) bool/* debug [protocol_interface/required_method]: ConnectInitiatorWithDeviceInfo */
	InitiatorDisconnected(initiatorMUID MIDICIInitiatiorMUID /* typedef */)/* debug [protocol_interface/required_method]: InitiatorDisconnected */
	// Optional methods
	HandleDataForProfileOnChannelData(aProfile IMIDICIProfile, channel MIDIChannelNumber /* typedef */, inData objc.IObject /* cross-framework: NSData */)
	HasHandleDataForProfileOnChannelData() bool
	WillSetProfileOnChannelEnabled(aProfile IMIDICIProfile, channel MIDIChannelNumber /* typedef */, shouldEnable bool) bool
	HasWillSetProfileOnChannelEnabled() bool
}

// MIDICIProfileResponderDelegate is a delegate implementation builder for the PMIDICIProfileResponderDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MIDICIProfileResponderDelegate struct {
	_HandleDataForProfileOnChannelData func(aProfile IMIDICIProfile, channel MIDIChannelNumber /* typedef */, inData objc.IObject /* cross-framework: NSData */)
	_WillSetProfileOnChannelEnabled func(aProfile IMIDICIProfile, channel MIDIChannelNumber /* typedef */, shouldEnable bool) bool
	_ConnectInitiatorWithDeviceInfo func(initiatorMUID MIDICIInitiatiorMUID /* typedef */, deviceInfo IMIDICIDeviceInfo) bool
	_InitiatorDisconnected func(initiatorMUID MIDICIInitiatiorMUID /* typedef */)
}

// SetHandleDataForProfileOnChannelData sets the handler for the HandleDataForProfileOnChannelData delegate method.
//
// Processes MIDI data for a profile and channel.
func (d *MIDICIProfileResponderDelegate) SetHandleDataForProfileOnChannelData(f func(aProfile IMIDICIProfile, channel MIDIChannelNumber /* typedef */, inData objc.IObject /* cross-framework: NSData */)) {
	d._HandleDataForProfileOnChannelData = f
}

// SetWillSetProfileOnChannelEnabled sets the handler for the WillSetProfileOnChannelEnabled delegate method.
//
// Provides an opportunity to perform an action before the system sets the profile.
func (d *MIDICIProfileResponderDelegate) SetWillSetProfileOnChannelEnabled(f func(aProfile IMIDICIProfile, channel MIDIChannelNumber /* typedef */, shouldEnable bool) bool) {
	d._WillSetProfileOnChannelEnabled = f
}

// SetConnectInitiatorWithDeviceInfo sets the handler for the ConnectInitiatorWithDeviceInfo delegate method.
//
// Enables a MIDI-CI initiator to create a session or reject the connection attempt.
func (d *MIDICIProfileResponderDelegate) SetConnectInitiatorWithDeviceInfo(f func(initiatorMUID MIDICIInitiatiorMUID /* typedef */, deviceInfo IMIDICIDeviceInfo) bool) {
	d._ConnectInitiatorWithDeviceInfo = f
}

// SetInitiatorDisconnected sets the handler for the InitiatorDisconnected delegate method.
//
// Provides an opportunity to perform an action after the system disconnects the initiator.
func (d *MIDICIProfileResponderDelegate) SetInitiatorDisconnected(f func(initiatorMUID MIDICIInitiatiorMUID /* typedef */)) {
	d._InitiatorDisconnected = f
}

// HandleDataForProfileOnChannelData implements the PMIDICIProfileResponderDelegate interface.
func (d *MIDICIProfileResponderDelegate) HandleDataForProfileOnChannelData(aProfile IMIDICIProfile, channel MIDIChannelNumber /* typedef */, inData objc.IObject /* cross-framework: NSData */) {
	if d._HandleDataForProfileOnChannelData != nil {
		d._HandleDataForProfileOnChannelData(aProfile, channel, inData)
	}
}

// HasHandleDataForProfileOnChannelData returns true if a handler for HandleDataForProfileOnChannelData has been set.
func (d *MIDICIProfileResponderDelegate) HasHandleDataForProfileOnChannelData() bool {
	return d._HandleDataForProfileOnChannelData != nil
}

// WillSetProfileOnChannelEnabled implements the PMIDICIProfileResponderDelegate interface.
func (d *MIDICIProfileResponderDelegate) WillSetProfileOnChannelEnabled(aProfile IMIDICIProfile, channel MIDIChannelNumber /* typedef */, shouldEnable bool) bool {
	if d._WillSetProfileOnChannelEnabled != nil {
		return d._WillSetProfileOnChannelEnabled(aProfile, channel, shouldEnable)
	}
	var zero bool
	return zero
}

// HasWillSetProfileOnChannelEnabled returns true if a handler for WillSetProfileOnChannelEnabled has been set.
func (d *MIDICIProfileResponderDelegate) HasWillSetProfileOnChannelEnabled() bool {
	return d._WillSetProfileOnChannelEnabled != nil
}

// ConnectInitiatorWithDeviceInfo implements the PMIDICIProfileResponderDelegate interface.
func (d *MIDICIProfileResponderDelegate) ConnectInitiatorWithDeviceInfo(initiatorMUID MIDICIInitiatiorMUID /* typedef */, deviceInfo IMIDICIDeviceInfo) bool {
	if d._ConnectInitiatorWithDeviceInfo != nil {
		return d._ConnectInitiatorWithDeviceInfo(initiatorMUID, deviceInfo)
	}
	var zero bool
	return zero
}

// HasConnectInitiatorWithDeviceInfo returns true if a handler for ConnectInitiatorWithDeviceInfo has been set.
func (d *MIDICIProfileResponderDelegate) HasConnectInitiatorWithDeviceInfo() bool {
	return d._ConnectInitiatorWithDeviceInfo != nil
}

// InitiatorDisconnected implements the PMIDICIProfileResponderDelegate interface.
func (d *MIDICIProfileResponderDelegate) InitiatorDisconnected(initiatorMUID MIDICIInitiatiorMUID /* typedef */) {
	if d._InitiatorDisconnected != nil {
		d._InitiatorDisconnected(initiatorMUID)
	}
}

// HasInitiatorDisconnected returns true if a handler for InitiatorDisconnected has been set.
func (d *MIDICIProfileResponderDelegate) HasInitiatorDisconnected() bool {
	return d._InitiatorDisconnected != nil
}
