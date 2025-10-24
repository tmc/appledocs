// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata_test

import (
	"github.com/tmc/appledocs/generated/coredata"
)

// Suppress unused import errors
var _ = coredata.NewPersistentCloudKitContainer

// ExamplePersistentCloudKitContainer_AcceptShareInvitations demonstrates using AcceptShareInvitations on a PersistentCloudKitContainer instance.
// Accepts one or more invitations to participate in sharing using the specified metadata.
func ExamplePersistentCloudKitContainer_AcceptShareInvitations() {
	obj := coredata.NewPersistentCloudKitContainer()
	obj.AcceptShareInvitations()
	// Output:
}

// ExamplePersistentCloudKitContainer_FetchParticipants demonstrates using FetchParticipants on a PersistentCloudKitContainer instance.
// Fetches all participants that match the specified critieria.
func ExamplePersistentCloudKitContainer_FetchParticipants() {
	obj := coredata.NewPersistentCloudKitContainer()
	obj.FetchParticipants()
	// Output:
}

// ExamplePersistentCloudKitContainer_PersistUpdatedShare demonstrates using PersistUpdatedShare on a PersistentCloudKitContainer instance.
// Saves the share record and schedules it for export to iCloud.
func ExamplePersistentCloudKitContainer_PersistUpdatedShare() {
	obj := coredata.NewPersistentCloudKitContainer()
	obj.PersistUpdatedShare()
	// Output:
}

// ExamplePersistentCloudKitContainer_Share demonstrates using Share on a PersistentCloudKitContainer instance.
// Associates the specified managed objects with a new or existing share record.
func ExamplePersistentCloudKitContainer_Share() {
	obj := coredata.NewPersistentCloudKitContainer()
	obj.Share()
	// Output:
}
