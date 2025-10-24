// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASECardioidDirectivityModelParameters

// ExampleNewPHASECardioidDirectivityModelParametersWithSubbandParameters demonstrates how to create a PHASECardioidDirectivityModelParameters instance using NewPHASECardioidDirectivityModelParametersWithSubbandParameters.
// Creates an object that directs sound in a heart-shaped curve surrounding a sound source.
func ExampleNewPHASECardioidDirectivityModelParametersWithSubbandParameters() {
	_ = phase.NewPHASECardioidDirectivityModelParametersWithSubbandParameters(
		[]phase.PHASECardioidDirectivityModelSubbandParameters{}, // subbandParameters []PHASECardioidDirectivityModelSubbandParameters
	)
	// Output:
}
