// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit_test

import (
	"github.com/tmc/appledocs/generated/healthkit"
)

// Suppress unused import errors
var _ = healthkit.NewHKWorkoutSession

// ExampleNewHKWorkoutSessionWithActivityTypeLocationType demonstrates how to create a HKWorkoutSession instance using NewHKWorkoutSessionWithActivityTypeLocationType.
// Returns a newly instantiated workout session.
func ExampleNewHKWorkoutSessionWithActivityTypeLocationType() {
	_ = healthkit.NewHKWorkoutSessionWithActivityTypeLocationType(
		healthkit.HKWorkoutActivityType{}, // activityType HKWorkoutActivityType
		healthkit.HKWorkoutSessionLocationType{}, // locationType HKWorkoutSessionLocationType
	)
	// Output:
}
// ExampleHKWorkoutSession_End demonstrates using End on a HKWorkoutSession instance.
// Ends the workout session.
func ExampleHKWorkoutSession_End() {
	obj := healthkit.NewHKWorkoutSession()
	obj.End()
	// Output:
	}

// ExampleHKWorkoutSession_Pause demonstrates using Pause on a HKWorkoutSession instance.
// Pauses the workout session.
func ExampleHKWorkoutSession_Pause() {
	obj := healthkit.NewHKWorkoutSession()
	obj.Pause()
	// Output:
	}

// ExampleHKWorkoutSession_Prepare demonstrates using Prepare on a HKWorkoutSession instance.
// Prepares the workout session.
//
// Note: This example is not executed because Prepare crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleHKWorkoutSession_Prepare() {
	obj := healthkit.NewHKWorkoutSession()
	obj.Prepare()
	}

// ExampleHKWorkoutSession_Resume demonstrates using Resume on a HKWorkoutSession instance.
// Resumes the workout session.
func ExampleHKWorkoutSession_Resume() {
	obj := healthkit.NewHKWorkoutSession()
	obj.Resume()
	// Output:
	}

