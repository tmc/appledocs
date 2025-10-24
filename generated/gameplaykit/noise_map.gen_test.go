// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewNoiseMap

// ExampleNewNoiseMap demonstrates how to create a NoiseMap instance.
// Initializes a noise map with a constant noise value of zero throughout.
func ExampleNewNoiseMap() {
	_ = gameplaykit.NewNoiseMap()
	// Output:
}

// ExampleNewNoiseMapWithNoise demonstrates how to create a NoiseMap instance using NewNoiseMapWithNoise.
// Initializes a noise map by sampling from the specified noise object.
func ExampleNewNoiseMapWithNoise() {
	_ = gameplaykit.NewNoiseMapWithNoise(
		gameplaykit.GKNoise{}, // noise GKNoise
	)
	// Output:
}
