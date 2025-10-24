// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASEConeDirectivityModelParameters

// ExampleNewPHASEConeDirectivityModelParametersWithSubbandParameters demonstrates how to create a PHASEConeDirectivityModelParameters instance using NewPHASEConeDirectivityModelParametersWithSubbandParameters.
// Creates an object that directs sound in a cone-shaped curve that extends from a sound source.
func ExampleNewPHASEConeDirectivityModelParametersWithSubbandParameters() {
	_ = phase.NewPHASEConeDirectivityModelParametersWithSubbandParameters(
		[]phase.PHASEConeDirectivityModelSubbandParameters{}, // subbandParameters []PHASEConeDirectivityModelSubbandParameters
	)
	// Output:
}
