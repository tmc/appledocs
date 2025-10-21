// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASESource

// ExampleNewPHASESourceWithEngine demonstrates how to create a PHASESource instance using NewPHASESourceWithEngine.
// Creates a single point in the environment from which sound emanates.
func ExampleNewPHASESourceWithEngine() {
	_ = phase.NewPHASESourceWithEngine(
		phase.PHASEEngine{}, // engine PHASEEngine
	)
	// Output:
}
// ExampleNewPHASESourceWithEngineShapes demonstrates how to create a PHASESource instance using NewPHASESourceWithEngineShapes.
// Creates a voluminous area in the environment from which sound emanates.
func ExampleNewPHASESourceWithEngineShapes() {
	_ = phase.NewPHASESourceWithEngineShapes(
		phase.PHASEEngine{}, // engine PHASEEngine
		[]phase.PHASEShape{}, // shapes []PHASEShape
	)
	// Output:
}
