// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewPath

// ExampleNewPathWithGraphNodesRadius demonstrates how to create a Path instance using NewPathWithGraphNodesRadius.
// Initializes a path using the positions of the specified graph nodes.
func ExampleNewPathWithGraphNodesRadius() {
	_ = gameplaykit.NewPathWithGraphNodesRadius(
		[]gameplaykit.IGraphNode{}, // graphNodes []IGraphNode
		0.0, // radius float32
	)
	// Output:
}
