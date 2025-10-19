// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)


// ExampleNewCustomMigrationStageWithCurrentModelReferenceNextModelReference demonstrates how to create a CustomMigrationStage instance using NewCustomMigrationStageWithCurrentModelReferenceNextModelReference.
// Creates a custom migration stage with the specified source and destination model references.
func ExampleNewCustomMigrationStageWithCurrentModelReferenceNextModelReference() {
	_ = coredata.NewCustomMigrationStageWithCurrentModelReferenceNextModelReference(
		nil, // currentModel unsafe.Pointer
		nil, // nextModel unsafe.Pointer
	)
	// Output:
}


