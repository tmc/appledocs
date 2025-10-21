// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKSyncEnginePendingZoneDelete

// ExampleNewCKSyncEnginePendingZoneDeleteWithZoneID demonstrates how to create a CKSyncEnginePendingZoneDelete instance using NewCKSyncEnginePendingZoneDeleteWithZoneID.
// Creates a pending zone delete for the specified record zone identifier.
func ExampleNewCKSyncEnginePendingZoneDeleteWithZoneID() {
	_ = cloudkit.NewCKSyncEnginePendingZoneDeleteWithZoneID(
		cloudkit.CKRecordZoneID{}, // zoneID CKRecordZoneID
	)
	// Output:
}
