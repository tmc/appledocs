// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit_test

import (
	"github.com/tmc/appledocs/generated/eventkit"
)

// Suppress unused import errors
var _ = eventkit.NewEKRecurrenceEnd



// ExampleNewEKRecurrenceEndWithOccurrenceCount demonstrates how to create a EKRecurrenceEnd instance using NewEKRecurrenceEndWithOccurrenceCount.
// Initializes and returns a count-based recurrence end with a given maximum occurrence count.
func ExampleNewEKRecurrenceEndWithOccurrenceCount() {
	_ = eventkit.NewEKRecurrenceEndWithOccurrenceCount(
		0, // occurrenceCount uint
	)
	// Output:
}


