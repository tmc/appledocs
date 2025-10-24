// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewAnimation

// ExampleNewAnimationWithDurationAnimationCurve demonstrates how to create a Animation instance using NewAnimationWithDurationAnimationCurve.
// Returns an   object initialized with the specified duration and animation-curve values.
func ExampleNewAnimationWithDurationAnimationCurve() {
	_ = appkit.NewAnimationWithDurationAnimationCurve(
		0.0, // duration float64
		appkit.AnimationCurve{}, // animationCurve AnimationCurve
	)
	// Output:
}
// ExampleAnimation_ClearStartAnimation demonstrates using ClearStartAnimation on a Animation instance.
// Clears linkage to another animation that causes the receiver to start.
func ExampleAnimation_ClearStartAnimation() {
	obj := appkit.NewAnimation()
	obj.ClearStartAnimation()
	// Output:
	}

// ExampleAnimation_ClearStopAnimation demonstrates using ClearStopAnimation on a Animation instance.
// Clears linkage to another animation that causes the receiver to stop.
func ExampleAnimation_ClearStopAnimation() {
	obj := appkit.NewAnimation()
	obj.ClearStopAnimation()
	// Output:
	}

// ExampleAnimation_StartAnimation demonstrates using StartAnimation on a Animation instance.
// Starts the animation represented by the receiver.
func ExampleAnimation_StartAnimation() {
	obj := appkit.NewAnimation()
	obj.StartAnimation()
	// Output:
	}

// ExampleAnimation_StopAnimation demonstrates using StopAnimation on a Animation instance.
// Stops the animation represented by the receiver.
func ExampleAnimation_StopAnimation() {
	obj := appkit.NewAnimation()
	obj.StopAnimation()
	// Output:
	}

