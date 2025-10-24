// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKModifySubscriptionsOperation

// ExampleNewCKModifySubscriptionsOperation demonstrates how to create a CKModifySubscriptionsOperation instance.
// Creates an empty modify subscriptions operation.
func ExampleNewCKModifySubscriptionsOperation() {
	_ = cloudkit.NewCKModifySubscriptionsOperation()
	// Output:
}
// ExampleNewCKModifySubscriptionsOperationWithSubscriptionsToSaveSubscriptionIDsToDelete demonstrates how to create a CKModifySubscriptionsOperation instance using NewCKModifySubscriptionsOperationWithSubscriptionsToSaveSubscriptionIDsToDelete.
// Creates an operation for saving and deleting the specified subscriptions.
func ExampleNewCKModifySubscriptionsOperationWithSubscriptionsToSaveSubscriptionIDsToDelete() {
	_ = cloudkit.NewCKModifySubscriptionsOperationWithSubscriptionsToSaveSubscriptionIDsToDelete(
		[]cloudkit.CKSubscription{}, // subscriptionsToSave []CKSubscription
		[]cloudkit.string{}, // subscriptionIDsToDelete []string
	)
	// Output:
}
