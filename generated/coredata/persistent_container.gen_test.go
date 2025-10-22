// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewPersistentContainer

// ExampleNewPersistentContainerWithName demonstrates how to create a PersistentContainer instance using NewPersistentContainerWithName.
// Creates a container with the specified name.
func ExampleNewPersistentContainerWithName() {
	_ = coredata.NewPersistentContainerWithName(
		"name", // name string
	)
	// Output:
}
// ExampleNewPersistentContainerWithNameManagedObjectModel demonstrates how to create a PersistentContainer instance using NewPersistentContainerWithNameManagedObjectModel.
// Create a container with the specified name and managed object model.
func ExampleNewPersistentContainerWithNameManagedObjectModel() {
	_ = coredata.NewPersistentContainerWithNameManagedObjectModel(
		"name", // name string
		coredata.NSManagedObjectModel{}, // model NSManagedObjectModel
	)
	// Output:
}
