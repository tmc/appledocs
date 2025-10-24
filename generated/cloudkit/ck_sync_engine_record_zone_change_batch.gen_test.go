// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKSyncEngineRecordZoneChangeBatch

// ExampleNewCKSyncEngineRecordZoneChangeBatchWithRecordsToSaveRecordIDsToDeleteAtomicByZone demonstrates how to create a CKSyncEngineRecordZoneChangeBatch instance using NewCKSyncEngineRecordZoneChangeBatchWithRecordsToSaveRecordIDsToDeleteAtomicByZone.
// Creates a batch of records to modify.
func ExampleNewCKSyncEngineRecordZoneChangeBatchWithRecordsToSaveRecordIDsToDeleteAtomicByZone() {
	_ = cloudkit.NewCKSyncEngineRecordZoneChangeBatchWithRecordsToSaveRecordIDsToDeleteAtomicByZone(
		[]cloudkit.ICKRecord{}, // recordsToSave []ICKRecord
		[]cloudkit.ICKRecordID{}, // recordIDsToDelete []ICKRecordID
		false, // atomicByZone bool
	)
	// Output:
}
