// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKSyncEnginePendingRecordZoneChange

// ExampleNewCKSyncEnginePendingRecordZoneChangeWithRecordIDType demonstrates how to create a CKSyncEnginePendingRecordZoneChange instance using NewCKSyncEnginePendingRecordZoneChangeWithRecordIDType.
// Creates a record zone change of the specified type for the given record.
func ExampleNewCKSyncEnginePendingRecordZoneChangeWithRecordIDType() {
	_ = cloudkit.NewCKSyncEnginePendingRecordZoneChangeWithRecordIDType(
		cloudkit.CKRecordID{}, // recordID CKRecordID
		cloudkit.CKSyncEnginePendingRecordZoneChangeType{}, // type CKSyncEnginePendingRecordZoneChangeType
	)
	// Output:
}
