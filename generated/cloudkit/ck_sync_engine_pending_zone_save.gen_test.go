// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKSyncEnginePendingZoneSave

// ExampleNewCKSyncEnginePendingZoneSaveWithZone demonstrates how to create a CKSyncEnginePendingZoneSave instance using NewCKSyncEnginePendingZoneSaveWithZone.
// Creates a pending zone save for the specified record zone.
func ExampleNewCKSyncEnginePendingZoneSaveWithZone() {
	_ = cloudkit.NewCKSyncEnginePendingZoneSaveWithZone(
		cloudkit.CKRecordZone{}, // zone CKRecordZone
	)
	// Output:
}
