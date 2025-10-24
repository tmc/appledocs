// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKAcceptSharesOperation

// ExampleNewCKAcceptSharesOperation demonstrates how to create a CKAcceptSharesOperation instance.
// Creates an operation for accepting shares.
func ExampleNewCKAcceptSharesOperation() {
	_ = cloudkit.NewCKAcceptSharesOperation()
	// Output:
}
// ExampleNewCKAcceptSharesOperationWithShareMetadatas demonstrates how to create a CKAcceptSharesOperation instance using NewCKAcceptSharesOperationWithShareMetadatas.
// Creates an operation for accepting the specified shares.
func ExampleNewCKAcceptSharesOperationWithShareMetadatas() {
	_ = cloudkit.NewCKAcceptSharesOperationWithShareMetadatas(
		[]cloudkit.ICKShareMetadata{}, // shareMetadatas []ICKShareMetadata
	)
	// Output:
}
