// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASEMaterial

// ExampleNewPHASEMaterialWithEnginePreset demonstrates how to create a PHASEMaterial instance using NewPHASEMaterialWithEnginePreset.
// Creates a material with the given preset.
func ExampleNewPHASEMaterialWithEnginePreset() {
	_ = phase.NewPHASEMaterialWithEnginePreset(
		phase.PHASEEngine{}, // engine PHASEEngine
		phase.PHASEMaterialPreset{}, // preset PHASEMaterialPreset
	)
	// Output:
}
