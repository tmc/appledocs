// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit_test

import (
	"github.com/tmc/appledocs/generated/eventkit"
)

// Suppress unused import errors
var _ = eventkit.NewEKAlarm

// ExampleNewEKAlarmWithRelativeOffset demonstrates how to create a EKAlarm instance using NewEKAlarmWithRelativeOffset.
// Creates and returns an alarm with a relative offset.
func ExampleNewEKAlarmWithRelativeOffset() {
	_ = eventkit.NewEKAlarmWithRelativeOffset(
		0.0, // offset float64
	)
	// Output:
}
