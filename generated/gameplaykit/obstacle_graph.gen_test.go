// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewObstacleGraph

// ExampleNewObstacleGraphWithObstaclesBufferRadius demonstrates how to create a ObstacleGraph instance using NewObstacleGraphWithObstaclesBufferRadius.
// Initializes a graph with the specified list of obstacles.
func ExampleNewObstacleGraphWithObstaclesBufferRadius() {
	_ = gameplaykit.NewObstacleGraphWithObstaclesBufferRadius(
		[]gameplaykit.PolygonObstacle{}, // obstacles []PolygonObstacle
		0.0, // bufferRadius float32
	)
	// Output:
}
