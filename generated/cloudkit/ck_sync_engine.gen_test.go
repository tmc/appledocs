// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKSyncEngine

// ExampleNewCKSyncEngineWithConfiguration demonstrates how to create a CKSyncEngine instance using NewCKSyncEngineWithConfiguration.
// Creates a sync engine with the specified configuration.
func ExampleNewCKSyncEngineWithConfiguration() {
	_ = cloudkit.NewCKSyncEngineWithConfiguration(
		cloudkit.CKSyncEngineConfiguration{}, // configuration CKSyncEngineConfiguration
	)
	// Output:
}
