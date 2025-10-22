// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASERandomNodeDefinition

// ExampleNewPHASERandomNodeDefinition demonstrates how to create a PHASERandomNodeDefinition instance.
// Creates a random node.
func ExampleNewPHASERandomNodeDefinition() {
	_ = phase.NewPHASERandomNodeDefinition()
	// Output:
}
// ExampleNewPHASERandomNodeDefinitionWithIdentifier demonstrates how to create a PHASERandomNodeDefinition instance using NewPHASERandomNodeDefinitionWithIdentifier.
// Creates a random node with the name you specify.
func ExampleNewPHASERandomNodeDefinitionWithIdentifier() {
	_ = phase.NewPHASERandomNodeDefinitionWithIdentifier(
		"identifier", // identifier string
	)
	// Output:
}
