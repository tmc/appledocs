// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit_test

import (
	"github.com/tmc/appledocs/generated/eventkit"
)

// Suppress unused import errors
var _ = eventkit.NewEKEvent

// ExampleEKEvent_Refresh demonstrates using Refresh on a EKEvent instance.
// Updates the event’s data with the current information in the Calendar database.
func ExampleEKEvent_Refresh() {
	obj := eventkit.NewEKEvent()
	_ = obj.Refresh()
	// Output:
	}

