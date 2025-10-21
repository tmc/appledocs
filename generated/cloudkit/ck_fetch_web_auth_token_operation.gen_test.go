// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKFetchWebAuthTokenOperation


// ExampleNewCKFetchWebAuthTokenOperationWithAPIToken demonstrates how to create a CKFetchWebAuthTokenOperation instance using NewCKFetchWebAuthTokenOperationWithAPIToken.
// Creates a fetch operation for the specified API token.
func ExampleNewCKFetchWebAuthTokenOperationWithAPIToken() {
	_ = cloudkit.NewCKFetchWebAuthTokenOperationWithAPIToken(
		"APIToken", // APIToken string
	)
	// Output:
}

// ExampleNewCKFetchWebAuthTokenOperation demonstrates how to create a CKFetchWebAuthTokenOperation instance.
// Creates an empty fetch operation.
func ExampleNewCKFetchWebAuthTokenOperation() {
	_ = cloudkit.NewCKFetchWebAuthTokenOperation()
	// Output:
}


