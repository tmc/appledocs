// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewVoronoiNoiseSource

// ExampleNewVoronoiNoiseSourceWithFrequencyDisplacementDistanceEnabledSeed demonstrates how to create a VoronoiNoiseSource instance using NewVoronoiNoiseSourceWithFrequencyDisplacementDistanceEnabledSeed.
// Initializes a Voronoi noise source with the specified parameters.
func ExampleNewVoronoiNoiseSourceWithFrequencyDisplacementDistanceEnabledSeed() {
	_ = gameplaykit.NewVoronoiNoiseSourceWithFrequencyDisplacementDistanceEnabledSeed(
		0.0, // frequency float64
		0.0, // displacement float64
		false, // distanceEnabled bool
		gameplaykit.int32 /* not a class type */{}, // seed int32 /* not a class type */
	)
	// Output:
}

