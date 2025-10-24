// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewBehavior

// ExampleNewBehaviorWithGoalWeight demonstrates how to create a Behavior instance using NewBehaviorWithGoalWeight.
// Creates a behavior with a single goal.
func ExampleNewBehaviorWithGoalWeight() {
	_ = gameplaykit.NewBehaviorWithGoalWeight(
		gameplaykit.GKGoal{}, // goal GKGoal
		0.0,                  // weight float32
	)
	// Output:
}

// ExampleNewBehaviorWithGoals demonstrates how to create a Behavior instance using NewBehaviorWithGoals.
// Creates a behavior with the specified goals.
func ExampleNewBehaviorWithGoals() {
	_ = gameplaykit.NewBehaviorWithGoals(
		[]gameplaykit.Goal{}, // goals []Goal
	)
	// Output:
}
