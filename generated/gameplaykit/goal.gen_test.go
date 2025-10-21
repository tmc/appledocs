// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewGoal

// ExampleNewGoalToFleeAgent demonstrates how to create a Goal instance using NewGoalToFleeAgent.
// Creates a goal whose effect is to move an agent away from the current position of the specified other agent.
func ExampleNewGoalToFleeAgent() {
	_ = gameplaykit.NewGoalToFleeAgent(
		gameplaykit.GKAgent{}, // agent GKAgent
	)
	// Output:
}
// ExampleNewGoalToSeekAgent demonstrates how to create a Goal instance using NewGoalToSeekAgent.
// Creates a goal whose effect is to move an agent toward the current position of the specified other agent.
func ExampleNewGoalToSeekAgent() {
	_ = gameplaykit.NewGoalToSeekAgent(
		gameplaykit.GKAgent{}, // agent GKAgent
	)
	// Output:
}
