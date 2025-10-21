// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASEMedium

// ExampleNewPHASEMediumWithEnginePreset demonstrates how to create a PHASEMedium instance using NewPHASEMediumWithEnginePreset.
// Creates a medium.
func ExampleNewPHASEMediumWithEnginePreset() {
	_ = phase.NewPHASEMediumWithEnginePreset(
		phase.PHASEEngine{}, // engine PHASEEngine
		phase.PHASEMediumPreset{}, // preset PHASEMediumPreset
	)
	// Output:
}
