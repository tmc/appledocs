// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit_test

import (
	"github.com/tmc/appledocs/generated/eventkit"
)

// Suppress unused import errors
var _ = eventkit.NewEKReminder

// ExampleNewEKReminderWithEventStore demonstrates how to create a EKReminder instance using NewEKReminderWithEventStore.
// Creates and returns a new reminder in the given event store.
func ExampleNewEKReminderWithEventStore() {
	_ = eventkit.NewEKReminderWithEventStore(
		eventkit.EKEventStore{}, // eventStore EKEventStore
	)
	// Output:
}
