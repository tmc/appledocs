// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewRandomSource

// ExampleNewRandomSource demonstrates how to create a RandomSource instance.
// Initializes a new random source object.
func ExampleNewRandomSource() {
	_ = gameplaykit.NewRandomSource()
	// Output:
}

// ExampleNewRandomSourceWithCoder demonstrates how to create a RandomSource instance using NewRandomSourceWithCoder.
func ExampleNewRandomSourceWithCoder() {
	_ = gameplaykit.NewRandomSourceWithCoder(
		gameplaykit.Coder{}, // aDecoder Coder
	)
	// Output:
}
