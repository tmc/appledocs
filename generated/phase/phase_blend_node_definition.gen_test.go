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
// ExampleNewPHASEBlendNodeDefinitionWithBlendMetaParameterDefinitionIdentifier demonstrates how to create a PHASEBlendNodeDefinition instance using NewPHASEBlendNodeDefinitionWithBlendMetaParameterDefinitionIdentifier.
// Creates a named blend node with a maxiumum blend range value.
func ExampleNewPHASEBlendNodeDefinitionWithBlendMetaParameterDefinitionIdentifier() {
	_ = phase.NewPHASEBlendNodeDefinitionWithBlendMetaParameterDefinitionIdentifier(
		phase.PHASENumberMetaParameterDefinition{}, // blendMetaParameterDefinition PHASENumberMetaParameterDefinition
		"identifier", // identifier string
	)
	// Output:
}
