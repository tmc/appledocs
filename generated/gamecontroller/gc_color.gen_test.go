// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller_test

import (
	"github.com/tmc/appledocs/generated/gamecontroller"
)

// Suppress unused import errors
var _ = gamecontroller.NewGCColor

// ExampleNewGCColorWithRedGreenBlue demonstrates how to create a GCColor instance using NewGCColorWithRedGreenBlue.
// Creates a color with the specified red, green, and blue values.
func ExampleNewGCColorWithRedGreenBlue() {
	_ = gamecontroller.NewGCColorWithRedGreenBlue(
		0.0, // red float32
		0.0, // green float32
		0.0, // blue float32
	)
	// Output:
}
