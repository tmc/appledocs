// Code generated from Apple documentation for SafetyKit. DO NOT EDIT.

package safetykit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PSACrashDetectionDelegate is the SACrashDetectionDelegate protocol interface.
//
// The protocol that an object adopts to receive Crash Detection events and changes to the authorization status.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - watchOS 10.1+
//
// See: doc://com.apple.safetykit/documentation/SafetyKit/SACrashDetectionDelegate
type PSACrashDetectionDelegate interface {
	// Optional methods
	CrashDetectionManagerDidDetectEvent(crashDetectionManager ISACrashDetectionManager, event ISACrashDetectionEvent)
	HasCrashDetectionManagerDidDetectEvent() bool
}

// SACrashDetectionDelegate is a delegate implementation builder for the PSACrashDetectionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SACrashDetectionDelegate struct {
	_CrashDetectionManagerDidDetectEvent func(crashDetectionManager ISACrashDetectionManager, event ISACrashDetectionEvent)
}

// SetCrashDetectionManagerDidDetectEvent sets the handler for the CrashDetectionManagerDidDetectEvent delegate method.
//
// Receive and process a Crash Detection event.
func (d *SACrashDetectionDelegate) SetCrashDetectionManagerDidDetectEvent(f func(crashDetectionManager ISACrashDetectionManager, event ISACrashDetectionEvent)) {
	d._CrashDetectionManagerDidDetectEvent = f
}

// CrashDetectionManagerDidDetectEvent implements the PSACrashDetectionDelegate interface.
func (d *SACrashDetectionDelegate) CrashDetectionManagerDidDetectEvent(crashDetectionManager ISACrashDetectionManager, event ISACrashDetectionEvent) {
	if d._CrashDetectionManagerDidDetectEvent != nil {
		d._CrashDetectionManagerDidDetectEvent(crashDetectionManager, event)
	}
}

// HasCrashDetectionManagerDidDetectEvent returns true if a handler for CrashDetectionManagerDidDetectEvent has been set.
func (d *SACrashDetectionDelegate) HasCrashDetectionManagerDidDetectEvent() bool {
	return d._CrashDetectionManagerDidDetectEvent != nil
}
