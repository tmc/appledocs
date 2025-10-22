// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASEDistanceModelFadeOutParameters

// ExampleNewPHASEDistanceModelFadeOutParametersWithCullDistance demonstrates how to create a PHASEDistanceModelFadeOutParameters instance using NewPHASEDistanceModelFadeOutParametersWithCullDistance.
// Creates a distance beyond which sound sources stop playing.
func ExampleNewPHASEDistanceModelFadeOutParametersWithCullDistance() {
	_ = phase.NewPHASEDistanceModelFadeOutParametersWithCullDistance(
		0.0, // cullDistance float64
	)
	// Output:
}
