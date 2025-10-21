// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit_test

import (
	"github.com/tmc/appledocs/generated/healthkit"
)

// Suppress unused import errors
var _ = healthkit.NewHKLiveWorkoutDataSource

// ExampleNewHKLiveWorkoutDataSourceWithHealthStoreWorkoutConfiguration demonstrates how to create a HKLiveWorkoutDataSource instance using NewHKLiveWorkoutDataSourceWithHealthStoreWorkoutConfiguration.
// Creates a new data source based on the provided workout configuration.
func ExampleNewHKLiveWorkoutDataSourceWithHealthStoreWorkoutConfiguration() {
	_ = healthkit.NewHKLiveWorkoutDataSourceWithHealthStoreWorkoutConfiguration(
		healthkit.HKHealthStore{}, // healthStore HKHealthStore
		healthkit.HKWorkoutConfiguration{}, // configuration HKWorkoutConfiguration
	)
	// Output:
}
