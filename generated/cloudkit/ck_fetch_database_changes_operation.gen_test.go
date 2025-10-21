// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKFetchDatabaseChangesOperation

// ExampleNewCKFetchDatabaseChangesOperation demonstrates how to create a CKFetchDatabaseChangesOperation instance.
// Creates an empty fetch database changes operation.
func ExampleNewCKFetchDatabaseChangesOperation() {
	_ = cloudkit.NewCKFetchDatabaseChangesOperation()
	// Output:
}
// ExampleNewCKFetchDatabaseChangesOperationWithPreviousServerChangeToken demonstrates how to create a CKFetchDatabaseChangesOperation instance using NewCKFetchDatabaseChangesOperationWithPreviousServerChangeToken.
// Creates an operation for fetching database changes.
func ExampleNewCKFetchDatabaseChangesOperationWithPreviousServerChangeToken() {
	_ = cloudkit.NewCKFetchDatabaseChangesOperationWithPreviousServerChangeToken(
		cloudkit.CKServerChangeToken{}, // previousServerChangeToken CKServerChangeToken
	)
	// Output:
}
