// Code generated from Apple documentation for SafetyKit. DO NOT EDIT.

package safetykit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PSAEmergencyResponseDelegate is the SAEmergencyResponseDelegate protocol interface.
//
// The interface for receiving updates about a requested emergency response action.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - watchOS 10.1+
//
// See: doc://com.apple.safetykit/documentation/SafetyKit/SAEmergencyResponseDelegate
type PSAEmergencyResponseDelegate interface {
	// Optional methods
	EmergencyResponseManagerDidUpdateVoiceCallStatus(emergencyResponseManager ISAEmergencyResponseManager, voiceCallStatus SAEmergencyResponseManagerVoiceCallStatus)
	HasEmergencyResponseManagerDidUpdateVoiceCallStatus() bool
}

// SAEmergencyResponseDelegate is a delegate implementation builder for the PSAEmergencyResponseDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SAEmergencyResponseDelegate struct {
	_EmergencyResponseManagerDidUpdateVoiceCallStatus func(emergencyResponseManager ISAEmergencyResponseManager, voiceCallStatus SAEmergencyResponseManagerVoiceCallStatus)
}

// SetEmergencyResponseManagerDidUpdateVoiceCallStatus sets the handler for the EmergencyResponseManagerDidUpdateVoiceCallStatus delegate method.
//
// Provides the voice call status to the delegate.
func (d *SAEmergencyResponseDelegate) SetEmergencyResponseManagerDidUpdateVoiceCallStatus(f func(emergencyResponseManager ISAEmergencyResponseManager, voiceCallStatus SAEmergencyResponseManagerVoiceCallStatus)) {
	d._EmergencyResponseManagerDidUpdateVoiceCallStatus = f
}

// EmergencyResponseManagerDidUpdateVoiceCallStatus implements the PSAEmergencyResponseDelegate interface.
func (d *SAEmergencyResponseDelegate) EmergencyResponseManagerDidUpdateVoiceCallStatus(emergencyResponseManager ISAEmergencyResponseManager, voiceCallStatus SAEmergencyResponseManagerVoiceCallStatus) {
	if d._EmergencyResponseManagerDidUpdateVoiceCallStatus != nil {
		d._EmergencyResponseManagerDidUpdateVoiceCallStatus(emergencyResponseManager, voiceCallStatus)
	}
}

// HasEmergencyResponseManagerDidUpdateVoiceCallStatus returns true if a handler for EmergencyResponseManagerDidUpdateVoiceCallStatus has been set.
func (d *SAEmergencyResponseDelegate) HasEmergencyResponseManagerDidUpdateVoiceCallStatus() bool {
	return d._EmergencyResponseManagerDidUpdateVoiceCallStatus != nil
}
