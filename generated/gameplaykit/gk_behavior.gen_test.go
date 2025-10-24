// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewBehavior

// ExampleNewBehaviorWithGoals demonstrates how to create a Behavior instance using NewBehaviorWithGoals.
// Creates a behavior with the specified goals.
func ExampleNewBehaviorWithGoals() {
	_ = gameplaykit.NewBehaviorWithGoals(
		[]gameplaykit.IGoal{}, // goals []IGoal
	)
	// Output:
}
// ExampleBehavior_RemoveAllGoals demonstrates using RemoveAllGoals on a Behavior instance.
// Removes all goals from the behavior.
func ExampleBehavior_RemoveAllGoals() {
	obj := gameplaykit.NewBehavior()
	obj.RemoveAllGoals()
	// Output:
	}

