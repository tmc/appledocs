// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewSpheresNoiseSource

// ExampleNewSpheresNoiseSourceWithFrequency demonstrates how to create a SpheresNoiseSource instance using NewSpheresNoiseSourceWithFrequency.
// Initializes a sphere noise source with the specified frequency.
func ExampleNewSpheresNoiseSourceWithFrequency() {
	_ = gameplaykit.NewSpheresNoiseSourceWithFrequency(
		0.0, // frequency float64
	)
	// Output:
}
