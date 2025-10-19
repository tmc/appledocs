// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit_test

import (
	"github.com/tmc/appledocs/generated/spritekit"
)

// Suppress unused import errors
var _ = spritekit.NewSKRange


// ExampleNewRangeWithConstantValue demonstrates how to create a SKRange instance using NewRangeWithConstantValue.
// Creates and initializes a new range object that specifies a constant value.
func ExampleNewRangeWithConstantValue() {
	_ = spritekit.NewRangeWithConstantValue(
		0.0, // value float64
	)
	// Output:
}

// ExampleNewRangeWithLowerLimit demonstrates how to create a SKRange instance using NewRangeWithLowerLimit.
// Creates and initializes a new range object that specifies only a minimum value.
func ExampleNewRangeWithLowerLimit() {
	_ = spritekit.NewRangeWithLowerLimit(
		0.0, // lower float64
	)
	// Output:
}

// ExampleNewSKRangeWithLowerLimitUpperLimit demonstrates how to create a SKRange instance using NewSKRangeWithLowerLimitUpperLimit.
// Initializes a new range object.
func ExampleNewSKRangeWithLowerLimitUpperLimit() {
	_ = spritekit.NewSKRangeWithLowerLimitUpperLimit(
		0.0, // lower float64
		0.0, // upper float64
	)
	// Output:
}

// ExampleNewRangeWithUpperLimit demonstrates how to create a SKRange instance using NewRangeWithUpperLimit.
// Creates and initializes a new range object that specifies only a maximum value.
func ExampleNewRangeWithUpperLimit() {
	_ = spritekit.NewRangeWithUpperLimit(
		0.0, // upper float64
	)
	// Output:
}

// ExampleNewRangeWithValueVariance demonstrates how to create a SKRange instance using NewRangeWithValueVariance.
// Creates and initializes a new range object using a value and a maximum distance from that value.
func ExampleNewRangeWithValueVariance() {
	_ = spritekit.NewRangeWithValueVariance(
		0.0, // value float64
		0.0, // variance float64
	)
	// Output:
}


