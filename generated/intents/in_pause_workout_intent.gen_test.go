// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents_test

import (
	"github.com/tmc/appledocs/generated/intents"
)

// Suppress unused import errors
var _ = intents.NewINPauseWorkoutIntent

// ExampleNewINPauseWorkoutIntentWithWorkoutName demonstrates how to create a INPauseWorkoutIntent instance using NewINPauseWorkoutIntentWithWorkoutName.
// Initializes an intent object with the specified workout name.
func ExampleNewINPauseWorkoutIntentWithWorkoutName() {
	_ = intents.NewINPauseWorkoutIntentWithWorkoutName(
		intents.INSpeakableString{}, // workoutName INSpeakableString
	)
	// Output:
}
