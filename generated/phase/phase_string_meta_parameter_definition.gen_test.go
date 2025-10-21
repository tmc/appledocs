// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASEStringMetaParameterDefinition


// ExampleNewPHASEStringMetaParameterDefinitionWithValueIdentifier demonstrates how to create a PHASEStringMetaParameterDefinition instance using NewPHASEStringMetaParameterDefinitionWithValueIdentifier.
// Creates a specification for a named textual metaparameter with the given value.
func ExampleNewPHASEStringMetaParameterDefinitionWithValueIdentifier() {
	_ = phase.NewPHASEStringMetaParameterDefinitionWithValueIdentifier(
		"value", // value string
		"identifier", // identifier string
	)
	// Output:
}

// ExampleNewPHASEStringMetaParameterDefinitionWithValue demonstrates how to create a PHASEStringMetaParameterDefinition instance using NewPHASEStringMetaParameterDefinitionWithValue.
// Creates a specification for a textual metaparameter with the given value.
func ExampleNewPHASEStringMetaParameterDefinitionWithValue() {
	_ = phase.NewPHASEStringMetaParameterDefinitionWithValue(
		"value", // value string
	)
	// Output:
}


