// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewNoise

// ExampleNewNoise demonstrates how to create a Noise instance.
func ExampleNewNoise() {
	_ = gameplaykit.NewNoise()
	// Output:
}

// ExampleNewNoiseWithComponentNoisesSelectionNoise demonstrates how to create a Noise instance using NewNoiseWithComponentNoisesSelectionNoise.
// Creates a noise object by combining the specified noise objects, using another noise object to select which regions of the output correspond to which input noise.
func ExampleNewNoiseWithComponentNoisesSelectionNoise() {
	_ = gameplaykit.NewNoiseWithComponentNoisesSelectionNoise(
		[]gameplaykit.Noise{}, // noises []Noise
		gameplaykit.GKNoise{}, // selectionNoise GKNoise
	)
	// Output:
}

// ExampleNewNoiseWithNoiseSource demonstrates how to create a Noise instance using NewNoiseWithNoiseSource.
// Initializes a noise object with the specified noise source.
func ExampleNewNoiseWithNoiseSource() {
	_ = gameplaykit.NewNoiseWithNoiseSource(
		gameplaykit.GKNoiseSource{}, // noiseSource GKNoiseSource
	)
	// Output:
}
