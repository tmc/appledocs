// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewDateInterval

// ExampleNewDateIntervalWithStartDateDuration demonstrates how to create a DateInterval instance using NewDateIntervalWithStartDateDuration.
// Initializes a date interval with a given start date and duration.
func ExampleNewDateIntervalWithStartDateDuration() {
	_ = foundation.NewDateIntervalWithStartDateDuration(
		foundation.NSDate{}, // startDate NSDate
		foundation.TimeInterval(0.0), // duration TimeInterval
	)
	// Output:
}
