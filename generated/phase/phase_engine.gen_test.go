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
// ExamplePHASEEngine_Pause demonstrates using Pause on a PHASEEngine instance.
// Pauses all audio playback.
func ExamplePHASEEngine_Pause() {
	obj := phase.NewPHASEEngine()
	obj.Pause()
	// Output:
	}

// ExamplePHASEEngine_Stop demonstrates using Stop on a PHASEEngine instance.
// Stops all audio playback.
func ExamplePHASEEngine_Stop() {
	obj := phase.NewPHASEEngine()
	obj.Stop()
	// Output:
	}

// ExamplePHASEEngine_Update demonstrates using Update on a PHASEEngine instance.
// Processes app commands and increments framework processing.
func ExamplePHASEEngine_Update() {
	obj := phase.NewPHASEEngine()
	obj.Update()
	// Output:
	}

