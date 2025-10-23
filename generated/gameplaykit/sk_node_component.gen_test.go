// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewSKNodeComponent

// ExampleNewSKNodeComponentWithNode demonstrates how to create a SKNodeComponent instance using NewSKNodeComponentWithNode.
// Initializes a component to manage the specified SpriteKit node.
func ExampleNewSKNodeComponentWithNode() {
	_ = gameplaykit.NewSKNodeComponentWithNode(
		gameplaykit.Node{}, // node Node
	)
	// Output:
}
