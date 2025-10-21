// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents_test

import (
	"github.com/tmc/appledocs/generated/intents"
)

// Suppress unused import errors
var _ = intents.NewINEndWorkoutIntent

// ExampleNewINEndWorkoutIntentWithWorkoutName demonstrates how to create a INEndWorkoutIntent instance using NewINEndWorkoutIntentWithWorkoutName.
// Initializes an intent object with the specified workout name.
func ExampleNewINEndWorkoutIntentWithWorkoutName() {
	_ = intents.NewINEndWorkoutIntentWithWorkoutName(
		intents.INSpeakableString{}, // workoutName INSpeakableString
	)
	// Output:
}
