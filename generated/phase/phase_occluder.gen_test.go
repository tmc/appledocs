// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASEOccluder

// ExampleNewPHASEOccluderWithEngineShapes demonstrates how to create a PHASEOccluder instance using NewPHASEOccluderWithEngineShapes.
// Creates an occluder with the given engine and shapes.
func ExampleNewPHASEOccluderWithEngineShapes() {
	_ = phase.NewPHASEOccluderWithEngineShapes(
		phase.PHASEEngine{}, // engine PHASEEngine
		[]phase.PHASEShape{}, // shapes []PHASEShape
	)
	// Output:
}
