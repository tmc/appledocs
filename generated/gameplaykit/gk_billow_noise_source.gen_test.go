// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewBillowNoiseSource

// ExampleNewBillowNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed demonstrates how to create a BillowNoiseSource instance using NewBillowNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed.
// Creates a billow noise source with the specified parameters.
func ExampleNewBillowNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed() {
	_ = gameplaykit.NewBillowNoiseSourceWithFrequencyOctaveCountPersistenceLacunaritySeed(
		0.0, // frequency float64
		10, // octaveCount int
		0.0, // persistence float64
		0.0, // lacunarity float64
		gameplaykit.int32 /* not a class type */{}, // seed int32 /* not a class type */
	)
	// Output:
}
