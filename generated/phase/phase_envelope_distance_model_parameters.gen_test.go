// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASEEnvelopeDistanceModelParameters

// ExampleNewPHASEEnvelopeDistanceModelParametersWithEnvelope demonstrates how to create a PHASEEnvelopeDistanceModelParameters instance using NewPHASEEnvelopeDistanceModelParametersWithEnvelope.
// Creates the distance model parameters with an envelope.
func ExampleNewPHASEEnvelopeDistanceModelParametersWithEnvelope() {
	_ = phase.NewPHASEEnvelopeDistanceModelParametersWithEnvelope(
		phase.PHASEEnvelope{}, // envelope PHASEEnvelope
	)
	// Output:
}
