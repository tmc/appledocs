// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewMersenneTwisterRandomSource

// ExampleNewMersenneTwisterRandomSource demonstrates how to create a MersenneTwisterRandomSource instance.
// Initializes a random source from a nondeterministic seed.
func ExampleNewMersenneTwisterRandomSource() {
	_ = gameplaykit.NewMersenneTwisterRandomSource()
	// Output:
}
// ExampleNewMersenneTwisterRandomSourceWithSeed demonstrates how to create a MersenneTwisterRandomSource instance using NewMersenneTwisterRandomSourceWithSeed.
// Initializes a random source with the specified seed value.
func ExampleNewMersenneTwisterRandomSourceWithSeed() {
	_ = gameplaykit.NewMersenneTwisterRandomSourceWithSeed(
		0, // seed uint64
	)
	// Output:
}
