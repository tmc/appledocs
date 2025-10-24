// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewFetchRequest

// ExampleNewFetchRequest demonstrates how to create a FetchRequest instance.
// Creates a new fetch request.
func ExampleNewFetchRequest() {
	_ = coredata.NewFetchRequest()
	// Output:
}

// ExampleNewFetchRequestWithEntityName demonstrates how to create a FetchRequest instance using NewFetchRequestWithEntityName.
// Initializes a fetch request configured with a given entity name.
func ExampleNewFetchRequestWithEntityName() {
	_ = coredata.NewFetchRequestWithEntityName(
		"entityName", // entityName string
	)
	// Output:
}
