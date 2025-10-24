// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit_test

import (
	"github.com/tmc/appledocs/generated/healthkit"
)

// Suppress unused import errors
var _ = healthkit.NewHKWorkoutBuilder

// ExampleHKWorkoutBuilder_DiscardWorkout demonstrates using DiscardWorkout on a HKWorkoutBuilder instance.
// Stops the collection of data and discards the current results without saving the workout.
//
// Note: This example is not executed because DiscardWorkout crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleHKWorkoutBuilder_DiscardWorkout() {
	obj := healthkit.NewHKWorkoutBuilder()
	obj.DiscardWorkout()
	}

