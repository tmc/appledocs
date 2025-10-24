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
		[]gameplaykit.IAgent{}, // agents []IAgent
		0.0, // maxDistance float32
		0.0, // maxAngle float32
	)
	// Output:
}
// ExampleNewGoalToAvoidAgentsMaxPredictionTime demonstrates how to create a Goal instance using NewGoalToAvoidAgentsMaxPredictionTime.
// Creates a goal whose effect is to make an agent avoid colliding with the specified other agents, taking into account the other agents’ movement.
func ExampleNewGoalToAvoidAgentsMaxPredictionTime() {
	_ = gameplaykit.NewGoalToAvoidAgentsMaxPredictionTime(
		[]gameplaykit.IAgent{}, // agents []IAgent
		0.0, // maxPredictionTime float64
	)
	// Output:
}
// ExampleNewGoalToAvoidObstaclesMaxPredictionTime demonstrates how to create a Goal instance using NewGoalToAvoidObstaclesMaxPredictionTime.
// Creates a goal whose effect is to make an agent avoid colliding with the specified static obstacles.
func ExampleNewGoalToAvoidObstaclesMaxPredictionTime() {
	_ = gameplaykit.NewGoalToAvoidObstaclesMaxPredictionTime(
		[]gameplaykit.IObstacle{}, // obstacles []IObstacle
		0.0, // maxPredictionTime float64
	)
	// Output:
}
// ExampleNewGoalToCohereWithAgentsMaxDistanceMaxAngle demonstrates how to create a Goal instance using NewGoalToCohereWithAgentsMaxDistanceMaxAngle.
// Creates a goal whose effect is to make an agent stay near the other agents in a specified group.
func ExampleNewGoalToCohereWithAgentsMaxDistanceMaxAngle() {
	_ = gameplaykit.NewGoalToCohereWithAgentsMaxDistanceMaxAngle(
		[]gameplaykit.IAgent{}, // agents []IAgent
		0.0, // maxDistance float32
		0.0, // maxAngle float32
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
// ExampleNewGoalToSeparateFromAgentsMaxDistanceMaxAngle demonstrates how to create a Goal instance using NewGoalToSeparateFromAgentsMaxDistanceMaxAngle.
// Creates a goal whose effect is to make an agent maintain the specified distance from other agents in a specified group.
func ExampleNewGoalToSeparateFromAgentsMaxDistanceMaxAngle() {
	_ = gameplaykit.NewGoalToSeparateFromAgentsMaxDistanceMaxAngle(
		[]gameplaykit.IAgent{}, // agents []IAgent
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
