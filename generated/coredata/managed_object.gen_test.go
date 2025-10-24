// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewManagedObject

// ExampleNewManagedObjectWithEntityInsertIntoManagedObjectContext demonstrates how to create a ManagedObject instance using NewManagedObjectWithEntityInsertIntoManagedObjectContext.
// Initializes a managed object from an entity description and inserts it into the specified managed object context.
func ExampleNewManagedObjectWithEntityInsertIntoManagedObjectContext() {
	_ = coredata.NewManagedObjectWithEntityInsertIntoManagedObjectContext(
		coredata.NSEntityDescription{},    // entity NSEntityDescription
		coredata.NSManagedObjectContext{}, // context NSManagedObjectContext
	)
	// Output:
}
