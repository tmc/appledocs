// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewCylindersNoiseSource

// ExampleNewCylindersNoiseSourceWithFrequency demonstrates how to create a CylindersNoiseSource instance using NewCylindersNoiseSourceWithFrequency.
// Initializes a cylinder noise source with the specified frequency.
func ExampleNewCylindersNoiseSourceWithFrequency() {
	_ = gameplaykit.NewCylindersNoiseSourceWithFrequency(
		0.0, // frequency float64
	)
	// Output:
}
