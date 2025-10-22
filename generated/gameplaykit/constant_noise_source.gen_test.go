// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewConstantNoiseSource

// ExampleNewConstantNoiseSourceWithValue demonstrates how to create a ConstantNoiseSource instance using NewConstantNoiseSourceWithValue.
// Initializes a noise source with the specified constant value.
func ExampleNewConstantNoiseSourceWithValue() {
	_ = gameplaykit.NewConstantNoiseSourceWithValue(
		0.0, // value float64
	)
	// Output:
}
