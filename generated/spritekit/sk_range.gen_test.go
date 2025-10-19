// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit_test

import (
	"github.com/tmc/appledocs/generated/spritekit"
)

// Suppress unused import errors
var _ = spritekit.NewSKRange


// ExampleNewSKRangeWithLowerLimit demonstrates how to create a SKRange instance using NewSKRangeWithLowerLimit.
// Creates and initializes a new range object that specifies only a minimum value.
func ExampleNewSKRangeWithLowerLimit() {
	_ = spritekit.NewSKRangeWithLowerLimit(
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

// ExampleNewSKRangeWithUpperLimit demonstrates how to create a SKRange instance using NewSKRangeWithUpperLimit.
// Creates and initializes a new range object that specifies only a maximum value.
func ExampleNewSKRangeWithUpperLimit() {
	_ = spritekit.NewSKRangeWithUpperLimit(
		0.0, // upper float64
	)
	// Output:
}

// ExampleNewSKRangeWithValueVariance demonstrates how to create a SKRange instance using NewSKRangeWithValueVariance.
// Creates and initializes a new range object using a value and a maximum distance from that value.
func ExampleNewSKRangeWithValueVariance() {
	_ = spritekit.NewSKRangeWithValueVariance(
		0.0, // value float64
		0.0, // variance float64
	)
	// Output:
}

// ExampleNewSKRangeWithConstantValue demonstrates how to create a SKRange instance using NewSKRangeWithConstantValue.
// Creates and initializes a new range object that specifies a constant value.
func ExampleNewSKRangeWithConstantValue() {
	_ = spritekit.NewSKRangeWithConstantValue(
		0.0, // value float64
	)
	// Output:
}


