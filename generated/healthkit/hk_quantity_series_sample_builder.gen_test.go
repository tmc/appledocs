// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit_test

import (
	"github.com/tmc/appledocs/generated/healthkit"
)

// Suppress unused import errors
var _ = healthkit.NewHKQuantitySeriesSampleBuilder

// ExampleHKQuantitySeriesSampleBuilder_Discard demonstrates using Discard on a HKQuantitySeriesSampleBuilder instance.
// Discards all previously collected data and invalidates the builder.
//
// Note: This example is not executed because Discard crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleHKQuantitySeriesSampleBuilder_Discard() {
	obj := healthkit.NewHKQuantitySeriesSampleBuilder()
	obj.Discard()
	}

