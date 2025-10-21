// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASESwitchNodeDefinition

// ExampleNewPHASESwitchNodeDefinitionWithSwitchMetaParameterDefinition demonstrates how to create a PHASESwitchNodeDefinition instance using NewPHASESwitchNodeDefinitionWithSwitchMetaParameterDefinition.
// Creates a node that invokes a child node based on the value of the given parameter.
func ExampleNewPHASESwitchNodeDefinitionWithSwitchMetaParameterDefinition() {
	_ = phase.NewPHASESwitchNodeDefinitionWithSwitchMetaParameterDefinition(
		phase.PHASEStringMetaParameterDefinition{}, // switchMetaParameterDefinition PHASEStringMetaParameterDefinition
	)
	// Output:
}

