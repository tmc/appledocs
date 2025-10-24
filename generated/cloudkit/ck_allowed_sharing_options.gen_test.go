// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKAllowedSharingOptions

// ExampleNewCKAllowedSharingOptionsWithAllowedParticipantPermissionOptionsAllowedParticipantAccessOptions demonstrates how to create a CKAllowedSharingOptions instance using NewCKAllowedSharingOptionsWithAllowedParticipantPermissionOptionsAllowedParticipantAccessOptions.
// Creates and initializes an allowed sharing options object.
func ExampleNewCKAllowedSharingOptionsWithAllowedParticipantPermissionOptionsAllowedParticipantAccessOptions() {
	_ = cloudkit.NewCKAllowedSharingOptionsWithAllowedParticipantPermissionOptionsAllowedParticipantAccessOptions(
		cloudkit.CKSharingParticipantPermissionOption{}, // allowedParticipantPermissionOptions CKSharingParticipantPermissionOption
		cloudkit.CKSharingParticipantAccessOption{}, // allowedParticipantAccessOptions CKSharingParticipantAccessOption
	)
	// Output:
}
