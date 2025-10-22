// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewGoal

// ExampleNewGoalToAlignWithAgentsMaxDistanceMaxAngle demonstrates how to create a Goal instance using NewGoalToAlignWithAgentsMaxDistanceMaxAngle.
// Creates a goal whose effect is to make an agent align its orientation with that of other agents in a specified group.
func ExampleNewGoalToAlignWithAgentsMaxDistanceMaxAngle() {
	_ = gameplaykit.NewGoalToAlignWithAgentsMaxDistanceMaxAngle(
		[]gameplaykit.Agent{}, // agents []Agent
		0.0, // maxDistance float32
		0.0, // maxAngle float32
	)
	// Output:
}
// ExampleNewGoalToCohereWithAgentsMaxDistanceMaxAngle demonstrates how to create a Goal instance using NewGoalToCohereWithAgentsMaxDistanceMaxAngle.
// Creates a goal whose effect is to make an agent stay near the other agents in a specified group.
func ExampleNewGoalToCohereWithAgentsMaxDistanceMaxAngle() {
	_ = gameplaykit.NewGoalToCohereWithAgentsMaxDistanceMaxAngle(
		[]gameplaykit.Agent{}, // agents []Agent
		0.0, // maxDistance float32
		0.0, // maxAngle float32
	)
	// Output:
}
// ExampleNewGoalToFleeAgent demonstrates how to create a Goal instance using NewGoalToFleeAgent.
// Creates a goal whose effect is to move an agent away from the current position of the specified other agent.
func ExampleNewGoalToFleeAgent() {
	_ = gameplaykit.NewGoalToFleeAgent(
		gameplaykit.GKAgent{}, // agent GKAgent
	)
	// Output:
}
// ExampleNewGoalToReachTargetSpeed demonstrates how to create a Goal instance using NewGoalToReachTargetSpeed.
// Creates a goal whose effect is to accelerate or decelerate an agent until it reaches the specified speed.
func ExampleNewGoalToReachTargetSpeed() {
	_ = gameplaykit.NewGoalToReachTargetSpeed(
		0.0, // targetSpeed float32
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
// ExampleNewGoalToSeparateFromAgentsMaxDistanceMaxAngle demonstrates how to create a Goal instance using NewGoalToSeparateFromAgentsMaxDistanceMaxAngle.
// Creates a goal whose effect is to make an agent maintain the specified distance from other agents in a specified group.
func ExampleNewGoalToSeparateFromAgentsMaxDistanceMaxAngle() {
	_ = gameplaykit.NewGoalToSeparateFromAgentsMaxDistanceMaxAngle(
		[]gameplaykit.Agent{}, // agents []Agent
		0.0, // maxDistance float32
		0.0, // maxAngle float32
	)
	// Output:
}
// ExampleNewGoalToWander demonstrates how to create a Goal instance using NewGoalToWander.
// Creates a goal whose effect is to make an agent wander aimlessly, moving forward and turning at random.
func ExampleNewGoalToWander() {
	_ = gameplaykit.NewGoalToWander(
		0.0, // speed float32
	)
	// Output:
}
