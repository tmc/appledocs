// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKFetchRecordsOperation

// ExampleNewCKFetchRecordsOperation demonstrates how to create a CKFetchRecordsOperation instance.
// Creates an empty fetch operation.
func ExampleNewCKFetchRecordsOperation() {
	_ = cloudkit.NewCKFetchRecordsOperation()
	// Output:
}
// ExampleNewCKFetchRecordsOperationWithRecordIDs demonstrates how to create a CKFetchRecordsOperation instance using NewCKFetchRecordsOperationWithRecordIDs.
// Creates a fetch operation for retrieving the records with the specified IDs.
func ExampleNewCKFetchRecordsOperationWithRecordIDs() {
	_ = cloudkit.NewCKFetchRecordsOperationWithRecordIDs(
		[]cloudkit.CKRecordID{}, // recordIDs []CKRecordID
	)
	// Output:
}
