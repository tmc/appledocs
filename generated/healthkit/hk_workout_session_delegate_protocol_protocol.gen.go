// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PHKWorkoutSessionDelegate is the HKWorkoutSessionDelegate protocol interface.
//
// The session delegate protocol that defines an interface for receiving notifications about errors and changes in the workout session’s state.
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS +
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.healthkit/documentation/HealthKit/HKWorkoutSessionDelegate
type PHKWorkoutSessionDelegate interface {
	// Required methods
	WorkoutSessionDidChangeToStateFromStateDate(workoutSession IHKWorkoutSession, toState HKWorkoutSessionState, fromState HKWorkoutSessionState, date objc.IObject /* cross-framework: NSDate */)/* debug [protocol_interface/required_method]: WorkoutSessionDidChangeToStateFromStateDate */
	WorkoutSessionDidFailWithError(workoutSession IHKWorkoutSession, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: WorkoutSessionDidFailWithError */
	// Optional methods
	WorkoutSessionDidBeginActivityWithConfigurationDate(workoutSession IHKWorkoutSession, workoutConfiguration IHKWorkoutConfiguration, date objc.IObject /* cross-framework: NSDate */)
	HasWorkoutSessionDidBeginActivityWithConfigurationDate() bool
	WorkoutSessionDidDisconnectFromRemoteDeviceWithError(workoutSession IHKWorkoutSession, error_ objc.IObject /* cross-framework: Error */)
	HasWorkoutSessionDidDisconnectFromRemoteDeviceWithError() bool
	WorkoutSessionDidEndActivityWithConfigurationDate(workoutSession IHKWorkoutSession, workoutConfiguration IHKWorkoutConfiguration, date objc.IObject /* cross-framework: NSDate */)
	HasWorkoutSessionDidEndActivityWithConfigurationDate() bool
	WorkoutSessionDidGenerateEvent(workoutSession IHKWorkoutSession, event IHKWorkoutEvent)
	HasWorkoutSessionDidGenerateEvent() bool
	WorkoutSessionDidReceiveDataFromRemoteWorkoutSession(workoutSession IHKWorkoutSession, data []foundation.Data)
	HasWorkoutSessionDidReceiveDataFromRemoteWorkoutSession() bool
}

// HKWorkoutSessionDelegate is a delegate implementation builder for the PHKWorkoutSessionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type HKWorkoutSessionDelegate struct {
	_WorkoutSessionDidBeginActivityWithConfigurationDate func(workoutSession IHKWorkoutSession, workoutConfiguration IHKWorkoutConfiguration, date objc.IObject /* cross-framework: NSDate */)
	_WorkoutSessionDidDisconnectFromRemoteDeviceWithError func(workoutSession IHKWorkoutSession, error_ objc.IObject /* cross-framework: Error */)
	_WorkoutSessionDidEndActivityWithConfigurationDate func(workoutSession IHKWorkoutSession, workoutConfiguration IHKWorkoutConfiguration, date objc.IObject /* cross-framework: NSDate */)
	_WorkoutSessionDidGenerateEvent func(workoutSession IHKWorkoutSession, event IHKWorkoutEvent)
	_WorkoutSessionDidReceiveDataFromRemoteWorkoutSession func(workoutSession IHKWorkoutSession, data []foundation.Data)
	_WorkoutSessionDidChangeToStateFromStateDate func(workoutSession IHKWorkoutSession, toState HKWorkoutSessionState, fromState HKWorkoutSessionState, date objc.IObject /* cross-framework: NSDate */)
	_WorkoutSessionDidFailWithError func(workoutSession IHKWorkoutSession, error_ objc.IObject /* cross-framework: Error */)
}

// SetWorkoutSessionDidBeginActivityWithConfigurationDate sets the handler for the WorkoutSessionDidBeginActivityWithConfigurationDate delegate method.
//
// Tells the delegate that a new workout session began.
func (d *HKWorkoutSessionDelegate) SetWorkoutSessionDidBeginActivityWithConfigurationDate(f func(workoutSession IHKWorkoutSession, workoutConfiguration IHKWorkoutConfiguration, date objc.IObject /* cross-framework: NSDate */)) {
	d._WorkoutSessionDidBeginActivityWithConfigurationDate = f
}

// SetWorkoutSessionDidDisconnectFromRemoteDeviceWithError sets the handler for the WorkoutSessionDidDisconnectFromRemoteDeviceWithError delegate method.
//
// Tells the delegate that the mirrored workout session disconnected from the primary session.
func (d *HKWorkoutSessionDelegate) SetWorkoutSessionDidDisconnectFromRemoteDeviceWithError(f func(workoutSession IHKWorkoutSession, error_ objc.IObject /* cross-framework: Error */)) {
	d._WorkoutSessionDidDisconnectFromRemoteDeviceWithError = f
}

// SetWorkoutSessionDidEndActivityWithConfigurationDate sets the handler for the WorkoutSessionDidEndActivityWithConfigurationDate delegate method.
//
// Tells the session that the current workout activity ended.
func (d *HKWorkoutSessionDelegate) SetWorkoutSessionDidEndActivityWithConfigurationDate(f func(workoutSession IHKWorkoutSession, workoutConfiguration IHKWorkoutConfiguration, date objc.IObject /* cross-framework: NSDate */)) {
	d._WorkoutSessionDidEndActivityWithConfigurationDate = f
}

// SetWorkoutSessionDidGenerateEvent sets the handler for the WorkoutSessionDidGenerateEvent delegate method.
//
// Tells the delegate that the system generated a workout event.
func (d *HKWorkoutSessionDelegate) SetWorkoutSessionDidGenerateEvent(f func(workoutSession IHKWorkoutSession, event IHKWorkoutEvent)) {
	d._WorkoutSessionDidGenerateEvent = f
}

// SetWorkoutSessionDidReceiveDataFromRemoteWorkoutSession sets the handler for the WorkoutSessionDidReceiveDataFromRemoteWorkoutSession delegate method.
//
// Passes data from the remote workout session to the session delegate.
func (d *HKWorkoutSessionDelegate) SetWorkoutSessionDidReceiveDataFromRemoteWorkoutSession(f func(workoutSession IHKWorkoutSession, data []foundation.Data)) {
	d._WorkoutSessionDidReceiveDataFromRemoteWorkoutSession = f
}

// SetWorkoutSessionDidChangeToStateFromStateDate sets the handler for the WorkoutSessionDidChangeToStateFromStateDate delegate method.
//
// Tells the delegate that the session’s state changed.
func (d *HKWorkoutSessionDelegate) SetWorkoutSessionDidChangeToStateFromStateDate(f func(workoutSession IHKWorkoutSession, toState HKWorkoutSessionState, fromState HKWorkoutSessionState, date objc.IObject /* cross-framework: NSDate */)) {
	d._WorkoutSessionDidChangeToStateFromStateDate = f
}

// SetWorkoutSessionDidFailWithError sets the handler for the WorkoutSessionDidFailWithError delegate method.
//
// Tells the delegate that the session failed with an error.
func (d *HKWorkoutSessionDelegate) SetWorkoutSessionDidFailWithError(f func(workoutSession IHKWorkoutSession, error_ objc.IObject /* cross-framework: Error */)) {
	d._WorkoutSessionDidFailWithError = f
}

// WorkoutSessionDidBeginActivityWithConfigurationDate implements the PHKWorkoutSessionDelegate interface.
func (d *HKWorkoutSessionDelegate) WorkoutSessionDidBeginActivityWithConfigurationDate(workoutSession IHKWorkoutSession, workoutConfiguration IHKWorkoutConfiguration, date objc.IObject /* cross-framework: NSDate */) {
	if d._WorkoutSessionDidBeginActivityWithConfigurationDate != nil {
		d._WorkoutSessionDidBeginActivityWithConfigurationDate(workoutSession, workoutConfiguration, date)
	}
}

// HasWorkoutSessionDidBeginActivityWithConfigurationDate returns true if a handler for WorkoutSessionDidBeginActivityWithConfigurationDate has been set.
func (d *HKWorkoutSessionDelegate) HasWorkoutSessionDidBeginActivityWithConfigurationDate() bool {
	return d._WorkoutSessionDidBeginActivityWithConfigurationDate != nil
}

// WorkoutSessionDidDisconnectFromRemoteDeviceWithError implements the PHKWorkoutSessionDelegate interface.
func (d *HKWorkoutSessionDelegate) WorkoutSessionDidDisconnectFromRemoteDeviceWithError(workoutSession IHKWorkoutSession, error_ objc.IObject /* cross-framework: Error */) {
	if d._WorkoutSessionDidDisconnectFromRemoteDeviceWithError != nil {
		d._WorkoutSessionDidDisconnectFromRemoteDeviceWithError(workoutSession, error_)
	}
}

// HasWorkoutSessionDidDisconnectFromRemoteDeviceWithError returns true if a handler for WorkoutSessionDidDisconnectFromRemoteDeviceWithError has been set.
func (d *HKWorkoutSessionDelegate) HasWorkoutSessionDidDisconnectFromRemoteDeviceWithError() bool {
	return d._WorkoutSessionDidDisconnectFromRemoteDeviceWithError != nil
}

// WorkoutSessionDidEndActivityWithConfigurationDate implements the PHKWorkoutSessionDelegate interface.
func (d *HKWorkoutSessionDelegate) WorkoutSessionDidEndActivityWithConfigurationDate(workoutSession IHKWorkoutSession, workoutConfiguration IHKWorkoutConfiguration, date objc.IObject /* cross-framework: NSDate */) {
	if d._WorkoutSessionDidEndActivityWithConfigurationDate != nil {
		d._WorkoutSessionDidEndActivityWithConfigurationDate(workoutSession, workoutConfiguration, date)
	}
}

// HasWorkoutSessionDidEndActivityWithConfigurationDate returns true if a handler for WorkoutSessionDidEndActivityWithConfigurationDate has been set.
func (d *HKWorkoutSessionDelegate) HasWorkoutSessionDidEndActivityWithConfigurationDate() bool {
	return d._WorkoutSessionDidEndActivityWithConfigurationDate != nil
}

// WorkoutSessionDidGenerateEvent implements the PHKWorkoutSessionDelegate interface.
func (d *HKWorkoutSessionDelegate) WorkoutSessionDidGenerateEvent(workoutSession IHKWorkoutSession, event IHKWorkoutEvent) {
	if d._WorkoutSessionDidGenerateEvent != nil {
		d._WorkoutSessionDidGenerateEvent(workoutSession, event)
	}
}

// HasWorkoutSessionDidGenerateEvent returns true if a handler for WorkoutSessionDidGenerateEvent has been set.
func (d *HKWorkoutSessionDelegate) HasWorkoutSessionDidGenerateEvent() bool {
	return d._WorkoutSessionDidGenerateEvent != nil
}

// WorkoutSessionDidReceiveDataFromRemoteWorkoutSession implements the PHKWorkoutSessionDelegate interface.
func (d *HKWorkoutSessionDelegate) WorkoutSessionDidReceiveDataFromRemoteWorkoutSession(workoutSession IHKWorkoutSession, data []foundation.Data) {
	if d._WorkoutSessionDidReceiveDataFromRemoteWorkoutSession != nil {
		d._WorkoutSessionDidReceiveDataFromRemoteWorkoutSession(workoutSession, data)
	}
}

// HasWorkoutSessionDidReceiveDataFromRemoteWorkoutSession returns true if a handler for WorkoutSessionDidReceiveDataFromRemoteWorkoutSession has been set.
func (d *HKWorkoutSessionDelegate) HasWorkoutSessionDidReceiveDataFromRemoteWorkoutSession() bool {
	return d._WorkoutSessionDidReceiveDataFromRemoteWorkoutSession != nil
}

// WorkoutSessionDidChangeToStateFromStateDate implements the PHKWorkoutSessionDelegate interface.
func (d *HKWorkoutSessionDelegate) WorkoutSessionDidChangeToStateFromStateDate(workoutSession IHKWorkoutSession, toState HKWorkoutSessionState, fromState HKWorkoutSessionState, date objc.IObject /* cross-framework: NSDate */) {
	if d._WorkoutSessionDidChangeToStateFromStateDate != nil {
		d._WorkoutSessionDidChangeToStateFromStateDate(workoutSession, toState, fromState, date)
	}
}

// HasWorkoutSessionDidChangeToStateFromStateDate returns true if a handler for WorkoutSessionDidChangeToStateFromStateDate has been set.
func (d *HKWorkoutSessionDelegate) HasWorkoutSessionDidChangeToStateFromStateDate() bool {
	return d._WorkoutSessionDidChangeToStateFromStateDate != nil
}

// WorkoutSessionDidFailWithError implements the PHKWorkoutSessionDelegate interface.
func (d *HKWorkoutSessionDelegate) WorkoutSessionDidFailWithError(workoutSession IHKWorkoutSession, error_ objc.IObject /* cross-framework: Error */) {
	if d._WorkoutSessionDidFailWithError != nil {
		d._WorkoutSessionDidFailWithError(workoutSession, error_)
	}
}

// HasWorkoutSessionDidFailWithError returns true if a handler for WorkoutSessionDidFailWithError has been set.
func (d *HKWorkoutSessionDelegate) HasWorkoutSessionDidFailWithError() bool {
	return d._WorkoutSessionDidFailWithError != nil
}
