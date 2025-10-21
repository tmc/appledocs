// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit_test

import (
	"github.com/tmc/appledocs/generated/eventkit"
)

// Suppress unused import errors
var _ = eventkit.NewEKEvent

// ExampleNewEKEventWithEventStore demonstrates how to create a EKEvent instance using NewEKEventWithEventStore.
// Creates and returns a new event belonging to a specified event store.
func ExampleNewEKEventWithEventStore() {
	_ = eventkit.NewEKEventWithEventStore(
		eventkit.EKEventStore{}, // eventStore EKEventStore
	)
	// Output:
}
