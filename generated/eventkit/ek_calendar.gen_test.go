// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit_test

import (
	"github.com/tmc/appledocs/generated/eventkit"
)

// Suppress unused import errors
var _ = eventkit.NewEKCalendar

// ExampleNewEKCalendarForEntityTypeEventStore demonstrates how to create a EKCalendar instance using NewEKCalendarForEntityTypeEventStore.
// Creates a new calendar that can contain the given entity type.
func ExampleNewEKCalendarForEntityTypeEventStore() {
	_ = eventkit.NewEKCalendarForEntityTypeEventStore(
		eventkit.EKEntityType{}, // entityType EKEntityType
		eventkit.EKEventStore{}, // eventStore EKEventStore
	)
	// Output:
}
// ExampleNewEKCalendarWithEventStore demonstrates how to create a EKCalendar instance using NewEKCalendarWithEventStore.
// Creates and returns a calendar belonging to a specified event store.
func ExampleNewEKCalendarWithEventStore() {
	_ = eventkit.NewEKCalendarWithEventStore(
		eventkit.EKEventStore{}, // eventStore EKEventStore
	)
	// Output:
}
