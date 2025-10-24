// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewManagedObjectContext

// ExampleNewManagedObjectContext demonstrates how to create a ManagedObjectContext instance.
func ExampleNewManagedObjectContext() {
	_ = coredata.NewManagedObjectContext()
	// Output:
}

// ExampleNewManagedObjectContextWithConcurrencyType demonstrates how to create a ManagedObjectContext instance using NewManagedObjectContextWithConcurrencyType.
// Creates a context that uses the specified concurrency type.
func ExampleNewManagedObjectContextWithConcurrencyType() {
	_ = coredata.NewManagedObjectContextWithConcurrencyType(
		coredata.ManagedObjectContextConcurrencyType{}, // ct ManagedObjectContextConcurrencyType
	)
	// Output:
}
