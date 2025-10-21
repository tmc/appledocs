// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKFetchShareParticipantsOperation

// ExampleNewCKFetchShareParticipantsOperation demonstrates how to create a CKFetchShareParticipantsOperation instance.
// Creates an empty operation.
func ExampleNewCKFetchShareParticipantsOperation() {
	_ = cloudkit.NewCKFetchShareParticipantsOperation()
	// Output:
}
// ExampleNewCKFetchShareParticipantsOperationWithUserIdentityLookupInfos demonstrates how to create a CKFetchShareParticipantsOperation instance using NewCKFetchShareParticipantsOperationWithUserIdentityLookupInfos.
// Creates an operation for generating share participants from the specified user data.
func ExampleNewCKFetchShareParticipantsOperationWithUserIdentityLookupInfos() {
	_ = cloudkit.NewCKFetchShareParticipantsOperationWithUserIdentityLookupInfos(
		[]cloudkit.CKUserIdentityLookupInfo{}, // userIdentityLookupInfos []CKUserIdentityLookupInfo
	)
	// Output:
}
