// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewGraph

// ExampleNewGraphWithNodes demonstrates how to create a Graph instance using NewGraphWithNodes.
// Initializes a graph with the specified list of nodes.
func ExampleNewGraphWithNodes() {
	_ = gameplaykit.NewGraphWithNodes(
		[]gameplaykit.IGraphNode{}, // nodes []IGraphNode
	)
	// Output:
}
