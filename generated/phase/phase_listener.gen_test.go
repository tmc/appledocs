// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASEListener

// ExampleNewPHASEListenerWithEngine demonstrates how to create a PHASEListener instance using NewPHASEListenerWithEngine.
// Creates a listener with the given engine.
func ExampleNewPHASEListenerWithEngine() {
	_ = phase.NewPHASEListenerWithEngine(
		phase.PHASEEngine{}, // engine PHASEEngine
	)
	// Output:
}
