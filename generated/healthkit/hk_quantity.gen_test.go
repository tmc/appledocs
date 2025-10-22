// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit_test

import (
	"github.com/tmc/appledocs/generated/healthkit"
)

// Suppress unused import errors
var _ = healthkit.NewHKQuantity

// ExampleNewHKQuantityWithUnitDoubleValue demonstrates how to create a HKQuantity instance using NewHKQuantityWithUnitDoubleValue.
// Instantiates and returns a new quantity object.
func ExampleNewHKQuantityWithUnitDoubleValue() {
	_ = healthkit.NewHKQuantityWithUnitDoubleValue(
		healthkit.HKUnit{}, // unit HKUnit
		0.0, // value float64
	)
	// Output:
}
