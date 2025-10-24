// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASENumberMetaParameterDefinition

// ExampleNewPHASENumberMetaParameterDefinitionWithValue demonstrates how to create a PHASENumberMetaParameterDefinition instance using NewPHASENumberMetaParameterDefinitionWithValue.
// Creates a specification for a metaparameter with the given numeric value.
func ExampleNewPHASENumberMetaParameterDefinitionWithValue() {
	_ = phase.NewPHASENumberMetaParameterDefinitionWithValue(
		0.0, // value float64
	)
	// Output:
}
// ExampleNewPHASENumberMetaParameterDefinitionWithValueMinimumMaximum demonstrates how to create a PHASENumberMetaParameterDefinition instance using NewPHASENumberMetaParameterDefinitionWithValueMinimumMaximum.
// Creates a specification for a metaparameter with the given numeric value and range.
func ExampleNewPHASENumberMetaParameterDefinitionWithValueMinimumMaximum() {
	_ = phase.NewPHASENumberMetaParameterDefinitionWithValueMinimumMaximum(
		0.0, // value float64
		0.0, // minimum float64
		0.0, // maximum float64
	)
	// Output:
}
