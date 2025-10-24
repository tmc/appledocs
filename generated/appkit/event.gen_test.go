// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewEvent

// ExampleNewEventWithCGEvent demonstrates how to create a Event instance using NewEventWithCGEvent.
// Creates and returns an event object for a Core Graphics event.
func ExampleNewEventWithCGEvent() {
	_ = appkit.NewEventWithCGEvent(
		appkit.EventRef /* not a class type */ {}, // cgEvent EventRef /* not a class type */
	)
	// Output:
}
