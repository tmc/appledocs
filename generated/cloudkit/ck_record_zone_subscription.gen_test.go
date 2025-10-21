// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKRecordZoneSubscription

// ExampleNewCKRecordZoneSubscriptionWithZoneID demonstrates how to create a CKRecordZoneSubscription instance using NewCKRecordZoneSubscriptionWithZoneID.
// Creates a subscription for all records in the specified record zone.
func ExampleNewCKRecordZoneSubscriptionWithZoneID() {
	_ = cloudkit.NewCKRecordZoneSubscriptionWithZoneID(
		cloudkit.CKRecordZoneID{}, // zoneID CKRecordZoneID
	)
	// Output:
}
