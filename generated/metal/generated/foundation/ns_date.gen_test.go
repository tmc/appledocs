// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewDate

// ExampleNewDate demonstrates how to create a Date instance.
// Returns a date object initialized to the current date and time.
func ExampleNewDate() {
	_ = foundation.NewDate()
	// Output:
}
// ExampleNewDateWithTimeIntervalSince1970 demonstrates how to create a Date instance using NewDateWithTimeIntervalSince1970.
// Returns a date object initialized relative to 00:00:00 UTC on 1 January 1970 by a given number of seconds.
func ExampleNewDateWithTimeIntervalSince1970() {
	_ = foundation.NewDateWithTimeIntervalSince1970(
		0.0, // secs float64
	)
	// Output:
}
// ExampleNewDateWithTimeIntervalSinceNow demonstrates how to create a Date instance using NewDateWithTimeIntervalSinceNow.
// Returns a date object initialized relative to the current date and time by a given number of seconds.
func ExampleNewDateWithTimeIntervalSinceNow() {
	_ = foundation.NewDateWithTimeIntervalSinceNow(
		0.0, // secs float64
	)
	// Output:
}
// ExampleNewDateWithTimeIntervalSinceReferenceDate demonstrates how to create a Date instance using NewDateWithTimeIntervalSinceReferenceDate.
// Returns a date object initialized relative to 00:00:00 UTC on 1 January 2001 by a given number of seconds.
func ExampleNewDateWithTimeIntervalSinceReferenceDate() {
	_ = foundation.NewDateWithTimeIntervalSinceReferenceDate(
		0.0, // ti float64
	)
	// Output:
}
