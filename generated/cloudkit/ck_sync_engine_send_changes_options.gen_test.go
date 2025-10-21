// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKSyncEngineSendChangesOptions

// ExampleNewCKSyncEngineSendChangesOptionsWithScope demonstrates how to create a CKSyncEngineSendChangesOptions instance using NewCKSyncEngineSendChangesOptionsWithScope.
func ExampleNewCKSyncEngineSendChangesOptionsWithScope() {
	_ = cloudkit.NewCKSyncEngineSendChangesOptionsWithScope(
		cloudkit.CKSyncEngineSendChangesScope{}, // scope CKSyncEngineSendChangesScope
	)
	// Output:
}
