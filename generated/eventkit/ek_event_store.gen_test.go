// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit_test

import (
	"github.com/tmc/appledocs/generated/eventkit"
)

// Suppress unused import errors
var _ = eventkit.NewEKEventStore

// ExampleNewEKEventStore demonstrates how to create a EKEventStore instance.
// Creates a new event store.
func ExampleNewEKEventStore() {
	_ = eventkit.NewEKEventStore()
	// Output:
}
// ExampleNewEKEventStoreWithSources demonstrates how to create a EKEventStore instance using NewEKEventStoreWithSources.
// Creates an event store that contains data for the specified sources.
func ExampleNewEKEventStoreWithSources() {
	_ = eventkit.NewEKEventStoreWithSources(
		[]eventkit.EKSource{}, // sources []EKSource
	)
	// Output:
}
