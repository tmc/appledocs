// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewBatchDeleteRequest

// ExampleNewBatchDeleteRequestWithFetchRequest demonstrates how to create a BatchDeleteRequest instance using NewBatchDeleteRequestWithFetchRequest.
// Creates a request that deletes the results of the specified fetch request.
func ExampleNewBatchDeleteRequestWithFetchRequest() {
	_ = coredata.NewBatchDeleteRequestWithFetchRequest(
		coredata.NSFetchRequest{}, // fetch NSFetchRequest
	)
	// Output:
}

// ExampleNewBatchDeleteRequestWithObjectIDs demonstrates how to create a BatchDeleteRequest instance using NewBatchDeleteRequestWithObjectIDs.
// Creates a request that deletes the managed objects with the specified identifiers.
func ExampleNewBatchDeleteRequestWithObjectIDs() {
	_ = coredata.NewBatchDeleteRequestWithObjectIDs(
		[]coredata.ManagedObjectID{}, // objects []ManagedObjectID
	)
	// Output:
}
