// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewQuadtree

// ExampleNewQuadtreeWithBoundingQuadMinimumCellSize demonstrates how to create a Quadtree instance using NewQuadtreeWithBoundingQuadMinimumCellSize.
// Initializes a quadtree with the specified dimensions.
func ExampleNewQuadtreeWithBoundingQuadMinimumCellSize() {
	_ = gameplaykit.NewQuadtreeWithBoundingQuadMinimumCellSize(
		gameplaykit.Quad{}, // quad Quad
		0.0,                // minCellSize float32
	)
	// Output:
}
