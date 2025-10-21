// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKRecordZoneID

// ExampleNewCKRecordZoneIDWithZoneNameOwnerName demonstrates how to create a CKRecordZoneID instance using NewCKRecordZoneIDWithZoneNameOwnerName.
// Creates a record zone ID with the specified name and owner.
func ExampleNewCKRecordZoneIDWithZoneNameOwnerName() {
	_ = cloudkit.NewCKRecordZoneIDWithZoneNameOwnerName(
		"zoneName", // zoneName string
		"ownerName", // ownerName string
	)
	// Output:
}
