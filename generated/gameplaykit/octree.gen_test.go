// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewOctree

// ExampleNewOctreeWithBoundingBoxMinimumCellSize demonstrates how to create a Octree instance using NewOctreeWithBoundingBoxMinimumCellSize.
// Initializes an octree with the specified dimensions.
func ExampleNewOctreeWithBoundingBoxMinimumCellSize() {
	_ = gameplaykit.NewOctreeWithBoundingBoxMinimumCellSize(
		gameplaykit.Box{}, // box Box
		0.0, // minCellSize float32
	)
	// Output:
}
