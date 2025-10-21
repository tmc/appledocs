// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewLinearCongruentialRandomSource

// ExampleNewLinearCongruentialRandomSource demonstrates how to create a LinearCongruentialRandomSource instance.
// Initializes a random source from a nondeterministic seed.
func ExampleNewLinearCongruentialRandomSource() {
	_ = gameplaykit.NewLinearCongruentialRandomSource()
	// Output:
}
// ExampleNewLinearCongruentialRandomSourceWithSeed demonstrates how to create a LinearCongruentialRandomSource instance using NewLinearCongruentialRandomSourceWithSeed.
// Initializes a random source with the specified seed value.
func ExampleNewLinearCongruentialRandomSourceWithSeed() {
	_ = gameplaykit.NewLinearCongruentialRandomSourceWithSeed(
		0, // seed uint64
	)
	// Output:
}
