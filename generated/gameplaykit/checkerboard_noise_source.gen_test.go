// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewCheckerboardNoiseSource

// ExampleNewCheckerboardNoiseSourceWithSquareSize demonstrates how to create a CheckerboardNoiseSource instance using NewCheckerboardNoiseSourceWithSquareSize.
// Initializes a checkerboard noise source with the specified square size.
func ExampleNewCheckerboardNoiseSourceWithSquareSize() {
	_ = gameplaykit.NewCheckerboardNoiseSourceWithSquareSize(
		0.0, // squareSize float64
	)
	// Output:
}
