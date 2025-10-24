// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PCKSyncEngineDelegate is the CKSyncEngineDelegate protocol interface.
//
// An interface for providing record data to a sync engine and customizing that engine’s behavior.
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+
//
// See: doc://com.apple.cloudkit/documentation/CloudKit/CKSyncEngineDelegate-3c38p
type PCKSyncEngineDelegate interface {
	// Required methods
	SyncEngineHandleEvent(syncEngine ICKSyncEngine, event ICKSyncEngineEvent)/* debug [protocol_interface/required_method]: SyncEngineHandleEvent */
	SyncEngineNextFetchChangesOptionsForContext(syncEngine ICKSyncEngine, context ICKSyncEngineFetchChangesContext) CKSyncEngineFetchChangesOptions/* debug [protocol_interface/required_method]: SyncEngineNextFetchChangesOptionsForContext */
	SyncEngineNextRecordZoneChangeBatchForContext(syncEngine ICKSyncEngine, context ICKSyncEngineSendChangesContext) CKSyncEngineRecordZoneChangeBatch/* debug [protocol_interface/required_method]: SyncEngineNextRecordZoneChangeBatchForContext */
}

// CKSyncEngineDelegate is a delegate implementation builder for the PCKSyncEngineDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CKSyncEngineDelegate struct {
	_SyncEngineHandleEvent func(syncEngine ICKSyncEngine, event ICKSyncEngineEvent)
	_SyncEngineNextFetchChangesOptionsForContext func(syncEngine ICKSyncEngine, context ICKSyncEngineFetchChangesContext) CKSyncEngineFetchChangesOptions
	_SyncEngineNextRecordZoneChangeBatchForContext func(syncEngine ICKSyncEngine, context ICKSyncEngineSendChangesContext) CKSyncEngineRecordZoneChangeBatch
}

// SetSyncEngineHandleEvent sets the handler for the SyncEngineHandleEvent delegate method.
//
// Tells the delegate to handle the specified sync event.
func (d *CKSyncEngineDelegate) SetSyncEngineHandleEvent(f func(syncEngine ICKSyncEngine, event ICKSyncEngineEvent)) {
	d._SyncEngineHandleEvent = f
}

// SetSyncEngineNextFetchChangesOptionsForContext sets the handler for the SyncEngineNextFetchChangesOptionsForContext delegate method.
func (d *CKSyncEngineDelegate) SetSyncEngineNextFetchChangesOptionsForContext(f func(syncEngine ICKSyncEngine, context ICKSyncEngineFetchChangesContext) CKSyncEngineFetchChangesOptions) {
	d._SyncEngineNextFetchChangesOptionsForContext = f
}

// SetSyncEngineNextRecordZoneChangeBatchForContext sets the handler for the SyncEngineNextRecordZoneChangeBatchForContext delegate method.
//
// Asks the delegate to provide the next set of record changes to send to the server.
func (d *CKSyncEngineDelegate) SetSyncEngineNextRecordZoneChangeBatchForContext(f func(syncEngine ICKSyncEngine, context ICKSyncEngineSendChangesContext) CKSyncEngineRecordZoneChangeBatch) {
	d._SyncEngineNextRecordZoneChangeBatchForContext = f
}

// SyncEngineHandleEvent implements the PCKSyncEngineDelegate interface.
func (d *CKSyncEngineDelegate) SyncEngineHandleEvent(syncEngine ICKSyncEngine, event ICKSyncEngineEvent) {
	if d._SyncEngineHandleEvent != nil {
		d._SyncEngineHandleEvent(syncEngine, event)
	}
}

// HasSyncEngineHandleEvent returns true if a handler for SyncEngineHandleEvent has been set.
func (d *CKSyncEngineDelegate) HasSyncEngineHandleEvent() bool {
	return d._SyncEngineHandleEvent != nil
}

// SyncEngineNextFetchChangesOptionsForContext implements the PCKSyncEngineDelegate interface.
func (d *CKSyncEngineDelegate) SyncEngineNextFetchChangesOptionsForContext(syncEngine ICKSyncEngine, context ICKSyncEngineFetchChangesContext) CKSyncEngineFetchChangesOptions {
	if d._SyncEngineNextFetchChangesOptionsForContext != nil {
		return d._SyncEngineNextFetchChangesOptionsForContext(syncEngine, context)
	}
	var zero CKSyncEngineFetchChangesOptions
	return zero
}

// HasSyncEngineNextFetchChangesOptionsForContext returns true if a handler for SyncEngineNextFetchChangesOptionsForContext has been set.
func (d *CKSyncEngineDelegate) HasSyncEngineNextFetchChangesOptionsForContext() bool {
	return d._SyncEngineNextFetchChangesOptionsForContext != nil
}

// SyncEngineNextRecordZoneChangeBatchForContext implements the PCKSyncEngineDelegate interface.
func (d *CKSyncEngineDelegate) SyncEngineNextRecordZoneChangeBatchForContext(syncEngine ICKSyncEngine, context ICKSyncEngineSendChangesContext) CKSyncEngineRecordZoneChangeBatch {
	if d._SyncEngineNextRecordZoneChangeBatchForContext != nil {
		return d._SyncEngineNextRecordZoneChangeBatchForContext(syncEngine, context)
	}
	var zero CKSyncEngineRecordZoneChangeBatch
	return zero
}

// HasSyncEngineNextRecordZoneChangeBatchForContext returns true if a handler for SyncEngineNextRecordZoneChangeBatchForContext has been set.
func (d *CKSyncEngineDelegate) HasSyncEngineNextRecordZoneChangeBatchForContext() bool {
	return d._SyncEngineNextRecordZoneChangeBatchForContext != nil
}
