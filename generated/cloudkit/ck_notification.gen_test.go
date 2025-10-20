// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKNotification


// ExampleNewCKNotificationFromRemoteNotificationDictionary demonstrates how to create a CKNotification instance using NewCKNotificationFromRemoteNotificationDictionary.
// Creates a new notification using the specified payload data.
func ExampleNewCKNotificationFromRemoteNotificationDictionary() {
	_ = cloudkit.NewCKNotificationFromRemoteNotificationDictionary(
		0, // notificationDictionary objc.ID
	)
	// Output:
}


