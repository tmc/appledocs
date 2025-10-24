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
// ExampleNoise_ApplyAbsoluteValue demonstrates using ApplyAbsoluteValue on a Noise instance.
// Replaces all negative values in the noise field with their positive absolute values.
func ExampleNoise_ApplyAbsoluteValue() {
	obj := gameplaykit.NewNoise()
	obj.ApplyAbsoluteValue()
	// Output:
	}

// ExampleNoise_Invert demonstrates using Invert on a Noise instance.
// Replaces all values in the noise field with their opposite, reversing the range of noise values.
func ExampleNoise_Invert() {
	obj := gameplaykit.NewNoise()
	obj.Invert()
	// Output:
	}

