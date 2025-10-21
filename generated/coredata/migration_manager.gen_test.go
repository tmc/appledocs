// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewMigrationManager

// ExampleNewMigrationManagerWithSourceModelDestinationModel demonstrates how to create a MigrationManager instance using NewMigrationManagerWithSourceModelDestinationModel.
// Initializes a migration manager instance with given source and destination models.
func ExampleNewMigrationManagerWithSourceModelDestinationModel() {
	_ = coredata.NewMigrationManagerWithSourceModelDestinationModel(
		coredata.NSManagedObjectModel{}, // sourceModel NSManagedObjectModel
		coredata.NSManagedObjectModel{}, // destinationModel NSManagedObjectModel
	)
	// Output:
}
