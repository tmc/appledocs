// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKRecordZone

// ExampleNewCKRecordZoneWithZoneID demonstrates how to create a CKRecordZone instance using NewCKRecordZoneWithZoneID.
// Creates a record zone object with the specified zone ID.
func ExampleNewCKRecordZoneWithZoneID() {
	_ = cloudkit.NewCKRecordZoneWithZoneID(
		cloudkit.CKRecordZoneID{}, // zoneID CKRecordZoneID
	)
	// Output:
}
// ExampleNewCKRecordZoneWithZoneName demonstrates how to create a CKRecordZone instance using NewCKRecordZoneWithZoneName.
// Creates a record zone object with the specified zone name.
func ExampleNewCKRecordZoneWithZoneName() {
	_ = cloudkit.NewCKRecordZoneWithZoneName(
		"zoneName", // zoneName string
	)
	// Output:
}
