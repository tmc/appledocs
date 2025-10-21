// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKQueryOperation

// ExampleNewCKQueryOperation demonstrates how to create a CKQueryOperation instance.
// Creates an empty query operation.
func ExampleNewCKQueryOperation() {
	_ = cloudkit.NewCKQueryOperation()
	// Output:
}
// ExampleNewCKQueryOperationWithCursor demonstrates how to create a CKQueryOperation instance using NewCKQueryOperationWithCursor.
// Creates an operation with additional results from a previous search.
func ExampleNewCKQueryOperationWithCursor() {
	_ = cloudkit.NewCKQueryOperationWithCursor(
		cloudkit.CKQueryCursor{}, // cursor CKQueryCursor
	)
	// Output:
}
// ExampleNewCKQueryOperationWithQuery demonstrates how to create a CKQueryOperation instance using NewCKQueryOperationWithQuery.
// Creates an operation that searches for records in the specified record zone.
func ExampleNewCKQueryOperationWithQuery() {
	_ = cloudkit.NewCKQueryOperationWithQuery(
		cloudkit.CKQuery{}, // query CKQuery
	)
	// Output:
}
