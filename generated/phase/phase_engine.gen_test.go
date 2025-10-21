// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASEEngine

// ExampleNewPHASEEngineWithUpdateMode demonstrates how to create a PHASEEngine instance using NewPHASEEngineWithUpdateMode.
// Creates an engine updated by the app or framework.
func ExampleNewPHASEEngineWithUpdateMode() {
	_ = phase.NewPHASEEngineWithUpdateMode(
		phase.PHASEUpdateMode{}, // updateMode PHASEUpdateMode
	)
	// Output:
}
// ExampleNewPHASEEngineWithUpdateModeRenderingMode demonstrates how to create a PHASEEngine instance using NewPHASEEngineWithUpdateModeRenderingMode.
// Creates a new engine that has both update and rendering modes.
func ExampleNewPHASEEngineWithUpdateModeRenderingMode() {
	_ = phase.NewPHASEEngineWithUpdateModeRenderingMode(
		phase.PHASEUpdateMode{}, // updateMode PHASEUpdateMode
		phase.PHASERenderingMode{}, // renderingMode PHASERenderingMode
	)
	// Output:
}
