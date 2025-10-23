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
// ExampleNewDateWithCoder demonstrates how to create a Date instance using NewDateWithCoder.
// Returns a date object initialized from data in the given unarchiver.
func ExampleNewDateWithCoder() {
	_ = foundation.NewDateWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
// ExampleNewDateWithString demonstrates how to create a Date instance using NewDateWithString.
// Returns a date object initialized with a date and time value specified by a given string in the international string representation format.
func ExampleNewDateWithString() {
	_ = foundation.NewDateWithString(
		"description", // description string
	)
	// Output:
}
// ExampleNewDateWithTimeIntervalSince1970 demonstrates how to create a Date instance using NewDateWithTimeIntervalSince1970.
// Returns a date object initialized relative to 00:00:00 UTC on 1 January 1970 by a given number of seconds.
func ExampleNewDateWithTimeIntervalSince1970() {
	_ = foundation.NewDateWithTimeIntervalSince1970(
		foundation.TimeInterval(0.0), // secs TimeInterval
	)
	// Output:
}
// ExampleNewDateWithTimeIntervalSinceDate demonstrates how to create a Date instance using NewDateWithTimeIntervalSinceDate.
// Returns a date object initialized relative to another given date by a given number of seconds.
func ExampleNewDateWithTimeIntervalSinceDate() {
	_ = foundation.NewDateWithTimeIntervalSinceDate(
		foundation.TimeInterval(0.0), // secsToBeAdded TimeInterval
		foundation.NSDate{}, // date NSDate
	)
	// Output:
}
// ExampleNewDateWithTimeIntervalSinceNow demonstrates how to create a Date instance using NewDateWithTimeIntervalSinceNow.
// Returns a date object initialized relative to the current date and time by a given number of seconds.
func ExampleNewDateWithTimeIntervalSinceNow() {
	_ = foundation.NewDateWithTimeIntervalSinceNow(
		foundation.TimeInterval(0.0), // secs TimeInterval
	)
	// Output:
}
// ExampleNewDateWithTimeIntervalSinceReferenceDate demonstrates how to create a Date instance using NewDateWithTimeIntervalSinceReferenceDate.
// Returns a date object initialized relative to 00:00:00 UTC on 1 January 2001 by a given number of seconds.
func ExampleNewDateWithTimeIntervalSinceReferenceDate() {
	_ = foundation.NewDateWithTimeIntervalSinceReferenceDate(
		foundation.TimeInterval(0.0), // ti TimeInterval
	)
	// Output:
}
