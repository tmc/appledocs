// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewMeshGraph

// ExampleMeshGraph_Triangulate demonstrates using Triangulate on a MeshGraph instance.
// Creates or updates the graph with a network of nodes that describes the open space around its obstacles.
func ExampleMeshGraph_Triangulate() {
	obj := gameplaykit.NewMeshGraph()
	obj.Triangulate()
	// Output:
	}

