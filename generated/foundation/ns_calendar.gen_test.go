// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewCalendar

// ExampleNewCalendarWithCalendarIdentifier demonstrates how to create a Calendar instance using NewCalendarWithCalendarIdentifier.
// Initializes a calendar according to a given identifier.
func ExampleNewCalendarWithCalendarIdentifier() {
	_ = foundation.NewCalendarWithCalendarIdentifier(
		foundation.CalendarIdentifier{}, // ident CalendarIdentifier
	)
	// Output:
}
// ExampleNewCalendarWithIdentifier demonstrates how to create a Calendar instance using NewCalendarWithIdentifier.
// Creates a new calendar specified by a given identifier.
func ExampleNewCalendarWithIdentifier() {
	_ = foundation.NewCalendarWithIdentifier(
		foundation.CalendarIdentifier{}, // calendarIdentifierConstant CalendarIdentifier
	)
	// Output:
}
