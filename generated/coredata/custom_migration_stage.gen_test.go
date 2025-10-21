// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewCustomMigrationStage

// ExampleNewCustomMigrationStageWithCurrentModelReferenceNextModelReference demonstrates how to create a CustomMigrationStage instance using NewCustomMigrationStageWithCurrentModelReferenceNextModelReference.
// Creates a custom migration stage with the specified source and destination model references.
func ExampleNewCustomMigrationStageWithCurrentModelReferenceNextModelReference() {
	_ = coredata.NewCustomMigrationStageWithCurrentModelReferenceNextModelReference(
		coredata.NSManagedObjectModelReference{}, // currentModel NSManagedObjectModelReference
		coredata.NSManagedObjectModelReference{}, // nextModel NSManagedObjectModelReference
	)
	// Output:
}
