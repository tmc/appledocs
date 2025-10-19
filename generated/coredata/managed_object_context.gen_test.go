// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)


// ExampleNewManagedObjectContextWithConcurrencyType demonstrates how to create a ManagedObjectContext instance using NewManagedObjectContextWithConcurrencyType.
// Creates a context that uses the specified concurrency type.
func ExampleNewManagedObjectContextWithConcurrencyType() {
	_ = coredata.NewManagedObjectContextWithConcurrencyType(
		nil, // ct unsafe.Pointer
	)
	// Output:
}


