// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents_test

import (
	"github.com/tmc/appledocs/generated/intents"
)

// Suppress unused import errors
var _ = intents.NewINShareFocusStatusIntent

// ExampleNewINShareFocusStatusIntentWithFocusStatus demonstrates how to create a INShareFocusStatusIntent instance using NewINShareFocusStatusIntentWithFocusStatus.
// Creates an intent with the specified focus status.
func ExampleNewINShareFocusStatusIntentWithFocusStatus() {
	_ = intents.NewINShareFocusStatusIntentWithFocusStatus(
		intents.INFocusStatus{}, // focusStatus INFocusStatus
	)
	// Output:
}
