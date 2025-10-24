// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PNISessionDelegate is the NISessionDelegate protocol interface.
//
// An object that monitors and reacts to session updates.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - watchOS 7.3+
//
// See: doc://com.apple.nearbyinteraction/documentation/NearbyInteraction/NISessionDelegate
type PNISessionDelegate interface {
	// Optional methods
	SessionDidGenerateShareableConfigurationDataForObject(session INISession, shareableConfigurationData objc.IObject /* cross-framework: NSData */, object ININearbyObject)
	HasSessionDidGenerateShareableConfigurationDataForObject() bool
	SessionDidInvalidateWithError(session INISession, error_ objc.IObject /* cross-framework: Error */)
	HasSessionDidInvalidateWithError() bool
	SessionDidRemoveNearbyObjectsWithReason(session INISession, nearbyObjects []NINearbyObject, reason NINearbyObjectRemovalReason)
	HasSessionDidRemoveNearbyObjectsWithReason() bool
	SessionDidUpdateNearbyObjects(session INISession, nearbyObjects []NINearbyObject)
	HasSessionDidUpdateNearbyObjects() bool
	SessionDidUpdateAlgorithmConvergenceForObject(session INISession, convergence INIAlgorithmConvergence, object ININearbyObject)
	HasSessionDidUpdateAlgorithmConvergenceForObject() bool
	SessionDidUpdateDLTDOAMeasurements(session INISession, measurements []NIDLTDOAMeasurement)
	HasSessionDidUpdateDLTDOAMeasurements() bool
	SessionDidStartRunning(session INISession)
	HasSessionDidStartRunning() bool
	SessionSuspensionEnded(session INISession)
	HasSessionSuspensionEnded() bool
	SessionWasSuspended(session INISession)
	HasSessionWasSuspended() bool
}

// NISessionDelegate is a delegate implementation builder for the PNISessionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type NISessionDelegate struct {
	_SessionDidGenerateShareableConfigurationDataForObject func(session INISession, shareableConfigurationData objc.IObject /* cross-framework: NSData */, object ININearbyObject)
	_SessionDidInvalidateWithError func(session INISession, error_ objc.IObject /* cross-framework: Error */)
	_SessionDidRemoveNearbyObjectsWithReason func(session INISession, nearbyObjects []NINearbyObject, reason NINearbyObjectRemovalReason)
	_SessionDidUpdateNearbyObjects func(session INISession, nearbyObjects []NINearbyObject)
	_SessionDidUpdateAlgorithmConvergenceForObject func(session INISession, convergence INIAlgorithmConvergence, object ININearbyObject)
	_SessionDidUpdateDLTDOAMeasurements func(session INISession, measurements []NIDLTDOAMeasurement)
	_SessionDidStartRunning func(session INISession)
	_SessionSuspensionEnded func(session INISession)
	_SessionWasSuspended func(session INISession)
}

// SetSessionDidGenerateShareableConfigurationDataForObject sets the handler for the SessionDidGenerateShareableConfigurationDataForObject delegate method.
//
// Provides configuration data to share with a third-party accessory.
func (d *NISessionDelegate) SetSessionDidGenerateShareableConfigurationDataForObject(f func(session INISession, shareableConfigurationData objc.IObject /* cross-framework: NSData */, object ININearbyObject)) {
	d._SessionDidGenerateShareableConfigurationDataForObject = f
}

// SetSessionDidInvalidateWithError sets the handler for the SessionDidInvalidateWithError delegate method.
//
// Notifies you of an invalidated session.
func (d *NISessionDelegate) SetSessionDidInvalidateWithError(f func(session INISession, error_ objc.IObject /* cross-framework: Error */)) {
	d._SessionDidInvalidateWithError = f
}

// SetSessionDidRemoveNearbyObjectsWithReason sets the handler for the SessionDidRemoveNearbyObjectsWithReason delegate method.
//
// Notifies you when the session removes one or more nearby objects.
func (d *NISessionDelegate) SetSessionDidRemoveNearbyObjectsWithReason(f func(session INISession, nearbyObjects []NINearbyObject, reason NINearbyObjectRemovalReason)) {
	d._SessionDidRemoveNearbyObjectsWithReason = f
}

// SetSessionDidUpdateNearbyObjects sets the handler for the SessionDidUpdateNearbyObjects delegate method.
//
// Notifies you when the session updates nearby objects.
func (d *NISessionDelegate) SetSessionDidUpdateNearbyObjects(f func(session INISession, nearbyObjects []NINearbyObject)) {
	d._SessionDidUpdateNearbyObjects = f
}

// SetSessionDidUpdateAlgorithmConvergenceForObject sets the handler for the SessionDidUpdateAlgorithmConvergenceForObject delegate method.
//
// Provides recommended actions the user can take to facilitate the framework’s Camera Assistance.
func (d *NISessionDelegate) SetSessionDidUpdateAlgorithmConvergenceForObject(f func(session INISession, convergence INIAlgorithmConvergence, object ININearbyObject)) {
	d._SessionDidUpdateAlgorithmConvergenceForObject = f
}

// SetSessionDidUpdateDLTDOAMeasurements sets the handler for the SessionDidUpdateDLTDOAMeasurements delegate method.
//
// Provides device ranging estimates for a Downlink Time-Difference-of-Arrival session.
func (d *NISessionDelegate) SetSessionDidUpdateDLTDOAMeasurements(f func(session INISession, measurements []NIDLTDOAMeasurement)) {
	d._SessionDidUpdateDLTDOAMeasurements = f
}

// SetSessionDidStartRunning sets the handler for the SessionDidStartRunning delegate method.
//
// Notifies the app when a session starts or resumes running.
func (d *NISessionDelegate) SetSessionDidStartRunning(f func(session INISession)) {
	d._SessionDidStartRunning = f
}

// SetSessionSuspensionEnded sets the handler for the SessionSuspensionEnded delegate method.
//
// Notifies you of the end of a session’s suspension.
func (d *NISessionDelegate) SetSessionSuspensionEnded(f func(session INISession)) {
	d._SessionSuspensionEnded = f
}

// SetSessionWasSuspended sets the handler for the SessionWasSuspended delegate method.
//
// Notifies you of a suspended session.
func (d *NISessionDelegate) SetSessionWasSuspended(f func(session INISession)) {
	d._SessionWasSuspended = f
}

// SessionDidGenerateShareableConfigurationDataForObject implements the PNISessionDelegate interface.
func (d *NISessionDelegate) SessionDidGenerateShareableConfigurationDataForObject(session INISession, shareableConfigurationData objc.IObject /* cross-framework: NSData */, object ININearbyObject) {
	if d._SessionDidGenerateShareableConfigurationDataForObject != nil {
		d._SessionDidGenerateShareableConfigurationDataForObject(session, shareableConfigurationData, object)
	}
}

// HasSessionDidGenerateShareableConfigurationDataForObject returns true if a handler for SessionDidGenerateShareableConfigurationDataForObject has been set.
func (d *NISessionDelegate) HasSessionDidGenerateShareableConfigurationDataForObject() bool {
	return d._SessionDidGenerateShareableConfigurationDataForObject != nil
}

// SessionDidInvalidateWithError implements the PNISessionDelegate interface.
func (d *NISessionDelegate) SessionDidInvalidateWithError(session INISession, error_ objc.IObject /* cross-framework: Error */) {
	if d._SessionDidInvalidateWithError != nil {
		d._SessionDidInvalidateWithError(session, error_)
	}
}

// HasSessionDidInvalidateWithError returns true if a handler for SessionDidInvalidateWithError has been set.
func (d *NISessionDelegate) HasSessionDidInvalidateWithError() bool {
	return d._SessionDidInvalidateWithError != nil
}

// SessionDidRemoveNearbyObjectsWithReason implements the PNISessionDelegate interface.
func (d *NISessionDelegate) SessionDidRemoveNearbyObjectsWithReason(session INISession, nearbyObjects []NINearbyObject, reason NINearbyObjectRemovalReason) {
	if d._SessionDidRemoveNearbyObjectsWithReason != nil {
		d._SessionDidRemoveNearbyObjectsWithReason(session, nearbyObjects, reason)
	}
}

// HasSessionDidRemoveNearbyObjectsWithReason returns true if a handler for SessionDidRemoveNearbyObjectsWithReason has been set.
func (d *NISessionDelegate) HasSessionDidRemoveNearbyObjectsWithReason() bool {
	return d._SessionDidRemoveNearbyObjectsWithReason != nil
}

// SessionDidUpdateNearbyObjects implements the PNISessionDelegate interface.
func (d *NISessionDelegate) SessionDidUpdateNearbyObjects(session INISession, nearbyObjects []NINearbyObject) {
	if d._SessionDidUpdateNearbyObjects != nil {
		d._SessionDidUpdateNearbyObjects(session, nearbyObjects)
	}
}

// HasSessionDidUpdateNearbyObjects returns true if a handler for SessionDidUpdateNearbyObjects has been set.
func (d *NISessionDelegate) HasSessionDidUpdateNearbyObjects() bool {
	return d._SessionDidUpdateNearbyObjects != nil
}

// SessionDidUpdateAlgorithmConvergenceForObject implements the PNISessionDelegate interface.
func (d *NISessionDelegate) SessionDidUpdateAlgorithmConvergenceForObject(session INISession, convergence INIAlgorithmConvergence, object ININearbyObject) {
	if d._SessionDidUpdateAlgorithmConvergenceForObject != nil {
		d._SessionDidUpdateAlgorithmConvergenceForObject(session, convergence, object)
	}
}

// HasSessionDidUpdateAlgorithmConvergenceForObject returns true if a handler for SessionDidUpdateAlgorithmConvergenceForObject has been set.
func (d *NISessionDelegate) HasSessionDidUpdateAlgorithmConvergenceForObject() bool {
	return d._SessionDidUpdateAlgorithmConvergenceForObject != nil
}

// SessionDidUpdateDLTDOAMeasurements implements the PNISessionDelegate interface.
func (d *NISessionDelegate) SessionDidUpdateDLTDOAMeasurements(session INISession, measurements []NIDLTDOAMeasurement) {
	if d._SessionDidUpdateDLTDOAMeasurements != nil {
		d._SessionDidUpdateDLTDOAMeasurements(session, measurements)
	}
}

// HasSessionDidUpdateDLTDOAMeasurements returns true if a handler for SessionDidUpdateDLTDOAMeasurements has been set.
func (d *NISessionDelegate) HasSessionDidUpdateDLTDOAMeasurements() bool {
	return d._SessionDidUpdateDLTDOAMeasurements != nil
}

// SessionDidStartRunning implements the PNISessionDelegate interface.
func (d *NISessionDelegate) SessionDidStartRunning(session INISession) {
	if d._SessionDidStartRunning != nil {
		d._SessionDidStartRunning(session)
	}
}

// HasSessionDidStartRunning returns true if a handler for SessionDidStartRunning has been set.
func (d *NISessionDelegate) HasSessionDidStartRunning() bool {
	return d._SessionDidStartRunning != nil
}

// SessionSuspensionEnded implements the PNISessionDelegate interface.
func (d *NISessionDelegate) SessionSuspensionEnded(session INISession) {
	if d._SessionSuspensionEnded != nil {
		d._SessionSuspensionEnded(session)
	}
}

// HasSessionSuspensionEnded returns true if a handler for SessionSuspensionEnded has been set.
func (d *NISessionDelegate) HasSessionSuspensionEnded() bool {
	return d._SessionSuspensionEnded != nil
}

// SessionWasSuspended implements the PNISessionDelegate interface.
func (d *NISessionDelegate) SessionWasSuspended(session INISession) {
	if d._SessionWasSuspended != nil {
		d._SessionWasSuspended(session)
	}
}

// HasSessionWasSuspended returns true if a handler for SessionWasSuspended has been set.
func (d *NISessionDelegate) HasSessionWasSuspended() bool {
	return d._SessionWasSuspended != nil
}
