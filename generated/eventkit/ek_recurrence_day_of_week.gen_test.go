// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit_test

import (
	"github.com/tmc/appledocs/generated/eventkit"
)

// Suppress unused import errors
var _ = eventkit.NewEKRecurrenceDayOfWeek

// ExampleNewEKRecurrenceDayOfWeek demonstrates how to create a EKRecurrenceDayOfWeek instance using NewEKRecurrenceDayOfWeek.
// Creates and returns a day of the week with a given day.
func ExampleNewEKRecurrenceDayOfWeek() {
	_ = eventkit.NewEKRecurrenceDayOfWeek(
		eventkit.EKWeekday{}, // dayOfTheWeek EKWeekday
	)
	// Output:
}
// ExampleNewEKRecurrenceDayOfWeekWeekNumber demonstrates how to create a EKRecurrenceDayOfWeek instance using NewEKRecurrenceDayOfWeekWeekNumber.
// Creates and returns an autoreleased day of the week with a given day and week number.
func ExampleNewEKRecurrenceDayOfWeekWeekNumber() {
	_ = eventkit.NewEKRecurrenceDayOfWeekWeekNumber(
		eventkit.EKWeekday{}, // dayOfTheWeek EKWeekday
		0, // weekNumber int
	)
	// Output:
}
// ExampleNewEKRecurrenceDayOfWeekWithDayOfTheWeekWeekNumber demonstrates how to create a EKRecurrenceDayOfWeek instance using NewEKRecurrenceDayOfWeekWithDayOfTheWeekWeekNumber.
// Initializes and returns a day of the week with a given day and week number.
func ExampleNewEKRecurrenceDayOfWeekWithDayOfTheWeekWeekNumber() {
	_ = eventkit.NewEKRecurrenceDayOfWeekWithDayOfTheWeekWeekNumber(
		eventkit.EKWeekday{}, // dayOfTheWeek EKWeekday
		0, // weekNumber int
	)
	// Output:
}
