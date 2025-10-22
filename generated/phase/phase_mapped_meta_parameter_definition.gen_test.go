// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASEMappedMetaParameterDefinition

// ExampleNewPHASEMappedMetaParameterDefinitionWithInputMetaParameterDefinitionEnvelopeIdentifier demonstrates how to create a PHASEMappedMetaParameterDefinition instance using NewPHASEMappedMetaParameterDefinitionWithInputMetaParameterDefinitionEnvelopeIdentifier.
// Creates a specification for a named metaparameter that the app plots on a graph defined by the given set of curves.
func ExampleNewPHASEMappedMetaParameterDefinitionWithInputMetaParameterDefinitionEnvelopeIdentifier() {
	_ = phase.NewPHASEMappedMetaParameterDefinitionWithInputMetaParameterDefinitionEnvelopeIdentifier(
		phase.PHASENumberMetaParameterDefinition{}, // inputMetaParameterDefinition PHASENumberMetaParameterDefinition
		phase.PHASEEnvelope{}, // envelope PHASEEnvelope
		"identifier", // identifier string
	)
	// Output:
}
