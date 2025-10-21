// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASEBlendNodeDefinition

// ExampleNewPHASEBlendNodeDefinitionWithBlendMetaParameterDefinition demonstrates how to create a PHASEBlendNodeDefinition instance using NewPHASEBlendNodeDefinitionWithBlendMetaParameterDefinition.
// Creates a blend node with a maxiumum blend range value.
func ExampleNewPHASEBlendNodeDefinitionWithBlendMetaParameterDefinition() {
	_ = phase.NewPHASEBlendNodeDefinitionWithBlendMetaParameterDefinition(
		phase.PHASENumberMetaParameterDefinition{}, // blendMetaParameterDefinition PHASENumberMetaParameterDefinition
	)
	// Output:
}
