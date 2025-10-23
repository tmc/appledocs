// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewDateInterval

// ExampleNewDateInterval demonstrates how to create a DateInterval instance.
// Initializes a date interval by setting the start and end date to the current date.
func ExampleNewDateInterval() {
	_ = foundation.NewDateInterval()
	// Output:
}
// ExampleNewDateIntervalWithCoder demonstrates how to create a DateInterval instance using NewDateIntervalWithCoder.
// Returns a date interval initialized from data in the given unarchiver.
func ExampleNewDateIntervalWithCoder() {
	_ = foundation.NewDateIntervalWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
// ExampleNewDateIntervalWithStartDateDuration demonstrates how to create a DateInterval instance using NewDateIntervalWithStartDateDuration.
// Initializes a date interval with a given start date and duration.
func ExampleNewDateIntervalWithStartDateDuration() {
	_ = foundation.NewDateIntervalWithStartDateDuration(
		foundation.NSDate{}, // startDate NSDate
		foundation.TimeInterval(0.0), // duration TimeInterval
	)
	// Output:
}
// ExampleNewDateIntervalWithStartDateEndDate demonstrates how to create a DateInterval instance using NewDateIntervalWithStartDateEndDate.
// Initializes a date interval from a given start date and end date.
func ExampleNewDateIntervalWithStartDateEndDate() {
	_ = foundation.NewDateIntervalWithStartDateEndDate(
		foundation.NSDate{}, // startDate NSDate
		foundation.NSDate{}, // endDate NSDate
	)
	// Output:
}
