// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewRandomDistribution


// ExampleNewRandomDistributionWithLowestValueHighestValue demonstrates how to create a RandomDistribution instance using NewRandomDistributionWithLowestValueHighestValue.
// Creates a random distribution with the specified lower and upper bounds, using the Arc4 randomizer.
func ExampleNewRandomDistributionWithLowestValueHighestValue() {
	_ = gameplaykit.NewRandomDistributionWithLowestValueHighestValue(
		0, // lowestInclusive int
		0, // highestInclusive int
	)
	// Output:
}


// ExampleNewRandomDistributionForDieWithSideCount demonstrates how to create a RandomDistribution instance using NewRandomDistributionForDieWithSideCount.
// Creates a random distribution equivalent to a die with the specified number of sides.
func ExampleNewRandomDistributionForDieWithSideCount() {
	_ = gameplaykit.NewRandomDistributionForDieWithSideCount(
		0, // sideCount int
	)
	// Output:
}


