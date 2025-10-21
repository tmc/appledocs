// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKShare

// ExampleNewCKShareWithRecordZoneID demonstrates how to create a CKShare instance using NewCKShareWithRecordZoneID.
// Creates a new share for the specified record zone.
func ExampleNewCKShareWithRecordZoneID() {
	_ = cloudkit.NewCKShareWithRecordZoneID(
		cloudkit.CKRecordZoneID{}, // recordZoneID CKRecordZoneID
	)
	// Output:
}
// ExampleNewCKShareWithRootRecord demonstrates how to create a CKShare instance using NewCKShareWithRootRecord.
// Creates a new share for the specified record.
func ExampleNewCKShareWithRootRecord() {
	_ = cloudkit.NewCKShareWithRootRecord(
		cloudkit.CKRecord{}, // rootRecord CKRecord
	)
	// Output:
}
// ExampleNewCKShareWithRootRecordShareID demonstrates how to create a CKShare instance using NewCKShareWithRootRecordShareID.
// Creates a new share for the specified record and record ID.
func ExampleNewCKShareWithRootRecordShareID() {
	_ = cloudkit.NewCKShareWithRootRecordShareID(
		cloudkit.CKRecord{}, // rootRecord CKRecord
		cloudkit.CKRecordID{}, // shareID CKRecordID
	)
	// Output:
}
