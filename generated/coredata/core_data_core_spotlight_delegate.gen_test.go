// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewCoreDataCoreSpotlightDelegate

// ExampleNewCoreDataCoreSpotlightDelegateForStoreWithDescriptionCoordinator demonstrates how to create a CoreDataCoreSpotlightDelegate instance using NewCoreDataCoreSpotlightDelegateForStoreWithDescriptionCoordinator.
// Creates a Core Spotlight delegate with the specified store description and coordinator.
func ExampleNewCoreDataCoreSpotlightDelegateForStoreWithDescriptionCoordinator() {
	_ = coredata.NewCoreDataCoreSpotlightDelegateForStoreWithDescriptionCoordinator(
		coredata.NSPersistentStoreDescription{}, // description NSPersistentStoreDescription
		coredata.NSPersistentStoreCoordinator{}, // psc NSPersistentStoreCoordinator
	)
	// Output:
}
// ExampleNewCoreDataCoreSpotlightDelegateForStoreWithDescriptionModel demonstrates how to create a CoreDataCoreSpotlightDelegate instance using NewCoreDataCoreSpotlightDelegateForStoreWithDescriptionModel.
// Creates a Core Spotlight delegate with the specified store description and managed object model.
func ExampleNewCoreDataCoreSpotlightDelegateForStoreWithDescriptionModel() {
	_ = coredata.NewCoreDataCoreSpotlightDelegateForStoreWithDescriptionModel(
		coredata.NSPersistentStoreDescription{}, // description NSPersistentStoreDescription
		coredata.NSManagedObjectModel{}, // model NSManagedObjectModel
	)
	// Output:
}
