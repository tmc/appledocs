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
		[]eventkit.IEKSource{}, // sources []IEKSource
	)
	// Output:
}
// ExampleEKEventStore_DefaultCalendarForNewReminders demonstrates using DefaultCalendarForNewReminders on a EKEventStore instance.
// Identifies the default calendar for adding reminders to, as specified by user settings.
func ExampleEKEventStore_DefaultCalendarForNewReminders() {
	obj := eventkit.NewEKEventStore()
	_ = obj.DefaultCalendarForNewReminders()
	// Output:
	}

// ExampleEKEventStore_RefreshSourcesIfNecessary demonstrates using RefreshSourcesIfNecessary on a EKEventStore instance.
// Pulls new data from remote sources, if necessary.
func ExampleEKEventStore_RefreshSourcesIfNecessary() {
	obj := eventkit.NewEKEventStore()
	obj.RefreshSourcesIfNecessary()
	// Output:
	}

// ExampleEKEventStore_Reset demonstrates using Reset on a EKEventStore instance.
// Reverts the event store to its saved state.
func ExampleEKEventStore_Reset() {
	obj := eventkit.NewEKEventStore()
	obj.Reset()
	// Output:
	}

