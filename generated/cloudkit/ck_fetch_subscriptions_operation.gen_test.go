// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKFetchSubscriptionsOperation

// ExampleNewCKFetchSubscriptionsOperation demonstrates how to create a CKFetchSubscriptionsOperation instance.
// Creates an empty fetch subscriptions operation.
func ExampleNewCKFetchSubscriptionsOperation() {
	_ = cloudkit.NewCKFetchSubscriptionsOperation()
	// Output:
}
// ExampleNewCKFetchSubscriptionsOperationWithSubscriptionIDs demonstrates how to create a CKFetchSubscriptionsOperation instance using NewCKFetchSubscriptionsOperationWithSubscriptionIDs.
// Creates an operation for fetching the specified subscriptions.
func ExampleNewCKFetchSubscriptionsOperationWithSubscriptionIDs() {
	_ = cloudkit.NewCKFetchSubscriptionsOperationWithSubscriptionIDs(
		[]cloudkit.string{}, // subscriptionIDs []string
	)
	// Output:
}
