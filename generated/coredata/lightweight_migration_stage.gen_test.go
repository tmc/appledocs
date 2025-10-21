// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewLightweightMigrationStage

// ExampleNewLightweightMigrationStageWithVersionChecksums demonstrates how to create a LightweightMigrationStage instance using NewLightweightMigrationStageWithVersionChecksums.
// Creates a lightweight migration stage with the specified version checksums.
func ExampleNewLightweightMigrationStageWithVersionChecksums() {
	_ = coredata.NewLightweightMigrationStageWithVersionChecksums(
		[]coredata.string{}, // versionChecksums []string
	)
	// Output:
}
