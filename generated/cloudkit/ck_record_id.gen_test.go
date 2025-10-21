// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKRecordID

// ExampleNewCKRecordIDWithRecordName demonstrates how to create a CKRecordID instance using NewCKRecordIDWithRecordName.
// Creates a new record ID with the specified name in the default zone.
func ExampleNewCKRecordIDWithRecordName() {
	_ = cloudkit.NewCKRecordIDWithRecordName(
		"recordName", // recordName string
	)
	// Output:
}
// ExampleNewCKRecordIDWithRecordNameZoneID demonstrates how to create a CKRecordID instance using NewCKRecordIDWithRecordNameZoneID.
// Creates a new record ID with the specified name and zone information.
func ExampleNewCKRecordIDWithRecordNameZoneID() {
	_ = cloudkit.NewCKRecordIDWithRecordNameZoneID(
		"recordName", // recordName string
		cloudkit.CKRecordZoneID{}, // zoneID CKRecordZoneID
	)
	// Output:
}
