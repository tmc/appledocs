// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewBatchUpdateRequest

// ExampleNewBatchUpdateRequestWithEntity demonstrates how to create a BatchUpdateRequest instance using NewBatchUpdateRequestWithEntity.
// Creates a batch-update request for a managed entity.
func ExampleNewBatchUpdateRequestWithEntity() {
	_ = coredata.NewBatchUpdateRequestWithEntity(
		coredata.NSEntityDescription{}, // entity NSEntityDescription
	)
	// Output:
}

// ExampleNewBatchUpdateRequestWithEntityName demonstrates how to create a BatchUpdateRequest instance using NewBatchUpdateRequestWithEntityName.
// Creates a batch-update request for a named managed entity.
func ExampleNewBatchUpdateRequestWithEntityName() {
	_ = coredata.NewBatchUpdateRequestWithEntityName(
		"entityName", // entityName string
	)
	// Output:
}
