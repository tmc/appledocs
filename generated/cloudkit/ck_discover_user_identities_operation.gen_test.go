// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKDiscoverUserIdentitiesOperation

// ExampleNewCKDiscoverUserIdentitiesOperation demonstrates how to create a CKDiscoverUserIdentitiesOperation instance.
// Creates an operation for discovering user identities.
func ExampleNewCKDiscoverUserIdentitiesOperation() {
	_ = cloudkit.NewCKDiscoverUserIdentitiesOperation()
	// Output:
}
// ExampleNewCKDiscoverUserIdentitiesOperationWithUserIdentityLookupInfos demonstrates how to create a CKDiscoverUserIdentitiesOperation instance using NewCKDiscoverUserIdentitiesOperationWithUserIdentityLookupInfos.
// Creates an operation for discovering the user identities of the specified lookup infos.
func ExampleNewCKDiscoverUserIdentitiesOperationWithUserIdentityLookupInfos() {
	_ = cloudkit.NewCKDiscoverUserIdentitiesOperationWithUserIdentityLookupInfos(
		[]cloudkit.CKUserIdentityLookupInfo{}, // userIdentityLookupInfos []CKUserIdentityLookupInfo
	)
	// Output:
}
