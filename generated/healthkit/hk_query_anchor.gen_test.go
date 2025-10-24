// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit_test

import (
	"github.com/tmc/appledocs/generated/healthkit"
)

// Suppress unused import errors
var _ = healthkit.NewHKQueryAnchor

// ExampleNewHKQueryAnchorFromValue demonstrates how to create a HKQueryAnchor instance using NewHKQueryAnchorFromValue.
// Returns an anchor object from the provided anchor value.
func ExampleNewHKQueryAnchorFromValue() {
	_ = healthkit.NewHKQueryAnchorFromValue(
		0, // value uint
	)
	// Output:
}
