// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewPersistentStoreCoordinator

// ExampleNewPersistentStoreCoordinatorWithManagedObjectModel demonstrates how to create a PersistentStoreCoordinator instance using NewPersistentStoreCoordinatorWithManagedObjectModel.
// Creates a persistent store coordinator with the specified managed object model.
func ExampleNewPersistentStoreCoordinatorWithManagedObjectModel() {
	_ = coredata.NewPersistentStoreCoordinatorWithManagedObjectModel(
		coredata.NSManagedObjectModel{}, // model NSManagedObjectModel
	)
	// Output:
}
