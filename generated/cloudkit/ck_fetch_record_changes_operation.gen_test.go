// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKFetchRecordChangesOperation

// ExampleNewCKFetchRecordChangesOperation demonstrates how to create a CKFetchRecordChangesOperation instance.
// Creates an empty fetch record changes operation.
func ExampleNewCKFetchRecordChangesOperation() {
	_ = cloudkit.NewCKFetchRecordChangesOperation()
	// Output:
}
// ExampleNewCKFetchRecordChangesOperationWithRecordZoneIDPreviousServerChangeToken demonstrates how to create a CKFetchRecordChangesOperation instance using NewCKFetchRecordChangesOperationWithRecordZoneIDPreviousServerChangeToken.
// Creates an operation for fetching changes in the specified record zone.
func ExampleNewCKFetchRecordChangesOperationWithRecordZoneIDPreviousServerChangeToken() {
	_ = cloudkit.NewCKFetchRecordChangesOperationWithRecordZoneIDPreviousServerChangeToken(
		cloudkit.CKRecordZoneID{}, // recordZoneID CKRecordZoneID
		cloudkit.CKServerChangeToken{}, // previousServerChangeToken CKServerChangeToken
	)
	// Output:
}
