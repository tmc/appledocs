// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewTimeZone

// ExampleNewTimeZoneForSecondsFromGMT demonstrates how to create a TimeZone instance using NewTimeZoneForSecondsFromGMT.
// Returns a time zone object offset from Greenwich Mean Time by a given number of seconds.
func ExampleNewTimeZoneForSecondsFromGMT() {
	_ = foundation.NewTimeZoneForSecondsFromGMT(
		0, // seconds int
	)
	// Output:
}

// ExampleNewTimeZoneWithAbbreviation demonstrates how to create a TimeZone instance using NewTimeZoneWithAbbreviation.
// Returns the time zone object identified by a given abbreviation.
func ExampleNewTimeZoneWithAbbreviation() {
	_ = foundation.NewTimeZoneWithAbbreviation(
		"abbreviation", // abbreviation string
	)
	// Output:
}

// ExampleNewTimeZoneWithName demonstrates how to create a TimeZone instance using NewTimeZoneWithName.
// Returns a time zone initialized with a given identifier.
func ExampleNewTimeZoneWithName() {
	_ = foundation.NewTimeZoneWithName(
		"tzName", // tzName string
	)
	// Output:
}




