// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewRidgedNoiseSource

// ExampleNewRidgedNoiseSourceWithFrequencyOctaveCountLacunaritySeed demonstrates how to create a RidgedNoiseSource instance using NewRidgedNoiseSourceWithFrequencyOctaveCountLacunaritySeed.
// Initializes a ridged noise source with the specified parameters.
func ExampleNewRidgedNoiseSourceWithFrequencyOctaveCountLacunaritySeed() {
	_ = gameplaykit.NewRidgedNoiseSourceWithFrequencyOctaveCountLacunaritySeed(
		0.0, // frequency float64
		10, // octaveCount int
		0.0, // lacunarity float64
		gameplaykit.int32 /* not a class type */{}, // seed int32 /* not a class type */
	)
	// Output:
}
