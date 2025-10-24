// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewRandomDistribution

// ExampleNewRandomDistributionForDieWithSideCount demonstrates how to create a RandomDistribution instance using NewRandomDistributionForDieWithSideCount.
// Creates a random distribution equivalent to a die with the specified number of sides.
func ExampleNewRandomDistributionForDieWithSideCount() {
	_ = gameplaykit.NewRandomDistributionForDieWithSideCount(
		10, // sideCount int
	)
	// Output:
}
// ExampleNewRandomDistributionWithLowestValueHighestValue demonstrates how to create a RandomDistribution instance using NewRandomDistributionWithLowestValueHighestValue.
// Creates a random distribution with the specified lower and upper bounds, using the Arc4 randomizer.
func ExampleNewRandomDistributionWithLowestValueHighestValue() {
	_ = gameplaykit.NewRandomDistributionWithLowestValueHighestValue(
		0, // lowestInclusive int
		0, // highestInclusive int
	)
	// Output:
}
// ExampleRandomDistribution_NextBool demonstrates using NextBool on a RandomDistribution instance.
// Generates and returns a new random Boolean value within the characteristics of the distribution.
func ExampleRandomDistribution_NextBool() {
	obj := gameplaykit.NewRandomDistribution()
	_ = obj.NextBool()
	// Output:
	}

// ExampleRandomDistribution_NextInt demonstrates using NextInt on a RandomDistribution instance.
// Generates and returns a new random integer within the bounds of the distribution.
func ExampleRandomDistribution_NextInt() {
	obj := gameplaykit.NewRandomDistribution()
	_ = obj.NextInt()
	// Output:
	}

// ExampleRandomDistribution_NextUniform demonstrates using NextUniform on a RandomDistribution instance.
// Generates and returns a new random floating-point value within the characteristics of the distribution.
func ExampleRandomDistribution_NextUniform() {
	obj := gameplaykit.NewRandomDistribution()
	_ = obj.NextUniform()
	// Output:
	}

