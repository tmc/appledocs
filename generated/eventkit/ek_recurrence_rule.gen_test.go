// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit_test

import (
	"github.com/tmc/appledocs/generated/eventkit"
)

// Suppress unused import errors
var _ = eventkit.NewEKRecurrenceRule

// ExampleNewEKRecurrenceRuleRecurrenceWithFrequencyIntervalEnd demonstrates how to create a EKRecurrenceRule instance using NewEKRecurrenceRuleRecurrenceWithFrequencyIntervalEnd.
// Initializes and returns a simple recurrence rule with a given frequency, interval, and end.
func ExampleNewEKRecurrenceRuleRecurrenceWithFrequencyIntervalEnd() {
	_ = eventkit.NewEKRecurrenceRuleRecurrenceWithFrequencyIntervalEnd(
		eventkit.EKRecurrenceFrequency{}, // type EKRecurrenceFrequency
		0, // interval int
		eventkit.EKRecurrenceEnd{}, // end EKRecurrenceEnd
	)
	// Output:
}
